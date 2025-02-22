-- +goose Up
-- +goose StatementBegin
create view sign.vwhugofeature (id, point, name, signs, state_name, state_slug, country_name, country_slug) as
SELECT f.id,
       f.point,
       f.name,
       signs.signs,
       s.name AS state_name,
       s.slug AS state_slug,
       c.name AS country_name,
       c.slug AS country_slug
FROM sign.feature f
         LEFT JOIN (SELECT hs.feature_id,
                           array_agg(hs.imageid::text) AS signs
                    FROM sign.highwaysign hs
                    GROUP BY hs.feature_id) signs ON f.id = signs.feature_id
         LEFT JOIN sign.admin_area_country c ON f.admin_area_country_id = c.id
         LEFT JOIN sign.admin_area_state s ON f.admin_area_state_id = s.id;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP view sign.vwhugofeature;
-- +goose StatementEnd
