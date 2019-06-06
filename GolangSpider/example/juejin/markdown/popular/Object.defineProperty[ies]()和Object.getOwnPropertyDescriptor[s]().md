# Object.defineProperty[ies]()和Object.getOwnPropertyDescriptor[s]() #

> 
> 
> 
> Object的defineProperty和defineProperties这两个方法在js中的重要性十分重要，主要功能就是用来定义或修改这些内部属性,与之相对应的getOwnPropertyDescriptor和getOwnPropertyDescriptors就是获取这行内部属性的描述。
> 
> 
> 

> 
> 
> 
> 下面文章我先介绍数据描述符和存取描述符的属性代表的含义，然后简单介绍以上四个方法的基本功能，这些如果了解可直接跳过，最后我会举例扩展及说明各内部属性在各种场景下产生的实际效果，那才是这篇文章的核心内容。本文章关于概念性的描述还是会尽量使用《javaScript高级教程》、MDN网站等概念，保证准确和易于大家理解，讲解部分则结合个人理解和举例说明。
> 
> 
> 

### 数据(数据描述符)属性 ###

> 
> 
> 
> 数据属性有4个描述内部属性的特性
> 
> 
> 
> ##### [[Configurable]] #####
> 
> 
> 
> 表示能否通过 [delete](
> https://link.juejin.im?target=https%3A%2F%2Fsegmentfault.com%2Fa%2F1190000010574280
> ) 删除此属性，能否修改属性的特性，或能否修改把属性修改为访问器属性，如果直接使用字面量定义对象，默认值为true。
> 
> 
> 
> ##### [[Enumerable]] #####
> 
> 
> 
> 表示该属性是否可枚举，即是否通过for-in循环或Object.keys()返回属性，如果直接使用字面量定义对象，默认值为true
> 
> 
> 
> ##### [[Writable]] #####
> 
> 
> 
> 能否修改属性的值，如果直接使用字面量定义对象，默认值为true
> 
> 
> 
> ##### [[Value]] #####
> 
> 
> 
> 该属性对应的值，默认为undefined
> 
> 

### 访问器(存取描述符)属性 ###

> 
> 
> 
> 访问器属性也有4个描述内部属性的特性
> 
> 
> 
> ##### [[Configurable]] #####
> 
> 
> 
> 和数据属性的[[Configurable]]一样，表示能否通过delete删除此属性，能否修改属性的特性，或能否修改把属性修改为访问器属性，如果直接使用字面量定义对象，默认值为true
> 
> 
> 
> 
> ##### [[Enumerable]] #####
> 
> 
> 
> 和数据属性的[[Configurable]]一样，表示该属性是否可枚举，即是否通过for-in循环或Object.keys()返回属性，如果直接使用字面量定义对象，默认值为true
> 
> 
> 
> 
> ##### [[GET]] #####
> 
> 
> 
> 一个给属性提供 getter 的方法(访问对象属性时调用的函数,返回值就是当前属性的值)，如果没有 getter 则为
> undefined。该方法返回值被用作属性值。默认为 undefined
> 
> 
> 
> ##### [[SET]] #####
> 
> 
> 
> 一个给属性提供 setter 的方法(给对象属性设置值时调用的函数)，如果没有 setter 则为
> undefined。该方法将接受唯一参数，并将该参数的新值分配给该属性。默认为 undefined
> 
> 

### 创建/修改/获取属性的方法 ###

##### YI、Object.defineProperty() #####

> 
> 
> 
> 功能： 方法会直接在一个对象上定义一个新属性，或者修改一个对象的现有属性， 并返回这个对象。如果不指定configurable, writable,
> enumerable ，则这些属性默认值为false，如果不指定value, get, set，则这些属性默认值为undefined ` 语法:
> Object.defineProperty(obj, prop, descriptor)` obj: 需要被操作的目标对象 prop:
> 目标对象需要定义或修改的属性的名称 descriptor: 将被定义或修改的属性的描述符
> 
> ` var obj = new Object(); Object.defineProperty(obj, 'name', {
> configurable: false, writable: true, enumerable: true, value: '张三' })
> console.log(obj.name) //张三 复制代码`

