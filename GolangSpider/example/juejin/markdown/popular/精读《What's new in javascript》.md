# 精读《What's new in javascript》 #

# 1. 引言 #

本周精读的内容是： [Google I/O 19]( https://link.juejin.im?target=https%3A%2F%2Fwww.youtube.com%2Fwatch%3Fv%3Dc0oy0vQKEZE ) 。

2019 年 Google I/O 介绍了一些激动人心的 JS 新特性，这些特性有些已经被主流浏览器实现，并支持 polyfill，有些还在草案阶段。

我们可以看到 JS 语言正变得越来越严谨，不同规范间也逐渐完成了闭环，而且在不断吸纳其他语言的优秀特性，比如 WeakRef，让 JS 在成为使用范围最广编程语言的同时，也越成为编程语言的集大成者，让我们有信心继续跟随 JS 生态，不用被新生的小语种分散精力。

# 2. 精读 #

本视频共介绍了 16 个新特性。

## private class fields ##

私有成员修饰符，用于 Class：

` class IncreasingCounter { #count = 0; get value() { return this.#count; } increment() { this.#count++; } } 复制代码`

通过 ` #` 修饰的成员变量或成员函数就成为了私有变量，如果试图在 Class 外部访问，则会抛出异常：

` const counter = new IncreasingCounter() counter.#count // -> SyntaxError counter.#count = 42 // -> SyntaxError 复制代码`

虽然 ` #` 这个关键字被吐槽了很多次，但结论已经尘埃落定了，只是个语法形式而已，不用太纠结。

目前仅 Chrome、Nodejs 支持。

## Regex matchAll ##

正则匹配支持了 ` matchAll` API，可以更方便进行正则递归了：

` const string = 'Magic hex number: DEADBEEF CAFE' const regex = /\b\p{ASCII_Hex_Digit}+\b/gu/ for (const match of string.matchAll(regex)) { console.log(match) } // Output: // ['DEADBEEF', index: 19, input: 'Magic hex number: DEADBEEF CAFE'] // ['CAFE', index: 28, input: 'Magic hex number: DEADBEEF CAFE'] 复制代码`

相比以前在 ` while` 语句里循环正则匹配，这个 API 真的是相当的便利。And more，还顺带提到了 ` Named Capture Groups` ，这个在之前的 [精读《正则 ES2018》]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fdt-fe%2Fweekly%2Fblob%2Fv2%2F091.%25E7%25B2%25BE%25E8%25AF%25BB%25E3%2580%258A%25E6%25AD%25A3%25E5%2588%2599%2520ES2018%25E3%2580%258B.md%2322-named-capture-groups ) 中也有提到，具体可以点过去阅读，也可以配合 ` matchAll` 一起使用。

## Numeric literals ##

大数字面量的支持，比如：

` 1234567890123456789 * 123 ; // -> 151851850485185200000 复制代码`

这样计算结果是丢失精度的，但只要在数字末尾加上 ` n` ，就可以正确计算大数了：

` 1234567890123456789 n * 123 n; // -> 151851850485185185047n 复制代码`

目前 BigInt 已经被 Chrome、Firefox、Nodejs 支持。

## BigInt formatting ##

为了方便阅读，大数还支持了国际化，可以适配成不同国家的语言表达形式：

` const nf = new Intl.NumberFormat( "fr" ); nf.format( 12345678901234567890 n); // -> '12 345 678 901 234 567 890' 复制代码`

记住 ` Intl` 这个内置变量，后面还有不少国际化用途。

同时，为了方便程序员阅读代码，大数还支持带下划线的书写方式：

` const nf = new Intl.NumberFormat( "fr" ); nf.format( 12345678901234567890 n); // -> '12 345 678 901 234 567 890' 复制代码`

目前已经被 Chrome、Firefox、Nodejs 支持。

## flat & flatmap ##

