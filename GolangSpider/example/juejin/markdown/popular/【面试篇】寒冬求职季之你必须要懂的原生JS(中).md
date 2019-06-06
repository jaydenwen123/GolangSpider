# 【面试篇】寒冬求职季之你必须要懂的原生JS(中) #

互联网寒冬之际，各大公司都缩减了HC，甚至是采取了“裁员”措施，在这样的大环境之下，想要获得一份更好的工作，必然需要付出更多的努力。

一年前，也许你搞清楚闭包，this，原型链，就能获得认可。但是现在，很显然是不行了。本文梳理出了一些面试中有一定难度的高频原生JS问题，部分知识点可能你之前从未关注过，或者看到了，却没有仔细研究，但是它们却非常重要。

本文将以 真实的面试题 的形式来呈现知识点，大家在阅读时， 建议不要先看我的答案， 而是自己先思考一番。尽管，本文所有的答案，都是我在翻阅各种资料，思考并验证之后，才给出的 (绝非复制粘贴而来) 。但因水平有限，本人的答案未必是最优的，如果您有更好的答案，欢迎在 issue 中留言。

本文篇幅较长，但是满满的都是干货！并且还埋伏了可爱的表情包，希望小伙伴们能够坚持读完。

写文超级真诚的小姐姐祝愿大家都能找到心仪的工作。

如果你还没读过上篇【上篇和中篇并无依赖关系，您可以读过本文之后再阅读上篇】，可戳 [【面试篇】寒冬求职季之你必须要懂的原生JS(上)]( https://juejin.im/post/5cab0c45f265da2513734390 )

![](https://user-gold-cdn.xitu.io/2019/4/22/16a42bddfb7fde39?imageView2/0/w/1280/h/960/ignore-error/1)

**小姐姐花了近百个小时才完成这篇文章，篇幅较长，希望大家阅读时多花点耐心，力求真正的掌握相关知识点。**

### 1.说一说JS异步发展史 ###

异步最早的解决方案是回调函数，如事件的回调，setInterval/setTimeout中的回调。但是回调函数有一个很常见的问题，就是回调地狱的问题(稍后会举例说明);

为了解决回调地狱的问题，社区提出了Promise解决方案，ES6将其写进了语言标准。Promise解决了回调地狱的问题，但是Promise也存在一些问题，如错误不能被try catch，而且使用Promise的链式调用，其实并没有从根本上解决回调地狱的问题，只是换了一种写法。

ES6中引入 Generator 函数，Generator是一种异步编程解决方案，Generator 函数是协程在 ES6 的实现，最大特点就是可以交出函数的执行权，Generator 函数可以看出是异步任务的容器，需要暂停的地方，都用yield语句注明。但是 Generator 使用起来较为复杂。

ES7又提出了新的异步解决方案:async/await，async是 Generator 函数的语法糖，async/await 使得异步代码看起来像同步代码，异步编程发展的目标就是让异步逻辑的代码看起来像同步一样。

> 
> 
> 
> 1.回调函数: callback
> 
> 

` //node读取文件 fs.readFile(xxx, 'utf-8' , function ( err, data ) { //code }); 复制代码`

回调函数的使用场景(包括但不限于):

* 事件回调
* Node API
* setTimeout/setInterval中的回调函数

异步回调嵌套会导致代码难以维护，并且不方便统一处理错误，不能try catch 和 回调地狱(如先读取A文本内容，再根据A文本内容读取B再根据B的内容读取C...)。

` fs.readFile(A, 'utf-8' , function ( err, data ) { fs.readFile(B, 'utf-8' , function ( err, data ) { fs.readFile(C, 'utf-8' , function ( err, data ) { fs.readFile(D, 'utf-8' , function ( err, data ) { //.... }); }); }); }); 复制代码`
> 
> 
> 
> 
> 2.Promise
> 
> 

Promise 主要解决了回调地狱的问题，Promise 最早由社区提出和实现，ES6 将其写进了语言标准，统一了用法，原生提供了Promise对象。

那么我们看看Promise是如何解决回调地狱问题的，仍然以上文的readFile为例。

` function read ( url ) { return new Promise ( ( resolve, reject ) => { fs.readFile(url, 'utf8' , (err, data) => { if (err) reject(err); resolve(data); }); }); } read(A).then( data => { return read(B); }).then( data => { return read(C); }).then( data => { return read(D); }).catch( reason => { console.log(reason); }); 复制代码`

想要运行代码看效果，请戳(小姐姐使用的是VS的 Code Runner 执行代码): [github.com/YvetteLau/B…]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2FYvetteLau%2FBlog%2Fblob%2Fmaster%2FJS%2FAsync%2Fpromise.js )

思考一下在Promise之前，你是如何处理异步并发问题的，假设有这样一个需求：读取三个文件内容，都读取成功后，输出最终的结果。有了Promise之后，又如何处理呢？代码可戳: [github.com/YvetteLau/B…]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2FYvetteLau%2FBlog%2Fblob%2Fmaster%2FJS%2FAsync%2Findex.js )

注: 可以使用 bluebird 将接口 promise化;

**引申:** Promise有哪些优点和问题呢？

> 
> 
> 
> 3.Generator
> 
> 

Generator 函数是 ES6 提供的一种异步编程解决方案，整个 Generator 函数就是一个封装的异步任务，或者说是异步任务的容器。异步操作需要暂停的地方，都用 yield 语句注明。

Generator 函数一般配合 yield 或 Promise 使用。Generator函数返回的是迭代器。对生成器和迭代器不了解的同学，请自行补习下基础。下面我们看一下 Generator 的简单使用:

` function * gen ( ) { let a = yield 111 ; console.log(a); let b = yield 222 ; console.log(b); let c = yield 333 ; console.log(c); let d = yield 444 ; console.log(d); } let t = gen(); //next方法可以带一个参数，该参数就会被当作上一个yield表达式的返回值 t.next( 1 ); //第一次调用next函数时，传递的参数无效 t.next( 2 ); //a输出2; t.next( 3 ); //b输出3; t.next( 4 ); //c输出4; t.next( 5 ); //d输出5; 复制代码`

为了让大家更好的理解上面代码是如何执行的，我画了一张图，分别对应每一次的next方法调用:

![](https://user-gold-cdn.xitu.io/2019/4/22/16a42bddfb8cdc0f?imageView2/0/w/1280/h/960/ignore-error/1)

仍然以上文的readFile为例，使用 Generator + co库来实现:

` const fs = require ( 'fs' ); const co = require ( 'co' ); const bluebird = require ( 'bluebird' ); const readFile = bluebird.promisify(fs.readFile); function * read ( ) { yield readFile(A, 'utf-8' ); yield readFile(B, 'utf-8' ); yield readFile(C, 'utf-8' ); //.... } co(read()).then( data => { //code }).catch( err => { //code }); 复制代码`

不使用co库，如何实现？能否自己写一个最简的my_co？请戳: [github.com/YvetteLau/B…]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2FYvetteLau%2FBlog%2Fblob%2Fmaster%2FJS%2FAsync%2Fgenerator.js )

PS: 如果你还不太了解 Generator/yield，建议阅读ES6相关文档。

> 
> 
> 
> 4.async/await
> 
> 

ES7中引入了 async/await 概念。async其实是一个语法糖，它的实现就是将Generator函数和自动执行器（co），包装在一个函数中。

async/await 的优点是代码清晰，不用像 Promise 写很多 then 链，就可以处理回调地狱的问题。错误可以被try catch。

` const fs = require ( 'fs' ); const bluebird = require ( 'bluebird' ); const readFile = bluebird.promisify(fs.readFile); async function read ( ) { await readFile(A, 'utf-8' ); await readFile(B, 'utf-8' ); await readFile(C, 'utf-8' ); //code } read().then( ( data ) => { //code }).catch( err => { //code }); 复制代码`

