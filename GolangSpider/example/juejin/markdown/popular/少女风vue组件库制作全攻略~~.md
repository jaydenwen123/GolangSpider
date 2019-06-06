# 少女风vue组件库制作全攻略~~ #

## 预览 ##

[组件库官网]( https://link.juejin.im?target=https%3A%2F%2Ffirenzia.github.io%2Fsakura-ui%2F )
[github地址]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2FFirenzia%2Fsakura-ui%2F )
如果喜欢各位小哥哥小姐姐给个小星星鼓励一下哈， 请勿在生产环境中使用，供学习&交流~~

![](https://user-gold-cdn.xitu.io/2019/5/14/16ab1f527248e169?imageslim)

![](https://user-gold-cdn.xitu.io/2019/5/14/16ab1f8f4d386114?imageslim)

## 完整项目目录结构 ##

git clone到本地安装依赖后，执行npm run serve进行本地组件库开发，npm run docs:dev进行组件库官网开发。一般在src/demo.vue进行单个组件测试通过后，再引入到.vuepress/components中放入组件库官网。

` ├─docs // vuepress开发目录 │ ├─.vuepress │ │ ├─components // 在markdown中可以使用的vue组件 │ │ ├─dist // vuepress打包目录 │ │ │ ├─assets │ │ │ │ ├─css │ │ │ │ ├─img │ │ │ │ └─js │ │ │ ├─components │ │ │ │ ├─basic │ │ │ │ ├─form │ │ │ │ ├─navigation │ │ │ │ ├─notice │ │ │ │ └─other │ │ │ └─guide │ │ │ │ │ ├─config.js // vurepess配置修改入口,包括左边sidebar,右上方nav导航菜单,favicon等 │ │ ├─style.style // 覆盖vuerpress默认主题样式 │ │ └─public //公共资源入口，如favicon │ ├─static │ │ ├─img │ │ └─js │ └─views // vuepress视图文件，格式是markdown │ ├─components │ │ ├─basic │ │ ├─form │ │ ├─navigation │ │ ├─notice │ │ └─other │ ├─design │ │ └─color │ └─guide ├─src // 组件库源码目录 │ ├─button │ ├─cascader │ ├─collapse │ ├─container │ ├─datepicker │ ├─form │ ├─icon │ ├─layout │ ├─notice │ ├─plugins │ ├─slide │ ├─tab │ ├─step │ ├─sticky │ └─index.js // 组件库源码组件入口文件,执行npm run build的目标文件 ├─package.json // 与npm发布相关，记录版本号，包入口文件地址 复制代码`

## 学习组件库制作会收获 ##

* 学习组件封装技能，良好的接口设计, 掌握组件设计套路
* 夯实js/css基础
* 深入对vue的理解

## 制作流程 ##

* 组件设计/开发
* 发布npm
* 制作官网展示

## 组件设计/开发 ##

### 频繁涉及到的vue api包括 ###

* $children ： 获取当前组件子组件。
* $parent： 获取当前组件父组件。
* $options: 用于当前 Vue 实例的初始化选项, 可以用此选项获得组件的name。
* $refs： 一个对象，持有注册过 ref 特性 的所有 DOM 元素和组件实例。
* $el: Vue 实例使用的根 DOM 元素。
* provide & inject ：这对选项需要一起使用，允许一个祖先组件向其所有子孙后代注入一个依赖,注意不是响应式的。注入的对象可以是个vue实例的eventBus。
* $on： 组件监听自定义事件。
* $emit： 组件触发自定义事件。
* .sync：语法糖，单向数据流中，父组件监听到子组件修改props的意图后父组件修改传入的props， 用了.sync不需要显式在父组件监听组件内部触发的自定义事件去修改值, 父组件只要写:x.sync="bindValue", 注意此时子组件触发的事件必须是"update:x"此语法糖才生效。
* updated 生命周期钩子函数， 由于数据更改导致的虚拟 DOM 重新渲染和打补丁 ，在这之后会调用该钩子, 在父子组件通信可能用到。
* beforeDestoryed/ destory 生命周期钩子函数，destory后组件的所有的事件监听器会被移除。 注意:如果是自己在组件内部对dom增加了事件监听，组件销毁的时候需要自己手动接触自己另外添加上去的监听程序。而且组件销毁，dom元素还被保留在页面，需要手动清除，可以调用原生js api, node.remove()清除dom节点。

### 原生js api包括： ###

* target.addEventListener(type, listener[, useCapture])/removeEventListener 由于这是 DOM2 规范的基本内容，几乎所有浏览器都支持这个，而且不需要特殊的跨浏览器兼容代码。
* Node.contains()返回的是一个布尔值，来表示传入的节点是否为该节点的后代节点。多用于事件监听判断是否点击了目标区域。
* window.scrollY 获取文档垂直方向的滚动距离。
* Element​.get​Bounding​Client​Rect() 返回元素的大小及其相对于视口的位置，返回一个对象，包括width/height/left/right/top/bottom。多用于计算定位。

### 技术点总结 ###

组件设计的思想包括单数据流/ eventBus事件中心，核心是组件通信。

* 单数据流: 数据的改变是单向的，即通过props的方式，只能让父组件来修改数据，子组件不能主动修改props。这样的例子如在collapse/tab/slide组件中，让父组件来控制选中的值。单向数据流的思想让数据修改更好设计，逻辑更加清晰。
* vue插件开发： 什么时候用插件开发? 当组件不是显式在代码中被调用，不是直接写在template中，而是通过调用Vue原型链上的方法被挂载到文档中。 比如modal模态框/toast弹窗。插件设计的基本思路是暴露一个install方法，这个方法中在vue原型链上增加一个自定义的方法X， X中引入组件a，通过Vue.extend(a)获得组件构造器Constructor, 在通过new Constructor({propsData})获得组件实例vm, 再挂载组件实例到文档中。

` import modal from '../notice/modal' export default { install (vue, options) { const Construtor = vue.extend(modal) let modalVm // 保证全局只有一个modal实例 let lastOption vue.prototype. $modal = (options) => { if (lastOption !== JSON.stringify(options)) { //! modalVm modalVm = new Construtor({ propsData: options }) modalVm. $mount () document.body.append(modalVm. $el ) } lastOption = JSON.stringify(options) modalVm.isVisible = true } } } 复制代码`

* eventBus： 什么时候用eventBus? 当状态的变化需要多个子组件被通知。如tab组件中，当选中的值发生变化，tab-head需要感知变化让提示的短线做个动画滑到选中的标签下，tab-item需要感知变化让文字变成选中样式，tab-pane需要感知变化让选中的面板出现。
* 递归：在级联组件的设计中用到。类似函数fn中用setTimout(fn,millseconds)调用自己实现setInterval的递归效果。 组件只要内部提供name属性，就可以递归地调用自身。 允许组件模板递归地调用自身。通过提供 name 选项，便于调试，在控制台可以看到可以获得更有语义信息的组件标签。
* 媒体查询 &flex布局：响应式布局的原理是媒体查询和百分比布局，介于某个尺寸的时候某个类名生效;跟布局相关的大部分用到flex,非常好用。详细看 [阮一峰老师教程]( https://link.juejin.im?target=http%3A%2F%2Fwww.ruanyifeng.com%2Fblog%2F2015%2F07%2Fflex-grammar.html%3Futm_source%3Dtuicool )

+----------+----------------------+----------+-------------+----------+----------------------+------+-------------------+
| 组件类型 |         组件         | 单数据流 | VUE插件开发 | EVENTBUS | 原生JS操作DOM & 事件 | 递归 | 媒体查询&FLEX布局 |
+----------+----------------------+----------+-------------+----------+----------------------+------+-------------------+
| 基础     | button按钮           | -        | -           | -        | -                    | -    | -                 |
| 基础     | icon图标             | -        | -           | -        | -                    | -    | -                 |
| 基础     | grid网格             | -        | -           | -        | -                    | -    | yes               |
| 基础     | layout布局           | -        | -           | -        | -                    | -    | yes               |
| 表单     | input输入框          | -        | -           | -        | -                    | -    | -                 |
| 表单     | cascader级联选择器   | yes      | -           | -        | -                    | yes  | -                 |
| 表单     | form表单             | -        | -           | -        | -                    | -    | -                 |
| 表单     | datepicker日期选择器 | -        | -           | -        | yes                  | -    | -                 |
| 导航     | tab标签页            | -        | -           | yes      | -                    | -    | -                 |
| 导航     | step步骤调           | -        | -           | -        | -                    | -    | -                 |
| 通知     | toast提示            | -        | yes         | -        | yes                  | -    | -                 |
| 通知     | popover弹出框        | -        | -           | -        | yes                  | -    | -                 |
| 通知     | modal模态框          | -        | yes         | -        | yes                  | -    | -                 |
| 其他     | collapse折叠面板     | yes      | -           | yes      | -                    | -    | -                 |
| 其他     | slide轮播图          | yes      | -           | -        | -                    | -    | -                 |
| 其他     | sticky粘滞           | -        | -           | -        | -                    | -    | -                 |
+----------+----------------------+----------+-------------+----------+----------------------+------+-------------------+

### 组件设计三要素 ###

* props：可以参考饿了么或者antd, 需要从用户的角度考虑怎么使用方便和扩展性好，一般需要校验类型和有效值，设置默认值。
* slot：插槽内容分发，使用作用域插槽让slot也可以获得组件内部方法，让用户自定义的内容页能调用组件内部方法，比如popover弹出框中用户想自己加个按钮手动调用关闭。
* event: 组件事件。从用户角度考虑，比如datepicker组件中用户想在日期面板被打开或这关闭的时候进行操作。这种一般用在交互类UI组件。

#### 举个例子 ####

复杂组件datepicker开发思路

![](https://user-gold-cdn.xitu.io/2019/5/14/16ab3cec51014934?imageslim)

* 

在原有的popover组件上开发
点击一个元素A(输入框)后可以弹出元素B（日期面板）

* 

生成日期面板
生成7*6=42个日期，6行是为了确保一个月都能在面板上完整显示。这里计算最方便的做法是用时间戳，计算出这个月第一天时间戳和这一天周几，就可以一次性计算出这42个日期。不用算上个月下个月分三段算，这样的问题是还要考虑边界情况，如刚好出现上一年下一年等，麻烦容易出bug。这42个日期我们在computed用visibleDays表示。

` visibleDays () { let { year, month } = this.display let defaultObj = new Date(year, month, 28) var curMonthFirstDay = helper.getMonthFirstDay(defaultObj) var curMonthFirstDayDay = helper.getDay(curMonthFirstDay) === 0 ? 7 : helper.getDay(curMonthFirstDay) let x = curMonthFirstDayDay - 1 // 前面需要补多少位 var arr = [] for ( let i = 0; i < 42; i++) { arr.push(new Date(curMonthFirstDay.getTime() + (-x + i) * 3600 * 24 * 1000)) } return arr }, 复制代码`
* 

props接受value, 类型是date
日期面板上的日期渲染的时候加上一个计算的class, 分别加上'today','selected-date','available','prev-month','next-month'，进行样式上的区分

* 

实现选中日期
告诉父组件修改数据意图让父组件修改传入的props，对应使用我们组件的时候使用, 这里的基础知识是组件上的v-model是个语法糖，v-model="x"会被解析成:value="a" @input="a=$event"。同时面板上输入框显示的数据也要跟着变化，所以这里用计算属性，如在computed中用formattedValue表示。

` formattedValue: { return this.value instanceof Date ? helper.getFormatDate(this.value) : '' } 复制代码`
* 

实现点击上一年/月，下一年/月
我们需要知道当前展示的是哪一年哪一个月，这个数据是组件内部维护的，所以在data申明一个display对象

` display: { year: (this.value && this.value.getFullYear()) || new Date().getFullYear(), month: (this.value && this.value.getMonth()) || new Date().getMonth() } 复制代码`

点击的时候即修改display对象的year/month，因为visibleDays也是计算属性，依赖display对象，所以点击上一年/月，下一年/月，渲染的日期也跟着变。

* 

实现选择年
年面板的制作，生成12个年，点击第1（12）个年渲染出上（下）12个年。这里只需要给渲染出来的年的第一个和最后一个dom元素绑定事件，事件监听程序传入当前点击的元素的值，即可计算出上或下一个12年。 同理点击年的时候用$emit通知父组件修改value

* 

实现选择月
直接写死12个月份，同理点击月的时候用$emit通知父组件修改value

* 

增加住面板上【今天】和【清空】的按钮
点击的时候用$emit通知父组件修改value，new Date()和''

* 

细节处理
用户选中完日期后要关闭面板 用户选了年后点击周围空白区域日期面板关闭，第二次点击进来应该默认展示日面板

* 

用户可以修改输入框里面的值，需要判断有效性
有效的话$emit通知父组件改值，无效的话当失去焦点的时候变回原来的值，这里需要用原生js去给input修改value。 注意这里直接改formattedValue的话无效，虽然输入框的值绑定了:value="formattedValue",但是因为formattedValue是计算属性，依赖于this.value，在用户输入无效值的情况下this.value不会改变，因此界面不会被更新，所以需要手动改value的值。

` set ValueManually ( $event ) { if (!helper.isValidDate( $event )) { this. $refs.inputWrapper. $refs.input.value = this.isDate(this.value) ? helper.getFormatDate(this.value) : '' return } this. $emit ( 'input' , new Date( $event )) } 复制代码`
* 

完善
给弹出日期面板和关闭日期面板增加组件自定义事件, 即调用$emit触发'showDatepicker'和'closeDatepicker'事件。

## 发布npm ##

* 使用vue cli3 的库模式打包代码，修改package.json 中的"build": "vue-cli-service build --target lib --name sakura src/index.js"，打包后输出umd构建版本, 参考 [vue cli]( https://link.juejin.im?target=https%3A%2F%2Fcli.vuejs.org%2Fzh%2Fguide%2Fbuild-targets.html%23%25E5%25BA%2593 ) 。 什么是umd? 统一模块定义，可以兼容common.js(node端规范)/ AMD(浏览器端规范)/ ES6(node端不完全支持)等多种模块化方案，确保代码在各种环境下能被运行。 ` File Size Gzipped dist/sakura.umd.min.js 13.28 kb 8.42 kb dist/sakura.umd.js 20.95 kb 10.22 kb dist/sakura.common.js 20.57 kb 10.09 kb dist/sakura.css 0.33 kb 0.23 kb 复制代码`
* 在package.json指明模块入口"main":"dist/sakura.umd.min.js" ` "name" : "heian-sakura-ui" , "version" : "0.0.6" , "private" : false , "main" : "dist/sakura.umd.min.js" , "description" : "an UI framework based on Vue.js" , 复制代码`
* 在npm 上注册一个用户
* 在命令行输入，注意每次发布都要修改package.json中的 "version": "0.0.x"，"private"必须设置成false才能发布 ` npm adduser // 提示输入注册的用户名 npm publish 复制代码`

## 官网制作 ##

使用vue press

* 

在原有项目中使用

` # 安装依赖 npm install -D vuepress # 创建一个 docs 目录 mkdir docs 复制代码`

在package.json中进行脚本配置

` { "scripts" : { "docs:dev" : "vuepress dev docs" , "docs:build" : "vuepress build docs" } } 复制代码`

然后运行npm run docs:dev即可访问

* 

简单配置
在docs/.vuepress下新建文件config.js

` module.exports = { base: '/sakura-ui/' , title: 'Sakura UI' , description: 'Inspiration from heian sakura' , head: [ [ 'link' , { rel: 'icon' , href: '/favicon.ico' }] ], themeConfig: { nav: [ { text: 'Home' , link: '/' }, { text: 'Github' , link: 'https://github.com/Firenzia/sakura-ui/' }, ], sidebar: [ { title: '开发指南' , collapsable: true , children: [ 'views/guide/install.md' , 'views/guide/get-started.md' ] }, { title: '设计' , collapsable: true , children: [ 'views/design/color/' , ] }, { title: '组件' , collapsable: true , children: [ 'views/components/basic/' , 'views/components/form/' , 'views/components/navigation/' , 'views/components/notice/' , 'views/components/other/' ] }, ] } } 复制代码`
* 

使用vue组件
官网中提到，所有在 .vuepress/components 中找到的 *.vue 文件将会自动地被注册为全局的异步组件，可以在markdown中引用, vue文件中的代码高亮我用的是vue-highlightjs [查看这里]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fmetachris%2Fvue-highlightjs )

* 

编写文档
由于所有的页面在生成静态 HTML 时都需要通过 Node.js 服务端渲染，对于SSR 不怎么友好的组件（比如包含了自定义指令），你可以将它们包裹在内置的 ClientOnly 组件中，而且注意因为是ssr,组件内部beforeCreate, created生命周期钩子函数访问不到浏览器 / DOM 的 API，只能在beforeMount和mounted中调用。

` --- title: 'Basic 基础' sidebarDepth: 2 --- ## Icon 图标 <ClientOnly> <sakura-icon/> <font size=5>Attributes</font> | 参数| 说明 | 类型 | 可选值 | 默认值 | | :------ | ------ | ------ | ------ | ------ | | name | 图标名称 | string |- | - | | color | 图标颜色, 支持常见颜色和十六进制颜色 | string |- | - | </ClientOnly> 复制代码`
* 

覆盖默认主题样式
在.vuepress下新增style.styl进行覆盖。

* 

部署到github
官网上介绍的很清楚， [点这里]( https://link.juejin.im?target=https%3A%2F%2Fvuepress.vuejs.org%2Fzh%2Fguide%2Fdeploy.html%23github-pages ) 。 在项目根目录下新增deploy.sh,windows下直接命令行运行./deploy.sh即可发布到github pages上。

## 结语 ##

如果你能看到这里，非常感谢，第一次写文章，希望大家多多提出意见。组件库还有很多细节需要完善，比如里面css的类名命名我没做的很规范，大部分组件都是自己测试没有测到复杂或特殊场景，还有很多功能还没支持。通过这段时间制作组件库，自己的技术有了一定提升，官网的展示融入了自己的一点想法和设计，希望大家喜欢~~ 谢谢！