-- +goose Up
-- +goose StatementBegin
create table sign.highwaysign_pending_changes
(
    id             integer generated always as identity
        constraint highwaysign_pending_changes_pk
            primary key,
    highwaysign_id integer
        constraint highwaysign_pending_changes_highwaysign_id_fk
            references sign.highwaysign,
    changed_on     timestamp
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table sign.highwaysign_pending_changes;
-- +goose StatementEnd
