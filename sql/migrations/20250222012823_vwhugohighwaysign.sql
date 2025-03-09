-- +goose Up
-- +goose StatementBegin
create view sign.vwhugohighwaysign
            (id, title, sign_description, feature_id, date_taken, imageid, flickrid, point, country_slug, state_slug,
             place_slug, county_slug, tags, categories, highways, is_to, image_height, image_width, quality)
as
SELECT hs.id,
       hs.title,
       hs.sign_description,
       hs.feature_id,
       hs.date_taken,
       hs.imageid,
       cast(hs.flickrid as text) as flickrid,
       cast(hs.point as geometry) as point,
       cast(aac.slug as text)     AS country_slug,
       cast(aas.slug as text)     AS state_slug,
       cast(placejoin.place_slug as text) as place_slug,
       cast(countyjoin.county_slug as text) as county_slug,
       cast(tagjoin.tags as text[]) as tags,
       cast(catjoin.tags as text[]) AS categories,
       cast(hwyjoin.highways as text[]) as highways,
       cast(hwyjoin.is_to as text[]) as is_to,
       hs.image_height,
       hs.image_width,
       hs.quality
FROM sign.highwaysign hs
         JOIN sign.admin_area_country aac ON aac.id = hs.admin_area_country_id
         LEFT JOIN sign.admin_area_state aas ON aas.id = hs.admin_area_state_id
         LEFT JOIN (SELECT p.id,
                           (s.slug::text || '_'::text) || p.slug::text AS place_slug
                    FROM sign.admin_area_place p
                             JOIN sign.admin_area_state s ON p.admin_area_stateid = s.id
                    WHERE p.admin_area_stateid = s.id) placejoin ON hs.admin_area_place_id = placejoin.id
         LEFT JOIN (SELECT c.id,
                           (s.slug::text || '_'::text) || c.slug::text AS county_slug
                    FROM sign.admin_area_county c
                             JOIN sign.admin_area_state s ON c.admin_area_stateid = s.id
                    WHERE c.admin_area_stateid = s.id) countyjoin ON hs.admin_area_county_id = countyjoin.id
         LEFT JOIN (SELECT ths.highwaysign_id,
                           array_agg(t.slug) AS tags
                    FROM sign.tag_highwaysign ths
                             JOIN sign.tag t ON ths.tag_id = t.id
                    GROUP BY ths.highwaysign_id) tagjoin ON hs.id = tagjoin.highwaysign_id
         LEFT JOIN (SELECT ths.highwaysign_id,
                           array_agg(t.slug) AS tags
                    FROM sign.tag_highwaysign ths
                             JOIN sign.tag t ON ths.tag_id = t.id
                    WHERE t.is_category = true
                    GROUP BY ths.highwaysign_id) catjoin ON hs.id = catjoin.highwaysign_id
         LEFT JOIN (SELECT hhs.highwaysign_id,
                           array_agg(h.slug ORDER BY ht.sort, h.sort_number) AS highways,
                           array_agg(h.slug) FILTER (WHERE hhs.is_to)        AS is_to
                    FROM sign.highwaysign_highway hhs
                             JOIN sign.highway h ON hhs.highway_id = h.id
                             JOIN sign.highway_type ht ON h.highway_type_id = ht.id
                    GROUP BY hhs.highwaysign_id) hwyjoin ON hs.id = hwyjoin.highwaysign_id;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop view sign.vwhugohighwaysign;
-- +goose StatementEnd
