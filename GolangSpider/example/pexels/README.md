# Pexels图片素材网 #

> The best free stock photos & videos shared by talented creators.

> 该网站主要提供一些免费的图片素材和视频，并支持不同尺寸的图片下载

网站首页：

![Pexels图片素材网](https://github.com/jaydenwen123/GolangSpider/blob/master/GolangSpider/example/pexels/example_images/home.png)

网站搜索接口：

![Pexels图片素材网](https://github.com/jaydenwen123/GolangSpider/blob/master/GolangSpider/example/pexels/example_images/search.png)


## 主要任务 ##

1.分析搜索接口

>//中文：
>
>zhPageUrl=`https://www.pexels.com/zh-cn/search/{keyword}/?format=html&page={page}&type=`
>
>//英文：
>
>enPageUrl=`https://www.pexels.com/search/{keyword}/?format=html&page={page}&type=`
>

2.提供输入关键词和中英文两种语言的支持

3.并发解析搜索到的图片

4.保存图片关键信息(图片尺寸、图片链接等)

5.并发下载图片保存到本地

>下载不同尺寸的图片链接：
>
> pictureUrl=`{url}?crop=entropy&cs=srgb&fit=crop&fm=jpg&h={height}&w={width}`
> 

## 爬虫成果 ##

**1.测试中文搜索词**

> 选择语言，并提供搜索词 选择`李连杰`(该网站中文的搜索词搜索到的图片不准确)
> 
> ![Pexels图片素材网](https://github.com/jaydenwen123/GolangSpider/blob/master/GolangSpider/example/pexels/example_images/zhtest1.png)

> 下载图片
> 
> ![Pexels图片素材网](https://github.com/jaydenwen123/GolangSpider/blob/master/GolangSpider/example/pexels/example_images/zhtest2.png)

> 下载日志
> 
> ![Pexels图片素材网](https://github.com/jaydenwen123/GolangSpider/blob/master/GolangSpider/example/pexels/example_images/zhtest3.png)

> 下载完毕
> 
> ![Pexels图片素材网](https://github.com/jaydenwen123/GolangSpider/blob/master/GolangSpider/example/pexels/example_images/zhtest4.png)

> 本地图片
> 
> ![Pexels图片素材网](https://github.com/jaydenwen123/GolangSpider/blob/master/GolangSpider/example/pexels/example_images/lilianjie.png)


**2.测试英文搜索词**

> 选择语言，并提供搜索词 选择`artificial intelligence`(该网站中文的搜索词搜索到的图片不准确)
> 
> ![Pexels图片素材网](https://github.com/jaydenwen123/GolangSpider/blob/master/GolangSpider/example/pexels/example_images/test1.png)

> 下载图片
> 
> ![Pexels图片素材网](https://github.com/jaydenwen123/GolangSpider/blob/master/GolangSpider/example/pexels/example_images/test2.png)

> 下载日志
> 
> ![Pexels图片素材网](https://github.com/jaydenwen123/GolangSpider/blob/master/GolangSpider/example/pexels/example_images/test3.png)

> 下载完毕
> 
> ![Pexels图片素材网](https://github.com/jaydenwen123/GolangSpider/blob/master/GolangSpider/example/pexels/example_images/test4.png)

> 本地图片
> 
> ![Pexels图片素材网](https://github.com/jaydenwen123/GolangSpider/blob/master/GolangSpider/example/pexels/example_images/artificial_intelligence.png)


## 涉及技术 ##

1. goroutine&channel并发
2. html解析库——goquery的使用
3. json解析库——gjson库的使用

## 参考资料 ##

1. [Go-Spider 大牛的爬虫网址](https://github.com/GopherCoder/Go-Spider)
2. [goquery github链接](https://github.com/PuerkitoBio/goquery)
3. [gjson github链接](https://github.com/tidwall/gjson)
4. [gobyexample 网址](https://gobyexample.com/)

