package models

import (
	"catdogs.club/back-end/configs/common"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"log"
	"os"
)

var engine *xorm.Engine

func InitModel() {
	engine, err := xorm.NewEngine("mysql", configs.GetDbAddr())
	f, err := os.Create(configs.GetSqlLogFile())
	if err != nil {
		log.Fatal(err)
	}
	logger := xorm.NewSimpleLogger(f)
	logger.ShowSQL(true)
	engine.SetLogger(logger)
	engine.SetMaxIdleConns(10)
	engine.SetMaxOpenConns(100)
	engine.Sync2(new(User))
}
