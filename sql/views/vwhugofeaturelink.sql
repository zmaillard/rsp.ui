create view sign.vwhugofeaturelink (id, from_feature, to_feature, road_name, highways, to_point, from_point) as
SELECT fl.id,
       fl.from_feature,
       fl.to_feature,
       fl.road_name,
       highways.highways,
       tf.point AS to_point,
       ff.point AS from_point
FROM sign.feature_link fl
         LEFT JOIN sign.feature tf ON fl.to_feature = tf.id
         LEFT JOIN sign.feature ff ON fl.from_feature = ff.id
         LEFT JOIN (SELECT flh.feature_link_id,
                           array_agg(h.slug) AS highways
                    FROM sign.feature_link_highway flh
                             JOIN sign.highway h ON flh.highway_id = h.id
                    GROUP BY flh.feature_link_id) highways ON fl.id = highways.feature_link_id;
