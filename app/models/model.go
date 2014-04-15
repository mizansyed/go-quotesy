package models

import "github.com/revel/revel"

type Model interface {
	SetFlashValues(c *revel.Flash)
	Validate(v *revel.Validation)
}
