# 16年毕业的前端er在杭州求职ing #

> 
> 
> 
> 来杭州也有一两个星期了，这个周末下雨，是在没地去，还是习惯性的打开电脑逛技术论坛，想想也是好久没有更新博文了。。。
> 
> 

## 背景 ##

因为曾经看过一篇文章 [面试分享：一年经验初探阿里巴巴前端社招]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fjawil%2Fblog%2Fissues%2F22 ) 所以来杭州也是带有目标的，网易！如果能有幸加入阿里，也是非常荣幸的。所以面试总是懒懒散散的，大概一天也就面试一家。

然后我的技术栈大概是react+node，GitHub地址： [Nealyang]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2FNealyang )

目前的状态是阿里通过了技术面和交叉面，大概下周一总监面+hr面。网易hr面结束了，在等通知，科大讯飞已经拿到offer了，还是比较不错的offer，别的创业公司、上市公司不管是人工智能还是智能家居也都基本拿了offer，但是。。。好吧，还是有着一颗对大厂的夙愿。或许就是大学埋下的吧。

因为个人比较懒得跑面试，所以很多公司的电面我都接了，但是很多公司的现场面试都没有去，哎呀呀，好吧，我懒~这里我大概回一下所有面试所问到的问题吧，因为之前没想去记录，所以很多我都忘记是哪一家了，索性一不做二不休，直接一股脑回忆下面试题吧。能想起来的我标注下是哪家公司。

## 面试 begin ##

大概我是一月分离开的环球网，然后再北京收拾东西总结知识，开始投递简历，大概都是boss直聘和拉钩上面

### HTML & CSS 部分 ###

* css常用布局

这个在面试上市公司和创业公司问的比较多。大概我会回答一些盒模型包括怪异盒模型，定位布局，流布局，浮动布局，flex和grid布局，包括还有三栏布局中的圣杯和双飞翼。这些都还比较熟悉，所以问到都还知道。其中flex布局问的比较多，阿里的交叉面还有别的公司有问到子元素的一些属性。

* BFC

这个滴滴面试的时候有问道（滴滴是破例让我加入流程中的，并且他们还招的技术栈是vue）一般在问清除浮动的时候会说一下

* 居中问题

这个应该是老生常谈的东西了，电话面试的时候有两家问到

* session、cookie、sessionStorage、localStorage等区别

这个也是上市公司和创业公司问到，大概就是说了下中间的区别，还有会简单说下cookie的属性，以及一些前端安全方面

* px/em/rem的区别

滴滴电面的时候问的，这个我也知道，大概说了下相对于父元素还是文档来确定大小之类的。

* animation和transiton的相关属性

这个我也就用了个大概，大概知道的简写位置和属性，当然，阿里一面还问到，为什么动画推荐用c3而不是js，这个问题当时并没有回答好，大概就是从性能上扯了扯，但是什么占用主线程以及浏览器对c3加速都没聊到。然后网易面试也问到了，然后我巴拉巴拉说了下后来查的相关东西。然后网易问了一句，浏览器怎么优化的动画。。。我。。。不知道。

* css编写注意事项

因为这个在之前团队里面没有明文规定，所以我也没总结过，大概说了下自己编码中的方式，和浏览器查抄的过程。

* css和HTML 问的的确都不是很多，然后还有什么标签，meta和media啥的，大概也就介绍了下，问的都不是很深，我也没有回答的很深。。。因为我HTML CSS真的一般般。

### JavaScript部分 ###

* JavaScript数据类型分哪些

这是一个初创公司电面的问题，问的都非常基础，比如css画三角形之类的。别说，之前没准备，还真的忘记了border怎么设置出现直角三角形还是等腰三角形。当然，这个类型还是。。。没得说的

* JavaScript闭包

这个应该问的都比较多，我之前总结过，以及常用的场景，也结合es6谈了下作用域和单例模式谈了下

* 前端跨域

这个我基本都知道，之前有在掘金上总结过，这个很多公司又问道，包括阿里、网易一面。一般方式我都知道，具体展开会把CORS跨域，heade信息字段都说了一遍。也不难

* JavaScript继承

