# 不好意思！耽误你的十分钟，让MVVM原理还给你 #

### 时间在嘀嗒嘀嗒的走着 ###

#### 既然来了就继续看看吧 ####

* 这篇文章其实没有什么鸟用，只不过对于现在的 **前端面试** 而言，已经是一个被问烦了的考点了
* 既然是考点，那么我就想简简单单的来给大家划一下重点

众所周知当下是MVVM盛行的时代，从早期的Angular到现在的React和Vue，再从最初的三分天下到现在的两虎相争。

无疑不给我们的开发带来了一种前所未有的新体验，告别了操作DOM的思维，换上了数据驱动页面的思想，果然时代的进步，改变了我们许多许多。

啰嗦话多了起来，这样不好。我们来进入今天的主题

### 划重点 ###

MVVM 双向数据绑定 在Angular1.x版本的时候通过的是 **脏值检测** 来处理

而现在无论是React还是Vue还是最新的Angular，其实实现方式都更相近了

那就是通过 **数据劫持+发布订阅模式**

真正实现其实靠的也是ES5中提供的 **Object.defineProperty** ，当然这是不兼容的所以Vue等只支持了IE8+

#### 为什么是它 ####

Object.defineProperty()说实在的我们大家在开发中确实用的不多，多数是修改内部特性，不过就是定义对象上的属性和值么？干嘛搞的这么费劲(纯属个人想法)

But在实现框架or库的时候却发挥了大用场了，这个就不多说了，只不过轻舟一片而已，还没到写库的实力

**知其然要知其所以然，来看看如何使用**

` let obj = {}; let song = '发如雪' ; obj.singer = '周杰伦' ; Object.defineProperty(obj, 'music' , { // 1. value: '七里香' , configurable: true , // 2. 可以配置对象，删除属性 // writable: true , // 3. 可以修改对象 enumerable: true , // 4. 可以枚举 // ☆ get, set 设置时不能设置writable和value，它们代替了二者且是互斥的 get () { // 5. 获取obj.music的时候就会调用get方法 return song; }, set (val) { // 6. 将修改的值重新赋给song song = val; } }); // 下面打印的部分分别是对应代码写入顺序执行 console.log(obj); // {singer: '周杰伦' , music: '七里香' } // 1 delete obj.music; // 如果想对obj里的属性进行删除，configurable要设为 true 2 console.log(obj); // 此时为 {singer: '周杰伦' } obj.music = '听妈妈的话' ; // 如果想对obj的属性进行修改，writable要设为 true 3 console.log(obj); // {singer: '周杰伦' , music: "听妈妈的话" } for ( let key in obj) { // 默认情况下通过defineProperty定义的属性是不能被枚举(遍历)的 // 需要设置enumerable为 true 才可以 // 不然你是拿不到music这个属性的，你只能拿到singer console.log(key); // singer, music 4 } console.log(obj.music); // '发如雪' 5 obj.music = '夜曲' ; // 调用 set 设置新的值 console.log(obj.music); // '夜曲' 6 复制代码`

以上是关于Object.defineProperty的用法

下面我们来写个实例看看，这里我们以Vue为参照去实现怎么写MVVM

` // index.html <body> <div id= "app" > <h1>{{song}}</h1> <p>《{{album.name}}》是{{singer}}2005年11月发行的专辑</p> <p>主打歌为{{album.theme}}</p> <p>作词人为{{singer}}等人。</p> 为你弹奏肖邦的{{album.theme}} </div> <!--实现的mvvm--> <script src= "mvvm.js" ></script> <script> // 写法和Vue一样 let mvvm = new Mvvm({ el: '#app' , data: { // Object.defineProperty(obj, 'song' , '发如雪' ); song: '发如雪' , album: { name: '十一月的萧邦' , theme: '夜曲' }, singer: '周杰伦' } }); </script> </body> 复制代码`

上面是html里的写法，相信用过Vue的同学并不陌生

那么现在就开始实现一个自己的MVVM吧

### 打造MVVM ###

` // 创建一个Mvvm构造函数 // 这里用es6方法将options赋一个初始值，防止没传，等同于options || {} function Mvvm(options = {}) { // vm. $options Vue上是将所有属性挂载到上面 // 所以我们也同样实现,将所有属性挂载到了 $options this. $options = options; // this._data 这里也和Vue一样 let data = this._data = this. $options.data; // 数据劫持 observe(data); } 复制代码`

#### 数据劫持 ####

为什么要做数据劫持？

