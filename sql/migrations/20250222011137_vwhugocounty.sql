-- +goose Up
-- +goose StatementBegin
create view sign.vwhugocounty(id, county_slug, county_name, state_slug, state_name, image_count) as
SELECT c.id,
       c.slug       AS county_slug,
       c.name       AS county_name,
       aas.slug     AS state_slug,
       aas.name     AS state_name,
       counts.count AS image_count
FROM sign.admin_area_county c
         LEFT JOIN (SELECT hs.admin_area_county_id,
                           count(*) AS count
                    FROM sign.highwaysign hs
                    WHERE hs.admin_area_county_id IS NOT NULL
                    GROUP BY hs.admin_area_county_id) counts ON c.id = counts.admin_area_county_id
         JOIN sign.admin_area_state aas ON aas.id = c.admin_area_stateid;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop view sign.vwhugocounty;
-- +goose StatementEnd
