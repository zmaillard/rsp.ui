-- name: GetStatesByCountry :many
SELECT * FROM sign.admin_area_state WHERE adminarea_country_id = $1;

-- name: GetStateByName :one
SELECT * FROM sign.admin_area_state WHERE name = $1 LIMIT 1;

-- name: GetCountry :one
SELECT * FROM sign.admin_area_country WHERE id = $1;

-- name: GetState :one
SELECT * FROM sign.admin_area_state WHERE id = $1;

-- name: GetAllCountries :many
SELECT * FROM sign.admin_area_country;


-- name: GetAllStates :many
SELECT * FROM sign.admin_area_state;

-- name: InsertAdminAreaCounty :one
INSERT INTO sign.admin_area_county (name, slug, admin_area_stateid) VALUES ($1, $2, $3) RETURNING *;

-- name: InsertAdminAreaPlace :one
INSERT INTO sign.admin_area_place (name, slug, admin_area_stateid) VALUES ($1, $2, $3) RETURNING *;


-- name: InsertAdminAreaState :one
INSERT INTO sign.admin_area_state (name, slug, adminarea_country_id) VALUES ($1, $2, $3) RETURNING *;

-- name: InsertAdminAreaStateWithSubdivision :one
INSERT INTO sign.admin_area_state (name, slug, subdivision_name, adminarea_country_id) VALUES ($1, $2, $3, $4) RETURNING *;


-- name: InsertAdminAreaCountry :one
INSERT INTO sign.admin_area_country (name, slug) VALUES ($1, $2) RETURNING *;


-- name: AdminAreaStateCountByName :one
SELECT COUNT(*) FROM sign.admin_area_state WHERE name = $1 AND adminarea_country_id = $2;

-- name: AdminAreaPlaceCountByName :one
SELECT COUNT(*) FROM sign.admin_area_place WHERE name = $1 AND admin_area_stateid = $2;

-- name: AdminAreaCountryCountByName :one
SELECT COUNT(*) FROM sign.admin_area_country WHERE name = $1;

-- name: AdminAreaCountyCountByName :one
SELECT COUNT(*) FROM sign.admin_area_county WHERE name = $1 AND admin_area_stateid = $2;

-- name: GetAdminAreaStateByName :one
SELECT * FROM sign.admin_area_state WHERE name = $1 AND adminarea_country_id = $2;

-- name: GetAdminAreaPlaceByName :one
SELECT * FROM sign.admin_area_place WHERE name = $1 AND admin_area_stateid = $2;

-- name: GetAdminAreaCountryByName :one
SELECT * FROM sign.admin_area_country WHERE name = $1;

-- name: GetAdminAreaCountyByName :one
SELECT * FROM sign.admin_area_county WHERE name = $1 AND admin_area_stateid = $2;