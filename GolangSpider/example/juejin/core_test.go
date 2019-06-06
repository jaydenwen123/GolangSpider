package juejin

import (
	"testing"
)

func TestGetPopularArticle(t *testing.T) {
	//初始化目录
	Init()
	GetPageArticles(1,popularParam,MARKDOWN_PUPULAR_DIR,recommandArticleDetailPath,recommandPageInfoPath)
	//GetPageArticles(1,hottestParam,MARKDOWN_HOT_DIR,recommandArticleDetailPath,recommandPageInfoPath)
	//GetPageArticles(1,newestParam,MARKDOWN_NEW_DIR,recommandArticleDetailPath,recommandPageInfoPath)
	//keyword="java"
	//GetPageArticles(1,searchThreeMonthArticlesParam,MARKDOWN_SEARCH_DIR,searchArticleDetailPath,searchArticlePageInfoPath)

}
