# React 基础知识总结 #

## 前言 ##

* emememe 今天在掘金贡献第一篇文章，自己习惯在有道云写笔记，也就直接复制过来了，所以排版大家就将就一下，如果大家看到有啥可以优化或者有更好的方法，麻烦评论分享下哦，感谢

## setState ##

* [imweb.io/topic/5b189…]( https://link.juejin.im?target=https%3A%2F%2Fimweb.io%2Ftopic%2F5b189d04d4c96b9b1b4c4ed6 ) 文档说明

#### 注意事项一： ####

* 

更新数据使用 setState方法， **==但是setState方法既是同步的方法也是异步方法：（再原生事件，和setTimeout中是同步事件，因为不需要走合成事件这一步，只有在合成事件和钩子函数中是“异步”的），setState后并不会马上跟新数据，即使在setState后打印this.state，依然不是最新的数据==** ；

* 

setState方法：有两个方法可以获取到更新后的数据：

* 

1，使用setState的第二个参数，回调函数，可以获取到更新后的state数

* 

2，【注意】：官方推荐使用这个方法： **==就是在生命周期 componentDidUpdate() 里面去获取，事实证明： componentDidUpdate()的执行顺序 比 setState的回调函数要快==**

* 

详细可以看： [juejin.im/post/5b45c5…]( https://juejin.im/post/5b45c57c51882519790c7441 )

* 

setState有两个参数

* 参数1：是要更新的数据，是一个对象
* 参数2：是一个回调函数，

` <!--默认数据--> this.state = { count : 1 } this.setState({ count : this.state.count + 1 }, () => { <!-- 回调函数可以获取到更新过的state --> <!--【注意】： 但是执行却是在componentDidMount()生命周期后面，所以官网也是建议在生命周期componentDidMount()里面获取更新 state的数据--> console.log(this.state) }) <!--生命周期--> componentDidUpdate() { console.log('componentDidUpdate') console.log(this.state) } 复制代码`

#### 注意事项二： ####

* 设想有一个需求，需要在在onClick里累加两次，如下：

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2b971216d1e77?imageView2/0/w/1280/h/960/ignore-error/1)

* setStater 内部执行是 类似 Object.assign 这样的浅拷贝，后面回把前面的值覆盖掉

` let state = Object.assign( { count : 2 }, { count : 2 } ); state === { count : 2 }; 复制代码`

* 所以这种情况需要使用函数提交数据

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2b97b4bd6206a?imageView2/0/w/1280/h/960/ignore-error/1)

#### 注意事项三： ####

* 在 componentDidMount()生命周期， setState更新数据，会触发二次的render ；渲染会发生在浏览器更新屏幕之前，请谨慎使用该模式，因为它会导致性能问题

## 生命周期 ##

#### 生命周期-组件更新后 componentDidUpdate() ####

* 

componentDidUpdate 会在更新后会被立即调用。首次渲染不会执行此方法

* 

当组件更新后，可以在此处对 DOM 进行操作。如果你对更新前后的 props 进行了比较， 动画效果。（例如，当 props 未发生变化时，则不会执行网络请求）

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2b984ba646bee?imageView2/0/w/1280/h/960/ignore-error/1)

* 

componentDidUpdate 三个参数

* 参数1：prevProps 表示更新前的 props 的值
* 参数2：prevState 表示更新前的 state 的值
* 参数3：snapshot 如果组件实现了 getSnapshotBeforeUpdate() 生命周期（不常用），则它的返回值将作为参数传递给 “snapshot”。否则此参数将为 undefined

* 

**==【注意】：componentDidUpdate 里面可以调用setState跟新数据，但是必须是在条件判断的情况下，否则会造成死循环==**

* 

**==【注意】：如果 shouldComponentUpdate() 返回值为 false，则不会调用 componentDidUpdate()==**

#### 生命周期-组件更新前 shouldComponentUpdate() 不常用 ####

* 

