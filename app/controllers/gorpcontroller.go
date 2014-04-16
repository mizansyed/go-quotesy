package controllers

import (
	"bytes"
	"database/sql"
	"fmt"
	"github.com/coopernurse/gorp"
	_ "github.com/mattn/go-sqlite3"
	"github.com/muizsyed/go-quotesy/app/models"
	"github.com/revel/revel"
	"github.com/revel/revel/modules/db/app"
	// "github.com/ziutek/mymysql"
	_ "io"
	"log"
	_ "net/http"
	"os"
	_ "path/filepath"
	_ "reflect"
	_ "runtime"
	_ "strings"
	_ "time"
)

var (
	Dbm *gorp.DbMap
)

func InitDB() {
	db.Init()

	Dbm = &gorp.DbMap{Db: db.Db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}

	// Enable trace - perhaps should make this configurable
	Dbm.TraceOn("[gorp]", log.New(os.Stdout, "Quotesy:", log.Lmicroseconds))

	// Add tables to DB Map
	Dbm.AddTableWithName(models.Language{}, "Language").SetKeys(true, "Id").SetVersionCol("Version")
	Dbm.AddTableWithName(models.Category{}, "Category").SetKeys(true, "Id").SetVersionCol("Version")

	fmt.Println("Create tables")

}

type GorpController struct {
	*revel.Controller
	Txn *gorp.Transaction
}

func (c *GorpController) GetCount(table string) int {

	var buffer bytes.Buffer

	buffer.WriteString("select count(*) from ")
	buffer.WriteString(table)

	fmt.Println(buffer.String())

	count64, err := c.Txn.SelectInt(buffer.String())

	if err != nil {
		panic(err)
	}

	fmt.Println("Converting count to int: languages")
	// Do I have to do this?
	count := int(count64)

	return count
}

func (c *GorpController) Begin() revel.Result {
	txn, err := Dbm.Begin()
	if err != nil {
		panic(err)
	}
	c.Txn = txn
	return nil
}

func (c *GorpController) Commit() revel.Result {
	if c.Txn == nil {
		return nil
	}
	if err := c.Txn.Commit(); err != nil && err != sql.ErrTxDone {
		panic(err)
	}
	c.Txn = nil
	return nil
}

func (c *GorpController) Rollback() revel.Result {
	if c.Txn == nil {
		return nil
	}
	if err := c.Txn.Rollback(); err != nil && err != sql.ErrTxDone {
		panic(err)
	}
	c.Txn = nil
	return nil
}
