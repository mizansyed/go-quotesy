
-- +goose Up
CREATE TABLE Language (
    language_id int NOT NULL AUTO_INCREMENT,
    name varchar(100),
    description varchar(1000),
	is_active tinyint,
	created bigint,
	updated bigint,
	version int,
	PRIMARY KEY(language_id)
);


-- +goose Down
DROP TABLE Language;

