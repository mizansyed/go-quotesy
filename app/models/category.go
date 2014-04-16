package models

import "github.com/revel/revel"
import "strconv"
import "fmt"
import "database/sql"

type Category struct {
	Id       int           `db:"category_id"`
	ParentId sql.NullInt64 `db:"parent_id"`
	Name     string        `db:"name"`
	IsActive bool          `db:"is_active"`
	Created  int64         `db:"created"`
	Updated  int64         `db:"updated"`
	Version  int           `db:"version"`
}

func (category Category) IdStr() string {
	return strconv.Itoa(category.Id)
}

func (category Category) SetFlashValues(c *revel.Flash) {
	c.Data["category.Name"] = category.Name

	if category.ParentId.Valid {
		c.Data["category.ParentId"] = fmt.Sprintf("%v", category.ParentId.Int64)
	}
	c.Data["category.IsActive"] = fmt.Sprintf("%v", category.IsActive)
}

func (category Category) Validate(v *revel.Validation) {
	v.Required(category.Name).Message("Category Name is required")
}
