# 【切图仔日常】浅谈CSS3动画之凌波微步--steps() #

# 背景 #

一日敲代码的我，得到一个需求：写一个10秒的倒计时。

用JavaScript定时器麻溜写完之后，恰好同事勇司机接完水。瞟了一眼，然后凑过来说，这个用CSS3也可以写，而且一行JavaScript都不用写。

"一行JavaScript都不用写，纯CSS3就可以写。CSS3有这么溜的操作！"

''对呀！CSS3 animation当中有一个steps()，你上网查一下就知道了！"

"涨姿势了！赶紧查阅一下！"

# animation-timing-function #

> 
> 
> 
> CSS animation-timing-function属性定义CSS动画在每一动画周期中执行的节奏。可能值为一或多个timing
> function(数学函数)对于关键帧动画来说，timing
> function作用于一个关键帧周期而非整个动画周期，即从关键帧开始开始，到关键帧结束结束。 定义于一个关键帧区块的缓动函数(animation
> timing function)应用到改关键帧； 另外，若该关键帧没有定义缓动函数，则使用定义于整个动画的缓动函数。 -- MDN Web 文档 (
> https://link.juejin.im?target=http%3A%2F%2Fwww.baidu.com%2Flink%3Furl%3Dmd0Ik_8COQRxAf1BLb94_nGf6ZhOB-1br1nHbI6q_I5Co7GiUTFboDtuKv3wfGlR
> )
> 
> 

当看到这一长串的CSS3属性，相信有部分童鞋们肯定会陌生这个属性。因为当时我第一眼看到也是一脸问号，表示只熟悉前面的 **animation** 。

