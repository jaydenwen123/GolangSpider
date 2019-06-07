package leetcode


//获取推荐的数据
func GetRecommanderJson() {
	SaveJsonDataByUrl(recommanderFavoriateUrl, favoriatesJsonfile)
	SaveJsonDataByUrl(recommanderHotUrl, hot100Jsonfile)
	SaveJsonDataByUrl(recommanderTencentUrl, tencent50Jsonfile)
	SaveJsonDataByUrl(recommanderTopUrl, topJsonfile)
}