等价于 lodash [flatten]( https://link.juejin.im?target=https%3A%2F%2Flodash.com%2Fdocs%2F4.17.11%23flatten ) 功能：

` const array = [ 1 , [ 2 , [ 3 ]]]; array.flat(); // -> [1, 2, [3]] 复制代码`

还支持自定义深度，如果支持 ` Infinity` 无限层级：

` const array = [ 1 , [ 2 , [ 3 ]]]; array.flat( Infinity ); // -> [1, 2, 3] 复制代码`

这样我们就可以配合 `.map` 使用：

` [ 2 , 3 , 4 ].map(duplicate).flat(); 复制代码`

因为这个用法太常见，js 内置了 ` flatMap` 函数代替 ` map` ，与上面的效果是等价的：

` [ 2 , 3 , 4 ].flatMap(duplicate); 复制代码`

目前已经被 Chrome、Firefox、Safari、Nodejs 支持。

## fromEntries ##

` fromEntries` 是 ` Object.fromEntries` 的语法，用来将对象转化为数组的描述：

` const object = { x : 42 , y : 50 , abc : 9001 }; const entries = Object.entries(object); // -> [['x', 42], ['y', 50]] 复制代码`

这样就可以对对象的 key 与 value 进行加工处理，并通过 ` fromEntries` API 重新转回对象：

` const object = { x : 42 , y : 50 , abc : 9001 } const result Object.fromEntries( Object.entries(object) .filter( ( [ key, value] ) => key.length === 1 ) .map( ( [ key, value ] ) => [ key, value * 2 ]) ) // -> { x: 84, y: 100 } 复制代码`

不仅如此，还可以将 object 快速转化为 Map:

` const map = new Map ( Object.entries(object)); 复制代码`

目前已经被 Chrome、Firefox、Safari、Nodejs 支持。

## Map to Object conversion ##

` fromEntries` 建立了 object 与 map 之间的桥梁，我们还可以将 Map 快速转化为 object：

` const objectCopy = Object.fromEntries(map); 复制代码`

目前已经被 Chrome、Firefox、Safari、Nodejs 支持。

## globalThis ##

> 
> 
> 
> 业务代码一般不需要访问全局的 window 变量，但是框架与库一般需要，比如 polyfill。
> 
> 

访问全局的 this 一般会做四个兼容，因为 js 在不同运行环境下，全局 this 的变量名都不一样：

` const getGlobalThis = () => { if ( typeof self !== "undefined" ) return self; // web worker 环境 if ( typeof window !== "undefined" ) return window ; // web 环境 if ( typeof global !== "undefined" ) return global; // node 环境 if ( typeof this !== "undefined" ) return this ; // 独立 js shells 脚本环境 throw new Error ( "Unable to locate global object" ); }; 复制代码`

因此整治一下规范也合情合理：

` globalThis; // 在任何环境，它就是全局的 this 复制代码`

目前已经被 Chrome、Firefox、Safari、Nodejs 支持。

## Stable sort ##

就是稳定排序结果的功能，比如下面的数组：

` const doggos = [ { name : "Abby" , rating : 12 }, { name : "Bandit" , rating : 13 }, { name : "Choco" , rating : 14 }, { name : "Daisy" , rating : 12 }, { name : "Elmo" , rating : 12 }, { name : "Falco" , rating : 13 }, { name : "Ghost" , rating : 14 } ]; doggos.sort( ( a, b ) => b.rating - a.rating); 复制代码`

最终排序结果可能如下：

` [ { name : "Choco" , rating : 14 }, { name : "Ghost" , rating : 14 }, { name : "Bandit" , rating : 13 }, { name : "Falco" , rating : 13 }, { name : "Abby" , rating : 12 }, { name : "Daisy" , rating : 12 }, { name : "Elmo" , rating : 12 } ]; 复制代码`

也可能如下：

` [ { name : "Ghost" , rating : 14 }, { name : "Choco" , rating : 14 }, { name : "Bandit" , rating : 13 }, { name : "Falco" , rating : 13 }, { name : "Abby" , rating : 12 }, { name : "Daisy" , rating : 12 }, { name : "Elmo" , rating : 12 } ]; 复制代码`

注意 ` choco` 与 ` Ghost` 的位置可能会颠倒，这是因为 JS 引擎可能只关注 ` sort` 函数的排序，而在顺序相同时，不会保持原有的排序规则。现在通过 **Stable sort** 规范，可以确保这个排序结果是稳定的。

目前已经被 Chrome、Firefox、Safari、Nodejs 支持。

## Intl.RelativeTimeFormat ##

` Intl.RelativeTimeFormat` 可以对时间进行语义化翻译：

` const rtf = new Intl.RelativeTimeFormat( "en" , { numeric : "auto" }); rtf.format( -1 , "day" ); // -> 'yesterday' rtf.format( 0 , "day" ); // -> 'today' rtf.format( 1 , "day" ); // -> 'tomorrow' rtf.format( -1 , "week" ); // -> 'last week' rtf.format( 0 , "week" ); // -> 'this week' rtf.format( 1 , "week" ); // -> 'next week' 复制代码`

不同语言体系下， ` format` 会返回不同的结果，通过控制 ` RelativeTimeFormat` 的第一个参数 ` en` 决定，比如可以切换为 ` ta-in` 。

## Intl.ListFormat ##

` ListFormat` 以列表的形式格式化数组：

` const lfEnglish = new Intl.ListFormat( "en" ); lfEnglish.format([ "Ada" , "Grace" ]); // -> 'Ada and Grace' 复制代码`

可以通过第二个参数指定连接类型：

` const lfEnglish = new Intl.ListFormat( "en" , { type : "disjunction" }); lfEnglish.format([ "Ada" , "Grace" ]); // -> 'Ada or Grace' 复制代码`

目前已经被 Chrome、Nodejs 支持。

## Intl.DateTimeFormat -> formatRange ##

` DateTimeFormat` 可以定制日期格式化输出：

` const start = new Date (startTimestamp); // -> 'May 7, 2019' const end = new Date (endTimestamp); // -> 'May 9, 2019' const fmt = new Intl.DateTimeFormat( "en" , { year : "numeric" , month : "long" , day : "numeric" }); const output = ` ${fmt.format(start)} - ${fmt.format(end)} ` ; // -> 'May 7, 2019 - May 9, 2019' 复制代码`

最后一句，也可以通过 ` formatRange` 函数代替：

` const output = fmt.formatRange(start, end); // -> 'May 7 - 9, 2019' 复制代码`

目前已经被 Chrome 支持。

## Intl.Locale ##

定义国际化本地化的相关信息：

` const locale = new Intl.Locale( "es-419-u-hc-h12" , { calendar : "gregory" }); locale.language; // -> 'es' locale.calendar; // -> 'gregory' locale.hourCycle; // -> 'h12' locale.region; // -> '419' locale.toString(); // -> 'es-419-u-ca-gregory-hc-h12' 复制代码`

目前已经被 Chrome、Nodejs 支持。

## Top-Level await ##

支持在根节点生效 ` await` ，比如：

` const result = await doSomethingAsync(); doSomethingElse(); 复制代码`

目前还没有支持。

## Promise.allSettled/Promise.any ##

` Promise.allSettled` 类似 ` Promise.all` 、 ` Promise.any` 类似 ` Promise.race` ，区别是，在 Promise reject 时， ` allSettled` 不会 reject，而是也当作 fulfilled 的信号。

举例来说：

` const promises = [ fetch( "/api-call-1" ), fetch( "/api-call-2" ), fetch( "/api-call-3" ) ]; await Promise.allSettled(promises); 复制代码`

即便某个 ` fetch` 失败了，也不会导致 ` reject` 的发生，这样在不在乎是否有项目失败，只要拿到都结束的信号的场景很有用。

对于 ` Promise.any` 则稍有不同：

` const promises = [ fetch( "/api-call-1" ), fetch( "/api-call-2" ), fetch( "/api-call-3" ) ]; try { const first = await Promise.any(promises); // Any of ths promises was fulfilled. console.log(first); } catch (error) { // All of the promises were rejected. } 复制代码`

只要有子项 fulfilled，就会完成 ` Promise.any` ，哪怕第一个 Promise reject 了，而第二个 Promise fulfilled 了， ` Promise.any` 也会 fulfilled，而对于 ` Promise.race` ，这种场景会直接 rejected。

如果所有子项都 rejected，那 ` Promise.any` 也只好 rejected 啦。

目前已经被 Chrome、Firefox 支持。

## WeakRef ##

WeakRef 是从 OC 抄过来的弱引用概念。

为了解决这个问题：当对象被引用后，由于引用的存在，导致对象无法被 GC。

所以如果建立了弱引用，那么对象就不会因为存在的这段引用关系而影响 GC 了！

具体用法是：

` const obj = {}; const weakObj = new WeakRef(obj); 复制代码`

使用 ` weakObj` 与 ` obj` 没有任何区别，唯一不同时， ` obj` 可能随时被 GC，而一旦被 GC，弱引用拿到的对象可能就变成 ` undefined` ，所以要做好错误保护。

# 3. 总结 #

JS 这几个特性提升了 JS 语言的成熟性、完整性，而且看到其访问控制能力、规范性、国际化等能力有着重加强，解决的都是 JS 最普遍遇到的痛点问题。

那么，这些 JS 特性中，你最喜欢哪一条呢？想吐槽哪一条呢？欢迎留言。

> 
> 
> 
> 讨论地址是： [精读《What's new in javascript》 · Issue #159 · dt-fe/weekly](
> https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fdt-fe%2Fweekly%2Fissues%2F159
> )
> 
> 

**如果你想参与讨论，请 [点击这里]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fdt-fe%2Fweekly ) ，每周都有新的主题，周末或周一发布。前端精读 - 帮你筛选靠谱的内容。**

> 
> 
> 
> 关注 **前端精读微信公众号**
> 
> 

![](https://user-gold-cdn.xitu.io/2019/5/27/16af7118a5947f57?imageView2/0/w/1280/h/960/ignore-error/1)

**special Sponsors**

* [DevOps 全流程平台]( https://link.juejin.im?target=https%3A%2F%2Fe.coding.net%2F%3Futm_source%3Dweekly )

> 
> 
> 
> 版权声明：自由转载-非商用-非衍生-保持署名（ [创意共享 3.0 许可证](
> https://link.juejin.im?target=https%3A%2F%2Fcreativecommons.org%2Flicenses%2Fby-nc-nd%2F3.0%2Fdeed.zh
> ) ）
> 
>