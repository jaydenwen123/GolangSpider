# ES6 完全使用手册 #

## 前言 ##

* 这里的 "ES6" 泛指 ES5 之后的新语法
* 这里的 "完全" 是指本文会不断更新
* 这里的 "使用" 是指本文会展示很多 ES6 的使用场景
* 这里的 "手册" 是指你可以参照本文将项目更多的重构为 ES6 语法

此外还要注意这里不一定就是正式进入规范的语法。

## 1. let 和 const ##

在我们开发的时候，可能认为应该默认使用 let 而不是 var，这种情况下，对于需要写保护的变量要使用 const。

然而另一种做法日益普及：默认使用 const，只有当确实需要改变变量的值的时候才使用 let。这是因为大部分的变量的值在初始化后不应再改变，而预料之外的变量的修改是很多 bug 的源头。

` // 例子 1-1 // bad var foo = 'bar' ; // good let foo = 'bar' ; // better const foo = 'bar' ; 复制代码`

## 2. 模板字符串 ##

### 1. 模板字符串 ###

需要拼接字符串的时候尽量改成使用模板字符串:

` // 例子 2-1 // bad const foo = 'this is a' + example; // good const foo = `this is a ${example} ` ; 复制代码`

### 2. 标签模板 ###

可以借助标签模板优化书写方式:

` // 例子 2-2 let url = oneLine ` www.taobao.com/example/index.html ?foo= ${foo} &bar= ${bar} ` ; console.log(url); // www.taobao.com/example/index.html?foo=foo&bar=bar 复制代码`

oneLine 的源码可以参考 [《ES6 系列之模板字符串》]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fmqyqingfeng%2FBlog%2Fissues%2F84 )

## 3. 箭头函数 ##

优先使用箭头函数，不过以下几种情况避免使用：

### 1. 使用箭头函数定义对象的方法 ###

` // 例子 3-1 // bad let foo = { value : 1 , getValue : () => console.log( this.value) } foo.getValue(); // undefined 复制代码`

### 2. 定义原型方法 ###

` // 例子 3-2 // bad function Foo ( ) { this.value = 1 } Foo.prototype.getValue = () => console.log( this.value) let foo = new Foo() foo.getValue(); // undefined 复制代码`

### 3. 作为事件的回调函数 ###

` // 例子 3-3 // bad const button = document.getElementById( 'myButton' ); button.addEventListener( 'click' , () => { console.log( this === window ); // => true this.innerHTML = 'Clicked button' ; }); 复制代码`

## 4. Symbol ##

### 1. 唯一值 ###

` // 例子 4-1 // bad // 1. 创建的属性会被 for-in 或 Object.keys() 枚举出来 // 2. 一些库可能在将来会使用同样的方式，这会与你的代码发生冲突 if (element.isMoving) { smoothAnimations(element); } element.isMoving = true ; // good if (element.__$jorendorff_animation_library$PLEASE_DO_NOT_USE_THIS_PROPERTY$isMoving__) { smoothAnimations(element); } element.__$jorendorff_animation_library$PLEASE_DO_NOT_USE_THIS_PROPERTY$isMoving__ = true ; // better var isMoving = Symbol ( "isMoving" ); ... if (element[isMoving]) { smoothAnimations(element); } element[isMoving] = true ; 复制代码`

### 2. 魔术字符串 ###

魔术字符串指的是在代码之中多次出现、与代码形成强耦合的某一个具体的字符串或者数值。

魔术字符串不利于修改和维护，风格良好的代码，应该尽量消除魔术字符串，改由含义清晰的变量代替。

` // 例子 4-1 // bad const TYPE_AUDIO = 'AUDIO' const TYPE_VIDEO = 'VIDEO' const TYPE_IMAGE = 'IMAGE' // good const TYPE_AUDIO = Symbol () const TYPE_VIDEO = Symbol () const TYPE_IMAGE = Symbol () function handleFileResource ( resource ) { switch (resource.type) { case TYPE_AUDIO: playAudio(resource) break case TYPE_VIDEO: playVideo(resource) break case TYPE_IMAGE: previewImage(resource) break default : throw new Error ( 'Unknown type of resource' ) } } 复制代码`