可执行代码，请戳： [github.com/YvetteLau/B…]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2FYvetteLau%2FBlog%2Fblob%2Fmaster%2FJS%2FAsync%2Fasync.js )

思考一下 async/await 如何处理异步并发问题的？ [github.com/YvetteLau/B…]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2FYvetteLau%2FBlog%2Fblob%2Fmaster%2FJS%2FAsync%2Findex.js )

如果你有更好的答案或想法，欢迎在这题目对应的github下留言： [说一说JS异步发展史]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2FYvetteLau%2FBlog%2Fissues%2F10 )

### 2.谈谈对 async/await 的理解，async/await 的实现原理是什么? ###

async/await 就是 Generator 的语法糖，使得异步操作变得更加方便。来张图对比一下:

![](https://user-gold-cdn.xitu.io/2019/4/22/16a42bddfc5c8515?imageView2/0/w/1280/h/960/ignore-error/1)

async 函数就是将 Generator 函数的星号（*）替换成 async，将 yield 替换成await。

> 
> 
> 
> 我们说 async 是 Generator 的语法糖，那么这个糖究竟甜在哪呢？
> 
> 

1）async函数内置执行器，函数调用之后，会自动执行，输出最后结果。而Generator需要调用next或者配合co模块使用。

2）更好的语义，async和await，比起星号和yield，语义更清楚了。async表示函数里有异步操作，await表示紧跟在后面的表达式需要等待结果。

3）更广的适用性。co模块约定，yield命令后面只能是 Thunk 函数或 Promise 对象，而async 函数的 await 命令后面，可以是 Promise 对象和原始类型的值。

4）返回值是Promise，async函数的返回值是 Promise 对象，Generator的返回值是 Iterator，Promise 对象使用起来更加方便。

> 
> 
> 
> async 函数的实现原理，就是将 Generator 函数和自动执行器，包装在一个函数里。
> 
> 

具体代码试下如下(和spawn的实现略有差异，个人觉得这样写更容易理解)，如果你想知道如何一步步写出 my_co ，可戳: [github.com/YvetteLau/B…]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2FYvetteLau%2FBlog%2Fblob%2Fmaster%2FJS%2FAsync%2Fmy_async.js )

` function my_co ( it ) { return new Promise ( ( resolve, reject ) => { function next ( data ) { try { var { value, done } = it.next(data); } catch (e){ return reject(e); } if (!done) { //done为true,表示迭代完成 //value 不一定是 Promise，可能是一个普通值。使用 Promise.resolve 进行包装。 Promise.resolve(value).then( val => { next(val); }, reject); } else { resolve(value); } } next(); //执行一次next }); } function * test ( ) { yield new Promise ( ( resolve, reject ) => { setTimeout(resolve, 100 ); }); yield new Promise ( ( resolve, reject ) => { // throw Error(1); resolve( 10 ) }); yield 10 ; return 1000 ; } my_co(test()).then( data => { console.log(data); //输出1000 }).catch( ( err ) => { console.log( 'err: ' , err); }); 复制代码`

如果你有更好的答案或想法，欢迎在这题目对应的github下留言： [谈谈对 async/await 的理解，async/await 的实现原理是什么?]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2FYvetteLau%2FBlog%2Fissues%2F11 )

### 3.使用 async/await 需要注意什么？ ###

* await 命令后面的Promise对象，运行结果可能是 rejected，此时等同于 async 函数返回的 Promise 对象被reject。因此需要加上错误处理，可以给每个 await 后的 Promise 增加 catch 方法；也可以将 await 的代码放在 ` try...catch` 中。
* 多个await命令后面的异步操作，如果不存在继发关系，最好让它们同时触发。
` //下面两种写法都可以同时触发 //法一 async function f1 ( ) { await Promise.all([ new Promise ( ( resolve ) => { setTimeout(resolve, 600 ); }), new Promise ( ( resolve ) => { setTimeout(resolve, 600 ); }) ]) } //法二 async function f2 ( ) { let fn1 = new Promise ( ( resolve ) => { setTimeout(resolve, 800 ); }); let fn2 = new Promise ( ( resolve ) => { setTimeout(resolve, 800 ); }) await fn1; await fn2; } 复制代码` * await命令只能用在async函数之中，如果用在普通函数，会报错。
* async 函数可以保留运行堆栈。
` /** * 函数a内部运行了一个异步任务b()。当b()运行的时候，函数a()不会中断，而是继续执行。 * 等到b()运行结束，可能a()早就* 运行结束了，b()所在的上下文环境已经消失了。 * 如果b()或c()报错，错误堆栈将不包括a()。 */ function b ( ) { return new Promise ( ( resolve, reject ) => { setTimeout(resolve, 200 ) }); } function c ( ) { throw Error ( 10 ); } const a = () => { b().then( () => c()); }; a(); /** * 改成async函数 */ const m = async () => { await b(); c(); }; m(); 复制代码`

报错信息如下，可以看出 async 函数可以保留运行堆栈。

![](https://user-gold-cdn.xitu.io/2019/4/22/16a42bddfc7d17e7?imageView2/0/w/1280/h/960/ignore-error/1)

如果你有更好的答案或想法，欢迎在这题目对应的github下留言： [使用 async/await 需要注意什么？]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2FYvetteLau%2FBlog%2Fissues%2F12 )

### 4.如何实现 Promise.race？ ###

在代码实现前，我们需要先了解 Promise.race 的特点：

* 

Promise.race返回的仍然是一个Promise. 它的状态与第一个完成的Promise的状态相同。它可以是完成（ resolves），也可以是失败（rejects），这要取决于第一个Promise是哪一种状态。

* 

如果传入的参数是不可迭代的，那么将会抛出错误。

* 

如果传的参数数组是空，那么返回的 promise 将永远等待。

* 

如果迭代包含一个或多个非承诺值和/或已解决/拒绝的承诺，则 Promise.race 将解析为迭代中找到的第一个值。

` Promise.race = function ( promises ) { //promises 必须是一个可遍历的数据结构，否则抛错 return new Promise ( ( resolve, reject ) => { if ( typeof promises[ Symbol.iterator] !== 'function' ) { //真实不是这个错误 Promise.reject( 'args is not iteratable!' ); } if (promises.length === 0 ) { return ; } else { for ( let i = 0 ; i < promises.length; i++) { Promise.resolve(promises[i]).then( ( data ) => { resolve(data); return ; }, (err) => { reject(err); return ; }); } } }); } 复制代码`

测试代码:

` //一直在等待态 Promise.race([]).then( ( data ) => { console.log( 'success ' , data); }, (err) => { console.log( 'err ' , err); }); //抛错 Promise.race().then( ( data ) => { console.log( 'success ' , data); }, (err) => { console.log( 'err ' , err); }); Promise.race([ new Promise ( ( resolve, reject ) => { setTimeout( () => { resolve( 100 ) }, 1000 ) }), new Promise ( ( resolve, reject ) => { setTimeout( () => { resolve( 200 ) }, 200 ) }), new Promise ( ( resolve, reject ) => { setTimeout( () => { reject( 100 ) }, 100 ) }) ]).then( ( data ) => { console.log(data); }, (err) => { console.log(err); }); 复制代码`

**引申:** Promise.all/Promise.reject/Promise.resolve/Promise.prototype.finally/Promise.prototype.catch 的实现原理，如果还不太会，戳: [Promise源码实现]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2FYvetteLau%2FBlog%2Fissues%2F2 )

如果你有更好的答案或想法，欢迎在这题目对应的github下留言： [如何实现 Promise.race？]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2FYvetteLau%2FBlog%2Fissues%2F13 )

