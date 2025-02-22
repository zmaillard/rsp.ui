-- +goose Up
-- +goose StatementBegin
create table sign.feature_link_alias
(
    id                    integer generated always as identity
        constraint feature_link_alias_pk
            primary key,
    name                  varchar(255) not null,
    feature_alias_type_id integer
        constraint feature_link_alias_feature_link_alias_type_id_fk
            references sign.feature_link_alias_type,
    feature_link_id       integer
        constraint feature_link_alias_feature_link_id_fk
            references sign.feature_link
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table sign.feature_link_alias;
-- +goose StatementEnd
