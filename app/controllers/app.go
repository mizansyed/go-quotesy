package controllers

import "github.com/revel/revel"

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	greeting := "Store quotes you want to remember!"
	return c.Render(greeting)
}
