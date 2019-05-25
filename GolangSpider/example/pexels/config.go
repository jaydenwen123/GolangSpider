package pexels

var (
	//原始的url，可以将format格式设置为html，则返回的就直接是html代码
	//而且seed参数要不要皆可以
	//https://www.pexels.com/zh-cn/search/%E8%82%8C%E8%82%89%E7%94%B7/
	//pageUrl=`https://www.pexels.com/{lanuage}/search/{keyword}/?format=js&seed=2019-05-24%2B04%3A35%3A11%2B%2B0000&page=1&type=`
	//中国
	zhPageUrl=`https://www.pexels.com/zh-cn/search/{keyword}/?format=html&page={page}&type=`
	//美国
	enPageUrl=`https://www.pexels.com/search/{keyword}/?format=html&page={page}&type=`
	//下载任意宽度和高度的图片
	pictureUrl=`{url}?crop=entropy&cs=srgb&fit=crop&fm=jpg&h={height}&w={width}`

)

//正则表达式解析的变量
var (
	totalEnRe=`<a class="rd__tabs__tab rd__tabs__tab--active".*?>(.*?)Photos</a>`
	totalZhRe=`<a class="rd__tabs__tab rd__tabs__tab--active".*?>(.*?)张照片</a>`
)


//目录变量
var (
	pictureBaseDir=`F:/Program Files/go/workspace/src/GolangSpider/example/pexels/picture`
)



const (
	//每页的图片数量
	PAGE_SIZE=30
	//图片的尺寸
	SMALL="small"
	ORIGINAL="original"
	MEDIUM="medium"
	LARGE="large"
)


const	(
	COOKIE=`__cfduid=dd6a23199d6acc1064024c5a96fee53391558686087; _fbp=fb.1.1558686123608.499776825; _ga=GA1.2.133995655.1558686128; _gid=GA1.2.1730743129.1558686128; _pexels_session=T0pEQ21SOFkvSkk2SWlqank2ZmxWKyt4cmNFUnVadXRNUXZmSHoxajMwL25mNGhFM1ZIQlFrMU1sOXM0anYzU01YMW5zb0ZjNDNreTNvdUwySVV5LzVQQ0xRbWFpU1JuVXQ2UHcrZ0kwckIycFYyRzFGSGxqem1DWnBuUFpVOE5YSG5Ca05MUkRaVDd3RkU2clc2eUk1Yk5nc01FbzB4eFZZNk4yUGtNT2orTXliNEZyeTY3cGdEdXhFeU5FK3h1NGw5dENzQWdZZWlKeEdaV3RJcjV5Zz09LS1WdVhKQ1E4VUwyTVZGMDIyaVd2eEVnPT0%3D--051604ea490533b15e3eeab4040ced75a19ac946; locale=zh-CN; _gat=1`
	ACCEPT=`text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8`
	USER_AGENT=`Mozilla/5.0 (Windows NT 6.3; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/72.0.3626.109 Safari/537.36`
)

var(
	headers=map[string]string{"cookie":COOKIE,"accept":ACCEPT,"user-agent":USER_AGENT}
	languages=map[int]string{1:"中文",2:"英文"}
	pageUrl=map[int]string{1:zhPageUrl,2:enPageUrl}
)

