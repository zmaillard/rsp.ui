-- +goose Up
-- +goose StatementBegin
create function public.slugify(value text) returns text
    immutable
    strict
    language sql
as
$$
    -- removes accents (diacritic signs) from a given string --
WITH "unaccented" AS (
    SELECT unaccent("value") AS "value"
),
     -- lowercases the string
     "lowercase" AS (
         SELECT lower("value") AS "value"
         FROM "unaccented"
     ),
     -- replaces anything that's not a letter, number, hyphen('-'), or underscore('_') with a hyphen('-')
     "hyphenated" AS (
         SELECT regexp_replace("value", '[^a-z0-9\\-_]+', '-', 'gi') AS "value"
         FROM "lowercase"
     ),
     -- trims hyphens('-') if they exist on the head or tail of the string
     "trimmed" AS (
         SELECT regexp_replace(regexp_replace("value", '\\-+$', ''), '^\\-', '') AS "value"
         FROM "hyphenated"
     )
SELECT "value" FROM "trimmed";
$$;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop function public.slugify;
-- +goose StatementEnd
