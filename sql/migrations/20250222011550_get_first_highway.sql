-- +goose Up
-- +goose StatementBegin
create function sign.get_first_highway(highway integer)
    returns TABLE(prev_feat integer, next_feat integer)
    language plpgsql
as
$$
BEGIN
    return query
        WITH
            links_per_highway AS (
                select
                    case when feature_link_highway.is_descending
                             then fl.to_feature else fl.from_feature
                        end as from_feature,
                    case when feature_link_highway.is_descending
                             then fl.from_feature  else fl.to_feature
                        end as to_feature
                from sign.feature_link_highway
                         inner join sign.feature_link fl on feature_link_highway.feature_link_id = fl.id
                where feature_link_highway.highway_id = highway ),
            all_features AS (
                select * from links_per_highway
                                  inner join sign.feature f1 on links_per_highway.from_feature = f1.id
                                  inner join sign.feature f2 on links_per_highway.to_feature = f2.id
                where from_feature not in (select  to_feature from links_per_highway))
        select from_feature as f, to_feature as t from all_features;
END
$$;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop function sign.get_first_highway;
-- +goose StatementEnd
