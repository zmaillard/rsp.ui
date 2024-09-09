# roadsign.pictures

A gallery of road signs from North America.  Hosted as a [Hugo](https://gohugo.io) static website, with content generated at build time from a PostgreSQL database using Go.

To run locally, you will need to have the following tools installed:
- [Hugo](https://gohugo.io/installation/).  
- NodeJS v.20.11.0 or later to manage Javascript and CSS assets.
- [Go](https://go.dev/doc/install) v1.20 or above to run the Go program to generate the content.

The other sub-projects have their own dependencies, and are listed below in the [Sub Projects](#sub-projects) section.

## Environment Variables

### Content Exporter
| Variable Name | Description |
----------------|------------------
| DB_USER | User with permissions to read from the database |
| DB_HOST | Host name of PostgreSQL server |
| DB_PASSWORD | Password for the user |
| DB_NAME | Database name that hosts sign database |
| DB_PORT | Port to connect to the PostgreSQL server |
| HUGO_PATH | Path to output content to.  Should correspond with the directory that Hugo reads content from. |

### Static Website
| Variable Name                   | Description                                         |
---------------------------------|-----------------------------------------------------
| HUGO_PARAMS_W3WAPIKEY           | API Token for [What3Words](https://what3words.com/) |
| HUGO_PARAMS_SEARCHURL           | Base URL for the search API                         |
| HUGO_PARAMS_RANDOMURL           | Base URL for [Random Sign API](random/README.md)    |
| HUGO_PARAMS_SEARCHINDEX         | Search index name for signs                         |
| HUGO_PARAMS_SEARCHINDEXHIGHWAY  | Search index name for highways                     |
| HUGO_PARAMS_SEARCHKEY           | Token with read access to the search index          |
| HUGO_PARAMS_SIGNBASEURL         | Base URL for the Roadsign Picture hosting           |
| HUGO_PARAMS_SHIELDBASEURL       | Base URL for the Highway Shields hosting            |
| HUGO_PARAMS_MAPBOXTOKEN         | API Token for [Mapbox](https://www.mapbox.com)      |
| HUGO_PARAMS_MAPTILE             | Url for [Sign Vector Tiles](tiles/README.md)        |

## Running Locally

The all the necessary CSS and JavaScript build tools will need to be installed the first name:
```bash
npm install
```

To export content, set the correct environment variables and run the following command:
```bash
go run main.go
```

To run the website locally, set the correct environment variables and run the following command:
```bash
hugo serve
```

If making changes to the CSS you will want to run the Tailwind service in a different terminal window:
```bash
npm run tailwind
```

## Sub Projects
- [Indexer](index/README.md)
- [Random Signs](random/README.md)
- [Map Tiles](tiles/README.md)

