package util

import (
	"io/ioutil"
	"jaytaylor.com/html2text"
	"os"
)
//根据html文件生成markdown文件
//htmlFile:html文件名
//markdownFile:markdown文件名
func MarkdownFromHtmlFile(htmlFile string,markdownFile string) error{
	hFile,err:=os.Open(htmlFile)
	if err != nil {
		//panic(err)
		return err
	}
	markText, err := html2text.FromReader(hFile,html2text.Options{PrettyTables: true})
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(markdownFile, []byte(markText), os.ModeType)
	if err != nil {
		return err
	}
	return nil
}

//根据html字符串生成markdown文件
//htmlStr：html字符串
//markdownFile：markdown文件名
func MarkdownFromHtmlString(htmlStr string,markdownFile string) error {
	markText, err := html2text.FromString(htmlStr, html2text.Options{PrettyTables: true})
	if err != nil {
		return err
	}
	err= ioutil.WriteFile(markdownFile, []byte(markText), os.ModeType)
	if err!=nil{
		return err
	}
	return nil
}

//根据html字节内容生成markdown文件
//htmlData：html字节内容
//markdownFile：markdown文件名
func MarkdownFromHtmlBytes(htmlData []byte,markdownFile string) error {
	return MarkdownFromHtmlString(string(htmlData),markdownFile)
}
