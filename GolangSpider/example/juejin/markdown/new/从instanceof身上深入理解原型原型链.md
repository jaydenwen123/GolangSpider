# 从instanceof身上深入理解原型/原型链 #

今天将延续 [这篇文章]( https://link.juejin.im?target=https%3A%2F%2Fluozongmin.com%2F2019%2F06%2F01%2F%25E7%2594%25B1%25E4%25B8%2580%25E6%25AE%25B5%25E4%25BB%25A3%25E7%25A0%2581%25E5%25BC%2595%25E5%258F%2591%25E7%259A%2584%25E5%2585%25B3%25E4%25BA%258EObject%25E5%2592%258CFunction%25E7%259A%2584%25E9%25B8%25A1%25E5%2592%258C%25E8%259B%258B%25E9%2597%25AE%25E9%25A2%2598%25E7%259A%2584%25E6%2580%259D%25E8%2580%2583%2F ) ，借助一个老朋友——instanceof运算符，将通过它以及结合多次讲的原型/原型链经典图来深入理解原型/原型链。

对于原始类型（primitive type）的值，即 ` string` / ` number` / ` boolean` ，你可以通过 ` typeof` 判断其类型，但是 ` typeof` 在判断到合成类型（complex type）的值的时候，返回值只有 ` object` / ` function` ，你不知道它到底是一个 ` object` 对象，还是数组，也不能判断出Object 下具体是什么细分的类型，比如 ` Array` 、 ` Date` 、 ` RegExp` 、 ` Error` 等。

官方对 ` instanceof` 运算符的解释是返回一个布尔值，表示对象是否为某个构造函数的实例。比如：

` function Foo ( ) {} var f1 = new Foo(); console.log(f1 instanceof Foo); // true console.log(f1 instanceof Object ); // true 复制代码`

上面代码中，对象 ` f1` 是构造函数 ` Foo` 的实例，所以返回 ` true` ，但是“f1 instanceof Object”为什么也是 ` true` 呢？

至于为什么等会再解释，先把 ` instanceof` 判断的规则告诉大家。根据以上代码看下图：

![](https://user-gold-cdn.xitu.io/2019/6/5/16b26904a597b27a?imageView2/0/w/1280/h/960/ignore-error/1)

` instanceof` 运算符的左边是实例对象，右边是构造函数，左边变量暂称为A，右边变量暂称为B。它会检查右边构造函数的原型对象（prototype），是否在左边对象的原型链上。

通俗一点来讲， ` instanceof` 的判断规则是： **` instanceof` 会检查整个原型链，将沿着A的 ` __proto__` 这条线来一直找，同时沿着B的 ` prototype` 这条线来一直找，直到能找到同一个引用，即同一个对象，那么就返回 ` true` 。如果找到终点还未重合，则返回 ` false`** 。即上图中的 ` f1` --> ` __proto__` 和 ` Foo` --> ` prototype` 指向同一个对象， ` console.log(f1 instanceof Foo)` 为 ` true` 。

按照以上规则，重新来看看“ f1 instanceof Object ”这句代码为什么是 ` true` ？ 根据上图很容易就能看出来， f1--> ` __proto__` --> ` __proto__` 和 ` Object` --> ` prototype` 指向同一个对象， ` console.log(f1 instanceof Object)` 为 ` true` 。

通过上面的规则，可以很好地解释一些比较怪异的现象，例如：

` console.log( Object instanceof Function ); // true console.log( Function instanceof Object ); // true console.log( Function instanceof Function ); // true console.log( Object instanceof Object ); // true 复制代码`

这些就是 [这篇文章]( https://link.juejin.im?target=https%3A%2F%2Fluozongmin.com%2F2019%2F06%2F01%2F%25E7%2594%25B1%25E4%25B8%2580%25E6%25AE%25B5%25E4%25BB%25A3%25E7%25A0%2581%25E5%25BC%2595%25E5%258F%2591%25E7%259A%2584%25E5%2585%25B3%25E4%25BA%258EObject%25E5%2592%258CFunction%25E7%259A%2584%25E9%25B8%25A1%25E5%2592%258C%25E8%259B%258B%25E9%2597%25AE%25E9%25A2%2598%25E7%259A%2584%25E6%2580%259D%25E8%2580%2583%2F ) 所讲的看似很混乱的东西，现在知道为何了吧。

但还有一种特殊情况，就是左边对象的原型链上，只有 ` null` 对象。这时， ` instanceof` 判断会失真。

` var obj = Object.create( null ); typeof obj // "object" Object.create( null ) instanceof Object // false 复制代码`

上面代码中， ` Object.create(null)` 返回一个新对象 ` obj` ，它的原型是 ` null` （ ` Object.create` 后续会有专门文章介绍）。右边的构造函数 ` Object` 的 ` prototype` 属性，不在左边的原型链上，因此 ` instanceof` 就认为 ` obj` 不是 ` Object` 的实例。但是，只要一个对象的原型不是 ` null` ， ` instanceof` 运算符的判断就不会失真。

说到这里，继续贴上这幅原型/原型链的经典图，是否现在看起来没那么复杂了呢。

![](https://user-gold-cdn.xitu.io/2019/6/5/16b2694e5e430d6d?imageView2/0/w/1280/h/960/ignore-error/1)

如果 [这篇文章]( https://link.juejin.im?target=https%3A%2F%2Fluozongmin.com%2F2019%2F06%2F01%2F%25E7%2594%25B1%25E4%25B8%2580%25E6%25AE%25B5%25E4%25BB%25A3%25E7%25A0%2581%25E5%25BC%2595%25E5%258F%2591%25E7%259A%2584%25E5%2585%25B3%25E4%25BA%258EObject%25E5%2592%258CFunction%25E7%259A%2584%25E9%25B8%25A1%25E5%2592%258C%25E8%259B%258B%25E9%2597%25AE%25E9%25A2%2598%25E7%259A%2584%25E6%2580%259D%25E8%2580%2583%2F ) 你看的比较仔细，再结合刚才介绍的 ` instanceof` 的概念规则，相信能看懂上面这张图的内容了。

那么问题又出来了。 ` instanceof` 这样设计，到底有什么用？到底 ` instanceof` 想表达什么呢？

这就要重点讲讲继承了，即 ` instanceof` 表示的就是一种继承关系，或者原型链的结构，请看后续文章介绍。

**如果觉得文章对你有些许帮助，欢迎在 [我的GitHub博客]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fmiqilin21%2Fmiqilin21.github.io ) 点赞和关注，感激不尽！**