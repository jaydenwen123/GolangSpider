# 【面试篇】寒冬求职之你必须要懂的Web安全 #

随着互联网的发展，各种Web应用变得越来越复杂，满足了用户的各种需求的同时，各种网络安全问题也接踵而至。作为前端工程师的我们也逃不开这个问题，今天一起看一看Web前端有哪些安全问题以及我们如何去检测和防范这些问题。非前端的攻击本文不会讨论(如SQL注入，DDOS攻击等)，毕竟后端也非本人擅长的领域。

QQ邮箱、新浪微博、YouTube、WordPress 和 百度 等知名网站都曾遭遇攻击，如果你从未有过安全方面的问题，不是因为你所开发的网站很安全，更大的可能是你的网站的流量非常低或者没有攻击的价值。

本文主要讨论以下几种攻击方式: XSS攻击、CSRF攻击以及点击劫持。

> 
> 
> 
> **希望大家在阅读完本文之后，能够很好的回答以下几个面试题。**
> 
> 

1.前端有哪些攻击方式？

2.什么是XSS攻击？XSS攻击有几种类型？如果防范XSS攻击？

3.什么是CSRF攻击？如何防范CSRF攻击

4.如何检测网站是否安全？

在开始之前，建议大家先clone代码，我为大家准备好了示例代码，并且写了详细的注释，大家可以对照代码来理解每一种攻击以及如何去防范攻击，毕竟看再多的文字，都不如实操。(ReadMe中详细得写了操作步骤)： [github.com/YvetteLau/B…]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2FYvetteLau%2FBlog%2Ftree%2Fmaster%2FSecurity )

![](https://user-gold-cdn.xitu.io/2019/5/11/16aa69598156339c?imageView2/0/w/1280/h/960/ignore-error/1)

## 1. XSS攻击 ##

XSS(Cross-Site Scripting，跨站脚本攻击)是一种代码注入攻击。攻击者在目标网站上注入恶意代码，当被攻击者登陆网站时就会执行这些恶意代码，这些脚本可以读取 cookie，session tokens，或者其它敏感的网站信息，对用户进行钓鱼欺诈，甚至发起蠕虫攻击等。

XSS 的本质是：恶意代码未经过滤，与网站正常的代码混在一起；浏览器无法分辨哪些脚本是可信的，导致恶意脚本被执行。由于直接在用户的终端执行，恶意代码能够直接获取用户的信息，利用这些信息冒充用户向网站发起攻击者定义的请求。

> 
> 
> 
> XSS分类
> 
> 

根据攻击的来源，XSS攻击可以分为存储型(持久性)、反射型(非持久型)和DOM型三种。下面我们来详细了解一下这三种XSS攻击：

> 
> 
> 
> ### 1.1 反射型XSS ###
> 
> 

当用户点击一个恶意链接，或者提交一个表单，或者进入一个恶意网站时，注入脚本进入被攻击者的网站。Web服务器将注入脚本，比如一个错误信息，搜索结果等，未进行过滤直接返回到用户的浏览器上。

> 
> 
> 
> 反射型 XSS 的攻击步骤：
> 
> 

* 攻击者构造出特殊的 ` URL` ，其中包含恶意代码。
* 用户打开带有恶意代码的 ` URL` 时，网站服务端将恶意代码从 ` URL` 中取出，拼接在 HTML 中返回给浏览器。
* 用户浏览器接收到响应后解析执行，混在其中的恶意代码也被执行。
* 恶意代码窃取用户数据并发送到攻击者的网站，或者冒充用户的行为，调用目标网站接口执行攻击者指定的操作。

反射型 XSS 漏洞常见于通过 ` URL` 传递参数的功能，如网站搜索、跳转等。由于需要用户主动打开恶意的 ` URL` 才能生效，攻击者往往会结合多种手段诱导用户点击。

POST 的内容也可以触发反射型 XSS，只不过其触发条件比较苛刻（需要构造表单提交页面，并引导用户点击），所以非常少见。

> 
> 
> 
> 查看反射型攻击示例
> 
> 

请戳： [github.com/YvetteLau/B…]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2FYvetteLau%2FBlog%2Ftree%2Fmaster%2FSecurity )

根据 ` README.md` 的提示进行操作(真实情况下是需要诱导用户点击的，上述代码仅是用作演示)。

注意 ` Chrome` 和 ` Safari` 能够检测到 ` url` 上的xss攻击，将网页拦截掉，但是其它浏览器不行，如 ` Firefox`

