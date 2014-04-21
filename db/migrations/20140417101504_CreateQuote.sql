
-- +goose Up
CREATE TABLE Quote (
	quote_id int NOT NULL AUTO_INCREMENT,
	language_id int NOT NULL,
	description varchar(1000),
	content varchar(1000),
    is_active tinyint,
	created bigint,
	updated bigint,
	version int,
	PRIMARY KEY(quote_id),
	FOREIGN KEY(language_id) REFERENCES Language(language_id)
);

CREATE TABLE QuoteCategory (
	quotecategory_id int NOT NULL AUTO_INCREMENT,
	quote_id int NOT NULL,
	category_id int NOT NULL,
	PRIMARY KEY(quotecategory_id),
	FOREIGN KEY(quote_id) REFERENCES Quote(quote_id),
	FOREIGN KEY(category_id) REFERENCES Category(category_id)
);

CREATE TABLE QuoteTag (
	quotetag_id int NOT NULL AUTO_INCREMENT,
	quote_id int NOT NULL,
	tag_id int NOT NULL,
	PRIMARY KEY(quotetag_id),
	FOREIGN KEY(quote_id) REFERENCES Quote(quote_id),
	FOREIGN KEY(tag_id) REFERENCES Tag(tag_id)
);

-- +goose Down
DROP TABLE QuoteTag, QuoteCategory, Quote;


