# 前端战五渣学React——JSX & React.createElement() & React.ReactElement()源码 #

![](https://user-gold-cdn.xitu.io/2019/6/5/16b28109fb26c85b?imageView2/0/w/1280/h/960/ignore-error/1)

> 
> 
> 
> 最近《一拳超人》动画更新第二季了，感觉打斗场面没有第一季那么烧钱了，但是剧情还挺好看的，就找了漫画来看。琦玉老师真的厉害！！！打谁都一拳，就喜欢看老师一拳把那些上来就吹牛逼的反派打的稀烂，专治各种不服！！！
> 
> 
> 

### 招聘AD ###

阿里巴巴集团核心前端岗位

薪资25到50

一年一般至少16个月工资

有意者微信联系：Dell-JS

# 正文 #

## 三大民工框架 ##

说到现在的前端，各种招聘JD上都会写

**“对主流框架（React/Vue/Angular）有了解，至少深入了解一种”**

或者是

**“精通MV*框架（React/Vue/Angular），至少熟练使用一种，有大型项目经验”**

从中我们可以看出现在前端在工作中使用的框架几乎形成了三足鼎立之势，形如当初的“三大民工漫画”——《海贼王》、《火影忍者》以及《死神》，而Angular又类似《死神》一样，国内人气低迷（我只是从招聘信息来看的。。。angular布道者勿喷）。而React凭借自己的灵活性和vue凭借简单好上手的优势，平分秋色。这回就来主要讲一讲React的一大核心概念——JSX，以及对应的 ` React.createElement()` 这个方法的源码阅读。

本文阅读需要具备以下知识储备：

* JavaScript基本语法，用js开发过项目最好
* 最好使用过react，没用过的😅可能。。。

## JSX ##

(了解的可以直接跳到下一节看React.createElement()的源码)

话不多说，让我们来实现一个功能：

**创建一个div标签，class名为“title”，内容为“你好 前端战五渣”**

看下面的代码⬇️

` <!DOCTYPE html> < html lang = "en" > < head > < meta charset = "UTF-8" > <!-- 引入react核心代码 --> < script src = "https://cdn.bootcss.com/react/16.8.6/umd/react.development.js" > </ script > <!-- 引入reactDom核心代码 --> < script src = "https://cdn.bootcss.com/react-dom/16.8.6/umd/react-dom.development.js" > </ script > <!-- 引入babel核心代码 --> < script src = "https://cdn.bootcss.com/babel-standalone/6.26.0/babel.min.js" > </ script > < title > JSX & React.createElement() </ title > </ head > < body > <!-- 使用javascript原生插入节点的根节点 --> < div id = "rootByJs" > </ div > <!-- 使用React.createElement()方法插入节点的根节点 --> < div id = "rootByReactCreateElement" > </ div > <!-- 使用JSX方法插入节点的根节点 --> < div id = "rootByJsx" > </ div > < script > // 原生方法插入 let htmlNode = document.createElement( 'div' ); htmlNode.innerHTML = '你好 前端战五渣' ; htmlNode.className = 'title' ; document.getElementById( 'rootByJs' ).appendChild(htmlNode); </ script > < script > // 使用React.createElement()方法插入 ReactDOM.render( React.createElement( 'div' , { className : "title" }, '你好 前端战五渣' ), document.getElementById( 'rootByReactCreateElement' ) ); </ script > < script type = "text/babel" > // 使用JSX方法插入 ReactDOM.render( < div className = "title" > 你好 前端战五渣 </ div > , document.getElementById( 'rootByJsx' ) ); </ script > </ body > </ html > 复制代码`

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2aa01692a3b85?imageView2/0/w/1280/h/960/ignore-error/1)

上面实现这个功能，用了三种方法，一种是js原生方法，一种是用react提供的createElement方法，还有最后一种使用JSX来实现的。

### 什么是JSX ###

其实jsx就是react这个框架提出的一种语法扩展，在react建议使用jsx，因为jsx可以清晰明了的描述DOM结构。可能到这里我们可能有人会说，这跟模板语言有什么区别呢？template也可以实现啊，但是JSX具有JavaScript的全部功能（官网这么说的🤦‍♀️）

一句话总结：JSX语法就是JavaScript和html可以混着写，灵活的一笔

