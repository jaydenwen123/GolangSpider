# 【译】11 种在大多数教程中找不到的 JavaScript 技巧 #

> 
> 
> 
> 译者：前端小智
> 
> 
> 
> 原文： [medium.com/@bretcamero…](
> https://link.juejin.im?target=https%3A%2F%2Fmedium.com%2F%40bretcameron%2F12-javascript-tricks-you-wont-find-in-most-tutorials-a9c9331f169d
> )
> 
> 

当我开始学习JavaScript时，我把我在别人的代码、code challenge网站以及我使用的教程之外的任何地方发现的每一个节省时间的技巧都列了一个清单。

在这篇文章中，我将分享11条我认为特别有用的技巧。这篇文章是为初学者准备的，但我希望即使是中级JavaScript开发人员也能在这个列表中找到一些新的东西。

**想阅读更多优质文章请猛戳 [GitHub博客]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fqq449245884%2Fxiaozhi ) ,一年百来篇优质文章等着你！**

## 1..过滤唯一值 ##

` Set` 对象类型是在ES6中引入的，配合展开操作 `...` 一起，我们可以使用它来创建一个新数组，该数组只有唯一的值。

` const array = [1, 1, 2, 3, 5, 5, 1] const uniqueArray = [...new Set(array)]; console.log(uniqueArray); // Result: [1, 2, 3, 5] 复制代码`

在ES6之前，隔离惟一值将涉及比这多得多的代码。

此技巧适用于包含基本类型的数组： ` undefined` ， ` null` ， ` boolean` ， ` string` 和 ` number` 。 （如果你有一个包含对象，函数或其他数组的数组，你需要一个不同的方法！）

##2. 与或运算

三元运算符是编写简单（有时不那么简单）条件语句的快速方法，如下所示：

` x > 100 ? 'Above 100' : 'Below 100'; x > 100 ? (x > 200 ? 'Above 200' : 'Between 100-200') : 'Below 100'; 复制代码`

但有时使用三元运算符处理也会很复杂。 相反，我们可以使用'与' ` &&` 和'或' ` ||` 逻辑运算符以更简洁的方式书写表达式。 这通常被称为“短路”或“短路运算”。

#### **它是怎么工作的** ####

假设我们只想返回两个或多个选项中的一个。

使用 ` &&` 将返回第一个条件为 ` 假` 的值。如果每个操作数的计算值都为 ` true` ，则返回最后一个计算过的表达式。

` let one = 1, two = 2, three = 3; console.log(one && two && three); // Result: 3 console.log(0 && null); // Result: 0 复制代码`

使用 ` ||` 将返回第一个条件为 ` 真` 的值。如果每个操作数的计算结果都为 ` false` ，则返回最后一个计算过的表达式。

` let one = 1, two = 2, three = 3; console.log(one || two || three); // Result: 1 console.log(0 || null); // Result: null 复制代码`

#### **例一** ####

假设我们想返回一个变量的长度，但是我们不知道变量的类型。

我们可以使用 ` if/else` 语句来检查 ` foo` 是可接受的类型，但是这可能会变得非常冗长。或运行可以帮助我们简化操作：

` return (foo || []).length 复制代码`

如果变量 ` foo` 是true，它将被返回。否则，将返回空数组的长度: ` 0` 。

#### **例二** ####

你是否遇到过访问嵌套对象属性的问题？ 你可能不知道对象或其中一个子属性是否存在，这可能会导致令人沮丧的错误。

假设我们想在 ` this.state` 中访问一个名为 ` data` 的属性，但是在我们的程序成功返回一个获取请求之前， ` data` 是未定义的。

根据我们使用它的位置，调用 ` this.state.data` 可能会阻止我们的应用程序运行。 为了解决这个问题，我们可以将其做进一步的判断：

` if (this.state.data) { return this.state.data; } else { return 'Fetching Data'; } 复制代码`

