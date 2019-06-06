# 后端转前端 起步 css理论 OOCSS/SMACSS 实战心得 #

> 
> 
> 
> ### 序言：无论什么工作，最好都带上自己的思想，明悟一条清晰的行为方针。 ###
> 
> 2019年，后端转前端工作数月后，认为前端开发样式书写太繁琐，样式之间的覆盖问题也让人头疼，身为一个懒人，得想办法简化自己的工作！

网上查看了2014年的一些css理论，如oocss/smacss。面向对象的css？感觉又回到了后端老本行了。一番仔细阅读后， **值得一试！** 以下为理论链接

****https://segmentfault.com/a/1190000000704006 ( https://link.juejin.im?target=https%3A%2F%2Fsegmentfault.com%2Fa%2F1190000000704006 )**
**

又经过数月的实战应用，已经逐渐适应了自己写的css"框架(?)"，并且也成功安利同事使用， **同事用了都说好！**

目录结构：

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2ade14f67f548?imageView2/0/w/1280/h/960/ignore-error/1)

首先，网上下载了一个reset.css，清理一些body边框，设置box-sizing之类的事情

其次，主要分为3类：

* lay-index 框架布局基础样式
* pro-index 项目功能性的组合样式
* fuc-less width/height/margin/padding等自动生成

最后，则是element-ui的组件样式重置，与主题切换。

## 一、lay-index核心代码之flex(less语法)： ##

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2b52dd834f770?imageView2/0/w/1280/h/960/ignore-error/1)

lay-index是布局css文件，而目前最实用的布局自然是 **flex弹性布局** ,还不了解flex的同学请戳： http://www.ruanyifeng.com/blog/2015/07/flex-grammar.html ( https://link.juejin.im?target=http%3A%2F%2Fwww.ruanyifeng.com%2Fblog%2F2015%2F07%2Fflex-grammar.html )

lay-index做的工作则是，将复杂flex语法属性简化，比如我要实现纵向自适应排列：

` < div class = "flex col" > < div class = "flex1" style = "width:100%" > 占据(100%-30px)*100/170的高度 </ div > < div class = "flex07" style = "width:100%" > 占据(100%-30px)*70/170的高度 </ div > < div class = "flex-none" style = "height:30px" > 占据30px的高度 </ div > </ div > 复制代码`

如果我希望一个div内容 水平居中和 **垂直居中** (画重点，要考！)

` < div class = "flex" > 内容居中了 </ div > 复制代码`

lay-indnx内其他代码，例如：

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2b41fbfee3ae9?imageView2/0/w/1280/h/960/ignore-error/1) ![](https://user-gold-cdn.xitu.io/2019/6/6/16b2b43e030e4737?imageView2/0/w/1280/h/960/ignore-error/1)

## 二、pro-index里的核心代码 ##

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2b542e10eeacc?imageView2/0/w/1280/h/960/ignore-error/1)

这里存储一些由项目产生的组件样式，方便下次开发快速应用。

` <div class= "in-flex pro-box width" > 内容居中了,我有边框了 </div> 复制代码`

## 三、fuc-less里的核心代码 ##

在框架完善过程中发现，逐渐添加的 单元样式，已经能满足日常的工作需要了，唯一不好处理的是：width、height、padding、margin这样的不定数值

为了代码写起来舒畅，于是利用less的循环自增语法，来解决这个问题。

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2b75da9a29311?imageView2/0/w/1280/h/960/ignore-error/1)

最终会得到：

.height30,.width200,.padding5,.marginT15 这些批量产生的css

## 以下为最终代码： ##

` <div class= "WH flex col" > <headerSub class= "width flex-none paddingL30" ></headerSub> <div class= "width flex1 flex row padding10 pro-box" > <leftMenu :leftMenu= "leftMenu" ></leftMenu> <router-view class= "flex1 height" ></router-view> </div> </div> 复制代码`

整个css框架，是在几个月内逐渐累积下来的，也经过了数次迭代，大概在做第三个项目时，发现框架已经没有什么需要新增或修改的东西了。

每次书写页面也比以前快了很多，也不用再记太多的css代码，甚至都不再需要UI出效果图，直接对着原型写前端。

### 我也仔细思考了下，为什么oocss并没有流行起来： ###

* 市面上有适应面更广，能直接使用的，诸如 bootstrap这样的框架
* oocss需要前端开发人员带着自己的想法，去推演一段时间才能成行
* 最终成型的结果，带有浓郁的个人风格，其它人不一定能完全适应
* 属于奇技淫巧，会忽视掉代码基础（长期简写，一些样式的完整写法我也忘了）

### **我又认真的考虑了下，我为什么要继续用这个框架：** ###

## **概因：天下武功，无坚不摧，唯快不破** ##

> 
> 最后，将此安利给那些在小公司独自奋战的前端们，希望你们用更少的时间完成工作，更多的时间学习进步

` console.log(`觉得有帮助的话，点个赞吧\(^o^)/~`,huhuche); 复制代码`

> 
> 
> 
> 
> 
> 
> 
> 
> 
> 
> 
> 
> 
> 
> 
> 
> 
>