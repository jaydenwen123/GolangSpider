# [译] 使用 VS Code 调试 Node.js 的超简单方法 #

> 
> 
> 
> * 原文地址： [The Absolute Easiest Way to Debug Node.js — with VSCode](
> https://link.juejin.im?target=https%3A%2F%2Fitnext.io%2Fthe-absolute-easiest-way-to-debug-node-js-with-vscode-2e02ef5b1bad
> )
> * 原文作者： [Paige Niedringhaus](
> https://link.juejin.im?target=https%3A%2F%2Fmedium.com%2F%40paigen11 )
> * 译文出自： [掘金翻译计划](
> https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fxitu%2Fgold-miner
> )
> * 本文永久链接： [github.com/xitu/gold-m…](
> https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fxitu%2Fgold-miner%2Fblob%2Fmaster%2FTODO1%2Fthe-absolute-easiest-way-to-debug-node-js-with-vscode.md
> )
> * 译者： [iceytea](
> https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Ficeytea%2F )
> * 校对者： [fireairforce](
> https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Ffireairforce ) , [cyz980908](
> https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fcyz980908 )
> 
> 
> 

> 
> 
> 
> 让我们面对现实吧...调试 Node.js 一直是我们心中的痛。
> 
> 

![](https://user-gold-cdn.xitu.io/2019/5/5/16a8711f67baa36b?imageView2/0/w/1280/h/960/ignore-error/1)

### 触达调试 Node.js 的痛点 ###

如果你曾经有幸为 Node.js 项目编写代码，那么当我说调试它以找到出错的地方并不是最简单的事情时，你就知道我在谈论什么。

不像浏览器中的 JavaScript，也不像有类似 IntelliJ 这样强大的 IDE 的 Java，你无法到处设置断点，刷新页面或者重启编译器，也无法慢慢审阅代码、检查对象、评估函数、查找变异或者遗漏的变量等。你无法那样去做，这简直太糟糕了。

但 Node.js 也是可以被调试的，只是需要多费些体力。让我们认真讨论这些可选方法，我会展示给你在我开发经历中遇到的最简单调试方法。

### 调试 Node.js 的一些可选方法 ###

有一些方式能调试有问题的 Node.js 程序。我把这些方法（包含详细链接）都列在了下面。如果你感兴趣，可以去了解下。

* **` Console.log()`** — 如果你曾经编写过 JavaScript 代码，那么这个可靠的备用程序真的不需要进一步解释。它被内置在 Node.js 并在终端中打印，就像内置到 JavaScript，并在浏览器控制台中打印一样。

在 Java 语言下，它是 ` System.out.println()` 。在 Python 语言下，它是 ` print()` 。你明白我的意思了吧。这是最容易实现的方法，也是用额外的行信息来“弄脏”干净代码的最快方法 —— 但它（有时）也可以帮助你发现和修复错误。

* **Node.js 文档 ` —-inspect`** — Node.js 文档撰写者本身明白调试不大简单，所以他们做了一些 [方便的参考]( https://link.juejin.im?target=https%3A%2F%2Fnodejs.org%2Fen%2Fdocs%2Fguides%2Fdebugging-getting-started%2F ) 帮助人们开始调试。

这很有用，但是老实说，除非你已经编写了一段时间的程序，否则它并不是最容易破译的。它们很快就进入了 UUIDs、WebSockets 和安全隐患的陷阱，我开始感到无所适从。我心里想：一定有一种不那么复杂的方法来做这件事。

* **Chrome DevTools** — [Paul Irish]( https://link.juejin.im?target=undefined ) 在 2016 年撰写了一篇有关使用 Chrome 开发者工具调试 Node.js 的 [博文]( https://link.juejin.im?target=https%3A%2F%2Fmedium.com%2F%40paul_irish%2Fdebugging-node-js-nightlies-with-chrome-devtools-7c4a1b95ae27 ) （并在 2018 年更新）。它看起来相当简单，对于调试来说是一个很大的进步。

半个小时之后，我仍然没有成功地将 DevTools 窗口连接到我的简单 Node 程序上，我不再那么肯定了。也许我只是不能按照说明去做，但是 Chrome DevTools 似乎让调试变得比它应该的更复杂。

* **JetBrains** — JetBrains 是我最喜欢的软件开发公司之一，也是 IntelliJ 和 WebStorm 的开发商之一。他们的工具有一个奇妙的插件生态系统，直到最近，他们还是我的首选 IDE。

有了这样一个专业用户基础，就出现了许多有用的文章，比如 [这一篇]( https://link.juejin.im?target=https%3A%2F%2Fwww.jetbrains.com%2Fhelp%2Fwebstorm%2Frunning-and-debugging-node-js.html ) ，它们调试 Node，但与 Node 文档和 Chrome DevTools 选项类似，这并不容易。你必须创建调试配置，附加正在运行的进程，并在 WebStorm 准备就绪之前在首选项中进行大量配置。

* **Visual Studio Code** — 这是我新的 Node 调试黄金标准。我从来没有想过我会这么说，但是我完全投入到 [VS Code]( https://link.juejin.im?target=https%3A%2F%2Fcode.visualstudio.com%2Fdownload ) 中，并且团队所做的每一个新特性的发布，都使我更加喜爱这个 IDE。

VS Code 做了其他所有选项在 [调试 Node.js]( https://link.juejin.im?target=https%3A%2F%2Fcode.visualstudio.com%2Fdocs%2Fnodejs%2Fnodejs-debugging%23_attaching-to-nodejs ) 都没能做到的事情，这让它变得傻瓜式简单。如果你想让你的调试变得更高级，这当然也是可以的，但是他们把它分解得足够简单，任何人都可以快速上手并运行，不论你对 IDE、Node 和编程的熟练度如何。这太棒了。

### 配置 VS Code 来调试 Node.js ###

![](https://user-gold-cdn.xitu.io/2019/5/5/16a8711f63b58e17?imageView2/0/w/1280/h/960/ignore-error/1)

好吧，让我们来配置 VS Code 来调试 Node。我假设你已经从 [这里]( https://link.juejin.im?target=https%3A%2F%2Fcode.visualstudio.com%2Fdownload ) 下载了 VS Code，开始配置它吧。

打开 ` Preferences > Settings` ，在搜索框中输入 ` node debug` 。在 ` Extensions` 选项卡下应该会有一个叫 ` Node debug` 的扩展。在这里点击第一个方框： **Debug > Node: Auto Attach** ，然后设置下拉框的选项为 ` on` 。你现在几乎已经配置完成了。是的，这相当的简单。

![这是当你点击 Settings 选项卡，你应该能看到的内容。设置第一个下拉框 **Debug > Node: Auto Attach** 选项为 `on`。](https://user-gold-cdn.xitu.io/2019/5/5/16a8711f67f3e2d6?imageView2/0/w/1280/h/960/ignore-error/1)

现在进入项目文件，然后通过点击文件的左侧边栏，在你想要看到代码暂停的地方设置一些断点。在终端内输入 ` node --inspect <FILE NAME>` 。现在看，神奇的事情发生了...

![看到红色断点了吗？看到终端中的 `node — inspect readFileStream.js` 了吗？就像这样。](https://user-gold-cdn.xitu.io/2019/5/5/16a8711f83af538e?imageView2/0/w/1280/h/960/ignore-error/1)

**VS Code 正在进行的代码调试**

如果你需要一个 Node.js 项目来测试它，可以 [在这里]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fpaigen11%2Ffile-read-challenge ) 下载我的 repo。它是用来测试使用 Node 传输大量数据的不同形式的，但是它在这个演示中非常好用。如果你想了解更多关于流数据节点和性能优化的内容，你可以点击 [这里]( https://link.juejin.im?target=https%3A%2F%2Fitnext.io%2Fusing-node-js-to-read-really-really-large-files-pt-1-d2057fe76b33 ) 和 [这里]( https://link.juejin.im?target=https%3A%2F%2Fitnext.io%2Fstreams-for-the-win-a-performance-comparison-of-nodejs-methods-for-reading-large-datasets-pt-2-bcfa732fa40e ) 。

当你敲击 ` Enter` 键时，你的 VS Code 终端底部会变成橙色，表示你处于调试模式，你的控制台会打印一些类似于 ` Debugger Attached` 的信息。

![橙色的工具栏和 `Debugger attached` 消息会告诉你 VS Code 正常运行在调试模式。](https://user-gold-cdn.xitu.io/2019/5/5/16a8711f86275df4?imageView2/0/w/1280/h/960/ignore-error/1)

当你看到这一幕发生时，恭喜你，你已经让 Node.js 运行在调试模式下啦！

至此，你可以在屏幕的左下角看到你设置的断点（而且你可以通过复选框切换这些断点的启用状态），而且，你可以像在浏览器中那样去调试。在 IDE 的顶部中心有小小的继续、步出、步入、重新运行等按钮，从而逐步完成代码。VS Code 甚至用黄色突出显示了你已经停止的断点和行，使其更容易被跟踪。

![单击顶部的继续按钮，从一个断点跳转到代码中的下一个断点。](https://user-gold-cdn.xitu.io/2019/5/5/16a8711f6870c5c5?imageView2/0/w/1280/h/960/ignore-error/1)

当你从一个断点切换到另一个断点时，你可以看到程序在 VS Code 底部的调试控制台中打印出一堆 ` console.log` ，黄色的高亮显示也会随之一起移动。

![如你所见，当我们暂停在断点上时，我们可以在 VS Code 的左上角看到可以在控制台中探索到的所有局部作用域信息。](https://user-gold-cdn.xitu.io/2019/5/5/16a87120b0a64df6?imageView2/0/w/1280/h/960/ignore-error/1)

正如你所看到的，随着程序的运行，调试控制台输出的内容越多，断点就越多，在此过程中，我可以使用 VS Code 左上角的工具在本地范围内探索对象和函数，就像我可以在浏览器中探索范围和对象一样。不错！

这很简单，对吧？

### 总结 ###

Node.js 的调试不需要像过去那样麻烦，也不需要在代码库中包含 500 多个 ` console.log` 来找出 bug 的位置。

Visual Studio Code 的 ` Debug > Node: Auto Attach` 设置使之成为过去，我对此非常感激。

再过几周，我将会写一些关于端到端测试的文章，使用 Puppeteer 和 headless Chrome，或者使用 Nodemailer 在 MERN 应用程序中重置密码，所以请关注我，以免错过。

感谢阅读，希望这篇文章能让你了解如何在 VS Code 的帮助下更容易、更有效地调试 Node.js 程序。非常感谢你给我的掌声和对我文章的分享！

**如果你喜欢阅读这篇文章，你可能也会喜欢我的其他文章：**

* [使用 Node.js 读取超大数据集和文件（第一部分）]( https://link.juejin.im?target=https%3A%2F%2Fitnext.io%2Fusing-node-js-to-read-really-really-large-files-pt-1-d2057fe76b33 )
* [Sequelize: Node.js SQL ORM 框架]( https://link.juejin.im?target=https%3A%2F%2Fmedium.com%2F%40paigen11%2Fsequelize-the-orm-for-sql-databases-with-nodejs-daa7c6d5aca3 )
* [流的胜利：用于读取大型数据集的 Node.js 方法的性能比较（第二部分）]( https://link.juejin.im?target=https%3A%2F%2Fitnext.io%2Fstreams-for-the-win-a-performance-comparison-of-nodejs-methods-for-reading-large-datasets-pt-2-bcfa732fa40e )

**参考资料和进阶资源：**

* Github, Node 读取文件 Repo： [github.com/paigen11/fi…]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fpaigen11%2Ffile-read-challenge )
* Node.js 文档 — 调试部分： [nodejs.org/en/docs/gui…]( https://link.juejin.im?target=https%3A%2F%2Fnodejs.org%2Fen%2Fdocs%2Fguides%2Fdebugging-getting-started%2F )
* Paul Irish’s：使用 Chrome DevTools 调试 Node.js： [medium.com/@paul_irish…]( https://link.juejin.im?target=https%3A%2F%2Fmedium.com%2F%40paul_irish%2Fdebugging-node-js-nightlies-with-chrome-devtools-7c4a1b95ae27 )
* JetBrains 提供的文档 — 《运行和调试 Node.js》 — [www.jetbrains.com/help/websto…]( https://link.juejin.im?target=https%3A%2F%2Fwww.jetbrains.com%2Fhelp%2Fwebstorm%2Frunning-and-debugging-node-js.html )
* Visual Studio Code 下载链接： [code.visualstudio.com/download]( https://link.juejin.im?target=https%3A%2F%2Fcode.visualstudio.com%2Fdownload )
* VS Code 调试 Node.js 文档： [code.visualstudio.com/docs/nodejs…]( https://link.juejin.im?target=https%3A%2F%2Fcode.visualstudio.com%2Fdocs%2Fnodejs%2Fnodejs-debugging%23_attaching-to-nodejs )

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