# 【愣锤笔记】温故而知新--ES6拾遗 #

> 
> 
> 
> ES是一把利器，也是一匹野马。
> 扎实好ES基础，会让人如虎添翼，你还在犹豫什么。
> 
> 

### const和let ###

* 暂时性死区：块级作用域内，在const和let声明变量之前该变量都是不可使用的。不管外部是否用同名变量已经声明。
* 不存在变量提升
* 不允许重复声明
* es5没有块级作用域，es6中使用const或let则自动会形成块级作用域
* 块级作用域可以达到和IIFE相同的效果

` // IIFE 写法 ( function () { var num = ...; ... }()); // 块级作用域写法 { let num = 123; } 复制代码`

* es6规范允许在块级作用域内声明函数，但是客户端（其他环境除外）的实现依旧是类似var的形式。所有考虑到兼容，不要在块级作用域内声明函数，即使需要函数也要写成表达式的形式（ ` let a = function(){}` ）
* es6的块级作用域必须要有大括号，否则引擎不会把它作为块级作用域：

` if ( true ) let a = 1; // 不存在块级作用域 复制代码`

* const常量，指的是常量对应的内存地址不得改变，而不是对应的值不得改变。所有把引用类型的数据设置为常量，其内部的值是可以改变的，因此需要小心：

` const a = {}; a.b = 13; // 没毛病，不会报错 const arr = []; arr.push(123); // 也没毛病 // 如果需要，可以将对象冻结 const a = Object.freeze({}) // 除了对象本身，其属性也应该冻结 var constantize = (obj) => { Object.freeze(obj); Object.keys(obj).forEach( (key, i) => { if ( typeof obj[key] === 'object' ) { constantize( obj[key] ); } }); }; 复制代码`

* es5的var/function声明的全局变量会自动挂载到全局对象中，但是es6的const/let/class声明的全局则不会。

### 解构 ###

数组解构

* 解构不到时变量的值为undefined

` var [a, c] = [2]; a // 2 c // undefined 复制代码`

* 等号右侧只要具有iterator解构的数据都可以进行数组解构
* 解构可以设置默认值，只有严格等于undefined时才会使用默认值:

` let [a, b = 1] = [2, null] // 1, null let [a, b = 1] = [2] // 2 1 let [a, b = 1] = [2, undefined] // 2 1 复制代码`

* 解构可以使用其他变量，但是需要先声明

` let [x = y, y = 1] = []; // 报错，使用y时，y还未声明 let [x = 1, y = x] = []; // x=1; y=1 复制代码`

对象解构

* 数组按照位置解构，对象按照同名属性解构
* 如果等号右侧不是对象或数组，解构时会先将其转换成对象。由于undefined/null无法转换成对象，对其解构时会报错：

` const {a, b} = undefined || []; // 个人推荐这种这种写法避免报错，a undefined, b undefined 复制代码`

解构其他用法

* 解构函数返回的多个值

` // 函数返回多个值只能通过数组或者对象的形式返回 const func = () => [1, {a: 1}, [1,2,3,4]]; const [first, second, third] = func(); // first 1,second {a: 1}, third [1,2,3,4] const func = () => { return { first: 1, second: { a: 1}, third: [1, 2, 3, 4, 5] } }; const {first, second, third} = func(); // first 1, second {a: 1}, third [1, 2, 3, 4, 5] 复制代码`

* 从导入的模块进行解构

` import { moduleA, moduleB } from 'someModule' ; 复制代码`

### 字符串 ###

* 字符串模板中，可以使用变量/函数等

` var func = function () { return '123' } var val = '456' console.log(`this is string ${func()} + ${val} `) // this is string 123 + 456 复制代码`

* 如果变量的不是字符串，会按照一般规则转换成字符串，例如对象会调用其 ` toString` 方法
* 字符串方法

