-- name: GetSignsByFeatureId :many
SELECT  id, flickrid, date_taken, date_added, title, sign_description, image_width, image_height, ST_X(point::geometry)::float as longitude, ST_Y(point::geometry)::float as latitude, imageid, lastsyncwithflickr, last_update, cropped_image_id, last_indexed, archived, feature_id, admin_area_country_id, admin_area_state_id, admin_area_county_id, admin_area_place_id
FROM sign.highwaysign WHERE feature_id = $1;

-- name: DeleteTagsOnSign :exec
DELETE FROM sign.tag_highwaysign WHERE highwaysign_id = $1;

-- name: DeleteHighwaysOnSign :exec
DELETE FROM sign.highwaysign_highway WHERE highwaysign_id = $1;

-- name: GetPendingChanges :many
SELECT id, highwaysign_id, changed_on FROM sign.highwaysign_pending_changes;

-- name: DeletePendingChange :exec
DELETE FROM sign.highwaysign_pending_changes WHERE id = $1;

-- name: GetSign :one
SELECT  id, flickrid, date_taken, date_added, title, sign_description, image_width, image_height, ST_X(point::geometry)::float as longitude, ST_Y(point::geometry)::float as latitude, imageid, lastsyncwithflickr, last_update, cropped_image_id, last_indexed, archived, feature_id, admin_area_country_id, admin_area_state_id, admin_area_county_id, admin_area_place_id
FROM sign.highwaysign WHERE id = $1;

-- name: GetSignByImageId :one
SELECT  id, flickrid, date_taken, date_added, title, sign_description, image_width, image_height, ST_X(point::geometry)::float as longitude, ST_Y(point::geometry)::float as latitude, imageid, lastsyncwithflickr, last_update, cropped_image_id, last_indexed, archived, feature_id, admin_area_country_id, admin_area_state_id, admin_area_county_id, admin_area_place_id
FROM sign.highwaysign WHERE imageid = $1;

-- name: GetAllSigns :many
SELECT  id, flickrid, date_taken, date_added, title, sign_description, image_width, image_height, ST_X(point::geometry)::float as longitude, ST_Y(point::geometry)::float as latitude, imageid, lastsyncwithflickr, last_update, cropped_image_id, last_indexed, archived, feature_id, admin_area_country_id, admin_area_state_id, admin_area_county_id, admin_area_place_id
FROM sign.highwaysign;

-- name: GetSignsByIds :many
SELECT id, flickrid, date_taken, date_added, title, sign_description, image_width, image_height, ST_X(point::geometry)::float as longitude, ST_Y(point::geometry)::float as latitude, imageid, lastsyncwithflickr, last_update, cropped_image_id, last_indexed, archived, feature_id, admin_area_country_id, admin_area_state_id, admin_area_county_id, admin_area_place_id
FROM sign.highwaysign WHERE id = ANY($1::int[]);

-- name: UpdateSignAdminAreas :exec
UPDATE sign.highwaysign SET admin_area_country_id = sqlc.arg(admin_area_country_id) , admin_area_state_id = sqlc.arg(admin_area_state_id), admin_area_county_id = sqlc.narg(admin_area_county_id), admin_area_place_id = sqlc.narg(admin_area_place_id), last_update = sqlc.arg(last_update) WHERE id = sqlc.arg(highway_sign_id);

-- name: CreateSign :one
INSERT INTO sign.highwaysign (date_taken, date_added, title, sign_description, image_width, image_height, point, imageid, feature_id, admin_area_country_id, admin_area_state_id, admin_area_county_id, admin_area_place_id) VALUES ($1, $2, $3, $4, $5, $6, ST_PointFromText($7), $8, $9, $10, $11, $12, $13) RETURNING id;

-- name: UpdateSignLocation :exec
UPDATE sign.highwaysign SET point = ST_PointFromText($1) WHERE id = $2;

-- name: CreateTag :one
INSERT INTO sign.tag (name, slug, flickr_only) VALUES ($1, $2, $3) RETURNING *;

-- name: GetTagByName :one
SELECT * FROM sign.tag WHERE name = $1;

-- name: AddTagToSign :one
INSERT INTO sign.tag_highwaysign (tag_id, highwaysign_id) VALUES ($1, $2) RETURNING *;

-- name: GetTagsStartWith :many
SELECT * FROM sign.tag WHERE name ilike $1;

-- name: GetTagById :one
SELECT * FROM sign.tag WHERE id = $1;

-- name: GetCategories :many
SELECT * FROM sign.tag WHERE is_category = true;

-- name: UpdateCategoryDetails :exec
UPDATE sign.tag SET category_details = sqlc.arg(category_details)  WHERE id = sqlc.arg(tag_id);

-- name: UpdateLastUpdated :exec
UPDATE sign.highwaysign SET last_update = $2 WHERE id = $1;


-- name: UpdateSignTitle :exec
UPDATE sign.highwaysign SET title = $2, last_update = $3 WHERE id = $1;

-- name: UpdateSignDescription :exec
UPDATE sign.highwaysign SET sign_description = $2, last_update = $3 WHERE id = $1;

