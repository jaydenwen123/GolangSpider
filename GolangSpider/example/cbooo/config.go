package cbooo

const (
	BASE_PATH="example/cbooo/data/"
	JSON_BOXOFFICE = BASE_PATH+"boxoffice.json"
	CSV_BOXOFFICE  = BASE_PATH+"boxoffice.csv"
	XML_BOXOFFICE  = BASE_PATH+"boxoffice.xml"
	JSON_TODAY_PAIPIAN=BASE_PATH+"today_paipian.json"
	JSON_COMING=BASE_PATH+"coming.json"
	XML_COMING=BASE_PATH+"coming.xml"
	JOSN_RECOMMENDER=BASE_PATH+"recommender.json"
	JSON_MOVIE=BASE_PATH+"movie.json"
)

//正则表达式变量
var (
	//定义替换的字符
	urlTmpStr="$"
	//url变量
	cboooUrl="http://www.cbooo.cn/"
	movieImgUrlTemplate="http://www.cbooo.cn/moviepic/$.jpg"
	movieDetailUrlTemplate="http://www.cbooo.cn/m/$"

	//正则表达式变量
	//boxoffice:CBO票房榜数据
	titleRe           = `<h1 onclick="GotoUrl\('/realtime'\)" class="trtop">(.*?)</h1>`
	todayBigBarRe     = `<h6><span id="week"></span>(.*?)</h6>`
	rankNo            = `<td.*?style='width:60px'>(.*?)</td>`
	movieIdRe		=`<tr.*?id='(\d{0,8})' onmouseover='ChangeTopR\(this.id\)'`
	movieNameRe       = `<td style='width:150px'>(.*?)</td>`
	realSaleRe        = `<td style='width:120px'>(.*?)</td>`
	saleRatioRe       = `<td style='width:80px'>(.*?)</td>`
	accumulateRatioRe = `<td style='width:120px'>(.*?)</td>`
	movieRatioRe      = `<td style='width:80px'>(.*?)</td>`
	publishDaysRe     = `<td style='width:70px;'>(.*?)</td>`

	//即将上映数据
	//当正则表达式中存在括号时，需要特殊处理
	comingTitleRe	=`<h1 onclick="GotoUrl\('/comming'\)" class="trtop">(.*?)</h1>`
	//正则表达式多行匹配：(?s:(.*?))
	comingMovieRe=`<tr id='(\d{0,8})' onclick='gotom\(this.id\)' class='trtop'>(?s:(.*?))</tr>`
	comingMovieInfoRe=`td>(.*?)</td><tdtitle='.*?'>(.*?)</td><td>(.*?)</td><td>(.*?)</td>`

	//提取电影详细信息正则表达式
	moviePosterRe=`<img style="height:260px;" src="(.*?)".*?/>`
	movieDetailBlockRe=`<div class="cont">(?s:(.*?))</div>`
	//[复仇者联盟4：终局之战 2019 Avengers:Endgame 2247.5万 393406.5万 科幻/动作/冒险 181min 2019-4-24（中国） 3D/IMAX 美国 http://www.cbooo.cn/c/6 中国电影集团公司]
//<h2>(.*?)<span>（(.*?)）</span><p>(.*?)</p></h2><p><spanclass="m-span">今日实时票房<br/>(.*?)</span><spanclass="m-span">累计票房<br/>(.*?)</span></p><p>类型：(.*?)</p><p>片长：(.*?)</p><p>上映时间：(.*?)</p><p>制式：(.*?)</p><p>国家及地区：(.*?)</p><p>发行公司：<atarget="_blank"href="(.*?)"title=".*?">(.*?)</a></p>
	movieDetailRe=`<h2>(.*?)<span>（(.*?)）</span><p>(.*?)</p></h2><p><spanclass="m-span">今日实时票房<br/>(.*?)</span><spanclass="m-span">累计票房<br/>(.*?)</span></p><p>类型：(.*?)</p><p>片长：(.*?)</p><p>上映时间：(.*?)</p><p>制式：(.*?)</p><p>国家及地区：(.*?)</p><p>发行公司：<atarget="_blank"href="(.*?)"title=".*?">(.*?)</a></p>`
	moviePersonBlockRe=`<dd>(?s:(.*?))</dd>`
	//http://www.cbooo.cn/p/216930
	moviePersonDetailRe=`<p><atarget="_blank"href="(.*?)"title=".*?">(.*?)</a><span></span></p>`
)
