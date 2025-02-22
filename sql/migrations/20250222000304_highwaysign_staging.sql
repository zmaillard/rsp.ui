-- +goose Up
-- +goose StatementBegin
create table sign.highwaysign_staging
(
    id           serial,
    image_width  integer,
    image_height integer,
    date_taken   timestamp,
    imageid      bigint,
    latitude     double precision,
    longitude    double precision
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table sign.highwaysign_staging;
-- +goose StatementEnd
