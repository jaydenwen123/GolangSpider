# 爬取掘金文章 #
----------


> [官网：https://juejin.im](https://juejin.im)
> 
> ![掘金主页](https://github.com/jaydenwen123/GolangSpider/blob/master/GolangSpider/example/juejin/images/home.png)

## 主要任务 ##
1. 爬取掘金**最新**、**最热**、**热榜(3天、7天、30天)**文章信息
2. 根据关键词搜索**文章(一天、一周，三月)**、**用户**、**标签**、**综合数据**，爬取文章信息
3. 获取所有的标签信息
4. 批量**关注标签**
5. 爬取**某一标签的全部文章信息**
6. 将爬取的文章保存成**markdown格式**存储

## 后台接口分析 ##

### 1.后台首页接口 ###

> ![掘金主页](https://github.com/jaydenwen123/GolangSpider/blob/master/GolangSpider/example/juejin/images/home.png)


**1.1热门内容接口**

> https://web-api.juejin.im/query  
> post  
> **请求头：**  
> Content-Type: application/json  
> X-Agent: Juejin/Web   
> **参数：**  
> `{"operationName":"","query":"","variables":{"first":20,"after":"","order":"POPULAR"},"extensions":{"query":{"id":"21207e9ddb1de777adeaca7a2fb38030"}}}`    
> 

**1.2热榜内容接口** 

> https://web-api.juejin.im/query  
> post  
> **请求头：**  
> Content-Type: application/json  
> X-Agent: Juejin/Web  
> **参数：**  
> `{"operationName":"","query":"","variables":{"first":20,"after":"","order":"THREE_DAYS_HOTTEST"},"extensions":{"query":{"id":"21207e9ddb1de777adeaca7a2fb38030"}}}`  


**1.3最新内容接口**

> https://web-api.juejin.im/query  
> post  
> **请求头：**  
> Content-Type: application/json  
> X-Agent: Juejin/Web  
> **参数：**  
> `{"operationName":"","query":"","variables":{"first":20,"after":"","order":"NEWEST"},"extensions":{"query":{"id":"21207e9ddb1de777adeaca7a2fb38030"}}}`  


### 2.文章详情接口 ###

> ![文章详情](https://github.com/jaydenwen123/GolangSpider/blob/master/GolangSpider/example/juejin/images/article_detail.png)


**2.1文章详情接口** 
 
> https://juejin.im/post/{id}  
> id:5cf61ed3e51d4555fd20a2f3  
> get请求    

**2.2文章内容** 
 
>   
	<h1 class="article-title" data-v-3f6f7ca1>如何提升JSON.stringify()的性能？</h1>  
	    <div data-id="5cf7ae1b6fb9a07ef06f830a" itemprop="articleBody" class="article-content" data-v-3f6f7ca1>  
	    .........content.....
	    </div>  


### 3.搜索接口 ###

> ![文章详情](https://github.com/jaydenwen123/GolangSpider/blob/master/GolangSpider/example/juejin/images/search.png)


**3.1掘金搜索接口**

> https://web-api.juejin.im/query  
> post  
> **请求头：**  
> User-Agent: Mozilla/5.0 (Windows NT 6.3; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/72.0.3626.109 Safari/537.36
> X-Agent: Juejin/Web  
> **参数：** 
>    
> `{"operationName":"","query":"","variables":{"type":"ALL","query":"golang","after":"","period":"ALL","first":20},"extensions":{"query":{"id":"d9997080c3d67a02bfdae094729fed3b"}}}`
> `{"operationName":"","query":"","variables":{"type":"ALL","query":"golang","after":"","period":"M3","first":20},"extensions":{"query":{"id":"d9997080c3d67a02bfdae094729fed3b"}}}`
>
> //type:ALL/ARTICLE/TAG/USER

### 4.标签 ###

> ![标签](https://github.com/jaydenwen123/GolangSpider/blob/master/GolangSpider/example/juejin/images/tag.png)

**4.1获取标签信息**

> https://gold-tag-ms.juejin.im/v1/tags/type/hot/page/1/pageSize/40
> Origin: https://juejin.im
> Referer: https://juejin.im/subscribe/all
> User-Agent: Mozilla/5.0 (Windows NT 6.3; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.169 Safari/537.36
> X-Juejin-Client: 1559818729874
> X-Juejin-Src: web
> X-Juejin-Token: eyJhY2Nlc3NfdG9rZW4iOiI3d2MzSG9Sb0JOeEV3dnpkIiwicmVmcmVzaF90b2tlbiI6ImdhbklJaE9LdnRJVWdBSkUiLCJ0b2tlbl90eXBlIjoibWFjIiwiZXhwaXJlX2luIjoyNTkyMDAwfQ==
> X-Juejin-Uid: 5ce8befdf265da1bd1463390


**4.2关注标签**

> addTagUrl=`https://gold-tag-ms.juejin.im/v1/tag/subscribe/5597a23fe4b08a686ce5a7c4`  
> //PUT  
> **请求头：**  
> User-Agent: Mozilla/5.0 (Windows NT 6.3; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.169 Safari/537.36
> X-Juejin-Client: 1559818729874  
> X-Juejin-Src: web  
> X-Juejin-Token:   eyJhY2Nlc3NfdG9rZW4iOiI3d2MzSG9Sb0JOeEV3dnpkIiwicmVmcmVzaF90b2tlbiI6ImdhbklJaE9LdnRJVWdBSkUiLCJ0b2tlbl90eXBlIjoibWFjIiwiZXhwaXJlX2luIjoyNTkyMDAwfQ==  
> X-Juejin-Uid: 5ce8befdf265da1bd1463390  


**4.3获取标签全部文章**

> ![文章详情](https://github.com/jaydenwen123/GolangSpider/blob/master/GolangSpider/example/juejin/images/tag_article.png)

> **get请求**
> 
> `https://timeline-merger-ms.juejin.im/v1/get_tag_entry?src=web&uid=5ce8befdf265da1bd1463390&device_id=1559818729874&token=eyJhY2Nlc3NfdG9rZW4iOiI3d2MzSG9Sb0JOeEV3dnpkIiwicmVmcmVzaF90b2tlbiI6ImdhbklJaE9LdnRJVWdBSkUiLCJ0b2tlbl90eXBlIjoibWFjIiwiZXhwaXJlX2luIjoyNTkyMDAwfQ%3D%3D&tagId=5597a063e4b08a686ce57030&page=0&pageSize=20&sort=rankIndex`
> 


## 成果展现 ##

**1.文章列表**
> ![文章列表](https://github.com/jaydenwen123/GolangSpider/blob/master/GolangSpider/example/juejin/images/article_list.png)


**2.文章详情**
> ![文章详情](https://github.com/jaydenwen123/GolangSpider/blob/master/GolangSpider/example/juejin/images/article_show.png)

**3.下载日志**
> ![下载日志](https://github.com/jaydenwen123/GolangSpider/blob/master/GolangSpider/example/juejin/images/downloadlog.png)

**4.文章简要信息**
> ![文章简要信息](https://github.com/jaydenwen123/GolangSpider/blob/master/GolangSpider/example/juejin/images/article_info.png)

## 关键技术 ##
1. golang html转markdown
2. 正则表达式提取文章html数据
3. http不同方法PUT、DELETE、GET、POST发送请求
4. gjson解析json数据

## 参考资料 ##
1. [html2text：https://jaytaylor.com/html2text](https://jaytaylor.com/html2text)
2. [gjson：https://github.com/tidwall/gjson](https://github.com/tidwall/gjson)

## 待优化的点 ##
1. 将文章数据保存到ElasticSearch中，通过web界面提供搜索接口
2. 文章保存成markdown时，代码格式比较乱，后期考虑优化
3. 采用redis对爬过的文章去重操作