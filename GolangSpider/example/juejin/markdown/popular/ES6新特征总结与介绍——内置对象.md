# ES6新特征总结与介绍——内置对象 #

### 一、新增 ###

#### （一）Symbol ####

ES6 引入了一种新的原始数据类型 Symbol ，表示独一无二的值，最大的用法是用来定义对象的唯一属性名。

` a = Symbol( 'sss' ) b = Symbol( 'sss' ) a === b // false 复制代码`

##### 1. 作为属性名 #####

* 由于每一个 Symbol 的值都是不相等的，所以 Symbol 作为对象的属性名，可以保证属性不重名。
* Symbol 作为对象属性名时不能用.运算符，要用[]。因为.运算符后面是字符串。
* Symbol 值作为属性名时，该属性是公有属性不是私有属性，可以在类的外部访问。但是不会出现在 for...in 、 for...of 的循环中，也不会被 Object.keys() 、 Object.getOwnPropertyNames() 返回。如果要读取到一个对象的 Symbol 属性，可以通过 Object.getOwnPropertySymbols() 和 Reflect.ownKeys() 取到。

##### 2. 定义常量 #####

##### 3. Symbol.for() #####

返回由给定的 key 找到的 symbol，否则就是返回新创建的 symbol。

##### 4. Symbol.keyFor() #####

Symbol.keyFor() 返回一个已登记的 Symbol 类型值的 key。

` let sym = Symbol.for( "foo" ) Symbol.keyFor(sym) // "foo" 复制代码`

#### （二）Map 与 Set ####

ES6 引入了四种新的原始数据结构。

##### 1. Set #####

ES6 提供了新的数据结构 Set。它类似于数组，但是成员的值都是唯一的，没有重复的值。

* Set 结构的实例有以下属性:

* Set.prototype.size：返回Set实例的成员总数。

* 四个操作方法:

* add(value)：添加某个值，返回 Set 结构本身。
* delete(value)：删除某个值，返回一个布尔值，表示删除是否成功。
* has(value)：返回一个布尔值，表示该值是否为Set的成员。
* clear()：清除所有成员，没有返回值。

* 遍历操作:

* keys()：返回键名的遍历器
* values()：返回键值的遍历器
* entries()：返回键值对的遍历器
* forEach()：使用回调函数遍历每个成员

##### 2. WeakSet #####

首先，WeakSet 的成员只能是对象，而不能是其他类型的值。

其次，WeakSet 中的对象都是弱引用，即垃圾回收机制不考虑WeakSet对该对象的引用，也就是说，如果其他对象都不再引用该对象，那么垃圾回收机制会自动回收该对象所占用的内存。(WeakMap同)

* 没有Size属性
* 操作方法：

* WeakSet.prototype.add(value)：向 WeakSet 实例添加一个新成员。
* WeakSet.prototype.delete(value)：清除 WeakSet 实例的指定成员。
* WeakSet.prototype.has(value)：返回一个布尔值，表示某个值是否在。

* 不可遍历

##### 3. Map #####

它类似于对象，也是键值对的集合，但是“键”的范围不限于字符串，各种类型的值（包括对象）都可以当作键。

* Map 结构的实例有以下属性:

* size 属性：返回 Map 结构的成员总数。

* Map 结构的实例有以下操作方法：

* set(key, value)：set方法设置键名key对应的键值为value，然后返回整个 Map 结构。
* get(key)：get方法读取key对应的键值，如果找不到key，返回undefined。
* delete(value)：删除某个值，返回一个布尔值，表示删除是否成功。
* has(value)：has方法返回一个布尔值，表示某个键是否在当前 Map 对象之中。
* clear()：清除所有成员，没有返回值。

* 遍历操作：

* keys()：返回键名的遍历器。
* values()：返回键值的遍历器。
* entries()：返回所有成员的遍历器。
* forEach()：遍历 Map 的所有成员。

##### 4. WeakMap #####

首先，WeakMap只接受对象作为键名（null除外），不接受其他类型的值作为键名。

其次，WeakMap的键名所指向的对象，不计入垃圾回收机制。

* 没有Size属性
* WeakMap 结构的实例有以下操作方法：

* get(key)
* set(key, value)
* has(value)
* delete(value)

* 不可遍历

##### 5. 应用 #####

Map和Set中对象的引用都是强类型化的，并不会允许垃圾回收。这样一来，如果Map和Set中引用了不再需要的大型对象，如已经从DOM树中删除的DOM元素，那么其回收代价是昂贵的。WeakMap和WeakSet储存 DOM节点，而不用担心这些节点从文档移除时，会引发内存泄漏。

#### （三）Proxy 与 Reflect ####

