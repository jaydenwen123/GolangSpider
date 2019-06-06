# React 事件绑定有好几种方式，我快懵了 #

事件处理程序(Event handlers)就是每当有事件被触发时决定要执行的动作或行为。

在 React 应用中，事件名使用小驼峰格式书写，意思就是 ` onclick` 要写成 ` onClick` 。

React 实现的 [合成事件]( https://link.juejin.im?target=https%3A%2F%2Fzh-hans.reactjs.org%2Fdocs%2Fevents.html ) 机制给 React 应用和接口带来了一致性，同时具备高性能的优点。它通过将事件标准化来达到一致性，从而达到在不同浏览器和平台上都具有相同的属性。

合成事件是浏览器的原生事件的跨浏览器包装器。除兼容所有浏览器外，它还拥有和浏览器原生事件相同的接口，包括 ` stopPropagation()` 和 ` preventDefault()` 。

合成事件的高性能是通过自动使用事件代理实现的。事实上，React 并不将事件处理程序绑定到节点自身，而是将一个事件监听器绑定到 document 的根元素上。每当有事件被触发，React 就将事件监听器映射到恰当的组件元素上。

### 监听事件 ###

在 React 中监听事件简单如下：

` class ShowAlert extends React. Component { showAlert() { alert( "Im an alert" ); } render() { return < button onClick = {this.showAlert} > show alert </ button > ; } } 复制代码`

在上面的例子中， ` onClick` 属性是我们的事件处理程序，它被添加到目标元素上以便在元素被点击时执行要被执行的函数。 ` onClick` 属性设置了 ` showAlert` 函数。

简单点儿说，就是无论何时点击该按钮， ` showAlert` 函数都会被调用并显示一条信息。

### 绑定方法 ###

在 JavaScript 中，类方法并不是默认绑定的。因此，将函数绑定到类的实例上十分重要。

#### 在 ` render()` 中绑定 ####

这种方式是在 ` render` 函数中调用 ` bind` 方法来进行绑定的:

` class ChangeInput extends Component { constructor (props) { super (props); this.state = { name : "" }; } changeText(event) { this.setState({ name : event.target.value }); } render() { return ( < div > < label htmlFor = "name" > Enter Text here </ label > < input type = "text" id = "name" onChange = {this.changeText.bind(this)} /> < h3 > {this.state.name} </ h3 > </ div > ); } } 复制代码`

在上面的例子中，我们使用 ` onChange` 事件处理程序在 input 输入框上监听键盘事件，该操作是通过在 render 函数中绑定完成的。该方法需要在 render 函数中调用 `.bind(this)` 。

为啥呢？

任何 ES6 类中的方法都是普通的 JavaScript 函数，因此其都从 [Function]( https://link.juejin.im?target=https%3A%2F%2Fdeveloper.mozilla.org%2Fen-US%2Fdocs%2FWeb%2FJavaScript%2FReference%2FGlobal_Objects%2FFunction ) 的原型上继承 ` bind()` 方法。现在，当我们在 JSX 内部调用 ` onChange` 时， ` this` 会指向我们的组件实例。

使用这种方法可能会造成一些潜在的性能问题，因为函数在每次 render 后都被重新分配一次。这种性能代价可能在小型 React 应用上并不明显，但在较大应用上就要值得注意了。

#### 在 ` constructor()` 中绑定 ####

如果说在 render 中绑定的方法不适合你，你也可以在 ` constructor()` 中进行绑定。例子如下：

` class ChangeInput extends Component { constructor (props) { super (props); this.state = { name : "" }; this.changeText = this.changeText.bind( this ); } changeText(event) { this.setState({ name : event.target.value }); } render() { return ( < div > < label htmlFor = "name" > Enter Text here </ label > < input type = "text" id = "name" onChange = {this.changeText} /> < h3 > {this.state.name} </ h3 > </ div > ); } } 复制代码`

如你所见， ` changeText` 函数绑定在 ` constructor` 上：

` this.changeText = this.changeText.bind( this ) 复制代码`

等号左边的 ` this.changeText` 指向 ` changeText` 方法，由于该步操作是在 ` constructor` 中完成的，所以 ` this` 指向 ` ChangeInput` 组件。

等号右边的 ` this.changeText` 指向相同的 ` changeText` 方法，但是我们现在在其上调用了 `.bind()` 方法。

括号中的 ` this` 作为我们传入 `.bind()` 中的参数指向的是上下文(context)，也就是指向 ` ChangeInput` 组件。

同样值得注意的是如果 ` changeText` 不绑定到组件实例上的话，它就不能访问 ` this.setState` ，因为这时的 ` this` 会是 ` undefined` 。

#### 使用箭头函数进行绑定 ####

绑定事件处理函数的另一种方式是通过箭头函数。通过这种 ES7 的类特性(实验性的 [public class fields 语法]( https://link.juejin.im?target=https%3A%2F%2Fbabeljs.io%2Fdocs%2Fen%2Fbabel-plugin-proposal-class-properties ) )，我们可以在方法定义时就进行事件绑定，例子如下：

` class ChangeInput extends Component { handleEvent = event => { alert( "I was clicked" ); }; render() { return ( < div > < button onClick = {this.handleEvent} > Click on me </ button > </ div > ); } } 复制代码`

箭头函数表达式语法比传统的函数表达式简短且没有自己的 ` this` 、 ` arguments` 、 ` super` 和 ` new.target` 。

在上面的例子中，一旦组件被创建， ` this.handleEvent` 就不会再有变化了。这种方法非常简单且容易阅读。

同在 ` render` 函数中绑定的方法一样，这种方法也有其性能代价。

#### 使用匿名箭头函数在 ` render` 中绑定 ####

绑定的时候定义一个匿名函数（箭头函数），将匿名函数与元素绑定，而不是直接绑定事件处理函数，这样 ` this` 在匿名函数中就不是 ` undefined` 了：

` class LoggingButton extends React. Component { handleClick() { console.log( 'this is:' , this ); } render() { return ( < button onClick = {(e) => this.handleClick(e)}> Click me </ button > ); } } 复制代码`

这种方法有一个问题，每次 ` LoggingButton` 渲染的时候都会创建一个新的 ` callback` 匿名函数。在这个例子中没有问题，但是如果将这个 ` onClick` 的值作为 ` props` 传给子组件的时候，将会导致子组件重新 ` render` ，所以不推荐。

### 定制组件和事件 ###

每当谈到 React 中的事件，只有 DOM 元素可以有事件处理程序。举个栗子，比如说现在有个组件叫 ` CustomButton` ，其有一个 ` onClick` 事件，但是当你点击时不会有任何反应，原因就是前面所说的。

那么我们要如何处理定制组件中的事件绑定呢？

答案就是通过在 ` CustomButton` 组件中渲染一个 DOM 元素然后把 ` onClick` 作为 prop 传进去， ` CustomButton` 组件实际上只是为点击事件充当了传递介质。

` class CustomButton extends Component { render() { const { onPress, children } = this.props; return ( < button type = "button" onClick = {onPress} > {children} </ button > ); } } class ChangeInput extends Component { handleEvent = () => { alert( "I was clicked" ); }; render() { return ( < div > < CustomButton onPress = {this.handleEvent} > Click on me </ CustomButton > </ div > ); } } 复制代码`

本例中， ` CustomButton` 组件接收了一个叫 ` onPress` 的prop，然后又给 ` button` 传入了一个 ` onClick` 。

### 向事件处理程序传递参数 ###

给事件处理程序传递额外参数是很常见的，比如 ` id` 是行 ID，下面两种传参方式都可以：

` <button onClick={(e)=>this.deleteRow(id, e)}>Delete Row</button> <button onClick={this.deleteRow.bind(this, id)}>Delete Row</button> 复制代码`

上面两种传参方式是等价的。

在两种传参方式中，参数 ` e` 都代表 React 事件将要作为参数 ` id` 后面的第二个参数来传递。区别在于使用箭头函数时，参数 ` e` 要显示传递，但是使用 ` bind` 的话，事件对象以及更多的参数将会被隐式的进行传递。

### 参考 ###

[A guide to React onClick event handlers]( https://link.juejin.im?target=https%3A%2F%2Fblog.logrocket.com%2Fa-guide-to-react-onclick-event-handlers-d411943b14dd )

[事件处理]( https://link.juejin.im?target=https%3A%2F%2Fzh-hans.reactjs.org%2Fdocs%2Fhandling-events.html )

[React 事件绑定的正确姿势]( https://link.juejin.im?target=https%3A%2F%2Fwww.kawabangga.com%2Fposts%2F3369 )