# Android Gradle使用总结（2） #

### Android Gradle使用篇： ###

#### [Android Gradle使用总结（1）]( https://juejin.im/post/5cebe7cf6fb9a07efc49686c ) ####

#### [Android Gradle使用总结（2）]( https://juejin.im/post/5cf5240e51882520724c8301 ) ####

## 一、自定义Android gradle工程 ##

### 1、defaultConfig默认配置 ###

> 
> 
> 
> ` defaultConfig` 是默认的配置，是一个 ` ProductFlavour` 。对于多渠道打包，等情况，如果不针对自定义的是一个 `
> ProductFlavour` 单独配置的话，则默认使用 ` defaultConfig` 的配置
> 
> 

#### （1）applicationId ####

指定生成的app包名，对应值是 ` String` ，如： ` applicationId "com.android.xx"`

#### （2）minSdkVersion ####

App最低支持的Android操作系统版本，对应值是 ` int` 型（即对于sdk 的ApiLevel）如， ` minSdkVersion 25`

#### （3）targetSdkVersion ####

配置当前是基于哪个sdk版本进行开发，可选值与 minSdkVersion一样

#### （4）versionCode ####

内部版本号，对应值是 ` int` 型，用于配置Android App内部版本号，通常用于版本的升级

#### （5）versionName ####

用于配置app的版本名称，对应值是 ` String` ，如"v1.0.0"，让用户知道当前app版本。versionCode是内部使用，versionName是外部使用，一起配合完成app的版本升级控制。

#### （6）testApplicationId ####

配置测试app的包名，默认是 **applicationId + ".test"** ，一般情况下，用默认的即可。

#### （7）testInstrumentationRunner ####

用于配置单元测试时使用的Runner，默认使用android.test.InstrumentationTestRunner。

#### （8）signingConfig ####

**配置默认的签名信息** ，对生成的app进行签名。对应值是 ` SigningConfig` 对象。 具体使用，查看 **配置签名信息** 部分 。

#### （9）proguardFile 和 proguardFiles ####

两者都是配置混淆规则文件，区别是proguardFile接受一个文件对象，而proguardFiles可以同时接受多个。

#### （10）multiDexEnabled ####

是否启动自动拆分多个Dex的功能，用于突破方法超过65535的设置， **后面再具体介绍**

` multiDexEnabled true`

### 2、配置签名信息 ###

> 
> 
> 
> Android Gradle 提供了 ` signingConfigs{}` 配置块，用于生成多个签名配置信息，其类型是 `
> NamedDomainObjectContainer` ，因此，我们在 ` signingConfigs{}` 中定义的都是一个 `
> SigningConfig` 的对象实例。
> 
> 

一个SigningConfig，即签名配置，可配置的元素有： ` storeFile` : 签名文件位置； ` storeType` ：签名证书类型（可不填）； ` storePassword` ：签名证书密码； ` keyAlias` ：签名证书密钥别名； ` keyPassword` ：签名证书中的密钥的密码。如：

` signingConfigs{ release{ //生成了release的签名配置 storeFile file( '../John.jks' ) storePassword '12345678' keyAlias 'key0' keyPassword '12345678' } } 复制代码`
> 
> 
> 
> 
> 注：debug签名文件一般位于$HOME/.android/debug.keystore
> 
> 

对于生成的签名配置，可以应用到 ` defaultConfig` 中的 ` signingConfig` 进行默认签名配置或应用到 ` buildTypes` 中针对构建类型进行签名配置。如：

` buildTypes { release { signingConfig signingConfigs.release //配置release类型签名 minifyEnabled false proguardFiles getDefaultProguardFile( 'proguard-android.txt' ), 'proguard-rules.pro' } } defaultConfig { applicationId "com.example.john.tapeview" minSdkVersion 24 targetSdkVersion 28 versionCode 1 versionName "1.0" signingConfig signingConfigs.release //配置默认签名 testInstrumentationRunner "android.support.test.runner.AndroidJUnitRunner" } 复制代码`

### 3、构建类型（buildTypes） ###

> 
> 
> 
> ` buildTypes{}` 和 ` signingConfigs{}` 一样，都是Android 的一个方法，接收到也是域对象 `
> NamedDomainObjectContainer` ，添加的每一个都是 ` BuildType` 类型。
> 
> 

每一个BuildType都会生成一个SourceSet，默认位置是src//，因此，可以单独为其指定Java源代码和res资源等。且，新增的BuildType不能命名为 main 或 androidTest（因android默认已经生成），且相互之间不能同名。

#### （1）applicationIdSuffix ####

用于配置基于applicationId的后缀，如applicationId为 com.John.gradle.sample , 在debug的buildType中指定 ` applicationIdSuffix ".debug"` ，则生成的debug的包名为com.John.gradle.sample.debug

#### （2）signingConfig ####

配置签名信息，使用同前面提到的 defaultConfig一样。

#### （3）启动zipAlign优化 ####

zipalign是android提供的整理优化apk的工具，能提高系统和应用的运行效率，更快读写apk资源，降低内存，所以，对于发布的apk，一般开启 ` zipAlignEnabled true //true为开启`

