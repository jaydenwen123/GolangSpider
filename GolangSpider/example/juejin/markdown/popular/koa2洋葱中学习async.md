# koa2洋葱中学习async #

因为记性差所以记录下来，看源码更多的时候看一种思想，koa2使用了很多es6的语法，所以就以看源码来学习语法吧。 看完koa的源码，最核心的一块就是洋葱式调用。所以我就单独吧这块的代码简化下来，以便自己以后突然想起来再看的时候不需要重新去捋思路了。

* 首先创建一个存储函数的数组middleware
* app.use将中间件存入middleware数组，数组准备好以后启动服务
* 启动服务以后调用回调，将middleware数组里中间件递归执行一次
* 执行结束后返回Promise
* 我就以这个顺序把函数独立出来，并且删了一些原作者验证的代码和console代码， 这里假设每个方法都是属于app这个实例的，就不以app.来写了。独立出来

### middleware是在实例上创建了一个空数组， ###

` use(fn) { if (typeof fn !== 'function' ) throw new TypeError( 'middleware must be a function!' ); if (isGeneratorFunction(fn)) { fn = convert(fn); //转换为async 函数 } this.middleware.push(fn); return this; //返回新实例 } 复制代码`

### use方法就是将所有中间件放入middleware里。 ###

` app.listen(3000, () => { }) 复制代码`

### 数组准备好以后启动服务 ###

` listen(...args) { const server = http.createServer(this.callback()); return server.listen(...args); } 复制代码`

### 创建node服务器调用回调 ###

` callback () { const fn = compose(this.middleware); if (!this.listeners( 'error' ).length) this.on( 'error' , this.onerror); const handleRequest = (req, res) => { const ctx = this.createContext(req, res); return this.handleRequest(ctx, fn); }; return handleRequest; } 复制代码`

### callback在回调里调用compose方法 ###

` function compose (middleware) { //middleware 回调数组 return function (context, next) { // last called middleware # let index = -1; return dispatch(0) } //返回一个prm状态 } 复制代码`

### 返回一个Promise对象 告诉递归调用完成的结果 ###

` function dispatch (i) { if (i <= index) return Promise.reject(new Error( 'next() called multiple times' )) //判断如果i小于index直接返回错误 index = i; let fn = middleware[i]; //取出数组中对应下标函数 if (i === middleware.length) fn = next; if (!fn) return Promise.resolve(); try { return Promise.resolve(fn(context, function next () { return dispatch(i + 1) //递归调用 })) } catch (err) { return Promise.reject(err) } } 因为是async 所以当 await 接收到返回的Promise结果以后就会逐个执行下去 ， 也就是说当async函数被逐个执行完毕以后返回一个Promise对象，那么就会从async函数最后一个await逐个向上返回调用，直到所有await执行完毕。 这就是洋葱式调用 async function (){ await fun1(async function () { await fun2(async function () { await fun3(async function (){ return Promise.resolve(); }); }); }); }; 上一级的await一直在等待下一级的返回结果，所以逐级向下执行，在等到执行了Promise.resolve();有了返回结果以后再逐级向上返回 // 复制代码`

### 递归调用的具体执行 ###

` handleRequest(ctx, fnMiddleware) { return fnMiddleware(ctx).then(handleResponse).catch(onerror); //fnMiddleware 就是递归执行完毕以后返回prm对象接收一个函数 } //callback里调永compose函数并将middleware传递过去 复制代码`

//最后一行其实就是返回的fn对象

### 总结 ###

koa的思想其实就是运用了es7的新特性async函数的新特性，逐个等待异步的结果，一旦下层的返回结果就会逐个告诉转达给上层，就形成了洋葱式的调用，所以要读懂源码必须了解async函数，所以在读码的时候也同时学习复习了async函数，还是很不错的。如果不懂这个函数也许会在递归调用那块搞晕，会不明白递归调用完了以后为什么会按顺序返回结果。我觉得搞懂这个就算是搞懂koa2的核心思想了，其他的都是一些封装函数，可以慢慢看

` app.use(async (ctx, next) => { console.log(1) //典型语法await在等待 await next(); console.log(2) }); app.use(async (ctx, next) => { console.log(3) //典型语法await在等待 await next(); console.log(5) }); 返回结果 是 1,3,5,2， 这就是洋葱 这里附上经典的图吧 复制代码`

![](https://user-gold-cdn.xitu.io/2019/6/5/16b2649dfb65d697?imageView2/0/w/1280/h/960/ignore-error/1) ![](https://user-gold-cdn.xitu.io/2019/6/5/16b264a141256db5?imageView2/0/w/1280/h/960/ignore-error/1)