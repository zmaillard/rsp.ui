-- +goose Up
-- +goose StatementBegin
create schema sign;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop schema sign;
-- +goose StatementEnd
