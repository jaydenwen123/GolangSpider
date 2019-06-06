# 【拒绝拖延】常见的JavaScript内存泄露原因及解决方案 #

### 前言 ###

内存泄漏指由于疏忽或错误造成程序未能释放已经不再使用的内存。内存泄漏并非指内存在物理上的消失，而是应用程序分配某段内存后，由于设计错误，导致在释放该段内存之前就失去了对该段内存的控制，从而造成了内存的浪费。这里就讲一些常见会带来内存泄露的原因。

### 0. 全局变量 ###

> 
> 
> 
> JavaScript自由的其中一种方式是它可以处理没有声明的变量：一个未声明的变量的引用在全局对象中创建了一个新变量。在浏览器的环境中，全局对象是window。
> 
> 
> 

` function foo ( ) { name = '前端曰' ； } // 其实是把name变量挂载在window对象上 function foo ( ) { window.name = '前端曰' ； } // 又或者 function foo ( ) { this.name = '前端曰' ； } foo() // 其实这里的this就是指向的window对象 复制代码`

这样无意中一个意外的全局变量就被创建了，为了阻止这种错误发生，在你的Javascript文件最前面添加 ` 'use strict;'` 。这开启了解析JavaScript的阻止意外全局的更严格的模式。或者自己注意好变量的定义！

### 1. 循环引用 ###

在js的内存管理环境中，对象 A 如果有访问对象 B 的权限，叫做对象 A 引用对象 B。引用计数的策略是将“对象是否不再需要”简化成“对象有没有其他对象引用到它”，如果没有对象引用这个对象，那么这个对象将会被回收 。

` function func ( ) { let obj1 = {}; let obj2 = {}; obj1.a = obj2; // obj1 引用 obj2 obj2.a = obj1; // obj2 引用 obj1 } 复制代码`

当函数 func 执行结束后，返回值为 undefined，所以整个函数以及内部的变量都应该被回收，但根据引用计数方法，obj1 和 obj2 的引用次数都不为 0，所以他们不会被回收。要解决循环引用的问题，最好是在不使用它们的时候手工将它们设为空。

解决方案： ` obj1` 和 ` obj2` 都设为 ` null` 。

### 2. 老生常谈的闭包 ###

闭包：匿名函数可以访问父级作用域的变量。

` var names = ( function ( ) { var name = 'js-say' ; return function ( ) { console.log(name); } })() 复制代码`

闭包会造成对象引用的生命周期脱离当前函数的上下文，如果闭包如果使用不当，可以导致环形引用（circular reference），类似于死锁，只能避免，无法发生之后解决，即使有垃圾回收也还是会内存泄露。

### 3. 被遗忘的延时器/定时器 ###

在我们的日常需求中，可能会经常试用到 ` setInterval/setTimeout` ，但是使用完之后通常忘记清理。

` var someResource = getData(); setInterval( function ( ) { var node = document.getElementById( 'Node' ); if (node) { // 处理 node 和 someResource node.innerHTML = JSON.stringify(someResource)); } }, 1000 ); 复制代码`

` setInterval/setTimeout` 中的 ` this` 指向的是window对象，所以内部定义的变量也挂载到了全局； ` if` 内引用了 ` someResource` 变量，如果没有清除 ` setInterval/setTimeout` 的话 ` someResource` 也得不到释放；同理其实 ` setTimeout` 也一样。所以我们用完需要记得去 ` clearInterval/clearTimeout` 。

### 4. DOM引起的内存泄露 ###

* 未清除DOM引用

` var refA = document.getElementById( 'refA' ); document.body.removeChild(refA); // #refA不能回收，因为存在变量refA对它的引用。将其对#refA引用释放，但还是无法回收#refA。 复制代码`

解决方案： ` refA = null` 。

* DOM对象添加的属性是一个对象的引用

` var MyObject = {}; document.getElementById( 'myDiv' ).myProp = MyObject; 复制代码`

解决方案：在页面 ` onunload` 事件中释放 ` document.getElementById('myDiv').myProp = null;` 。

DOM被删除或清空没有清楚绑定事件这种情况应该是比较常见的，同时也应该是比较容易被忽略的。

* 给DOM对象绑定事件

` var btn = document.getElementById( "myBtn" ); btn.onclick = function ( ) { document.getElementById( "myDiv" ).innerHTML = "wechat: js-say" ; } document.body.removeChild(btn); btn = null ; 复制代码`

这里把DOM移除了，但是绑定的事件仍没被移除，会引起内存泄露所以需要清除事件。

` var btn = document.getElementById( "myBtn" ); btn.onclick = function ( ) { btn.onclick = null ; document.getElementById( "myDiv" ).innerHTML = "wechat: js-say" ; } document.body.removeChild(btn); btn = null ; 复制代码`

### 小广告 ###

**我自己运营的公众号，记录我自己的成长！**

每周至少一篇博客，拒绝拖延从我做起！

公众号：前端曰

公众号ID：js-say

ps：是(yue)不是(ri)

![](https://user-gold-cdn.xitu.io/2019/5/23/16ae4b504a87a348?imageView2/0/w/1280/h/960/ignore-error/1)