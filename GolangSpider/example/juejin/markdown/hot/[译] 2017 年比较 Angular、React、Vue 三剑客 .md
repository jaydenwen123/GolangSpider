# [译] 2017 年比较 Angular、React、Vue 三剑客 #

> 
> 
> 
> * 原文地址： [Angular vs. React vs. Vue: A 2017 comparison](
> https://link.juejin.im?target=https%3A%2F%2Fmedium.com%2Funicorn-supplies%2Fangular-vs-react-vs-vue-a-2017-comparison-c5c52d620176
> )
> * 原文作者： [Jens Neuhaus](
> https://link.juejin.im?target=https%3A%2F%2Fmedium.com%2F%40jensneuhaus%3Fsource%3Dpost_header_lockup
> )
> * 译文出自： [掘金翻译计划](
> https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fxitu%2Fgold-miner
> )
> * 本文永久链接： [github.com/xitu/gold-m…](
> https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fxitu%2Fgold-miner%2Fblob%2Fmaster%2FTODO%2Fangular-vs-react-vs-vue-a-2017-comparison.md
> )
> * 译者： [Raoul1996](
> https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fraoul1996 )
> * 校对者： [caoyi0905](
> https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fcaoyi0905 ) 、 [PCAaron](
> https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2FPCAaron )
> 
> 
> 

# 2017 年比较 Angular、React、Vue 三剑客 #

