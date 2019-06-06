# Java RESTful Web Service落地 #

## 为什么要提出REST架构？ ##

Web Service协议栈的设计者没有充分认识到Web基础架构的巨大优点，甚至可以说并没有理解HTTP协议究竟是用来做什么的、为何要如此设计。在Web Service协议栈的设计之中，仍然有深深的企业应用痕迹。Web Servcie虽然宣称能够很好地支持交互操作，然而因为协议栈的复杂性很高，在实战中互操作性并不好（例如升级过程困难而且复杂）。此外，Web Service仅仅将HTTP协议当做一种传输协议来使用，还依赖XML这种冗余度很高的文本格式，这导致Web Service应用性能底下。很多开发团队宁可使用Hessian等轻量级的RPC协议，也不愿意使用Web Service。Java世界急需一套新的规范来取代JAX-WS。这套新的规范就是JAX-RS。——《Java RESTful Web Service实战》

## REST架构是如何推导的？ ##

[infoqstatic.com/resource/ar…]( https://link.juejin.im?target=http%3A%2F%2Finfoqstatic.com%2Fresource%2Farticles%2Funderstanding-restful-style%2Fzh%2Fresources%2Frestrest.zip )

![](https://user-gold-cdn.xitu.io/2019/6/5/16b27897f20d196f?imageView2/0/w/1280/h/960/ignore-error/1)

REST 架构风格最重要的架构约束有 6 个：

客户 - 服务器（Client-Server） 通信只能由客户端单方面发起，表现为请求 - 响应的形式。

无状态（Stateless） 通信的会话状态（Session State）应该全部由客户端负责维护。

缓存（Cache） 响应内容可以在通信链的某处被缓存，以改善网络效率。

统一接口（Uniform Interface） 通信链的组件之间通过统一的接口相互通信，以提高交互的可见性。

分层系统（Layered System） 通过限制组件的行为（即，每个组件只能“看到”与其交互的紧邻层），将架构分解为若干等级的层。

按需代码（Code-On-Demand，可选） 支持通过下载并执行一些代码（例如 Java Applet、Flash 或 JavaScript），对客户端的功能进行扩展。

## RPC架构和REST架构的区别是什么？ ##

REST 支持抽象的工具是资源，RPC 支持抽象的工具是过程。REST 风格的架构建模是以名词为核心的，RPC 风格的架构建模是以动词为核心的。

RPC 中没有统一接口的概念。不同的 API，接口设计风格可以完全不同。RPC 也不支持操作语义对于中间组件的可见性。

RPC 中没有使用超文本，响应的内容中只包含消息本身。REST 使用了超文本，可以实现更大粒度的交互，交互的效率比 RPC 更高。

REST 支持数据流和管道，RPC 不支持数据流和管道。

因为使用了平台中立的消息，RPC 风格的耦合度比 DO 风格要小一些，但是 RPC 风格也常常会带来客户端与服务器端的紧耦合。支持统一接口 + 超文本驱动的 REST 风格，可以达到最小的耦合度。

## REST架构落地的4个级别是什么？ ##

[martinfowler.com/articles/ri…]( https://link.juejin.im?target=https%3A%2F%2Fmartinfowler.com%2Farticles%2FrichardsonMaturityModel.html )

第一层次（Level 0）的 Web API 服务只是使用 HTTP 作为传输方式。

第二层次（Level 1）的 Web API 服务引入了资源的概念。每个资源有对应的标识符和表达。

第三层次（Level 2）的 Web API 服务使用不同的 HTTP 方法来进行不同的操作，并且使用 HTTP 状态码来表示不同的结果。

第四层次（Level 3）的 Web API 服务使用 HATEOAS。在资源的表达中包含了链接信息。客户端可以根据链接来发现可以执行的动作。

通常情况下，伪 RESTful API 都是基于第一层次与第二层次设计的。例如，我们的 Web API 中使用各种动词，例如 get_menu 和 save_menu ，而真正意义上的 RESTful API 需要满足第三层级以上。

![](https://user-gold-cdn.xitu.io/2019/6/5/16b277ce30f5654b?imageView2/0/w/1280/h/960/ignore-error/1)

## 亚马逊RESTful落地 ##

![](https://user-gold-cdn.xitu.io/2019/6/5/16b277834339f4e8?imageView2/0/w/1280/h/960/ignore-error/1)