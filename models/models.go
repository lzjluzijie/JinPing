package models

import (
	"log"

	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"

	//mysql
	_ "github.com/go-sql-driver/mysql"

	//sqlite
	_ "github.com/mattn/go-sqlite3"
)

var x *xorm.Engine

func init() {
	//engine, err := xorm.NewEngine("mysql", "root:password@/message?charset=utf8")
	engine, err := xorm.NewEngine("sqlite3", "./jinping.db")
	if err != nil {
		panic(err)
		return
	}

	x = engine

	x.SetLogger(nil)
	x.SetMapper(core.GonicMapper{})

	err = x.Sync2()
	if err != nil {
		log.Fatal(err.Error())
		return
	}
}
