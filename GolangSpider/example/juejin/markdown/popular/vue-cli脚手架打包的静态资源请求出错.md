# vue-cli脚手架打包的静态资源请求出错 #

### vue-cli脚手架打包的静态资源请求出错 ###

#### 问题 ####

* vue-cli默认配置打包后部署到特定路径下静态资源路径错误问题。
* 静态资源打包使用相对路径后css文件引入大图片路径错误问题

使用vue-cli2脚手架生成的默认打包配置文件，npm run build打包，部署项目到特定路径下：//ip:port/test/index.html 此时访问//ip:port/test/index.html可以正常访问，但是引用的js和css等文件服务器响应为404，此时我们查看资源请求路径：

` http://ip:port/static/css/app.[ hash ].css http://ip:port/static/js/app.[ hash ].js 复制代码`

可以看出，上面的静态资源访问路径是不正确的，我们正确的请求路径应该是

` http://ip:port/ test /static/css/app.[ hash ].css http://ip:port/ test /static/js/app.[ hash ].js 复制代码`

#### 原因 ####

可以看出导致资源加载失败的原因是路径错误，我们可以移步看看index.html文件，

` <!DOCTYPE html> <html> <head> <title>project</title> <link href=/static/css/app.css rel=stylesheet> </head> <body> <div id=app></div> <script type =text/javascript src=/static/js/app.js> </script> ... </body> </html> 复制代码`

可以看出引入的css和js都是使用的绝对根目录路径，因此将项目部署到特定目录下，其引入的资源路径无法被正确解析。

#### 解决 ####

在webpack打包时，使用相对路径来处理静态资源，修改build中资源发布路径配置(build/config.js);

` build: { ... // Paths assetsRoot: path.resolve(__dirname, '../dist' ), assetsSubDirectory: 'static' , assetsPublicPath: './' , ... } 复制代码`

将assetsPublicPath: '/'，更改为assetsPublicPath: './'，再进行打包，并将资源部署到特定路径下，然后访问，此时index.html可以正常访问，同时js和css资源也可以正常加载访问了。

#### css中引入assets目录下的图片资源出错 ####

我们经常这样引用一个img图片

` background: url( 'static/img/bg.png' ); 复制代码`

但是打包后看到这个图片的引用地址是这样的。

` http://ip:port/ test /static/css/static/img/bg.png 复制代码`

可以看出css中图片的路径存在问题了，分析打包过程，css是在js中引入的或是写在vue文件中的，css文件首先被less，postcss等处理，处理后会被ExtractTextPlugin处理，ExtractTextPlugin将js中的css全部抽离至app.css文件中。

##### 解决方法一 #####

将options.extract设置为false

` options.extract: false , 复制代码`

关闭抽离css功能，再次打包并部署，此时你会发现没有css文件了，css文件全部在app.js文件中，通过js将css注入到index.html文件中，此时图片的访问路径是相当index.html文件的，所以可以正常访问

#### 解决方案二 ####

设置ExtractTextPlugin插件中的publicPath ExtractTextPlugin插件是为了将css从js文件中抽离出来，我们可以通过配置ExtractTextPlugin的静态资源路径参数来达到同样的效果，build目录下的utils.js文件，修改publicPath: '../../';

` // Extract CSS when that option is specified // ( which is the case during production build) if (!options.extract) { return ExtractTextPlugin.extract({ use: loaders, fallback: 'vue-style-loader' , publicPath: '../../' }) } else { return [ 'vue-style-loader' ].concat(loaders) } 复制代码`

再次打包部署，发现此时的图片访问路径为'../../static/img/bg.png';

> 
> 
> 
> publicPath 属性值为打包后的 app.css文件至index.html文件的相对路径
> 
>