package cbooo

import (
	"github.com/jaydenwen123/go-util"
	"strings"
)

//根据影片的id得到影片海报的url
func getPosterUrl(id string) string {
	return strings.Replace(movieImgUrlTemplate, urlTmpStr, id, -1)
}

//根据影片Id得到影片的详细信息的url
func getDetailUrl(id string) string {
	return strings.Replace(movieDetailUrlTemplate, urlTmpStr, id, -1)
}

//提取电影信息
func ParseMovieDetail(content string) *MovieInfo{

	movie:=&MovieInfo{MovieDetail:&MovieDetail{}}
	//1.获取电影海报
	setMoviePoster(content,movie)
	//2.获取电影详细信息
	setMovieDetail(content,movie)
	//3.获取主演和导演信息
	setMoviePerson(content,movie)
	return movie
}
//解析电影的导演和主演信息
func setMoviePerson(content string, movie *MovieInfo) {
	moviePersonBlock := util.MatchTarget(moviePersonBlockRe, content)
	for index,item:=range moviePersonBlock{
		//moviePersonBlock = util.TrimSpace(item)
		//fmt.Println(moviePersonBlock)
		moviePersons := util.TrimSpace(item[1])
		targets := util.MatchTarget(moviePersonDetailRe, moviePersons)
		if index==0{
			//导演：
			directors:=[]MoviePerson{}
			for _,target:=range targets{
				//fmt.Println(target[1:])
				directors=append(directors,MoviePerson{PersonName:strings.Replace(target[2],"&#183;",".",-1),
					LinkUrl:target[1]} )
			}
			movie.MovieDetail.Director=directors
		}else if index==1{
			//主演：
			actors:=[]MoviePerson{}
			for _,target:=range targets{
				//fmt.Println(target[1:])
				actors=append(actors,MoviePerson{PersonName:strings.Replace(target[2],"&#183;",".",-1),
					LinkUrl:target[1]} )
			}
			movie.MovieDetail.Actor=actors
		}else{
			break
		}
	}
}
//解析影片的海报图片url
func setMovieDetail(content string,movie *MovieInfo) {
	movieDetailBlock := util.MatchStringValue(movieDetailBlockRe, content)
	//去除空格
	movieDetailBlock = util.TrimSpace(movieDetailBlock)
	//fmt.Println(movieDetailBlock)
	//开始匹配
	movieDet:=util.MatchTarget(movieDetailRe,movieDetailBlock)[0][1:]
	arrLen:=len(movieDet)
	//fmt.Println(movieDet)
	//movieDet=movieDet[0][1:]
	//[复仇者联盟4：终局之战
	// 2019
	// Avengers:Endgame
	// 2247.5万
	// 393406.5万
	// 科幻/动作/冒险
	// 181min
	// 2019-4-24（中国 ）
	// 3D/IMAX
	// 美国
	// http://www.cbooo.cn/c/6
	// 中国电影集团公司]


	movie.MovieName=movieDet[0]
	movie.NickName=movieDet[2]
	movie.MovieType=movieDet[arrLen-7]
	if arrLen>11{
		movie.MovieDetail.TodayRealSale=movieDet[arrLen-9]
	}
	movie.MovieDetail.AccumulateSale=movieDet[arrLen-8]
	movie.MovieDetail.Duration=movieDet[arrLen-6]
	movie.MovieDetail.MovieDate=movieDet[arrLen-5]
	movie.MovieDetail.MovieFormat=movieDet[arrLen-4]
	movie.MovieDetail.MovieCountry=movieDet[arrLen-3]
	movie.MovieDetail.MovieYear=movieDet[1]
	//公司（url）
	movie.MovieDetail.MovieCompany=movieDet[arrLen-1]+"("+movieDet[arrLen-2]+")"
}

//解析影片的详细信息如导演、主演、片长等信息
func setMoviePoster(content string,movie *MovieInfo)  {
	movie.MoviePoster=util.MatchStringValue(moviePosterRe,content)
}

