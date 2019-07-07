package english



//1.每日英语听力的链接
var (
	categoryUrl = `http://www.eudic.cn/ting/channel?id=8a695f40-1da1-11e6-bcc9-000c29ffef9b&type=category`
	//获取一级栏目和二级栏目的链接
	//一级栏目：
	channelUrl = `http://www.eudic.cn/ting/channel?id=e5708ee5-f9a2-11e6-9e96-000c29ffef9b&type=tag`
	//二级栏目：
	articleUrlTemplate = `http://www.eudic.cn/ting/article?id={id}`
	//文章详情：
	detailUrl = `http://www.eudic.cn/webting/desktopplay?id=bcc705dd-1f34-11e7-be40-000c29ffef91&token=
QYN+eyJ0b2tlbiI6IiIsInVzZXJpZCI6IiIsInVybHNpZ24iOiJ1SDVtT0tWc2JnU09pTkhyUkkzS0Job
3BBVDA9IiwidCI6IkFCSU1UVTRNakUyTXpnME9BPT0ifQ%3D%3D`
)
//请求头
var (
	//User-Agent:
	UserAgent=`Mozilla/5.0 (Windows NT 6.3; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.169 Safari/537.36`
	headers=map[string]string{
		"User-Agent":UserAgent,
	}
)

var (
	audioLinkRe="Webting_play.initPlayPage\\(\".+?\"\\);"
)
//采用goquery来提取数据
//一级栏目和二级栏目
//`<dl class="cl_item">
//            <dt class="cl_title">用户上传</dt>
//            <div class="keylist-warp">
//                <div class="inner_warp">
//                        <dd><a href="/ting/channel?id=961207a2-a0d2-11e6-8a0e-000c29ffef9b&type=tag" title="考试一点通" alt="考试一点通">考试一点通</a></dd>
//                        <dd><a href="/ting/channel?id=9f8090ed-a0d2-11e6-8a0e-000c29ffef9b&type=tag" title="大学四六八" alt="大学四六八">大学四六八</a></dd>
//                        <dd><a href="/ting/channel?id=adc2111d-a0d2-11e6-8a0e-000c29ffef9b&type=tag" title="教材资料库" alt="教材资料库">教材资料库</a></dd>
//                        <dd><a href="/ting/channel?id=b6fd0686-a0d2-11e6-8a0e-000c29ffef9b&type=tag" title="实用口语汇" alt="实用口语汇">实用口语汇</a></dd>
//                        <dd><a href="/ting/channel?id=befe8cf9-a0d2-11e6-8a0e-000c29ffef9b&type=tag" title="文学大转盘" alt="文学大转盘">文学大转盘</a></dd>
//                        <dd><a href="/ting/channel?id=df537dc2-a0ce-11e6-8a0e-000c29ffef9b&type=tag" title="影视直通车" alt="影视直通车">影视直通车</a></dd>
//</div><div class='inner_warp'>                </div>
//                <div class="clear"></div>
//            </div>
//            <div class="clear"></div>
//        </dl>
//`

const(
	PREFIX="http://www.eudic.cn"
)