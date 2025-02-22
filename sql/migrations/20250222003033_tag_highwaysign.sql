-- +goose Up
-- +goose StatementBegin
create table sign.tag_highwaysign
(
    id             integer generated always as identity
        constraint tag_highwaysign_pk
            primary key,
    tag_id         integer
        constraint tag_highwaysign_tag_id_fk
            references sign.tag,
    highwaysign_id integer
        constraint tag_highwaysign_highwaysign_id_fk
            references sign.highwaysign
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table sign.tag_highwaysign;
-- +goose StatementEnd
