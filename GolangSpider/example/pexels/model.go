package pexels

import (
	"github.com/astaxie/beego/logs"
	"strconv"
	"strings"
)

//Type
type Type struct {
	Height	int`json:"height"`
	Width 	int`json:"width"`
}



//图片的尺寸
func NewType(hw string) *Type {
	arr := strings.Split(hw, "x")
	if len(arr)==2{
		height,err:=strconv.Atoi(arr[0])
		if err!=nil{
			logs.Error("解析图片的高度错误.使用默认的高度下载图片",err.Error())
			height=1280
		}
		width,err:=strconv.Atoi(arr[1])
		if err!=nil{
			logs.Error("解析图片的宽度错误.使用默认的宽度下载图片",err.Error())
			width=960
		}
		return &Type{Height: height, Width: width}
	}
	return nil
}


//图片信息
type Picture struct {
	//图片尺寸
	Types map[string]*Type
	//图片名称
	Names map[string]string
	//链接
	DownloadUrls map[string]string
	//图片Id
	Id string
	//图片链接
}

func NewPicture(id string,types map[string]string,baseUrl string)  *Picture{
	picture := &Picture{}
	picture.Id=id
	picture.Types=make(map[string]*Type,0)
	picture.Names=make(map[string]string,0)
	picture.DownloadUrls=make(map[string]string,0)
	for key,typ:=range types{
		Type:=NewType(typ)
		picture.Types[key]=Type
		picture.Names[key]=generateName(id,typ)
		picture.DownloadUrls[key]=generateUrl(baseUrl,Type)
	}
	return picture
}
//生成下载的url
func generateUrl(url string, typ *Type) string {
	downloadUrl:=strings.Replace(pictureUrl,"{url}",url,-1)
	downloadUrl=strings.Replace(downloadUrl,"{height}",strconv.Itoa(typ.Height),-1)
	downloadUrl=strings.Replace(downloadUrl,"{width}",strconv.Itoa(typ.Width),-1)
	return downloadUrl
}
//12312[245-123].jpg
func generateName(id string, Type string) string {
	return id+"["+strings.Replace(Type,"x","-",-1)+"].jpg"
}

//页数据
type PageContent struct {
	Index int
	Data string
}