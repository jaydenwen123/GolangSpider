# 古诗文网 #

网站链接：[古诗文网](https://www.gushiwen.org/)

![古诗文网](https://github.com/jaydenwen123/GolangSpider/blob/master/GolangSpider/example/gushiwen/images/home.png)

## 该实例主要的工作： ##

1. 爬取该网站所有的诗文类型
2. 爬取所有类型的诗文链接
3. 爬取每首诗文的详细信息并保存到JSON文件和PostgreSQL数据库中

## 爬取成果展示 ##

1.朝代信息

>![朝代信息](https://github.com/jaydenwen123/GolangSpider/blob/master/GolangSpider/example/gushiwen/images/dynasty.png)

2.诗文分类信息

>![诗文分类信息](https://github.com/jaydenwen123/GolangSpider/blob/master/GolangSpider/example/gushiwen/images/typejson.png)

3.诗人信息

>![诗人信息](https://github.com/jaydenwen123/GolangSpider/blob/master/GolangSpider/example/gushiwen/images/poemer.png)

4.诗文详细信息

>![诗文详细信息](https://github.com/jaydenwen123/GolangSpider/blob/master/GolangSpider/example/gushiwen/images/poemjson.png)

5.数据库存储信息

>![数据库存储信息](https://github.com/jaydenwen123/GolangSpider/blob/master/GolangSpider/example/gushiwen/images/poem_postgresql.png)

## 所涉及的技术 ##

1. 正则表达式
2. Golang操作数据库
3. gorm 库的使用
4. PostgreSQL数据库
5. goroutine和channel
6. Golang处理JSON文件

## 待完成的工作 ##

1. 对爬取的古诗文进行去重(考虑redis)
2. 爬取古诗文完整的译文、注释、赏析信息
3. 增加根据朝代或者诗人维度爬取诗文信息

## 参考资料 ##

1. [gorm开源orm库](https://gorm.io/docs/)
2. [gobyexample](https://gobyexample.com/)
3. [golang中文手册](https://studygolang.com/pkgdoc)
4. [GoSpider](https://github.com/GopherCoder/Go-Spider)