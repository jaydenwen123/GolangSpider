# 【Step-By-Step】一周面试题深入解析 / 周刊01 #

> 
> 
> 
> 一周汇总
> 
> 

1.如何正确判断this的指向？

2.JS中原始类型有哪几种？null 是对象吗？原始数据类型和复杂数据类型有什么区别？

3.说一说你对HTML5语义化的理解

4.如何让 (a == 1 && a == 2 && a == 3) 的值为true？

5.防抖(debounce)函数的作用是什么？有哪些应用场景，请实现一个防抖函数。

### 关于【Step-By-Step】 ###

**不积跬步无以至千里。**

[Step-By-Step]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2FYvetteLau%2FStep-By-Step ) (点击进入项目) 是我于 ` 2019-05-20` 开始的一个项目，项目愿景：一步一个脚印，量变引起质变。

[Step-By-Step]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2FYvetteLau%2FStep-By-Step ) 仅会在工作日发布面试题，主要考虑到部分小伙伴平时工作较为繁忙，或周末有出游计划。每个周末我会仔细阅读大家的答案，整理最一份较优答案出来，因本人水平有限，有误的地方，大家及时指正。参与答题的小伙伴，可以对比自己的回答。

**答题不是目的** ，不希望大家仅仅是简单的搜索答案，复制粘贴到issue下。更多的是希望大家及时查漏补缺 / 巩固相关知识。

