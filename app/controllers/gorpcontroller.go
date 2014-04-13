package controllers

import (
	"fmt"
	"database/sql"
	"github.com/coopernurse/gorp"
	 _ "github.com/mattn/go-sqlite3"
	"github.com/revel/revel"
	"github.com/revel/revel/modules/db/app"
	"github.com/muizsyed/go-quotesy/app/models"
	_ "runtime"
	_"io"
	_"net/http"
	_"os"
	_"path/filepath"
	_"reflect"
	_"runtime"
	_"strings"
	_"time"
	"bytes"
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

