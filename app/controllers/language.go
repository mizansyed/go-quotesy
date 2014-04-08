package controllers

import "github.com/revel/revel"

type Language struct {
	*revel.Controller
}

func (c Language) Index() revel.Result {
	greeting := "Should show all the languages!"
	return c.Render(greeting)
}
