-- +goose Up
-- +goose StatementBegin
create table sign.feature_link_alias_type
(
    id   integer not null
        constraint feature_link_alias_type_pk
            primary key,
    name varchar(255)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table sign.feature_link_alias_type;
-- +goose StatementEnd
