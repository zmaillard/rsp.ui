-- +goose Up
-- +goose StatementBegin
create view sign.vwfeaturelinkwithhighway(id, link, highways, temp_placeholder) as
SELECT fl.id,
       fl.link,
       array_to_string(i.hwyarr, '::'::text) AS highways,
       fl.temp_placeholder
FROM sign.feature_link fl
         JOIN (SELECT flh.feature_link_id,
                      array_agg(h.highway_name::text ||
                                CASE
                                    WHEN flh.is_descending THEN '^^DESC'::text
                                    ELSE ''::text
                                    END) AS hwyarr
               FROM sign.feature_link_highway flh
                        JOIN sign.highway h ON flh.highway_id = h.id
               GROUP BY flh.feature_link_id) i ON fl.id = i.feature_link_id;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP view sign.vwfeaturelinkwithhighway;
-- +goose StatementEnd