![](https://user-gold-cdn.xitu.io/2019/4/22/16a42bddfe3a2cc2?imageView2/0/w/1280/h/960/ignore-error/1)

### 5.可遍历数据结构的有什么特点？ ###

一个对象如果要具备可被 for...of 循环调用的 Iterator 接口，就必须在其 Symbol.iterator 的属性上部署遍历器生成方法(或者原型链上的对象具有该方法)

**PS:** 遍历器对象根本特征就是具有next方法。每次调用next方法，都会返回一个代表当前成员的信息对象，具有value和done两个属性。

` //如为对象添加Iterator 接口; let obj = { name : "Yvette" , age : 18 , job : 'engineer' , [ Symbol.iterator]() { const self = this ; const keys = Object.keys(self); let index = 0 ; return { next() { if (index < keys.length) { return { value : self[keys[index++]], done : false }; } else { return { value : undefined , done : true }; } } }; } }; for ( let item of obj) { console.log(item); //Yvette 18 engineer } 复制代码`

使用 Generator 函数(遍历器对象生成函数)简写 Symbol.iterator 方法，可以简写如下:

` let obj = { name : "Yvette" , age : 18 , job : 'engineer' , * [ Symbol.iterator] () { const self = this ; const keys = Object.keys(self); for ( let index = 0 ;index < keys.length; index++) { yield self[keys[index]]; //yield表达式仅能使用在 Generator 函数中 } } }; 复制代码`
> 
> 
> 
> 
> 原生具备 Iterator 接口的数据结构如下。
> 
> 

* Array
* Map
* Set
* String
* TypedArray
* 函数的 arguments 对象
* NodeList 对象
* ES6 的数组、Set、Map 都部署了以下三个方法: entries() / keys() / values()，调用后都返回遍历器对象。

如果你有更好的答案或想法，欢迎在这题目对应的github下留言： [可遍历数据结构的有什么特点？]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2FYvetteLau%2FBlog%2Fissues%2F14 )

### 6.requestAnimationFrame 和 setTimeout/setInterval 有什么区别？使用 requestAnimationFrame 有哪些好处？ ###

在 requestAnimationFrame 之前，我们主要使用 setTimeout/setInterval 来编写JS动画。

编写动画的关键是循环间隔的设置，一方面，循环间隔足够短，动画效果才能显得平滑流畅；另一方面，循环间隔还要足够长，才能确保浏览器有能力渲染产生的变化。

大部分的电脑显示器的刷新频率是60HZ，也就是每秒钟重绘60次。大多数浏览器都会对重绘操作加以限制，不超过显示器的重绘频率，因为即使超过那个频率用户体验也不会提升。因此，最平滑动画的最佳循环间隔是 1000ms / 60 ，约为16.7ms。

setTimeout/setInterval 有一个显著的缺陷在于时间是不精确的，setTimeout/setInterval 只能保证延时或间隔不小于设定的时间。因为它们实际上只是把任务添加到了任务队列中，但是如果前面的任务还没有执行完成，它们必须要等待。

requestAnimationFrame 才有的是系统时间间隔，保持最佳绘制效率，不会因为间隔时间过短，造成过度绘制，增加开销；也不会因为间隔时间太长，使用动画卡顿不流畅，让各种网页动画效果能够有一个统一的刷新机制，从而节省系统资源，提高系统性能，改善视觉效果。

综上所述，requestAnimationFrame 和 setTimeout/setInterval 在编写动画时相比，优点如下:

1.requestAnimationFrame 不需要设置时间，采用系统时间间隔，能达到最佳的动画效果。

2.requestAnimationFrame 会把每一帧中的所有DOM操作集中起来，在一次重绘或回流中就完成。

3.当 requestAnimationFrame() 运行在后台标签页或者隐藏的 ` <iframe>` 里时，requestAnimationFrame() 会被暂停调用以提升性能和电池寿命（大多数浏览器中）。

requestAnimationFrame 使用(试试使用requestAnimationFrame写一个移动的小球，从A移动到B初):

` function step ( timestamp ) { //code... window.requestAnimationFrame(step); } window.requestAnimationFrame(step); 复制代码`

如果你有更好的答案或想法，欢迎在这题目对应的github下留言： [requestAnimationFrame 和 setTimeout/setInterval 有什么区别？使用 requestAnimationFrame 有哪些好处？]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2FYvetteLau%2FBlog%2Fissues%2F15 )

### 7.JS 类型转换的规则是什么？ ###

类型转换的规则三言两语说不清，真想哇得一声哭出来~