Proxy 与 Reflect 是 ES6 为了操作对象引入的 API 。

##### 1. Proxy #####

Proxy 可以对目标对象的读取、函数调用等操作进行拦截，然后进行操作处理。它不直接操作对象，而是像代理模式，通过对象的代理对象进行操作，在进行这些操作时，可以添加一些需要的额外操作。

* 一个 Proxy 对象由两个部分组成：target 、 handler 。在通过 Proxy 构造函数生成实例对象时，需要提供这两个参数。 target 即目标对象， handler 是一个对象，声明了代理 target 的指定行为。
` let handle = { get: function (target,key){ console.log( 'setting' + key) return target[key] // 不是target.key }, set : function (target, key, value) { console.log( 'setting ' + key) target[key] = value } } let proxy = new Proxy(target, handler) proxy.name // 实际执行 handler.get proxy.age = 25 // 实际执行 handler.set // getting name // setting age // 25 复制代码` * Proxy 支持的拦截操作一览

* get(target, propKey, receiver)：get方法用于拦截某个属性的读取操作，可以接受三个参数，依次为目标对象、属性名和 proxy 实例本身
* set(target, propKey, value, receiver)：set方法用来拦截某个属性的赋值操作，返回一个布尔值。可以接受四个参数，依次为目标对象、属性名、属性值和 Proxy 实例本身。
* apply方法拦截函数的调用、call和apply操作。apply方法可以接受三个参数，分别是目标对象、目标对象的上下文对象（this）和目标对象的参数数组。
* has(target, propKey)：has方法用来拦截HasProperty操作，即判断对象是否具有某个属性时，这个方法会生效，返回一个布尔值。典型的操作就是in运算符。has方法可以接受两个参数，分别是目标对象、需查询的属性名。
* construct(target, args)：construct方法用于拦截new命令。construct方法可以接受两个参数，分别是目标对象、构造函数的参数对象。
* deleteProperty(target, propKey)：deleteProperty方法用于拦截delete操作，如果这个方法抛出错误或者返回false，当前属性就无法被delete命令删除，返回一个布尔值。
* defineProperty(target, propKey, propDesc)：defineProperty方法拦截了Object.defineProperty操作，返回一个布尔值。
* getOwnPropertyDescriptor(target, propKey)：getOwnPropertyDescriptor方法拦截Object.getOwnPropertyDescriptor()，返回一个属性描述对象或者undefined。
* getPrototypeOf(target)：getPrototypeOf方法主要用来拦截获取对象原型，返回一个对象。
* isExtensible(target)：拦截Object.isExtensible(proxy)，返回一个布尔值。
* ownKeys(target)：ownKeys方法用来拦截对象自身属性的读取操作，返回一个数组。
* preventExtensions(target)：preventExtensions方法拦截Object.preventExtensions()。该方法必须返回一个布尔值，否则会被自动转为布尔值。
* setPrototypeOf(target, proto)：拦截Object.setPrototypeOf(proxy, proto)，返回一个布尔值。

* 目标对象内部的this关键字会指向 Proxy 代理。

##### 2. Reflect #####

* Reflect对象的设计目的有这样几个。

* 将Object对象的一些明显属于语言内部的方法（比如Object.defineProperty），放到Reflect对象上。
* 修改某些Object方法的返回结果，让其变得更合理。比如，Object.defineProperty(obj, name, desc)在无法定义属性时，会抛出一个错误，而Reflect.defineProperty(obj, name, desc)则会返回false。

` try { Object.defineProperty(target, property, attributes) // success } catch (e) { // failure } if (Reflect.defineProperty(target, property, attributes)) { // success } else { // failure } 复制代码`

* 让Object操作都变成函数行为。某些Object操作是命令式，比如name in obj和delete obj[name]，而Reflect.has(obj, name)和Reflect.deleteProperty(obj, name)让它们变成了函数行为。

` // 老写法 'assign' in Object // true // 新写法 Reflect.has(Object, 'assign' ) // true 复制代码`

* Reflect对象的方法与Proxy对象的方法一一对应，只要是Proxy对象的方法，就能在Reflect对象上找到对应的方法。这就让Proxy对象可以方便地调用对应的Reflect方法，完成默认行为，作为修改行为的基础。也就是说，不管Proxy怎么修改默认行为，你总可以在Reflect上获取默认行为。

` Proxy(target, { set : function (target, name, value, receiver) { var success = Reflect.set(target, name, value, receiver) if (success) { console.log( 'property ' + name + ' on ' + target + ' set to ' + value) } return success } }) 复制代码` * Reflect 支持的拦截操作一览

