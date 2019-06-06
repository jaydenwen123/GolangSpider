# React渲染过程源码分析 #

## 什么是虚拟DOM(Virtual DOM) ##

在传统的开发模式中，每次需要进行页面更新的时候都需要我们手动的更新DOM：

![](https://user-gold-cdn.xitu.io/2019/5/19/16acf64f73adc630?imageView2/0/w/1280/h/960/ignore-error/1)

在前端开发中，最应该避免的就是DOM的更新，因为DOM更新是极其耗费性能的，有过操作DOM经历的都应该知道，修改DOM的代码也非常冗长，也会导致项目代码阅读困难。在React中，把真是得DOM转换成JavaScript对象树，这就是我们说的 **虚拟DOM** ，它并不是真正的DOM，只是存有渲染真实DOM需要的属性的对象。

![](https://user-gold-cdn.xitu.io/2019/5/19/16acf6a26ae798de?imageView2/0/w/1280/h/960/ignore-error/1)

### 虚拟DOM的好处 ###

虽然虚拟DOM会提升一定得性能但是并不明显，因为每次需要更新的时候Virtual DOM需要比较两次的DOM有什么不同，然后批量更新，这也是需要资源的。

Virtual真实的好处其实是，他可以实现跨平台，我们所熟知的react-native就是基于VirtualDOM来实现的。

## Virtual DOM实现 ##

现在我们根据源码来分析一下Virtual DOM的构建过程。

**JSX和React.createElement**

在看源码之前，现在回顾一下React中创建组件的两种方式。

1.JSX

` function App () { return ( <div>Hello React</div> ); } 复制代码`

2.React.createElement

` const App = React.createElement( 'div' , null, 'Hello React' ); 复制代码`

这里多说一句其实JSX只不过是React.createElement的语法糖，在编译的时候babel会将JSX转换成为使用React.createElement的形式，因为JSX语法更加符合我们日常开发的习惯，所以我们在写React的时候更多的是使用JSX语法进行编写。

### React.createElement都做了什么 ###

下面粘贴一段React.createElement的源码来分析:

` ReactElement.createElement = function ( type , config, children) { //初始化参数 var propName; var props = {}; var key = null; var ref = null; var self = null; var source = null; if (config != null) { // 如果存在config，则提取里面的内容 if (hasValidRef(config)) { ref = config.ref; } if (hasValidKey(config)) { key = '' + config.key; } self = config.__self === undefined ? null : config.__self; source = config.__source === undefined ? null : config.__source; // 将新添加的元素更新到新的props中 for (propName in config) { if ( hasOwnProperty.call(config, propName) && !RESERVED_PROPS.hasOwnProperty(propName) ) { props[propName] = config[propName]; } } } //如果只有一个children参数，那么指直接赋值给children //否则合并处理children var childrenLength = arguments.length - 2; if (childrenLength === 1) { props.children = children; } else if (childrenLength > 1) { var childArray = Array(childrenLength); for (var i = 0; i < childrenLength; i++) { childArray[i] = arguments[i + 2]; } props.children = childArray; } // 如果某个prop为空，且存在默认的prop，则将默认的prop赋值给props if ( type && type.defaultProps) { var defaultProps = type.defaultProps; for (propName in defaultProps) { if (props[propName] === undefined) { props[propName] = defaultProps[propName]; } } } //返回一个ReactElement实例对象，这个可以理解就是我们说的虚拟DOM return ReactElement( type , key, ref, self, source , ReactCurrentOwner.current, props, ); }; 复制代码`

#### ReactElement与其中的安全机制 ####

看到这里我们不禁好奇上述代码中返回的ReactElement到底是个什么东西呢？其实ReactElement就只是我们常说的虚拟DOM，ReactElement主要包含了这个DOM节点的类型（type）、属性（props）和子节点（children）。ReactElement只是包含了DOM节点的数据，还没有注入对应的一些方法来完成React框架的功能。

现在来看一下ReactElement的源码部分 ：

` var ReactElement = function ( type , key, ref, self, source , owner, props) { var element = { // react中防止XSS注入的变量，也是标志这个是react元素的变量，稍后会讲 $ $typeof : REACT_ELEMENT_TYPE, // 构建属于这个元素的属性值 type : type , key: key, ref: ref, props: props, // 记录一下创建这个元素的组件 _owner: owner, }; return element; }; 复制代码`

上述代码可以看出来，ReactElement其实就是装有各种属性的一个大对象而已。

#### $$typeof ####

首先我们现在控制台打印一下react.createElement的结果：

![](https://user-gold-cdn.xitu.io/2019/5/21/16ad6385cd8ce65d?imageView2/0/w/1280/h/960/ignore-error/1)

WHAT？？？这个变量是什么???

其实$$typeof是为了安全问题引入的变量，什么安全问题呢？那就是 **XSS**

我们都知道React.createElement方法的第三个参数是允许用户输入自定义组件的,那么设想一下，如果前端允许用户输入下面一段代码：

` var input = "{" type ": " div ", " props ": {" dangerouslySetInnerHTML ": {" __html ": " <script>alert( 'hey' )</script> "}}}" " //然后我们开始用输入的值创建ReactElement，就变成了下面这个样子 React.createElement('div', null, input); 复制代码`

至此XSS注入就达成目的啦。

**那么$$typeof这个变量是怎么做到安全认证的呢？？？**

` var REACT_ELEMENT_TYPE = (typeof Symbol === 'function' && Symbol.for && Symbol.for( 'react.element' )) || 0xeac7; ReactElement.isValidElement = function (object) { return ( typeof object === 'object' && object !== null && object.$ $typeof === REACT_ELEMENT_TYPE ); }; 复制代码`

首先$$typeof是Symbol类型的变量，是无法通过json对象转成字符串，所以就如果只是简单的json拷贝，是没有办法通过ReactElement.isValidElement的验证的，ReactElement.isValidElement会将不带有$$typeof变量的元素全部丢掉不用。

## React的render过程 ##

现在通过源码来看一下react中从定义完组件之后render到页面的过程。

#### 1.ReactDOM.render ####

当我们想要将一个组件渲染到页面上需要调用ReactDOM.render(element,container,[callback])方法，现在我们就从这个方法入手一步一步来看源码：

` var ReactDOM = { findDOMNode: findDOMNode, render: ReactMount.render, unmountComponentAtNode: ReactMount.unmountComponentAtNode, version: ReactVersion }; 复制代码`

从上面代码我们可以看到，我们经常调用的ReactDOM.render，其实是在调用ReactMount的render方法。所以我们现在来看ReactMount中的render方法都做了些什么。

` /src/renderers/dom/client/ReactMount.js render: function (nextElement, container, callback) { return ReactMount._renderSubtreeIntoContainer( null, nextElement, container, callback, ); } 复制代码`

#### 2._renderSubtreeIntoContainer ####

现在我们终于找到了源头，那就是_renderSubtreeIntoContainer方法，我们在来看一下它是怎么样定义的，可以根据下面代码中的注释一步一步的来看：

` _renderSubtreeIntoContainer: function ( parentComponent, nextElement, container, callback, ) { // 检验传入的callback是否符合标准，如果不符合，validateCallback会throw出 //一个错误(内部调用了node_modules/fbjs/lib/invariant有invariant方法) ReactUpdateQueue.validateCallback(callback, 'ReactDOM.render' ); // 此处的TopLevelWrapper，只不过是将你传进来的 type ，进行一层包裹，并赋值ID，并会在TopLevelWrapper.render方法中返回你传入的值 // 具体看源码，，所以个这东西只是一个包裹层 var nextWrappedElement = React.createElement(TopLevelWrapper, { child: nextElement, }); //判断之前是否渲染过此元素，如果有返回此元素，如果没有返回null var prevComponent = getTopLevelWrapperInContainer(container); if (prevComponent) { var prevWrappedElement = prevComponent._currentElement; var prevElement = prevWrappedElement.props.child; // 判断是否需要更新组件 if (shouldUpdateReactComponent(prevElement, nextElement)) { var publicInst = prevComponent._renderedComponent.getPublicInstance(); var updatedCallback = callback && function () { callback.call(publicInst); }; // 如果需要更新则调用组件更新方法，直接返回更新后的组件 ReactMount._updateRootComponent( prevComponent, nextWrappedElement, nextContext, container, updatedCallback, ); return publicInst; } else { // 不需要更新组件，那就把之前的组件卸载掉 ReactMount.unmountComponentAtNode(container); } } // 返回当前容器的DOM节点，如果没有container返回null var reactRootElement = getReactRootElementInContainer(container); // 返回上面reactRootElement的data-reactid var containerHasReactMarkup =reactRootElement && !!internalGetID(reactRootElement); // 判断当前容器是不是有身为react元素的子元素 var containerHasNonRootReactChild = hasNonRootReactChild(container); // 得到是否应该重复使用的标记变量 var shouldReuseMarkup = containerHasReactMarkup && !prevComponent && !containerHasNonRootReactChild; // 将一个新的组件渲染到真是得DOM上 var component = ReactMount._renderNewRootComponent( nextWrappedElement, container, shouldReuseMarkup, nextContext, )._renderedComponent.getPublicInstance(); // 如果有callback函数那就执行这个回调函数，并且将其this只想component if (callback) { callback.call(component); } // 返回组件 return component; }, 复制代码`

根据上面的注释可以很容易理解上面的代码,现在我们总结一下_renderSubtreeIntoContainer方法的执行过程：

` 1.校验传入callback的格式是否符合规范 2.用TopLevelWrapper包裹层(带有reactID)包裹传入的type，这里说明一下，react.createElement这个方法的type值可以有三种分别是，原生标签的标签名字符串('div'、'span')、react component 、react fragment 3.判断是否渲染过此次准备渲染的元素，如果渲染过，则判断是否需要更新。 3.1 如果需要更新则调用更新方法，并且直接将更新后的组件返回 3.2 如果不需要更新，则卸载老组件 4.如果没渲染过，则处理shouldReuseMarkup变量 5.调用ReactMount._renderNewRootComponent将组将更新到DOM(此函数后面会分析) 6.返回组件 复制代码`

#### 3.ReactMount._renderNewRootComponent(渲染组件，批次装载) ####

上面说到其实在_renderSubtreeIntoContainer方法中，最后使用了ReactMount._renderNewRootComponent进行进行组件的渲染，接下来我们看一下该方法的源码：

` _renderNewRootComponent: function ( nextElement, container, shouldReuseMarkup, context, ) { // 监听window上面的滚动事件，缓存滚动变量，保证在滚动的时候页面不会触发重排 ReactBrowserEventEmitter.ensureScrollValueMonitoring(); //获取组件实例 var componentInstance = instantiateReactComponent(nextElement, false ); // 批处理，初始化render的过程是异步的，但是在render的时候componentWillMount或者componentDidMount生命中其中 // 可能会执行更新变量的操作，这是react会将这些操作通过当前批次策略，统一处理。 ReactUpdates.batchedUpdates( batchedMountComponentIntoNode, // * componentInstance, container, shouldReuseMarkup, context, ); var wrapperID = componentInstance._instance.rootID; instancesByReactRootID[wrapperID] = componentInstance; // 返回实例 return componentInstance; } 复制代码`

还是先来总结一下上面代码的过程：

` 1.监听滚动事件，缓存变量，避免滚动带来的重排 2.初始化组件实例 3.批量执行更新操作 复制代码`

##### react四大类组件 #####

在上面代码执行过程的2中调用instantiateReactComponent创建了，组件的实例，其实组件类型有四种，具体看下图：

![](https://user-gold-cdn.xitu.io/2019/5/23/16ae0aec45a7e1e4?imageView2/0/w/1280/h/960/ignore-error/1) 在这里我们还是看一下它的具体实现，然后分析一下过程：

` function instantiateReactComponent(node, shouldHaveDebugID) { var instance; if (node === null || node === false ) { // 空组件 instance = ReactEmptyComponent.create(instantiateReactComponent); } else if (typeof node === 'object' ) { var element = node; if (typeof element.type === 'string' ) { // 原生DOM instance = ReactHostComponent.createInternalComponent(element); } else if (isInternalComponentType(element.type)) { instance = new element.type(element); } else { // react组件 instance = new ReactCompositeComponentWrapper(element); } } else if (typeof node === 'string' || typeof node === 'number' ) { // 文本字符串 instance = ReactHostComponent.createInstanceForText(node); } else { } return instance; } 1.node为空时初始化空组件ReactEmptyComponent.create(instantiateReactComponent) 2.node类型是对象时，即是DOM标签或者自定义组件，那么如果element的类型是字符串，则初始化DOM标签组件ReactNativeComponent.createInternalComponent，否则初始化自定义组件ReactCompositeComponentWrapper 3.当node是字符串或者数字时，初始化文本组件ReactNativeComponent.createInstanceForText 4.其他情况不处理 复制代码`

##### 批次装载 #####

在_renderNewRootComponent代码中有一个方法后面我是打了星号的，batchedUpdate方法的第一个参数其实是个callback，这里也就是batchedMountComponentIntoNode，从方法名就可以很容易看出来他是一个批次装载组件的方法，他是定义在ReactMount上面的，来看一下他的具体实现吧。

` function batchedMountComponentIntoNode( componentInstance, container, shouldReuseMarkup, context, ) { // 在batchedMountComponentIntoNode中，使用transaction.perform调用mountComponentIntoNode让其基于事务机制进行调用 var transaction = ReactUpdates.ReactReconcileTransaction.getPooled( !shouldReuseMarkup && ReactDOMFeatureFlags.useCreateElement, ); transaction.perform( mountComponentIntoNode, null, componentInstance, container, transaction, shouldReuseMarkup, context, ); ReactUpdates.ReactReconcileTransaction.release(transaction); } 复制代码`

事务机制以后再进行分析，这里就直接来看mountComponentIntoNode是如何将组件渲染成DOM节点的吧。

#### 4.生成DOM(mountComponentIntoNode) ####

mountComponentIntoNode这个函数主要就是装载组件，并且将其插入到DOM中，话不多说，直接上源码，然后根据源码一步步的分析:

` /** * Mounts this component and inserts it into the DOM. * * @param {ReactComponent} componentInstance The instance to mount. * @param {DOMElement} container DOM element to mount into. * @param {ReactReconcileTransaction} transaction * @param {boolean} shouldReuseMarkup If true , do not insert markup */ function mountComponentIntoNode( wrapperInstance, container, transaction, shouldReuseMarkup, context, ) { var markup = ReactReconciler.mountComponent( wrapperInstance, transaction, null, ReactDOMContainerInfo(wrapperInstance, container), context, ); wrapperInstance._renderedComponent._topLevelWrapper = wrapperInstance; ReactMount._mountImageIntoNode( markup, container, wrapperInstance, shouldReuseMarkup, transaction, ); } 复制代码`

可以看到mountComponentIntoNode方法首先调用了ReactReconciler.mountComponent方法，而在ReactReconciler.mountComponent方法中其实是调用了上面四种react组件的mountComponent方法，前面的就不说了，我们直接来看一下四种组件中的mountComponent方法都干了什么吧。

` /src/renderers/dom/shared/ReactDOMComponent.js mountComponent: function ( transaction, hostParent, hostContainerInfo, context, ) { var props = this._currentElement.props; switch (this._tag) { case 'audio' : case 'form' : case 'iframe' : case 'img' : case 'link' : case 'object' : case 'source' : case 'video' : .... // 创建容器 var mountImage; var ownerDocument = hostContainerInfo._ownerDocument; var el; if (this._tag === 'script' ) { var div = ownerDocument.createElement( 'div' ); var type = this._currentElement.type; div.innerHTML = `< ${type} ></ ${type} >`; el = div.removeChild(div.firstChild); } else if (props.is) { el = ownerDocument.createElement(this._currentElement.type, props.is); } else { el = ownerDocument.createElement(this._currentElement.type); } } // 更新props，第一个参数是上次的props，第二个参数是最新的props，如果上一次的props为空那么就是新建状态 this._updateDOMProperties(null, props, transaction); // 生成DOMLazyTree对象 var lazyTree = DOMLazyTree(el); // 处理孩子节点 this._createInitialChildren(transaction, props, context, lazyTree); mountImage = lazyTree; // 返回容器 return mountImage; } 复制代码`

总结一下上述代码的执行过程，在这里我只截取了初次渲染时候执行的代码： 1.对特殊的标签进行处理，并且调用方法给出相应警告 2.创建DOM节点 3.调用_updateDOMProperties方法来处理props 4.生成DOMLazyTree 5.通过DOMLazyTree调用_createInitialChildren处理孩子节点。然后返回DOM节点

下面我们来看一下这个DOMLazyTree方法都干了些什么，还是上源码：

` function queueChild(parentTree, childTree) { if ( enable Lazy) { parentTree.children.push(childTree); } else { parentTree.node.appendChild(childTree.node); } } function queueHTML(tree, html) { if ( enable Lazy) { tree.html = html; } else { set InnerHTML(tree.node, html); } } function queueText(tree, text) { if ( enable Lazy) { tree.text = text; } else { set TextContent(tree.node, text); } } function toString () { return this.node.nodeName; } function DOMLazyTree(node) { return { node: node, children: [], html: null, text: null, toString, }; } DOMLazyTree.queueChild = queueChild; DOMLazyTree.queueHTML = queueHTML; DOMLazyTree.queueText = queueText; 复制代码`

从上述代码可以看到DOMLazyTree其实就是一个用来包裹节点信息的对象，里面有孩子节点，html节点，文本节点，并且提供了将这些节点插入到真是DOM中的方法，现在我们来看一下在_createInitialChildren方法中它是如何来使用这个lazyTree对象的：

` _createInitialChildren: function (transaction, props, context, lazyTree) { var innerHTML = props.dangerouslySetInnerHTML; if (innerHTML != null) { if (innerHTML.__html != null) { DOMLazyTree.queueHTML(lazyTree, innerHTML.__html); } } else { var contentToUse = CONTENT_TYPES[typeof props.children] ? props.children : null; var childrenToUse = contentToUse != null ? null : props.children; if (contentToUse != null) { if (contentToUse !== '' ) { DOMLazyTree.queueText(lazyTree, contentToUse); } } else if (childrenToUse != null) { var mountImages = this.mountChildren( childrenToUse, transaction, context, ); for (var i = 0; i < mountImages.length; i++) { DOMLazyTree.queueChild(lazyTree, mountImages[i]); } } } } 复制代码`

判断当前节点的dangerouslySetInnerHTML属性、孩子节点是否为文本和其他节点分别调用DOMLazyTree的queueHTML、queueText、queueChild.

##### ReactCompositeComponent #####

在实例调用mountComponent时，在这里额外的说一下这个函数的执行过程，ReactCompositeComponent也就是我们说的react自定义组件，起主要的执行过程如下：

` 1.处理props、contex等变量，调用构造函数创建组件实例 2.判断是否为无状态组件，处理state 3.调用performInitialMount生命周期，处理子节点，获取markup。 4.调用componentDidMount生命周期 复制代码`

在performInitialMount函数中，首先调用了componentWillMount生命周期，由于自定义的React组件并不是一个真实的DOM，所以在函数中又调用了孩子节点的mountComponent。这也是一个递归的过程，当所有孩子节点渲染完成后，返回markup并调用componentDidMount.

#### 5.渲染DOM(_mountImageIntoNode) ####

在上述mountComponentIntoNode中最后一步是执行_mountImageIntoNode方法，在该方法中核心的渲染方法就是insertTreeBefore，我们直接来看这个方法的源码，然后进行分析：

` var insertTreeBefore = function ( parentNode, tree, referenceNode, ) { if ( tree.node.nodeType === DOCUMENT_FRAGMENT_NODE_TYPE || (tree.node.nodeType === ELEMENT_NODE_TYPE && tree.node.nodeName.toLowerCase() === 'object' && (tree.node.namespaceURI == null || tree.node.namespaceURI === DOMNamespaces.html)) ) { insertTreeChildren(tree); parentNode.insertBefore(tree.node, referenceNode); } else { parentNode.insertBefore(tree.node, referenceNode); insertTreeChildren(tree); } } function insertTreeChildren(tree) { if (! enable Lazy) { return ; } var node = tree.node; var children = tree.children; if (children.length) { for (var i = 0; i < children.length; i++) { insertTreeBefore(node, children[i], null); } } else if (tree.html != null) { set InnerHTML(node, tree.html); } else if (tree.text != null) { set TextContent(node, tree.text); } } 复制代码`

1.该方法首先就是判断当前节点是不是fragment节点或者Object插件 2.如果满足条件1，首先调用insertTreeChildren将此节点的孩子节点渲染到当前节点上，再将渲染完的节点插入到html 3.如果不满足1，是其他节点，先将节点插入到插入到html，再调用insertTreeChildren将孩子节点插入到html

在此过程中已经一次调用了setInnerHTML或setTextContent来分别渲染html节点和文本节点。

### 结尾 ###

上述文章就是react的初次渲染过程分析，如果有哪些地方写的不对，欢迎在评论中讨论。本文代码采用的react15中的代码，和react最新版本代码会有一些的出入。