# vue keep-alive组件的使用 #

### **一、问题触发并解决** ###

最近自己在写vue练习，内容相对简单，主要是对vue进行熟悉和相关问题发现，查漏补缺。简单说下练习的项目内容及问题的产生：

练习使用的vue-cli 2.0脚手架搭建，内容就是简单的仿音乐播放app，功能也相对简单，大体就是单页router切换各个专辑列表，点击进入专辑内容页面，单击歌曲名称可以进行播放、暂停、下一曲功能。

简单的背景介绍完了，说下问题产生的情形：在从整个歌曲列表页点击跳转到单个专辑列表页，然后点击返回按钮返回歌曲列表页时，页面保存了之前的浏览位置，但是接口重新请求了数据，因为歌曲列表页有滚动加载效果，所以数据获取在vuex里用了数组的concat方法，导致返回请求的数据重新加载了列表里，而v-for循环由于key值有了重复，导致控制台报错；说起来可能比较难懂，上一些基本的代码部分：

` vuex里获取列表数据： GET_COLLECTION_LIST(state, val) { state.collectionList = state.collectionList.concat(val)} 歌曲列表里created获取数据，mounted监听滚动事件滚动加载，destroyed取消监听： created (){ this.collectionListMethod({ limit : this.limit, offset: this.offset})}, mounted (){ window.addEventListener( 'scroll' , this.isScrollBot)} destroyed (){ window.removeEventListener( 'scroll' , this.isScrollBot)} 专辑列表页返回使用的是 this. $router.go(-1) 复制代码`

有问题就要解决问题，通过查询了解到keep-alive可以对数据进行缓存，路由切换的时候可以使用缓存不用重新触发接口请求，貌似可以很完美解决问题，实践出真知，在代码里添加：

` 首先在歌曲列表路由里添加meta:{keepAlive: true }为后面的router-view是否需要缓存提供标识 { path: '/songLis' , component: SongLis, meta: { keepAlive: true }} 然后在router-view外层添加keep-alive，并根据meta参数判断是否所有路由都需要缓存 <keep-alive> <router-view v-if= " $route.meta.keepAlive" /> </keep-alive> <router-view v-if= "! $route.meta.keepAlive" /> 复制代码`

添加完毕，回到页面看效果！可喜可贺的是控制台不报错了，说明keep-alive起作用了，撒花庆祝~~~ 但是事情并没有那么简单，当我们从专辑列表切到其他路由，滚动鼠标会发现滚动加载的监听事件没有取消，组件destroyed的时候明明取消监听了啊！而且再切回到专辑列表页的时候，滚动加载反而不执行了，表示一脸懵~

经过查看 [vue文档]( https://link.juejin.im?target=https%3A%2F%2Fcn.vuejs.org%2Fv2%2Fapi%2F%23keep-alive ) ，发现文档有这么一句话：

` **<keep-alive>**` **包裹动态组件时，会缓存不活动的组件实例，而不是销毁它们。**

也就是说使用keep-alive的组件不会触发destroyed钩子方法，这就是取消监听失败的原因。根据文档介绍，keep-alive切换时，会触发自己的activeted和deactiveted两个钩子函数，可以理解为vue的created和destroyed两个钩子，因此需要修改代码，在deactivated时候取消监听，同时在activated的时候恢复监听，否则切回去的时候滚动监听也不会有效果：

` //keep-alive钩子函数，组件恢复时触发 activated (){ window.addEventListener( 'scroll' , this.isScrollBot)}, //keep-alive钩子函数，组件变为不可用时触发 deactivated (){ window.removeEventListener( 'scroll' , this.isScrollBot)} 复制代码`

修改后的效果完全符合预期，切换路由页面保留当位置，不会重复请求接口而且也不会在专辑列表组件外部触发滚动加载。

### **二、关于keep-alive** ###

既然用到了keep-alive，就通过文档简单总结下相关使用：

上面已经说过，通过keep-alive包裹的组件，在不活动时会缓存组件实例，不会对组件进行销毁，再次处于活动状态时，会读取缓存的内容并保存组件状态，不用重复请求接口获取数据。

**（1）最基本使用**

` 基本用法： <keep-alive> <component :is= "view" ></component> </keep-alive> 也可以根据条件判断： <keep-alive> <comp -a v-if= "a > 1" ></comp -a > <comp-b v-else></comp-b> </keep-alive> 这个时候触发两个组件切换时，组件内的数据和状态都会得到保存，如果有input输入框，输入框内容会保留 复制代码`

**（2）有条件的缓存**

keep-alive提供了include、exclude、max三个参数，前两个允许组件有条件的进行缓存，两者都可以接受字符串、正则、数组形式；max表示最多可以缓存多少个组件，接受一个number类型。

` <keep-alive include= "a,b" > <component :is= "view" ></component> </keep-alive> 此时只有a、b两个组件会触发keep-alive，此处的名字是组件内部name对应名字，如果name不存在，会查找父组件里components里注册的名字，如果也不存在则不会匹配。 正则和数组方法同上，但是需要用v-bind:include= '' 方式。 <keep-alive :max= "10" > <component :is= "view" ></component> </keep-alive> 复制代码`

**（3）钩子函数**

两个钩子函数activated和deactivated，上面已经说过，分别在组件恢复活动状态和组件失去活动状态时触发，可以起到类似created和destroyed钩子作用。

以上是对keep-alive的个人使用和理解，如有不足还望指正~