这个我之前也总结过相关文章，网易的一面第二个面试官问了，我大概从原型继承、构造函数继承、组合继承、寄生组合继承优缺点和实现方式都说了下，还有es6的实现方式。一般这样就回答差不多了。后来网易还接着问，es5如何实现super关键字。看过babel转换后代码，但是这个。。真的忘记看了。大概说了下自己的实现思路，也就是装饰着模式。然后也就浑过这题了。

* JavaScript的节流和防抖

滴滴一面问到了，阿里交叉面的时候聊业务场景的时候，也有问到。之前看过文章，自己项目中也用过，所以大概知道些

* JavaScript的事件

阿里交叉面问到的js事件执行机制。我大概谈了下event loop，microtask，task queue。然后事件委托、捕获、冒泡、目标阶段大概谈了下，也顺道谈了下target和currentTarget。

* ajax请求方式

因该算是考察基础功吧，谈了下XMLHTTPRequest的过程，readyState的几种类型和代表的意思。以及浏览器兼容性的处理方案。

* js判断数据类型的方法

貌似有两家公司问到，大概说了下typeof、instanceof、constructor、prototype等判断方式，注意事项以及优缺点。应该回答的还可以

* 函数声明和变量声明

这个大概我也知道，还说了下es6的相关东西

* this指向的问题

这个我也总结相关文章，大概说了下四种绑定规则，还说下new的执行过程以及箭头函数注意事项

* 

面向对象的理解 滴滴一面问的，大概说了下理解以及实现，从封装、继承和多态上说了下es5和es6的实现方式

* 

对于js这门语言你认为怎么样

哇，这个问题问的真的大。有看过《JavaScript语言精粹》，大概说了哪些弱类型语言通病，因为之前搞过Java，所以综合对比了下，同时也说了这些诟病怎么解决。应该会的面试官还是挺满意的

* es6相关知识点

这个应该回答的都不是很深入，大概我都用过，promise的实现方式也研究过，但是不记得哪一家公司问到generator的怎么实现的。大概从iterator上简单说了自己的方案，然后说没看过。然后对于别的其实问的不是很多。基本套路就是es6了解过吗？用过哪些语法。后面具体可能会说下哪一个新特性的实现方式或者转向babel、webpack的相关面试。

### React部分 ###

* react部分必考的肯定有生命周期

这里我大概说了下每一个生命周期，es5、es6的两种书写方式，以及每一个生命周期我们一般用来做些什么操作

* setState是异步的还是同步的

阿里一面的时候问到的，我大概说了两种setState设置方式，以及表现为同步的那种设置方式展开说了下

* 子组件和父组件componentDidMount哪一个先执行

这个也大概从生命周期分期了下。话说我到现在还不知道自己回答的对不对，技友们，你们觉得呢？

* redux的一般流程

这个我比较熟悉，一带说了下所有的技术栈，以及react-redux的原理、高阶组件、以及redux-saga的实现原理。（逮住会的，都啪啪啪说出来，自己掌握点节奏。但是要适当，比如问到我es6，我啦啦啦说了一二十分钟，一般面试官会有点不耐烦。所以视情况而定）

* 如何设计一些组件，原则是什么，你写过什么自豪或者眼前一亮的组件

阿里一面以及一家上市公司也闻到过这类似的问题，大概从组合、复用、重复、测试、维护等方面说了下

* a组件在b组件内，c组件在a组件内，如何让他渲染出来，a组件和c组件同级

阿里面试的时候问到的问题，想了一会，说了不会。后来查了下，大概可以通过react16中返回不带包裹元素的组件来实现。因为和阿里一面面试官后来聊得比较开心，加了微信，还斗胆为了下他，他说还有曲线救国的实现方式

* react组件的优化

从pureRenderMixin、ShouldComponentUpdate等方面说了下，以及组件的设计和木偶组建的函数编写方式说了下

* react组件的通信

这个大搞几种方式也都说了下，prop，context（顺道扯了react-redux的context实现方式）、redux甚至广播都说了一遍

* react 的virtual dom和diff算法的实现方式

阿里交叉面问的，直接说实现方式源码没有看过，但是大概说了下原理和步骤，具体代码怎么写的不知道。

* MVC、MVVM了解么，数据双向绑定和单向绑定实现方式

