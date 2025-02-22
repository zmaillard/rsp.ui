-- +goose Up
-- +goose StatementBegin
create view sign.vwhugohighwaytype
            (id, highway_type_name, highway_type_slug, sort, imagecount, imageid, highways, country) as
SELECT ht.id,
       ht.highway_type_name,
       ht.slug  AS highway_type_slug,
       ht.sort,
       signcounts.imagecount,
       highwaysign.imageid,
       highwayagg.highways,
       aac.slug AS country
FROM sign.highway_type ht
         JOIN sign.admin_area_country aac ON aac.id = ht.admin_area_country_id
         LEFT JOIN (SELECT signtype.highway_type_id,
                           count(signtype.highwaysign_id) AS imagecount
                    FROM (SELECT DISTINCT hsh.highwaysign_id,
                                          h.highway_type_id
                          FROM sign.highwaysign_highway hsh
                                   JOIN sign.highway h ON h.id = hsh.highway_id) signtype
                    GROUP BY signtype.highway_type_id) signcounts ON ht.id = signcounts.highway_type_id
         LEFT JOIN sign.highwaysign ON ht.display_image_id = highwaysign.id
         LEFT JOIN (SELECT highway.highway_type_id,
                           array_agg(highway.slug) AS highways
                    FROM sign.highway
                    GROUP BY highway.highway_type_id) highwayagg ON ht.id = highwayagg.highway_type_id;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop view sign.vwhugohighwaytype;
-- +goose StatementEnd
