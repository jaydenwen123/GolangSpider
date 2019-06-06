# 用 JS 代码解释 Java Stream #

## 引言 ##

前不久，公司后端同事找到我，邀请我在月会上分享函数式编程，我说你还是另请高明吧…… 我也不是谦虚，我一个前端页面仔，怎么去给以 Java 后端开发为主的技术部讲函数式编程呢？但是同事说你还是试试吧。然后我就去先试着准备下。

由于我最近在学函数式领域建模(Functional Domain Modeling)，一开始我想讲下 Scala，然后我找到了 Functional and Reactive Domain Modeling 这本书。但是转念一想，我都没写过后端，对后端的业务场景根本不了解，看本书就去讲，有些不尊重听众智商了。最后我打算在 Java 里找些接地气的内容来分享，然后我就去学了 Java，发现了 Java 里面有 Stream 这么好的函数式一级语言支持。

最终经过谨慎考虑后，我还是没有去分享。我没有业务上的理解，仅仅去讲些语法特性和奇技淫巧，有些班门弄斧了。我在学习 Stream 的过程中有一些发现，现在把这些发现分享出来。

本文代码在我上一篇文章中出现过，但这次解释更详细。

## JavaScript 更适合用来学习一些 FP 概念 ##

本文无意抛出 JavaScript 和 Java 谁比谁好这种无聊的论断，我仅想指出，JS 这种弱类型语言可以避免一些语法噪音，让初学者在学习一些 FP 概念时，快速直抵核心，理解本质。

在学 Java Stream 时，要先学 lambda expression, method reference, 然后学 Interface 和 Functional Interface, 学这些仅仅为了支持 lambda 传参。而用 JS 的话，函数顺手就写，顺手就传，不用那么多准备仪式。

再次声明一下，Java 这种面向对象特性有其强大之处。这里仅仅指出在学习理解一些 FP 概念上，JS 更简单灵活。

## 理解 Stream ##

Stream 可以拆解成三个部分：