` // 字符串是否包含某个字符串 // 可接受第二个参数表示从下标n开始往后查找 var str = 'abcdefg' ; str.includes( 'c' , 1); // true // 字符开始是否包含某个字符串 // 第二个参数表示从下标n开始是否包含某个字符串 str.startsWith( 'b' ) // false str.startsWith( 'b' , 1) // true // 结尾是否包含 str.endsWith( 'b' ) // false str.endsWith( 'b' , 2) // true ，表示前n个字符的结尾是否包含查询的字符串 str.endsWith( 'b' , 1) // false // 返回一个重复n次的新字符串 str.repeat(2) // 'abcdefgabcdefg' // 补全字符串（返回补全后的新字符串，不改变原字符串） var str = 'a' ; str.padStart(4, 'b' ) // bbba，在前面补 str.padEnd(4, 'b' ) // abbb。 后补 str.padStart(4) // ' a' , 默认补空格 // 消除前/尾空格 (返回新字符，不改变原字符串) var str = ' a ' ; str.trimStart() // "a " str.trimEnd() // " a" 复制代码`

### 数值 ###

* 新增方法，建议放在Number对象上使用（标准也在逐步减少全局方法，使语言逐渐模块化）

` // 判断有限数 Number.isFinite(123) // true Number.isFinite(Infinity) // false Number.isFinite( '123' ) // false ，非数字类型一律返回 false // parseInt和parseFloat移植到了Number对象上 Number.parseInt() Number.parseFloat() // 判断整数 Number.isInteger(2) // true Number.isInteger(2.0) // true js中整数和浮点数采用的是同样的储存方法，所以 2 和 2.0 被视为同一个值 复制代码`
> 
> 
> 
> 
> JavaScript 采用 IEEE 754 标准, 数值存储为64位双精度格式,
> 数值精度最多可以达到 53 个二进制位（1 个隐藏位与 52 个有效位）。
> 如果数值的精度超过这个限度，第54位及后面的位就会被丢弃，这种情况下，Number.isInteger可能会误判。
> 
> 

* js中的最小常量, 表示1与大于1的最小浮点数的差值
` Number.EPSILON` // 2.220446049250313e-16
引入该常量是为浮点数计算，设置一个误差范围：

` 0.1 + 0.2 === 0.3 // false // 可以设计如下函数，误差小于某个差值时自动忽略 const isEqualInAcceptErrorRange = (a, b) => Math.abs(a - b) < Number.EPSILON * Math.pow(2, 2); isEqualInAcceptErrorRange(0.1 + 0.2, 0.3) // true 复制代码`

* Math扩展

` Math.trunc(1.111) //1， 返回去除小数后的值，等于1.111 | 0 Matn.sign(1.111) // +1，判断一个数是正数还是负数,-1负数，0，-0，其他值反水NaN 复制代码`

### 函数 ###

* 函数可以写默认参数

` // 不能有同名函数参数 // 已声明的参数变量在函数体内不可以用 let /const重复声明 // 参数表达式是惰性求值的，每次调用函数时都会重新计算参数的值（也只会在每次调用时才计算参数的值） function func(a = 1, b = 2, c) { } 复制代码`

* 与对象解构一起使用

` // 与对象解构一起使用，但是如果函数调用时没传参数， // 那么从undefined进行解构则会报错 const func = ({ x, y = 1 }) => {}; func({}) // 正常 func() // 报错 // 与对象解构+函数参数默认值一起使用 const func2 = ({ x, y = 1 } = {}) => {}; func2() // 正常 复制代码`

* 一旦使用了函数默认参数，函数的length属性将失真，length本质是函数预期传入的参数
* rest参数

` // 只能使用在最后，后面不能再有参数 // rest参数是一个真正的数组，arguments是类数组 const func = function (...args) { return args // 等同于 // return Array.prototype.slice.call(arguments) } 复制代码`

* 箭头函数返回对象时必须加小括号

` const func5 = () => ({a: 1, b: 2}) 复制代码`

* 箭头函数的this指定义时的上下文中的this(因为箭头函数没有this，所有this自然是其外部的this)，而普通函数中this则指向运行时的上下文
* 箭头函数不可以作为构造函数使用
* 箭头函数无arguments对象，可以哟过rest参数替代
* 不应该使用箭头函数的场景:

