-- +goose Up
-- +goose StatementBegin
create table sign.feature_type
(
    id   serial
        constraint feature_type_pk
            primary key,
    name varchar(255) not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP table sign.feature_type;
-- +goose StatementEnd