##### ER、Object.defineProperties() #####

> 
> 
> 
> 功能： 方法直接在一个对象上定义一个或多个新的属性或修改现有属性，并返回该对象。 `
> 语法：Object.defineProperties(obj,props)` obj : 将要被添加属性或修改属性的对象。 props :
> 该对象的一个或多个键值对定义了将要为对象添加或修改的属性的具体配置
> 
> ` var obj = new Object(); Object.defineProperties(obj, { name: { value:
> '张三', configurable: false, writable: true, enumerable: true }, age: {
> value: 18, configurable: true } }) console.log(obj.name, obj.age) // 张三,
> 18 复制代码`

##### SAN、Object.getOwnPropertyDescriptor() #####

> 
> 
> 
> 功能： 该方法返回指定对象上的一个自有属性对应的属性描述（自有属性指的是直接赋予该对象的属性，不需要从原型链上进行查找的属性） ` 语法:
> Object.getOwnPropertyDescriptor(obj, prop)` //obj: 需要查找的目标对象, prop:
> 目标对象内属性名称
> 
> ` var person = { name: '张三', age: 18 } var desc =
> Object.getOwnPropertyDescriptor(person, 'name'); console.log(desc) 结果如下 //
> { // configurable: true, // enumerable: true, // writable: true, // value:
> "张三" // } 复制代码`

### SI、Object. getOwnPropertyDescriptors() ###

> 
> 
> 
> 功能：所指定对象的所有自身属性的描述符，如果没有任何自身属性，则返回空对象。 ` 语法：
> Object.getOwnPropertyDescriptors(obj)` //obj: 需要查找的目标对象
> 
> ` var person = { name: '张三', age: 18 } var desc =
> Object.getOwnPropertyDescriptors(person); console.log(desc) // age: { //
> value: 18, writable: true, enumerable: true, configurable: true} // name:
> { // value: "张三", writable: true, enumerable: true, configurable: true} //
> __proto__: Object 复制代码`

### 各种场景下描述符属性的扩展示例讲解 ###

#####.configurable #####

> 
> 
> 
> 如果设置configurable属性为false，则不可使用delete操作符（在严格模式下抛出错误）， ` 修改所有内部属性值会抛出错误`
> 
> 
> 
> ###### 在对象中添加一个数据描述符属性 ######
> 
> ` var person = {}; Object.defineProperty(person, 'name', { configurable:
> false, value: 'John' }) ; delete person.name // 严格模式下抛出错误
> console.log(person.name) // 'John' 没有删除 Object.defineProperty(person,
> 'name', { configurable: true //报错 }); Object.defineProperty(person,
> 'name', { enumerable: 2 //报错 }); Object.defineProperty(person, 'name', {
> writable: true //报错 }); Object.defineProperty(person, 'name', { value: 2
> //报错 }); 复制代码`
> 
> 注意：以上是·最开始定义属性描述符·时,writabl默认为false,才会出现上述效果,如果writable定义为true,
> 则可以修改[[writable]]和[[value]]属性值,修改另外两个属性值报错：
> 
> ` var obj = {}; Object.defineProperty(obj, 'a', { configurable: false,
> writable: true, value: 1 }); Object.defineProperty(obj, 'a', { //
> configurable: true, //报错 // enumerable: true, //报错 writable: false, value:
> 2 }); var d = Object.getOwnPropertyDescriptor(obj, 'a') console.log(d); //
> { // value: 2, // writable: false, // } 复制代码`
> 
> ###### 在对象中添加存取描述符属性 ######
> ` var obj = {}; var aValue; //如果不初始化变量, 不给下面的a属性设置值,直接读取会报错aValue is not
> defined var b; Object.defineProperty(obj, 'a', { configurable : true,
> enumerable : true, get: function() { return aValue }, set:
> function(newValue) { aValue = newValue; b = newValue + 1 } })
> console.log(b) // undefined console.log(obj.a) // undefined,
> 当读取属性值时，调用get方法,返回undefined obj.a = 2; // 当设置属性值时,调用set方法,aValue为2
> console.log(obj.a) // 2 读取属性值,调用get方法,此时aValue为2 console.log(b) // 3
> 再给obj.a赋值时,执行set方法,b的值被修改为2,额外说一句,vue中的计算属性就是利用setter来实现的 复制代码`
> 
> 注意:
> 
> 
> 
> * getter和setter可以不同时使用，但在严格模式下只使用一个，会抛出错误。
> * 数据描述符与存取描述符不可以混用，否则会抛出错误。
> * 使用 var定义的任何变量，其 ` configurable` 属性值都为 ` false` ,定义对象也是一样.
> 
> 
> 