组件跟新前调用，必须返回true || false， 否则会报错， 如果返回true，则会重新渲染页面，返回false， 就会阻止后面的渲染，不会跟新页面，可以适当做一些优化，但是不会阻止子组件渲染，除非子组件也返回false

* 

如果要做优化，那就把this.props 和 nextProps || this.state 和 nextState 做比较

* 

不建议在 shouldComponentUpdate() 中进行深层比较或使用 JSON.stringify()。这样非常影响效率，且会损害性能。

* 

shouldComponentUpdate 两个参数

* 参数1：nextProps 表示更新后的 props 的值
* 参数2：nextState 表示更新后的 state 的值

* 

react内部有内置的优化方式 PureComponent 其实跟 Component一样，只不过PureComponent里面对shouldComponentUpdate做了一层浅优化，但是一般只有展示组件才会用到这个

#### 生命周期-组件销毁 componentWillUnmount() ####

* 

会在调用 render 方法之前调用，并且在初始挂载及后续更新时都会被调用。它应返回一个对象来更新 state，如果返回 null 则不更新任何内容

* 

会在组件卸载及销毁之前直接调用。在此方法中执行必要的清理操作，例如，清除定时器， 事件等。

* 

componentWillUnmount() 中不应调用 setState()，因为该组件将永远不会重新渲染。组件实例卸载后，将永远不会再挂载它。

#### 生命周期-组件更新前 static getDerivedStateFromProps() ####

* 

static getDerivedStateFromProps(props, state) 两个参数

* 参数1：props 表示更新后的 props 的值
* 参数2：state 表示子组件的 state 的值

* 

这个生命周期包含了旧的生命周期componentWillReceiveProps， 所以两个不能同时存在页面

* 

这个生命周期的作用在于把 父组件props的值赋给子组件的state， 但是这个并不是双向数据流，所以父组件跟新了，子组件并不会跟新，这个方法就像vue 的watch，监听props的变化，然后赋值给子组件的state

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2b9916024fdfa?imageView2/0/w/1280/h/960/ignore-error/1)

* 一般受控组件不需要用到这个生命周期

#### 生命周期-捕获所有组件报错 ####

* 

如果一个 class 组件中定义了 static getDerivedStateFromError() 或 componentDidCatch() 这两个生命周期方法中的任意一个（或两个）时，

* 

那么它就变成一个错误边界。== **当抛出错误后，请使用 static getDerivedStateFromError() 渲染备用 UI ， 使用 componentDidCatch() 打印错误信息** ==

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2b99738c568dc?imageView2/0/w/1280/h/960/ignore-error/1)

* 调用

## ![](https://user-gold-cdn.xitu.io/2019/6/6/16b2b99a8ad3cd45?imageView2/0/w/1280/h/960/ignore-error/1) ##

#### defaultProps 默认的 props 的值 ####

* 默认的props的值，适用于一些props 需要默认值，但是又没有传递进来的值

` <!--组件 class 名--> ChildIndex. defaultProps = { total : 39 } 复制代码`

## props - 父子组件之间的传值方式(只读) ##

* 父组件通过传递函数给子组件，子组件通过调用这个函数，把值传给父组件

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2b99f0b1732ba?imageView2/0/w/1280/h/960/ignore-error/1)

#### 方式一： ####

* 子组件直接调用父组件的函数

* 情况一：如果不传参数，直接this.props.clickChange， 就可以
* 情况二：需要传递参数，必须使用箭头函数包裹起来，不然会报错；因为子组件是通过原生事件调用父组件的方法，那我们传递参数的时候，已经是返回值， 这个是不是一个函数，原生事件接受的是一个回调函数哦

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2b9a2fd64c9a2?imageView2/0/w/1280/h/960/ignore-error/1)

#### 方式二： ####

* 子组件添加一个方法，当事件触发的时候调用这个方法，再通过这个方法调用父组件的方法

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2b9a5f9e0d441?imageView2/0/w/1280/h/960/ignore-error/1)

## context 里面的数据能被随意接触就能被随意修改 ##

* 相当于作用域的顶层变量 只要是这个组件内所有的子组件，都可以获取到这个context的值

