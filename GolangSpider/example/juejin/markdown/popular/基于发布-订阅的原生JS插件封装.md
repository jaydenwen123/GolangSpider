# 基于"发布-订阅"的原生JS插件封装 #

大家好，我是神三元。 今天我们来做一个小玩意，用原生JS封装一个动画插件。效果如下：

![](https://user-gold-cdn.xitu.io/2019/6/3/16b1c3c212ee7299?imageslim) 这个飞驰的小球看起来是不是特有灵性呢？没错，它就是用原生JS实现的。 接下来，就让我们深入细节，体会其中的奥秘。相信这个实现的过程，会比动画本身更加精彩!

#### 一、需求分析 ####

封装一个插件，将小球的DOM对象作为参数传入，使得小球在鼠标按下和放开后能够运动，在水平方向做匀减速直线运动，初速度为鼠标移开瞬间的速度，在竖直方向的运动类似于自由落体运动。并且，小球的始终在不离开浏览器的边界运动，碰到边界会有如图的反弹效果。

#### 二、梳理思路 ####

分析这样的一个过程，其中大致会经历一下的关键步骤：

* 1、鼠标按下时，记录小球的初始位置信息
* 2、鼠标按下后滑动，记录松开鼠标瞬间的移动速度
* 3、鼠标松开后，在水平方向上，让小球根据刚刚记录的移动速度进行匀减速运动，竖直方向设定一个竖直向下的加速度，开始运动。
* 4、水平方向速度减为0时，水平方向运动停止;竖直方向速度减为0或者足够小时，竖直方向运动停止。

#### 三、难点分析 ####

看到这里，估计你的思路清晰了不少，但可能还是有一些比较难以搞定的问题。

首先，你怎么拿到松开手瞬间的小球移动速度？如何去表达出这个加速度的效果？

在实现方面，这是非常重要的问题。不过，其实非常的简单。

浏览器本身就是存在反应时间的，你可以把它当做一个摄像机,在给DOM元素绑定了事件之后,每隔一段时间(一般非常的短,根据不同浏览器厂商和电脑性能而定，这里我用到chrome，保守估计为20ms)会给这个元素拍张照，记录它的状态。在按下鼠标之后的拖动过程中，事实上会给元素拍摄无数张照片。如果现在每经过一段时间,我记录当下当前照片与上一段照片的位置差，那么最后一次拍照和倒数第二次拍照的小球位置差距，是不是就可以作为离开的瞬时速度呢？当然可以啦。废话不多说，上图：

![](https://user-gold-cdn.xitu.io/2019/6/3/16b1c8a6588a4343?imageView2/0/w/1280/h/960/ignore-error/1)

同样，对实现加速度的效果，首先弄清一个问题，什么是速度？速度就是单位时间内运动的距离，这里暂且把它当做20ms内的距离，那么我每次拍照时，将这个距离增加或减少一个值，这个值就是加速度。

#### 四、初步实现 ####

当大部分问题考虑清楚之后,现在开始实现。 首先是基本的样式，比较简单。

` <!DOCTYPE html> < html > < head > < meta charset = "UTF-8" > < title > 狂奔的小球 </ title > < link rel = "stylesheet" href = "css/reset.min.css" > < style > html , body { height : 100% ; overflow : hidden; } #box { position : absolute; top : 100px ; left : 100px ; width : 150px ; height : 150px ; border-radius : 50% ; background : lightcoral; cursor : move; z-index : 0 ; } </ style > </ head > < body > < div id = "box" > </ div > </ body > </ html > 复制代码`

现在来完成核心的JS代码，采用ES6语法

` //drag.js class Drag { //ele为传入的DOM对象 constructor (ele) { //初始化参数 this.ele = ele; [ 'strX' , 'strY' , 'strL' , 'strT' , 'curL' , 'curT' ].forEach( item => { this [item] = null ; }); //为按下鼠标绑定事件,事件函数一定要绑定this,在封装过程中this统一指定为实例对象，下不赘述 this.DOWN = this.down.bind( this ); this.ele.addEventListener( 'mousedown' , this.DOWN); } down(ev) { let ele = this.ele; this.strX = ev.clientX; //鼠标点击处到浏览器窗口最左边的距离 this.strY = ev.clientY; //鼠标点击处到浏览器窗口最上边的距离 this.strL = ele.offsetLeft; //元素到浏览器窗口最左边的距离 this.strT = ele.offsetTop; //元素到浏览器窗口最上边的距离 this.MOVE = this.move.bind( this ); this.UP = this.up.bind( this ); document.addEventListener( 'mousemove' , this.MOVE); document.addEventListener( 'mouseup' , this.UP); //flag //清理上一次点击形成的一些定时器和变量 clearInterval( this.flyTimer); this.speedFly = undefined ; clearInterval( this.dropTimer); } move(ev) { let ele = this.ele; this.curL = ev.clientX - this.strX + this.strL; this.curT = ev.clientY - this.strY + this.strT; ele.style.left = this.curL + 'px' ; ele.style.top = this.curT + 'px' ; //flag //功能: 记录松手瞬间小球的速度 if (! this.lastFly) { this.lastFly = ele.offsetLeft; this.speedFly = 0 ; return ; } this.speedFly = ele.offsetLeft - this.lastFly; this.lastFly = ele.offsetLeft; } up(ev) { //给前两个事件解绑 document.removeEventListener( 'mousemove' , this.MOVE); document.removeEventListener( 'mouseup' , this.UP); //flag //水平方向 this.horizen.call( this ); this.vertical.call( this ); } //水平方向的运动 horizen() { let minL = 0 , maxL = document.documentElement.clientWidth - this.ele.offsetWidth; let speed = this.speedFly; speed = Math.abs(speed); this.flyTimer = setInterval( () => { speed *=.98 ; Math.abs(speed) <= 0.1 ? clearInterval( this.flyTimer): null ; //小球当前到视口最左端的距离 let curT = this.ele.offsetLeft; curT += speed; //小球到达视口最右端，反弹 if (curT >= maxL) { this.ele.style.left = maxL + 'px' ; speed *= -1 ; return ; } //小球到达视口最右端，反弹 if (curT <= minL) { this.ele.style.left = minL + 'px' ; speed *= -1 ; return ; } this.ele.style.left = curT + 'px' ; }, 20 ); } //竖直方向的运动 vertical() { let speed = 9.8 , minT = 0 , maxT = document.documentElement.clientHeight - this.ele.offsetHeight, flag = 0 ; this.dropTimer = setInterval( () => { speed += 10 ; speed *=.98 ; Math.abs(speed) <= 0.1 ? clearInterval( this.dropTimer): null //小球当前到视口最左端的距离 let curT = this.ele.offsetTop; curT += speed; //小球飞到视口顶部，反弹 if (curT >= maxT) { this.ele.style.top = maxT + 'px' ; speed *= -1 ; return ; } //小球落在视口底部，反弹 if (curT <= minT) { this.ele.style.top = minT + 'px' ; speed *= -1 ; return ; } this.ele.style.top = curT + 'px' ; }, 20 ); } } window.Drag = Drag; 复制代码`

到此，完整的效果就出来了，你可以自己复制体验一下。

#### 四、采用发布-订阅 ####

估计读完这段代码，你也体会到了这个功能的实现是非常容易实现的。但是实际上，作为一个插件的标准来讲，这段代码是存在一些潜在的问题的，这些问题并不是逻辑上的问题，而是设计问题。直白一点说，其实是它的扩展性不强，倘若我要对某一个效果进行重新调整或者直接重写效果，我需要再这繁重的代码里面去搜索和修改。

因此，我们这里的目的并不只是提供一个功能，它绝不只是一个玩具，我们应当思考，如何将它做的更有通用性，能够得到最大程度的复用。 这里，我想引用软件工程领域耳熟能详的SOLID设计原则中的O部分————开放封闭原则。

` 开放封闭原则主要体现在两个方面： 对扩展开放，意味着有新的需求或变化时，可以对现有代码进行扩展，以适应新的情况。 对修改封闭，意味着类一旦设计完成，就可以独立完成其工作，而不要对类进行任何修改。 复制代码`

我们希望尽可能少地对类本身进行修改，因为你无法预测具体的功能会如何变化。

那怎么解决这个问题呢？很简单，对扩展开放，我们就将具体的效果代码以扩展的方式提供，对类扩展，而不是全部放在类里面。 我们的具体做法就是采用发布-订阅模式。

` 发布—订阅模式又叫观察者模式，它定义对象间的一对多的依赖关系，当一个对象的状态发生改变时，所有依赖于它的对象都将得到通知。 复制代码`

拿刚刚实现的功能来说，在对象创建的时候，我就开辟一个池子，将需要执行的方法放进这个池子，当鼠标按下的时候，我把池子里面的函数拿过来依次执行，对于鼠标松开就再创建一个池子，同理，这就是发布-订阅。

jQuery里面有现成的发布订阅方法。

` //开辟一个容器 let $plan = $.callBack(); //往容器里面添加函数 $plan.add( function ( x, y ) { console.log(x, y); }) $plan.add( function ( x, y ) { console.log(y, x); }) $plan.fire( 10 , 20 ); //会输出10，20 20,10 //$plan.remove(function)用来从容器中删除某个函数 复制代码`

现在我们不妨原生JS手写一下简单的发布-订阅,让我们原生撸到底

` //subscribe.js class Subscribe { constructor () { //创建容器 this.pond = []; } //向容器中增加方法，注意去重 add(fn) { let pond = this.pond, isExist = false ; //去重环节 pond.forEach( item => item === fn ? isExist = true : null ); !isExist ? pond.push(fn) : null ; } remove(fn) { let pond = this.pond; pond.forEach( ( item, index ) => { if (item === fn) { //提一下我在这里遇到的坑，这里如果写item=null是无效的 //例子：let a = {name: funtion(){}}; //let b = a.name; //这个时候操作b的值对于a的name属性是没有影响的 pond[index] = null ; } }) } fire(...arg) { let pond = this.pond; for ( let i = 0 ; i < pond.length; i++) { let item = pond[i]; //如果itme为空了,最好把它删除掉 if (item === null ) { pond.splice(i, 1 ); //如果用了splice要防止数组塌陷问题，即删除了一个元素后，后面所有元素的索引默认都会减1 i--; continue ; } item(...arg); } } } window.Subscribe = Subscribe; 复制代码` ` //测试一下 let subscribe = new Subscribe(); let fn1 = function fn1 ( x, y ) { console.log( 1 , x, y); }; let fn2 = function fn2 ( ) { console.log( 2 ); }; let fn3 = function fn3 ( ) { console.log( 3 ); subscribe.remove(fn1); subscribe.remove(fn2); }; let fn4 = function fn4 ( ) { console.log( 4 ); }; subscribe.add(fn1); subscribe.add(fn1); subscribe.add(fn2); subscribe.add(fn1); subscribe.add(fn3); subscribe.add(fn4); setInterval( () => { subscribe.fire( 100 , 200 ); }, 1000 ); 复制代码`

结果：

![](https://user-gold-cdn.xitu.io/2019/6/3/16b1d164273c52f4?imageView2/0/w/1280/h/960/ignore-error/1) 确定过眼神，你就是对的Subscribe。(手动滑稽)

#### 五、优化代码 ####

` //Drag.js if ( typeof Subscribe === 'undefined' ) { throw new ReferenceError ( '没有引入subscribe.js!' ); } class Drag { constructor (ele) { this.ele = ele; [ 'strX' , 'strY' , 'strL' , 'strT' , 'curL' , 'curT' ].forEach( item => { this [item] = null ; }); this.subDown = new Subscribe; this.subMove = new Subscribe; this.subUp = new Subscribe; //=>DRAG-START this.DOWN = this.down.bind( this ); this.ele.addEventListener( 'mousedown' , this.DOWN); } down(ev) { let ele = this.ele; this.strX = ev.clientX; this.strY = ev.clientY; this.strL = ele.offsetLeft; this.strT = ele.offsetTop; this.MOVE = this.move.bind( this ); this.UP = this.up.bind( this ); document.addEventListener( 'mousemove' , this.MOVE); document.addEventListener( 'mouseup' , this.UP); this.subDown.fire(ele, ev); } move(ev) { let ele = this.ele; this.curL = ev.clientX - this.strX + this.strL; this.curT = ev.clientY - this.strY + this.strT; ele.style.left = this.curL + 'px' ; ele.style.top = this.curT + 'px' ; this.subMove.fire(ele, ev); } up(ev) { document.removeEventListener( 'mousemove' , this.MOVE); document.removeEventListener( 'mouseup' , this.UP); this.subUp.fire( this.ele, ev); } } window.Drag = Drag; 复制代码` ` //dragExtend.js function extendDrag ( drag ) { //鼠标按下 let stopAnimate = function stopAnimate ( curEle ) { clearInterval(curEle.flyTimer); curEle.speedFly = undefined ; clearInterval(curEle.dropTimer); }; //鼠标移动 let computedFly = function computedFly ( curEle ) { if (!curEle.lastFly) { curEle.lastFly = curEle.offsetLeft; curEle.speedFly = 0 ; return ; } curEle.speedFly = curEle.offsetLeft - curEle.lastFly; curEle.lastFly = curEle.offsetLeft; }; //水平方向的运动 let animateFly = function animateFly ( curEle ) { let minL = 0 , maxL = document.documentElement.clientWidth - curEle.offsetWidth, speed = curEle.speedFly; curEle.flyTimer = setInterval( () => { speed *=.98 ; Math.abs(speed) <= 0.1 ? clearInterval(animateFly): null ; let curT = curEle.offsetLeft; curT += speed; if (curT >= maxL) { curEle.style.left = maxL + 'px' ; speed *= -1 ; return ; } if (curT <= minL) { curEle.style.left = minL + 'px' ; speed *= -1 ; return ; } curEle.style.left = curT + 'px' ; }, 20 ); }; //竖直方向的运动 let animateDrop = function animateDrop ( curEle ) { let speed = 9.8 , minT = 0 , maxT = document.documentElement.clientHeight - curEle.offsetHeight; curEle.dropTimer = setInterval( () => { speed += 10 ; speed *=.98 ; Math.abs(speed) <= 0.1 ? clearInterval(animateFly): null ; let curT = curEle.offsetTop; curT += speed; if (curT >= maxT) { curEle.style.top = maxT + 'px' ; speed *= -1 ; return ; } if (curT <= minT) { curEle.style.top = minT + 'px' ; speed *= -1 ; return ; } curEle.style.top = curT + 'px' ; }, 20 ); }; drag.subDown.add(stopAnimate); drag.subMove.add(computedFly); drag.subUp.add(animateFly); drag.subUp.add(animateDrop); }; 复制代码`

在html文件中加入如下script

` < script > //原生JS 小技巧: //直接写box跟document.getElementById('box')是一样的效果 let drag = new Drag(box); extendDrag(drag); </ script > 复制代码`

接下来，你就能重新看到那个活泼的小球啦。

#### 六、结(chui)语(niu) ####

恭喜你，读到了这里，相当不容易啊。先为你点个赞！

在这里我并不是简单讲讲效果的实现、贴贴代码就过去了，而是带你体验了封装插件的整个过程。有了发布-订阅的场景，理解这个设计思想就更加容易了。其实你看在这个过程中，功能并没有添加多少，但是这波操作确实值得，因为它让整个代码更加的灵活。回过头看，比如DOM2的事件池机制，vue的生命周期钩子等等，你就会明白它们为什么要这么设计，原理上和这次封装没有区别，这样一想，很多东西就更加清楚了。

在我看来，无论你是做哪个端的开发工作，其实大部分业务场景、大部分流行的框架技术都很可能会在若干年后随风而逝，但真正留下来的、伴随你一生的东西是编程思想。在我的理解中，编程的意义远不止造轮子，写插件，来显得自己金玉其外，而是留心思考，提炼出一些思考问题的方式，从而在某个确定的时间点让你拥有极其敏锐的判断，来指导和优化你下一步的决策，而不是纵身于飞速迭代的技术浪潮，日渐焦虑。我觉得这是一个程序员应该追求的东西。