![](https://user-gold-cdn.xitu.io/2019/6/2/16b18a1ad2892759?imageView2/0/w/1280/h/960/ignore-error/1)

` animation-timing-function` 其实不记得很正常，这个属于 **animation** 的单独一个属性，就好比 ` background-image` 这种css属性大家很少单独使用，为了书写代码方便而更多使用 ` background： url('test.png') center no-repeat；` 这种简写定义整个背景。

同理在大家对于 **animation** 使用过程当中，普遍采用 ` animation: move 5s infinite ease both;` 这种类简写定义整个动画属性。这样的简写模式，简洁、高效何乐不为呢。

![引用XXXschool网站截图](https://user-gold-cdn.xitu.io/2019/6/2/16b18a1ad67acd76?imageView2/0/w/1280/h/960/ignore-error/1)

借用XXXschool网站上给的使用方法，但是并没有找到我们想要的 steps()。所以让我们切换到MDN上面的看一下属性介绍：

![MDN语法相关参数](https://user-gold-cdn.xitu.io/2019/6/2/16b18a1ada51d2bc?imageView2/0/w/1280/h/960/ignore-error/1)

当我切换到MDN上面就可以看到语法介绍上面有steps(),而且整体数量级别也不是一点差别，描述非常详细。

PS: 当然可能这XXXschool上面更新可能比较慢一点，但是作为日用简单的查询还是可以用，深入研究探索还是建议去MDN、W3.org相关网站上面查阅。

# 缓动函数与阶梯函数 #

总的来说 ` animation-timing-function` 基本就是分为两大类型函数：

## 1、缓动函数： ##

它描述了在一个过渡或动画中一维数值的改变速度。这实质上让你可以自己定义一个加速度曲线，以便动画的速度在动画的过程中可以进行改变。缓动函数指定动画效果在执行时的速度，使其看起来更加真实。

例如现实物体照着一定节奏移动，并不是一开始就移动很快的。当我们打开抽屉时，首先会让它加速，然后慢下来。当某个东西往下掉时，首先是越掉越快，撞到地上后回弹，最终才又碰触地板。

前端做过渡的动画大部分都是采用的是这个缓动函数，比较书面一点称呼就是：**立方贝塞尔曲线（cubic Bézier curves）**的子集

![立方贝赛尔曲线](https://user-gold-cdn.xitu.io/2019/6/2/16b18a1adb53c1ae?imageView2/0/w/1280/h/960/ignore-error/1)

这里就不继续述说缓动函数了，如果各位小伙伴感兴趣推荐两个网站供小伙伴折腾：

[缓动函数速查表]( https://link.juejin.im?target=http%3A%2F%2Feasings.net%2Fzh-cn ) [贝塞尔曲线在线生成工具]( https://link.juejin.im?target=http%3A%2F%2Fcubic-bezier.com%2F )

## 2、阶梯函数 ##

> 
> 
> 
> 数学中，一个实数函数被称为阶段函数（或者阶梯函数），则它可以被写作：有限的间隔指标函数的线性组合。不正规的说法是，一个阶段函数就是一个分段常值函数，只是含有的阶段很多但是有限。
> --wiki百科
> 
> 

![图片来自wiki百科](https://user-gold-cdn.xitu.io/2019/6/2/16b18a1adf15b1bb?imageView2/0/w/1280/h/960/ignore-error/1)

看完这些书面化的介绍，学渣的洪荒之内爆发了，搞一个steps()要这么难搞？

![](https://user-gold-cdn.xitu.io/2019/6/2/16b18a1b01d5d72d?imageView2/0/w/1280/h/960/ignore-error/1)

# steps()属性 #

吃包辣条压压惊，查找资料继续翻阅文档。几番查找，大致整理如下：

首先缓动函数就像一个人走路一样，同样的一段路。他可以迅速的走过去，也可以快速的跑过去，也可以跑一半的距离然后再走过去。所以无论如何控制我们选用那一种运动方式，移动到始终都是靠这个人双腿一步步移动完成的。

而阶梯函数就像一个失去双腿，但是却意外获得瞬间移动的人。他可以瞬间移动终点，他也可以瞬间移动一半距离，然后再瞬移到终点。所以阶梯函数着重控制的分几次瞬移完成。

是不是有那么一点感觉，不像刚才那个看到函数那么的头疼了！

![](https://user-gold-cdn.xitu.io/2019/6/2/16b18a1b05a9be0d?imageView2/0/w/1280/h/960/ignore-error/1)

## 基本语法： ` steps(<integer>[, [ start | end ] ]?)` ##

参数一： **integer** 中文意思就是 **整数** ,由此说明第一个参数是整数。主要用来指定函数间隔的数量，且必须是正整数(大于0)。 参数二： 可选 ，接受 **start** 和 **end** 两个值，指定在每个间隔的起点或是终点发生阶跃变化，默认为 **end** 。

**Keyword values** : ` start-step` 、 ` end-step` 另外相信记忆力比较好的童鞋，这个时候会注意到上述两个关键词。 ` step-start` 等同于 ` steps(1,start)` ，动画分成1步，动画执行时为开始左侧端点的部分为开始； ` step-end` 等同于 ` steps(1,end)` ：动画分成一步，动画执行时以结尾端点为开始，默认值为 **end** 。

这时借用W3.org文档提供的图表来看，对比wiki的图表来看，结合我的说明此时小伙伴们的思路应该会更加清晰。

![图片来自w3.org](https://user-gold-cdn.xitu.io/2019/6/2/16b18a1b0808ee8a?imageView2/0/w/1280/h/960/ignore-error/1)

## 动画卡通人物： ##

扯了半天的理论，没有实际的案例估计小伙伴们也记不住这些东西，那么先来一个栗子！

![敲黑板，赶紧麻溜的回神](https://user-gold-cdn.xitu.io/2019/6/2/16b18a1b19993553?imageView2/0/w/1280/h/960/ignore-error/1)

各大网站大致逛了一下，发现SegmentFault网站上有位哥们翻译国外写teps()里面插入的案例不错，其中一个例子用来帮助整理和学习非常的不错。

传送门： [【译】css动画里的steps()用法详解]( https://link.juejin.im?target=https%3A%2F%2Fsegmentfault.com%2Fa%2F1190000007042048 )

如下一幅图是由10个小可爱卡通人物组成的，那么我们的任务就是让这个卡通人物动起来。

![雪碧动画图](https://user-gold-cdn.xitu.io/2019/6/2/16b18a1b23e07fe9?imageView2/0/w/1280/h/960/ignore-error/1)

思路分析：

整幅图是由10个相同大小的矩形组成，每一个矩形都恰好完整包裹卡通人物，一个标准完美的雪碧图。接下来就将这个雪碧图作为背景插入展示的DIV盒模型当中。（PS：这个DIV盒模型与小矩形的宽高一样）

之后通过background-position去控制雪碧图位置定位，例如、0s显示默认第一个动作，然后100ms之后，background-position,瞬间改版了X轴方向，显示的第二个动作，然后又过了100ms，又展示第三张图片。依次类推，利用人眼的 **视觉暂留** 从而形成动画的效果。

> 
> 
> 
> 视觉暂留(Persistence of vision)  
> 现象是光对视网膜所产生的视觉在光停止作用后，仍保留一段时间的现象，其具体应用是电影的拍摄和放映。原因是由视神经的反应速度造成的。是动画、电影等视觉媒体形成和传播的根据。视觉实际上是靠眼睛的晶状体成像，感光细胞感光，并且将光信号转换为神经电流，传回大脑引起人体视觉。感光细胞的感光是靠一些感光色素，感光色素的形成是需要一定时间的，这就形成了视觉暂停的机理。
> --摘自网络
> 
> 

所以抛开重复的兼容代码，10行左右代码的就能实现一个动画。

`.hi { width: 50px; height: 72px; // 矩形的宽高 background-image: url( "http://s.cdpn.io/79/sprite-steps.png" ); animation: play .8s steps(10) infinite; // 分10步完成0到-500px完成并无限重复 } @keyframes play { from { background-position: 0px; } to { background-position: -500px; } } 复制代码`

最终效果图：

![招手小人.gif](https://user-gold-cdn.xitu.io/2019/6/2/16b18a1b341d5997?imageslim)

仅仅是代码上就应该比JavaScript要简洁的多，是不是感觉开启了新的世界大门！

![](https://user-gold-cdn.xitu.io/2019/6/2/16b18a1b4116993e?imageView2/0/w/1280/h/960/ignore-error/1)

# 倒计时案例： #

理论 + demo 都看了一遍，那么就开始解决实际工作需求：用steps()函数敲一个倒计时demo出来。

![日常需求](https://user-gold-cdn.xitu.io/2019/6/2/16b18a1b646a7103?imageView2/0/w/1280/h/960/ignore-error/1)

思考点：

通过需求来看，这个10秒倒计时需要控制到10ms的级别.第四位数，每10ms变化一次。然后第三位数每间隔100ms变化一次，第二位数字每间隔1s变化一次，第一位数字不去做变化。

理清楚思路之后就好办了，接下来code部分了。

在敲代码之前，先做好一张等比距离的雪碧图：

![需求雪碧图](https://user-gold-cdn.xitu.io/2019/6/2/16b18a1b52cd8ce2?imageView2/0/w/1280/h/960/ignore-error/1)

> 
> 
> 
> PS：因为之前的需求是做成的一张白色透明图片，而文章背景颜色比较淡。所以临时截了一个黑色背景图作为参考
> 
> 

**HTML代码：**

` <div class= "time" > <div class= "s10 num" ></div> <div class= "s0 num" ></div> <div class= "spot" >:</div> <div class= "ms0 num" ></div> <div class= "ms10 num" ></div> </div> 复制代码`

**CSS代码：**

`.time { /*display: none;*/ position: absolute; top: 2.05rem; left: 4.73rem; width: 3.22rem; height: 0.57rem; transform: scale(2); } .num { position: absolute; top: 0.2rem; width: 0.28rem; height: 0.37rem; background: url(img/timeNum.png) no-repeat center; background-size: cover; background-position: 0 0; } .spot { position: absolute; top: 0.06rem; left: 1.70rem; color: #fff; font-size: 24px; line-height: 0.57rem; } .s10 { left: 0.9rem; } .s0 { left: 1.25rem; -webkit-animation: sMove 10s steps(10, start); } .ms0 { left: 2rem; -webkit-animation: sMove 1s steps(10, start) infinite; } .ms10 { left: 2.3rem; -webkit-animation: sMove 100ms steps(10, start) infinite ; } @-webkit-keyframes sMove { 0% { background-position: -2.80rem 0; } 100% { background-position: 0 0; } } 复制代码`

最终效果图：

![Css3倒计时.gif](https://user-gold-cdn.xitu.io/2019/6/2/16b18a1b692124af?imageslim)

通过代码可以看到除去CSS代码当中一样基本的样式部分。 **s0** 、 **ms0** 、 **ms10** 这三个元素基本的animation部分基本都是一样，只是完成的时间以及和循环的次数不一样。是不是感觉不可思议，这么简单飘逸就搞定了？没错，就是那么飘逸！

![](https://user-gold-cdn.xitu.io/2019/6/2/16b18a1b6e3190d1?imageView2/0/w/1280/h/960/ignore-error/1)

**完成时间差异：** 因为雪碧图有10个数字，一次只是展示1/10，所以要分10步才能完成一张图的循环。那么如同之前的思路分析 **ms10** 走一步需要10ms，走完一个循环10步就需要100ms， **ms0** 就需要1s走完一个循环， **s0** 就需要10s才能走完一个循环。

**循环次数** 有小伙伴 **ms10** 、 **ms0** 可能会问，为什么不给定相应的循环次数，而是给无限循环呢。当然给你固定的参数也是可以，我之所以给这个无限循环次数，是因为我偷懒不想算他们的循环次数。

我更习惯给 **s0** 加一个监听事件 **webkitAnimationEnd** ，因为 **s0** 没有指定循环次数所以默认值为循环一次，当 **s0** 动画执行完毕之后，我只需要监听 **webkitAnimationEnd** 便开始下一段程序。

# 小结： #

好了！以上就是鄙人对于CSS3动画当中steps()一点简单理解，权当抛了一下砖头引大神们的玉。后续如何用steps()玩出各种花样就看各位看官自己了！

原创文章，文笔有限，才疏学浅，文中若有不正之处，再次再次再次欢迎各位啪啪的打脸赐教。（有句话说的好，重要的词得说三遍。）

PS：对于里面完整blogDemo代码感兴趣的，可以从本人github上阅览。 [源码demo传送门地址]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2FCheDabang%2FblogCode )

![个人敲鼓](https://user-gold-cdn.xitu.io/2019/6/3/16b18fd2c65763e1?imageslim)

我是车大棒，我为我自己插眼。