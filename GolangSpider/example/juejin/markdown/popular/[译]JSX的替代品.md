# [译]JSX的替代品 #

## 前言 ##

JSX现在是一种非常受欢迎的选择，用户在各种框架中进行模板模式开发，而不仅仅是在React中，但是，如果你不喜欢它，或者有一个你想要避免使用它的项目，或者只是好奇如何在没有它的情况下编写您的React代码呢？最简单的答案是阅读 [官方文档]( https://link.juejin.im?target=https%3A%2F%2Freactjs.org%2Fdocs%2Freact-without-jsx.html ) ，但是它有点短。

我们有更多的选择

` 免责声明：就个人而言，我喜欢JSX并在我的React项目中使用它，但是，我稍微研究了这个主题，发现了一些成果，并想分享给大家。`

## 什么是JSX ##

首先，我们需要了解JSX是什么，以便用纯JavaScript代码来编写匹配的代码。

JSX是一种特定于域的语言，这意味着我们需要使用JSX转换代码以获得常规JavaScript，否则浏览器将无法理解我们的代码。在未来，如果你的目标浏览器不是完全支持所有的JSX语法转换，您就无法完全删除转换，这可能是一个问题。

理解JSX的最佳方法可能是使用babel-repl实际执行此操作。您需要单击presets（应该在左侧面板中）并且选择react，以便正确解析它。之后，您将能够在右侧实时查看JavaScript代码。例如，您可以写入以下内容：

![](https://user-gold-cdn.xitu.io/2019/5/17/16ac4f4d69ba0fc2?imageView2/0/w/1280/h/960/ignore-error/1)

其实这段原本写法为。

![](https://user-gold-cdn.xitu.io/2019/5/17/16ac4f59d7443789?imageView2/0/w/1280/h/960/ignore-error/1)

您可以看到每<%tag%>的结构，都被函数调用 [React.createElement]( https://link.juejin.im?target=https%3A%2F%2Freactjs.org%2Fdocs%2Freact-api.html%23createelement ) 替换。

第一个参数是徐建活具有内置标记值的字符串（如div或span），第二个参数是关于options 所有其余参数都被视为子项。

我强烈建议您使用不同的树来玩，例如，看看React如何使用true或false值、数组、组件等呈现属性：即使您只使用JSX和漂亮的内容，它也很有用。

` 要深入阅读JSX，有一个[官方文档](https://reactjs.org/docs/jsx-in-depth.html)页面`

## 重命名 ##

虽然生成的代码完全有效，并且我们可以用这种方式编写所有的React代码，但这种方法存在一些问题。

第一个问题是它非常冗长。就像真人一样很啰嗦，这里的主要罪犯是React.createElement。因此，第一个解决方案是把它保存到一个变量，通常命名为h，类似 [hyperscript]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fhyperhype%2Fhyperscript ) 。这将为您节省大量文本，并使其更具可读性。为了说明这一点，让我们重写一下过去的例子：

![](https://user-gold-cdn.xitu.io/2019/5/17/16ac4fcdb8848859?imageView2/0/w/1280/h/960/ignore-error/1)

## Hyperscript ##

如果您使用过上面任何一个例子用于开发的话，您会发现它有几个缺陷。首先，React.createElement函数需要 3个参数，所以如果没有属性，你必须提供null，并且className可能是最常见的属性，每次都需要编写一个新对象。

作为替代方案，您可以使用 [react-hyperscript]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fmlmorg%2Freact-hyperscript ) 库。它不需要提供空道具，也允许您以类似emmet的样式

` div#main.content- > <div id="main" class="content"> 复制代码`

指定类和ID 。这样子改版的话，我们的代码会变的更精炼：

![](https://user-gold-cdn.xitu.io/2019/5/17/16ac4ff988aa4434?imageView2/0/w/1280/h/960/ignore-error/1)

## HTM ##

如果您不反对JSX本身，但不喜欢转换代码的必要性，那么有一个名为htm的项目。它的宗旨与JSX一样（并且看起来很想通），但是它使用模板文字。这肯定会增加一些开销（你必须在运行时解析这些模板）。但是它在某些情况下可能是值得的。

它通过包装元素函数来工作，React.createElement在我们的例子中，它可以是任何其他具有类似API的库，并返回一个函数，仅在运行时。它将解析我们的模板并返回与babel完全相同的代码。

![](https://user-gold-cdn.xitu.io/2019/5/17/16ac5014662ae5c0?imageView2/0/w/1280/h/960/ignore-error/1) 正如您所看到的，它几乎与真正的JSX相同，我们只需要以稍微不同一点的方式插入变量。

但是，这些主要是细节，如果你想在没有任何工具设置的情况下展示如何使用React，这可能很方便。

## 类似Lisp的语法 ##

这个想法类似于hyperscript，然而，这是一个值得关注的优雅方法。还有许多其他类似的辅助库，因此，在选择哪个上，完全取决于主观。 它可能会为您自己的项目提供一些灵感。

库 [ijk]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Flukejacksonn%2Fijk ) 带来了仅使用数组编写模板的想法，使用位置作为参数。主要优点是你不需要经常写h（是的，甚至可以重复！）。以下是如何使用它的示例：

![](https://user-gold-cdn.xitu.io/2019/5/17/16ac503a98acf770?imageView2/0/w/1280/h/960/ignore-error/1)

## END ##

本文并未说明您不应该使用JSX，或者它是否是一个坏主意。但是，您可能想知道如何在没有它的情况下编写代码，以及代码的外观如何，本文的目的仅仅是回答这个问题。

来源： [blog.bloomca.me/2019/02/23/…]( https://link.juejin.im?target=https%3A%2F%2Fblog.bloomca.me%2F2019%2F02%2F23%2Falternatives-to-jsx.html )