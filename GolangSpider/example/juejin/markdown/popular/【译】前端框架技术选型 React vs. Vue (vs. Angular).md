# 【译】前端框架技术选型 React vs. Vue (vs. Angular) #

> 
> 
> 
> 这是该系列文章的第2部分：“Fundbox的前端技术选型”。第1部分介绍了Fundbox的技术现状以及我们重新设计它的动机。第2部分介绍了选择新框架背后的考虑：是迁移到React，Vue还是Angular。第3部分描述了我们如何从Angular迁移到Vue，同时保证我们产品发展路线不受影响。
> 
> 
> 

## Overview ##

重新考虑前端技术选型需要大量思考，讨论，决策，规划，管理和实施。我们首先需要做出的决定之一是选择一个前端框架来重新设计我们的产品。

我们研究了几个月来保证我们得出一个的更好决策。进行讨论，建立概念证明，与其他公司相关经验的同事进行面谈，并阅读大量在线材料。

在本文中，我将比较选型过程中的入围者。我们决定从以下几个框架中选出我们的下一个基础框架：Angular，React和Vue。

## 目标 ##

我们的目标是构建一个全新的，现代化的，快速可靠的平台，为我们当前和未来的所有前端应用程序提供服务。

## 候选框架 ##

* React
* Vue
* Angular

### Angular ###

