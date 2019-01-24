package models

import (
	"time"

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
	engine, err := xorm.NewEngine("sqlite3", "./jp.db")
	if err != nil {
		panic(err)
	}
	x = engine

	x.TZLocation, err = time.LoadLocation("Asia/Shanghai")
	if err != nil {
		panic(err)
	}

	x.SetLogger(nil)
	x.SetMapper(core.GonicMapper{})

	err = x.Sync2(new(Comment), new(Topic), new(User))
	if err != nil {
		panic(err)
	}
}
