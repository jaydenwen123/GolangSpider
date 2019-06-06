# Element(React)源码分析系列4--Radio组件 #

## 前言 ##

学习就好比是座大山，人们沿着不同的路登山，分享着自己看到的风景。你不一定能看到别人看到的风景，体会到别人的心情。只有自己去登山，才能看到不一样的风景，体会更加深刻。一千个读者就有一千个哈姆雷特，但是莎士比亚心中的哈姆雷特肯定只有一个。就好比element源码只有一个，但每个人看到的都是不一样的风景。element源码解读是一个系列，每一个组件细腻的背后都能看到前端工程师付出的心血，本篇带来的是Element源码分析系列4-Radio(单选框)

## 简介 ##

单选框这个组件看似简单，实则知识点众多，较为复杂，如果写一个html的原生单选框，那确实很简单，但是封装一个完整的单选组件就不那么简单了。element团队在整个radio组件的设计和构思真可谓十分的细致，没有一行多余的代码。从基础的css到选择逻辑都无不彰显其巧夺天工的思想。

## 划重点 ##

### 1、如何隐藏原始radio标签默认样式 ###

我们都知道原生的radio标签很丑，样式在各个浏览器不统一，作为一个组件不可能就这样屈服，况且人这种物种总是希望统驾驭所有的物质之上，包括代码。所以必须自己实现所有radio按钮的样式，那么如何自己实现一个可控制的radio标签呢？看element是怎么实现的。

![radio](https://user-gold-cdn.xitu.io/2019/6/6/16b2b50567e6a8a3?imageView2/0/w/1280/h/960/ignore-error/1) 看图不难分析出radio组件的html样式,整个大盒子被一个laber标签包裹着，laber里分为两个span,第一个span是用来展示选择按钮的，第二个span是来描述选项的。第一个span里应该包括input(用来做radio)和span(隐藏真正的input)两个标签。其相关的html结构如下：

` <label className= 'el-radio' > <span> <span className= "el-radio__inner" ></span> <input type = "radio" className= "el-radio__original" /> </span> <span className= "el-radio__label" > {children || value} </span> </label> 复制代码`

问题一：如何做到隐藏原始radio标签默认样式？作者巧用opacity:0，真正的input透明度为0，且是绝对定位脱离文档流，因此不占空间且我们看不到，注意不是display:none或者visibility:hidden,如果是none或者hidden的话则无法触发鼠标点击了，所以只有opacity:0才能达到目的.

`.el-radio__original { opacity: 0; outline: 0; position: absolute; z-index: -1; top: 0; left: 0; right: 0; bottom: 0; margin: 0; } 复制代码`

### 2、如何点击选择按钮展示不同的状态 ###

如何做到点击旋钮显示不同的状态。通过radio标签可分析出，点击即可以理解为一个事件(这里用onChange表示)。通过点击事件传递一个props值(checked)显示不同的状态。所以在input上一定会有一个事件和一个所要传递的状态。

` <input type = "radio" className= "el-radio__original" checked={checked} // 传递的状态 onChange={this.onChange.bind(this)} // 事件 /> 复制代码`

那么接下来就是定义事件了，其核心代码为：

` constructor(props) { super(props); this.state = { checked: this.getChecked(props) // 定义一个state，受props影响 }; } componentWillReceiveProps(props) { const checked = this.getChecked(props); if (this.state.checked !== checked) { // 维持每次点击时只选中一个 this.setState({ checked }); } } getChecked(props) { return Boolean(props.checked) // 根据传递的props，输出 true / false } onChange(e) { const checked = e.target.checked; // 定义被点击项的checked为 true if (checked) { if (this.props.onChange) { this.props.onChange(this.props.value); // 向外暴露一个onChange事件并携带value值 } } this.setState({ checked }); // 更新被点击项的state为 true } 复制代码`

如果你对上述描述的还是雨里雾里，那么下面这张图可以更好的帮组你理解element-radio标签事件的逻辑：

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2b7684a36dcb5?imageView2/0/w/1280/h/960/ignore-error/1) 从图中我们可以看出，radio是如何做到点击改变状态，并且又如何维护每次点击完后radio的状态。

* 传入checked作为props决定每个radio的状态
* 根据传入的props,定义一个内部的state,跟props的checked保持一致
* 点击事件改变当前的radio的状态，更新state中的checked,并且向外传递一个value
* 外面接受到传来的value，重新定义props
* 通过生命周期函数更新内部的state，让其保持每个radio一致

### 3、Radio.Group如何做到单选框组 ###

单选框组故名思议是将所有的radio包裹一层，由最外层原始事件来决定每个radio的状态。故需要用到React的两个API：React.Children.map、React.cloneElement

技术扩展：

* React.Children 提供了用于处理 this.props.children 不透明数据结构的实用方法。

> 
> 
> 
> React.Children.map(children, function[(thisArg)])
> 
> 

在 children 里的每个直接子节点上调用一个函数，并将 this 设置为 thisArg。如果 children 是一个数组，它将被遍历并为数组中的每个子节点调用该函数。如果子节点为 null 或是 undefined，则此方法将返回 null 或是 undefined，而不会返回数组。

* React.cloneElement: 以 element 元素为样板克隆并返回新的 React 元素。返回元素的 props 是将新的 props 与原始元素的 props 浅层合并后的结果。

> 
> 
> 
> React.cloneElement(element, [props], [...children])
> 
> 

先了解这两个API，在来看看element中Radio.Group的源码就很简单来。

` <div ref= 'RadioGroup' className= 'el-radio-group' > { React.Children.map(this.props.children, element => { if (!element) { return null } return React.cloneElement(element, Object.assign({}, element.props, { onChange: this.onChange.bind(this), value: this.props.value, })) }) } </div> 复制代码`

通过遍历所有的radio组件，克隆出相应的radio组件，并为其附上radio的一些属性，包括方法和value。

## 结语 ##

纵观整个radio的设计和实现，每个设计的过程都十分的微妙，尤其是radio如何更新状态，在平时业务中，我们也经常会遇到这种需求，但是真正的实现起来却没有这么简洁。源码的学习或多或少可以让我们增强自己的代码能力和业务能力。那么跟随我的脚步，带你一步一步解析整个element源码的奇思妙想。山不再高，有仙则灵；水不在深，有心则成。后续将推出更多源码解析文章。源码请👇这里 [element-radio源码]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2FWsmDyj%2Fr-react%2Ftree%2Fmaster%2Felement-react%2Fsrc%2Fcomponents%2Fradio )