已经过去的一周中，大家回答的都非常认真，希望小伙伴们能够一如既往的坚持。同样也欢迎更多的小伙伴一起参与进来，如果想 **加群** 学习，扫码 **二维码** ( https://link.juejin.im?target=https%3A%2F%2Fuser-gold-cdn.xitu.io%2F2019%2F5%2F26%2F16af37df7708267b%3Fw%3D243%26amp%3Bh%3D245%26amp%3Bf%3Djpeg%26amp%3Bs%3D36789 ) (点击查看)，添加我为好友，验证信息为 **加入组织** ，我拉你进群。

### 1.如何正确判断this的指向？(2019-05-20) ###

如果用一句话说明 this 的指向，那么即是: 谁调用它，this 就指向谁。

但是仅通过这句话，我们很多时候并不能准确判断 this 的指向。因此我们需要借助一些规则去帮助自己：

this 的指向可以按照以下顺序判断:

#### 全局环境中的 this ####

浏览器环境：无论是否在严格模式下，在全局执行环境中（在任何函数体外部）this 都指向全局对象 ` window` ;

node 环境：无论是否在严格模式下，在全局执行环境中（在任何函数体外部），this 都是空对象 ` {}` ;

#### 是否是 ` new` 绑定 ####

如果是 ` new` 绑定，并且构造函数中没有返回 function 或者是 object，那么 this 指向这个新对象。如下:

> 
> 
> 
> 构造函数返回值不是 function 或 object。
> 
> 

` function Super ( age ) { this.age = age; } let instance = new Super( '26' ); console.log(instance.age); //26 复制代码`
> 
> 
> 
> 
> 构造函数返回值是 function 或 object，这种情况下 this 指向的是返回的对象。
> 
> 

` function Super ( age ) { this.age = age; let obj = { a : '2' }; return obj; } let instance = new Super( 'hello' ); console.log(instance.age); //undefined 复制代码`

你可以想知道为什么会这样？我们来看一下 ` new` 的实现原理:

* 创建一个新对象。
* 这个新对象会被执行 ` [[原型]]` 连接。
* 属性和方法被加入到 this 引用的对象中。并执行了构造函数中的方法.
* 如果函数没有返回其他对象，那么 this 指向这个新对象，否则 this 指向构造函数中返回的对象。
` function new ( func ) { let target = {}; target.__proto__ = func.prototype; let res = func.call(target); //排除 null 的情况 if (res && typeof (res) == "object" || typeof (res) == "function" ) { return res; } return target; } 复制代码`

#### 函数是否通过 call,apply 调用，或者使用了 bind 绑定，如果是，那么this绑定的就是指定的对象【归结为显式绑定】。 ####

` function info ( ) { console.log( this.age); } var person = { age : 20 , info } var age = 28 ; var info = person.info; info.call(person); //20 info.apply(person); //20 info.bind(person)(); //20 复制代码`

这里同样需要注意一种 **特殊** 情况，如果 call,apply 或者 bind 传入的第一个参数值是 ` undefined` 或者 ` null` ，严格模式下 this 的值为传入的值 null /undefined。非严格模式下，实际应用的默认绑定规则，this 指向全局对象(node环境为global，浏览器环境为window)

` function info ( ) { //node环境中:非严格模式 global，严格模式为null //浏览器环境中:非严格模式 window，严格模式为null console.log( this ); console.log( this.age); } var person = { age : 20 , info } var age = 28 ; var info = person.info; //严格模式抛出错误； //非严格模式，node下输出undefined（因为全局的age不会挂在 global 上） //非严格模式。浏览器环境下输出 28（因为全局的age会挂在 window 上） info.call( null ); 复制代码`

#### 隐式绑定，函数的调用是在某个对象上触发的，即调用位置上存在上下文对象。典型的隐式调用为: ` xxx.fn()` ####

` function info ( ) { console.log( this.age); } var person = { age : 20 , info } var age = 28 ; person.info(); //20;执行的是隐式绑定 复制代码`

#### 默认绑定，在不能应用其它绑定规则时使用的默认规则，通常是独立函数调用。 ####

非严格模式： node环境，执行全局对象 global，浏览器环境，执行全局对象 window。

严格模式：执行 undefined

` function info ( ) { console.log( this.age); } var age = 28 ; //严格模式；抛错 //非严格模式，node下输出 undefined（因为全局的age不会挂在 global 上） //非严格模式。浏览器环境下输出 28（因为全局的age不会挂在 window 上） //严格模式抛出，因为 this 此时是 undefined info(); 复制代码`

#### 箭头函数的情况： ####

箭头函数没有自己的this，继承外层上下文绑定的this。

` let obj = { age : 20 , info : function ( ) { return () => { console.log( this.age); //this继承的是外层上下文绑定的this } } } let person = { age : 28 }; let info = obj.info(); info(); //20 let info2 = obj.info.call(person); info2(); //28 复制代码`
> 
> 
> 
> 
> [点击查看更多](
> https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2FYvetteLau%2FStep-By-Step%2Fissues%2F1
> )
> 
> 

### 2.JS中原始类型有哪几种？null 是对象吗？原始数据类型和复杂数据类型有什么区别？(2019-05-21) ###

#### 目前，JS原始类型有六种，分别为: ####

* Boolean
* String
* Number
* Undefined
* Null
* Symbol(ES6新增)

ES10新增了一种基本数据类型：BigInt

复杂数据类型只有一种: Object

null 不是一个对象，尽管 ` typeof null` 输出的是 ` object` ，这是一个历史遗留问题，JS 的最初版本中使用的是 32 位系统，为了性能考虑使用低位存储变量的类型信息，000 开头代表是对象， ` null` 表示为全零，所以将它错误的判断为 ` object` 。

#### 基本数据类型和复杂数据类型的区别为: ####

* 内存的分配不同

* 

基本数据类型存储在栈中。

* 

复杂数据类型存储在堆中，栈中存储的变量，是指向堆中的引用地址。

* 访问机制不同

* 基本数据类型是按值访问
* 复杂数据类型按引用访问，JS不允许直接访问保存在堆内存中的对象，在访问一个对象时，首先得到的是这个对象在栈内存中的地址，然后再按照这个地址去获得这个对象中的值。

* 复制变量时不同(a=b)

* 基本数据类型：a=b;是将b中保存的原始值的副本赋值给新变量a，a和b完全独立，互不影响
* 复杂数据类型：a=b;将b保存的对象内存的引用地址赋值给了新变量a;a和b指向了同一个堆内存地址，其中一个值发生了改变，另一个也会改变。

` let b = { age : 10 } let a = b; a.age = 20 ; console.log(b); //{ age: 20 } 复制代码` * 参数传递的不同(实参/形参)

函数传参都是按值传递(栈中的存储的内容)：基本数据类型，拷贝的是值；复杂数据类型，拷贝的是引用地址

` //基本数据类型 let b = 10 function change ( info ) { info= 20 ; } //info=b;基本数据类型，拷贝的是值得副本，二者互不干扰 change(b); console.log(b); //10 复制代码` ` //复杂数据类型 let b = { age : 10 } function change ( info ) { info.age = 20 ; } //info=b;根据第三条差异，可以看出，拷贝的是地址的引用，修改互相影响。 change(b); console.log(b); //{ age: 20 } 复制代码`
> 
> 
> 
> 
> [点击查看更多](
> https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2FYvetteLau%2FStep-By-Step%2Fissues%2F5
> )
> 
> 

### 3.说一说你对HTML5语义化的理解(2019-05-22) ###

语义化意味着顾名思义，HTML5的语义化指的是合理正确的使用语义化的标签来创建页面结构，如 header,footer,nav，从标签上即可以直观的知道这个标签的作用，而不是滥用div。

#### 语义化的优点有: ####

* 代码结构清晰，易于阅读，利于开发和维护
* 方便其他设备解析（如屏幕阅读器）根据语义渲染网页。
* 有利于搜索引擎优化（SEO），搜索引擎爬虫会根据不同的标签来赋予不同的权重

#### 语义化标签主要有： ####

* title：主要用于页面的头部的信息介绍
* header：定义文档的页眉
* nav：主要用于页面导航
* main：规定文档的主要内容。对于文档来说应当是唯一的。它不应包含在文档中重复出现的内容，比如侧栏、导航栏、版权信息、站点标志或搜索表单。
* article：独立的自包含内容
* h1~h6：定义标题
* ul: 用来定义无序列表
* ol: 用来定义有序列表
* address：定义文档或文章的作者/拥有者的联系信息。
* canvas：用于绘制图像
* dialog：定义一个对话框、确认框或窗口
* aside：定义其所处内容之外的内容。 ` <aside>` 的内容可用作文章的侧栏。
* section：定义文档中的节（section、区段）。比如章节、页眉、页脚或文档中的其他部分。
* figure：规定独立的流内容（图像、图表、照片、代码等等）。figure 元素的内容应该与主内容相关，但如果被删除，则不应对文档流产生影响。
* details：描述文档或者文档某一部分细节
* mark：义带有记号的文本。

#### 语义化应用 ####

例如使用这些可视化标签，构建下面的页面结构：

![](https://user-gold-cdn.xitu.io/2019/5/26/16af36698b85075f?imageView2/0/w/1280/h/960/ignore-error/1)

对于早期不支持 HTML5 的浏览器，如IE8及更早之前的版本，我们可以引入 html5shiv 来支持。

> 
> 
> 
> [点击查看更多](
> https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2FYvetteLau%2FStep-By-Step%2Fissues%2F8
> )
> 
> 

### 4.如何让 (a == 1 && a == 2 && a == 3) 的值为true？ ###

#### 利用隐式转换规则 ####

` ==` 操作符在左右数据类型不一致时，会先进行隐式转换。

` a == 1 && a == 2 && a == 3` 的值意味着其不可能是基本数据类型。因为如果 a 是 null 或者是 undefined bool类型，都不可能返回true。

因此可以推测 a 是复杂数据类型，JS 中复杂数据类型只有 ` object` ，回忆一下，Object 转换为原始类型会调用什么方法？

* 

如果部署了 ` [Symbol.toPrimitive]` 接口，那么调用此接口，若返回的不是基本数据类型，抛出错误。

* 

如果没有部署 ` [Symbol.toPrimitive]` 接口，那么根据要转换的类型，先调用 ` valueOf` / ` toString`

* 非Date类型对象， ` hint` 是 ` default` 时，调用顺序为： ` valueOf` >>> ` toString` ，即 ` valueOf` 返回的不是基本数据类型，才会继续调用 ` toString` ，如果 ` toString` 返回的还不是基本数据类型，那么抛出错误。
* 如果 ` hint` 是 ` string` (Date对象默人的hint是string) ，调用顺序为： ` toString` >>> ` valueOf` ，即 ` toString` 返回的不是基本数据类型，才会继续调用 ` valueOf` ，如果 ` valueOf` 返回的还不是基本数据类型，那么抛出错误。
* 如果 ` hint` 是 ` number` ，调用顺序为： ` valueOf` >>> ` toString`

` var obj = { [ Symbol.toPrimitive](hint) { console.log(hint); return 10 ; }, valueOf() { console.log( 'valueOf' ); return 20 ; }, toString() { console.log( 'toString' ); return 'hello' ; } } console.log(obj + 'yvette' ); //default //如果没有部署 [Symbol.toPrimitive]接口，调用顺序为`valueOf` >>> `toString` console.log(obj == 'yvette' ); //default //如果没有部署 [Symbol.toPrimitive]接口，调用顺序为`valueOf` >>> `toString` console.log(obj * 10 ); //number //如果没有部署 [Symbol.toPrimitive]接口，调用顺序为`valueOf` >>> `toString` console.log( Number (obj)); //number //如果没有部署 [Symbol.toPrimitive]接口，调用顺序为`valueOf` >>> `toString` console.log( String (obj)); //string //如果没有部署 [Symbol.toPrimitive]接口，调用顺序为`toString` >>> `valueOf` 复制代码`

那么对于这道题，只要 ` [Symbol.toPrimitive]` 接口，第一次返回的值是 1，然后递增，即成功成立。

` let a = { [ Symbol.toPrimitive]: ( function ( hint ) { let i = 1 ; //闭包的特性之一：i 不会被回收 return function ( ) { return i++; } })() } console.log(a == 1 && a == 2 && a == 3 ); //true 复制代码`

调用 ` valueOf` 接口的情况：

` let a = { valueOf : ( function ( ) { let i = 1 ; //闭包的特性之一：i 不会被回收 return function ( ) { return i++; } })() } console.log(a == 1 && a == 2 && a == 3 ); //true 复制代码`

另外，除了i自增的方法外，还可以利用 正则，如下

` let a = { reg : /\d/g , valueOf () { return this.reg.exec( 123 )[ 0 ] } } console.log(a == 1 && a == 2 && a == 3 ); //true 复制代码`

调用 ` toString` 接口的情况，不再做说明。

#### 利用数据劫持 ####

使用 ` Object.defineProperty` 定义的属性，在获取属性时，会调用 ` get` 方法。利用这个特性，我们在 ` window` 对象上定义 ` a` 属性，如下：

` let i = 1 ; Object.defineProperty( window , 'a' , { get : function ( ) { return i++; } }); console.log(a == 1 && a == 2 && a == 3 ); //true 复制代码`

ES6 新增了 ` Proxy` ，此处我们同样可以利用 ` Proxy` 去实现，如下：

` let a = new Proxy ({}, { i : 1 , get : function ( ) { return () => this.i++; } }); console.log(a == 1 && a == 2 && a == 3 ); // true 复制代码`

#### 数组的 ` toString` 接口默认调用数组的 ` join` 方法，重写数组的 ` join` 方法。 ####

` let a = [ 1 , 2 , 3 ]; a.join = a.shift; console.log(a == 1 && a == 2 && a == 3 ); //true 复制代码`

#### 利用 ` with` 关键字 ####

我本人对 ` with` 向来是敬而远之的。不过 ` with` 的确也是此题方法之一：

` let i = 0 ; with ({ get a() { return ++i; } }) { console.log(a == 1 && a == 2 && a == 3 ); //true } 复制代码`
> 
> 
> 
> 
> [点击查看更多](
> https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2FYvetteLau%2FStep-By-Step%2Fissues%2F9
> )
> 
> 

### 5.防抖(debounce)函数的作用是什么？有哪些应用场景，请实现一个防抖函数。 ###

#### 防抖函数的作用 ####

防抖函数的作用就是控制函数在一定时间内的执行次数。防抖意味着N秒内函数只会被执行一次，如果N秒内再次被触发，则重新计算延迟时间。

举例说明：小思最近在减肥，但是她非常贪吃。为此，与其男朋友约定好，如果10天不吃零食，就可以购买一个包(不要问为什么是包，因为包治百病)。但是如果中间吃了一次零食，那么就要重新计算时间，直到小思坚持10天没有吃零食，才能购买一个包。所以，管不住嘴的小思，没有机会买包(悲伤的故事)...这就是 **防抖** 。

不管吃没吃零食，每10天买一个包，中间想买包，忍着，等到第十天的时候再买，这种情况是 **节流** 。如何控制女朋友的消费，各位攻城狮们，get到了吗？防抖可比节流有效多了！

#### 防抖应用场景 ####

* 搜索框输入查询，如果用户一直在输入中，没有必要不停地调用去请求服务端接口，等用户停止输入的时候，再调用，设置一个合适的时间间隔，有效减轻服务端压力。
* 表单验证
* 按钮提交事件。
* 浏览器窗口缩放，resize事件等。

#### 防抖函数实现 ####

` immediate` 为 true 时，表示开始会立即触发一次。

` immediate` 为 false 时，表示最后一次一定会触发。

` lodash` 中的 ` debounce` 的第三个参数 ` option` ，提供了 ` leading` 和 ` trailing` 两个参数。

` function debounce ( func, wait, immediate=true ) { let timeout, context, args; // 延迟执行函数 const later = () => setTimeout( () => { // 延迟函数执行完毕，清空定时器 timeout = null // 延迟执行的情况下，函数会在延迟函数中执行 // 使用到之前缓存的参数和上下文 if (!immediate) { func.apply(context, args); context = args = null ; } }, wait); let debounced = function (...params ) { if (!timeout) { timeout = later(); if (immediate) { //立即执行 func.apply( this , params); } else { //闭包 context = this ; args = params; } } else { clearTimeout(timeout); timeout = later(); } } debounced.cancel = function ( ) { clearTimeout(timeout); timeout = null ; }; return debounced; }; 复制代码`
> 
> 
> 
> 
> [点击查看更多](
> https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2FYvetteLau%2FStep-By-Step%2Fissues%2F10
> )
> 
> 

参考文章：

[1] [www.ecma-international.org/ecma-262/6.…]( https://link.juejin.im?target=https%3A%2F%2Fwww.ecma-international.org%2Fecma-262%2F6.0%2F%23sec-completion-record-specification-type )

[2] [嗨，你真的懂this吗？]( https://juejin.im/post/5c96d0c751882511c832ff7b )

[3] [【面试篇】寒冬求职季之你必须要懂的原生JS(上)]( https://juejin.im/post/5cab0c45f265da2513734390 )

[4] [【面试篇】寒冬求职季之你必须要懂的原生JS(中)]( https://juejin.im/post/5cbd1e33e51d45789161d053 )

[5] [digcss.com/throttle-th…]( https://link.juejin.im?target=https%3A%2F%2Fdigcss.com%2Fthrottle-throttling-and-shake-debouncing-details%2F )

谢谢各位小伙伴愿意花费宝贵的时间阅读本文，如果本文给了您一点帮助或者是启发，请不要吝啬你的赞和Star，您的肯定是我前进的最大动力。 [github.com/YvetteLau/B…]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2FYvetteLau%2FBlog )

> 
> 
> 
> 关注公众号，加入交流群。
> 
> 

![](https://user-gold-cdn.xitu.io/2019/5/14/16ab621cfa97956f?imageView2/0/w/1280/h/960/ignore-error/1)