package controllers

import "github.com/revel/revel"
import "github.com/muizsyed/go-quotesy/app/models"

type App struct {
		GorpController
}

func (c App) Index() revel.Result {
	
	
	quotes, err := c.Txn.Select(models.Quote{}, "select * from Quote order by created desc limit 5")
	
	if err != nil {
		panic(err);
	}
	
	return c.Render(quotes)
}
