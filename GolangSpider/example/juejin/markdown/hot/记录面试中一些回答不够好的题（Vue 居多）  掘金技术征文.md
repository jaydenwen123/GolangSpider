# 记录面试中一些回答不够好的题（Vue 居多） | 掘金技术征文 #

## 相关问题 ##

* flex 布局 与 grid 布局。
* 实现 Vue SSR 。
* 从 SPA 使用最小成本迁移到 SSR 。
* 实现方法： （未完成） 根据指定元素，在数组里面找出 ff 数组（ff 数组这个名字是我瞎说的）。比如数组 [2, 3, 6, 7] ，指定元素 7，则 ff 数组是 [2, 2, 3]（2+2+3 = 7）和 [7]。若指定元素 6，则 ff 数组为 [2, 2, 2], [3, 3], 和 [6] 。
* 实现 Promise.finally。
* 另一种方式实现 Vue 的响应式原理。
* Vue 组件 data 为什么必须是函数。
* Vue computed 实现。
* diff 算法实现。
* Vue complier 实现。
* 快排及其优化。
* 缓存算法实现及其优化（缓存算法简单模型：假设可以缓存三个数据，请求前三个数据时，直接进缓存列表，当请求第四个数据时，若命中缓存，将被缓存的数据放入缓存列表头部，否则把新加入的数据放入缓存列表头部，淘汰最后一个数据）。
* 怎么快速定位哪个组件出现性能问题。
* http 状态码 202, 204 。
* WebSocket 。
* 尽可能多的说出你对 Electron 的理解。

## 相关解答 ##

### flex 布局 与 grid 布局 ###

这个问题比较简单，用 flex 与 grid 实现如下即可：

