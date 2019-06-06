# 浅谈CSS三栏布局（包括双飞翼布局和圣杯布局） #

# CSS三栏布局 #

## 写在正文之前 ##

* PS：真想说一句，写博客，真香！（虽然我知道根本没人看）

## 三栏布局的概念 ##

* 三栏布局概念听起来很简单，就是让三列从左到右排列，左边区域和右边区域定宽，而中间内容区域宽度自适应。就像下图这样： ![在这里插入图片描述](https://user-gold-cdn.xitu.io/2019/6/5/16b26cef181089c0?imageView2/0/w/1280/h/960/ignore-error/1) 当然要注意：我们这里所说的中间部分宽度自适应就是随着屏幕的大小改变而自己适应的过程。这也是三栏布局产生的原因。简单的来说呢，就是两边栏固定，中间栏自适应。这种布局很古老，但依旧非常经典，因为有好多地方都存在它的身影，包括一些大厂面试的时候还是很喜欢问这个问题的。

## 三栏布局的具体实现及原理分析 ##

### 第一种：浮动三栏布局 ###

* 先给出html部分代码： ` < div class = "wrapper" > < div class = "left" > left </ div > < div class = "right" > right </ div > < div class = "content" > content </ div > </ div > 复制代码`
* 下面是css部分代码： `.wrapper { text-align : center; color : #fff ; overflow : hidden; /* 这里清除浮动，因为触发了BFC */ line-height : 200px ; }.left { float : left; width : 200px ; height : 200px ; background-color : red; }.right { float : right; width : 200px ; height : 200px ; background-color : blue; }.content { height : 200px ; margin : 0 200px ; background-color : lime; } 复制代码`

* 效果如下： ![在这里插入图片描述](https://user-gold-cdn.xitu.io/2019/6/5/16b26cef1886def2?imageView2/0/w/1280/h/960/ignore-error/1)
* 那么这种方法是很简单的，相信各位瞬间就能看懂。不过也有一些需要注意的细节点：

* 这种方法里面，DOM节点中的 left, right, content 这三个块是不能换顺序的，也就是说这种方法的缺点很明显：浏览器自上而下解析代码渲染DOM树，那么content内容区域不能被优先渲染出来。
* 至于为什么不能换顺序，大家也很清楚：因为我们要让他们在一行内展示，那么必须让左右这两个块漂浮起来不占原来的位置了，才能让content区域跻身而入，当然浮动的问题我就不需要再多说。。。

* 所以总结起来就是： **可以实现效果，但是不完美，因为不能优先渲染content区域** 。

### 第二种：定位三栏布局 ###

* html代码： ` < div class = "wrapper" > < div class = "content" > content </ div > < div class = "left" > left </ div > < div class = "right" > right </ div > </ div > 复制代码`
* css代码： `.wrapper { width : 100% ; line-height : 200px ; color : #fffdef ; text-align : center; position : relative; }.content { margin : 0 200px ; background-color : lime; height : 200px ; }.left { position : absolute; top : 0 ; left : 0 ; width : 200px ; height : 200px ; background-color : red; }.right { position : absolute; top : 0 ; right : 0 ; width : 200px ; height : 200px ; background-color : blue; } 复制代码`
* 效果如下： ![在这里插入图片描述](https://user-gold-cdn.xitu.io/2019/6/5/16b26cef187c34ca?imageView2/0/w/1280/h/960/ignore-error/1)
* 好吧，其实效果跟前面是一样的，这种方法也相对比较简单，实现原理和前面的浮动实现也差不多。不过他的盒子高度靠content去撑起来，而它的优点就是content这次可以被优先渲染出来了，因为它被排在了第一位。真是可喜可贺！

### 第三种：用BFC原理做三栏布局 ###

* BFC特性有一点是触发了BFC的盒子不会和浮动的盒子重叠，也就是说触发BFC的盒子不会被浮动的盒子盖住，那么问题解决一半了，我们看一下具体的实现：
* html代码： ` < div class = "wrapper" > < div class = "left" > left </ div > < div class = "right" > right </ div > < div class = "content" > content </ div > </ div > 复制代码`
* css代码： `.wrapper { text-align : center; color : #fffdef ; width : 100% ; line-height : 200px ; font-size : 40px ; }.left { float : left; width : 200px ; height : 200px ; background-color : red; }.right { float : right; width : 200px ; height : 200px ; background-color : blue; }.content { height : 200px ; background-color : lime; /* 这会形成BFC区域，不会与浮动的元素重叠 */ overflow : hidden; } 复制代码` 效果图如下： ![在这里插入图片描述](https://user-gold-cdn.xitu.io/2019/6/5/16b26cef18ab88b4?imageView2/0/w/1280/h/960/ignore-error/1) 好吧，效果又是一样的，不过本着假装自己认真的原则又上传了一下。。。不过这个多了一个 ` overflow: hidden` ，没有了 ` margin: 0 200px` 了呢。在这里是这样的， ` overflow: hidden` 是会触发BFC的，那么此时content区域就是一个BFC区域了，那么他不会被浮动的元素盖住，所以，它的content的宽度就是从不被left盖住的位置开始到被right盖住的区域之前这段宽度，不用再像方法一那样把margin给左右的值空出来这个区域，所以也形成了宽度自适应，搞定！不过它的缺点也显而易见，和方法一的一样：不能优先渲染content区域。

### 方法四：圣杯布局（重点） ###

* 我们终于等到了重头戏之一：圣杯布局。其实所谓圣杯布局只不过比起三栏布局多了一个需求：要求content区域优先渲染。那大家可能比较疑惑：为啥叫圣杯布局？因为它长得像圣杯啊，看下图： ![在这里插入图片描述](https://user-gold-cdn.xitu.io/2019/6/5/16b26cef18c2e7f0?imageView2/0/w/1280/h/960/ignore-error/1) 圣杯的两只手就像布局中的左边栏和右边栏一样，杯子主体就像content内容区域一样。具体的实现如下：
* html代码： ` < div class = "main clearfix" > < div class = "content" > content </ div > < aside class = "sidebar-left" > left </ aside > < aside class = "sidebar-right" > right </ aside > </ div > 复制代码`
* css代码： `.main { padding : 0 200px 0 150px ; } body { color : #fff ; font-size : 40px ; background-color : #666 ; font-family : Arial; text-align : center; }.content ,.sidebar-left ,.sidebar-right { float : left; position : relative; height : 400px ; line-height : 400px ; }.content { width : 100% ; background-color : #f5c531 ; }.sidebar-left { width : 150px ; background-color : #a0c263 ; margin-left : - 100% ; left : - 150px ; }.sidebar-right { background-color : #a0c263 ; width : 200px ; margin-right : - 200px ; /*margin-left: -200px;*/ /*right: -200px;*/ } 复制代码` 看下效果图： ![在这里插入图片描述](https://user-gold-cdn.xitu.io/2019/6/5/16b26cef1ccbbcda?imageView2/0/w/1280/h/960/ignore-error/1) 左边栏150px，右边栏200px，中间content自适应宽度，完美解决，这就是所说的圣杯布局，那么这个原理是什么呢？其实就是基于两条：1. 浮动 2. ` margin` 负值。 我们现在看代码： * 首先给他们的父级元素main一个左padding值和右padding值，留出来给左右边栏的地方。
* 给这三个元素浮动起来，然后直接给content元素宽度100%，宽度全都给content了，那left和right就会掉下去。
* 先给content宽度100%的时候，确实left和right是要掉下去的，就像下面这样： ![在这里插入图片描述](https://user-gold-cdn.xitu.io/2019/6/5/16b26cef45f9f81a?imageView2/0/w/1280/h/960/ignore-error/1) 那咋把left和right这两个烦人的玩意弄到他们的位置上呢？这里就要用到 ` margin` 的负值，结合浮动使用，那么我们知道浮动起来的几个元素会拍成一行内，而宽度不够才会掉下去，这里就是这种情况，但是如果我给left块一个 ` margin-left` 负值到一个界限，他就会回到上一行，因为它会向左走，左边没地方了，而且给的是 ` margin` 值，那么他就会贴到上一行，因为他们是一起浮动的，上一行同样能够接纳他（这里我这样想：比如说宽度足够的时候，那么我给它 ` margin-left` 负值，它不是一样要往左走吗？一样的道理，它在下一行往左走到没地方走的时候自然就回到了上一行了，因为他们是一起浮动的），这里我给了它 ` margin-left: -100%` ，所以它就到了这个位置 ![在这里插入图片描述](https://user-gold-cdn.xitu.io/2019/6/5/16b26cef5408c2bc?imageView2/0/w/1280/h/960/ignore-error/1) 因为 ` margin` 的百分比的值是父级盒子宽度的百分比，给了-100%，自然向左走一个父级的宽度，就到了这，那么这个位置不是我们想要的，我们要他在左边的空出来的地方待着，那很简单： ` left: -150px` 搞定： ![在这里插入图片描述](https://user-gold-cdn.xitu.io/2019/6/5/16b26cef467a0f42?imageView2/0/w/1280/h/960/ignore-error/1) 现在就差right没搞定了，这个其实很简单，只需要一句代码： ` margin-right: -200px` 搞定： ![在这里插入图片描述](https://user-gold-cdn.xitu.io/2019/6/5/16b26cef5d4b203a?imageView2/0/w/1280/h/960/ignore-error/1) 其实这里可以理解为：本来他就应该在上面，但是由于宽度不够，无奈被挤下来了，但是现在我只要给 ` margin-right` 一个负值（这个值大小要超过自身的宽度）他就会上去，继续被接纳，（这个值如果给的没有自身宽度大，那么就相当于不被接纳，还要被排斥下来）所以正好排在了空缺的地方。 注意看我给right的代码里面最后两句注释上的代码，如果用这两句，可以达到和上面同样的效果，原理也相同：先给一个 ` margin-left: -200px` （正好给200px的话就是正好贴在在父级盒子的最右边，因为正好刚刚能上来，如果再多给点值还会再往左走，但是这里不需要的，我只是在尽量去把这个原理阐述明白），被上面接纳，来到了它的父级盒子的贴着最右边： ![在这里插入图片描述](https://user-gold-cdn.xitu.io/2019/6/5/16b26cef7c05a572?imageView2/0/w/1280/h/960/ignore-error/1) 然后再给一个 ` right: -200px` 让它向右走200px，搞定！这个其实就和我给left盒子的方法原理一样。最终效果图： ![在这里插入图片描述](https://user-gold-cdn.xitu.io/2019/6/5/16b26cef7df26d61?imageView2/0/w/1280/h/960/ignore-error/1)

### 方法五：双飞翼布局（重点） ###

* 双飞翼布局其实是根据圣杯布局演化出来的一种布局。具体实现如下：
* html代码： ` < div class = "main clearfix" > < div class = "content-wrapper" > < div class = "content" > content </ div > </ div > < aside class = "sidebar-left" > left </ aside > < aside class = "sidebar-right" > right </ aside > </ div > 复制代码`
* css代码： ` body { color : #fff ; font-size : 40px ; font-family : Arial; background-color : #666 ; text-align : center; }.sidebar-left ,.sidebar-right { float : left; height : 400px ; line-height : 400px ; }.content-wrapper { width : 100% ; float : left; }.content { margin : 0 200px 0 150px ; background-color : #f5c531 ; height : 400px ; line-height : 400px ; }.sidebar-left { width : 150px ; background-color : #a0c263 ; margin-left : - 100% ; }.sidebar-right { background-color : #a0c263 ; width : 200px ; margin-left : - 200px ; } 复制代码` 老规矩，用图说话： ![在这里插入图片描述](https://user-gold-cdn.xitu.io/2019/6/5/16b26cef87cbf493?imageView2/0/w/1280/h/960/ignore-error/1) 其实我们可以看到，这里面和圣杯布局差别不是很大，我简单说一下里面的细节差距吧： * 这里面left直接给 ` margin-left: -100%` 就能到想要的位置，为什么呢？因为要注意一点，在这里，left父级盒子是宽度100%的，不再是圣杯布局里面的留出来左右padding值的父级自适应宽度的盒子，这里面使用content盒子的左右 ` margin` 值留出来的定宽，所以直接就能把left盒子定到想要的位置
* 那么同理，right盒子也是因为这样，所以直接 ` margin-left: -200px` 正好贴到父级盒子最右边，就能到想要的位置了。

### 方法六：Flex三栏布局 ###

* 利用Flex弹性盒子可以很轻松的完成三栏布局，并且能够达到content优先渲染的要求。
* html结构：

` < div class = "flex-box" > < div class = "flex-content flex-item" > 我是内容 </ div > < div class = "flex-left flex-item" > 我是左边栏 </ div > < div class = "flex-right flex-item" > 我是右边栏 </ div > </ div > 复制代码`

* css结构：

` body { color : #fff ; font-size : 30px ; text-align : center; }.flex-box { display : flex; line-height : 600px ; }.flex-left { width : 200px ; height : 600px ; background-color : lime; order : 0 ; }.flex-content { order : 1 ; width : 100% ; background-color : purple; }.flex-right { width : 200px ; height : 600px ; background-color : blue; order : 2 ; } 复制代码`

效果图就不贴了。。。太累了。这里主要利用flex布局中，给子元素添加的属性order，这个order属性意为在主轴方向的排列中显示的优先级，值越小，优先级越高，所以可以达到HTML结构中content最先渲染，却又能让其在中间部分显示的效果。

#### 总结 ####

就写这6种三栏布局的方式吧，都是很简单的。默默地喝掉最后一口咖啡，继续奋战。