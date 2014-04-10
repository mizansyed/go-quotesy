package controllers

import (
	"fmt"
	"database/sql"
	"github.com/coopernurse/gorp"
	 _ "github.com/mattn/go-sqlite3"
	"github.com/revel/revel"
	"github.com/revel/revel/modules/db/app"
	"github.com/muizsyed/go-quotesy/app/models"
)

var (
	Dbm *gorp.DbMap
)

func InitDB() {
	db.Init()
	Dbm = &gorp.DbMap{Db: db.Db, Dialect: gorp.SqliteDialect{}}
	//Dbm.TraceOn("", log.New(os.Stdout, "gorptest: ", log.Lmicroseconds))
	Dbm.AddTableWithName(models.Language{}, "language").SetKeys(true, "Id")
	
	fmt.Println("Create tables")
	
}

type GorpController struct {
	*revel.Controller
	Txn *gorp.Transaction
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
