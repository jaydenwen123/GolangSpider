# 青桔单车 chameleon 跨平台实践分享 #

今 天Chameleon社区公众号收到了来自不懂小彬@青桔单车的投稿，让我们一起看一下青桔单车关于使用Chameleon的经验分享。

**
**

### ▍目录 ###

* 

前言

* 

背景

* 

行业现状——百家争鸣

* 

业务要求——高效稳定多入口

* 

框架期望

* 

实践

* 

跨端技术方案设计

* 

跨平台框架——chameleon

* 

青桔单车业务技术架构

* 

多端界面一致性

* 

差异化（定制化）

* 

工程化

* 

数据 mock

* 

CML 配置

* 

框架设计

* 

组件调用-父子组件通信

* 

数据的管理-store

* 

cml 框架设计

* 

性能

* 

性能提升

* 

包大小

* 

总结

### 前言 ###

近些年，整个前端领域发展迅速，效率型的前端框架也层出不穷，每个团队选择的技术解决方案都不太一致，因为互联网的特性及中国自身的特色，各个产品对于多端的投放的需求是一致的。像小程序这种跨端场景和现有的开发方式也不一样，为了满足业务的需求技术人员在日常开发中会因为投放平台的不一样而进行多套代码的维护，效率较低且成本也高。跨端框架应景而生（这里就不过多介绍各个跨端的框架了），在我看来 **大前端是趋势，跨端是趋势的第一步** ，而对于技术人员来说在不影响功能体验的情况下能解决维护多套代码的痛点非常重要。

### 背景 ###

#### 行业现状——百家争鸣 ####

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2b393543a145e?imageView2/0/w/1280/h/960/ignore-error/1)

经过将近两年的发展，小程序已经深入用户的日常生活，小程序应用数量超过了百万量级，覆盖200多个细分的行业，日活用户达到两个亿。与此同时，像支付宝、百度、头条、手Q等等都开始了自家的小程序生态，百家争鸣应景而生。青桔单车作为便民的出行工具，对于用户使用方式上也是成本越低体验越友好，即用即走的小程序已然成为平台选择的趋势。

#### 业务要求——高效稳定多入口 ####

高效、稳定、多入口就是业务现在的要求，青桔单车是日活相对较高的小程序（目前在阿拉丁小程序TOP榜前10），这也要求我们对小程序的性能（加载、渲染、响应的时间）、稳定及安全有较高的标准。

同时业务也需要在各个平台上获得更多的入口，这就直接导致我们在选择框架，业务开发时要求比较严谨。

**![](https://user-gold-cdn.xitu.io/2019/6/6/16b2b393542ead27?imageView2/0/w/1280/h/960/ignore-error/1)**

#### 框架期望 ####

从用户角度出发，为了减少用户使用成本（下载安装或更新APP），我们选择了市场上比较符合单车特性的平台作为入口。那么这时候对于研发来说就会有很多问题，我们在选择框架时，会对以下点有较高的期望或要求:

* **各端开发效果要高度一致，不能开发一次后用大量时间兼容多端**
* **针对不同的端有差异化的实现方式，能较容易去扩展，以应对产品对不同端的差异化需求**
* **工程化要好，因为前端过程式的开发比较低效** **，需要使用软件工程的技术和方法进行开发、管理**
* **框架设计，接入不用改动太多代码，具有较好的抽象度**
* **对性能上要求很高，包大小、加载、渲染、响应都需要严格考虑**

### 实践 ###

#### 跨端技术方案设计 ####

##### 跨平台框架——chameleon #####

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2b39356145022?imageView2/0/w/1280/h/960/ignore-error/1)

引用chameleon官网的简介： **Chameleon 不仅仅是跨端解决方案，让开发者高效、低成本开发多端原生应用。基于优秀的前端打包工具Webpack，吸收了业内多年来积累的最有用的工程化设计，提供了前端基础开发脚手架命令工具，帮助端开发者从开发、联调、测试、上线等全流程高效的完成业务开发。** 从chameleon框架的架构设计图看：

* 在各个平台下chameleon基于其本身的运行池增加一层包括路由、自定义生命周期、组件、数据事件绑定及管理、样式渲染等相对完整的DSL；
* chameleon-tools做编译（基于webpack），提供语法检查、转译、依赖组装等；
* chameleon-UI、chameleon-API做各个平台样式的不一致、API差异化的抹平；
* 提供多态协议的方式解决组件及API的差异化；

综合来看chameleon的设计模式比较适合我们做跨端项目的开发，提升我们的效率，不维护多套重复的代码……