` var cat = { count: 0, add: () => { // 由于对象构不成单独的作用域，所以如果写成箭头函数， // 则this在此时指向全局了，而不是该对象 this.count ++ } } cat.add() // cat.count并没有变化 // 事件中需要this指向当前元素时，也不可以使用箭头 button.addEventListener( 'click' , () => { // 得不到期望的结果 this.classList.toggle( 'on' ); }); 复制代码`

* 箭头函数可以嵌套

` // 例如定义一个管道函数，即前一个函数返回的值作为下一个函数的参数 const pipeline = (...funcs) => v => funcs.reduce((a, b) => b(a), v); 复制代码`

### 数组 ###

* 数组扩展运算符只能用在函数中，用于将数组转换成逗号分割的参数序列
* 扩展运算符可以替代apply来进行传递参数

` var arg = [1, 2, 3, 4] func.apply(null, arg) func(...arg) Max.max.apply(null, arg) Math.max(...arg) 复制代码`

* 合并数组

` // 注意是浅拷贝 var arr = [...arr1, ..arr2, ..arr3] 复制代码`

* 所有定义了Iterator接口的数据都可以使用扩展运算符转换成真正的数组

` [...Iterator] 复制代码`

* Array.from()将累数组对象和Iterator接口的数据解构转化成真正的数组

` var divs = document.getElementsByTagName( 'div' ) Array.from(divs) // 接受第二个参数，用于处理每一项数据,类似于map Array.from(arrayLike, x => x * x); 复制代码`

* Array.of()返回参数组成的新数组
* 扩展方法

` // 拷贝数组部分内容覆盖到数组其他位置 // 参数，覆盖开始的位置（含）/拷贝开始的位置（含）/拷贝结束的位置（不含） // 负数表示倒数 [1,2,3,4,5].copyWithin(1, 2, 3) // [1, 3, 3, 4, 5] // 找到某个值，否则返回undefined // 接受第二个参数绑定回调函数的上下文 [1,2,3,4].find((e) => e === 3) // 3 // 找到符合条件的第一个值的下标,否则返回-1 var arr = [1,2,3,4,5] arr.findIndex(e => e > 4) // 4 // 填充数组 // 参数：填充值，填充开始位置，填充结束位置(不含) // 只会填充已有的值 [1,3,4,5,6,7].fill( 'a' , 3, 11111) // [1, 3, 4, "a" , "a" , "a" ] // 遍历键名，arr.keys() var arr = [ 'a' , 'b' , 'c' , 'd' ] for ( let index of arr.keys()) { console.log(index) } // 遍历值，arr.values() var arr = [ 'a' , 'b' , 'c' , 'd' ] for ( let index of arr.values()) { console.log(index) } // 遍历键值对，arr.entries() var arr = [ 'a' , 'b' , 'c' , 'd' ] for ( let index of arr.entries()) { console.log(index) } // 是否包含某个值 [1, 2, 3].includes(2) // true // falt数组，默认flat一层 [1, [2,3], [[4, [5, [6]]]]].flat() // 只有第一层 嵌套的被flat // 等价于 [1, [2,3], [[4, [5, [6]]]]].flat(1) // 接收一个参数表示要flat的层数，如果想flat全部层级 [1, [2,3], [[4, [5, [6]]]]].flat(Infinity) // [1, 2, 3, 4, 5, 6] Infinity表示flat全部层级 // flatMap先对数组每项的值进行回调函数的处理后再flat // 只能拉伸一层 [2, 3, 4].flatMap((x) => [x, x * 2]) // 空位 // 空位的值不是undefined，[,,,,]或者new Array(5)都会产生空位 // es6明确规定空位会转换成undefined 各种方法对空位的处理不一致，所以要绝对避免开发中出现空位 复制代码`

### 对象 ###

* 表达式作为对象的属性名

` var name = 'xiaoming' ; var o = { [ 'liu' + name]: 1 } o.liuxiaoming // 1 复制代码`

* 获取属性的描述信息

` Object.getOwnPropertyDescriptor({a: 1}, 'a' ) 复制代码`

* 对象属性的遍历

