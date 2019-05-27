# Somusic #

Somusic下载地址：[Somusic下载地址](https://github.com/jaydenwen123/Somusic/raw/master/somusic.exe "Somusic下载地址")

Somusic英文版文档:[Somusic英文版文档](https://github.com/jaydenwen123/Somusic/blob/master/README.md "Somusic英文版文档")



>Github项目地址：https://github.com/jaydenwen123/Somusic
>
>码云项目地址：https://gitee.com/jaydenwen/Somusic


这是命令行音乐下载器，它包含许多功能，如来自kugou网站的搜索歌曲和mv，下载歌曲和单个或批量的mv，列表搜索歌曲或mvs，显示下载歌曲或mv。以上功能有匹配的命令 你可以使用帮助或h来找到doc.最后遗留一个播放音乐和播放MV的功能。我将在后续完善这个功能。

## 谁应该看到这个项目？ ##

 - 如果你想通过实际项目快速学习golang;
 - 如果您想更好地了解命令行软件的运行方式;
 - 如果您想快速进入网络爬虫世界;
 - 如果您对音乐网站感兴趣，并且想要自己构建音乐播放器;
 - 这个项目对你来说绝对不错。


## 你能得到什么？ ##

1. 用Golang语言爬取音乐网站数据。
2. 通过Golang goroutine和Channel并发下载文件（不仅包括mp3文件和mp4文件，还包括二进制文件和文本文件。）。
3. 掌握JSON和Golang结构体，接口，http，正则以及Golang基础知识的技能。
4. 练习使用优秀的golang开源库，例如：goquery，gjson ......
5. 了解命令行工具的工作原理。 如windows cmd，Golang Web Framework beego的蜜蜂工具。
6. 分析Http关于音乐网站的接口。
7. 通过Golang熟悉网络爬虫的技能。

# 开启旅程#

## 安装##

> 1.创建一个用于存储项目的目录。例如：`cd d:\golang\workspac\.`
>
> 2.你应该执行这个命令`git clone https://github.com/jaydenwen123/Somusic.git`
>
> 3.如果要将此项目移动到％gopath％，则可以将其移动到gopath的src目录中。
>

现在将检索项目到您的本地目录，您可以开始您的旅行。


## 帮助文档##
1.somusic支持许多功能，也可以匹配它的命令。 下表中列出了所有支持的功能。

| 命令 | 参数|		功能 	|	说明  |
|:---------:|:------------|:------------|:--------------|
| gboard |无参数 | 下载排行榜歌曲 | 下载kugou排行榜的歌曲 |
| lsong | [max songid] or <start-end> | 显示搜索到的歌曲 | 按照升序的方式显示音乐列表（注意：显示的歌曲是搜索到但未下载） |
| lmv | [max mvid] or <first-end> | 显示搜索到的MV |	按照升序的方式显示MV列表信息，这个命令和lsong比较相似 |
| gsong | [songid] or <first1-end1,first2-end2...> or <songid1,...,first1-end1,songid2,songid3...> | 从远程服务器下载指定的搜索到的歌曲 | 按照指定的范围下载音乐。支持下载单曲，批量（第一段）歌曲，不连续（songid1，songid5，songid8，...）歌曲和混合以上所有方式进行并发下载 |
| gmv |[mvid] or <first1-end1,first2-end2...> or <mvid1,...,first1-end1,mvid5,mvid7...> |	下载MV歌曲 |	这个命令和gsong命令比较像 |
| psong |[songid] |	播放选择的歌曲 | 这个功能还没有完成。将在不久的将来填写 |
| pmv |[mvid] |	播放选择的MV |	这个功能暂时还没有实现。将在后续实现 |
| qsong |[keyword] |	搜索歌曲 | 根据提供的关键词信息搜索歌曲 |
| qmv |[keyword] |	搜索MV | 格局提供的关键词搜索MV|
| ssong |无参数 |	查看本地已经下载的歌曲 |	显示下载的歌曲列表，您可以使用该歌曲在列表中播放歌曲 |
| smv |无参数 |	查看本地已经下载的MV |	显示下载的MV列表，您可以使用该歌曲在列表中播放MV，这个命令和ssong命令相似 |
| chstyle |[new style string] |	改变somusic命令行的风格 |利用新的风格替换掉旧的风格。这个命令与`style`命令相同。|
| style |[new style string] | 改变somusic命令行的风格 |改变somusic命令行的风格 |
| chdelimiter |[new delimiter chars] |  改变somusic命令行的分隔符|改变命令行的分隔符|
| delimiter |[new delimiter chars] | 改变somusic命令行的分隔符|它将替换掉旧的分隔符.这个命令功能与命令`chdelimiter`相同|
| mvpath |无参数 |	查看当前保存MV的路径 |	查看当前保存MV的路径 |
| songpath |无参数 |	查看当前保存歌曲的路径 |	查看当前保存歌曲的路径 |
| chmvpath |[newmvpath]|	切换存放下载MV的路径 |	更改保存下载mv路径。 使用`~`恢复默认目录 |
| chsongpath |[newsongpath] |	切换存放下载歌曲的路径 |	更改保存下载歌曲路径。 使用`~`恢复默认目录 |
| help or h |无参数|	查看帮助文档|	查看帮助文档 |
| quit or CTRL+C |无参数 |	退出程序 |	退出程序 |
| exit or CTRL+C |无参数 |	退出程序 |	退出程序 |
| cls or clear | 无参数 |清除日志信息 |	清除日志信息。在当前版本中，它只支持windows 清除日志信息.后续版本将添加linux清除日志功能 |


2.以下是帮助文档的图片，该文档在goland ide中运行。
![help document](https://uploadfiles.nowcoder.com/images/20190526/5246296_1558868698516_D20B2D73793D5521A11332723BF277B0 "help1.png") 
![help document](https://uploadfiles.nowcoder.com/images/20190526/5246296_1558868756725_AB641EF9D7AC6895F27071EA1B468587 "help2.png") 
![help document](https://uploadfiles.nowcoder.com/images/20190526/5246296_1558868789022_DD0F2189C31C1CF62847873ECA1B70C2 "help3.png") 

## 用法 ##

在这个部分。 我将使用搜索歌曲关键字：`bigbig`和`天使的翅膀`，以及搜索MV关键字：'Falling Down`和`小幸运`作为示例来说明如何使用somusic程序。

1.**search song with keyword.**

>command: `qsong bigbig`(bigbig)
>![qsong](https://uploadfiles.nowcoder.com/images/20190526/5246296_1558868949409_96FBD661ACE76EAF7FE77581509E8970 "qsong.png") 
>command: `qsong 天使的翅膀`(天使的翅膀)
>![qsong](https://uploadfiles.nowcoder.com/images/20190526/5246296_1558868983767_4F07BDB926DA87664420EC67992185DB "qsong天使的翅膀.png") 

2.**search mv with keyword.**

>command: `qmv falling down`(falling down)
>![qmv](https://uploadfiles.nowcoder.com/images/20190526/5246296_1558869020707_6EDD29937D481AC716449AD56BABC91B "qmv.png") 
>command: `qmv 小幸运`(小幸运)
>![qmv](https://uploadfiles.nowcoder.com/images/20190526/5246296_1558869049894_332CA1DD40D2D72783D9D40ED3E0A755 "qmv小幸运.png") 


3.**list the searched song information.**

>command: `lsong`(小幸运)
>![lsong](https://uploadfiles.nowcoder.com/images/20190526/5246296_1558869080229_4CD9F007037B6EABFC6D1E6F61C7E10F "lsong.png") 
>command: `lsong 11`(big big world)
>![lsong](https://uploadfiles.nowcoder.com/images/20190526/5246296_1558869103797_18C4FF3D0CB1DCA0BBE6671396B8844E "lsong2.png") 
>command: `lsong `(天使的翅膀)
>![lsong](https://uploadfiles.nowcoder.com/images/20190526/5246296_1558869134347_E287CE7BA65CFDCFC548F38FD0DD764A "lsong天使的翅膀.png") 

4.**list the searched mv informtion.**

>command: `lmv`(Falling Down)
>![lmv](https://uploadfiles.nowcoder.com/images/20190526/5246296_1558869216508_C91931FF3A287844CDDC9FD5DA798BD4 "lmv.png") 
>![lmv](https://uploadfiles.nowcoder.com/images/20190526/5246296_1558869239116_89FAA092CC35813D9F081D74BAC8D1F2 "lmv2.png") 
>command:`lmv`(小幸运)
>![lmv](https://uploadfiles.nowcoder.com/images/20190526/5246296_1558869273132_BC70B3B886A60D6D4FBE887D2ED97525 "lmv小幸运.png") 

5.**download the searched song.**

>command: `gsong 3,6`(big big world)
>![gsong](https://uploadfiles.nowcoder.com/images/20190526/5246296_1558869294818_51F0EF08BFEECF1F36B2A45581904EF9 "gsong.png") 

6.**download the searched mv.**

>command: `gmv 1-10`(Falling Down)
>![gmv](https://uploadfiles.nowcoder.com/images/20190526/5246296_1558869319068_5C9B32F32770A8666F6668ED89258025 "gmv.png") 
>![gmv](https://uploadfiles.nowcoder.com/images/20190526/5246296_1558869342865_DE7F06A161C09F6153353B2F3F3A0AE4 "gmv2.png") 
>command:`gmv 1-5`(小幸运)
>![gmv](https://uploadfiles.nowcoder.com/images/20190526/5246296_1558869385494_FF7835517063F821EBD774A83A456AD5 "gmv小幸运.png") 
>![gmv](https://uploadfiles.nowcoder.com/images/20190526/5246296_1558869406482_E2DB4CA516648B6B98545A294193201E "gmv小幸运2.png") 

7.**show the local downloaded songs.**

>command: `ssong`
>![ssong](https://uploadfiles.nowcoder.com/images/20190526/5246296_1558869431753_7B8036CADBA27FC02240E19B16D9F37A "ssong .png") 
>![ssong](https://uploadfiles.nowcoder.com/images/20190526/5246296_1558869480423_7B8036CADBA27FC02240E19B16D9F37A "ssong.png") 

8.**show the local donwloaded mvs.**

>command: `smv`
>![smv](https://uploadfiles.nowcoder.com/images/20190526/5246296_1558869499768_16A2DC30AF7CE024C1583FB415344427 "smv.png") 

9.**show the current saved download songs' directory.**

>command: `songpath`
>![songpath](https://uploadfiles.nowcoder.com/images/20190526/5246296_1558869534188_52B40C4555C92EEF8CC8EB66642217C1 "songpath.png") 

10.**show the current saved download mvs' directory.**

>command: `mvpath`
>![mvpath](https://uploadfiles.nowcoder.com/images/20190526/5246296_1558869555324_B9FD0958ABD03995D8C4D91B41866817 "mvpath.png") 
>

11.**change the saved download mvs' directory.**

>command: `chmvpath D:\歌曲`
>![chmvpath](https://uploadfiles.nowcoder.com/images/20190526/5246296_1558869573384_AE8B0BCF8EE9FF922CBF69F7EAE11593 "chmvpath.png") 

12.**change the saved download mvs' directory.**

>command: `chsongpath D:\歌曲`
>![chsongpath](https://uploadfiles.nowcoder.com/images/20190526/5246296_1558869590618_24C7158D7DCF3F9D932B41CEF72E879E "chsongpath.png") 

13.**change the program command line style.**

>command: `style mimusic`
>![style](https://uploadfiles.nowcoder.com/images/20190526/5246296_1558869614215_B16A32699B181BE29CDC669B82EC36A5 "style.png") 

14.**change the program command line delimiter.**

>command: `delimiter #`
>![delimiter](https://uploadfiles.nowcoder.com/images/20190526/5246296_1558869631904_31B5D191D90DAD72F992CF81692EFC01 "delimiter.png") 

15.**show or find the help document.**

>command: `help`
>![help document](https://uploadfiles.nowcoder.com/images/20190526/5246296_1558868698516_D20B2D73793D5521A11332723BF277B0 "help1.png") 
>![help document](https://uploadfiles.nowcoder.com/images/20190526/5246296_1558868756725_AB641EF9D7AC6895F27071EA1B468587 "help2.png") 
>![help document](https://uploadfiles.nowcoder.com/images/20190526/5246296_1558868789022_DD0F2189C31C1CF62847873ECA1B70C2 "help3.png") 

16.**quit or exit the program.**

>command: `exit`
>![exit](https://uploadfiles.nowcoder.com/images/20190526/5246296_1558869714978_CA8A52904DDD4F2A403E8AC93DF09FE5 "exit.png") 

17.**clear the log information.**

>command: `cls`
>![cls](https://uploadfiles.nowcoder.com/images/20190526/5246296_1558869734360_CCA12B690F232201D8D1A8AE3372F004 "cls.png") 


# 参考资料 #
1. gjson(https://github.com/tidwall/gjson)
2. goquery(https://github.com/PuerkitoBio/goquery)
3. gorm(https://github.com/jinzhu/gorm)
4. beego orm(https://github.com/astaxie/beego/orm)
5. beego logs(https://github.com/astaxie/beego/logs)
6. regexp standard library(https://studygolang.com/pkgdoc)
7. net/http standard library(https://studygolang.com/pkgdoc)
8. channel&goroutine(https://gobyexample.com)

# 需要改进什么 #

 -  1.播放歌曲或在现实中播放MV。
 -  2.将变量配置到文件中。如保存下载歌曲目录和mv目录，软件命令行样式和分隔符。
 -  3.添加缓存模块。它可以提高somusic的性能。

>Github项目地址: https://github.com/jaydenwen123/Somusic
>
>码云项目地址：https://gitee.com/jaydenwen/Somusic
>


