-- +goose Up
-- +goose StatementBegin
create table sign.highwaysign_highway
(
    id             integer generated always as identity
        constraint highwaysign_highway_pk
            primary key,
    highway_id     integer
        constraint highwaysign_highway_highway_id_fk
            references sign.highway,
    highwaysign_id integer
        constraint highwaysign_highway_highwaysign_id_fk
            references sign.highwaysign,
    is_to          boolean
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP table sign.highwaysign_highway;
-- +goose StatementEnd
