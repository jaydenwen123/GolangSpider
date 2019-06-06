# 一个合格的中级前端工程师必须要掌握的 28 个 JavaScript 技巧 #

![](https://user-gold-cdn.xitu.io/2019/5/30/16b06aa76f8b5b33?imageView2/0/w/1280/h/960/ignore-error/1)

## 前言 ##

> 
> 
> 
> 文中代码对应的详细注释和具体使用方法都放在我的 github 上，源代码在底部连接
> 
> 

## 1.判断对象的数据类型 ##

![](https://user-gold-cdn.xitu.io/2019/5/30/16b08a60a70f5077?imageView2/0/w/1280/h/960/ignore-error/1)

使用 Object.prototype.toString 配合闭包，通过传入不同的判断类型来返回不同的判断函数，一行代码，简洁优雅灵活（注意传入 type 参数时首字母大写）

**不推荐将这个函数用来检测可能会产生包装类型的基本数据类型上,因为 call 会将第一个参数进行装箱操作**

## 2. ES5 实现数组 map 方法 ##

![](https://user-gold-cdn.xitu.io/2019/5/30/16b089d425e4a174?imageView2/0/w/1280/h/960/ignore-error/1)

值得一提的是，map 的第二个参数为第一个参数回调中的 this 指向，如果第一个参数为箭头函数，那设置第二个 this 会因为箭头函数的词法绑定而失效

另外就是对稀疏数组的处理，通过 hasOwnProperty 来判断当前下标的元素是否存在与数组中(感谢评论区的朋友)

## 3. 使用 reduce 实现数组 map 方法 ##

![](https://user-gold-cdn.xitu.io/2019/5/30/16b06a9d567bab69?imageView2/0/w/1280/h/960/ignore-error/1)

## 4. ES5 实现数组 filter 方法 ##

![](https://user-gold-cdn.xitu.io/2019/5/30/16b089dd88abc7e9?imageView2/0/w/1280/h/960/ignore-error/1)

## 5. 使用 reduce 实现数组 filter 方法 ##

![](https://user-gold-cdn.xitu.io/2019/5/30/16b06a9d56d1a917?imageView2/0/w/1280/h/960/ignore-error/1)

## 6. ES5 实现数组的 some 方法 ##

![](https://user-gold-cdn.xitu.io/2019/6/2/16b16d392dbd40e2?imageView2/0/w/1280/h/960/ignore-error/1)

执行 some 方法的数组如果是一个空数组，最终始终会返回 ` false` ，而另一个数组的 every 方法中的数组如果是一个空数组，会始终返回 ` true`

## 7. ES5 实现数组的 reduce 方法 ##

![](https://user-gold-cdn.xitu.io/2019/6/2/16b16d431e7028e7?imageView2/0/w/1280/h/960/ignore-error/1)

因为可能存在稀疏数组的关系，所以 reduce 需要保证跳过稀疏元素，遍历正确的元素和下标(感谢@神三元的提供的代码)

## 8. 使用 reduce 实现数组的 flat 方法 ##

![](https://user-gold-cdn.xitu.io/2019/5/30/16b06a9da00a4224?imageView2/0/w/1280/h/960/ignore-error/1)

因为 selfFlat 是依赖 this 指向的，所以在 reduce 遍历时需要指定 selfFlat 的 this 指向，否则会默认指向 window 从而发生错误

原理通过 reduce 遍历数组，遇到数组的某个元素仍是数组时，通过 ES6 的扩展运算符对其进行降维（ES5 可以使用 concat 方法），而这个数组元素可能内部还嵌套数组，所以需要递归调用 selfFlat

同时原生的 flat 方法支持一个 depth 参数表示降维的深度，默认为 1 即给数组降一层维度

![](https://user-gold-cdn.xitu.io/2019/5/30/16b06a9da1074f15?imageView2/0/w/1280/h/960/ignore-error/1)

传入 Inifity 会将传入的数组变成一个一维数组

![](https://user-gold-cdn.xitu.io/2019/5/30/16b06a9da7d2da09?imageView2/0/w/1280/h/960/ignore-error/1)

原理是每递归一次将 depth 参数减 1，如果 depth 参数为 0 时，直接返回原数组

## 9. 实现 ES6 的 class 语法 ##

![](https://user-gold-cdn.xitu.io/2019/6/2/16b16d2e5fbdde5d?imageView2/0/w/1280/h/960/ignore-error/1)

ES6 的 class 内部是基于寄生组合式继承，它是目前最理想的继承方式，通过 Object.create 方法创造一个空对象，并将这个空对象继承 Object.create 方法的参数，再让子类（subType）的原型对象等于这个空对象，就可以实现子类实例的原型等于这个空对象，而这个空对象的原型又等于父类原型对象（superType.prototype）的继承关系

而 Object.create 支持第二个参数，即给生成的空对象定义属性和属性描述符/访问器描述符，我们可以给这个空对象定义一个 constructor 属性更加符合默认的继承行为，同时它是不可枚举的内部属性（enumerable:false）

而 ES6 的 class 允许子类继承父类的静态方法和静态属性，而普通的寄生组合式继承只能做到实例与实例之间的继承，对于类与类之间的继承需要额外定义方法，这里使用 Object.setPrototypeOf 将 superType 设置为 subType 的原型，从而能够从父类中继承静态方法和静态属性

## 10. 函数柯里化 ##

![](https://user-gold-cdn.xitu.io/2019/5/30/16b06a9dc5e210a1?imageView2/0/w/1280/h/960/ignore-error/1)

使用方法：

![](https://user-gold-cdn.xitu.io/2019/5/30/16b06a9dd090eecc?imageView2/0/w/1280/h/960/ignore-error/1)

柯里化是函数式编程的一个重要技巧，将使用多个参数的一个函数转换成一系列使用一个参数的函数的技术

函数式编程另一个重要的函数 compose，能够将函数进行组合，而组合的函数只接受一个参数，所以如果有接受多个函数的需求并且需要用到 compose 进行函数组合，就需要使用柯里化对准备组合的函数进行部分求值，让它始终只接受一个参数

借用 [冴羽博客]( https://juejin.im/post/59a8ccf8f265da2498241625 ) 中的一个例子

![](https://user-gold-cdn.xitu.io/2019/5/30/16b06a9dd66bd4df?imageView2/0/w/1280/h/960/ignore-error/1)

## 11. 函数柯里化（支持占位符） ##

![](https://user-gold-cdn.xitu.io/2019/6/2/16b176609052c6a3?imageView2/0/w/1280/h/960/ignore-error/1)

使用方法：

![](https://user-gold-cdn.xitu.io/2019/5/30/16b06a9deb6779e0?imageView2/0/w/1280/h/960/ignore-error/1)

![](https://user-gold-cdn.xitu.io/2019/5/30/16b06a9df3e80535?imageView2/0/w/1280/h/960/ignore-error/1)

通过占位符能让柯里化更加灵活，实现思路是，每一轮传入的参数先去填充上一轮的占位符，如果当前轮参数含有占位符，则放到内部保存的数组末尾，当前轮的元素不会去填充当前轮参数的占位符，只会填充之前传入的占位符

## 12. 偏函数 ##

![](https://user-gold-cdn.xitu.io/2019/5/30/16b06a9df55ef5b4?imageView2/0/w/1280/h/960/ignore-error/1)

使用方法：

![](https://user-gold-cdn.xitu.io/2019/5/30/16b06a9e05d4ac3a?imageView2/0/w/1280/h/960/ignore-error/1)

偏函数和柯里化概念类似，个人认为它们区别在于偏函数会固定你传入的几个参数，再一次性接受剩下的参数，而函数柯里化会根据你传入参数不停的返回函数，直到参数个数满足被柯里化前函数的参数个数

Function.prototype.bind 函数就是一个偏函数的典型代表，它接受的第二个参数开始，为预先添加到绑定函数的参数列表中的参数，与 bind 不同的是，上面的这个函数同样支持占位符

## 13. 斐波那契数列及其优化 ##

![](https://user-gold-cdn.xitu.io/2019/6/2/16b1799c4567157e?imageView2/0/w/1280/h/960/ignore-error/1)

利用函数记忆，将之前运算过的结果保存下来，对于频繁依赖之前结果的计算能够节省大量的时间，例如斐波那契数列，缺点就是闭包中的 obj 对象会额外占用内存

## 14. 实现函数 bind 方法 ##

![](https://user-gold-cdn.xitu.io/2019/5/30/16b06a9e43bd35eb?imageView2/0/w/1280/h/960/ignore-error/1)

函数的 bind 方法核心是利用 call，同时考虑了一些其他情况，例如

* bind 返回的函数被 new 调用作为构造函数时，绑定的值会失效并且改为 new 指定的对象
* 定义了绑定后函数的 length 属性和 name 属性（不可枚举属性）
* 绑定后函数的原型需指向原来的函数

## 15. 实现函数 call 方法 ##

![](https://user-gold-cdn.xitu.io/2019/5/30/16b06a9e37001a80?imageView2/0/w/1280/h/960/ignore-error/1)

原理就是将函数作为传入的上下文参数（context）的属性执行，这里为了防止属性冲突使用了 ES6 的 Symbol 类型

## 16. 简易的 CO 模块 ##

![](https://user-gold-cdn.xitu.io/2019/5/30/16b06a9e54563b48?imageView2/0/w/1280/h/960/ignore-error/1)

使用方法：

![](https://user-gold-cdn.xitu.io/2019/5/30/16b06a9e5a219b49?imageView2/0/w/1280/h/960/ignore-error/1)

run 函数接受一个生成器函数，每当 run 函数包裹的生成器函数遇到 yield 关键字就会停止，当 yield 后面的 promise 被解析成功后会自动调用 next 方法执行到下个 yield 关键字处，最终就会形成每当一个 promise 被解析成功就会解析下个 promise，当全部解析成功后打印所有解析的结果，衍变为现在用的最多的 async/await 语法

## 17. 函数防抖 ##

![](https://user-gold-cdn.xitu.io/2019/5/30/16b06a9e791beedf?imageView2/0/w/1280/h/960/ignore-error/1)

leading 为是否在进入时立即执行一次， trailing 为是否在事件触发结束后额外再触发一次，原理是利用定时器，如果在规定时间内再次触发事件会将上次的定时器清除，即不会执行函数并重新设置一个新的定时器，直到超过规定时间自动触发定时器中的函数

同时通过闭包向外暴露了一个 cancel 函数，使得外部能直接清除内部的计数器

## 18. 函数节流 ##

![](https://user-gold-cdn.xitu.io/2019/5/30/16b06a9e860d7c7d?imageView2/0/w/1280/h/960/ignore-error/1)

和函数防抖类似，区别在于内部额外使用了时间戳作为判断，在一段时间内没有触发事件才允许下次事件触发

## 19. 图片懒加载 ##

![](https://user-gold-cdn.xitu.io/2019/5/30/16b06a9e883485e5?imageView2/0/w/1280/h/960/ignore-error/1)

getBoundClientRect 的实现方式，监听 scroll 事件（建议给监听事件添加节流），图片加载完会从 img 标签组成的 DOM 列表中删除，最后所有的图片加载完毕后需要解绑监听事件

![](https://user-gold-cdn.xitu.io/2019/5/30/16b06a9ea8913738?imageView2/0/w/1280/h/960/ignore-error/1)

intersectionObserver 的实现方式，实例化一个 IntersectionObserver ，并使其观察所有 img 标签

当 img 标签进入可视区域时会执行实例化时的回调，同时给回调传入一个 entries 参数，保存着实例观察的所有元素的一些状态，比如每个元素的边界信息，当前元素对应的 DOM 节点，当前元素进入可视区域的比率，每当一个元素进入可视区域，将真正的图片赋值给当前 img 标签，同时解除对其的观察

## 20. new 关键字 ##

![](https://user-gold-cdn.xitu.io/2019/5/30/16b06a9eaceb83de?imageView2/0/w/1280/h/960/ignore-error/1)

## 21. 实现 Object.assign ##

![](https://user-gold-cdn.xitu.io/2019/5/30/16b06a9eb12702a0?imageView2/0/w/1280/h/960/ignore-error/1)

Object.assign 的原理可以参考我另外一篇 [博客]( https://juejin.im/post/5c6234f16fb9a049a81fcca5#heading-34 )

## 22. instanceof ##

![](https://user-gold-cdn.xitu.io/2019/5/30/16b06a9eb217b698?imageView2/0/w/1280/h/960/ignore-error/1)

原理是递归遍历 right 参数的原型链，每次和 left 参数作比较，遍历到原型链终点时则返回 false，找到则返回 true

## 23. 私有变量的实现 ##

![](https://user-gold-cdn.xitu.io/2019/5/30/16b06a9ebaeee920?imageView2/0/w/1280/h/960/ignore-error/1)

使用 Proxy 代理所有含有 ` _` 开头的变量，使其不可被外部访问

![](https://user-gold-cdn.xitu.io/2019/5/30/16b06a9ebef85aed?imageView2/0/w/1280/h/960/ignore-error/1)

通过闭包的形式保存私有变量，缺点在于类的所有实例访问的都是同一个私有变量

![](https://user-gold-cdn.xitu.io/2019/5/30/16b06a9edc4661c7?imageView2/0/w/1280/h/960/ignore-error/1)

另一种闭包的实现，解决了上面那种闭包的缺点，每个实例都有各自的私有变量，缺点是舍弃了 class 语法的简洁性，将所有的特权方法（访问私有变量的方法）都保存在构造函数中

![](https://user-gold-cdn.xitu.io/2019/5/30/16b06a9ee0a7b7ec?imageView2/0/w/1280/h/960/ignore-error/1)

通过 WeakMap 和闭包，在每次实例化时保存当前实例和所有私有变量组成的对象，外部无法访问闭包中的 WeakMap，使用 WeakMap 好处在于当没有变量引用到某个实例时，会自动释放这个实例保存的私有变量，减少内存溢出的问题

## 24. 洗牌算法 ##

早前的 chrome 对于元素小于 10 的数组会采用插入排序，这会导致对数组进行的乱序并不是真正的乱序，即使最新的版本 chrome 采用了原地算法使得排序变成了一个稳定的算法，对于乱序的问题仍没有解决

![](https://user-gold-cdn.xitu.io/2019/5/30/16b06a9ee7f628db?imageView2/0/w/1280/h/960/ignore-error/1)

![](https://user-gold-cdn.xitu.io/2019/5/30/16b06a9eeefdecc9?imageView2/0/w/1280/h/960/ignore-error/1)

通过洗牌算法可以达到真正的乱序，洗牌算法分为原地和非原地，图一是原地的洗牌算法，不需要声明额外的数组从而更加节约内存占用率，原理是依次遍历数组的元素，将当前元素和之后的所有元素中随机选取一个，进行交换

## 25. 单例模式 ##

![](https://user-gold-cdn.xitu.io/2019/5/30/16b06a9ef0559e13?imageView2/0/w/1280/h/960/ignore-error/1)

通过 ES6 的 Proxy 拦截构造函数的执行方法来实现的单例模式

## 26. promisify ##

![](https://user-gold-cdn.xitu.io/2019/5/30/16b06a9ef50d20a8?imageView2/0/w/1280/h/960/ignore-error/1)

使用方法：

![](https://user-gold-cdn.xitu.io/2019/5/30/16b06a9f1d929229?imageView2/0/w/1280/h/960/ignore-error/1)

promisify 函数是将回调函数变为 promise 的辅助函数，适合 error-first 风格（nodejs）的回调函数，原理是给 error-first 风格的回调无论成功或者失败，在执行完毕后都会执行最后一个回调函数，我们需要做的就是让这个回调函数控制 promise 的状态即可

这里还用了 Proxy 代理了整个 fs 模块，拦截 get 方法，使得不需要手动给 fs 模块所有的方法都包裹一层 promisify 函数，更加的灵活

## 27. 优雅的处理 async/await ##

![](https://user-gold-cdn.xitu.io/2019/5/31/16b0bc9cc35dd332?imageView2/0/w/1280/h/960/ignore-error/1)

使用方法：

![](https://user-gold-cdn.xitu.io/2019/5/31/16b0bca88167ffce?imageView2/0/w/1280/h/960/ignore-error/1)

无需每次使用 async/await 都包裹一层 try/catch ，更加的优雅，这里提供另外一个思路，如果使用了 webpack 可以编写一个 loader，分析 AST 语法树，遇到 await 语法，自动注入 try/catch，这样连辅助函数都不需要使用

## 28. 发布订阅 EventEmitter ##

![](https://user-gold-cdn.xitu.io/2019/6/1/16b10e6615419ca0?imageView2/0/w/1280/h/960/ignore-error/1)

通过 on 方法注册事件，trigger 方法触发事件，来达到事件之间的松散解耦，并且额外添加了 once 和 off 辅助函数用于注册只触发一次的事件以及注销事件

## 29. 实现 JSON.stringify（附加） ##

使用 JSON.stringify 将对象转为 JSON 字符串时，一些非法的数据类型会失真，主要表现如下

* 如果对象含有 toJSON 方法会调用 toJSON
* 在数组中
* * 存在 Undefined/Symbol/Function 数据类型时会变为 null

* * 存在 Infinity/NaN 也会变成 null

* 在对象中
* * 属性值为 Undefined/Symbol/Function 数据类型时，属性和值都不会转为字符串

* * 属性值为 Infinity/NaN ，属性值会变为 null

* 日期数据类型的值会调用 toISOString
* 非数组/对象/函数/日期的复杂数据类型会变成一个空对象
* 循环引用会抛出错误

另外 JSON.stringify 还可以传入第二第三个可选参数，有兴趣的朋友可以深入了解

实现代码较长，这里我直接贴上对应源代码地址 [JSON.stringify]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fyeyan1996%2FJavaScript%2Fblob%2Fmaster%2Fjson.js )

## 源代码 ##

[源代码]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fyeyan1996%2FJavaScript ) ，欢迎 star，以后有后续的技巧会第一时间添加新的内容

## 参考资料 ##

[JavaScript 专题之函数柯里化]( https://juejin.im/post/598d0b7ff265da3e1727c491 )

[JavaScript专题之函数组合]( https://juejin.im/post/59a8ccf8f265da2498241625 )

[JavaScript 专题之函数记忆]( https://juejin.im/post/59af56a96fb9a0248f4aadb8 )

[ES6 系列之私有变量的实现]( https://juejin.im/post/5bf41990e51d4552ee424d2c )

[JavaScript专题之乱序]( https://juejin.im/post/59deda706fb9a045204b3625 )