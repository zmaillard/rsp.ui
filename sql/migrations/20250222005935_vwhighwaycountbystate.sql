-- +goose Up
-- +goose StatementBegin
create view sign."vwHighwayCountByState"(admin_area_state_id, highway_id, image_count) as
SELECT h.admin_area_state_id,
       hh.highway_id,
       count(*) AS image_count
FROM sign.highwaysign_highway hh
         JOIN sign.highwaysign h ON h.id = hh.highwaysign_id
GROUP BY h.admin_area_state_id, hh.highway_id;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop view sign."vwHighwayCountByState";
-- +goose StatementEnd
