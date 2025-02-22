-- +goose Up
-- +goose StatementBegin
create table sign.feature
(
    id                    serial
        constraint feature_pk
            primary key,
    point                 geography,
    name                  varchar(254),
    admin_area_country_id integer
        constraint feature_admin_area_country_id_fk
            references sign.admin_area_country,
    admin_area_state_id   integer
        constraint feature_admin_area_state_id_fk
            references sign.admin_area_state,
    featured              boolean default false,
    feature_type_id       integer
);

create index idx_feature_geo
    on sign.feature using gist (point);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP table sign.feature;
-- +goose StatementEnd
