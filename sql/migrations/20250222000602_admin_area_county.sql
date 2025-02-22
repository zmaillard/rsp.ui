-- +goose Up
-- +goose StatementBegin
create table sign.admin_area_county
(
    id                 integer generated always as identity
        constraint admin_area_county_pk
            primary key,
    name               varchar(255),
    slug               varchar(255),
    admin_area_stateid integer
        constraint admin_area_county_admin_area_state_id_fk
            references sign.admin_area_state,
    image_count        integer
);

create index admin_area_county_name_index
    on sign.admin_area_county (name);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table sign.admin_area_county;
-- +goose StatementEnd
