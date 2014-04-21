package models

import "github.com/revel/revel"
import _ "time"
import "fmt"
import "strconv"

type Language struct {
	Id          int    `db:"language_id"`
	Name        string `db:"name"`
	Description string `db:"description"`
	IsActive    bool   `db:"is_active"`
	Created     int64  `db:"created"`
	Updated     int64  `db:"updated"`
	Version     int    `db:"version"`
}

func (language Language) IdStr() string {
	return strconv.Itoa(language.Id)
}

func (language Language) SetFlashValues(c *revel.Flash) {
	c.Data["language.Name"] = language.Name
	c.Data["language.Description"] = language.Description
	c.Data["language.IsActive"] = fmt.Sprintf("%v", language.IsActive)
}

func (language Language) Validate(v *revel.Validation) {
	v.Required(language.Name).Message("Name is required")
	//v.Required(booking.Hotel)
	//v.Required(booking.CheckInDate)
	//v.Required(booking.CheckOutDate)

	//v.Match(booking.CardNumber, regexp.MustCompile(`\d{16}`)).
	//	Message("Credit card number must be numeric and 16 digits")

	//v.Check(booking.NameOnCard,
	//	revel.Required{},
	//	revel.MinSize{3},
	//	revel.MaxSize{70},
	//)
}