### 3. 私有变量 ###

Symbol 也可以用于私有变量的实现。

` // 例子 4-3 const Example = ( function ( ) { var _private = Symbol ( 'private' ); class Example { constructor () { this [_private] = 'private' ; } getName() { return this [_private]; } } return Example; })(); var ex = new Example(); console.log(ex.getName()); // private console.log(ex.name); // undefined 复制代码`

## 5. Set 和 Map ##

### 1. 数组去重 ###

` // 例子 5-1 [...new Set (array)] 复制代码`

### 2. 条件语句的优化 ###

` // 例子 5-2 // 根据颜色找出对应的水果 // bad function test ( color ) { switch (color) { case 'red' : return [ 'apple' , 'strawberry' ]; case 'yellow' : return [ 'banana' , 'pineapple' ]; case 'purple' : return [ 'grape' , 'plum' ]; default : return []; } } test( 'yellow' ); // ['banana', 'pineapple'] 复制代码` ` // good const fruitColor = { red : [ 'apple' , 'strawberry' ], yellow : [ 'banana' , 'pineapple' ], purple : [ 'grape' , 'plum' ] }; function test ( color ) { return fruitColor[color] || []; } 复制代码` ` // better const fruitColor = new Map () .set( 'red' , [ 'apple' , 'strawberry' ]) .set( 'yellow' , [ 'banana' , 'pineapple' ]) .set( 'purple' , [ 'grape' , 'plum' ]); function test ( color ) { return fruitColor.get(color) || []; } 复制代码`

## 6. for of ##

### 1. 遍历范围 ###

for...of 循环可以使用的范围包括：

* 数组
* Set
* Map
* 类数组对象，如 arguments 对象、DOM NodeList 对象
* Generator 对象
* 字符串

### 2. 优势 ###

ES2015 引入了 for..of 循环，它结合了 forEach 的简洁性和中断循环的能力：

` // 例子 6-1 for ( const v of [ 'a' , 'b' , 'c' ]) { console.log(v); } // a b c for ( const [i, v] of [ 'a' , 'b' , 'c' ].entries()) { console.log(i, v); } // 0 "a" // 1 "b" // 2 "c" 复制代码`

### 3. 遍历 Map ###

` // 例子 6-2 let map = new Map (arr); // 遍历 key 值 for ( let key of map.keys()) { console.log(key); } // 遍历 value 值 for ( let value of map.values()) { console.log(value); } // 遍历 key 和 value 值(一) for ( let item of map.entries()) { console.log(item[ 0 ], item[ 1 ]); } // 遍历 key 和 value 值(二) for ( let [key, value] of data) { console.log(key) } 复制代码`

## 7. Promise ##

### 1. 基本示例 ###

` // 例子 7-1 // bad request(url, function ( err, res, body ) { if (err) handleError(err); fs.writeFile( '1.txt' , body, function ( err ) { request(url2, function ( err, res, body ) { if (err) handleError(err) }) }) }); // good request(url) .then( function ( result ) { return writeFileAsynv( '1.txt' , result) }) .then( function ( result ) { return request(url2) }) .catch( function ( e ) { handleError(e) }); 复制代码`

### 2. finally ###

` // 例子 7-2 fetch( 'file.json' ) .then( data => data.json()) .catch( error => console.error(error)) .finally( () => console.log( 'finished' )); 复制代码`

## 8. Async ##

### 1. 代码更加简洁 ###

