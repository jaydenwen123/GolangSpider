package cbooo

import (
	"GolangSpider/common"
	"GolangSpider/util"
	"encoding/csv"
	"github.com/astaxie/beego/logs"
	"os"
	"regexp"
	"strconv"
	"strings"
)

//中国票房
func Main() {
	//url := "http://www.cbooo.cn"
	SpiderCboooHome()
	SpiderMovieInfo("670808",true)
	SpiderMovieInfo("681412",true)
	SpiderMovieInfo("682238",true)
}

func SpiderMovieInfo(movieId string,downloadPoster bool) {
	//1.发送请求
	detailUrl:=getDetailUrl(movieId)
	_, content := common.Request(detailUrl)
	//fmt.Println(content)
	//content=TrimSpace(content)
	//fmt.Println(content)
	movie := ParseMovieDetail(content)
	//补充两个影片信息
	movie.MovieId,_=strconv.Atoi(movieId)
	movie.MovieDetail.MovieUrl=detailUrl
	//fmt.Printf("%#v",movie)
	util.Save2FormatJsonFile(movie, BASE_PATH+movieId+".json","\t")
	if downloadPoster {
		if err:=util.Download(movie.MoviePoster,BASE_PATH+movieId+".jpg");err!=nil{
			logs.Error("download movie poster failed")
		}else{
			logs.Debug("download movie poster success")
		}
	}
}

func SpiderCboooHome() {
	//1.发送http请求
	_, content := common.Request(cboooUrl)
	//fmt.Println(content)
	//解析CBO实时票房榜数据
	SpiderBoxOffice(content)
	//解析即将上映的影片数据
	SpiderComing(content)
}

func SpiderComing(content string) {
	coming := ParseComingMovie(content)
	//保存数据到文件
	util.Save2JsonFile(coming, JSON_COMING)
	util.Save2XmlFile1(coming, XML_COMING)
}

//解析即将上映部分的数据
func ParseComingMovie(content string) *Coming {
	coming := &Coming{}
	submatchs := util.MatchTarget(comingTitleRe, content)
	//标题
	coming.Title = submatchs[0][1]
	//fmt.Println(submatchs[0][1])
	//fmt.Println("---------------------------")
	//解析电影数据
	submatchs = util.MatchTarget(comingMovieRe, content)
	comingInfos := make([]*ComingInfo, len(submatchs))
	for index, item := range submatchs {
		//影片Id
		//fmt.Println(item[1])
		comingInfo := ParseSingleComingMovieInfo(item[2])
		//创建即将上映影片对象
		id, _ := strconv.Atoi(item[1])
		comingInfos[index] = &ComingInfo{Movie:
		&MovieInfo{MovieId: id,
			MovieName:   comingInfo[1],
			MovieType:   comingInfo[2],
			MoviePoster: strings.Replace(movieImgUrlTemplate, urlTmpStr, item[1], -1)},
			ComingDay:   comingInfo[0],
			TicketIndex: comingInfo[3]}
		//comingInfos[index].TicketIndex,_=strconv.ParseFloat(comingInfo[3],64)
	}
	coming.ComingInfos = comingInfos
	return coming
}

//解析单独一部影片信息
//影片信息，顺序为[上映日期、影片名称、影片类别、购票指数]
func ParseSingleComingMovieInfo(item string) []string {
	//先正则去除空格
	s := util.TrimSpace(item)
	//fmt.Println(s)
	//影片信息，顺序为[上映日期、影片名称、影片类别、购票指数]
	infos := util.MatchTarget(comingMovieInfoRe, s)
	//fmt.Println(infos[0][1:])
	return infos[0][1:]
}


func SpiderBoxOffice(content string) {
	//2.正则表达式解析
	boxOffice := ParseRealTimeWithRegexp(content)
	//3.保存成json文件
	util.Save2JsonFile(boxOffice, JSON_BOXOFFICE)
	//4.保存成csv文件
	Save2CsvFile(boxOffice, CSV_BOXOFFICE)
	//5.保存成xml文件
	util.Save2XmlFile1(boxOffice, XML_BOXOFFICE)
	//util.Save2XmlFile2(boxOffice,"boxoffice2.xml")
	//测试结构体转换为xml字符串
	//xmlStr := util.Obj2XmlStr(boxOffice)
	//fmt.Println(xmlStr)
	//xmlStr=util.Obj2XmlStrIndent(boxOffice," ","\t")
	//fmt.Println(xmlStr)
}

