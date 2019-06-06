# 近一万字的ES6语法知识点补充 #

![](https://user-gold-cdn.xitu.io/2019/2/12/168df9eaae700fc2?imageView2/0/w/1280/h/960/ignore-error/1)

# 前言 #

ECMAScript 6.0（简称ES6），作为下一代JavaScript的语言标准正式发布于2015 年 6 月，至今已经发布3年多了，但是因为蕴含的语法之广，完全消化需要一定的时间，这里我总结了部分ES6，以及ES6以后新语法的知识点，使用场景，希望对各位有所帮助

**本文讲着重是对ES6语法特性的补充，不会讲解一些API层面的语法，更多的是发掘背后的原理，以及ES6到底解决了什么问题**

**如有错误,欢迎指出,将在第一时间修改,欢迎提出修改意见和建议**

话不多说开始ES6之旅吧~~~

# let/const（常用） #

let,const用于声明变量，用来替代老语法的var关键字，与var不同的是，let/const会创建一个块级作用域（通俗讲就是一个花括号内是一个新的作用域）

这里外部的console.log(x)拿不到前面2个块级作用域声明的let:

![](https://user-gold-cdn.xitu.io/2019/2/12/168df9eaaef8cc17?imageView2/0/w/1280/h/960/ignore-error/1)

在日常开发中多存在于使用 **if/for** 关键字结合let/const创建的块级作用域， **值得注意的是使用let/const关键字声明变量的for循环和var声明的有些不同**

![](https://user-gold-cdn.xitu.io/2019/2/12/168df9eaaed7a489?imageView2/0/w/1280/h/960/ignore-error/1)

for循环分为3部分，第一部分包含一个变量声明，第二部分包含一个循环的退出条件，第三部分包含每次循环最后要执行的表达式，也就是说第一部分在这个for循环中只会执行一次var i = 0，而后面的两个部分在每次循环的时候都会执行一遍

![](https://user-gold-cdn.xitu.io/2019/2/12/168df9eab004895e?imageView2/0/w/1280/h/960/ignore-error/1)

而使用使用let/const关键字声明变量的for循环，除了会创建块级作用域，let/const还会将它绑定到每个循环中，确保对上个循环结束时候的值进行重新赋值

什么意思呢？简而言之就是每次循环都会声明一次（对比var声明的for循环只会声明一次），可以这么理解let/const中的for循环

给每次循环创建一个块级作用域:

![](https://user-gold-cdn.xitu.io/2019/2/12/168df9eab036229e?imageView2/0/w/1280/h/960/ignore-error/1)

### 暂时性死区 ###

使用let/const声明的变量，从一开始就形成了封闭作用域，在声明变量之前是无法使用这个变量的，这个特点也是为了弥补var的缺陷（var声明的变量有变量提升）

![](https://user-gold-cdn.xitu.io/2019/2/19/169065634290c00f?imageView2/0/w/1280/h/960/ignore-error/1)

剖析暂时性死区的原理， **其实let/const同样也有提升的作用** ，但是和var的区别在于

* 

var在创建时就被初始化，并且赋值为undefined

* 

let/const在进入块级作用域后，会因为提升的原因先创建，但不会被初始化，直到声明语句执行的时候才被初始化，初始化的时候如果使用let声明的变量没有赋值，则会默认赋值为undefined，而const必须在初始化的时候赋值。而创建到初始化之间的代码片段就形成了暂时性死区

引用 [一篇博客]( https://link.juejin.im?target=https%3A%2F%2Fsegmentfault.com%2Fa%2F1190000008213835 ) 对于ES6标准翻译出来的一段话

> 
> 
> 
> 由let/const声明的变量，当它们包含的词法环境(Lexical
> Environment)被实例化时会被创建，但只有在变量的词法绑定(LexicalBinding)已经被求值运算后，才能够被访问
> 
> 

回到例子，这里因为使用了let声明了变量name,在代码执行到if语句的时候会先进入预编译阶段，依次创建块级作用域,词法环境,name变量（没有初始化）,随后进入代码执行阶段，只有在运行到let name语句的时候变量才被初始化并且默认赋值为undefined，但是因为暂时性死区导致在运行到声明语句 **之前** 使用到了name变量，所以报错了

![](https://user-gold-cdn.xitu.io/2019/2/12/168e0aa07109eb76?imageView2/0/w/1280/h/960/ignore-error/1)

上面这个例子,因为使用var声明变量,会有变量提升,同样也是发生在预编译阶段,var会提升到当前函数作用域的顶部并且默认赋值为undefined,如果这几行代码是在全局作用域下,则name变量会直接提升到全局作用域,随后进入执行阶段执行代码,name被赋值为"abc",并且可以成功打印出字符串abc

相当于这样

![](https://user-gold-cdn.xitu.io/2019/2/12/168e0abf693ab68b?imageView2/0/w/1280/h/960/ignore-error/1)

暂时性死区其实是为了防止ES5以前在变量声明前就使用这个变量,这是因为var的变量提升的特性导致一些不熟悉var原理的开发者习以为常的以为变量可以先使用在声明,从而埋下一些隐患

关于JS预编译和JS的3种作用域(全局,函数,块级)这里也不赘述了,否则又能写出几千字的博客,有兴趣的朋友自行了解一下,同样也有助于了解JavaScript这门语言

### const ###

使用const关键字声明一个常量，常量的意思是不会改变的变量，const和let的一些区别是

* const声明变量的时候必须赋值，否则会报错，同样使用const声明的变量被修改了也会报错

![](https://user-gold-cdn.xitu.io/2019/2/12/168df9ead3323a16?imageView2/0/w/1280/h/960/ignore-error/1)

* const声明变量不能改变，如果声明的是一个引用类型，则不能改变它的内存地址（这里牵扯到JS引用类型的特点，有兴趣可以看我另一篇博客 [对象深拷贝和浅拷贝]( https://juejin.im/post/5c26dd8fe51d4570c053e08b ) ）

![](https://user-gold-cdn.xitu.io/2019/2/12/168df9eaff4bf8e6?imageView2/0/w/1280/h/960/ignore-error/1)

有些人会有疑问，为什么日常开发中没有显式的声明块级作用域，let/const声明的变量却没有变为全局变量

![](https://user-gold-cdn.xitu.io/2019/2/12/168df9eb0879648b?imageView2/0/w/1280/h/960/ignore-error/1)

这个其实也是let/const的特点，ES6规定它们不属于顶层全局变量的属性，这里用chrome调试一下

![](https://user-gold-cdn.xitu.io/2019/2/12/168df9eb0a403854?imageView2/0/w/1280/h/960/ignore-error/1)

可以看到使用let声明的变量x是在一个叫script作用域下的，而var声明的变量因为变量提升所以提升到了全局变量window对象中，这使我们能放心的使用新语法，不用担心污染全局的window对象

### 建议 ###

在日常开发中，我的建议是全面拥抱let/const，一般的变量声明使用let关键字，而当声明一些配置项（类似接口地址，npm依赖包，分页器默认页数等一些一旦声明后就不会改变的变量）的时候可以使用const，来显式的告诉项目其他开发者，这个变量是不能改变的(const声明的常量建议使用全大写字母标识,单词间用下划线)，同时也建议了解var关键字的缺陷（变量提升，污染全局变量等），这样才能更好的使用新语法

# 箭头函数（常用） #

ES6 允许使用 **箭头** （=>）定义函数

箭头函数对于使用function关键字创建的函数有以下区别

* 

箭头函数没有arguments（建议使用更好的语法，剩余运算符替代）

* 

箭头函数没有prototype属性，不能用作构造函数（不能用new关键字调用）

* 

箭头函数没有自己this，它的this是词法的，引用的是上下文的this，即在你写这行代码的时候就箭头函数的this就已经和外层执行上下文的this绑定了(这里个人认为并不代表完全是静态的,因为外层的上下文仍是动态的可以使用call,apply,bind修改,这里只是说明了箭头函数的this始终等于它上层上下文中的this)

![](https://user-gold-cdn.xitu.io/2019/2/12/168df9eb0c098401?imageView2/0/w/1280/h/960/ignore-error/1)

因为setTimeout会将一个匿名的回调函数推入异步队列，而回调函数是具有全局性的，即在非严格模式下this会指向window，就会存在丢失变量a的问题，而如果使用箭头函数，在书写的时候就已经确定它的this等于它的上下文（这里是makeRequest的函数执行上下文，相当于将箭头函数中的this绑定了makeRequest函数执行上下文中的this）因为是controller对象调用的makeRequest函数，所以this就指向了controller对象中的a变量

箭头函数的this指向即使使用call,apply,bind也无法改变（这里也验证了为什么ECMAScript规定不能使用箭头函数作为构造函数，因为它的this已经确定好了无法改变）

### 建议 ###

箭头函数替代了以前需要显式的声明一个变量保存this的操作，使得代码更加的简洁

ES5写法:

![](https://user-gold-cdn.xitu.io/2019/2/12/168df9eb128894f7?imageView2/0/w/1280/h/960/ignore-error/1)

ES6箭头函数:

![](https://user-gold-cdn.xitu.io/2019/2/12/168df9eb2071ae0b?imageView2/0/w/1280/h/960/ignore-error/1)

再来看一个例子

![](https://user-gold-cdn.xitu.io/2019/2/12/168e0cd852903baa?imageView2/0/w/1280/h/960/ignore-error/1)

值得注意的是makeRequest后面的function不能使用箭头函数，因为这样它就会再使用上层的this，而再上层是全局的执行上下文，它的this的值会指向window,所以找不到变量a返回undefined

在数组的迭代中使用箭头函数更加简洁，并且省略了return关键字

![](https://user-gold-cdn.xitu.io/2019/2/12/168df9eb3a73b69d?imageView2/0/w/1280/h/960/ignore-error/1)

不要在可能改变this指向的函数中使用箭头函数，类似Vue中的methods,computed中的方法,生命周期函数，Vue将这些函数的this绑定了当前组件的vm实例，如果使用箭头函数会强行改变this，因为箭头函数优先级最高（无法再使用call,apply,bind改变指向）

![](https://user-gold-cdn.xitu.io/2019/2/12/168e0d5ae99d7e7c?imageView2/0/w/1280/h/960/ignore-error/1)

在把箭头函数作为日常开发的语法之前,个人建议是去了解一下箭头函数的是如何绑定this的,而不只是当做省略function这几个单词拼写,毕竟那才是ECMAScript真正希望解决的问题

# iterator迭代器 #

iterator迭代器是ES6非常重要的概念，但是很多人对它了解的不多，但是它却是另外4个ES6常用特性的实现基础（解构赋值，剩余/扩展运算符，生成器，for of循环），了解迭代器的概念有助于了解另外4个核心语法的原理，另外ES6新增的Map,Set数据结构也有使用到它，所以我放到前面来讲

对于可迭代的数据解构，ES6在内部部署了一个[Symbol.iterator]属性，它是一个函数，执行后会返回iterator对象（也叫迭代器对象），而生成iterator对象[Symbol.iterator]属性叫iterator接口,有这个接口的数据结构即被视为可迭代的

数组中的Symbol.iterator方法(iterator接口)默认部署在数组原型上:

![](https://user-gold-cdn.xitu.io/2019/2/12/168df9eb2a2e09da?imageView2/0/w/1280/h/960/ignore-error/1)

默认部署iterator接口的数据结构有以下几个，注意普通对象默认是没有iterator接口的（可以自己创建iterator接口让普通对象也可以迭代）

* Array
* Map
* Set
* String
* TypedArray（类数组）
* 函数的 arguments 对象
* NodeList 对象

iterator迭代器是一个对象，它具有一个next方法所以可以这么调用

![](https://user-gold-cdn.xitu.io/2019/2/12/168df9eb360753c9?imageView2/0/w/1280/h/960/ignore-error/1)

next方法返回又会返回一个对象，有value和done两个属性，value即每次迭代之后返回的值，而done表示是否还需要再次循环，可以看到当value为undefined时，done为true表示循环终止

梳理一下

* 可迭代的数据结构会有一个[Symbol.iterator]方法
* [Symbol.iterator]执行后返回一个iterator对象
* iterator对象有一个next方法
* 执行一次next方法(消耗一次迭代器)会返回一个有value,done属性的对象

借用 [冴羽博客]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fmqyqingfeng%2FBlog%2Fissues%2F90 ) 中ES5实现的迭代器可以更加深刻的理解迭代器是如何生成和消耗的

![](https://user-gold-cdn.xitu.io/2019/3/19/16994a020924cb2e?imageView2/0/w/1280/h/960/ignore-error/1)

# 解构赋值（常用） #

解构赋值可以直接使用对象的某个属性，而不需要通过属性访问的形式使用，对象解构原理个人认为是通过寻找相同的属性名，然后原对象的这个属性名的值赋值给新对象对应的属性

![](https://user-gold-cdn.xitu.io/2019/2/12/168df9eb466d9679?imageView2/0/w/1280/h/960/ignore-error/1)

这里左边真正声明的其实是titleOne,titleTwo这两个变量，然后会根据左边这2个变量的位置寻找右边对象中title和test[0]中的title对应的值，找到字符串abc和test赋值给titleOne,titleTwo（如果没有找到会返回undefined）

**数组解构的原理其实是消耗数组的迭代器，把生成对象的value属性的值赋值给对应的变量**

数组解构的一个用途是交换变量，避免以前要声明一个临时变量值存储值

ES6交换变量:

![](https://user-gold-cdn.xitu.io/2019/2/12/168df9eb57d68be3?imageView2/0/w/1280/h/960/ignore-error/1)

### 建议 ###

同样建议使用，因为解构赋值语意化更强，对于作为对象的函数参数来说，可以减少形参的声明，直接使用对象的属性（如果嵌套层数过多我个人认为不适合用对象解构，不太优雅）

一个常用的例子是Vuex中actions中的方法会传入2个参数，第一个参数是个对象，你可以随意命名，然后使用<名字>.commit的方法调用commit函数，或者使用对象解构直接使用commit

不使用对象解构:

![](https://user-gold-cdn.xitu.io/2019/2/12/168df9eb4dca1cbf?imageView2/0/w/1280/h/960/ignore-error/1)

使用对象解构:

![](https://user-gold-cdn.xitu.io/2019/2/12/168df9eb58a3cd0d?imageView2/0/w/1280/h/960/ignore-error/1)

另外可以给使用axios的响应结果进行解构(axios默认会把真正的响应结果放在data属性中)

![](https://user-gold-cdn.xitu.io/2019/2/12/168df9eb5c2686d1?imageView2/0/w/1280/h/960/ignore-error/1)

# 剩余/扩展运算符（常用） #

剩余/扩展运算符同样也是ES6一个非常重要的语法，使用3个点（...），后面跟着一个含有iterator接口的数据结构

### 扩展运算符 ###

以数组为例,使用扩展运算符使得可以"展开"这个数组，可以这么理解，数组是存放元素集合的一个容器，而使用扩展运算符可以将这个容器拆开，这样就只剩下元素集合，你可以把这些元素集合放到另外一个数组里面

![](https://user-gold-cdn.xitu.io/2019/2/12/168df9eb69467378?imageView2/0/w/1280/h/960/ignore-error/1)

扩展运算符可以代替ES3中数组原型的concat方法

![](https://user-gold-cdn.xitu.io/2019/2/14/168eb00e090d1386?imageView2/0/w/1280/h/960/ignore-error/1)

这里将arr1,arr2通过扩展运算符展开,随后将这些元素放到一个新的数组中,相对于concat方法语义化更强

### 剩余运算符 ###

剩余运算符最重要的一个特点就是替代了以前的arguments

访问函数的arguments对象是一个很昂贵的操作，以前的arguments.callee,arguments.caller都被废止了，建议在支持ES6语法的环境下不要在使用arguments对象，使用剩余运算符替代（箭头函数没有arguments，必须使用剩余运算符才能访问参数集合）

![](https://user-gold-cdn.xitu.io/2019/2/12/168df9eb71537568?imageView2/0/w/1280/h/960/ignore-error/1)

剩余运算符可以和数组的解构赋值一起使用，但是必须放在 **最后一个** ，因为剩余运算符的原理其实是利用了数组的迭代器，它会消耗3个点后面的数组的所有迭代器，读取所有迭代器生成对象的value属性，剩运算符后不能在有解构赋值，因为剩余运算符已经消耗了所有迭代器，而数组的解构赋值也是消耗迭代器，但是这个时候已经没有迭代器了，所以会报错

![](https://user-gold-cdn.xitu.io/2019/2/12/168df9eb6122c311?imageView2/0/w/1280/h/960/ignore-error/1)

这里first会消耗右边数组的一个迭代器，...arr会消耗剩余所有的迭代器，而第二个例子...arr直接消耗了所有迭代器，导致last没有迭代器可供消耗了，所以会报错，因为这是毫无意义的操作

**剩余运算符和扩展运算符的区别就是，剩余运算符会收集这些集合，放到右边的数组中，扩展运算符是将右边的数组拆分成元素的集合，它们是相反的**

### 在对象中使用扩展运算符 ###

这个是ES9的语法，ES9中支持在对象中使用扩展运算符，之前说过数组的扩展运算符原理是消耗所有迭代器，但对象中并没有迭代器，我个人认为可能是实现原理不同，但是仍可以理解为将键值对从对象中拆开，它可以放到另外一个普通对象中

![](https://user-gold-cdn.xitu.io/2019/2/12/168df9eb76380489?imageView2/0/w/1280/h/960/ignore-error/1)

其实它和另外一个ES6新增的API相似，即Object.assign，它们都可以合并对象，但是还是有一些不同Object.assign会触发目标对象的setter函数，而对象扩展运算符不会，这个我们放到后面讨论

### 建议 ###

使用扩展运算符可以快速的将类数组转为一个真正的数组

![](https://user-gold-cdn.xitu.io/2019/2/12/168df9eb82d7b593?imageView2/0/w/1280/h/960/ignore-error/1)

合并多个数组

![](https://user-gold-cdn.xitu.io/2019/2/12/168df9eb854df5fb?imageView2/0/w/1280/h/960/ignore-error/1)

函数柯里化

![](https://user-gold-cdn.xitu.io/2019/2/12/168df9eb8982ea6e?imageView2/0/w/1280/h/960/ignore-error/1)

# 对象属性/方法简写(常用) #

### 对象属性简写 ###

es6允许当对象的属性和值相同时，省略属性名

![](https://user-gold-cdn.xitu.io/2019/2/12/168df9eb8cbe0f2d?imageView2/0/w/1280/h/960/ignore-error/1)

需要注意的是

* **省略的是属性名而不是值**
* 值必须是一个变量

对象属性简写经常与解构赋值一起使用

![](https://user-gold-cdn.xitu.io/2019/2/12/168df9eb9359bc8b?imageView2/0/w/1280/h/960/ignore-error/1)

结合上文的解构赋值，这里的代码会其实是声明了x,y,z变量，因为bar函数会返回一个对象，这个对象有x,y,z这3个属性，解构赋值会寻找等号右边表达式的x,y,z属性，找到后赋值给声明的x,y,z变量

### 方法简写 ###

es6允许当一个对象的属性的值是一个函数（即是一个方法），可以使用简写的形式

![](https://user-gold-cdn.xitu.io/2019/2/12/168df9eba5a9ac98?imageView2/0/w/1280/h/960/ignore-error/1)

在Vue中因为都是在vm对象中书写方法，完全可以使用方法简写的方式书写函数

![](https://user-gold-cdn.xitu.io/2019/2/12/168df9eba26780ea?imageView2/0/w/1280/h/960/ignore-error/1)

# for ... of循环 #

for ... of是作为ES6新增的遍历方式,允许遍历一个含有iterator接口的数据结构并且返回各项的值,和ES3中的for ... in的区别如下

* 

for ... of只能用在可迭代对象上,获取的是迭代器返回的value值,for ... in 可以获取所有对象的键名

* 

for ... in会遍历对象的整个原型链,性能非常差不推荐使用,而for ... of只遍历当前对象不会遍历它的原型链

* 

对于数组的遍历,for ... in会返回数组中所有可枚举的属性(包括原型链上可枚举的属性),for ... of只返回数组的下标对应的属性值

for ... of循环的原理其实也是利用了可迭代对象内部部署的iterator接口,如果将for ... of循环分解成最原始的for循环,内部实现的机制可以这么理解

![](https://user-gold-cdn.xitu.io/2019/2/12/168df9eba84cafbe?imageView2/0/w/1280/h/960/ignore-error/1)

可以看到只要满足第二个条件(iterator.next()存在且res.done为true)就可以一直循环下去,并且每次把迭代器的next方法生成的对象赋值给res,然后将res的value属性赋值给for ... of第一个条件中声明的变量即可,res的done属性控制是否继续遍历下去

for... of循环同时支持break,continue,return(在函数中调用的话)并且可以和对象解构赋值一起使用

![](https://user-gold-cdn.xitu.io/2019/2/12/168df9ebb0fe5712?imageView2/0/w/1280/h/960/ignore-error/1)

arr数组每次使用for ... of循环都返回一对象({a:1},{a:2},{a:3}),然后会经过对象解构,寻找属性为a的值,赋值给obj.a,所以在每轮循环的时候obj.a会分别赋值为1,2,3

# Promise（常用） #

Promise作为ES6中推出的新的概念，改变了JS的异步编程，现代前端大部分的异步请求都是使用Promise实现，fetch这个web api也是基于Promise的，这里不得简述一下之前统治JS异步编程的回调函数，回调函数有什么缺点，Promise又是怎么改善这些缺点

### 回调函数 ###

众所周知，JS是单线程的，因为多个线程改变DOM的话会导致页面紊乱，所以设计为一个单线程的语言，但是浏览器是多线程的，这使得JS同时具有异步的操作，即定时器，请求，事件监听等，而这个时候就需要一套事件的处理机制去决定这些事件的顺序，即Event Loop（事件循环），这里不会详细讲解事件循环，只需要知道，前端发出的请求，一般都是会进入浏览器的http请求线程，等到收到响应的时候会通过回调函数推入异步队列，等处理完主线程的任务会读取异步队列中任务，执行回调

在《你不知道的JavaScript》下卷中，这么介绍

使用回调函数处理异步请求相当于把你的回调函数置于了一个黑盒，虽然你声明了等到收到响应后执行你提供的回调函数,可是你并不知道这个第三方库会在什么具体会怎么执行回调函数

使用第三方的请求库你可能会这么写:

![](https://user-gold-cdn.xitu.io/2019/2/12/168e0109224ac6ef?imageView2/0/w/1280/h/960/ignore-error/1)

收到响应后，执行后面的回调打印字符串，但是如果这个第三方库有类似超时重试的功能，可能会执行多次你的回调函数，如果是一个支付功能，你就会发现你扣的钱可能就不止1000元了-.-

第二个众所周知的问题就是，在回调函数中再嵌套回调函数会导致代码非常难以维护，这是人们常说的“回调地狱”

![](https://user-gold-cdn.xitu.io/2019/2/12/168df9ebc8e9311d?imageView2/0/w/1280/h/960/ignore-error/1)

另外你使用的第三方ajax库还有可能并没有提供一些错误的回调，请求失败的一些错误信息可能会被吞掉，而你确完全不知情(nodejs提供了err-first风格的回调,即异步操作的第一个回调永远是错误的回调处理,但是你还是不能保证所有的库都提供了发送错误时的执行的回调函数)

总结一下回调函数的一些缺点

* 

多重嵌套，导致回调地狱

* 

代码跳跃，并非人类习惯的思维模式

* 

信任问题，你不能把你的回调完全寄托与第三方库，因为你不知道第三方库到底会怎么执行回调（多次执行）

* 

第三方库可能没有提供错误处理

* 

不清楚回调是否都是异步调用的（可以同步调用ajax，在收到响应前会阻塞整个线程，会陷入假死状态，非常不推荐）

` xhr.open( "GET" , "/try/ajax/ajax_info.txt" , false ); //通过设置第三个async为 false 可以同步调用ajax 复制代码`

### Promise ###

针对回调函数这么多缺点，ES6中引入了一个新的概念Promise，Promise是一个构造函数，通过new关键字创建一个Promise的实例，来看看Promise是怎么解决回调函数的这些问题

![](https://user-gold-cdn.xitu.io/2019/2/12/168df9ebca9f5e7b?imageView2/0/w/1280/h/960/ignore-error/1)

**Promise并不是回调函数的衍生版本，而是2个概念，所以需要将之前的回调函数改为支持Promise的版本，这个过程成为"提升"，或者"promisory"，现代MVVM框架常用的第三方请求库axios就是一个典型的例子，另外nodejs中也有bluebird，Q等**

* 多重嵌套，导致回调地狱

Promise在设计的时候引入了链式调用的概念，每个then方法 **同样也是一个Promise** ，因此可以无限链式调用下去

![](https://user-gold-cdn.xitu.io/2019/2/12/168df9ebced32817?imageView2/0/w/1280/h/960/ignore-error/1)

配合箭头函数，明显的比之前回调函数的多层嵌套优雅很多

* 代码跳跃，并非人类习惯的思维模式

Promise使得能够同步思维书写代码，上述的代码就是先请求3000端口，得到响应后再请求3001，再请求3002，再请求3003，而书写的格式也是符合人类的思维，从先到后

* 信任问题，你不能把你的回调完全寄托与第三方库，因为你不知道第三方库到底会怎么执行回调（多次执行）

Promise本身是一个状态机，具有以下3个状态

* pending（等待）
* fulfilled（成功）
* rejected（拒绝）

当请求发送没有得到响应的时候为pending状态，得到响应后会resolve(决议)当前这个Promise实例,将它变为fulfilled/rejected(大部分情况会变为fulfilled)

当请求发生错误后会执行reject(拒绝)将这个Promise实例变为rejected状态

一个Promise实例的状态只能从pending => fulfilled 或者从 pending => rejected，即当一个Promise实例从pending状态改变后，就不会再改变了（不存在fulfilled => rejected 或 rejected => fulfilled）

而Promise实例必须主动调用then方法，才能将值从Promise实例中取出来（前提是Promise不是pending状态），这一个“主动”的操作就是解决这个问题的关键，即第三方库做的只是改变Promise的状态，而响应的值怎么处理，这是开发者主动控制的，这里就实现了控制反转，将原来第三方库的控制权转移到了开发者上

![](https://user-gold-cdn.xitu.io/2019/2/12/168df9ebcee4bf6f?imageView2/0/w/1280/h/960/ignore-error/1)

* 第三方库可能没有提供错误处理

Promise的then方法会接受2个函数，第一个函数是这个Promise实例被resolve时执行的回调，第二个函数是这个Promise实例被reject时执行的回调，而这个也是开发者主动调用的

使用Promise在异步请求发送错误的时候，即使没有捕获错误，也不会阻塞主线程的代码（准确的来说，异步的错误都不会阻塞主线程的代码）

![](https://user-gold-cdn.xitu.io/2019/2/12/168df9ebd0318392?imageView2/0/w/1280/h/960/ignore-error/1)

* 不清楚回调是否都是异步调用的

Promise在设计的时候保证所有响应的处理回调都是异步调用的，不会阻塞代码的执行，Promise将then方法的回调放入一个叫微任务的队列中（MicroTask），确保这些回调任务在同步任务执行完以后再执行，这部分同样也是事件循环的知识点，有兴趣的朋友可以深入研究一下

对于第三个问题中,为什么说执行了resolve函数后"大部分情况"会进入fulfilled状态呢?考虑以下情况

![](https://user-gold-cdn.xitu.io/2019/2/14/168eb59272b3154f?imageView2/0/w/1280/h/960/ignore-error/1)

![](https://user-gold-cdn.xitu.io/2019/2/14/168eb5ae056a4568?imageView2/0/w/1280/h/960/ignore-error/1)

(这里用一个定时器在下轮事件循环中打印这个Promise实例的状态,否则会是pending状态)

很多人认为promise中调用了resolve函数则这个promise一定会进入fulfilled状态,但是这里可以看到,即使调用了resolve函数,仍返回了一个拒绝状态的Promise,原因是因为如果在一个promise的resolve函数中又传入了一个Promise,会展开传入的这个promise

这里因为传入了一个拒绝状态的promise,resolve函数展开这个promise后,就会变成一个拒绝状态的promise,所以把resolve理解为决议比较好一点

等同于这样

![](https://user-gold-cdn.xitu.io/2019/2/14/168eb60451743e8a?imageView2/0/w/1280/h/960/ignore-error/1)

### 建议 ###

在日常开发中，建议全面拥抱新的Promise语法，其实现在的异步编程基本也都使用的是Promise

建议使用ES7的async/await进一步的优化Promise的写法，async函数始终返回一个Promise，await可以实现一个"等待"的功能，async/await被成为异步编程的终极解决方案，即用同步的形式书写异步代码，并且能够更优雅的实现异步代码顺序执行以及在发生异步的错误时提供更精准的错误信息,详细用法可以看阮老师的ES6标准入门

![](https://user-gold-cdn.xitu.io/2019/2/21/1690e7a5ff97364e?imageView2/0/w/1280/h/960/ignore-error/1)

关于Promise还有很多很多需要讲的，包括它的静态方法all，race，resolve，reject，Promise的执行顺序，Promise嵌套Promise，thenable对象的处理等，碍于篇幅这里只介绍了一下为什么需要使用Promise。但很多开发者在日常使用中只是了解这些API，却不知道Promise内部具体是怎么实现的，遇到复杂的异步代码就无从下手，非常建议去了解一下Promise A+的规范，自己实现一个Promise

# ES6 Module(常用) #

在ES6 Module出现之前，模块化一直是前端开发者讨论的重点，面对日益增长的需求和代码，需要一种方案来将臃肿的代码拆分成一个个小模块，从而推出了AMD,CMD和CommonJs这3种模块化方案，前者用在浏览器端，后面2种用在服务端，直到ES6 Module出现

**ES6 Module默认目前还没有被浏览器支持，需要使用babel，在日常写demo的时候经常会显示这个错误**

![](https://user-gold-cdn.xitu.io/2019/2/12/168df9ebeb392397?imageView2/0/w/1280/h/960/ignore-error/1)

可以在script标签中使用tpye="module"在 **同域** 的情况下可以解决（非同域情况会被同源策略拦截，webstorm会开启一个同域的服务器没有这个问题，vscode貌似不行）

![](https://user-gold-cdn.xitu.io/2019/2/12/168df9ebedf62b58?imageView2/0/w/1280/h/960/ignore-error/1)

ES6 Module使用import关键字导入模块，export关键字导出模块，它还有以下特点

* 

ES6 Module是静态的，也就是说它是在编译阶段运行，和var以及function一样具有提升效果（这个特点使得它支持tree shaking）

* 

自动采用严格模式（顶层的this返回undefined）

* 

ES6 Module支持使用export {<变量>}导出具名的接口，或者export default导出匿名的接口

module.js导出:

![](https://user-gold-cdn.xitu.io/2019/2/12/168df9ebef0bed93?imageView2/0/w/1280/h/960/ignore-error/1)

a.js导入:

![](https://user-gold-cdn.xitu.io/2019/2/12/168df9ebf3ade585?imageView2/0/w/1280/h/960/ignore-error/1)

这两者的区别是，export {<变量>}导出的是一个变量的引用，export default导出的是一个值

什么意思呢，就是说在a.js中使用import导入这2个变量的后，在module.js中因为某些原因x变量被改变了，那么会立刻反映到a.js，而module.js中的y变量改变后，a.js中的y还是原来的值

module.js:

![](https://user-gold-cdn.xitu.io/2019/2/12/168df9ec0ee34ff0?imageView2/0/w/1280/h/960/ignore-error/1)

a.js:

![](https://user-gold-cdn.xitu.io/2019/2/12/168df9ec1621b560?imageView2/0/w/1280/h/960/ignore-error/1)

可以看到给module.js设置了一个一秒后改变x,y变量的定时器,在一秒后同时观察导入时候变量的值,可以发现x被改变了,但y的值仍是20,因为y是通过export default导出的,在导入的时候的值相当于只是导入数字20,而x是通过export {<变量>}导出的,它导出的是一个变量的引用,即a.js导入的是当前x的值,只关心 **当前** x变量的值是什么,可以理解为一个"活链接"

export default这种导出的语法其实只是指定了一个命名导出,而它的名字叫default,换句话说,将模块的导出的名字重命名为default,也可以使用import <变量> from <路径> 这种语法导入

module.js导出:

![](https://user-gold-cdn.xitu.io/2019/2/12/168df9ec0954d5de?imageView2/0/w/1280/h/960/ignore-error/1)

a.js导入:

![](https://user-gold-cdn.xitu.io/2019/2/12/168df9ebf3ade585?imageView2/0/w/1280/h/960/ignore-error/1)

**但是由于是使用export {<变量>}这种形式导出的模块,即使被重命名为default,仍然导出的是一个变量的引用**

这里再来说一下目前为止主流的模块化方案ES6 Module和CommonJs的一些区别

* 

CommonJs输出的是一个值的拷贝,ES6 Module通过export {<变量>}输出的是一个变量的引用,export default输出的是一个值

* 

CommonJs运行在服务器上,被设计为运行时加载,即代码执行到那一行才回去加载模块,而ES6 Module是静态的输出一个接口,发生在编译的阶段

* 

CommonJs在第一次加载的时候运行一次并且会生成一个缓存,之后加载返回的都是缓存中的内容

### import() ###

关于ES6 Module静态编译的特点,导致了无法动态加载,但是总是会有一些需要动态加载模块的需求,所以现在有一个提案,使用把import作为一个函数可以实现动态加载模块,它返回一个Promise,Promise被resolve时的值为输出的模块

![](https://user-gold-cdn.xitu.io/2019/2/12/168df9ec16da2511?imageView2/0/w/1280/h/960/ignore-error/1)

![](https://user-gold-cdn.xitu.io/2019/2/12/168df9ec186c4715?imageView2/0/w/1280/h/960/ignore-error/1)

使用import方法改写上面的a.js使得它可以动态加载(使用静态编译的ES6 Module放在条件语句会报错,因为会有提升的效果,并且也是不允许的),可以看到输出了module.js的一个变量x和一个默认输出

Vue中路由的懒加载的ES6写法就是使用了这个技术,使得在路由切换的时候能够动态的加载组件渲染视图

# 函数默认值 #

ES6允许在函数的参数中设置默认值

ES5写法:

![](https://user-gold-cdn.xitu.io/2019/2/12/168df9ec37511eed?imageView2/0/w/1280/h/960/ignore-error/1)

ES6写法:

![](https://user-gold-cdn.xitu.io/2019/2/12/168df9ec2bcb6af0?imageView2/0/w/1280/h/960/ignore-error/1)

相比ES5,ES6函数默认值直接写在参数上,更加的直观

如果使用了函数默认参数,在函数的参数的区域(括号里面),它会作为一个单独的 **块级作用域** ,并且拥有let/const方法的一些特性,比如暂时性死区

![](https://user-gold-cdn.xitu.io/2019/2/12/168df9ec1f367047?imageView2/0/w/1280/h/960/ignore-error/1)

这里当运行func的时候,因为没有传参数,使用函数默认参数,y就会去寻找x的值,在沿着词法作用域在外层找到了值为1的变量x

再来看一个例子

![](https://user-gold-cdn.xitu.io/2019/2/12/168df9ec35ac50b5?imageView2/0/w/1280/h/960/ignore-error/1)

这里同样没有传参数,使用函数的默认赋值,x通过词法作用域找到了变量w,所以x默认值为2,y同样通过词法作用域找到了刚刚定义的x变量,y的默认值为3,但是在解析到z = z + 1这一行的时候,JS解释器先会去解析z+1找到相应的值后再赋给变量z,但是因为暂时性死区的原因(let/const"劫持"了这个块级作用域,无法在声明之前使用这个变量,上文有解释),导致在let声明之前就使用了变量z,所以会报错

这样理解函数的默认值会相对容易一些

![](https://user-gold-cdn.xitu.io/2019/2/13/168e47dff16b7e77?imageView2/0/w/1280/h/960/ignore-error/1)

当传入的参数为undefined时才使用函数的默认值(显式传入undefined也会触发使用函数默认值,传入null则不会触发)

在举个例子:

![](https://user-gold-cdn.xitu.io/2019/2/12/168df9ec395bebe1?imageView2/0/w/1280/h/960/ignore-error/1)

这里借用阮一峰老师书中的一个例子,func的默认值为一个函数,执行后返回foo变量,而在函数内部执行的时候,相当于对foo变量的一次变量查询(LHS查询),而查询的起点是这个单独的块级作用域,即JS解释器不会去查询去函数内部查询变量foo,而是沿着词法作用域先查看同一作用域(前面的函数参数)中有没有foo变量,再往函数的外部寻找foo变量,最终找不到所以报错了,这个也是函数默认值的一个特点

![](https://user-gold-cdn.xitu.io/2019/3/18/1698f6f2b3f8adb6?imageView2/0/w/1280/h/960/ignore-error/1)

通过debugger可以更加直观的发现在这个函数内部可以通过词法作用域访问func函数,foo变量,还有this,但是当查看func函数的词法作用域时,发现它只能访问到Global,即全局作用域,foo变量并不存在于它的词法作用域中

### 函数默认值配合解构赋值 ###

![](https://user-gold-cdn.xitu.io/2019/2/12/168df9ec50606fd1?imageView2/0/w/1280/h/960/ignore-error/1)

第一行给func函数传入了2个空对象,所以函数的第一第二个参数都不会使用函数默认值,然后函数的第一个参数会尝试解构对象,提取变量x,因为第一个参数传入了一个空对象,所以解构不出变量x,但是这里又在内层设置了一个默认值,所以x的值为10,而第二个参数同样传了一个空对象,不会使用函数默认值,然后会尝试解构出变量y,发现空对象中也没有变量y,但是y没有设置默认值所以解构后y的值为undefined

第二行第一个参数显式的传入了一个undefined,所以会使用函数默认值为一个空对象,随后和第一行一样尝试解构x发现x为undefined,但是设置了默认值所以x的值为10,而y和上文一样为undefined

第三行2个参数都会undefined,第一个参数和上文一样,第二个参数会调用函数默认值,赋值为{y:10},然后尝试解构出变量y,即y为10

第四行和第三行相同,一个是显式传入undefined,一个是隐式不传参数

第五行直接使用传入的参数,不会使用函数默认值,并且能够顺利的解构出变量x,y

# Proxy #

Proxy作为一个"拦截器",可以在目标对象前架设一个拦截器,他人访问对象,必须先经过这层拦截器,Proxy同样是一个构造函数,使用new关键字生成一个拦截对象的实例,ES6提供了非常多对象拦截的操作,几乎覆盖了所有可能修改目标对象的情况(Proxy一般和Reflect配套使用,前者拦截对象,后者返回拦截的结果,Proxy上有的的拦截方法Reflect都有)

![](https://user-gold-cdn.xitu.io/2019/2/12/168df9ec5addffad?imageView2/0/w/1280/h/960/ignore-error/1)

### Object.definePropery ###

提到Proxy就不得不提一下ES5中的Object.defineProperty,这个api可以给一个对象添加属性以及这个属性的属性描述符/访问器(这2个不能共存,同一属性只能有其中一个),属性描述符有configurable,writable,enumerable,value这4个属性,分别代表是否可配置,是否只读,是否可枚举和属性的值,访问器有configurable,enumerable,get,set,前2个和属性描述符功能相同,后2个都是函数,定义了get,set后对元素的读写操作都会执行后面的getter/setter函数,并且覆盖默认的读写行为

![](https://user-gold-cdn.xitu.io/2019/2/12/168df9ec56c9a8f0?imageView2/0/w/1280/h/960/ignore-error/1)

定义了obj中a属性的表示为只读,且不可枚举,obj2定义了get,但没有定义set表示只读,并且读取obj2的b属性返回的值是getter函数的返回值

ES5中的Object.defineProperty这和Proxy有什么关系呢?个人理解Proxy是Object.defineProperty的增强版,ES5只规定能够定义属性的属性描述符或访问器.而Proxy增强到了13种,具体太多了我就不一一放出来了,这里我举几个比较有意思的例子

### handler.apply ###

apply可以让我们拦截一个函数(JS中函数也是对象,Proxy也可以拦截函数)的执行,我们可以把它用在函数节流中

![](https://user-gold-cdn.xitu.io/2019/2/12/168df9ec6239f7d9?imageView2/0/w/1280/h/960/ignore-error/1)

调用拦截后的函数:

![](https://user-gold-cdn.xitu.io/2019/2/13/168e52a570e4d7d1?imageView2/0/w/1280/h/960/ignore-error/1)

### handler.contruct ###

contruct可以拦截通过new关键字调用这个函数的操作,我们可以把它用在单例模式中

![](https://user-gold-cdn.xitu.io/2019/2/12/168df9ec6bf3b894?imageView2/0/w/1280/h/960/ignore-error/1)

这里通过一个闭包保存了instance变量,每次使用new关键字调用被拦截的函数后都会查看这个instance变量,如果存在就返回闭包中保存的instance变量,否则就新建一个实例,这样可以实现全局只有一个实例

### handler.defineProperty ###

defineProperty可以拦截对这个对象的Object.defineProerty操作

**注意对象内部的默认的[[SET]]操作(即对这个对象的属性赋值)会间接触发defineProperty和getOwnPropertyDescriptor这2个拦截方法**

![](https://user-gold-cdn.xitu.io/2019/2/12/168df9ec717f7a0c?imageView2/0/w/1280/h/960/ignore-error/1)

这里有几个知识点

* 这里使用了递归的操作,当需要访问对象的属性时候,会判断代理的对象属性的值仍是一个可以代理的对象就递归的进行代理,否则通过错误捕获执行默认的get操作
* 定义了defineProperty的拦截方法,当对这个代理对象的某个属性进行赋值的时候会执行对象内部默认的[[SET]]操作进行赋值,这个操作会间接触发defineProperty这个方法,随后会执行定义的callback函数

这样就实现了无论对象嵌套多少层,只要有属性进行赋值就会触发get方法,对这层对象进行代理,随后触发defineProperty执行callback回调函数

### 其他的使用场景 ###

Proxy另外还有很多功能,比如在实现验证器的时候,可以将业务逻辑和验证器分离达到解耦,通过get拦截对私有变量的访问实现私有变量,拦截对象做日志记录，实现微信api的promise化等

### Vue ###

尤大预计2019年下半年发布Vue3.0,其中一个核心的功能就是使用Proxy替代Object.defineProperty

我相信了解过一点Vue响应式原理的人都知道Vue框架在对象拦截上的一些不足

` <template> <div> <div>{{arr}}</div> <div>{{obj}}</div> <button @click= "handleClick" >修改arr下标</button> <button @click= "handleClick2" >创建obj的属性</button> </div> </template> <script> export default { name: "index" , data () { return { arr:[1,2,3], obj:{ a:1, b:2 } } }, methods: { handleClick () { this.arr[0] = 10 console.log(this.arr) }, handleClick2 () { this.obj.c = 3 console.log(this.obj) } }, } </script> 复制代码`

![](https://user-gold-cdn.xitu.io/2019/2/12/168df9ec77c2b72f?imageslim)

可以看到这里数据改变了,控制台打印出了新的值,但是视图没有更新,这是因为Vue内部使用Object.defineProperty进行的数据劫持,而这个API无法探测到 **对象根属性的添加和删除,以及直接给数组下标进行赋值** ,所以不会通知渲染watcher进行视图更新,而理论上这个API也无法探测到数组的一系列方法(push,splice,pop),但是Vue框架修改了数组的原型,使得在调用这些方法修改数据后会执行视图更新的操作

` //源码位置:src/core/observer/array.js methodsToPatch.forEach( function (method) { // cache original method var original = arrayProto[method]; def(arrayMethods, method, function mutator () { var args = [], len = arguments.length; while ( len-- ) args[ len ] = arguments[ len ]; var result = original.apply(this, args); var ob = this.__ob__; var inserted; switch (method) { case 'push' : case 'unshift' : inserted = args; break case 'splice' : inserted = args.slice(2); break } if (inserted) { ob.observeArray(inserted); } // notify change ob.dep.notify(); //这一行就会主动调用notify方法,会通知到渲染watcher进行视图更新 return result }); }); 复制代码`

在 [掘金翻译的尤大Vue3.0计划]( https://juejin.im/post/5bb719b9f265da0ab915dbdd ) 中写到

> 
> 
> 
> 3.0 将带来一个基于 Proxy 的 observer 实现，它可以提供覆盖语言 (JavaScript——译注) 全范围的响应式能力，消除了当前
> Vue 2 系列中基于 Object.defineProperty 所存在的一些局限，如： 对属性的添加、删除动作的监测 对数组基于下标的修改、对于
> .length 修改的监测 对 Map、Set、WeakMap 和 WeakSet 的支持
> 
> 

Proxy就没有这个问题,并且还提供了更多的拦截方法,完全可以替代Object.defineProperty,唯一不足的也就是浏览器的支持程度了(IE:谁在说我?)

所以要想深入了解Vue3.0实现机制,学会Proxy是必不可少的

# Object.assign #

这个ES6新增的Object静态方法允许我们进行多个对象的合并

![](https://user-gold-cdn.xitu.io/2019/2/12/168df9ec7eb6298e?imageView2/0/w/1280/h/960/ignore-error/1)

可以这么理解,Object.assign遍历需要合并给target的对象(即sourece对象的集合)的属性,用 **等号** 进行赋值,这里遍历{a:1}将属性a和值数字1赋值给target对象,然后再遍历{b:2}将属性b和值数字2赋值给target对象

这里罗列了一些这个API的需要注意的知识点

* 

Object.assign是浅拷贝,对于值是引用类型的属性,拷贝仍旧的是它的引用

* 

可以拷贝Symbol属性

* 

不能拷贝不可枚举的属性

* 

Object.assign保证target始终是一个对象,如果传入一个基本类型,会转为基本包装类型,null/undefined没有基本包装类型,所以传入会报错

* 

source参数如果是不可枚举的数据类型会忽略合并(字符串类型被认为是可枚举的,因为内部有iterator接口)

* 

因为是用 **等号** 进行赋值,如果被赋值的对象的属性有setter函数会触发setter函数,同理如果有getter函数,也会调用赋值对象的属性的getter函数(这就是为什么Object.assign无法合并对象属性的访问器,因为它会直接执行对应的getter/setter函数而不是合并它们,如果需要合并对象属性的getter/setter函数,可以使用ES7提供的Object.getOwnPropertyDescriptors和Object.defineProperties这2个API实现)

![](https://user-gold-cdn.xitu.io/2019/2/13/168e64e6f6872a09?imageView2/0/w/1280/h/960/ignore-error/1)

![](https://user-gold-cdn.xitu.io/2019/2/13/168e64bd2eb1e640?imageView2/0/w/1280/h/960/ignore-error/1)

可以看到这里成功的复制了obj对象中a属性的getter/setter

为了加深了解我自己模拟了Object.assign的实现,可供参考

![](https://user-gold-cdn.xitu.io/2019/2/12/168df9ec8acf44b3?imageView2/0/w/1280/h/960/ignore-error/1)

这里有一个坑不得不提,对于target参数传入一个字符串,内部会转换为基本包装类型,而字符串基本包装类型的属性是只读的(属性描述符的writable属性为false),这里感谢 [木易杨的专栏]( https://juejin.im/post/5c31e5c4e51d45524975d05a )

![](https://user-gold-cdn.xitu.io/2019/2/13/168e66d1185ec325?imageView2/0/w/1280/h/960/ignore-error/1)

![](https://user-gold-cdn.xitu.io/2019/2/13/168e66541e1da350?imageView2/0/w/1280/h/960/ignore-error/1)

打印对象属性的属性描述符可以看到下标属性的值都是只读的,即不能再次赋值,所以尝试以下操作会报错

![](https://user-gold-cdn.xitu.io/2019/2/13/168e65efa1ab106f?imageView2/0/w/1280/h/960/ignore-error/1)

字符串abc会转为基本包装类型,然后将字符串def合并给这个基本包装类型的时候会将字符串def展开,分别将字符串def赋值给基本包装类型abc的0,1,2属性,随后就会在赋值的时候报错(非严格模式下会只会静默处理,ES6的Object.assign默认开启了严格模式)

### 和ES9的对象扩展运算符对比 ###

ES9支持在对象上使用扩展运算符,实现的功能和Object.assign相似,唯一的区别就是在含有getter/setter函数的对象的属性上有所区别

![](https://user-gold-cdn.xitu.io/2019/2/14/168e9b0427f2d2ec?imageView2/0/w/1280/h/960/ignore-error/1)

![](https://user-gold-cdn.xitu.io/2019/2/14/168e9b09b5b84f7d?imageView2/0/w/1280/h/960/ignore-error/1)

(最后一个字符串get可以忽略,这是控制台为了显示a变量触发的getter函数)

分析一下这个例子

ES9:

* 会合并2个对象,并且只触发2个对象对应属性的getter函数
* 相同属性的后者覆盖了前者,所以a属性的值是第二个getter函数return的值

ES6:

* 同样会合并这2个对象,并且只触发了obj上a属性的setter函数而不会触发它的getter函数(结合上述Object.assgin的内部实现理解会容易一些)
* obj上a属性的setter函数替代默认的赋值行为,导致obj2的a属性不会被复制过来

除去对象属性有getter/setter的情况,Object.assgin和对象扩展运算符功能是相同的,两者都可以使用,两者都是浅拷贝,使用ES9的方法相对简洁一点

### 建议 ###

* Vue中重置data中的数据

这个是我最常用的小技巧,使用Object.assign可以将你目前组件中的data对象和组件默认初始化状态的data对象中的数据合并,这样可以达到初始化data对象的效果

![](https://user-gold-cdn.xitu.io/2019/2/12/168df9ec96c22302?imageView2/0/w/1280/h/960/ignore-error/1)

在当前组件的实例中$data属性保存了当前组件的data对象,而$options是当前组件实例初始化时的一些属性,其中有个data方法,即在在组件中写的data函数,执行后会返回一个初始化的data对象,然后将这个初始化的data对象合并到当前的data来初始化所有数据

* 给对象合并需要的默认属性

![](https://user-gold-cdn.xitu.io/2019/2/12/168df9eca34511aa?imageView2/0/w/1280/h/960/ignore-error/1)

可以封装一个函数,外层声明一个DEFAULTS常量,options为每次传入的动态配置,这样每次执行后会合并一些默认的配置项

* 在传参的时候可以多个数据合并成一个对象传给后端

![](https://user-gold-cdn.xitu.io/2019/2/12/168df9ec9ea1c032?imageView2/0/w/1280/h/960/ignore-error/1)

# 参考资料 #

* 

[阮一峰：ES6标准入门]( https://link.juejin.im?target=http%3A%2F%2Fes6.ruanyifeng.com )

* 

[慕课网：ES6零基础教学]( https://link.juejin.im?target=https%3A%2F%2Fcoding.imooc.com%2Fclass%2F98.html )

* 

你不知道的JavaScript下卷