# vue，angular，react框架对比 #

首先，我们先了解什么是MVX框架模式？

##### MVX框架模式：MVC+MVP+MVVM #####

1.MVC：Model(模型)+View(视图)+controller(控制器)，主要是基于分层的目的，让彼此的职责分开。 View通过Controller来和Model联系，Controller是View和Model的协调者，View和Model不直接联系，基本联系都是单向的。 用户User通过控制器Controller来操作模板Model从而达到视图View的变化。

2.MVP：是从MVC模式演变而来的，都是通过Controller/Presenter负责逻辑的处理+Model提供数据+View负责显示。 在MVP中，Presenter完全把View和Model进行了分离，主要的程序逻辑在Presenter里实现。 并且，Presenter和View是没有直接关联的，是通过定义好的接口进行交互，从而使得在变更View的时候可以保持Presenter不变。 MVP模式的框架：Riot,js。

3.MVVM：MVVM是把MVC里的Controller和MVP里的Presenter改成了ViewModel。Model+View+ViewModel。 View的变化会自动更新到ViewModel,ViewModel的变化也会自动同步到View上显示。 这种自动同步是因为ViewModel中的属性实现了Observer，当属性变更时都能触发对应的操作。 MVVM模式的框架有：AngularJS+Vue.js和Knockout+Ember.js后两种知名度较低以及是早起的框架模式。

#### Vue.js是什么？ ####

看到了上面的框架模式介绍，我们可以知道它是属于MVVM模式的框架。那它有哪些特性呢？
其实Vue.js不是一个框架，因为它只聚焦视图层，是一个构建数据驱动的Web界面的库。
Vue.js通过简单的API（应用程序编程接口）提供高效的数据绑定和灵活的组件系统。
Vue.js的特性如下：

> 
> 
> 
> 1.轻量级的框架
> 2.双向数据绑定
> 3.指令
> 4.插件化
> 
> 
> 

#### Vue.js与其他框架的区别？ ####

##### 1.与AngularJS的区别
#####

相同点：

> 
> 
> 
> * 都支持指令：内置指令和自定义指令。
> 
> * 都支持过滤器：内置过滤器和自定义过滤器。
> 
> * 都支持双向数据绑定。
> 
> * 都不支持低端浏览器。
> 
> 
> 
> 

不同点：

> 
> 
> 
> * 1.AngularJS的学习成本高，比如增加了Dependency Injection特性，而Vue.js本身提供的API都比较简单、直观。
> 
> * 2.在性能上，AngularJS依赖对数据做脏检查，所以Watcher越多越慢。
> 
> * Vue.js使用基于依赖追踪的观察并且使用异步队列更新。所有的数据都是独立触发的。
> 
> * 对于庞大的应用来说，这个优化差异还是比较明显的。
> 
> 
> 
> 

##### 2.与React的区别
#####

相同点：

> 
> 
> 
> * React采用特殊的JSX语法，Vue.js在组件开发中也推崇编写.vue特殊文件格式，对文件内容都有一些约定，两者都需要编译后使用。
> * 中心思想相同：一切都是组件，组件实例之间可以嵌套。
> * 都提供合理的钩子函数，可以让开发者定制化地去处理需求。
> * 都不内置列数AJAX，Route等功能到核心包，而是以插件的方式加载。
> * 在组件开发中都支持mixins的特性。
> 
> 
> 

不同点：

> 
> 
> 
> * React依赖Virtual DOM,而Vue.js使用的是DOM模板。React采用的Virtual DOM会对渲染出来的结果做脏检查。
> * Vue.js在模板中提供了指令，过滤器等，可以非常方便，快捷地操作DOM。
> 
> 
>