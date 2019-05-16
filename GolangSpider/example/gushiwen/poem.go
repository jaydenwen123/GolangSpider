package gushiwen

import (
	"GolangSpider/example/gushiwen/db"
	"GolangSpider/util"
	"fmt"
	"github.com/astaxie/beego/logs"
	"io/ioutil"
)

func Main()  {
	//古诗文网：https://www.gushiwen.org/
	//爬取诗文思路
	//首先根据诗文类型，找到所属类型的诗文，提取诗文ID，然后在爬取每篇诗文的信息
	//采用go协程并发的爬取

	types, _, _, _ := SpiderShiwenKindUrl()
	//1.初始化
	db.InitDB(dbDialect, dbInfo)
	//saveFile(types, authors, dynastys, styles)
	//2.根据诗文类型爬取所有的诗文
	SpiderAllShiwensByTypes(types)

	//将 example/gushiwen/data/poems 目录下的所有文件读取进来，然后写入到数据库中

	//SaveDirFilesToDB()

}

func SaveDirFilesToDB() {
	files, err := ioutil.ReadDir(DUMP_POEM_BASE_PATH)
	if err != nil {
		logs.Error("read Dir error.", err.Error())
		panic(err.Error())
	}
	for _, file := range files {
		//DUMP_POEM_BASE_PATH+"一七令·山.json"
		//fmt.Println(DUMP_POEM_BASE_PATH+file.Name())
		FromJsonFileToDB(DUMP_POEM_BASE_PATH + file.Name())
	}
}

func FromJsonFileToDB(filePath string) {
	poem := &db.Poem{Author: &db.Author{}}
	util.LoadObjectFromJsonFile(filePath, poem)
	//2.存储诗词信息到数据库中
	DumpPoemToDatabase(poem, "")
}
//并发的爬取所有类型的诗文
func SpiderAllShiwensByTypes(types map[string]string) {
	typesChan := make([]chan bool, len(types))
	index := 0
	for _, item := range types {
		typesChan[index] = make(chan bool)
		go HandlePoemType(item, typesChan[index])
		index++
	}
	for i := 0; i < len(types); i++ {
		select {
		case <-typesChan[i]:
			fmt.Println("handle type poems finish")
		}
	}
}
//爬取一种类型的诗文
func HandlePoemType(item string,doneFlag chan bool) {
//https://so.gushiwen.org/shiwenv_6105b29267b5.aspx
	urls := SpiderShiwenUrlByType(item)
	doFlag :=make([]chan bool,len(urls))
	for index, url := range urls {
		doFlag[index]=make(chan bool)
		go SavePoemToFile(url,doFlag[index])
	}
	for i:=0;i<len(urls);i++{
		select {
		case <-doFlag[i]:
			fmt.Println("parse poem finish")
		}
	}
	doneFlag<-true
}
//保存一首诗到文件中
func SavePoemToFile(url string,doFlag chan bool) {
	poem := SpiderShiwenContent(url)
	//保存到文件
	util.Save2JsonFile(poem, DUMP_POEM_BASE_PATH+poem.Title+".json")
	//保存到数据库
	db.DB.Save(poem)
	doFlag<-true
}

func saveFile(types map[string]string, authors map[string]string, dynastys map[string]string, styles map[string]string) {
	util.Save2JsonFile(types, JSON_SHIWEN_TYPE)
	util.Save2JsonFile(authors, JSON_SHIWEN_AUTHOR)
	util.Save2JsonFile(dynastys, JSON_SHIWEN_DYNASTY)
	util.Save2JsonFile(styles, JSON_SHIWEN_STYLE)
}
