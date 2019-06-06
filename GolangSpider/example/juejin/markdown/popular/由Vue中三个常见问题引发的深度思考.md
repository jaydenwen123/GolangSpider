# 由Vue中三个常见问题引发的深度思考 #

### 前言 ###

工作中我们通过搜索引擎或者官方文档很容易就会知道一个语法怎么使用，但是你知道其中的原理吗？我想有一部分同学应该做不到清楚的说明其实现原理。众所周知，如今技术更新迭代速度很快，据 Vue 作者尤雨溪表示 Vue3.x 会在今年的下半年发布正式版本，视频地址在这里 [ VUE CONF 杭州之 3.0 进展 ]( https://link.juejin.im?target=https%3A%2F%2Fv.qq.com%2Fx%2Fpage%2Fk0802iqtskt.html ) 。如果你在使用 Vue 或者正在学习 Vue，那么我建议你观看一遍完整的视频，这可能对你以后接触 3.0 版本有很大帮助。这篇文章是基于 Vue2.x 的版本讲述的。主要是围绕以下三个问题展开讨论。

> 
> 
> 
> * 为什么 ` data` 要写成函数，而不允许写成对象？
> * Vue 中常说的数据劫持到底是什么？
> * Vue 实例中数组改变 ` length` 或下标赋值为什么不能更新视图？
> 
> 
> 

### Tips ###

> 
> 
> 
> 如果你已经掌握了这三个问题的原因和原理；
> 
> 
> 
> 如果你觉得你不需要掌握原理会用即可；
> 
> 
> 
> 抖个机灵——请看本文最后一行。
> 
> 

接下来，我们就对这三个问题一一解答。

### 问题一 ###

> 
> 
> 
> #### 为什么 ` data` 要写成函数，而不允许写成对象？ ####
> 
> 

想要理解这个问题，我们首先要知道以下三点。

#### 注意要点 ####

> 
> * ` data` 是 Vue 实例上的一个属性。2. 对象是对于内存地址的引用。3. 函数有自己的作用域空间。
> 

第一点无可厚非， ` data` 属性附着于 Vue 实例上。

第二点，JS 的数据类型分为基本类型和引用类型，基本类型存储在栈内存中，引用类型存储在堆内存中，并且引用类型指向的是栈内存中的堆区地址。下面两个例子可以帮助你清晰地理解这句话。

#### 基本类型赋值 ####

` var a = 10 ; var b = 10 ; var c = a; console.log(a === b); // true a ++ ; console.log(a); // 11 console.log(c); // 10 复制代码`

这段代码分别给 a、b 赋值 10，a 和 b 是全等的。然后用 a 来初始化 c，那么 c 的值也是 10。但 c 中的 10 与 a 中的是完全独立的，该值只是 a 中的值的一个副本，此后， 这两个变量可以参加任何操作而相互不受影响。具体位置如下示意图。

![](https://user-gold-cdn.xitu.io/2019/5/27/16af9736eb0b8d88?imageView2/0/w/1280/h/960/ignore-error/1)

#### 引用类型赋值 ####

` var a = {}; var b = {}; var c = a; console.log(a === b); // false a.name = 'Marry' ; a.say = () => console.log( 'Hi Marry!' ); console.log(c.name); // 'Marry' console.log(c.say()); // 'Hi Marry!' 复制代码`

上面这段代码。首先声明了a、b两个空对象，然后把 a 赋值给 c。因为对象是对栈内存的地址的引用，所以不同的对象的地址是不同的，所以他们不是全等的。接着给 a 新增加属性和方法，c 同样可以拥有此属性和方法，主要是因为 c 和 a 指向堆内存中的同一个地址。其关系图如下示意图所示。

![](https://user-gold-cdn.xitu.io/2019/5/27/16af981109d71087?imageView2/0/w/1280/h/960/ignore-error/1)

至于第三点，大多数有 JS 基础的同学应该都能理解，每个函数都有自己的作用域。

以上是对三个注意点的说明，那么接下来我们就以两个例子解释问题一： 为什么 ` data` 要写成函数，而不允许写成对象？

#### data 为对象示例代码 ####

` function MyCompnent ( ) {} MyCompnent.prototype.data = { age : 12 }; var JackMa = new MyCompnent(); var PonyMa = new MyCompnent(); console.log(JackMa.data.age === PonyMa.data.age); // true JackMa.data.age = 13 ; console.log( 'JackMa ' + JackMa.data.age + '岁；' + 'PonyMa ' + PonyMa.data.age + '岁' ); // JackMa 13岁；PonyMa 13岁 复制代码`

上面的示例中，我们创建一个构造函数 MyCompnent，它充当的角色相当于 Vue，在他的原型属性上声明一个 ` data` 属性，其实也相当于 ` Vue.$data` 。接着声明两个实例，改变其中一个实例，另外一个实例也会跟着改变，这个道理其实和引用类型赋值大同小异。

#### data 为函数的示例代码 ####

` function MyCompnent ( ) { this.data = this.data(); } MyCompnent.prototype.data = function ( ) { return { age : 12 } }; var JackMa = new MyCompnent(); var PonyMa = new MyCompnent(); console.log(JackMa.data.age === PonyMa.data.age); // true JackMa.data = { age : 13 }; console.log( 'JackMa ' + JackMa.data.age + '岁；' + 'PonyMa ' + PonyMa.data.age + '岁' ); // JackMa 13岁；PonyMa 12岁 复制代码`

上述代码模拟了 Vue 实例上的 ` data` 为函数的时候，如果改变一个实例的 ` data` 属性的值，那么不会影响到另外一个实例上的 ` data` 的值。

#### Tips ####

> 
> 
> 
> 面试过程中经常会被问到 JS
> 数据类型问题，如果你只回答基本类型和引用类型可能你只能得到一半分数，但是如果你能把存储位置等要点回答出来并且举例说明，想必是很加分的。
> 
> 

#### 小结 ####

Vue 里面 ` data` 属性之所以不能写成对象的格式，是因为对象是对地址的引用，而不是独立存在的。如果一个.vue 文件有多个子组件共同接收一个变量的话，改变其中一个子组件内此变量的值，会影响其他组件的这个变量的值。如果写成函数的话，那么他们有一个作用域的概念在里面，相互隔阂，不受影响。

### 问题二 ###

> 
> 
> 
> #### Vue 中常说的数据劫持到底是什么？ ####
> 
> 

相信大多数用过或者了解 Vue 的同学都听过 数据劫持 ，进一步问为什么可能你也能答出一二，例如 getter、setter 之类。今天我就系统地和你说一下数据劫持之美。首先我们先看一看下图。

![](https://user-gold-cdn.xitu.io/2019/5/27/16af9a047a409676?imageView2/0/w/1280/h/960/ignore-error/1)

上图完整的描述了 Vue 运行的机制，首先数据发生改变，就会经过 ` Data` 处理，然后 ` Dep` 会发出通知( ` notify` )，告诉 ` Watcher` 有数据发生了变化，接着 ` Watcher` 会传达给渲染函数跟他说有数据变化了，可以渲染视图了(数据驱动视图)，进而渲染函数执行 ` render` 方法去更新 ` VNODE` ，也就是我们说的虚拟DOM，最后虚拟DOM根据最优算法，去局部更新需要渲染的视图。这里的 ` Data` 就做了我们今天要说的事——数据劫持。

想要更深入地理解如何劫持，我们就需要看源码实现。

#### Observer ####

` /** * Vue中的每一个变量都是由 Observer 构造函数生成的。 * 细心的你可能会发现，你打印出来任何一个Vue上的引用类型属性，后面都有 __ob__: Observer 的字样。 */ var Observer = function Observer ( value ) { this.value = value; // 这里把发布者 Dep 注册了 this.dep = new Dep(); // ··· // 此处调用 walk this.walk(value); }; 复制代码` ` /** * 此处会将 obj 里面的每一个值用 defineReactive$$1 处理，而它就是今晚的主角。 */ Observer.prototype.walk = function walk ( obj ) { var keys = Object.keys(obj); for ( var i = 0 ; i < keys.length; i++) { defineReactive$$ 1 (obj, keys[i]); } }; 复制代码`

#### 数据劫持 ####

` /** * 这个函数就是数据劫持的根据地，里面为对象重写了 get 和 set 方法以及固有属性 enumerable 等。 * */ function defineReactive$$1 ( obj, key, val, customSetter, shallow ) { // ··· Object.defineProperty(obj, key, { enumerable : true , configurable : true , get : function reactiveGetter ( ) { // ··· return value }, set : function reactiveSetter ( newVal ) { // ··· // 此处最为关键，这个函数的主要作用就是通过 notify 告诉 Watcher 有数据变化了。 dep.notify(); } }); } 复制代码`

#### Dep ####

` /** * subs 是所有 Watcher 的收集器，类型为数组；notify 实则是调用了每个Watcher的 update方法 。 */ Dep.prototype.notify = function notify ( ) { var subs = this.subs.slice(); // ··· for ( var i = 0 , l = subs.length; i < l; i++) { subs[i].update(); } }; 复制代码`

#### Watcher ####

` /** * 更新视图的最直观的方法就是 Watcher 上的 update 方法 ， Dep subs 反复调用 * 这里最终都是调用 run 方法。 */ Watcher.prototype.update = function update ( ) { /* istanbul ignore else */ if ( this.lazy) { this.dirty = true ; } else if ( this.sync) { this.run(); } else { queueWatcher( this ); } }; 复制代码` ` /** * run 方法内的 cb 方法建立 Watcher 和 VNode 之间的关联关系，从而引发视图的更新。 */ Watcher.prototype.run = function run ( ) { // ··· if ( this.user) { // ··· this.cb.call( this.vm, value, oldValue); // ··· } else { this.cb.call( this.vm, value, oldValue); } // ··· }; 复制代码`

#### 小结 ####

至此，我们不但了解了数据劫持的的原理，还知道了谁去劫持，劫持过后做了什么，是谁引发的视图更新等等。是不是对 Vue 的运行机制更明白了一些呢？

### 问题三 ###

> 
> 
> 
> #### Vue 实例中数组改变 ` length` 或下标赋值为什么不能更新视图？ ####
> 
> 

#### 情景再现 ####

![](https://user-gold-cdn.xitu.io/2019/5/28/16afbf6d5a61d28c?imageView2/0/w/1280/h/960/ignore-error/1)

上图示例中，为方便调试在 mounted 周期内执行 ` windows.vm = this;` 。 ` week` 包含周一到周五五个元素，我们尝试改变 ` week` 的 ` length` 为 3 以及给它下标为 4 的元素赋值一个 ` 周八` ，结果都没有生效。那怎么可以生效呢？请看下图。

![](https://user-gold-cdn.xitu.io/2019/5/28/16afbfbbf876130a?imageView2/0/w/1280/h/960/ignore-error/1) 主要是因为我们调用了一个数组的内置方法 ` push` ，如果你愿意尝试，你会发现调用数组的 ` slice` 方法是不行的。只要是因为 Vue 提取了数组的可以改变原数组的原生方法，进行了再加工。只有经过 Vue 处理过的方法才有更新视图的能力。下面我将从内置方法和源码的角度给大家说明这个结论。

#### 内置方法 ####

![](https://user-gold-cdn.xitu.io/2019/5/28/16afc0137c8c52ca?imageView2/0/w/1280/h/960/ignore-error/1) 上图中我们展开 week 数组，发现在第一个 ` __proto__` 里面内置了 ` pop` 、 ` push` 等多个数组的方法；在第二个 ` __proto__` 不但有上面几种还有更多其他的方法。由此可见，Vue 是对数组的 Api 进行了劫持。

#### 源码解析 ####

` var methodsToPatch = [ 'push' , 'pop' , 'shift' , 'unshift' , 'splice' , 'sort' , 'reverse' ]; /** * Intercept mutating methods and emit events * 此方法主要作用就是遍历数组局部方法，调用的同时去调用 dep 的 notify 通知 Watcher 进而更新视图 */ methodsToPatch.forEach( function ( method ) { // cache original method var original = arrayProto[method]; def(arrayMethods, method, function mutator ( ) { var args = [], len = arguments.length; while ( len-- ) args[ len ] = arguments [ len ]; var result = original.apply( this , args); var ob = this.__ob__; var inserted; switch (method) { case 'push' : case 'unshift' : inserted = args; break case 'splice' : inserted = args.slice( 2 ); break } if (inserted) { ob.observeArray(inserted); } // 这一步至关重要 ob.dep.notify(); return result }); }); 复制代码`

#### 小结 ####

到这里我们知道了，Vue 劫持了数组可以改变原数组的 Api，使得每次调用都会执行 ` dep.notify()` 方法进而去更新视图。

### 结束语 ###

分享的时光永远这么短暂，但我还要对你说：工作中经常遇到此类问题，希望我们多问自己一个为什么？去研究它到底是怎么实现的，掌握设计理念，学习设计思想，而不是仅限于知道如何使用。这样自己才会有更多的成长！

最后，如果你愿意 点个赞 ，这将给我很大的动力！