` // 父组件必须先定义context的类型 static childContextTypes = { themeColor : PropTypes.string } // getChildContext 这个方法就是设置 context 的过程，它返回的对象就是 context（也就是上图中处于中间的方块），所有的子组件都可以访问到这个对象。 // 我们用 this.state.themeColor 来设置了 context 里面的 themeColor getChildContext () { return { themeColor : this.state.themeColor } } // 子组件也必须定义这个 static contextTypes = { store : PropTypes.object } // 获取context的值，必须带上context this.context.themeColor 复制代码`

## 条件渲染 ##

* 条件渲染，首字母必须大写，不然react不会渲染，会报错，函数的首字母也要大写

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2b9aa69a62afa?imageView2/0/w/1280/h/960/ignore-error/1)

#### react 事件调用的三种方式 ####

* 之前群里有小伙说，在腾讯面试的时候，问到方法一和方法二的区别，哪个好？腾讯告知说方法二好，方法一好像有丢失this的bug；（具体百度）

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2b9adfa3c817e?imageView2/0/w/1280/h/960/ignore-error/1)

## 注意事项 ##

* textarea、select、input通常都是受控组件，而且值都是要通过change事件改变，所以结合起来，只能通过value 和 onChange 来做控
* **textarea 如果字段只是只读，必须设置readOnly ，如果要改变，必须设置onChange事件**

* value值不能直接写入，只能通过value设置

` < textarea value = {this.state.value} readOnly /> 复制代码`

* **select 设置值只能通过value 和 onChange事件，option中的selected 是默认选中，但是并不会使用 selected 属性**

` < select value = {this.state.value} onChange = {this.handleChange} > < option value = "grapefruit" > 葡萄柚 </ option > < option value = "lime" > 柠檬 </ option > </ select > 复制代码`

* **input 设置值只能通过value 和 onChange事件**

` < input type = "text" value = {this.state.value} onChange = {this.handleChange} /> 复制代码`

* **name属性得用法，实现多个input传值案例**
* 如果做多个input呢，把字段当作name属性传入，通过event.target.name 获取name属性, 把name 当作是state 里面得key，存数据

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2b9b42d7d6585?imageView2/0/w/1280/h/960/ignore-error/1)

* 

**label 的 for属性 应该写成 htmlFor，不然会报错哦**

* 

**<React.Fragment></React.Fragment> 像vue的模板一样，不会渲染出来，适合某些特殊的场景，也可以使用新语法<></>(很多不支持新语法，不支持key)**

* 

**组件的名字必须大写，首字符也要大写，不然react会误认为是html标签**

* 

**className每个元素只能使用一次，如果想要多个样式切换，可以使用模板字符串 ``**

` < div className = { ` home-foot-dialog ${ this.state.isShowDialog ? ' home-foot--show ' : ' home-foot--hide '}`} > 复制代码`

* style 最好是先定义，在赋值给dom, 而且style 使用 ''字符串，是默认不会渲染的，否则每次componentDidMount这个生命周期内都是拿到已经渲染好的style

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2b9b7b761716d?imageView2/0/w/1280/h/960/ignore-error/1)

* children 类似 vue的slot插槽，可以用来放入不同的内容

## redux and react-redux（简约版） ##

#### 不使用装饰器的用法： ####

* 1，首先创建store.js 文件，里面存放各种返回的类型数据

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2b9bf0f26c603?imageView2/0/w/1280/h/960/ignore-error/1)

* 2，在index.js 引入redux react-redux，并且创建store

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2b9c381b7d094?imageView2/0/w/1280/h/960/ignore-error/1)

* 3，在要存储数据的页面引用react-redux 中的connect 方法，然后dispatch数据

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2b9c6f7f51a70?imageView2/0/w/1280/h/960/ignore-error/1)

* 4，在要到的数据的页面引用react-redux 中的connect 方法

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2b9ca50bb0e39?imageView2/0/w/1280/h/960/ignore-error/1)

#### react-redux 中的 connect 方法两个参数 ####

* 

