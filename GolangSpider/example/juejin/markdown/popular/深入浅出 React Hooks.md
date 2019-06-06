# 深入浅出 React Hooks #

# ![image.png](https://user-gold-cdn.xitu.io/2019/6/3/16b1aec6de4f8227?imageView2/0/w/1280/h/960/ignore-error/1) #

> 
> 
> 
> 直播回放链接： [云栖社区](
> https://link.juejin.im?target=https%3A%2F%2Fyq.aliyun.com%2Farticles%2F700174
> ) ( [@x-cold](
> https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fx-cold ) )
> 
> 

## React Hooks 是什么？ ##

Hooks 顾名思义，字面意义上来说就是 React 钩子的概念。通过一个 case 我们对 React Hooks 先有一个第一印象。

假设现在要实现一个计数器的组件。如果使用组件化的方式，我们需要做的事情相对更多一些，比如说声明 state，编写计数器的方法等，而且需要理解的概念可能更多一些，比如 Javascript 的类的概念，this 上下文的指向等。

[示例]( https://link.juejin.im?target=https%3A%2F%2Fcodepen.io%2Fx-cold%2Fpen%2FJqjZKR )

` import React, { Component } from 'react' ; import ReactDOM from 'react-dom' ; class Counter extends React. Component { state = { count : 0 } countUp = () => { const { count } = this.state; this.setState({ count : count + 1 }); } countDown = () => { const { count } = this.state; this.setState({ count : count - 1 }); } render() { const { count } = this.state; return ( < div > < button onClick = {this.countUp} > + </ button > < h1 > {count} </ h1 > < button onClick = {this.countDown} > - </ button > </ div > ) } } ReactDOM.render( < Counter /> , document.getElementById('root')); 复制代码`

使用 React Hooks，我们可以这么写。

[示例]( https://link.juejin.im?target=https%3A%2F%2Fcodepen.io%2Fx-cold%2Fpen%2FZNEReY )

` import React, { useState } from 'react' ; import ReactDOM from 'react-dom' ; function Counter ( ) { const [count, setCount] = useState( 0 ); return ( < div > < button onClick = {() => setCount(count + 1)}>+ </ button > < h1 > {count} </ h1 > < button onClick = {() => setCount(count - 1)}>- </ button > </ div > ) } ReactDOM.render( < Counter /> , document.getElementById('root')); 复制代码`

通过上面的例子，显而易见的是 React Hooks 提供了一种简洁的、函数式（FP）的程序风格，通过纯函数组件和可控的数据流来实现状态到 UI 的交互（MVVM）。

### Hooks API ###

* [Basic Hooks]( https://link.juejin.im?target=https%3A%2F%2Freactjs.org%2Fdocs%2Fhooks-reference.html%23basic-hooks )

* ` useState` ( https://link.juejin.im?target=https%3A%2F%2Freactjs.org%2Fdocs%2Fhooks-reference.html%23usestate )
* ` useEffect` ( https://link.juejin.im?target=https%3A%2F%2Freactjs.org%2Fdocs%2Fhooks-reference.html%23useeffect )
* ` useContext` ( https://link.juejin.im?target=https%3A%2F%2Freactjs.org%2Fdocs%2Fhooks-reference.html%23usecontext )

* [Additional Hooks]( https://link.juejin.im?target=https%3A%2F%2Freactjs.org%2Fdocs%2Fhooks-reference.html%23additional-hooks )

* ` useReducer` ( https://link.juejin.im?target=https%3A%2F%2Freactjs.org%2Fdocs%2Fhooks-reference.html%23usereducer )
* ` useCallback` ( https://link.juejin.im?target=https%3A%2F%2Freactjs.org%2Fdocs%2Fhooks-reference.html%23usecallback )
* ` useMemo` ( https://link.juejin.im?target=https%3A%2F%2Freactjs.org%2Fdocs%2Fhooks-reference.html%23usememo )
* ` useRef` ( https://link.juejin.im?target=https%3A%2F%2Freactjs.org%2Fdocs%2Fhooks-reference.html%23useref )
* ` useImperativeHandle` ( https://link.juejin.im?target=https%3A%2F%2Freactjs.org%2Fdocs%2Fhooks-reference.html%23useimperativehandle )
* ` useLayoutEffect` ( https://link.juejin.im?target=https%3A%2F%2Freactjs.org%2Fdocs%2Fhooks-reference.html%23uselayouteffect )
* ` useDebugValue` ( https://link.juejin.im?target=https%3A%2F%2Freactjs.org%2Fdocs%2Fhooks-reference.html%23usedebugvalue )

### useState ###

useState 是最基本的 API，它传入一个初始值，每次函数执行都能拿到新值。

` import React, { useState } from 'react' ; import ReactDOM from 'react-dom' ; function Counter ( ) { const [count, setCount] = useState( 0 ); return ( < div > < button onClick = {() => setCount(count + 1)}>+ </ button > < h1 > {count} </ h1 > < button onClick = {() => setCount(count - 1)}>- </ button > </ div > ) } ReactDOM.render( < Counter /> , document.getElementById('root')); 复制代码`

需要注意的是，通过 useState 得到的状态 count，在 Counter 组件中的表现为一个常量，每一次通过 setCount 进行修改后，又重新通过 useState 获取到一个新的常量。

### useReducer ###

useReducer 和 useState 几乎是一样的，需要外置外置 reducer (全局)，通过这种方式可以对多个状态同时进行控制。仔细端详起来，其实跟 redux 中的数据流的概念非常接近。

` import { useState, useReducer } from 'react' ; import ReactDOM from 'react-dom' ; function reducer ( state, action ) { switch (action.type) { case 'up' : return { count : state.count + 1 }; case 'down' : return { count : state.count - 1 }; } } function Counter ( ) { const [state, dispatch] = useReducer(reducer, { count : 1 }) return ( < div > {state.count} < button onClick = {() => dispatch({ type: 'up' })}>+ </ button > < button onClick = {() => dispatch({ type: 'down' })}>+ </ button > </ div > ); } ReactDOM.render( < Counter /> , document.getElementById('root')); 复制代码`

### useEffect ###

一个至关重要的 Hooks API，顾名思义，useEffect 是用于处理各种状态变化造成的副作用，也就是说只有在特定的时刻，才会执行的逻辑。

` import { useState, useEffect } from 'react' ; import ReactDOM from 'react-dom' ; function Example ( ) { const [count, setCount] = useState( 0 ); // => componentDidMount/componentDidUpdate useEffect( () => { // update document.title = `You clicked ${count} times` ; // => componentWillUnMount return function cleanup ( ) { document.title = 'app' ; } }, [count]); return ( < div > < p > You clicked {count} times </ p > < button onClick = {() => setCount(count + 1)}> Click me </ button > </ div > ); } ReactDOM.render( < Example /> , document.getElementById('root')); 复制代码`

### useMemo ###

useMemo 主要用于渲染过程优化，两个参数依次是计算函数（通常是组件函数）和依赖状态列表，当依赖的状态发生改变时，才会触发计算函数的执行。如果没有指定依赖，则每一次渲染过程都会执行该计算函数。

` const memoizedValue = useMemo( () => computeExpensiveValue(a, b), [a, b]); 复制代码` ` import { useState, useMemo } from 'react' ; import ReactDOM from 'react-dom' ; function Time ( ) { return < p > {Date.now()} </ p > ; } function Counter ( ) { const [count, setCount] = useState( 0 ); const memoizedChildComponent = useMemo( ( count ) => { return < Time /> ; }, [count]); return ( < div > < h1 > {count} </ h1 > < button onClick = {() => setCount(count + 1)}>+ </ button > < div > {memoizedChildComponent} </ div > </ div > ); } ReactDOM.render( < Counter /> , document.getElementById('root')); 复制代码`

### useContext ###

context 是在外部 create ，内部 use 的 state，它和全局变量的区别在于，如果多个组件同时 useContext，那么这些组件都会 rerender，如果多个组件同时 useState 同一个全局变量，则只有触发 setState 的当前组件 rerender。

[示例-未使用 useContext]( https://link.juejin.im?target=https%3A%2F%2Fcodepen.io%2Fx-cold%2Fpen%2FOYJGKQ )

` import { useState, useContext, createContext } from 'react' ; import ReactDOM from 'react-dom' ; // 1. 使用 createContext 创建上下文 const UserContext = new createContext(); // 2. 创建 Provider const UserProvider = props => { let [username, handleChangeUsername] = useState( '' ); return ( <UserContext.Provider value={{ username, handleChangeUsername }}> {props.children} </UserContext.Provider> ); }; // 3. 创建 Consumer const UserConsumer = UserContext.Consumer; // 4. 使用 Consumer 包裹组件 const Pannel = () => ( <UserConsumer> {({ username, handleChangeUsername }) => ( <div> <div>user: {username}</div> <input onChange={e => handleChangeUsername(e.target.value)} /> </div> )} </UserConsumer> ); const Form = () => <Pannel />; const App = () => ( <div> <UserProvider> <Form /> </UserProvider> </div> ); ReactDOM.render(<App />, document.getElementById('root')); 复制代码`

[示例 - 使用 useContext]( https://link.juejin.im?target=https%3A%2F%2Fcodepen.io%2Fx-cold%2Fpen%2FGaRLqZ%3Feditors%3D0010 )

` import { useState, useContext, createContext } from 'react' ; import ReactDOM from 'react-dom' ; // 1. 使用 createContext 创建上下文 const UserContext = new createContext(); // 2. 创建 Provider const UserProvider = props => { let [username, handleChangeUsername] = useState( '' ); return ( <UserContext.Provider value={{ username, handleChangeUsername }}> {props.children} </UserContext.Provider> ); }; const Pannel = () => { const { username, handleChangeUsername } = useContext(UserContext); // 3. 使用 Context return ( <div> <div>user: {username}</div> <input onChange={e => handleChangeUsername(e.target.value)} /> </div> ); }; const Form = () => <Pannel />; const App = () => ( <div> <UserProvider> <Form /> </UserProvider> </div> ); ReactDOM.render(<App />, document.getElementById('root')); 复制代码`

### useRef ###

useRef 返回一个可变的 ref 对象，其 .current 属性初始化为传递的参数（initialValue）。返回的对象将持续整个组件的生命周期。事实上 useRef 是一个非常有用的 API，许多情况下，我们需要保存一些改变的东西，它会派上大用场的。

[示例]( https://link.juejin.im?target=https%3A%2F%2Fcodepen.io%2Fx-cold%2Fpen%2FEzxMPw )

` function TextInputWithFocusButton ( ) { const inputEl = useRef( null ); const onButtonClick = () => { // `current` points to the mounted text input element inputEl.current.focus(); }; return ( <> <input ref={inputEl} type="text" /> <button onClick={onButtonClick}>Focus the input</button> </> ); } 复制代码`

## React 状态共享方案 ##

说到状态共享，最简单和直接的方式就是通过 props 逐级进行状态的传递，这种方式耦合于组件的父子关系，一旦组件嵌套结构发生变化，就需要重新编写代码，维护成本非常昂贵。随着时间的推移，官方推出了各种方案来解决状态共享和代码复用的问题。

### Mixins ###

![image.png](https://user-gold-cdn.xitu.io/2019/6/3/16b1aec6deac2e80?imageView2/0/w/1280/h/960/ignore-error/1)
React 中，只有通过 createClass 创建的组件才能使用 mixins。这种高耦合，依赖难以控制，复杂度高的方式随着 ES6 的浪潮逐渐淡出了历史舞台。

### HOC ###

高阶组件源于函数式编程，由于 React 中的组件也可以视为函数（类），因此天生就可以通过 HOC 的方式来实现代码复用。可以通过属性代理和反向继承来实现，HOC 可以很方便的操控渲染的结果，也可以对组件的 props / state 进行操作，从而可以很方便的进行复杂的代码逻辑复用。

` import React from 'react' ; import PropTypes from 'prop-types' ; // 属性代理 class Show extends React. Component { static propTypes = { children : PropTypes.element, visible : PropTypes.bool, }; render() { const { visible, children } = this.props; return visible ? children : null ; } } // 反向继承 function Show2 ( WrappedComponent ) { return class extends WrappedComponent { render() { if ( this.props.visible === false ) { return null ; } else { return super.render(); } } } } function App ( ) { return ( < Show visible = {Math.random() > 0.5}>hello </ Show > ); } 复制代码`

Redux 中的状态复用是一种典型的 HOC 的实现，我们可以通过 compose 来将数据组装到目标组件中，当然你也可以通过装饰器的方式进行处理。

` import React from 'react' ; import { connect } from 'react-redux' ; // use decorator @connect( state => ({ name : state.user.name })) class App extends React. Component { render() { return < div > hello, {this.props.name} </ div > } } // use compose connect( ( state ) => ({ name : state.user.name }))(App); 复制代码`

### Render Props ###

显而易见，renderProps 就是一种将 render 方法作为 props 传递到子组件的方案，相比 HOC 的方案，renderProps 可以保护原有的组件层次结构。

` import React from 'react' ; import ReactDOM from 'react-dom' ; import PropTypes from 'prop-types' ; // 与 HOC 不同，我们可以使用具有 render prop 的普通组件来共享代码 class Mouse extends React. Component { static propTypes = { render : PropTypes.func.isRequired } state = { x : 0 , y : 0 }; handleMouseMove = ( event ) => { this.setState({ x : event.clientX, y : event.clientY }); } render() { return ( < div style = {{ height: ' 100 %' }} onMouseMove = {this.handleMouseMove} > {this.props.render(this.state)} </ div > ); } } function App ( ) { return ( < div style = {{ height: ' 100 %' }}> < Mouse render = {({ x , y }) => ( // render prop 给了我们所需要的 state 来渲染我们想要的 < h1 > The mouse position is ({x}, {y}) </ h1 > )}/> </ div > ); } ReactDOM.render( < App /> , document.getElementById('root')); 复制代码`

### Hooks ###

通过组合 Hooks API 和 React 内置的 Context，从前面的示例可以看到通过 Hook 让组件之间的状态共享更清晰和简单。

## React Hooks 设计理念 ##

### 基本原理 ###

![image.png](https://user-gold-cdn.xitu.io/2019/6/3/16b1aec6ec1b8f0c?imageView2/0/w/1280/h/960/ignore-error/1)

` function FunctionalComponent ( ) { const [state1, setState1] = useState( 1 ); const [state2, setState2] = useState( 2 ); const [state3, setState3] = useState( 3 ); } 复制代码`

![image.png](https://user-gold-cdn.xitu.io/2019/6/3/16b1aec6e1ae41e1?imageView2/0/w/1280/h/960/ignore-error/1)

` { memoizedState : 'foo' , next : { memoizedState : 'bar' , next : { memoizedState : 'bar' , next : null } } } 复制代码`

### 函数式贯彻到底 ###

#### capture props ####

函数组件天生就是支持 props 的，基本用法上和 class 组件没有太大的差别。需要注意的两个区别是：

* class 组件 props 挂载在 this 上下文中，而函数式组件通过形参传入；
* 由于挂载位置的差异，class 组件中如果 this 发生了变化，那么 this.props 也会随之改变；而在函数组件里 props 始终是不可变的，因此遵守 capture value 原则（即获取的值始终是某一时刻的），Hooks 也遵循这个原则。

通过一个 [示例]( https://link.juejin.im?target=https%3A%2F%2Fcodesandbox.io%2Fs%2Fpjqnl16lm7 ) 来理解一下 capture value，我们可以通过 useRef 来规避 capture value，因为 useRef 是可变的。

#### state ####

+----------+--------------------------------------------------+--------------------------------------------------------------------------------------+
|          |                    CLASS 组件                    |                                       函数组件                                       |
+----------+--------------------------------------------------+--------------------------------------------------------------------------------------+
| 创建状态 | this.state = {}                                  | useState, useReducer                                                                 |
| 修改状态 | this.setState()                                  | set function                                                                         |
| 更新机制 | 异步更新，多次修改合并到上一个状态，产生一个副本 | 同步更新，直接修改为目标状态                                                         |
| 状态管理 | 一个 state 集中式管理多个状态                    | 多个 state，可以通过                                                                 |
|          |                                                  | useReducer                                                                           |
|          |                                                  | 进行状态合并（手动）                                                                 |
| 性能     | 高                                               | 如果 useState                                                                        |
|          |                                                  | 初始化状态需要通过非常复杂的计算得到，请使用函数的声明方式，否则每次渲染都会重复执行 |
|          |                                                  |                                                                                      |
+----------+--------------------------------------------------+--------------------------------------------------------------------------------------+

#### 生命周期 ####

* componentDidMount / componentDidUpdate / componentWillUnMount

useEffect 在每一次渲染都会被调用，稍微包装一下就可以作为这些生命周期使用；

* shouldComponentUpdate

通常我们优化组件性能时，会优先采用纯组件的方式来减少单个组件的渲染次数。

` class Button extends React. PureComponent {} 复制代码`

React Hooks 中可以采用 useMemo 代替，可以实现仅在某些数据变化时重新渲染组件，等同于自带了 shallowEqual 的 shouldComponentUpdate。

#### 强制渲染 forceUpdate ####

由于默认情况下，每一次修改状态都会造成重新渲染，可以通过一个不使用的 set 函数来当成 forceUpdate。

` const forceUpdate = () => useState( 0 )[ 1 ]; 复制代码`

### 实现原理 ###

## 基于 Hooks，增强 Hooks ##

### 来一套组合拳吧！ ###

由于每一个 Hooks API 都是纯函数的概念，使用时更关注输入 (input) 和输出 (output)，因此可以更好的通过组装函数的方式，对不同特性的基础 Hooks API 进行组合，创造拥有新特性的 Hooks。

* useState 维护组件状态
* useEffect 处理副作用
* useContext 监听 provider 更新变化

### useDidMount ###

` import { useEffect } from 'react' ; const useDidMount = fn => useEffect( () => fn && fn(), []); export default useDidMount; 复制代码`

### ###

### useDidUpdate ###

` import { useEffect, useRef } from 'react' ; const useDidUpdate = ( fn, conditions ) => { const didMoutRef = useRef( false ); useEffect( () => { if (!didMoutRef.current) { didMoutRef.current = true ; return ; } // Cleanup effects when fn returns a function return fn && fn(); }, conditions); }; export default useDidUpdate 复制代码`

### useWillUnmount ###

在讲到 useEffect 时已经提及过，其允许返回一个 cleanup 函数，组件在取消挂载时将会执行该 cleanup 函数，因此 useWillUnmount 也能轻松实现~

` import { useEffect } from 'react' ; const useWillUnmount = fn => useEffect( () => () => fn && fn(), []); export default useWillUnmount; 复制代码`

### useHover ###

[示例]( https://link.juejin.im?target=https%3A%2F%2Fcodepen.io%2Fx-cold%2Fpen%2FjoOXxK )

` // lib/onHover.js import { useState } from 'react' ; const useHover = () => { const [hovered, set] = useState( false ); return { hovered, bind : { onMouseEnter : () => set( true ), onMouseLeave : () => set( false ), }, }; }; export default useHover; 复制代码` ` import { useHover } from './lib/onHover.js' ; function Hover ( ) { const { hovered, bind } = useHover(); return ( < div > < div {...bind }> hovered: {String(hovered)} </ div > </ div > ); } 复制代码`

### useField ###

[示例]( https://link.juejin.im?target=https%3A%2F%2Fcodepen.io%2Fx-cold%2Fpen%2FrgNPNB )

` // lib/useField.js import { useState } from 'react' ; const useField = ( initial ) => { const [value, set] = useState(initial); return { value, set, reset : () => set(initial), bind : { value, onChange : e => set(e.target.value), }, }; } export default useField; 复制代码` ` import { useField } from 'lib/useField' ; function Input { const { value, bind } = useField( 'Type Here...' ); return ( < div > input text: {value} < input type = "text" {...bind } /> </ div > ); } function Select() { const { value, bind } = useField('apple') return ( < div > selected: {value} < select {...bind }> < option value = "apple" > apple </ option > < option value = "orange" > orange </ option > </ select > </ div > ); } 复制代码`

## 注意事项 ##

* Hook 的使用范围：函数式的 React 组件中、自定义的 Hook 函数里；
* Hook 必须写在函数的最外层，每一次 useState 都会改变其下标 (cursor)，React 根据其顺序来更新状态；
* 尽管每一次渲染都会执行 Hook API，但是产生的状态 (state) 始终是一个常量（作用域在函数内部）；

## 结语 ##

React Hooks 提供为状态管理提供了新的可能性，尽管我们可能需要额外去维护一些内部的状态，但是可以避免通过 renderProps / HOC 等复杂的方式来处理状态管理的问题。Hooks 带来的好处如下：

* 更细粒度的代码复用，并且不会产生过多的副作用
* 函数式编程风格，代码更简洁，同时降低了使用和理解门槛
* 减少组件嵌套层数
* 组件数据流向更清晰

事实上，通过定制各种场景下的自定义 Hooks，能让我们的应用程序更方便和简洁，组件的层次结构也能保证完好，还有如此令人愉悦的函数式编程风格，Hooks 在 React 16.8.0 版本已经正式发布稳定版本，现在开始用起来吧！！！

## 参考资料 ##

* [hooks-reference]( https://link.juejin.im?target=https%3A%2F%2Freactjs.org%2Fdocs%2Fhooks-reference.html )
* [react-hooks-lib]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fbeizhedenglong%2Freact-hooks-lib )
* [【React深入】从Mixin到HOC再到Hook]( https://juejin.im/post/5cad39b3f265da03502b1c0a )
* [How Are Function Components Different from Classes?]( https://link.juejin.im?target=https%3A%2F%2Foverreacted.io%2Fhow-are-function-components-different-from-classes%2F )
* [under-the-hood-of-reacts-hooks-system]( https://link.juejin.im?target=https%3A%2F%2Fmedium.com%2Fthe-guild%2Funder-the-hood-of-reacts-hooks-system-eb59638c9dba )
* [阅读源码后，来讲讲React Hooks是怎么实现的]( https://juejin.im/post/5bdfc1c4e51d4539f4178e1f )

![image.png](https://user-gold-cdn.xitu.io/2019/6/3/16b1aec6e222ed1d?imageView2/0/w/1280/h/960/ignore-error/1)