* 观察对象，给对象增加Object.defineProperty
* vue特点是不能新增不存在的属性 不存在的属性没有get和set
* 深度响应 因为每次赋予一个新对象时会给这个新对象增加defineProperty(数据劫持)

多说无益，一起看代码

` // 创建一个Observe构造函数 // 写数据劫持的主要逻辑 function Observe(data) { // 所谓数据劫持就是给对象增加get, set // 先遍历一遍对象再说 for ( let key in data) { // 把data属性通过defineProperty的方式定义属性 let val = data[key]; observe(val); // 递归继续向下找，实现深度的数据劫持 Object.defineProperty(data, key, { configurable: true , get () { return val; }, set (newVal) { // 更改值的时候 if (val === newVal) { // 设置的值和以前值一样就不理它 return ; } val = newVal; // 如果以后再获取值(get)的时候，将刚才设置的值再返回去 observe(newVal); // 当设置为新值后，也需要把新值再去定义成属性 } }); } } // 外面再写一个函数 // 不用每次调用都写个new // 也方便递归调用 function observe(data) { // 如果不是对象的话就直接 return 掉 // 防止递归溢出 if (!data || typeof data !== 'object' ) return ; return new Observe(data); } 复制代码`

以上代码就实现了数据劫持，不过可能也有些疑惑的地方比如：递归

再来细说一下为什么递归吧，看这个栗子

` let mvvm = new Mvvm({ el: '#app' , data: { a: { b: 1 }, c: 2 } }); 复制代码`

我们在控制台里看下

