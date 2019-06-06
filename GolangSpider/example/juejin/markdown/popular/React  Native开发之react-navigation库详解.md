# React Native开发之react-navigation库详解 #

![](https://user-gold-cdn.xitu.io/2019/6/4/16b1fc6db1a06a16?imageView2/0/w/1280/h/960/ignore-error/1) 众所周知，在多页面应用程序中，页面的跳转是通过路由或导航器来实现的。在0.44版本之前，开发者可以直接使用官方提供的Navigator组件来实现页面的跳转，不过从0.44版本开始，Navigator被官方从react native的核心组件库中剥离出来，放到react-native-deprecated-custom-components的模块中。 如果开发者需要继续使用Navigator，则需要先使用yarn add react-native-deprecated-custom-components命令安装后再使用。不过，官方并不建议开发者这么做，而是建议开发者直接使用导航库react-navigation。react-navigation是React Native社区非常著名的页面导航库，可以用来实现各种页面的跳转操作。 目前，react-navigation支持三种类型的导航器，分别是StackNavigator、TabNavigator和DrawerNavigator。具体区别如下：

* StackNavigator：包含导航栏的页面导航组件，类似于官方的Navigator组件。
* TabNavigator：底部展示tabBar的页面导航组件。
* DrawerNavigator：用于实现侧边栏抽屉页面的导航组件。

需要说明的是，由于react-navigation在3.x版本进行了较大的升级，所以在使用方式上与2.x版本会有很多的不同。 和其他的第三方插件库一样，使用之前需要先在项目汇中添加react-navigation依赖，安装的命令如下：

` yarn add react-navigation //或者 npm install react-navigation --save 复制代码`

安装完成之后，可以在package.json文件的dependencies节点看到react-navigation的依赖信息。

` "react-navigation" : "^3.8.1" 复制代码`

由于react-navigation依赖于react-native-gesture-handler库，所以还需要安装react-native-gesture-handler，安装的命令如下：

` yarn add react-native-gesture-handler //获取 npm install --save react-native-gesture-handle 复制代码`

同时，由于react-native-gesture-handler需要依赖原生环境，所以在需要使用link命令链接原生依赖，命令如下：

` react-native link react-native-gesture-handler 复制代码`

为了保证react-native-gesture-handler能够成功的运行在Android系统上，需要在Android工程的MainActivity.java中添加如下代码：

` public class MainActivity extends ReactActivity { ... @Override protected ReactActivityDelegate createReactActivityDelegate () { return new ReactActivityDelegate(this, getMainComponentName()) { @Override protected ReactRootView createRootView () { return new RNGestureHandlerEnabledRootView(MainActivity.this); } }; } } 复制代码`

然后，就可以使用react-navigation进行页面导航功能开发，如图7-12所示，是使用createStackNavigator实现页面导航的示例。

![在这里插入图片描述](https://user-gold-cdn.xitu.io/2019/6/4/16b1fc66ffe5ca05?imageView2/0/w/1280/h/960/ignore-error/1) 在createStackNavigator模式下，为了方便对页面进行统一管理，首先新建一个RouterConfig.js文件，并使用createStackNavigator注册页面。对于应用的初始页面还需要使用initialRouteName进行申明。同时，导航器栈还需要使用createAppContainer函数进行包裹。例如：

` import {createAppContainer,createStackNavigator} from 'react-navigation' ; import MainPage from './MainPage' import DetailPage from "./DetailPage" ; const AppNavigator = createStackNavigator({ MainPage: MainPage, DetailPage:DetailPage },{ initialRouteName: "MainPage" , }, ); export default createAppContainer(AppNavigator); 复制代码`

其中，createStackNavigator用于配置栈管理的页面，它支持的配置选项有：

* path：路由中设置的路径映射配置。
* initialRouteName：设置栈管理方式的默认页面，且此默认页面必须是路由配置中的某一个。
* initialRouteParams：初始路由参数。
* defaultNavigationOptions：用于配置导航栏的默认导航选项。
* mode：定义渲染和页面跳转的样式，选项有card和modal，默认为card。
* headerMode：定义返回上级页面时动画效果，选项有float、screen和none。

最后，在入口文件中以组件的方式引入StackNavigatorPage.js文件即可。例如：

` import StackNavigatorPage from './src/StackNavigatorPage' export default class App extends Component<Props> { render () { return ( <StackNavigatorPage/> ); } } 复制代码`

要实现页面的栈管理功能或跳转功能，还需要再至少新建两个子页面，例如MainPage.js和DetailPage.js。

` export default class MainPage extends PureComponent { static navigationOptions = { header: null, //默认页面去掉导航栏 }; render () { const {navigate} = this.props.navigation; return ( <View> <TouchableOpacity onPress={() => { navigate( 'DetailPage' )}}> <Text style={styles.textStyle}>跳转详情页</Text> </TouchableOpacity> </View> ); } } export default class DetailPage extends PureComponent { static navigationOptions = { title: '详情页' , }; render () { let url = 'http://www.baidu.com' ; return ( <View> <WebView style={{width: '100%' ,height: '100%' }} source ={{uri: url}}/> </View> ); } } 复制代码`

除了示例中使用到的navigationOptions属性，StackNavigator导航器支持的navigationOptions属性还包括：

* header：设置导航属性，如果设置为null则隐藏顶部导航栏。
* headerTitle：设置导航栏标题。
* headerBackImage：设置后退按钮的自定义图片。
* headerBackTitle：设置跳转页面左侧返回箭头后面的文字，默认是上一个页面的标题。
* headerTruncatedBackTitle：设置上个页面标题不符合返回箭头后面的文字时显示的文字。
* headerRight：设置导航栏右侧展示的React组件。
* headerLeft：设置标题栏左侧展示的React组件。
* headerStyle：设置导航条的样式，如背景色、宽高等。
* headerTitleStyle：设置导航栏的文字样式。
* headerBackTitleStyle：设置导航栏上【返回】文字的样式。
* headerLeftContainerStyle：自定义导航栏左侧组件容器的样式，例如增加padding值。
* headerRightContainerStyle：自定义导航栏右侧组件容器的样式，例如增加 padding值。
* headerTitleContainerStyle：自定义 导航栏标题组件容器的样式，例如增加 padding值。
* headerTintColor：设置导航栏的颜色。
* headerPressColorAndroid：设置导航栏被按下时的颜色纹理，Android需要版本大于5.0。
* headerTransparent：设置标题背景是否透明。
* gesturesEnabled：设置是否可以使用手势关闭当前页面，iOS默认开启，Android默认关闭。

除了可以实现路由管理和页面跳转操作外，还可以使用react-navigation实现顶部和底部的Tab切换，如图7-13所示。

![在这里插入图片描述](https://user-gold-cdn.xitu.io/2019/6/4/16b1fc670008aaa3?imageView2/0/w/1280/h/960/ignore-error/1) 如果要实现底部选项卡切换功能，可以直接使用react-navigation提供的createBottomTabNavigator接口，并且此导航器需要使用createAppContainer函数包裹后才能作为React组件被正常调用。例如：

` import React, {PureComponent} from 'react' ; import {StyleSheet, Image} from 'react-native' ; import {createAppContainer, createBottomTabNavigator} from 'react-navigation' import Home from './tab/HomePage' import Mine from './tab/MinePage' const BottomTabNavigator = createBottomTabNavigator( { Home: { screen: Home, navigationOptions: () => ({ tabBarLabel: '首页' , tabBarIcon:({focused})=>{ if (focused){ return ( <Image/> //选中的图片 ) } else { return ( <Image/> //默认图片 ) } } }), }, Mine: { screen: Mine, navigationOptions: () => ({ tabBarLabel: '我的' , tabBarIcon:({focused})=>{ … } }) } }, { //默认参数设置 initialRouteName: 'Home' , tabBarPosition: 'bottom' , showIcon: true , showLabel: true , pressOpacity: 0.8, tabBarOptions: { activeTintColor: 'green' , style: { backgroundColor: '#fff' , }, } } ); const AppContainer = createAppContainer(BottomTabNavigator); export default class TabBottomNavigatorPage extends PureComponent { render () { return ( <AppContainer/> ); } } 复制代码`

当然，除了支持创建底部选项卡之外，react-navigation还支持创建顶部选项卡，此时只需要使用react-navigation提供的createMaterialTopTabNavigator即可。如果要使用实现抽屉式菜单功能，还可以使用react-navigation提供的createDrawerNavigator。

附： [react-navigation官网]( https://link.juejin.im?target=https%3A%2F%2Freactnavigation.org%2F )