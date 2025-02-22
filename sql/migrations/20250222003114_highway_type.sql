-- +goose Up
-- +goose StatementBegin
create table sign.highway_type
(
    id                    integer generated always as identity
        constraint highway_type_pk
            primary key,
    highway_type_name     varchar(50),
    sort                  integer,
    slug                  varchar(50),
    display_image_id      integer
        constraint highway_type_highwaysign_id_fk
            references sign.highwaysign,
    image_count           integer,
    admin_area_country_id integer
        constraint highway_type_admin_area_country_id_fk
            references sign.admin_area_country
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP table sign.highway_type;
-- +goose StatementEnd
