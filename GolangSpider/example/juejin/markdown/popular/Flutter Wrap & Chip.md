# Flutter Wrap & Chip #

在写业务的时候，难免有***搜索***功能，一般搜索功能的页面如下：

![](https://user-gold-cdn.xitu.io/2019/6/5/16b24fc2fa829c5e?imageView2/0/w/1280/h/960/ignore-error/1)

那如果在Android上面写的话，

一般来讲是一个 RecyclerView + 自动换行的 layoutManager + 自定义的background。

当然这个自动换行的LayoutManager 还得自己定义，去算坐标。

那 Flutter 提供给我们很方便的控件 Wrap + Chip，用这两个就能轻松实现上诉效果。

先来说一下Wrap。

### Wrap ###

看名字我们也能看出来，它就是一个包裹住子布局的 Widget，并且可以自动换行。

先看一下构造函数，来确定一下需要传入什么参数：

` Wrap({ Key key, this.direction = Axis.horizontal, // 子布局排列方向 this.alignment = WrapAlignment.start, // 子布局对齐方向 this.spacing = 0.0 , // 间隔 this.runAlignment = WrapAlignment.start, // run 可以理解为新的一行，新的一行的对齐方向 this.runSpacing = 0.0 , // 两行的间距 this.crossAxisAlignment = WrapCrossAlignment.start, this.textDirection, this.verticalDirection = VerticalDirection.down, List <Widget> children = const <Widget>[], }) : super (key: key, children: children); 复制代码`

最基本的我们只需要传入一个children即可，但是我们想要好的效果，一般都会传入 spacing 和 runSpacing。

### Chip ###

下面看一下 Chip，Chip可以理解为碎片的意思，还是先来看一下构造函数：

` Chip({ Key key, this.avatar, //左侧Widget，一般为小图标 @required this.label, //标签 this.labelStyle, this.labelPadding, this.deleteIcon //删除图标 this.onDeleted //删除回调，为空时不显示删除图标 this.deleteIconColor //删除图标的颜色 this.deleteButtonTooltipMessage //删除按钮的tip文字 this.shape //形状 this.clipBehavior = Clip.none, this.backgroundColor //背景颜色 this.padding, // padding this.materialTapTargetSize //删除图标material点击区域大小 }) 复制代码`

可以看到这里东西还是不少的，最基本的要传入一个label。

label 一般就为我们的 text，先来看一下只定义一个 label 的效果：

![](https://user-gold-cdn.xitu.io/2019/6/5/16b24fc30f59c636?imageView2/0/w/1280/h/960/ignore-error/1)

下面再加入 avatar：

![](https://user-gold-cdn.xitu.io/2019/6/5/16b24fc3059e36ba?imageView2/0/w/1280/h/960/ignore-error/1)

再来加入 delete：

![](https://user-gold-cdn.xitu.io/2019/6/5/16b24fc2f49dda3b?imageView2/0/w/1280/h/960/ignore-error/1)

这里注意，一定要设置上 onDeleted 参数，否则不显示delete控件。

### 编写 '历史搜索' ###

前面都是在 children 里添加widget 来达到目的，不好做删除工作。

现在我们来使用 ListView，并添加删除事件。

代码如下：

` import 'package:flutter/material.dart' ; class WrapPage extends StatefulWidget { @override _WrapPageState createState() => _WrapPageState(); } class _WrapPageState extends State < WrapPage > { // 生成历史数据 final List < String > _list = List < String >.generate( 10 , (i) => 'chip$i' ); @override Widget build(BuildContext context) { return Scaffold( appBar: AppBar( title: Text( 'WrapPage' ), ), body: Wrap( spacing: 10 , runSpacing: 5 , children: _list.map<Widget>((s) { return Chip( label: Text( '$s' ), avatar: Icon(Icons.person), deleteIcon: Icon( Icons.close, color: Colors.red, ), onDeleted: () { setState(() { _list.remove(s); // 删除事件 }); }, ); }).toList() )); } } 复制代码`

效果如下：

![](https://user-gold-cdn.xitu.io/2019/6/5/16b24fc2f6b91f27?imageslim)

### 总结 ###

Flutter 提供给我们很多好用的 widget， 我们只需要组合起来就可以达到我们的目的。

完整代码已经传至GitHub： [github.com/wanglu1209/…]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fwanglu1209%2FWFlutterDemo )

![](https://user-gold-cdn.xitu.io/2019/5/31/16b0e6e72c8141db?imageView2/0/w/1280/h/960/ignore-error/1)