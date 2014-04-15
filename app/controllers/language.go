package controllers

import "github.com/revel/revel"
import "github.com/muizsyed/go-quotesy/app/models"
import "github.com/muizsyed/go-quotesy/app/routes"
import "fmt"
import "strconv"
import "time"

type Language struct {
	GorpController
}

func (c Language) Index(page int) revel.Result {

	if page < 1 {
		page = 1
	}

	var err error

	count := c.GetCount("Language")

	offset := 10 * (page - 1)
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
	language.SetFlashValues(&c.Flash)

	c.Session["language-version"] = strconv.Itoa(language.Version)

	return c.Render(language)
}

func (c Language) UpdateLanguage(language models.Language) revel.Result {

	language.Validate(c.Validation)

	var oldlanguage = c.Session["language-version"]

	oldVersion, _ := strconv.Atoi(oldlanguage)

	language.Version = oldVersion
	language.Updated = time.Now().UnixNano()

	if c.Validation.HasErrors() {

		language.SetFlashValues(&c.Flash)
		c.RenderArgs["language"] = language
		//c.Validation.Keep()
		c.FlashParams()
		return c.RenderTemplate(c.Name + "/Edit." + c.Request.Format)
	}

	var err error

	_, err = c.Txn.Update(&language)

	if err != nil {
		panic(err)
	}

	c.Flash.Success("Updated Langauge, " + language.Name)

	return c.Redirect(routes.Language.Index(1))
}

func (c Language) New() revel.Result {

	c.RenderArgs["action"] = "New"

	return c.RenderTemplate(c.Name + "/Edit." + c.Request.Format)
}

func (c Language) AddLanguage(language models.Language) revel.Result {

	language.Validate(c.Validation)

	if c.Validation.HasErrors() {

		c.RenderArgs["action"] = "New"
		//c.Validation.Keep()
		c.FlashParams()

		return c.RenderTemplate(c.Name + "/Edit." + c.Request.Format)
	}

	language.Created = time.Now().UnixNano()
	language.Updated = language.Created

	var err error

	err = c.Txn.Insert(&language)
	if err != nil {
		panic(err)
	}

	c.Flash.Success("Added Langauge, " + language.Name)

	return c.Redirect(routes.Language.Index(1))
}

func (c Language) Delete(id int) revel.Result {

	var language models.Language

	err := c.Txn.SelectOne(&language, "select * from Language where language_id=?", id)

	if err != nil {
		panic(err)
	}

	_, err = c.Txn.Delete(&language)

	if err != nil {
		panic(err)
	}

	c.Flash.Success(fmt.Sprintf("Deleted %s", language.Name))

	return nil

}