如果不希望被前端拿到cookie，后端可以设置 ` httpOnly` (不过这不是 ` XSS攻击` 的解决方案，只能降低受损范围)

> 
> 
> 
> 如何防范反射型XSS攻击
> 
> 

**对字符串进行编码。**

对url的查询参数进行转义后再输出到页面。

` app.get( '/welcome' , function ( req, res ) { //对查询参数进行编码，避免反射型 XSS攻击 res.send( ` ${ encodeURIComponent (req.query.type)} ` ); }); 复制代码`
> 
> 
> 
> 
> ### 1.2 DOM 型 XSS ###
> 
> 

DOM 型 XSS 攻击，实际上就是前端 ` JavaScript` 代码不够严谨，把不可信的内容插入到了页面。在使用 `.innerHTML` 、 `.outerHTML` 、 `.appendChild` 、 ` document.write()` 等API时要特别小心，不要把不可信的数据作为 HTML 插到页面上，尽量使用 `.innerText` 、 `.textContent` 、 `.setAttribute()` 等。

> 
> 
> 
> DOM 型 XSS 的攻击步骤：
> 
> 

* 攻击者构造出特殊数据，其中包含恶意代码。
* 用户浏览器执行了恶意代码。
* 恶意代码窃取用户数据并发送到攻击者的网站，或者冒充用户的行为，调用目标网站接口执行攻击者指定的操作。

> 
> 
> 
> 如何防范 DOM 型 XSS 攻击
> 
> 

防范 DOM 型 XSS 攻击的核心就是对输入内容进行转义(DOM 中的内联事件监听器和链接跳转都能把字符串作为代码运行，需要对其内容进行检查)。

1.对于 ` url` 链接(例如图片的 ` src` 属性)，那么直接使用 ` encodeURIComponent` 来转义。

2.非 ` url` ，我们可以这样进行编码：

` function encodeHtml ( str ) { return str.replace( /"/g , '&quot;' ) .replace( /'/g , '&apos;' ) .replace( /</g , '&lt;' ) .replace( />/g , '&gt;' ); } 复制代码`

DOM 型 XSS 攻击中，取出和执行恶意代码由浏览器端完成，属于前端 JavaScript 自身的安全漏洞。

> 
> 
> 
> 查看DOM型XSS攻击示例(根据readme提示查看)
> 
> 

请戳： [github.com/YvetteLau/B…]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2FYvetteLau%2FBlog%2Ftree%2Fmaster%2FSecurity )

> 
> 
> 
> ### 1.3 存储型XSS ###
> 
> 

恶意脚本永久存储在目标服务器上。当浏览器请求数据时，脚本从服务器传回并执行，影响范围比反射型和DOM型XSS更大。存储型XSS攻击的原因仍然是没有做好数据过滤：前端提交数据至服务端时，没有做好过滤；服务端在接受到数据时，在存储之前，没有做过滤；前端从服务端请求到数据，没有过滤输出。

> 
> 
> 
> 存储型 XSS 的攻击步骤：
> 
> 

* 攻击者将恶意代码提交到目标网站的数据库中。
* 用户打开目标网站时，网站服务端将恶意代码从数据库取出，拼接在 HTML 中返回给浏览器。
* 用户浏览器接收到响应后解析执行，混在其中的恶意代码也被执行。
* 恶意代码窃取用户数据并发送到攻击者的网站，或者冒充用户的行为，调用目标网站接口执行攻击者指定的操作。

这种攻击常见于带有用户保存数据的网站功能，如论坛发帖、商品评论、用户私信等。

> 
> 
> 
> 如何防范存储型XSS攻击：
> 
> 

* 前端数据传递给服务器之前，先转义/过滤(防范不了抓包修改数据的情况)
* 服务器接收到数据，在存储到数据库之前，进行转义/过滤
* 前端接收到服务器传递过来的数据，在展示到页面前，先进行转义/过滤

> 
> 
> 
> 查看存储型XSS攻击示例(根据Readme提示查看)
> 
> 

请戳： [github.com/YvetteLau/B…]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2FYvetteLau%2FBlog%2Ftree%2Fmaster%2FSecurity )

**除了谨慎的转义，我们还需要其他一些手段来防范XSS攻击:**

**1.Content Security Policy**

在服务端使用 HTTP的 ` Content-Security-Policy` 头部来指定策略，或者在前端设置 ` meta` 标签。

例如下面的配置只允许加载同域下的资源：

` Content-Security-Policy: default -src 'self' 复制代码` ` < meta http-equiv = "Content-Security-Policy" content = "form-action 'self';" > 复制代码`

