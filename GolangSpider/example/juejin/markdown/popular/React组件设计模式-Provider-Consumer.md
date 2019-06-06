# React组件设计模式-Provider-Consumer #

我们都知道，基于props做组件的跨层级数据传递是非常困难并且麻烦的，中间层组件要为了传递数据添加一些无用的props。 而React自身早已提供了context API来解决这种问题，但是16.3.0之前官方都建议不要使用，认为会迟早会被废弃掉。说归说，很多库已经采用了 context API。可见呼声由多么强烈。终于在16.3.0之后的版本，React正式提供了稳定的context API，本文中的示例基于v16.3.0之后的context API。

## 概念 ##

首先要理解上下文（context）的作用以及提供者和消费者分别是什么，同时要思考这种模式解决的是什么问题（跨层级组件通信）。

context做的事情就是创建一个上下文对象，并且对外暴露 ` 提供者（通常在组件树中上层的位置）` 和 ` 消费者` ，在上下文之内的所有子组件， 都可以访问这个上下文环境之内的数据，并且不用通过props。可以理解为有一个集中管理state的对象，并限定了这个对象可访问的范围， 在范围之内的子组件都能获取到它内部的值。

提供者为消费者提供context之内的数据，消费者获取提供者为它提供的数据，自然就解决了上边的问题。

## 用法 ##

这里要用到一个小例子，功能就是主题颜色的切换。效果如图：

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2c429ad2f6fd7?imageslim)

根据上边的概念和功能，分解一下要实现的步骤：

* 创建一个上下文，来提供给我们提供者和消费者
* 提供者提供数据
* 消费者获取数据

这里的文件组织是这样的：

` ├─context.js // 存放context的文件 │─index.js // 根组件，Provider所在的层级 │─Page.js // 为了体现跨层级通信的添加的一个中间层级组件，子组件为Title和Paragraph │─Title.js // 消费者所在的层级 │─Paragraph.js // 消费者所在的层级 复制代码`

### 创建一个上下文 ###

` import React from 'react' const ThemeContext = React.createContext() export const ThemeProvider = ThemeContext.Provider export const ThemeConsumer = ThemeContext.Consumer 复制代码`

这里， ` ThemeContext` 就是一个被创建出来的上下文，它内部包含了两个属性，看名字就可以知道，一个是提供者一个是消费者。 Provider和Consumer是成对出现的，每一个Provider都会对应一个Consumer。而每一对都是由React.createContext()创建出来的。

### page组件 ###

没啥好说的，就是一个容器组件而已

` const Page = () => <> <Title/> <Paragraph/> </> 复制代码`

### 提供者提供数据 ###

提供者一般位于比较上边的层级，ThemeProvider 接受的value就是它要提供的上下文对象。

` // index.js import { ThemeProvider } from './context' render () { const { theme } = this.state return <ThemeProvider value={{ themeColor: theme }}> <Page/> </ThemeProvider> } 复制代码`

### 消费者获取数据 ###

在这里，消费者使用了renderProps模式，Consumer会将上下文的数据作为参数传入renderProps渲染的函数之内，所以这个函数内才可以访问上下文的数据。

` // Title.js 和 Paragraph的功能是一样的，代码也差不多，所以单放了Title.js import React from 'react' import { ThemeConsumer } from './context' class Title extends React.Component { render () { return <ThemeConsumer> { theme => <h1 style={{ color: theme.themeColor }}> title </h1> } </ThemeConsumer> } } 复制代码`

## 关于嵌套上下文 ##

此刻你可能会产生疑问，就是应用之内不可能只会有一个context。那多个context如果发生嵌套了怎么办？

### v16.3.0之前的版本 ###

其实v16.3.0之前版本的React的context的设计上考虑到了这种场景。只不过实现上麻烦点。来看一下具体用法： 和当前版本的用法不同的是，Provider和Consumer不是成对被创建的。

Provider是一个普通的组件，当然，是需要位于Consumer组件的上层。要创建它，我们需要用到两个方法：

* getChildContext: 提供 ` 自身范围` 上下文的数据
* childContextTypes：声明 ` 自身范围` 的上下文的结构

` class ThemeProvider extends React.Component { getChildContext () { return { theme: this.props.value }; } render () { return ( <React.Fragment> {this.props.children} </React.Fragment> ); } } ThemeProvider.childContextTypes = { theme: PropTypes.object }; 复制代码`

再看消费者，需要用到 ` contextTypes` ，来声明接收的上下文的结构。

` const Title = (props, context) => { const {textColor} = context.theme; return ( <p style={{color: color}}> 我是标题 </p> ); }; Title.contextTypes = { theme: PropTypes.object }; 复制代码`

最后的用法：

` <ThemeProvider value={{color: 'green' }} > <Title /> </ThemeProvider> 复制代码`

回到嵌套的问题上，大家看出如何解决的了吗？

Provider做了两件事，提供context数据，然后。又声明了这个context范围的数据结构。而Consumer呢，通过contextTypes定义接收到的context数据结构。 也就相当于Consumer指定了要接收哪种结构的数据，而这种结构的数据又是由某个Provider提前定义好的。通过这种方式，再多的嵌套也不怕，Consumer只要定义 接收谁声明的context的结构就好了。如果不定义的话，是接收不到context的数据的。

### v16.3.0之后的版本 ###

v16.3.0之后的版本使用起来比以前简单了很多。解决嵌套问题的方式也更优雅。由于Provider和Consumer是成对地被创建出来的。即使这一对的Provider于另一对的 Consumer的数据结构和值的类型相同，这个Consumer也让能访问那个Provider的上下文。这便是解决方法。

## 总结 ##

对于这个context这个东西。我感觉还是不要在应用里大量使用。就像React-Redux的Provider，或者antd的LocalProvider，差不多用一次就够，因为用多会使应用里很混乱， 组件之间的依赖关系变得复杂。但是React为我们提供的这个api还是可以看到它自身还是想弥补其状态管理的短板的，况且Hooks中的useReducer出现后，更说明了这一点。