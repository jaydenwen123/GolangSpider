# 前端进阶之 let、const、var #

> 
> 
> 
> * 作者：陈大鱼头
> * github： [KRISACHAN](
> https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2FKRISACHAN )
> * 链接： [github.com/YvetteLau/S…](
> https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2FYvetteLau%2FStep-By-Step%2Fissues%2F16%23issuecomment-498046553
> )
> * 背景：最近高级前端工程师 **刘小夕** 在 **github** 上开了个每个工作日布一个前端相关题的 **repo** ，怀着学习的心态我也参与其中，以下为我的回答，如果有不对的地方，非常欢迎各位指出。
> 
> 
> 
> 

## 什么是提升？（hoisting） ##

### 函数提升与变量提升 ###

首先我们来看一段代码

` console.log(变量) // undefined var 变量 = 1 复制代码`

从上面的代码来看，虽然变量还没有被声明，但是我们却可以使用，这种情况就叫做变量提升。

再来一段代码

` console.log(变量) // ƒ 变量() {} function 变量( ) {} var 变量 = 1 复制代码`

上面的代码叫做函数提升，函数提升跟变量提升差不多，就是函数提升优先级比变量高。

**从上可知，使用 ` var` 声明的变量会被提升到作用域的顶部。**

## let、const、var的区别 ##

### let、const、var的提升 ###

首先我们再来看一段代码

` var a = 1 let b = 1 const c = 1 console.log( window.b) // undefined console.log( window.c) // undefined function test ( ) { console.log(a) let a } test() 复制代码`

首先在全局作用域下使用 ` let` 和 ` const` 声明变量，变量并不会被挂载到 ` window` 上，这一点就和 ` var` 声明有了区别。

再者当我们在声明 ` a` 之前如果使用了 ` a` ，就会出现报错的情况。

首先报错的原因是因为存在暂时性死区，我们不能在声明前就使用变量，这也是 ` let` 和 ` const` 优于 ` var` 的一点。然后这里你认为的提升和 ` var` 的提升是有区别的，虽然变量在编译的环节中被告知在这块作用域中可以访问，但是访问是受限制的。

## let、const、var 创建的不同 ##

**` let`** 和 **` const`** 声明定义了作用于 **正在运行的执行上下文（running execution context）** 的 **词法环境（LexicalEnvironment）** 的变量。

**` let`** 和 **` const`** 声明的变量是在词法环境实例化时创建的，但是给变量赋值的原生功能 **` LexicalBinding`** 以及变量初始化的功能 **` Initializer`** 是在之后执行的，而不是在创建变量时，所以在执行之前无法以任何方式访问它们，这就是 **暂时性死区** 。

var语句声明了作用于 **正在运行的执行上下文（running execution context）** 的 **变量环境（VariableEnvironment）** 的变量。

**` var`** 声明的变量同样是在词法环境实例化时创建的，并且创建时就赋值有 **` undefined`** ，在任何的 **变量环境（VariableEnvironment） **中，** ` var` 变量** 的绑定可以出现多个，但是最终值是由 赋值时确定的，而不是创建变量时。

### 词法环境（LexicalEnvironment）与变量环境（VariableEnvironment） ###

* **词法环境（LexicalEnvironment）** ：简单来说就是 **ECMASCRIPT** 中的书写语法的上下文语法环境。
* **变量环境（VariableEnvironment）** ：简单来说就是 **执行上下文** 中专门存储变量跟该变量赋值的一个对象。

**总结： ` let` ` const` 跟 ` var` 不同的原因是 ` let` ` const` 的创建是基于词法环境，而 ` var` 是基于变量环境。。用通俗的话来说就是，不是同一个系统的...**

如果你、喜欢探讨技术，或者对本文有任何的意见或建议，你可以扫描下方二维码，关注微信公众号“ **鱼头的Web海洋** ”，随时与鱼头互动。欢迎！衷心希望可以遇见你。

![](https://user-gold-cdn.xitu.io/2019/5/18/16aca8a03abb7c13?imageView2/0/w/1280/h/960/ignore-error/1)