滴滴一面问的，实现方式还是说了不知道，然后说了下MVC和MVVM的设计模式，因为之前用过angular1，大概就说下脏检查步骤以及view-model的作用

* react-router实现方式，单页面应用相关东西

大概说了下react-router的一般使用方式，以及没有使用react-router的时候如何利用h5 的history API来实现路由跳转等。

* react的ssr了解么？大概怎么实现

阿里的一面问的，在github上写过demo，但是没有用过别的第三方库，这里我就大概说了下webpack的配置项以及大概的实现思路和注意事项。

* react大概也就问了这么写，别的就是具体的业务场景改怎么写代码怎么分析，比较不大众，这里我就我细说了。其实也就考验你的项目经验吧。当然，还有一些react Native的面试题，比如常用组件，和原生如何通信之类的，这些就有赞问的多，但是因为RN玩的不是很透彻，所以对于交互原理都不是很明白。

### 浏览器 ###

* http三次握手后拿到HTML，浏览器怎么加载

阿里的一面问的问题，这个我之前在环球做过相关技术分享，所以大概都知道，从过程到不同内核差异（差异部分简单提了下）说了下dom、CSSDom以及paint等过程。然后面试官接着问如何防止repaint和reflow。大概从引起repaint和reflow等操作上说了下避免。网易的一面也问到了repaint和reflow。

* 前端优化一般都做哪些

这个之前总结过，雅虎的军规啥的。以及首屏优化。然后面试跟了些预加载http head信息相关的，这个没怎么看，回答的不是很好

* 浏览器缓存

这个我也做了相关的技术分享，也看过《图解http》大概从http 1.0和1.1都说了下，其中有一家公司问到200 From cache和200 ok区别（有赞），这个还真的忽略了，后来查了下大概了解了。其实也就是强缓存

* http常见状态码

从100~500 大概也说了十几种。其实也就是《图解http》中的东西，当时还刻意背了下

* http2.0相关

网易一面问题，说了下2.0的采用二进制格式、多路复用、报文头压缩、服务器主动推送还扯了websocket的相关内容 [WebSocket：5分钟从入门到精通]( https://juejin.im/post/5a4e6a43f265da3e303c4787 ) 。然后网易接着问，报文头怎么压缩的？我。。。？？不知道。。。然后大概也问了下https的TLS/SSL,之前看过漫画的htts的相关东西，大概说了下漫画里面的故事~

* post、get区别

这个回答的不是很好，也是一个大厂问的题目，我回答的都是表象。后来我看了一篇文章，大概知道了。 [99%的人都理解错了HTTP中GET与POST的区别]( https://link.juejin.im?target=https%3A%2F%2Fmp.weixin.qq.com%2Fs%3F__biz%3DMzI3NzIzMzg3Mw%3D%3D%26amp%3Bmid%3D100000054%26amp%3Bidx%3D1%26amp%3Bsn%3D71f6c214f3833d9ca20b9f7dcd9d33e4%23rd )

* 别的我也不记得了，回头想起来在来补充吧

### 构建工具 ###

* 编写过webpack的扩展嘛，Plugin或者loader

这个我看过一本书《深入浅出webpack》，所以基本都能回答上来。包括原理和编写loader、Plugin注意事项。当然，我自己没有写过。。。 [《深入浅出webpack》]( https://link.juejin.im?target=http%3A%2F%2Fwebpack.wuhaolin.cn%2F )

* babel 问的不多，但是我也准备了，包括每一个包的作用和内部转换过程，不记得哪家公司问了，大概我也就说了下babel转换的过程。

## 结束语 ##

下周起阿里终面，网易等通知。别的公司基本offer也都拿到了，但是大厂毕竟大厂，基本拿到的offer都过期了。。。并没有办法，毕竟有个大厂梦。好吧，其实还是挺幸运的，阿里用人部门主管leader和一面面试官，加了微信都聊得挺开心的。还感谢主管GitHub follow了我。期望加入~

基本想到的就是就这么些，后面如果想到再来补充吧。

ps：最近掘金 关注涨的有点多，也不知道为啥，最近也没有怎么学习，所以就匆匆忙忙整理这些不知道是不是干活的干活。如果整理的不好，还望见谅。后续结束求职经历，在来细细雕琢。