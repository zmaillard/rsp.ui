-- +goose Up
-- +goose StatementBegin
create view sign.vwhugofeaturelink (id, from_feature, to_feature, road_name, highways, to_point, from_point) as
SELECT fl.id,
       fl.from_feature,
       fl.to_feature,
       fl.road_name,
       cast(highways.highways as text[]) as highways,
       cast(tf.point as geometry) AS to_point,
       cast(ff.point AS geometry)  AS from_point
FROM sign.feature_link fl
         LEFT JOIN sign.feature tf ON fl.to_feature = tf.id
         LEFT JOIN sign.feature ff ON fl.from_feature = ff.id
         LEFT JOIN (SELECT flh.feature_link_id,
                           array_agg(h.slug) AS highways
                    FROM sign.feature_link_highway flh
                             JOIN sign.highway h ON flh.highway_id = h.id
                    GROUP BY flh.feature_link_id) highways ON fl.id = highways.feature_link_id;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop view sign.vwhugofeaturelink;
-- +goose StatementEnd
