-- +goose Up
-- +goose StatementBegin
create view sign.vwindexsign
            (imageid, title, sign_description, date_taken, country_slug, country_name, state_slug, state_name,
             county_name, county_slug, place_name, place_slug, tagitems, hwys, point, last_indexed, last_update,
             quality)
as
SELECT hs.imageid,
       hs.title,
       hs.sign_description,
       hs.date_taken,
       aac.slug   AS country_slug,
       aac.name   AS country_name,
       aas.slug   AS state_slug,
       aas.name   AS state_name,
       aacnt.name AS county_name,
       aacnt.slug AS county_slug,
       aaplc.name AS place_name,
       aaplc.slug AS place_slug,
       taglist.tagitems,
       hwylist.hwys,
       hs.point,
       hs.last_indexed,
       hs.last_update,
       hs.quality
FROM sign.highwaysign hs
         JOIN sign.admin_area_country aac ON aac.id = hs.admin_area_country_id
         JOIN sign.admin_area_state aas ON aas.id = hs.admin_area_state_id
         LEFT JOIN sign.admin_area_county aacnt ON aacnt.id = hs.admin_area_county_id
         LEFT JOIN sign.admin_area_place aaplc ON aaplc.id = hs.admin_area_place_id
         LEFT JOIN (SELECT tag_highwaysign.highwaysign_id,
                           array_agg(DISTINCT tag.name) AS tagitems
                    FROM sign.tag_highwaysign
                             JOIN sign.tag ON tag_highwaysign.tag_id = tag.id
                    GROUP BY tag_highwaysign.highwaysign_id) taglist ON hs.id = taglist.highwaysign_id
         LEFT JOIN (SELECT highwaysign_highway.highwaysign_id,
                           json_agg(json_build_object('slug', highway.slug, 'name', highway.highway_name)) AS hwys
                    FROM sign.highwaysign_highway
                             JOIN sign.highway ON highwaysign_highway.highway_id = highway.id
                    GROUP BY highwaysign_highway.highwaysign_id) hwylist ON hs.id = hwylist.highwaysign_id;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop view sign.vwindexsign;
-- +goose StatementEnd
