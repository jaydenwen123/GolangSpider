# Android 富文本处理 #

### 最近遇到了一些富文本处理的坑，特此分享下 ###

### TextView ###

` Html.fromHtml(data.getPro_job())` 这个方式最简单，缺陷是不能解析 ` ul` 、 ` li` 等类型标签。

### RichText ###

` RichText.from(data.getPro_job()).into(wvDes);`

[RichText]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fzzhoujay%2FRichText ) 是Android平台下的富文本解析器，支持Html和Markdown，这样就可以解析 ` ul` 等标签，但缺陷是字体样式加载会有问题，比如说加粗、颜色等。

### WebView ###

建议使用腾讯x5浏览器，使用方法可查看另一篇博客 [【Android Web】腾讯X5浏览器的集成与常见问题]( https://juejin.im/post/5c75e708e51d453ee8185d7b )

` webView.loadDataWithBaseURL("", data.getPro_intro(), "text/html", "utf-8", null);`

` webview` 在解析标签上没有问题，但是又引发了另一个问题， ` ScrollView` 嵌套下，底部会有一大片的空白。

![](https://user-gold-cdn.xitu.io/2019/6/3/16b1c9da5dd35381?imageView2/0/w/1280/h/960/ignore-error/1)

这个问题有2种处理方式

* 获取 ` html` 的内容高度，然后再设置给 ` webview`

` wvDes.setWebViewClient( new WebViewClient() { @Override public void onPageFinished (WebView webView, String s) { super.onPageFinished(webView, s); wvDes.loadUrl( "javascript:app.resize(document.body.getBoundingClientRect().height)" ); } }); wvDes.addJavascriptInterface( new WebViewJavaScriptFunction() { @Override public void onJsFunctionCalled (String tag) { } @JavascriptInterface public void resize ( int height) { wvDes.postDelayed( new Runnable() { @Override public void run () { LinearLayout.LayoutParams params = new LinearLayout.LayoutParams(LinearLayout.LayoutParams.MATCH_PARENT, height + 80 ); wvDes.setLayoutParams(params); } }, 100 ); } }, "app" ); 复制代码`

但是，经测试，在部分华为手机上， ` document.body.getBoundingClientRect().height` 获取的高度不正确，只能无奈放弃。

* addView

` mParentLayout.removeAllViews(); WebView webView = new WebView( this ); webView.setLayoutParams( new LinearLayout.LayoutParams(ViewGroup.LayoutParams.MATCH_PARENT, ViewGroup.LayoutParams.WRAP_CONTENT)); webView.loadDataWithBaseURL( "" , data.getPro_intro(), "text/html" , "utf-8" , null ); mParentLayout.addView(webView); 复制代码`

目前众多机型测试中没有发现问题，有问题的小伙伴，可以在下面留言。

### 你的认可，是我坚持更新博客的动力，如果觉得有用，就请点个赞，谢谢 ###