为 web 应用选择 JavaScript 开发框架是一件很费脑筋的事。现如今 [Angular]( https://link.juejin.im?target=https%3A%2F%2Fangular.io%2F ) 和 [React]( https://link.juejin.im?target=https%3A%2F%2Ffacebook.github.io%2Freact%2F ) 非常流行，并且最近出现的新贵 [VueJS]( https://link.juejin.im?target=https%3A%2F%2Fvuejs.org%2F ) 同样博得了很多人的关注。更重要的是，这只是一些 [新起之秀]( https://link.juejin.im?target=https%3A%2F%2Fhackernoon.com%2Ftop-7-javascript-frameworks-c8db6b85f1d0 ) 。

![Javascripts in 2017 —— things aren’t easy these days!](https://user-gold-cdn.xitu.io/2017/11/16/15fc436d1d260ee3?imageView2/0/w/1280/h/960/ignore-error/1) Javascripts in 2017 —— things aren’t easy these days!

那么我们如何选择使用哪个框架呢？列出他们的优劣是极好的。我们将按照先前文章的方式去做，“ [共有 9 步：为 Web 应用选择一个技术栈]( https://link.juejin.im?target=https%3A%2F%2Fmedium.com%2Funicorn-supplies%2F9-steps-how-to-choose-a-technology-stack-for-your-web-application-a6e302398e55 ) ”。

## 在开始之前 —— 是否应用单页 Web 应用开发？ ##

首先你需要弄明白你需要单页面应用程序（SPA）还是多页面的方式。关于这个问题的详细内容请阅读我的博客文章，“ [单页面应用程序（SPA）与多页 Web 应用程序（MPA）]( https://link.juejin.im?target=https%3A%2F%2Fmedium.com%2Funicorn-supplies%2Fangular-vs-react-vs-vue-a-2017-comparison-c5c52d620176%23 ) ”（即将推出，请关注我 [Twitter]( https://link.juejin.im?target=http%3A%2F%2Fwww.twitter.com%2Fjensneuhaus%2F ) 的更新）。

## 今日首发：Angular，React 和 Vue ##

首先，我想从 **生命周期和战略考虑** 角度讨论。然后，我们再讨论这三个 JavaScript 框架的 **功能和概念** 。最后，我们再做 **结论** 。

以下是我们今天要解决的问题：

* **这些框架或库有多成熟** ？
* 这些框架只会 **火热一时** 吗？
* **这些框架相应的社区规模有多大，能得到多少帮助** ？
* 找到每个框架开发者 **容易吗** ？
* 这些框架的 **基本编程概念** 是什么？
* **对于小型或大型应用程序** ，框架是否易用？
* 每个框架 **学习曲线** 什么样？
* 你期望这些框架的 **性能** 怎么样？
* 在哪能 **仔细了解底层原理** ？
* 你可以用你选择的框架 **开发** 吗？

准备好，听我娓娓道来！

## 生命周期与战略考虑 ##

![比较 React、Angular 和 Vue](https://user-gold-cdn.xitu.io/2017/11/16/15fc436d1c9a7d92?imageView2/0/w/1280/h/960/ignore-error/1) 比较 React、Angular 和 Vue

### 一些历史 ###

**Angular** 是基于 TypeScript 的 Javascript 框架。由 Google 进行开发和维护，它被描述为“超级厉害的 JavaScript [MVW]( https://link.juejin.im?target=https%3A%2F%2Fplus.google.com%2F%2BAngularJS%2Fposts%2FaZNVhj355G2 ) 框架”。Angular（也被称为 “Angular 2+”，“Angular 2” 或者 “ng2”）已被重写，是与 AngularJS（也被称为 “Angular.js” 或 “AngularJS 1.x”）不兼容的后续版本。当 AngularJS（旧版本）最初于2010年10月发布时，仍然在 [修复 bug]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fangular%2Fangular.js ) ，等等 —— 新的 Angular（sans JS）于 2016 年 9 月推出版本 2。最新的主版本是 4， [因为版本 3 被跳过了]( https://link.juejin.im?target=http%3A%2F%2Fwww.infoworld.com%2Farticle%2F3150716%2Fapplication-development%2Fforget-angular-3-google-skips-straight-to-angular-4.html ) 。Google，Vine，Wix，Udemy，weather.com，healthcare.gov 和 Forbes 都使用 Angular（根据 [madewithangular]( https://link.juejin.im?target=https%3A%2F%2Fwww.madewithangular.com%2F ) ， [stackshare]( https://link.juejin.im?target=https%3A%2F%2Fstackshare.io%2Fangular-2 ) 和 [libscore.com]( https://link.juejin.im?target=http%3A%2F%2Flibscore.com%2F%23angular ) 提供的数据）。

**React** 被描述为 “用于构建用户界面的 JavaScript 库”。React 最初于 2013 年 3 月发布，由 Facebook 进行开发和维护，Facebook 在多个页面上使用 React 组件（但不是作为单页应用程序）。根据 [Chris Cordle]( https://link.juejin.im?target=https%3A%2F%2Fmedium.com%2F%40chriscordle ) [这篇文章]( https://link.juejin.im?target=https%3A%2F%2Fmedium.com%2F%40chriscordle%2Fwhy-angular-2-4-is-too-little-too-late-ea86d7fa0bae ) 的统计，React 在 Facebook 上的使用远远多于 Angular 在 Google 上的使用。React 还被 Airbnb，Uber，Netflix，Twitter，Pinterest，Reddit，Udemy，Wix，Paypal，Imgur，Feedly，Stripe，Tumblr，Walmart 等使用（根据 [Facebook]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Ffacebook%2Freact%2Fwiki%2FSites-Using-React ) , [stackshare]( https://link.juejin.im?target=https%3A%2F%2Fstackshare.io%2Freact ) 和 [libscore.com]( https://link.juejin.im?target=http%3A%2F%2Flibscore.com%2F%23React ) 提供的数据）。

Facebook 正在开发 **React Fiber** 。它会改变 React 的底层 - 渲染速度应该会更快 - 但是在变化之后，版本会向后兼容。Facebook 将会在 2017 年 4 月的开发者大会上 [讨论]( https://link.juejin.im?target=https%3A%2F%2Fdevelopers.facebook.com%2Fvideos%2Ff8-2017%2Fthe-evolution-of-react-and-graphql-at-facebook-and-beyond%2F ) 新变化，并发布一篇非官方的 [关于新架构的文章]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Facdlite%2Freact-fiber-architecture ) 。React Fiber 可能与 React 16 一起发布。

**Vue** 是 2016 年发展最为迅速的 JS 框架之一。Vue 将自己描述为一款“用于构建直观，快速和组件化交互式界面的 [MVVM]( https://link.juejin.im?target=https%3A%2F%2Fen.wikipedia.org%2Fwiki%2FModel%25E2%2580%2593view%25E2%2580%2593viewmodel ) 框架”。它于 2014 年 2 月首次由 Google 前员工 [Evan You]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fyyx990803 ) 发布（顺便说一句：尤雨溪那时候发表了一篇 [vue 发布首周的营销活动和数据]( https://link.juejin.im?target=http%3A%2F%2Fblog.evanyou.me%2F2014%2F02%2F11%2Ffirst-week-of-launching-an-oss-project%2F ) 的博客文章）。尤其是考虑到 Vue 在没有大公司的支持的情况下，作为一个人开发的框架还能获得这么多的吸引力，这无疑是非常成功的。尤雨溪目前有一个包含数十名核心开发者的团队。2016 年，版本 2 发布。Vue 被阿里巴巴，百度，Expedia，任天堂，GitLab 使用 — 可以在 [madewithvuejs.com]( https://link.juejin.im?target=https%3A%2F%2Fmadewithvuejs.com%2F ) 找到一些小型项目的列表。

Angular 和 Vue 都遵守 **MIT license** 许可，而 React 遵守 **[BSD3-license]( https://link.juejin.im?target=https%3A%2F%2Fen.wikipedia.org%2Fwiki%2FBSD_licenses%233-clause_license_.28.22BSD_License_2.0.22.2C_.22Revised_BSD_License.22.2C_.22New_BSD_License.22.2C_or_.22Modified_BSD_License.22.29 ) 许可证** 。在专利文件上有很多讨论。 [James Ide]( https://link.juejin.im?target=https%3A%2F%2Fmedium.com%2F%40ji ) （前 Facebook 工程师）解释专利文件背后的 [原因和历史]( https://link.juejin.im?target=https%3A%2F%2Fmedium.com%2F%40ji%2Fthe-react-license-for-founders-and-ctos-b38d2538f3e5 ) ：Facebook 的专利授权是在保护自己免受专利诉讼的能力的同时分享其代码。专利文件被更新了一次，有些人声称，如果你的公司不打算起诉 Facebook，那么使用 React 是可以的。你可以 [在 Github 的这个 issue 上]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Ffacebook%2Freact%2Fissues%2F7293 ) 查看讨论。我不是律师，所以如果 React 许可证对你或你的公司有问题，你应该自己决定。关于这个话题还有很多文章： [Dennis Walsh]( https://link.juejin.im?target=https%3A%2F%2Fmedium.com%2F%40dwalsh.sdlr ) 写到， [你为什么不该害怕]( https://link.juejin.im?target=https%3A%2F%2Fmedium.com%2F%40dwalsh.sdlr%2Freact-facebook-and-the-revokable-patent-license-why-its-a-paper-25c40c50b562 ) 。 [Raúl Kripalani]( https://link.juejin.im?target=https%3A%2F%2Fmedium.com%2F%40raulk ) 警告： [反对创业公司使用 React]( https://link.juejin.im?target=https%3A%2F%2Fmedium.com%2F%40raulk%2Fif-youre-a-startup-you-should-not-use-react-reflecting-on-the-bsd-patents-license-b049d4a67dd2 ) ，他还写了一篇 [备忘录概览]( https://link.juejin.im?target=https%3A%2F%2Fmedium.com%2F%40raulk%2Ffurther-notes-and-questions-arising-from-facebooks-bsd-3-strong-patent-retaliation-license-c6386e8e1d60 ) 。此外，Facebook 上还有一个最新的声明： [解释 React 的许可证]( https://link.juejin.im?target=https%3A%2F%2Fcode.facebook.com%2Fposts%2F112130496157735%2Fexplaining-react-s-license%2F ) 。

### 核心开发 ###

如前所述，Angular 和 React 得到大公司的支持和使用。Facebook，Instagram 和 WhatsApp 正在它们的页面使用 React。Google 在很多项目中使用 Angular，例如，新的 Adwords 用户界面是使用 [Angular 和 Dart]( https://link.juejin.im?target=http%3A%2F%2Fnews.dartlang.org%2F2016%2F03%2Fthe-new-adwords-ui-uses-dart-we-asked.html%3Fm%3D1 ) 。然而，Vue 是由一群通过 Patreon 和其他赞助方式支持的个人实现的，是好坏你自己确定。 [Matthias Götzke]( https://link.juejin.im?target=https%3A%2F%2Fmedium.com%2F%40mgoetzke ) 认为 Vue 小团队的好处是 [用了更简洁和更少的过度设计的代码或 API]( https://link.juejin.im?target=https%3A%2F%2Fmedium.com%2F%40mgoetzke%2Fsome-people-have-been-asking-about-the-dependability-of-vue-jss-9dc2842b3709 ) 。

我们来看看一些统计数据：Angular 在团队介绍页 [列出 36 人]( https://link.juejin.im?target=https%3A%2F%2Fangular.io%2Fabout%3Fgroup%3DAngular ) ，Vue [列出 16 人]( https://link.juejin.im?target=https%3A%2F%2Fvuejs.org%2Fv2%2Fguide%2Fteam.html ) ，而 React 没有团队介绍页。 **在 Github 上** ，Angular 有 25,000+ 的 star 和 463 位代码贡献者，React 有 70,000+ 的 star 和 1,000+ 位代码贡献者，而 Vue 有近 60,000 的 star 和只有 120 位代码贡献者。你也可以查看 [Angular，React 和 Vue 的 Github Star 历史]( https://link.juejin.im?target=http%3A%2F%2Fwww.timqian.com%2Fstar-history%2F%23facebook%2Freact%26amp%3Bangular%2Fangular%26amp%3Bvuejs%2Fvue ) 。又一次说明 Vue 的趋势似乎很好。根据 [bestof.js]( https://link.juejin.im?target=https%3A%2F%2Fbestof.js.org%2Ftags%2Fframework%2Ftrending%2Flast-3-months ) 提供的数据显示，在过去三个月 Angular 2 平均每天获得 31 个 star，React 74 个，Vue.JS 107 个。

![Angular，React 与 Due 的 Github Star 历史 (数据来源)](https://user-gold-cdn.xitu.io/2017/11/16/15fc436d1d7105c8?imageView2/0/w/1280/h/960/ignore-error/1) Angular，React 与 Due 的 Github Star 历史 (数据来源)

[数据来源]( https://link.juejin.im?target=http%3A%2F%2Fwww.timqian.com%2Fstar-history%2F%23facebook%2Freact%26amp%3Bangular%2Fangular%26amp%3Bvuejs%2Fvue )

**更新** : 感谢 [Paul Henschel]( https://link.juejin.im?target=https%3A%2F%2Fmedium.com%2F%40drcmda ) 提出的 [npm 趋势]( https://link.juejin.im?target=http%3A%2F%2Fwww.npmtrends.com%2Fangular-vs-react-vs-vue-vs-%40angular%2Fcore ) 。npm 趋势显示了 npm 包的下载次数，相对比单独地看 Github star 更有用：

![在过去 2 年，npm 包的下载次数](https://user-gold-cdn.xitu.io/2017/11/16/15fc436d1e44c6c6?imageView2/0/w/1280/h/960/ignore-error/1) 在过去 2 年，npm 包的下载次数

### 市场生命周期 ###

由于各种名称和版本，很难在 Google 趋势中比较 Angular，React 和 Vue。一种近似的方法可以是“互联网与技术”类别中的搜索。结果如下：

![](https://user-gold-cdn.xitu.io/2017/11/16/15fc436d1f10f410?imageView2/0/w/1280/h/960/ignore-error/1)

Vue 没有在 2014 年之前创建 - 所以这里有什么不对劲。La Vue是法语的 “view” ，“sight” 或 “opinion”。也许就是这样。“VueJS” 和 “Angular” 或 “React” 的比较也是不公平的，因为 VueJS 几乎没有搜索到任何结果。

那我们试试别的吧。ThoughtWorks 的 [Technology Radar]( https://link.juejin.im?target=https%3A%2F%2Fwww.thoughtworks.com%2Fde%2Fradar%23 ) 技术随时间推移的变化。ThoughtWorks 的 [Technology Radar]( https://link.juejin.im?target=https%3A%2F%2Fwww.thoughtworks.com%2Fde%2Fradar%23 ) 随着时间推移，技术的演进过程给人深刻的印象。Redux 是 [在采用阶段]( https://link.juejin.im?target=https%3A%2F%2Fwww.thoughtworks.com%2Fde%2Fradar%2Flanguages-and-frameworks%2Fredux ) （被 ThoughtWorks 项目采用的！），它在许多 ThoughtWorks 项目中的价值是不可估量的。Vue.js 是 [在试用阶段]( https://link.juejin.im?target=https%3A%2F%2Fwww.thoughtworks.com%2Fde%2Fradar%2Flanguages-and-frameworks%2Fvue-js ) （被试着用的）。Vue被描述为具有平滑学习曲线的，轻量级并具灵活性的Angular的替代品。Angular 2 是 [正在处于评估阶段]( https://link.juejin.im?target=https%3A%2F%2Fwww.thoughtworks.com%2Fde%2Fradar%2Flanguages-and-frameworks%2Fangular-2 ) 使用 —— 已被 ThoughtWork 团队成功实践，但是还没有被强烈推荐。

根据 [2017 年 Stackoverflow 的最新调查]( https://link.juejin.im?target=https%3A%2F%2Finsights.stackoverflow.com%2Fsurvey%2F2017%23most-loved-dreaded-and-wanted ) ，被调查的开发者中，喜爱 React 有 67%，喜欢 AngularJS 的有 52%。“没有兴趣在开发中继续使用”的开发者占了更高的数量，AngularJS（48%）和 React（33%）。在这两种情况下，Vue 都不在前十。然后是 statejs.com 关于比较 “ [前端框架]( https://link.juejin.im?target=http%3A%2F%2Fstateofjs.com%2F2016%2Ffrontend%2F ) ” 的调查。最有意思的事实是：React 和 Angular 有 100% 的认知度，23% 的受访者不了解 Vue。关于满意度，92% 的受访者愿意“再次使用” React ，Vue 有 89% ,而 Angular 2 只有 65%。

客户满意度调查呢？ [Eric Elliott]( https://link.juejin.im?target=https%3A%2F%2Fmedium.com%2F%40_ericelliott ) 于 2016 年 10 月开始评估 Angular 2 和 React。只有 38% 的受访者会再次使用 Angular 2，而 84% 的人会再次使用 React。

### 长期支持和迁移 ###

Facebook [在其设计原则中指出]( https://link.juejin.im?target=https%3A%2F%2Ffacebook.github.io%2Freact%2Fcontributing%2Fdesign-principles.html%23stability ) ，React API 非常稳定。还有一些脚本可以帮助你从当前的API移到更新的版本：请查阅 [react-codemod]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Freactjs%2Freact-codemod ) 。迁移是非常容易的，没有这样的东西（需要）作为长期支持的版本。在 Reddit 这篇文章中指出，人们看到到升级 [从来不是问题]( https://link.juejin.im?target=https%3A%2F%2Fwww.reddit.com%2Fr%2Freactjs%2Fcomments%2F5a45ai%2Fis_react_a_good_choice_for_a_stable_longterm_app%2F ) 。React 团队写了一篇关于他们 [版本控制方案]( https://link.juejin.im?target=https%3A%2F%2Ffacebook.github.io%2Freact%2Fblog%2F2016%2F02%2F19%2Fnew-versioning-scheme.html ) 的博客文章。当他们添加弃用警告时，在下一个主要版本中的行为发生更改之前，他们会保留当前版本的其余部分。没有计划更改为新的主要版本 - v14 于 2015 年 10 月发布，v15 于 2016 年 4 月发布，而 v16 还没有发布日期。（译者注： [v16 于 2017 年 9 月底发布]( https://link.juejin.im?target=https%3A%2F%2Freactjs.org%2Fblog%2F2017%2F09%2F26%2Freact-v16.0.html ) ）最近 [React核心开发人员指出]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Ffacebook%2Freact%2Fissues%2F8854%23issuecomment-312527769 ) ，升级不应该是一个问题。

关于 Angular，从 v2 发布开始，有一篇 [关于版本管理和发布 Angular]( https://link.juejin.im?target=http%3A%2F%2Fangularjs.blogspot.de%2F2016%2F10%2Fversioning-and-releasing-angular.html ) 的博客文章。每六个月会有一次重大更新，至少有六个月的时间（两个主要版本）。在文档中有一些实验性的 API 被标记为较短的弃用期。目前还没有官方公告，但 [根据这篇文章]( https://link.juejin.im?target=https%3A%2F%2Fwww.infoq.com%2Fnews%2F2017%2F04%2Fng-conf-2017-keynote ) ，Angular 团队已经宣布了以 Angular 4 开始的长期支持版本。这些将在下一个主要版本发布之后至少一年得到支持。这意味着至少在 **2018 年 9 月** 之前，将支持 Angular 4，并提供 bug 修复和重要补丁。在大多数情况下，将 Angular 从 v2 更新到 v4 与更新 Angular 依赖关系一样简单。Angular 还提供了有关是否需要进一步更改的 [信息指南]( https://link.juejin.im?target=https%3A%2F%2Fangular-update-guide.firebaseapp.com%2F ) 。

Vue 1.x 到 2.0 的更新过程对于一个小应用程序来说应该很容易 - 开发者团队已经声称 90% 的 API 保持不变。在控制台上有一个很好的升级 - 诊断迁移 - 辅助工具。一位开发人员 [指出]( https://link.juejin.im?target=https%3A%2F%2Fnews.ycombinator.com%2Fitem%3Fid%3D13151966 ) ，从 v1 到 v2 的更新在大型应用程序中仍然没有挑战。不幸的是，关于 LTS 版本的下一个主要版本或计划信息没有清晰的（公共）路径。

还有一件事：Angular 是一个完整的框架，提供了很多捆绑在一起的东西。React 比 Angular 更灵活，你可能会使用更多独立的，不稳定的，快速更新的库 - 这意味着你需要自己维护相应的更新和迁移。如果某些包不再被维护，或者其他一些包在某些时候成为事实上的标准，这也可能是不利的。

### 人力资源与招聘 ###

如果你的团队有不需要了解更多 Javascript 技术的 HTML 开发人员，则最好选择 Angular 或 Vue。React 需要了解更多的 JavaScript 技术（我们稍后再谈）。

你的团队有工作时可以敲代码的设计师吗？Reddit 上的用户 “pier25” 指出，如果你在 Facebook 工作， [每个人都是一个资深开发者，React 是有意义的]( https://link.juejin.im?target=https%3A%2F%2Fwww.reddit.com%2Fr%2Fwebdev%2Fcomments%2F5ho71i%2Fwhy_we_chose_vuejs_over_react%2Fdeuynwc%2F ) 。然而事实上，你不会总是找到一个可以修改 JSX 的设计师，因此使用 HTML 模板将会更容易。

Angular 框架的好处是来自另一家公司的新的 Angular 2 开发人员将很快熟悉所有必要的约定。React 项目在架构决策方面各不相同，开发人员需要熟悉特定的项目设置。

如果你的开发人员具有面向对象的背景或者不喜欢 Javascript，Angular 也是很好的选择。为了推动这一点，这里是 [Mahesh Chand 引述]( https://link.juejin.im?target=http%3A%2F%2Fwww.c-sharpcorner.com%2Farticle%2Fangular-2-or-react-for-decision-makers%2F ) ：

> 
> 
> 
> 我不是一个 JavaScript 开发人员。我的背景是使用 “真正的” 软件平台构建大型企业系统。我从 1997 年开始使用 C，C
> ++，Pascal，Ada 和 Fortran 构建应用程序。（...）我可以清楚地说，JavaScript 对我来说简直是胡言乱语。作为
> Microsoft MVP 和专家，我对 TypeScript 有很好的理解。我也不认为 Facebook 是一家软件开发公司。但是，Google
> 和微软已经是最大的软件创新者。我觉得使用 Google 和微软强大支持的产品会更舒服。另外（...）与我的背景，我知道微软对 TypeScript
> 有更宏伟的蓝图。
> 
> 

emmmmmmmm...... 我应该提到的，Mahesh 是微软的区域总监。

## React，Angular 和 Vue 的比较 ##

### 组件 ###

我们所讨论的框架都是基于组件的。一个组件得到一个输入，并且在一些内部的行为/计算之后，它返回一个渲染的 UI 模板（一个登录/注销区或一个待办事项列表项）作为输出。定义的组件应该易于在网页或其他组件中重用。例如，你可以使用具有各种属性（列，标题信息，数据行等）的网格组件（由一个标题组件和多个行组件组成），并且能够在另一个页面上使用具有不同数据集的组件。这里有一篇 [关于组件的综合性文章]( https://link.juejin.im?target=https%3A%2F%2Fderickbailey.com%2F2015%2F08%2F26%2Fbuilding-a-component-based-web-ui-with-modern-javascript-frameworks%2F ) ，如果你想了解更多这方面的信息。

React 和 Vue 都擅长处理组件：小型的无状态的函数接收输入和返回元素作为输出。

### Typescript，ES6 与 ES5 ###

React 专注于使用 Javascript ES6。Vue 使用 Javascript ES5 或 ES6。

Angular 依赖于 TypeScript。这在相关示例和开源项目中提供了更多的一致性（React 示例可以在 ES5 或 ES6 中找到）。这也引入了像装饰器和静态类型的概念。静态类型对于代码智能工具非常有用，比如自动重构，跳转到定义等等 - 它们也可以减少应用程序中的错误数量，尽管这个话题当然没有共识。 [Eric Elliott]( https://link.juejin.im?target=https%3A%2F%2Fmedium.com%2F%40_ericelliott ) 在他的文章 “ [静态类型的令人震惊的秘密]( https://link.juejin.im?target=https%3A%2F%2Fmedium.com%2Fjavascript-scene%2Fthe-shocking-secret-about-static-types-514d39bf30a3 ) ” 中不同意上面的观点。Daniel C Wang 表示， [使用静态类型并没有什么坏处]( https://link.juejin.im?target=https%3A%2F%2Fmedium.com%2F%40danwang74%2Fthe-economics-between-testing-and-types-4a3f8c8a86eb ) ，同时有测试驱动开发（TDD）和静态类型挺好的。

你也应该知道你 [可以使用 Flow 在 React 中启用类型检查]( https://link.juejin.im?target=https%3A%2F%2Fwww.sitepoint.com%2Fwriting-better-javascript-with-flow%2F ) 。这是 Facebook 为 JavaScript 开发的静态类型检查器。Flow [也可以集成到 VueJS 中]( https://link.juejin.im?target=https%3A%2F%2Falligator.io%2Fvuejs%2Fcomponents-flow%2F ) 。

如果你在用 TypeScript 编写代码，那么你不需要再编写标准的 JavaScript 了。尽管它在不断发展，但与整个 JavaScript 语言相比，TypeScript 的用户群仍然很小。一个风险可能是你正在向错误的方向发展，因为 TypeScript 可能 - 也许不太可能 - 随着时间的推移也会消失。此外，TypeScript 为项目增加了很多（学习）开销 - 你可以在 [Eric Elliott]( https://link.juejin.im?target=https%3A%2F%2Fmedium.com%2F%40_ericelliott ) 的 [Angular 2 vs. React 比较]( https://link.juejin.im?target=https%3A%2F%2Fmedium.com%2Fjavascript-scene%2Fangular-2-vs-react-the-ultimate-dance-off-60e7dfbc379c ) 中阅读更多关于这方面的内容。

**更新** : [James Ravenscroft]( https://link.juejin.im?target=https%3A%2F%2Fmedium.com%2F%40jrwebdev ) 在对这篇文章的评论中写道， [TypeScript 对 JSX 有一流的支持]( https://link.juejin.im?target=https%3A%2F%2Fmedium.com%2F%40jrwebdev%2Fid-argue-that-if-you-love-typescript-then-react-may-be-a-better-choice-ceec950ee543 ) - 可以无缝地对组件进行类型检查。所以，如果你喜欢 TypeScript 并且你想使用 React，这应该不成问题。

### 模板 —— JSX 还是 HTML ###

React 打破了长期以来的最佳实践。几十年来，开发人员试图分离 UI 模板和内联的 Javascript 逻辑，但是使用 JSX，这些又被混合了。这听起来很糟糕，但是你应该听彼得·亨特（Peter Hunt）的演讲 “ [React：反思最佳实践]( https://link.juejin.im?target=https%3A%2F%2Fwww.youtube.com%2Fwatch%3Fv%3Dx7cQ3mrcKaY ) ”（2013 年 10 月）。他指出，分离模板和逻辑仅仅是技术的分离，而不是关注的分离。你应该构建组件而不是模板。组件是可重用的、可组合的、可单元测试的。

JSX 是一个类似 HTML 语法的可选预处理器，并随后在 JavaScript 中进行编译。JSX 有一些怪癖 —— 例如，你需要使用 className 而不是 class，因为后者是 Javascript 的保留字。JSX 对于开发来说是一个很大的优势，因为代码写在同一个地方，可以在代码完成和编译时更好地检查工作成果。当你在 JSX 中输入错误时，React 将不会编译，并打印输出错误的行号。Angular 2 在运行时静默失败（如果使用 Angular 中的预编译，这个参数可能是无效的）。

JSX 意味着 React 中的所有内容都是 Javascript -- 用于JSX模板和逻辑。 [Cory House]( https://link.juejin.im?target=https%3A%2F%2Fmedium.com%2F%40housecor ) 在 [2016 年 1 月的文章]( https://link.juejin.im?target=https%3A%2F%2Fmedium.freecodecamp.org%2Fangular-2-versus-react-there-will-be-blood-66595faafd51 ) 中指出：“Angular 2 继续把 'JS' 放到 HTML 中。React 把 'HTML' 放到 JS 中。“这是一件好事，因为 Javascript 比 HTML 更强大。

Angular 模板使用特殊的 Angular 语法（比如 ngIf 或 ngFor）来增强 HTML。虽然 React 需要 JavaScript 的知识，但 Angular 会迫使你学习 [Angular 特有的语法]( https://link.juejin.im?target=https%3A%2F%2Fangular.io%2Fguide%2Fcheatsheet ) 。

Vue 具有“ [单个文件组件]( https://link.juejin.im?target=https%3A%2F%2Fvuejs.org%2Fv2%2Fguide%2Fsingle-file-components.html ) ”。这似乎是对于关注分离的权衡 - 模板，脚本和样式在一个文件中，但在三个不同的有序部分中。这意味着你可以获得语法高亮，CSS 支持以及更容易使用预处理器（如 Jade 或 SCSS）。我已经阅读过其他文章，JSX 更容易调试，因为 Vue 不会显示不规范 HTML 的语法错误。这是不正确的，因为 Vue [转换 HTML 来渲染函数]( https://link.juejin.im?target=https%3A%2F%2Fvuejs.org%2Fv2%2Fguide%2Frender-function.html ) - 所以错误显示没有问题（感谢 [Vinicius Reis]( https://link.juejin.im?target=https%3A%2F%2Fmedium.com%2F%40luizvinicius73 ) 的评论和更正！）。

旁注：如果你喜欢 JSX 的思路，并想在 Vue 中使用它，可以使用 [babel-plugin-transform-vue-jsx]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fvuejs%2Fbabel-plugin-transform-vue-jsx ) 。

### 框架和库 ###

Angular 是一个框架而不是一个库，因为它提供了关于如何构建应用程序的强有力的约束，并且还提供了更多开箱即用的功能。Angular 是一个 “完整的解决方案” - 功能齐全，你可以愉快的开始开发。你不需要研究库，路由解决方案或类似的东西 - 你只要开始工作就好了。

另一方面，React 和 Vue 是很灵活的。他们的库可以和各种包搭配。（在 [npm]( https://link.juejin.im?target=https%3A%2F%2Fwww.npmjs.com%2Fsearch%3Fq%3Dreact%26amp%3Bpage%3D1%26amp%3Branking%3Dpopularity ) 上有很多 React 的包，但 Vue 的包比较少，因为毕竟这个框架还比较新）。有了 React，你甚至可以交换库本身的 API 兼容替代品，如 [Inferno]( https://link.juejin.im?target=https%3A%2F%2Finfernojs.org%2F ) 。然而，灵活性越大，责任就越大 - React 没有规则和有限的指导。每个项目都需要决定架构，而且事情可能更容易出错。

另一方面，Angular 还有一个令人困惑的构建工具，样板，检查器（linter）和时间片来处理。如果使用项目初始套件或样板，React 也是如此。他们自然是非常有帮助的，但是 React 可以开箱即用，这也许是你应该学习的方式。有时，在 JavaScript 环境中工作要使用各种工具被称为 “Javascript 疲劳症”。 [Eric Clemmons]( https://link.juejin.im?target=https%3A%2F%2Fmedium.com%2F%40ericclemmons ) 在他的 [文章]( https://link.juejin.im?target=https%3A%2F%2Fmedium.com%2F%40ericclemmons%2Fjavascript-fatigue-48d4011b6fc4 ) 中说：

> 
> 
> 
> 当开始使用框架，还有一堆安装的工具，你可能会不习惯。这些都是框架生成的。很多开发人员不明白，框架内部发生了什么 ——
> 或者需要花费很多时间才能搞明白。
> 
> 

Vue 似乎是三个框架中最轻量的。GitLab 有一篇 [关于 Vue.js（2016 年 10 月）的决定的博客文章]( https://link.juejin.im?target=https%3A%2F%2Fabout.gitlab.com%2F2016%2F10%2F20%2Fwhy-we-chose-vue%2F ) ：

> 
> 
> 
> Vue.js 完美的兼顾了它将为你做什么和你需要做什么。（...）Vue.js 始终是可及的，一个坚固，但灵活的安全网，保证编程效率和把操作 DOM
> 造成的痛苦降到最低。
> 
> 

他们喜欢简单易用 —— 源代码非常易读，不需要任何文档或外部库。一切都非常简单。Vue.js “对任何东西都不做大的假设”。还有一个 [关于 GitLab 决定的播客节目]( https://link.juejin.im?target=https%3A%2F%2Fwww.youtube.com%2Fwatch%3Fv%3Dioogrvs2Ejc%23action%3Dshare ) 。

另一个来自 Pixeljets 的 [关于向 Vue 转变]( https://link.juejin.im?target=http%3A%2F%2Fpixeljets.com%2Fblog%2Fwhy-we-chose-vuejs-over-react%2F ) 的博文。React “是 JS 界在 [意识层面]( https://link.juejin.im?target=https%3A%2F%2Fen.wikipedia.org%2Fwiki%2FSingle_source_of_truth ) 向前迈出的一大步，它以一种实用简洁的方式向人们展示了真正的函数式编程。和 Vue 相比，React 的一大缺点是由于 JSX 的限制，组件的粒度会更小。这里是文章的引述：

> 
> 
> 
> 对于我和我的团队来说，代码的可读性是很重要的，但编写代码很有趣也是非常重要的。在实现真正简单的计算器小部件时创建 6
> 个组件并不奇怪。在许多情况下，在维护，修改或对某个小部件进行可视化检查方面也是不好的，因为你需要绕过多个文件/函数并分别检查每个小块的
> HTML。再次，我不是建议写巨石 - 我建议在日常开发中使用组件而不是微组件。
> 
> 

关于 [Hacker news]( https://link.juejin.im?target=https%3A%2F%2Fnews.ycombinator.com%2Fitem%3Fid%3D13151317 ) 和 [Reddit]( https://link.juejin.im?target=https%3A%2F%2Fwww.reddit.com%2Fr%2Fwebdev%2Fcomments%2F5ho71i%2Fwhy_we_chose_vuejs_over_react%2F ) 上的博客文章有趣的讨论 - 有来自 Vue 的持异议者和进一步支持者的争论。

### 状态管理和数据绑定 ###

构建用户界面很困难，因为处处都有状态 - 随着时间的推移而变化的数据带来了复杂性。定义的状态工作流程对于应用程序的增长和复杂性有很大的帮助。对于复杂度不大的应用程序，就不必定义的状态流了，像原生 JS 就足够了。

它是如何工作的？组件在任何时间点描述 UI。当数据改变时，框架重新渲染整个 UI 组件 - 显示的数据始终是最新的。我们可以把这个概念称为“ UI 即功能”。

React 经常与 Redux 在一起使用。 **Redux** 以三个 [基本原则]( https://link.juejin.im?target=http%3A%2F%2Fredux.js.org%2Fdocs%2Fintroduction%2FThreePrinciples.html ) 来自述：

* 单一数据源（Single source of truth）
* State 是只读的（State is read-only）
* 使用纯函数执行修改（Changes are made with pure functions）

换句话说：整个应用程序的状态存储在单个 store 的状态树中。这有助于调试应用程序，一些功能更容易实现。状态是只读的，只能通过 action 来改变，以避免竞争条件（这也有助于调试）。编写 Reducer 来指定如何通过 action 来转换 state。

大多数教程和样板文件都已经集成了 Redux，但是如果没有它，你可以使用 React（你可能不需要在你的项目中使用 Redux）。Redux 在代码中引入了复杂性和相当强的约束。如果你正在学习React，那么在你要使用 Redux 之前，你应该考虑学习纯粹的 React。你绝对应该阅读 [Dan Abramov]( https://link.juejin.im?target=https%3A%2F%2Fmedium.com%2F%40dan_abramov ) 的“ [你可能不需要 Redux]( https://link.juejin.im?target=https%3A%2F%2Fmedium.com%2F%40dan_abramov%2Fyou-might-not-need-redux-be46360cf367 ) ”。

[有些开发人员]( https://link.juejin.im?target=https%3A%2F%2Fnews.ycombinator.com%2Fitem%3Fid%3D13151577 ) 建议使用 **[Mobx]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fmobxjs%2Fmobx ) 代替 Redux** 。你可以把它看作是一个 “自动的 Redux”，这使得事情一开始就更容易使用和理解。如果你想了解，你应该从 [介绍]( https://link.juejin.im?target=https%3A%2F%2Fmobxjs.github.io%2Fmobx%2Fgetting-started.html ) 开始。你也可以阅读 Robin 的 [Redux 和 MobX 的比较]( https://link.juejin.im?target=https%3A%2F%2Fwww.robinwieruch.de%2Fredux-mobx-confusion%2F ) 。他还提供了有关 [从 Redux 迁移到 MobX]( https://link.juejin.im?target=https%3A%2F%2Fwww.robinwieruch.de%2Fmobx-react%2F ) 的信息。如果你想查找其他 Flux 库， [这个列表]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fvoronianski%2Fflux-comparison ) 非常有用。如果你是来自 MVC 的世界，那么你应该阅读 [Mikhail Levkovsky]( https://link.juejin.im?target=https%3A%2F%2Fmedium.com%2F%40mlovekovsky ) 的文章“ [Redux 中的思考（当你所知道的是 MVC）]( https://link.juejin.im?target=https%3A%2F%2Fmedium.com%2Fp%2Fthinking-in-redux-when-all-youve-known-is-mvc-c78a74d35133%3Fsource%3Duser_popover ) ”。

Vue 可以使用 Redux，但它提供了 [Vuex]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fvuejs%2Fvuex ) 作为自己的解决方案。

React 和 Angular 之间的巨大差异是 **单向与双向绑定** 。当 UI 元素（例如，用户输入）被更新时，Angular 的双向绑定改变 model 状态。React 只有一种方法：先更新 model，然后渲染 UI 元素。Angular 的方式实现起来代码更干净，开发人员更容易实现。React 的方式会有更好的数据总览，因为数据只能在一个方向上流动（这使得调试更容易）。

这两个概念各有优劣。你需要了解这些概念，并确定这是否会影响你选择框架。文章“ [双向数据绑定：Angular 2 和 React]( https://link.juejin.im?target=https%3A%2F%2Fwww.accelebrate.com%2Fblog%2Ftwo-way-data-binding-angular-2-and-react%2F ) ”和 [这个 Stackoverflow 上的问题]( https://link.juejin.im?target=https%3A%2F%2Fstackoverflow.com%2Fquestions%2F34519889%2Fcan-anyone-explain-the-difference-between-reacts-one-way-data-binding-and-angula ) 都提供了一个很好的解释。 [在这里]( https://link.juejin.im?target=http%3A%2F%2Fn12v.com%2F2-way-data-binding%2F ) 你可以找到一些交互式的代码示例（3 年前的示例（，只适用于 Angular 1 和 React）。最后，Vue 支持 [单向绑定和双向绑定]( https://link.juejin.im?target=https%3A%2F%2Fmedium.com%2Fjs-dojo%2Fexploring-vue-js-reactive-two-way-data-binding-da533d0c4554 ) （默认为单向绑定）。

如果你想进一步阅读，这有一篇长文，是有关状态的不同类型和 [Angular 应用程序中的状态管理]( https://link.juejin.im?target=https%3A%2F%2Fblog.nrwl.io%2Fmanaging-state-in-angular-applications-22b75ef5625f ) （ [Victor Savkin]( https://link.juejin.im?target=https%3A%2F%2Fmedium.com%2F%40vsavkin ) ）。

### 其他的编程概念 ###

Angular 包含依赖注入（dependency injection），即一个对象将依赖项（服务）提供给另一个对象（客户端）的模式。这导致更多的灵活性和更干净的代码。文章 “ [理解依赖注入]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fangular%2Fangular.js%2Fwiki%2FUnderstanding-Dependency-Injection ) ” 更详细地解释了这个概念。

[模型 - 视图 - 控制器模式]( https://link.juejin.im?target=https%3A%2F%2Fen.wikipedia.org%2Fwiki%2FModel%25E2%2580%2593view%25E2%2580%2593controller ) （MVC）将项目分为三个部分：模型，视图和控制器。Angular（MVC 模式的框架）有开箱即用的 MVC 特性。React 只有 V —— 你需要自己解决 M 和 C。

### 灵活性与精简到微服务 ###

你可以通过仅仅添加 React 或 Vue 的 JavaScript 库到你的源码中的方式去使用它们。但是由于 Angular 使用了 TypeScript，所以不能这样使用 Angular。

现在我们正在更多地转向微服务和微应用。React 和 Vue 通过只选择真正需要的东西，你可以更好地控制应用程序的大小。它们提供了更灵活的方式去把一个老应用的一部分从单页应用（SPA）转移到微服务。Angular 最适合单页应用（SPA），因为它可能太臃肿而不能用于微服务。

正如 [Cory House]( https://link.juejin.im?target=https%3A%2F%2Fmedium.com%2F%40housecor ) 所说:

> 
> 
> 
> JavaScript 发展速度很快，而且 React 可以让你将应用程序的一小部分替换成更好用的 JS 库，而不是期待你的框架能够创新。 **小巧，可组合的单一用途工具的理念永远不会过时**
> 。
> 
> 

有些人对非单页的网站也使用 React（例如复杂的表单或向导）。甚至 Facebook 都没有把 React 用在 Facebook 的主页，而是用在特定的页面，实现特定的功能。

### 体积和性能 ###

任何框架都不会十全十美：Angular 框架非常臃肿。gzip 文件大小为 143k，而 Vue 为 23K，React 为 43k。

为了提高性能，React 和 Vue 都使用了虚拟 DOM（Virtual DOM）。如果你对此感兴趣，可以阅读 [虚拟 DOM 和 DOM 之间的差异]( https://link.juejin.im?target=http%3A%2F%2Freactkungfu.com%2F2015%2F10%2Fthe-difference-between-virtual-dom-and-dom%2F ) 以及 [react.js 中虚拟 DOM 的实际优势]( https://link.juejin.im?target=https%3A%2F%2Fwww.accelebrate.com%2Fblog%2Fthe-real-benefits-of-the-virtual-dom-in-react-js%2F ) 。此外，虚拟 DOM 的作者之一在 Stackoverflow 上回答了 [性能的相关问题]( https://link.juejin.im?target=https%3A%2F%2Fstackoverflow.com%2Fquestions%2F21109361%2Fwhy-is-reacts-concept-of-virtual-dom-said-to-be-more-performant-than-dirty-mode ) 。

为了检查性能，我看了一下最佳的 [js 框架基准]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fkrausest%2Fjs-framework-benchmark ) 。你可以自己下载并运行它，或者查看 [交互式结果表]( https://link.juejin.im?target=http%3A%2F%2Fwww.stefankrause.net%2Fjs-frameworks-benchmark6%2Fwebdriver-ts-results%2Ftable.html ) 。

![](https://user-gold-cdn.xitu.io/2017/11/16/15fc436d20796233?imageView2/0/w/1280/h/960/ignore-error/1)

Angular，React 和 Vue 性能比较（ [源文件]( https://link.juejin.im?target=http%3A%2F%2Fwww.stefankrause.net%2Fjs-frameworks-benchmark6%2Fwebdriver-ts-results%2Ftable.html ) ）

![](https://user-gold-cdn.xitu.io/2017/11/16/15fc436e2c06ae4b?imageView2/0/w/1280/h/960/ignore-error/1)

内存分配（ [源文件]( https://link.juejin.im?target=http%3A%2F%2Fwww.stefankrause.net%2Fjs-frameworks-benchmark6%2Fwebdriver-ts-results%2Ftable.html ) ）

总结一下：Vue 有着很好的性能和高深的内存分配技巧。如果比较快慢的话，这些框架都非常接近（比如 [Inferno]( https://link.juejin.im?target=http%3A%2F%2Fwww.stefankrause.net%2Fjs-frameworks-benchmark6%2Fwebdriver-ts-results%2Ftable.html ) ）。请记住，性能基准只能作为考虑的附注，而不是作为判断标准。

### 测试 ###

Facebook [使用 Jest ]( https://link.juejin.im?target=http%3A%2F%2Ffacebook.github.io%2Fjest%2F ) 来测试其 React 代码。这里有篇 [Jest 和 Mocha 之间的比较]( https://link.juejin.im?target=https%3A%2F%2Fspin.atomicobject.com%2F2017%2F05%2F02%2Freact-testing-jest-vs-mocha%2F ) 的文章 —— 还有一篇关于 [Enzyme 和 Mocha 如何一起使用]( https://link.juejin.im?target=https%3A%2F%2Fsemaphoreci.com%2Fcommunity%2Ftutorials%2Ftesting-react-components-with-enzyme-and-mocha ) 的文章。Enzyme 是 Airbnb 使用的 JavaScript 测试工具（与 Jest，Karma 和其他测试框架一起使用）。如果你想了解更多，有一些关于在 React（ [这里]( https://link.juejin.im?target=https%3A%2F%2Fmedium.com%2F%40bruderstein%2Fthe-missing-piece-to-the-react-testing-puzzle-c51cd30df7a0 ) 和 [这里]( https://link.juejin.im?target=http%3A%2F%2Freactkungfu.com%2F2015%2F07%2Fapproaches-to-testing-react-components-an-overview%2F ) ）测试的旧文章。

Angular 2 中使用 **Jasmine** 作为测试框架。 [Eric Elliott]( https://link.juejin.im?target=https%3A%2F%2Fmedium.com%2F%40_ericelliott ) 在一篇文章中抱怨说 Jasmine “有数百种测试和断言的方式，需要仔细阅读每一个，来了解它在做什么”。输出也是非常臃肿和难以阅读。有关 Angular 2 [与 Karma]( https://link.juejin.im?target=https%3A%2F%2Fmedium.com%2F%40laco0416%2Fsetting-up-angular-2-testing-environment-with-karma-and-webpack-e9b833befd99 ) 和 [Mocha]( https://link.juejin.im?target=https%3A%2F%2Fmedium.com%2F%40PeterNagyJob%2Fangular2-configuration-and-unit-testing-with-mocha-and-chai-4ada9484e569 ) 的整合的一些有用的文章。这里有一个关于 [Angular 2 测试策略]( https://link.juejin.im?target=https%3A%2F%2Fwww.youtube.com%2Fwatch%3Fv%3DC0F2E-PRm44 ) 的旧视频（从2015年起）。

Vue 缺乏测试指导，但是 Evan 在 2017 年的展望中写道， [团队计划在这方面开展工作]( https://link.juejin.im?target=https%3A%2F%2Fmedium.com%2Fthe-vue-point%2Fvue-in-2016-8df71d98bfb3 ) 。他们推荐使用 [Karma]( https://link.juejin.im?target=http%3A%2F%2Fkarma-runner.github.io%2F1.0%2Findex.html ) 。 [Vue 和 Jest 结合使用]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Flocoslab%2Fvue-jest-utils ) ，还有 [avoriaz 作为测试工具]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Feddyerburgh%2Favoriaz ) 。

### 通用与原生 app ###

通用 app 正在将应用程序引入 web、搬上桌面，同样将深入原生 app 的世界。

React 和 Angular 都支持原生开发。Angular 拥有用于原生应用的 [NativeScript]( https://link.juejin.im?target=https%3A%2F%2Fdocs.nativescript.org%2Ftutorial%2Fng-chapter-0 ) （由 Telerik 支持）和用于混合开发的 Ionic 框架。借助 React，你可以试试 [react-native-renderer]( https://link.juejin.im?target=http%3A%2F%2Fangularjs.blogspot.de%2F2016%2F04%2Fangular-2-react-native.html ) 来构建跨平台的 iOS 和 Android 应用程序，或者用 [react-native]( https://link.juejin.im?target=https%3A%2F%2Ffacebook.github.io%2Freact-native%2F ) 开发原生 app。许多 app（包括 Facebook；查看更多的 [展示]( https://link.juejin.im?target=https%3A%2F%2Ffacebook.github.io%2Freact-native%2Fshowcase.html ) ）都是用 react-native 构建的。

Javascript 框架在客户端上渲染页面。这对于性能，整体用户体验和 SEO 是不利的。服务器端预渲染是一个好办法。所有这三个框架都有相应的库来实现服务端渲染。React 有 next.js，Vue 有 nuxt.js，而 Angular 有...... [Angular Universal]( https://link.juejin.im?target=https%3A%2F%2Funiversal.angular.io%2F ) 。

### 学习曲线 ###

Angular 的学习曲线确实很陡。它有全面的文档，但你仍然可能被吓哭，因为 [说起来容易做起来难]( https://link.juejin.im?target=https%3A%2F%2Fwww.reddit.com%2Fr%2Fwebdev%2Fcomments%2F5ho71i%2Fwhy_we_chose_vuejs_over_react%2Fdb1vppj%2F ) 。即使你对 Javascript 有深入的了解，也需要了解框架的底层原理。去初始化项目是很神奇的，它会引入很多的包和代码。因为有一个大的，预先存在的生态系统，你需要随着时间的推移学习，这很不利。另一方面，由于已经做出了很多决定，所以在特定情况下可能会很好。对于 React，你可能需要针对第三方库进行大量重大决策。仅仅 React 中就有 16 种 [不同的 flux 软件包来用于状态管理]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fvoronianski%2Fflux-comparison ) 可供选择。

Vue 学习起来很容易。公司转向 Vue 是因为它对初级开发者来说似乎更容易一些。这里有一片说他们团队为什么 [从 Angular 转到 Vue]( https://link.juejin.im?target=https%3A%2F%2Fmedium.com%2F%40Hemantisme%2Fmoving-from-angular-to-vue-a-vuetiful-journey-c29842ab2039 ) 的文章。 [另一位用户]( https://link.juejin.im?target=https%3A%2F%2Fnews.ycombinator.com%2Fitem%3Fid%3D13151716 ) 表示，他公司的 React 应用程序非常复杂，以至于新开发人员无法跟上代码。有了 Vue，初级和高级开发人员之间的差距缩小了，他们可以更轻松地协作，减少 bug，减少解决问题的时间。

有些人说他们用 React 做的东西比用 Vue 做的更好。如果你是一个没有经验的 Javascript 开发人员 - 或者如果你在过去十年中主要使用 jQuery，那么你应该考虑使用 Vue。转向 React 时，思维方式的转换更为明显。Vue 看起来更像是简单的 Javascript，同时也引入了一些新的概念：组件，事件驱动模型和单向数据流。这同样是很小的部分。

同时，Angular 和 React 也有自己的实现方式。它们可能会限制你，因为你需要调整自己的做法，才能顺畅的开发。这可能是一个缺点，因为你不能随心所欲，而且学习曲线陡峭。这也可能是一个好处，因为你在学习技术时必须学习正确的概念。用 Vue，你可以用老方法来做。这一开始可能会比较容易上手，但长此以往会出现问题。

在调试方面，React 和 Vue 的黑魔法更少是一个加分项。找出 bug 更容易，因为需要看的地方少了，堆栈跟踪的时候，自己的代码和那些库之间有更明显的区别。使用 React 的人员报告说，他们永远不必阅读库的源代码。但是，在调试 Angular 应用程序时，通常需要调试 Angular 的内部来理解底层模型。从好的一面来看，从 Angular 4 开始，错误信息应该更清晰，更具信息性。

### Angular, React 和 Vue 底层原理 ###

你想自己阅读源代码吗？你想看看事情到底是怎么样的吗？

可能首先要查看 Github 仓库: React（ [github.com/facebook/re…]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Ffacebook%2Freact ) ）、Angular（ [github.com/angular/ang…]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fangular%2Fangular ) ）和 Vue（ [github.com/vuejs/vue]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fvuejs%2Fvue ) ）。

语法看起来怎么样？ValueCoders [比较 Angular，React 和 Vue 的语法]( https://link.juejin.im?target=http%3A%2F%2Fwww.valuecoders.com%2Fblog%2Ftechnology-and-apps%2Fvue-js-comparison-angular-react%2F ) 。

在生产环境中查看也很容易 —— 连同底层的源代码。 [TodoMVC]( https://link.juejin.im?target=http%3A%2F%2Ftodomvc.com%2F ) 列出了几十个相同的 Todo 应用程序，用不同的 Javascript 框架编写 —— 你可以比较 [Angular]( https://link.juejin.im?target=http%3A%2F%2Ftodomvc.com%2Fexamples%2Fangularjs ) ， [React]( https://link.juejin.im?target=http%3A%2F%2Ftodomvc.com%2Fexamples%2Freact%2F%23%2F ) 和 [Vue]( https://link.juejin.im?target=http%3A%2F%2Ftodomvc.com%2Fexamples%2Fvue%2F ) 的解决方案。 [RealWorld]( https://link.juejin.im?target=https%3A%2F%2Frealworld.io%2F%23 ) 创建了一个真实世界的应用程序（中仿），他们已经准备好了 [Angular]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fgothinkster%2Fangular-realworld-example-app ) （4+）和 [React]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fgothinkster%2Freact-redux-realworld-example-app ) （带 Redux ）的解决方案。 [Vue]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fmchandleraz%2Frealworld-vue ) 的开发正在进行中。

你可以看到许多真实的 app，以下是 React 的方案：

* [Do]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2F1ven%2Fdo ) （一款很好用的笔记管理 app，用 **React 和 Redux** 实现）
* [sound-redux]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fandrewngu%2Fsound-redux ) （用 React 和 Redux 实现的 Soundcloud 客户端）
* [Brainfock]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2FBrainfock%2FBrainfock ) （用 React 实现的项目和团队管理解决方案）
* [react-hn]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Finsin%2Freact-hn ) 和 [react-news]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fechenley%2Freact-news ) （仿 Hacker news）
* [react-native-whatsapp-ui]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fhimanshuchauhan%2Freact-native-whatsapp-ui ) 和 [教程]( https://link.juejin.im?target=https%3A%2F%2Fwww.codementor.io%2Fcodementorteam%2Fbuild-a-whatsapp-messenger-clone-in-react-part-1-4l2o0waav ) （仿 Whatsapp 的 react-native 版）
* [phoenix-trello]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fbigardone%2Fphoenix-trello%2Fblob%2Fmaster%2FREADME.md ) （仿 Trello）
* [slack-clone]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Favrj%2Fslack-clone ) 和 [其他教程]( https://link.juejin.im?target=https%3A%2F%2Fmedium.com%2F%40benhansen%2Flets-build-a-slack-clone-with-elixir-phoenix-and-react-part-1-project-setup-3252ae780a1 ) (仿Slack)

以下是 Angular 版的 app：

* [angular2-hn]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fhousseindjirdeh%2Fangular2-hn ) 和 [hn-ng2]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fhswolff%2Fhn-ng2 ) （仿 Hacker News， [另一个由 Ashwin Sureshkumar 创建的很好的教程]( https://link.juejin.im?target=https%3A%2F%2Fmedium.com%2F%40Sureshkumar_Ash%2Fangular-2-hackernews-clone-dynamic-components-routing-params-and-refactor-340773d82e6f ) ）
* [Redux-and-angular-2]( https://link.juejin.im?target=https%3A%2F%2Fmedium.com%2F%40Sureshkumar_Ash%2Fangular-2-hackernews-clone-dynamic-components-routing-params-and-refactor-340773d82e6f ) （仿 Twitter）

以下是 Vue 版的 app：

* [vue-hackernews-2.0]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fvuejs%2Fvue-hackernews-2.0 ) 和 [Loopa news]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2FAngarsk8%2Floopa-news ) （仿Hacker News）
* [vue-soundcloud]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fmul14%2Fvue-soundcloud ) （Soundcloud 演示）

## 总结 ##

### 现在决定使用哪个框架 ###

React，Angular 和 Vue 都很酷，而且没有一个能明显的超过对方。相信你的直觉。 [最后一点有趣的玩世不恭的言辞]( https://link.juejin.im?target=https%3A%2F%2Fwildermuth.com%2F2017%2F02%2F12%2FWhy-I-Moved-to-Vue-js-from-Angular-2%23comment-3153455874 ) 可能会有助于你的决定：

> 
> 
> 
> 这个肮脏的小秘密就是大多数 “现代 JavaScript 开发” 与实际构建网站无关 ——
> 它正在构建可供构建可供人们使用的库或者包，这些人可以为编写教程和教授课程的人构建框架。我不确定任何人实际上正在为实际用户建立任何交互。
> 
> 

当然，这是夸张的，但是可能有一点点道理。是的，Javascript生态系统中有很多杂音。在你搜索的过程中，你可能会发现很多其他有吸引力的选项 —— 尽量不要被最新，最闪亮的框架蒙蔽。

### 我应该选什么？ ###

如果你在Google工作： **Angular**

如果你喜欢 TypeScript： **Angular（ [或 React]( https://link.juejin.im?target=https%3A%2F%2Fmedium.com%2F%40jrwebdev%2Fid-argue-that-if-you-love-typescript-then-react-may-be-a-better-choice-ceec950ee543 ) ）**

如果你喜欢面向对象编程（OOP）: **Angular**

如果你需要指导手册，架构和帮助： **Angular**

如果你在Facebook工作： **React**

如果你喜欢灵活性： **React**

如果你喜欢大型的技术生态系统： **React**

如果你喜欢在几十个软件包中进行选择： **React**

如果你喜欢JS和“一切都是 Javascript 的方法”： **React**

如果你喜欢真正干净的代码： **Vue**

如果你想要最平缓的学习曲线： **Vue**

如果你想要最轻量级的框架： **Vue**

如果你想在一个文件中分离关注点： **Vue**

如果你一个人工作，或者有一个小团队： **Vue（或 React）**

如果你的应用程序往往变得非常大： **Angular（或 React）**

如果你想用 react-native 构建一个应用程序： **React**

如果你想在圈子中有很多的开发者： **Angular 或 React**

如果你与设计师合作，并需要干净的 HTML 文件： **Angular or Vue**

如果你喜欢 Vue 但是害怕有限的技术生态系统： **React**

如果你不能决定，先学习 **React** ，然后 **Vue** ，然后 **Angular** 。

**所以，你做出选择了吗？**

![Yeeesss，你做到了！](https://user-gold-cdn.xitu.io/2017/11/16/15fc436e2c180643?imageView2/0/w/1280/h/960/ignore-error/1) Yeeesss，你做到了！

很好！阅读关于如何 **开始 Angular，React 或 Vue** 开发（即将推出，在 [Twitter]( https://link.juejin.im?target=http%3A%2F%2Fwww.twitter.com%2Fjensneuhaus%2F ) 上关注我的更新）。

### 更多资源 ###

* [React JS，Angular 和 Vue JS —— 快速开始和比较]( https://link.juejin.im?target=https%3A%2F%2Fwww.udemy.com%2Fangular-reactjs-vuejs-quickstart-comparison%2F ) （对这三个框架进行了 8 小时的介绍和比较)
* [Angular React（和 Vue）- DEAL破坏者]( https://link.juejin.im?target=https%3A%2F%2Fhackernoon.com%2Fangular-vs-react-the-deal-breaker-7d76c04496bc ) （一个简短但很好的比较 [Dominik T]( https://link.juejin.im?target=https%3A%2F%2Fmedium.com%2F%40dominik.t ) ）
* [Angular 2 和 React —— 终极之舞]( https://link.juejin.im?target=https%3A%2F%2Fmedium.com%2Fjavascript-scene%2Fangular-2-vs-react-the-ultimate-dance-off-60e7dfbc379c ) （ [Eric Elliott]( https://link.juejin.im?target=https%3A%2F%2Fmedium.com%2F%40_ericelliott ) 一个很好的比较）
* [React Angular Ember 和 Vue.js]( https://link.juejin.im?target=https%3A%2F%2Fmedium.com%2F%40gsari%2Freact-vs-angular-vs-ember-vs-vue-js-e186c0afc1be ) （ [Gökhan Sari]( https://link.juejin.im?target=https%3A%2F%2Fmedium.com%2F%40gsari ) 的三种框架的比较）
* [React 和 Angular]( https://link.juejin.im?target=https%3A%2F%2Fwww.sitepoint.com%2Freact-vs-angular%2F ) （两个框架的明确比较）
* [Vue 可以战胜 React 吗？]( https://link.juejin.im?target=https%3A%2F%2Frubygarage.org%2Fblog%2Fvuejs-vs-react-battle ) （很多代码示例的一个很好的比较）
* [10 个理由，为什么我从 Angular 转到 React]( https://link.juejin.im?target=https%3A%2F%2Fwww.robinwieruch.de%2Freasons-why-i-moved-from-angular-to-react%2F ) （Robin Wieruch 另一个很好的对比）
* [所有的JavaScript框架都很糟糕]( https://link.juejin.im?target=https%3A%2F%2Fmedium.com%2F%40mattburgess%2Fall-javascript-frameworks-are-terrible-e68d8865183e ) （ [Matt Burgess]( https://link.juejin.im?target=https%3A%2F%2Fmedium.com%2F%40mattburgess ) 对所有主要框架的大肆抨击）

**感谢您的关注。我忘了重要的事吗？你有不同的意见吗？我总是很高兴得到反馈。**

**在 Twitter 上关注我的更新和获取更多内容：** [@jensneuhaus]( https://link.juejin.im?target=http%3A%2F%2Fwww.twitter.com%2Fjensneuhaus%2F ) —— 🙌

> 
> 
> 
> [掘金翻译计划](
> https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fxitu%2Fgold-miner
> ) 是一个翻译优质互联网技术文章的社区，文章来源为 [掘金]( https://juejin.im ) 上的英文分享文章。内容覆盖 [Android](
> https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fxitu%2Fgold-miner%23android
> ) 、 [iOS](
> https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fxitu%2Fgold-miner%23ios
> ) 、 [React](
> https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fxitu%2Fgold-miner%23react
> ) 、 [前端](
> https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fxitu%2Fgold-miner%23%E5%89%8D%E7%AB%AF
> ) 、 [后端](
> https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fxitu%2Fgold-miner%23%E5%90%8E%E7%AB%AF
> ) 、 [产品](
> https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fxitu%2Fgold-miner%23%E4%BA%A7%E5%93%81
> ) 、 [设计](
> https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fxitu%2Fgold-miner%23%E8%AE%BE%E8%AE%A1
> ) 等领域，想要查看更多优质译文请持续关注 [掘金翻译计划](
> https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fxitu%2Fgold-miner
> ) 、 [官方微博](
> https://link.juejin.im?target=http%3A%2F%2Fweibo.com%2Fjuejinfanyi ) 、 [知乎专栏](
> https://link.juejin.im?target=https%3A%2F%2Fzhuanlan.zhihu.com%2Fjuejinfanyi
> ) 。
> 
>