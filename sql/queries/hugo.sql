-- name: GetHugoCountries :many
SELECT id, country_name, country_slug, subdivision_name, image_count, states, featured, highway_types FROM sign.vwhugocountry;

-- name: GetHugoCounties :many
SELECT id, county_name, county_slug, image_count, state_name, state_slug FROM sign.vwhugocounty WHERE image_count is not null;

-- name: GetHugoHighways :many
SELECT id, highway_name, slug, sort_number, image_name, highway_type_slug, highway_type_name, cast (states as text[]), cast (counties as text[]), cast (places as text[]), cast (previous_features as int[]), cast (next_features as int[]), display_name FROM sign.vwhugohighway;


-- name: GetHugoHighwaySigns :many
SELECT id, title, sign_description, feature_id, date_taken, imageid, flickrid, point, country_slug, state_slug, place_slug, county_slug, tags, categories, highways, is_to, image_height, image_width, quality FROM sign.vwhugohighwaysign;

-- name: GetHugoHighwayTypes :many
SELECT id, highway_type_name, highway_type_slug, sort, coalesce(imagecount,0), imageid, cast(highways as text[]), country FROM sign.vwhugohighwaytype;

-- name: GetHugoPlaces :many
SELECT id, place_name, place_slug, image_count, state_name, state_slug FROM sign.vwhugoplace where image_count is not null;

-- name: GetHugoStates :many
SELECT id, state_name, state_slug, subdivision_name, image_count, highways, featured, country_slug, counties, places, categories, highway_names FROM sign.vwhugostate;

-- name: GetHugoFeatures :many
SELECT id, cast(point as geometry), name, cast(signs as text[]), state_name, state_slug, country_name, country_slug, highway_names FROM sign.vwhugofeature;

-- name: GetHugoFeatureLinks :many
SELECT id, from_feature, to_feature, road_name, highways, to_point, from_point, highway_name FROM sign.vwhugofeaturelink;

-- name: GetHugoTags :many
select id, name, slug, is_category, category_details from sign.tag;

-- name: GetHugoHighwayNames :many
select hn.id, sign.slugify(hn.name) as slug, hn.name, aas.name as state_name, aas.slug as state_slug from sign.highway_name hn inner join sign.admin_area_state aas on hn.state_id = aas.id;