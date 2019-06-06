# Android AOSP基础（四）Source Insight和Android Studio导入系统源码 #

> 
> 
> 
> 本文首发于微信公众号「刘望舒」
> 
> 

关联系列
[Android AOSP基础系列]( https://link.juejin.im?target=http%3A%2F%2Fliuwangshu.cn%2Ftags%2FAOSP%25E5%259F%25BA%25E7%25A1%2580%2F )
[Android系统启动系列]( https://link.juejin.im?target=http%3A%2F%2Fliuwangshu.cn%2Ftags%2FAndroid%25E7%25B3%25BB%25E7%25BB%259F%25E5%2590%25AF%25E5%258A%25A8%2F )
[应用进程启动系列]( https://link.juejin.im?target=http%3A%2F%2Fliuwangshu.cn%2Ftags%2FAndroid%25E5%25BA%2594%25E7%2594%25A8%25E8%25BF%259B%25E7%25A8%258B%2F )
[Android深入四大组件系列]( https://link.juejin.im?target=http%3A%2F%2Fliuwangshu.cn%2Ftags%2FAndroid%25E6%25B7%25B1%25E5%2585%25A5%25E5%259B%259B%25E5%25A4%25A7%25E7%25BB%2584%25E4%25BB%25B6%2F )
[Android深入理解Context系列]( https://link.juejin.im?target=http%3A%2F%2Fliuwangshu.cn%2Ftags%2FAndroid%25E6%25B7%25B1%25E5%2585%25A5%25E7%2590%2586%25E8%25A7%25A3Context%2F )
[Android深入理解JNI系列]( https://link.juejin.im?target=http%3A%2F%2Fliuwangshu.cn%2Ftags%2FAndroid%25E6%25B7%25B1%25E5%2585%25A5%25E7%2590%2586%25E8%25A7%25A3JNI%2F )
[Android解析WindowManager]( https://link.juejin.im?target=http%3A%2F%2Fliuwangshu.cn%2Ftags%2FWindowManager%2F )
[Android解析WMS系列]( https://link.juejin.im?target=http%3A%2F%2Fliuwangshu.cn%2Ftags%2FWindowManagerService%2F )
[Android解析AMS系列]( https://link.juejin.im?target=http%3A%2F%2Fliuwangshu.cn%2Ftags%2FActivityManagerService%2F )
[Android包管理机制系列]( https://link.juejin.im?target=http%3A%2F%2Fliuwangshu.cn%2Ftags%2FAndroid%25E5%258C%2585%25E7%25AE%25A1%25E7%2590%2586%25E6%259C%25BA%25E5%2588%25B6%2F )
[Android输入系统系列]( https://link.juejin.im?target=http%3A%2F%2Fliuwangshu.cn%2Ftags%2FAndroid%25E8%25BE%2593%25E5%2585%25A5%25E7%25B3%25BB%25E7%25BB%259F%2F )

### **前言** ###

在上一篇文章 [Android AOSP基础（三）Android系统源码的整编和单编]( https://link.juejin.im?target=http%3A%2F%2Fliuwangshu.cn%2Fframework%2Faosp%2F3-compiling-aosp.html ) 中，我们对系统源码进行了编译，这篇文章我们接着来学习如何将系统源码导入到编辑器中，以便于查看和调试源码。关于查看源码，可以使用Android Studio、Eclipse、Sublime、Source Insight等软件，这里我推荐使用Source Insight，但是有的同学可能不是很习惯，而且Source Insight是Windows平台的软件，Mac平台用不了，那么使用Android Studio是一个不错的选择，而且使用Android Studio还可以调试源码。

### **1. Source Insight导入系统源码** ###

在《Android进阶解密》的第一章，我讲解了Source Insight如何导入系统源码，可能有的同学没有买这本书，这里再来讲一遍。 Source Insight只能查看源码，不能调试源码，如果只想在Source Insight中查看源码，可以直接从百度网盘： [pan.baidu.com/s/1ngsZs]( https://link.juejin.im?target=https%3A%2F%2Fpan.baidu.com%2Fs%2F1ngsZs ) 将源码下载下来。如果想在Android Studio中查看源码，那么最好还是在Linux环境下将AOSP源码下载下来。

#### **新建源码项目** ####

安装软件后，首先要新建源码项目。通过菜单项Project→New Project会弹出提示框。

![](https://user-gold-cdn.xitu.io/2019/6/6/16b29833a5eb6d1f?imageView2/0/w/1280/h/960/ignore-error/1) 这里我们指定源码项目的名称为Android_8.0.0，然后点击OK按钮进入“New Project Settings”界面。 ![](https://user-gold-cdn.xitu.io/2019/6/6/16b29833b42af421?imageView2/0/w/1280/h/960/ignore-error/1)

上图箭头指向的Browse按钮来选择本地系统源码所在的路径，比如我的系统源码路径为：D:/Android/android-8.0.0_r1 。选择好加载路径后点击OK按钮会进入“Add and Remove Project Files”界面，在这个界面可以向项目中添加整个Android系统源码，也可以只把源码部分目录添加到项目中，以后再根据需要添加其他目录。如果向项目添加整个Android系统源码加载时会非常慢，这里我们只添加如下源码目录：frameworks/、libcore/、packages/、system/、art/和libnativehelper/，这几个目录基本上可以满足日常的系统源码阅读了，如下图所示。

![](https://user-gold-cdn.xitu.io/2019/6/6/16b29833b3b8389f?imageView2/0/w/1280/h/960/ignore-error/1) 点击Add Tree按钮就会将选择的目录源码加载到Android_8.0.0项目中，这个时候会弹出加载进度条，加载完毕后点击窗口的关闭按钮就可以了。

#### **定位文件** ####

Source Insight的定位文件功能十分强大，我们只需要知道源码文件名就可以轻松找到它，比如我们要找MediaPlayer.java，只要在文件搜索框输入MediaPlayer.java即可：

![](https://user-gold-cdn.xitu.io/2019/6/6/16b29833b399412b?imageView2/0/w/1280/h/960/ignore-error/1)

#### **全局搜索** ####

Source Insight另一个好用的功能就是全局搜索，默认快捷键为：CTRL+/，或者点击最上面工具栏类似R的图标。在Search in的输入选项中我们可以自定义搜索的范围，比如我们想查找所有Java文件中引用MediaPlayer类的情况，就可以像下图一样进行操作。

![](https://user-gold-cdn.xitu.io/2019/6/6/16b29833abeb3d40?imageView2/0/w/1280/h/960/ignore-error/1)

当然，Source Insights的功能远不只以上几种，相信随着使用次数的增多，你就会熟练掌握它的大部分功能，这里就不过多介绍了。

### **2. Android Studio导入系统源码** ###

Source Insight导入源码不需要对源码进行编译，但是Android Studio导入整个系统源码需要对源码进行编译，生成AS的项目配置文件。

#### **生成AS的项目配置文件** ####

如果你整编过源码，查看out/host/linux-x86/framework/idegen.jar是否存在，如果不存在，进入源码根目录执行如下的命令：

` source build/envsetup.sh lunch [选择整编时选择的参数或者数字] mmm development/tools/idegen/ 复制代码`

如果没整编过源码，可以直接执行如下命令单编idegen模块：

` source build/ensetup.sh make idegen 复制代码`

关于Android系统源码的编译可以查看 [Android AOSP基础（三）Android系统源码的整编和单编]( https://link.juejin.im?target=http%3A%2F%2Fliuwangshu.cn%2Fframework%2Faosp%2F3-compiling-aosp.html ) 这篇文章。

idegen模块编译成功后，会在 out/host/linux-x86/framework目录下生成idegen.jar，执行如下命令：

` sudo development/tools/idegen/idegen.sh 复制代码`

这时会在源码根目录生成android.iml 和 android.ipr 两个文件，这两个文件一般是只读模式，这里建议改成可读可写，否则，在更改一些项目配置的时候可能会出现无法保存的情况。

` sudo chmod 777 android.iml sudo chmod 777 android.ipr 复制代码`

#### **配置AS的项目配置文件** ####

由于要将所有源码导入AS会导致第一次加载很慢，可以在android.iml中修改excludeFolder配置，将不需要看的源码排除掉。等源码项目加载完成后，还可以通过AS对Exclude的Module进行调整。如果你的电脑的性能很好，可以不用进行配置。 在android.iml中搜索excludeFolder，在下面加入这些配置。

` < excludeFolder url = "file://$MODULE_DIR$/bionic" /> < excludeFolder url = "file://$MODULE_DIR$/bootable" /> < excludeFolder url = "file://$MODULE_DIR$/build" /> < excludeFolder url = "file://$MODULE_DIR$/cts" /> < excludeFolder url = "file://$MODULE_DIR$/dalvik" /> < excludeFolder url = "file://$MODULE_DIR$/developers" /> < excludeFolder url = "file://$MODULE_DIR$/development" /> < excludeFolder url = "file://$MODULE_DIR$/device" /> < excludeFolder url = "file://$MODULE_DIR$/docs" /> < excludeFolder url = "file://$MODULE_DIR$/external" /> < excludeFolder url = "file://$MODULE_DIR$/hardware" /> < excludeFolder url = "file://$MODULE_DIR$/kernel" /> < excludeFolder url = "file://$MODULE_DIR$/out" /> < excludeFolder url = "file://$MODULE_DIR$/pdk" /> < excludeFolder url = "file://$MODULE_DIR$/platform_testing" /> < excludeFolder url = "file://$MODULE_DIR$/prebuilts" /> < excludeFolder url = "file://$MODULE_DIR$/sdk" /> < excludeFolder url = "file://$MODULE_DIR$/system" /> < excludeFolder url = "file://$MODULE_DIR$/test" /> < excludeFolder url = "file://$MODULE_DIR$/toolchain" /> < excludeFolder url = "file://$MODULE_DIR$/tools" /> < excludeFolder url = "file://$MODULE_DIR$/.repo" /> 复制代码`

#### **导入系统源代码到AS中** ####

在AS安装目录的bin目录下，打开studio64.vmoptions文件，根据自己电脑的实际情况进行设置，这里修改为如下数值：

` -Xms1024m -Xmx1024m 复制代码`

如果你是在VirtualBox中下载的系统源码，那么将VirtualBox中的系统源码拷贝到共享文件夹中，这样源码就会自动到Windows或者Mac上，如果你不知道如何设置VirtualBox共享文件夹，可以查看 [Android AOSP基础（一）VirtualBox 安装 Ubuntu]( https://link.juejin.im?target=http%3A%2F%2Fliuwangshu.cn%2Fframework%2Faosp%2F1-install-ubuntu.html ) 这篇文章。 通过AS的Open an existing Android Studio project选项选择android.ipr 就可以导入源码，这里我用了大概7分钟就导入完毕。导入后工程目录切换为Project选项就可以查看源码。

![](https://user-gold-cdn.xitu.io/2019/6/6/16b29833adc6f6d6?imageView2/0/w/1280/h/960/ignore-error/1)

#### **配置项目的JDK、SDK** ####

由于我们下载的是9.0的AOSP源码，SDK版本也应该对应为API 28，如果没有就去SDK Manager下载即可。 点击File -> Project Structure-->SDKs配置项目的JDK、SDK。 创建一个新的JDK,这里取名为1.8(No Libraries)，删除其中classpath标签页下面的所有jar文件。

![](https://user-gold-cdn.xitu.io/2019/6/6/16b29834756ef22c?imageView2/0/w/1280/h/960/ignore-error/1)

接着设置将Android SDK的Java SDK设置为1.8(No Libraries)，这样Android源码使用的Java就是Android源码中的。

![](https://user-gold-cdn.xitu.io/2019/6/6/16b29834a64f2a37?imageView2/0/w/1280/h/960/ignore-error/1)

确保的项目的SDK为源码对应的SDK。

![](https://user-gold-cdn.xitu.io/2019/6/6/16b29834ad537fe6?imageView2/0/w/1280/h/960/ignore-error/1)

#### **Exclude不需要的代码目录** ####

File -> Project Structure -> Modules中可以通过Excluded来筛选代码目录，比如我们选择bionic目录，点击Excluded，bionic目录会变为橙色，bionic字段会出现在右侧视图中，说明该目录已经被Excluded掉，通俗来讲就是被排除在工程之外。如果不希望bionic目录被Excluded掉，再次点击Excluded，bionic目录会变为灰色。

![](https://user-gold-cdn.xitu.io/2019/6/6/16b29834b3184440?imageView2/0/w/1280/h/960/ignore-error/1)

### **总结** ###

这篇我们学习了Source Insight和Android Studio导入系统源码的方法，但是具体的查看源码的方式没有讲解，这些需要读者在使用中逐步去掌握，下一篇我们会学习如何使用Android Studio去调试系统源码。

分享大前端、Android、Java等技术，助力5万程序员成长进阶。

![](https://user-gold-cdn.xitu.io/2018/8/21/1655afa15727cd03?imageView2/0/w/1280/h/960/ignore-error/1)