-- name: UpdateSignFeature :exec
UPDATE sign.highwaysign SET feature_id = $2, last_update = $3 WHERE id = $1;

-- name: UpdateFeaturedSignForState :exec
UPDATE sign.admin_area_state SET featured_sign_id = $2 WHERE id = $1;

-- name: UpdateFeaturedSignForHighwayType :exec
UPDATE sign.highway_type SET display_image_id = $2 WHERE id = $1;



-- name: GetSignDetails :one
SELECT  highwaysign.id, flickrid, date_taken, date_added, title, sign_description, image_width, image_height, ST_X(point::geometry)::float as longitude, ST_Y(point::geometry)::float as latitude, imageid, lastsyncwithflickr, last_update, cropped_image_id, last_indexed, archived, feature_id, s.name as state, c.name as country, co.name as county, p.name as place
FROM sign.highwaysign INNER JOIN sign.admin_area_state s on highwaysign.admin_area_state_id = s.id
                      INNER JOIN sign.admin_area_country c on highwaysign.admin_area_country_id = c.id
                      LEFT JOIN sign.admin_area_county co on highwaysign.admin_area_county_id = co.id
                      LEFT JOIN sign.admin_area_place p on highwaysign.admin_area_place_id = p.id
WHERE highwaysign.id = $1;

-- name: GetSignTags :many
SELECT t.name FROM sign.tag_highwaysign ths
    INNER JOIN sign.tag t on ths.tag_id = t.id
         WHERE ths.highwaysign_id = $1;


-- name: GetSignsOnHighway :many
SELECT  h.id, flickrid, date_taken, date_added, title, sign_description, image_width, image_height, ST_X(point::geometry)::float as longitude, ST_Y(point::geometry)::float as latitude, imageid, lastsyncwithflickr, last_update, cropped_image_id, last_indexed, archived, feature_id, admin_area_country_id, admin_area_state_id, admin_area_county_id, admin_area_place_id
FROM sign.highwaysign_highway hsh
    INNER JOIN sign.highwaysign h on h.id = hsh.highwaysign_id
WHERE hsh.highway_id = $1;

-- name: GetAllSignsByStateSearch :many
SELECT  id, flickrid, date_taken, date_added, title, sign_description, image_width, image_height, ST_X(point::geometry)::float as longitude, ST_Y(point::geometry)::float as latitude, imageid, lastsyncwithflickr, last_update, cropped_image_id, last_indexed, archived, feature_id, admin_area_country_id, admin_area_state_id, admin_area_county_id, admin_area_place_id, array_to_json(highway_tags.tags) as tags
FROM sign.highwaysign
         LEFT OUTER JOIN (
    SELECT ths.highwaysign_id, ARRAY_AGG(t.name) as tags FROM sign.tag_highwaysign ths
                                                                  INNER JOIN sign.tag t on t.id = ths.tag_id
    GROUP BY ths.highwaysign_id
) highway_tags ON highwaysign.id = highway_tags.highwaysign_id
WHERE admin_area_state_id = $1
ORDER BY date_taken DESC, id
LIMIT $2 OFFSET $3;

-- name: GetAllSignsByCountySearch :many
SELECT  id, flickrid, date_taken, date_added, title, sign_description, image_width, image_height, ST_X(point::geometry)::float as longitude, ST_Y(point::geometry)::float as latitude, imageid, lastsyncwithflickr, last_update, cropped_image_id, last_indexed, archived, feature_id, admin_area_country_id, admin_area_state_id, admin_area_county_id, admin_area_place_id, array_to_json(highway_tags.tags) as tags
FROM sign.highwaysign
         LEFT OUTER JOIN (
    SELECT ths.highwaysign_id, ARRAY_AGG(t.name) as tags FROM sign.tag_highwaysign ths
                                                                  INNER JOIN sign.tag t on t.id = ths.tag_id
    GROUP BY ths.highwaysign_id
) highway_tags ON highwaysign.id = highway_tags.highwaysign_id
WHERE admin_area_county_id = $1
ORDER BY date_taken DESC, id
LIMIT $2 OFFSET $3;

-- name: GetAllSignsByCountrySearch :many
SELECT  id, flickrid, date_taken, date_added, title, sign_description, image_width, image_height, ST_X(point::geometry)::float as longitude, ST_Y(point::geometry)::float as latitude, imageid, lastsyncwithflickr, last_update, cropped_image_id, last_indexed, archived, feature_id, admin_area_country_id, admin_area_state_id, admin_area_county_id, admin_area_place_id, array_to_json(highway_tags.tags) as tags
FROM sign.highwaysign as hs
         LEFT OUTER JOIN (
    SELECT ths.highwaysign_id, ARRAY_AGG(t.name) as tags FROM sign.tag_highwaysign ths
                                                                  INNER JOIN sign.tag t on t.id = ths.tag_id
    GROUP BY ths.highwaysign_id
) highway_tags ON hs.id = highway_tags.highwaysign_id
WHERE admin_area_country_id = $1
ORDER BY date_taken DESC, id
LIMIT $2 OFFSET $3;