# React组件设计模式-Render-props #

写业务时，我们经常需要抽象一些使用频率较高的逻辑，但是除了高阶组件可以抽象逻辑，RenderProps也是一种比较好的方法。

RenderProps，顾名思义就是将组件的props渲染出来。实际上是让组件的props接收函数，由函数来渲染内容。将通用的逻辑 抽象在该组件的内部，然后依据业务逻辑来调用函数（props内渲染内容的函数），从而达到重用逻辑的目的。

### 简单实现 ###

我们先看一个最简单的RenderProps模式的实现：

` const RenderProps = props => <> {props.children(props)} </> 复制代码`

在这里，RenderProps组件的子组件是一个函数 ` props.children(props)` ，而props.children返回的是UI元素。

使用时的代码如下：

` <RenderProps> {() => <>Hello RenderProps</>} </RenderProps> 复制代码`

以上未作任何的业务逻辑处理，有什么用处呢？我们可以想到，可以在 RenderProps 组件中去用代码操控返回的结果。 以最常见的用户登录逻辑为例，希望在登陆之后才可以看到内容，否则展示请登录：

` const Auth = props => { const userName = getUserName() if (userName) { const allProps = {userName, ...props} return <> {props.children(allProps)} </> } else { return <>请登录</> } } <Auth> {({userName}) => <>Hello！{userName}</>} </Auth> 复制代码`

` props.children(allProps)` 就相当于Auth组件嵌套的 ` ({userName}) => <>Hello！{userName}</>`

上边的例子中，用户若已经登陆，getUserName返回用户名，否则返回空。这样我们就可以判断返回哪些内容了。 当然，上边通过renderProps传入了userName,这属于Auth组件的增强功能。

### 函数名不仅可以是children ###

平时一般使用的时候，props.children都是具体的组件实例，但上边的实现是基于以函数为子组件（ ` children(props)` ），被调用返回UI。 同样，可以调用props中的任何函数。还是以上边的逻辑为例：

` const Auth = props => { const userName = 'Mike' if (userName) { const allProps = { userName, ...props } return <>{props.login(allProps)}</> } else { return <> {props.noLogin(props)} </> } } 复制代码`

使用方法如下：

` <Auth login={({userName}) => <h1>Hello {userName}</h1>} noLogin={() => <h1>please login</h1>} /> 复制代码`

这里，Auth组件的props接收两个函数： **login(表示已经登录)** ， **noLogin(表未登录)** ， Auth组件内部，通过判断是否登陆来决定显示哪个组件。

### 总结 ###

render-props作为一种抽象通用逻辑的方法，其本身也会遇到像高阶组件那样层层嵌套的问题。

` <GrandFather> {Props => { <Father> {props => { <Son {...props} />; }} </Father>; }} </GrandFather> 复制代码`

但和高阶组件不同的是，由于渲染的是函数(高阶组件渲染的是组件)，就为利用compose提供了机会。例如 ` react-powerplugin` 。

` import { compose } from 'react-powerplug' const ComposeComponent = compose( <GrandFather />, <Father /> ) <ComposeComponent> {props => { <Son {...props} />; }} <ComposeComponent/> 复制代码`

还有 ` Epitath` 也提供了一种新模式来解决这个问题。这部分展开来说的话是另一个话题了，我也在摸索中。