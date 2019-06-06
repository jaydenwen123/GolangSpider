# 微信小程序-沪江英语+商城tap切换功能 #

首先说段pi话，我是一个菜鸡前端小白，我不仅是菜鸡前端我还没过英语四级，现在大三了有点着急，但是我相信我会过的，所以我选择仿了一个跟英语有关的小程序想着能够潜移默化的受到影响，这个小程序就是标题所说的“ **沪江英语** ”(黑体加粗)。这是我做的第一个小程序，但是你们不能因为我菜就不往下看了说不定就能有些收获呢你说是吧（嘿嘿嘿嘿嘿）

# 项目预览及开始前的准备 #

## 项目预览 ##

![](https://user-gold-cdn.xitu.io/2019/6/4/16b2035f24b9f667?imageView2/0/w/1280/h/960/ignore-error/1)

## 开始前准备及使用工具 ##

* 申请账号：进入 [微信公众平台]( https://link.juejin.im?target=https%3A%2F%2Fmp.weixin.qq.com%2Fwxamp%2Fhome%2Fguide%3Flang%3Dzh_CN%26amp%3Btoken%3D465288549 ) 根据指示申请注册账号
* 开发工具： [VScode]( https://link.juejin.im?target=https%3A%2F%2Fcode.visualstudio.com%2FDownload ) 和 [微信开发者工具]( https://link.juejin.im?target=https%3A%2F%2Fdevelopers.weixin.qq.com%2Fminigame%2Fdev%2Fdevtools%2Fdownload.html )
* 使用的文档及第三方工具

* [EasyMock]( https://link.juejin.im?target=https%3A%2F%2Fwww.easy-mock.com%2F ) (一个可视化,并且能快速生成模拟数据的第三方服务)
* 小程序云开发的数据库功能和存储功能(在这里我主要用来存储图片，利用其生成的图片地址)
* [VantWeapp]( https://link.juejin.im?target=https%3A%2F%2Fgitee.com%2Fmirrors%2FVant-Weapp ) 以及其 [开发文档]( https://link.juejin.im?target=https%3A%2F%2Fyouzan.github.io%2Fvant-weapp%2F%23%2Fintro )
* [微信官方文档]( https://link.juejin.im?target=https%3A%2F%2Fdevelopers.weixin.qq.com%2Fminiprogram%2Fdev%2Fframework%2F )
* [markMan]( https://link.juejin.im?target=http%3A%2F%2Fwww.getmarkman.com%2F ) (可以进行图片测量和标注的工具)
* [iconfont]( https://link.juejin.im?target=https%3A%2F%2Fwww.iconfont.cn%2Fhome%2Findex%3Fspm%3Da313x.7781069.1998910419.2 ) (阿里巴巴矢量图标库，有大量icon资源)
* [wxParse]( https://link.juejin.im?target=https%3A%2F%2Fuser-gold-cdn.xitu.io%2F2019%2F6%2F5%2F16b237556d140339 ) (富文本解析组件，能够显示富文本而且它的思路是将整个HTML String转换成一个Node数组 一句话就是能够把html直接渲染到小程序中)

* 

# 项目完成过程 #

刚开始的时候看到原小程序如下图所示

![](https://user-gold-cdn.xitu.io/2019/6/4/16b20b7b31d9270b?imageslim) 我的天，这些都可以点而且每个页面都不一样，我的天，是不是每个页面都要跳转，是不是要建好多个页面，这也太烦了吧，我当时就想着有没有别的方法可以在同一个页面实现这个功能，一开始我没有想出来，因为我连数据该怎么建都不知道，但是我没有放弃，我思前想后头皮发麻，终于我选择了先切页面下面这一段是介绍页面结构的，再后面就是数据的构建以及逻辑了(不想看下面一段内容的请直接跳到下下段内容)

## 页面结构 ##

` "pages" : [ "pages/index/index" , //首页 "pages/theme/theme" , //主题页 "pages/articlels/articlels" , //文章列表页 "pages/content/content" , //文章详情页 "pages/search/search" //搜索页面 ], 复制代码`

### 首页 ###

![](https://user-gold-cdn.xitu.io/2019/6/4/16b21a9b64ebafe7?imageView2/0/w/1280/h/960/ignore-error/1)

` <!-- index.wxml --> <view class= "container" > <!-- 搜索 --> <view class= "search" bindtap= "toSearch" > <van-search value= "{{ value }}" placeholder= "搜索" background= "#49b849" /> </view> <!-- 导航 --> <view class= "weui-grids" > <block wx: for = "{{navigation}}" wx:key= "{{item.id}}" > <navigator url= "" class= "weui-grid" hover-class= "none" url= "../articlels/articlels?id={{item.id}}&navigationText={{item.navigationText}}" > <image class= "weui-grid__icon" src= "{{item.typePic}}" /> <view class= "weui-grid__label" >{{item.typeName}}</view> </navigator> </block> </view> <!-- 广告 --> <swiper class= "banner" indicator-dots= "true" autoplay= "true" interval= "5000" duration= "500" > <block wx: for = "{{advertPic}}" wx:key= "index" > <swiper-item> <image src= "{{item}}" class= "slide-image" /> </swiper-item> </block> </swiper> <!-- --> <scroll-view scroll-y= "true" > <block wx: for = "{{article}}" wx:key= "index" wx: for -item= "article" > <view class= "article" > <view class= "article-title" > <view class= "article-title__text" >{{article.title}}</view> <view class= "article-title__time" >{{article.time}}</view> </view> <!-- --> <view class= "article-column" > <image class= "article-column__img" src= "{{article.imgUrl}}" /> <view class= "article-column__text" >{{article.text}}</view> </view> <!-- --> <view class= "article-content" > <navigator wx: for = "{{article.typeList}}" wx:key= "{{typeList.lId}}" wx: for -item= "typeList" url= "../articlels/articlels?id={{typeList.lId}}&navigationText={{typeList.navigationText}}" class= "article-list" hover-class= "none" > <view class= "article-list__titlt" >{{ type List.listTitle}}</view> <view class= "article-list__des" >{{ type List.listdes}}</view> </navigator> </view> </view> </block> </scroll-view> </view> 复制代码` ` /**index.wxss**/ page { height: 100%; background-color: #f8f8f8; } .contaner{ width: 100%; height: 100%; box-sizing: border-box; } /* 搜索 */ .search{ width: 100%; height: 108rpx; } .van-search{ position: absolute; left: 0; right: 0; } .van-search__content{ height: 54rpx; align-items: center; } /* 导航 */ .weui-grids{ width: 100%; background-color: #ffffff; margin-bottom: 18.75rpx; } .weui-grid{ width: 25%; /*每份占父容器宽度的25%，一排可容纳四个*/ text-align: center; display: inline-block; /*设置为行内块级元素多个元素可在一排显示且设置宽高*/ } .weui-grid__icon{ width: 54rpx; height: 46rpx; margin-top: 53rpx; margin-bottom: 21rpx; } .weui-grid__label{ font-size: 22rpx; color: #333333; margin-bottom: 46rpx; } /* 广告 */ .banner{ width: 100%; height: 150rpx; margin-bottom: 21rpx; } .slide-image{ width: 100%; height: 170rpx; } /* */ .article{ width: 100%; padding: 28rpx 37.5rpx 0; margin-bottom:21rpx; background: #fff; /* padding-top: 28rpx; */ box-sizing: border-box; } .article-title{ position: relative; width: 100%; height: 73rpx; color: #797979; } .article-title__text{ font-size: 29rpx; position: relative; top: 0; left: 0; } .article-title__time{ line-height: 29rpx; font-size: 20rpx; align-items: center; position: absolute; top: 0; right: 0; } /* */ .article-column{ width: 100%; margin-bottom: 53rpx; position: relative; } .article-column__img{ width: 100%; height: 290rpx; } .article-column__text{ position: absolute; font-size: 31rpx; color: #ffffff; left: 21rpx; bottom: 25rpx; padding-right: 21rpx; } /* */ .article-content{ width: 100%; display: flex; flex-flow: row wrap; /* justify-content: space-around; */ align-content: flex-start; } .article-list{ width: 33.33%; flex: 0 0 auto; text-align: center; margin-bottom: 65.625rpx; } .article-list__titlt{ color: #000000; line-height: 26rpx; font-size: 26rpx; margin-bottom: 15.625rpx; } .article-list__des{ line-height: 20.79rpx; font-size: 20.79rpx; color: rgba(0,0,0,0.7); } 复制代码`

### 文章列表页 ###

![](https://user-gold-cdn.xitu.io/2019/6/4/16b21e61f8343f7f?imageView2/0/w/1280/h/960/ignore-error/1)

` <view class= "articlePage" > <view class= "articlesection" > <navigator url= "../content/content?id={{item.articleId}}&totalId={{totalId}}&typeListId={{typeListId}}" class= "articlesection-list" hover-class= "none" wx: for = "{{articles}}" wx:key= "articleId" > <view class= "list-Left" > <view class= "list-Left__title" >{{item.articledesc}}</view> <view class= "list-Left__num" >{{item.num}}</view> </view> <view class= "list-Right" > <image class= "list-ight__img" src= "{{item.img}}" /> </view> </navigator> </view> </view> 复制代码` ` /* miniprogram/pages/articlels/articlels.wxss */ page{ width: 100%; } .articlePage{ width: 100%; background-color: #fff; } .articlesection{ padding: 21.857rpx 35.42rpx 0; box-sizing: border-box; width: 100%; } .articlesection-list{ width: 100%; display: flex; justify-content: space-between; flex-direction: row; border-bottom: 1px solid #f3f3f3; padding-bottom: 37.5rpx; margin-top:37.5rpx ; } .list-Left{ flex: 1; } .list-Left__title{ width: 398rpx; display: -webkit-box; -webkit-box-orient: vertical; -webkit-line-clamp: 2; overflow: hidden; font-size: 31.25rpx; color: #333333; margin-bottom: 90rpx; } .list-Left__num{ font-size: 23.75rpx; color: #cccccc; } .list-Right{ width: 200rpx; height: 151rpx; } .list-ight__img{ width: 200rpx; height: 151rpx; } 复制代码`

### 文章详情页 ###

![](https://user-gold-cdn.xitu.io/2019/6/4/16b2200a354c3503?imageView2/0/w/1280/h/960/ignore-error/1)

` <import src= "../../wxParse/wxParse.wxml" /> <view class= "content" > <scroll-view scroll-y= "true" enable -back-to-top= "true" class= "artContent" > <view class= "artContent_header" > <text class= "artContent_header--title" >{{title}}</text> <view class= "artContent_header--bottom" > <view class= "artContent_header--author" >{{author}}</view> <view class= "artContent_header--time" >{{time}}</view> </view> </view> <view class= "artContent-body" > <template is= "wxParse" data= "{{wxParseData:content.nodes}}" /> </view> </scroll-view> <view class= "footer" > <view class= "footer-follow" > <image class= "footer-follow__icon" bindtap= "backHome" src= "../../images/back.png" /> <view class= "footer-follow__text" >回主页</view> </view> <view class= "footer-share" bindtap= "nextArticle" data-id= '{{id}}' > <image class= "footer-share__icon" src= "../../images/next.png" /> <view class= "footer-share__text" >下一篇</view> </view> </view> </view> 复制代码` ` /* miniprogram/pages/content/content.wxss */ @import "../../wxParse/wxParse.wxss" ; .content{ width: 100%; } .artContent{ width: 100%; background-color: #fbfbfb; padding: 52rpx 36.45rpx 0; box-sizing: border-box; } .artContent_header{ height: 198rpx; display: flex; flex-direction: column; justify-content: space-between; font-weight: 500; border-bottom: 1px solid #ededed; } .artContent_header--title{ font-size: 40rpx; color: #333333; } .artContent_header--bottom{ display: flex; flex-direction: row; justify-content: space-between; font-size: 22rpx; color: #d1d1d1; margin-bottom: 24rpx; } .artContent-body{ margin-top: 59.375rpx; width: 100%; margin-bottom: 147rpx; } .footer{ position: fixed; bottom: 0; width: 100%; background-color: #fff; padding: 16.67rpx 0; } .footer{ position: fixed; bottom: 0; width: 100%; display: flex; flex-direction: row; height: 94.79rpx; background-color: #fff; justify-content: space-around; align-items: center; } .footer-follow,.footer-share{ text-align: center; } .footer-follow__icon,.footer-share__icon{ width: 31.25rpx; height: 31.25rpx; } .footer-follow__text,.footer-share__text{ font-size: 16.67rpx; color: #8a8a8a; } .langs_en{ line-height: 57rpx; font-size: 31.25rpx; margin-bottom: 41.67rpx; } .langs_cn{ color: #666666; font-size: 28.125rpx; line-height: 55.2rpx; margin-bottom: 41.67rpx; } 复制代码`

### 搜索页面 ###

![](https://user-gold-cdn.xitu.io/2019/6/4/16b2209b69419214?imageslim) 由动图可知搜索页面由四个部分组成：搜索框+搜索获得内容列表+文章列表(完全引用了文章列表页面的结构,但是没有封装成组件所以大家先将就着看)+ 搜索示例，我直接上代码吧

` <van-search value= "{{ value }}" placeholder= "请输入搜索关键词" show-action bind :search= "onSearch" bind :cancel= "onCancel" bind :change= "searchInput" /> <scroll-view scroll-y class= "search-list {{is_hidden?'hidden':''}}" > <block wx: for = "{{search_list}}" wx:key= "{{item.articleId}}" > <text class= "search-item" bindtap= "showItemDetail" data-articledesc= "{{item.articledesc}}" >{{item.articledesc}}</text> </block> </scroll-view> <block wx: if = "{{articles==''}}" > <view class= "case" > <view class= "case-item" wx: for = "{{word}}" wx:key= "index" bindtap= "showItemDetail" data-articledesc= "{{item}}" > <text>{{item}}</text> </view> </view> </block> <view class= "articlePage" > <view class= "articlesection" > <navigator url= "../content/content?id={{item.articleId}}&totalId={{totalId}}&typeListId={{item.typeListId}}" class= "articlesection-list" hover-class= "none" wx: for = "{{articles}}" wx:key= "item.articleId" > <view class= "list-Left" > <view class= "list-Left__title" >{{item.articledesc}}</view> <view class= "list-Left__num" >{{item.num}}</view> </view> <view class= "list-Right" > <image class= "list-ight__img" src= "{{item.img}}" /> </view> </navigator> </view> </view> 复制代码` ` page { width: 100%; } .case{ width: 100%; display: flex; flex-direction: row; flex-wrap: wrap; justify-content:space-around; align-items: center; } .case-item{ width: 30%; height: 70rpx; margin-bottom: 20rpx; display: flex; flex-direction: row; justify-content: center; align-items: center; } .case-item text{ font-size: 35rpx; width: 140rpx; background-color: #2cb62c; /* border-radius: 50%; */ color:rgb(116, 18, 62); border: 1px solid rgba(0, 0, 0, .3); text-align: center; border-radius: 20rpx; } /* */ /* */ .search-list{ width: 100%; height: 50vh; display: flex; flex-direction: column; position: fixed; z-index: 2; background: #fff; border-bottom: 1rpx solid #eee; } .search-list .search-item{ display: inline-block; width: 100%; height: 80rpx; line-height: 80rpx; padding: 8rpx 30rpx; border-bottom: 1rpx solid #eee; font-size: 20rpx; white-space: nowrap; overflow: hidden; text-overflow: ellipsis; } .search-list .search-item:last-child{ border-bottom: none; } .search-list.hidden{ display: none; } /* */ /* */ .articlePage{ width: 100%; background-color: #fff; } .articlesection{ padding: 21.857rpx 35.42rpx 0; box-sizing: border-box; width: 100%; } .articlesection-list{ width: 100%; display: flex; justify-content: space-between; flex-direction: row; border-bottom: 1px solid #f3f3f3; padding-bottom: 37.5rpx; margin-top:37.5rpx ; } .list-Left{ flex: 1; } .list-Left__title{ width: 398rpx; display: -webkit-box; -webkit-box-orient: vertical; -webkit-line-clamp: 2; overflow: hidden; font-size: 31.25rpx; color: #333333; margin-bottom: 90rpx; } .list-Left__num{ font-size: 23.75rpx; color: #cccccc; } .list-Right{ width: 200rpx; height: 151rpx; } .list-ight__img{ width: 200rpx; height: 151rpx; } 复制代码`

主题页没啥可以介绍的所以就不介绍了

![](https://user-gold-cdn.xitu.io/2019/6/4/16b22112bd732c85?imageView2/0/w/1280/h/960/ignore-error/1) **接下来重点来了小程序数据的构建还有功能逻辑**

## 数据的构建 ##

### Easy-Mock ###

一开始本小辣鸡就说了连怎么建数据都不知道，后来每天盯着沪江英语小程序的页面看，点了一遍又一遍，格物致知，终于，我想通了

![](https://user-gold-cdn.xitu.io/2019/6/4/16b2216074487bb0?imageView2/0/w/1280/h/960/ignore-error/1) 正经的我要开始告诉你们了。首先我主页的数据是在 [Easy-Mock]( https://link.juejin.im?target=https%3A%2F%2Fwww.easy-mock.com%2F ) 里面创建的接口数据

![](https://user-gold-cdn.xitu.io/2019/6/4/16b221d519db516b?imageView2/0/w/1280/h/960/ignore-error/1) 主要由三部分组成： **navigation** 为首页顶部的文章类型列表的数据， **advertPic** 为swiper广告图片数据， **article** 为flex布局部分的数据此数组内部还嵌套了数组 放上我建的 [EsayMock数据]( https://link.juejin.im?target=https%3A%2F%2Fwww.easy-mock.com%2Fmock%2F5cdb60fd19aa86341d1d37ab%2Fenglish%2Fllm ) 接口的链接可以点进去去看一眼。 每个 **navigation** 以及 **article** 中的每一项里面的 **typeList** 中的item都有一个id/lId，这个id主要是用来从云数据库从取得文章列表内容(articlels页面)中的数据的，其他的数据都是用来渲染页面的。接下来讲一下嵌套数据渲染的使用

` <block wx: for = "{{article}}" wx:key= "index" wx: for -item= "article" > <view class= "article" > <view class= "article-title" > <view class= "article-title__text" >{{article.title}}</view> <view class= "article-title__time" >{{article.time}}</view> </view> <!-- --> <view class= "article-column" > <image class= "article-column__img" src= "{{article.imgUrl}}" /> <view class= "article-column__text" >{{article.text}}</view> </view> <!-- --> <view class= "article-content" > <navigator wx: for = "{{article.typeList}}" wx:key= "{{typeList.lId}}" wx: for -item= "typeList" url= "../articlels/articlels?id={{typeList.lId}}&navigationText={{typeList.navigationText}}" class= "article-list" hover-class= "none" > <view class= "article-list__titlt" >{{ type List.listTitle}}</view> <view class= "article-list__des" >{{ type List.listdes}}</view> </navigator> </view> </view> </block> 复制代码`

如上代码，第一层的循环在上述代码第一行

` <block wx: for = "{{article}}" wx:key= "index" wx: for -item= "article" > 复制代码`

第二层循环在第......倒数第七行

` <navigator wx: for = "{{article.typeList}}" wx:key= "{{typeList.lId}}" wx: for -item= "typeList" 复制代码`

我们将article中的typList遍历了出来。 **需要注意** 的是我把两处代码中的item都通过wx:for-item="xxx*"将item改名，两处不是都需要改，但必须通过wx:for-item="xxx"将一处的item更名，防止渲染时都使用item.xxx进行渲染导致混乱

### 云开发数据库 ###

第一次自己建数据库中的数据的时候，我真的是一个加号一个加号的点，一个字段一个字段的敲，真的！是真的敲的我头皮发麻，精神失常，差点放弃了这个小程序，直到我发现导入那两个字。于是我在桌面上建了记事本把它改成了**.json**格式，里面的数据都用对象的形式敲出来，但是一定要注意格式(细心，一定要细心)，否则导入时会失败。数据格式及内容如下

![](https://user-gold-cdn.xitu.io/2019/6/4/16b22eb311ebcd27?imageView2/0/w/1280/h/960/ignore-error/1) 此数据在集合名为 ShanghaiEnglish 当中所有数据都包含在articles中及下标为0的数据中

![](https://user-gold-cdn.xitu.io/2019/6/4/16b22ed1cf1aeb9a?imageView2/0/w/1280/h/960/ignore-error/1) 然后articles中有41条数据article为文章列表内容，id这个字段用于逻辑匹配

![](https://user-gold-cdn.xitu.io/2019/6/4/16b22f0c6ed8e6b1?imageView2/0/w/1280/h/960/ignore-error/1) article里面的articleid时文章id，typeListId与上一个数据表中的id一致

### 云存储 ###

这个小程序中我主要用云存储储存图片来获取图片地址

![](https://user-gold-cdn.xitu.io/2019/6/4/16b22f4c4d1e7a4c?imageView2/0/w/1280/h/960/ignore-error/1) ![](https://user-gold-cdn.xitu.io/2019/6/4/16b22f4d939e7c61?imageView2/0/w/1280/h/960/ignore-error/1) 点击第一张图的左侧文件名部分会出现第二张图中的信息复制下载地址可以在浏览器中访问图片

由于数据实在太多手撸了几千行假数据(爬虫真的是很好的东西，可惜我还不太会，等我写完这个再学)，不想继续造假了，所以我放过了自己，页面有些地方不好看，你也放过我

![](https://user-gold-cdn.xitu.io/2019/6/4/16b22f85f590de3e?imageView2/0/w/1280/h/960/ignore-error/1)

## 页面逻辑的实现 ##

### 首页跳转至文章列表页 ###

![](https://user-gold-cdn.xitu.io/2019/6/4/16b22fe907837f10?imageslim) 我就直接上代码

` import { API_BASE } from '../../config/api' 复制代码`

由于我把EasyMock的放在了config文件夹的api.js中首先在index.js中引入EasyMock

` Page({ data: { navigation:[], advertPic:[], article:[] }, toSearch: function () { wx.navigateTo({ url: '/pages/search/search' , }) }, onPullDownRefresh () { wx.showLoading({ title: '玩命加载中' , }) wx.request({ url: API_BASE, success: (res) => { wx.hideLoading(); wx.stopPullDownRefresh() } }) }, onLoad: function () { wx.showLoading({ title: '玩命加载中' , }) self = this wx.request({ url: API_BASE, data: {}, method: 'GET' , // OPTIONS, GET, HEAD, POST, PUT, DELETE, TRACE, CONNECT // header: {}, // 设置请求的 header success: function (res){ // success console.log(res) self.setData({ advertPic:res.data.data.advertPic, navigation:res.data.data.navigation, article:res.data.data.article }) wx.hideLoading(); }, fail: function () { wx.showModal({ title: '提示' , content: 您的网络状态不佳, showCancel: false }) }, complete: function () { // complete } }) 复制代码`

onPullDownRefresh实现下拉刷新，当数据没有请求到的时候，现实玩命加载中，当数据请求成功以后消失。 在小程序onLoad(页面加载)生命周期中，用 wx.request API请求easymock中的数据，并将数据放入data中，请求前进行 wx.showLoading请求成功 wx.hideLoading();失败则显示网络状况不佳(自从写了这个小程序我就晓得了这些个小程序，要真的是数据有问题请求不到，就说我网不好，我可能以前都被骗到过，我太单纯了)

` <view class= "article-content" > <navigator wx: for = "{{article.typeList}}" wx:key= "{{typeList.lId}}" wx: for -item= "typeList" url= "../articlels/articlels?id={{typeList.lId}}&navigationText={{typeList.navigationText}}" class= "article-list" hover-class= "none" > <view class= "article-list__titlt" >{{ type List.listTitle}}</view> <view class= "article-list__des" >{{ type List.listdes}}</view> </navigator> </view> 复制代码`

由于从首页点进去以后标题不同所以在页面跳转中传了navigationText， 又由于需要在下一个页面判断这是在首页的哪个list上点进来的，所以又传了个id到articlels页面

` // miniprogram/pages/articlels/articlels.js const db = wx.cloud.database(); //获取数据库引用 复制代码` ` data: { }, /** * 生命周期函数--监听页面加载 */ onLoad: function (options) { wx.showLoading({ title: '玩命加载中' , }) let that = this // console.log(options) wx.setNavigationBarTitle({ // 设置当前标题 title: options.navigationText }) db.collection( "ShanghaiEnglish" ).get().then(res => { let data = res.data[0].articles.find((item) => { return item.id == options.id; // console.log(data.article) }) that.setData({ articles:data.article, totalId:data.article.length, type ListId:data.article.typeListId }) }) wx.hideLoading(); }, 复制代码`

在onLoad中 通过options.xxx获取上一个页面传过来的值。如在

` wx.setNavigationBarTitle({ // 设置当前标题 title: options.navigationText }) 复制代码`

通过options.navigationText获取navigationTitle的值并通过 wx.setNavigationBarTitle这个API进行设置

` db.collection( "ShanghaiEnglish" ).get().then(res => { let data = res.data[0].articles.find((item) => { return item.id == options.id; // console.log(data.article) }) 复制代码`

这段代码主要获取数据库中下标为0的数据(所有数据)中的articles里的id字段与index页面传入的id值相等的整个article数组，即对应的文章列表 然后再通过

` that.setData({ articles:data.article, totalId:data.article.length, type ListId:data.article.typeListId }) 复制代码`

将数据放入data中。articles是文章列表，用来渲染页面的。totalId表示的是articles的文章总数(不写文章不知道，一写文章就发现应该把名字设置为total没有Id的)，typeListId就是文章属于哪一个类型的id啦。

` <navigator url= "../content/content?id={{item.articleId}}&totalId={{totalId}}&typeListId={{typeListId}}" class= "articlesection-list" hover-class= "none" wx: for = "{{articles}}" wx:key= "articleId" > <view class= "list-Left" > <view class= "list-Left__title" >{{item.articledesc}}</view> <view class= "list-Left__num" >{{item.num}}</view> </view> <view class= "list-Right" > <image class= "list-ight__img" src= "{{item.img}}" /> </view> </navigator> 复制代码`

然后又传了值给下一个页面。本段结束。请听下回分解。看累了的话您也去休息会儿

![](https://user-gold-cdn.xitu.io/2019/6/4/16b2336c2060dcf6?imageView2/0/w/1280/h/960/ignore-error/1)

### 文章列表页到content页以及下一篇功能的实现 ###

![](https://user-gold-cdn.xitu.io/2019/6/4/16b23309da8e1a48?imageslim)

![](https://user-gold-cdn.xitu.io/2019/6/4/16b2331797da19c1?imageslim) 我好累啊我不想说了，但是我一得说完。 上回说到通过跳转传了值，来我们来看看怎么用的。

` that.setData({ id:options.id, totalId:options.totalId, type ListId:options.typeListId }) db.collection( "ShanghaiEnglish" ).get().then(res => { let data = res.data[0].articles.find((item) => { return item.id == options.typeListId; }) that.setData({ article:data.article }) let id = that.data.id that.setData({ author: that.data.article[id].author, time:that.data.article[id].day, title:that.data.article[id].articledesc, content:that.data.article[id].content, }) var content = that.data.content; WxParse.wxParse( 'content' , 'html' , content, that, 5); wx.setNavigationBarTitle({ // 设置当前标题 title: that.data.title }) 复制代码`

又双叒来接收传过来的值了，又要从数据库里取值了，然后你会发现取值的这部分和上个页面一毛一样，为什么呢，因为我这个小程序不是又了思路再写的，是一个页面一个页面慢慢实现的。而我又懒得封装了。所以只能骗自己你们这样看不用去别的页面找方法，理解起来更方便。 我们假装以及将完了取数据。然后我们通过上个页面传过来的文章自己本身的id其实也就是下标去完数据源里面塞数据就像这样。

` let id = that.data.id that.setData({ author: that.data.article[id].author, time:that.data.article[id].day, title:that.data.article[id].articledesc, content:that.data.article[id].content, }) 复制代码`

然后内容部分就渲染出来了，然后讲下一部分

![](https://user-gold-cdn.xitu.io/2019/6/5/16b2344281501003?imageView2/0/w/1280/h/960/ignore-error/1)

下一篇功能的实现

` <view class= "footer-share" bindtap= "nextArticle" data-id= '{{id}}' > <image class= "footer-share__icon" src= "../../images/next.png" /> <view class= "footer-share__text" >下一篇</view> </view> 复制代码` ` nextArticle: function (e) { let that = this let currentTargetID = e.currentTarget.dataset.id let totalId = that.data.totalId if (currentTargetID < totalId - 1){ let nextTarget = Number(currentTargetID)+1 wx.navigateTo({ url: 'content?id=' + that.data.article[nextTarget].articleId + '&totalId=' + that.data.totalId + '&typeListId=' + that.data.article[nextTarget].typeListId }) } else { wx.showModal({ title: '提示' , content: '别贪心哦~已经没有内容了' , showCancel: false , success: function (res) { } }) return ; } 复制代码`

我通过给下一篇的view绑定点击事件然后还给了个data-id='{{id}}'用来点击时获取此文章的id属性其实也就是下标。本来不用这个直接用that.data.id就行了，但是我觉得这样洋气一点就加了。 通过判断currentTargetID也就是点击下一篇时当时文章的id(下标)时候小于文章列表中总篇数减一。(为啥减一呢？？因为我是先把下标加了一再进行跳转操作的。)，然后就像这样

` wx.navigateTo({ url: 'content?id=' + that.data.article[nextTarget].articleId + '&totalId=' + that.data.totalId + '&typeListId=' + that.data.article[nextTarget].typeListId }) 复制代码`

把下一篇文章的值传到这个页面。然后又进行上一阶段的朴实的操作。最后呢就是没有文章了给了个提示。这段结束。谢谢大家。下回再见

### 搜索页面的实现 ###

` <van-search value= "{{ value }}" placeholder= "请输入搜索关键词" show-action bind :search= "onSearch" bind :cancel= "onCancel" bind :change= "searchInput" /> 复制代码`

首先我给这个组件绑定了三个bind:search="onSearch" bind:cancel="onCancel" bind:change="searchInput"事件，分别时搜索，取消还有输入框改变。我们一个一个来看

再首先

` onLoad: function (options) { let that = this db.collection( "ShanghaiEnglish" ).get().then(res => { let data = res.data[0].articles const flatterArticle = data.reduce((pre, next) => { return pre.concat(next.article); }, []) that.setData({ allAticle: flatterArticle }) console.log(flatterArticle) }) }, 复制代码`

再onLoad中把所有articles中的所有article中的每一条数据都取了出来，并组成了一个新的数组allAticle。

` // 输入时匹配含有输入内容的字段 searchInput: function (e) { console.log(e.detail) // 将搜索内容存入缓存 wx.setStorageSync( 'keywords' , e.detail); // 调用getList方法搜索数据 let search_list = this.getList(e.detail);//通过getList方法查询标题中包含此内容的所有文章 if (e.detail == "" ) { search_list = []; this.data.is_hidden = true ; } else { this.data.is_hidden = false ; } this.setData({ search_list, is_hidden: this.data.is_hidden }); }, 复制代码`

图上注释的挺详细的，那么问题来了getList方法是啥？请往下看一丢丢

` getList(attr) { let self = this return self.data.allAticle.filter(item => { return item.articledesc.toString().toLowerCase().indexOf(attr) > -1; }); }, 复制代码`

**模糊查询** 圈起来要考。这个代码的意思就是在allArticle中查找标题中有这个attr的所有项。上面调用这个方法传了个e.detail(输入的内容)实参，所以也就是返回标题中含有输入内容的所有项。

` onSearch(e) { // 将缓存中的keywords值赋值给keywords const keywords = wx.getStorageSync( 'keywords' ); wx.showLoading({ title: '请稍等' , }); set Timeout(() => { this.setData({ articles: this.getList(keywords),//articles为所有包含这个keywords的项 is_hidden: true }); wx.hideLoading(); }, 500); }, 复制代码`

上面我觉得我标注的还蛮清楚的我就不讲了

` onCancel () { wx.switchTab({ url: '/pages/index/index' , }) }, 复制代码`

onCancel就是用来返回主页面的，这里需要注意的就是，跳转到tabbar页面必须用wx.switchTab。

### 内容页中的文章内容的完成 ###

一开始看到内容页里面没篇文章结构都有不一样的地点的时候我就佛了，不晓得怎么完成，直到我遇到了它 wxPaser它能将html的结构转换为微信小程序的组件并渲染到页面。先在 [gitup]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Ficindy%2FwxParse ) 中下载wxPaser，然后放入目录中做如下配置

` <import src= "../../wxParse/wxParse.wxml" /> 复制代码` ` @import "../../wxParse/wxParse.wxss" ; 复制代码` ` var WxParse = require( '../../wxParse/wxParse.js' ) 复制代码`

分别在content的wxml，wxss和js中引入wxPaser 然后再在js中

` var content = that.data.content; WxParse.wxParse( 'content' , 'html' , content, that, 5); 复制代码`

WxParse.wxParse(bindName , type, data, target,imagePadding)

* 1.bindName绑定的数据名(必填)
* 2.type可以为html或者md(必填)
* 3.data为传入的具体数据(必填)
* 4.target为Page对象,一般为this(必填)
* 5.imagePadding为当图片自适应是左右的单一padding(默认为0,可选)
* 再在wxml中使用转换过后的数据

` <view class= "artContent-body" > <template is= "wxParse" data= "{{wxParseData:content.nodes}}" /> </view> 复制代码`

还可以在wxss中给你的htmlclass类加上样式

## 最后还没完还另外加一个商城tap切换功能 ##

![](https://user-gold-cdn.xitu.io/2019/6/5/16b2389cd2379adb?imageslim) 看到很多人用是否hidden来操作，我觉得tab一多太麻烦了，所以发动了一下聪明的小脑袋瓜

` <view class= "list {{curIndex === index ? 'listActive' : ''}}" bindtap= "toList" data-id= "{{item.id}}" >{{item.name}}</view> 复制代码` ` toList: function (e) { let that = this let currentId = e.currentTarget.dataset.id that.setData({ curIndex: e.currentTarget.dataset.id, type List: that.data.typeLists[currentId].typeList }) }, 复制代码`

主要的代码就上面几行，代码我放在gitup了可以自己拿哦，随便拿不要客气

# 结语 #

[gitup地址]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fyuhan0927%2Flalala ) 我的第一篇文章终于写完了！！！(撒花撒花撒花～) 希望 天下无双仪态万千母仪天下活力四射魅力无边婀娜多姿人间极品美得超越了凡间所有事物气若幽兰莲步生花沉鱼落雁闭月羞花国色天香倾国倾城花容月貌艳如桃李千娇百媚如花似玉楚楚可人亭亭玉立绝世容颜眉清目秀清的 **小姐姐** 和 一表人才博学多才彬彬有礼才貌双全鹤立鸡群翩翩少年足智多谋仪表不凡气宇轩昂眉清目秀玉树临风剑眉星眸清新俊逸挺鼻薄唇风流风倜傥潇洒英俊古雕刻画淡定优雅飘逸宁人探扇浅笑俊美无涛气宇轩昂风度翩翩仪表堂堂的 **小哥哥** 能给我一个 **star！！** (加粗)也希望各位能不吝赐教，带我和我的技术脱离底层走向巅峰 爱您！尊敬您！