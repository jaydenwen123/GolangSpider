# addEvent.js源码解析 #

![露露](https://user-gold-cdn.xitu.io/2019/6/5/16b252ca49eec58e?imageView2/0/w/1280/h/960/ignore-error/1) 露露

**前言：**
看两三遍即可。

在看 jQuery 源码时，发现了这段注释：

` //源码5235行
/*
* Helper functions for managing events -- not part of the public interface.
* Props to Dean Edwards' addEvent library for many of the ideas.
*/
jQuery. event = { }
复制代码`

Dean Edwards 的 [addEvent.js]( https://link.juejin.im?target=http%3A%2F%2Fdean.edwards.name%2Fweblog%2F2005%2F10%2Fadd-event%2F ) 库为 jQuery 的 **事件绑定** 提供了很多想法，我们就来看下 2005 年的 addEvent.js 。

**〇、先复习下 Object 的内存地址指向的知识**

` let a={ }
let b=a
b[ 0 ]= '111'

console.log(a, 'a55' ) //{0:'111'}
复制代码`

b 改变属性，a 也会改变，因为 b 与 a 指向同一地址（b=a）

**一、addEvent()**
**作用：**
为目标元素绑定事件（如 click）

**源码：**

` //addEvent即为DOM元素绑定事件

// a counter used to create unique IDs
//为每一个事件添加唯一的id
addEvent.guid = 1 ;

function addEvent (element, type, handler) {
// assign each event handler a unique ID
//如果用户自定义的回调函数没有$$guid的话，
//为每个handler添加唯一的id
if (!handler.$$guid) handler.$$guid = addEvent.guid++;
// create a hash table of event types for the element
//为目标元素添加events属性
if (!element.events) element.events = {};
// create a hash table of event handlers for each element/event pair
//events是目标元素绑定的事件的集合
//第一次使用addEvent绑定往往是undefined
let handlers = element.events[type];
if (!handlers) {
handlers = element.events[type] = {};
// store the existing event handler (if there is one)
//handlers的0属性作为目标元素的原生绑定事件
if (element[ "on" + type]) {
handlers[ 0 ] = element[ "on" + type];
}
}
// store the event handler in the hash table
//将每次绑定的事件保存到handler的id属性中
handlers[handler.$$guid] = handler;

// assign a global event handler to do all the work
//为目标元素的原生事件绑定处理程序（handleEvent）
element[ "on" + type] = handleEvent;
}
复制代码`

**解析：**
注意 ` let handlers = element.events[type]` ，handlers 与 element.events[type] 指向同一地址，接下来都是对 handlers 进行操作，实际上也就是对 element.events 进行操作

**关键！**
addEvent 的目的是让目标元素的结构如下：

` element:{
//依次执行了 click 对象里的每个 handler
onclick:handleEvent(MouseEvent)
events:{
click:{
// 如果先onclick了，就会在这里
// 0 : onclick（也就是 function (e) {console. log ( '原生点击了one' )}）
1 : handler（也就是 function () {console. log ( '点击了one' )}  ）
2 : handler （ xxx ）
3 : ...
.. : ...
},
foucus:{

},
..... :{

}

}
}
复制代码`

**注意里面的注释。**
（1）可以看到通过 addEvent 绑定的'click'事件并不是真的绑定在 element 上，而是把绑定的事件处理程序（handler）都放到了 element 的 events 上，即 **绑定事件和目标元素的分离**

（2）由 handleEvent 来 **统一执行 click 事件**

**二、handleEvent()**
**作用：**
执行事件的处理程序

**源码：**

` //执行事件的处理程序
function handleEvent ( event ) {
// grab the event object (IE uses a global event object)
console.log( event , 'event55' )
//event即MouseEvent,原生的鼠标事件集合
event = event || window. event ;
// get a reference to the hash table of event handlers
//找到相同事件的处理程序
//注意这个this，指的是目标元素
//找到handlers集合
let handlers = this.events[ event.type];
// execute each event handler

//依次执行事件的处理程序
for ( let i in handlers) {
this.$$handleEvent = handlers[i];
this.$$handleEvent( event );
}
}
复制代码`

**解析：**
（1）看到上面的 **关键** ，再看 handleEvent 就简单地多，就是把 element->events->click 里的事件处理程序都执行下

（2）主要看for 里面，多了个handlers 和 handleEvent， **我感觉是多余的（不知道下面说的对不对，因为想到了中学解数学题，当你觉得题干中的条件你用不到时，往往是自己想错了）** ，因为直接这么写就好了啊：

**写法一：**

` function handleEvent ( event ) {
event = event || window. event ;
for ( let i in this.events[ event.type]) {
this.events[ event.type][i]( event )
}
}
复制代码`

**写法二：**

` function handleEvent ( event ) {
event = event || window. event ;
let handlers = this.events[ event.type];
for ( let i in handlers) {
handlers[i].call( this , event )
}
}
复制代码`

根据 ` this` 是 ` 谁调用的，this就是谁` 的原则，那么 ` handleEvent` 的 ` this` 是 ` element["on"+type]` ，即 ` element.onclick` 在调用 ` handleEvent` 方法，既然是 ` element` 的属性 ` onclick` 在调用的话，那么执行的上下文就是 ` element` ， ` this` 即 ` element`

Dean Edwards用 ` this.$$handleEvent` 来执行 ` handler` ，目的也是保证正确的作用域，即 ` this`

**三、removeEvent**

` //移除监听事件
function removeEvent ( element, type, handler ) {
// delete the event handler from the hash table
console.log(handler, 'handler52' )
if (element.events && element.events[type]) {
delete element.events[type][handler.$$guid];
}
}
复制代码`

**注意：** 不把自定义事件包成一个变量，是移除不了监听事件的：

` let one= document.querySelector( "#one" )
addEvent(one, 'click' , function ( ) {
console.log( '点击了one' )
})
removeEvent(one, 'click' , function ( ) {
console.log( '点击了one' )
})
复制代码`

写成这样才可以：

` let one= document.querySelector( "#one" )
let handler= function ( ) {
console.log( '点击了one' )
}
addEvent(one, 'click' ,handler)
removeEvent(one, 'click' ,handler)
复制代码`

**四、实验**

` <div id= "one" >干我</div>

let one=document.querySelector( "#one" )

one.onclick= function (e) {
console. log ( '原生点击了one' )
}

addEvent(one, 'click' , function () {
console. log ( '点击了one' )
})
复制代码`

输出： ` 原生点击了one 点击了one`

` <div id= "one" >干我</div>

let one=document.querySelector( "#one" )

addEvent(one, 'click' , function () {
console. log ( '点击了one' )
})

one.onclick= function (e) {
console. log ( '原生点击了one' )
}
复制代码`

输出： ` 原生点击了one` ，原因是后面的 one.onclick 覆盖了 addEvent() 里的事件绑定

当然 jQuery 是都会触发的。

**五、内存泄漏**
简单说：
**element只绑定一次 ` onclick` ，只绑定一次 ` events` 。**
并通过 ` guid` 来为每一个 ` handler` 定一个 ` id` ，然后依次添加进 ` events.click` 中，并通过 ` onclick` 执行

如果一直调用onclick来绑定事件的话，内存开销会很大。

**最后：**
完整代码请看 [github.com/AttackXiaoJ…]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2FAttackXiaoJinJin%2FjQueryExplain%2Fblob%2Fmaster%2FaddEvent.js.html )

![](https://user-gold-cdn.xitu.io/2019/6/3/16b1cc0686322d7c?imageView2/0/w/1280/h/960/ignore-error/1)

（完）