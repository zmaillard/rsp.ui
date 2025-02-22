-- +goose Up
-- +goose StatementBegin
create table sign.feature_link
(
    id               integer generated always as identity
        constraint feature_link_pk
            primary key,
    from_feature     integer
        constraint feature_link_from_feature__fk
            references sign.feature,
    to_feature       integer
        constraint feature_link_to_feature__fk
            references sign.feature,
    road_name        varchar(255),
    temp_placeholder varchar(255),
    link             geography
);

create index idx_feature_link_geo
    on sign.feature_link using gist (link);

create trigger featurelink_feature_changes
    before update
        of from_feature, to_feature
    on sign.feature_link
    for each row
execute procedure sign.generate_new_link();

create trigger featurelink_feature_create
    before insert
    on sign.feature_link
    for each row
execute procedure sign.generate_new_link();
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop trigger featurelink_feature_create on sign.feature_link;
drop trigger featurelink_feature_changes on sign.feature_link;
drop table sign.feature_link;
-- +goose StatementEnd
