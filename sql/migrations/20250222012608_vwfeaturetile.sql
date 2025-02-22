-- +goose Up
-- +goose StatementBegin
create view sign.vwfeaturetile(id, point, name, admin_area_country_id, admin_area_state_id) as
SELECT feature.id,
       feature.point::geometry AS point,
       feature.name,
       feature.admin_area_country_id,
       feature.admin_area_state_id
FROM sign.feature;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop view sign.vwfeaturetile;
-- +goose StatementEnd
