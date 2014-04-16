package controllers

import "github.com/revel/revel"
import "github.com/muizsyed/go-quotesy/app/models"
import "github.com/muizsyed/go-quotesy/app/routes"
import "fmt"
import "strconv"
import "time"

type Category struct {
	GorpController
}

func (c Category) Index(page int) revel.Result {

	if page < 1 {
		page = 1
	}

	var err error

	count := c.GetCount("Category")

	offset := 10 * (page - 1)
	categories, err := c.Txn.Select(models.Category{}, "select * from Category order by category_id limit 10 offset ?", offset)

	if err != nil {
		panic(err)
	}

	return c.Render(categories, count, page)
}

func (c Category) GetCategories(exclusion *int) []interface{} {

	var categories []interface{}
	var err error

	if exclusion == nil {
		categories, err = c.Txn.Select(models.Category{}, "select * from Category order by name")
	} else {
		categories, err = c.Txn.Select(models.Category{}, "select * from Category where category_id <> ? order by name", *exclusion)
	}

	if err != nil {
		panic(err)
	}

	return categories
}

func (c Category) New() revel.Result {

	c.RenderArgs["action"] = "New"
	c.RenderArgs["categories"] = c.GetCategories(nil)

	return c.RenderTemplate(c.Name + "/Edit." + c.Request.Format)
}

func (c Category) Add(category models.Category) revel.Result {

	fmt.Println("Entered AddCategory")
	category.Validate(c.Validation)

	if c.Validation.HasErrors() {
		c.RenderArgs["categories"] = c.GetCategories(nil)
		c.RenderArgs["action"] = "New"
		c.FlashParams()

		return c.RenderTemplate(c.Name + "/Edit." + c.Request.Format)
	}

	category.Created = time.Now().UnixNano()
	category.Updated = category.Created

	var err error

	err = c.Txn.Insert(&category)
	if err != nil {
		panic(err)
	}

	c.Flash.Success("Added Category, " + category.Name)

	return c.Redirect(routes.Category.Index(1))
}

func (c Category) Edit(id int) revel.Result {

	var category models.Category

	err := c.Txn.SelectOne(&category, "select * from Category where category_id=?", id)

	if err != nil {
		panic(err)
	}

	// Dont like the following - but need to set the values initially
	category.SetFlashValues(&c.Flash)

	c.RenderArgs["categories"] = c.GetCategories(&id)
	c.Session["category-version"] = strconv.Itoa(category.Version)

	return c.Render(category)
}

func (c Category) Update(category models.Category) revel.Result {

	fmt.Println("Parent Id")
	fmt.Println(c.Request.Form)
	category.Validate(c.Validation)

	var oldSessionVersion = c.Session["category-version"]

	oldVersion, _ := strconv.Atoi(oldSessionVersion)

	category.Version = oldVersion
	category.Updated = time.Now().UnixNano()

	if c.Validation.HasErrors() {

		category.SetFlashValues(&c.Flash)
		c.RenderArgs["categories"] = c.GetCategories(&category.Id)
		c.RenderArgs["category"] = category
		//c.Validation.Keep()
		c.FlashParams()
		return c.RenderTemplate(c.Name + "/Edit." + c.Request.Format)
	}

	var err error

	_, err = c.Txn.Update(&category)

	if err != nil {
		panic(err)
	}

	c.Flash.Success("Updated Category, " + category.Name)

	return c.Redirect(routes.Category.Index(1))
}
