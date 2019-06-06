# Vue底层架构及其应用 #

> 
> 
> 
> 阅读时间大约16~22min
> 作者：汪汪
> 个人主页： [www.zhihu.com/people/wang…](
> https://link.juejin.im?target=https%3A%2F%2Fwww.zhihu.com%2Fpeople%2Fwang-wang-69-90-11%2Factivities
> )
> 
> 

# 一、前言 #

市面上有很多基于vue的core和compile做出的优化开源框架，为非Web场景引入了Vue的能力，因此学习成本低，受到广大开发者的欢迎，下面大体列一下我所了解到的，有更优秀的欢迎大家评论指出

+--------------------+---------+
|        分类        |  技术   |
+--------------------+---------+
| 跨平台native       | weex    |
| 小程序             | mpvue   |
| 服务端渲染         | Vue SSR |
| 小程序多端统一框架 | uni-app |
+--------------------+---------+

至于提供类Vue开发体验的框架就数不胜数了，如小程序框架--wepy，

从其他的方面看，github日榜，Vue每天都有过100的star，足见其火热程度，这也是为什么大家都争先恐后的在非web领域提供Vue的支持。那么Vue的底层架构及其应用就尤为重要了

# 二、Vue底层架构 #

了解Vue的底层架构，是为非web领域提供Vue能力的大前提。Vue核心分为三大块：core，compiler，platform，下面分别介绍其原理及带来的能力。

## 1、core ##

core是Vue的灵魂所在，正是core实现了通过vnode方式，递归生成指定平台视图并在数据变动时，自动diff更新视图，也正是因为VNode机制，使得core是平台无关的，就算core的功能在于UI渲染。