JSX的优点呢？

* 可以在js中写更加语义化且简单易懂的标签
* 更加简洁
* 结合原生js的语法

（也有人说jsx写起来很乱，仁者见仁智者见智吧）

### JSX和React.createElement()的关系 ###

那我们知道了JSX是什么，可是这跟我们这回要说的 ` React.createElement()` 方法有什么关系呢？先来回顾一个面试会问的问题“你能说说vue和react有什么区别吗”，有一个区别就是在使用webpack打包的过程中，vue是用vue-loader来处理 `.vue` 后缀的文件，而react在打包的时候，是通过babel来转换的，因为react的组件说白了还是 `.js` 或者 `.jsx` ，是扩展的js语法，所以是通过babel转换成浏览器识别的es5或者其他版本的js

那我们来看看jsx的语法通过babel转换会变成什么样⬇️

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2adce6f24e05a?imageView2/0/w/1280/h/960/ignore-error/1)

我们可以看到通过babel转换以后，我们的JSX语法中的标签会被转换成一个 ` React.createElement()` 并传入对应的参数

` ReactDOM.render( < div className = "title" > hello gedesiwen </ div > , document.getElementById( 'rootByJsx' ) ); 复制代码`

变~

` ReactDOM.render( React.createElement( 'div' , { className : 'title' }, 'hello gedesiwen' ), document.getElementById( 'rootByJsx' ) ); 复制代码`

这我们看见了jsx变成了 ` React.createElement()`

### 多个子节点 ###

上面的代码中，我们只是有一个子节点，就是文本节点“你好 前端战五渣”，那如果我们有很多个呢

我们在React组件中代码是这样的⬇️

` import DragonBall from './dragonBall' ; let htmlNode = ( < Fragment > < DragonBall name = "孙悟空" /> < div className = "hello" key = {1} > hello </ div > < div className = "world" key = {2} > world </ div > </ Fragment > ) ReactDOM.render( htmlNode, document.getElementById('rootByJsx') ); 复制代码`

我们的节点中包括DragonBall组件，还有Fragment，并且还有两个div

Fragment是干什么的呢？？？这就是JSX语法的一个规则，我们只能有一个根节点，如果我们有两个并列的div，但是直接写并列的两个div会报错，我们就只能在外面套一层div，但是我们不想创建不用的标签，这时候我们就能使用Fragment，他不会被渲染出来

> 
> 
> 
> React 中的一个常见模式是一个组件返回多个元素。Fragments 允许你将子列表分组，而无需向 DOM 添加额外节点。 ————react文档
> 
> 
> 

那上面这段我们通过babel会转换成这样⬇️

` var htmlNode = React.createElement( Fragment, null , React.createElement(_dragonBall.default, { name : "saiyajin" }), React.createElement( "div" , { className : "hello" , key : 1 }, "hello" ), React.createElement( "div" , { className : "world" , key : 2 }, "world" ) ); ReactDOM.render(htmlNode, document.getElementById( 'rootByJsx' )); 复制代码`

这就是我们转换完的js，那我们的 ` React.createElement()` 方法到底做了什么呢

# React.createElement()源码 #

首先我们需要从github上把 [react的源码,v16.8.6]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Ffacebook%2Freact ) 拉下来

然后我们找到在文件 ` /packages/react/src/ReactElement.js` 这个文件中就有我们需要的 ` React.createElement()` 方法

**（代码中左右判断 ` __DEV__` 的代码，不做考虑）**

先上完整的方法代码，伴有注释