但这似乎很重复。 ' ` 或'` 运算符提供了更简洁的解决方案：

` return (this.state.data || 'Fetching Data'); 复制代码`

#### **一个新特性: Optional Chaining** ####

过去在 Object 属性链的调用中，很容易因为某个属性不存在而导致之后出现 ` Cannot read property xxx of undefined` 的错误。

那 ` optional chaining` 就是添加了 ` ?.` 这么个操作符，它会先判断前面的值，如果是 ` null` 或 ` undefined` ，就结束调用、返回 ` undefined` 。

例如，我们可以将上面的示例重构为 ` this.state.data?.()` 。或者，如果我们主要关注 ` state` 是否已定义，我们可以返回 ` this.state？.data` 。

该提案目前处于第1阶段，作为一项实验性功能。 你可以在 [这里]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Ftc39%2Fproposal-optional-chaining ) 阅读它，你现在可以通过Babel使用你的JavaScript，将 [@babel/plugin-proposal-optional-chaining]( https://link.juejin.im?target=https%3A%2F%2Fbabeljs.io%2Fdocs%2Fen%2Fbabel-plugin-proposal-optional-chaining ) 添加到你的 `.babelrc` 文件中。

## 3.转换为布尔值 ##

除了常规的布尔值 ` true` 和 ` false` 之外，JavaScript还将所有其他值视为 **‘truthy’** 或**‘falsy’**。

除非另有定义，否则 JavaScript 中的所有值都是'truthy'，除了 ` 0` ， ` “”` ， ` null` ， ` undefined` ， ` NaN` ，当然还有 ` false` ，这些都是**'falsy'**

我们可以通过使用负算运算符轻松地在 ` true` 和 ` false` 之间切换。它也会将类型转换为“boolean”。

` const isTrue = !0; const isFalse = !1; const alsoFalse = !!0; console.log(isTrue); // Result: true console.log(typeof true); // Result: "boolean" 复制代码`

## 4. 转换为字符串 ##

要快速地将数字转换为字符串，我们可以使用连接运算符 ` +` 后跟一组空引号 ` ""` 。

` const val = 1 + ""; console.log(val); // Result: "1" console.log(typeof val); // Result: "string" 复制代码`

## 5. 转换为数字 ##

使用加法运算符 ` +` 可以快速实现相反的效果。

` let int = "15"; int = +int; console.log(int); // Result: 15 console.log(typeof int); Result: "number" 复制代码`

这也可以用于将布尔值转换为数字，如下所示

` console.log(+true); // Return: 1 console.log(+false); // Return: 0 复制代码`

在某些上下文中， ` +` 将被解释为连接操作符，而不是加法操作符。当这种情况发生时(你希望返回一个整数，而不是浮点数)，您可以使用两个波浪号: ` ~~` 。

连续使用两个波浪有效地否定了操作，因为 ` — ( — n — 1) — 1 = n + 1 — 1 = n` 。 换句话说， ` ~—16` 等于 ` 15。`

` const int = ~~"15" console.log(int); // Result: 15 console.log(typeof int); Result: "number" 复制代码`

虽然我想不出很多用例，但是按位NOT运算符也可以用在布尔值上： ` ~true = -2` 和 ` ~false = -1` 。

## 6.性能更好的运算 ##

从ES7开始，可以使用指数运算符 ` **` 作为幂的简写，这比编写 ` Math.pow(2, 3)` 更快。 这是很简单的东西，但它之所以出现在列表中，是因为没有多少教程更新过这个操作符。

` console.log(2 ** 3); // Result: 8 复制代码`

这不应该与通常用于表示指数的^符号相混淆，但在JavaScript中它是按位 ` 异或` 运算符。

在ES7之前，只有以 ` 2` 为基数的幂才存在简写，使用按位左移操作符 ` <<`

` Math.pow(2, n); 2 << (n - 1); 2**n; 复制代码`

例如， ` 2 << 3 = 16` 等于 ` 2 ** 4 = 16` 。

## 7. 快速浮点数转整数 ##

