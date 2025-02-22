-- name: DeleteStaging :exec
DELETE FROM sign.highwaysign_staging WHERE id = $1;

-- name: InsertStaging :one
INSERT INTO sign.highwaysign_staging (image_width, image_height, date_taken, imageid, latitude, longitude) VALUES (sqlc.arg(imageWidth), sqlc.arg(imageHeight), sqlc.arg(datetaken), sqlc.arg(imageid), sqlc.narg(latitude), sqlc.narg(longitude)) RETURNING *;

-- name: GetStaging :one
SELECT id, image_width, image_height, date_taken, imageid, latitude, longitude FROM sign.highwaysign_staging WHERE id = $1;


-- name: GetAllStaging :many
SELECT id, image_width, image_height, date_taken, imageid, latitude, longitude FROM sign.highwaysign_staging order by date_taken;