![](https://user-gold-cdn.xitu.io/2019/4/22/16a42bde024df49d?imageView2/0/w/1280/h/960/ignore-error/1)

JS中类型转换分为 强制类型转换 和 隐式类型转换 。

* 

通过 Number()、parseInt()、parseFloat()、toString()、String()、Boolean(),进行强制类型转换。

* 

逻辑运算符(&&、 ||、 !)、运算符(+、-、*、/)、关系操作符(>、 <、 <= 、>=)、相等运算符(==)或者 if/while 的条件，可能会进行隐式类型转换。

**强制类型转换**

> 
> 
> 
> 1.Number() 将任意类型的参数转换为数值类型
> 
> 

规则如下:

* 如果是布尔值，true和false分别被转换为1和0
* 如果是数字，返回自身
* 如果是 null，返回 0
* 如果是 undefined，返回 ` NAN`
* 如果是字符串，遵循以下规则: * 如果字符串中只包含数字(或者是 ` 0X` / ` 0x` 开头的十六进制数字字符串，允许包含正负号)，则将其转换为十进制
* 如果字符串中包含有效的浮点格式，将其转换为浮点数值
* 如果是空字符串，将其转换为0
* 如不是以上格式的字符串，均返回 ` NaN`

* 如果是Symbol，抛出错误
* 如果是对象，则调用对象的 ` valueOf()` 方法，然后依据前面的规则转换返回的值。如果转换的结果是 ` NaN` ，则调用对象的 ` toString()` 方法，再次依照前面的规则转换返回的字符串值。

部分内置对象调用默认的 ` valueOf` 的行为:

+----------+----------------------------------------------+
|   对象   |                    返回值                    |
+----------+----------------------------------------------+
| Array    | 数组本身（对象类型）                         |
| Boolean  | 布尔值（原始类型）                           |
| Date     | 从 UTC 1970 年 1 月 1                        |
|          | 日午夜开始计算，到所封装的日期所经过的毫秒数 |
| Function | 函数本身（对象类型）                         |
| Number   | 数字值（原始类型）                           |
| Object   | 对象本身（对象类型）                         |
| String   | 字符串值（原始类型）                         |
+----------+----------------------------------------------+

` Number ( '0111' ); //111 Number ( '0X11' ) //17 Number ( null ); //0 Number ( '' ); //0 Number ( '1a' ); //NaN Number ( -0X11 ); //-17 复制代码`
> 
> 
> 
> 
> 2.parseInt(param, radix)
> 
> 

如果第一个参数传入的是字符串类型:

* 忽略字符串前面的空格，直至找到第一个非空字符，如果是空字符串，返回NaN
* 如果第一个字符不是数字符号或者正负号，返回NaN
* 如果第一个字符是数字/正负号，则继续解析直至字符串解析完毕或者遇到一个非数字符号为止

如果第一个参数传入的Number类型:

* 数字如果是0开头，则将其当作八进制来解析(如果是一个八进制数)；如果以0x开头，则将其当作十六进制来解析

如果第一个参数是 null 或者是 undefined，或者是一个对象类型：

* 返回 NaN

如果第一个参数是数组： 1. 去数组的第一个元素，按照上面的规则进行解析

如果第一个参数是Symbol类型： 1. 抛出错误

如果指定radix参数，以radix为基数进行解析

` parseInt ( '0111' ); //111 parseInt ( 0111 ); //八进制数 73 parseInt ( '' ); //NaN parseInt ( '0X11' ); //17 parseInt ( '1a' ) //1 parseInt ( 'a1' ); //NaN parseInt ([ '10aa' , 'aaa' ]); //10 parseInt ([]); //NaN; parseInt(undefined); 复制代码`
> 
> 
> 
> 
> parseFloat
> 
> 

规则和 ` parseInt` 基本相同，接受一个Number类型或字符串，如果是字符串中，那么只有第一个小数点是有效的。

> 
> 
> 
> toString()
> 
> 

规则如下:

* 如果是Number类型，输出数字字符串
* 如果是 null 或者是 undefined，抛错
* 如果是数组，那么将数组展开输出。空数组，返回 ` ''`
* 如果是对象，返回 ` [object Object]`
* 如果是Date, 返回日期的文字表示法
* 如果是函数，输出对应的字符串(如下demo)
* 如果是Symbol，输出Symbol字符串

` let arry = []; let obj = { a : 1 }; let sym = Symbol ( 100 ); let date = new Date (); let fn = function ( ) { console.log( '稳住，我们能赢！' )} let str = 'hello world' ; console.log([].toString()); // '' console.log([ 1 , 2 , 3 , undefined , 5 , 6 ].toString()); //1,2,3,,5,6 console.log(arry.toString()); // 1,2,3 console.log(obj.toString()); // [object Object] console.log(date.toString()); // Sun Apr 21 2019 16:11:39 GMT+0800 (CST) console.log(fn.toString()); // function () {console.log('稳住，我们能赢！')} console.log(str.toString()); // 'hello world' console.log(sym.toString()); // Symbol(100) console.log( undefined.toString()); // 抛错 console.log( null.toString()); // 抛错 复制代码`
> 
> 
> 
> 
> String()
> 
> 

` String()` 的转换规则与 ` toString()` 基本一致，最大的一点不同在于 ` null` 和 ` undefined` ，使用 String 进行转换，null 和 undefined对应的是字符串 ` 'null'` 和 ` 'undefined'`

> 
> 
> 
> Boolean
> 
> 

除了 undefined、 null、 false、 ''、 0(包括 +0，-0)、 NaN 转换出来是false，其它都是true.

**隐式类型转换**

> 
> 
> 
> && 、|| 、 ! 、 if/while 的条件判断
> 
> 

需要将数据转换成 Boolean 类型，转换规则同 Boolean 强制类型转换

> 
> 
> 
> 运算符: + - * /
> 
> 

` +` 号操作符，不仅可以用作数字相加，还可以用作字符串拼接。

仅当 ` +` 号两边都是数字时，进行的是加法运算。如果两边都是字符串，直接拼接，无需进行隐式类型转换。

除了上面的情况外，如果操作数是对象、数值或者布尔值，则调用toString()方法取得字符串值(toString转换规则)。对于 undefined 和 null，分别调用String()显式转换为字符串，然后再进行拼接。

` console.log({}+ 10 ); //[object Object]10 console.log([ 1 , 2 , 3 , undefined , 5 , 6 ] + 10 ); //1,2,3,,5,610 复制代码`

` -` 、 ` *` 、 ` /` 操作符针对的是运算，如果操作值之一不是数值，则被隐式调用Number()函数进行转换。如果其中有一个转换除了为NaN，结果为NaN.

> 
> 
> 
> 关系操作符: ==、>、< 、<=、>=
> 
> 

` >` , ` <` ， ` <=` ， ` >=`

* 如果两个操作值都是数值，则进行数值比较
* 如果两个操作值都是字符串，则比较字符串对应的字符编码值
* 如果有一方是Symbol类型，抛出错误
* 除了上述情况之外，都进行Number()进行类型转换，然后再进行比较。

**注：** NaN是非常特殊的值，它不和任何类型的值相等，包括它自己，同时它与任何类型的值比较大小时都返回false。

` console.log( 10 > {}); //返回false. /** *{}.valueOf ---> {} *{}.toString() ---> '[object Object]' ---> NaN *NaN 和 任何类型比大小，都返回 false */ 复制代码`
> 
> 
> 
> 
> 相等操作符： ` ==`
> 
> 

* 如果类型相同，无需进行类型转换。
* 如果其中一个操作值是 null 或者是 undefined，那么另一个操作符必须为 null 或者 undefined 时，才返回 true，否则都返回 false.
* 如果其中一个是 Symbol 类型，那么返回 false.
* 两个操作值是否为 string 和 number，就会将字符串转换为 number
* 如果一个操作值是 boolean，那么转换成 number
* 如果一个操作值为 object 且另一方为 string、number 或者 symbol，是的话就会把 object 转为原始类型再进行判断(调用object的valueOf/toString方法进行转换)

> 
> 
> 
> 对象如何转换成原始数据类型
> 
> 

如果部署了 [Symbol.toPrimitive] 接口，那么调用此接口，若返回的不是基础数据类型，抛出错误。

如果没有部署 [Symbol.toPrimitive] 接口，那么先返回 valueOf() 的值，若返回的不是基础类型的值，再返回 toString() 的值，若返回的不是基础类型的值， 则抛出异常。

` //先调用 valueOf, 后调用 toString let obj = { [ Symbol.toPrimitive]() { return 200 ; }, valueOf() { return 300 ; }, toString() { return 'Hello' ; } } //如果 valueOf 返回的不是基本数据类型，则会调用 toString， //如果 toString 返回的也不是基本数据类型，会抛出错误 console.log(obj + 200 ); //400 复制代码`

![](https://user-gold-cdn.xitu.io/2019/4/22/16a42bde28971b0d?imageView2/0/w/1280/h/960/ignore-error/1)

如果你有更好的答案或想法，欢迎在这题目对应的github下留言： [JS 类型转换的规则是什么？]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2FYvetteLau%2FBlog%2Fissues%2F16 )

### 8.简述下对 webWorker 的理解？ ###

HTML5则提出了 Web Worker 标准，表示js允许多线程，但是子线程完全受主线程控制并且不能操作dom，只有主线程可以操作dom，所以js本质上依然是单线程语言。

web worker就是在js单线程执行的基础上开启一个子线程，进行程序处理，而不影响主线程的执行，当子线程执行完之后再回到主线程上，在这个过程中不影响主线程的执行。子线程与主线程之间提供了数据交互的接口postMessage和onmessage，来进行数据发送和接收。

` var worker = new Worker( './worker.js' ); //创建一个子线程 worker.postMessage( 'Hello' ); worker.onmessage = function ( e ) { console.log(e.data); //Hi worker.terminate(); //结束线程 }; 复制代码` ` //worker.js onmessage = function ( e ) { console.log(e.data); //Hello postMessage( "Hi" ); //向主进程发送消息 }; 复制代码`

仅是最简示例代码，项目中通常是将一些耗时较长的代码，放在子线程中运行。

如果你有更好的答案或想法，欢迎在这题目对应的github下留言： [简述下对 webWorker 的理解]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2FYvetteLau%2FBlog%2Fissues%2F17 )

### 9.ES6模块和CommonJS模块的差异？ ###

* 

ES6模块在编译时，就能确定模块的依赖关系，以及输入和输出的变量。

CommonJS 模块，运行时加载。

* 

ES6 模块自动采用严格模式，无论模块头部是否写了 ` "use strict";`

* 

require 可以做动态加载，import 语句做不到，import 语句必须位于顶层作用域中。

* 

ES6 模块中顶层的 this 指向 undefined，CommonJS 模块的顶层 this 指向当前模块。

* 

CommonJS 模块输出的是一个值的拷贝，ES6 模块输出的是值的引用。

