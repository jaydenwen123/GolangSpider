package gushiwen

//常用的url
const(

	SHIWEN_URL="https://www.gushiwen.org/shiwen/"
	POEM_BASE_URL="https://so.gushiwen.org"
	SHIWEN_BASE_URL="https://www.gushiwen.org"
	DUMP_BASE_PATH="example/gushiwen/data/"
	DUMP_POEM_BASE_PATH="example/gushiwen/data/poems/"
	JSON_SHIWEN_TYPE=DUMP_BASE_PATH+"types.json"
	JSON_SHIWEN_AUTHOR=DUMP_BASE_PATH+"author.json"
	JSON_SHIWEN_DYNASTY=DUMP_BASE_PATH+"dynasty.json"
	JSON_SHIWEN_STYLE=DUMP_BASE_PATH+"style.json"
)

var(
	//解析诗文类型、作者、朝代、形式
	shiwenKindRe=`<div class="cont">(?s:(.*?))</div>`
	singleShiwenRe=`<a.*?href="(.*?)">(.*?)</a>`
	typePoemRe=`<span><a href="(.*?)" target="_blank">(.*?)</a>(.*?)</span>`

	//解析诗文的标题
	titleRe=`<h1 style="font-size:20px; line-height:22px; height:22px; margin-bottom:10px;">(.*?)</h1>`
	//解析诗文的朝代
	dynastyRe=`<p class="source"><a href=".*?">(.*?)</a><span>：</span><a href=".*?">.*?</a> </p>`
	//解析诗文的内容
	contentRe=`<div class="contson" id="(.*?)">(?s:(.*?))</div>`
	//解析诗文的译文

	//数据库配置信息
	dbInfo="host=127.0.0.1 port=5432 user=postgres dbname=shiwen password=wen6224261995 sslmode=disable"
	dbDialect="postgres"
)