` /** * React的创建元素方法 * @param type 标签名字符串(如’div‘或'span')，也可以是React组件类型，或是React fragment类型 * @param config 包含元素各个属性键值对的对象 * @param children 包含元素的子节点或者子元素 */ function createElement ( type, config, children ) { let propName; // 声明一个变量，储存后面循环需要用到的元素属性 const props = {}; // 储存元素属性的键值对集合 let key = null ; // 储存元素的key值 let ref = null ; // 储存元素的ref属性 let self = null ; // 下面文章介绍 let source = null ; // 下面文章介绍 if (config != null ) { // 判断config是否为空，看看是不是没有属性 // hasValidRef()这个方法就是判断config有没有ref属性，有的话就赋值给之前定义好的ref变量 if (hasValidRef(config)) { ref = config.ref; } // hasValidKey()这个方法就是判断config有没有key属性，有的话就赋值给之前定义好的key变量 if (hasValidKey(config)) { key = '' + config.key; // key值看来还给转成了字符串😳 } // __self和__source下面文章做介绍，实际也没搞明白是干嘛的 self = config.__self === undefined ? null : config.__self; source = config.__source === undefined ? null : config.__source; // 现在就是要把config里面的属性都一个一个挪到props这个之前声明好的对象里面 for (propName in config) { if ( // 判断某个config的属性是不是原型上的 hasOwnProperty.call(config, propName) && // 这行判断是不是原型链上的属性 !RESERVED_PROPS.hasOwnProperty(propName) // 不能是原型链上的属性，也不能是key，ref，__self以及__source ) { props[propName] = config[propName]; // 乾坤大挪移，把config上的属性一个一个转到props里面 } } } // 处理除了type和config属性剩下的其他参数 const childrenLength = arguments.length - 2 ; // 抛去type和config，剩下的参数个数 if (childrenLength === 1 ) { // 如果抛去type和config，就只剩下一个参数，就直接把这个参数的值赋给props.children props.children = children; // 一个参数的情况一般是只有一个文本节点 } else if (childrenLength > 1 ) { // 如果不是一个呢？？ const childArray = Array (childrenLength); // 声明一个有剩下参数个数的数组 for ( let i = 0 ; i < childrenLength; i++) { // 然后遍历，把每个参数赋值到上面声明的数组里 childArray[i] = arguments [i + 2 ]; } props.children = childArray; // 最后把这个数组赋值给props.children } // 所以props.children要不是一个字符串，要不就是一个数组 // 如果有type并且type有defaultProps属性就执行下面这段 // 那defaultProps属性是啥呢？？ // 如果传进来的是一个组件，而不是div或者span这种标签，可能就会有props，从父组件传进来的值如果没有的默认值 if (type && type.defaultProps) { const defaultProps = type.defaultProps; for (propName in defaultProps) { // 遍历，然后也放到props里面 if (props[propName] === undefined ) { props[propName] = defaultProps[propName]; } } } // 所以props里面存的是config的属性值，然后还有children的属性，存的是字符串或者数组，还有一部分defaultProps的属性 // 然后返回一个调用ReactElement执行方法，并传入刚才处理过的参数 return ReactElement( type, key, ref, self, source, ReactCurrentOwner.current, props, ); } 复制代码`

` React.createElement()` 方法的代码加注释就是上面这个，小伙伴们应该都能看懂了吧，只是其中其中还有 ` __self` 、 ` __source` 以及 ` type.defaultProps` 没有讲清楚，那我们下面会讲到，我们可以先来看看这个最后返回的 ` ReactElement()` 方法

# ReactElement()源码 #

这个方法很简单，就是添加一个判断为react元素类型的值，然后返回，

` /** * @param {*} type * @param {*} props * @param {*} key * @param {string|object} ref * @param {*} owner * @param {*} self A *temporary* helper to detect places where `this` is * different from the `owner` when React.createElement is called, so that we * can warn. We want to get rid of owner and replace string `ref`s with arrow * functions, and as long as `this` and owner are the same, there will be no * change in behavior. * * 这虽然说了用于判断this指向的，但是。。。。。方法里面也没有用到，不知道是干嘛的😳😳😳😳 * * @param {*} source An annotation object (added by a transpiler or otherwise) * indicating filename, line number, and/or other information. * * 这个参数一样。。。。也没有用到啊。。。那我传进来是干嘛的，什么注释对象。。😳😳😳搞不懂 * * @internal */ const ReactElement = function ( type, key, ref, self, source, owner, props ) { const element = { $$typeof : REACT_ELEMENT_TYPE, // 声明一下是react的元素类型 type: type, key : key, ref : ref, props : props, _owner : owner, }; return element; }; 复制代码`

## __self和__source ##

刚看到React.createElement()方法里面就用到了 ` __self` 和 ` __source` 两个属性，当时还去查了一下react的文档