CommonJS 模块输出的是值的拷贝，也就是说，一旦输出一个值，模块内部的变化就影响不到这个值。如：

` //name.js var name = 'William' ; setTimeout( () => name = 'Yvette' , 200 ); module.exports = { name }; //index.js const name = require ( './name' ); console.log(name); //William setTimeout( () => console.log(name), 300 ); //William 复制代码`

对比 ES6 模块看一下:

ES6 模块的运行机制与 CommonJS 不一样。JS 引擎对脚本静态分析的时候，遇到模块加载命令 import ，就会生成一个只读引用。等到脚本真正执行时，再根据这个只读引用，到被加载的那个模块里面去取值。

` //name.js var name = 'William' ; setTimeout( () => name = 'Yvette' , 200 ); export { name }; //index.js import { name } from './name' ; console.log(name); //William setTimeout( () => console.log(name), 300 ); //Yvette 复制代码`

如果你有更好的答案或想法，欢迎在这题目对应的github下留言： [ES6模块和CommonJS模块的差异？]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2FYvetteLau%2FBlog%2Fissues%2F18 )

### 10.浏览器事件代理机制的原理是什么？ ###

在说浏览器事件代理机制原理之前，我们首先了解一下事件流的概念，早期浏览器，IE采用的是事件冒泡事件流，而Netscape采用的则是事件捕获。"DOM2级事件"把事件流分为三个阶段，捕获阶段、目标阶段、冒泡阶段。现代浏览器也都遵循此规范。

![](https://user-gold-cdn.xitu.io/2019/4/22/16a42bde5aadba59?imageView2/0/w/1280/h/960/ignore-error/1)

> 
> 
> 
> 那么事件代理是什么呢？
> 
> 

事件代理又称为事件委托，在祖先级DOM元素绑定一个事件，当触发子孙级DOM元素的事件时，利用事件冒泡的原理来触发绑定在祖先级DOM的事件。因为事件会从目标元素一层层冒泡至document对象。

> 
> 
> 
> 为什么要事件代理？
> 
> 

* 

添加到页面上的事件数量会影响页面的运行性能，如果添加的事件过多，会导致网页的性能下降。采用事件代理的方式，可以大大减少注册事件的个数。

* 

事件代理的当时，某个子孙元素是动态增加的，不需要再次对其进行事件绑定。

* 

不用担心某个注册了事件的DOM元素被移除后，可能无法回收其事件处理程序，我们只要把事件处理程序委托给更高层级的元素，就可以避免此问题。

> 
> 
> 
> 如将页面中的所有click事件都代理到document上:
> 
> 

addEventListener 接受3个参数，分别是要处理的事件名、处理事件程序的函数和一个布尔值。布尔值默认为false。表示冒泡阶段调用事件处理程序，若设置为true，表示在捕获阶段调用事件处理程序。

` document.addEventListener( 'click' , function ( e ) { console.log(e.target); /** * 捕获阶段调用调用事件处理程序，eventPhase是 1; * 处于目标，eventPhase是2 * 冒泡阶段调用事件处理程序，eventPhase是 1； */ console.log(e.eventPhase); }); 复制代码`

如果你有更好的答案或想法，欢迎在这题目对应的github下留言： [浏览器事件代理机制的原理是什么？]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2FYvetteLau%2FBlog%2Fissues%2F19 )

![](https://user-gold-cdn.xitu.io/2019/4/22/16a42bde5d8322d4?imageView2/0/w/1280/h/960/ignore-error/1)

### 11.js如何自定义事件？ ###

> 
> 
> 
> 自定义 DOM 事件(不考虑IE9之前版本)
> 
> 

自定义事件有三种方法,一种是使用 ` new Event()` , 另一种是 ` createEvent('CustomEvent')` , 另一种是 ` new customEvent()`

* 使用 ` new Event()`

获取不到 ` event.detail`

` let btn = document.querySelector( '#btn' ); let ev = new Event( 'alert' , { bubbles : true , //事件是否冒泡;默认值false cancelable: true , //事件能否被取消;默认值false composed: false }); btn.addEventListener( 'alert' , function ( event ) { console.log(event.bubbles); //true console.log(event.cancelable); //true console.log(event.detail); //undefined }, false ); btn.dispatchEvent(ev); 复制代码` * 使用 ` createEvent('CustomEvent')` (DOM3)

要创建自定义事件，可以调用 ` createEvent('CustomEvent')` ，返回的对象有 initCustomEvent 方法，接受以下四个参数:

* type: 字符串，表示触发的事件类型，如此处的'alert'
* bubbles: 布尔值： 表示事件是否冒泡
* cancelable: 布尔值，表示事件是否可以取消
* detail: 任意值，保存在 event 对象的 detail 属性中

` let btn = document.querySelector( '#btn' ); let ev = btn.createEvent( 'CustomEvent' ); ev.initCustomEvent( 'alert' , true , true , 'button' ); btn.addEventListener( 'alert' , function ( event ) { console.log(event.bubbles); //true console.log(event.cancelable); //true console.log(event.detail); //button }, false ); btn.dispatchEvent(ev); 复制代码` * 使用 ` new customEvent()` (DOM4)

使用起来比 ` createEvent('CustomEvent')` 更加方便

` var btn = document.querySelector( '#btn' ); /* * 第一个参数是事件类型 * 第二个参数是一个对象 */ var ev = new CustomEvent( 'alert' , { bubbles : 'true' , cancelable : 'true' , detail : 'button' }); btn.addEventListener( 'alert' , function ( event ) { console.log(event.bubbles); //true console.log(event.cancelable); //true console.log(event.detail); //button }, false ); btn.dispatchEvent(ev); 复制代码`
> 
> 
> 
> 
> 自定义非 DOM 事件(观察者模式)
> 
> 

EventTarget类型有一个单独的属性handlers，用于存储事件处理程序（观察者）。

addHandler() 用于注册给定类型事件的事件处理程序；

fire() 用于触发一个事件；

removeHandler() 用于注销某个事件类型的事件处理程序。

` function EventTarget ( ) { this.handlers = {}; } EventTarget.prototype = { constructor :EventTarget, addHandler : function ( type,handler ) { if ( typeof this.handlers[type] === "undefined" ){ this.handlers[type] = []; } this.handlers[type].push(handler); }, fire : function ( event ) { if (!event.target){ event.target = this ; } if ( this.handlers[event.type] instanceof Array ){ const handlers = this.handlers[event.type]; handlers.forEach( ( handler )=> { handler(event); }); } }, removeHandler : function ( type,handler ) { if ( this.handlers[type] instanceof Array ){ const handlers = this.handlers[type]; for ( var i = 0 ,len = handlers.length; i < len; i++){ if (handlers[i] === handler){ break ; } } handlers.splice(i, 1 ); } } } //使用 function handleMessage ( event ) { console.log(event.message); } //创建一个新对象 var target = new EventTarget(); //添加一个事件处理程序 target.addHandler( "message" , handleMessage); //触发事件 target.fire({ type : "message" , message : "Hi" }); //Hi //删除事件处理程序 target.removeHandler( "message" ,handleMessage); //再次触发事件，没有事件处理程序 target.fire({ type : "message" , message : "Hi" }); 复制代码`

如果你有更好的答案或想法，欢迎在这题目对应的github下留言： [js如何自定义事件？]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2FYvetteLau%2FBlog%2Fissues%2F20 )

### 12.跨域的方法有哪些？原理是什么？ ###

知其然知其所以然，在说跨域方法之前，我们先了解下什么叫跨域，浏览器有同源策略，只有当“协议”、“域名”、“端口号”都相同时，才能称之为是同源，其中有一个不同，即是跨域。

