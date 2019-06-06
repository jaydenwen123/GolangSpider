# 你真的会用 Promise 吗？ #

![](https://user-gold-cdn.xitu.io/2019/6/5/16b26e1a98d4b04f?imageView2/0/w/1280/h/960/ignore-error/1)

# 前言：回调地狱 #

试想一下，有 3 个异步请求，第二个需要依赖第一个请求的返回结果，第三个需要依赖第二个请求的返回结果，一般怎么做？

` try { // 请求1 $.ajax({ url : url1, success : function ( data1 ) { // 请求2 try { $.ajax({ url : url1, data : data1, success : function ( data2 ) { try { // 请求3 $.ajax({ url : url1, data : data2, success : function ( data3 ) { // 后续业务逻辑... } }); } catch (ex3){ // 请求3的异常处理 } } }) } catch (ex){ // 请求2的异常处理 } } }) } catch (ex1){ // 请求1的异常处理 } 复制代码`

显然，如果再加上复杂的业务逻辑、异常处理，代码会更臃肿。在一个团队中，对这种代码的 review 和维护将会很痛苦。

回调地狱带来的负面作用有以下几点：

* 代码臃肿。
* 可读性、可维护性差。
* 耦合度高、可复用性差。
* 容易滋生 bug。
* 异常处理很恶心，只能在回调里处理异。

![](https://user-gold-cdn.xitu.io/2019/6/5/16b26e35a8f8ce35?imageView2/0/w/1280/h/960/ignore-error/1)

# Promise 全解 #

## 什么是 Promise？ ##

* Promise 是一种异步编程解决方案，避免回调地狱，可以把异步代码写得像同步一样。
* Promise 是一个对象，用于表示一个异步操作的最终状态（完成或失败），以及该异步操作的结果值。
* Promise 是一个代理（代理一个值），被代理的值在Promise对象创建时可能是未知的。它允许你为异步操作的成功和失败分别绑定相应的处理方法（handlers）。 这让异步方法可以像同步方法那样返回值，但并不是立即返回最终执行结果，而是一个能代表未来出现的结果的 promise 对象。

` var promise1 = new Promise ( function ( resolve, reject ) { setTimeout( function ( ) { resolve( 'foo' ); }, 300 ); }); promise1.then( function ( value ) { console.log(value); // after 300ms, expected output: "foo" }); 复制代码`

## Promise 核心特性？ ##

* 

一个 Promise 有 3 种状态：

* pending: 初始状态，既不是成功，也不是失败状态。
* fulfilled: 意味着操作成功完成。
* rejected: 意味着操作失败。

pending 状态的 Promise 可能会变为fulfilled 状态，也可能变为 rejected 状态。

![](https://user-gold-cdn.xitu.io/2019/6/5/16b26e54eeb0ba5a?imageView2/0/w/1280/h/960/ignore-error/1)

* 

Promise 对象的状态，只有内部能够改变（而且只能改变一次），不受外界影响。

* 

对象的状态一旦改变，就不会再变，任何时候都可以得到这个结果。 Promise 对象的状态改变，只有两种可能：从 Pending 变为 Resolved 和从 Pending 变为 Rejected。一旦状态发生改变，状态就凝固了，会一直保持这个结果。

![](https://user-gold-cdn.xitu.io/2019/6/5/16b26e679fc3beae?imageView2/0/w/1280/h/960/ignore-error/1)

` const p = new Promise ( ( resolve, reject )=> { resolve( "resolved first time!" ); // 只有第一次有效 resolve( "resolved second time!" ); reject( "rejected!" ); }); p.then( ( data )=> console.log(data), (error)=> console.log(error) ); 复制代码`

![](https://user-gold-cdn.xitu.io/2019/6/5/16b26e6da6433838?imageView2/0/w/1280/h/960/ignore-error/1)

## Promise API ##

` // 1. 构造方法 const p = new Promise ( ( resolve, reject ) => { /* executor*/ // 1.1. Promise构造函数执行时立即调用 executor 函数; // 1.2. resolve 和 reject 函数被调用时，分别将promise的状态改为fulfilled（完成）或rejected（失败） // 1.3. 如果在executor函数中抛出一个错误，那么该promise 状态为rejected。 // 1.4. executor函数的返回值被忽略。 }); // 2.原型方法 Promise.prototype.catch(onRejected) Promise.prototype.then(onFulfilled, onRejected) // 3.静态方法 Promise.all(iterable); Promise.race(iterable); Promise.reject(reason); Promise.resolve(value); 复制代码`

## 示例：用 Promise 和 XMLHttpRequest 加载图像 ##

` function imgLoad ( url ) { // Create new promise with the Promise() constructor; // This has as its argument a function // with two parameters, resolve and reject return new Promise ( function ( resolve, reject ) { // Standard XHR to load an image var request = new XMLHttpRequest(); request.open( 'GET' , url); request.responseType = 'blob' ; // When the request loads, check whether it was successful request.onload = function ( ) { if (request.status === 200 ) { // If successful, resolve the promise by passing back the request response resolve(request.response); } else { // If it fails, reject the promise with a error message reject( Error ( 'Image didn\'t load successfully; error code:' + request.statusText)); } }; request.onerror = function ( ) { // Also deal with the case when the entire request fails to begin with // This is probably a network error, so reject the promise with an appropriate message reject( Error ( 'There was a network error.' )); }; // Send the request request.send(); }); } // Get a reference to the body element, and create a new image object var body = document.querySelector( 'body' ); var myImage = new Image(); // Call the function with the URL we want to load, but then chain the // promise then() method on to the end of it. This contains two callbacks imgLoad( 'myLittleVader.jpg' ).then( function ( response ) { // The first runs when the promise resolves, with the request.response // specified within the resolve() method. var imageURL = window.URL.createObjectURL(response); myImage.src = imageURL; body.appendChild(myImage); // The second runs when the promise // is rejected, and logs the Error specified with the reject() method. }, function ( Error ) { console.log( Error ); }); 复制代码`

## Promise 与事件循环机制 ##

Event Loop 中的事件，分为 MacroTask（宏任务）和 MicroTask（微任务）。

* MacroTask: setTimeout, setInterval, setImmediate, requestAnimationFrame, I/O, UI rendering
* MicroTask: process.nextTick, Promises, Object.observe, MutationObserver

通俗来说，MacroTasks 和 MicroTasks 最大的区别在它们会被放置在不同的任务调度队列中。

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2a378f87ecb97?imageView2/0/w/1280/h/960/ignore-error/1) 每一次事件循环中，主进程都会先执行一个MacroTask 任务，这个任务就来自于所谓的MacroTask Queue队列；当该 MacroTask 执行完后，Event loop 会立马调用 MicroTask 队列的任务，直到消费完所有的 MicroTask，再继续下一个事件循环。

> 
> 
> 
> 注：async/await 本质上还是基于Promise的一些封装，而Promise是属于微任务的一种。所以在使用 await 关键字与
> Promise.then 效果类似。即：async 函数在 await 之前的代码都是同步执行的，可以理解为await之前的代码属于new
> Promise时传入的代码，await之后的所有代码都是在Promise.then中的回调；
> 
> 

## Promise 常见面试题目 ##

### 题目：写出运行结果 ###

` setTimeout( function ( ) { console.log( 1 ); }, 0 ) new Promise ( function ( resolve ) { console.log( 2 ); resolve(); console.log( 3 ); }).then( function ( ) { console.log( 4 ); }) console.log( 5 ); 复制代码`

答案 & 解析：

` // 解析： // 1. new Promise(fn)后，函数fn会立即执行； // 2. fn在执行过程中，由于调用了resolve，使得Promise立即转换为resolve状态， // 这也促使p.then(fn)中的函数fn被立即放入microTask队列中，因此fn将会在 // 本轮事件循环的结束时执行，而不是下一轮事件循环的开始时执行； // 3. setTimeout属于macroTask，是在下一轮事件循环中执行； //答案： // 2 3 5 4 1 复制代码`

### 题目：写出运行结果 ###

` Promise.resolve( 1 ) .then( ( res ) => { console.log(res); return 2 ; }) .catch( ( res ) => { console.log(res); return 3 ; }) .then( ( res ) => { console.log(res); }); 复制代码`

答案 & 解析：

` // 解析：每次调用p.then或p.catch都会返回一个新的promise， // 从而实现了链式调用；第一个.then中未抛出异常， // 所以不会被.catch语句捕获，会正常进入第二个.then执行； // 答案：1 2 复制代码`

### 题目：写出运行结果 ###

` Promise.resolve() .then( () => { return new Error ( 'error!' ) }) .then( res => { console.log( 'then: ' , res) }) .catch( err => { console.log( 'catch: ' , err) }); 复制代码`

答案 & 解析：

` // 解析：在 .then 或 .catch 中 return 一个 error 对象并不会抛出错误， // 所以不会被后续的 .catch 捕获； // 答案： then : Error: error! // at ... // at ... 复制代码`

### 题目：写出运行结果 ###

` Promise.resolve( 1 ) .then( 2 ) .then( Promise.resolve( 3 )) .then( console.log); 复制代码`

答案 & 解析：

` // 解析：p.then、.catch 的入参应该是函数，传入非函数则会发生值穿透； // 答案：1 复制代码`

### 题目：写出运行结果 ###

` Promise.resolve() .then( value => { throw new Error ( 'error' ); }, reason => { console.error( 'fail1:' , reason); } ) .catch( reason => { console.error( 'fail2:' , reason); } ); 复制代码`

答案 & 解析：

` // 解析：.then可以接收两个参数：.then(onResolved, onRejected) // .catch是.then的语法糖：.then(onRejected) ==> .then(null, onRejected) // 答案：fail2: Error: error // at ..... // at ..... 复制代码`

### 题目：写出运行结果 ###

` console.log( 1 ); new Promise ( function ( resolve, reject ) { reject(); resolve(); }).then( function ( ) { console.log( 2 ); }, function ( ) { console.log( 3 ); }); console.log( 4 ); 复制代码`

答案 & 解析：

` // 解析：Promise状态的一旦变成resolved或rejected， // Promise的状态和值就固定下来了， // 不论你后续再怎么调用resolve或reject方法， // 都不能改变它的状态和值。 // // 答案：1 4 3 复制代码`

### 题目：写出运行结果 ###

` new Promise ( resolve => { // p1 resolve( 1 ); // p2 Promise.resolve().then( () => { console.log( 2 ); // t1 }); console.log( 4 ) }).then( t => { console.log(t); // t2 }); console.log( 3 ); 复制代码`

答案 & 解析：

` // 解析： // 1. new Promise(fn), fn 立即执行，所以先输出 4； // 2. p1和p2的Promise在执行 then 之前都已处于resolve状态， // 故按照 then 执行的先后顺序，将t1、t2放入microTask中等待执行； // 3. 完成执行console.log(3)后，macroTask执行结束，然后microTask // 中的任务t1、t2依次执行，所以输出3、2、1； // 答案： // 4 3 2 1 复制代码`

### 题目：写出运行结果 ###

` Promise.reject( 'a' ) .then( () => { console.log( 'a passed' ); }) .catch( () => { console.log( 'a failed' ); }); Promise.reject( 'b' ) .catch( () => { console.log( 'b failed' ); }) .then( () => { console.log( 'b passed' ); }) 复制代码`

答案 & 解析：

` // 解析：p.then(fn)、p.catch(fn)中的fn都是异步执行，上述代码可理解为： // set Timeout( function (){ // set Timeout( function (){ // console.log( 'a failed' ); // }); // }); // set Timeout( function (){ // console.log( 'b failed' ); // // set Timeout( function (){ // console.log( 'b passed' ); // }); // }); // 答案：b failed // a failed // b passed 复制代码`

### 题目：写出运行结果 ###

` async function async1 ( ) { console.log( 'async1 start' ) await async2() console.log( 'async1 end' ) } async function async2 ( ) { console.log( 'async2' ) } console.log( 'script start' ); setTimeout( function ( ) { console.log( 'settimeout' ) }) async1(); new Promise ( function ( resolve ) { console.log( 'promise1' ); resolve(); }).then( function ( ) { console.log( 'promise2' ); }) console.log( 'script end' ); 复制代码`

答案：（不解析了，大家研究一下）

` script start async1 start async2 promise1 script end promise2 async1 end settimeout 复制代码`

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2a3ee7719e4d2?imageView2/0/w/1280/h/960/ignore-error/1)

# 自己实现一版 Promise #

Promise有很多社区规范，如 Promise/A、Promise/B、Promise/D 以及 Promise/A 的升级版 Promise/A+；Promise/A+ 是 ES6 Promises 的前身，而且网络上有很多可供学习、参考的开源实现（例如：Adehun、bluebird、Q、ypromise等）。

## Promise 的规范去哪找？ ##

> 
> 
> 
> Promise/A+ 规范：
> [github.com/promises-ap…](
> https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fpromises-aplus%2Fpromises-spec
> )
> 
> 

## 如何保证自己实现的 Promise 符合规范？ ##

用官方的Promise规范测试集，测试自己的实现。

> 
> 
> 
> Promise/A+ 规范测试集：
> [github.com/promises-ap…](
> https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fpromises-aplus%2Fpromises-tests
> )
> 
> 

## 开始编码 ##

### 识别核心接口 ###

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2a415689fba24?imageView2/0/w/1280/h/960/ignore-error/1)

` 可以看出，共需实现7个接口； 复制代码`

### 分析接口间联系 ###

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2a427c248455b?imageView2/0/w/1280/h/960/ignore-error/1)

` 可以看出，7个接口中，只有构造函数RookiePromise和成员函数then算核心接口，其他接口均可通过这两个接口实现； 复制代码`

### 仔细阅读官方规范，逐条合规编码 ###

#### 构建主框架 ####

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2a42f90afad9e?imageView2/0/w/1280/h/960/ignore-error/1)

