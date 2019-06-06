# 【Flutter】如何写一个Flutter自动打包成iOS代码模块的脚本 #

相信很多使用原生+Flutter的iOS项目都会遇到混合开发的集成问题，也有大神写了一些解决方案，下面就记录一下我的心路历程：

## 前期准备 ##

开始之前，我先拜读了一些大神的文章（这里只挑出对我帮助最大的）：

[Flutter混合开发组件化与工程化架构]( https://link.juejin.im?target=http%3A%2F%2Fzhengxiaoyong.com%2F2018%2F12%2F16%2FFlutter%25E6%25B7%25B7%25E5%2590%2588%25E5%25BC%2580%25E5%258F%2591%25E7%25BB%2584%25E4%25BB%25B6%25E5%258C%2596%25E4%25B8%258E%25E5%25B7%25A5%25E7%25A8%258B%25E5%258C%2596%25E6%259E%25B6%25E6%259E%2584%2F )

[混沌初始，iOS现有项目集成Flutter]( https://link.juejin.im?target=https%3A%2F%2Fwww.jianshu.com%2Fp%2Ff78fe35e5bbf )

## 方案筛选 ##

经过探索，结合项目的实际情况（我司的项目采用模块化开发，pods方式集成），有下面的两个方案：

* 使用google的集成方案
* 将flutter编译后的产物打包到一个新的子模块中，并在其中实现对应的接口和交互逻辑。

无论是从使用方便的角度还是对代码的侵入程度来看，采用方案2都是顺理成章的。

## 实现方式 ##

写脚本之前，先来看看打包步骤（其实就是收集编译产物）

### 打包步骤 ###

打包的步骤如下：

1、flutter build ios。

2、进入对应flutter项目的../ios(或者.ios)/Flutter/文件夹，找到App.framwork和Flutter.framework以及FlutterPluginRegistrant文件夹，拷贝到子模块中。

3、进入对应flutter项目的build/ios/Debug-iphoneos(或者Release-iphoneos)，拷贝各个依赖库的.a文件到子模块中。

4、如果项目有依赖到第三方的插件（一般来说都会有），需要根据.flutter-plugins文件中的路径到对应模块的源代码的ios/Classes中拷贝各个依赖库的.h文件（这一步可以说是相当繁琐了）。

### 使用自动化脚本执行上面的操作 ###

以下代码均在flutter工程的目录下操作（请确保flutter的插件没有源码依赖）：

首先是打包：

` echo "===清理flutter历史编译===" flutter clean echo "===重新生成plugin索引===" flutter packages get echo "===生成App.framework和flutter_assets===" flutter build ios --debug # 或者 flutter build ios --release 复制代码`

然后是拷贝App.framework和Flutter.framework：

` framework_dir=... echo "===copy App.framework和Flutter.framework===" cp -r "./.ios/Flutter/App.framework" $framework_dir cp -r "./.ios/Flutter/engine/Flutter.framework" $framework_dir 复制代码`

拷贝注册器：

` classes_dir=... echo "===copy注册器：FlutterPluginRegistrant===" regist_dir= "./.ios/Flutter/FlutterPluginRegistrant/Classes/" model_class= ${classes_dir} /flutter/FlutterPluginRegistrant mkdir -p $model_class cp -r $regist_dir $model_class 复制代码`

接下来是各个插件的拷贝，这里有一个小坑，由于各个插件是被打包成.a的形式进行引入，那么很可能导致模拟器上无法运行的问题，需要打一个真机包和一个模拟器包，并将它们进行合并，才能在使用过程中用模拟器进行调试：

` # 合成的.a文件缓存 temp_dir=... mkdir -p ${temp_dir} current_path= " $PWD " # 执行clean并重新编译pods部分 cd.ios/Pods /usr/bin/env xcrun xcodebuild clean /usr/bin/env xcrun xcodebuild build -configuration Release ARCHS= 'arm64 armv7' BUILD_AOT_ONLY=YES VERBOSE_SCRIPT_LOGGING=YES -workspace Runner.xcworkspace -scheme Runner BUILD_DIR=../build/ios -sdk iphoneos # 遍历.flutter-plugins文件 cat .flutter-plugins | while read line do array=( ${line//=/ } ) plugin_name= ${array[0]} echo "===修改注册器（修正引用）===" perl -pi -e "s|\< ${plugin_name} \/|\"|g" ${model_class} /GeneratedPluginRegistrant.m perl -pi -e "s|.h\>|.h\"|g" ${model_class} /GeneratedPluginRegistrant.m temp_library= ${temp_dir} /lib ${plugin_name}.a echo ">>>生成lib ${plugin_name}.a<<<" cd.ios/Pods /usr/bin/env xcrun xcodebuild build -configuration Release ARCHS= 'arm64 armv7' -target ${plugin_name} BUILD_DIR=../../build/ios -sdk iphoneos -quiet /usr/bin/env xcrun xcodebuild build -configuration Debug ARCHS= 'x86_64' -target ${plugin_name} BUILD_DIR=../../build/ios -sdk iphonesimulator -quiet echo ">>>合并lib ${plugin_name}.a<<<" lipo -create "../../build/ios/Debug-iphonesimulator/ ${plugin_name} /lib ${plugin_name}.a" "../../build/ios/Release-iphoneos/ ${plugin_name} /lib ${plugin_name}.a" -o $temp_library cd $current_path if [[ -f " $temp_library " ]]; then echo "===copy ${plugin_name} ===" plugin= ${framework_dir} / ${plugin_name} rm -rf $plugin mkdir -p $plugin cp -f $temp_library $plugin classes= ${array[1]} ios/Classes class= $dest_dir /Classes/flutter/ ${plugin_name} rm -rf $class mkdir -p $class for header in `find " $classes " -name *.h`; do cp -f $header $class done fi done rm -rf ${temp_dir} fi 复制代码`

可能你会注意到上面的 ` ===修改注册器（修正引用）===` 内容，这是由于我们将flutter的插件部分的文件直接打包成静态库，并将其.h文件全部导入到模块内部了，因此自动生成的 ` #import <xxx/xxx.h>` 模式不适用了（会报错哟），需要改成 ` #import "xxx.h"` 的形式进行引用。

到此，flutter的自动打包脚本基本完成，上面的各个文件路径需要根据实际情况进行调整，改为适合自己的项目的路径。

希望这篇能够帮助大家少走弯路。