//将实时票房数据写入到csv文件中
func Save2CsvFile(bo *BoxOffice, filename string) {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_RDWR, 0766)
	if err != nil {
		logs.Error("open file error")
		panic(err.Error())
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	//写入标题部分
	if err = writer.Write([]string{bo.Title}); err != nil {
		logs.Error("write header error:", err.Error())
	}
	if err = writer.Write([]string{bo.TodayBigBar}); err != nil {
		logs.Error("write header error:", err.Error())
	}

	//写入主要的数据
	for index, mi := range bo.BoxOfficeInfos {
		if index == 0 {
			//写入票房的表头
			if err = writer.Write(mi.GetCsvHeader()); err != nil {
				logs.Error("write csv header error:", err.Error())
			}
		}
		//写入主体数据
		if err = writer.Write(mi.ValueToStrArray()); err != nil {
			logs.Error("write csv header error:", err.Error())
		}
	}
	writer.Flush()
	logs.Debug("write csv file success")
}

//解析CBO实时票房排名
func ParseRealTimeWithRegexp(content string) *BoxOffice {

	boxOffice := &BoxOffice{}
	//fmt.Println("------------title--------------")
	cp := regexp.MustCompile(titleRe)
	submatchs := cp.FindAllStringSubmatch(content, -1)
	for _, item := range submatchs {
		for _, data := range item {
			boxOffice.Title = data
		}
	}

	//fmt.Println("---------todayBigBarRe----------")
	cp = regexp.MustCompile(todayBigBarRe)
	submatchs = cp.FindAllStringSubmatch(content, -1)
	//fmt.Println(submatchs)
	if len(submatchs) == 1 && len(submatchs[0]) == 2 {
		boxOffice.TodayBigBar = strings.TrimSpace(submatchs[0][1])
	}

	//构建BoxOfficeInfos

	// fmt.Println("--------------rankNo------------")
	submatchs = util.MatchTarget(rankNo, content)
	BoxOfficeInfos := make([]*BoxOfficeInfo, len(submatchs))
	for index, item := range submatchs {
		BoxOfficeInfos[index] = &BoxOfficeInfo{}
		BoxOfficeInfos[index].RankNo, _ = strconv.Atoi(item[1])
	}

	//fmt.Println("--------------movieIdRe------------")
	submatchs = util.MatchTarget(movieIdRe, content)
	for index, item := range submatchs {
		movieId, _ := strconv.Atoi(item[1])
		BoxOfficeInfos[index].Movie = &MovieInfo{
			MovieId:     movieId,
			MoviePoster: getPosterUrl(item[1])}
	}
	//fmt.Println("--------------movieNameRe------------")
	submatchs = util.MatchTarget(movieNameRe, content)
	for index, item := range submatchs {
		BoxOfficeInfos[index].Movie.MovieName = item[1]
	}
	//fmt.Println("--------------RealSaleRe------------")
	submatchs = util.MatchTarget(realSaleRe, content)
	for index, item := range submatchs {
		if index%2 == 0 {
			BoxOfficeInfos[index/2].RealSale, _ = strconv.ParseFloat(item[1], 64)
		}
	}
	//fmt.Println("---------------SaleRatioRe-----------")
	submatchs = util.MatchTarget(saleRatioRe, content)
	for index, item := range submatchs {
		if index%2 == 0 {
			BoxOfficeInfos[index/2].SaleRatio = item[1]
		}
	}
	//fmt.Println("---------------AccumulateRatioRe-----------")
	submatchs = util.MatchTarget(accumulateRatioRe, content)
	for index, item := range submatchs {
		if index%2 == 1 {
			BoxOfficeInfos[index/2].AccumulateSale, _ = strconv.ParseFloat(item[1], 64)
		}
	}
	//fmt.Println("---------------MovieRatioRe-----------")
	submatchs = util.MatchTarget(movieRatioRe, content)
	for index, item := range submatchs {
		if index%2 == 1 {
			BoxOfficeInfos[index/2].MovieRatio = item[1]
		}
	}
	//fmt.Println("---------------PublishDaysRe-----------")
	submatchs = util.MatchTarget(publishDaysRe, content)
	for index, item := range submatchs {
		BoxOfficeInfos[index].PublishDays, _ = strconv.Atoi(item[1])
	}
	boxOffice.BoxOfficeInfos = BoxOfficeInfos
	return boxOffice
}

