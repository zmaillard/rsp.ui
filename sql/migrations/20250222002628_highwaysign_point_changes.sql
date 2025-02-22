-- +goose Up
-- +goose StatementBegin
create trigger highwaysign_point_changes
    before update
        of point
    on sign.highwaysign
    for each row
execute procedure sign.log_highwaysign_changes();
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop trigger highwaysign_point_changes on sign.highwaysign;
-- +goose StatementEnd
