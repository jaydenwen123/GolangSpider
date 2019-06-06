# Android刘海屏、水滴屏全面屏适配方案 #

> 
> 
> 
> 我将适配方案整理后，封装成了一个库并上传至github，可参考使用
> 
> 
> 
> 项目地址: [github.com/smarxpan/No…](
> https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fsmarxpan%2FNotchScreenTool
> )
> 
> 

市面上的屏幕尺寸和全面屏方案五花八门。

这里我使用了小米的图来说明：

![](https://user-gold-cdn.xitu.io/2019/6/4/16b21c10a9ee7b56?imageView2/0/w/1280/h/960/ignore-error/1)

上述两种屏幕都可以统称为刘海屏，不过对于右侧较小的刘海，业界一般称为水滴屏或美人尖。为便于说明，后文提到的「刘海屏」「刘海区」都同时指代上图两种屏幕。

## 当我们在谈屏幕适配时，我们在谈什么 ##

* 适应更长的屏幕
* 防止内容被刘海遮挡

其中第一点是所有应用都需要适配的，对应下文的 ` 声明最大长宽比`

而第二点，如果应用本身不需要全屏显示或使用沉浸式状态栏，是不需要适配的。

针对需要适配第二点的应用，需要获取刘海的位置和宽高，然后将显示内容避开即可。

## 声明最大长宽比 ##

以前的普通屏长宽比为16：9，全面屏手机的屏幕长宽比增大了很多，如果不适配的话就会类似下面这样：

![](https://user-gold-cdn.xitu.io/2019/6/4/16b21c1686dab0df?imageView2/0/w/1280/h/960/format/png/ignore-error/1)

黑色区域为未利用的区域。

### 适配方式 ###

适配方式有两种：

* 

将targetSdkVersion版本设置到API 24及以上

这个操作将会为 ` <application>` 标签隐式添加一个属性， ` android:resizeableActivity="true"` , 该属性的作用后面将详细说明。

* 

在 ` <application>` 标签中增加属性： ` android:resizeableActivity="false"`

同时在节点下增加一个 ` meta-data` 标签：

` <!-- Render on full screen up to screen aspect ratio of 2.4 --> <!-- Use a letterbox on screens larger than 2.4 --> <meta-data android:name="android.max_aspect" android:value="2.4" /> 复制代码`

### 原理说明 ###

这里涉及到的知识点是android:resizeableActivity属性。

在 Android 7.0（API 级别 24）或更高版本的应用，android:resizeableActivity属性默认为true（对应适配方式1）。这个属性是控制多窗口显示的，决定当前的应用或者Activity是否支持多窗口。

> 
> 
> 
> [多窗口支持](
> https://link.juejin.im?target=https%3A%2F%2Fdeveloper.android.com%2Fguide%2Ftopics%2Fui%2Fmulti-window.html%3Fhl%3Dzh-CN
> )
> 
> 

在清单的 ` <activity>` 或 ` <application>` 节点中设置该属性，启用或禁用多窗口显示：

` android:resizeableActivity=["true" | "false"] 复制代码`

如果该属性设置为 true，Activity 将能以分屏和自由形状模式启动。 如果此属性设置为 false，Activity 将不支持多窗口模式。 如果该值为 false，且用户尝试在多窗口模式下启动 Activity，该 Activity 将全屏显示。

适配方式2即为设置屏幕的最大长宽比，这是官方提供的设置方式。

**如果设置了最大长宽比，必须 ` android:resizeableActivity="false"` 。 否则最大长宽比没有任何作用。**

## 适配刘海屏 ##

### Android9.0及以上适配 ###

Android P（9.0）开始，官方提供了适配异形屏的方式。

> 
> 
> 
> [Support display cutouts](
> https://link.juejin.im?target=https%3A%2F%2Fdeveloper.android.com%2Fguide%2Ftopics%2Fdisplay-cutout%3Fhl%3Dzh-CN
> )
> 
> 

通过全新的 DisplayCutout 类，可以确定非功能区域的位置和形状，这些区域不应显示内容。 要确定这些凹口屏幕区域是否存在及其位置，请使用 getDisplayCutout() 函数。

* 

全新的窗口布局属性 layoutInDisplayCutoutMode 让您的应用可以为设备凹口屏幕周围的内容进行布局。 您可以将此属性设为下列值之一：

* ` LAYOUT_IN_DISPLAY_CUTOUT_MODE_DEFAULT`
* ` LAYOUT_IN_DISPLAY_CUTOUT_MODE_SHORT_EDGES`
* ` LAYOUT_IN_DISPLAY_CUTOUT_MODE_NEVER`

默认值是 ` LAYOUT_IN_DISPLAY_CUTOUT_MODE_DEFAULT` ，刘海区域不会显示内容，需要将值设置为 ` LAYOUT_IN_DISPLAY_CUTOUT_MODE_SHORT_EDGES`

* 

您可以按如下方法在任何运行 Android P 的设备或模拟器上模拟屏幕缺口：

* 启用开发者选项。
* 在 Developer options 屏幕中，向下滚动至 Drawing 部分并选择 Simulate a display with a cutout。
* 选择凹口屏幕的大小。

* 

适配参考：

` // 延伸显示区域到刘海 WindowManager.LayoutParams lp = window.getAttributes(); lp.layoutInDisplayCutoutMode = WindowManager.LayoutParams.LAYOUT_IN_DISPLAY_CUTOUT_MODE_SHORT_EDGES; window.setAttributes(lp); // 设置页面全屏显示 final View decorView = window.getDecorView(); decorView.setSystemUiVisibility(View.SYSTEM_UI_FLAG_LAYOUT_FULLSCREEN | View.SYSTEM_UI_FLAG_LAYOUT_STABLE); 复制代码`

其中延伸显示区域到刘海的代码，也可以通过修改Activity或应用的style实现，例如：

` <?xml version="1.0" encoding="utf-8"?> <resources> <style name="AppTheme" parent="xxx"> <item name="android:windowLayoutInDisplayCutoutMode">shortEdges</item> </style> </resources> 复制代码`

### Android O 适配 ###

因Google官方的适配方案到Android P才推出，因此在Android O设备上，各家厂商有自己的实现方案。

我这里主要适配了华为、小米、oppo，这三家都给了完整的解决方案。至于vivo，vivo给了判断是否刘海屏的API，但是没用设置刘海区域显示到API，因此无需适配。

#### 适配华为Android O设备 ####

方案一：

* 

具体方式如下所示：

` <meta-data android:name="android.notch_support" android:value="true"/> 复制代码`
* 

对Application生效，意味着该应用的所有页面，系统都不会做竖屏场景的特殊下移或者是横屏场景的右移特殊处理：

` <application android:allowBackup="true" android:icon="@mipmap/ic_launcher" android:label="@string/app_name" android:roundIcon="@mipmap/ic_launcher_round" android:testOnly="false" android:supportsRtl="true" android:theme="@style/AppTheme"> <meta-data android:name="android.notch_support" android:value="true"/> <activity android:name=".MainActivity"> <intent-filter> <action android:name="android.intent.action.MAIN"/> <category android:name="android.intent.category.LAUNCHER"/> </intent-filter> </activity> 复制代码`
* 

对Activity生效，意味着可以针对单个页面进行刘海屏适配，设置了该属性的Activity系统将不会做特殊处理：

` <application android:allowBackup="true" android:icon="@mipmap/ic_launcher" android:label="@string/app_name" android:roundIcon="@mipmap/ic_launcher_round" android:testOnly="false" android:supportsRtl="true" android:theme="@style/AppTheme"> <activity android:name=".MainActivity"> <intent-filter> <action android:name="android.intent.action.MAIN"/> <category android:name="android.intent.category.LAUNCHER"/> </intent-filter> </activity> <activity android:name=".LandscapeFullScreenActivity" android:screenOrientation="sensor"> </activity> <activity android:name=".FullScreenActivity"> <meta-data android:name="android.notch_support" android:value="true"/> </activity> </application> 复制代码`

方案二

对Application生效，意味着该应用的所有页面，系统都不会做竖屏场景的特殊下移或者是横屏场景的右移特殊处理

> 
> 
> 
> 我的 [NotchScreenTool](
> https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fsmarxpan%2FNotchScreenTool
> ) 中使用的就是方案二，如果需要针对Activity，建议自行修改。
> 
> 

* 

设置应用窗口在华为刘海屏手机使用刘海区

` /*刘海屏全屏显示FLAG*/ public static final int FLAG_NOTCH_SUPPORT=0x00010000; /** * 设置应用窗口在华为刘海屏手机使用刘海区 * @param window 应用页面window对象 */ public static void setFullScreenWindowLayoutInDisplayCutout(Window window) { if (window == null) { return; } WindowManager.LayoutParams layoutParams = window.getAttributes(); try { Class layoutParamsExCls = Class.forName("com.huawei.android.view.LayoutParamsEx"); Constructor con=layoutParamsExCls.getConstructor(LayoutParams.class); Object layoutParamsExObj=con.newInstance(layoutParams); Method method=layoutParamsExCls.getMethod("addHwFlags", int.class); method.invoke(layoutParamsExObj, FLAG_NOTCH_SUPPORT); } catch (ClassNotFoundException | NoSuchMethodException | IllegalAccessException |InstantiationException | InvocationTargetException e) { Log.e("test", "hw add notch screen flag api error"); } catch (Exception e) { Log.e("test", "other Exception"); } } 复制代码`
* 

清除添加的华为刘海屏Flag，恢复应用不使用刘海区显示

` /** * 设置应用窗口在华为刘海屏手机使用刘海区 * @param window 应用页面window对象 */ public static void setNotFullScreenWindowLayoutInDisplayCutout (Window window) { if (window == null) { return; } WindowManager.LayoutParams layoutParams = window.getAttributes(); try { Class layoutParamsExCls = Class.forName("com.huawei.android.view.LayoutParamsEx"); Constructor con=layoutParamsExCls.getConstructor(LayoutParams.class); Object layoutParamsExObj=con.newInstance(layoutParams); Method method=layoutParamsExCls.getMethod("clearHwFlags", int.class); method.invoke(layoutParamsExObj, FLAG_NOTCH_SUPPORT); } catch (ClassNotFoundException | NoSuchMethodException | IllegalAccessException |InstantiationException | InvocationTargetException e) { Log.e("test", "hw clear notch screen flag api error"); } catch (Exception e) { Log.e("test", "other Exception"); } } 复制代码`

#### 适配小米Android O设备 ####

* 

判断是否是刘海屏

` private static boolean isNotch() { try { Method getInt = Class.forName("android.os.SystemProperties").getMethod("getInt", String.class, int.class); int notch = (int) getInt.invoke(null, "ro.miui.notch", 0); return notch == 1; } catch (Throwable ignore) { } return false; } 复制代码`
* 

设置显示到刘海区域

` @Override public void setDisplayInNotch(Activity activity) { int flag = 0x00000100 | 0x00000200 | 0x00000400; try { Method method = Window.class.getMethod("addExtraFlags", int.class); method.invoke(activity.getWindow(), flag); } catch (Exception ignore) { } } 复制代码`
* 

获取刘海宽高

` public static int getNotchHeight(Context context) { int resourceId = context.getResources().getIdentifier("notch_height", "dimen", "android"); if (resourceId > 0) { return context.getResources().getDimensionPixelSize(resourceId); } return 0; } public static int getNotchWidth(Context context) { int resourceId = context.getResources().getIdentifier("notch_width", "dimen", "android"); if (resourceId > 0) { return context.getResources().getDimensionPixelSize(resourceId); } return 0; } 复制代码`

#### 适配oppoAndroid O设备 ####

* 

判断是否是刘海屏

` @Override public boolean hasNotch(Activity activity) { boolean ret = false; try { ret = activity.getPackageManager().hasSystemFeature("com.oppo.feature.screen.heteromorphism"); } catch (Throwable ignore) { } return ret; } 复制代码`
* 

获取刘海的左上角和右下角的坐标

` /** * 获取刘海的坐标 * <p> * 属性形如：[ro.oppo.screen.heteromorphism]: [378,0:702,80] * <p> * 获取到的值为378,0:702,80 * <p> * <p> * (378,0)是刘海区域左上角的坐标 * <p> * (702,80)是刘海区域右下角的坐标 */ private static String getScreenValue() { String value = ""; Class<?> cls; try { cls = Class.forName("android.os.SystemProperties"); Method get = cls.getMethod("get", String.class); Object object = cls.newInstance(); value = (String) get.invoke(object, "ro.oppo.screen.heteromorphism"); } catch (Throwable ignore) { } return value; } 复制代码`

Oppo Android O机型不需要设置显示到刘海区域，只要设置了应用全屏就会默认显示。

因此Oppo机型必须适配。

## 适配总结 ##

根据上述功能，我将其整理成了一个依赖库： [NotchScreenTool]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fsmarxpan%2FNotchScreenTool )

使用起来很简单：

` // 支持显示到刘海区域 NotchScreenManager.getInstance().setDisplayInNotch(this); // 获取刘海屏信息 NotchScreenManager.getInstance().getNotchInfo(this, new INotchScreen.NotchScreenCallback() { @Override public void onResult(INotchScreen.NotchScreenInfo notchScreenInfo) { Log.i(TAG, "Is this screen notch? " + notchScreenInfo.hasNotch); if (notchScreenInfo.hasNotch) { for (Rect rect : notchScreenInfo.notchRects) { Log.i(TAG, "notch screen Rect = " + rect.toShortString()); } } } }); 复制代码`

获取刘海区域信息后就可以根据自己应用的需要，来避开重要的控件。

详情可参考我项目中的代码。

## 参考链接 ##

> 
> 
> 
> [声明受限屏幕支持：声明最大长宽比](
> https://link.juejin.im?target=%255Bhttps%3A%2F%2Fdeveloper.android.com%2Fguide%2Fpractices%2Fscreens-distribution%3Fhl%3Dzh-CN%23MaxAspectRatio%255D(https%3A%2F%2Fdeveloper.android.com%2Fguide%2Fpractices%2Fscreens-distribution%3Fhl%3Dzh-CN%23MaxAspectRatio)
> )
> 
> 
> 
> [Android 8.1 兼容性定义](
> https://link.juejin.im?target=https%3A%2F%2Fsource.android.google.cn%2Fcompatibility%2Fandroid-cdd.html%3Fhl%3Dzh-cn
> )
> 
> 
> 
> [多窗口支持](
> https://link.juejin.im?target=https%3A%2F%2Fdeveloper.android.com%2Fguide%2Ftopics%2Fui%2Fmulti-window.html%3Fhl%3Dzh-CN
> )
> 
> 
> 
> [Support display cutouts](
> https://link.juejin.im?target=https%3A%2F%2Fdeveloper.android.com%2Fguide%2Ftopics%2Fdisplay-cutout%3Fhl%3Dzh-CN
> )
> 
> 
> 
> [华为刘海屏手机安卓O版本适配指导](
> https://link.juejin.im?target=https%3A%2F%2Fdeveloper.huawei.com%2Fconsumer%2Fcn%2Fdevservice%2Fdoc%2F50114
> )
> 
> 
> 
> [OPPO凹形屏适配说明](
> https://link.juejin.im?target=https%3A%2F%2Fopen.oppomobile.com%2Fwiki%2Fdoc%23id%3D10159
> )
> 
> 
> 
> [vivo 全面屏应用适配指南](
> https://link.juejin.im?target=https%3A%2F%2Fdev.vivo.com.cn%2FdocumentCenter%2Fdoc%2F103
> )
> 
> 
> 
> [小米刘海屏水滴屏 Android O 适配](
> https://link.juejin.im?target=https%3A%2F%2Fdev.mi.com%2Fconsole%2Fdoc%2Fdetail%3FpId%3D1293
> )
> 
>