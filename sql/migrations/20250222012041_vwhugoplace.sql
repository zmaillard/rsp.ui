-- +goose Up
-- +goose StatementBegin
create view sign.vwhugoplace(id, place_slug, place_name, state_slug, state_name, image_count) as
SELECT p.id,
       p.slug       AS place_slug,
       p.name       AS place_name,
       aas.slug     AS state_slug,
       aas.name     AS state_name,
       counts.count AS image_count
FROM sign.admin_area_place p
         LEFT JOIN (SELECT hs.admin_area_place_id,
                           count(*) AS count
                    FROM sign.highwaysign hs
                    WHERE hs.admin_area_place_id IS NOT NULL
                    GROUP BY hs.admin_area_place_id) counts ON p.id = counts.admin_area_place_id
         JOIN sign.admin_area_state aas ON aas.id = p.admin_area_stateid;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop view sign.vwhugoplace;
-- +goose StatementEnd
