-- +goose Up
-- +goose StatementBegin
create table sign.admin_area_state
(
    id                   integer generated always as identity
        constraint admin_area_state_pk
            primary key,
    name                 varchar(255),
    subdivision_name     varchar(255),
    slug                 varchar(255),
    adminarea_country_id integer
        constraint admin_area_state_admin_area_country_id_fk
            references sign.admin_area_country,
    featured_sign_id     integer,
    image_count          integer
);

create index admin_area_state_name_index
    on sign.admin_area_state (name);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP table sign.admin_area_state;
-- +goose StatementEnd
