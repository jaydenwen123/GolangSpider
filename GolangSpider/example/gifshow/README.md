# 爬取快手短视频 #
----------


> 通过利用fidder手机抓包来实现  
> 网站：https://live.kuaishou.com/profile/yunyinyue1  
> ![快手](https://github.com/jaydenwen123/GolangSpider/blob/master/GolangSpider/example/gifshow/images/web.png)

## 主要任务 ##
1. 爬取热门推荐的**视频信息**并**下载视频**

## 快手接口分析 ##

### 1.热门接口 ###

> //快手热门视频接口分析  
> **POST请求**  
> 
> 
	http://api.gifshow.com/rest/n/feed/hot?app=0&kpf=ANDROID_PHONE&ver=6.5&c=HUAWEI&mod=HUAWEI%28BLN-AL40%29&appver=6.5.1.9253&ftt=&isp=CMCC&kpn=KUAISHOU&lon=0&language=zh-cn&sys=ANDROID_8.0.0&max_memory=384&ud=19523983&country_code=cn&oc=HUAWEI&hotfix_ver=&did_gt=1558500302521&iuid=&extId=84ae6d25936c4eab63503b236a35ad34&net=WIFI&did=ANDROID_d5c03a4945197bbf&lat=0 

> **请求参数：**  
> Cookie: token=647e4728fff6462ca87784fa7cced390-19523983  
> User-Agent: kwai-android  
> Connection: keep-alive  
> Accept-Language: zh-cn  
> **X-REQUESTID: 641013794**  
> Content-Type: application/x-www-form-urlencoded  
> Content-Length: 367  
> Host: api.gifshow.com  
> Accept-Encoding: gzip  




## 成果展现 ##

**1.视频列表**
> ![视频列表](https://github.com/jaydenwen123/GolangSpider/blob/master/GolangSpider/example/gifshow/images/video_list.png)


**2.视频播放画面**
> ![视频播放画面](https://github.com/jaydenwen123/GolangSpider/blob/master/GolangSpider/example/gifshow/images/video_detail.png)

**3.下载日志**
> ![下载日志](https://github.com/jaydenwen123/GolangSpider/blob/master/GolangSpider/example/gifshow/images/download_log.png)

**4.视频简要信息**
> ![视频简要信息](https://github.com/jaydenwen123/GolangSpider/blob/master/GolangSpider/example/gifshow/images/hot_video_json.png)

## 关键技术 ##
1. channel&goroutine
2. 正则表达式提取文章html数据
3. gjson解析json数据

## 参考资料 ##
1. [gjson：https://github.com/tidwall/gjson](https://github.com/tidwall/gjson)

## 待优化的点 ##
1. 将文章数据保存到ElasticSearch中，通过web界面提供搜索接口
2. 根据用户ID，批量下载用户所有视频
3. 支持搜索接口，搜索用户
4. 快手、抖音等短视频都存在app请求加密，后期需要考虑着手加密算法部分