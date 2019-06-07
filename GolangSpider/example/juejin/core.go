package juejin

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/jaydenwen123/go-util"
	"github.com/pkg/errors"
	"github.com/tidwall/gjson"
	"os"
	"time"
)

func init() {

}
func Init()  {
	util.InitDir(MARKDOWN_BASE_DIR)
	util.InitDir(MARKDOWN_HOT_DIR)
	util.InitDir(MARKDOWN_NEW_DIR)
	util.InitDir(MARKDOWN_PUPULAR_DIR)
	util.InitDir(MARKDOWN_SEARCH_DIR)
}

//1.爬取文章列表
func GetArticle(param *PostQueryBody) (string ,error){
	jsonStr := util.RequestJsonWithPost(juejinUrl, headers, param.String())

	//fmt.Println(jsonStr)
	if gjson.Valid(jsonStr){
		return jsonStr,nil
	}
	return "",errors.New("get article error.")
}
//2.解析文章Id
func ParseArticleInfo(articles string,articlePath,pagePath string) ([]*Article,*PageInfo) {
	parse := gjson.Parse(articles)
	articleNodes := parse.Get(articlePath)
	pageInfoNode:=parse.Get(pagePath)
	if articleNodes.IsArray(){
		Articles:=make([]*Article,0)
		pageInfo:=&PageInfo{
			HasNextPage:pageInfoNode.Get("hasNextPage").Bool(),
			EndCursor:pageInfoNode.Get("endCursor").String(),}
		for _,article:=range articleNodes.Array(){
			id:=article.Get("id").String()
			title:=article.Get("title").String()
			originalUrl:=article.Get("originalUrl").String()
			commentsCount:=article.Get("commentsCount").Int()
			likeCount:=article.Get("likeCount").Int()
			Articles=append(Articles,&Article{
				Title:title,
				Id:id,
				OriginalUrl:originalUrl,
				CommentCount:int(commentsCount),
				LikeCount:int(likeCount),
			})
		}
		return Articles,pageInfo
	}
	return nil,nil
}

//3.爬取文章详情信息
func SpiderArticleDetail(articleUrl string) string{
	_, htmlStr:= util.RequestWithHeader(articleUrl, headers)
	return htmlStr
}
//4.解析文章html内容
func ParseArticleDetail(htmlStr string)  (string,error){
	articleDetail := util.MatchStringValue(articleDetailRe, htmlStr)
	if articleDetail==""{
		return "",errors.New("parse the article detail error.")
	}
	return articleDetail,nil
}
//5.将文章内容保存为markdown文档存储
func SaveArticleAsMarkdown(articleHtmlStr string,markdownFile string) error {
	err := util.MarkdownFromHtmlString(articleHtmlStr, markdownFile)
	return err
}

func GetPageArticles(currentPage int,param *PostQueryBody,dirName string,articlePath,pagePath string)  {
	//第一步
	articles, err := GetArticle(param)
	if err!=nil{
		logs.Error(err.Error())
	}
	//第二步
	articleInfos,pageInfo := ParseArticleInfo(articles,articlePath,pagePath)
	if articleInfos!=nil{
		fmt.Println("正在处理第",currentPage,"页的数据")
		SavePageArticlesAsMarkdown(dirName, currentPage, articleInfos)
		//递归的获取数据,获取前十页的数据
		//&& currentPage<MAX_PAGE
		if pageInfo.HasNextPage {
			//更新post请求参数，然后递归获取数据
			param.Variables.After = pageInfo.EndCursor
			//递归获取数据
			GetPageArticles(currentPage+1, param, dirName, articlePath, pagePath)
		}
	} else {
		logs.Error("解析文章出错。。。")
	}

}

func SavePageArticlesAsMarkdown(dirName string, currentPage int, articleInfos []*Article) {
	filename := dirName + string(os.PathSeparator) + "article" + fmt.Sprintf("%d", currentPage) + ".json"
	util.Save2JsonFile(articleInfos, filename)
	//第三步
	for _, articleInfo := range articleInfos {
		htmlStr := SpiderArticleDetail(articleInfo.OriginalUrl)
		//第四步
		articleDetail, err := ParseArticleDetail(htmlStr)
		if err != nil {
			logs.Error(err.Error())
			continue
		}
		//第五步
		filename = dirName + string(os.PathSeparator) + util.TrimIllegalStr(articleInfo.Title, "") + ".md"
		err = SaveArticleAsMarkdown(articleDetail, filename)
		if err != nil {
			logs.Error("保存", filename, "失败错误信息：", err.Error())
		} else {
			logs.Info("保存", filename, "成功")
		}
		time.Sleep(time.Millisecond * 100)
	}
}