##### Writable #####

> 
> 
> 
> 当Writable为false（并且configurable为true）,[value]可以通过
> defineProperty修改，但不能直接赋值修改。
> 
> ` var obj = {}; Object.defineProperty(obj, 'a', { configurable: true,
> enumerable: false, writable: false, value: 1 });
> Object.defineProperty(obj, 'a', { configurable: false, enumerable: true,
> writable: false , value: 2 }); var d =
> Object.getOwnPropertyDescriptor(obj, 'a') console.log(d); // 结果如下 // { //
> value: 2, // writable: false, // enumerable: true, // configurable: false
> // } 但是如果直接复制修改 var obj = {} Object.defineProperty(obj, 'a', {
> configurable: true, enumerable: false, writable: false, value: 1 });
> obj.a=2; var d = Object.getOwnPropertyDescriptor(obj, 'a') console.log(d);
> // 结果如下 // { // value: 1, // 没有做出修改 // writable: false, // enumerable:
> true, // configurable: false // } 复制代码`

##### Enumerable #####

> 
> 
> 
> 直接上例子
> 
> ` var obj = {}; Object.defineProperties(obj, { a: { value: 1, enumerable:
> false }, b: { value: 2, enumerable: true }, c: { value: 3, enumerable:
> false } }) obj.d = 4; //等同于 //Object.defineProperty(obj, 'd', { //
> configurable: true, // enumerable: true, // writable: true, // value: 4
> //}) for(var key in obj) { console.log(key); // 打印一次b, 一次d,
> a和c属性enumerable为false，不可被枚举 } var arr = Object.keys(obj);
> console.log(arr); // ['b', 'd'] 复制代码`

##### get 和 set #####

> 
> 
> 
> 简易的数据双向绑定
> 
> ` //html <body> <p> input1=> <input type="text" id="input1"> </p> <p>
> input2=> <input type="text" id="input2"> </p> <div> 我每次比input1的值加1=> <span
> id="span"></span> </div> </body> //js var oInput1 =
> document.getElementById('input1'); var oInput2 =
> document.getElementById('input2'); var oSpan =
> document.getElementById('span'); var obj = {};
> Object.defineProperties(obj, { val1: { configurable: true, get: function()
> { oInput1.value = 0; oInput2.value = 0; oSpan.innerHTML = 0; return 0 },
> set: function(newValue) { oInput2.value = newValue; oSpan.innerHTML =
> Number(newValue) ? Number(newValue) : 0 } }, val2: { configurable: true,
> get: function() { oInput1.value = 0; oInput2.value = 0; oSpan.innerHTML =
> 0; return 0 }, set: function(newValue) { oInput1.value = newValue;
> oSpan.innerHTML = Number(newValue)+1; } } }) oInput1.value = obj.val1;
> oInput1.addEventListener('keyup', function() { obj.val1 = oInput1.value;
> }, false) oInput2.addEventListener('keyup', function() { obj.val2 =
> oInput2.value; }, false) 复制代码`

###### 原文链接地址: [segmentfault.com/a/119000001…]( https://link.juejin.im?target=https%3A%2F%2Fsegmentfault.com%2Fa%2F1190000011294519%3Futm_source%3Dtag-newest%23articleHeader15 ) ######