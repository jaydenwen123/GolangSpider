# <link>标签的几个用法，帮助提高页面性能 #

#### 写在前面 ####

> 
> 
> 
> 本文首发于公众号：符合预期的CoyPan
> 
> 

**HTML** 中****元素规定了外部资源与当前文档的关系。最常见的用法，是用来链接一个外部的样式表，比如：

` < link href = "main.css" rel = "stylesheet" > 复制代码`

link标签还能做一些其他的事情，来帮助我们提高页面性能。

#### link标签的使用 ####

来看一下link标签除了链接外部样式表之外的一些使用场景。

##### DNS Prefetch #####

DNS预解析。

这个大多数人都知道，用法也很简单：

` < link rel = "dns-prefetch" href = "//example.com" > 复制代码`

DNS解析，简单来说就是把域名转化为ip地址。我们在网页里使用域名请求其他资源的时候，都会先被转化为ip地址，再发起链接。dns-prefeth使得转化工作提前进行了，缩短了请求资源的耗时。

什么时候使用呢？当我们页面中使用了其他域名的资源时，比如我们的静态资源都放在cdn上，那么我们可以对cdn的域名进行预解析。浏览器的支持情况也不错。

![](https://user-gold-cdn.xitu.io/2019/6/3/16b1d65ed4f70160?imageView2/0/w/1280/h/960/ignore-error/1)

##### Preconnect #####

预链接。

使用方法如下：

` < link rel = "preconnect" href = "//example.com" > < link rel = "preconnect" href = "//cdn.example.com" crossorigin > 复制代码`

我们访问一个站点时，简单来说，都会经过以下的步骤：

* DNS解析
* TCP握手
* 如果为Https站点，会进行TLS握手

使用preconnect后，浏览器会针对特定的域名，提前初始化链接(执行上述三个步骤)，节省了我们访问第三方资源的耗时。需要注意的是，我们一定要确保preconnect的站点是网页必需的，否则会浪费浏览器、网络资源。

浏览器的支持情况也不错：

![](https://user-gold-cdn.xitu.io/2019/6/3/16b1d6642ffd359a?imageView2/0/w/1280/h/960/ignore-error/1)

##### Prefetch #####

预拉取。

使用方法如下：

` < link rel = "prefetch" href = "//example.com/next-page.html" as = "document" crossorigin = "use-credentials" > < link rel = "prefetch" href = "/library.js" as = "script" > 复制代码`

link标签里的as参数可以有以下取值：

` audio: 音频文件 video: 视频文件 Track: 网络视频文本轨道 script: javascript文件 style: css样式文件 font: 字体文件 image: 图片 fetch: XHR、Fetch请求 worker: Web workers embed: 多媒体<embed>请求 object: 多媒体<object>请求 document : 网页 复制代码`

预拉取用于标识从当前网站跳转到下一个网站可能需要的资源，以及本网站应该获取的资源。这样可以在将来浏览器请求资源时提供更快的响应。

如果正确使用了预拉取，那么用户在从当前页面前往下一个页面时，可以很快得到响应。但是如果错误地使用了预拉取，那么浏览器就会下载额外不需要的资源，影响页面性能，并且造成网络资源浪费。

这里需要注意的是，使用了prefetch，资源仅仅被提前下载，下载后不会有任何操作，比如解析资源。

![](https://user-gold-cdn.xitu.io/2019/6/3/16b1d66837576523?imageView2/0/w/1280/h/960/ignore-error/1)

##### Prerender #####

预渲染。

` < link rel = "prerender" href = "//example.com/next-page.html" > 复制代码`

prerender比prefetch更进一步。不仅仅会下载对应的资源，还会对资源进行解析。解析过程中，如果需要其他的资源， **可能** 会直接下载这些资源。这样，用户在从当前页面跳转到目标页面时，浏览器可以更快的响应。

浏览器的支持情况如下：

![](https://user-gold-cdn.xitu.io/2019/6/3/16b1d66cdc58498e?imageView2/0/w/1280/h/960/ignore-error/1)

#### Resource Hints ####

上面的四种用法，其实就是： **Resource Hints** 。

Resource Hints ，翻译过来是【资源提示】。w3c的概括为：

> 
> 
> 
> This specification defines the ` dns-prefetch` (
> https://link.juejin.im?target=https%3A%2F%2Fwww.w3.org%2FTR%2Fresource-hints%2F%23dfn-dns-prefetch
> ) , ` preconnect` (
> https://link.juejin.im?target=https%3A%2F%2Fwww.w3.org%2FTR%2Fresource-hints%2F%23dfn-preconnect
> ) , ` prefetch` (
> https://link.juejin.im?target=https%3A%2F%2Fwww.w3.org%2FTR%2Fresource-hints%2F%23dfn-prefetch
> ) , and ` prerender` (
> https://link.juejin.im?target=https%3A%2F%2Fwww.w3.org%2FTR%2Fresource-hints%2F%23dfn-prerender
> ) relationships of the HTML Link Element ( ` <link>` ). These primitives
> enable the developer, and the server generating or delivering the
> resources, to assist the user agent in the decision process of which
> origins it should connect to, and which resources it should fetch and
> preprocess to improve page performance.
> 
> 

> 
> 
> 
> 此规范定义HTML链接元素（ ` <link>` ）的DNS预取、预连接、预取和预渲染关系。这些原语使开发人员和生成或传递资源的服务器能够帮助用户代理决定应该连接到哪个源，以及应该获取哪些资源，并进行预处理以提高页面性能。
> 
> 
> 

更多详细内容，可以在w3c的草案中查看： [www.w3.org/TR/resource…]( https://link.juejin.im?target=https%3A%2F%2Fwww.w3.org%2FTR%2Fresource-hints%2F )

####Resource Hints使用方法

除了上面介绍的使用link标签的使用方法，还可以直接通过http header的方式使用。例如可以使用下面的header:

` Link: < https: // widget.com > ; rel=dns-prefetch Link: < https: // example.com > ; rel=preconnect Link: < https: // example.com / next-page.html > ; rel=prerender; Link: < https: // example.com / logo-hires.jpg > ; rel=prefetch; as=image; 复制代码`

还可以在javascript使用：

` var hint = document.createElement( "link" ); hint.rel = "prefetch" ; hint.as = "document" ; hint.href = "/article/part3.html" ; document.head.appendChild(hint); 复制代码`

#### Resource Hints总结 ####

上文介绍了DNS Prefetch，Preconnect， Prefetch，Prerender。这四种hint的功能逐渐递进：

* 

Dns Prefetch进行DNS预查询。

* 

Preconnect进行预链接。在一些重定向技术中，Preconnect可以让浏览器和最终目标源更早建立连接。

* 

Prefetch进行预下载。比如说，我们可以根据用户行为猜测其下一步操作，然后动态预获取所需资源，并且不用担心该资源被解析(执行)而影响页面当前功能。

* 

Prerender不仅仅提前下载资源，还会提前直接解析(执行)资源。如果我们对下一个页面进行Prerender，用户在打开下一个页面时，就会感觉很流畅了。

需要注意的是，浏览器对于Resource Hints的实现并不是想象中的那样简单直接。Resource Hints只是一些『提示』，浏览器可以采用我们的提示，但是具体怎么实现还是由浏览器自己来决定的。比如，如果当前CPU压力大，网络阻塞时，你使用了Prefetch，那么浏览器可能仅仅会只对dns进行预解析，并不会下载资源。

#### 写在后面 ####

本文介绍了link标签的四种使用方法，最终引出了Resource Hints的概念。Resource Hints可以帮助我们提高页面的性能。但是这只是理论上的，真正的收益还需要在实际业务中去探索、验证。

符合预期。