#### 编写状态转换逻辑 ####

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2a433853c9b0f?imageView2/0/w/1280/h/960/ignore-error/1)

> 
> 
> 
> Promise 对象的状态改变，只有两种可能：pending -> fulfilled 和 pending ->
> rejected。只要这两种情况发生，状态就凝固了，不会再变了，会一直保持这个结果；
> ——《ES6 标准入门（第三版）》
> 
> 

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2a4424d2d8099?imageView2/0/w/1280/h/960/ignore-error/1) 注：_notify函数用作异步执行传入的函数数组以及参数；代码中将_callbacks、_errbacks传给_notify函数后立即清空，是为了保证_callbacks、_errbacks至多被执行一次；

#### 实现 then 接口 ####

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2a44b47fea215?imageView2/0/w/1280/h/960/ignore-error/1)

#### 实现resolve(promise, x)接口 ####

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2a451d553ec45?imageView2/0/w/1280/h/960/ignore-error/1)

#### 完整 RookiePromise 源码实现 ####

` /** * 2.1. Promise States * A promise must be in one of three states: * pending, fulfilled, or rejected. */ const STATE_PENDING = "pending" ; const STATE_FULFILLED = "fulfilled" ; const STATE_REJECTED = "rejected" ; function RookiePromise ( fn ) { this._state = STATE_PENDING; this._value = undefined ; this._callbacks = []; this._errorbacks = []; /** * 2.3. The Promise Resolution Procedure * The promise resolution procedure is an abstract operation * taking as input a promise and a value, which we denote as * [[Resolve]](promise, x) */ var executed = false ; // 用于保证resolve接口只有第一次被触发时有效； function resolve ( promise, x ) { if (executed){ return ; } executed = true ; var innerResolve = ( promise, x ) => { if (promise === x){ // 2.3.1. If promise and x refer to the same object, // reject promise with a TypeError as the reason. this._reject( new TypeError ( "出错了, promise === x, 会造成死循环!" )); } else if (x instanceof RookiePromise){ // 2.3.2. If x is a promise, adopt its state [3.4]: // 2.3.2.1. If x is pending, promise must remain pending until x is fulfilled or rejected. // 2.3.2.2. If/when x is fulfilled, fulfill promise with the same value. // 2.3.2.3. If/when x is rejected, reject promise with the same reason. if (x._state == STATE_PENDING){ x.then( ( value ) => { innerResolve(promise, value); }, (reason) => { this._reject(reason); }); } else if (x._state == STATE_FULFILLED){ this._fulfill(x._value); } else if (x._state == STATE_REJECTED){ this._reject(x._value); } } else if (x && ( typeof x == "function" || typeof x == "object" )){ // 2.3.3. Otherwise, if x is an object or function, try { // 2.3.3.1. Let then be x.then. let then = x.then; if ( typeof then === "function" ){ //thenable var executed = false ; try { // 2.3.3.3. If then is a function, call it with x as this, // first argument resolvePromise, and // second argument rejectPromise, // where: then.call(x, (value) => { // 2.3.3.3.3. If both resolvePromise and rejectPromise are called, // or multiple calls to the same argument are made, // the first call takes precedence, and any further calls are ignored. if (executed){ return ; } executed = true ; // 2.3.3.3.1. If/when resolvePromise is called with a value y, // run [[Resolve]](promise, y). innerResolve(promise, value); }, (reason) => { // 2.3.3.3.3. If both resolvePromise and rejectPromise are called, // or multiple calls to the same argument are made, // the first call takes precedence, and any further calls are ignored. if (executed){ return ; } executed = true ; // 2.3.3.3.2. If/when rejectPromise is called with a reason r, // reject promise with r. this._reject(reason); }); } catch (e){ // 2.3.3.3.4. If calling then throws an exception e, // 2.3.3.3.4.1. If resolvePromise or rejectPromise have been called, ignore it. if (executed){ return ; } // 2.3.3.3.4.2. Otherwise, reject promise with e as the reason. throw e; } } else { // 2.3.3.4. If then is not a function, fulfill promise with x. this._fulfill(x); } } catch (ex){ // 2.3.3.2. If retrieving the property x.then results in a thrown exception e, // reject promise with e as the reason. this._reject(ex); } } else { // 2.3.4. If x is not an object or function, fulfill promise with x. this._fulfill(x); } }; innerResolve(promise, x) } function reject ( promise, reason ) { this._reject(reason); } resolve = resolve.bind( this , this ); // 通过bind模拟规范中的 [[Resolve]](promise, x) 行为 reject = reject.bind( this , this ); fn(resolve, reject); // new RookiePromise((resolve, reject) => { ... }) } /** * 2.1. Promise States * * A promise must be in one of three states: pending, fulfilled, or rejected. * * 2.1.1. When pending, a promise: * 2.1.1.1 may transition to either the fulfilled or rejected state. * 2.1.2. When fulfilled, a promise: * 2.1.2.1 must not transition to any other state. * 2.1.2.2 must have a value, which must not change. * 2.1.3. When rejected, a promise: * 2.1.3.1 must not transition to any other state. * 2.1.3.2 must have a reason, which must not change. * * Here, “must not change” means immutable identity (i.e. ===), * but does not imply deep immutability. */ RookiePromise.prototype._fulfill = function ( value ) { if ( this._state == STATE_PENDING){ this._state = STATE_FULFILLED; this._value = value; this._notify( this._callbacks, this._value); this._errorbacks = []; this._callbacks = []; } } RookiePromise.prototype._reject = function ( reason ) { if ( this._state == STATE_PENDING){ this._state = STATE_REJECTED; this._value = reason; this._notify( this._errorbacks, this._value); this._errorbacks = []; this._callbacks = []; } } RookiePromise.prototype._notify = function ( fns, param ) { setTimeout( () => { for ( var i= 0 ; i<fns.length; i++){ fns[i](param); } }, 0 ); } /** * 2.2. The then Method * A promise’s then method accepts two arguments: * promise.then(onFulfilled, onRejected) */ RookiePromise.prototype.then = function ( onFulFilled, onRejected ) { // 2.2.7. then must return a promise [3.3]. // promise2 = promise1.then(onFulFilled, onRejected); // return new RookiePromise( ( resolve, reject )=> { // 2.2.1. Both onFulfilled and onRejected are optional arguments: // 2.2.1.1. If onFulfilled is not a function, it must be ignored. // 2.2.1.2. If onRejected is not a function, it must be ignored. if ( typeof onFulFilled == "function" ){ this._callbacks.push( function ( value ) { try { // 2.2.5. onFulfilled and onRejected must be called as functions (i.e. with no this value) var value = onFulFilled(value); resolve(value); } catch (ex){ // 2.2.7.2. If either onFulfilled or onRejected throws an exception e, // promise2 must be rejected with e as the reason. reject(ex); } }); } else { // 2.2.7.3. If onFulfilled is not a function and promise1 is fulfilled, // promise2 must be fulfilled with the same value as promise1. this._callbacks.push(resolve); // 值穿透 } if ( typeof onRejected == "function" ){ this._errorbacks.push( function ( reason ) { try { // 2.2.5. onFulfilled and onRejected must be called as functions (i.e. with no this value) var value = onRejected(reason); resolve(value); } catch (ex){ // 2.2.7.2. If either onFulfilled or onRejected throws an exception e, // promise2 must be rejected with e as the reason. reject(ex); } }); } else { // 2.2.7.4. If onRejected is not a function and promise1 is rejected, // promise2 must be rejected with the same reason as promise1. this._errorbacks.push(reject); // 值穿透 } // 2.2.6. then may be called multiple times on the same promise. // 2.2.6.1. If/when promise is fulfilled, all respective onFulfilled callbacks must // execute in the order of their originating calls to then. // 2.2.6.2. If/when promise is rejected, all respective onRejected callbacks must // execute in the order of their originating calls to then. if ( this._state == STATE_REJECTED){ // 2.2.4. onFulfilled or onRejected must not be called until the // execution context stack contains only platform code. this._notify( this._errorbacks, this._value); this._errorbacks = []; this._callbacks = []; } else if ( this._state == STATE_FULFILLED){ // 2.2.4. onFulfilled or onRejected must not be called until the // execution context stack contains only platform code. this._notify( this._callbacks, this._value); this._errorbacks = []; this._callbacks = []; } }); }; RookiePromise.prototype.catch = function ( onRejected ) { return this.then( null , onRejected); }; RookiePromise.resolve = function ( value ) { return new RookiePromise( ( resolve, reject ) => resolve(value)); }; RookiePromise.reject = function ( reason ) { return new RookiePromise( ( resolve, reject ) => reject(reason)); }; RookiePromise.all = function ( values ) { return new Promise ( ( resolve, reject ) => { var result = [], remaining = values.length; function resolveOne ( index ) { return function ( value ) { result[index] = value; remaining--; if (!remaining){ resolve(result); } }; } for ( var i = 0 ; i < values.length; i++) { RookiePromise.resolve(values[i]).then(resolveOne(i), reject); } }); }; RookiePromise.race = function ( values ) { return new Promise ( ( resolve, reject ) => { for ( var i = 0 ; i < values.length; i++) { RookiePromise.resolve(values[i]).then(resolve, reject); } }); }; module.exports = RookiePromise; 复制代码`