参数1： mapStateToProps 获取值，把值放入props中

* 

参数2： mapDispatchToProps 设置值，通过dispatch把值存入redux中

* 

== **【注意】：使用react-redux 必须使用 connect 方包裹组件** ==

* 

【注意】：刷新后，redux数据会没有

## redux and react-redux（完整版） ##

* 创建 redux目录

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2b9cdc2239ac7?imageView2/0/w/1280/h/960/ignore-error/1)

* 创建 actionTypes.js 文件， 里面是各种type类型

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2b9d004fcaf2f?imageView2/0/w/1280/h/960/ignore-error/1)

* 创建 actions.js 文件，里面都是 dispatch 的方法，对应的每一个commit，而且类型已经定义好，就不用再去定义类型，不过参数名字要统一一个

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2b9d31b25d966?imageView2/0/w/1280/h/960/ignore-error/1)

* 创建 reducers.js 文件，返回对应的数据；
* combineReducers 可以合并数据，那就可以定义不同的module，引入就可以用

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2b9db05fde501?imageView2/0/w/1280/h/960/ignore-error/1)

* 创建 store.js ，创建Store，最后再index.js引入就好

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2b9dde6701415?imageView2/0/w/1280/h/960/ignore-error/1)

#### 使用装饰器 ####

* 先在package.json 里面，找到babel的配置，可以添加这个插件 @babel/plugin-proposal-decorators

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2b9e1944f8697?imageView2/0/w/1280/h/960/ignore-error/1)

#### 用法： ####

* dispatch state数据， 把数据绑在props上面

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2b9e6e8e51802?imageView2/0/w/1280/h/960/ignore-error/1)

* 获取state的数据，把数据绑在props上面

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2b9e9afeb80c9?imageView2/0/w/1280/h/960/ignore-error/1)

` <!--引入对应的函数--> import { connect } from 'react-redux' import { deviceQuery } from '../../redux/actions' import { withRouter } from 'react-router-dom' /** * 装饰器 (切记，不能添加分号，不然会报错) * @param {mapStateToProps} state 获取state的数据，然后把数据放在props上面 * @param {mapDispatchToProps} dispatch state数据， 把数据绑在props上面 */ @connect( null , dispatch => ({ onDeviceQuery(item) { dispatch( deviceQuery({ paramsOne : item.paramsOne, paramsTwo : item.paramsTwo, line : item.line })) } })) @withRouter <!--上面栗子 === 下面的栗子--> @connect( state => ({ device :state.device}), dispatch=>({ deviceQuery(item){ dispatch( deviceQuery({ paramsOne : item.paramsOne, paramsTwo : item.paramsTwo, line : item.line }) )} }) )(withRouter(HomeMap) <!--获取值--> @connect( state => ({ device : state.device}), null ) @withRouter <!--上面栗子 === 下面的栗子--> @connect( state => ({ device :state.device}), null , )(withRouter(HomeMap) 复制代码`

## react-dom-route ##

* react-dom-route4的路由是按照模块管理，看路由看成每个模块

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2b9ee175d2f8a?imageView2/0/w/1280/h/960/ignore-error/1)

* 

路由跳转，需要引入Link 和 withRouter

* 

使用withRouter方法包裹组件，这样才能获取 history属性

* 

Link 会渲染成a标签，可以携带参数跳转

` import { Link, withRouter } from 'react-router-dom' <!--路由跳转 - 带参数路由 --> < Link to = "/analysis" className = "tool-data--href" > </ Link > < Link to = { `/ board /${ params `}> 每天上网用户统计 </ Link > <!--使用withRouter方法包裹组件，这样才能获取 history属性--> withRouter(HomeMap) this.props.history.push("/device"); 复制代码`

* 

exact 严格匹配路由

* 

Switch：使用Switch组件来包裹一组。 会遍历自身的子元素（即路由）并对第一个匹配当前路径的元素进行渲染

* 

react-dom-route有两种模式：

* BrowserRouter： 用来管理动态请求
* HashRouter： 用于静态页面，是最好的选择