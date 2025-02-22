-- +goose Up
-- +goose StatementBegin
create view sign.vwhighwaysigntile
            (id, flickrid, date_taken, date_added, title, sign_description, image_width, image_height, point, imageid,
             lastsyncwithflickr, last_update, cropped_image_id, last_indexed, archived, feature_id,
             admin_area_country_id, admin_area_state_id, admin_area_county_id, admin_area_place_id)
as
SELECT highwaysign.id,
       highwaysign.flickrid,
       highwaysign.date_taken,
       highwaysign.date_added,
       highwaysign.title,
       highwaysign.sign_description,
       highwaysign.image_width,
       highwaysign.image_height,
       highwaysign.point::geometry AS point,
       highwaysign.imageid,
       highwaysign.lastsyncwithflickr,
       highwaysign.last_update,
       highwaysign.cropped_image_id,
       highwaysign.last_indexed,
       highwaysign.archived,
       highwaysign.feature_id,
       highwaysign.admin_area_country_id,
       highwaysign.admin_area_state_id,
       highwaysign.admin_area_county_id,
       highwaysign.admin_area_place_id
FROM sign.highwaysign;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop view sign.vwhighwaysigntile;
-- +goose StatementEnd
