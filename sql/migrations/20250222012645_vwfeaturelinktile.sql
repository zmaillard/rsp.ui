-- +goose Up
-- +goose StatementBegin
create view sign.vwfeaturelinktile(id, from_feature, to_feature, road_name, temp_placeholder, link) as
SELECT feature_link.id,
       feature_link.from_feature,
       feature_link.to_feature,
       feature_link.road_name,
       feature_link.temp_placeholder,
       feature_link.link::geometry AS link
FROM sign.feature_link;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP view sign.vwfeaturelinktile;
-- +goose StatementEnd