![](https://user-gold-cdn.xitu.io/2018/3/4/161ef8a6b83341fb?imageView2/0/w/1280/h/960/ignore-error/1)

实现方式如下：

` < html > < head > < style > /* flex */.box { display : flex; flex-wrap : wrap; width : 100% ; }.box div { width : calc (100% / 3 - 2px); height : 100px ; border : 1px solid black; } /* grid */.box { display : grid; grid-template-columns : 1 fr 1 fr 1 fr; width : 100% ; }.box div { height : 100px ; border : 1px solid black; } </ style > < head > < body > < div class = "box" > < div > </ div > < div > </ div > < div > </ div > < div > </ div > </ div > < body > </ html > 复制代码`

grid 学习：https://www.jianshu.com/p/d183265a8dad

### 实现 Vue SSR ###

一些想法写在下题。

### 从 SPA 使用最小成本迁移到 SSR ###

Vue SSR 的好处就不多说了，这有一篇相关文章 [服务端渲染与客户端渲染]( https://link.juejin.im?target=https%3A%2F%2Fjkchao.cn%2Farticle%2F5a11155fb520d115154c8fa1 ) 。 简单的总结下 Vue SSR 的实现。 有一张实现图：

![](https://user-gold-cdn.xitu.io/2018/3/4/161ef7bf329e8812?imageView2/0/w/1280/h/960/ignore-error/1)

其基本实现原理：

* app.js 作为客户端与服务端的公用入口，导出 Vue 根实例，供客户端 entry 与服务端 entry 使用。客户端 entry 主要作用挂载到 DOM 上，服务端 entry 除了创建和返回实例，还进行路由匹配与数据预获取。
* webpack 为客服端打包一个 Client Bundle ，为服务端打包一个 Server Bundle 。
* 服务器接收请求时，会根据 url，加载相应组件，获取和解析异步数据，创建一个读取 Server Bundle 的 BundleRenderer，然后生成 html 发送给客户端。
* 客户端混合，客户端收到从服务端传来的 DOM 与自己的生成的 DOM 进行对比，把不相同的 DOM 激活，使其可以能够响应后续变化，这个过程称为 [客户端激活 ]( https://link.juejin.im?target=https%3A%2F%2Fssr.vuejs.org%2Fzh%2Fhydration.html ) 。为确保混合成功，客户端与服务器端需要共享同一套数据。在服务端，可以在渲染之前获取数据，填充到 stroe 里，这样，在客户端挂载到 DOM 之前，可以直接从 store 里取数据。首屏的动态数据通过 ` window.__INITIAL_STATE__` 发送到客户端。

Vue SSR 的实现，主要就是把 Vue 的组件输出成一个完整 HTML, [vue-server-renderer]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fvuejs%2Fvue%2Ftree%2Fdev%2Fpackages%2Fvue-server-renderer ) 就是干这事的。

纯客户端输出过程有一个 complier 过程（「 [下题]( #header-11 ) 」中有一个简单描述），主要作用是将 template 转化成 render 字符串 。

Vue SSR 需要做的事多点（输出完整 HTML），除了 complier -> vnode，还需如数据获取填充至 HTML、客户端混合（hydration）、缓存等等。

相比于其他模板引擎（ejs, jade 等），最终要实现的目的是一样的，性能上可能要差点。

参考：

* https://ssr.vuejs.org/zh/
* https://segmentfault.com/a/1190000006701796

### ff 数组 ###

### 实现 Promise.finally ###

finally 方法用于指定不管 Promise 对象最后状态如何，都会执行的操作，使用方法如下：

` Promise.then( result => { ··· }) .catch( error => { ··· }) .finally( () => { ··· }) 复制代码`

finally 特点：

* 不接收任何参数。
* finally 本质上是 then 方法的特例。

` Promise.prototype.finally = function ( callback ) { let P = this.constructor return this.then( value => P.resolve(callback()).then( () => value), reason => P.resolve(callback()).then( () => { throw reason }) ) } 复制代码`

### 另一种方式实现 Vue 的响应式原理 ###

Vue 的响应式原理是使用 Object.defineProperty 追踪依赖，当属性被访问或改变时通知变化。

有两个不足之处：

* 不能检测到增加或删除的属性。
* 数组方面的变动，如根据索引改变元素，以及直接改变数组长度时的变化，不能被检测到。

原因差不多，无非就是没有被 getter/setter 。

第一个比较容易理解，为什么数组长度不能被 getter/setter ？

在知乎上找了一个答案：如果你知道数组的长度，理论上是可以预先给所有的索引设置 getter/setter 的。但是一来很多场景下你不知道数组的长度，二来，如果是很大的数组，预先加 getter/setter 性能负担较大。

现在有一个替代的方案 Proxy，但这东西兼容性不好，迟早要上的。

Proxy，在目标对象之前架设一层拦截。具体，可以参考 http://es6.ruanyifeng.com/#docs/reference

### Vue 组件 data 为什么必须是函数 ###

理解两点：

* 每个组件都是 Vue 的实例。
* 组件共享 data 属性，当 data 的值是同一个引用类型的值时，改变其中一个会影响其他。

### Vue computed 实现 ###

> 
> 
> 
> 这个题目有两家问了，感觉都不是答得很好。
> 
> 

从两个问题出发：

* 建立与其他属性（如：data、 Store）的联系；
* 属性改变后，通知计算属性重新计算。

实现时，主要如下

* 初始化 data， 使用 Object.defineProperty 把这些属性全部转为 getter/setter。
* 初始化 computed, 遍历 computed 里的每个属性，每个 computed 属性都是一个 watch 实例。每个属性提供的函数作为属性的 getter，使用 Object.defineProperty 转化。
* Object.defineProperty getter 依赖收集。用于依赖发生变化时，触发属性重新计算。
* 若出现当前 computed 计算属性嵌套其他 computed 计算属性时，先进行其他的依赖收集。

参考：https://segmentfault.com/a/1190000010408657

### diff 算法实现 ###

> 
> 
> 
> 以前写过两篇文章讨论这个算法的实现，没想到过的太久，忘记了。（文章地址：https://github.com/jkchao/blog/issues/3
> ，https://github.com/jkchao/blog/issues/4） 。 也好，称此机会总结下
> 
> 

diff 的实现主要通过两个方法，patchVnode 与 updateChildren 。

patchVnode 有两个参数，分别是老节点 oldVnode, 新节点 vnode 。主要分五种情况：

* if (oldVnode === vnode)，他们的引用一致，可以认为没有变化。
* if(oldVnode.text !== null && vnode.text !== null && oldVnode.text !== vnode.text)，文本节点的比较，需要修改，则会调用Node.textContent = vnode.text。
* if( oldCh && ch && oldCh !== ch ), 两个节点都有子节点，而且它们不一样，这样我们会调用 updateChildren 函数比较子节点，这是diff的核心，后边会讲到。
* if (ch)，只有新的节点有子节点，调用createEle(vnode)，vnode.el已经引用了老的dom节点，createEle函数会在老dom节点上添加子节点。
* if (oldCh)，新节点没有子节点，老节点有子节点，直接删除老节点。

updateChildren 是关键，这个过程可以概括如下：

![](https://user-gold-cdn.xitu.io/2018/3/4/161ef7bf338881f1?imageView2/0/w/1280/h/960/ignore-error/1)

oldCh 和 newCh 各有两个头尾的变量 StartIdx 和 EndIdx ，它们的2个变量相互比较，一共有4种比较方式。如果 4 种比较都没匹配，如果设置了key，就会用key进行比较，在比较的过程中，变量会往中间靠，一旦 StartIdx > EndIdx 表明 oldCh 和 newCh 至少有一个已经遍历完了，就会结束比较。

### Vue complier 实现 ###

> 
> 
> 
> 以前写过一篇 「 [Vue 生面周期总结的文章](
> https://link.juejin.im?target=https%3A%2F%2Fjkchao.cn%2Farticle%2F59d6e93c7e2ee06d412efef9
> ) 」的文章，里面提到了 complier 的作用，没有做深入了解。。。
> 
> 

模板解析这种事，本质是将数据转化为一段 html ，最开始出现在后端，经过各种处理吐给前端。随着各种 mv* 的兴起，模板解析交由前端处理。 总的来说，Vue complier 是将 template 转化成一个 render 字符串。 可以简单理解成以下步骤：

* parse 过程，将 template 利用正则转化成 AST 抽象语法树。
* optimize 过程，标记静态节点，后 diff 过程跳过静态节点，提升性能。
* generate 过程，生成 render 字符串。

参考：

* https://segmentfault.com/a/1190000006990480
* https://github.com/answershuto/learnVue/blob/master/docs/%E8%81%8A%E8%81%8AVue%E7%9A%84template%E7%BC%96%E8%AF%91.MarkDown

### 快排及其优化 ###

> 
> 
> 
> 前端对算法的要求还是比较低的，但也是必不可少的一部分。
> 
> 

找到一篇比较不错的文章：https://www.cnblogs.com/zichi/p/4788953.html

### 缓存算法实现及其优化 ###

最简单的一种思路就是使用数组存储，然后让我优化。 我。。。一脸懵逼。 有兴趣的同学可以参考这个： http://www.cnblogs.com/dolphin0520/p/3749259.html 。

ps: 看来我得补补数据结构和算法相关的知识了。

### 怎么快速定位哪个组件出现性能问题 ###

当面试官问这个问题，没有 get 到面试官的点，扯了一堆乱七八糟没用的 - -。 后来面试官说主要是用 timeline 工具。 大意是通过 timeline 来查看每个函数的调用时常，定位出哪个函数的问题，从而能判断哪个组件出了问题。

附上两个使用 timeline 的文章：

* https://juejin.im/post/5a6e78abf265da3e3f4cf085
* https://developers.google.cn/web/tools/chrome-devtools/?hl=zh-cn

### http 状态码 202, 204 ###

> 
> 
> 
> 面试官不知道为何扯到了 202, 204。。。好像是由自己带进坑的。- -
> 
> 

202: 服务器已接受请求，但尚未处理。 204: 服务器成功处理了请求，没有返回任何内容。

这些状态码感觉只要能记住常用的就 ok 了，当然还得了解 200 +, 300+, 400+, 500+ 代表什么意思。

### WebSocket ###

> 
> 
> 
> WebSocket 应该算是一个比较常问的面试点，如果问的不深的话，应该比较好回答。
> 
> 

由于 http 存在一个明显的弊端（消息只能有客户端推送到服务器端，而服务器端不能主动推送到客户端），导致如果服务器如果有连续的变化，这时只能使用轮询，而轮询效率过低，并不适合。于是 WebSocket 被发明出来。

相比与 http 具有以下有点：

* 支持双向通信，实时性更强；
* 可以发送文本，也可以二进制文件；
* 协议标识符是 ws，加密后是 wss ；
* 较少的控制开销。连接创建后，ws客户端、服务端进行数据交换时，协议控制的数据包头部较小。在不包含头部的情况下，服务端到客户端的包头只有2~10字节（取决于数据包长度），客户端到服务端的的话，需要加上额外的4字节的掩码。而HTTP协议每次通信都需要携带完整的头部；
* 支持扩展。ws协议定义了扩展，用户可以扩展协议，或者实现自定义的子协议。（比如支持自定义压缩算法等）
* 无跨域问题。

实现比较简单，服务端库如 ` socket.io` 、 ` ws` ，可以很好的帮助我们入门。而客户端也只需要参照 api 实现即可。

参考：

* http://www.ruanyifeng.com/blog/2017/05/websocket.html
* https://www.cnblogs.com/chyingp/p/websocket-deep-in.html

### 尽可能多的说出你对 Electron 的理解 ###

> 
> 
> 
> 以前写过一篇简单的关于 [electron-vue](
> https://link.juejin.im?target=https%3A%2F%2Fjkchao.cn%2Farticle%2F59f721eb836b695d8fffd53d
> ) 的文章，没想到真有面试官问，而且问的挺深的。
> 
> 

最最重要的一点，electron 实际上是一个套了 Chrome 的 node 程序。

所以应该是从两个方面说开来：

* Chrome （无各种兼容性问题）；
* Node （Node 能做的它也能做）。

Chrome 没什么好说的，是个前端都懂。

Node 方面可说的就多了。

有个面试官问我，在 electron 怎么解决跨域问题？

在我自己的项目里，确实遇到了这个问题，可惜选择了一个不怎么好的方法的方法，设置 nginx 。

为什么不好，如果项目是公司的，还需要运维同学帮忙。- -

也聊到了使用 CORS 允许跨域，也觉得不好，因为需要后端接口处理。 一脸懵逼的我，直到面试官提醒使用 node 来代理以下，才恍然大悟。（原来还可以这种操作。。。。）

当然也可以连接数据库，上家公司本来打算要做一个 electron 配合连接数据库的桌面应用。（还没开始做就离职了- -） 挺可惜的，当时数据库都已经选择好了，leveldb 或者 lowdb ，觉得应该不难。

附上两个 electron 配合数据库使用的链接：

* https://github.com/typicode/lowdb/issues/169
* https://github.com/Level/electron-demo

功力不足，难免有错误之处，还望多多指出。

掘金技术证文活动链接: https://juejin.im/post/5aaf2a95f265da239b413aa1