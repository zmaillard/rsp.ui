-- +goose Up
-- +goose StatementBegin
create table sign.admin_area_country
(
    id               serial
        constraint admin_area_country_pk
            primary key,
    name             varchar(255),
    subdivision_name varchar(255),
    slug             varchar(255),
    featured_sign_id integer,
    image_count      integer
);

create index admin_area_country_name_index
    on sign.admin_area_country (name);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table sign.admin_area_country;
-- +goose StatementEnd
