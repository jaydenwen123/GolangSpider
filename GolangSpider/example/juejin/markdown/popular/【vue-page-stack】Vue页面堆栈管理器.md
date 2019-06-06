# 【vue-page-stack】Vue页面堆栈管理器 #

## Vue页面堆栈管理器 ##

> 
> 
> 
> A vue page stack manager Vue页面堆栈管理器 [vue-page-stack](
> https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fhezhongfeng%2Fvue-page-stack
> )
> 
> 

![](https://user-gold-cdn.xitu.io/2019/6/4/16b221ba81f9e7b1?imageslim)

> 
> 
> 
> 示例展示了一般的前进、后退（有activited）和replace的场景，同时还展示了同一个路由可以存在多层的效果（输入必要信息）
> 
> 

**目前版本还没有经过整体业务的测试，欢迎有同样需求的进行试用**

[预览]( https://link.juejin.im?target=https%3A%2F%2Fhezhongfeng.github.io%2Fvue-page-stack-example%2F )

[源码]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fhezhongfeng%2Fvue-page-stack )

## 需求分析 ##

由于重度使用了Vue全家桶在 ` web App` 、公众号和原生Hybrid开发，所以很自然的会遇到页面跳转与回退这方面的问题。

场景举例：

* 列表页进入详情页，然后回退
* 某操作页A需要在下一页面B选择，选择后需要退回到A页面（A页面还要知道选择了什么）
* 在任意页面进入到登录页面，登录或者注册成功后返回到原页面，并且要保证继续回退是不会到登陆页面的
* 支持浏览器的 ` back` 和 ` forward` (微信或者小程序很有用)
* 在进入、退出或者某些特殊页面的时候添加一些动画，比如模仿ios的默认动画（进入是页面从右向左平移，退出是页面从左向右平移）

## 尝试过的方法 ##

尝试了以下方法，但是都没有达到我的预期

### keep-alive ###

一般是使用两个 ` router-view` 通过route信息和keep-alive控制页面是否缓存，这样存在两个问题：

* keep-alive对相同的页面只会存储一次，不会有两个版本的相同页面
* 两个router-view之间没有办法使用 ` transition` 等动画

### CSS配合嵌套route ###

曾经在查看 ` cube-ui` 的例子的时候，发现他们的例子好像解决了页面缓存的问题，我借鉴(copy)了他们的处理方式，升级了一下，使用CSS和嵌套route的方式实现了基本的需求。 但是也有缺点：

* 我必须严格按照页面的层级来写我的route
* 很多页面在多个地方需要用到，我必须都得把路由配上(例如商品详情页面，会在很多个地方有入口)

## 功能说明 ##

* 在vue-router上扩展，原有导航逻辑不需改变
* ` push` 或者 ` forward` 的时候重新渲染页面，Stack中会添加新渲染的页面
* ` back` 或者 ` go(负数)` 的时候不会重新渲染，从Stack中读取先前的页面，会保留好先前的内容状态，例如表单内容，滚动条滑动的位置等
* ` back` 或者 ` go(负数)` 的时候会把不用的页面从Stack中移除
* ` replace` 会更新Stack中页面信息
* 回退到之前页面的时候有activited钩子函数触发
* 支持浏览器的后退，前进事件
* 支持响应路由参数的变化，例如从 /user/foo 导航到 /user/bar，组件实例会被复用
* 可以在前进和后退的时候添加不同的动画，也可以在特殊页面添加特殊的动画

## 安装和用法 ##

### 安装 ###

` npm install vue-page-stack # OR yarn add vue-page-stack 复制代码`

### 使用 ###

` import Vue from 'vue' import VuePageStack from 'vue-page-stack' ; // vue-router是必须的 Vue.use(VuePageStack, { router }); 复制代码` ` // App.vue <template> <div id= "app" > <vue-page-stack> <router-view ></router-view> </vue-page-stack> </div> </template> <script> export default { name: 'App' , data () { return { }; }, components: {}, created () {}, methods: {} }; </script> 复制代码`

## API ##

### 注册 ###

注册的时候可以指定VuePageStack的名字和keyName，一般不需要

` Vue.use(VuePageStack, { router, name: 'VuePageStack' , keyName: 'stack-key' }); 复制代码`

### 前进和后退 ###

想在前进、后退或者特殊路由添加一些动画，可以在 ` router-view` 的页面通过watch ` $route` ，通过 ` stack-key-dir(自定义keyName这里也随之变化)` 参数判断此时的方向，可以参考 [实例]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fhezhongfeng%2Fvue-page-stack-example )

## 相关说明 ##

### keyName ###

为什么会给路由添加 ` keyName` 这个参数，是为了支持浏览器的后退，前进事件，这个特点在微信公众号和小程序很重要

### 原理 ###

获取当前页面Stack部分参考了keep-alive的部分

## 结束语 ##

> 
> 
> 
> 念念不忘，必有回响
> 
> 

这个插件存在我心中很久了，断断续续做了好久，终于被我搞定了，真的非常开心。

目前版本还没有经过整体业务的测试，欢迎有同样需求的进行试用，有任何的意见或者建议，欢迎在 [Github]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fhezhongfeng%2Fvue-page-stack ) 提issue和PR，感谢你的支持和贡献。

这个插件同时借鉴了 [vue-navigation]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fzack24q%2Fvue-navigation ) 和 [vue-nav]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fnearspears%2Fvue-nav ) ，很感谢他们给的灵感。