` // 例子 8-1 // good function fetch ( ) { return ( fetchData() .then( () => { return "done" }); ) } // better async function fetch ( ) { await fetchData() return "done" }; 复制代码` ` // 例子 8-2 // good function fetch ( ) { return fetchData() .then( data => { if (data.moreData) { return fetchAnotherData(data) .then( moreData => { return moreData }) } else { return data } }); } // better async function fetch ( ) { const data = await fetchData() if (data.moreData) { const moreData = await fetchAnotherData(data); return moreData } else { return data } }; 复制代码` ` // 例子 8-3 // good function fetch ( ) { return ( fetchData() .then( value1 => { return fetchMoreData(value1) }) .then( value2 => { return fetchMoreData2(value2) }) ) } // better async function fetch ( ) { const value1 = await fetchData() const value2 = await fetchMoreData(value1) return fetchMoreData2(value2) }; 复制代码`

### 2. 错误处理 ###

` // 例子 8-4 // good function fetch ( ) { try { fetchData() .then( result => { const data = JSON.parse(result) }) .catch( ( err ) => { console.log(err) }) } catch (err) { console.log(err) } } // better async function fetch ( ) { try { const data = JSON.parse( await fetchData()) } catch (err) { console.log(err) } }; 复制代码`

### 3. "async 地狱" ###

` // 例子 8-5 // bad ( async () => { const getList = await getList(); const getAnotherList = await getAnotherList(); })(); // good ( async () => { const listPromise = getList(); const anotherListPromise = getAnotherList(); await listPromise; await anotherListPromise; })(); // good ( async () => { Promise.all([getList(), getAnotherList()]).then(...); })(); 复制代码`

## 9. Class ##

构造函数尽可能使用 Class 的形式

` // 例子 9-1 class Foo { static bar () { this.baz(); } static baz () { console.log( 'hello' ); } baz () { console.log( 'world' ); } } Foo.bar(); // hello 复制代码` ` // 例子 9-2 class Shape { constructor (width, height) { this._width = width; this._height = height; } get area() { return this._width * this._height; } } const square = new Shape( 10 , 10 ); console.log(square.area); // 100 console.log(square._width); // 10 复制代码`

## 10.Decorator ##

### 1. log ###

` // 例子 10-1 class Math { @log add(a, b) { return a + b; } } 复制代码`

log 的实现可以参考 [《ES6 系列之我们来聊聊装饰器》]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fmqyqingfeng%2FBlog%2Fissues%2F109 )

### 2. autobind ###

` // 例子 10-2 class Toggle extends React. Component { @autobind handleClick() { console.log( this ) } render() { return ( < button onClick = {this.handleClick} > button </ button > ); } } 复制代码`

autobind 的实现可以参考 [《ES6 系列之我们来聊聊装饰器》]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fmqyqingfeng%2FBlog%2Fissues%2F109 )

### 3. debounce ###

` // 例子 10-3 class Toggle extends React. Component { @debounce( 500 , true ) handleClick() { console.log( 'toggle' ) } render() { return ( < button onClick = {this.handleClick} > button </ button > ); } } 复制代码`

debounce 的实现可以参考 [《ES6 系列之我们来聊聊装饰器》]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fmqyqingfeng%2FBlog%2Fissues%2F109 )

### 4. React 与 Redux ###

` // 例子 10-4 // good class MyReactComponent extends React. Component {} export default connect(mapStateToProps, mapDispatchToProps)(MyReactComponent); // better @connect(mapStateToProps, mapDispatchToProps) export default class MyReactComponent extends React. Component {}; 复制代码`

## 11. 函数 ##

### 1. 默认值 ###

` // 例子 11-1 // bad function test ( quantity ) { const q = quantity || 1 ; } // good function test ( quantity = 1 ) { ... } 复制代码` ` // 例子 11-2 doSomething({ foo : 'Hello' , bar : 'Hey!' , baz : 42 }); // bad function doSomething ( config ) { const foo = config.foo !== undefined ? config.foo : 'Hi' ; const bar = config.bar !== undefined ? config.bar : 'Yo!' ; const baz = config.baz !== undefined ? config.baz : 13 ; } // good function doSomething ( { foo = 'Hi' , bar = 'Yo!' , baz = 13 } ) { ... } // better function doSomething ( { foo = 'Hi' , bar = 'Yo!' , baz = 13 } = {} ) { ... } 复制代码` ` // 例子 11-3 // bad const Button = ( {className} ) => { const classname = className || 'default-size' ; return < span className = {classname} > </ span > }; // good const Button = ( {className = 'default-size' } ) => ( < span className = {classname} > </ span > ); // better const Button = ( {className} ) => <span className={className}> </ span > } Button.defaultProps = { className : 'default-size' } 复制代码` ` // 例子 11-4 const required = () => { throw new Error ( 'Missing parameter' )}; const add = ( a = required( ), b = required () ) => a + b; add( 1 , 2 ) // 3 add( 1 ); // Error: Missing parameter. 复制代码`

