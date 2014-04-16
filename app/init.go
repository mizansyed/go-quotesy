package app

import (
	"bytes"
	"database/sql"
	"fmt"
	"github.com/revel/revel"
	_ "html"
	"html/template"
	"math"
	"reflect"
	"strconv"
)

var NullInt64Binder = revel.Binder{
	Bind: func(params *revel.Params, name string, typ reflect.Type) reflect.Value {
		vals, ok := params.Values[name]
		if !ok || len(vals) == 0 {
			return reflect.Zero(typ)
		}

		intValue, err := strconv.ParseInt(vals[0], 10, 64)
		if err != nil {
			fmt.Println(err)
			return reflect.Zero(typ)
		}
		pValue := reflect.New(typ)

		// Set the NullInt64 type
		pValue.Elem().Set(reflect.ValueOf(sql.NullInt64{intValue, true}))
		return pValue.Elem()
	},
	Unbind: func(output map[string]string, key string, val interface{}) {
		output[key] = fmt.Sprintf("%d", val)
	},
}

func init() {
	revel.TypeBinders[reflect.TypeOf(sql.NullInt64{})] = NullInt64Binder
	// Filters is the default set of global filters.
	revel.Filters = []revel.Filter{
		revel.PanicFilter,             // Recover from panics and display an error page instead.
		revel.RouterFilter,            // Use the routing table to select the right Action
		revel.FilterConfiguringFilter, // A hook for adding or removing per-Action filters.
		revel.ParamsFilter,            // Parse parameters into Controller.Params.
		revel.SessionFilter,           // Restore and write the session cookie.
		revel.FlashFilter,             // Restore and write the flash cookie.
		revel.ValidationFilter,        // Restore kept validation errors and save new ones from cookie.
		revel.I18nFilter,              // Resolve the requested language
		HeaderFilter,                  // Add some security based headers
		revel.InterceptorFilter,       // Run interceptors around the action.
		revel.CompressFilter,          // Compress the result.
		revel.ActionInvoker,           // Invoke the action.
	}

	/*revel.TemplateFuncs["field_val"] = func (f *revel.Field) template.HTML {

		var str string

		if f.FlashExists() {
			str = f.Flash()
		} else {
			str = fmt.Sprintf("%v", f.Value())
		}

		return template.HTML(str)
	}*/

	revel.TemplateFuncs["display_error"] = func(f *revel.Field) template.HTML {

		var str string

		if f.Error != nil {
			str = fmt.Sprintf(`<div class="error-container"><span class="error">%s</span></div>`, f.Error)
		}

		return template.HTML(str)
	}

	revel.TemplateFuncs["pager"] = func(val string, recordCount int, recordsPerPage int, pageIndex int) template.HTML {

		var buffer bytes.Buffer

		pages := int(math.Ceil(float64(recordCount) / float64(recordsPerPage)))

		buffer.WriteString("<ul class='pagination'>")

		if pageIndex > 1 {
			buffer.WriteString(fmt.Sprintf("<li><a href='/%s/%d'>&laquo;</a></li>", val, pageIndex-1))
		}

		for i := 1; i <= pages; i++ {
			buffer.WriteString(fmt.Sprintf("<li><a href='/%s/%d'>%d</a></li>", val, i, i))
		}

		if pageIndex < pages {
			buffer.WriteString(fmt.Sprintf("<li><a href='/%s/%d'>&raquo;</a></li>", val, pageIndex+1))
		}
		buffer.WriteString("</ul>")

		return template.HTML(buffer.String())
	}

	// register startup functions with OnAppStart
	// ( order dependent )
	// revel.OnAppStart(InitDB())
	// revel.OnAppStart(FillCache())
}

// TODO turn this into revel.HeaderFilter
// should probably also have a filter for CSRF
// not sure if it can go in the same filter or not
var HeaderFilter = func(c *revel.Controller, fc []revel.Filter) {
	// Add some common security headers
	c.Response.Out.Header().Add("X-Frame-Options", "SAMEORIGIN")
	c.Response.Out.Header().Add("X-XSS-Protection", "1; mode=block")
	c.Response.Out.Header().Add("X-Content-Type-Options", "nosniff")

	fc[0](c, fc[1:]) // Execute the next filter stage.
}
