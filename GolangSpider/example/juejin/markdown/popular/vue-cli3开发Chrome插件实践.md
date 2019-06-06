# vue-cli3开发Chrome插件实践 #

之前找了不少如何开发谷歌插件的文章，结果发现都是些很基础的内容，并没有写到如何快速编译打包插件。我就在想为什么不能通过webpack来打包插件呢？如果通过webpack编译的话，就能使开发过程变得更舒服，使文件结构趋向模块化，并且打包的时候直接编译压缩代码。后来发现了 ` vue-cli-plugin-chrome-ext` 插件，通过这个插件能很方便地用 ` vue-cli3` 来开发谷歌插件，并能直接引用各种UI框架跟npm插件。

> 
> 
> 
> tip：如果你没接触过谷歌插件开发的话建议先看看基础文档：
> 
> 
> 
> * [Chrome 插件开发中文文档](
> https://link.juejin.im?target=https%3A%2F%2Fcrxdoc-zh.appspot.com%2Fextensions%2F
> )
> * [Chrome插件开发全攻略](
> https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fsxei%2Fchrome-plugin-demo.git
> )
> 
> 
> 

## 搭建环境 ##

* 创建一个 ` vue-cli3` 项目： ` vue create vue-extension` ，对话流程选择默认就行。
* 进入项目 ` cd vue-extension`
* 安装 ` vue-cli-plugin-chrome-ext` 插件： ` vue add chrome-ext` ,根据安装对话选项设置好。
* 删除 ` vue-cli3` 无用文件跟文件夹： ` src/main.js` 、 ` src/components`

## 运行项目 ##

* 

` npm run build-watch` 运行开发环境，对修改文件进行实时编译并自动在根目录下生成 ` dist` 文件夹，然后在浏览器上加载 ` dist` 文件夹完成插件安装。(注意： 修改 ` background` 文件跟 ` manifest.json` 文件并不能实时刷新代码，需要重新加载插件才行。 后面已经有实时刷新的解决方法)

