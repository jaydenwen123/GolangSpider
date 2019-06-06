# 探索async #

` async` 本质是什么? 其实就是 ` generator` 的语法糖，虽然 ` ES6` 已经实现了 ` async` 和 ` generator` ，但是在生产环境中都是经过 ` babel` 编译成 ` promise`.

### 简单的 ` async` ###

` async function p1 (){ console.log(1) return 1 } // 等效 function p1 (){ console.log(1) return Promise.resolve(1) } 复制代码`

### 一般情况的 ` async` ###

` async function p2 (){ console.log(2) let a = await new Promise(resolve => set Timeout(() =>resolve(1), 3000) ) console.log(a); return 2 }; // 等效 function p2 (){ console.log(2) return Promise.resolve(new Promise(resolve => set Timeout(() =>resolve(1), 3000)).then(res => { let a = res; console.log(a) return Promise.resolve(2) })) } 复制代码`

### 循环中的 ` async` ###

里面如何实现这个我还真不清楚，但是可以通过队列模拟其实现

` let p1 = () => new Promise((resolve => set Timeout(() => resolve(1), 1000))); let p2 = () => new Promise((resolve => set Timeout(() => resolve(2), 2000))); let p3 = () => new Promise((resolve => set Timeout(() => resolve(3), 3000))); let ps = [p1, p2, p3]; async function p (){ for ( let i = 0; i < 3; i++) { let a = await ps[i](); console.log(a) } } // 等效 async function p (){ let queue = []; for ( let i = 0; i < 3; i++) { queue.push(() => ps[i]().then(res => { let a = res; console.log(a) })) } queue.reduce((p1, p2) => p1.then(res => p2()) , Promise.resolve()); }; 复制代码`

可以看出来 ` async` 简化了 ` Promise` ,大部分场景下 ` Promise` 其实也够用了，但是在链式调用的场景下使用 ` async` 非常简洁. 下面抛出一个问题

` let p1 = new Promise((resolve) => resolve(1)) let p2 = new Promise((resolve) => { resolve(p1); new Promise(resolve => resolve()).then(() => console.log(2)); }) p2.then(res => console.log(1)) // 为什么先执行2然后执行1? // 提示: 把resolve(p1)转换一下就很清楚了 复制代码`