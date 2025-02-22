-- name: UpdateAllHighwaySortingsOnSign :exec
UPDATE sign.highwaysign_highway SET highway_id = sqlc.arg(to_highway)::int WHERE highwaysign_id = sqlc.arg(highway_sign)::int AND highway_id = sqlc.arg(from_highway)::int;


-- name: GetScopes :many
SELECT id, scope FROM sign.highway_scope;

-- name: GetScope :one
SELECT id, scope FROM sign.highway_scope WHERE id = $1;

-- name: GetScopeByName :one
SELECT id, scope FROM sign.highway_scope WHERE scope = $1;

-- name: GetHighway :one
SELECT * FROM sign.highway h WHERE h.id = $1;

-- name: GetHighwayByName :one
SELECT * FROM sign.highway h WHERE h.highway_name = $1;

-- name: GetAllHighways :many
SELECT * FROM sign.highway h;

-- name: GetHighwaysOnSign :many
SELECT hsh.is_to, h.id, h.highway_name FROM sign.highwaysign_highway hsh
    INNER JOIN sign.highway h on h.id = hsh.highway_id
    WHERE hsh.highwaysign_id = $1;

-- name: GetHighwayType :one
SELECT id, highway_type_name, sort, slug, display_image_id, image_count, admin_area_country_id
FROM sign.highway_type WHERE id = $1;

-- name: GetHighwayTypeByName :one
SELECT id, highway_type_name, sort, slug, display_image_id, image_count, admin_area_country_id
FROM sign.highway_type WHERE highway_type_name = $1;

-- name: GetAllHighwayTypes :many
SELECT id, highway_type_name, sort, slug, display_image_id, image_count, admin_area_country_id FROM sign.highway_type;

-- name: CreateHighway :one
INSERT INTO sign.highway (highway_name, highway_type_id, admin_area_state_id, admin_area_country_id, image_name, slug, scope_id, date_added, sort_number)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING *;

-- name: CreateHighwayType :one
INSERT INTO sign.highway_type (highway_type_name, admin_area_country_id, slug, sort)
VALUES ($1, $2, $3, $4) RETURNING *;

-- name: InsertHighwaySorting :one
INSERT INTO sign.highwaysign_highway (highway_id, highwaysign_id, is_to) VALUES ($1, $2, $3) RETURNING *;

-- name: RemoveHighwaySorting :exec
DELETE FROM sign.highwaysign_highway WHERE highway_id = $1 AND  highwaysign_id = $2;

-- name: GetHighwaysForStateAndCountry :many
SELECT * FROM sign.highway WHERE admin_area_state_id = $1 OR (admin_area_state_id is null AND admin_area_country_id = $2) ORDER BY highway_name;

-- name: GetHighwaysStartWith :many
SELECT * FROM sign.highway WHERE highway_name ILIKE $1 ORDER BY highway_name;

-- name: UpdateImage :exec
UPDATE sign.highway SET image_name = $1 WHERE id = $2;

-- name: UpdateHighwayType :exec
UPDATE sign.highway SET highway_type_id = $1 WHERE id = $2;
