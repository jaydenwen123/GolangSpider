# 了解js运行机制——微任务与宏任务 #

由一道面试题引发的思考。

` setTimeout( function ( ) { console.log( 4 ); }, 0 ); new Promise ( function ( reslove ) { console.log( 1 ); reslove(); }).then( function ( data ) { console.log( 3 ); }); console.log( 2 ); 复制代码`

会输出：1，2，3，4。我们来想一下为什么。

浏览器中的事件循环 eventLoop，分为同步执行栈和异步队列，首先会执行同步的任务，当同步任务执行完之后会从异步队列中取异步任务拿到同步执行栈中进行执行。

在取异步队列时，还会有一个区分，就是区分微任务和宏任务。

* microtask：微任务，优先级高，并且可以插队，不是先定义先执行。包括：promise 中的 then，observer，MutationObserver，setImmediate
* macrotask：宏任务，优先级低，先定义的先执行。包括：ajax，setTimeout，setInterval，事件绑定，postMessage，MessageChannel（用于消息通讯）

因为微任务的优先级较高，所以会先将微任务的异步任务取出来进行执行，当微任务的任务都执行完毕之后，会将宏任务中的任务取出来执行。

我们这次来看一下上面的题，promise 中是同步任务，promise 的 .then 中是异步任务，并且是微任务。使用 setTimeout 是宏任务，即使是延时为 0，也是宏任务。

所以上面的执行顺序就是先将 setTimeout 加入到异步队列的宏任务池中，然后执行 promise 中的 ` console.log(1)` ，再将 promise 的.then 加到异步队列中微任务池中，再执行同步任务 ` console.log(2)` ，当同步任务都执行完之后，去微任务中的任务，执行 ` console.log(3)` ，微任务执行完之后取宏任务，执行 ` console.log(4)` 。所以顺序就是：1，2，3，4。

**扩展：**

将这道面试题进行一些改造：

` setTimeout( function ( ) { console.log( 4 ); }, 0 ); new Promise ( function ( reslove ) { console.log( 1 ); setTimeout( function ( ) { reslove( 'done' ); }, 0 ); reslove( 'first' ); }).then( function ( data ) { console.log(data); }); console.log( 2 ); 复制代码`

这个时候就会输出：1，2，first，4，没有输出 done，有些人可能会想，应该输出 1，2，first，4，done，这个时候你就要知道，当使用 reslove 之后，promise 的状态就从 pedding 变成了 resolve， **promise 的状态更改之后就不能再更改了** ，所以 ` reslove('done')` 就不会执行了。

当我们把 ` reslove('done')` 改成 ` console.log('done')` 的时候，他就会输出 1，2，first，4，done 了。

再做一个更复杂的变更：

` setTimeout( function ( ) { console.log( 1 ); }, 0 ); new Promise ( function ( reslove ) { console.log( 2 ); reslove( 'p1' ); new Promise ( function ( reslove ) { console.log( 3 ); setTimeout( function ( ) { reslove( 'setTimeout2' ); console.log( 4 ); }, 0 ); reslove( 'p2' ); }).then( function ( data ) { console.log(data); }); setTimeout( function ( ) { reslove( 'setTimeout1' ); console.log( 5 ); }, 0 ); }).then( function ( data ) { console.log(data); }); console.log( 6 ); 复制代码`

输出的结果是：2，3，6，p2，p1，1，4，5。

先执行同步的任务，new Promise 中的都是同步任务，所以先输出 2，3，6，然后再执行微任务的，微任务可以插队，所以并不是先定义的 p1 先执行，而且先将 p2 执行，然后执行 p1，当微任务都执行完成之后，执行宏任务，宏任务依次输出 1，4，5，promise 的状态不可以变更，所以 setTimeout1 和 setTimeout2 不会输出。