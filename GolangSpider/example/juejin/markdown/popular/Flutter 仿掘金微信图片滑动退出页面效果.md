# Flutter 仿掘金微信图片滑动退出页面效果 #

[extended_image]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Ffluttercandies%2Fextended_image ) 相关文章

* [Flutter 什么功能都有的Image]( https://juejin.im/post/5c867112f265da2dd427a340 )
* [Flutter 可以缩放拖拽的图片]( https://juejin.im/post/5ca758916fb9a05e1c4d01bb )
* [Flutter 仿掘金微信图片滑动退出页面效果]( https://juejin.im/post/5cf62ab0e51d45776031afb2 )

( https://link.juejin.im?target=https%3A%2F%2Fpub.dartlang.org%2Fpackages%2Fextended_image )

![pub package](https://img.shields.io/pub/v/extended_image.svg) ( https://link.juejin.im?target=https%3A%2F%2Fpub.dartlang.org%2Fpackages%2Fextended_image )

这个需求在做 [extended_image]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Ffluttercandies%2Fextended_image ) 的时候就有上帝客户提过了，一直都没有时间去考虑实现。最近思考了一下，把效果给实现了。

![](https://user-gold-cdn.xitu.io/2019/6/5/16b2580722f2a715?imageslim)

### 首先把你的页面用ExtendedImageSlidePage包一下 ###

` var page = ExtendedImageSlidePage( child: PicSwiper( index, listSourceRepository .map<PicSwiperItem>( (f) => PicSwiperItem(f.imageUrl, des: f.title)) .toList(), ), //pageGestureAxis: PageGestureAxis.horizontal, ); 复制代码`

ExtendedImageGesturePage的参数

+----------------------------+-----------------------------------------------------------------------+-----------------------------------+
|         PARAMETER          |                              DESCRIPTION                              |              DEFAULT              |
+----------------------------+-----------------------------------------------------------------------+-----------------------------------+
| child                      | 需要包裹的页面                                                        | -                                 |
| slidePageBackgroundHandler | 在滑动页面的时候根据Offset自定义整个页面的背景色                      | defaultSlidePageBackgroundHandler |
| slideScaleHandler          | 在滑动页面的时候根据Offset自定义整个页面的缩放值                      | defaultSlideScaleHandler          |
| slideEndHandler            | 滑动页面结束的时候计算是否需要pop页面                                 | defaultSlideEndHandler            |
| slideAxis                  | 滑动页面的方向（both,horizontal,vertical）,掘金是vertical，微信是Both | both                              |
| resetPageDuration          | 滑动结束，如果不pop页面，整个页面回弹动画的时间                       | milliseconds: 500                 |
| slideType                  | 滑动整个页面还是只是图片(wholePage/onlyImage)                         | SlideType.onlyImage               |
+----------------------------+-----------------------------------------------------------------------+-----------------------------------+

下面是默认实现，你也可以根据你的喜好，来定义属于自己方式

` Color defaultSlidePageBackgroundHandler( {Offset offset, Size pageSize, Color color, SlideAxis pageGestureAxis}) { double opacity = 0.0 ; if (pageGestureAxis == SlideAxis.both) { opacity = offset.distance / (Offset(pageSize.width, pageSize.height).distance / 2.0 ); } else if (pageGestureAxis == SlideAxis.horizontal) { opacity = offset.dx.abs() / (pageSize.width / 2.0 ); } else if (pageGestureAxis == SlideAxis.vertical) { opacity = offset.dy.abs() / (pageSize.height / 2.0 ); } return color.withOpacity(min( 1.0 , max( 1.0 - opacity, 0.0 ))); } bool defaultSlideEndHandler( {Offset offset, Size pageSize, SlideAxis pageGestureAxis}) { if (pageGestureAxis == SlideAxis.both) { return offset.distance > Offset(pageSize.width, pageSize.height).distance / 3.5 ; } else if (pageGestureAxis == SlideAxis.horizontal) { return offset.dx.abs() > pageSize.width / 3.5 ; } else if (pageGestureAxis == SlideAxis.vertical) { return offset.dy.abs() > pageSize.height / 3.5 ; } return true ; } double defaultSlideScaleHandler( {Offset offset, Size pageSize, SlideAxis pageGestureAxis}) { double scale = 0.0 ; if (pageGestureAxis == SlideAxis.both) { scale = offset.distance / Offset(pageSize.width, pageSize.height).distance; } else if (pageGestureAxis == SlideAxis.horizontal) { scale = offset.dx.abs() / (pageSize.width / 2.0 ); } else if (pageGestureAxis == SlideAxis.vertical) { scale = offset.dy.abs() / (pageSize.height / 2.0 ); } return max( 1.0 - scale, 0.8 ); } 复制代码`

### 确保你的页面是透明背景的 ###

如果你设置 slideType =SlideType.onlyImage, 请确保的你页面是透明的，毕竟没法操控你页面上的颜色

### Push一个透明的页面 ###

这里我把官方的MaterialPageRoute 和CupertinoPageRoute拷贝出来了， 改为TransparentMaterialPageRoute/TransparentCupertinoPageRoute，因为它们的opaque不能设置为false

` Navigator.push( context, Platform.isAndroid ? TransparentMaterialPageRoute(builder: (_) => page) : TransparentCupertinoPageRoute(builder: (_) => page), ); 复制代码`

嗯应该还算使用简单吧？群里的小伙伴吐槽表情包太多，不让放，蓝瘦香菇。

### 实现中的一些坑 ###

#### 1.手势跟缩放拖拽以及PageView之前的关系和冲突 ####

开始我的思路是想在ExtendedImageSlidePage 注册手势监听事件，然后ExtendedImageGesture里面当条件满足(达到边界/无法拖拽)的时候通知 ExtendedImageSlidePage 开始滑动页面手势了，可以阻止ExtendedImageSlidePage的child的hittest。

但是在实际中发现，在ExtendedImageGesture手势未完成之前（手指抬起）,ExtendedImageSlidePage 也是获取不到任何手势，而且IgnorePointer 也是不会生效的

后面干脆直接把手势接收都放ExtendedImageGesture里面了，直接通知ExtendedImageSlidePage进行translate和scale

#### 2.透明页面 ####

TransparentMaterialPageRoute/TransparentCupertinoPageRoute 因为需要整个页面是透明的，所以重写了官方的。

但是在pop页面的时候还是有不满意的地方，比如ios上面有个从左到右Shadow，安卓上面整个页面也有Shadow。

通过修改官方源码，去掉了这些效果，感兴趣的小伙伴可以查看 [extended_image_slide_page_route.dart]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Ffluttercandies%2Fextended_image%2Fblob%2Fmaster%2Flib%2Fsrc%2Fgesture%2Fextended_image_slide_page_route.dart ) （最近拉面批评代码上太多，差评，那个啥代码就不贴了 ）

### 缺陷 ###

设置SlideType.onlyImage 因为没法操控页面上面的其他内容，所以只能靠改变背景色，暂时没想到更好的方法来给页面上的其他东西做动画

### 关于extended_image的readme ###

最近重新整理了一下readme，因为大家老是吐槽不容易看，希望新的readme能帮助大家更好地使用这个组件，感谢 [财经龙大佬]( https://juejin.im/user/593f4df9ac502e006b56268a ) 百忙当中帮忙格式readme，懒惰的程序猿，readme都要大佬帮忙弄，羞愧。。。

最后放上 [extended_image]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Ffluttercandies%2Fextended_image ) ，如果你有什么不明白或者对这个方案有什么改进的地方，请告诉我，欢迎加入 [Flutter Candies]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Ffluttercandies ) ，一起生产可爱的Flutter 小糖果(QQ群:181398081)

最最后放上 [Flutter Candies]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Ffluttercandies ) 全家桶，真香。

![](https://user-gold-cdn.xitu.io/2019/5/29/16b02e0775f4af97?imageView2/0/w/1280/h/960/ignore-error/1)