前端和服务端设置 CSP 的效果相同，但是 ` meta` 无法使用 ` report`

> 
> 
> 
> 更多的设置可以查看 [Content-Security-Policy
> ](
> https://link.juejin.im?target=https%3A%2F%2Fdeveloper.mozilla.org%2Fzh-CN%2Fdocs%2FWeb%2FHTTP%2FHeaders%2FContent-Security-Policy__by_cnvoid
> )
> 
> 

严格的 CSP 在 XSS 的防范中可以起到以下的作用：

* 禁止加载外域代码，防止复杂的攻击逻辑。
* 禁止外域提交，网站被攻击后，用户的数据不会泄露到外域。
* 禁止内联脚本执行（规则较严格，目前发现 GitHub 使用）。
* 禁止未授权的脚本执行（新特性，Google Map 移动版在使用）。
* 合理使用上报可以及时发现 XSS，利于尽快修复问题。

**2.输入内容长度控制**

对于不受信任的输入，都应该限定一个合理的长度。虽然无法完全防止 XSS 发生，但可以增加 XSS 攻击的难度。

**3.输入内容限制**

对于部分输入，可以限定不能包含特殊字符或者仅能输入数字等。

**4.其他安全措施**

* HTTP-only Cookie: 禁止 JavaScript 读取某些敏感 Cookie，攻击者完成 XSS 注入后也无法窃取此 Cookie。
* 验证码：防止脚本冒充用户提交危险操作。

> 
> 
> 
> ### 1.4 XSS 检测 ###
> 
> 

读到这儿，相信大家已经知道了什么是XSS攻击，XSS攻击的类型，以及如何去防范XSS攻击。但是有一个非常重要的问题是：我们如何去检测XSS攻击，怎么知道自己的页面是否存在XSS漏洞？

很多大公司，都有专门的安全部门负责这个工作，但是如果没有安全部门，作为开发者本身，该如何去检测呢？

1.使用通用 XSS 攻击字串手动检测 XSS 漏洞

如: ` jaVasCript:/*-/*`/*\`/*'/*"/**/(/* */oNcliCk=alert() )//%0D%0A%0d%0a//</stYle/</titLe/</teXtarEa/</scRipt/--!>\x3csVg/<sVg/oNloAd=alert()//>\x3e`

能够检测到存在于 HTML 属性、HTML 文字内容、HTML 注释、跳转链接、内联 JavaScript 字符串、内联 CSS 样式表等多种上下文中的 XSS 漏洞，也能检测 eval()、setTimeout()、setInterval()、Function()、innerHTML、document.write() 等 DOM 型 XSS 漏洞，并且能绕过一些 XSS 过滤器。

` <img src=1 onerror=alert(1)>`

2.使用第三方工具进行扫描（详见最后一个章节）