## 12. 拓展运算符 ##

### 1. arguments 转数组 ###

` // 例子 12-1 // bad function sortNumbers ( ) { return Array.prototype.slice.call( arguments ).sort(); } // good const sortNumbers = (...numbers ) => numbers.sort(); 复制代码`

### 2. 调用参数 ###

` // 例子 12-2 // bad Math.max.apply( null , [ 14 , 3 , 77 ]) // good Math.max(...[ 14 , 3 , 77 ]) // 等同于 Math.max( 14 , 3 , 77 ); 复制代码`

### 3. 构建对象 ###

剔除部分属性，将剩下的属性构建一个新的对象

` // 例子 12-3 let [a, b, ...arr] = [ 1 , 2 , 3 , 4 , 5 ]; const { a, b, ...others } = { a : 1 , b : 2 , c : 3 , d : 4 , e : 5 }; 复制代码`

有条件的构建对象

` // 例子 12-4 // bad function pick ( data ) { const { id, name, age} = data const res = { guid : id } if (name) { res.name = name } else if (age) { res.age = age } return res } // good function pick ( {id, name, age} ) { return { guid : id, ...(name && {name}), ...(age && {age}) } } 复制代码`

合并对象

` // 例子 12-5 let obj1 = { a : 1 , b : 2 , c : 3 } let obj2 = { b : 4 , c : 5 , d : 6 } let merged = {...obj1, ...obj2}; 复制代码`

### 4. React ###

将对象全部传入组件

` // 例子 12-6 const parmas = { value1 : 1 , value2 : 2 , value3 : 3 } <Test {...parmas} /> 复制代码`

### 13. 双冒号运算符 ###

` // 例子 13-1 foo::bar; // 等同于 bar.bind(foo); foo::bar(...arguments); // 等同于 bar.apply(foo, arguments ); 复制代码`

如果双冒号左边为空，右边是一个对象的方法，则等于将该方法绑定在该对象上面。

` // 例子 13-2 var method = obj::obj.foo; // 等同于 var method = ::obj.foo; let log = :: console.log; // 等同于 var log = console.log.bind( console ); 复制代码`

## 14. 解构赋值 ##

### 1. 对象的基本解构 ###

` // 例子 14-1 componentWillReceiveProps(newProps) { this.setState({ active : newProps.active }) } componentWillReceiveProps({active}) { this.setState({active}) } 复制代码` ` // 例子 14-2 // bad handleEvent = () => { this.setState({ data : this.state.data.set( "key" , "value" ) }) }; // good handleEvent = () => { this.setState( ( {data} ) => ({ data : data.set( "key" , "value" ) })) }; 复制代码` ` // 例子 14-3 Promise.all([ Promise.resolve( 1 ), Promise.resolve( 2 )]) .then( ( [x, y] ) => { console.log(x, y); }); 复制代码`

### 2. 对象深度解构 ###

` // 例子 14-4 // bad function test ( fruit ) { if (fruit && fruit.name) { console.log (fruit.name); } else { console.log( 'unknown' ); } } // good function test ( {name} = {} ) { console.log (name || 'unknown' ); } 复制代码` ` // 例子 14-5 let obj = { a : { b : { c : 1 } } }; const { a : { b : {c = '' } = '' } = '' } = obj; 复制代码`

### 3. 数组解构 ###

` // 例子 14-6 // bad const spliteLocale = locale.splite( "-" ); const language = spliteLocale[ 0 ]; const country = spliteLocale[ 1 ]; // good const [language, country] = locale.splite( '-' ); 复制代码`

