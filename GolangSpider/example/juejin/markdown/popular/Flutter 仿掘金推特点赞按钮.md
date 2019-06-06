# Flutter 仿掘金推特点赞按钮 #

( https://link.juejin.im?target=https%3A%2F%2Fpub.dartlang.org%2Fpackages%2Flike_button )

![pubpackage](https://img.shields.io/pub/v/like_button.svg) ( https://link.juejin.im?target=https%3A%2F%2Fpub.dartlang.org%2Fpackages%2Flike_button )

做这个按钮的起因是，产品突然说要改一下点赞效果，像下面这样。

![](https://user-gold-cdn.xitu.io/2019/5/29/16b02a3ab259e538?imageslim)

但是我心中的审美不是这样的，不可以，坚决不行。 想起来 [群花拉面]( https://juejin.im/user/5b96160b5188255c5b5c1bab ) 写过一个 [用Flutter实现一个仿Twitter的点赞效果]( https://juejin.im/post/5bf01b7d51882516fa638069 ).

看了下代码， 站在大佬们的肩膀 [jd-alexander]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fjd-alexander%2FLikeButton ) and [吉原拉面]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fyumi0629%2FFlutterUI%2Ftree%2Fmaster%2Flib%2Flikebutton ) ，把代码魔改了一下，开源真香，并且添加了点赞数字动画。

把效果给产品看了一下

![](https://user-gold-cdn.xitu.io/2019/5/30/16b087dfeb9634e0?imageslim)

直接让产品改变主意，这个效果好，全平台都改。。。

![](https://user-gold-cdn.xitu.io/2019/5/29/16b029a4179e6435?imageView2/0/w/1280/h/960/ignore-error/1)

## 怎么使用 ##

如果什么都不设置，默认是Icons.favorite的效果

` LikeButton(), 复制代码`

当然你也可以自定义一些东西

` LikeButton( size: buttonSize, circleColor: CircleColor(start: Color( 0xff00ddff ), end: Color( 0xff0099cc )), bubblesColor: BubblesColor( dotPrimaryColor: Color( 0xff33b5e5 ), dotSecondaryColor: Color( 0xff0099cc ), ), likeBuilder: ( bool isLiked) { return Icon( Icons.home, color: isLiked ? Colors.deepPurpleAccent : Colors.grey, size: buttonSize, ); }, likeCount: 665 , countBuilder: ( int count, bool isLiked, String text) { var color = isLiked ? Colors.deepPurpleAccent : Colors.grey; Widget result; if (count == 0 ) { result = Text( "love" , style: TextStyle(color: color), ); } else result = Text( text, style: TextStyle(color: color), ); return result; }, ), 复制代码`

下面是参数的介绍

+----------------------------+----------------------------------------------------------------------------------------+--------------------------------------+
|            参数            |                                          描述                                          |                 默认                 |
+----------------------------+----------------------------------------------------------------------------------------+--------------------------------------+
| size                       | 点赞Widget的大小                                                                       |                                 30.0 |
| animationDuration          | isLiked变化动画的时间长度                                                              | const Duration(milliseconds:         |
|                            |                                                                                        | 1000)                                |
| bubblesSize                | 大小泡泡的最外圈的大小                                                                 | size * 2.0                           |
| bubblesColor               | 泡泡的颜色，由4个颜色组成                                                              | const                                |
|                            |                                                                                        | BubblesColor(dotPrimaryColor: const  |
|                            |                                                                                        | Color(0xFFFFC107),dotSecondaryColor: |
|                            |                                                                                        | const                                |
|                            |                                                                                        | Color(0xFFFF9800),dotThirdColor:     |
|                            |                                                                                        | const                                |
|                            |                                                                                        | Color(0xFFFF5722),dotLastColor:      |
|                            |                                                                                        | const Color(0xFFF44336),)            |
| circleSize                 | 动画时候的圈圈的最终大小                                                               | size * 0.8                           |
| circleColor                | 动画时候的圈圈的颜色，有开始颜色和结束颜色组成                                         | const CircleColor(start: const       |
|                            |                                                                                        | Color(0xFFFF5722), end: const        |
|                            |                                                                                        | Color(0xFFFFC107)                    |
| onTap                      | 点击按钮回调                                                                           | 你可以在这里处理请求，改变数据       |
| isLiked                    | 是否喜欢                                                                               | false                                |
| likeCount                  | 喜欢数量，如果未null的话，不显示likeCount部分                                          | null                                 |
| mainAxisAlignment          | MainAxisAlignment值，like_button是由Row组成的                                          | MainAxisAlignment.center             |
| likeBuilder                | 生成like                                                                               | null                                 |
|                            | widget的回调，返回值为Widget,可以根据isLiked，来生成不同样式的UI                       |                                      |
| countBuilder               | 生成count                                                                              | null                                 |
|                            | widget的回调，返回值为Widget,可以根据isLiked/likeCount，来生成不同样式的UI             |                                      |
| likeCountAnimationDuration | likeCount变化动画的时间长度                                                            | const Duration(milliseconds:         |
|                            |                                                                                        | 500)                                 |
| likeCountAnimationType     | likeCount动画的类型。none没有动画；part动画只作用在变化的部分；all动画作用在整个数字上 | LikeCountAnimationType.part          |
| likeCountPadding           | likeCount widget的padding                                                              | const EdgeInsets.only(left:          |
|                            |                                                                                        | 3.0)                                 |
+----------------------------+----------------------------------------------------------------------------------------+--------------------------------------+

## 在什么时候去通知服务端,更新数据 ##

在点击的时候，注册onTap回调。

如果你请求成功，返回成功的isLiked值;

如果失败了，请返回null，这样ui就不会变化，建议这个时候Toast告知一下。

` LikeButton( onTap: ( bool isLiked) { return onLikeButtonTap(isLiked, item); },) 复制代码` ` Future< bool > onLikeButtonTap( bool isLiked, TuChongItem item) { ///send your request here /// final Completer< bool > completer = new Completer< bool >(); Timer( const Duration (milliseconds: 200 ), () { item.is_favorite = !item.is_favorite; item.favorites = item.is_favorite ? item.favorites + 1 : item.favorites - 1 ; // if your request is failed,return null, completer.complete(item.is_favorite); }); return completer.future; } 复制代码`

最后放上 [like_button]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Ffluttercandies%2Flike_button ) ，如果你有什么不明白或者对这个方案有什么改进的地方，请告诉我，欢迎加入 [Flutter Candies]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Ffluttercandies ) ，一起生产可爱的Flutter 小糖果(QQ群:181398081)

最最后放上 [Flutter Candies]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Ffluttercandies ) 全家桶，真香。

![](https://user-gold-cdn.xitu.io/2019/5/29/16b02e0775f4af97?imageView2/0/w/1280/h/960/ignore-error/1)