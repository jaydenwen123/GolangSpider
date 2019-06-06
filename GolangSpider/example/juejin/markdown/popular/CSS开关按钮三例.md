# CSS开关按钮三例 #

我们将使用纯CSS打造一些切换开关并使其拥有类似于checkbox的用户体验。

很多时候我们都需要用户通过勾选/取消checkbox来表明他们对一些问题的答案。我们设置了一个标签，一个checkbox，并在提交表单后获取checkbox值，以查看用户是否已经选中或取消选中该checkbox。我们都知道默认的的checkbox长啥样，而且还不能通过纯CSS的方式来设置checkbox的样式。这种元素的样式由每个浏览器引擎单独管理（每个浏览器下面checkbox的样式都可能不一样）。于是，有一个更统一的界面岂不是会更好？

不要急！一个小小的CSS技巧可以帮助我们解决这个问题。通过将:checkded, :before和:after伪类结合到我们的checkbox上，我们可以实现一些漂亮并拥有平滑过渡效果的切换型开关。没有黑魔法...仅仅是CSS的魅力。下面让我们开始吧。

## HTML ##

需要用到的HTML并不是我们之前没见过的，也就是一个标准的checkbox结合一个label。我们用一个div将checkox和label包裹起来，并给这个div添加了一个switch的样式类。

label的样式则会使用input + label选择器来定位，那样label就不需要自己的样式类名了。现在让我们来看下下面的HTML结构：

` <div class = "switch" > <input id= "cmn-toggle-1" class = "cmn-toggle cmn-toggle-round" type = "checkbox" > <label for = "cmn-toggle-1" ></label> </div> <div class = "switch" > <input id= "cmn-toggle-4" class = "cmn-toggle cmn-toggle-round-flat" type = "checkbox" > <label for = "cmn-toggle-4" ></label> </div> <div class = "switch" > <input id= "cmn-toggle-7" class = "cmn-toggle cmn-toggle-yes-no" type = "checkbox" > <label for = "cmn-toggle-7" data-on= "Yes" data-off= "No" ></label> </div> 复制代码`

这里没什么特别的。对于CSS,我们希望真实的checkbox被隐藏在屏幕和视线之外。基本上所有的样式都被加在label上。这样做很方便，因为点击label实际上会勾选/取消勾选checkbox。我们将用下面的CSS来实现切换开关：

`.cmn-toggle { position : absolute; margin-left : - 9999px ; visibility : hidden; }.cmn-toggle + label { display : block; position : relative; cursor : pointer; outline : none; user-select : none; } 复制代码`

## 样式一 ##

