package cbooo

import "fmt"

type Coming struct {
	Title string
	ComingInfos []*ComingInfo

}

//即将上映影片信息
type ComingInfo struct {
	ComingDay string
	Movie *MovieInfo
	TicketIndex	string
}
//电影信息
type MovieInfo struct {
	MovieId		int
	MovieName	string
	NickName	string
	MovieType	string
	MoviePoster	string
	MovieDetail	*MovieDetail
}

//电影详细信息如片长、导演、主演、简介、制片国家等信息
type	MovieDetail struct {
	MovieCountry	string
	MovieCompany	string
	Duration	string
	MovieDate	string
	TodayRealSale	string
	AccumulateSale	string
	Director	[]MoviePerson
	Actor		[]MoviePerson
	MovieUrl	string
	MovieFormat string
	MovieYear	string
}
//导演或者主演信息
type MoviePerson struct {
	PersonName	string
	LinkUrl	string
}

//票房信息
type BoxOffice struct {
	Title       string
	TodayBigBar string
	BoxOfficeInfos  []*BoxOfficeInfo
}

//影片信息
type BoxOfficeInfo struct {
	//影片排名
	RankNo int
	//影片名称
	Movie *MovieInfo
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

func (this *BoxOfficeInfo) GetCsvHeader() []string {
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

func (this *BoxOfficeInfo) ValueToStrArray() []string {
	strArr := []string{}
	strArr = append(append(append(append(append(append(append(strArr,
		fmt.Sprintf("%d", this.RankNo)),
		this.Movie.MovieName),
		fmt.Sprintf("%.1f", this.RealSale)),
		this.SaleRatio),
		fmt.Sprintf("%.1f", this.AccumulateSale)),
		this.MovieRatio),
		fmt.Sprintf("%d", this.PublishDays))
	return strArr

}