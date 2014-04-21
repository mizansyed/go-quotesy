package controllers

import "github.com/revel/revel"
import "github.com/muizsyed/go-quotesy/app/models"
import "github.com/muizsyed/go-quotesy/app/routes"
import "fmt"
import _"strconv"
import "time"
import "regexp"
import "bytes"

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

func (c Quote) CreateGetTagStatement(slugs []string) string {
	
	var buffer bytes.Buffer
	
	buffer.WriteString("SELECT * FROM Tag WHERE Slug IN (")
	
	for i := 0; i < len(slugs); i++ {
		
		if i > 0 {
			buffer.WriteString(",")
		}
		
		buffer.WriteString(`"`)
		buffer.WriteString(slugs[i])
		buffer.WriteString(`"`)
		
		
	}
	
	buffer.WriteString(")")
	
	return buffer.String()
}

func (c Quote) CreateInsertTagStatement(tags []string, slugs []string) string {
	var buffer bytes.Buffer
	
	buffer.WriteString(	"INSERT INTO Tag (Name, Slug) VALUES ")
	
	for i := 0; i < len(slugs); i++ {
		
		if i > 0 {
			buffer.WriteString(" , ")
		}
		
		buffer.WriteString(	`("`)
		buffer.WriteString(tags[i])
		buffer.WriteString(`","`)
		buffer.WriteString(slugs[i])
		buffer.WriteString(	`")`)		
		
	}
	
	buffer.WriteString("ON DUPLICATE KEY UPDATE count=count+1; ")
	
	return buffer.String()
}

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
	
	
	var err error	
	var tags []models.Tag
	var linkQuoteCategory models.QuoteCategory
	
	
	// Set date time - for new quote
	quote.Quote.Created = time.Now().UnixNano()
	quote.Quote.Updated = quote.Quote.Created
	
	// Insert quote
	err = c.Txn.Insert(&quote.Quote)
	if err != nil {
		panic(err)
	}
	
	// create slugs from tags
	slugs := c.Slug(quote.Tags)
	
	// If we have any slugs
	if (len(slugs) > 0 ) {
		
		insertTagsString := c.CreateInsertTagStatement(quote.Tags, slugs)
		
		// insert tags for them where they dont exist
		_, err = c.Txn.Exec(insertTagsString)
		
		if err != nil {
			panic(err)
		}
		
		// we need the new ids - lets get tags bag
		getTagsString := c.CreateGetTagStatement(slugs)
		
		_, err = c.Txn.Select(&tags, getTagsString)
	
		if err != nil {
			panic(err)
		}
	}
	
	// for all the tags that need association
	for i := 0; i < len(tags); i++ {
		 
		fmt.Println(tags[i].Name)
		
	}
	
	// Link to a category - Soon this will be multiple
	
	if (quote.CategoryId.Valid) {
	
		// Get the category and add a count
	
		linkQuoteCategory.QuoteId = quote.Quote.Id
		linkQuoteCategory.CategoryId = int(quote.CategoryId.Int64)
		
		err = c.Txn.Insert(&linkQuoteCategory)
		
		if err != nil {
			panic(err)
		}
	}
	
	
	c.Flash.Success("Added Quote")
	
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