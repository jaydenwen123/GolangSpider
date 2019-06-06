# ElementUI文档中忽略的内容补充 #

## 前言 ##

虽然ElementUI文档已经十分详细，但是难免会有一点点遗漏的地方。本文介绍了笔者使用Element的经验以及文档中忽略或简要介绍的内容（笔者想了好久，就写了一点点）。如果你有什么需要补充的，不妨评论区告诉我吧。

## [官方 FAQ]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2FElemeFE%2Felement%2Fblob%2Fdev%2FFAQ.md ) ##

建议每一个使用 ElementUI 的人都去读一读，官方也很无奈啊。有些可能是被问烦了，有些还是真的挺有用的

eg: 给组件绑定的事件为什么无法触发？

在 Vue 2.0 中，为自定义组件绑定原生事件必须使用 .native 修饰符： ` <my-component @click.native="handleClick">Click Me</my-component>`

从易用性的角度出发，我们对 Button 组件进行了处理，使它可以监听 click 事件： ` <el-button @click="handleButtonClick">Click Me</el-button>`

但是对于其他组件，还是需要添加 .native 修饰符。

## Input 实时响应用户的输入 ##

` <el-input type= "text" v-model= "test" @change= "change" > </ el-input > 复制代码`

使用 change 监听时，input 框的值改变不能触发 change 事件，但是这时如果是 input 输入框失焦确能触发事件。总的来说就是 change 事件只在 input 值改变并且失去焦点的时候才会触发，其它情况不触发事件

change 事件现在仅在输入框失去焦点或用户按下回车时触发，与原生 input 元素一致。如果需要实时响应用户的输入，可以使用 input 事件。

详见其 [更新说明]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2FElemeFE%2Felement%2Fblob%2Fd41afe0183058876d1b2ebcbee82ed11a1029212%2FCHANGELOG.zh-CN.md%23%25E9%259D%259E%25E5%2585%25BC%25E5%25AE%25B9%25E6%2580%25A7%25E6%259B%25B4%25E6%2596%25B0-3 )

## Input 使用 v-model [使用修饰符]( https://link.juejin.im?target=https%3A%2F%2Fcn.vuejs.org%2Fv2%2Fguide%2Fforms.html%23%25E4%25BF%25AE%25E9%25A5%25B0%25E7%25AC%25A6 ) ##

如果要自动过滤用户输入的首尾空白字符，可以给 v-model 添加 trim 修饰符：

` <el-input v-model.trim="input" placeholder="请输入内容"></el-input>`

想自动将用户的输入值转为数值类型，可以给 v-model 添加 number 修饰符：

` <el-input type="number" v-model.number="numberValidateForm.age" autocomplete="off"></el-input>`

## 表单验证 ##

ElementUI 表单验证使用 [async-validator]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fyiminghe%2Fasync-validator ) ，表单元素的 type 有 string，number，boolean，method，regexp，integer，float，array，object，enum，date，url，hex，email

有些还是很常用的，比如 url 验证，date 验证，email 验证。但是文档里面没有写，我有时候会记不住。哈哈哈。所以我写了一个 [表单代码生成器]( https://link.juejin.im?target=https%3A%2F%2Fso-easy.cc%2Fvue-element-form-editor%2F%23%2F ) ，这样就不用去记住了。

## select 远程搜索组件回显 ##

element-ui 当你的选项是固定的时候，它会基于你选中的 value，回显对应的 label，但是远程搜索组件由于 ` options` 不固定，回显就是一个问题。

解决的方法就是传入已选中的值的 ` options` 传入，比如我有一个组件 ` ArticleSelect` ，我选中的 id 值为 [ 1,2 ] ，如果不做处理的话，这个组件就不会回显。仅干巴巴的显示 1，2 两个 tag。但是我可以通过把选中的值的 ` options` (值为 ` [{value:1,label:'第一篇'},{value:2,label:'第二篇'}]` ) 传入这个组件，实现回显显示标题。

但，可能有人就问了，select 组件远程搜索 options 不是会随着搜索的关键词而动态变化么，为什么这样可以？我们看一下 ElementUI select 组件设置选中值的代码：

` setSelected() { // 省略不是多选的情况的代码 // 多选 let result = []; if ( Array.isArray( this.value)) { this.value.forEach( value => { // 注意到这里是push操作 result.push( this.getOption(value)); }); } this.selected = result; // 设置完成之后重新计算选项框的高度 this.$nextTick( () => { this.resetInputHeight(); }); } 复制代码`

由代码可知， Element 设置 选中的值是一个 push 操作，所以 options 后续改变也不会影响我选中的值，完美解决了我的需求

## 自定义 element-ui 样式 ##

> 
> 
> 
> 这一节我是从 [这里](
> https://link.juejin.im?target=https%3A%2F%2Fpanjiachen.github.io%2Fvue-element-admin-site%2Fzh%2Fguide%2Fessentials%2Fstyle.html%23%25E8%2587%25AA%25E5%25AE%259A%25E4%25B9%2589-element-ui-%25E6%25A0%25B7%25E5%25BC%258F
> ) 抄来的，就不自己写了
> 
> 

### CSS 命名空间 ###

现在我们来说说怎么覆盖 element-ui 样式。由于 element-ui 的样式我们是在全局引入的，所以你想在某个页面里面覆盖它的样式就不能加 scoped，但你又想只覆盖这个页面的 element 样式，你就可在它的父级加一个 class，用命名空间来解决问题。

`.article-page { /* 你的命名空间 */ .el-tag { /* element-ui 元素*/ margin-right: 0px; } } 复制代码`

当然也可以使用深度作用选择器 下文会介绍

### 父组件改变子组件样式 [深度选择器]( https://link.juejin.im?target=https%3A%2F%2Fvue-loader.vuejs.org%2Fguide%2Fscoped-css.html%23mixing-local-and-global-styles ) ###

当你子组件使用了 scoped 但在父组件又想修改子组件的样式可以，通过 >>> 来实现：

` <style scoped> .a >>> .b { /* ... */ } </style> 复制代码`

将会编译成

`.a [data-v-f3f3eg9].b { /* ... */ } 复制代码`

如果你使用了一些预处理的东西，如 sass 你可以通过 /deep/ 来代替 >>> 实现想要的效果。

所以你想覆盖某个特定页面 element 的 button 的样式，你可以这样做：

`.xxx-container >>> .el-button{ xxxx } 复制代码`

## 终极技巧 ##

如果有时候你不幸遇到了 ElementUI 的 bug（虽然大部分情况是你自己的问题），给组件添加 key，更新 key 值，强制更新组件，一般可以解决问题。

## 关于我 ##

一个一年小前端，关注我的微信公众号，和我一起交流，我会尽我所能，并且看看我能成长成什么样子吧。交个朋友吧！

![](https://user-gold-cdn.xitu.io/2019/5/16/16ac0e7b1f78bede?imageView2/0/w/1280/h/960/ignore-error/1)