` var o = {a: 1} for / in // 自身和继承的除symbols外的所有属性 Object.keys(o) // 自身的除Symbols和不可枚举外的所有属性 Object.getOwnPropertyNames(o) // 自身的除Symbols外所有属性 Object.getOwnPropertySymbols(o) // 所有Symbols属性 Reflect.ownKeys(o) // 自身所有属性 复制代码`

* super关键字

` // es6新增super指向对象的原型对象 // 只能用在对象的方法中，否则报错 var o = {a: 1} var o2 = { a: 2, b () { return super.a } } // 将o2的原型设置为o Object.setPrototypeOf(o2, o) o2.b() // 1 复制代码`

* 对象的解构，用法和数组一样，也是浅拷贝，不能解构undefined/null
* 解构可以解构到原型上继承到值，但是扩展运算符到解构无法解构到原型上继承到值

` var o = Object.create({a: 1}) var {a} = o // a 1 var {...rest} = {} // rest {} 复制代码`

* 对象到方法

` // 是否相等 Object.is(a, b) // 判断两个值是否完全相等，比===强在+0-0，NaN的判断 // 对象浅拷贝 Object.assign(target, o2, o3, o4) // 将o234合并到target对象上，同名属性替换 // 获取对象单个属性的描述对象 Object.getOwnPropertyDescriptor(o2, 'a' ) // 获取对象所有属性的描述对象,注意两者的区别，一个加s，一个没有s Object.getOwnPropertyDescriptors(o2) // 设置对象的原型 var o1 = {} var o2 = {a: 1} Object.setPrototypeOf(o1, o2) // 把o1的原型设置为o2 // 读取对象的原型对象 Object.getPrototypeOf(o1) 复制代码`

### Symbol ###

* Symbol是新增的第七种原始数据类型，凡是属性名为Symbol类型的，可以保证不与其他属性名冲突

` var s1 = Symbol() var s2 = Symbol( 's2' ) // 为s2增加一个描述 s2.description // 获取s2的描述 // 作为属性名使用 var o = { [s2]: 456 } o[s1] = 123 // 不能使用点运算符，否则会认为是字符串 复制代码`

* 

Symbol不会被 ` for...in` 、 ` for...of` 循环，也不会被 ` Object.keys()` 、 ` Object.getOwnPropertyNames()` 、 ` JSON.stringify()` 返回

* 

Symbol.for()/Symbol.keyfor()

` // Symbol.for接收一个参数，每次先搜索有没有存在这个Symbol，有则返回，无则创建再返回 var s1 = Symbol.for( 's1' ) var s2 = Symbol.for( 's1' ) s1 === s2 // Symbol.keyFor(s1)返回一个Symbol.for()登记过的值的key // 参数是一个Symbol.for()注册的数据类型 Symbol.keyFor(s1) // "s1" 复制代码`

* 内置的Symbol值

` // 对象的Symbol.hasInstance指向对象内部一个方法，使用instanceOf判断一个变量是否是该对象的实例时，会自动调用该对象的这个方法去验证 // 比如，foo instanceof Foo在语言内部，实际调用的是Foo[Symbol.hasInstance](foo) class MyClass { [Symbol.hasInstance] (arr) { return arr instanceof Array } } [] instanceof new MyClass() // true 'aa' instanceof new MyClass() // false var o = {} o instanceof new MyClass() // false 复制代码`

### Set ###

` // 类似与不能重复的数组 var s = new Set() var s2 = new Set([1, 2, 3]) s2.add(4) // 利用Set去重 [...new Set([1,2,3,4,5,5,5,5])] // 获取长度 s.size // 删除 s.delete(1) // 删除成功返回 true ，否则返回 false // 清除所有 s.clear() // 检测是否含有 s.has(2) // 返回 true / false // 转换成数组 Array.from(new Set([1,2,3,4])) // Set可以直接使用 for Each方法 // filter/map等只能间接使用 s.forEach((value, key) => {}) s = new Set([...s].filter(e => {})) // 利用Set实现并集 var s1 = new Set([1,2,3]) var s2 = new Set([2,3,4]) new Set([...s1, ...s2]) // 交集 new Set([...s1].filter(e => s2.has(e))) 复制代码`

### Map ###

