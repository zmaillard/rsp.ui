create view sign.vwhugostate
            (id, state_name, state_slug, subdivision_name, image_count, highways, places, counties, featured,
             country_slug) as
SELECT state.id,
       state.name             AS state_name,
       state.slug             AS state_slug,
       state.subdivision_name,
       statecount.image_count,
       highways.statehwys     AS highways,
       places.stateplaces     AS places,
       counties.statecounties AS counties,
       hs.imageid             AS featured,
       c.slug                 AS country_slug
FROM sign.admin_area_state state
         JOIN (SELECT highwaysign.admin_area_state_id,
                      count(*) AS image_count
               FROM sign.highwaysign
               GROUP BY highwaysign.admin_area_state_id) statecount ON statecount.admin_area_state_id = state.id
         JOIN (SELECT hs_1.admin_area_state_id,
                      array_agg(DISTINCT h.slug) AS statehwys
               FROM sign.highwaysign_highway hsh
                        JOIN sign.highway h ON h.id = hsh.highway_id
                        JOIN sign.highwaysign hs_1 ON hsh.highwaysign_id = hs_1.id
               GROUP BY hs_1.admin_area_state_id) highways ON highways.admin_area_state_id = state.id
         LEFT JOIN (SELECT p.admin_area_stateid,
                           json_agg(json_build_object('slug', p.slug, 'name', p.name)) AS stateplaces
                    FROM sign.admin_area_place p
                    GROUP BY p.admin_area_stateid) places ON places.admin_area_stateid = state.id
         LEFT JOIN (SELECT c_1.admin_area_stateid,
                           json_agg(json_build_object('slug', c_1.slug, 'name', c_1.name)) AS statecounties
                    FROM sign.admin_area_county c_1
                    GROUP BY c_1.admin_area_stateid) counties ON counties.admin_area_stateid = state.id
         LEFT JOIN sign.highwaysign hs ON state.featured_sign_id = hs.id
         JOIN sign.admin_area_country c ON state.adminarea_country_id = c.id;
