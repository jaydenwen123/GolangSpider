# [译] Python 架构相关：我们需要更多吗？ #

> 
> 
> 
> * 原文地址： [Python Architecture Stuff: do we need more?](
> https://link.juejin.im?target=http%3A%2F%2Fwww.obeythetestinggoat.com%2Fpython-architecture-stuff-do-we-need-more.html
> )
> * 原文作者： [Harry](
> https://link.juejin.im?target=http%3A%2F%2Fwww.obeythetestinggoat.com%2Fauthor%2Fharry.html
> )
> * 译文出自： [掘金翻译计划](
> https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fxitu%2Fgold-miner
> )
> * 本文永久链接： [github.com/xitu/gold-m…](
> https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fxitu%2Fgold-miner%2Fblob%2Fmaster%2FTODO1%2Fpython-architecture-stuff-do-we-need-more.md
> )
> * 译者： [QiaoN](
> https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2FQiaoN )
> * 校对者：
> 
> 
> 

最近，我一直在学习有关应用程序架构的很多新东西。亲爱的读者，我想知道你是否对这些想法感兴趣，以及我们是否应该尝试围绕它构建更多的资源（博客、会谈等）。

## 我们应该如何构建一个应用程序来充分利用测试？ ##

对我而言一切都始于这个问题。在 [我的书尾]( https://link.juejin.im?target=https%3A%2F%2Fwww.obeythetestinggoat.com%2Fbook%2Fchapter_hot_lava.html ) ，我用结束章节讨论了如何充分利用你的测试，在单元、集成和端到端的测试中做出权衡，并对我没有真正理解的一些主题做出些模糊浅显的指示：端口与适配器（ports and adapters），六角架构（hexagonal architecture），函数式内核/命令式外壳（functional core imperative shell），干净架构（the clean architecture），等等。

从那之后，我和一个正在用 Python 积极实现这些模式的 [技术团队]( https://link.juejin.im?target=https%3A%2F%2Fio.made.com%2F ) 达成合作。其实，这些架构模式并不是什么新鲜事，人们多年来一直在用 Java 和 C# 进行探索。只是我对它们很陌生……从个人经验而言，我在这里可能会有些深入（我对你的反应很感兴趣），但它们对 Python 社区的大部分人可能也是个新鲜事？

随着我们的成熟，确实能感觉到越来越多的曾经的小项目和大胆的初创公司变成了复杂的业务和（偷偷的说）商业软件，所以这些东西可能会变得越来越显著。

我最初从测试的角度看待它，正确的架构真得可以帮助你充分利用测试，通过分离出一个业务逻辑核心（“领域模型”）并让其摆脱所有的架构依赖，它能完全的用快速、灵活的单元测试进行测试。最终让人感觉 [测试金字塔]( https://link.juejin.im?target=https%3A%2F%2Fmartinfowler.com%2Farticles%2Fpractical-test-pyramid.html ) 是一个可实现的目标，而非一个奢望。

## 关于该主题的经典书籍（均为 Java） ##

[Evans 的领域驱动设计（DDD）]( https://link.juejin.im?target=https%3A%2F%2Fdomainlanguage.com%2Fddd%2F ) 和 [Fowler 的架构模式]( https://link.juejin.im?target=https%3A%2F%2Fwww.martinfowler.com%2Fbooks%2Feaa.html ) 都是很经典的书籍，任何对这些感兴趣的人都应该阅读。但如果你像我一样，费力阅读那些 public static void main AbstractFactoryManager之类的东西实在让人有点烦。也许一些更轻量级的、Python 化的介绍能让人感觉更加合理，少点云里雾里？

## Python 领域中的一些现有资源： ##

Made 的首席架构师，尊敬的 Bob 先生，就我们现在讨论的问题写过一个分为四部分的博客系列。我开始时特别喜欢阅读它。这系列是 DDD 基本概念、端口与适配器/依赖倒置、和某种程度上的事件驱动架构的快速使用介绍。都是 Python 适用。（触发警告：Type Hints）。

* [Python 中具有命令处理模式（Command Handler pattern）的端口和适配器]( https://link.juejin.im?target=https%3A%2F%2Fio.made.com%2Fintroducing-command-handler%2F )
* [Python 中的库（Repository）和工作单元模式]( https://link.juejin.im?target=https%3A%2F%2Fio.made.com%2Frepository-and-unit-of-work-pattern-in-python%2F )
* [命令和查询，处理程序（Handler）和视图]( https://link.juejin.im?target=https%3A%2F%2Fio.made.com%2Fcommands-and-queries-handlers-and-views%2F )
* [为什么要用领域事件（Domain Events）？]( https://link.juejin.im?target=https%3A%2F%2Fio.made.com%2Fwhy-use-domain-events%2F )

在 io.made.com 上还有很多，但以上四篇为主要内容。我们希望得到一些关于它们的反馈，哪些被阐述到了，哪些需要进一步解释，等等。

另：一个去年圣诞节及时发布的的书，Leonardo Giordani 的 [Python 干净架构（Clean Architectures in Python）]( https://link.juejin.im?target=https%3A%2F%2Fleanpub.com%2Fclean-architectures-in-python ) 。这本书是两本书合二为一，第一部分是 TDD 的介绍，但第二部分有四章介绍了与我在这里讨论的类似的模式。

我也很喜欢一年前 David Seddon 的一个演讲 [岩石河：如何构建你的 Django 单体应用（monolith）]( https://link.juejin.im?target=http%3A%2F%2Fseddonym.me%2Ftalks%2F2017-12-12-rocky-river%2F ) ，显示出其他人开始思考我们如何超越基本的 Django 模型/视图/模板架构。

在 Valentin Ignatev 的 [ DDD 资源列表]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fvalignatev%2Fddd-dynamic ) 还有更多的内容，这是我最近在推特上看到的。似乎很多人对此都有想法。

## 号召行动：这个东西有趣吗？ ##

Bob已经得到一些对他博客帖子很好的反馈，Leonardo 也有了一些很不错的初始销量，所以我感觉到 Python 社区的一些兴趣，但是我想对它进行一个理性考察。

* 这些东西有趣或者有意义吗？你要了解更多吗？
* 你用 Python 正在做的事超出 “基础网页应用开发” 或 “数据管道（Data Pipeline）”的范围了吗？你是否发现编写快速单元测试很困难？你是否开始想把你的业务逻辑从任何你使用的框架中解放出来？
* 你是否在使用 DDD 或任何 Python 的经典模式？你可能已有所有的答案，愿意告诉我吗？或者你只想告诉我一些适用你的答案和事情？
* 你认为这些东西听起来很抽象且没有意义吗？也许 Made.com 在 Python 领域里有点像大纲，因为我们用 Python 编写物流/ERP/企业软件，而这一切和你日常工作非常不同吗？
* 从这些主题的新指南而言，你认为 Python/动态语言社区最受益的是什么？

我很乐意听到你的意见。文末评论开放，或者你也可以在推特上给我留言 [@hjwp]( https://link.juejin.im?target=https%3A%2F%2Ftwitter.com%2Fhjwp ) 。

## 我说过读经典书籍了没？去阅读经典吧。 ##

* 企业应用架构模式（Patterns of Enterprise Architecture），Martin Fowler 著， [amazon.com]( https://link.juejin.im?target=https%3A%2F%2Famzn.to%2F2U6HTZN ) / [.co.uk]( https://link.juejin.im?target=https%3A%2F%2Famzn.to%2F2R0WkN3 )
* 领域驱动设计（Domain-Driven Design），Eric Evans 著， [amazon.com]( https://link.juejin.im?target=https%3A%2F%2Famzn.to%2F2U6HTZN ) / [.co.uk]( https://link.juejin.im?target=https%3A%2F%2Famzn.to%2F2R0WkN3 )

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