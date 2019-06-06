# JS与OC交互 #

#### UIWebView拦截URL ####

* 

原理：

* js通过加载url方式被webView拦截，这时候看如果是自己定义的scheme请求就不让webView继续加载请求，否则就继续加载请求。
* webView看加载的请求的host是哪种host进行分别处理。
* 处理oc代码。
* 之后调用stringByEvaluatingJavaScriptFromString调用js代码。

* 

注意：

* js调用oc属于异步方式
* oc调用js属于同步方式，且必须在主线程加载，如果js代码比较耗时那么可能会卡顿主线程

* 

[代码地址]( https://link.juejin.im?target=https%3A%2F%2Fpan.baidu.com%2Fs%2F1PGyq8rvah_dAVdWvycsLlA )

#### UIWebview利用jsc库 ####

* 

原理:

* html调用方法： function locationClick() { getLocation('A','B','C'); }
* oc中delegate回调：
` - (void)webViewDidFinishLoad:(UIWebView *)webView{ JSContext *context = [self valueForKeyPath:@ "documentView.webView.mainFrame.javaScriptContext" ]; } 复制代码` * 通过分析context[@"getLocation"]来判断方法名。
* 通过NSArray *arrArgs = [JSContext currentArguments];获取参数

* 

注意：

* js执行时候会进入context的回调，该回调block是在子线程中的。如果更细ui要在主线程。

* 

[代码地址]( https://link.juejin.im?target=https%3A%2F%2Fpan.baidu.com%2Fs%2F1PGyq8rvah_dAVdWvycsLlA )

#### WKWebView拦截URL ####

* 

原理：

* js通过加载url方式被webView拦截，这时候看如果是自己定义的scheme请求就不让webView继续加载请求，否则就继续加载请求。
* webView看加载的请求的host是哪种host进行分别处理。
* 处理oc代码。
* 之后调用evaluateJavaScript调用js代码。

* 

注意：

* js调用oc属于异步方式
* oc调用js属于同步方式，且必须在主线程加载，如果js代码比较耗时那么可能会卡顿主线程
* WKWebView有个处理js弹窗的代理方法，这个方法必须要实现，如果不实现js的弹窗将会无效。

* 

WKWebView和UIWebView的比较

* wk更节省内存
* wk加载速度刚快
* wk解决了内存泄露问题
* wk刚好适配了ios8+

* 

[代码地址]( https://link.juejin.im?target=https%3A%2F%2Fpan.baidu.com%2Fs%2F1FX41xiteQQo1BTXMAdIa3Q )

#### WKWebView messageHandle方式 ####

* 原理： * js通过调用方法window.webkit.messageHandlers.getLocation.postMessage({A:'a',B:'b'});其中getLocation为name，{A:'a',B:'b'}相当于参数
* oc通过 [self.webView.configuration.userContentController addScriptMessageHandler:self name:obj];相当于注册监听
* oc会进入回调userContentController:(WKUserContentController *)userContentController didReceiveScriptMessage:(WKScriptMessage *)message。
* 通过分析message.name来判断是js调用的那个方法，通过分析message.body来获取参数。

[代码地址]( https://link.juejin.im?target=https%3A%2F%2Fpan.baidu.com%2Fs%2F1FX41xiteQQo1BTXMAdIa3Q )

#### 可以利用三方框架实现交互 ####

* WebViewJavascriptBridge支持UIWebView，WKWebView
* 利用这个框架可以实现oc与js互调

#### 可以利用前端框架 ####

* RN可以实现
* Cordova可以实现

### webView实现全包裹 ###

* ios要想做到内容全包裹，必须借助js，不像android...
* js获取内容高度方法为：
` CGFloat height = [[webView stringByEvaluatingJavaScriptFromString:@ "document.body.scrollHeight" ] float Value]; 复制代码`