![](https://user-gold-cdn.xitu.io/2019/5/30/16b076587cd2726e?imageView2/0/w/1280/h/960/ignore-error/1)

我将从如下几个方面来说明core

* 挂载
* 指令
* Vnode----划重点
* 组件实例vm及vm间的关系
* nextTick
* Watcher----划重点
* vnode diff算法----划重点
* core总结

### 1.1 挂载 ###

将vnode生成的具体平台元素append到已知节点上。我们拿web平台举例，用vnode通过document.createElement生成dom，然后在append到文档树中某个节点上。后面我们也会经常说到挂载组件，它指的就是执行组件对应render生成vnode，然后遍历vnode生成具体平台元素，组件的根节点元素会被append到父元素上。

### 1.2 指令 ###

指令在Vue中是具有特定含义的属性，指令分两类，一类是编译时处理，在生成的render函数上体现，如：v-if，v-for，另外一类是运行时使用，更多的是对生成的具体平台元素操作，web平台的话就是对dom的操作

### 1.3 VNode---------划重点 ###

vnode是虚拟node节点，是具体平台元素对象的进一步抽象（简化版），每一个平台元素对应一个vnode，可通过vnode结构完整还原具体平台元素结构。 下面以web平台来解释vnode。对于web，假定有如下结构：

` < div class = "box" @ click = "onClick" > ------------------对应一个vnode < p class = "content" > 哈哈 </ p > -------对应一个vnode < TestComps > </ TestComps > ----------自定义组件同样对应一个vnode < div > </ div > -----------------------对应一个vnode </ div > 复制代码`

经过Vue的compile模块将生成渲染函数，执行这个渲染函数就会生成对应的vnode结构：

` //这里我只列出关键的vnode信息 { tag : 'div' , data :{ attr :{}, staticClass : 'box' , on :{ click :onClick}}, children :[{ tag : 'p' , data :{ attr :{}, staticClass : 'content' , on :{}}, children :[{ tag : '' , data :{}, text : '哈哈' }] },{ tag : 'div' , data :{ attr :{}, on :{}}, },{ tag : 'TestComps' , data :{ attr :{}, hook :{ init :fn, prepatch :fn, insert :fn, destroy :fn } }, }] } 复制代码`

最外层的div对应一个vnode，包含三个孩子vnode，注意自定义组件也对应一个vnode，不过这个vnode上挂着组件实例

### 1.4 组件实例vm及vm间的关系---------划重点 ###

组件实例其实就是Vue实例对象，只有自定义组件才会有，平台相关元素是没有的，要看懂Vue的core，明白下面这个关系很重要。现在，让我们来直观感受下：

假定有如下结构的模板，元素上的vnode表示生成的对应vnode名称：

` // new Vue的template，对应的实例记为vm1 <div vnode1> <p vnode2></p> <TestComps vnode3 test Attr= "hahha" @click= "clicked" :username= "username" :password= "password" ></TestComps> </div> 复制代码` ` // TestComps的template，对应的实例记为vm2 <div vnode4> <span vnode5></span> <p vnode6></p> </div> 复制代码` ` // 生成的vnode关系树为 vnode1={ tag: 'div' , children:[vnode2,vnode3] } vnode3={ tag: 'TestComps' , children:undefined, parent:undefined } vnode4={ tag: 'div' , children:[vnode5,vnode6], parent:vnode3 //这一点关系很重要 } 复制代码` ` // 生成的vm关系树为 vm1={ $data :{password: "123456" ,username: "aliarmo" }, //组件对应state $props :{} //使用组件时候传下来到模板里面的数据 $attrs :{}, $children :[vm2], $listeners :{} $options : { components: {} parent: undefined //父组件实例 propsData: undefined //使用组件时候传下来到模板里面的数据 _parentVnode: undefined } $parent :undefiend //当前组件的父组件实例 $refs :{} //当前组件里面包含的dom引用 $root :vm1 //根组件实例 $vnode :undefined //组件被引用时候的那个vnode，比如<TestComps></TestComps> _vnode:vnode1 //当前组件模板根元素所对应的vnode对象 } vm2={ $data :{} //组件对应state $props :{password: "123456" ,username: "aliarmo" } //使用组件时候传下来到模板里面的数据 $attrs :{ test Attr: 'hahha' }, $children :[], $listeners :{click:fn} $options : { components: {} parent: vm1 //父组件实例 propsData: {password: "123456" ,username: "aliarmo" } //使用组件时候传下来到模板里面的数据 _parentVnode: vnode3 } $parent :vm1 //当前组件的父组件实例 $refs :{} //当前组件里面包含的dom引用 $root :vm1 //根组件实例 $vnode :vnode3 //组件被引用时候的那个vnode，比如<TestComps></TestComps> _vnode:vnode4 //当前组件模板根元素所对应的vnode对象 } 复制代码`

### 1.5 nextTick ###

它可以让我们在下一个事件循环做一些操作，而非在本次循环，用于异步更新，原理在于microtask和macrotask

让我们来看段代码：

` new Promise ( resolve => { return 123 }).then( data => { console.log( 'step2' ,data) }) console.log( 'step1' ) 复制代码`

结果是先输出 step1，然后在step2，resolve的promise是一个microtask，同步代码是macrotask

` // 在Vue中 this.username= 'aliarmo' // 可以触发更新 this.pwd= '123' // 同样可以触发更新 复制代码`

那同时改变两个state，是否会触发两次更新呢，并不会，因为this.username触发更新的回调会被放入一个通过Promise或者MessageChannel实现的microtask中，亦或是setTimeout实现的macrotask，总之到了下一个事件循环。

### 1.6 Watcher---------划重点 ###

一个组件对应一个watcher，在挂载组件的时候创建这个观察者，组件的state，包含data，props都是被观察者，被观察者的任何变化会被通知到观察者，被观察者的变动导致观察者执行的动作是 ` vm._update(vm._render(), hydrating)` ,组件重新render生成vnode并patch。

**明白这个关系很重要** ：观察者包含对变动做出响应的定义，一个组件对应一个观察者对应组件里面的所有被观察者，被观察者可能被用于其他组件，那么一个被观察者会对应多个观察者，当被观察者发生变动时，通知到所有观察者做出更新响应。

![](https://user-gold-cdn.xitu.io/2019/5/30/16b077503e0013de?imageView2/0/w/1280/h/960/ignore-error/1)

组件A的state1发生了变化，那会导致观察了这个state1的watcher收到变动通知，会导致组件A重新渲染生成新的vnode，在组件A新vnode和老的vnode patch的过程中，会updateChildrenComponent，也就是导致子组件B的props被重新设置一个新值，因为子组件B是有观察传入的state1的，因此会通知到相应watcher，导致子组件B的更新

**整个watcher体系的建立过程** ：

* 创建组件实例的时候会对data和props进行observer，
* 对传入的props进行浅遍历，重新设定属性的属性描述符get和set，如果props的某个属性值为对象，那么这个对象在父组件是被深度observe过的，所以props是浅遍历
* observer会深度遍历data，对data所包含属性重新定义，即defineReactive，重新设定属性描述符的get和set
* 在mountComponent的时候，会new Wacther，当前watcher实例会被pushTarget，设定为目标watcher，然后执行 ` vm._update(vm._render(), hydrating)` ，执行render函数导致属性的get函数被调用，每个属性会对应一个dep实例，在这个时候，dep实例关联到组件对应的watcher，实现依赖收集，关联后popTarget。
* 如果有子组件，会导致子组件的实例化，重新执行上述步骤

**state变动响应过程** ：

* 当state变动后，调用属性描述符的set函数，dep会通知到关联的watcher进入到nextTick任务里面，这个watcher实例的run函数包含 ` vm._update(vm._render(), hydrating)` ，执行这个run函数，导致重新生成vnode，进行patch，经过diff，达到更新UI目的

**父组件state变化如何导致子组件也发生变化？**

父组件state更新后，会导致渲染函数重新执行，生成新的vnode，在oldVnode和newVnode patch的过程中，如果遇到的是组件vnode，会updateChildrenComponent，这里面做的操作就是更新子组件的props，因为子组件是有监听props属性的变动的，导致子组件re-render

**父组件传入一个对象给子组件，子组件改变传入的对象props，父组件又是如何被更新到的？**

大前提：如果父组件传给子组件的props中有对象，那么子组件接收到的是这个对象的引用。也就是ParentComps中的this.person和SubComps中的this.person指向同一个对象

` // 假定父组件传person对象给子组件SubComps Vue.component('ParentComps',{ data(){ return { person:{ username:'aliarmo', pwd:123 } } }, template:` < div > < p > {{person.username}} </ p > < SubComps :person = "person" /> </ div > ` }) 复制代码`

现在我们在SubComps里面，更新person对象的某个属性,如：this.person.username='wmy' 这样会导致ParentComps和SubComps的更新，为什么呢？

因为Vue在ParentComps中会深度递归观察对象的每个属性，在第一次执行ParentComps的render的时候，绑定ParentComps的Watcher，传入到SubComps后，不会对传入的对象在进行观察，在第一次执行SubComps的render的时候，会绑定到SubComps的Watcher，因此当SubComps改变了this.person.username的值，会通知到两个Watcher，导致更新。这很好的解释了凭空在传入的props属性对象上挂载新的属性不触发渲染，因为传入的props属性对象是在父组件被观察的。

### 1.7 vnode diff算法---------划重点 ###

当组件的state发生变化，重新执行渲染函数生成新的vnode，然后将新生成的vnode与老的vnode进行对比，以最小的代价更新原有视图。 **diff算法的原理是通过移动、新增、删除和替换oldChildrenVnodes对应的结构来生成newChildrenVnodes对应的结构，并且每个老的元素只能被复用一次，老元素最终的位置取决于当前新的vnode。要明确传入diff算法的是两个sameVnode的孩子节点，从两者的开头和结尾位置，同时往中间靠，直到两者中的一个到达中间。**

PS：oldChildrenVnodes表示老的孩子vnode节点集合，newChildrenVnodes表示state变化后生成的新的孩子vnode节点集合

说这个算法之前，先得明白如何判断两个vnode为sameVnode，我只大体列一下：

* vnode的key值相等，例如 ` <Comps1 key="key1" />` , ` <Comps2 key="key2" />` ，key值就不相等, ` <Comps1 key="key1" />` , ` <Comps2 key="key1" />` , key值就是相等的, ` <div></div>` , ` <p></p>` ,这两个的key值是undefined，key值相等，这个是sameVnode的大前提。
* vnode的tag相同，都是注释或者都不是注释，同时定义或未定义data，标签为input则type必须相同，还有些其他的条件跟我们不太相关就不列出来了。

**整个vnode diff流程**

大前提，要看懂这个vnode diff，务必先明白vnode是啥，如何生成的，vnode与elm的关系，详情请看上面的vnode概念

* 如果两个vnode是sameVnode，则进行patch vnode
* patch vnode过程

（1）首先vnode的elm指向oldVnode的elm

（2）使用vnode的数据更新elm的attr，class，style，domProps，events等

（3）如果vnode是文本节点，则直接设置elm的text，结束

（4）如果vnode是非文本节点&&有孩子&&oldVnode没有孩子，则elm直接append

（5）如果vnode是非文本节点&&没有孩子&&oldVnode有孩子,则直接移除elm的孩子节点

（6）如果非文本节点&&都有孩子节点，则updateChildren，进入diff 算法， **前面5个步骤排除了不能进行diff情况**

* diff 算法，这里以web平台为例

**这里还有强调下，传入diff算法的是两个sameVnode的孩子节点** ，那么如何用newChildrenVnodes替换oldChildrenVnodes，最简单的方式莫过于，遍历newChildrenVnodes，直接重新生成这个html片段，皆大欢喜。但是这样做会 不断的createElement，对性能有影响，于是前辈们就想出了这个diff算法。

（1）取两者最左边的节点，判断是否为sameVnode，如果是则进行上述的 **第二步patch vnode过程** ，整个流程走完后，此时elm的class，style，events等已经更新了，elm的children结构也通过前面说的整个流程得到了更新，这时候就看是否需要移动这个elm了，因为都是孩子的最左边节点，因此位置不变，最左边节点位置向前移动一步

（2）如果不是（1）所述case，取两者最右边的节点，跟（1）的判定流程一样，不过是最右边节点位置向前移动一步

（3）如果不是（1）（2）所述case，取oldChildrenVnodes最左边节点和newChildrenVnodes最右边节点，跟（1）的判定流程一样，不过， **elm的位置需要移动到oldVnode最右边elm的右边** ，因为vnode取的是最右边节点，如果与oldVnode的最右边节点是sameVnode的话，位置是不用改变的，因此newChildrenVnodes的最右节点和oldChildrenVnodes的最右节点位置是对应的，但由于是复用的oldChildrenVnodes的最左边节点，oldChildrenVnodes最右边节点还没有被复用，因此不能替换掉，所以移动到oldChildrenVnodes最右边elm的右边。然后oldChildrenVnodes最左边节点位置向前移动一步，newChildrenVnodes最右边节点位置向前移动一步

（4）如果不是（1）（2）（3）所述case，取oldChildrenVnodes最右边节点和newChildrenVnodes最左边节点，跟（1）的判定流程一样，不过， **elm的位置需要移动到oldChildrenVnodes最左边elm的左边** ，因为vnode取的是最左边节点，如果与oldChildrenVnodes的最左边节点是sameVnode的话，位置是不用改变的，因此newChildrenVnodes的最左节点和oldChildrenVnodes的最左节点位置是对应的，但由于是复用的oldChildrenVnodes的最右边节点，oldChildrenVnodes最左边节点还没有被复用，因此不能替换掉，所以移动到oldChildrenVnodes最左边elm的左边。然后oldChildrenVnodes最右边节点位置向前移动一步，newChildrenVnodes最左边节点位置向前移动一步

（5）如果不是（1）（2）（3）（4）所述case，在oldChildrenVnodes中寻找与newChildrenVnodes最左边节点是sameVnode的oldVnode，如果没有找到，则用这个新的vnode创建一个新element，插入位置如后所述，如果找到了，则跟（1）的判定流程一样，不过 **插入的位置是oldChildrenVnodes的最左边节点的左边** ，因为如果newChildrenVnodes最左边节点与oldChildrenVnodes最左边节点是sameVnode的话，位置是不用变的，并且复用的是oldChildrenVnodes中找到的oldVNode的elm。被复用过的oldVnode后面不会再被取出来。然后newChildrenVnodes最左边节点位置向前移动一步

（6）经过上述步骤，oldChildrenVnodes或者newChildrenVnodes的最左节点与最右节点重合，退出循坏

（7）如果是oldChildrenVnodes的最左节点与最右节点先重合，说明newChildrenVNodes还有节点没有被插入，递归创建这些节点对应元素，然后 **插入到oldChildrenVnodes的最左节点的右边或者最右节点的左边** ，因为是从两者的开始和结束位置向中间靠拢，想想，如果newChildrenVNodes剩余的第一个节点与oldChildrenVnodes的最左边节点为sameVnode的话，位置是不用变的

（8）如果是newChildrenVnodes的最左节点与最右节点先重合，说明oldChildrenVnodes中有一段结构没有被复用，开始和结束位置向中间靠拢，因此 **没有被复用的位置是oldChildrenVnodes的最左边和最右边之间节点** ，删除节点对应的elm即可。

举个栗子来描述下具体的diff过程（web平台）：

` // 有Vue模板如下 <div> ------ oldVnode1，newVnode1，element1 <span v-if= "isShow1" ></span> -------oldVnode2，newVnode2，element2 <div :key= "key" ></div> -------oldVnode3，newVnode3，element3 <p></p> -------oldVnode4，newVnode4，element4 <div v-if= "isShow2" ></div> -------oldVnode5，newVnode5，element5 </div> // 如果 isShow1= true ，isShow2= true ，key= "aliarmo" 那么模板将会渲染成如下： <div> <span></span>--------------element2 <div key= "aliarmo" ></div>----------element3 <p></p>-------------element4 <div></div>----------element5 </div> // 改变state，isShow1= false ，isShow2= true ，key= "wmy" ，那么模板将会渲染成如下： <div> <div key= "wmy" ></div>------------element6 <p></p>-------------------element4 <div></div>---------element5 </div> 复制代码`

那么，改变state后的dom结构是如何生成的？

如上图，在 ` isShow1=true，isShow2=true，key="aliarmo"` 条件下，生成的vnode结构是： ` oldVnode1，oldVnodeChildren=[oldVnode2,oldVnode3,oldVnode4,oldVnode5]`

对应的dom结构为：

![](https://user-gold-cdn.xitu.io/2019/5/30/16b07752b8c15015?imageView2/0/w/1280/h/960/ignore-error/1)

改变state为 ` isShow1=false，isShow2=true，key="wmy"` 后，生成的新vnode结构是

` newVnode1，newVnodeChildren=[newVnode3,newVnode4,newVnode5]`

最左边两个新老vnode对比，也就是 ` oldVnode2` ， ` newVnode3` ，不是 ` sameVnode` ，

那最右边两个新老vnode对比，也就是 ` oldVnode5` ， ` newVnode5` ，是 ` sameVnode` ，不用移动原来的Element5所在位置，原有dom结构未发生变化，

最左边两个新老vnode对比，也就是 ` oldVnode2` ， ` newVnode3` ，不是 ` sameVnode` ，

那最右边两个新老vnode对比，也就是 ` oldVnode4` ， ` newVnode4` ，是 ` sameVnode` ，不用移动原来的Element4所在位置，原有dom结构未发生变化，

最左边两个新老vnode对比，也就是 ` oldVnode2` ， ` newVnode3` ，不是 ` sameVnode` ，

那最右边两个新老vnode对比，也就是 ` oldVnode3` ， ` newVnode3` ，由于key值不同，不是 ` sameVnode` ，

当前最左边和最右边对比， ` oldVnode2` ， ` newVnode3` ，不是 ` sameVnode`

当前最右边和最左边对比， ` oldVnode5` ， ` newVnode3` ，不是 ` sameVnode`

在遍历oldVnodeChildren，寻找与 ` newVnode3` 为 ` sameVnode` 的 ` oldVnode` ，没有找到，则用 ` newVnode3` 创建一个新的元素 ` Element6` ，插入到当前 ` oldVnode2` 所对应元素的最左边，dom结构发生变化 ` newVnodeChildren` 两头重合，退出循环，删除剩余未被复用元素 ` Element2` ， ` Element3`

### 1.8 core总结 ###

现在我们终于可以理一下，从new Vue()开始，core里面发生了些什么

* new Vue()或者new自定义组件构造函数（继承自Vue）
* 初始化，props，methods，computed，data，watch，并给state加上Observe，调用生命周期created
* 开始mount组件，mount之前确保render函数的生成
* new Watcher，导致render和patch，注意一个watcher对应一个组件，watcher对变化的响应是重新执行render生成vnode进行patch
* render在当前组件上下文（组件实例）执行，生成对应的vnode结构
* 如果没有oldVnode，那patch就是深度遍历vnode，生成具体的平台元素，给具体的平台元素添加属性和绑定事件，调用自定义指令提供的钩子函数，并append到已存在的元素上，在遍历的过程中，如果遇到的是自定义组件，则从 **步骤1** 开始重复
* 如果有oldVnode，那patch就是利用vnode diff算法在原有的平台元素上进行修修补补，不到万不得已不创建新的平台元素
* state发生变化，通知到state所在组件对应的watcher，重新执行render生成vnode进行patch，也就是回到 **步骤4**

## 2、compiler ##

Vue的compiler部分负责对template的编译，生成render和staticRender函数，编译一次永久使用，所以一般我们在构建的时候就做了这件事情，以提高页面性能。执行render和staticRender函数可以生成VNode，从而为core提供这一层抽象。

` template ==》 AST ==》 递归ATS生成render和staticRender ==》VNode`

**（1）template转化成AST过程**

先让我们来直观感受下AST，它描述了下面的template结构

` // Vue模板 let template = ` < div class = "Test" :class = "classObj" v-show = "isShow" > {{username}}:{{password}} < div > < span > hahhahahha </ span > </ div > < div v-if = "isVisiable" @ click = "onClick" > </ div > < div v-for = "item in items" > {{item}} </ div > </ div > ` 复制代码`

![](https://user-gold-cdn.xitu.io/2019/5/30/16b077941343a1e8?imageView2/0/w/1280/h/960/ignore-error/1)

**下面描述下template转为AST的简要过程：**

* 如果template是以<开始的字符串，则判断是评论，还是Doctype还是结束标签，或者是开始标签，这里只说处理开始和结束标签。

（1）如果是开始标签，则处理类似于下面字符串

` < div class = "Test" :class = "classObj" v-show = "isShow" > 复制代码`

通过正则可以很容易解析出tag，所有属性列表，再对属性列表进行分类，分别解析出v-if,v-for等指令，事件，特殊属性等，template去除被解析的部分，回到步骤1

（2）如果是结束标签，则处理类似于下面字符串，同样template去除被解析的部分，回到步骤1

` </ div > 复制代码` * 如果不是第一种情况，说明是字符串，处理类似于下面的插值字符串或者纯文本，同样template去除被解析的部分，回到步骤1
` {{username}}:{{password}} 或者 用户名：密码 复制代码` * 如果template为空，解析结束

**（2）AST生成render和staticRender**

主要是遍历ast（有兴趣的同学可以自己体验下，如：遍历AST生成还原上述模板，相信会有不一样的体验），根据每个节点的属性拼接渲染函数的字符串，如：模板中有v-if="isVisiable"，那么AST中这个节点就会有一个if属性，这样，在创建这个节点对应的VNode的时候，就会有

` (isVisiable) ? _c( 'div' ) : _e() 复制代码`

在with的作用下， ` isVisiable` 的值决定了VNode是否生成。当然，对于一些指令，在编译时是处理不了的，会在生成VNode的时候挂载在VNode上，解析VNode时再进行进一步处理，比如v-show，v-on。

下面是上面模板生成的render和staticRender：

` // render函数 ( function anonymous ( ) { with ( this ) { return _c( 'div' , { directives : [{ name : "show" , rawName : "v-show" , value : (isShow), expression : "isShow" }], staticClass : "Test" , class : classObj }, [_v( "\n " + _s(username) + ":" + _s(password) + "\n " ), _m( 0 ), _v( " " ), (isVisiable) ? _c( 'div' , { on : { "click" : onClick } }) : _e(), _v( " " ), _l((items), function ( item ) { return _c( 'div' , [_v(_s(item))]) })], 2 ) } } ) // staticRender ( function anonymous ( ) { with ( this ) { return _c( 'div' , [_c( 'span' , [_v( "hahhahahha" )])]) } } ) 复制代码`

其中this是组件实例，_c、_v分别用于创建VNode和字符串，像username和password是在定义组件时候传入的state并被挂载在this上。

## 3、platform ##

platform模块与具体平台相关，我们可以在这里定义平台相关接口传入runtime和compile，以实现具体平台的定制化，因此为其他平台带来Vue能力，大部分工作在这里。

需要传入runtime的是 **如何创建具体的平台元素，平台元素之间的关系以及如何append，insert，remove平台元素等，元素生成后需要进行的属性，事件监听等** 。拿web平台举例，我们需要传入document.createElement，document.createTextNode，遍历vnode的时候生成HTML元素；挂载时需要的insertBefore；state发生变化导致vnode diff时的remove，append等。还有生成HTML元素后，用setAttribute和removeAttribute操作属性；addEventListener和removeEventListener进行事件监听；提供一些有利于web平台使用的自定义组件和指令等等

需要传入compile的是对 **某些特殊属性或者指令在编译时的处理** 。如web平台，需要对class，style，model的特殊处理，以区别于一般的HTML属性；提供web平台专用指令，v-html（编译后其实是绑定元素的innerHTML），v-text（编译后其实是绑定元素的textContent），v-model，这些指令依赖于具体的平台元素。

# 三、应用 #

说了这么多，最终目的是为了 **复用Vue的core和compile，以期在其他的平台上带来Vue或者类Vue的开发体验** ，前面也说了很多复用成功的例子，如果你要为某一平台带来Vue的开发体验，可以进行参考。大前端概念下，指不定那天，汽车显示屏，智能手表等终端界面就可以用Vue来进行开发，爽歪歪。那么，如何复用呢？当然我只说必选的，你可以定制更多更复杂的功能方便具体平台使用。

* 

定义vnode生成具体平台元素所需要的nodeOps，也就是元素的增删改查，对于web来说nodeOps是要创建，移动真正的dom对象，如果是其他平台，可自行定义这些元素的操作方法；

* 

定义vnode生成具体平台元素所需要的modules，对于web来说，modules是用来操作dom属性的方法；

* 

具体平台需要定义一个自己的$mount方法给Vue，挂载组件到已存在的元素上；

* 

还有一些方法，如isReservedTag，是否为保留的标签名，自定义组件名称不能与这些相同；mustUseProp，判定元素的某个属性必须要跟组件state绑定，不能绑定一个常量等等；

# 四、总结 #

软件行业有一句名言，没有什么问题是添加一层抽象层不能解决的，如果有那就两层，Vue的设计正是如此（可能Vue也是借鉴别人的），compile的AST抽象层衔接了模板语法和render函数，经过VNode这个抽象层让core剥离了具体平台。

这篇文章的最终目的是为了让大家了解到复用Vue源码的core和compile进行二次开发，可以做到为具体平台带来Vue或者类Vue的开发体验。当然，这只是一种思路，说不定哪天，Vue风格开发体验失宠了，那么我们也可以把这种思路应用到新开发风格上。

![](https://user-gold-cdn.xitu.io/2019/5/30/16b07e9c262ab4e1?imageView2/0/w/1280/h/960/ignore-error/1)
> 
> 
> 
> 
> 关注【IVWEB社区】公众号获取每周最新文章，通往人生之巅！
> 
>