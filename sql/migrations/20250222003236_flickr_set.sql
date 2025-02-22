-- +goose Up
-- +goose StatementBegin
create table sign.flickr_set
(
    id                serial
        constraint flickr_set_pk
            primary key,
    highway_id        integer
        constraint flickr_set_highway_id_fk
            references sign.highway,
    flickr_set_id     varchar(25),
    primary_flickr_id varchar(25),
    last_updated      date
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP table sign.flickr_set;
-- +goose StatementEnd
