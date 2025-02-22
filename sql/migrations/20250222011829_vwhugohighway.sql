-- +goose Up
-- +goose StatementBegin
create view sign.vwhugohighway
            (id, highway_name, slug, sort_number, image_name, highway_type_slug, highway_type_name, states, counties,
             places, previous_features, next_features)
as
SELECT hwy.id,
       hwy.highway_name,
       hwy.slug,
       hwy.sort_number,
       hwy.image_name,
       ht.slug               AS highway_type_slug,
       ht.highway_type_name,
       hwyplaces.allstates   AS states,
       hwyplaces.allcounties AS counties,
       hwyplaces.allplaces   AS places,
       features.fromfeat     AS previous_features,
       features.tofeat       AS next_features
FROM sign.highway hwy
         JOIN sign.highway_type ht ON hwy.highway_type_id = ht.id
         LEFT JOIN (SELECT h.id,
                           array_agg(orf.prev_feat) AS fromfeat,
                           array_agg(orf.next_feat) AS tofeat
                    FROM sign.highway h,
                         LATERAL sign.get_first_highway(h.id) gf(prev_feat, next_feat),
                         LATERAL sign.get_ordered_features(h.id, gf.prev_feat) orf(prev_feat, next_feat)
                    GROUP BY h.id) features ON hwy.id = features.id
         LEFT JOIN (SELECT hs.highway_id,
                           array_agg(DISTINCT places.slug) FILTER (WHERE places.slug IS NOT NULL)     AS allplaces,
                           array_agg(DISTINCT states.slug) FILTER (WHERE states.slug IS NOT NULL)     AS allstates,
                           array_agg(DISTINCT counties.slug) FILTER (WHERE counties.slug IS NOT NULL) AS allcounties
                    FROM sign.highwaysign_highway hs
                             JOIN sign.highwaysign h ON h.id = hs.highwaysign_id
                             LEFT JOIN (SELECT place.id                                            AS pid,
                                               (state.slug::text || '_'::text) || place.slug::text AS slug
                                        FROM sign.admin_area_place place
                                                 JOIN sign.admin_area_state state ON place.admin_area_stateid = state.id) places
                                       ON h.admin_area_place_id = places.pid
                             LEFT JOIN (SELECT county.id                                            AS cid,
                                               (state.slug::text || '_'::text) || county.slug::text AS slug
                                        FROM sign.admin_area_county county
                                                 JOIN sign.admin_area_state state ON county.admin_area_stateid = state.id) counties
                                       ON h.admin_area_county_id = counties.cid
                             LEFT JOIN sign.admin_area_state states ON h.admin_area_state_id = states.id
                    GROUP BY hs.highway_id) hwyplaces ON hwy.id = hwyplaces.highway_id
WHERE hwyplaces.allstates IS NOT NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop view sign.vwhugohighway;
-- +goose StatementEnd
