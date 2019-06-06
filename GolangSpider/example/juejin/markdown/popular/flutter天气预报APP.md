# flutter天气预报APP #

## 前言 ##

这是一个使用flutter写的天气预报APP，主要使用了以下几个插件，入门级练练手。 ` dio` :网络请求插件，用于获取天气信息 ` fluttertoast` :弹出toast提示信息 ` shared_preferences` :简单的数据存储，用于保存设置过的天气预报信息 ` intl` :日期格式化 项目GitHub地址： [d9l_weather]( https://link.juejin.im?target=%255Bhttps%3A%2F%2Fgithub.com%2Fhuang-weilong%2Fd9l_weather%255D(https%3A%2F%2Fgithub.com%2Fhuang-weilong%2Fd9l_weather) )

## 界面 ##

首先搜集一些天气预报APP的设计稿，确定一下自己的界面。看到有很多好看的，但是并不想做的太复杂。于是选择了一些简洁的，背景图也都去掉了。然后在PS里大概出一个界面，如下：

![首页图.jpg](https://user-gold-cdn.xitu.io/2019/6/5/16b263a396fd0f50?imageView2/0/w/1280/h/960/ignore-error/1) 然后分析一下这个页面，大主体就是一个Column布局排列下来，中间穿插Row布局。详细的布局这里就不写，可以查看源码 [home_page.dart]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fhuang-weilong%2Fd9l_weather%2Fblob%2Fmaster%2Flib%2Fhome_page.dart ) 界面堆好之后，再做一个搜索城市的页面，一个搜索框加列表就可以了，怎么简单怎么来。页面通过右上角的设置按钮进入。

## 接口 ##

页面都写好之后就需要把数据替换成真实数据了，这里使用了 [和风天气]( https://link.juejin.im?target=https%3A%2F%2Fwww.heweather.com%2F ) 的API获取天气数据，注册之后就能使用。但是普通用户有些接口是不能用的，但是对这个APP来说，能够查到天气信息以及足够了。 新建一个 ` dio_client.dart` 文件，里面放所有API请求方法，这里写成单例模式，如下。

` class DioClient { factory DioClient() => _getInstance(); static DioClient get instance => _getInstance(); static DioClient _instance; // 单例对象 static DioClient _getInstance() { if (_instance == null ) { _instance = DioClient._internal(); } return _instance; } DioClient._internal(); } 复制代码`

在main函数中初始化单例对象

` DioClient(); 复制代码`

使用方法

` DioClient().getRealTimeWeather(); 复制代码`

如获取实时天气的方法如下：

` Future<RealTimeWeather> getRealTimeWeather( String cid) async { String url = rootUrl + '/now' ; try { Response response = await Dio(). get (url, options: options, queryParameters: { 'location' : cid, // 查询的城市id 'key' : key, }); // 根据API返回的参数定义的model RealTimeWeather realTimeWeather; realTimeWeather = RealTimeWeather.fromJson(response.data[ 'HeWeather6' ].first); if (realTimeWeather.status.contains( 'permission' )) { return realTimeWeather; } realTimeWeather.basic = Basic.fromJson(realTimeWeather.mBasic); realTimeWeather.update = Update.fromJson(realTimeWeather.mUpdate); realTimeWeather.now = Now.fromJson(realTimeWeather.mNow); return realTimeWeather; } catch (e) { print ( 'getRealTimeWeather error= $e' ); return null ; } } 复制代码`

## shared_preferences ##

在搜索完一个城市的天气后，需要把这个城市的 ` id` 保存在 ` shared_preferences` 中，这样关闭app下次再打开的时候才能显示上一次查询的城市天气，或者需要保存多个城市天气预报的时候，也可以保存。 保存只需要一行代码：

` SpClient.sp.setString( 'cid' , cid); 复制代码`

` shared_preferences` 的使用也是使用了单例模式，和 ` dio_client.dart` 一样

## 最后 ##

这个项目很简单，也只是用了很少的东西，主要是练练手吧。也没太多东西能够介绍的。后面有时间的话会继续完善，比如加上国际化、使用状态管理，多城市保存等等。感兴趣的可以关注一下。GitHub给star支持，谢谢！ 最后再放一下本项目的GitHub地址 [d9l_weather]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fhuang-weilong%2Fd9l_weather )

### 其他 ###

#### [flutter版的文件管理器项目地址]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fhuang-weilong%2Fflutter_file_manager.git ) ####

#### [flutter入门widget的使用。]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fhuang-weilong%2Fflutter_widgets ) 带你认识flutter widgets。根据flutter中文网widgets目录进行编写的一个库。 ####