![](https://user-gold-cdn.xitu.io/2019/5/28/16afc60a0ddda5ca?imageView2/0/w/1280/h/960/ignore-error/1)

* 

` npm run build` 运行生产环境编译打包，将所有文件进行整合打包。

## 引入element UI ##

目前的插件加载到浏览器后弹出页面是这种界面：

![](https://user-gold-cdn.xitu.io/2019/5/28/16afc6125e788226?imageView2/0/w/1280/h/960/ignore-error/1) 平时我们肯定要引入好看的UI框架的，在这里我们可以引入 ` element-ui` ，首先安装：

` npm install element-ui 复制代码`

考虑到插件打包后的文件大小，最后通过按需加载的方式来引入组件，按照 ` element-ui` 官方的按需加载方法，要先安装 ` babel-plugin-component` 插件:

` npm install babel-plugin-component -D 复制代码`

然后，将 ` babel.config.js` 修改为：

` module.exports = { presets : [ '@vue/app' ], "plugins" : [ [ "component" , { "libraryName" : "element-ui" , "styleLibraryName" : "theme-chalk" } ] ] } 复制代码`

接下来修改 ` popup` 相关文件引入所需组件， ` src/popup/index.js` 内容:

` import Vue from "vue" ; import AppComponent from "./App/App.vue" ; Vue.component( "app-component" , AppComponent); import { Card } from 'element-ui' ; Vue.use(Card); new Vue({ el : "#app" , render : createElement => { return createElement(AppComponent); } }); 复制代码`

` src/popup/App/App.vue` 内容：

` < template > < el-card class = "box-card" > < div slot = "header" class = "clearfix" > < span > 卡片名称 </ span > < el-button style = "float: right; padding: 3px 0" type = "text" > 操作按钮 </ el-button > </ div > < div v-for = "o in 4" :key = "o" class = "text item" > {{'列表内容 ' + o }} </ div > </ el-card > </ template > < script > export default { name : 'app' , } </ script > < style >.box-card { width : 300px ; } </ style > 复制代码`

渲染效果：

![](https://user-gold-cdn.xitu.io/2019/5/28/16afc618df22a4a8?imageView2/0/w/1280/h/960/ignore-error/1)

当然，不仅仅是插件内部的页面，还可以将 ` element-ui` 组件插入到 ` content` 页面。

## ` content.js` 使用 ` element-ui` 组件 ##

` content.js` 主要作用于浏览网页，对打开的网页进行插入、修改 ` DOM` ，对其进行操作交互。别觉得 ` element-ui` 只能配合 ` vue` 使用，其实就是一个前端框架，只要我们引入了就能使用， ` webpack` 会自动帮我们抽离出来编译打包。

> 
> 
> 
> 根据评论的朋友提示，可以通过Chrome插件的 ` chrome.extension.getURL` API来引入字体文件，解决 `
> element-ui` 无法引入相对路径的字体文件问题。
> 
> 

首先我们创建 ` src/content/index` 文件，在里面我们通过 ` chrome.extension.getURL` API来引入插件的字体文件，内容：

` import { Message, MessageBox } from 'element-ui' ; // 通过Chrome插件的API加载字体文件 ( function insertElementIcons ( ) { let elementIcons = document.createElement( 'style' ) elementIcons.type = 'text/css' ; elementIcons.textContent = ` @font-face { font-family: "element-icons"; src: url(' ${ window.chrome.extension.getURL( "fonts/element-icons.woff" )} ') format('woff'), url(' ${ window.chrome.extension.getURL( "fonts/element-icons.ttf " )} ') format('truetype'); /* chrome, firefox, opera, Safari, Android, iOS 4.2+*/ } ` document.head.appendChild(elementIcons); })(); MessageBox.alert( '这是一段内容' , '标题名称' , { confirmButtonText : '确定' , callback : action => { Message({ type : 'info' , message : `action: ${ action } ` }); } }) 复制代码`

` vue.config.js` 增加 ` content.js` 文件的打包配置，因为 ` content.js` （ ` background.js` 同样可以只生成js文件）只有js文件，不用像app模式那样打包生成相应的 ` html` 文件。这里我们还要对字体打包配置处理下，因为默认打包后的文件名是带有hash值的，在这里我们要去除掉，完整内容如下：

` const CopyWebpackPlugin = require ( "copy-webpack-plugin" ); const path = require ( "path" ); // Generate pages object const pagesObj = {}; const chromeName = [ "popup" , "options" ]; chromeName.forEach( name => { pagesObj[name] = { entry : `src/ ${name} /index.js` , template : "public/index.html" , filename : ` ${name}.html` }; }); // 生成manifest文件 const manifest = process.env.NODE_ENV === "production" ? { from : path.resolve( "src/manifest.production.json" ), to : ` ${path.resolve( "dist" )} /manifest.json` } : { from : path.resolve( "src/manifest.development.json" ), to : ` ${path.resolve( "dist" )} /manifest.json` }; const plugins = [ CopyWebpackPlugin([ manifest ]) ] module.exports = { pages : pagesObj, // // 生产环境是否生成 sourceMap 文件 productionSourceMap: false , configureWebpack : { entry : { 'content' : './src/content/index.js' }, output : { filename : 'js/[name].js' }, plugins : [CopyWebpackPlugin(plugins)] }, css : { extract : { filename : 'css/[name].css' // chunkFilename: 'css/[name].css' } }, chainWebpack : config => { // 处理字体文件名，去除hash值 const fontsRule = config.module.rule( 'fonts' ) // 清除已有的所有 loader。 // 如果你不这样做，接下来的 loader 会附加在该规则现有的 loader 之后。 fontsRule.uses.clear() fontsRule.test( /\.(woff2?|eot|ttf|otf)(\?.*)?$/i ) .use( 'url' ) .loader( 'url-loader' ) .options({ limit : 1000 , name : 'fonts/[name].[ext]' }) } }; 复制代码`

最后在 ` manifest.development.json` 加载 ` content.js` 文件，以及将字体文件添加到网页可加载资源字段 ` web_accessible_resources` 里去：

` { "manifest_version" : 2 , "name" : "vue-extension" , "description" : "a chrome extension with vue-cli3" , "version" : "0.0.1" , "options_page" : "options.html" , "browser_action" : { "default_popup" : "popup.html" }, "content_security_policy" : "script-src 'self' 'unsafe-eval'; object-src 'self'" , "content_scripts" : [{ "matches" : [ "*://*.baidu.com/*" ], "js" : [ "js/content.js" ], "run_at" : "document_end" }], "web_accessible_resources" : [ "fonts/*" ] } 复制代码`

然后浏览器重新加载插件后打开 ` https://www.baidu.com/` 网址后可看到：

![](https://user-gold-cdn.xitu.io/2019/5/28/16afc61ec6a202b3?imageView2/0/w/1280/h/960/ignore-error/1)

## 实时刷新插件 ##

之前写的时候发现单纯地通过 ` vue-cli3` 更新代码并不能使插件的 ` background.js` 、 ` content.js` 也跟着更新，因为代码已经加载到浏览器了，浏览器并不会监听这些文件的变化。也是通过评论的朋友推荐，通过 [crx-hotreload]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fxpl%2Fcrx-hotreload ) 可以完美实现实时刷新功能，而不用重新手动加载。代码还简单易用，在这里我们直接将代码复制到 ` src/utils/hot-reload.js` 文件：

` // 代码来源：https://github.com/xpl/crx-hotreload/edit/master/hot-reload.js const filesInDirectory = dir => new Promise ( resolve => dir.createReader().readEntries( entries => Promise.all(entries.filter( e => e.name[ 0 ] !== '.' ).map( e => e.isDirectory ? filesInDirectory(e) : new Promise ( resolve => e.file(resolve)) )) .then( files => [].concat(...files)) .then(resolve) ) ) const timestampForFilesInDirectory = dir => filesInDirectory(dir).then( files => files.map( f => f.name + f.lastModifiedDate).join()) const reload = () => { window.chrome.tabs.query({ active : true , currentWindow : true }, tabs => { // NB: see https://github.com/xpl/crx-hotreload/issues/5 if (tabs[ 0 ]) { window.chrome.tabs.reload(tabs[ 0 ].id) } window.chrome.runtime.reload() }) } const watchChanges = ( dir, lastTimestamp ) => { timestampForFilesInDirectory(dir).then( timestamp => { if (!lastTimestamp || (lastTimestamp === timestamp)) { setTimeout( () => watchChanges(dir, timestamp), 1000 ) // retry after 1s } else { reload() } }) } window.chrome.management.getSelf( self => { if (self.installType === 'development' ) { window.chrome.runtime.getPackageDirectoryEntry( dir => watchChanges(dir)) } }) 复制代码`

然后在 ` vue.config.js` 对热刷新代码进行处理，如果是开发环境的话就将其复制到 ` assets` 文件夹里面：

` // vue.config.js... // 在这段下面添加 const plugins = [ CopyWebpackPlugin([ manifest ]) ] // 开发环境将热加载文件复制到dist文件夹 if (process.env.NODE_ENV !== 'production' ) { plugins.push( CopyWebpackPlugin([{ from : path.resolve( "src/utils/hot-reload.js" ), to : path.resolve( "dist" ) }]) ) } ... 复制代码`

最后只要在开发环境 ` manifest.development.json` 里配置一下，将 ` hot-reload.js` 加载到 ` background` 运行环境中即可：

` "background": { "scripts": ["hot-reload.js"] } 复制代码`

## 添加打包文件大小预览配置 ##

既然用了 ` vue-cli3` 了，怎能不继续折腾呢，我们平时用 ` webpack` 开发肯定离不开打包组件预览功能，才能分析哪些组件占用文件大，该有的功能一个都不能少😎。这么实用的功能，实现起来也无非就是添加几行代码的事：

` // vue.config.js module.export = { /* ... */ chainWebpack: config => { // 查看打包组件大小情况 if (process.env.npm_config_report) { // 在运行命令中添加 --report参数运行， 如：npm run build --report config .plugin( 'webpack-bundle-analyzer' ) .use( require ( 'webpack-bundle-analyzer' ).BundleAnalyzerPlugin) } } } 复制代码`

就辣么简单，然后运行 ` npm run build --report` 看看效果：

![](https://user-gold-cdn.xitu.io/2019/5/28/16afc622c80b67a2?imageView2/0/w/1280/h/960/ignore-error/1)

## 自动打包zip ##

做过Chrome插件的都知道，提交到谷歌插件市场的话需要打包为zip文件才行。如果每次我们都需要将编译文件打包成zip的话就太麻烦了，这种每次都要经历的重复流程当然是交给程序来处理啦。 想打包zip的话首先要安装 ` zip-webpack-plugin` 插件到开发环境：

` npm install --save-dev zip-webpack-plugin 复制代码`

然后添加打包代码到 ` vue.config.js` :

` // vue.config.js... // 开发环境将热加载文件复制到静态文件夹(在这段下面添加) if (process.env.NODE_ENV !== 'production' ) { /*...*/ } // 生产环境下打包dist为zip if (process.env.NODE_ENV === 'production' ) { plugins.push( new ZipPlugin({ path : path.resolve( "dist" ), filename : 'dist.zip' , }) ) } ... 复制代码`

搞定收工！

## 结语 ##

事实证明， ` vue-cli3` 很强大， ` vue` 相关的插件并不是不能应用于开发浏览器插件， ` element-ui` 也不仅限于 ` vue` 的运用。只有你想不到，没有做不到的事😁。

> 
> 
> 
> tip：如果你懒得从头开始搭建模板的话也可以从GitHub拉取 [vue-extension-template](
> https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2FMrli2016%2Fvue-extension-template.git
> ) 。
> 
>