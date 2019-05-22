package kugou

import (
	"fmt"
	"time"
)

//定义解析榜单的变量
var	(
	//榜单歌曲选择路径
	//body > div.pc_temp_wrap.pc_temp_2col_critical > div > div.pc_temp_side > div.pc_rank_sidebar.pc_rank_sidebar_first
	//body > div.pc_temp_wrap.pc_temp_2col_critical > div > div.pc_temp_side > div.pc_rank_sidebar.pc_rank_sidebar_first > ul > li:nth-child(3)
	boardPath=`body > div.pc_temp_wrap.pc_temp_2col_critical > div > div.pc_temp_side > div.pc_rank_sidebar > ul  a`
)

//定义保存歌曲的目录
var (
	boardSaveDir="example/kugou/board/"
	likeSaveDir="F:/Program Files/go/workspace/src/GolangSpider/example/kugou/like/"
	downloadSaveSongDir="F:/Program Files/go/workspace/src/GolangSpider/example/kugou/download/song/"
	downloadSaveMVDir="F:/Program Files/go/workspace/src/GolangSpider/example/kugou/download/mv/"
)

//定义通用的url
var(
	//定义获取所有榜单html内容的url
	boardUrl=`https://www.kugou.com/yy/html/rank.html`
	//下载单首歌曲的url，直接替换掉歌曲的hash值即可
	songInfoTemplateUrl=`https://wwwapi.kugou.com/yy/index.php?r=play/getdata&hash={}&platid=4`
	//搜索歌曲的url
	searchUrl=`https://songsearch.kugou.com/song_search_v2?keyword={}&page=1&pagesize=$&platform=WebFilter`

	//下载歌曲MV的链接,其中hash和key是需要替换的，key是采用md5加密算法加密的加密规则为：
	//hash=".strtoupper($mv_hash).
	//key=".md5(strtoupper($mv_hash)."kugoumvcloud").
	//golang加密实现如下
	//data:=[]byte(strings.ToUpper("E23438F0448B148E595816DCC15D0D6E")+"kugoumvcloud")
	//val := md5.Sum(data)
	//fmt.Printf("%x",val)
	mvSearchUrl=`https://mvsearch.kugou.com/mv_search?keyword={keyword}&page=1&pagesize={pagesize}&platform=WebFilter`
	mvInfoUrl=`http://trackermv.kugou.com/interface/index/cmd=100&hash={hash}&key={key}&pid=6&ext=mp4&ismp3=0`

	HEADER=map[string]string{"Content-Type": CONTENT_TYPE, "Cookie": COOKIE,
		"Refer":REFER,
		"origin":ORIGIN}
)

//定义发送http请求的请求头信息
var (
	//kg_mid=dc8c2ae8999da9ab67910eac60b6faed; kg_dfid=3dTilM3ox9ex0DBwK11ykf6l; Hm_lvt_aedee6983d4cfc62f509129360d6bb3d=1557830630,1557887161,1558002009; kg_dfid_collect=d41d8cd98f00b204e9800998ecf8427e; Hm_lpvt_aedee6983d4cfc62f509129360d6bb3d=1558010436
	COOKIE="kg_mid=dc8c2ae8999da9ab67910eac60b6faed; kg_dfid=3dTilM3ox9ex0DBwK11ykf6l; Hm_lvt_aedee6983d4cfc62f509129360d6bb3d=1557830630,1557887161,1558002009; kg_dfid_collect=d41d8cd98f00b204e9800998ecf8427e; Hm_lpvt_aedee6983d4cfc62f509129360d6bb3d="+fmt.Sprintf("%d",time.Now().Unix())
	REFER="http://www.kugou.com/"
	CONTENT_TYPE="application/json"
	ORIGIN=`http://www.kugou.com/`


)
//定义根据goquery解析歌曲json信息中的歌曲名、图片链接、以及下载链接
const (
	SONG_NAME_PATH="data.audio_name"
	IMG_URL_PATH="data.img"
	PLAY_URL_PATH="data.play_url"


)


//定义下载搜索歌曲用的信息
var  (
	downloadMaxCount=20
)