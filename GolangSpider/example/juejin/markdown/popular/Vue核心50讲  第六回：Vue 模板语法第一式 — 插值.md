# Vue核心50讲 | 第六回：Vue 模板语法第一式 — 插值 #

书接上文，上一回咱们讲到了 Vue 的黑魔法 —— 模板语法，但是咱们只是解释了什么是模板语法。关于 Vue 的模板语法要怎么使用，具体又包含哪些内容，咱们现在是一概不知。那么接下来，就让咱们慢慢走进 Vue 的模板语法，首先要来讲的就是模板语法的第一式 —— 插值。

## 什么是插值 ##

那么问题来了，什么是插值啊？每次遇到一个新的概念，心都好累！关键是 Vue 官方并没有给出解释，而是直接给了具体的用法而已。

![](https://user-gold-cdn.xitu.io/2019/6/5/16b27fb855af0f69?imageView2/0/w/1280/h/960/ignore-error/1)

上一回咱们讲到模板语法的时候，说了模板语法通过 Vue 框架就可以绑定到对应 Vue 实例的数据内容。换句话讲，模板语法就是把 Vue 实例的数据展示在 HTML 网页中。而现在咱们说的插值，就是把这些数据展示文本内容、浏览器解析后的结果、HTML 标签属性之类的。

比如下面这段代码示例，就是把数据展示成文本内容的：

` <div id= "app" > <h2>{{ message }}</h2> </div> <script src= "scripts/vue.js" ></script> <script> // 创建Vue的实例对象 var app = new Vue({ el: '#app' , data: { // 存储在Vue的实例对象中 message: '前端课湛' } }); </script> 复制代码`

上面这段代码中的 ` {{ message }}` 就是插值了。

Vue 模板语法的插值有几种用法呢？一共是四种：

* 文本插值
* HTML 插值
* v-bind 指令
* 表达式

接下来，咱们分别地来说一说吧。

## 文本插值 ##

说到这文本插值啊，它可是 Vue 实现数据绑定中最常见的一种形式。不仅如此，文本插值还有一个名字，叫做“Mustache”语法。又是个新概念，心太累了！说白了，就是一对花括号的写法，就像 {{ message }} 这种就是文本插值了。其实很简单吧？！

` <div id= "app" > <h2>{{ message }}</h2> </div> <script src= "scripts/vue.js" ></script> <script> // 创建Vue的实例对象 var app = new Vue({ el: '#app' , data: { // 存储在Vue的实例对象中 message: '前端课湛' } }); </script> 复制代码`

刚刚用到的这段代码就是文本插值的示例代码啦。不过，这块需要注意一下哈！就是 Vue 实例中的 message 的值变化的时候，对应插值处也会跟着更新。比如，上面这段代码运行之后，咱们在浏览器的控制台来改变 message 的值，插值处也会改变的。

如果想要插值处不跟着数据变化而变化的话，咱们也是有办法的。Vue 提供了 v-once 指令，这个指令就可以不让插值处跟着数据变化二变化了。

这时候你可能会吐槽了，心真尼玛累！咋又来个新概念，啥是指令啊？简单来说，指令就是 Vue 提供的 HTML 标签的自定义属性，你可以先这么理解。

闲言少叙，书归正传。

上面那段代码示例，咱们只需要在对应的标签中定义 v-once 指令就行了。就像下面这样示儿的：

` <div id= "app" > <h2 v-once>{{ message }}</h2> </div> <script src= "scripts/vue.js" ></script> <script> // 创建Vue的实例对象 var app = new Vue({ el: '#app' , data: { // 存储在Vue的实例对象中 message: '前端课湛' } }); </script> 复制代码`

上面这段代码有兴趣你自己运行去吧，自己看看效果。

## HTML 插值 ##

文本插值说明白了之后，咱们来看看 HTML 插值吧。说到这个 HTML 插值吧，我是觉得就是增强版的文本插值。为啥这么说呢？因为文本插值只能把文本展示在插值处，但是如果数据本身就是一段 HTML 代码呢？文本插值就会原封不动地把这段 HTML 代码展示出来。不信？咱来试试：

` <div id= "app" > {{ message }} </div> <script src= "scripts/vue.js" ></script> <script> // 创建Vue的实例对象 var app = new Vue({ el: '#app' , data: { // 存储在Vue的实例对象中 message: '<span style="color: lightcoral">前端课湛</span>' } }); </script> 复制代码`

上面这段代码示例中的数据就是 ` <span style="color: lightcoral">前端课湛</span>` 这样一段 HTML 代码，文本插值会把这段 HTML 代码原封不动地展示出来。

![](https://user-gold-cdn.xitu.io/2019/6/5/16b27fe851d56cea?imageView2/0/w/1280/h/960/ignore-error/1)

这个时候，要想浏览器去解析这段 HTML 代码的话，就需要使用 Vue 提供的 HTML 插值了。这个 HTML 插值其实就是 v-html 指令。

` <div id= "app" > <span v-html=message></span> </div> <script src= "scripts/vue.js" ></script> <script> // 创建Vue的实例对象 var app = new Vue({ el: '#app' , data: { // 存储在Vue的实例对象中 message: '<span style="color: lightcoral">前端课湛</span>' } }); </script> 复制代码`

这样修改代码之后，运行的时候浏览器就会解析这段 HTML 代码了。

![](https://user-gold-cdn.xitu.io/2019/6/5/16b27ff283135b50?imageView2/0/w/1280/h/960/ignore-error/1)

HTML 插值在使用的时候也需要注意一个问题，就是 HTML 插值可能会引起 XSS 攻击。啥？这一回心真真尼玛累啊！XSS 攻击咱就不解释了，好奇的话自行百度吧。

所以，啥意思呢？就是要告诉你啊，如果你要用 HTML 插值的话，那这个数据是不能交给用户提供的。不然，就很危险了！

## v-bind 指令 ##

讲到这儿啊，你是不是觉得 Vue 的模板语法还是挺强大的？可是还不够呢！HTML 标签还有属性呢，这个能不能也同样实现绑定到 Vue 实例的数据呢？答案是肯定的。不是肯定的咱说它干啥，真是的。

那具体怎么来实现呢？这就需要用到 Vue 的 v-bind 指令啦。废话不多说，咱们来看一段示例代码吧。

` <div id= "app" > <span v-bind:title= "message" ></span> </div> <script src= "scripts/vue.js" ></script> <script> new Vue({ el: '#app' , data: { message: 'this is message.' } }); </script> 复制代码`

上面这段代码运行之后，就会把数据绑定到 v-bind 指令的 title 属性上。

![](https://user-gold-cdn.xitu.io/2019/6/5/16b27fffe9cceed9?imageView2/0/w/1280/h/960/ignore-error/1)

牛掰吧？！现在文本内容、HTML 代码解析，还有 HTML 标签的属性都有了。

## 表达式 ##

不仅如此呢！Vue 还允许使用 JavaScript 表达式呢。关于啥是 JavaScript 的表达式，如果你知不道的话，那你为啥要学 Vue 呢？！拉出去枪毙五分钟，再重新去补 JavaScript 基础去！

比如下面这些都是 JavaScript 表达式：

` {{ number + 1 }} {{ ok ? 'YES' : 'NO' }} {{ message.split( '' ).reverse().join( '' ) }} <div v-bind:id= "'list-' + id" ></div> 复制代码`

这里还要说一个需要注意的地方，Vue 支持表达式是有个限制的，就是每个数据的绑定只能支持单个表达式。换句话讲，多个表达式或者语句啥的都不支持呢！

` <!-- 这是语句，不是表达式 --> {{ var a = 1 }} <!-- 流控制也不会生效，请使用三元表达式 --> {{ if (ok) { return message } }} 复制代码`

好吧！一下子说了这么多，这下你也该满意了吧？！