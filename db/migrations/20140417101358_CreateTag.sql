
-- +goose Up
CREATE TABLE Tag (
	tag_id int NOT NULL AUTO_INCREMENT,
    name varchar(100),
	slug varchar(100) ,
	count int NOT NULL,
	PRIMARY KEY(tag_id, slug),
	CONSTRAINT uc_Slug UNIQUE (slug)
);

-- +goose Down
DROP TABLE Tag;