那么同源策略的作用是什么呢？同源策略限制了从同一个源加载的文档或脚本如何与来自另一个源的资源进行交互。这是一个用于隔离潜在恶意文件的重要安全机制。

那么我们又为什么需要跨域呢？一是前端和服务器分开部署，接口请求需要跨域，二是我们可能会加载其它网站的页面作为iframe内嵌。

> 
> 
> 
> **跨域的方法有哪些？**
> 
> 

> 
> 
> 
> 常用的跨域方法
> 
> 

* jsonp

尽管浏览器有同源策略，但是 ` <script>` 标签的 src 属性不会被同源策略所约束，可以获取任意服务器上的脚本并执行。jsonp 通过插入script标签的方式来实现跨域，参数只能通过url传入，仅能支持get请求。

实现原理:

Step1: 创建 callback 方法

Step2: 插入 script 标签

Step3: 后台接受到请求，解析前端传过去的 callback 方法，返回该方法的调用，并且数据作为参数传入该方法

Step4: 前端执行服务端返回的方法调用

下面代码仅为说明 jsonp 原理，项目中请使用成熟的库。分别看一下前端和服务端的简单实现：

` //前端代码 function jsonp ( {url, params, cb} ) { return new Promise ( ( resolve, reject ) => { //创建script标签 let script = document.createElement( 'script' ); //将回调函数挂在 window 上 window [cb] = function ( data ) { resolve(data); //代码执行后，删除插入的script标签 document.body.removeChild(script); } //回调函数加在请求地址上 params = {...params, cb} //wb=b&cb=show let arrs = []; for ( let key in params) { arrs.push( ` ${key} = ${params[key]} ` ); } script.src = ` ${url} ? ${arrs.join( '&' )} ` ; document.body.appendChild(script); }); } //使用 function sayHi ( data ) { console.log(data); } jsonp({ url : 'http://localhost:3000/say' , params : { //code }, cb : 'sayHi' }).then( data => { console.log(data); }); 复制代码` ` //express启动一个后台服务 let express = require ( 'express' ); let app = express(); app.get( '/say' , (req, res) => { let {cb} = req.query; //获取传来的callback函数名，cb是key res.send( ` ${cb} ('Hello!')` ); }); app.listen( 3000 ); 复制代码`

从今天起，jsonp的原理就要了然于心啦~

![](https://user-gold-cdn.xitu.io/2019/4/22/16a42bde62fd6497?imageView2/0/w/1280/h/960/ignore-error/1)

* cors

jsonp 只能支持 get 请求，cors 可以支持多种请求。cors 并不需要前端做什么工作。

> 
> 
> 
> 简单跨域请求:
> 
> 

只要服务器设置的Access-Control-Allow-Origin Header和请求来源匹配，浏览器就允许跨域

* 请求的方法是get，head或者post。
* Content-Type是application/x-www-form-urlencoded, multipart/form-data 或 text/plain中的一个值，或者不设置也可以，一般默认就是application/x-www-form-urlencoded。
* 请求中没有自定义的HTTP头部，如x-token。(应该是这几种头部 Accept，Accept-Language，Content-Language，Last-Event-ID，Content-Type）
` //简单跨域请求 app.use( ( req, res, next ) => { res.setHeader( 'Access-Control-Allow-Origin' , 'XXXX' ); }); 复制代码`
> 
> 
> 
> 
> 带预检(Preflighted)的跨域请求
> 
> 

不满于简单跨域请求的，即是带预检的跨域请求。服务端需要设置 Access-Control-Allow-Origin (允许跨域资源请求的域) 、 Access-Control-Allow-Methods (允许的请求方法) 和 Access-Control-Allow-Headers (允许的请求头)

` app.use( ( req, res, next ) => { res.setHeader( 'Access-Control-Allow-Origin' , 'XXX' ); res.setHeader( 'Access-Control-Allow-Headers' , 'XXX' ); //允许返回的头 res.setHeader( 'Access-Control-Allow-Methods' , 'XXX' ); //允许使用put方法请求接口 res.setHeader( 'Access-Control-Max-Age' , 6 ); //预检的存活时间 if (req.method === "OPTIONS" ) { res.end(); //如果method是OPTIONS，不做处理 } }); 复制代码`

更多CORS的知识可以访问: [HTTP访问控制（CORS）
]( https://link.juejin.im?target=https%3A%2F%2Fdeveloper.mozilla.org%2Fzh-CN%2Fdocs%2FWeb%2FHTTP%2FAccess_control_CORS )

* nginx 反向代理

使用nginx反向代理实现跨域，只需要修改nginx的配置即可解决跨域问题。

A网站向B网站请求某个接口时，向B网站发送一个请求，nginx根据配置文件接收这个请求，代替A网站向B网站来请求。 nginx拿到这个资源后再返回给A网站，以此来解决了跨域问题。

例如nginx的端口号为 8090，需要请求的服务器端口号为 3000。（localhost:8090 请求 localhost:3000/say）

nginx配置如下:

` server { listen 8090; server_name localhost; location / { root /Users/liuyan35/Test/Study/CORS/1-jsonp; index index.html index.htm; } location /say { rewrite ^/say/(.*)$ / $1 break ; proxy_pass http://localhost:3000; add_header 'Access-Control-Allow-Origin' '*' ; add_header 'Access-Control-Allow-Credentials' 'true' ; add_header 'Access-Control-Allow-Methods' 'GET, POST, OPTIONS' ; } # others } 复制代码` * websocket

Websocket 是 HTML5 的一个持久化的协议，它实现了浏览器与服务器的全双工通信，同时也是跨域的一种解决方案。

Websocket 不受同源策略影响，只要服务器端支持，无需任何配置就支持跨域。

前端页面在 8080 的端口。

` let socket = new WebSocket( 'ws://localhost:3000' ); //协议是ws socket.onopen = function ( ) { socket.send( 'Hi,你好' ); } socket.onmessage = function ( e ) { console.log(e.data) } 复制代码`

服务端 3000端口。可以看出websocket无需做跨域配置。

` let WebSocket = require ( 'ws' ); let wss = new WebSocket.Server({ port : 3000 }); wss.on( 'connection' , function ( ws ) { ws.on( 'message' , function ( data ) { console.log(data); //接受到页面发来的消息'Hi,你好' ws.send( 'Hi' ); //向页面发送消息 }); }); 复制代码` * postMessage

postMessage 通过用作前端页面之前的跨域，如父页面与iframe页面的跨域。window.postMessage方法，允许跨窗口通信，不论这两个窗口是否同源。

话说工作中两个页面之前需要通信的情况并不多，我本人工作中，仅使用过两次，一次是H5页面中发送postMessage信息，ReactNative的webview中接收此此消息，并作出相应处理。另一次是可轮播的页面，某个轮播页使用的是iframe页面，为了解决滑动的事件冲突，iframe页面中去监听手势，发送消息告诉父页面是否左滑和右滑。

> 
> 
> 
> 子页面向父页面发消息
> 
> 

父页面

` window.addEventListener( 'message' , (e) => { this.props.movePage(e.data); }, false ); 复制代码`

子页面(iframe):

` if ( /*左滑*/ ) { window.parent && window.parent.postMessage( -1 , '*' ) } else if ( /*右滑*/ ){ window.parent && window.parent.postMessage( 1 , '*' ) } 复制代码`
> 
> 
> 
> 
> 父页面向子页面发消息
> 
> 

父页面:

` let iframe = document.querySelector( '#iframe' ); iframe.onload = function ( ) { iframe.contentWindow.postMessage( 'hello' , 'http://localhost:3002' ); } 复制代码`

子页面:

` window.addEventListener( 'message' , function ( e ) { console.log(e.data); e.source.postMessage( 'Hi' , e.origin); //回消息 }); 复制代码` * node 中间件

node 中间件的跨域原理和nginx代理跨域，同源策略是浏览器的限制，服务端没有同源策略。

node中间件实现跨域的原理如下:

1.接受客户端请求

2.将请求 转发给服务器。

3.拿到服务器 响应 数据。

4.将 响应 转发给客户端。

> 
> 
> 
> 不常用跨域方法
> 
> 

以下三种跨域方式很少用，如有兴趣，可自行查阅相关资料。

* 

window.name + iframe

* 

location.hash + iframe

* 

document.domain (主域需相同)

如果你有更好的答案或想法，欢迎在这题目对应的github下留言： [跨域的方法有哪些？原理是什么？]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2FYvetteLau%2FBlog%2Fissues%2F21 )

### 13.js异步加载的方式有哪些？ ###

* 

` <script>` 的 defer 属性，HTML4 中新增

* 

` <script>` 的 async 属性，HTML5 中新增

` <script>` 标签打开defer属性，脚本就会异步加载。渲染引擎遇到这一行命令，就会开始下载外部脚本，但不会等它下载和执行，而是直接执行后面的命令。

defer 和 async 的区别在于: defer要等到整个页面在内存中正常渲染结束，才会执行；

async一旦下载完，渲染引擎就会中断渲染，执行这个脚本以后，再继续渲染。defer是“渲染完再执行”，async是“下载完就执行”。

如果有多个 defer 脚本，会按照它们在页面出现的顺序加载。

多个async脚本是不能保证加载顺序的。

* 动态插入 script 脚本
` function downloadJS ( ) { varelement = document.createElement( "script" ); element.src = "XXX.js" ; document.body.appendChild(element); } //何时的时候，调用上述方法 复制代码` * 有条件的动态创建脚本

如页面 onload 之后，

如果你有更好的答案或想法，欢迎在这题目对应的github下留言： [js异步加载的方式有哪些？]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2FYvetteLau%2FBlog%2Fissues%2F22 )

### 14.下面代码a在什么情况中打印出1？ ###

` //? if (a == 1 && a == 2 && a == 3 ) { console.log( 1 ); } 复制代码`

1.在类型转换的时候，我们知道了对象如何转换成原始数据类型。如果部署了 [Symbol.toPrimitive]，那么返回的就是 的返回值。当然，我们也可以把此函数部署在valueOf或者是toString接口上，效果相同。

` //利用闭包延长作用域的特性 let a = { [ Symbol.toPrimitive]: ( function ( ) { let i = 1 ; return function ( ) { return i++; } })() } 复制代码`

（1）. 比较 a == 1 时，会调用 [Symbol.toPrimitive]，此时 i 是 1，相等。 （2）. 继续比较 a == 2,调用 [Symbol.toPrimitive]，此时 i 是 2，相等。 （3）. 继续比较 a == 3,调用 [Symbol.toPrimitive]，此时 i 是 3，相等。

2.利用Object.defineProperty在window/global上定义a属性，获取a属性时，会调用get.

` let val = 1 ; Object.defineProperty( window , 'a' , { get : function ( ) { return val++; } }); 复制代码`

3.利用数组的特性。

` var a = [ 1 , 2 , 3 ]; a.join = a.shift; 复制代码`

数组的 ` toString` 方法返回一个字符串，该字符串由数组中的每个元素的 toString() 返回值经调用 join() 方法连接（由逗号隔开）组成。

因此，我们可以重新 join 方法。返回第一个元素，并将其删除。

如果你有更好的答案或想法，欢迎在这题目对应的github下留言： [下面代码a在什么情况中打印出1？ ]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2FYvetteLau%2FBlog%2Fissues%2F23 )

