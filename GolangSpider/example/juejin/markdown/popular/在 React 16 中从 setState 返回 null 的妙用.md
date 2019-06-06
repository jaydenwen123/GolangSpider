# 在 React 16 中从 setState 返回 null 的妙用 #

> 
> 
> 
> 翻译：疯狂的技术宅
> 
> 
> 
> 原文： [blog.logrocket.com/returning-n…](
> https://link.juejin.im?target=https%3A%2F%2Fblog.logrocket.com%2Freturning-null-from-setstate-in-react-16-5fdb1c35d457%2F
> )
> 
> 
> 
> **未经允许严禁转载**
> 
> 

### 概述 ###

在 React 16 中为了防止不必要的 DOM 更新，允许你决定是否让 `.setState` 更来新状态。在调用 `.setState` 时返回 ` null` 将不再触发更新。

我们将通过重构一个 mocktail （一种不含酒精的鸡尾酒）选择程序来探索它是如何工作的，即使我们选择相同的 mocktail 两次也会更新。

![我们的 mocktail 选择程序](https://user-gold-cdn.xitu.io/2019/6/4/16b222ffc6288793?imageslim)

目录结构如下所示：

` src |-> App.js |-> Mocktail.js |-> index.js |-> index.css |-> Spinner.js 复制代码`

### 我们的程序如何工作 ###

我们的程序将显示一个被选中的 mocktail。可以通过单击按钮来选择或切换 mocktail。这时会加载一个新的 mocktail，并在加载完成后渲染出这个 mocktail 的图像。

` App` 组件的父组件有 ` mocktail` 状态和 ` updateMocktail` 方法，用于处理更新 mocktail。

` import React, { Component } from 'react' ; import Mocktail from './Mocktail' ; class App extends Component { state = { mocktail : '' } updateMocktail = mocktail => this.setState({ mocktail }) render() { const mocktails = [ 'Cosmopolitan' , 'Mojito' , 'Blue Lagoon' ]; return ( <React.Fragment> <header> <h1>Select Your Mocktail</h1> <nav> { mocktails.map((mocktail) => { return <button key={mocktail} value={mocktail} type="button" onClick={e => this.updateMocktail(e.target.value)}>{mocktail}</button> }) } </nav> </header> <main> <Mocktail mocktail={this.state.mocktail} /> </main> </React.Fragment> ); } } export default App; 复制代码`

在 ` button` 元素的 ` onClick` 事件上调用 ` updateMocktail` 方法， ` mocktail` 状态被传递给子组件 ` Mocktail` 。

` Mocktail` 组件有一个名为 ` isLoading` 的加载状态，当其为 ` true` 时会渲染 ` Spinner` 组件。

` import React, { Component } from 'react' ; import Spinner from './Spinner' ; class Mocktail extends Component { state = { isLoading : false } componentWillReceiveProps() { this.setState({ isLoading : true }); setTimeout( () => this.setState({ isLoading : false }), 500 ); } render() { if ( this.state.isLoading) { return < Spinner /> } return ( <React.Fragment> <div className="mocktail-image"> <img src={`img/${this.props.mocktail.replace(/ +/g, "").toLowerCase()}.png`} alt={this.props.mocktail} /> </div> </React.Fragment> ); } } export default Mocktail; 复制代码`

在 ` Mocktail` 组件的 ` componentWillReceiveProps` 生命周期方法中调用 ` setTimeout` ，将加载状态设置为 ` true` 达 500 毫秒。

每次使用新的 ` mocktail` 状态更新 ` Mocktail` 组件的 props 时，它会用半秒钟显示加载动画，然后渲染 mocktail 图像。

### 问题 ###

现在的问题是，即使状态没有改变， ` mocktail` 状态也会被更新，同时触发重新渲染 ` Mocktail` 组件。

例如每当单击 **Mojito** 按钮时，我们都会看到程序对 Mojito 图像进行了不必要地重新渲染。 React 16 对状态性能进行了改进，如果新的状态值与其现有值相同的话，通过在 ` setState` 中返回 ` null` 来防止来触发更新。

![img](https://user-gold-cdn.xitu.io/2019/6/4/16b222ffc646a843?imageslim)

### 解决方案 ###

以下是我们将要遵循的步骤，来防止不必要的重新渲染：

* 检查新的状态值是否与现有值相同
* 如果值相同，我们将返回 ` null`
* 返回 ` null` 将不会更新状态和触发组件重新渲染

首先，在 ` app` 组件的 ` updateMocktail` 方法中，创建一个名为 ` newMocktail` 的常量，并用传入的 ` mocktail` 值为其赋值。

` updateMocktail = mocktail => { const newMocktail = mocktail; this.setState({ mocktail }) } 复制代码`

因为我们需要基于之前的状态检查和设置状态，而不是传递 ` setState` 和 ` object` ，所以我们需要传递一个以前的状态作为参数的函数。然后检查 ` mocktail` 状态的新值是否与现有值相同。

如果值相同， ` setState` 将返回 ` null` 。否则 ` setState` 返回更新的 ` mocktail` 状态，这将触发使用新状态重新渲染 ` Mocktail` 组件。

` updateMocktail = mocktail => { const newMocktail = mocktail; this.setState( state => { if (state.mocktail === newMocktail) { return null ; } else { return { mocktail }; } }) } 复制代码`

![img](https://user-gold-cdn.xitu.io/2019/6/4/16b222ffc6b3b8ad?imageslim)

现在单击按钮仍会加载其各自的 mocktail 图像。但是，如果我们再次单击同一个mocktail按钮，React 不会重新渲染 ` Mocktail` 组件，因为 ` setState` 返回 ` null` ，所以状态没有改变，也就不会触发更新。

我在下面的两个 GIF 中突出显示了 React DevTools 中的更新：

![没有从 setState 返回 null](https://user-gold-cdn.xitu.io/2019/6/4/16b222ffc6af6619?imageslim)

没有从 setState 返回 null

![从 setState 返回 null 之后](https://user-gold-cdn.xitu.io/2019/6/4/16b222ffc6ca85a3?imageslim)

从 setState 返回 null 之后
> 
> 
> 
> **注意：**我在这里换了一个深色主题，以便更容易观察到 React DOM 中的更新。
> 
> 

### 总结 ###

本文介绍了在 React 16 中怎样从 ` setState` 返回 ` null` 。我在下面的 ` CodeSandbox` 中添加了 ` mocktail` 选择程序的完整代码，供你使用和 fork。

CodeSandbox： [codesandbox.io/embed/vj8wk…]( https://link.juejin.im?target=https%3A%2F%2Fcodesandbox.io%2Fembed%2Fvj8wk0mzjy )

通过使用 ` null` 可以防止不必要的状态更新和重新渲染，这样使我们的程序执行得更快，从而改善程序的用户体验。

用户偶然发现我们的产品，他们对产品的看法直接反映了对公司及其产品的看法，因此我们需要以自然和直观的方式围绕用户的期望去构建体验。

希望本文能够对你有所帮助。感谢阅读！

## 欢迎关注前端公众号：前端先锋，获取前端工程化使用工具包。 ##

![](https://user-gold-cdn.xitu.io/2019/6/4/16b22313e64c35e3?imageView2/0/w/1280/h/960/ignore-error/1)