` // 键值对的集合，类似于对象，但是键不再只是字符串 var m = new Map() var m2 = new Map([[1, 2], [ 'key' , 345]]) // 如果键是基本数据类型，只要严格相等则认为是同一个键 // 如果键是引用类型，只有是同一个引用才判定是同一个键 // 新增数据， set 返回当前map对象，所以支持链式调用 // 有则覆盖，无则生成 m.set(2, 3) m.set([ 'aa' ], 1234) m.set({}, 543).set( 'b' , 123) // 读取键值，无则返回undefined m.get(2) // 获取长度 m.size // 删除 m.delete(键名) // 返回 true / false // 清除所有 m.clear() // 判断是否含有某个键 m.has() // 遍历,遍历顺序就是插入顺序 m.forEach() m.keys() m.values() m.entries() 复制代码`

### Proxy ###

* Proxy代理，给原目标设置一个代理器来控制访问。代理器只对Proxy返回的实例有效，对原目标对象无效

` var obj = {x: 1} var proxyObj = new Proxy(obj, { get () { return 2222; } }) console.log(proxyObj.x); // 2222 // Proxy的实例可以作为其他对象的原型对象使用 var obj2 = Object.create(proxyObj) console.log(obj2.x) // 2222 复制代码`

* 代理器拦截设置（即new Proxy的第二个参数）

` // get拦截对象属性的读取 // 接收三个参数：拦截目标，拦截属性，proxy实例 // 例如，利用proxy生成dom节点的函数 const dom = new Proxy({}, { get (target, key, proxySelf) { return function (attrs = {}, ...children) { let el = document.createElement(key); for ( let i of Object.keys(attrs)) { el.setAttribute(i, attrs[i]) } for ( let child of children) { if (typeof child === 'string' ) { child = document.createTextNode(child) } el.appendChild(child) }; return el; } } }); // 调用生成节点 const el = dom.div({}, 'Hello, my name is ' , dom.a({href: '//example.com' }, 'Mark' ), '. I like:' , dom.ul({}, dom.li({}, 'The web' ), dom.li({}, 'Food' ), dom.li({}, '…actually that\' s it ') ) ); document.body.appendChild(el); // set拦截对象属性的设置 // 接收四个参数：目标对象，设置的属性，属性值，proxy实例 var proxyHandler = { get(target, key, value, proxySelf) { if (key[0] === ' _ ') { throw new Error(`${key}是一个内部属性`) } return target[key] }, set(target, key, value, proxySelf) { if (key[0] === ' _ ') { throw new Error(`${key}是一个内部属性`) } target[key] = value } } var objProxy = new Proxy({}, proxyHandler) objProxy.a = 123 objProxy.a // 123 objProxy._a // 报错内部属性 // has拦截HasProperty操作，例如典型的' a ' in obj，但是不对for/in生效 // 第一个参数源对象，第二个参数属性 // construct 拦截new操作 // 三个参数：原目标对象/函数参数/proxy实例本身 var Person = function() {} var PersonProxy = new Proxy(Person, { construct (target, args, proxySelf) { console.log(' 拦截了new操作 ') return new target(args) } }) new PersonProxy() // deleteProperty拦截delete操作，如果返回false则无法删除 // 参数 target，key // defineProperty拦截Object.defineProperty操作 // 返回false则该操作无效 // 参数： target, key, proxySelf // getOwnPropertyDescriptor拦截Object.getOwnPropertyDescriptor()操作 // 其他拦截方法 - getPrototypeOf - isExtensible - ownKeys - preventExtensions - setPrototypeOf // 取消proxy的代理器，用于在完成代理后，立即收回代理权 // Proxy.revocable返回一个对象，该对象的proxy属性是该代理的实例，revoke是一个可以收回proxy代理权的函数。 var Person = function() {} var {proxy, revoke} = Proxy.revocable(Person, { construct (target, args, proxySelf) { console.log(' 拦截了new操作 ') return new target(args) } }) new proxy() // 正常实例化 revoke() // 取消其proxy的代理权 new proxy() // 无法实例化 // proxy实例之后，实例中的this指代proxy实例 复制代码`

### class类 ###

