# Vue 框架原理相关知识点 #

## vue生命周期原理，钩子函数 ##

create和mounted 的区别

参考

* [juejin.im/post/5afd7e…]( https://juejin.im/post/5afd7eb16fb9a07ac5605bb3 )
* [juejin.im/post/5c6d48…]( https://juejin.im/post/5c6d48e36fb9a049eb3c84ff )

![image](https://user-gold-cdn.xitu.io/2019/6/5/16b269796c87dd13?imageView2/0/w/1280/h/960/ignore-error/1)

![image](https://user-gold-cdn.xitu.io/2019/6/5/16b269862e6f3331?imageView2/0/w/1280/h/960/ignore-error/1)

## MVVM框架的设计理念 ##

## 为什么选择vue ##

为什么使用vue，首先要看和其他框架React/Angular的对比

### React ###

React 的特别是使用 JSX，有些人喜欢用，有些人不喜欢？看它的语法就知道

* 一个render函数，里面又放html代码，又放 JS 代码。逻辑不能使用 if-else，只能使用一堆三元运算符。 css也可以当成对象属性放进去，揉在一块，虽然最后他们会编译成纯JS，反正我个人是比较喜欢 JS/CSS/HTML 分开写。
* 过渡的工具依赖
` ReactDOM.render( < div > < h1 > {1+1} </ h1 > </ div > , document.getElementById( 'example' ) ); var myStyle = { fontSize : 100 , color : '#FF0000' }; ReactDOM.render( < h1 style = {myStyle} > 菜鸟教程 </ h1 > , document.getElementById( 'example' ) ); ReactDOM.render( < div > < h1 > {i == 1 ? 'True!' : 'False'} </ h1 > </ div > , document.getElementById( 'example' ) ); 复制代码`

### Vue ###

**特点**

* 从React那里借鉴了组件化、prop、单向数据流、性能、虚拟渲染，并意识到状态管理的重要性。
* 从Angular那里借鉴了模板，并赋予了更好的语法，以及双向数据绑定（在单个组件里）。
* 它不强制使用某种编译器，所以你完全可以在遗留代码里使用Vue，并对之前乱糟糟的jQuery代码进行改造。

**不足**

* 模板的运行时错误描述不够直观
* 年轻，社区生态不够完善

### 为什么使用框架 ###

参考

* [www.zcfy.cc/article/the…]( https://link.juejin.im?target=https%3A%2F%2Fwww.zcfy.cc%2Farticle%2Fthe-deepest-reason-why-modern-javascript-frameworks-exist )
* [medium.com/dailyjs/the…]( https://link.juejin.im?target=https%3A%2F%2Fmedium.com%2Fdailyjs%2Fthe-deepest-reason-why-modern-javascript-frameworks-exist-933b86ebc445 )
* 现代 js 框架解决的主要问题是保持 UI 与状态同步。
* 使用原生 JavaScript 编写复杂、高效而又易于维护的 UI 界面几乎是不可能的。
* Web components 并未提供解决同步问题的方案。
* 使用现有的虚拟 DOM 库去搭建自己的框架并不困难。但并不建议这么做！

## 双向绑定原理，diff算法内部实现 ##

双向绑定原理：依赖收集、发布订阅

## vue事件机制 ##

## 从template转换成真实DOM的实现机制 ##

## nextTick原理 ##