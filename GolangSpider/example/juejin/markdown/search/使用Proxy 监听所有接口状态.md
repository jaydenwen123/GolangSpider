# 使用Proxy 监听所有接口状态 #

### 代理为我解决了什么 ###

在开发项目过程中几乎所有接口都需要知道它的返回状态，比如失败或者成功，在移动端通常后台会返回结果，而我们只需要一个弹窗来弹出来结果就可以了。但是这个弹窗如果在整个项目里需要手动去每一个都定义，那是非常庞大的代码量，而且维护起来非常的麻烦。通常做法就是绑定在原型上一个公共方法，比如this.message('后台返回接口信息')。 这样看似省力了很多其实还是很麻烦。 如果使用了proxy做一个全局代理，那么就完全不一样了。不管任何一个api都会将状态传递个这个代理中心，并且由代理中心直接反应结果。

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2adfe439b1565?imageView2/0/w/1280/h/960/ignore-error/1)

` import Vue from 'vue' import {ToastPlugin} from 'vux' import api from './api/api' //引入封装好的api模块，和使用的toast弹出窗，弹窗可以选择任何框架的看起来比较好看的弹窗组件 Vue.use(ToastPlugin); //toast初始化 let vm = new Vue(); //创建实例，是因为toast弹窗依赖它所以这里要创建个实例，去调用弹窗用 Vue.prototype.dilog = function (value) { vm. $vux.toast.show({ text: value || "业务处理成功" , type : 'text' , width: "5rem" , position: 'middle' }); }; //陷阱，只要接口状态改变就会调用此方法 var interceptor = { set : function (recObj, key, value) { vm.dilog(value); //弹出层，value就是api返回的状态值 return this } }; //创建代理以进行侦听 var proxyEngineer = new Proxy(api, interceptor); Vue.prototype.api = proxyEngineer; //将api替换为新的实例 复制代码`

之所以这样做，是因为创建好的封装好的api文件里，不应该在去引入一个vue实例了，如果不用代理，直接在api文件里引入vue那将是巨大的消耗。

` class API { constructor (){ this.massages = "业务处理成功!" ; //定义信息状态属性 //当前接口错误提示 this.code= '000000' || '999999' } post(params, callback, dailog, errcallback = function () { //错误信息回调}) { //dailog 是是否需要在初始化弹窗，比如一个列表通常不需要加载完了弹出一个加载成功，或者获取数据成功什么的。Boolean，通常只需要在点击某事件时候使用，或者是初///始化数据报错使用 //this.code 代表状态码 let config={}; config.data = params.data||{}; var url = ` ${base} ${params.url}.do`; var dailog = dailog; //封装了axios的post方法 return axios.post(url, config.data, config, dailog).then(res => { let rst = res.data; if (rst.code === '000000' || rst.code === '999999' ) { callback&&callback(rst.result||{}); if (dailog) { //根据dailog 值来判断需不需要弹窗 this.massages = rst.message; } } else { errcallback && errcallback(); this.massages=rst.message; //监听massages的变化 } //这里如果返回this返回的是代理对象的this return res }).catch(e => { console.log(e) }) } } const api = new API(); export default api //代码核心地方其实就是在类上定义了信息字段，通过massages值变化来反馈信息 复制代码`

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2b3768bd154ac?imageView2/0/w/1280/h/960/ignore-error/1) 我所使用的toast效果。

` this.api.post(params, res => { //你需要执行的逻辑 // 再也不需要写什么 //this. $msg (res.value) 这类的代码，代理已经都帮你处理完了 }) 复制代码`

这就是我在实际中用到的代理，这个方法不管在多页面还是单页面都适用。当然代码有些粗糙，也没做过多限制，只是说了下思想。以防自己忘记。
顺带说下代理这个特性这里就把《Understanding ECMAScript 6》这本书的内容拿来用了，并稍微添加一些自己的理解。只做记录。

### 代理与反射是什么？ ###

通过调用 new Proxy() ，你可以创建一个代理用来替代另一个对象（被称为目标），这个代 理对目标对象进行了虚拟，因此该代理与该目标对象表面上可以被当作同一个对象来对待。 代理允许你拦截在目标对象上的底层操作，而这原本是 JS 引擎的内部能力。拦截行为使用了 一个能够响应特定操作的函数（被称为陷阱）。 拦截器的概念比较重要。 Reflect 拦截器有一些反射接口，

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2b4385cd22189?imageView2/0/w/1280/h/960/ignore-error/1) 拦截的作用其实就是重写内置对象的特定方法。

### 创建一个简单的代理 ###

` let target = {}; let proxy = new Proxy(target, {}); proxy.name = "proxy" ; console.log(proxy.name); // "proxy" console.log(target.name); // "proxy" target.name = "target" ; console.log(proxy.name); // "target" console.log(target.name); // "target" proxy拦截代理了target 复制代码`

### 我主要使用到的点 ###

` let target = { name: "target" }; let proxy = new Proxy(target, { set ( trap Target, key, value, receiver) { // 忽略已有属性，避免影响它们 if (! trap Target.hasOwnProperty(key)) { if (isNaN(value)) { throw new TypeError( "Property must be a number." ); } } // 添加属性 return Reflect.set( trap Target, key, value, receiver); } }); // 添加一个新属性 proxy.count = 1; console.log(proxy.count); // 1 console.log(target.count); // 1 // 你可以为 name 赋一个非数值类型的值，因为该属性已经存在 proxy.name = "proxy" ; console.log(proxy.name); // "proxy" console.log(target.name); // "proxy" // 抛出错误 proxy.anotherName = "proxy" ; 复制代码`

另外vue3.0 的响应式也是使用的代理