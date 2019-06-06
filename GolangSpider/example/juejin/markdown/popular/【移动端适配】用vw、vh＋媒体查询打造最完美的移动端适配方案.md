# 【移动端适配】用vw、vh＋媒体查询打造最完美的移动端适配方案 #

> 
> 
> 
> 更新——前两天写的这篇文章未曾想能获得这番热烈反响，首先感谢大家的支持与抬爱，菜鸡表示诚惶诚恐。我也是刚入掘金不久，本意是想在这个有原则的前端社区写点文章，把平时积累多总结。一来有助于督促自己，二是希望能给有需要的朋友给予帮助。希望大佬们能多多发表意见或者建议，一起成长，进步！望轻拍，感激(ಥ
> _ ಥ)
> 
> 

> 
> 
> 
> 一篇真正看了就会用的vw、wh适配教程
> 
> 

从古老的的百分比布局+px+媒体查询到rem布局，一直没找到心仪的移动端适配方案。网上搜索的教程质量也是参差不齐(要么配置过于繁琐，要么一篇文章到处抄袭)，反正我看完了总有一种无从下手的无奈。所幸，经同事推荐找到一款完美的插件。欣喜之余，以作记录。同时希望能给需要的朋友提供帮助。

## 移动端适配神器——postcss-px-to-viewport ##

这里不多介绍vw、vh属性，毕竟网上一搜一大把，本文章只有最纯粹的干货。只需要通过包管理工具安装postcss-px-to-viewport插件后进行简单配置就可以在页面直接使用px单位，项目编译后自动转换为对应的vw或vh属性

### px转vw、vh ###

#### 1. 在Github搜索 ` postcss-px-to-viewport` ####

选择星星最多的

![](https://user-gold-cdn.xitu.io/2019/5/31/16b0ccf5ded9cc87?imageView2/0/w/1280/h/960/ignore-error/1)

英语渣表示看到有中文文档很是兴奋

![](https://user-gold-cdn.xitu.io/2019/5/31/16b0ccf5e8e08235?imageView2/0/w/1280/h/960/ignore-error/1)

#### 2. 安装插件 ####

` npm install postcss-px-to-viewport --save-dev 复制代码`

#### 3. 配置参数 ####

![](https://user-gold-cdn.xitu.io/2019/5/31/16b0ccf5e8d0b512?imageView2/0/w/1280/h/960/ignore-error/1)

这里以vue cli3.x版本做参考，在package.json中配置

![](https://user-gold-cdn.xitu.io/2019/5/31/16b0ccf5e73b530e?imageView2/0/w/1280/h/960/ignore-error/1)

以上，现在代码中使用px就能直接转为对应的vw、vh属性了

## 通过媒体查询处理边界情况 ##

> 
> 
> 
> 一般来说使用px转为vw、vh就可以应付99%的移动端适配了，但偶尔会有个别情况需要使用媒体查询适配小屏分辨率
> 
> 

比如以iphone6为基准布局，看起来毫无问题

![](https://user-gold-cdn.xitu.io/2019/5/31/16b0ccf5e8cc2f71?imageView2/0/w/1280/h/960/ignore-error/1)

但在如iphone5等320像素的分辨率下就会有些瑕疵

![](https://user-gold-cdn.xitu.io/2019/5/31/16b0ccf5e9c4d6c5?imageView2/0/w/1280/h/960/ignore-error/1)

明显看到，字体重叠了。这时就可以请出法宝。用媒体查询解决

![](https://user-gold-cdn.xitu.io/2019/5/31/16b0ccf65320cd2c?imageView2/0/w/1280/h/960/ignore-error/1)

代码意思是，当用户手机分辨率(宽度)为320像素到340像素之间的时候做兼容处理，下面来看下处理后的效果

![](https://user-gold-cdn.xitu.io/2019/5/31/16b0ccf651a10f44?imageView2/0/w/1280/h/960/ignore-error/1)

完美解决

## 总结 ##

至于vw、vh属性的兼容性，从https://caniuse.com/网站给出的数据来看，pc端也许有点差强人意，但手机上基本可以放心使用了。(顺带吐槽一下浏览器兼容性真是阻碍技术发展的碍脚石)

![](https://user-gold-cdn.xitu.io/2019/5/31/16b0ccf63a87de68?imageView2/0/w/1280/h/960/ignore-error/1)