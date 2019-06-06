package juejin

//请求头：
var (
	ContentType=`application/json`
	XAgent=`Juejin/Web`
	UserAgent=`Mozilla/5.0 (Windows NT 6.3; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/72.0.3626.109 Safari/537.36`
	XJuejinClient= `1559818729874`
	XJuejinSrc= `web`
	XJuejinToken=`eyJhY2Nlc3NfdG9rZW4iOiI3d2MzSG9Sb0JOeEV3dnpkIiwicmVmcmVzaF90b2tlbiI6ImdhbklJaE9LdnRJVWdBSkUiLCJ0b2tlbl90eXBlIjoibWFjIiwiZXhwaXJlX2luIjoyNTkyMDAwfQ==`
	XJuejinUid=`5ce8befdf265da1bd1463390`
	headers=map[string]string{
		`Content-Type`:ContentType,
		`X-Agent`:XAgent,
		`User-Agent`:UserAgent,
		`X-Juejin-Client`:XJuejinClient,
		`X-Juejin-Src`:XJuejinSrc,
		`X-Juejin-Token`:XJuejinToken,
		`X-Juejin-Uid`:XJuejinUid,
		}
)

const (
	//定义请求参数的值
	POPULAR="POPULAR"
	NEWEST="NEWEST"
	THREE_DAYS_HOTTEST="THREE_DAYS_HOTTEST"
	WEEKLY_HOTTEST="WEEKLY_HOTTEST"
	MONTHLY_HOTTEST="MONTHLY_HOTTEST"
	HOTTEST="HOTTEST"

	//搜索类型ALL/ARTICLE/TAG/USER
	SEARCH_TYPE_USER="USER"
	SEARCH_TYPE_TAG="TAG"
	SEARCH_TYPE_ARTICLE="ARTICLE"
	SEARCH_TYPE_ALL="ALL"

	SEARCH_DEFAULT_PERIOD="ALL"
	//文章的period

	//1天内
	SEARCH_ARTICLE_DAY="D1"
	//1周内
	SEARCH_ARTICLE_WEEK="W1"
	//3个月内
	SEARCH_ARTICLE_THREE_MONTH="M3"

	RECOMMAND_ID="21207e9ddb1de777adeaca7a2fb38030"
	SEARCH_ID="d9997080c3d67a02bfdae094729fed3b"
	FIRST=20
	PAGESIZE=20
	MAX_PAGE=5
)

