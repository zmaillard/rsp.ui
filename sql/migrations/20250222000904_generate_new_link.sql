-- +goose Up
-- +goose StatementBegin
create function sign.generate_new_link() returns trigger
    language plpgsql
as
$$
BEGIN
    IF NEW.from_feature is not null and NEW.to_feature is not null THEN
        NEW.link := (SELECT ST_MakeLine(point ORDER BY order_field)::geography
                     FROM (
                              select 1 as order_field, point::geometry from sign.feature where id = NEW.from_feature
                              UNION ALL
                              select 2 as order_field, point::geometry from sign.feature where id = NEW.to_feature ) fquery);
    END IF;

    RETURN NEW;
end;
$$;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop function sign.generate_new_link();
-- +goose StatementEnd
