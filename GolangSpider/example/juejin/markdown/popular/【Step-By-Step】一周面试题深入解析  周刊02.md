# 【Step-By-Step】一周面试题深入解析 / 周刊02 #

### 关于【Step-By-Step】 ###

> 
> 
> 
> [Step-By-Step](
> https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2FYvetteLau%2FStep-By-Step
> ) (点击进入项目) 是我于 ` 2019-05-20` 开始的一个项目，每个工作日发布一道面试题。
> 
> 
> 
> 每个周末我会仔细阅读大家的答案，整理最一份较优答案出来，因本人水平有限，有误的地方，大家及时指正。
> 
> 

> 
> 
> 
> 如果想 **加群** 学习，扫码 **二维码** (
> https://link.juejin.im?target=https%3A%2F%2Fuser-gold-cdn.xitu.io%2F2019%2F5%2F26%2F16af37df7708267b%3Fw%3D243%26amp%3Bh%3D245%26amp%3Bf%3Djpeg%26amp%3Bs%3D36789
> ) (点击查看)，添加我为好友，验证信息为 **加入组织** ，我拉你进群。
> 
> 

[【Step-By-Step】一周面试题深入解析 / 周刊 01]( https://juejin.im/post/5cea6e5fe51d45775e33f4de )

> 
> 
> 
> 本周面试题一览:
> 
> 

* 节流(throttle)函数的作用是什么？有哪些应用场景，请实现一个节流函数
* 说一说你对JS执行上下文栈和作用域链的理解？
* 什么是BFC？BFC的布局规则是什么？如何创建BFC？
* let、const、var 的区别有哪些？
* 深拷贝和浅拷贝的区别是什么？如何实现一个深拷贝？

### 6. 节流(throttle)函数的作用是什么？有哪些应用场景，请实现一个节流函数。(2019-05-27) ###

#### 节流函数的作用 ####

节流函数的作用是规定一个单位时间，在这个单位时间内最多只能触发一次函数执行，如果这个单位时间内多次触发函数，只能有一次生效。

举例说明：小明的妈妈和小明约定好，如果小明在周考中取得满分，那么当月可以带他去游乐场玩，但是一个月最多只能去一次。

这其实就是一个节流的例子，在一个月的时间内，去游乐场最多只能触发一次。即使这个时间周期内，小明取得多次满分。

#### 节流应用场景 ####

1.按钮点击事件

2.拖拽事件

3.onScoll

4.计算鼠标移动的距离(mousemove)

#### 节流函数实现 ####

> 
> 
> 
> 利用时间戳实现
> 
> 

` function throttle ( func, delay ) { var lastTime = 0 ; function throttled ( ) { var context = this ; var args = arguments ; var nowTime = Date.now(); if (nowTime > lastTime + delay) { func.apply(context, args); lastTime = nowTime; } } //节流函数最终返回的是一个函数 return throttled; } 复制代码`
> 
> 
> 
> 
> 利用定时器实现
> 
> 

` function throttle ( func, delay ) { var timeout = null ; function throttled ( ) { var context = this ; var args = arguments ; if (!timeout) { timeout = setTimeout( () => { func.apply(context, args); clearTimeout(timeout); timeout= null }, delay); } } return throttled; } 复制代码`

时间戳和定时器的方式都没有考虑最后一次执行的问题，比如有个按钮点击事件，设置的间隔时间是1S，在第0.5S，1.8S，2.2S点击，那么只有0.5S和1.8S的两次点击能够触发函数执行，而最后一次的2.2S会被忽略。

> 
> 
> 
> 组合实现，允许设置第一次或者最后一次是否触发函数执行
> 
> 

` function throttle ( func, wait, options ) { var timeout, context, args, result; var previous = 0 ; if (!options) options = {}; var later = function ( ) { previous = options.leading === false ? 0 : Date.now() || new Date ().getTime(); timeout = null ; result = func.apply(context, args); if (!timeout) context = args = null ; }; var throttled = function ( ) { var now = Date.now() || new Date ().getTime(); if (!previous && options.leading === false ) previous = now; var remaining = wait - (now - previous); context = this ; args = arguments ; if (remaining <= 0 || remaining > wait) { if (timeout) { clearTimeout(timeout); timeout = null ; } previous = now; result = func.apply(context, args); if (!timeout) context = args = null ; } else if (!timeout && options.trailing !== false ) { // 判断是否设置了定时器和 trailing timeout = setTimeout(later, remaining); } return result; }; throttled.cancel = function ( ) { clearTimeout(timeout); previous = 0 ; timeout = context = args = null ; }; return throttled; } 复制代码`

使用很简单:

` btn.onclick = throttle(handle, 1000 , { leading : true , trailing : true }); 复制代码`
> 
> 
> 
> 
> [点击查看更多](
> https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2FYvetteLau%2FStep-By-Step%2Fissues%2F12
> )
> 
> 

### 7. 说一说你对JS执行上下文栈和作用域链的理解？(2019-05-28) ###

在开始说明JS上下文栈和作用域之前，我们先说明下JS上下文以及作用域的概念。

#### [JS执行上下文]( https://link.juejin.im?target=https%3A%2F%2Ftc39.github.io%2Fecma262%2F%3Fnsukey%3DrQHqMrFpKq6JJN%252F%252FOeubPCslaSTSRyuc%252FXCznnIDze1SGzwva5SZtzixJ13p2gAlxua95Xa7fraZXwj5tyLRDK33%252BpNhyfKR%252FxyzhWNyB%252FqaIlsDGyQBckNoHQGPveOB24M%252BcK%252FgF8Tg1ehUGLWiCvumxdgcQwZOWj2BGfD3n%252FY%253D%23sec-execution-contexts ) ####

执行上下文就是当前 JavaScript 代码被解析和执行时所在环境的抽象概念， JavaScript 中运行任何的代码都是在执行上下文中运行。

> 
> 
> 
> 执行上下文类型分为：
> 
> 

* 全局执行上下文
* 函数执行上下文
* eval函数执行上下文(不被推荐)

执行上下文创建过程中，需要做以下几件事:

* 创建变量对象：首先初始化函数的参数arguments，提升函数声明和变量声明。
* 创建作用域链（Scope Chain）：在执行期上下文的创建阶段，作用域链是在变量对象之后创建的。
* 确定this的值，即 ResolveThisBinding

#### 作用域 ####

**作用域** 负责收集和维护由所有声明的标识符（变量）组成的一系列查询，并实施一套非常严格的规则，确定当前执行的代码对这些标识符的访问权限。—— 摘录自《你不知道的JavaScript》(上卷)

作用域有两种工作模型：词法作用域和动态作用域，JS采用的是 **词法作用域** 工作模型，词法作用域意味着作用域是由书写代码时变量和函数声明的位置决定的。( ` with` 和 ` eval` 能够修改词法作用域，但是不推荐使用，对此不做特别说明)

> 
> 
> 
> 作用域分为：
> 
> 

* 全局作用域
* 函数作用域
* 块级作用域

#### JS执行上下文栈(后面简称执行栈) ####

执行栈，也叫做调用栈，具有 **LIFO** (后进先出) 结构，用于存储在代码执行期间创建的所有执行上下文。

> 
> 
> 
> 规则如下：
> 
> 

* 首次运行JavaScript代码的时候,会创建一个全局执行的上下文并Push到当前的执行栈中，每当发生函数调用，引擎都会为该函数创建一个新的函数执行上下文并Push当前执行栈的栈顶。
* 当栈顶的函数运行完成后，其对应的函数执行上下文将会从执行栈中Pop出，上下文的控制权将移动到当前执行栈的下一个执行上下文。

以一段代码具体说明：

` function fun3 ( ) { console.log( 'fun3' ) } function fun2 ( ) { fun3(); } function fun1 ( ) { fun2(); } fun1(); 复制代码`

` Global Execution Context` (即全局执行上下文)首先入栈，过程如下：

![](https://user-gold-cdn.xitu.io/2019/6/2/16b1775c9934039d?imageView2/0/w/1280/h/960/ignore-error/1)

伪代码:

` //全局执行上下文首先入栈 ECStack.push(globalContext); //执行fun1(); ECStack.push(<fun1> functionContext); //fun1中又调用了fun2; ECStack.push(<fun2> functionContext); //fun2中又调用了fun3; ECStack.push(<fun3> functionContext); //fun3执行完毕 ECStack.pop(); //fun2执行完毕 ECStack.pop(); //fun1执行完毕 ECStack.pop(); //javascript继续顺序执行下面的代码，但ECStack底部始终有一个 全局上下文（globalContext）; 复制代码`

#### 作用域链 ####

作用域链就是从当前作用域开始一层一层向上寻找某个变量，直到找到全局作用域还是没找到，就宣布放弃。这种一层一层的关系，就是作用域链。

如：

` var a = 10 ; function fn1 ( ) { var b = 20 ; console.log(fn2) function fn2 ( ) { a = 20 } return fn2; } fn1()(); 复制代码`

fn2作用域链 = [fn2作用域, fn1作用域，全局作用域]

![](https://user-gold-cdn.xitu.io/2019/6/2/16b1775c974b1f50?imageView2/0/w/1280/h/960/ignore-error/1)

> 
> 
> 
> [点击查看更多](
> https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2FYvetteLau%2FStep-By-Step%2Fissues%2F14
> )
> 
> 

### 8. 什么是BFC？BFC的布局规则是什么？如何创建BFC？(2019-05-29) ###

#### 什么是BFC ####

BFC 是 Block Formatting Context 的缩写，即块格式化上下文。我们来看一下CSS2.1规范中对 BFC 的说明。

> 
> 
> 
> Floats, absolutely positioned elements, block containers (such as
> inline-blocks, table-cells, and table-captions) that are not block boxes,
> and block boxes with 'overflow' other than 'visible' (except when that
> value has been propagated to the viewport) establish new block formatting
> contexts for their contents.
> 
> 
> 
> 浮动、绝对定位的元素、非块级盒子的块容器（如inline-blocks、table-cells 和 table-captions），以及 `
> overflow` 的值不为 ` visible` （该值已传播到视区时除外）为其内容建立新的块格式上下文。
> 
> 

因此，如果想要深入的理解BFC，我们需要了解以下两个概念：

> 
> 
> 
> 1.Box
> 
> 

> 
> 
> 
> 2.Formatting Context
> 
> 

#### Box ####

Box 是 CSS 布局的对象和基本单位，页面是由若干个Box组成的。

元素的类型 和 ` display` 属性，决定了这个 Box 的类型。不同类型的 Box 会参与不同的 Formatting Context。

#### Formatting Context ####

Formatting Context 是页面的一块渲染区域，并且有一套渲染规则，决定了其子元素将如何定位，以及和其它元素的关系和相互作用。

Formatting Context 有 BFC (Block formatting context)，IFC (Inline formatting context)，FFC (Flex formatting context) 和 GFC (Grid formatting context)。FFC 和 GFC 为 CC3 中新增。

#### [BFC布局规则]( https://link.juejin.im?target=https%3A%2F%2Fwww.w3.org%2FTR%2F2011%2FREC-CSS2-20110607%2Fvisuren.html%23block-formatting ) ####

* BFC内，盒子依次垂直排列。
* BFC内，两个盒子的垂直距离由 ` margin` 属性决定。属于同一个BFC的两个相邻Box的margin会发生重叠【符合合并原则的margin合并后是使用大的margin】
* BFC内，每个盒子的左外边缘接触内部盒子的左边缘（对于从右到左的格式，右边缘接触）。即使在存在浮动的情况下也是如此。除非创建新的BFC。
* BFC的区域不会与float box重叠。
* BFC就是页面上的一个隔离的独立容器，容器里面的子元素不会影响到外面的元素。反之也如此。
* 计算BFC的高度时，浮动元素也参与计算。

#### 如何创建BFC ####

* 根元素
* 浮动元素（float 属性不为 none）
* position 为 absolute 或 relative
* overflow 不为 visible 的块元素
* display 为 inline-block, table-cell, table-caption

#### BFC的应用 ####

1.防止 margin 重叠

` < style >.a { height : 100px ; width : 100px ; margin : 50px ; background : pink; } </ style > < body > < div class = "a" > </ div > < div class = "a" > </ div > </ body > 复制代码`

两个div直接的 ` margin` 是50px，发生了 ` margin` 的重叠。

![](https://user-gold-cdn.xitu.io/2019/6/2/16b1775c97052243?imageView2/0/w/1280/h/960/ignore-error/1)

> 
> 
> 
> 根据BFC规则，同一个BFC内的两个两个相邻Box的 ` margin` 会发生重叠，因此我们可以在div外面再嵌套一层容器，并且触发该容器生成一个
> BFC，这样 ` <div class="a"></div>` 就会属于两个 BFC，自然也就不会再发生 ` margin` 重叠
> 
> 

` < style >.a { height : 100px ; width : 100px ; margin : 50px ; background : pink; }.container { overflow : auto; /*触发生成BFC*/ } </ style > < body > < div class = "container" > < div class = "a" > </ div > </ div > < div class = "a" > </ div > </ body > 复制代码`

![](https://user-gold-cdn.xitu.io/2019/6/2/16b1775c9a13484b?imageView2/0/w/1280/h/960/ignore-error/1)

2.清除内部浮动

` < style >.a { height : 100px ; width : 100px ; margin : 10px ; background : pink; float : left; }.container { width : 120px ; border : 2px solid black; } </ style > < body > < div class = "container" > < div class = "a" > </ div > </ div > </ body > 复制代码`

![](https://user-gold-cdn.xitu.io/2019/6/2/16b1775ca6136431?imageView2/0/w/1280/h/960/ignore-error/1)

> 
> 
> 
> container 的高度没有被撑开，如果我们希望 container 的高度能够包含浮动元素，那么可以创建一个新的 BFC，因为根据 BFC
> 的规则，计算 BFC 的高度时，浮动元素也参与计算。
> 
> 

` < style >.a { height : 100px ; width : 100px ; margin : 10px ; background : pink; float : left; }.container { width : 120px ; display : inline-block; /*触发生成BFC*/ border : 2px solid black; } </ style > 复制代码`

![](https://user-gold-cdn.xitu.io/2019/6/2/16b1775c984cecf6?imageView2/0/w/1280/h/960/ignore-error/1)

3.自适应多栏布局

` < style > body { width : 500px ; }.a { height : 150px ; width : 100px ; background : pink; float : left; }.b { height : 200px ; background : blue; } </ style > < body > < div class = "a" > </ div > < div class = "b" > </ div > </ body > 复制代码`

![](https://user-gold-cdn.xitu.io/2019/6/2/16b1775cc7bbe073?imageView2/0/w/1280/h/960/ignore-error/1)

> 
> 
> 
> 根据规则，BFC的区域不会与float box重叠。因此，可以触发生成一个新的BFC，如下：
> 
> 

` < style >.b { height : 200px ; overflow : hidden; /*触发生成BFC*/ background : blue; } </ style > 复制代码`

![](https://user-gold-cdn.xitu.io/2019/6/2/16b1775cc7f30da0?imageView2/0/w/1280/h/960/ignore-error/1)

> 
> 
> 
> [点击查看更多](
> https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2FYvetteLau%2FStep-By-Step%2Fissues%2F15
> )
> 
> 

### 9. let、const、var 的区别有哪些？(2019-05-30) ###

+----------+----------+------------+----------+--------------+--------+----------+
| 声明方式 | 变量提升 | 暂时性死区 | 重复声明 | 块作用域有效 | 初始值 | 重新赋值 |
+----------+----------+------------+----------+--------------+--------+----------+
| var      | 会       | 不存在     | 允许     | 不是         | 非必须 | 允许     |
| let      | 不会     | 存在       | 不允许   | 是           | 非必须 | 允许     |
| const    | 不会     | 存在       | 不允许   | 是           | 必须   | 不允许   |
+----------+----------+------------+----------+--------------+--------+----------+

#### 1.let/const 定义的变量不会出现变量提升，而 var 定义的变量会提升。 ####

` a = 10 ; var a; //正常 复制代码` ` a = 10 ; let a; //ReferenceError 复制代码`

#### 2.相同作用域中，let 和 const 不允许重复声明，var 允许重复声明。 ####

` let a = 10 ; var a = 20 ; //抛出异常：SyntaxError: Identifier 'a' has already been declared 复制代码`

#### 3.cosnt 声明变量时必须设置初始值 ####

` const a; //SyntaxError: Missing initializer in const declaration 复制代码`

#### 4.const 声明一个只读的常量，这个常量不可改变。 ####

这里有一个非常重要的点即是：复杂数据类型，存储在栈中的是堆内存的地址，存在栈中的这个地址是不变的，但是存在堆中的值是可以变得。有没有相当常量指针/指针常量~

` const a = 20 ; const b = { age : 18 , star : 500 } 复制代码`

一图胜万言，如下图所示，不变的是栈内存中 a 存储的 20，和 b 中存储的 0x0012ff21（瞎编的一个数字）。而 {age: 18, star: 200} 是可变的。思考下如果想希望一个对象是不可变的，应该用什么方法？

![](https://user-gold-cdn.xitu.io/2019/6/2/16b1775cc80345b3?imageView2/0/w/1280/h/960/ignore-error/1)

#### 5.let/const 声明的变量仅在块级作用域中有效。而 var 声明的变量在块级作用域外仍能访问到。 ####

` { let a = 10 ; const b = 20 ; var c = 30 ; } console.log(a); //ReferenceError console.log(b); //ReferenceError console.log(c); //30 复制代码`

在 let/const 之前，最早学习JS的时候，也曾被下面这个问题困扰：

期望： ` a[0]() 输出 0 , a[1]() 输出 1 , a[2]() 输出 2 , ...`

` var a = []; for ( var i = 0 ; i < 10 ; i++) { a[i] = function ( ) { console.log(i); }; } a[ 6 ](); // 10 复制代码`

![](https://user-gold-cdn.xitu.io/2019/6/2/16b17a2978042a96?imageView2/0/w/1280/h/960/ignore-error/1)

虽然后来知道了为什么，但是想要得到自己需要的结果，还得整个闭包，我...我做错了什么，要这么对我...

` var a = []; for ( var i = 0 ; i < 10 ; i++) { a[i] = ( function ( j ) { return function ( ) { console.log(j); } })(i) } a[ 6 ](); // 6 复制代码`

有了 let 之后，终于不要这么麻烦了。

` var a = []; for ( let i = 0 ; i < 10 ; i++) { a[i] = function ( ) { console.log(i); }; } a[ 6 ](); // 6 复制代码`

美滋滋，有没有~

美是美了，但是总得问自己为什么吧~

` var i` 为什么输出的是 10，这是因为 i 在全局范围内都是有效的，相当于只有一个变量 i，等执行到 ` a[6]()` 的时候，这个 i 的值是什么？请大声说出来。

再看 let , 我们说 let 声明的变量仅在块级作用域内有效，变量i是let声明的，当前的 i 只在本轮循环有效，所以每一次循环的 i 其实都是一个新的变量。有兴趣的小伙伴可以查看 babel 编译后的代码。

#### 6.顶层作用域中 var 声明的变量挂在window上(浏览器环境) ####

` var a = 10 ; console.log( window.a); //10 复制代码`

#### 7.let/const有暂时性死区的问题，即let/const 声明的变量，在定义之前都是不可用的。如果使用会抛出错误。 ####

只要块级作用域内存在let命令，它所声明的变量就“绑定”（binding）这个区域，不再受外部的影响。

` var a = 10 ; if ( true ) { a = 20 ; // ReferenceError let a; } 复制代码`

在代码块内，使用 let/const 命令声明变量之前，该变量都是不可用的，也就意味着 typeof 不再是一个百分百安全的操作。

` console.log( typeof b); //undefined console.log(a); //ReferenceError let a = 10 ; 复制代码`
> 
> 
> 
> 
> [点击查看更多](
> https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2FYvetteLau%2FStep-By-Step%2Fissues%2F16
> )
> 
> 

### 10. 深拷贝和浅拷贝的区别是什么？如何实现一个深拷贝？(2019-05-31) ###

深拷贝和浅拷贝是针对复杂数据类型来说的。

#### 深拷贝 ####

> 
> 
> 
> 深拷贝复制变量值，对于非基本类型的变量，则递归至基本类型变量后，再复制。
> 深拷贝后的对象与原来的对象是完全隔离的，互不影响，对一个对象的修改并不会影响另一个对象。
> 
> 

#### 浅拷贝 ####

> 
> 
> 
> 浅拷贝是会将对象的每个属性进行依次复制，但是当对象的属性值是引用类型时，实质复制的是其引用，当引用指向的值改变时也会跟着变化。
> 
> 

可以使用 ` for in` 、 ` Object.assign` 、 扩展运算符 `...` 、 ` Array.prototype.slice()` 、 ` Array.prototype.concat()` 等，例如:

` let obj = { name : 'Yvette' , age : 18 , hobbies : [ 'reading' , 'photography' ] } let obj2 = Object.assign({}, obj); let obj3 = {...obj}; obj.name = 'Jack' ; obj.hobbies.push( 'coding' ); console.log(obj); //{ name: 'Jack', age: 18,hobbies: [ 'reading', 'photography', 'coding' ] } console.log(obj2); //{ name: 'Yvette', age: 18,hobbies: [ 'reading', 'photography', 'coding' ] } console.log(obj3); //{ name: 'Yvette', age: 18,hobbies: [ 'reading', 'photography', 'coding' ] } 复制代码`

可以看出浅拷贝只最第一层属性进行了拷贝，当第一层的属性值是基本数据类型时，新的对象和原对象互不影响，但是如果第一层的属性值是复杂数据类型，那么新对象和原对象的属性值其指向的是同一块内存地址。来看一下使用 ` for in` 实现浅拷贝。

` let obj = { name : 'Yvette' , age : 18 , hobbies : [ 'reading' , 'photography' ] } let newObj = {}; for ( let key in obj){ newObj[key] = obj[key]; //这一步不需要多说吧，复杂数据类型栈中存的是对应的地址，因此赋值操作，相当于两个属性值指向同一个内存空间 } console.log(newObj); //{ name: 'Yvette', age: 18, hobbies: [ 'reading', 'photography' ] } obj.age = 20 ; obj.hobbies.pop(); console.log(newObj); //{ name: 'Yvette', age: 18, hobbies: [ 'reading' ] } 复制代码`

#### 深拷贝实现 ####

> 
> 
> 
> 1.深拷贝最简单的实现是: ` JSON.parse(JSON.stringify(obj))`
> 
> 

` let obj = { name : 'Yvette' , age : 18 , hobbies : [ 'reading' , 'photography' ] } let newObj = JSON.parse( JSON.stringify(obj)); //newObj和obj互不影响 obj.hobbies.push( 'coding' ); console.log(newObj); //{ name: 'Yvette', age: 18, hobbies: [ 'reading', 'photography' ] } 复制代码`

` JSON.parse(JSON.stringify(obj))` 是最简单的实现方式，但是有一点缺陷：

1.对象的属性值是函数时，无法拷贝。

` let obj = { name : 'Yvette' , age : 18 , hobbies : [ 'reading' , 'photography' ], sayHi : function ( ) { console.log(sayHi); } } let newObj = JSON.parse( JSON.stringify(obj)); console.log(newObj); //{ name: 'Yvette', age: 18, hobbies: [ 'reading', 'photography' ] } 复制代码`

2.原型链上的属性无法获取

` function Super ( ) { } Super.prototype.location = 'NanJing' ; function Child ( name, age, hobbies ) { this.name = name; this.age = age; } Child.prototype = new Super(); let obj = new Child( 'Yvette' , 18 ); console.log(obj.location); //NanJing let newObj = JSON.parse( JSON.stringify(obj)); console.log(newObj); //{ name: 'Yvette', age: 18} console.log(newObj.location); //undefined;原型链上的属性无法获取 复制代码`

3.不能正确的处理 Date 类型的数据

4.不能处理 RegExp

5.会忽略 symbol

6.会忽略 undefined

` let obj = { time : new Date (), reg : /\d{3}/ , sym : Symbol ( 10 ), name : undefined } let obj2 = JSON.parse( JSON.stringify(obj)); console.log(obj2); //{ time: '2019-06-02T08:16:44.625Z', reg: {} } 复制代码`
> 
> 
> 
> 
> 2.实现一个 deepClone 函数
> 
> 

* 如果是基本数据类型，直接返回
* 如果是 ` RegExp` 或者 ` Date` 类型，返回对应类型
* 如果是复杂数据类型，递归。
` function deepClone ( obj ) { //递归拷贝 if (obj instanceof RegExp ) return new RegExp (obj); if (obj instanceof Date ) return new Date (obj); if (obj === null || typeof obj !== 'object' ) { //如果不是复杂数据类型，直接返回 return obj; } /** * 如果obj是数组，那么 obj.constructor 是 [Function: Array] * 如果obj是对象，那么 obj.constructor 是 [Function: Object] */ let t = new obj.constructor(); for ( let key in obj) { //如果 obj[key] 是复杂数据类型，递归 if (obj.hasOwnProperty(key)){ //是否是自身的属性 t[key] = deepClone(obj[key]); } } return t; } 复制代码`

测试:

` function Super ( ) { } Super.prototype.location = 'NanJing' ; function Child ( name, age, hobbies ) { this.name = name; this.age = age; this.hobbies = hobbies; } Child.prototype = new Super(); let obj = new Child( 'Yvette' , 18 , [ 'reading' , 'photography' ]); obj.sayHi = function ( ) { console.log( 'hi' ); } console.log(obj.location); //NanJing let newObj = deepClone(obj); console.log(newObj); // console.log(newObj.location); //NanJing 可以获取到原型链上的属性 newObj.sayHi(); //hi 函数属性拷贝正常 复制代码`
> 
> 
> 
> 
> 3.循环引用
> 
> 

前面的deepClone没有考虑循环引用的问题，例如对象的某个属性，是这个对象本身。

` function deepClone ( obj, hash = new WeakMap( )) { //递归拷贝 if (obj instanceof RegExp ) return new RegExp (obj); if (obj instanceof Date ) return new Date (obj); if (obj === null || typeof obj !== 'object' ) { //如果不是复杂数据类型，直接返回 return obj; } if (hash.has(obj)) { return hash.get(obj); } /** * 如果obj是数组，那么 obj.constructor 是 [Function: Array] * 如果obj是对象，那么 obj.constructor 是 [Function: Object] */ let t = new obj.constructor(); hash.set(obj, t); for ( let key in obj) { //如果 obj[key] 是复杂数据类型，递归 if (obj.hasOwnProperty(key)){ //是否是自身的属性 if (obj[key] && typeof obj[key] === 'object' ) { t[key] = deepClone(obj[key], hash); } else { t[key] = obj[key]; } } } return t; } 复制代码`

测试代码:

` const obj1 = { name : 'Yvetet' , sayHi : function ( ) { console.log( 'Hi' ) }, time : new Date (), info : { } } obj1.circle = obj1; obj1.info.base = obj1; obj1.info.name = obj1.name; console.log(deepClone(obj1)); 复制代码`
> 
> 
> 
> 
> [点击查看更多](
> https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2FYvetteLau%2FStep-By-Step%2Fissues%2F17
> )
> 
> 

### 参考文章： ###

[1] [www.ecma-international.org/ecma-262/6.…]( https://link.juejin.im?target=https%3A%2F%2Fwww.ecma-international.org%2Fecma-262%2F6.0%2F%23sec-completion-record-specification-type )

[2] [【译】理解 Javascript 执行上下文和执行栈]( https://juejin.im/post/5bdfd3e151882516c6432c32 )

[3] [css-tricks.com/debouncing-…]( https://link.juejin.im?target=https%3A%2F%2Fcss-tricks.com%2Fdebouncing-throttling-explained-examples%2F )

[4] [github.com/mqyqingfeng…]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fmqyqingfeng%2FBlog%2Fissues%2F4 )

[5] [www.cnblogs.com/coco1s/p/40…]( https://link.juejin.im?target=https%3A%2F%2Fwww.cnblogs.com%2Fcoco1s%2Fp%2F4017544.html )

[6] [www.cnblogs.com/wangfupeng1…]( https://link.juejin.im?target=https%3A%2F%2Fwww.cnblogs.com%2Fwangfupeng1988%2Fp%2F4000798.html )

[7] [www.w3.org/TR/2011/REC…]( https://link.juejin.im?target=https%3A%2F%2Fwww.w3.org%2FTR%2F2011%2FREC-CSS2-20110607%2Fvisuren.html%23block-boxes )

[8] [github.com/mqyqingfeng…]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fmqyqingfeng%2FBlog%2Fissues%2F4 )

谢谢各位小伙伴愿意花费宝贵的时间阅读本文，如果本文给了您一点帮助或者是启发，请不要吝啬你的赞和Star，您的肯定是我前进的最大动力。 [github.com/YvetteLau/B…]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2FYvetteLau%2FBlog )

![](https://user-gold-cdn.xitu.io/2019/5/11/16aa69598012fc3c?imageView2/0/w/1280/h/960/ignore-error/1)

> 
> 
> 
> 关注公众号，加入技术交流群。
> 
> 

![](https://user-gold-cdn.xitu.io/2019/6/2/16b1775cda2f992a?imageView2/0/w/1280/h/960/ignore-error/1)