##### 青桔单车业务技术架构 #####

青桔单车简单的业务流程图：

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2b39355d26fa6?imageView2/0/w/1280/h/960/ignore-error/1)

青桔单车业务相对复杂，包括登录、认证、电子围栏、宣教、状态扭转等超过 30 个页面（不包括各种H5实现），用户主动打开/微信扫码进入小程序，完成登录后开始扫码开锁，开锁成功后 === 发单成功，用户开始骑行，骑行结束后完成支付整个流程结束（这里只提到核心流程），因为业务需要，我们需要维护微信、支付宝、高德（快应用、百度接入中）等众多小程序，对于研发来说最快的是copy一套代码然后针对性进行修改，当进行新功能开发的时候就痛苦了。

基于业务和多端的差异抹平方案考虑，最终CML青桔单车小程序技术方案设计图如下

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2b39357c582d0?imageView2/0/w/1280/h/960/ignore-error/1)

从青桔单车现在模块看，为了能真正的实现跨端开发，需要解决各平台间的差异问题：

* 各平台提供给小程序的SDK差异化，例如微信、支付宝的getAuthCode，API方法名不一样，返回也不一样，类似的问题如何解决？
* 业务中台API根据渠道/接入方案不同造成的问题如何解决？
* 页面内部-组件调用-父子组件通信、store数据的管理如何处理？
* 硬件底层实现方案如何解决？例如微信、支付宝BLE底层实现的差异；百度小程序IOS不支持BLE，类似问题如何解决？
* 各平台规则限制导致功能无法使用，例如webview某些端考虑稳定性限制了h5 — data上行到端的问题如何解决？
* 其他问题，utils组件差异化……

#### 多端界面一致性 ####

由于各端组件化的实现方式不一样(微信webcomponents、百度模拟的组件化、支付宝是react框架)，多端界面的一致性是一个比较麻烦的事情，从cml本身的设计及实际的体验来说，不论样式的单位换算还是组件的统一封装都做了较好的统一。