文档中也没有说是干嘛用的，然后查了一下issues

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2bebf0e4bedb7?imageView2/0/w/1280/h/960/ignore-error/1)

发现是这哥们提交的commit，好😳😳😳他说 ` _self` 是用来判断this和owner是不是同一个指向巴拉巴拉的，他还说 ` __source` 是什么注释对象，我也没看懂是干嘛的。。。。然后继续往下看，看到 ` React.createElement()` 方法返回 ` ReactElement()` 方法，并且把这些都传进去了。。。。

ReactElement源码中竟然没有用这两个参数

大哥你开心就好🤡🤡🤡

看到这篇文章的大佬有知道是干嘛的可以告诉我。。。。我反正现在是懵逼的😶😶😶

## type.defaultProps ##

这个是什么呢，我们来看一段代码吧

` import React, { Component } from 'react' ; import ReactDom from 'react-dom' ; class DragonBall extends Component { render() { return ( < div > {this.props.name} </ div > ) } } ReactDom.render( < DragonBall /> , document.getElementById('root')) 复制代码`

如果我这个DragonBall组件需要展示从props传过来，如果我们没传呢，就会是 ` undefined` ，就什么都不显示，如果我们想设置默认值呢，可以这么写⬇️

` import React, { Component } from 'react' ; import ReactDom from 'react-dom' ; class DragonBall extends Component { render() { return ( < div > {this.props.name || '戈德斯文'} </ div > ) } } ReactDom.render( < DragonBall /> , document.getElementById('root')) 复制代码`

就是像上面这样写，这样我们就进行了一次判断，如果 ` props.name` 如果没有的话，就显示后面的“戈德斯文”，那还有没有什么别的办法呢？？

想也知道啊，肯定就是我们说的这个 ` defaultProps` ，这个东西怎么用呢⬇️

` import React, { Component } from 'react' ; import ReactDom from 'react-dom' ; class DragonBall extends Component { render() { return ( < div > {this.props.name} </ div > ) } } DragonBall.defaultProps = { name : '戈德斯文' } ReactDom.render( < DragonBall /> , document.getElementById('root')) 复制代码`

我们只需要这样设置就可以，如果我们页面中很多地方需要用到 ` props` 传进来的值，就不需要每个用到 ` props` 值的地方都进行一次判断了

所以，在React.createElement()源码中

` if (type && type.defaultProps) { const defaultProps = type.defaultProps; for (propName in defaultProps) { // 遍历，然后也放到props里面 if (props[propName] === undefined ) { props[propName] = defaultProps[propName]; } } } 复制代码`

这段代码就是把默认的props重新赋值。

# 回到开始 #

经过 ` React.createElement()` 方法处理，并且经过 ` ReactElement()` 方法洗礼，我们最开始的

` let htmlNode = React.createElement( Fragment, null , React.createElement(_dragonBall.default, null ), React.createElement( "div" , null , "hello" ), React.createElement( "div" , null , "world" ) ); ReactDOM.render(htmlNode, document.getElementById( 'rootByJsx' )); 复制代码`

最后到底是变成什么样的呢？

` { "key" : null , "ref" : null , "props" : { "children" : [{ "key" : null , "ref" : null , "props" : { "name" : "saiyajin" }, "_owner" : null , "_store" : {} }, { "type" : "div" , "key" : "1" , "ref" : null , "props" : { "className" : "hello" , "children" : "hello" }, "_owner" : null , "_store" : {} }, { "type" : "div" , "key" : "2" , "ref" : null , "props" : { "className" : "world" , "children" : "world" }, "_owner" : null , "_store" : {} }] }, "_owner" : null , "_store" : {} } 复制代码`

然后再经过 ` ReactDom.render()` 方法渲染到页面上

#### ps：端午节快乐~~回家过节喽 ####

我是前端战五渣，一个前端界的小学生。

# 参考 #

* [《剖析 React 源码：先热个身》]( https://juejin.im/post/5cbae9a8e51d456e2809fba3 )
* [React:issues adding __self and __source special props]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Ffacebook%2Freact%2Fpull%2F4596 )
* [《Understanding React Default Props》]( https://link.juejin.im?target=https%3A%2F%2Fblog.bitsrc.io%2Funderstanding-react-default-props-5c50401ed37d )