### RookiePromise 编码小结 ###

RookiePromise的结构是按照Promise/A+规范中对then、resolve接口的描述组织的；优点是编码过程直观，缺点是innerResolve函数篇幅太长、头重脚轻，不够和谐；相信各位可以写出更漂亮的版本；

## 测试正确性 ##

### 安装 Promise/A+测试工具 ###

> 
> 
> 
> npm install –save promises-aplus-tests
> 
> 

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2a46e5a7f4f68?imageView2/0/w/1280/h/960/ignore-error/1)

### 编写 RookiePromise 的测试适配器 ###

RookiePromise需要额外提供3个静态接口，供Promise/A+自动测试工具调用；

` /** * In order to test your promise library, * you must expose a very minimal adapter interface. * These are written as Node.js modules with a few well-known exports: * * resolved(value): creates a promise that is resolved with value. * rejected(reason): creates a promise that is already rejected with reason. * deferred(): creates an object consisting of { promise, resolve, reject }: * promise is a promise that is currently in the pending state. * resolve(value) resolves the promise with value. * reject(reason) moves the promise from the pending state to the rejected state, * with rejection reason reason. * * https://github.com/promises-aplus/promises-tests */ var RookiePromise = require ( './RookiePromise.js' ); RookiePromise.resolved = RookiePromise.resolve; RookiePromise.rejected = RookiePromise.reject; RookiePromise.deferred = function ( ) { let defer = {}; defer.promise = new RookiePromise( ( resolve, reject ) => { defer.resolve = resolve; defer.reject = reject; }); return defer; } module.exports = RookiePromise 复制代码`

