package other

import (
	"GolangSpider/GolangSpider/util"
	"github.com/astaxie/beego/logs"
	"testing"
)

func TestHtml2Markdown(t *testing.T) {
	//file, err := os.Open("test.html")
	//if err!=nil{
	//	logs.Error("open file err:",err.Error())
	//}
	//inputHTML,err:=ioutil.ReadFile("example/juejin/test2.html")
	//if err!=nil{
	//	logs.Error("read file error.",err.Error())
	//}
	//text, err := html2text.FromString(string(inputHTML), html2text.Options{PrettyTables: true})
	//if err != nil {
	//	logs.Error("xxxxx")
	//	panic(err)
	//}
	//fmt.Println(text)
	// ioutil.WriteFile("example/juejin/test2.md", []byte(text), 0755)
	err := util.MarkdownFromHtmlFile("test.html", "test.md")
	//example/juejin/test.html
	if err != nil {
		logs.Error("转换markdown文档失败")
		panic(err)
	}else{
		logs.Info("转换markdown文档成功")
	}
}