![clipboard.png](https://user-gold-cdn.xitu.io/2019/6/4/16b2015dbdfea4f1?imageView2/0/w/1280/h/960/ignore-error/1)

此时label充当容器的角色，并拥有宽和高。我们还给它设置了一个背景颜色来模拟我们的切换开关的边界。:before元素模拟开关内部的浅灰色区域（开关打开时背景颜色会过渡到绿色）。:after元素才是真正的圆形开关，它的层级高于一切，在点击时的时候它将从左滑动到右。我们将给:after元素添加一个box-shadow使它看起来更加立体。当input接受:checked伪类时，我们将平滑的改变:before元素的背景颜色和:after元素的位置。CSS如下：

` input.cmn-toggle-round + label { padding : 2px ; width : 120px ; height : 60px ; background-color : #dddddd ; border-radius : 60px ; } input.cmn-toggle-round + label :before , input.cmn-toggle-round + label :after { display : block; position : absolute; top : 1px ; left : 1px ; bottom : 1px ; content : "" ; } input.cmn-toggle-round + label :before { right : 1px ; background-color : #f1f1f1 ; border-radius : 60px ; transition : background 0.4s ; } input.cmn-toggle-round + label :after { width : 58px ; background-color : #fff ; border-radius : 100% ; box-shadow : 0 2px 5px rgba (0, 0, 0, 0.3); transition : margin 0.4s ; } input.cmn-toggle-round :checked + label :before { background-color : #8ce196 ; } input.cmn-toggle-round :checked + label :after { margin-left : 60px ; } 复制代码`

## 样式二 ##

![clipboard.png](https://user-gold-cdn.xitu.io/2019/6/4/16b2015dc144c796?imageView2/0/w/1280/h/960/ignore-error/1)

接下来的这个例子和上面的例子非常相似，主要的区别在于它的外观表现。它符合现代网站平滑扁平化趋势，但是就功能而言和例1一样。下面的CSS仅仅改变了toggle的表现风格，其他的都是一样的。

` input.cmn-toggle-round-flat + label { padding : 2px ; width : 120px ; height : 60px ; background-color : #dddddd ; border-radius : 60px ; transition : background 0.4s ; } input.cmn-toggle-round-flat + label :before , input.cmn-toggle-round-flat + label :after { display : block; position : absolute; content : "" ; } input.cmn-toggle-round-flat + label :before { top : 2px ; left : 2px ; bottom : 2px ; right : 2px ; background-color : #fff ; border-radius : 60px ; transition : background 0.4s ; } input.cmn-toggle-round-flat + label :after { top : 4px ; left : 4px ; bottom : 4px ; width : 52px ; background-color : #dddddd ; border-radius : 52px ; transition : margin 0.4s , background 0.4s ; } input.cmn-toggle-round-flat :checked + label { background-color : #8ce196 ; } input.cmn-toggle-round-flat :checked + label :after { margin-left : 60px ; background-color : #8ce196 ; } 复制代码`

## 样式三 ##

![clipboard.png](https://user-gold-cdn.xitu.io/2019/6/4/16b2015db5732c69?imageView2/0/w/1280/h/960/ignore-error/1)

现在，我们要做一点不一样的事了。我们将会创建一个翻转风格的switcher开关。默认视图为灰色，并显示“No”（或任何表示未选中的内容），勾选后的视图则为绿色，并显示“Yes”。当点击label时，swithcer会沿Y轴翻转180度。我们将使用“data-attributes”来填充未选中/已选中时内容。这些“data-attributes”在HTML中由“data-on”和“data-off”指定，他们将分别填充到:after和:before两个伪元素中。请注意:after伪元素中的backface-visiibility属性，由于起点是-180度，通过这个属性可以隐藏背面的内容。

` input.cmn-toggle-yes-no + label { padding : 2px ; width : 120px ; height : 60px ; } input.cmn-toggle-yes-no + label :before , input.cmn-toggle-yes-no + label :after { display : block; position : absolute; top : 0 ; left : 0 ; bottom : 0 ; right : 0 ; color : #fff ; font-family : "Roboto Slab" , serif; font-size : 20px ; text-align : center; line-height : 60px ; } input.cmn-toggle-yes-no + label :before { background-color : #dddddd ; content : attr (data-off); transition : transform 0.5s ; backface-visibility : hidden; } input.cmn-toggle-yes-no + label :after { background-color : #8ce196 ; content : attr (data-on); transition : transform 0.5s ; transform : rotateY (180deg); backface-visibility : hidden; } input.cmn-toggle-yes-no :checked + label :before { transform : rotateY (180deg); } input.cmn-toggle-yes-no :checked + label :after { transform : rotateY (0); } 复制代码`

## 浏览器兼容性 ##

上面的这些在浏览器兼容方面的要求是，IE8及以下的浏览器不能识别:checked伪类，因此你需要检测浏览器，如果是老旧的IE，则直接回退到原始的checkbox，css transitions 属性不支持IE9及以下浏览器，但这仅仅会影响切换过程中的过渡部分，除此之外没有其他毛病能够正常工作。

## 总结 ##

这是一个关于一些很好的CSS切换开关示例！这种技术使得一切完全复合语义，不会增加任何疯狂的标记，并且用纯CSS就可以完成。当然，你需要注意浏览器兼容性情况，但是你可以使用条件样式来兼容旧版浏览器，使用上面提到的例子，并不会产生什么不足之处。