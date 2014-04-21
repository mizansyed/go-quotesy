package controllers

import "github.com/revel/revel"
import "github.com/muizsyed/go-quotesy/app/models"
import _"github.com/muizsyed/go-quotesy/app/routes"
import _"fmt"
import _"strconv"
import _"time"

type Tag struct {
	GorpController
}


func (c Tag) Filter(name string) revel.Result {	
	
	tags, err := c.Txn.Select(models.Json{}, "select tag_id value, name data from Tag where name like ?","%"+name+"%")
	
	if err != nil {
		panic(err)
	}
	
	return c.RenderJson(tags)
}