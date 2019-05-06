package cbooo

import (
	"GolangSpider/common"
	"GolangSpider/util"
	"encoding/csv"
	"fmt"
	"github.com/astaxie/beego/logs"
	"os"
	"regexp"
	"strconv"
	"strings"
)

//票房信息
type BoxOffice struct {
	Title       string
	TodayBigBar string
	MovieInfos  []*MovieInfo
}

//影片信息
type MovieInfo struct {
	//影片排名
	RankNo int
	//影片名称
	MovieName string
	//实时票房(万)
	RealSale float64
	//票房占比(%)
	SaleRatio string
	//累计票房(万)
	AccumulateSale float64
	//排片占比(%)
	MovieRatio string
	//上映天数
	PublishDays int
}

func (this *MovieInfo) GetCsvHeader() []string {
	csvHeader := []string{
		"影片排名",
		"影片名称",
		"实时票房（万）",
		"票房占比",
		"累计票房（万）",
		"排片占比",
		"上映天数"}
	return csvHeader
}

func (this *MovieInfo) ValueToStrArray() []string {
	strArr := []string{}
	strArr = append(append(append(append(append(append(append(strArr,
		fmt.Sprintf("%d", this.RankNo)),
		this.MovieName),
		fmt.Sprintf("%.1f", this.RealSale)),
		this.SaleRatio),
		fmt.Sprintf("%.1f", this.AccumulateSale)),
		this.MovieRatio),
		fmt.Sprintf("%d", this.PublishDays))
	return strArr

}

const (
	JSON_BOXOFFICE = "boxoffice.json"
	CSV_BOXOFFICE  = "boxoffice.csv"
	XML_BOXOFFICE  = "boxoffice.xml"
)

//正则表达式变量
var (
	titleRe           = `<h1 onclick="GotoUrl\('/realtime'\)" class="trtop">(.*?)</h1>`
	todayBigBarRe     = `<h6><span id="week"></span>(.*?)</h6>`
	rankNo            = `<td.*?style='width:60px'>(.*?)</td>`
	movieNameRe       = `<td style='width:150px'>(.*?)</td>`
	realSaleRe        = `<td style='width:120px'>(.*?)</td>`
	saleRatioRe       = `<td style='width:80px'>(.*?)</td>`
	accumulateRatioRe = `<td style='width:120px'>(.*?)</td>`
	movieRatioRe      = `<td style='width:80px'>(.*?)</td>`
	publishDaysRe     = `<td style='width:70px;'>(.*?)</td>`
)

//中国票房
func Main() {
	url := "http://www.cbooo.cn"
	//1.发送http请求
	_, content := common.Request(url)
	//fmt.Println(string(content))
	//2.正则表达式解析
	boxOffice := ParseRealTimeWithRegexp(content)
	//for _, movieInfo := range movieInfos {
	//	fmt.Printf("%#v", movieInfo)
	//	fmt.Println()
	//}
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
	for index, mi := range bo.MovieInfos {
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

	//构建MovieInfos

	// fmt.Println("--------------rankNo------------")
	submatchs = util.MatchTarget(rankNo, content)
	movieInfos := make([]*MovieInfo, len(submatchs))
	for index, item := range submatchs {
		movieInfos[index] = &MovieInfo{}
		movieInfos[index].RankNo, _ = strconv.Atoi(item[1])
	}
	//fmt.Println("--------------movieNameRe------------")
	submatchs = util.MatchTarget(movieNameRe, content)
	for index, item := range submatchs {
		movieInfos[index].MovieName = item[1]
	}
	//fmt.Println("--------------RealSaleRe------------")
	submatchs = util.MatchTarget(realSaleRe, content)
	for index, item := range submatchs {
		if index%2 == 0 {
			movieInfos[index/2].RealSale, _ = strconv.ParseFloat(item[1], 64)
		}
	}
	//fmt.Println("---------------SaleRatioRe-----------")
	submatchs = util.MatchTarget(saleRatioRe, content)
	for index, item := range submatchs {
		if index%2 == 0 {
			movieInfos[index/2].SaleRatio = item[1]
		}
	}
	//fmt.Println("---------------AccumulateRatioRe-----------")
	submatchs = util.MatchTarget(accumulateRatioRe, content)
	for index, item := range submatchs {
		if index%2 == 1 {
			movieInfos[index/2].AccumulateSale, _ = strconv.ParseFloat(item[1], 64)
		}
	}
	//fmt.Println("---------------MovieRatioRe-----------")
	submatchs = util.MatchTarget(movieRatioRe, content)
	for index, item := range submatchs {
		if index%2 == 1 {
			movieInfos[index/2].MovieRatio = item[1]
		}
	}
	//fmt.Println("---------------PublishDaysRe-----------")
	submatchs = util.MatchTarget(publishDaysRe, content)
	for index, item := range submatchs {
		movieInfos[index].PublishDays, _ = strconv.Atoi(item[1])
	}
	boxOffice.MovieInfos = movieInfos
	return boxOffice
}