如果希望将浮点数转换为整数，可以使用 ` Math.floor()` 、 ` Math.ceil()` 或 ` Math.round()` 。但是还有一种更快的方法可以使用 ` |` (位或运算符)将浮点数截断为整数。

` console.log(23.9 | 0); // Result: 23 console.log(-23.9 | 0); // Result: -23 复制代码`

` |` 的行为取决于处理的是正数还是负数，所以最好只在确定的情况下使用这个快捷方式。

如果 ` n` 为正，则 ` n | 0` 有效地向下舍入。 如果 ` n` 为负数，则有效地向上舍入。 更准确地说，此操作将删除小数点后面的任何内容，将浮点数截断为整数。

你可以使用 ` ~~` 来获得相同的舍入效果，如上所述，实际上任何位操作符都会强制浮点数为整数。这些特殊操作之所以有效，是因为一旦强制为整数，值就保持不变。

#### **删除最后一个数字** ####

` 按位或` 运算符还可以用于从整数的末尾删除任意数量的数字。这意味着我们不需要使用这样的代码来在类型之间进行转换。

` let str = "1553"; Number(str.substring(0, str.length - 1)); 复制代码`

相反，按位或运算符可以这样写：

` console.log(1553 / 10 | 0) // Result: 155 console.log(1553 / 100 | 0) // Result: 15 console.log(1553 / 1000 | 0) // Result: 1 复制代码`

## 8. 类中的自动绑定 ##

我们可以在类方法中使用ES6箭头表示法，并且通过这样做可以隐含绑定。 这通常会在我们的类构造函数中保存几行代码，我们可以愉快地告别重复的表达式，例如 ` this.myMethod = this.myMethod.bind(this)`

` import React, { Component } from React; export default class App extends Compononent { constructor(props) { super(props); this.state = {}; } myMethod = () => { // This method is bound implicitly! } render() { return ( <> <div> {this.myMethod()} </div> </> ) } }; 复制代码`

## 9. 数组截断 ##

如果要从数组的末尾删除值，有比使用 ` splice()` 更快的方法。

例如，如果你知道原始数组的大小，您可以重新定义它的 ` length` 属性，就像这样

` let array = [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]; array.length = 4; console.log(array); // Result: [0, 1, 2, 3] 复制代码`

这是一个特别简洁的解决方案。但是，我发现 ` slice()` 方法的运行时更快。如果速度是你的主要目标，考虑使用：

` let array = [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]; array = array.slice(0, 4); console.log(array); // Result: [0, 1, 2, 3] 复制代码`

## 10. 获取数组中的最后一项 ##

数组方法 ` slice()` 可以接受负整数，如果提供它，它将接受数组末尾的值，而不是数组开头的值。

` let array = [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]; console.log(array.slice(-1)); // Result: [9] console.log(array.slice(-2)); // Result: [8, 9] console.log(array.slice(-3)); // Result: [7, 8, 9] 复制代码`

## 11.格式化JSON代码 ##

最后，你之前可能已经使用过 ` JSON.stringify` ，但是您是否意识到它还可以帮助你缩进JSON？

` stringify()` 方法有两个可选参数：一个 ` replacer` 函数，可用于过滤显示的JSON和一个空格值。

` console.log(JSON.stringify({ alpha: 'A', beta: 'B' }, null, '\t')); // Result: // '{ // "alpha": A, // "beta": B // }' 复制代码`

**你的点赞是我持续分享好东西的动力，欢迎点赞！**

## 交流 ##

干货系列文章汇总如下，觉得不错点个Star，欢迎 加群 互相学习。

> 
> 
> 
> [github.com/qq449245884…](
> https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fqq449245884%2Fxiaozhi
> )
> 
> 

我是小智，公众号「大迁世界」作者， **对前端技术保持学习爱好者。我会经常分享自己所学所看的干货** ，在进阶的路上，共勉！

关注公众号，后台回复 **福利** ，即可看到福利，你懂的。

![](https://user-gold-cdn.xitu.io/2019/5/28/16afc0443b42d848?imageView2/0/w/1280/h/960/ignore-error/1)