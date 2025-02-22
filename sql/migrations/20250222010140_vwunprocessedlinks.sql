-- +goose Up
-- +goose StatementBegin
create view sign."vwUnprocessedLinks"(id, from_feature, to_feature, road_name, temp_placeholder, link) as
SELECT feature_link.id,
       feature_link.from_feature,
       feature_link.to_feature,
       feature_link.road_name,
       feature_link.temp_placeholder,
       feature_link.link
FROM sign.feature_link
         LEFT JOIN sign.feature_link_highway ON feature_link.id = feature_link_highway.feature_link_id
WHERE feature_link.road_name IS NULL
  AND feature_link_highway.id IS NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP view sign."vwUnprocessedLinks";
-- +goose StatementEnd