![](https://user-gold-cdn.xitu.io/2019/5/11/16aa69598bde01d1?imageView2/0/w/1280/h/960/ignore-error/1)

## 2. CSRF ##

CSRF（Cross-site request forgery）跨站请求伪造：攻击者诱导受害者进入第三方网站，在第三方网站中，向被攻击网站发送跨站请求。利用受害者在被攻击网站已经获取的注册凭证，绕过后台的用户验证，达到冒充用户对被攻击的网站执行某项操作的目的。

> 
> 
> 
> 典型的CSRF攻击流程：
> 
> 

* 受害者登录A站点，并保留了登录凭证（Cookie）。
* 攻击者诱导受害者访问了站点B。
* 站点B向站点A发送了一个请求，浏览器会默认携带站点A的Cookie信息。
* 站点A接收到请求后，对请求进行验证，并确认是受害者的凭证，误以为是无辜的受害者发送的请求。
* 站点A以受害者的名义执行了站点B的请求。
* 攻击完成，攻击者在受害者不知情的情况下，冒充受害者完成了攻击。

**一图胜千言：**

![](https://user-gold-cdn.xitu.io/2019/5/15/16abb8d5ab69386f?imageView2/0/w/1280/h/960/ignore-error/1)

> 
> 
> 
> CSRF的特点
> 
> 

1.攻击通常在第三方网站发起，如图上的站点B，站点A无法防止攻击发生。

2.攻击利用受害者在被攻击网站的登录凭证，冒充受害者提交操作；并不会去获取cookie信息(cookie有同源策略)

3.跨站请求可以用各种方式：图片URL、超链接、CORS、Form提交等等(来源不明的链接，不要点击)

> 
> 
> 
> 运行代码，更直观了解一下
> 
> 

用户 loki 银行存款 10W。

用户 yvette 银行存款 1000。

我们来看看 yvette 如何通过 ` CSRF` 攻击，将 loki 的钱偷偷转到自己的账户中，并根据提示，查看如何去防御CSRF攻击。

![](https://user-gold-cdn.xitu.io/2019/5/11/16aa695981d3c956?imageView2/0/w/1280/h/960/ignore-error/1)

请戳： [github.com/YvetteLau/B…]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2FYvetteLau%2FBlog%2Ftree%2Fmaster%2FSecurity ) [根据readme中的CSRF部分进行操作]

#### CSRF 攻击防御 ####

**1. 添加验证码(体验不好)**

验证码能够防御CSRF攻击，但是我们不可能每一次交互都需要验证码，否则用户的体验会非常差，但是我们可以在转账，交易等操作时，增加验证码，确保我们的账户安全。

**2. 判断请求的来源：检测Referer(并不安全，Referer可以被更改)**

` `Referer` 可以作为一种辅助手段，来判断请求的来源是否是安全的，但是鉴于 `Referer` 本身是可以被修改的，因为不能仅依赖于 `Referer` 复制代码`

**3. 使用Token(主流)**

` CSRF攻击之所以能够成功，是因为服务器误把攻击者发送的请求当成了用户自己的请求。那么我们可以要求所有的用户请求都携带一个CSRF攻击者无法获取到的Token。服务器通过校验请求是否携带正确的Token，来把正常的请求和攻击的请求区分开。跟验证码类似，只是用户无感知。 - 服务端给用户生成一个token，加密后传递给用户 - 用户在提交请求时，需要携带这个token - 服务端验证token是否正确 复制代码`

**4. Samesite Cookie属性**

为了从源头上解决这个问题，Google起草了一份草案来改进HTTP协议，为Set-Cookie响应头新增Samesite属性，它用来标明这个 Cookie是个“同站 Cookie”，同站Cookie只能作为第一方Cookie，不能作为第三方Cookie，Samesite 有两个属性值，分别是 Strict 和 Lax。

部署简单，并能有效防御CSRF攻击，但是存在兼容性问题。

**Samesite=Strict**

` Samesite=Strict` 被称为是严格模式,表明这个 Cookie 在任何情况都不可能作为第三方的 Cookie，有能力阻止所有CSRF攻击。此时，我们在B站点下发起对A站点的任何请求，A站点的 Cookie 都不会包含在cookie请求头中。

` **Samesite=Lax** `Samesite=Lax` 被称为是宽松模式，与 Strict 相比，放宽了限制，允许发送安全 HTTP 方法带上 Cookie，如 `Get` / `OPTIONS` 、`HEAD` 请求. 但是不安全 HTTP 方法，如： `POST`, `PUT`, `DELETE` 请求时，不能作为第三方链接的 Cookie 复制代码`

为了更好的防御CSRF攻击，我们可以组合使用以上防御手段。

## 3. 点击劫持 ##

点击劫持是指在一个Web页面中隐藏了一个透明的iframe，用外层假页面诱导用户点击，实际上是在隐藏的frame上触发了点击事件进行一些用户不知情的操作。

#### 典型点击劫持攻击流程 ####

* 攻击者构建了一个非常有吸引力的网页【不知道哪些内容对你们来说有吸引力，我就不写页面了，偷个懒】
* 将被攻击的页面放置在当前页面的 ` iframe` 中
* 使用样式将 iframe 叠加到非常有吸引力内容的上方
* 将iframe设置为100%透明
* 你被诱导点击了网页内容，你以为你点击的是***，而实际上，你成功被攻击了。

#### 点击劫持防御 ####

**1. frame busting**

Frame busting

` if ( top.location != window.location ){ top.location = window.location } 复制代码`

需要注意的是: HTML5中iframe的 ` sandbox` 属性、IE中iframe的 ` security` 属性等，都可以限制iframe页面中的JavaScript脚本执行，从而可以使得 frame busting 失效。

**2. X-Frame-Options**

X-FRAME-OPTIONS是微软提出的一个http头，专门用来防御利用iframe嵌套的点击劫持攻击。并且在IE8、Firefox3.6、Chrome4以上的版本均能很好的支持。

可以设置为以下值:

* 

DENY: 拒绝任何域加载

* 

SAMEORIGIN: 允许同源域下加载

* 

ALLOW-FROM: 可以定义允许frame加载的页面地址

### 安全扫描工具 ###

上面我们介绍了几种常见的前端安全漏洞，也了解一些防范措施，那么我们如何发现自己网站的安全问题呢？没有安全部门的公司可以考虑下面几款开源扫码工具：

#### 1. [Arachni]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2FArachni%2Farachni ) ####

Arachni是基于Ruby的开源，功能全面，高性能的漏洞扫描框架，Arachni提供简单快捷的扫描方式，只需要输入目标网站的网址即可开始扫描。它可以通过分析在扫描过程中获得的信息，来评估漏洞识别的准确性和避免误判。

Arachni默认集成大量的检测工具，可以实施 代码注入、CSRF、文件包含检测、SQL注入、命令行注入、路径遍历等各种攻击。

同时，它还提供了各种插件，可以实现表单爆破、HTTP爆破、防火墙探测等功能。

针对大型网站，该工具支持会话保持、浏览器集群、快照等功能，帮助用户更好实施渗透测试。

#### 2. [Mozilla HTTP Observatory]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fmozilla%2Fhttp-observatory%2F ) ####

Mozilla HTTP Observatory，是Mozilla最近发布的一款名为Observatory的网站安全分析工具，意在鼓励开发者和系统管理员增强自己网站的安全配置。用法非常简单：输入网站URL，即可访问并分析网站HTTP标头，随后可针对网站安全性提供数字形式的分数和字母代表的安全级别。

> 
> 
> 
> 检查的主要范围包括：
> 
> 

* Cookie
* 跨源资源共享（CORS）
* 内容安全策略（CSP）
* HTTP公钥固定（Public Key Pinning）
* HTTP严格安全传输（HSTS）状态
* 是否存在HTTP到HTTPs的自动重定向
* 子资源完整性（Subresource Integrity）
* X-Frame-Options
* X-XSS-Protection

#### 3. [w3af]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fandresriancho%2Fw3af ) ####

W3af是一个基于Python的Web应用安全扫描器。可帮助开发人员，有助于开发人员和测试人员识别Web应用程序中的漏洞。

扫描器能够识别200多个漏洞，包括跨站点脚本、SQL注入和操作系统命令。

> 
> 
> 
> ### 后续写作计划(写作顺序不定) ###
> 
> 

1.《寒冬求职季之你必须要懂的原生JS》(下)

2.《寒冬求职季之你必须要知道的CSS》

3.《寒冬求职季之你必须要懂的一些浏览器知识》

4.《寒冬求职季之你必须要知道的性能优化》

5.《寒冬求职季之你必须要懂的webpack原理》

**针对React技术栈:**

1.《寒冬求职季之你必须要懂的React》系列

2.《寒冬求职季之你必须要懂的ReactNative》系列

![](https://user-gold-cdn.xitu.io/2019/5/11/16aa69598012fc3c?imageView2/0/w/1280/h/960/ignore-error/1)

编写本文，虽然花费了很多时间，但是在这个过程中，我也学习到了很多知识，谢谢各位小伙伴愿意花费宝贵的时间阅读本文，如果本文给了您一点帮助或者是启发，请不要吝啬你的赞和Star，您的肯定是我前进的最大动力。 [github.com/YvetteLau/B…]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2FYvetteLau%2FBlog )

参考文章:

[1] [珠峰架构课(墙裂推荐)]( https://link.juejin.im?target=http%3A%2F%2Fwww.zhufengpeixun.cn%2Fmain%2Fcourse%2Findex.html )

[2] [如何防止CSRF攻击？]( https://link.juejin.im?target=https%3A%2F%2Ftech.meituan.com%2F2018%2F10%2F11%2Ffe-security-csrf.html )

[3] [谈谈对 Web 安全的理解]( https://link.juejin.im?target=https%3A%2F%2Fzhuanlan.zhihu.com%2Fp%2F25486768%3Fgroup_id%3D820705780520079360 )

[4] [程序员必须要了解的web安全]( https://juejin.im/post/5b4e0c936fb9a04fcf59cb79 )

[5] [Cookie的SameSite属性]( https://juejin.im/post/5c8a33dcf265da2dc538fc7d )

[6] [github.com/OWASP/Cheat…]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2FOWASP%2FCheatSheetSeries )

> 
> 
> 
> 关注小姐姐的公众号。
> 
> 

![](https://user-gold-cdn.xitu.io/2019/5/14/16ab621cfa97956f?imageView2/0/w/1280/h/960/ignore-error/1)