来点gif图，预览一下青桔单车小程序基于cml改版后三端的效果吧：

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2b3936247d3a6?imageslim) ![](https://user-gold-cdn.xitu.io/2019/6/6/16b2b3938a71115d?imageslim) ![](https://user-gold-cdn.xitu.io/2019/6/6/16b2b3938bd64fa9?imageslim) ![](https://user-gold-cdn.xitu.io/2019/6/6/16b2b39394a6468e?imageslim) ![](https://user-gold-cdn.xitu.io/2019/6/6/16b2b39395e3c154?imageslim)

#### 差异化（定制化） ####

我们按照业务场景拆分了几个公用的模块包括用户相关、发单、硬件/蓝牙通信、订单管理、营销等，每个模块都单独store、action及暴露的commonApi，配合各个页面逻辑实现整个产品功能，针对差异化我们列举一个登录的例子通过多态方式来兼容微信、支付宝登录接口。我们在项目中src/componets中建立一个API的空间作多态管理API，针对login我们按照cml的规范建立一个login.interface文件

实现如下：

` <script cml-type= "interface" > // 定义入参 type RESObject = { code: string, errMsg: string } type done = (res: RESObject) => Void interface MethodsInterface { login(timeout: String, opt: done ): Void } </script> <script cml-type= "wx" > class Login implements LoginInterface { login (timeout, opt) { wx.login(timeout, code => done ({code: '' 12 ', errMsg: ' 成功 '}))； } } export default new Login(); </script> <script cml-type="alipay"> class Login implements LoginInterface { login (timeout, opt) { my.getAuthCode(timeout, code => done({code: ' '12' , errMsg: '成功' }))； } } export default new Login(); </script> 复制代码`

这里想着重提一个非常容易被忽略的问题：interface中一定要定义输入输出，规范各端实现规范，这是为了避免在修改某个端的方法入参（增加或减少）却没有考虑其他端的实现，如果全量测试会浪费很大人力且也不能保证都能覆盖到测试，为了可维护性，建议从一开始就坚持写多态组件的interface，在程序层面上亡羊补牢有时候真的为时已晚……

接口定义后就可以用统一的方式进行调用

` import login from '../../components/Api/login/login.interface' ; export const login = function ({commit}) { return new Promise((resolve, reject) => { login.login() .then(({code}) => { commit(types.SET_USER_CODE, {code}); resolve({code}); }) .catch(e => { reject(); }); }); }; 复制代码`

再比如针对蓝牙通信的API微信需要端来做ArrayBuffer到HexString的转换而支付宝不需要，这里我们也采用多态进行接口方式抽离，微信进行转换，支付宝直接return原数据，整个BLE的过程非常复杂，为了提高连接，通信的成功率我们做了很多优化，如果直接在代码中hack会影响整个流程甚至造成整个蓝牙动作的不稳定，影响开锁率，做这层多态封装既不影响原有逻辑，改动也相对较少成本很低。

另一方面，PM如果在某一端提个需求，例如针对支付宝用户在登录时候加一些特殊功能，我们能在不影响其他流程比进行改造，同时这是在物理上进行了隔离，可以发布单独npm包，可维护性比较高。

#### 工程化 ####

工程化是使用软件工程的技术和方法对项目的开发、上线和维护进行管理，因为前端过程式的开发比较低效，可以通过模块化、组件化、本地开发、上线部署自动化来提高研发效率。

常用执行命令

* cml dev `编译全部`，cml wx dev `编译微信`，启动开发模式，监听文件变化动态打包
* cml build `编译全部`，cml wx dev `编译微信`, 构建生产环境

##### 数据mock #####

前端开发的过程中，数据mock是一个比较重要的功能，在验证逻辑、研发效率上及线上线下环境环境切换都起着很重要的作用。cml这里也提供数据mock的功能。

` import cml from "chameleon-api" ; cml.get({ url: '/api/getUserInfo' }) .then(res => { // ...省略部分实现逻辑 cml.setStorage( 'user' , ...res) }, err => { cml.showToast({ message: JSON.stringify(err), duration: 2000 }) }); 复制代码`

调用方法的参数url中只需要写api的路径。那么本地dev开发模式如何mock这个api请求以及build线上模式如何请求线上地址，就需要在配置文件中配置apiPrefix。

` // 设置api请求前缀 const test ApiPrefix = 'http://test.api.com' ; const apiPrefix = 'http://prod.api.com' ; cml.config.merge({ wx: { dev: { apiPrefix: test ApiPrefix }, build: { apiPrefix } } }) 复制代码`

cml支持本地mock，使用方式是在 ` /mock/api/` 文件夹下创建mock数据的js文件，启动dev模式后（apiPrefix留空表示本地ip+端口），可以直接实现本地mock。

##### CML配置 #####

1.chameleon的构建过程是配置化的，项目的根目录下提供一个chameleon.config.js文件，在该文件中可以使用全局对象cml的api去操作配置对象，针对不同端的配置方法，添加一个端还是比较简单的，在编译层配置的扩充性也有暴露

2.页面路由配置，chameleon项目内置了一套各端统一的路由管理方式，为区分小程序，用url表示h5，path表示小程序，同时提供mock服务接口，在路由的管理上相对容易，在小程序端构建时会将 ` router.config.json` 的内容，插入到app.json的pages字段，实现小程序端的路由。

#### 框架设计 ####

##### 组件调用-父子组件通信 #####

针对首页的消息流卡片我们会有登录、认证、订单、骑行卡等等不同的形式，我们封装了一个组件来处理，静态效果如下：

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2b3939f26dfe7?imageView2/0/w/1280/h/960/ignore-error/1)

相对来说这个组件比较通用，我们大致分为2种类型，通知型与动作型，如上图2就是一个纯显示的消息，其他的是带有按钮的消息，在组件设计上我们根据调用组件时传的type不一样来区分，组件使用上用props进行data传递。

通知型：TipsCard

动作型：ActionCard

下面是动作型的实现方式：

` <template> <view class= "action-card" > <view class= "{{'content '+(remindActive && 'bigSmall')}}" > // ... </view> </view> </template> <script> class ActionCard { mounted () { EventBus.on(REMIND_CARD, ({index}) => { if (index === this.index) this.remind(); }); } methods = { callback (e) { EventBus.emit(ACTION_CARD, {index: this.index}); } } } export default new ActionCard() </script> 复制代码`

调用上，我们动态传递 component is 中的type来指定不同的消息流类型

在自定义事件的处理上，我们通过c-bind:action绑定了一个componentAction，通过EventBus事件来传递执行

` <template> <view> <component is= "{{item.type+'-card'}}" /> // ... </view> </template> <script> class Home { beforeMount () { this.tipsList = [] EventBus.on(ACTION_CARD, (data) => { this.componentAction(data) }); } methods = { componentAction(data) { ...// 执行实例 } } } export default new Home(); </script> 复制代码`