` const getname = "GET_USER_NAME" ; class Person { // 实例属性也可以写在最顶部，此时不需要加this time = new Date(); // 构造函数里面的内容 constructor (x, y) { this.x = x; this.y = y; console.log(new.target === Person) // } // 类原型上的方法 toString () { console.log(this.x + ' - ' + this.y); } // 静态方法，只有类能访问的，实例不能访问 // 因此这里的x，y都是undefined，因为x,y属性都是实例属性 static toString () { console.log( '静态方法：' , this.x + ' - ' + this.y) } // 可以使用表达式命名 [getName] () { } } const p1 = Person( 'mack' , 25); p1.toString(); // mack - 25 Person.toString() // 静态方法： undefined - undefined 复制代码`

* class中的this指向类的实例
* 类的属性名可以采取表达式
* 可以对属性设置get/set拦截
* 如果不定义constructor函数，则类会默认增加一个constructor函数，并返回类的实例（即this）。如果在constructor中显示返回一个对象，会导致实例结果不再是类的实例。
* es6为new引入了target属性，返回实例的构造函数，如果不是通过new调用的，则返回undefined

` function Person () { // 以前的写法 if (!(this instanceof Person)) { return new Person() } // 现在可以通过new.target if (new.target !== Person) { return new Person() } } 复制代码`

### extends继承 ###

` // 定义父类 class Parent { constructor (name, age) { this.name = name; this.age = age; } static getName () { console.log(this.name); return this.name; } getAge () { console.log( 'age: ' , this.age); } } // 定义子类，子类继承父类 class Child extends Parent { constructor (...arg) { // 通过super继承父类的属性和方法 // 必须通过super继承后，才能使用子类自己的this // 否则报错，得不到子类的this对象 super(...arg); // es5的实现方式Parent.apply(this, arg) // this.name = name; // this.age = age; } } var p1 = new Child( 'xiaoming' , 24); Child.getName() // 父类的静态方法，会被子类继承 p1.getAge() // 父类的原型方法，会被子类的实例继承 复制代码`

* ` Object.getPrototypeOf(Child)` 获取父类，可以用来判断一个类是否继承另一个类
* super作为函数时指代父类的constructor，只能用在constructor中
* super作为对象使用时，在普通方法中指代父类原型，在静态方法中，指代父类。
* 通过super对象调用父类方法时，父类方法中的this，此时指向子类。

### 导入导出 ###

` const a = 5 const b = function (){} const c = class {} // 导出 export const age = '123' export { a, b, c } // 导出多个，建议采取该方式，放在文件底部，一眼就看清楚有哪些导出 export { a as bobyAge } // 导出时定义别名 // 导入 import { a, b, c } from './someJs' // 需要和导出的变量名对应 import { a as otherName } './someJs' // 可以起一个别名 import * as types from './someJs' // 导入所有变量，并起一个types变量名 复制代码`

* 多次导入同一个文件，只会执行一次
* export导出的内容，import时需要加{}
* import在编译时执行，不是在运行时执行，因此不能使用变量
* import存在声明提升，因此最好不要和require时使用还对其有依赖关系

` const a = 123; const f = function () {} const o = {} const C = Class {} // 默认导出 // 一个模块文件只能有一个默认导出 export default a; // 或 export default f // 或 export default { a, f, o, C } // 导入 import newName from './some.js' // 导入时不用和导出的变量名对应，可以随便起一个名字 import _ from 'underscore' ; // 同时导入默认内容和其他内容 import o, {a, b, c} from './some.js' 复制代码`

* 导入导出的复合写法

` // 先导入后导出，此时并没有把a，b导入到当前模块，只相当于对外做了一个转发 export {a, b} from './some.js' // 也可以使用别名 export {a as otherName} from './some.js' 复制代码`

### 结束语 ###

本文内容为温习es6内容，对遗漏知识点进行记录，便于日后翻阅、加深记忆。 内容拜读的阮一峰大大的 [ES6入门]( https://link.juejin.im?target=http%3A%2F%2Fes6.ruanyifeng.com%2F%23README )

> 
> 
> 
> 百尺竿头、日进一步。
> 我是愣锤，一名前端爱好者。
> 欢迎批评与交流。
> 
>