package models

import _"database/sql"

type Quote struct {
	Id       	int     		`db:"quote_id"`
	LanguageId 	int	 			`db:"language_id"`
	Description string 			`db:"description"`
	Content  	string			`db:"content"`
	IsActive 	bool 	 		`db:"is_active"`
	Created  	int64 	    	`db:"created"`
	Updated  	int64 			`db:"updated"`
	Version  	int   			`db:"version"`
}

type QuoteVM struct {
	
	Quote Quote
	CategoryId int
	Tags []string
}

type QuoteCategory struct {
	Id int `db:"quotecategory_id"`
	QuoteId int `db:"quote_id"`
	CategoryId int `db:"category_id"`
}

type QuoteTag struct {
	Id int `db:"quotetag_id"`
	QuoteId int `db:"quote_id"`
	TagId int `db:"tag_id"`
}