# CSS：你真的会用 z-index 吗？ #

![](https://user-gold-cdn.xitu.io/2019/5/29/16b02fb552c03f44?imageView2/0/w/1280/h/960/ignore-error/1)

# 你真的会用 z-index 么？ #

` 如果你的 css 里面存在大量这样的代码： z-index：66、666、999、9999 可能你还不太理解 z-index 复制代码`

![](https://user-gold-cdn.xitu.io/2019/5/29/16b02fc71f0e0684?imageView2/0/w/1280/h/960/ignore-error/1)

![](https://user-gold-cdn.xitu.io/2019/5/29/16b02fcc1277c0e6?imageView2/0/w/1280/h/960/ignore-error/1)

# HTML 元素是处于三维空间中 #

所有的盒模型元素都处于三维坐标系中，除了我们常用的横坐标和纵坐标，盒模型元素还可以沿着“z 轴”层叠摆放，当他们相互覆盖时，z 轴顺序就变得十分重要。

但“z 轴”顺序，不完全由 z-index 决定，在层叠比较复杂的 HTML 元素上使用 z-index 时，结果可能让人觉得困惑，甚至不可思议。这是由复杂的元素排布规则导致的。

![](https://user-gold-cdn.xitu.io/2019/5/29/16b0300df37b3bda?imageView2/0/w/1280/h/960/ignore-error/1)

## 不含 z-index 元素如何堆叠？ ##

> 
> 
> 
> 当没有元素包含z-index属性时，元素按照如下顺序堆叠（从底到顶顺序）：
> 
> * 根元素（ ` <html>` ）的背景和边界；
> * 位于普通流中的后代“无定位块级元素”，按它们在HTML中的出现顺序堆叠；
> * 后代中的“定位元素”，按它们在HTML中的出现顺序堆叠；
> 
> 
> 注意：普通流中的“无定位块级元素”始终先于“定位元素”渲染，并出现在“定位元素”下层，即便它们在HTML结构中出现的位置晚于定位元素也是如此。
> 
> 

![](https://user-gold-cdn.xitu.io/2019/5/29/16b0304af2806b4d?imageView2/0/w/1280/h/960/ignore-error/1)

` <!DOCTYPE html> <html> <head><style type = "text/css" > b { font-family: sans-serif; } div { padding: 10px; border: 1px dashed; text-align: center; } .static { position: static; height: 80px; background-color: #ffc; border-color: #996; } .absolute { position: absolute; width: 150px; height: 350px; background-color: #fdd; border-color: #900; opacity: 0.7; } .relative { position: relative; height: 80px; background-color: #cfc; border-color: #696; opacity: 0.7; } #abs1 { top: 10px; left: 10px; } #rel1 { top: 30px; margin: 0px 50px 0px 50px; } #rel2 { top: 15px; left: 20px; margin: 0px 50px 0px 50px; } #abs2 { top: 10px; right: 10px; } #sta1 { background-color: #ffc; margin: 0px 50px 0px 50px; } </style></head> <body> <div id= "abs1" class= "absolute" > <b>DIV #1</b><br />position: absolute;</div> <div id= "rel1" class= "relative" > <b>DIV #2</b><br />position: relative;</div> <div id= "rel2" class= "relative" > <b>DIV #3</b><br />position: relative;</div> <div id= "abs2" class= "absolute" > <b>DIV #4</b><br />position: absolute;</div> <div id= "sta1" class= "static" > <b>DIV #5</b><br />position: static;</div> </body></html> 复制代码`

![](https://user-gold-cdn.xitu.io/2019/5/29/16b0307946e768e6?imageView2/0/w/1280/h/960/ignore-error/1)

## float 如何影响堆叠？ ##

> 
> 
> 
> 对于浮动的块元素来说，层叠顺序变得有些不同。浮动块元素被放置于非定位块元素与定位块元素之间：
> 
> * 根元素（ ` <html>` ）的背景和边界；
> * 位于普通流中的后代“无定位块级元素”，按它们在HTML中的出现顺序堆叠；
> * 浮动块元素；<<<<
> * 位于普通流中的后代“无定位行内元素”；
> * 后代中的“定位元素”，按它们在HTML中的出现顺序堆叠；
> 

![](https://user-gold-cdn.xitu.io/2019/5/29/16b030b9d909dd16?imageView2/0/w/1280/h/960/ignore-error/1)

` <!DOCTYPE html><html><head> <meta charset= "UTF-8" > <title>Stacking and float </title> <style type = "text/css" > div { padding: 10px; text-align: center; } b { font-family: sans-serif; } #abs1 { position: absolute; width: 150px; height: 200px; top: 20px; right: 160px; border: 1px dashed #900; background-color: #fdd; } #sta1 { height: 100px; border: 1px dashed #996; background-color: #ffc; margin: 0px 10px 0px 10px; text-align: left; } #flo1 { margin: 0px 10px 0px 20px; float : left; width: 150px; height: 200px; border: 1px dashed #090; background-color: #cfc; } #flo2 { margin: 0px 20px 0px 10px; float : right; width: 150px; height: 200px; border: 1px dashed #090; background-color: #cfc; } #abs2 { position: absolute; width: 300px; height: 100px; top: 150px; left: 100px; border: 1px dashed #990; background-color: #fdd; } </style></head> <body> <div id= "abs1" > <b>DIV #1</b><br />position: absolute;</div> <div id= "flo1" > <b>DIV #2</b><br />float: left;</div> <div id= "flo2" > <b>DIV #3</b><br />float: right;</div> <br/> <div id= "sta1" > <b>DIV #4</b><br />no positioning</div> <div id= "abs2" > <b>DIV #5</b><br />position: absolute;</div> </body> </html> 复制代码`

![](https://user-gold-cdn.xitu.io/2019/5/29/16b030bbb02e863b?imageView2/0/w/1280/h/960/ignore-error/1)

## z-index 如何影响堆叠？ ##

> 
> 
> 
> z-index 属性指定了一个具有定位属性的元素及其子代元素的 z-order。 当元素之间重叠的时候，z-order
> 决定哪一个元素覆盖在其余元素的上方显示。 通常来说 z-index 较大的元素会覆盖较小的一个。
> 
> 
> 
> 对于一个已经定位的元素（即position属性值不是static的元素），z-index 属性指定：
> 
> * 元素在当前堆叠上下文中的堆叠层级。
> * 元素是否创建一个新的本地堆叠上下文。
> 

![](https://user-gold-cdn.xitu.io/2019/5/29/16b030e1886ee0e6?imageView2/0/w/1280/h/960/ignore-error/1)

` <!DOCTYPE html> <html> <head><style type = "text/css" > div { opacity: 0.7; font: 12px Arial; } span.bold { font-weight: bold; } #normdiv { z-index: 8; height: 70px; border: 1px dashed #999966; background-color: #ffffcc; margin: 0px 50px 0px 50px; text-align: center; } #reldiv1 { z-index: 3; height: 100px; position: relative; top: 30px; border: 1px dashed #669966; background-color: #ccffcc; margin: 0px 50px 0px 50px; text-align: center; } #reldiv2 { z-index: 2; height: 100px; position: relative; top: 15px; left: 20px; border: 1px dashed #669966; background-color: #ccffcc; margin: 0px 50px 0px 50px; text-align: center; } #absdiv1 { z-index: 5; position: absolute; width: 150px; height: 350px; top: 10px; left: 10px; border: 1px dashed #990000; background-color: #ffdddd; text-align: center; } #absdiv2 { z-index: 1; position: absolute; width: 150px; height: 350px; top: 10px; right: 10px; border: 1px dashed #990000; background-color: #ffdddd; text-align: center; } </style></head> <body> <br /><br /> <div id= "absdiv1" > <br /><span class= "bold" >DIV #1</span> <br />position: absolute; <br />z-index: 5; </div> <div id= "reldiv1" > <br /><span class= "bold" >DIV #2</span> <br />position: relative; <br />z-index: 3; </div> <div id= "reldiv2" > <br /><span class= "bold" >DIV #3</span> <br />position: relative; <br />z-index: 2; </div> <div id= "absdiv2" > <br /><span class= "bold" >DIV #4</span> <br />position: absolute; <br />z-index: 1; </div> <div id= "normdiv" > <br /><span class= "bold" >DIV #5</span> <br />no positioning <br />z-index: 8; </div> </body></html> 复制代码`

![](https://user-gold-cdn.xitu.io/2019/5/29/16b030e9a2841c16?imageView2/0/w/1280/h/960/ignore-error/1)

# 堆叠上下文（Stacking Context） #

## 什么是堆叠上下文？ ##

层叠上下文（Stacking Context）是HTML元素的三维概念，这些HTML元素在一条假想的相对于面向（电脑屏幕的）视窗或者网页的用户的z轴上延伸，HTML元素依据其自身属性按照优先级顺序占用层叠上下文的空间。

在层叠上下文中，其子元素的 z-index 值只在父级层叠上下文中有意义。子级层叠上下文被自动视为父级层叠上下文的一个独立单元。

* 层叠上下文可以包含在其他层叠上下文中，并且一起创建一个有层级的层叠上下文。
* 每个层叠上下文完全独立于它的兄弟元素：当处理层叠时只考虑子元素。
* 每个层叠上下文是自包含的：当元素的内容发生层叠后，整个该元素将会 在父层叠上下文中 按顺序进行层叠。

## 如何形成堆叠上下文？ ##

* 根元素 (HTML)
* 定位元素（relative、absolute），并且 z-index 不为 auto；
* opacity 小于 1 时；
* transform 不为 none 时；
* z-index 不为 auto 的 flex-item；

注：层叠上下文的层级是 HTML 元素层级的一个层级，因为只有某些元素才会创建层叠上下文。可以这样说，没有创建自己的层叠上下文的元素 将被父层叠上下文包含。

![](https://user-gold-cdn.xitu.io/2019/5/29/16b031109c66f488?imageView2/0/w/1280/h/960/ignore-error/1)

` <!DOCTYPE html> <html> <head><style type = "text/css" > * { margin: 0; } html { padding: 20px; font: 12px/20px Arial, sans-serif; } div { opacity: 0.7; position: relative; } h1 { font: inherit; font-weight: bold; } #div1, #div2 { border: 1px dashed #696; padding: 10px; background-color: #cfc; } #div1 { z-index: 5; margin-bottom: 190px; } #div2 { z-index: 2; } #div3 { z-index: 4; opacity: 1; position: absolute; top: 40px; left: 180px; width: 330px; border: 1px dashed #900; background-color: #fdd; padding: 40px 20px 20px; } #div4, #div5 { border: 1px dashed #996; background-color: #ffc; } #div4 { z-index: 6; margin-bottom: 15px; padding: 25px 10px 5px; } #div5 { z-index: 1; margin-top: 15px; padding: 5px 10px; } #div6 { z-index: 3; position: absolute; top: 20px; left: 180px; width: 150px; height: 125px; border: 1px dashed #009; padding-top: 125px; background-color: #ddf; text-align: center; } </style></head> <body> <br /><br /> <div id= "div1" > <h1>Division Element #1</h1> <code>position: relative;<br/> z-index: 5; </div> <div id= "div2" > <h1>Division Element #2</h1> <code>position: relative;<br/> z-index: 2; </div> <div id= "div3" > <div id= "div4" > <h1>Division Element #4</h1> <code>position: relative;<br/> z-index: 6; </div> <h1>Division Element #3</h1> <code>position: absolute;<br/> z-index: 4; <div id= "div5" > <h1>Division Element #5</h1> <code>position: relative;<br/> z-index: 1; </div> <div id= "div6" > <h1>Division Element #6</h1> <code>position: absolute;<br/> z-index: 3; </div> </div> </body></html> 复制代码`

![](https://user-gold-cdn.xitu.io/2019/5/29/16b03117461de04d?imageView2/0/w/1280/h/960/ignore-error/1)

## 堆叠上下文如何影响堆叠？ ##

* 元素的 background 和 borders;
* 拥有负堆叠层级（negative stack levels）的子层叠上下文（child stacking contexts）
* 在文档流中的（in-flow），非行内级的（non-inline-level），非定位（non-positioned）的后代元素
* 非定位的浮动元素
* 在文档流中的（in-flow），行内级的（inline-level），非定位（non-positioned）的后代元素，包括行内块级元素（inline blocks）和行内表格元素（inline tables）
* 堆叠层级为 0 的子堆叠上下文（child stacking contexts）和堆叠层级为 0 的定位的后代元素
* 堆叠层级为正的子堆叠上下文

上述关于层次的绘制规则递归地适用于任何层叠上下文。

![](https://user-gold-cdn.xitu.io/2019/5/29/16b031303258e491?imageView2/0/w/1280/h/960/ignore-error/1)

` <!DOCTYPE html> <html> <head><style type = "text/css" > div { font: 12px Arial; } span.bold { font-weight: bold; } div.lev1 { width: 250px; height: 70px; position: relative; border: 2px outset #669966; background-color: #ccffcc; padding-left: 5px; } #container1 { z-index: 1; position: absolute; top: 30px; left: 75px; } div.lev2 { opacity: 0.9; width: 200px; height: 60px; position: relative; border: 2px outset #990000; background-color: #ffdddd; padding-left: 5px; } #container2 { z-index: 1; position: absolute; top: 20px; left: 110px; } div.lev3 { z-index: 10; width: 100px; position: relative; border: 2px outset #000099; background-color: #ddddff; padding-left: 5px; } </style></head> <body> <br /> <div class= "lev1" > <span class= "bold" >LEVEL #1</span> <div id= "container1" > <div class= "lev2" > <br/><span class= "bold" >LEVEL #2</span> <br/>z-index: 1; <div id= "container2" > <div class= "lev3" ><span class= "bold" >LEVEL #3</span></div> <div class= "lev3" ><span class= "bold" >LEVEL #3</span></div> <div class= "lev3" ><span class= "bold" >LEVEL #3</span></div> <div class= "lev3" ><span class= "bold" >LEVEL #3</span></div> <div class= "lev3" ><span class= "bold" >LEVEL #3</span></div> <div class= "lev3" ><span class= "bold" >LEVEL #3</span></div> <div class= "lev3" ><span class= "bold" >LEVEL #3</span></div> <div class= "lev3" ><span class= "bold" >LEVEL #3</span></div> <div class= "lev3" ><span class= "bold" >LEVEL #3</span></div> <div class= "lev3" ><span class= "bold" >LEVEL #3</span></div> <div class= "lev3" ><span class= "bold" >LEVEL #3</span></div> </div> </div> <div class= "lev2" > <br/><span class= "bold" >LEVEL #2</span> <br/>z-index: 1; </div> </div> </div> <div class= "lev1" > <span class= "bold" >LEVEL #1</span> </div> <div class= "lev1" > <span class= "bold" >LEVEL #1</span> </div> <div class= "lev1" > <span class= "bold" >LEVEL #1</span> </div> </body></html> 复制代码`

![](https://user-gold-cdn.xitu.io/2019/5/29/16b03135bd818b83?imageView2/0/w/1280/h/960/ignore-error/1)

# 最佳实践（不犯二准则） #

> 
> 
> 
> 对于非浮层元素，避免设置 z-index 值，z-index 值没有任何道理需要超过2，原因是：
> 
> * 定位元素一旦设置了 z-index 值，就从普通定位元素变成了层叠上下文元素，相互间的层叠顺序就发生了根本的变化，很容易出现设置了巨大的
> z-index 值也无法覆盖其他元素的问题。
> * 避免 z-index “一山比一山高”的样式混乱问题。此问题多发生在多人协作以及后期维护的时候。例如，A小图标定位，习惯性写了个
> z-index: 9；B一看，自己原来的实现被覆盖了，立马写了个 z-index: 99；结果比弹框组件层级还高，立马弹框组件来一个
> z-index: 99999……显然，最后项目的 z-index 层级管理就是一团糟。
> 
> 
> CSS世界
> 
> 

# 考核（荔枝FM面试题） #

写出6个div元素的堆叠顺序，最上面的在第一个位置，例如: .one .two .three .four .five .six（z-index）；

html:

` <div class= "one" > <div class= "two" ></div> <div class= "three" ></div> </div> <div class= "four" > <div class= "five" ></div> <div class= "six" ></div> </div> 复制代码`

css:

`.one { position: relative; z-index: 2; .two { z-index: 6; } .three { position: absolute; z-index: 5; } } .four { position: absolute; z-index: 1; .five {} .six { position: absolute; top: 0; left: 0; z-index: -1; } } 复制代码`

答案：

`.three .two .one .five .six .four 复制代码`

参考：

> 
> 
> 
> 《CSS世界》
> 
> 
> 
> Understanding CSS z-index： [developer.mozilla.org/zh-CN/docs/…](
> https://link.juejin.im?target=https%3A%2F%2Fdeveloper.mozilla.org%2Fzh-CN%2Fdocs%2FWeb%2FGuide%2FCSS%2FUnderstanding_z_index
> )
> 
> 
> 
> CSS2.2，Layered presentation： [www.w3.org/TR/CSS22/vi…](
> https://link.juejin.im?target=https%3A%2F%2Fwww.w3.org%2FTR%2FCSS22%2Fvisuren.html%23layers
> )
> 
> 
> 
> z-index： [developer.mozilla.org/en-US/docs/…](
> https://link.juejin.im?target=https%3A%2F%2Fdeveloper.mozilla.org%2Fen-US%2Fdocs%2FWeb%2FCSS%2Fz-index
> )
> 
> 
> 
> Elaborate description of Stacking Contexts： [www.w3.org/TR/CSS22/zi…](
> https://link.juejin.im?target=https%3A%2F%2Fwww.w3.org%2FTR%2FCSS22%2Fzindex.html
> )
> 
> 
> 
> How z-index Works： [bitsofco.de/how-z-index…](
> https://link.juejin.im?target=https%3A%2F%2Fbitsofco.de%2Fhow-z-index-works%2F
> )
> 
> 
> 
> What No One Told You About Z-Index： [philipwalton.com/articles/wh…](
> https://link.juejin.im?target=https%3A%2F%2Fphilipwalton.com%2Farticles%2Fwhat-no-one-told-you-about-z-index%2F
> )
> 
> 
> 
> What You May Not Know About the Z-Index Property： [webdesign.tutsplus.com/articles/wh…](
> https://link.juejin.im?target=https%3A%2F%2Fwebdesign.tutsplus.com%2Farticles%2Fwhat-you-may-not-know-about-the-z-index-property--webdesign-16892%3F_ga%3D2.155470971.1526956850.1558937624-731871869.1558937624
> )
> 
> 
> 
> Z-Index And The CSS Stack: Which Element Displays First?： [vanseodesign.com/css/css-sta…](
> https://link.juejin.im?target=http%3A%2F%2Fvanseodesign.com%2Fcss%2Fcss-stack-z-index%2F
> )
> 
> ![](https://user-gold-cdn.xitu.io/2019/5/27/16af831597dc6484?imageView2/0/w/1280/h/960/ignore-error/1)
> 社区以及公众号发布的文章，100%保证是我们的原创文章，如果有错误，欢迎大家指正。
> 
> 
> 
> 

> 
> 
> 
> 文章首发在WebJ2EE公众号上，欢迎大家关注一波，让我们大家一起学前端~~~
> 
> 

![](https://user-gold-cdn.xitu.io/2019/5/22/16aded0040c28b43?imageView2/0/w/1280/h/960/ignore-error/1)

> 
> 
> 
> 再来一波号外，我们成立WebJ2EE公众号前端吹水群，大家不管是看文章还是在工作中前端方面有任何问题，我们都可以在群内互相探讨，希望能够用我们的经验帮更多的小伙伴解决工作和学习上的困惑,欢迎加入。
> 
> 
> 

![](https://user-gold-cdn.xitu.io/2019/5/30/16b07838d4ca6b30?imageView2/0/w/1280/h/960/ignore-error/1)