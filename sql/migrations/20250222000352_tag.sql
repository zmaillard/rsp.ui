-- +goose Up
-- +goose StatementBegin
create table sign.tag
(
    id          integer generated always as identity
        constraint tag_pk
            primary key,
    name        varchar(255),
    slug        varchar(50),
    flickr_only boolean,
    category_details varchar(255),
    is_category boolean default false not null
);

create index tag_name_index
    on sign.tag (name);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP table sign.tag;
-- +goose StatementEnd
