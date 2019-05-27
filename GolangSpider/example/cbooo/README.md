# Cbooo中国票房 #

中国票房链接：[中国票房CBO](http://www.cbooo.cn/)

在该例子中，主要完成以下三类数据的爬取：

## 1.中国实时票房榜 ##

### 将爬取的信息存储为JSON文件、XML文件、CSV文件保存到本地 ###

![中国实时票房榜](https://github.com/jaydenwen123/GolangSpider/blob/master/GolangSpider/example/cbooo/images/realrank.png)

### 爬取的实时票房榜的数据信息 ###

**中国实时票房榜JSON数据** 

![中国实时票房榜](https://github.com/jaydenwen123/GolangSpider/blob/master/GolangSpider/example/cbooo/images/boxoffice_json.png)

**中国实时票房榜XML数据** 

![中国实时票房榜](https://github.com/jaydenwen123/GolangSpider/blob/master/GolangSpider/example/cbooo/images/boxoffice_xml.png)

**中国实时票房榜CSV数据** 

![中国实时票房榜](https://github.com/jaydenwen123/GolangSpider/blob/master/GolangSpider/example/cbooo/images/boxoffice_csv.png)


## 2.即将上映 ##

### 将爬取的该信息保存成JSON文件、XML文件 ###

![即将上映](https://github.com/jaydenwen123/GolangSpider/blob/master/GolangSpider/example/cbooo/images/incoming.png)

### 爬取的即将上映的数据信息 ###

**即将上映JSON数据** 

![即将上映](https://github.com/jaydenwen123/GolangSpider/blob/master/GolangSpider/example/cbooo/images/coming_json.png)

**即将上映XML数据** 

![即将上映](https://github.com/jaydenwen123/GolangSpider/blob/master/GolangSpider/example/cbooo/images/coming_xml.png)


## 3.电影信息 ##

### 通过爬虫爬取电影的详细信息，保存成JOSN文件并下载电影海报图片 ###

![电影信息](https://github.com/jaydenwen123/GolangSpider/blob/master/GolangSpider/example/cbooo/images/incoming.png)
>  
### 爬取后的电影数据 ###

**爬取后的电影JSON数据**

![电影信息](https://github.com/jaydenwen123/GolangSpider/blob/master/GolangSpider/example/cbooo/images/movieinfo_json.png)

### 电影海报 ###

**爬取后的电影海报**

![电影海报](https://github.com/jaydenwen123/GolangSpider/blob/master/GolangSpider/example/cbooo/data/682238.jpg)

# 主要涉及的技术 #

1. golang Http
2. JSON、XML、CSV文件的处理
3. 正则表达式解析