### 执行测试 ###

> 
> 
> 
> npx promises-aplus-testsRookiePromiseTestAdapter.js > log.txt
> 
> 

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2a47b1908df5e?imageView2/0/w/1280/h/960/ignore-error/1) 完美通过测试，RookiePromise 是符合 Promise/A+规范的！！！

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2a47ebeb9ae95?imageView2/0/w/1280/h/960/ignore-error/1) 参考：

> 
> 
> 
> 《ES6 标准入门（第三版）》
> 《深入理解ES6》
> MDN（Promise）：
> [developer.mozilla.org/en-US/docs/…](
> https://link.juejin.im?target=https%3A%2F%2Fdeveloper.mozilla.org%2Fen-US%2Fdocs%2FWeb%2FJavaScript%2FReference%2FGlobal_Objects%2FPromise
> )
> Promise 示例（Promise 和 XMLHttpRequest 加载图像）：
> [github.com/mdn/js-exam…](
> https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fmdn%2Fjs-examples%2Ftree%2Fmaster%2Fpromises-test
> )
> States and Fates：
> [github.com/domenic/pro…](
> https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fdomenic%2Fpromises-unwrapping%2Fblob%2Fmaster%2Fdocs%2Fstates-and-fates.md
> )
> Promise/A+规范文档：
> [github.com/promises-ap…](
> https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fpromises-aplus%2Fpromises-spec
> )
> Promise/A+规范测试集：
> [github.com/promises-ap…](
> https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fpromises-aplus%2Fpromises-tests
> )
> 符合Promise/A+规范的一些开源实现：
> [github.com/promises-ap…](
> https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fpromises-aplus%2Fpromises-spec%2Fblob%2Fmaster%2Fimplementations.md
> )
> 
> 

![](https://user-gold-cdn.xitu.io/2019/5/27/16af831597dc6484?imageView2/0/w/1280/h/960/ignore-error/1)

> 
> 
> 
> 社区以及公众号发布的文章，100%保证是我们的原创文章，如果有错误，欢迎大家指正。
> 
> 

> 
> 
> 
> 文章首发在WebJ2EE公众号上，欢迎大家关注一波，让我们大家一起学前端~~~
> 
> 

![](https://user-gold-cdn.xitu.io/2019/5/22/16aded0040c28b43?imageView2/0/w/1280/h/960/ignore-error/1)

> 
> 
> 
> 再来一波号外，我们成立WebJ2EE公众号前端吹水群，大家不管是看文章还是在工作中前端方面有任何问题，我们都可以在群内互相探讨，希望能够用我们的经验帮更多的小伙伴解决工作和学习上的困惑,欢迎加入。
> 
> 
> 

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2a4c8b373bf4e?imageView2/0/w/1280/h/960/ignore-error/1)