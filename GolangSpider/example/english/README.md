# 每日英语听力 #

爬取每日英语听力上面的所有资料信息（链接、阅读数、日期等），并下载所有的听力中的官方提供下载的视频文件和音频文件以及文章内容。

> [每日英语听力](http://www.eudic.cn/ting/)
> 
> ![每日英语听力首页](https://github.com/jaydenwen123/GolangSpider/blob/master/GolangSpider/example/english/images/home.png)

## 主要工作 ##
1. 爬取一级栏目、二级栏目信息
2. 爬取文章列表信息、文章详情数据
3. 下载音频文件、视频文件、以及文章内容


## 后台接口分析 ##

**1.每日英语听力的链接** 

http://www.eudic.cn/ting/channel?id=8a695f40-1da1-11e6-bcc9-000c29ffef9b&type=category

### 获取一级栏目和二级栏目的链接 ###

**栏目详情**

> ![栏目详情](https://github.com/jaydenwen123/GolangSpider/blob/master/GolangSpider/example/english/images/category_detail.png) 

**1.一级栏目（channel）：**



    http://www.eudic.cn/ting/channel?id=e5708ee5-f9a2-11e6-9e96-000c29ffef9b&type=tag

> ![一级栏目页](https://github.com/jaydenwen123/GolangSpider/blob/master/GolangSpider/example/english/images/category.png) 

**2.二级栏目（文章列表）：**


    http://www.eudic.cn/ting/article?id=30c7c56b-1b3a-11e7-be40-000c29ffef9b
  

> ![二级栏目](https://github.com/jaydenwen123/GolangSpider/blob/master/GolangSpider/example/english/images/article.png)  
>  

**3.文章详情：**

	http://www.eudic.cn/webting/desktopplay?id=bcc705dd-1f34-11e7-be40-000c29ffef91&token=
	QYN+eyJ0b2tlbiI6IiIsInVzZXJpZCI6IiIsInVybHNpZ24iOiJ1SDVtT0tWc2JnU09pTkhyUkkzS0Job
	3BBVDA9IiwidCI6IkFCSU1UVTRNakUyTXpnME9BPT0ifQ%3D%3D

> ![音频文章详情](https://github.com/jaydenwen123/GolangSpider/blob/master/GolangSpider/example/english/images/audio_content.png)   
> ![视频文章详情](https://github.com/jaydenwen123/GolangSpider/blob/master/GolangSpider/example/english/images/video_content.png)   
  

## 涉及技术 ##
1. goquery库使用
2. 正则表达式提取文章html数据
4. golang存储json文件