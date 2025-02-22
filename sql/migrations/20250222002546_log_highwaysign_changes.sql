-- +goose Up
-- +goose StatementBegin
create function sign.log_highwaysign_changes() returns trigger
    language plpgsql
as
$$
BEGIN
    IF NEW.point <> OLD.point THEN
        INSERT INTO sign.highwaysign_pending_changes(highwaysign_id, changed_on)
        VALUES(OLD.id,now());
    END IF;

    RETURN NEW;
end;
$$;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop function sign.log_highwaysign_changes();
-- +goose StatementEnd