* Reflect.get(target, name, receiver)
* Reflect.set(target, name, value, receiver)
* Reflect.apply(target, thisArg, args)
* Reflect.has(target, name)
* Reflect.construct(target, args)
* Reflect.deleteProperty(target, name)
* Reflect.defineProperty(target, name, desc)
* Reflect.getOwnPropertyDescriptor(target, name)
* Reflect.getPrototypeOf(target)
* Reflect.isExtensible(target)
* Reflect.ownKeys(target)
* Reflect.preventExtensions(target)
* Reflect.setPrototypeOf(target, prototype)

### 二、扩展 ###

#### （一）ES6 字符串 ####

##### 1. 子串的识别 #####

三个方法都可以接受两个参数，需要搜索的字符串，和可选的搜索起始位置索引。

这三个方法只返回布尔值，如果需要知道子串的位置，还是得用 indexOf 和 lastIndexOf 。

* includes()：返回布尔值，判断是否找到参数字符串。
* startsWith()：返回布尔值，判断参数字符串是否在原字符串的头部。
* endsWith()：返回布尔值，判断参数字符串是否在原字符串的尾部。

##### 2. 字符串重复 #####

repeat()：返回新的字符串，表示将字符串重复指定次数返回。

` 'x'.repeat(3) // "xxx" 复制代码`

##### 3. 字符串补全 #####

接受两个参数，第一个参数是指定生成的字符串的最小长度，第二个参数是用来补全的字符串。如果没有指定第二个参数，默认用空格填充。

* padStart：返回新的字符串，表示用参数字符串从头部补全原字符串。
* padEnd：返回新的字符串，表示用参数字符串从头部补全原字符串。
` 'x'.padStart(5, 'ab' ) // 'ababx' 'x'.padStart(4, 'ab' ) // 'abax' 'x'.padEnd(5, 'ab' ) // 'xabab' 'x'.padEnd(4, 'ab' ) // 'xaba' 复制代码`

##### 4. 字符串消除 #####

它们的行为与trim()一致，trimStart()消除字符串头部的空格，trimEnd()消除尾部的空格。它们返回的都是新字符串，不会修改原始字符串。

* trimStart()
* trimEnd()

##### 5. 模板字符串 #####

* 可插入变量和表达式。
* 换行和空格都是会被保留

##### 6. 标签模板 #####

标签模板是一个函数的调用，其中调用的参数是模板字符串。

` alert`Hello world!` // alert( 'Hello world!' ) 复制代码`

#### （二）ES6 数值 ####

##### 1. 二进制和八进制表示法 #####

ES6 提供了二进制和八进制数值的新的写法，分别用前缀0b（或0B）和0o（或0O）表示。

##### 2. Number.isFinite(), Number.isNaN() #####

它们与传统的全局方法isFinite()和isNaN()的区别在于，传统方法先调用Number()将非数值的值转为数值，再进行判断，而这两个新方法只对数值有效，对于非数值一律返回false。

##### 3. Number.parseInt(), Number.parseFloat() #####

ES6 将全局方法parseInt()和parseFloat()，移植到Number对象上面，行为完全保持不变。

##### 4. Number.isInteger() #####

Number.isInteger()用来判断一个数值是否为整数。

##### 5. Number.EPSILON #####

ES6 在Number对象上面，新增一个极小的常量Number.EPSILON。根据规格，它表示 1 与大于 1 的最小浮点数之间的差。

##### 6. 安全整数和 Number.isSafeInteger() #####

JavaScript 能够准确表示的整数范围在-2^53到2^53之间（不含两个端点），超过这个范围，无法精确表示这个值。 Number.isSafeInteger()则是用来判断一个整数是否落在这个范围之内。

##### 7. Math 对象的扩展 #####

ES6 在 Math 对象上新增了 17 个数学相关的静态方法，这些方法只能在 Math 中调用。

* 普通计算

* Math.cbrt：用于计算一个数的立方根。
* Math.imul：两个数以 32 位带符号整数形式相乘的结果，返回的也是一个 32 位的带符号整数。
* Math.hypot：用于计算所有参数的平方和的平方根。
* Math.clz32：用于返回数字的32 位无符号整数形式的前导0的个数。

* 数字处理

* Math.trunc：用于返回数字的整数部分。
* Math.fround：用于获取数字的32位单精度浮点数形式。

* 判断

* Math.sign：判断数字的符号（正、负、0）。
* Math.expm1()：用于计算 e 的 x 次方减 1 的结果，即 Math.exp(x) - 1 。
* Math.log1p(x)：用于计算1 + x 的自然对数，即 Math.log(1 + x) 。

* 双曲函数方法