Angular在我们的选型过程中被提前放弃了,主要由于以下两个主要原因（更详细的推理可以在“ [Why move then?]( https://link.juejin.im?target=https%3A%2F%2Fmedium.com%2Ffundbox-engineering%2Fstate-of-front-end-at-fundbox-ed3f6427f8e4%230e5b ) ” 这里阅读）：

* 

Angular是我们目前采用的框架，使用v1版本。Angular v2引入了许多改进，但它不向后兼容。这意味着升级到最新的Angular不亚于换到其他框架。这也导致了开发人员对这个框架失去了信心，使用Angular的想法在Fundbox乃至整个行业中大幅下降。

* 

Angular逐渐发展成为一个可以帮助构建复杂系统的大框架，但对于构建启动快速变化的UI却没那么有用。React和Vue更加轻量，组件化意味着小巧，自主，封装，因此易于重复使用。如果我们从头开始开发新的基础架构（而不需要从现有基础架构迁移），我们也可以考虑使用Angular。但在我们的例子中，它不合适。

无论如何，我们正在重写我们应用程序的重要部分，这是我们学习新技术的绝佳机会。通过这种方式，我们可以拓宽知识面，丰富开发人员的经验。

## React VS Vue ##

接下来留给我们选择的就只有React和Vue了，我们主要从以下几个方面对这两个框架进行对比：

* **学习曲线**
拥有丰富开发经验人员学习框架是否足够简单？
* **代码风格**
框架的代码和约定的可读性和直观性如何？
* **单个文件组件**
在框架中浏览和维护组件有多直观？
* **性能**
使用框架构建的应用程序的性能如何？
框架的大小和运行内存占用有多大？
* **灵活性**
框架提供了多少功能？
有多少功能是强制性的？
定制框架有多容易？
* **工具**
框架可用的工具有哪些？
框架有多少稳定的插件？
* **移动端支持**
框架是否支持除Web以外的更多应用程序？
它是否提供了构建Native应用程序的方法？
* **社区**
框架的社区有多大？
社区是团结还是支离破碎的？
* **成熟度**
框架有多成熟？
它经过了多长时间的生产验证？
它的未来发展路线是否清晰？
* **支持**
框架背后的团队有多大？
是否有公司或组织负责该框架？
* **招聘人才**
聘请具有该框架使用经验的开发人员有多容易？

## 学习曲线 ##

### React ###

React官方文档提供了一些编写得很好的入门指南，并为新手提供了一个入门演练教程。具有一些前端框架经验的开发人员可以在几个小时内理解框架的核心原则。

官方的React文档很详细，但不像Vue的官方文档那样清晰有序。文档涵盖了必要的入门教程和核心概念等，但文档中缺少介绍框架的边界。随着项目变得更大，这些边界可能转换为痛点。

` React不是一个包含所有功能的框架。它的核心理念是精益，只关注于核心部分，其他功能通过引用第三方解决方案解决。` 这增加了学习曲线的一些复杂性，因为它根据您在整个过程中对框架所做的选择而有所不同。

### Vue ###

Vue可以直接在HTML页面中通过资源的方式加载，只需几分钟，整个库无需构建便可以使用了。这让我们可以在任何时候编写Vue应用程序。

> 
> 
> 
> 编辑：感谢David Bismut指出React也有办法将其作为独立的JS文件添加到页面中，无需构建步骤： [https//reactjs.org/docs/add-react-to-a-website.html](
> https://link.juejin.im?target=https%2F%2Freactjs.org%2Fdocs%2Fadd-react-to-a-website.html
> )
> 
> 

因为Vue借鉴了React和Angular的一些概念，开发人员更容易学习Vue。Vue的官方文档编写得非常好，甚至涵盖了开发Vue应用程序过程中偶然发现的问题。

Vue的定义比React更严格; 这也意味着它更具稳定性。值得注意的是，在Vue中，许多问题直接在其文档中得到解答，而不需要在其他地方搜索。

## 代码风格 ##

### React ###

React引入了一系列基于函数式编程的概念，简化了开发UI优先应用程序的过程。最值得注意的是：

* JSX，这是一种用JavaScript编写HTML的方法。JSX是React作为函数式编程的强大推动者的补充，具有重大意义。
* 它的组件生命周期提供了一种直观的方式来连接组件“生命”中的特定事件（创建，更新等）

### Vue ###

作为一个比React和Angular更年轻的框架，Vue从各个方面吸收了很多好的理念，混合了函数式和面向对象编程(OOP)。

默认情况下，Vue的编码风格在某些方面与Angular有点相似，但也消除了Angular的大部分痛点。Vue将HTML，JS和CSS分开，就像Web开发人员已经习惯了多年的编码风格。但假如你喜欢JSX这种风格，它也可以替换使用。所以它不会强制改变你的代码风格。

Vue对组件生命周期的考虑比React更直观。一般来说，Vue的API比React更宽泛但更简单。

## 单个文件组件 ##

### React ###

使用JSX，React中的单个文件组件完全是作为JavaScript模块编写的，意味着React具有编写HTML，CSS和JavaScript的特定方式。

所有功能采用JavaScript编写，意味着更少的bug，因为它减轻了在组件内部创建动态HTML的负担。作为替代，使用JSX时我们可以使用原生JavaScript生成模板。

也就是说，React的特殊语法需要更多的学习和练习才能在React中编写组件。

### Vue ###

Vue中的单个文件组件分为三个独立的部分： ` <template>` ， ` <script>` 和 ` <style>` ，每个都包含相应类型的代码，因此Web开发人员感觉更自然。

作为一个渐进式框架，Vue可以轻松定制。例如，作为可配置项，可以使用JSX而不是 ` <template>` 模板。 另外，另一个例子，只需在 ` <style>` 标记中添加 ` lang =“scss”` 属性，就可以编写SCSS而不是纯CSS。类似地，通过将 ` scoped` 属性添加到 ` <style>` 标记，Vue组件将实现开箱即用的CSS（也称为CSS模块）。

### 性能 ###

> 
> 
> 
> 编辑：值得一提的是，性能有点难以衡量，因为它在很大程度上取决于您正在构建的特定应用程序，甚至更多取决于您如何构建它。
> 
> 

### React ###

库大小（通过网络/未压缩）：32.5KB / 101.2KB

比较DOM操作，React的整体性能非常好。它比Angular快得多，但比Vue慢一点。

React提供了对开箱即用的服务器端渲染（SSR）的支持，并且可能对某些类型的实现很有用。

内置支持 ` bundling` 和 ` tree-shaking` ，最大限度地减少最终用户的资源负担。

### Vue ###

库大小（通过网络/未压缩）：31KB / 84.4KB

除了成为最快的Vue之外，Vue还是一个渐进式框架，从头开始构建以逐步采用。核心库仅专注于视图层，易于获取并与其他库或现有项目集成。

与React类似，Vue内置了对 ` bundling` 和 ` tree-shaking` 的支持，可最大限度地减少最终用户的资源负担。

![](https://user-gold-cdn.xitu.io/2019/5/31/16b0d3cc1a6a4913?imageView2/0/w/1280/h/960/ignore-error/1) [www.stefankrause.net/js-framewor…]( https://link.juejin.im?target=https%3A%2F%2Fwww.stefankrause.net%2Fjs-frameworks-benchmark7%2Ftable.html )

## 灵活性 ##

### React ###

React专注于UI，因此我们能从它上面获取的最根本功能是构建用户界面。

更高级的功能不在React官方库的提供的功能范围之内，如状态管理。大多数React应用程序都使用Redux进行状态管理，MobX作为React伴侣也倍受关注。

甚至React路由器也不是官方软件包，而是由React团队支持的第三方软件包。

### Vue ###

作为一个渐进式框架，Vue只允许使用其最基本的功能来构建应用程序，但如果需要，它还提供您开箱即用的大部分内容：用于状态管理的Vuex，用于应用程序URL管理的Vue路由器，Vue-SSR 用于服务器端渲染。

无论好坏，Vue比React更稳固。

## 工具 ##

### React ###

React有一个名为 ` create-react-app` 的第三方CLI工具，可帮助在React项目中构建新的应用程序和组件。

CLI工具还支持运行端到端和单元测试，代码检查和本地开发服务器功能。

React为主要IDE提供了很好的官方和社区支持。

### Vue ###

Vue有一个名为 ` Vue CLI` 的官方CLI工具。与React的create-react-app非常相似，Vue CLI工具提供了新应用程序的支架。

此外，Vue对所有重要的IDE都有很好的支持（不如React，但支持WebStorm和VSCode）。

## 移动端支持 ##

### React ###

React有一个用于构建本机移动应用程序的端口，它被称为 ` React Native` ，它是当前的“write once (in JavaScript), use many (in native iOS and Android)”的解决方案。

有大量的线上应用是使用React Native构建的。

### Vue ###

对于Vue，构建Mobile Native应用程序的方式有很多种。不像 ` React Native` ，使用Vue构建Native应用并没有明确的引领者。

` NativeScript` 是这些选项的领先者（它也是Angular的优先解决方案），除此之外还有 ` Weex` 和 ` Quasar` 。

## 社区 ##

### React ###

在StackOverflow中，有大约 ` 88,000` 个问题用 ` #reactjs` 标记。 有超过40,000个npm软件包可供React开发人员安装。

在最新的前端工具调查中，超过40％的受访者表示他们对使用React感到满意。

在GitHub中，React repo拥有近100,000颗星。

React的社区确实更为庞大，但缺点是比Vue更加分散，而且很难找到常见问题的直接答案。

### Vue ###

在StackOverflow中，有 ` 18,000` 个标记为#vue 有近10,000个npm软件包可供安装。

在最新的调查中，17％的受访者表示他们对使用Vue感到满意。但事实上，与React相比，对学习Vue感兴趣的开发人数增加了一倍，因此Vue开发人员的市场需求在未来可能会增长得比React更快。

GitHub中的 ` Vue` 项目start数已经超过了100,000，超过了React。

由于其出色的文档，Vue开发中的问题的大多数答案可以立即在文档中找到，并且跟社区答案基本是一致的。

## 成熟度 ##

### React ###

React于2013年3月发布（撰写本文时为5年）。

据SimilarTech称，React正在被205,000个独立网站使用，每月增长2.46％。

React在生产方面经过了很好的测试，超过了Vue。React建立了一个庞大的社区，这主要得益于其所有者Facebook的维护。

### Vue ###

Vue于2014年2月发布（撰写本文时为4岁）。

根据SimilarTech，Vue正在被26,000个独立网站使用。每月增长3.34％。

Vue在一年半前成为业界标杆并被广泛使用，包括一些大公司，如GitLab，阿里巴巴，百度等。事实证明，Vue的运转和更新都是稳定的。

## 支持 ##

### React ###

React是由Facebook创建和维护的框架。在Facebook，有一个团队定期支持React（React也被用于Facebook内的许多项目）。

据称，Facebook的React团队规模包括10名专职开发人员。但值得注意的是，Facebook研发部门的多个团队正在将React用于内部和外部项目，并且每个团队都可以将变更请求推送到库中。

React不具有实际的路线图，基于RFC规则，具体解释在 [这里]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Freactjs%2Frfcs ) 。

### Vue ###

Vue是一个独立的库，由 ` Evan You` 创立。他还负责管理Vue的维护及其路线图。

Vue团队规模包括23名专职开发人员。

Vue的高级路线图可以在他们的GitHub仓库中查看。

## 人才招聘 ##

### React ###

作为目前最流行的框架，如果React开发人员的市场中，React具有优势。

此外，通过学习React，开发人员的简历更具闪光点，因为他们很容易从这个流行框架中获得相关的宝贵经验。

### Vue ###

Vue是前端行业的新“热点”。当然，炒作也有一些缺点，但Vue长期以来一直在获得稳定的牵引力，开发人员急于将Vue项目作为 [错失恐惧症-FOMO]( https://link.juejin.im?target=https%3A%2F%2Fen.wikipedia.org%2Fwiki%2FFear_of_missing_out ) 的一部分。现在，找到有Vue经验的开发人员并不难。

## 各自优点 ##

### React ###

* 行业标准。
* 受欢迎，熟练使用框架的前端开发工程师更多。
* 更容易招聘到优秀的工程师。
* 拥有强大背景的公司支持，更加安全的未来和稳定。
* 更庞大的社区，大量的工具和包。
* web和移动应用可以共享一些代码。

### Vue ###

* 内置核心模块（Vuex，路由器等）并且运转非常棒。
* 面向“未来”，而不是“当前”。
* 更独特; 引领潮流而不是跟随它。
* 更快的上手，FED(前端开发)和BED(后端开发)都会在Vue代码中感觉很自然，速度很快。
* 更好全栈支持; 允许跨产品开发。

## 各自缺点 ##

### React ###

* 保持FED和BED之间的界限; React需要学到很多才能成为专家。
* 需要更多的时间训练开发者。
* 交付较慢（至少对于全新启动的复杂项目来说）。

### Vue ###

* 走进一个更具实验性的土地，没有风险，但前卫。
* 更难找到经验丰富的Vue开发者。
* 可用的插件和工具更少，社区更小。
* 不像React，开发者无法获得最流行框架的相关经验。

## 更多 ##

很长一段时间里，React和Angular是框架游戏中的主要参与者，而许多（很多！）框架每隔一天都会出现，并试图进入名人堂却没有成功; 除了-Vue。Vue之所以吸引人和受欢迎可以在文章，教程，POC和浏览器开发者社区的参考文献这几个方面看到。

React是行业中的潮流引领者和高级技能。React在今天是一个明确的引领者，无论是在行业使用上面还是社区受欢迎程度。它非常受欢迎的，说实话，它的高地位是应得的。 React可以轻松构建复杂而直观的Web和移动应用程序，但它具有成本 - 框架复杂性和样例复杂。基础知识相对直观，但大型React项目往往变得复杂。社区中的碎片化也是它一大确定。React引入许多新规范对其学习曲线有一些负面影响。

Vue更精简，是一个直接且新颖的框架，值得在舞台上占据一席之地，因为它非常简单易学，样例代码非常简单，性能高，灵活且完整。今天的许多网络应用程序可以使用Vue比使用React更快地构建。Vue很有趣，开发很简单。

最近前端社区内赞扬Vue的讨论在稳定的增长，意味着Vue将很快变得像React一样受欢迎。

## 参看资料 ##

* [How popular is VueJS in the industry? Will becoming a Vue expert be useful, career-wise?]( https://link.juejin.im?target=https%3A%2F%2Fwww.quora.com%2FHow-popular-is-VueJS-in-the-industry-Will-becoming-a-Vue-expert-be-useful-career-wise ) (Quora)
* [Comparison with Other Frameworks]( https://link.juejin.im?target=https%3A%2F%2Fvuejs.org%2Fv2%2Fguide%2Fcomparison.html ) (in Vue’s documentation)
* [Vue.js or React? Which you would chose and why?]( https://link.juejin.im?target=https%3A%2F%2Fwww.reddit.com%2Fr%2Fjavascript%2Fcomments%2F8o781t%2Fvuejs_or_react_which_you_would_chose_and_why%2Fe01qn55%2F ) (Reddit discussion)
* [Choosing the Right JavaScript Framework for Your Next Web Application ]( https://link.juejin.im?target=https%3A%2F%2Fwww.telerik.com%2Fcampaigns%2Fkendo-ui%2Fwp-choosing-js-framework%3Futm_medium%3Dcpm%26amp%3Butm_source%3Dcarbonads%26amp%3Butm_campaign%3Dkendo-ui-jquery-whitepaper-choosejsframework ) (Progress)
* [Angular vs. React vs. Vue: A 2017 comparison]( https://link.juejin.im?target=https%3A%2F%2Fmedium.com%2Funicorn-supplies%2Fangular-vs-react-vs-vue-a-2017-comparison-c5c52d620176 ) (Medium article)
* [Can Vue Fight for the Throne with React?]( https://link.juejin.im?target=https%3A%2F%2Frubygarage.org%2Fblog%2Fvuejs-vs-react-battle ) (RubyGarage blog)
* [React vs Vue.JS: Which Front-end Framework to Choose in 2018]( https://link.juejin.im?target=https%3A%2F%2Fexpertise.jetruby.com%2Freact-vs-vue-js-which-front-end-framework-to-choose-in-2018-2a62a1fe76f9 ) (jetruby blog)
* [Angular 5 vs. React vs. Vue]( https://link.juejin.im?target=https%3A%2F%2Fitnext.io%2Fangular-5-vs-react-vs-vue-6b976a3f9172 ) (ITNEXT blog)

感谢阅读

原文： [React vs. Vue (vs. Angular)]( https://link.juejin.im?target=https%3A%2F%2Fmedium.com%2Ffundbox-engineering%2Freact-vs-vue-vs-angular-163f1ae7be56 )