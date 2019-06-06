# 基于vue做一个禅道主页拖拽效果 #

# 不bb先看效果 #

![预览图](https://user-gold-cdn.xitu.io/2019/5/26/16af4728aaae8fec?imageslim)

[源码地址]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fhecun0000%2Fvue-dnd-kon )

# bb两句 #

最近在做一个基于vue的后台管理项目。平时项目进度统计就在上禅道上进行。so~ 然后领导就感觉这个拖拽效果还行，能不能加到咱们项目里面。 既然领导发话，那就开干。。

所有技术：vue + vuedraggable

拖动的实现基于 [vuedraggable]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2FSortableJS%2FVue.Draggable ) 的插件开发。

主页为两栏流式布局，每一个组件可以在上下拖动，也可以左右拖动。

![页面布局](https://user-gold-cdn.xitu.io/2019/5/26/16af480b7dcc4314?imageView2/0/w/1280/h/960/ignore-error/1)

# 基本步骤 #

## 布局 ##

这块布局为最为普通的两栏布局，这里采用flex布局。左边自适应，右边为固定宽。

`.layout-container { display : flex;.left { flex : 1 ; margin-right : 40px ; }.right { width : 550px ; } } 复制代码`

## 拖拽实现 ##

这里使用 ` vuedraggable` 插件。需要在组件里面引入使用。 ` draggable` 相当于拖拽容器，这块很明显需要两个拖拽的容器。所以分别在 `.left` `.right` 中添加两个拖拽容器。在默认情况下，这里已经可以进行拖拽了。插件的效果还是很强大。

` <div class= "layout-container" > <!--左栏--> <div class= "left" > <draggable v-bind= "dragOptions" class= "list-group" :list= "item" > // ... 拖拽元素或组件 </draggable> </div> <!--右栏--> <div class= "right" > <draggable v-bind= "dragOptions" class= "list-group" :list= "item" > // ... 拖拽元素或组件 </draggable> </div> </div> <script> import draggable from "vuedraggable" ; export default { components: {draggable}, computed: { dragOptions () { return { animation: 30, handle: ".drag-handle" , group: "description" , ghostClass: "ghost" , chosenClass: "sortable" , forceFallback: true }; } } }; </script> 复制代码`

但是， 和我想要的效果还是相差一点。

### 左右拖动 与 仅标题栏拖动 ###

这块只需要配置相关的配置项就可以比较简单。 左右拖动需要给拖拽容器指定相同的 ` group` 属性。指定标题元素拖动需要配置 ` handle` 为可拖动元素的选择器名称。

下面简单介绍下常用的配置项：

* **disabled** ：boolean 定义是否此sortable对象是否可用，为true时sortable对象不能拖放排序等功能，为false时为可以进行排序，相当于一个开关；
* **group** : 用处是为了设置可以拖放容器时使用，若两个容器该配置项相同，则可以相互拖动；
* **animation** ：number 单位：ms，定义排序动画的时间；
* **handle** ：selector 格式为简单css选择器的字符串，使列表单元中符合选择器的元素成为拖动的手柄，只有按住拖动手柄才能使列表单元进行拖动；
* **filter** ：selector 格式为简单css选择器的字符串，定义哪些列表单元不能进行拖放，可设置为多个选择器，中间用“，”分隔；
* **draggable** ：selector 格式为简单css选择器的字符串，定义哪些列表单元可以进行拖放
* **ghostClass** ：selector 格式为简单css选择器的字符串，当拖动列表单元时会生成一个副本作为影子单元来模拟被拖动单元排序的情况，此配置项就是来给这个影子单元添加一个class，我们可以通过这种方式来给影子元素进行编辑样式；
* **chosenClass** ：selector 格式为简单css选择器的字符串，当选中列表单元时会给该单元增加一个class；
* **forceFallback** ：boolean 如果设置为true时，将不使用原生的html5的拖放，可以修改一些拖放中元素的样式等；
* **fallbackClass** ：string 当forceFallback设置为true时，拖放过程中鼠标附着单元的样式；

采用相关配置如下：

` computed: { dragOptions() { return { animation : 30 , handle : ".drag-handle" , group : "description" , ghostClass : "ghost" , chosenClass : "sortable" , forceFallback : true }; } } 复制代码`

### 拖动时样式调整 ###

在拖动的时候，我们需要做三个事情。拖动时，拖动元素只显示标题栏，两栏内列表只显示标题元素以及将要移动的位置变灰。

* 

拖动元素只显示标题栏： 在默认情况下，会开启 ` html5` 元素的拖动效果。这里明显不需要。 ` forceFallback` 改为 ` false` 则可以关闭 ` html5` 的默认效果。顺便通过 ` chosenClass: "sortable"` 修改拖动元素class 类名。直接用css进行隐藏

`.sortable {.component-box { display : none; height : 0 ; } } 复制代码`
* 

两栏内列表只显示标题元素 这里我借助两个事件实现。

* onStart：function 列表单元拖动开始的回调函数
* onEnd：function 列表单元拖放结束后的回调函数

` <div class= "layout-container" :class= "{drag:dragging}" > //... </div> 复制代码` ` data() { return { dragging : false }; }, methods : { onStart() { this.dragging = true ; }, onEnd() { this.dragging = false ; } } 复制代码` `.drag {.component-box { display : none; } } 复制代码`

在开始拖动的时候给 `.layout-container` 添加 `.drag` 的 class 名。拖动结束时，移除class名。

* 

将要移动的位置变灰
这里需要用到上面 ` ghostClass: "ghost"` 配置项。并添加相应的css。

`.ghost {.drag-handle { background : rgb( 129 , 168 , 187 ); } } 复制代码`

好了基本已经实现了。。。

![](https://user-gold-cdn.xitu.io/2019/5/26/16af4ba45784237d?imageView2/0/w/1280/h/960/ignore-error/1)

### 展示动态组件 ###

接下来就是数据的动态展示了。 这里需要vue中的动态组件了。。附上官方文档连接 [点击查看]( https://link.juejin.im?target=https%3A%2F%2Fcn.vuejs.org%2Fv2%2Fguide%2Fcomponents-dynamic-async.html ) 。

然后里面每个拖动的元素的内容都写成组件，搭配动态组件实现自由拖动。

` // 将所用组件引入 import { timeline, calendar, welcome, carousel, imgs, KonList } from "@/components/DragComponents" ; components: { draggable, timeline, calendar, welcome, carousel, imgs, KonList } 复制代码`

配合 ` v-for` 对数据进行循环，然后进行动态展示。

` <component :is= "element.name" /> 复制代码`

这块涉及到数据格式相关的，可以直接看文末的代码。。。 这里就就不展开说了。。

## 数据保持 ##

在拖动结束后，我们需要将拖动的顺序缓存在前端，当下次进入后，可以继续使用拖动后的数据。

` // 获取新的布局 getLayout() { let myLayout = JSON.parse( window.localStorage.getItem( "kon" )); if (!myLayout || Object.keys(myLayout).length === 0 ) myLayout = this.layout; const newLayout = {}; for ( const side in myLayout) { newLayout[side] = myLayout[side].map( i => { return this.componentList.find( c => c.id === i); }); } this.mainData = newLayout; }, // 设置新的布局 setLayout() { const res = {}; for ( const side in this.mainData) { const item = this.mainData[side].map( i => i.id); res[side]=item; } window.localStorage.setItem( "kon" , JSON.stringify(res)); } 复制代码`

这样我只需要在 ` mounted` 中获取新的布局。。

` mounted() { this.getLayout(); } 复制代码`

在拖动结束后，设置新的布局

` onEnd() { this.dragging = false ; this.setLayout(); } 复制代码`

在项目中，还是建议配合后端进行用户布局的数据存储，每次拖动后将新的布局数据请求接口保存在数据库，同时存入缓存中。当再次进入页面的时候，读取缓存中的数据，没有的话请求后端的接口拿到用户的布局，然后再次存入缓存中。有的话直接读取缓存中的数据。

# 最后说两句 #

其实上面的效果也不是特别难，简单花点时间，看看相关文档，就能做出来，，记录在掘金上面，只是想和大家分享我的思路。同时希望和大家一起交流，一起进步。

![](https://user-gold-cdn.xitu.io/2019/5/26/16af4d5c43e1a09a?imageView2/0/w/1280/h/960/ignore-error/1)

**生活不易，大家加油**

附上源码： [项目地址]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fhecun0000%2Fvue-dnd-kon )

` < template > < div :class = "{drag:dragging}" > < div class = "layout-container" > < div :class = "key" v-for = "(item, key) in mainData" :key = "key" > < draggable v-bind = "dragOptions" class = "list-group" :list = "item" @ end = "onEnd" @ start = "onStart" > < transition-group name = "list" > < div class = "list-group-item" v-for = "(element, index) in item" :key = "index" > < div class = "drag-handle" > {{ element.title }} </ div > < div class = "component-box" > < component :is = "element.name" /> </ div > </ div > </ transition-group > </ draggable > </ div > </ div > </ div > </ template > < script > import draggable from "vuedraggable" ; import { timeline, calendar, welcome, carousel, imgs, KonList } from "@/components/DragComponents" ; export default { components : { draggable, timeline, calendar, welcome, carousel, imgs, KonList }, data() { return { dragging : false , componentList : [ { name : "KonList" , title : "追番地址" , id : "5" }, { name : "imgs" , title : "五月最强新番" , id : "4" }, { name : "timeline" , title : "日程组件" , id : "2" }, { name : "carousel" , title : "走马灯组件" , id : "1" }, { name : "calendar" , title : "日历组件" , id : "3" } ], layout : { left : [ "5" , "4" ], right : [ "2" , "1" , "3" ] }, mainData : {} }; }, computed : { dragOptions() { return { animation : 30 , handle : ".drag-handle" , group : "description" , ghostClass : "ghost" , chosenClass : "sortable" , forceFallback : true }; } }, mounted() { this.getLayout(); }, methods : { onStart() { this.dragging = true ; }, onEnd() { this.dragging = false ; this.setLayout(); }, getLayout() { let myLayout = JSON.parse( window.localStorage.getItem( "kon" )); if (!myLayout || Object.keys(myLayout).length === 0 ) myLayout = this.layout; const newLayout = {}; for ( const side in myLayout) { newLayout[side] = myLayout[side].map( i => { return this.componentList.find( c => c.id === i); }); } this.mainData = newLayout; }, setLayout() { const res = {}; for ( const side in this.mainData) { const item = this.mainData[side].map( i => i.id); res[side]=item; } window.localStorage.setItem( "kon" , JSON.stringify(res)); } } }; </ script > < style lang = "scss" scoped >.layout-container { height: 100%; display: flex; .left { flex: 1; margin-right: 40px; } .right { width: 550px; } .list-group-item { margin-bottom: 20px; border-radius: 6px; overflow: hidden; background: #fff; } .component-box { padding: 20px; } .drag-handle { cursor: move; height: 40px; line-height: 40px; color: #fff; font-weight: 700; font-size: 16px; padding: 0 20px; background: #6cf; } } .drag { .component-box { display: none; } } .list-enter-active { transition: all .3s linear; } .list-enter, .list-leave-to { opacity: .5; } .sortable { .component-box { display: none; height: 0; } } .list-group { > span { display: block; min-height: 20px; } } .ghost { .drag-handle { background: rgb(129, 168, 187); } } </ style > 复制代码`