# ES6新特征总结与介绍——异步编程 #

### 一、Generator ###

#### （一）基本概念 ####

语法上，Generator 函数是一个状态机，封装了多个内部状态。执行Generator函数会返回一个遍历器对象，也就是说，Generator函数除了状态机，还是一个遍历器对象生成函数。返回的遍历器对象，可以依次遍历 Generator 函数内部的每一个状态。

形式上，Generator 函数是一个普通函数，但是有两个特征。一是，function关键字与函数名之间有一个星号；二是，函数体内部使用yield表达式，定义不同的内部状态。

` function * helloWorldGenerator () { yield 'hello' ; yield 'world' ; return 'ending' ; } var hw = helloWorldGenerator(); 复制代码`

上面代码定义了一个 Generator 函数helloWorldGenerator，它内部有两个yield表达式（hello和world），即该函数有三个状态：hello，world 和 return 语句（结束执行）。

Generator 函数的调用方法与普通函数一样，也是在函数名后面加上一对圆括号。不同的是，调用Generator函数后，该函数并不执行，返回的也不是函数运行结果，而是一个指向内部状态的指针对象，也就是遍历器对象（Iterator Object）。

下一步，必须调用遍历器对象的next方法，使得指针移向下一个状态。也就是说，每次调用next方法，内部指针就从函数头部或上一次停下来的地方开始执行，直到遇到下一个yield表达式（或return语句）为止。换言之，Generator函数是分段执行的，yield表达式是暂停执行的标记，而next方法可以恢复执行。

#### （二）yield ####

由于 Generator 函数返回的遍历器对象，只有调用next方法才会遍历下一个内部状态，所以其实提供了一种可以暂停执行的函数。yield表达式就是暂停标志。

#### （三）Generator.prototype.next() ####

遍历器对象的next方法的运行逻辑如下。

* 遇到yield表达式，就暂停执行后面的操作，并将紧跟在yield后面的那个表达式的值，作为返回的对象的value属性值。
* 下一次调用next方法时，再继续往下执行，直到遇到下一个yield表达式。
* 如果没有再遇到新的yield表达式，就一直运行到函数结束，直到return语句为止，并将return语句后面的表达式的值，作为返回的对象的value属性值。
* 如果该函数没有return语句，则返回的对象的value属性值为undefined。

#### （四）Generator.prototype.throw() ####

Generator 函数返回的遍历器对象，都有一个throw方法，可以在函数体外抛出错误，然后在 Generator 函数体内捕获。

` var g = function * () { try { yield; } catch (e) { console.log( '内部捕获' , e); } }; var i = g(); i.next(); try { i.throw( 'a' ); i.throw( 'b' ); } catch (e) { console.log( '外部捕获' , e); } // 内部捕获 a // 外部捕获 b 复制代码`

上面代码中，遍历器对象i连续抛出两个错误。第一个错误被Generator函数体内的catch语句捕获。i第二次抛出错误，由于Generator函数内部的catch语句已经执行过了，不会再捕捉到这个错误了，所以这个错误就被抛出了 Generator 函数体，被函数体外的catch语句捕获。

#### （五）Generator.prototype.return() ####

Generator 函数返回的遍历器对象，还有一个return方法，可以返回给定的值，并且终结遍历 Generator 函数。

` function * gen () { yield 1; yield 2; yield 3; } var g = gen(); g.next() // { value: 1, done : false } g.return( 'foo' ) // { value: "foo" , done : true } g.next() // { value: undefined, done : true } 复制代码`

如果return方法调用时，不提供参数，则返回值的value属性为undefined。

如果 Generator 函数内部有try...finally代码块，且正在执行try代码块，那么return方法会推迟到finally代码块执行完再执行。

### 二、Promise ###

#### （一）Promise 状态 ####

Promise 异步操作有三种状态：pending（进行中）、fulfilled（已成功）和rejected（已失败）。除了异步操作的结果，任何其他操作都无法改变这个状态。

Promise 对象只有：从 pending 变为 fulfilled 和从 pending 变为 rejected 的状态改变。只要处于 fulfilled 和 rejected ，状态就不会再变了即 resolved（已定型）。

#### （二）Promise.prototype.then() ####

then 方法接收两个函数作为参数，then方法的第一个参数是resolved状态的回调函数，第二个参数（可选）是rejected状态的回调函数。两个函数只会有一个被调用。

` const promise = new Promise( function (resolve, reject) { if (/* 异步操作成功 */){ resolve(value) } else { reject(error) } }) promise.then( function (value) { // success }, function (error) { // failure }) 复制代码`

#### （三）Promise.prototype.catch() ####

Promise.prototype.catch方法是.then(null, rejection)或.then(undefined, rejection)的别名，用于指定发生错误时的回调函数。 一般来说，不要在then方法里面定义 Reject 状态的回调函数（即then的第二个参数），总是使用catch方法。

