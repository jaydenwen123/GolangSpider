# Flutter 混合开发FlutterBoost iOS 接入流程 #

紧接着上次的FlutterBoost Android版本接入，这次主要讲iOS相关的接入

### 1.创建Flutter module ###

这个步骤前面的Android版本一样

` flutter create -t module flutter_module 复制代码`

### 2.iOS开始接入 ###

#### 2.1 Pod集成 ####

现在一般的iOS应用都是用cocopod集成的，一般都有对应的Podfile文件，在对应的Podfile文件末尾处加入以下代码

` flutter_application_path = '../flutter_module' eval (File.read(File.join(flutter_application_path, '.ios' , 'Flutter' , 'podhelper.rb' )), binding) 复制代码`

最好也和Android一样，分开两个工程，iOS工程和flutter功能是平级的，这样互不影响

之后再iOS工程目录下执行 pod install 命令，会在pod下面的Development Pods文件下面看到Flutter 和FlutterPluginRegistrant 两个文件。

如果出现啥错误，记得在工程的BuildSettings 下面检查Enable BitCode是否为NO。

#### 2.2添加编译脚本 ####

` " $FLUTTER_ROOT /packages/flutter_tools/bin/xcode_backend.sh" build " $FLUTTER_ROOT /packages/flutter_tools/bin/xcode_backend.sh" embed 复制代码`

在BuildPhases 栏下，点击左上角的加号(+) 选择New Run Script Phase填入以上脚本

之后执行Build 编译，项目应该能运行起来，如果出现执行上面的步骤。

### 3.混编代码集成 ###

修改AppDelegate.h/m文件

` #import <UIKit/UIKit.h> #import <Flutter/Flutter.h> @interface AppDelegate : FlutterAppDelegate <UIApplicationDelegate> @property (strong, nonatomic) UIWindow *window; + (AppDelegate *)sharedAppDelegate; @end 复制代码`

h头文件需要集成FlutterAppDelegate

` #import <FlutterPluginRegistrant/GeneratedPluginRegistrant.h> #import "AppDelegate.h" #import "AppDelegate+Init.h" @interface AppDelegate () @end @implementation AppDelegate - (BOOL)application:(UIApplication *)application didFinishLaunchingWithOptions:(NSDictionary *)launchOptions { self.window = [[UIWindow alloc] initWithFrame:[[UIScreen mainScreen] bounds]]; self.window.backgroundColor = [UIColor whiteColor]; [self initConfigWithOptions:launchOptions]; [self.window makeKeyAndVisible]; [GeneratedPluginRegistrant registerWithRegistry:self]; return YES; } 复制代码`

在AppDelegate.m文件的didFinishLaunchingWithOptions方法中加入插件集成的方法

` [GeneratedPluginRegistrant registerWithRegistry:self]; 复制代码`

增加ViewController的业务跳转

` #import <Flutter/Flutter.h> #import "ViewController.h" @implementation ViewController - (void)viewDidLoad { [super viewDidLoad]; UIButton *button = [UIButton buttonWithType:UIButtonTypeCustom]; [button addTarget:self action:@selector(handleButtonAction) for ControlEvents:UIControlEventTouchUpInside]; [button set Title:@ "点我" for State:UIControlStateNormal]; [button set BackgroundColor:[UIColor redColor]]; button.frame = CGRectMake(80.0, 210.0, 160.0, 40.0); [self.view addSubview:button]; } - (void)handleButtonAction { FlutterViewController* flutterViewController = [[FlutterViewController alloc] init]; [self presentViewController:flutterViewController animated: false completion:nil]; } @end 复制代码`

这样即可点击跳转到Flutter默认生成的main界面

### 4.FlutterBoost接入 ###

#### 4.1 Flutter 工程接入FlutterBoost ####

在对应的pubspec.yaml文件中加入依赖，pubspec.yaml就是一个配置文件

` flutter_boost: git: url: 'https://github.com/alibaba/flutter_boost.git' ref: '0.0.411' 复制代码`

之后调用Package get，右上角即可查看，之后还是在命令行工具下在flutte_module 根目下，执行flutter build ios 以及在iOS的根目录下执行pod install 使iOS和flutter都添加FlutterBoost插件。

#### 4.2Flutter中main.dart文件中配置 ####

可以参考前面的Android版本

#### 4.3 iOS工程的修改 ####

#### 4.3.1 添加libc++ ####

需要 将libc++ 加入 "Linked Frameworks and Libraries" 这个主要是项目的General 的Linked Frameworks and Libraries 栏下，点击加号（+）搜索libc++,找到libc++.tbd即可

#### 4.3.2 修改AppDelegate.h/m文件 ####

` #import <UIKit/UIKit.h> #import <flutter_boost/FlutterBoost.h> @interface AppDelegate : FLBFlutterAppDelegate <UIApplicationDelegate, XGPushDelegate> @property (strong, nonatomic) UIWindow *window; + (AppDelegate *)sharedAppDelegate; @end 复制代码`

需要继承FLBFlutterAppDelegate ，而对应的.m文件可保持不变或者去掉

` [GeneratedPluginRegistrant registerWithRegistry:self]; 复制代码`

都可以

#### 4.3.3 实现FLBPlatform协议 ####

应用程序实现FLBPlatform协议方法，可以使用官方demo中的DemoRouter

` @interface DemoRouter : NSObject<FLBPlatform> @property (nonatomic,strong) UINavigationController *navigationController; + (DemoRouter *)sharedRouter; @end @implementation DemoRouter - (void)openPage:(NSString *)name params:(NSDictionary *)params animated:(BOOL)animated completion:(void (^)(BOOL))completion { if ([params[@ "present" ] boolValue]){ FLBFlutterViewContainer *vc = FLBFlutterViewContainer.new; [vc set Name:name params:params]; [self.navigationController presentViewController:vc animated:animated completion:^{}]; } else { FLBFlutterViewContainer *vc = FLBFlutterViewContainer.new; [vc set Name:name params:params]; [self.navigationController pushViewController:vc animated:animated]; } } - (void)closePage:(NSString *)uid animated:(BOOL)animated params:(NSDictionary *)params completion:(void (^)(BOOL))completion { FLBFlutterViewContainer *vc = (id)self.navigationController.presentedViewController; if ([vc isKindOfClass:FLBFlutterViewContainer.class] && [vc.uniqueIDString isEqual: uid]){ [vc dismissViewControllerAnimated:animated completion:^{}]; } else { [self.navigationController popViewControllerAnimated:animated]; } } @end 复制代码`

也可以自己根据此修改。其中的openPage 方法会接收来至flutter-->native以及native-->flutter的页面跳转，可以根据用户自由的书写

#### 4.3.5 初始化FlutterBoost ####

` [FlutterBoostPlugin.sharedInstance startFlutterWithPlatform：router onStart：^（FlutterViewController * fvc）{ }]; 复制代码`

官方demo是在AppDelegate中初始化的，可以修改FLBPlatform协议实现类完成对应的操作对应的初始化做

### 5.页面跳转 ###

Native-->Flutter

` FLBFlutterViewContainer *vc = FLBFlutterViewContainer.new; [vc set Name:name params:params]; [self.navigationController presentViewController:vc animated:animated completion:^{}]; 复制代码`

Flutter-->Native

` FlutterBoost.singleton.openPage( "pagename" , {}, true ); 复制代码`

最终都会跳转到FLBPlatform 协议实现类的openPage 方法中，很多操作都是在FLBPlatform协议实现类中，包括页面跳转，关闭，以及对应的一些Flutter 和Native通信相关的。