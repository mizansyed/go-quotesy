package controllers

import "github.com/revel/revel"
import "github.com/muizsyed/go-quotesy/app/models"
import "github.com/muizsyed/go-quotesy/app/routes"
import "fmt"

type Language struct {
	GorpController
}

func (c Language) Index(page int) revel.Result {
	
	if page < 1 {
		page = 1
	}
	
	var err error

	count := c.GetCount("Language")
	
	offset:= 10 * (page - 1)
	languages, err := c.Txn.Select(models.Language{}, "select * from Language order by language_id limit 10 offset ?", offset)
	
	if err != nil {
		panic(err)
	}
	
	return c.Render(languages, count, page)
}

func (c Language) Edit(id int) revel.Result {
	
	var language models.Language
	
	err := c.Txn.SelectOne(&language, "select * from Language where language_id=?", id)
	
	if err != nil {
		panic(err)
	}
	
	// Dont like the following - but need to set the values initially
	c.Flash.Data["language.Name"] = language.Name
	c.Flash.Data["language.Description"] = language.Description
	c.Flash.Data["language.IsActive"] = fmt.Sprintf("%v", language.IsActive)
	
	c.RenderArgs["action"] = "Edit";
	c.RenderArgs["language"] = language;
	c.RenderArgs["id"] = id;
	
	return c.RenderTemplate(c.Name + "/New." + c.Request.Format)

}

func (c Language) UpdateLanguage(language models.Language) revel.Result {
	
	// OPTIMISTIC LOCKING?
	var err error
	
	_, err = c.Txn.Update(&language)
	
	if err != nil {
		panic(err)
	}
	
	c.Flash.Success("Updated Langauge, " + language.Name)
	
	return c.Redirect(routes.Language.Index(1))
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
	
	c.Flash.Success("Added Langauge, " + language.Name)
	
	return c.Redirect(routes.Language.Index(1))
}