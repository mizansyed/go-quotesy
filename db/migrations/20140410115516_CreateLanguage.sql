
-- +goose Up
CREATE TABLE language (
    language_id integer primary key not null,
    name text,
    description text,
	is_active int
);


-- +goose Down
DROP TABLE language;