最终实现的效果如下图

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2b393a41de14d?imageslim) ![](https://user-gold-cdn.xitu.io/2019/6/6/16b2b393d1f112e7?imageslim)

##### 数据的管理-store #####

cml框架作为数据驱动视图的框架，应用由数据驱动，数据逻辑变得很复杂，必需要一个好用高效的数据管理框架，我们可能在各个页面都需要mapState，mapAction来做组件的状态获取，事件分发等

` import createStore from 'chameleon-store' ; import user from './user' ; import location from './location' ; // …………这里省略其他状态管理器 const store = createStore({ modules: { user, location, bicycle, // …………这里省略其他状态管理器 } }); export default store; 复制代码`

这里就列出对user状态的入口操作

` import mutations from './mutations' ; import * as actions from './actions' ; import * as getters from './getters' ; const user = { state: { }, mutations: { ...mutations }, actions: { ...actions }, getters: { ...getters } }; export default user; 复制代码`

cml提供了的chameleon-store模块包，且用法、写法和vuex一样，这点非常感人，手动点赞

##### cml框架设计 #####

cml提供了计算属性——computed、侦听属性——watch、类vuex的数据管理及DSL语法，从学习成本上来看较小，符合当前的开发习惯。

` class Home { data = { remindActive: false } watch = { lockState: function (cur, old) { return this.remindActive } } computed = { ...store.mapState([ 'location' , 'user' ]) } } export default new Home(); </script> 复制代码`

#### 性能 ####

基于青桔单车日活相对较高的特性，高效、稳定、多入口就是业务现在的要求，这要求我们对小程序的性能（加载、渲染、响应的时间）、稳定及安全有较高的标准。

##### 性能提升 #####

性能上，cml做了array diff，因为小程序的主要运行性能瓶颈是webview和js虚拟机的传输性能，cml尽可能diff出修改的部分进行传输，性能相对会更好，贴一下源码

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2b393d5cf6feb?imageView2/0/w/1280/h/960/ignore-error/1)

源码实现比较简单，但带来实际的用处还是比较大的，单车小程序有不少列表页面，当对于列表data进行重新赋值先进入diff函数过滤出实际改变的每一项，再进行view-render，性能上会提高不少

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2b393db4811ae?imageView2/0/w/1280/h/960/ignore-error/1)

##### 包大小 #####

包大小上，cml相对之前开发的版本并没有大的变化，1.2mb的编译后包大小基本保持不变，这归功于代码的压缩及其本身的runtime代码并不大，给个好评。

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2b393def354d5?imageView2/0/w/1280/h/960/ignore-error/1)

### 总结 ###

**“苦尽甘来”** 是青桔单车小程序接入chameleon跨端框架的总体感受。

苦：这里的”苦“我们到现在看其实是因为一种为未来做铺垫而需要做更多工作的“苦”，原本只需要考虑一端的CASE但现在要自己抽象出来一系列例如接口参数，通用组件，以及需要掌握更多的平台技术方案的“苦”。

甘：当我们通过chameleon将青桔单车小程序上线后，应对业务抛出来的新需求，我们再也不用维护多套代码了，也把因为维护多套代码导致可能某一端不一致性的风险给彻底排除了。不仅仅是RD同学受益，QA同学验收的时候因为是一套代码逻辑所以可以节省一半的时间。本来3天要完成的需求现在1.5天就可以完成也极大的满足了业务方的需求，我们认为这就是技术驱动业务发展的一个例子，技术不能只是一个工具而是要从业务角度出发去实现才有真正的价值。

cml做了DSL编译转化从根本上实现了各端的统一，引入组件多态、方法多态并抽象各端的配置可以更灵活抹平差异，方便配置、扩展。通过深入了解chameleon框架及项目实践，我们基于cml可以再抽象出一些公共层的组件、接口，当然这里是指的跨平台的组件、接口，因为只有这样才能更大化提升开发的效率，给业务带来更多的可能。

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2b3944a461d88?imageView2/0/w/1280/h/960/ignore-error/1)

我们在选择框架的时候要考虑不仅要考虑其本身的性能、稳定、安全，包大小等，还需要看一下复用性及发展，比如是否能基于框架本身扩展出适合业务甚至公司级别的通用组件、API，是否方便随着业务的变化而能做到及时响应式的扩展、变动。

到这里青桔单车chameleon跨端实践的介绍结束啦，有表述不好的希望多多包含并提出建议我们及时改进，谢谢。

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2b3130ae58465?imageView2/0/w/1280/h/960/format/png/ignore-error/1)