![](https://user-gold-cdn.xitu.io/2019/4/22/16a42bde699666d3?imageView2/0/w/1280/h/960/ignore-error/1) ![](https://user-gold-cdn.xitu.io/2019/4/22/16a42bde6c3954d4?imageView2/0/w/1280/h/960/ignore-error/1)

### 15.下面这段代码的输出是什么？ ###

` function Foo ( ) { getName = function ( ) { console.log( 1 )}; return this ; } Foo.getName = function ( ) { console.log( 2 )}; Foo.prototype.getName = function ( ) { console.log( 3 )}; var getName = function ( ) { console.log( 4 )}; function getName ( ) { console.log( 5 )}; Foo.getName(); getName(); Foo().getName(); getName(); new Foo.getName(); new Foo().getName(); new new Foo().getName(); 复制代码`

**说明：**一道经典的面试题，仅是为了帮助大家回顾一下知识点，加深理解，真实工作中，是不可能这样写代码的，否则，肯定会被打死的。

1.首先预编译阶段，变量声明与函数声明提升至其对应作用域的最顶端。

因此上面的代码编译后如下(函数声明的优先级先于变量声明):

` function Foo ( ) { getName = function ( ) { console.log( 1 )}; return this ; } function getName ( ) { console.log( 5 )}; //函数优先(函数首先被提升) var getName; //重复声明，被忽略 Foo.getName = function ( ) { console.log( 2 )}; Foo.prototype.getName = function ( ) { console.log( 3 )}; getName = function ( ) { console.log( 4 )}; 复制代码`

2. ` Foo.getName()` ;直接调用Foo上getName方法，输出2

3. ` getName()` ;输出4，getName被重新赋值了

4. ` Foo().getName()` ;执行Foo()，window的getName被重新赋值，返回this;浏览器环境中，非严格模式，this 指向 window，this.getName();输出为1.

如果是严格模式，this 指向 undefined，此处会抛出错误。

如果是node环境中，this 指向 global，node的全局变量并不挂在global上，因为global.getName对应的是undefined，不是一个function，会抛出错误。

5. ` getName()` ;已经抛错的自然走不动这一步了；继续浏览器非严格模式；window.getName被重新赋过值，此时再调用，输出的是1

6. ` new Foo.getName()` ;考察 [运算符优先级]( https://link.juejin.im?target=https%3A%2F%2Fdeveloper.mozilla.org%2Fzh-CN%2Fdocs%2FWeb%2FJavaScript%2FReference%2FOperators%2FOperator_Precedence ) 的知识，new 无参数列表，对应的优先级是18；成员访问操作符 `.` , 对应的优先级是 19。因此相当于是 ` new (Foo.getName)()` ;new操作符会执行构造函数中的方法，因此此处输出为 2.

7. ` new Foo().getName()` ；new 带参数列表，对应的优先级是19，和成员访问操作符 `.` 优先级相同。同级运算符，按照从左到右的顺序依次计算。 ` new Foo()` 先初始化 Foo 的实例化对象，实例上没有getName方法，因此需要原型上去找，即找到了 ` Foo.prototype.getName` ，输出3

8. ` new new Foo().getName()` ; new 带参数列表，优先级19，因此相当于是 ` new (new Foo()).getName()` ；先初始化 Foo 的实例化对象，然后将其原型上的 getName 函数作为构造函数再次 new ，输出3

因此最终结果如下:

` Foo.getName(); //2 getName(); //4 Foo().getName(); //1 getName(); //1 new Foo.getName(); //2 new Foo().getName(); //3 new new Foo().getName(); //3 复制代码`

如果你有更好的答案或想法，欢迎在这题目对应的github下留言： [下面这段代码的输出是什么？]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2FYvetteLau%2FBlog%2Fissues%2F24 )

### 16.实现双向绑定 Proxy 与 Object.defineProperty 相比优劣如何? ###

* 

Object.definedProperty 的作用是劫持一个对象的属性，劫持属性的getter和setter方法，在对象的属性发生变化时进行特定的操作。而 Proxy 劫持的是整个对象。

* 

Proxy 会返回一个代理对象，我们只需要操作新对象即可，而 ` Object.defineProperty` 只能遍历对象属性直接修改。

* 

Object.definedProperty 不支持数组，更准确的说是不支持数组的各种API，因为如果仅仅考虑arry[i] = value 这种情况，是可以劫持的，但是这种劫持意义不大。而 Proxy 可以支持数组的各种API。

* 

尽管 Object.defineProperty 有诸多缺陷，但是其兼容性要好于 Proxy.

PS: Vue2.x 使用 Object.defineProperty 实现数据双向绑定，V3.0 则使用了 Proxy.

` //拦截器 let obj = {}; let temp = 'Yvette' ; Object.defineProperty(obj, 'name' , { get() { console.log( "读取成功" ); return temp }, set(value) { console.log( "设置成功" ); temp = value; } }); obj.name = 'Chris' ; console.log(obj.name); 复制代码`

**PS:** Object.defineProperty 定义出来的属性，默认是不可枚举，不可更改，不可配置【无法delete】

我们可以看到 Proxy 会劫持整个对象，读取对象中的属性或者是修改属性值，那么就会被劫持。但是有点需要注意，复杂数据类型，监控的是引用地址，而不是值，如果引用地址没有改变，那么不会触发set。

` let obj = { name : 'Yvette' , hobbits : [ 'travel' , 'reading' ], info : { age : 20 , job : 'engineer' }}; let p = new Proxy (obj, { get(target, key) { //第三个参数是 proxy， 一般不使用 console.log( '读取成功' ); return Reflect.get(target, key); }, set(target, key, value) { if (key === 'length' ) return true ; //如果是数组长度的变化，返回。 console.log( '设置成功' ); return Reflect.set([target, key, value]); } }); p.name = 20 ; //设置成功 p.age = 20 ; //设置成功; 不需要事先定义此属性 p.hobbits.push( 'photography' ); //读取成功;注意不会触发设置成功 p.info.age = 18 ; //读取成功;不会触发设置成功 复制代码`

最后，我们再看下对于数组的劫持，Object.definedProperty 和 Proxy 的差别

Object.definedProperty 可以将数组的索引作为属性进行劫持，但是仅支持直接对 arry[i] 进行操作，不支持数组的API，非常鸡肋。

` let arry = [] Object.defineProperty(arry, '0' , { get() { console.log( "读取成功" ); return temp }, set(value) { console.log( "设置成功" ); temp = value; } }); arry[ 0 ] = 10 ; //触发设置成功 arry.push( 10 ); //不能被劫持 复制代码`

Proxy 可以监听到数组的变化，支持各种API。注意数组的变化触发get和set可能不止一次，如有需要，自行根据key值决定是否要进行处理。

` let hobbits = [ 'travel' , 'reading' ]; let p = new Proxy (hobbits, { get(target, key) { // if(key === 'length') return true; //如果是数组长度的变化，返回。 console.log( '读取成功' ); return Reflect.get(target, key); }, set(target, key, value) { // if(key === 'length') return true; //如果是数组长度的变化，返回。 console.log( '设置成功' ); return Reflect.set([target, key, value]); } }); p.splice( 0 , 1 ) //触发get和set，可以被劫持 p.push( 'photography' ); //触发get和set p.slice( 1 ); //触发get；因为 slice 是不会修改原数组的 复制代码`

如果你有更好的答案或想法，欢迎在这题目对应的github下留言： [实现双向绑定 Proxy 与 Object.defineProperty 相比优劣如何?]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2FYvetteLau%2FBlog%2Fissues%2F25 )

