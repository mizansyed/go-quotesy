package controllers

import "github.com/revel/revel"
import "github.com/muizsyed/go-quotesy/app/models"
import _ "fmt"

type Language struct {
	GorpController
}

func (c Language) Index() revel.Result {
	greeting := "Should show all the languages!"
	return c.Render(greeting)
}

func (c Language) New() revel.Result {
	return c.Render()
}

func (c Language) AddLanguage(language models.Language) revel.Result {
	
	var err error
	
	err = c.Txn.Insert(&language)
	if err != nil {
		panic(err)
	}
	
	return c.Redirect("new")
}