package english

import (
	"github.com/astaxie/beego/logs"
	"github.com/jaydenwen123/go-util"
	"io/ioutil"
	"log"
	"testing"
)

func TestParseCategory(t *testing.T) {
	//tests := []struct {
	//	name    string
	//	want    *Category
	//	wantErr bool
	//}{
	//	// TODO: Add test cases.
	//	{name:"test1",
	//		want:nil,
	//	wantErr:false},
	//}
	//for _, tt := range tests {
	//	t.Run(tt.name, func(t *testing.T) {
	//		got, err := ParseCategory()
	//		if (err != nil) != tt.wantErr {
	//			t.Errorf("ParseCategory() error = %v, wantErr %v", err, tt.wantErr)
	//			return
	//		}
	//		if !reflect.DeepEqual(got, tt.want) {
	//			t.Errorf("ParseCategory() = %v, want %v", got, tt.want)
	//		}
	//	})
	//}
	categories, err := ParseCategory()
	if err != nil {
		t.Errorf("%s","there is error!")
	}
	//t.Logf("\n%#v\n",categories)
	//log.Printf("%#v",categories)
	for _,category:=range categories{
		ParseChannelInfo(category)
		log.Printf("%s\n",category.Name)
		for _,channel:=range category.Channels{
			log.Println(channel.Name," ",channel.Link)
			ParseArticleInfo(channel)
			for _,article:=range channel.Articles{
				log.Println(article.Name," ",article.Link)
				for _,articleDetail:=range article.ArticleDetails{
					ParseArticleContent(articleDetail)
					log.Println(articleDetail.Name," ",articleDetail.Link," ",articleDetail.Times," ",articleDetail.Date," ")
					content:=articleDetail.Content
					log.Printf("%#v",content)
					if content.IsVideo{
						if err := util.Download(content.File, content.Name+".mp4");err!=nil{
							logs.Error("下载视频文件失败")
						}
					}else{
						if err := util.Download(content.File, content.Name+".mp3");err!=nil{
							logs.Error("下载音频文件失败")
						}
					}
					if err := ioutil.WriteFile(content.Name+".txt", []byte(content.Text), 0755);err!=nil{
						logs.Error("下载文本文件失败")
					}
				}
				break
			}
			break
		}
		break
	}

	util.Save2JsonFile(categories,"category.json")
}

