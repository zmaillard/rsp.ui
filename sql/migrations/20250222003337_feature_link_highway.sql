-- +goose Up
-- +goose StatementBegin
create table sign.feature_link_highway
(
    id              integer generated always as identity
        constraint feature_link_highway_pk
            primary key,
    highway_id      integer
        constraint feature_link_highway_highway_id_fk
            references sign.highway,
    feature_link_id integer
        constraint feature_link_highway_feature_link_id_fk
            references sign.feature_link,
    is_descending   boolean
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table sign.feature_link_highway;
-- +goose StatementEnd
