# Vue核心50讲 | 第四回：Vue 官方赠送的橙色装备，岂能不要 #

书接上文，上一回咱们快速的了解了 Vue 的生命周期，知道了在 Vue 的生命周期中存在三个比较重要的阶段，分别是 Created、Mounted 和 Destroyed。接下来，咱们就来说一说 Vue 官方赠送的橙色装备 —— vue-devtools。

说到 vue-devtools，使用 Vue 开发的时候 Vue 官方推荐在浏览器安装 Vue Devtools 这个工具。这个时候可能你会问了，这个工具是干啥用的吖？别急，且听我慢慢道来。

咱先来说一说怎么来安装这个工具。想要安装 Vue Devtools 工具，可以访问 [github.com/vuejs/vue-d…]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fvuejs%2Fvue-devtools ) 地址，里面有比较详细的介绍。怎么滴呢？因为 Vue 官方已经把 Vue Devtools 工具开源并托管在全球最大同性社交网站 GitHub 上了。

![](https://user-gold-cdn.xitu.io/2019/6/3/16b1defe57011176?imageView2/0/w/1280/h/960/ignore-error/1)

地址告诉你了，再来说一说 Vue Devtools 工具是干啥用的吧。Vue Devtools 工具提供了一个友好的界面，在这个界面中可以审查和调试 Vue 代码。

说到这儿啊，咱得多说两句了。为啥呢？很多人对开发来说都存在着一个误区，这个误区就是认为程序员只要开发程序，敲代码就好了。其实不然，程序员几乎每天都要面对着一个问题，就是怎么解决各种各样的代码问题。这个时候，代码的审查和调试工具就尤为重要。

这么说吧！就像你吃饭，但也得把它们拉出去，这就说明筷子和马桶对你来说是同等的重要。当然了，这个比喻不是那么地恰当！你自己理解就好。

## 安装 Vue Devtools 工具 ##

咱们闲言少叙，书归正传。

接下来，咱们再来说一说怎么来安装 Vue Devtools 工具。第一种方式，是最简单最直接最暴力的。就是通过 Vue 官方提供的链接，直接安装对应的版本。

![](https://user-gold-cdn.xitu.io/2019/6/3/16b1df0658ba0477?imageView2/0/w/1280/h/960/ignore-error/1)

> 
> 
> 
> 这里需要说明一点的，就是 Chrome 浏览器的版本是直接访问 Chrome 应用商店的。为啥要单独说这个，因为 Chrome
> 应用商店在国内是被墙的，所以你懂的。
> 
> 

这个你是不是很想骂街，我第一次知道的时候，也想骂街。但是别急，咱还有 PlanB 呢！

就是通过 Git 工具把 Vue Devtools 工具的开源项目 clone 到本地，进行编译再自己安装到浏览器上。具体怎么做呢？

* 将 vue-devtools 项目 clone 到本地目录
` git clone https://github.com/vuejs/vue-devtools.git 复制代码` * 使用 npm 来安装所需要的所有包
` npm install 复制代码` * 编译项目的所有文件
` npm run build 复制代码` * 

把编译好的文件，安装到 Chrome 浏览器中

输入地址 进入扩展程序页面，点击“加载已解压的扩展程序”按钮。

![](https://user-gold-cdn.xitu.io/2019/6/3/16b1df307a361989?imageView2/0/w/1280/h/960/ignore-error/1)

裤裆里面着火，当然了！Vue Devtools 工具的开源项目中也提供一些其他方式的使用方式，有兴趣就自行研究吧！咱就不再这儿多费口舌了。

## Vue Devtools工具的注意事项 ##

说到这儿呢！基本上，你应该可以成功地安装上 Vue Devtools 工具了。什么？还没安装成功？！拉出去枪毙五分钟！

安装成功之后，咱们还得说一说 Vue Devtools 工具使用的一些注意事项。

第一呢，就是 Vue 核心库的文件类似于 jQuery，也是提供了两个文件，一个开发者版，一个生产版（压缩之后的）。如果你使用的是生产版本的 Vue 核心库文件的话，Vue Devtools 这个工具默认是被禁用的。换句话讲，你要想使用 Vue Devtools 工具的话，就得使用开发者版的 Vue 核心库文件。

为啥要这么做？这是为了当你使用 Vue 开发的应用正式上线之后，来保护你的核心代码逻辑的。

再有呢，要通过“file://”协议打开的 Vue 开发的网页，需要在 Chrome 浏览器的扩展程序管理面板中选中此扩展程序的“允许访问文件网址”。

好了，关于 Vue Devtools 工具咱们基本上算是介绍完了。Vue 官方赠送的这个橙色装备你接收了吗？

[下一回：Vue 的初阶黑魔法 — 模板语法]( https://juejin.im/post/5cf676a76fb9a07ee27b029e )