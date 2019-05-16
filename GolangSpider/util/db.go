package util

import (
	"github.com/astaxie/beego/logs"
	"github.com/jinzhu/gorm"
)

func ConnectDB(dialect,dbInfo string,openLog bool) *gorm.DB{
	// sslmode=disable需要配置这个选项，不然会报错
	//db open error. pq: SSL is not enabled on the server
	db, err := gorm.Open(dialect, dbInfo)
	if err != nil {
		logs.Error("db open error.", err.Error())
		panic(err.Error())
	}
	//defer db.Close()
	//打印日志
	db.LogMode(openLog)
	//db.AutoMigrate(&Author{})
	//db.AutoMigrate(&Poem{})
	return db
}

func InitTables(tables []interface{},db *gorm.DB,isSingle bool)  {
	if isSingle{
		db.SingularTable(true)
	}
	//初始化表结构
	for _,table:=range tables{
		db.AutoMigrate(table)
	}
}