### 4. 变量重命名 ###

` // 例子 14-8 let { foo : baz } = { foo : 'aaa' , bar : 'bbb' }; console.log(baz); // "aaa" 复制代码`

### 5. 仅获取部分属性 ###

` // 例子 14-9 function test ( input ) { return [left, right, top, bottom]; } const [left, __, top] = test(input); function test ( input ) { return { left, right, top, bottom }; } const { left, right } = test(input); 复制代码`

## 15. 增强的对象字面量 ##

` // 例子 15-1 // bad const something = 'y' const x = { something : something } // good const something = 'y' const x = { something }; 复制代码`

动态属性

` // 例子 15-2 const x = { [ 'a' + '_' + 'b' ]: 'z' } console.log(x.a_b); // z 复制代码`

## 16. 数组的拓展方法 ##

### 1. keys ###

` // 例子 16-1 var arr = [ "a" , , "c" ]; var sparseKeys = Object.keys(arr); console.log(sparseKeys); // ['0', '2'] var denseKeys = [...arr.keys()]; console.log(denseKeys); // [0, 1, 2] 复制代码`

### 2. entries ###

` // 例子 16-2 var arr = [ "a" , "b" , "c" ]; var iterator = arr.entries(); for ( let e of iterator) { console.log(e); } 复制代码`

### 3. values ###

` // 例子 16-3 let arr = [ 'w' , 'y' , 'k' , 'o' , 'p' ]; let eArr = arr.values(); for ( let letter of eArr) { console.log(letter); } 复制代码`

### 4. includes ###

` // 例子 16-4 // bad function test ( fruit ) { if (fruit == 'apple' || fruit == 'strawberry' ) { console.log( 'red' ); } } // good function test ( fruit ) { const redFruits = [ 'apple' , 'strawberry' , 'cherry' , 'cranberries' ]; if (redFruits.includes(fruit)) { console.log( 'red' ); } } 复制代码`

### 5. find ###

` // 例子 16-5 var inventory = [ { name : 'apples' , quantity : 2 }, { name : 'bananas' , quantity : 0 }, { name : 'cherries' , quantity : 5 } ]; function findCherries ( fruit ) { return fruit.name === 'cherries' ; } console.log(inventory.find(findCherries)); // { name: 'cherries', quantity: 5 } 复制代码`

### 6. findIndex ###

` // 例子 16-6 function isPrime ( element, index, array ) { var start = 2 ; while (start <= Math.sqrt(element)) { if (element % start++ < 1 ) { return false ; } } return element > 1; } console.log([4, 6, 8, 12].findIndex(isPrime)); // -1, not found console.log([4, 6, 7, 12].findIndex(isPrime)); // 2 复制代码`

更多的就不列举了。

## 17. [optional-chaining]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Ftc39%2Fproposal-optional-chaining ) ##

举个例子：

` // 例子 17-1 const obj = { foo : { bar : { baz : 42 , }, }, }; const baz = obj?.foo?.bar?.baz; // 42 复制代码`

同样支持函数：

` // 例子 17-2 function test ( ) { return 42 ; } test?.(); // 42 exists?.(); // undefined 复制代码`

