# 今日头条屏幕适配方案落地研究 #

> 
> 
> 
> 原文作者:但，我知道，链接： www.cnblogs.com/haichao/p/1… [](
> https://link.juejin.im?target=https%3A%2F%2Fwww.cnblogs.com%2Fhaichao%2Fp%2F10020893.html
> )
> 
> 

### 目录 ###

* 前言 []( https://link.juejin.im?target=https%3A%2F%2Fwww.cnblogs.com%2Fhaichao%2Fp%2F10020893.html%23%25E5%2589%258D%25E8%25A8%2580 )
* 各平板数据比较 []( https://link.juejin.im?target=https%3A%2F%2Fwww.cnblogs.com%2Fhaichao%2Fp%2F10020893.html%23%25E5%2590%2584%25E5%25B9%25B3%25E6%259D%25BF%25E6%2595%25B0%25E6%258D%25AE%25E6%25AF%2594%25E8%25BE%2583 )
* 为什么看起来更小了？(头条方案跟最小宽度方案比较) []( https://link.juejin.im?target=https%3A%2F%2Fwww.cnblogs.com%2Fhaichao%2Fp%2F10020893.html%23%25E4%25B8%25BA%25E4%25BB%2580%25E4%25B9%2588%25E7%259C%258B%25E8%25B5%25B7%25E6%259D%25A5%25E6%259B%25B4%25E5%25B0%258F%25E4%25BA%2586%25E5%25A4%25B4%25E6%259D%25A1%25E6%2596%25B9%25E6%25A1%2588%25E8%25B7%259F%25E6%259C%2580%25E5%25B0%258F%25E5%25AE%25BD%25E5%25BA%25A6%25E6%2596%25B9%25E6%25A1%2588%25E6%25AF%2594%25E8%25BE%2583 )
* smallesWidth 方案迁移 []( https://link.juejin.im?target=https%3A%2F%2Fwww.cnblogs.com%2Fhaichao%2Fp%2F10020893.html%23smalleswidth-%25E6%2596%25B9%25E6%25A1%2588%25E8%25BF%2581%25E7%25A7%25BB )
* 优缺点 []( https://link.juejin.im?target=https%3A%2F%2Fwww.cnblogs.com%2Fhaichao%2Fp%2F10020893.html%23%25E4%25BC%2598%25E7%25BC%25BA%25E7%2582%25B9 )
* issue []( https://link.juejin.im?target=https%3A%2F%2Fwww.cnblogs.com%2Fhaichao%2Fp%2F10020893.html%23issue )
* 附录（适配核心代码） []( https://link.juejin.im?target=https%3A%2F%2Fwww.cnblogs.com%2Fhaichao%2Fp%2F10020893.html%23%25E9%2599%2584%25E5%25BD%2595%25E9%2580%2582%25E9%2585%258D%25E6%25A0%25B8%25E5%25BF%2583%25E4%25BB%25A3%25E7%25A0%2581 )

### 前言 ###

大家好，现在给大家推荐一种极低版本的 Android 屏幕适配方案，就是今日头条适配方案，“极低成本”这四个字正是今日头条的适配文章标题。

众所周知，安卓的屏幕碎片化极其严重，适配一直是从事安卓开发人员十分头疼的事情。前期，由于公司支持的平板款式单一，只需要做几款平板的适配即可，选用了 smalledtWidth(最小宽度)适配，但是这个方案在增加新屏幕时且原 dimens 文件无法很好适配时，就需要增加新屏幕的最小宽度 dimens 文件了，比较麻烦而且会增加项目大小(虽然只是几个文件)，而且这种屏幕适配极度依赖设备的屏幕密度，叫density。为了讲解更清楚，这里需要引入几个公式：

> 
> 
> 
> **px = density * dp** dp : 安卓开发人员常常挂在嘴上的长度单位 px : 设计人员眼中的长度单位 **density =
> dpi / 160** 因此， **px = dp * (dpi/160)** dpi : 根据屏幕真实分辨率和尺寸计算得出 举个例子：屏幕分辨率为
> 1920 * 1080，屏幕尺寸为5寸(屏幕斜边长度cm/0.3937), 则 **dpi = √(宽度²+ 高度²)/屏幕尺寸**
> 
> 

因此，屏幕密度至关重要，屏幕密度怎么来的？厂商写入一个 system/build.prop 文件，有时还会写错，就我们一款华为平板，获取的屏幕密度是2，但是手工测量并按公式得到实际屏幕密度是1.56。导致我们的适配方案在那款平板就失效了。

本人一直在寻找可以一劳永逸的屏幕适配方案，今日头条是选定基准分辨率，基于设备屏幕分辨率计算出新的屏幕密度进行适配，保证所有设备的显示效果一致，完美避开上面那款设备的问题。推荐给大家。

### 各平板数据比较 ###

首先，我详细记录了公司主流设备的参数，新方案肯定要对主流设备都能完美适配，这才是入门门槛。

+----------------+---------------+---------------------+----------------------+
|       Q        | 三星N5100-4.1 | 三星P355C-6.0(基准) |       华为-8.0       |
+----------------+---------------+---------------------+----------------------+
| 真实宽度(px)   |           800 |                 768 |                 1200 |
| 真实高度(px)   |          1280 |                1024 |                 1852 |
| 原始 density   |       1.33125 |                 1.0 | 2.0(不准，实际1.56 ) |
| new density    |       1.04166 | -                   |               1.5625 |
| new height(px) |          1066 | -                   |                 1600 |
+----------------+---------------+---------------------+----------------------+

可以看到横向是几种设备，竖向是一些参数，其中中英文混杂，这是为什么呢？这是我故意的，中文是设备原始参数，英文是根据今日头条方案原理计算的。因为，今日头条的目的是所有设备的显示效果一致。但是设备的分辨率是不同的，怎么显示一致呢？简单述之，就是缩放，按宽度缩放的。可能有人会有疑问，缩放后的效果图放不下，显示不完整怎么办？

我们看看上面的数据，可以看到按照三星6.0基准进行缩放，效果图在三星4.1这款设备宽度上的显示，是按768乘以new density ,也就是 1.04166 进行放大，不用按计算器了，就是800px，完美适配。那么高度呢，1024 也乘以 new density，发现是1066px，比实际高度像素值 1280px 小，不会出现显示不全的现象。可能有人会问了，这不是多出来了么，会不会留空白啊？对，好问题，所以合格的开发在竖向布局上增加自适应权重，以应对这种情况。当然，横向也需要考虑自适应权重。

同理，可得知效果图在华为8.0设备的宽度像素是 1600px， 也比实际设备宽度 1852px 小，也能显示完全。

### 为什么看起来更小了？(头条方案跟最小宽度方案比较) ###

对的，跟原先的比起来，是更小了，包括图片更小，文字更小。这是为什么呢？且听我细细道来... ...

大家都知道，安卓有 mdpi、hdpi、xhdpi后缀的文件，具体使用有 drawable-mdpi、drawable-hdpi，或者mipmap-mdpi、mipmap-hdpi, 又或者 values-mdpi、values-hdpi, 这些都是安卓自带的屏幕适配方案，只是不太好用吗，经常出问题。那么，这些文件都是怎么使用的呢，这又涉及到了屏幕密度这个属性，关联如下：

+-----------------+-------------+
|       DPI       |  屏幕密度   |
+-----------------+-------------+
| drawable-ldpi   |        0.75 |
| drawable-mdpi   | 1(baseline) |
| drawable-hdpi   |         1.5 |
| drawable-xhdpi  |           2 |
| drawable-xxhdpi |           3 |
| drawable-xxxdpi |           4 |
+-----------------+-------------+

* 平板A 三星平板5100 的屏幕密度是1.33125，大于mdpi，小于hdpi，向上取整，所以属于hdpi
* 平板B 三星平板P355C 的屏幕密度是1，属于mdpi
* ldpi:mdpi:hdpi:xhdpi:xxhdpi:xxxdpi = 0.75:1:1.5:2:3:4 = 3:4:6:8:12:16
* 上述比值乘以12，就是 36：48：72：144：192，刚好就是icon尺寸
* 我们会看到，最小宽度适配方案，values-hdpi 的值是 values-mdpi 的值乘以 0.8

#### 0.8 的参数 ####

* 宽高100dp的正方形图片，平板A会显示100px,平板B会乘以1.5，显示成150px，导致偏大
* 由于平板B的屏幕密度是 1.33125， 最好 显示成 100* 1.33125
* 1.33125/ 1.5 = 0.8875 约为 0.8

#### sw600dp-dpi ####

* sw : small width,就是最小宽度是600dp,
* px -> dp ： dp = px / density
* 平板A: 800 /1.33125 = 600.93
* 平板B: 768/1 = 768 上述两个平板，一个是600dp，一个是768dp,都是大于600dp,平板A使用sw600dp-hdpi,平板B使用sw600dp-mdpi

#### 最后称述 ####

平板A、B 同时显示一个 100px 的图片：

* 按最小宽度适配：100 * 1.5 * 0.8 = 120 ，图片会显示成 120px
* 按今日头条适配： 100 * 1.04166 = 104.166，图片会显示成 104.166 px
* 所以今日头条方案显示的图片就更小了。

那么，哪个更好呢？我们再来看看一个极端，显示一个 平板B 的填满宽度的图片， 768px：

* 按最小宽度适配：768px * 1.5 * 0.8 = 921.6px ，图片会显示成 921.6px, 远远超出平板A的尺寸，此时开发人员需要手动干预
* 按今日头条适配： 768px * 1.04166 = 799.99488，图片可以看成显示成 800 px
* 优点很明显，布局更简单

严谨的你，可能会问了，那显示超过768px呢？ 不好意思，我们的基准就是 768，不会超过他了。

### smallesWidth 方案迁移 ###

我们原项目使用的是 smallestWidth 方案，经试验迁移代价很低，经研究有如下两个方案。

* 删除所有适配 smallestWidth 的dimens 文件夹,只保留dp 值是1：1 的 dimens 文件即可；
* 不想删除亦可，将所有的 dimens 文件都覆盖成 dp 值是1：1 的 dimens 文件即可

### 优缺点 ###

#### 优点 ####

* 使用成本非常低，操作非常简单，使用该方案无需增加dimens 文件，修改代码，完虐其他屏幕适配方案
* 侵入性非常低，切换几乎瞬间完成，试错成本接近为0
* 修改的 density 是全局的，一次修改，终生受益。
* 不会有任何性能的损耗
* 今日头条 大厂保证

### 缺点 ###

1、 第三方布局库， 未按项目效果图布局，全局修改 density 导致修改第三方布局，造成显示界面问题 2、与 smallestwith 适配方案不兼容，切换回来比较麻烦

### issue ###

#### 一个 Bitmap 的density 问题 ####

在某处，开启今日头条适配方案，全局修改屏幕密度，获取 ImageView 的 Bitmap 的宽高，发现获取的宽高和实际的宽高(布局出来观察)不一致。经查阅源码，发现 Bitmap 也有一个 density, 怀疑未被修改。

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2a5cbc7394a80?imageView2/0/w/1280/h/960/ignore-error/1)

随决定，修改 sDefaultDensity 值，查阅代码，发现 sDefaultDensity 是静态私有，于是召唤反射大法

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2a5cbc7574567?imageView2/0/w/1280/h/960/ignore-error/1)

测试 Ok, 收工。

### 附录（适配核心代码） ###

* initAppDensity 方法 Application 调用，记录默认屏幕密度
* setDefault 和 setOrientation 方法 Activity 调用，设置新的屏幕密度
* resetAppOrientation 方法，恢复屏幕密度

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2a5cbc74332ca?imageView2/0/w/1280/h/960/ignore-error/1) ![](https://user-gold-cdn.xitu.io/2019/6/6/16b2a5cbc7227d17?imageView2/0/w/1280/h/960/ignore-error/1)

` // * ================================================ // * 本框架核心原理来自于 <a href= "https://mp.weixin.qq.com/s/d9QCoBP6kV9VSWvVldVVwA" >今日头条官方适配方案</a> // * <p> // * 本框架源码的注释都很详细, 欢迎阅读学习 // * <p> // * 任何方案都不可能完美, 在成本和收益中做出取舍, 选择出最适合自己的方案即可, 在没有更好的方案出来之前, 只有继续忍耐它的不完美, 或者自己作出改变 // * 既然选择, 就不要抱怨, 感谢 今日头条技术团队 和 张鸿洋 等人对 Android 屏幕适配领域的的贡献 // * <p> // * ================================================ // */ private static final int WIDTH = 1; private static final int HEIGHT = 2; private static final float DEFAULT_WIDTH = 768f; //默认宽度 private static final float DEFAULT_HEIGHT = 1024f; //默认高度 private static float appDensity; /** * 字体的缩放因子，正常情况下和density相等，但是调节系统字体大小后会改变这个值 */ private static float appScaledDensity; /** * 状态栏高度 */ private static int barHeight; private static DisplayMetrics appDisplayMetrics; private static float densityScale = 1.0f; /** * application 层调用，存储默认屏幕密度 * * @param application application */ public static void initAppDensity(@NonNull final Application application) { //获取application的DisplayMetrics appDisplayMetrics = application.getResources().getDisplayMetrics(); //获取状态栏高度 barHeight = getStatusBarHeight(application); if (appDensity == 0) { //初始化的时候赋值 appDensity = appDisplayMetrics.density; appScaledDensity = appDisplayMetrics.scaledDensity; //添加字体变化的监听 application.registerComponentCallbacks(new ComponentCallbacks () { @Override public void onConfigurationChanged(Configuration newConfig) { //字体改变后,将appScaledDensity重新赋值 if (newConfig != null && newConfig.fontScale > 0) { appScaledDensity = application.getResources().getDisplayMetrics().scaledDensity; } } @Override public void onLowMemory () { } }); } } /** * 此方法在BaseActivity中做初始化(如果不封装BaseActivity的话,直接用下面那个方法就好了) * * @param activity activity */ public static void set Default(Activity activity) { set AppOrientation(activity, WIDTH); } /** * 比如页面是上下滑动的，只需要保证在所有设备中宽的维度上显示一致即可， * 再比如一个不支持上下滑动的页面，那么需要保证在高这个维度上都显示一致 * * @param activity activity * @param orientation WIDTH HEIGHT */ public static void set Orientation(Activity activity, int orientation) { set AppOrientation(activity, orientation); } /** * 重设屏幕密度 * * @param activity activity * @param orientation WIDTH 宽，HEIGHT 高 */ private static void set AppOrientation(@NonNull Activity activity, int orientation) { float targetDensity; if (orientation == HEIGHT) { targetDensity = (appDisplayMetrics.heightPixels - barHeight) / DEFAULT_HEIGHT; } else { targetDensity = appDisplayMetrics.widthPixels / DEFAULT_WIDTH; } float targetScaledDensity = targetDensity * (appScaledDensity / appDensity); int targetDensityDpi = (int) (160 * targetDensity); // 最后在这里将修改过后的值赋给系统参数,只修改Activity的density值 DisplayMetrics activityDisplayMetrics = activity.getResources().getDisplayMetrics(); activityDisplayMetrics.density = targetDensity; activityDisplayMetrics.scaledDensity = targetScaledDensity; activityDisplayMetrics.densityDpi = targetDensityDpi; densityScale = appDensity / targetDensity; set BitmapDefaultDensity(activityDisplayMetrics.densityDpi); } /** * 重置屏幕密度 * * @param activity activity */ public static void resetAppOrientation(@NonNull Activity activity) { DisplayMetrics activityDisplayMetrics = activity.getResources().getDisplayMetrics(); activityDisplayMetrics.density = appDensity; activityDisplayMetrics.scaledDensity = appScaledDensity; activityDisplayMetrics.densityDpi = (int) (appDensity * 160); densityScale = 1.0f; set BitmapDefaultDensity(activityDisplayMetrics.densityDpi); } /** * 获取状态栏高度 * * @param context context * @ return 状态栏高度 */ private static int getStatusBarHeight(Context context) { int result = 0; int resourceId = context.getResources().getIdentifier( "status_bar_height" , "dimen" , "android" ); if (resourceId > 0) { result = context.getResources().getDimensionPixelSize(resourceId); } return result; } /** * 设置 Bitmap 的默认屏幕密度 * 由于 Bitmap 的屏幕密度是读取配置的，导致修改未被启用 * 所有，放射方式强行修改 * @param defaultDensity 屏幕密度 */ private static void set BitmapDefaultDensity(int defaultDensity) { //获取单个变量的值 Class clazz; try { clazz = Class.forName( "android.graphics.Bitmap" ); Field field = clazz.getDeclaredField( "sDefaultDensity" ); field.setAccessible( true ); field.set(null, defaultDensity); field.setAccessible( false ); } catch (ClassNotFoundException e) { } catch (NoSuchFieldException e) { } catch (IllegalAccessException e) { e.printStackTrace(); } } /** * 屏幕密度缩放系数 * * @ return 屏幕密度缩放系数 */ public static float getDensityScale () { return densityScale; } 复制代码`

### 阅读更多 ###

**面试官：你分析过线程池源码吗？** []( https://link.juejin.im?target=http%3A%2F%2Fmp.weixin.qq.com%2Fs%3F__biz%3DMzI3OTU0MzI4MQ%3D%3D%26amp%3Bmid%3D2247488047%26amp%3Bidx%3D1%26amp%3Bsn%3D986176b5348e3850e5b92e3ac5ea2a40%26amp%3Bchksm%3Deb477eb1dc30f7a7365cdc87e036260cc6c13ba521ed29a255b4ce62cda843b704c6cd48ab90%26amp%3Bscene%3D21%23wechat_redirect )

**屏幕适配之尺寸的相关概论《一》** []( https://link.juejin.im?target=http%3A%2F%2Fmp.weixin.qq.com%2Fs%3F__biz%3DMzI3OTU0MzI4MQ%3D%3D%26amp%3Bmid%3D2247485137%26amp%3Bidx%3D1%26amp%3Bsn%3D970c4aa3251853064406ed6c1cac1ed6%26amp%3Bchksm%3Deb476a4fdc30e359331f414f5312a96e0cfabaf630922d0c59ffa3a8499e4aa68d5e44c07854%26amp%3Bscene%3D21%23wechat_redirect )

**Android屏幕适配框架-(今日头条终极适配方案)** []( https://link.juejin.im?target=http%3A%2F%2Fmp.weixin.qq.com%2Fs%3F__biz%3DMzI3OTU0MzI4MQ%3D%3D%26amp%3Bmid%3D2247486334%26amp%3Bidx%3D1%26amp%3Bsn%3D1a3189739bc0ad984f67e18c09f35055%26amp%3Bchksm%3Deb4767e0dc30eef6fea34d891fca87b95edf5a834d5b8a3612f0e3700690e21607eb41499bf1%26amp%3Bscene%3D21%23wechat_redirect )

**高仿安卓「填空题」控件！手撸一个炫酷的View动效** []( https://link.juejin.im?target=http%3A%2F%2Fmp.weixin.qq.com%2Fs%3F__biz%3DMzI3OTU0MzI4MQ%3D%3D%26amp%3Bmid%3D2247488023%26amp%3Bidx%3D1%26amp%3Bsn%3Da325358a0dbd3c7fe1b4f8a5b890a17e%26amp%3Bchksm%3Deb477e89dc30f79f3d543ae42eac31ce7aa9bd0abefa2d83bd9d3c48c8b723244f3499cf3d96%26amp%3Bscene%3D21%23wechat_redirect )

**回顾我两个月面试阿里，携程，小红书，美团，网易等等** []( https://link.juejin.im?target=http%3A%2F%2Fmp.weixin.qq.com%2Fs%3F__biz%3DMzI3OTU0MzI4MQ%3D%3D%26amp%3Bmid%3D2247488104%26amp%3Bidx%3D1%26amp%3Bsn%3D8e59dbd80c530b8d8d877ff6b61a0a20%26amp%3Bchksm%3Deb477ef6dc30f7e0de3c072f190556f2619c453d7cc23750d240ce61c0793c37ba56921c6c90%26amp%3Bscene%3D21%23wechat_redirect )

![image](https://user-gold-cdn.xitu.io/2019/6/6/16b2a5ce1f15ff35?imageView2/0/w/1280/h/960/ignore-error/1)