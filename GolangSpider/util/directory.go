package util

import (
	"github.com/astaxie/beego/logs"
	"os"
)

//初始化目录,目录不存在则创建，如果存在则直接跳过
func InitDir(path string) error{
	if _, err := os.Stat(path);err!=nil && os.IsNotExist(err){
		err = os.MkdirAll(path, os.ModeDir)
		if err!=nil{
			logs.Error("init directory <",path,">error.",err.Error())
			return err
		}
		logs.Info("init directory<",path,"> success.")
		return err
	}
	return nil
}