![](https://user-gold-cdn.xitu.io/2019/4/4/169e7e003317c935?imageView2/0/w/1280/h/960/ignore-error/1)

如上图，Stream 可分为数据源，管道数据操作，数据消费三个部分。源数据生成后，流经管道，最终在管道终端被消费。最终的消费可能是数据流被聚合成新的数据集，也可能是逐个执行副作用( ` forEach` )。

下面用 JS 代码解释这三个部分。

## 数据源 ##

数据源可以理解成是一个 generator 函数被反复执行，生成数据流。什么是 generator 呢？最简单的 generator 可以是这样：

` const getRandNum = () => Math.floor( Math.random() * 100 ); 复制代码`

每次执行 ` getRandNum` 它都随机生成一个 0 到 100 的整数，它满足我们对 generator 行为的期待。可以看出 generator 其实就是一个动态生成数据的行为，那如果数据源是静态数据集，怎么得到这种动态 generator 呢？很简单：

` function getGeneratorFromList ( list ) { let index = 0 ; return function generate ( ) { if (index < list.length) { const result = list[index]; index += 1 ; return result; } }; } // 例子： const generate = getGeneratorFromList([ 1 , 2 , 3 ]); generate(); // 1 generate(); // 2 generate(); // 3 generate(); // undefined 复制代码`

给 ` getGeneratorFromList` 传个数组，它会返回一个 generator，这个 generator 每被执行一次都吐出传入数组当前遍历到的元素。这里只考虑数组，其它情况很容易扩展。

数据源部分就讲完了。其实很简单。

## 管道组合 ##

数据源生成后，就进入管道了。管道里对原 generator 进行各种转换，组合生成符合期待的 generator。注意上一句的描述，管道里仅仅是基于原 generator 函数生成新的 generator 函数， **计算行为并不会被触发** 。

我们给管道里塞入一些高阶操作函数，这些高阶操作函数接受前一个函数吐出的 generator，返回加入了新行为的 generator。我们先来定义一个 ` map` :

` function map ( mapping ) { return function ( generate ) { return function mappedGenerator ( ) { const value = generate(); if (value !== undefined ) { return mapping(value); } }; }; } 复制代码`

` map` 函数连续返回了两个函数。这里先简要解释下为什么这么做，可能会比较难懂，等你看了后面的代码再回过头来看会更容易理解些。当用户调用 ` map` 并将结果传入管道时， ` map` 返回了第二层函数。当管道组合被触发时，第二层函数被执行，最终 generator 函数被返回，然后返回的 generator 被传给管道里的下一个高阶操作函数。

再来定义一个 ` filter` ：

` function filter ( predicate ) { return function ( generate ) { return function filteredGenerator ( ) { const value = generate(); if (value !== undefined ) { if (predicate(value)) { return value; } return filteredGenerator(); } }; }; } 复制代码`

当判断条件不满足时，generator 会跳过当前值，持续递归，直到遍历到符合条件的值才将值返回出去。

` map` 和 ` filter` 是最重要的高阶操作函数，其它的操作函数就不展开解释了。

在提供了高阶操作函数之后，我们要提供一个函数将这些高阶函数组合起来。

` function Stream ( source ) { let initialGenerator; if ( Array.isArray(source)) { initGenerator = getGeneratorFromList(source); } else { initialGenerator = source; } function pipe (...operators ) { return operators.reduce( ( prev, current ) => current(prev), initialGenerator); } return { pipe }; } 复制代码`

本文就只考虑数据源是 generator 函数和数组两个情况了。这种考虑当然是不严谨的，但足以解释 Stream 的实现。

` pipe` 函数将操作函数从左往右依次执行，并将上一个函数执行的结果传给下一个函数。如此则完成了管道组合操作。

## 数据消费 ##

前面两步完成后，就剩下如何将动态的 generator 转换成静态的数据集了(forEach 执行副作用本文就不考虑了)。这一步也比较简单：

` function toList ( generate ) { const arr = []; let value = generate(); while (value !== undefined ) { arr.push(value); value = generate(); } return arr; } 复制代码`

这里就只展示生成数组了，其它数据类型读者可自行扩展。

至此，完整的 Stream 就实现了。当然，操作符支持上还不完整，但你能明白我的意思。

再定义一个 ` take` ，然后测试下：

` function take ( n ) { return function ( generate ) { let count = 0 ; return function ( ) { if (count < n) { count += 1 ; return generate(); } }; }; } Stream(getRandNum).pipe( filter( x => x % 2 === 1 ), take( 10 ), toList ); // => 10 个随机奇数 复制代码`

注意， ` getRandNum` 永远都不会返回 ` undefined` ，那为什么 ` toList` 没有进入死循环？这是因为 ` take` 给原 generator 加入了新的行为，让它只能返回 10 个有效值。这也是惰性求值的魅力。

更多高阶操作符如下：

` function skipWhile ( predicate ) { return function ( generate ) { let startTaking = false ; return function skippedGenerator ( ) { const value = generate(); if (value !== undefined ) { if (startTaking) { return value; } else if (!predicate(value)) { startTaking = true ; return value; } return skippedGenerator(); } }; }; } function takeWhile ( predicate ) { return function ( generate ) { return function ( ) { const value = generate(); if (value !== undefined ) { if (predicate(value)) { return value; } } }; }; } 复制代码`

## 后记 ##

Java 里面的 Stream 底层实现我并没有学习，我只学了下 API 用法。但是从 Stream API 的行为特性来推断，其底层实现应该和本文展示的 JS 实现思想是相通的。

前几天在 [知乎]( https://link.juejin.im?target=https%3A%2F%2Fzhuanlan.zhihu.com%2Fp%2F51549583 ) 上偶然发现有人未经我授权把我去年发表的《如何在 JS 里面消灭 for 循环》转载了。我看了下评论，比在掘金被骂的还惨。没想到在掘金被骂了一轮还要在知乎上再被骂一轮……

如果你写 Java，Java 8 给你提供了 Stream 你不用，偏偏要用 for 循环，还攻击用前者的人是在玩语法游戏，是不是很傻？