### 17. ` Object.is()` 与比较操作符 ` ===` 、 ` ==` 有什么区别？ ###

以下情况，Object.is认为是相等

` 两个值都是 undefined 两个值都是 null 两个值都是 true 或者都是 false 两个值是由相同个数的字符按照相同的顺序组成的字符串 两个值指向同一个对象 两个值都是数字并且 都是正零 +0 都是负零 -0 都是 NaN 都是除零和 NaN 外的其它同一个数字 复制代码`

Object.is() 类似于 ===，但是有一些细微差别，如下：

* NaN 和 NaN 相等
* -0 和 +0 不相等
` console.log( Object.is( NaN , NaN )); //true console.log( NaN === NaN ); //false console.log( Object.is( -0 , + 0 )); //false console.log( -0 === + 0 ); //true 复制代码`

Object.is 和 ` ==` 差得远了， ` ==` 在类型不同时，需要进行类型转换，前文已经详细说明。

如果你有更好的答案或想法，欢迎在这题目对应的github下留言： Object.is() 与比较操作符 ` ===` 、 ` ==` 有什么区别? ( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2FYvetteLau%2FBlog%2Fissues%2F26 )

### 18.什么是事件循环？Node事件循环和JS事件循环的差异是什么？ ###

最后一道题留给大家回答，再写下去，篇幅实在太长。

针对这道题，后面会专门写一篇文章~

**留下你的答案:** [什么是事件循环？Node事件循环和JS事件循环的差异是什么？]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2FYvetteLau%2FBlog%2Fissues%2F27 )

关于浏览器的event-loop可以看我之前的文章： [搞懂浏览器的EventLoop]( https://juejin.im/post/5c947bca5188257de704121d )

![](https://user-gold-cdn.xitu.io/2019/4/22/16a42bde889dd1a0?imageView2/0/w/1280/h/960/ignore-error/1)

> 
> 
> 
> 参考文章:
> 
> 

* [www.imooc.com/article/386…]( https://link.juejin.im?target=https%3A%2F%2Fwww.imooc.com%2Farticle%2F38600 )
* [es6.ruanyifeng.com/]( https://link.juejin.im?target=http%3A%2F%2Fes6.ruanyifeng.com%2F )
* [www.imooc.com/article/725…]( https://link.juejin.im?target=https%3A%2F%2Fwww.imooc.com%2Farticle%2F72500 )
* [www.cnblogs.com/LuckyWinty/…]( https://link.juejin.im?target=https%3A%2F%2Fwww.cnblogs.com%2FLuckyWinty%2Fp%2F5796190.html )
* [www.jianshu.com/p/a76dc7e0c…]( https://link.juejin.im?target=https%3A%2F%2Fwww.jianshu.com%2Fp%2Fa76dc7e0c5a1 )
* [www.v2ex.com/t/351261]( https://link.juejin.im?target=https%3A%2F%2Fwww.v2ex.com%2Ft%2F351261 )

> 
> 
> 
> ### 后续写作计划(写作顺序不定) ###
> 
> 

1.《寒冬求职季之你必须要懂的原生JS》(下)

2.《寒冬求职季之你必须要知道的CSS》

3.《寒冬求职季之你必须要懂的前端安全》

4.《寒冬求职季之你必须要懂的一些浏览器知识》

5.《寒冬求职季之你必须要知道的性能优化》

6.《寒冬求职季之你必须要懂的webpack原理》

**针对React技术栈:**

1.《寒冬求职季之你必须要懂的React》系列

2.《寒冬求职季之你必须要懂的ReactNative》系列

![](https://user-gold-cdn.xitu.io/2019/4/22/16a42bde9dcc8a73?imageView2/0/w/1280/h/960/ignore-error/1)

本文的写成耗费了非常多的时间，在这个过程中，我也学习到了很多知识，谢谢各位小伙伴愿意花费宝贵的时间阅读本文，如果本文给了您一点帮助或者是启发，请不要吝啬你的赞和Star，您的肯定是我前进的最大动力。 [github.com/YvetteLau/B…]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2FYvetteLau%2FBlog )

> 
> 
> 
> 关注小姐姐的公众号，和小姐姐一起学前端。
> 
> 

![](https://user-gold-cdn.xitu.io/2019/5/14/16ab621cfa97956f?imageView2/0/w/1280/h/960/ignore-error/1)