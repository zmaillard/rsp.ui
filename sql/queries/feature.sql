-- name: GetFeaturesWithinBuffer :many
SELECT id, ST_X(point::geometry) as longitude, ST_Y(point::geometry) as latitude, name, admin_area_country_id, admin_area_state_id, featured from sign.feature WHERE st_intersects(point, ST_Buffer(ST_GeogFromText($1)), $2 );

-- name: DeleteFeature :exec
DELETE FROM sign.feature WHERE id = $1;

-- name: ReverseFeatureLink :exec
UPDATE sign.feature_link SET to_feature = $1, from_feature = $2 WHERE id = $3;

-- name: UpdateFeatureLinkName :exec
UPDATE sign.feature_link SET road_name = $1 WHERE id = $2;

-- name: UpdateFeatured :exec
UPDATE sign.feature SET featured = $1 WHERE id = $2;


-- name: UpdateFeatureName :exec
UPDATE sign.feature SET name = $1 WHERE id = $2;

-- name: GetFeatureLinkHighwayDirection :one
SELECT is_descending from sign.feature_link_highway WHERE highway_id = $1 AND feature_link_id = $2;

-- name: ReverseFeatureLinkHighway :exec
UPDATE sign.feature_link_highway SET is_descending = NOT is_descending WHERE id = $1;

-- name: GetFeatureLinkHighway :one
SELECT * from sign.feature_link_highway WHERE highway_id = $1 AND feature_link_id = $2;

-- name: UpdateFeatureAdminArea :exec
UPDATE sign.feature SET admin_area_country_id = $1, admin_area_state_id = $2 WHERE id = $3;

-- name: GetFeatureLink :one
SELECT id, from_feature, to_feature, road_name from sign.feature_link WHERE id = $1;

-- name: GetFeatureLinksByIds :many
SELECT id, from_feature, to_feature, road_name from sign.feature_link WHERE id in ($1::int[]);

-- name: DeleteFeatureLink :exec
DELETE FROM sign.feature_link WHERE id = $1;

-- name: DeleteFeatureLinkHighway :exec
DELETE FROM sign.feature_link_highway WHERE feature_link_id = $1;

-- name: GetFeature :one
SELECT id, ST_X(point::geometry) as longitude, ST_Y(point::geometry) as latitude, name, admin_area_country_id, admin_area_state_id, featured  from sign.feature WHERE id = $1;

-- name: GetAllFeatures :many
SELECT id, ST_X(point::geometry) as longitude, ST_Y(point::geometry) as latitude, name, admin_area_country_id, admin_area_state_id, featured  from sign.feature;

-- name: CreateFeature :one
INSERT INTO sign.feature (point, name, admin_area_country_id, admin_area_state_id) VALUES (ST_PointFromText($1), $2, $3, $4) RETURNING id;

-- name: CreateFeatureLink :one
INSERT INTO sign.feature_link (from_feature, to_feature) VALUES ($1, $2) RETURNING id;

-- name: UpdateBeginAndEnd :exec
UPDATE sign.feature_link SET from_feature = $1, to_feature = $2 WHERE id = $3;

-- name: GetFeatureLinkHighways :many
SELECT flh.*, h.highway_name from sign.feature_link_highway flh INNER JOIN sign.highway h ON h.id = flh.highway_id WHERE feature_link_id = $1;

-- name: AddHighwayToFeatureLink :one
INSERT INTO sign.feature_link_highway (feature_link_id, highway_id, is_descending) VALUES ($1, $2, $3) RETURNING id;

-- name: RemoveHighwayFromFeatureLink :exec
DELETE FROM sign.feature_link_highway WHERE highway_id = $1 AND feature_link_id = $2;

-- name: GetAllFeaturedFeatures :many
SELECT id, ST_X(point::geometry) as longitude, ST_Y(point::geometry) as latitude, name, admin_area_country_id, admin_area_state_id, featured  from sign.feature WHERE featured = $1;

-- name: GetFeatureConnectedCount :one
SELECT COUNT(*) from sign.feature_link WHERE from_feature = $1 OR to_feature = $1;