var (
	//post请求
	//远程采用rpc通信，通信url和参数模板不变，其中内容改变
	juejinUrl=`https://web-api.juejin.im/query`
	//热门参数：
	popularParam=NewRecommandQueryParam(RECOMMAND_ID,FIRST,"",POPULAR)
	//最近参数：
	newestParam=NewRecommandQueryParam(RECOMMAND_ID,FIRST,"",NEWEST)
	//热榜参数：
	//全部
	hottestParam=NewRecommandQueryParam(RECOMMAND_ID,FIRST,"",HOTTEST)
	//3天内
	ThreeDaysHottestParam=NewRecommandQueryParam(RECOMMAND_ID,FIRST,"",THREE_DAYS_HOTTEST)
	//7天内
	WeeklyHottestParam=NewRecommandQueryParam(RECOMMAND_ID,FIRST,"",WEEKLY_HOTTEST)
	//30天内
	MonthlyHottestParam=NewRecommandQueryParam(RECOMMAND_ID,FIRST,"",MONTHLY_HOTTEST)

	//搜索关键词
	keyword="java"
	//搜索全部：
	searchALLParam=NewSearchQueryParam(SEARCH_ID,FIRST,"",SEARCH_DEFAULT_PERIOD,SEARCH_TYPE_ALL,keyword)
	//搜索用户
	searchUsersParam=NewSearchQueryParam(SEARCH_ID,FIRST,"",SEARCH_DEFAULT_PERIOD,SEARCH_TYPE_USER,keyword)
	//搜索标签
	searchTagsParam=NewSearchQueryParam(SEARCH_ID,FIRST,"",SEARCH_DEFAULT_PERIOD,SEARCH_TYPE_TAG,keyword)

	//搜索全部文章
	searchArticlesParam=NewSearchQueryParam(SEARCH_ID,FIRST,"",SEARCH_DEFAULT_PERIOD,SEARCH_TYPE_ARTICLE,keyword)

	//搜索一天内的文章
	searchDayArticlesParam=NewSearchQueryParam(SEARCH_ID,FIRST,"",SEARCH_ARTICLE_DAY,SEARCH_TYPE_ARTICLE,keyword)
	//搜索一周内的文章
	searchWeekArticlesParam=NewSearchQueryParam(SEARCH_ID,FIRST,"",SEARCH_ARTICLE_WEEK,SEARCH_TYPE_ARTICLE,keyword)
	//搜索三个月内的文章
	searchThreeMonthArticlesParam=NewSearchQueryParam(SEARCH_ID,FIRST,"",SEARCH_ARTICLE_THREE_MONTH,SEARCH_TYPE_ARTICLE,keyword)

	//{"operationName":"","query":"","variables":{"type":"ALL","query":"golang","after":"","period":"M3","first":20},"extensions":{"query":{"id":"d9997080c3d67a02bfdae094729fed3b"}}}
	//搜索类型
	//type:ALL/ARTICLE/TAG/USER

	//文章详情接口url，get请求
	detailUrl=`https://juejin.im/post/{id}`
	//5cf61ed3e51d4555fd20a2f3`
	//5cf61ed3e51d4555fd20a2f3

	recommandArticleDetailPath="data.articleFeed.items.edges.#.node"
	recommandPageInfoPath="data.articleFeed.items.pageInfo"
	searchArticleDetailPath="data.search.edges.#.node.entity"
	searchArticlePageInfoPath="data.search.pageInfo"

	//标签文章
	tagArticleDetailPath="d.entrylist"

	//文章内容
	articleDetailRe=`(<h1\s+class="article-title"[\s\S]+?</div>)</article>`
	//<h1 class="article-title" data-v-3f6f7ca1>如何提升JSON.stringify()的性能？</h1>
	//<div data-id="5cf7ae1b6fb9a07ef06f830a" itemprop="articleBody" class="article-content" data-v-3f6f7ca1>
	//.........content.....
	//</div>
)


var (
	//GET
	GetHotTagUrl=`https://gold-tag-ms.juejin.im/v1/tags/type/hot/page/{page}/pageSize/40`
	GetNewTagUrl=`https://gold-tag-ms.juejin.im/v1/tags/type/new/page/{page}/pageSize/40`
	//PUT
	addTagUrl=`https://gold-tag-ms.juejin.im/v1/tag/subscribe/{id}`
	//GET
	tagAllArticlesUrl=`https://timeline-merger-ms.juejin.im/v1/get_tag_entry?src=web&uid=5ce8befdf265da1bd1463390&device_id=1559818729874&token=eyJhY2Nlc3NfdG9rZW4iOiI3d2MzSG9Sb0JOeEV3dnpkIiwicmVmcmVzaF90b2tlbiI6ImdhbklJaE9LdnRJVWdBSkUiLCJ0b2tlbl90eXBlIjoibWFjIiwiZXhwaXJlX2luIjoyNTkyMDAwfQ%3D%3D&tagId=5597a063e4b08a686ce57030&page={page}&pageSize={pagesize}&sort=rankIndex`
)

//定义保存文章的目录
const (
	MARKDOWN_BASE_DIR="markdown"
	MARKDOWN_HOT_DIR=`markdown\hot`
	MARKDOWN_NEW_DIR=`markdown\new`
	MARKDOWN_PUPULAR_DIR=`markdown\popular`
	MARKDOWN_SEARCH_DIR=`markdown\search`
	MARKDOWN_TAG_DIR=`markdown\tag`
)