![](https://user-gold-cdn.xitu.io/2018/3/31/162797a1132d2905?imageView2/0/w/1280/h/960/ignore-error/1) 被标记的地方就是通过 **递归** observe(val)进行数据劫持添加上了get和set，递归继续向a里面的对象去定义属性，亲测通过可放心食用

接下来说一下observe(newVal)这里为什么也要递归

还是在可爱的控制台上，敲下这么一段代码 mvvm._data.a = {b:'ok'}

然后继续看图说话

![](https://user-gold-cdn.xitu.io/2018/3/31/1627983927aca52a?imageView2/0/w/1280/h/960/ignore-error/1) 通过observe(newVal)加上了 ![](https://user-gold-cdn.xitu.io/2018/3/31/1627983e0e528729?imageView2/0/w/1280/h/960/ignore-error/1) 现在大致明白了为什么要对设置的新值也进行递归observe了吧，哈哈，so easy

数据劫持已完成，我们再做个数据代理

#### 数据代理 ####

数据代理就是让我们每次拿data里的数据时，不用每次都写一长串，如mvvm._data.a.b这种，我们其实可以直接写成mvvm.a.b这种显而易见的方式

下面继续看下去，+号表示实现部分

` function Mvvm(options = {}) { // 数据劫持 observe(data); // this 代理了this._data + for ( let key in data) { Object.defineProperty(this, key, { configurable: true , get () { return this._data[key]; // 如this.a = {b: 1} }, set (newVal) { this._data[key] = newVal; } }); + } } // 此时就可以简化写法了 console.log(mvvm.a.b); // 1 mvvm.a.b = 'ok' ; console.log(mvvm.a.b); // 'ok' 复制代码`

写到这里数据劫持和数据代理都实现了，那么接下来就需要编译一下了，把{{}}里面的内容解析出来

#### 数据编译 ####

` function Mvvm(options = {}) { // observe(data); // 编译 + new Compile(options.el, this); } // 创建Compile构造函数 function Compile(el, vm) { // 将el挂载到实例上方便调用 vm. $el = document.querySelector(el); // 在el范围里将内容都拿到，当然不能一个一个的拿 // 可以选择移到内存中去然后放入文档碎片中，节省开销 let fragment = document.createDocumentFragment(); while (child = vm. $el.firstChild) { fragment.appendChild(child); // 此时将el中的内容放入内存中 } // 对el里面的内容进行替换 function replace(frag) { Array.from(frag.childNodes).forEach(node => { let txt = node.textContent; let reg = /\{\{(.*?)\}\}/g; // 正则匹配{{}} if (node.nodeType === 3 && reg.test(txt)) { // 即是文本节点又有大括号的情况{{}} console.log(RegExp. $1 ); // 匹配到的第一个分组 如： a.b, c let arr = RegExp. $1.split( '.' ); let val = vm; arr.forEach(key => { val = val[key]; // 如this.a.b }); // 用trim方法去除一下首尾空格 node.textContent = txt.replace(reg, val).trim(); } // 如果还有子节点，继续递归replace if (node.childNodes && node.childNodes.length) { replace(node); } }); } replace(fragment); // 替换内容 vm. $el.appendChild(fragment); // 再将文档碎片放入el中 } 复制代码`

看到这里在面试中已经可以初露锋芒了，那就一鼓作气，做事做全套，来个一条龙

现在数据已经可以编译了，但是我们手动修改后的数据并没有在页面上发生改变

下面我们就来看看怎么处理，其实这里就用到了特别常见的设计模式，发布订阅模式

#### 发布订阅 ####

发布订阅主要靠的就是数组关系，订阅就是放入函数，发布就是让数组里的函数执行

` // 发布订阅模式 订阅和发布 如[fn1, fn2, fn3] function Dep () { // 一个数组(存放函数的事件池) this.subs = []; } Dep.prototype = { addSub(sub) { this.subs.push(sub); }, notify () { // 绑定的方法，都有一个update方法 this.subs.forEach(sub => sub.update()); } }; // 监听函数 // 通过Watcher这个类创建的实例，都拥有update方法 function Watcher(fn) { this.fn = fn; // 将fn放到实例上 } Watcher.prototype.update = function () { this.fn(); }; let watcher = new Watcher(() => console.log(111)); // let dep = new Dep(); dep.addSub(watcher); // 将watcher放到数组中,watcher自带update方法， => [watcher] dep.addSub(watcher); dep.notify(); // 111, 111 复制代码`

#### 数据更新视图 ####

* 现在我们要订阅一个事件，当数据改变需要重新刷新视图，这就需要在replace替换的逻辑里来处理
* 通过new Watcher把数据订阅一下，数据一变就执行改变内容的操作

` function replace(frag) { // 省略... // 替换的逻辑 node.textContent = txt.replace(reg, val).trim(); // 监听变化 // 给Watcher再添加两个参数，用来取新的值(newVal)给回调函数传参 + new Watcher(vm, RegExp. $1 , newVal => { node.textContent = txt.replace(reg, newVal).trim(); + }); } // 重写Watcher构造函数 function Watcher(vm, exp, fn) { this.fn = fn; + this.vm = vm; + this.exp = exp; // 添加一个事件 // 这里我们先定义一个属性 + Dep.target = this; + let arr = exp.split( '.' ); + let val = vm; + arr.forEach(key => { // 取值 + val = val[key]; // 获取到this.a.b，默认就会调用get方法 + }); + Dep.target = null; } 复制代码`

当获取值的时候就会自动调用get方法，于是我们去找一下数据劫持那里的get方法

` function Observe(data) { + let dep = new Dep(); // 省略... Object.defineProperty(data, key, { get () { + Dep.target && dep.addSub(Dep.target); // 将watcher添加到订阅事件中 [watcher] return val; }, set (newVal) { if (val === newVal) { return ; } val = newVal; observe(newVal); + dep.notify(); // 让所有watcher的update方法执行即可 } }) } 复制代码`

当set修改值的时候执行了dep.notify方法，这个方法是执行watcher的update方法，那么我们再对update进行修改一下

` Watcher.prototype.update = function () { // notify的时候值已经更改了 // 再通过vm, exp来获取新的值 + let arr = this.exp.split( '.' ); + let val = this.vm; + arr.forEach(key => { + val = val[key]; // 通过get获取到新的值 + }); this.fn(val); // 将每次拿到的新值去替换{{}}的内容即可 }; 复制代码`

现在我们数据的更改可以修改视图了，这很good，还剩最后一点，我们再来看看面试常考的双向数据绑定吧

#### 双向数据绑定 ####

` // html结构 <input v-model= "c" type = "text" > // 数据部分 data: { a: { b: 1 }, c: 2 } function replace(frag) { // 省略... + if (node.nodeType === 1) { // 元素节点 let nodeAttr = node.attributes; // 获取dom上的所有属性,是个类数组 Array.from(nodeAttr).forEach(attr => { let name = attr.name; // v-model type let exp = attr.value; // c text if (name.includes( 'v-' )){ node.value = vm[exp]; // this.c 为 2 } // 监听变化 new Watcher(vm, exp, function (newVal) { node.value = newVal; // 当watcher触发时会自动将内容放进输入框中 }); node.addEventListener( 'input' , e => { let newVal = e.target.value; // 相当于给this.c赋了一个新值 // 而值的改变会调用 set ， set 中又会调用notify，notify中调用watcher的update方法实现了更新 vm[exp] = newVal; }); }); + } if (node.childNodes && node.childNodes.length) { replace(node); } } 复制代码`

大功告成，面试问Vue的东西不过就是这个罢了，什么双向数据绑定怎么实现的，问的一点心意都没有，差评！！！

**大官人请留步** ，本来应该收手了，可临时起意(手痒)，再写点功能吧，再加个computed(计算属性)和mounted(钩子函数)吧

#### computed(计算属性) && mounted(钩子函数) ####

` // html结构 <p>求和的值是{{sum}}</p> data: { a: 1, b: 9 }, computed: { sum () { return this.a + this.b; }, noop () {} }, mounted () { set Timeout(() => { console.log( '所有事情都搞定了' ); }, 1000); } function Mvvm(options = {}) { // 初始化computed,将this指向实例 + initComputed.call(this); // 编译 new Compile(options.el, this); // 所有事情处理好后执行mounted钩子函数 + options.mounted.call(this); // 这就实现了mounted钩子函数 } function initComputed () { let vm = this; let computed = this. $options.computed; // 从options上拿到computed属性 {sum: ƒ, noop: ƒ} // 得到的都是对象的key可以通过Object.keys转化为数组 Object.keys(computed).forEach(key => { // key就是sum,noop Object.defineProperty(vm, key, { // 这里判断是computed里的key是对象还是函数 // 如果是函数直接就会调get方法 // 如果是对象的话，手动调一下get方法即可 // 如： sum () { return this.a + this.b;},他们获取a和b的值就会调用get方法 // 所以不需要new Watcher去监听变化了 get: typeof computed[key] === 'function' ? computed[key] : computed[key].get, set () {} }); }); } 复制代码`

写了这些内容也不算少了，最后做一个形式上的总结吧

### 总结 ###

通过自己实现的mvvm一共包含了以下东西

* 通过Object.defineProperty的get和set进行数据劫持
* 通过遍历data数据进行数据代理到this上
* 通过{{}}对数据进行编译
* 通过发布订阅模式实现数据与视图同步
* 通过通过通过，收了，感谢大官人的留步了

### 补充 ###

针对以上代码在实现编译的时候还是会有一些小bug，再次经过研究和高人指点，完善了编译，下面请看修改后的代码

**修复** ：两个相邻的{{}}正则匹配，后一个不能正确编译成对应的文本，如{{album.name}} {{singer}}

` function Compile(el, vm) { // 省略... function replace(frag) { // 省略... if (node.nodeType === 3 && reg.test(txt)) { function replaceTxt () { node.textContent = txt.replace(reg, (matched, placeholder) => { console.log(placeholder); // 匹配到的分组 如：song, album.name, singer... new Watcher(vm, placeholder, replaceTxt); // 监听变化，进行匹配替换内容 return placeholder.split( '.' ).reduce((val, key) => { return val[key]; }, vm); }); }; // 替换 replaceTxt(); } } } 复制代码`

上面代码主要实现依赖的是reduce方法，reduce 为数组中的每一个元素依次执行回调函数

如果还有不太清楚的，那我们单独抽出来reduce这部分再看一下

` // 将匹配到的每一个值都进行split分割 // 如: 'song'.split( '.' ) => [ 'song' ] => [ 'song' ].reduce((val, key) => val[key]) // 其实就是将vm传给val做初始值，reduce执行一次回调返回一个值 // vm[ 'song' ] => '周杰伦' // 上面不够深入，我们再来看一个 // 再如： 'album.name'.split( '.' ) => [ 'album' , 'name' ] => [ 'album' , 'name' ].reduce((val, key) => val[key]) // 这里vm还是做为初始值传给val，进行第一次调用，返回的是vm[ 'album' ] // 然后将返回的vm[ 'album' ]这个对象传给下一次调用的val // 最后就变成了vm[ 'album' ][ 'name' ] => '十一月的萧邦' return placeholder.split( '.' ).reduce((val, key) => { return val[key]; }, vm); 复制代码`

reduce的用处多多，比如计算数组求和是比较普通的方法了，还有一种比较好用的妙处是可以进行二维数组的展平(flatten)，各位不妨来看最后一眼

` let arr = [ [1, 2], [3, 4], [5, 6] ]; let flatten = arr.reduce((previous, current) => { return previous.concat(current); }); console.log(flatten); // [1, 2, 3, 4, 5, 6] // ES6中也可以利用...展开运算符来实现的，实现思路一样，只是写法更精简了 flatten = arr.reduce((a, b) => [...a, ...b]); console.log(flatten); // [1, 2, 3, 4, 5, 6] 复制代码`

再次感谢父老乡亲，兄弟姐妹们的观看了！这回真的是最后一眼了，已经到底了！