* Math.sinh(x): 用于计算双曲正弦。
* Math.cosh(x): 用于计算双曲余弦。
* Math.tanh(x): 用于计算双曲正切。
* Math.asinh(x): 用于计算反双曲正弦。
* Math.acosh(x): 用于计算反双曲余弦。
* Math.atanh(x): 用于计算反双曲正切。

* 指数运算符
` 1 ** 2 // 1 2 ** 2 ** 3 // 256 复制代码`

#### （三）ES6 数组 ####

##### 1.数组创建 #####

* 

Array.from()

Array.from方法用于将两类对象转为真正的数组：类似数组的对象（array-like object）和可遍历（iterable）的对象（包括 ES6 新增的数据结构 Set 和 Map）。

` let arrayLike = { '0' : 'a' , '1' : 'b' , '2' : 'c' , length: 3 }; // ES5的写法 var arr1 = [].slice.call(arrayLike) // [ 'a' , 'b' , 'c' ] // ES6的写法 let arr2 = Array.from(arrayLike) // [ 'a' , 'b' , 'c' ] 复制代码` * 

Array.of()

Array.of方法用于将一组值，转换为数组。

` Array.of(3, 11, 8) // [3,11,8] 复制代码`

##### 2.数组查找 #####

find：用于找出第一个符合条件的数组成员，然后返回该成员。如果没有符合条件的成员，则返回undefined。

findIndex：返回第一个符合条件的数组成员的位置，如果所有成员都不符合条件，则返回-1。

` [1, 4, -5, 10].find((n) => n < 0) // -5 [1, 4, -5, 10].findIndex((n) => n < 0) //2 复制代码`

##### 3.数组填充 #####

* 

copyWithin(target, start, end)

数组实例的copyWithin方法，在当前数组内部，将指定位置的成员复制到其他位置（会覆盖原有成员），然后返回当前数组。参数1：被修改的起始索引，参数2：被用来覆盖的数据的起始索引，参数3(可选)：被用来覆盖的数据的结束索引，默认为数组末尾。

` [1, 2, 3, 4, 5].copyWithin(0, 3, 4) // [4, 2, 3, 4, 5] 复制代码` * 

fill(value, start, end)

fill方法使用给定值，填充一个数组。

` [ 'a' , 'b' , 'c' ].fill(7, 1, 2) // [ 'a' , 7, 'c' ] 复制代码`

##### 4.数组遍历 #####

keys()是对键名的遍历、values()是对键值的遍历，entries()是对键值对的遍历。

##### 5.数组包含 #####

Array.prototype.includes方法返回一个布尔值，表示某个数组是否包含给定的值，与字符串的includes方法类似。

` [1, 2, 3].includes(2) // true 复制代码`

##### 6.嵌套数组转一维数组 #####

* 

flat()用于将嵌套的数组“拉平”，变成一维的数组。flat()方法的参数是一个整数，表示想要拉平的层数，默认为1。

* 

flatMap()方法对原数组的每个成员执行一个函数，然后对返回值组成的数组执行flat()方法。

` [1, 2, [3, [4, 5]]].flat(2) // [1, 2, 3, 4, 5] [2, 3, 4].flatMap((x) => [x, x * 2]) // [2, 4, 3, 6, 4, 8] 复制代码`

##### 7. ArrayBuffer #####

#### （四）ES6 对象 ####

##### 1. 对象字面量 #####

* 属性的简洁表示法
` const foo = 'bar' const baz = {foo} baz // {foo: "bar" } // 等同于 const baz = {foo: foo} 复制代码` * 属性名表达式
* 方法的 name 属性

##### 2. 对象的新方法 #####

* 

super 关键字

关键字super，指向当前对象的原型对象。

* 

Object.is(value1, value2)

用来比较两个值是否严格相等，与（===）基本类似。

* 

Object.assign()

用于对象的合并，将源对象（source）的所有可枚举属性，复制到目标对象（target）。

` const target = { a: 1 } const source 1 = { b: 2 } const source 2 = { c: 3 } Object.assign(target, source 1, source 2) target // {a:1, b:2, c:3} 复制代码` * 

Object.getOwnPropertyDescriptors()

返回指定对象所有自身属性（非继承属性）的描述对象。

* 

Object.setPrototypeOf()，Object.getPrototypeOf()

设置或读取一个对象的原型对象。

* 

Object.keys()，Object.values()，Object.entries()

返回一个数组，成员是参数对象自身的（不含继承的）所有可遍历（enumerable）属性的键名/键值/键值对数组

* 

Object.fromEntries()

Object.fromEntries()方法是Object.entries()的逆操作，用于将一个键值对数组转为对象。

#### （五）ES6 正则表达式 ####