` // bad promise .then( function (data) { // success }, function (err) { // error }); // good promise .then( function (data) { //cb // success }) .catch( function (err) { // error }); 复制代码`

#### （四）Promise.prototype.finally() ####

finally方法用于指定不管 Promise 对象最后状态如何，都会执行的操作。

` promise .then(result => {···}) .catch(error => {···}) .finally(() => {···}) 复制代码`

上面代码中，不管promise最后的状态，在执行完then或catch指定的回调函数以后，都会执行finally方法指定的回调函数。

#### （五）Promise.all() ####

Promise.all方法用于将多个 Promise 实例，包装成一个新的 Promise 实例。

` const p = Promise.all([p1, p2, p3]) 复制代码`

上面代码中，Promise.all方法接受一个数组作为参数，p1、p2、p3都是 Promise，p的状态由p1、p2、p3决定，分成两种情况。

* 只有p1、p2、p3的状态都变成fulfilled，p的状态才会变成fulfilled，此时p1、p2、p3的返回值组成一个数组，传递给p的回调函数。
* 只要p1、p2、p3之中有一个被rejected，p的状态就变成rejected，此时第一个被reject的实例的返回值，会传递给p的回调函数。

#### （六）Promise.race() ####

Promise.race方法同样是将多个 Promise 实例，包装成一个新的 Promise 实例。

` const p = Promise.race([p1, p2, p3]) 复制代码`

上面代码中，只要p1、p2、p3之中有一个实例率先改变状态，p的状态就跟着改变。那个率先改变的Promise实例的返回值，就传递给p的回调函数。

#### （七）Promise.resolve() ####

有时需要将现有对象转为 Promise 对象，Promise.resolve方法就起到这个作用。

* 

参数是一个 Promise 实例

如果参数是 Promise 实例，那么Promise.resolve将不做任何修改、原封不动地返回这个实例。

* 

参数是一个thenable对象

Promise.resolve方法会将这个对象转为 Promise 对象，然后就立即执行thenable对象的then方法。

* 

参数不是具有then方法的对象，或根本就不是对象

如果参数是一个原始值，或者是一个不具有then方法的对象，则Promise.resolve方法返回一个新的 Promise 对象，状态为resolved。

* 

不带有任何参数

Promise.resolve()方法允许调用时不带参数，直接返回一个resolved状态的 Promise 对象。

#### （八）Promise.reject() ####

Promise.reject(reason)方法也会返回一个新的 Promise 实例，该实例的状态为rejected。

### 三、async ###

#### （一）基本用法 ####

async函数返回一个 Promise 对象，可以使用then方法添加回调函数。当函数执行的时候，一旦遇到await就会先返回，等到异步操作完成，再接着执行函数体内后面的语句。

` async function getStockPriceByName(name) { const symbol = await getStockSymbol(name) const stockPrice = await getStockPrice(symbol) return stockPrice } getStockPriceByName( 'goog' ).then( function (result) { console.log(result) }) 复制代码`

#### （二）返回 Promise 对象 ####

async函数返回一个 Promise 对象。 async函数内部return语句返回的值，会成为then方法回调函数的参数。

` async function f () { return 'hello world' } f().then(v => console.log(v)) // "hello world" 复制代码`

上面代码中，函数f内部return命令返回的值，会被then方法回调函数接收到。

async函数内部抛出错误，会导致返回的 Promise 对象变为reject状态。抛出的错误对象会被catch方法回调函数接收到。

` async function f () { throw new Error( '出错了' ) } f().then( v => console.log(v), e => console.log(e) ) // Error: 出错了 复制代码`

#### （三）await 命令 ####

* 正常情况下，await命令后面是一个 Promise 对象，返回该对象的结果。如果不是 Promise 对象，就直接返回对应的值。
` async function f () { // 等同于 // return 123; return await 123; } f().then(v => console.log(v)) // 123 复制代码` * 另一种情况是，await命令后面是一个thenable对象（即定义then方法的对象），那么await会将其等同于 Promise 对象。

#### （四）注意点 ####

* await命令后面的Promise对象，运行结果可能是rejected，所以最好把await命令放在try...catch代码块中。
` // good async function myFunction () { try { await somethingThatReturnsAPromise(); } catch (err) { console.log(err); } } // bad async function myFunction () { await somethingThatReturnsAPromise() .catch( function (err) { console.log(err) }) } 复制代码` * 多个await命令后面的异步操作，如果不存在继发关系，最好让它们同时触发。
` // good let [foo, bar] = await Promise.all([getFoo(), getBar()]); // bad let fooPromise = getFoo(); let barPromise = getBar(); let foo = await fooPromise; let bar = await barPromise; 复制代码` * await命令只能用在async函数之中，如果用在普通函数，就会报错。
* async 函数可以保留运行堆栈.