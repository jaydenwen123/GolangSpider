# [译] HTTP/2 常见问题解答 #

> 
> 
> 
> * 原文地址： [HTTP/2 Frequently Asked Questions](
> https://link.juejin.im?target=https%3A%2F%2Fhttp2.github.io%2Ffaq%2F )
> * 原文作者： [HTTP/2](
> https://link.juejin.im?target=https%3A%2F%2Fhttp2.github.io )
> * 译文出自： [掘金翻译计划](
> https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fxitu%2Fgold-miner
> )
> * 本文永久链接： [github.com/xitu/gold-m…](
> https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fxitu%2Fgold-miner%2Fblob%2Fmaster%2FTODO1%2Fhttp-2-frequently-asked-questions.md
> )
> * 译者： [YueYong](
> https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2FYueYongDev )
> * 校对者： [Ranjay](
> https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2FjerryOnlyZRJ ) , [ziyin
> feng](
> https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2FFengziyin1234 )
> 
> 
> 

以下是有关 HTTP/2 的常见问题解答。

* [一般问题]( #%E4%B8%80%E8%88%AC%E9%97%AE%E9%A2%98 )

* [为什么要修订 HTTP ?]( #%E4%B8%BA%E4%BB%80%E4%B9%88%E8%A6%81%E4%BF%AE%E8%AE%A2-http )
* [谁制定了 HTTP/2??]( #%E8%B0%81%E5%88%B6%E5%AE%9A%E4%BA%86-http2 )
* [HTTP/2 与 SPDY 的关系是什么？]( #http2-%E4%B8%8E-spdy-%E7%9A%84%E5%85%B3%E7%B3%BB%E6%98%AF%E4%BB%80%E4%B9%88 )
* [究竟是 HTTP/2.0 还是 HTTP/2？]( #%E7%A9%B6%E7%AB%9F%E6%98%AF-http20-%E8%BF%98%E6%98%AF-http2 )
* [和 HTTP/1.x 相比 HTTP/2 的关键区别是什么?]( #%E5%92%8C-http1x-%E7%9B%B8%E6%AF%94-http2-%E7%9A%84%E5%85%B3%E9%94%AE%E5%8C%BA%E5%88%AB%E6%98%AF%E4%BB%80%E4%B9%88 )
* [为什么 HTTP/2 是二进制的?]( #%E4%B8%BA%E4%BB%80%E4%B9%88-http2-%E6%98%AF%E4%BA%8C%E8%BF%9B%E5%88%B6%E7%9A%84 )
* [为什么 HTTP/2 需要多路传输?]( #%E4%B8%BA%E4%BB%80%E4%B9%88-http2-%E9%9C%80%E8%A6%81%E5%A4%9A%E8%B7%AF%E4%BC%A0%E8%BE%93 )
* [为什么只需要一个 TCP 连接?]( #%E4%B8%BA%E4%BB%80%E4%B9%88%E5%8F%AA%E9%9C%80%E8%A6%81%E4%B8%80%E4%B8%AA-TCP-%E8%BF%9E%E6%8E%A5 )
* [服务器推送的好处是什么？]( #%E6%9C%8D%E5%8A%A1%E5%99%A8%E6%8E%A8%E9%80%81%E7%9A%84%E5%A5%BD%E5%A4%84%E6%98%AF%E4%BB%80%E4%B9%88 )
* [消息头为何需要压缩？]( #%E6%B6%88%E6%81%AF%E5%A4%B4%E4%B8%BA%E4%BD%95%E9%9C%80%E8%A6%81%E5%8E%8B%E7%BC%A9 )
* [为什么选择 HPACK？]( #%E4%B8%BA%E4%BB%80%E4%B9%88%E9%80%89%E6%8B%A9-HPACK )
* [HTTP/2 可以让 cookies （或者其他消息头）变得更好吗？]( #http2-%E5%8F%AF%E4%BB%A5%E8%AE%A9-cookies-%E6%88%96%E8%80%85%E5%85%B6%E4%BB%96%E6%B6%88%E6%81%AF%E5%A4%B4%E5%8F%98%E5%BE%97%E6%9B%B4%E5%A5%BD%E5%90%97 )
* [非浏览器用户的 HTTP 是什么样的？]( #%E9%9D%9E%E6%B5%8F%E8%A7%88%E5%99%A8%E7%94%A8%E6%88%B7%E7%9A%84-HTTP-%E6%98%AF%E4%BB%80%E4%B9%88%E6%A0%B7%E7%9A%84 )
* [HTTP/2 需要加密吗？]( #http2-%E9%9C%80%E8%A6%81%E5%8A%A0%E5%AF%86%E5%90%97 )
* [HTTP/2 是怎么提高安全性的呢？]( #Hhttp2-%E6%98%AF%E6%80%8E%E4%B9%88%E6%8F%90%E9%AB%98%E5%AE%89%E5%85%A8%E6%80%A7%E7%9A%84%E5%91%A2 )
* [我现在可以使用 HTTP/2 吗？]( #%E6%88%91%E7%8E%B0%E5%9C%A8%E5%8F%AF%E4%BB%A5%E4%BD%BF%E7%94%A8-http2-%E5%90%97 )
* [HTTP/2 将会取代 HTTP/1.x 吗？]( #http2-%E5%B0%86%E4%BC%9A%E5%8F%96%E4%BB%A3-http1x-%E5%90%97 )
* [HTTP/3 会出现吗？]( #http3-%E4%BC%9A%E5%87%BA%E7%8E%B0%E5%90%97 )

* [实现过程中的问题]( #%E5%AE%9E%E7%8E%B0%E8%BF%87%E7%A8%8B%E4%B8%AD%E7%9A%84%E9%97%AE%E9%A2%98 )

* [为什么规则会围绕消息头帧的数据接续？]( #%E4%B8%BA%E4%BB%80%E4%B9%88%E8%A7%84%E5%88%99%E4%BC%9A%E5%9B%B4%E7%BB%95%E6%B6%88%E6%81%AF%E5%A4%B4%E5%B8%A7%E7%9A%84%E6%95%B0%E6%8D%AE%E6%8E%A5%E7%BB%AD )
* [HPACK 状态的最小和最大尺寸是多少？]( #hpack-%E7%8A%B6%E6%80%81%E7%9A%84%E6%9C%80%E5%B0%8F%E5%92%8C%E6%9C%80%E5%A4%A7%E5%B0%BA%E5%AF%B8%E6%98%AF%E5%A4%9A%E5%B0%91 )
* [我怎样才能避免保持 HPACK 状态？]( #%E6%88%91%E6%80%8E%E6%A0%B7%E6%89%8D%E8%83%BD%E9%81%BF%E5%85%8D%E4%BF%9D%E6%8C%81-hpack-%E7%8A%B6%E6%80%81 )
* [为什么会有一个单独的压缩/流程控制上下文？]( #%E4%B8%BA%E4%BB%80%E4%B9%88%E4%BC%9A%E6%9C%89%E4%B8%80%E4%B8%AA%E5%8D%95%E7%8B%AC%E7%9A%84%E5%8E%8B%E7%BC%A9%E6%B5%81%E7%A8%8B%E6%8E%A7%E5%88%B6%E4%B8%8A%E4%B8%8B%E6%96%87 )
* [为什么在 HPACK 中有 EOS 的符号？]( #%E4%B8%BA%E4%BB%80%E4%B9%88%E5%9C%A8-hpack-%E4%B8%AD%E6%9C%89-eos-%E7%9A%84%E7%AC%A6%E5%8F%B7 )
* [实现 HTTP/2 的时候我可以不用去实现 HTTP/1.1 吗？]( #%E5%AE%9E%E7%8E%B0-http2-%E7%9A%84%E6%97%B6%E5%80%99%E6%88%91%E5%8F%AF%E4%BB%A5%E4%B8%8D%E7%94%A8%E5%8E%BB%E5%AE%9E%E7%8E%B0-http11-%E5%90%97 )
* [5.3.2节中的优先级示例是否正确？]( #532%E8%8A%82%E4%B8%AD%E7%9A%84%E4%BC%98%E5%85%88%E7%BA%A7%E7%A4%BA%E4%BE%8B%E6%98%AF%E5%90%A6%E6%AD%A3%E7%A1%AE )
* [HTTP/2 连接中需要 TCP_NODELAY 吗？]( #http2-%E8%BF%9E%E6%8E%A5%E4%B8%AD%E9%9C%80%E8%A6%81-tcp_nodelay-%E5%90%97 )

* [部署问题]( #%E9%83%A8%E7%BD%B2%E9%97%AE%E9%A2%98 )

* [我该怎么调试加密过的 HTTP/2 ？]( #%E6%88%91%E8%AF%A5%E6%80%8E%E4%B9%88%E8%B0%83%E8%AF%95%E5%8A%A0%E5%AF%86%E8%BF%87%E7%9A%84-http2 )
* [我该怎么使用 HTTP/2 的服务端推送？]( #%E6%88%91%E8%AF%A5%E6%80%8E%E4%B9%88%E4%BD%BF%E7%94%A8-http2-%E7%9A%84%E6%9C%8D%E5%8A%A1%E7%AB%AF%E6%8E%A8%E9%80%81 )

## 一般问题 ##

### 为什么要修订 HTTP? ###

HTTP/1.1 已经在 Web 上服役了十五年以上，但其劣势也开始显现。

加载一个网页比以往更加耗费资源（详见 [HTTP Archive’s page size statistics]( https://link.juejin.im?target=http%3A%2F%2Fhttparchive.org%2Ftrends.php%23bytesTotal%26amp%3BreqTotal ) ）。与此同时，有效地加载所有这些静态资源变得非常困难，因为事实上，HTTP 只允许每个 TCP 连接有一个未完成的请求。

在过去，浏览器使用多个 TCP 连接来发出并行请求。然而这种做法是有限制的。如果使用了太多的连接，就会产生相反的效果（TCP 拥塞控制将被无效化，导致的用塞事件将会损害性能和网络）。而且从根本上讲这对其他程序来说也是不公平的(因为浏览器会占用许多本不该属于他的资源)。

同时，大量的请求意味着“在线上”有大量重复的数据。

这两个因素都意味着 HTTP/1.1 请求有很多与之相关的开销;如果请求太多，则会影响性能。

这使得业界在有哪些是最好的实践上达成共识，它们包括，比如，Spriting（图片合并）、data: inlining（数据内嵌）、Domain Sharding（域名分片）和 Concatenation（文件合并）等。这些不规范的解决方案说明了协议本身存在一些潜在问题，并且在使用的时候会出现很多问题。

### 谁制定了 HTTP/2? ###

HTTP/2 是由 [IETF]( https://link.juejin.im?target=http%3A%2F%2Fwww.ietf.org%2F ) 的 [HTTP 工作组]( https://link.juejin.im?target=https%3A%2F%2Fhttpwg.github.io%2F ) 开发的，该组织负责维护 HTTP 协议。该组织由众多 HTTP 实现者、用户、网络运营商和 HTTP 专家组成。

值得注意的是，虽然 [工作组的邮件列表]( https://link.juejin.im?target=http%3A%2F%2Flists.w3.org%2FArchives%2FPublic%2Fietf-http-wg%2F ) 托管在 W3C 网站上，不过这并不是 W3C 的功劳。但是， Tim Berners-Lee 和 W3C TAG与 WG 的进度保持一致。

许多人为这项工作做出了自己的贡献，尤其是一些来自“大”项目的工程师，例如 Firefox、Chrome、Twitter、Microsoft 的 HTTP stack、Curl 和 Akamai。以及若干 Python、Ruby 和 NodeJS 的 HTTP 实现者。

为了更好的了解有关 IETF 的信息，你可以访问 [Tao of the IETF]( https://link.juejin.im?target=http%3A%2F%2Fwww.ietf.org%2Ftao.html ) ；你也可以在 [Github 的贡献者图表]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fhttp2%2Fhttp2-spec%2Fgraphs%2Fcontributors ) 上查看有哪些人为该项目做出了贡献，同样的，你也可以在 [implementation list]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fhttp2%2Fhttp2-spec%2Fwiki%2FImplementations ) 上查看谁正在参与该项目。

### HTTP/2 与 SPDY 的关系是什么？ ###

HTTP/2 第一次出现并被讨论的时候，SPDY 正得到厂商 (像 Mozilla 和 nginx)的青睐和支持，并被看成是 HTTP/1.x 基础上的重大改善。

在不断的征求建议以及投票选择之后， [SPDY/2]( https://link.juejin.im?target=http%3A%2F%2Ftools.ietf.org%2Fhtml%2Fdraft-mbelshe-httpbis-spdy-00 ) 被选为 HTTP/2 的基础。从那时起，根据工作组的讨论和用户的反馈，它已经有了很多变化。

在整个过程中，SPDY 的核心开发人员参与了 HTTP/2 的开发，其中包括 Mike Belshe 和 Roberto Peon。

2015 年 2 月，谷歌 [宣布计划]( https://link.juejin.im?target=https%3A%2F%2Fblog.chromium.org%2F2015%2F02%2Fhello-http2-goodbye-spdy.html ) 取消对 SPDY 的支持，转而支持 HTTP/2。

### 究竟是 HTTP/2.0 还是 HTTP/2？ ###

工作组决定删除次要版本（“.0”），因为它在 HTTP/1.x 中造成了很多混乱。也就是说，HTTP 的版本仅代表它的兼容性，不表示它的特性和“亮点”。

### 和 HTTP/1.x 相比 HTTP/2 的关键区别是什么? ###

在高版本 HTTP/2 中：

* 是二进制的，代替原有的文本
* 是多路复用的，代替原来的序列和阻塞机制
* 所以可以在一个连接中并行处理
* 压缩头部信息减小开销
* 允许服务器主动推送应答到客户端的缓存中

### 为什么 HTTP/2 是二进制的? ###

和 HTTP/1.x 这样的文本协议相比，二进制协议解析起来更高效、“线上”更紧凑，更重要的是错误更少。因为它们对如空白字符的处理、大小写、行尾、空链接等的处理很有帮助。

举个栗子 🌰，HTTP/1.1 定义了 [四种不同的方法来解析一条消息]( https://link.juejin.im?target=http%3A%2F%2Fwww.w3.org%2FProtocols%2Frfc2616%2Frfc2616-sec4.html%23sec4.4 ) ；而在HTTP/2中，仅需一个代码路径即可。

HTTP/2 在 telnet 中不可用，但是我们已经有一些工具可以提供支持，例如 [Wireshark plugin]( https://link.juejin.im?target=https%3A%2F%2Fbugs.wireshark.org%2Fbugzilla%2Fshow_bug.cgi%3Fid%3D9042 ) 。

### 为什么 HTTP/2 需要多路传输？ ###

HTTP/1.x 有个问题叫“队头阻塞（head-of-line blocking）”，它是指在一次连接（connection）中，只提交一个请求的效率比较高，多了就会变慢。

HTTP/1.1 尝试使用管线化（pipelining）来解决这个问题，但是效果并不理想（对于数据量较大或者速度较慢的响应，依旧会阻碍排在他后面的请求）。此外，由于许多网络媒介（intermediary）和服务器不能很好的支持管线化，导致其部署起来也是困难重重。

这也就迫使客户端使用一些启发式的方法（基本靠猜）来决定通过哪些连接提交哪些请求；由于一个页面加载的数据量，往往比可用连接能处理的数据量的 10 倍还多，对性能产生极大的负面影响，结果经常引起瀑布式阻塞（waterfall of blocked requests）。

而多路传输（Multiplexing）能很好的解决这些问题，因为它能同时处理多个消息的请求和响应；甚至可以在传输过程中将一个消息跟另外一个掺杂在一起。

所以在这种情况下，客户端只需要一个连接就能加载一个页面。

### 为什么只需要一个 TCP 连接？ ###

如果使用 HTTP/1，浏览器打开每个点（origin）就需要 4 到 8 个连接（Connection）。而现在很多网站都使用多点传输（multiple origins），也就是说，光加载一个网页，打开的连接数量就超过 30 个。

一个应用同时打开这么多连接，已经远远超出了当初设计 TCP 时的预想；同时，因为每个连接都会响应大量的数据，使其可以造成网络缓存溢出的风险，结果可能导致网络堵塞和数据重传。

此外，使用这么多连接还会强占许多网络资源。这些资源都是从那些“遵纪守法”的应用那“偷”的（VoIP 就是个很好的例子）。

### 服务器推送的好处是什么？ ###

当浏览器请求页面时，服务器发送 HTML 作为响应，然后需要等待浏览器解析 HTML 并发出对所有嵌入资源的请求，然后才能开始发送 JavaScript，图像和 CSS。

服务器推送服务通过“推送”那些它认为客户端将会需要的内容到客户端的缓存中，以此来避免往返的延迟。

但是，推送的响应并不是“万金油”，如果使用不当，可能会损害性能。正确使用服务器推送是一个长期的实验及研究领域。

### 消息头为何需要压缩？ ###

来自 Mozilla 的 Patrick McManus 通过计算消息头对平均页面负载的印象，对此进行了形象且充分的说明。

假定一个页面有 80 个资源需要加载（这个数量对于今天的 Web 而言还是挺保守的），而每一次请求都有 1400 字节的消息头（这同样也并不少见，因为 Cookie 和引用等东西的存在），至少要 7 到 8 个来回去“在线”获得这些消息头。这还不包括响应时间——那只是从客户端那里获取到它们所花的时间而已。

这全都由于 TCP 的 [慢启动]( https://link.juejin.im?target=http%3A%2F%2Fen.wikipedia.org%2Fwiki%2FSlow-start ) 机制，它根据可以确认的数据包数量对新连接上发送数据的进行限制 — 这有效地限制了最初的几次来回可以发送的数据包数量。

相比之下，即使是头部轻微的压缩也可以是让那些请求只需一个来回就能搞定——有时候甚至一个包就可以了。

这些额外的开销是相当多的，特别是当你考虑对移动客户端的影响的时候。这些往返的延迟，即使在网络状况良好的情况下，也高达数百毫秒。

### 为什么选择 HPACK？ ###

SPDY/2 提出在每一方都使用一个单独的 GZIP 上下文用于消息头压缩，这实现起来很容易，也很高效。

从那时起，一个重要的攻击方式 [CRIME]( https://link.juejin.im?target=http%3A%2F%2Fen.wikipedia.org%2Fwiki%2FCRIME ) 诞生了，这种方式可以攻击加密文件内部的所使用的压缩流（如GZIP）。

使用 CRIME，那些具备向加密数据流中注入数据能力的攻击者获得了“探测”明文并还原的可能性。因为是 Web，JavaScript 使其成为了可能，而且已经有了通过对受到 TLS 保护的 HTTP 资源的使用CRIME来还原出 cookies 和认证令牌（Toekn）的案例。

因此，我们不应该使用 GZIP 进行压缩。由于找不到其它适合在这种用例下使用的安全有效的算法，所以我们创造了一种新的，针对消息头的，进行粗粒度操作的压缩模式；因为HTTP消息头并不常常需要改变，我们仍然可以得到很好的压缩效率，而且更加的安全。

### HTTP/2 可以让 cookies（或者其他消息头）变得更好吗？ ###

这一努力被许可在网络协议的一个修订版本上运行 – 例如，HTTP 消息头、方法等等如何才能在不改变 HTTP 语义的前提下放到“网络上”。

这是因为 HTTP 的应用非常广泛。如果我们使用了这个版本的 HTTP，它就会引入一种新的状态机制（例如之前讨论过的例子）或者改变其核心方法（幸好，这还没有发生过），这可能就意味着新的协议将不会兼容现有的 Web 内容。

具体地，我们是想要能够从 HTTP/1 转移到 HTTP/2，并且不会有信息的丢失。如果我们开始”清理”消息头（大多数人都认为现在的 HTTP 消息头简直是一团糟)，我们就不得不去面对现有 Web 的诸多问题。

这样做只会对新协议的普及造成麻烦。

总而言之， [工作组]( https://link.juejin.im?target=https%3A%2F%2Fhttpwg.github.io%2F ) 会对所有的 HTTP 负责，而不仅仅只是 HTTP/2。因此，我们才可以在版本独立的新机制下运作，只要它们也能同现有的网络向下兼容。

### 非浏览器用户的 HTTP 是什么样的？ ###

如果非浏览器应用已经使用过 HTTP 的话，那他们也应该可以使用 HTTP/2。

先前收到过 HTTP “APIs” 在 HTTP/2 中具有良好性能等特点这样的反馈，那是因为 API 的设计不需要考虑类似请求开销这样一些事情。

话虽如此，我们正在考虑的改进重点是典型的浏览用例，因为这是协议主要的使用场景。

我们的章程里面是这样说的：

正在组织的规范需要满足现在已经普遍部署了的 HTTP 的功能要求；具体来说主要包括，Web 浏览（桌面端和移动端），非浏览器（“HTTP APIs” 形式的），Web 服务（大范围的），还有各种网络中介（借助代理，企业防火墙，反向代理以及内容分发网络实现的）。同样的，对 HTTP/1.x 当前和未来的语义扩展 (例如，消息头，方法，状态码，缓存指令) 都应该在新的协议中支持。

值得注意的是，这里没有包括将 HTTP 用于非特定行为所依赖的场景中（例如超时，连接状态以及拦截代理）。这些可能并不会被最终的产品启用。

### HTTP/2 需要加密吗？ ###

不需要。在激烈的讨论后，工作组没有就新协议是否使用加密（如 TLS）而达成共识。

不过，有些观点认为只有在加密连接上使用时才会支持 HTTP/2，而目前还没有浏览器支持未加密的 HTTP/2。

### HTTP/2 是怎么提高安全性的呢？ ###

HTTP/2 定义了所需的 TLS 文档，包括版本，密码套件黑名单和使用的扩展。

细节详见 [相关规范]( https://link.juejin.im?target=http%3A%2F%2Fhttp2.github.io%2Fhttp2-spec%2F%23TLSUsage ) 。

还有对于一些额外机制的讨论，例如对 HTTP:// URLs（所谓的“机会主义加密”）使用 TLS；详见 [RFC 8164]( https://link.juejin.im?target=https%3A%2F%2Ftools.ietf.org%2Fhtml%2Frfc8164 ) 。

### 我现在可以使用 HTTP/2 吗？ ###

浏览器中，最新版本的 Edge、Safari、Firefox 和 Chrome都支持 HTTP/2。其他基于 Blink 的浏览器也将支持HTTP/2（例如 Opera 和 Yandex 浏览器）。详见 [caniuse]( https://link.juejin.im?target=http%3A%2F%2Fcaniuse.com%2F%23feat%3Dhttp2 ) 。

还有几个可用的服务器（包括来自 [Akamai]( https://link.juejin.im?target=https%3A%2F%2Fhttp2.akamai.com%2F ) ， [Google]( https://link.juejin.im?target=https%3A%2F%2Fwww.google.com%2F ) 和 [Twitter]( https://link.juejin.im?target=https%3A%2F%2Ftwitter.com%2F ) 的主要站点的 beta 支持），以及许多可以部署和测试的开源实现。

有关详细信息，请参阅 [实现列表]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fhttp2%2Fhttp2-spec%2Fwiki%2FImplementations ) 。

### HTTP/2 将会取代 HTTP/1.x 吗？ ###

工作组的目的是让那些使用 HTTP/1.x 的人也可以使用 HTTP/2，并能获得 HTTP/2 所带来的好处。他们说过，由于人们部署代理和服务器的方式不同，我们不能强迫整个世界进行迁移，所以 HTTP/1.x 仍有可能要使用了一段时间。

### HTTP/3 会出现吗？ ###

如果通过 HTTP/2 引入的沟通协作机制运行良好，支持新版本的 HTTP 就会比过去更加容易。

## 实现过程中的问题 ##

### 为什么规则会围绕消息头帧的数据接续？ ###

数据接续的存在是由于一个值（例如 cookie）可以超过 16kb，这意味着它不可能全部装进一个帧里面。

所以就决定以最不容易出错的方式让所有的消息头数据以一个接一个帧的方式传递，这样就使得对消息头的解码和缓冲区的管理变得更加容易。

### HPACK 状态的最小和最大尺寸是多少？ ###

接收一方总是会控制 HPACK 中内存的使用量, 并且最小能设置到 0，最大则要看 SETTING 帧中能表示的最大整型数是多少，目前是 2^32 - 1。

### 我怎样才能避免保持 HPACK 状态？ ###

发送一个 SETTINGS 帧，将状态尺寸（SETTINGS_HEADER_TABLE_SIZE）设置到 0，然后 RST 所有的流，直到一个带有 ACT 设置位的 SETTINGS 帧被接收。

### 为什么会有一个单独的压缩/流程控制上下文？ ###

简单说一下。

原来的提案里面提到了流分组这个概念，它可以共享上下文，进行流控制等等。尽管那样有利于代理（也有利于用户体验），但是这样做相应也会增加一点复杂度。所以我们就决定先以一个简单的东西开始，看看它会有多糟糕的问题，并且在未来的协议版本中解决这些问题（如果有的话）。

### 为什么在 HPACK 中有 EOS 的符号？ ###

由于 CPU 效率和安全的原因，HPACK 的霍夫曼编码填充了霍夫曼编码字符串的下一个字节边界。因此对于任何特定的字符串可能需要 0-7 个比特的填充。

如果单独考虑霍夫曼解码，任何比所需要的填充长的符号都可以正常工作。但是，HPACK 的设计允许按字节对比霍夫曼编码的字符串。通过填充 EOS 符号需要的比特，我们确保用户在做霍夫曼编码字符串字节级比较时是相等的。反之，许多 headers 可以在不需要霍夫曼解码的情况下被解析。

### 实现 HTTP/2 的时候我可以不用去实现 HTTP/1.1 吗？ ###

通常/大部分时候可以。

对于运行在 TLS（ ` h2` ）之上的 HTTP/2 而言，如果你没有实现 ` http1.1` 的 ALPN 标识，那你就不需要支持任何 HTTP/1.1 的特性。

对于运行在 TCP（ ` h2c` ）之上的 HTTP/2 而言，你需要实现最初始的升级（Upgrade）请求。

只支持 ` h2c` 的客户端需要生成一个针对 OPTIONS 的请求，因为 ` “*”` 或者一个针对 “/” 的 HEAD 请求，他们相当安全，并且也很容易构建。仅仅只希望实现 HTTP/2 的客户端应当把没有带上 101 状态码的 HTTP/1.1 响应看做错误处理。

只支持 ` h2c` 的服务器可以使用一个固定的 101 响应来接收一个包含升级（Upgrade）消息头字段的请求。没有 ` h2c` 的升级令牌的请求可以使用一个包含了 Upgrade 消息头字段的 505（HTTP 版本不支持）状态码来拒绝。那些不希望处理 HTTP/1.1 响应的服务器应该在发送了带有鼓励用户在升级了的 HTTP/2 连接上重试的连接序言之后立即用带有 REFUSED_STREAM 错误码拒绝该请求的第一份数据流.

### 5.3.2节中的优先级示例是否正确？ ###

不，那是正确的。流 B 的权重为 4，流 C 的权重为 12。为了确定每个流接收的可用资源的比例，将所有权重（16）相加并将每个流权重除以总权重。因此，流 B 接收四分之一的可用资源，流 C 接收四分之三。因此，正如规范所述： [流 B 理想地接收分配给流 C 的资源的三分之一]( https://link.juejin.im?target=http%3A%2F%2Fhttp2.github.io%2Fhttp2-spec%2F%23rfc.section.5.3.2 ) 。

### HTTP/2 连接中需要 TCP_NODELAY 吗？ ###

是的,有可能。即使对于仅使用单个流下载大量数据的客户端，仍然需要一些数据包以相反的方向发回以实现最大传输速度。在没有设置 TCP_NODELAY（仍然允许 Nagle 算法）的情况下，可以传输的数据包将被延迟一段时间以允许它们与后续分组合并。

例如，如果这样一个数据包告诉对等端有更多可用的窗口来发送数据，那么将其发送延迟数毫秒（或更长时间）会对高速连接造成严重影响。

## 部署问题 ##

### 我该怎么调试加密过的 HTTP/2？ ###

存取应用程序数据的方法很多，最简单的方法是使用 [NSS keylogging]( https://link.juejin.im?target=https%3A%2F%2Fdeveloper.mozilla.org%2Fen-US%2Fdocs%2FMozilla%2FProjects%2FNSS%2FKey_Log_Format ) 配上 Wireshark 插件（包含在最新开发版中）。这种方法对 Firefox 和 Chrome 都适用。

### 我该怎么使用 HTTP/2 的服务端推送？ ###

HTTP/2 服务器推送允许服务器向客户端提供内容而无需等待请求。这可以提高检索资源的时间，特别是对于具有大 [带宽延迟产品]( https://link.juejin.im?target=https%3A%2F%2Fen.wikipedia.org%2Fwiki%2FBandwidth-delay_product ) 的连接，其中网络往返时间占了在资源上花费的大部分时间。

推送基于请求内容而变化的资源可能是不明智的。目前，浏览器只会推送请求，如果他们不这样做，就会提出匹配的请求（详见 [Section 4 of RFC 7234]( https://link.juejin.im?target=https%3A%2F%2Ftools.ietf.org%2Fhtml%2Frfc7234%23section-4 ) ）。

有些缓存不考虑所有请求头字段的变化，即使它们列在 ` Vary` header 字段中。为了使推送资源被接收的可能性最大化，内容协商是最好的选择。基于 ` accept-encoding` 报头字段的内容协商受到缓存的广泛尊重，但是其他报头字段可能不受支持。

> 
> 
> 
> 如果发现译文存在错误或其他需要改进的地方，欢迎到 [掘金翻译计划](
> https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fxitu%2Fgold-miner
> ) 对译文进行修改并 PR，也可获得相应奖励积分。文章开头的 **本文永久链接** 即为本文在 GitHub 上的 MarkDown 链接。
> 
> 

> 
> 
> 
> [掘金翻译计划](
> https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fxitu%2Fgold-miner
> ) 是一个翻译优质互联网技术文章的社区，文章来源为 [掘金]( https://juejin.im ) 上的英文分享文章。内容覆盖 [Android](
> https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fxitu%2Fgold-miner%23android
> ) 、 [iOS](
> https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fxitu%2Fgold-miner%23ios
> ) 、 [前端](
> https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fxitu%2Fgold-miner%23%25E5%2589%258D%25E7%25AB%25AF
> ) 、 [后端](
> https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fxitu%2Fgold-miner%23%25E5%2590%258E%25E7%25AB%25AF
> ) 、 [区块链](
> https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fxitu%2Fgold-miner%23%25E5%258C%25BA%25E5%259D%2597%25E9%2593%25BE
> ) 、 [产品](
> https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fxitu%2Fgold-miner%23%25E4%25BA%25A7%25E5%2593%2581
> ) 、 [设计](
> https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fxitu%2Fgold-miner%23%25E8%25AE%25BE%25E8%25AE%25A1
> ) 、 [人工智能](
> https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fxitu%2Fgold-miner%23%25E4%25BA%25BA%25E5%25B7%25A5%25E6%2599%25BA%25E8%2583%25BD
> ) 等领域，想要查看更多优质译文请持续关注 [掘金翻译计划](
> https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fxitu%2Fgold-miner
> ) 、 [官方微博](
> https://link.juejin.im?target=http%3A%2F%2Fweibo.com%2Fjuejinfanyi ) 、 [知乎专栏](
> https://link.juejin.im?target=https%3A%2F%2Fzhuanlan.zhihu.com%2Fjuejinfanyi
> ) 。
> 
>