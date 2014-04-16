
-- +goose Up
CREATE TABLE Category (
	category_id int NOT NULL AUTO_INCREMENT,
	parent_id int NULL,
    name text,
	is_active int,
	created int,
	updated int,
	version int,
	PRIMARY KEY(category_id),
	FOREIGN KEY (parent_id) REFERENCES Category (category_id)
);


-- +goose Down
DROP TABLE Category;

