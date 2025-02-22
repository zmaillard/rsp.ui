-- +goose Up
-- +goose StatementBegin
create table sign.highwaysign
(
    id                    serial
        constraint highwaysign_pk
            primary key,
    flickrid              varchar(25),
    date_taken            timestamp,
    date_added            timestamp,
    title                 varchar(255),
    sign_description      text,
    image_width           integer,
    image_height          integer,
    point                 geography,
    imageid               bigint,
    lastsyncwithflickr    date,
    last_update           timestamp,
    cropped_image_id      bigint,
    last_indexed          timestamp,
    archived              boolean,
    feature_id            integer
        constraint highwaysign_feature_id_fk
            references sign.feature,
    admin_area_country_id integer
        constraint highwaysign_admin_area_country_id_fk
            references sign.admin_area_country,
    admin_area_state_id   integer
        constraint highwaysign_admin_area_state_id_fk
            references sign.admin_area_state,
    admin_area_county_id  integer
        constraint highwaysign_admin_area_county_id_fk
            references sign.admin_area_county,
    admin_area_place_id   integer
        constraint highwaysign_admin_area_place_id_fk
            references sign.admin_area_place,
    quality integer default 0 not null

);

alter table sign.admin_area_country
    add constraint admin_area_country_highwaysign_id_fk
        foreign key (featured_sign_id) references sign.highwaysign;

alter table sign.admin_area_state
    add constraint admin_area_state_highwaysign_id_fk
        foreign key (featured_sign_id) references sign.highwaysign;

create index highwaysign_imageid_index
    on sign.highwaysign (imageid);

create index idx_highwaysign_geo
    on sign.highwaysign using gist (point);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE sign.admin_area_country DROP CONSTRAINT admin_area_country_highwaysign_id_fk;
ALTER TABLE sign.admin_area_state DROP CONSTRAINT admin_area_state_highwaysign_id_fk;

drop table sign.highwaysign;
-- +goose StatementEnd
