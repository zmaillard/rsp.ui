-- +goose Up
-- +goose StatementBegin
create view sign.vwhugocountry
            (id, country_name, country_slug, subdivision_name, image_count, states, featured, highway_types) as
SELECT country.id,
       country.name              AS country_name,
       country.slug              AS country_slug,
       country.subdivision_name,
       countrycount.image_count,
       states.countrystates      AS states,
       hs.imageid                AS featured,
       highwaytypes.countrytypes AS highway_types
FROM sign.admin_area_country country
         JOIN (SELECT highwaysign.admin_area_country_id,
                      count(*) AS image_count
               FROM sign.highwaysign
               GROUP BY highwaysign.admin_area_country_id) countrycount
              ON countrycount.admin_area_country_id = country.id
         LEFT JOIN (SELECT s.adminarea_country_id,
                           json_agg(json_build_object('slug', s.slug, 'name', s.name)) AS countrystates
                    FROM sign.admin_area_state s
                    GROUP BY s.adminarea_country_id) states ON states.adminarea_country_id = country.id
         LEFT JOIN sign.highwaysign hs ON country.featured_sign_id = hs.id
         LEFT JOIN (SELECT uniquetypes.admin_area_country_id,
                           json_agg(json_build_object('slug', uniquetypes.slug, 'name',
                                                      uniquetypes.highway_type_name)) AS countrytypes
                    FROM (SELECT DISTINCT hs_1.admin_area_country_id,
                                          highway_type.highway_type_name,
                                          highway_type.slug
                          FROM sign.highwaysign_highway hsh
                                   JOIN sign.highway h ON h.id = hsh.highway_id
                                   JOIN sign.highway_type ON h.highway_type_id = highway_type.id
                                   JOIN sign.highwaysign hs_1 ON hsh.highwaysign_id = hs_1.id) uniquetypes
                    GROUP BY uniquetypes.admin_area_country_id) highwaytypes
                   ON country.id = highwaytypes.admin_area_country_id;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop view sign.vwhugocountry;
-- +goose StatementEnd
