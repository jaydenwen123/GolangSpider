# 如何写一个 vue 插件 #

> 
> 
> 
> 原文地址： [lastingman.com/2019/06/05/…](
> https://link.juejin.im?target=https%3A%2F%2Flastingman.com%2F2019%2F06%2F05%2Fhow-to-write-a-vue-plugin%2F
> )
> 
> 

## Vue 插件的作用 ##

* 添加全局的方法或者属性。
* 添加全局资源，列如 directives， filters，transitions。
* 通过全局混入来添加一些组件选项。
* 添加 Vue 实例方法，通过把它们添加到 ` Vue.prototype` 来实现。
* 一个库，提供自己的 API，同时提供上面提到的一个或多个功能。

## 开发插件 ##

在开发项目的时候，我们一般都用 vue-cli 来避免繁琐的 webpack 配置和 template 配置。但是官方 cli3 现在并不支持搭建 plugin 开发的项目。

还好，已经有大神（Kazupon）走在了我们前面，我们就用现成的 [vue-cli-plugin-p11n]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fkazupon%2Fvue-cli-plugin-p11n ) 。

如果你没有安装 vue-cli，请先安装

` npm i -g @vue/cli 复制代码`

首先，搭建项目

` vue create [your plugin name] && cd [your plugin name] vue add p11n 复制代码`

这样我们就有了一个初始化的插件开发环境。

### install 方法 ###

开发 vue 插件其实就是写一个 install 方法，然后把这个方法暴露出来给你的用户，他们就可以用 ` Vue.use(plugin)` 载入插件了。

> 
> 
> 
> 借用 vue 官方 API 上的解释： 如果插件是一个对象，必须提供 install 方法。如果插件是一个函数，它会被作为 install
> 方法。install 方法调用时，会将 Vue 作为参数传入。 该方法需要在调用 ` new Vue()` 之前被调用。 当 install
> 方法被同一个插件多次调用，插件将只会被安装一次。
> 
> 

` export const install = function ( Vue, options ) { // Vue 就是 vue 实例 // options 就是 Vue.use(plugin，options) 传入的第二个参数 options // 1. 添加全局方法或属性 Vue.myGlobalMethod = function ( ) { // 逻辑... } // 2. 添加全局资源 Vue.directive( 'my-directive' , { bind (el, binding, vnode, oldVnode) { // 逻辑... } ... }) // 3. 注入组件选项 Vue.mixin({ created : function ( ) { // 逻辑... } ... }) // 4. 添加实例方法 Vue.prototype.$myMethod = function ( methodOptions ) { // 逻辑... } } 复制代码`

## 发布插件 ##

p11n 已经帮助我们部署好了大部分 package.json 配置，只需要自己填写好 name，author，license，repository，description，keywords 这几个选项就可以了。

` # login npm npm login # patch version npm version patch # publish npm publish --access public 复制代码`

我自己写了一个非常简单的插件 [vue-chart]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Flastingman%2Fvue-chart ) ，可以作为参考。