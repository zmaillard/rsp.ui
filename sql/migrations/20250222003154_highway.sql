-- +goose Up
-- +goose StatementBegin
create table sign.highway
(
    id                    serial
        constraint highway_pk
            primary key,
    highway_name          varchar(50),
    scope_id              integer
        constraint highway_highway_scope_id_fk
            references sign.highway_scope,
    slug                  varchar(30),
    highway_type_id       integer
        constraint highway_highway_type_id_fk
            references sign.highway_type,
    image_name            varchar(50),
    date_added            date,
    sort_number           integer,
    admin_area_country_id integer
        constraint highway_admin_area_country_id_fk
            references sign.admin_area_country,
    admin_area_state_id   integer
        constraint highway_admin_area_state_id_fk
            references sign.admin_area_state
);

create index highway_highway_name_index
    on sign.highway (highway_name);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP table sign.highway;
-- +goose StatementEnd
