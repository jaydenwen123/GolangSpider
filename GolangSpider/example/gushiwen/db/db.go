package db

import (
	"GolangSpider/util"
	"github.com/jinzhu/gorm"
)

var(
	DB=&gorm.DB{}
	tables=[]interface{}{&Author{},
		&Poem{}}
)

func InitDB(dbDialect,dbInfo string)  {
	//logs.Info("to connect postgreSQL")

	//1.连接
	DB=util.ConnectDB(dbDialect,dbInfo,true)
	//2.初始化数据库
	initTable(tables,DB,true)
}

func initTable(tables []interface{},db *gorm.DB,isSingle bool)  {
	//2.初始化表
	util.InitTables(tables,db,isSingle)
}