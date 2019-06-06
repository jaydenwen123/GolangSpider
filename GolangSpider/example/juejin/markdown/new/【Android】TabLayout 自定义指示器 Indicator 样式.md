# 【Android】TabLayout 自定义指示器 Indicator 样式 #

[本文CSDN博客]( https://link.juejin.im?target=https%3A%2F%2Fblog.csdn.net%2Fu013719138%2Farticle%2Fdetails%2F89964674 )

[本文简书]( https://link.juejin.im?target=https%3A%2F%2Fwww.jianshu.com%2Fp%2F6b3dc9d82634 )

在布局里加入 TabLayout，默认是下划线的样式，可以使用 ` tabIndicatorGravity` 属性设置为： ` bottom` （默认值，可以不用设置，指示器显示在底部）、 ` top` （指示器显示在顶部）、 ` center` （指示器显示在中间）、 ` stretch` （指示器高度拉伸铺满 item）。

` <android.support.design.widget.TabLayout android:id="@+id/tl" android:layout_width="match_parent" android:layout_height="wrap_content" app:tabIndicatorColor="@color/colorPrimary" app:tabIndicatorFullWidth="true" <!-- 设置 Indicator 高度 --> app:tabIndicatorHeight="2dp" app:tabMode="scrollable" /> 复制代码`

### 1. "app:tabIndicatorFullWidth" 属性 ###

注意 ` app:tabIndicatorFullWidth="true"` 属性，设为 **true** ，是 Indicator 充满 item 的宽度：

![app:tabIndicatorFullWidth=](https://user-gold-cdn.xitu.io/2019/6/5/16b274dfd290d941?imageslim)

设为 **false** 是 Indicator 保持和 item 的内容宽度一致：

![app:tabIndicatorFullWidth=](https://user-gold-cdn.xitu.io/2019/6/5/16b274dfefdd45ee?imageslim)

### 2. 给 Indicator 设置边距 ###

网上的做法一般是通过反射来设置 Indicator 的宽度，可以参见博客： [关于Android改变TabLayout 下划线(Indicator)宽度实践总结]( https://link.juejin.im?target=https%3A%2F%2Fwww.jianshu.com%2Fp%2F83922d08250b )

不过我觉得可以使用 ` layer-list` 来实现。 在 ` drawable` 文件夹下新建一个 ` indicator.xml` 文件：

` <?xml version="1.0" encoding="utf-8"?> <layer-list xmlns:android="http://schemas.android.com/apk/res/android"> <item <!-- 设置左边距 --> android:left="15dp" <!-- 设置右边距 --> android:right="15dp"> <!-- 注：这里需要一个空的 <shape /> 标签，否则会报错 --> <shape /> </item> </layer-list> 复制代码`

**需要注意的是在 shape 里设置颜色是无效的，需要在布局文件里设置 Indicator 颜色。**

在 ` TabLayout` 布局里添加 Indicator 的样式：

` <android.support.design.widget.TabLayout android:id="@+id/tl" android:layout_width="match_parent" android:layout_height="wrap_content" <!-- 设置 Indicator 高度 --> app:tabIndicatorHeight="2dp" <!-- 设置 Indicator 颜色 --> app:tabIndicatorColor="@color/colorPrimary" <!-- 设置 Indicator 的样式 --> app:tabIndicator="@drawable/indicator" app:tabMode="scrollable" /> 复制代码`

![设置app:tabIndicator](https://user-gold-cdn.xitu.io/2019/6/5/16b274dfefe438e9?imageslim)

### 3. 给 Indicator 设置圆角 ###

如果不需要边距，只需要圆角，可以配合 ` app:tabIndicatorFullWidth` 属性，使用 ` shape` 设置 ` app:tabIndicator` 来实现圆角即可，无需使用 ` layer-list` ，代码就不用贴了吧~

这里为了使效果看得明显一点，把 Indicator 的高度设置为 5dp。 给 Indicator 添加了 5dp 的圆角：

` <? xml version= "1.0" encoding= "utf-8" ?> < layer-list xmlns:android = "http://schemas.android.com/apk/res/android" > < item android:left = "15dp" android:right = "15dp" > < shape > < corners android:radius = "5dp" /> </ shape > </ item > </ layer-list > 复制代码`

![带圆角的下划线](https://user-gold-cdn.xitu.io/2019/6/5/16b274dfefec06e5?imageslim)

### 4. 给 Indicator 设置宽高 ###

#### 4.1 在 ` <shape>` 的 ` <size>` 标签里设置宽高： ####

` <? xml version= "1.0" encoding= "utf-8" ?> < layer-list xmlns:android = "http://schemas.android.com/apk/res/android" > <!-- 若不设置 gravity，则 Indicator 宽度会填满整个 item --> < item android:gravity = "center_horizontal" > < shape > < corners android:radius = "5dp" /> < size android:width = "20dp" android:height = "5dp" /> </ shape > </ item > </ layer-list > 复制代码`

#### 4.2 在 API 23 以上，支持直接在 ` layer-list` 里给 ` <item>` 标签设置宽度和高度： ####

` <?xml version="1.0" encoding="utf-8"?> <layer-list xmlns:android="http://schemas.android.com/apk/res/android"> <item android:width="20dp" android:height="5dp" <!-- 若不设置 gravity 则默认是居左显示，需要设置为水平居中显示 --> android:gravity="center_horizontal"> <shape> <corners android:radius="5dp" /> </shape> </item> </layer-list> 复制代码`

![在这里插入图片描述](https://user-gold-cdn.xitu.io/2019/6/5/16b274dfefbc6b01?imageslim)

### 5. tabIndicator 属性源码分析 ###

` TabLayout` 的 ` tabIndicator` 属性里设置的 ` layer-list` 不支持设置颜色。 我们查看一下 ` TabLayout` 的源码，搜索 ` TabLayout_tabIndicator` ：

` this.setSelectedTabIndicator(MaterialResources.getDrawable(context, a, styleable.TabLayout_tabIndicator)); 复制代码`

` setSelectedTabIndicator()` 方法：

![setSelectedTabIndicator()](https://user-gold-cdn.xitu.io/2019/6/5/16b274e007e80a35?imageView2/0/w/1280/h/960/ignore-error/1) 搜索一下使用 ` tabSelectedIndicator` 的地方，在 ` SlidingTabIndicator` 类里的 ` draw()` 方法里： 第 1 处 ` tabSelectedIndicator` ![在这里插入图片描述](https://user-gold-cdn.xitu.io/2019/6/5/16b274e01a95065f?imageView2/0/w/1280/h/960/ignore-error/1) 再看一下 ` selectedIndicatorHeight` 是什么： ![在这里插入图片描述](https://user-gold-cdn.xitu.io/2019/6/5/16b274e0216d363e?imageView2/0/w/1280/h/960/ignore-error/1) ![在这里插入图片描述](https://user-gold-cdn.xitu.io/2019/6/5/16b274e023dae4fd?imageView2/0/w/1280/h/960/ignore-error/1) ` selectedIndicatorHeight` 是在布局里给 ` TabLayout` 设置的 ` tabIndicatorHeight` 属性。

可见如果我们在布局里给 ` TabLayout` 设置了 ` tabIndicatorHeight` 属性，则 Indicator 高度优先取 ` tabIndicatorHeight` 设置的高度；否则才会取咱们自定义的 ` drawable` 里的高度。

继续，第 2 处 ` tabSelectedIndicator`

![在这里插入图片描述](https://user-gold-cdn.xitu.io/2019/6/5/16b274e027867149?imageView2/0/w/1280/h/960/ignore-error/1) 黄色方框里可以发现，为什么之前在 ` drawable` 里设置的颜色无效了，因为使用的是 ` TabLayout_tabIndicatorColor` 属性里设置的颜色，所以 ` <stroke>` 也无效，只保留了整体的形状样式。

### 6. 自定义复杂的 Indicator 样式 ###

如果需要复杂一点的样式，比如 ` <stroke>` 。 先写一个 tab 被选中时的样式 ` indicator.xml` ：

` <?xml version="1.0" encoding="utf-8"?> <layer-list xmlns:android="http://schemas.android.com/apk/res/android"> <item <!-- 设置边距 --> android:bottom="8dp" android:left="8dp" android:right="8dp" android:top="8dp"> <shape> <!-- 设置圆角 --> <corners android:radius="5dp" /> <!-- 设置边框 --> <stroke android:width="1dp" android:color="@color/colorAccent" /> </shape> </item> </layer-list> 复制代码`

还需要一个 ` selector.xml` ：

` <? xml version= "1.0" encoding= "utf-8" ?> < selector xmlns:android = "http://schemas.android.com/apk/res/android" > < item android:drawable = "@drawable/indicator" android:state_selected = "true" /> </ selector > 复制代码`

接下来，我们要设置的是 ` tabBackground` ，也就是 tab 标签的背景，而不再是 ` tabIndicator` ，所以要把 Indicator 的高度设为 0 ，不使用 tab 原生的 Indicator。

**这里还要注意一下 ` tabRippleColor` 属性，是设置点击 tab 标签时的波纹颜色，不设置的时候，默认是灰色的，文章前面的截图里有显示效果。如果想去掉这个效果，设置颜色为透明即可。**

` <android.support.design.widget.TabLayout android:id="@+id/tl" android:layout_width="match_parent" android:layout_height="wrap_content" <!-- 使用我们自定义的点击样式 --> app:tabBackground="@drawable/selector" <!-- tabIndicator 高度设为 0 --> app:tabIndicatorHeight="0dp" app:tabMode="scrollable" <!-- 设置点击时的波纹颜色为透明 --> app:tabRippleColor="@android:color/transparent" /> 复制代码`

![在这里插入图片描述](https://user-gold-cdn.xitu.io/2019/6/5/16b274e034e400e6?imageslim) 想实现更复杂的效果，可以使用 [MagicIndicator]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fhackware1993%2FMagicIndicator )

附上一个效果图，感觉还是很酷炫的：

![在这里插入图片描述](https://user-gold-cdn.xitu.io/2019/6/5/16b274e03afbffe8?imageslim)