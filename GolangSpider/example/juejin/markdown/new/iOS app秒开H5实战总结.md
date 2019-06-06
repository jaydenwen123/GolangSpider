# iOS app秒开H5实战总结 #

在 [《iOS app秒开H5优化探索》]( https://juejin.im/post/5c9c664ff265da611624764d ) 一文中简单介绍了优化的方案以及一些知识点，本文继续介绍使用WKURLSchemeHandler拦截加载离线包优化打开速度的一些细节以及注意事项，阅读本文前请先大概了解一下上篇文章的内容以及WKURLSchemeHandler的基本用法。

## 离线包下载优化 ##

在上一篇 [《iOS app秒开H5优化]( https://juejin.im/post/5c9c664ff265da611624764d ) [探索]( https://juejin.im/post/5c9c664ff265da611624764d ) 》中，离线包下载处理有很多不合理的地方，如资源分散下载，不仅增加后续更新逻辑的复杂度，而且会造成系统资源浪费。为此可以把所有资源文件（js/css/html等）整合成zip包，一次性下载至本地，使用SSZipArchive解压到指定位置，更新version即可。 此外，下载时机在app启动和前后台切换都做一次检查更新，效果更好。

` NSURLSession *session = [NSURLSession sharedSession]; NSURLSessionDownloadTask *downLoadTask = [session downloadTaskWithRequest:request completionHandler:^(NSURL * _Nullable location, NSURLResponse * _Nullable response, NSError * _Nullable error) { if (!location) { return ; } //下载成功，移除旧资源 [fileManager removeFileAtPath:dirPath fileExtesion:nil]; //脚本临时存放路径 NSString *downloadTmpPath = [NSString stringWithFormat:@ "%@pkgfile_%@.zip" , NSTemporaryDirectory(), version]; // 文件移动到指定目录中 NSError *saveError; [fileManager moveItemAtURL:location toURL:[NSURL fileURLWithPath:downloadTmpPath] error:&saveError]; //解压zip BOOL success = [SSZipArchive unzipFileAtPath:downloadTmpPath toDestination:dirPath]; if (!success) { LogError(@ "pkgfile: unzip file error" ); [fileManager removeItemAtPath:downloadTmpPath error:nil]; [fileManager removeFileAtPath:dirPath fileExtesion:nil]; return ; } //更新版本号 [[NSUserDefaults standardUserDefaults] set Value:version for Key:pkgfileVisionKey]; [[NSUserDefaults standardUserDefaults] synchronize]; //清除临时文件和目录 [fileManager removeItemAtPath:downloadTmpPath error:nil]; }]; [downLoadTask resume]; [session finishTasksAndInvalidate]; 复制代码`

## WKWebView复用池 ##

在调试过程中，发现首次加载页面时间比后续打开时间都慢很多，原因预计是 webView 首次初始化时候需要启动资源和服务较多，于是尝试预先初始化 webView 复用方案，速度会快很多。

WKWebView复用池原理：预选准备两个NSMutableSet<WKWebView *>，一个正被使用visiableWebViewSet、一个空闲待用reusableWebViewSet，在+ (void)load初始化一个WKWebView，并加入reusableWebViewSet中，当H5页面需要使用时，从reusableWebViewSet中取出并放入visiableWebViewSet中，使用完（dealloc）放回reusableWebViewSet中。若该WKWebView异常则抛弃重新创建WKWebView，以免发生一些莫名其妙的问题。

#### 1、初始化 ####

` + (void)load { __block id observer = [[NSNotificationCenter defaultCenter] addObserverForName:UIApplicationDidFinishLaunchingNotification object:nil queue:nil usingBlock:^(NSNotification * _Nonnull note) { dispatch_async(dispatch_get_main_queue(), ^{ WKWebView *webview = [[WKWebView alloc] init]; [self->_reusableWebViewSet addObject:webview]; }); [[NSNotificationCenter defaultCenter] removeObserver:observer]; }]; } 复制代码`

#### 2、获取复用池中的webview ####

` - (WKWebView *)getReusedWebViewForHolder:(id)holder { if (!holder) { #if DEBUG NSLog(@ "WKWebViewPool must have a holder" ); #endif return nil; } WKWebView *webView; dispatch_semaphore_wait(_lock, DISPATCH_TIME_FOREVER); if (_reusableWebViewSet.count > 0) { webView = [_reusableWebViewSet anyObject]; [_reusableWebViewSet removeObject:webView]; [_visiableWebViewSet addObject:webView]; } else { [_visiableWebViewSet removeAllObjects]; webView = [[WKWebView alloc] init]; [_visiableWebViewSet addObject:webView]; } webView.holderObject = holder; dispatch_semaphore_signal(_lock); return webView; } 复制代码`

其中holder使用runtime为WKWebView添加的属性，传入使用复用池的当前VC即可，以供后续回收判断复用池是否正在使用。

#### 3、用完回收 ####

` - (void)recycleReusedWebView:(WKWebView *)webView { if (!webView) { return ; } dispatch_semaphore_wait(_lock, DISPATCH_TIME_FOREVER); if ([_visiableWebViewSet containsObject:webView]) { //将webView重置为初始状态 [webView webViewEndReuse]; [_visiableWebViewSet removeObject:webView]; [_reusableWebViewSet addObject:webView]; } else { if (![_reusableWebViewSet containsObject:webView]) { #if DEBUG NSLog(@ "Don't use the webView" ); #endif } } dispatch_semaphore_signal(_lock); } 其中webViewEndReuse为WKWebView的扩展方法： - (void)webViewEndReuse { self.holderObject = nil; if ([self isKindOfClass:[WKWebView class]]) { WKWebView *webView = (WKWebView *)self.webView; webView.delegate = nil; webView.scrollView.delegate = nil; [webView stopLoading]; [webView set UIDelegate:nil]; [webView loadHTMLString:@ "" baseURL:nil]; } } 复制代码`

复用池原理很简单，此外对收到内存警告后清除web缓存等回收处理等等，此处不再赘述。

## WebViewController改造 ##

通常项目中处理H5页面都会放在统一的WebViewController中，所以要结合开关、要优化的业务来分开复用池和普通webView的使用，以免出问题。

#### 1、替换url scheme ####

` NSString *urlString = @ "https://www.test.com/abc?id=123456" ; if ([YH_Global sharedInstance].isGrassLocalOpen && SYSTEM_VERSION_GREATER_THAN_OR_EQUAL_TO(@ "11.0" )) { urlString = [urlString stringByReplacingOccurrencesOfString:@ "https" withString:@ "customScheme" ]; } WebViewController *vc = [[WebViewController alloc] initWithUrl:urlString]; [self.navigationController pushViewController:vc animated:YES]; 复制代码`

此处替换url scheme http(s)为自定义协议，使用拦截生效。 此处需要特别说明的是，前端H5请求的js、css等资源使用自适应的协议，如： ` src='//www.test.com/abc.js'` ，这样native端使用不同scheme请求，H5就会使用对应的scheme进行请求加载。另外一个重要的点是，前端的ajax请求，scheme使用http(s)，native不要拦截，完全交给H5与服务器交互，这样就不会发生发送post请求，body丢失的情况。

#### 2、初始化webView ####

` - (instancetype)initWithUrl:(NSString *)url { if (self = [super init]) { if ([self checkMatchingWithUrl:url]) {//符合条件，使用复用池 self.webView = [[WKWebViewPool sharedInstance] getReusedWebViewForHolder:self]; } self.url = url; } return self; } 复制代码`

此处在initWithUrl中，而不在viewDidLoad中获取webView，是因为在init中，页面打开速度会快很多。

#### 3、预先添加数据脚本，提升体验 ####

这一步根据笔者公司的app的业务特性所有：用户社区帖子列表（native） => 帖子详情（H5实现）=> 个人中心等（H5）。从列表点击进入H5详情时，预先将帖子的部分数据，如头像、首图缩略图、内容等传给前端(，前端拿到数据，预先加载这部分数据，同时对首图缩略图增加渐变出现的效果，这时打开H5，页面从模糊的缩略图渐变至高清大图，以达到原生打开页面的体验（文末的最终效果图）。注意，这里图片传给前端的是url，并不是图片数据，下文会继续说明如何使用图片数据。

native与H5交互的部分代码：

` Model *modelMake = model;//列表点击的item数据 NSString *key = [NSString stringWithFormat:@ "native_list_%@" , modelMake.articleId]; NSData *data = [NSJSONSerialization dataWithJSONObject:[modelMake dictionaryValue] options:NSJSONWritingPrettyPrinted error:nil]; NSString *value = [[NSString alloc] initWithData:data?data:[NSData data] encoding:NSUTF8StringEncoding]; NSString *javaScript = [NSString stringWithFormat:@ "!window.predatas && (window.predatas = []);predatas.push({key: \"%@\", value: %@ })" , key, value]; WKUserContentController *userContentController = wkWebView.configuration.userContentController; WKUserScript *userScript = [[WKUserScript alloc] initWithSource:javaScript injectionTime:WKUserScriptInjectionTimeAtDocumentStart for MainFrameOnly:NO]; [userContentController addUserScript:userScript]; 复制代码`

#### 4、侧滑返回 ####

由于业务使用H5开发，从列表到详情再到个人中心，这时侧滑会直接回到列表页，并不像原生导航那样一层层返回。解决这个问题，首先想到使用WKWebView的allowsBackForwardNavigationGestures属性，结合webView的goBack方法，的确可以层层侧滑返回，但是最后出现会先回到第一次打开的详情页面，然后才会回到列表的情况以及一些其他异常问题。尝试了一些方案后，最终采用自己添加手势实现侧滑返回功能。

手势创建：

` self.leftSwipGes = [[UIScreenEdgePanGestureRecognizer alloc] initWithTarget:self action:@selector(leftSwipGesAction:)]; self.leftSwipGes.edges = UIRectEdgeLeft; self.leftSwipGes.delegate = self; [self.webView addGestureRecognizer:self.leftSwipGes]; 复制代码`

实现：

` - (void)leftSwipGesAction:(UIScreenEdgePanGestureRecognizer *)ges { if (UIGestureRecognizerStateEnded == ges.state) { if (self.webView.backForwardList.backList.count > 0) { WKBackForwardListItem *item = webView.backForwardList.backList.lastObject; if (![self.webView.URL.absoluteString isEqualToString:self.url]) { [webView goToBackForwardListItem:item]; } else { [self nativeBack:nil]; [webView goToBackForwardListItem:item]; } } else { [self nativeBack:nil]; } } } 复制代码`

其中，nativeBack()为native的返回方法。原理：侧滑时，当前webview的url不是初始H5页面的url时，webView的backForwardList退后一级，当退到初始页面时，直接返回列表。此外，注意处理自定义手势跟其他手势冲突的问题；同时还要禁用系统的侧滑返回，以及禁用FDFullscreenPopGesture等第三方库的侧滑返回。

## 拦截加载离线包 ##

前提创建WKWebview时注册好自定义协议，具体结合自己项目实现，只要保证创建WKWebView时注册即可：

` WKWebViewConfiguration *configuration = [WKWebViewConfiguration new]; [configuration set URLSchemeHandler:[CustomURLSchemeHandler new] for URLScheme: @ "customScheme" ]; WKWebView *webView = [[WKWebView alloc] initWithFrame:self.view.bounds configuration:configuration]; 复制代码`

#### 拦截 ####

上文也分析了，打开一个H5页面会有一段时间白屏，是因为它做了很多事情：

` 初始化 webview -> 请求页面 -> 下载数据 -> 解析HTML -> 请求 js/css 资源 -> dom 渲染 -> 解析 JS 执行 -> JS 请求数据 -> 解析渲染 -> 下载渲染图片`

所以当打开以自定义协议customScheme为scheme的H5页面时，webview请求页面，native会依次收到html、js、css、图片类型的拦截响应：

` - (void)webView:(WKWebView *)webView startURLSchemeTask:(id <WKURLSchemeTask>)urlSchemeTask API_AVAILABLE(ios(11.0)){ NSDictionary *headers = urlSchemeTask.request.allHTTPHeaderFields; NSString *accept = headers[@ "Accept" ]; //当前的requestUrl的scheme都是customScheme NSString *requestUrl = urlSchemeTask.request.URL.absoluteString; NSString *fileName = [[requestUrl componentsSeparatedByString:@ "?" ].firstObject componentsSeparatedByString:@ "/" ].lastObject; //Intercept and load local resources. if ((accept.length >= @ "text".length && [accept rangeOfString:@ "text/html" ].location != NSNotFound)) {//html 拦截 [self loadLocalFile:fileName urlSchemeTask:urlSchemeTask]; } else if ([self isMatchingRegularExpressionPattern:@ "\\.(js|css)" text:requestUrl]) {//js、css [self loadLocalFile:fileName urlSchemeTask:urlSchemeTask]; } else if (accept.length >= @ "image".length && [accept rangeOfString:@ "image" ].location != NSNotFound) {//image NSString *replacedStr = [requestUrl stringByReplacingOccurrencesOfString:kUrlScheme withString:@ "https" ]; NSString *key = [[SDWebImageManager sharedManager] cacheKeyForURL:[NSURL URLWithString:replacedStr]]; [[SDWebImageManager sharedManager].imageCache queryCacheOperationForKey:key done :^(UIImage * _Nullable image, NSData * _Nullable data, SDImageCacheType cacheType) { if (image) { NSData *imgData = UIImageJPEGRepresentation(image, 1); NSString *mimeType = [self getMimeTypeWithFilePath:fileName] ?: @ "image/jpeg" ; [self resendRequestWithUrlSchemeTask:urlSchemeTask mimeType:mimeType requestData:imgData]; } else { [self loadLocalFile:fileName urlSchemeTask:urlSchemeTask]; } }]; } else {// return an empty json. NSData *data = [NSJSONSerialization dataWithJSONObject:@{ } options:NSJSONWritingPrettyPrinted error:nil]; [self resendRequestWithUrlSchemeTask:urlSchemeTask mimeType:@ "text/html" requestData:data]; } } //Load local resources, eg: html、js、css... - (void)loadLocalFile:(NSString *)fileName urlSchemeTask:(id <WKURLSchemeTask>)urlSchemeTask API_AVAILABLE(ios(11.0)){ if (fileName.length == 0 || !urlSchemeTask) { return ; } //If the resource do not exist, re-send request by replacing to http(s). NSString *filePath = [kGrassH5ResourcesFiles stringByAppendingPathComponent:fileName]; if (![[NSFileManager defaultManager] fileExistsAtPath:filePath]) { if ([replacedStr hasPrefix:kUrlScheme]) { replacedStr = [replacedStr stringByReplacingOccurrencesOfString:kUrlScheme withString:@ "https" ]; } NSMutableURLRequest *request = [NSMutableURLRequest requestWithURL:[NSURL URLWithString:replacedStr]]; NSURLSession *session = [NSURLSession sessionWithConfiguration:[NSURLSessionConfiguration defaultSessionConfiguration]]; NSURLSessionDataTask *dataTask = [session dataTaskWithRequest:request completionHandler:^(NSData * _Nullable data, NSURLResponse * _Nullable response, NSError * _Nullable error) { [urlSchemeTask didReceiveResponse:response]; [urlSchemeTask didReceiveData:data]; if (error) { [urlSchemeTask didFailWithError:error]; } else { [urlSchemeTask didFinish]; NSString *accept = urlSchemeTask.request.allHTTPHeaderFields[@ "Accept" ]; if (!(accept.length >= @ "image".length && [accept rangeOfString:@ "image" ].location != NSNotFound)) { //图片不下载 [data writeToFile:filePath atomically:YES]; } } }]; [dataTask resume]; [session finishTasksAndInvalidate]; } else { NSData *data = [NSData dataWithContentsOfFile:filePath options:NSDataReadingMappedIfSafe error:nil]; [self resendRequestWithUrlSchemeTask:urlSchemeTask mimeType:[self getMimeTypeWithFilePath:filePath] requestData:data]; } } - (void)resendRequestWithUrlSchemeTask:(id <WKURLSchemeTask>)urlSchemeTask mimeType:(NSString *)mimeType requestData:(NSData *)requestData API_AVAILABLE(ios(11.0)) { if (!urlSchemeTask || !urlSchemeTask.request || !urlSchemeTask.request.URL) { return ; } NSString *mimeType_local = mimeType ? mimeType : @ "text/html" ; NSData *data = requestData ? requestData : [NSData data]; NSURLResponse *response = [[NSURLResponse alloc] initWithURL:urlSchemeTask.request.URL MIMEType:mimeType_local expectedContentLength:data.length textEncodingName:nil]; [urlSchemeTask didReceiveResponse:response]; [urlSchemeTask didReceiveData:data]; [urlSchemeTask didFinish]; } } 复制代码`

我这里只简单贴了部分拦截资源请求后的处理代码：收到拦截请求后，先获取本地资源包对应的资源，转换成data回传给webView进行渲染处理；若本地没有，则customScheme替换成https的url重发请求通知webview，这就是基本流程。实际开发调试过程中还有很多细节需要处理，如本地资源没有时，根据服务器预先下发的匹配规则重发请求；又如加载替换使用不同的html，又如打开页面一直白屏等等问题，这里就不列出了。

**但还要特别说明两点：**

1、代码中替换图片的逻辑，先查找本地图片的目的是为了实现上文所说的 **WebViewController改造第三条：预先添加数据脚本，提升体验** ，获取列表中已展示缩略图的SDWebImage缓存传给webView进行预加载，以实现渐变出现的效果。然后本地就重发请求通知webview。 到这里，你应该明白，优化实现秒开，中心思想就是要减少资源的网络请求，把第一页要展示的原素尽量预先加载。

2、在测试过程中，在一些机型较差的机器上，频繁快速的打开H5页面，会出现崩溃。查阅WKURLSchemeTask的官方解释：

> 
> 
> 
> An exception will be thrown if you try to send a new response object after
> the task has already been completed.
> An exception will be thrown if your app has been told to stop loading this
> task via the registered WKURLSchemeHandler object.
> 
> 

经分析，发现在处理本地不存在的图片时，先判断本地是否存在而后又发起请求，时间跨度比较长，当前urlSchemeTask由于某些原因提前结束了（会收到stopURLSchemeTask回调）,这时重发的请求又访问了WKURLSchemeTask的实例方法（didReceiveResponse等）就导致了崩溃。
解决办法：新增NSMutableDictionary成员变量，以当前的urlSchemeTask做key，拦截开始时设置YES，收到停止通知时设置NO，每次通知webview前判断当前的urlSchemeTask是否结束，提前结束了就不做处理。 这么做，带来的影响就是当前的图片会不显示，退出再次进来还是会出现的，结合出现的异常场景以及发生崩溃，这点影响还是可以接受的。

` - (void)webView:(WKWebView *)webView startURLSchemeTask:(id <WKURLSchemeTask>)urlSchemeTask API_AVAILABLE(ios(11.0)){ dispatch_sync(self.serialQueue, ^{ [self.holderDicM set Object:@(YES) for Key:urlSchemeTask.description]; }); } - (void)webView:(WKWebView *)webView stopURLSchemeTask:(id <WKURLSchemeTask>)urlSchemeTask API_AVAILABLE(ios(11.0)){ dispatch_sync(self.serialQueue, ^{ [self.holderDicM set Object:@(NO) for Key:urlSchemeTask.description]; }); } 复制代码`

注意要添加串行队列对数据进行保护，防止多线程同时访问修改数据，造成数据异常。

## 总结 ##

到这里，优化基本上完成，打开H5页面确实快了很多。我们的方案大致就是这样，这个肯定不是最优的方案，多多少少会有些问题，我相信读者会有更好的优化方案，或者遇到上述出现的问题有更合理的解决方法，欢迎大家一起讨论。

最后展示一下，我们优化后的打开H5页面的效果（iPhone 7）：

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2b6730f1b5f67?imageslim)