#### （4）资源清理shrinkResources ####

用于配置是否自动清理未使用的资源，true为开启，false为关闭。需 **结合混淆** 使用。

#### （5）使用混淆 minifyEnabled ####

启用混淆可以优化代码，同时结合 ` shrinkResources` 清理资源，可以缩小apk包，还可以混淆代码。一般发布的都是要混淆的。

开启混淆： ` minifyEnabled true //true为开启` ；使用 ` proguardFile` 和 ` proguardFiles` 设置混淆规则文件，两者区别如前面defaultConfig中说明。

**注：**对于多渠道打包，每个productFlavor都可以单独配置混淆规则文件（通过各自的 ` proguardFile` 和 ` proguardFiles` ）。

#### （6）multiDexEnabled ####

是否启动自动拆分多个Dex的功能，用于突破方法超过65535的设置， **后面再具体介绍**

` multiDexEnabled true`

#### （7）debuggable 和 jniDebuggable ####

debuggable：是否生成一个可供调试的apk，对应值为boolean类型；

jniDebuggable：是否生成一个可供调试Jni（C/C++）代码的apk，对应值为boolean类型；

## 二、Android gradle高级自定义 ##

### 1、批量修改生成的apk文件名 ###

Android Gradle提供了3个属性，applicationVariants（仅仅适用于Android 应用Gradle插件），和libraryVariants（仅仅适用于Android库gradle插件），testVariants（以上两种都适用）

applicationVariants是一个集合，每一个元素都是一个生成的产物，即 xxxRelease 和 xxxDebug 等（即生成apk） ，它有一个outputs作为输出集合。遍历，若名字以.apk结尾，那么就是要修改的apk输出。

` //动态改变生成的apk的名字： 项目名_渠道名_v版本名_构建日期.apk applicationVariants.all { variant -> variant.outputs.all { output -> if (output.outputFile != null && output.outputFile.name.endsWith( '.apk' ) && 'release' ==(variant.buildType.name)) { def flavorName = variant.flavorName.startsWith( "_" ) ? variant.flavorName.substring(1) : variant.flavorName def apkFileName = "Example92_${flavorName}_v${variant.versionName}_${buildTime()}.apk" outputFileName = apkFileName } } } ... def static buildTime() { def date = new Date() def formattedDate = date.format( 'yyyyMMdd' ) return formattedDate } 复制代码`

### 2、多方式配置 VersionCode 和 VersionName ###

#### （1）通过应用脚本插件（apply from） ####

* 新建一个 ` version.gradle` （当前存放在项目路径下的gradle文件夹中），并设置自定义属性

