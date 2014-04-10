package models



type Language struct
{
	Id int  `db:"language_id"`
	Name string `db:"name"`
	Description string `db:"description"`
	IsActive bool `db:"is_active"`
}