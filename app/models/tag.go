package models

//import "github.com/revel/revel"
//import "strconv"
//import "fmt"
//import "database/sql"

type Tag struct {
	Id       	int   	        `db:"tag_id"`
	Name     	string	        `db:"name"`
	Slug 		string			`db:"slug"`
    Count   	int				`db:"count"`
}
