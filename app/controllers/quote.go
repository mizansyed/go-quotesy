package controllers

import "github.com/revel/revel"
import "github.com/muizsyed/go-quotesy/app/models"
import "github.com/muizsyed/go-quotesy/app/routes"
import "fmt"
import _"strconv"
import "time"
import "regexp"

type Quote struct {
	GorpController
}

func (c Quote) New() revel.Result {

	c.RenderArgs["action"] = "New"
	c.RenderArgs["languages"] = c.GetLanguages(nil)

	return c.RenderTemplate(c.Name + "/Edit." + c.Request.Format)
}

/*
func (c Quote) Slug(tag string) string {
	
}
*/

func (c Quote) Slug(tags []string) []string {
	
	var slugs = make([]string,len(tags))
	
	unwantedchars, _ := regexp.Compile("[^a-zA-Z0-9]")
	excess_underscores, _ := regexp.Compile("[-]{2,}")
	
	for i := 0; i < len(tags); i++ {
	    slugs[i] = unwantedchars.ReplaceAllLiteralString(tags[i], "-")
		slugs[i] = excess_underscores.ReplaceAllLiteralString(slugs[i], "-")
	}
	
	return slugs;
}

func (c Quote) Add(quote models.QuoteVM) revel.Result {
	fmt.Println("Content")
	fmt.Println(quote.Quote.Content)
	fmt.Println(quote.CategoryId)
	
	// create slugs from tags
	slugs := c.Slug(quote.Tags)
	
	// Set date time
	quote.Quote.Created = time.Now().UnixNano()
	quote.Quote.Updated = quote.Quote.Created
	
	
	err := c.Txn.Insert(&quote.Quote)
	if err != nil {
		panic(err)
	}

	// Link to a category - Soon this will be multiple

	var linkQuoteCategory models.QuoteCategory
	
	linkQuoteCategory.QuoteId = quote.Quote.Id
	linkQuoteCategory.QuoteId = quote.CategoryId
	
	err = c.Txn.Insert(&linkQuoteCategory)
	
	fmt.Println(quote.Tags)
	fmt.Println(slugs)
	
	for i := 0; i < len(slugs); i++ {
		// If tag is new
		
		//
	}
	
	
	c.Flash.Success("Added Quote")

	
	
    
	
	/*err := c.Txn.Exec()
	
	INSERT INTO table (id,a,b,c,d,e,f,g)
VALUES (1,2,3,4,5,6,7,8) 
ON DUPLICATE KEY
    UPDATE a=a*/
	
	return c.Redirect(routes.Quote.New())
}

func (c Quote) GetLanguages(exclusion *int) []interface{} {

	var languages []interface{}
	var err error

	if exclusion == nil {
		languages, err = c.Txn.Select(models.Language{}, "select * from Language order by name")
	} else {
		languages, err = c.Txn.Select(models.Language{}, "select * from Language where language_id <> ? order by name", *exclusion)
	}

	if err != nil {
		panic(err)
	}

	return languages
}

/*
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
*/