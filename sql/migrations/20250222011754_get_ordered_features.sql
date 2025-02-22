-- +goose Up
-- +goose StatementBegin
create function sign.get_ordered_features(highway integer, feature integer)
    returns TABLE(prev_feat integer, next_feat integer)
    language plpgsql
as
$$
BEGIN
    return query
        WITH RECURSIVE
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
            ordered_features(from_feature_id, to_feature_id) AS
                (
                    select from_feature, to_feature FROM links_per_highway WHERE from_feature = feature
                    UNION ALL
                    SELECT from_feature, to_feature FROM links_per_highway l, ordered_features o
                    WHERE o.to_feature_id = l.from_feature
                )
        SELECT from_feature_id, to_feature_id from ordered_features;
END
$$;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop function sign.get_ordered_features;
-- +goose StatementEnd
