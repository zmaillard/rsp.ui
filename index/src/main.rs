use tokio_postgres::types::Json;
use anyhow::Result;
use config::Config;
use serde::{Deserialize, Serialize};
use geo_types::{{Point as GeoPoint}};
use meilisearch_sdk::Client;
use time::PrimitiveDateTime;
use tokio;

#[derive(Serialize, Debug)]
pub struct HighwayIndex {
    id: String,
    name: String,
    slug: String,
    image_name: String,
    url: String,
}

#[derive(Serialize, Debug)]
pub struct SignIndex {
    id: String,
    title: String,
    description: String,
    highways: Vec<Highway>,
    #[serde(rename="_geo")]
    point: Point,
    #[serde(rename="dateTaken")]
    date_taken: PrimitiveDateTime,
    country: Locality,
    county: Option<Locality>,
    state: Locality,
    place: Option<Locality>,
    url: String,
    quality: i32,
}

#[derive(Serialize, Debug)]
pub struct Locality {
    name: String,
    slug: String,
}

#[derive(Serialize, Deserialize, Debug)]
pub struct Highway {
    name: String,
    slug: String,
}

#[derive(Serialize, Deserialize, Debug)]
pub struct Point {
    lng: f64,
    lat: f64,
}

#[derive(Debug, Deserialize)]
pub struct AppSettings {
    postgres_host: String,
    postgres_user: String,
    postgres_password: String,
    postgres_port: String,
    postgres_db: String,
    meilisearch_host: String,
    meilisearch_key: String,
    image_hosting_url: String,
}

impl AppSettings {
    pub fn new() -> Self {
        let settings = Config::builder()
            .add_source(config::Environment::with_prefix("RSP"))
            .build()
            .unwrap();

        settings.try_deserialize::<Self>().unwrap()
    }
    pub fn database_connection_string(&self) -> String {
        format!("user={} password={} host={} port={} dbname={} sslmode=require",
                self.postgres_user,
                self.postgres_password,
                self.postgres_host,
                self.postgres_port,
                self.postgres_db,
        )
    }

    pub fn build_image_url(&self, image_id: &i64) -> String {
        format!("{}{}/{}_l.jpg", self.image_hosting_url, image_id ,image_id)
    }

    pub fn meilisearch_client(&self) -> meilisearch_sdk::client::Client {
        meilisearch_sdk::client::Client::new(
            &self.meilisearch_host,
            Some(&self.meilisearch_key),
        )
    }
}
const FLUSH:usize = 25;

#[tokio::main(flavor= "current_thread")]
async fn main() -> Result<()>{
    let app_settings = AppSettings::new();

    let mut roots = rustls::RootCertStore::empty();

    for cert in rustls_native_certs::load_native_certs().expect("could not load platform certs") {
        roots.add(cert).unwrap();
    }

    // Setup TLS
    let config = rustls::ClientConfig::builder()
        .with_root_certificates(roots)
        .with_no_client_auth();

    let tls = tokio_postgres_rustls::MakeRustlsConnect::new(config);

    let (client, connection) = tokio_postgres::connect(app_settings.database_connection_string().as_str(), tls).await?;

    tokio::spawn(async move {
        if let Err(e) = connection.await {
            eprintln!("connection error: {}", e);
        }
    });

    let meilisearch = app_settings.meilisearch_client();
    highway_index(&meilisearch, &client).await?;
    sign_index(&meilisearch, &client, &app_settings).await?;

    Ok(())
}

fn build_locality(name: Option<&str>, slug: Option<&str>) -> Option<Locality> {
    match (name, slug) {
        (Some(name), Some(slug)) =>
            Some(Locality {
                name: name.to_string(),
                slug: slug.to_string(),
            }),
        _ =>  None,
    }
}

async fn highway_index(meilisearch_client: &Client, psql_client: &tokio_postgres::Client) -> Result<()> {
    let highway_index = meilisearch_client.index("highway");
    let mut row_vect:Vec<HighwayIndex> = Vec::new();
    for row in psql_client.query("select id, highway_name, image_name, slug, case when image_name = '' then '' when image_name is null then '' else 'https://shield.roadsign.pictures/Shields/' || highway.image_name end from sign.highway where id in (select distinct highway_id from sign.highwaysign_highway)", &[]).await? {
        let id: i32 = row.get(0);
        let name: &str = row.get(1);
        let image_name: &str = row.get(2);
        let slug: &str = row.get(3);
        let url: &str = row.get(4);
        let highway = HighwayIndex {
            id: id.to_string(),
            name: name.to_string(),
            url: url.to_string(),
            slug: slug.to_string(),
            image_name: image_name.to_string(),
        };
        row_vect.push(highway);

        if row_vect.len() % FLUSH == 0 {
            highway_index.add_or_update(&row_vect,Some("id")).await?;
            row_vect.clear();
        }
    }

    if row_vect.len() > 0 {
        highway_index.add_or_update(&row_vect,Some("id")).await?;
    }

    Ok(())
}

async fn sign_index(meilisearch_client: &Client, psql_client: &tokio_postgres::Client, app_settings: &AppSettings) -> Result<()> {
    let sign_index = meilisearch_client.index("signs");

    let mut row_vect:Vec<SignIndex> = Vec::new();
    for row in psql_client.query("SELECT imageid, title, sign_description, date_taken, country_slug, country_name, state_slug, state_name, county_name, county_slug, place_name, place_slug, hwys, point::geometry::point, quality FROM rsp.sign.vwindexsign where last_indexed is null or last_indexed < last_update", &[]).await? {
        let imageid: i64 = row.get(0);

        let title: &str = row.get(1);
        let sign_description: &str = row.get(2);
        let date_taken: PrimitiveDateTime = row.get(3);
        let country_slug: &str = row.get(4);
        let country_name: &str = row.get(5);
        let state_slug: &str = row.get(6);
        let state_name: &str = row.get(7);
        let county_name: Option<&str> = row.get(8);
        let county_slug: Option<&str> = row.get(9);
        let place_name: Option<&str> = row.get(10);
        let place_slug: Option<&str> = row.get(11);
        let hwys: Option<Json<Vec<Highway>>> = row.get(12);
        let point: GeoPoint = row.get(13);
        let point = Point {
            lng: point.x(),
            lat: point.y(),
        };

        let hwy_res: Vec<Highway> = match hwys {
            Some(hwy) => hwy.0,
            None => Vec::new(),
        };

        let quality: i32 = row.get(14);

        let place: Option<Locality> = build_locality(place_name, place_slug);

        let county: Option<Locality> = build_locality(county_name, county_slug);

        let url = app_settings.build_image_url(&imageid);
        let sign = SignIndex {
            id: imageid.to_string(),
            title: title.to_string(),
            description: sign_description.to_string(),
            highways: hwy_res,
            point,
            date_taken,
            country: Locality {
                name: country_name.to_string(),
                slug: country_slug.to_string(),
            },
            county,
            state: Locality {
                name: state_name.to_string(),
                slug: state_slug.to_string(),
            },
            place,
            url,
            quality,
        };

        row_vect.push(sign);

        if row_vect.len() % FLUSH == 0 {
            sign_index.add_or_update(&row_vect,Some("id")).await?;
            row_vect.clear();
        }
    }

    if row_vect.len() > 0 {
        sign_index.add_or_update(&row_vect,Some("id")).await?;
    }

    // Update the last_indexed column
    psql_client.execute("UPDATE sign.highwaysign SET last_indexed = now() WHERE last_indexed is null or last_indexed < last_update", &[]).await?;

    Ok(())
}