需要添加 [@babel/plugin-proposal-optional-chaining]( https://link.juejin.im?target=https%3A%2F%2Fbabeljs.io%2Fdocs%2Fen%2Fbabel-plugin-proposal-optional-chaining%23example ) 插件支持

## 18. logical-assignment-operators ##

` // 例子 18-1 a ||= b; obj.a.b ||= c; a &&= b; obj.a.b &&= c; 复制代码`

Babel 编译为：

` var _obj$a, _obj$a2; a || (a = b); (_obj$a = obj.a).b || (_obj$a.b = c); a && (a = b); (_obj$a2 = obj.a).b && (_obj$a2.b = c); 复制代码`

出现的原因：

` // 例子 18-2 function example ( a = b ) { // a 必须是 undefined if (!a) { a = b; } } function numeric ( a = b ) { // a 必须是 null 或者 undefined if (a == null ) { a = b; } } // a 可以是任何 falsy 的值 function example ( a = b ) { // 可以，但是一定会触发 setter a = a || b; // 不会触发 setter，但可能会导致 lint error a || (a = b); // 就有人提出了这种写法： a ||= b; } 复制代码`

需要 [@babel/plugin-proposal-logical-assignment-operators]( https://link.juejin.im?target=https%3A%2F%2Fbabeljs.io%2Fdocs%2Fen%2Fbabel-plugin-proposal-logical-assignment-operators ) 插件支持

## 19. nullish-coalescing-operator ##

` a ?? b // 相当于 (a !== null && a !== void 0 ) ? a : b 复制代码`

举个例子：

` var foo = object.foo ?? "default" ; // 相当于 var foo = (object.foo != null ) ? object.foo : "default" ; 复制代码`

需要 [@babel/plugin-proposal-nullish-coalescing-operator]( https://link.juejin.im?target=https%3A%2F%2Fbabeljs.io%2Fdocs%2Fen%2Fbabel-plugin-proposal-nullish-coalescing-operator ) 插件支持

## 20. pipeline-operator ##

` const double = ( n ) => n * 2 ; const increment = ( n ) => n + 1 ; // 没有用管道操作符 double(increment(double( 5 ))); // 22 // 用上管道操作符之后 5 |> double |> increment |> double; // 22 复制代码`

## 其他 ##

新开了 [知乎专栏]( https://link.juejin.im?target=https%3A%2F%2Fzhuanlan.zhihu.com%2Fc_1042806379215601664 ) ，大家可以在更多的平台上看到我的文章，欢迎关注哦~

## 参考 ##

* [ES6 实践规范]( https://juejin.im/post/5934ff6d2f301e005861422f )
* [babel 7 简单升级指南]( https://juejin.im/post/5b87cab1e51d4538ac05dc54 )
* [不得不知的 ES6 小技巧]( https://link.juejin.im?target=https%3A%2F%2Fmp.weixin.qq.com%2Fs%2FNAbvosbJ4utOgFaM-6wO4Q%3F )
* [深入解析 ES6：Symbol]( https://link.juejin.im?target=http%3A%2F%2Fbubkoo.com%2F2015%2F07%2F24%2Fes6-in-depth-symbols%2F )
* [什么时候你不能使用箭头函数？]( https://link.juejin.im?target=https%3A%2F%2Fzhuanlan.zhihu.com%2Fp%2F26540168 )
* [一些使 JavaScript 更加简洁的小技巧]( https://link.juejin.im?target=https%3A%2F%2Fwww.css88.com%2Farchives%2F9868 )
* [几分钟内提升技能的 8 个 JavaScript 方法]( https://link.juejin.im?target=https%3A%2F%2Fwww.css88.com%2Farchives%2F9916 )
* [[译] 如何使用 JavaScript ES6 有条件地构造对象]( https://juejin.im/post/5bb47db76fb9a05d071953ea )
* [5 个技巧让你更好的编写 JavaScript(ES6) 中条件语句]( https://link.juejin.im?target=https%3A%2F%2Fwww.css88.com%2Farchives%2F9865 )
* [ES6 带来的重大特性 – JavaScript 完全手册（2018版）]( https://link.juejin.im?target=https%3A%2F%2Fwww.css88.com%2Farchives%2F9958 )

## ES6 系列 ##

ES6 系列目录地址： [github.com/mqyqingfeng…]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fmqyqingfeng%2FBlog )

ES6 系列预计写二十篇左右，旨在加深 ES6 部分知识点的理解，重点讲解块级作用域、标签模板、箭头函数、Symbol、Set、Map 以及 Promise 的模拟实现、模块加载方案、异步处理等内容。

如果有错误或者不严谨的地方，请务必给予指正，十分感谢。如果喜欢或者有所启发，欢迎 star，对作者也是一种鼓励。