![](https://user-gold-cdn.xitu.io/2019/6/1/16b10d4de7bd4bde?imageView2/0/w/1280/h/960/ignore-error/1)

` ext{ appVersionCode = 1 appVersionName = '1.0.2' } 复制代码` * 在rootProject中应用该脚本文件
` apply from : 'gradle/version.gradle' 复制代码` * 可在任意子工程中引用该设置：

![](https://user-gold-cdn.xitu.io/2019/6/1/16b10d81daa5e17f?imageView2/0/w/1280/h/960/ignore-error/1)

#### （2）直接在属性文件（xxx.properties）中定义使用 ####

* 在工程目录下的gradle.properties中定义 ![](https://user-gold-cdn.xitu.io/2019/6/1/16b10d95d9f303cf?imageView2/0/w/1280/h/960/ignore-error/1) 3.可在任意子工程中引用该设置：

![](https://user-gold-cdn.xitu.io/2019/6/1/16b10da4442522a5?imageView2/0/w/1280/h/960/ignore-error/1)

#### （3）从 git 的 tag 上获取 VersionName 和 versionCode ####

` /** * 以git tag的数量作为其版本号 * @ return tag的数量 */ def static getAppVersionCode (){ def stdout = new ByteArrayOutputStream() exec { command Line 'git' , 'tag' , '--list' standardOutput = stdout } return stdout.toString().split( "\n" ).size() } /** * 从git tag中获取应用的版本名称 * @ return git tag的名称 */ def static getAppVersionName (){ def stdout = new ByteArrayOutputStream() exec { command Line 'git' , 'describe' , '--abbrev=0' , '--tags' standardOutput = stdout } return stdout.toString().replaceAll( "\n" , "" ) } 复制代码`

### 3、动态配置AndroidManifest文件 ###

> 
> 
> 
> 在构建过程，动态修改 ` AndroidManifest` 文件中的内容
> 
> 

android gradle 提供了非常便捷的方法替换AndroidManifest文件中的内容，即manifestPlaceholder、Manifest占位符。

` manifestPlaceholders` 是 **ProductFlavor的一个属性，Map类型** ，所以可以同时配置多个占位符。

* 定义占位符：
`... productFlavors{ google{ manifestPlaceholders.put( "UMENG_CHANNEL" , "google" ) } baidu{ manifestPlaceholders.put( "UMENG_CHANNEL" , "baidu" ) } } ... 复制代码` * 在AndroidManifes中使用这个占位符 "UMENG_CHANNEL"
`... <application> ... <meta-data android:value="${UMENG_CHANNEL}" android:name = "UMENG_CHANNEL"/> ... </application> ... 复制代码` * 亦可以迭代修改占位符
`... productFlavors{ google{ } baidu{ } } productFlavors.all{flavor-> manifestPlaceholders.put( "UMENG_CHANNEL" ,flavor.name) } ... 复制代码`

### 4、自定义BuildConfig ###

Android gradle提供了 ` buildConfigField(Srting type,String name ,String value)` 来实现添加常量到 ` BuildConfig` 中。 **一般在 ` productFlavors` 或 ` buildTypes` 中使用** :

**注：对于value是String类型的，双引号一定得加上，这里写了什么，buildConfig会原封不动放上去**

`... productFlavors{ google{ //对于value是String类型的，双引号一定得加上，这里写了什么，buildConfig会原封不动放上去 buildConfigField 'String' , 'name' , '"google"' } baidu{ buildConfigField 'String' , 'name' , '"baidu"' } } ... 复制代码`

![](https://user-gold-cdn.xitu.io/2019/6/1/16b113537eb65731?imageView2/0/w/1280/h/960/ignore-error/1)

### 5、动态添加自定义的资源 ###

Android gradle提供了 ` resValue(String type,String name,String value)` 来 添加资源。 **只能在 ` productFlavors` 或 ` buildTypes` 中使用**

**注：这里写了什么，编译出来后也会原封不动放上去**

![](https://user-gold-cdn.xitu.io/2019/6/1/16b113e1078dc813?imageView2/0/w/1280/h/960/ignore-error/1)

### 6、Dex选项配置 ###

` android { ... dexOptions { incremental true //android studio3.x以上已废弃，是否开启dx增量模式 javaMaxHeapSize '4g' //提高java执行dx命令时的最大分配内存 jumboMode true //强制开启jumboMode模式，让方法超过65535也构建成功 pre } ... } 复制代码`

### 7、编译选项和adb操作选项 ###

` android { ... compileOptions { encoding = 'utf-8' sourceCompatibility JavaVersion.VERSION_1_8 targetCompatibility JavaVersion.VERSION_1_8 } adbOptions{ timeOutInMs = 5000 //设置超时时间，单位 ms installOptions '-r' , '-s' //安装选项，一般不使用 } ... } 复制代码`

### 8、突破方法超过65535的限制 ###

* 

在 ` defaultConfig` 或 ` buildTypes` 或 ` productFlavors` 中，开启 ` multiDexEnabled` ： ` multiDexEnabled true`

* 

增加依赖： ` implementation 'com.android.support:multidex:1.0.3'`

* 

在自定义的 ` Application` 中设置（5.0以上天然支持）：

使自定义的Application继承 ` MultiDexApplication`

![](https://user-gold-cdn.xitu.io/2019/6/1/16b1162752bee901?imageView2/0/w/1280/h/960/ignore-error/1)

**或直接在java代码中修改：**

`... @Override protected void attachBaseContext (Context base) { super.attachBaseContext(base); MultiDex.install( this ); } ... 复制代码`

## 多渠道构建 ##

> 
> 
> 
> 一个构建产物（apk），即Build Variant = BuildType +
> ProductFlavor，其中ProductFlavor又进一步划分dimensions，即 **Build Variant = BuildType
> + ProductFlavor's dimensions**
> 
> 

### 1、多渠道构建ProductFlavor属性 ###

`... productFlavors{ google{ applicationId "com.john.xxx" } } ... 复制代码`

#### （1）applicationId，跟前面所述一样。 ####

#### （2）manifestPlaceholders，跟前面所述一样。 ####

#### （3）multiDexEnabled，跟前面所述一样。 ####

#### （4）proguardFiles，跟前面所述一样。 ####

#### （5）signingConfig，跟前面所述一样。 ####

#### （6）testApplicationId，跟前面所述一样。 ####

#### （7）versionCode 和 versionName，跟前面所述一样。 ####

#### （8）dimension ####

> 
> 
> 
> 作为 ` ProductFlavor` 的维度，可理解为 ` ProductFlavor` 的分组
> 
> 

如，free 和 paid 属于version分组，而x86 和 arm属于架构分组：

` android{ ... flavorDimensions "version" , "abi" //定义分组 productFlavors{ free{ dimension 'version' } paid{ dimension 'version' } x86{ dimension 'abi' } arm{ dimension 'abi' } } ... } 复制代码`

对于以上的设置， **Build Variant = BuildType + ProductFlavor's dimensions** ，则生成的assemble任务， 即variant：

* ArmFreeDebug
* ArmFreeRelaese
* ArmPaidDebug
* ArmPaidRelease
* X86FreeDebug
* X86FreeRelaese
* X86PaidDebug
* X86PaidRelease

## 参考链接 ##

[Android Developer官网]( https://link.juejin.im?target=https%3A%2F%2Fdeveloper.android.com%2Fstudio%2Fbuild%3Fhl%3Dzh-cn )