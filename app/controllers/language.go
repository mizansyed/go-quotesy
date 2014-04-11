package controllers

import "github.com/revel/revel"
import "github.com/muizsyed/go-quotesy/app/models"
import _ "fmt"

type Language struct {
	GorpController
}

func (c Language) Index(page int) revel.Result {
	
	var err error
	
	//var langs []models.Language
	
	languages, err := c.Txn.Select(models.Language{}, `select * from Language`)
	
	if err != nil {
		panic(err)
	}
	
	//greeting := "Should show all the languages!"
	return c.Render(languages)
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