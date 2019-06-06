# 6个必须掌握的基础配置 - [webpack第一篇] #

## webpack实战系列全目录 ##

* 
* [ webpack 12个常见的实际场景]( https://juejin.im/post/5cda6f67e51d453ccd246501 )
* webpack15个常见的优化策略【敬请期待】
* webpack从0打造兼容ie8的脚手架【敬请期待】
* webpack面试全总结【敬请期待】

## 前言 ##

最近这段时间 ，一直在研究webpack相关的一些知识点，同时，公司正在做兼容ie8的官网，所以借此，把webpack相关知识点进行总结，同时最终目的是使用webpack4.x从0打造一个兼容ie8的脚手架，这样以后如果有这样的兼容浏览器的需求，大家就可以直接像使用vue-cli等脚手架一样，直接安装就可以生成模版文件（虽然可能都2019了，但还有兼容ie8的需求，内心一万个....）

本节，我们说一下webpack最基础的6个配置项：

> 
> * entry入口配置
> * output输出配置
> * module
> * resolve
> * plugin
> * devServer
> * 三种易混淆知识点
> 

## 一. entry入口配置 ##

首先，要大家清楚一点：webpack是采用模块化的思想，所有的文件或者配置都是一个个的模块，同时所有模块联系在一起，可以理解为就是一个简单的树状结构，那么最顶层的入口是什么呢？答案就是entry，

所以，webpack在执行构建的时候，第一步就是找到入口，从入口开始，寻找，遍历，递归解析出所有入口依赖的模块。

实例代码如下：

` module.exports = { entry: { app: './src/main.js' }, //还有output,module等其他配置项 } 复制代码`

我们来说明一下entry的类型：

![](https://user-gold-cdn.xitu.io/2019/5/14/16ab67d8c177b1b9?imageView2/0/w/1280/h/960/ignore-error/1) 当然，除了上面三种静态类型，我们还可以动态配置entry: 即采用箭头函数动态返回。 ![](https://user-gold-cdn.xitu.io/2019/5/14/16ab68054178f422?imageView2/0/w/1280/h/960/ignore-error/1)

此处，关于entry, 我们只需要记住，它有多种配置类型，而且可以动态配置，可以为入口设置别名，具体用法在实际开发中再查即可。

## 二. output输出配置 ##

前面说到entry可以是一个字符串，数组，也可以是对象，但是output只是一个对象，里面包含一系列输出配置项。

## 三. module ##

module 主要用于配置处理模块的规则，主要有三点：

#### 1. 配置loader ####

我们通过代码来说明：

` module: { rules: [ { test : /\.js$/, include: path.resolve(__dirname, 'src' ), //use可以是普通字符串数组，也可以是对象数组 use: [ 'babel-loader?cacheDirectory' ], use: [ { loader: 'babel-laoder' , options: { cacheDirectory: true ,// }, enforce: 'post' } ] }, { test : /\.scss$/, use: [ 'style-loader' , 'css-loader' , 'sass-loader' ], exclude: path.resolve(__dirname, 'node_modules' ) } { //对非文本文件采用 file-loader 加载 test : /\ . (gif Ipng Ijpe?g Ieot Iwoff Ittf Isvg Ipdf) $/, use: [’ file-loader ’], }， //配置更多的其他loader ] } 复制代码`

属性说明：

* test/include/exclude：表示匹配该loader的文件或者文件范围
* use：表示使用什么loader，它可以是一个字符串，也可以是字符串数组，也可以是对象数组，那多个loader时，执行顺序是从右向左，当然，也可以使用enforce去强制让某个loader的执行顺序放到最前面或者最后面。
* cacheDirectory : 表示传给 babel-loader 的参数，用于缓存 babel 的编译结果，加快编译的速度。
* enforce: post表示将该loader的执行顺序放到最前面，pre则相反。
* 多个loader时的处理顺序：为从后到前，即先交给 sass-loader 处理，再将结果交给 css-loader,最后交给 style-loader

#### 2. 配置noParse ####

noParse可以用于让webpack忽略哪些没有采用模块化的文件，不对这些文件进行编译处理，这样做可以提高构建性能，因为例如一些库：如jquey本身是没有采用模块化标注的，让webpack去解析这些文件即耗时，也没什么意义。

` module: { rules: [], noParse: /jquery|lodash/, noParse: (content) => { return /jquery/.test(content); } } 复制代码`

说明：

* noParse的值可以是正则表达式，也可以是一个函数
* 被忽略的文件里不应该包含 import、 require、 define 等模块化 语句，不 然会导致在构建出的代码中包含无法在浏览器环境下执行的模块化语句 。

#### 3. 配置parser ####

因为 Webpack 是以模块化的 JavaScript 文件为入口的，所以内置了对模块化 JavaScript 的解析功能，支持 AMO, CornmonJS、 SystemJS、 ES6。 parser 属性可以更细粒度地配置 哪些模块语法被解析、哪些不被解析。同 noParse 配置项的区别在于， parser 可以精确到 语法层 面，而 noParse 只能控制哪些文件不被解析。

` module: { rules: [ test : /\.js$/， use: [ ’ babel-loader ’], parse: [ amd: false ， //禁用AMD commonjs : false , //禁用 CommonJS system : false , //禁用 SystemJS harmony: false ， //禁用 ES6 import/ export requireinclude: false , // 禁用require.include requireEnsure: false , // 禁用require.ensure requireContext: false , // 禁用require.context browserify: false , //禁用 browserify requireJs : false , //禁用 requirejs: false , //禁用requirejs ] ] } 复制代码`

说明：

* parse是与noParse同级的属性，当然也可以嵌套到rules，表示针对与某个loader应用该属性的规则。
* 目前只要明白parse属性，是用于声明哪些模块语法被解析，哪些不被解析即可

## 四. resolve ##

resolve配置webpack如何去寻找模块所对应的文件, 我们平时通过import导入的模块，resolve可以告诉webpack如何去解析导入的模块，

#### 1. alias: 配置路径别名 ####

` resolve: { alias : { 'components' : './src/components/' , 'react$' : '/path/to/react.js' } } 复制代码`

配置以后，我们就可以通过：

* 

import Button from 'components/button'， 实际上就是 import Button from ' ./src/components/button ' ，

* 

react$只会命中以 react 结尾的导入语句，即只会将 import ’ react ’关键 字替换 成 import ’ / path/to/react .min.j s’ 。

#### 2. extensions：用于配置模块文件的后缀列表。 ####

我们可能在导入模块的时候，都遇到这种情况，例如 require (’. /data ’); 此时，我们发现导入的文件其实是没有后缀名的，为什么不用写后缀名呢？原因就是我们配置了resolve-extensions 后缀列表。默认是:

` resolve : { extension: [ '.js' , '.json' ] } 复制代码`

也就是说，当遇到 require (’. /data ’)这样的导入语句时， Webpack会先寻找./ data . js 文件，如果该文件不存在 ， 就去寻找 . /data .json 文件，如果还 是找不到，就报错 。

假如我们想让 Webpack优先使用目录下的 Typescript文件，则可以这样配置:

` resolve : { extension: [’.ts’,’.j5 ’,’.json’] } 复制代码`

#### 3. modules ####

resolve.modules 配置 Webpack 去哪些目录下寻找第三方模块，默认只会去 node modules 目录下寻 找 。有时我们的项目里会有一些模块被其他模块大量依赖和导入，由于 其他模块 的位置不定，针对不同的文件都要计算被导入的模块文件的相对路径 ，这个路径 有时会很长，

例如：就像import’../../../components/button’，这时可以利用modules 配置项优化 。假如那些被大量导入的模块都在./ src/components 目录下，则将 modules配置成、

` resolve: { modules : [’./ src/cornponents ’, ' node modules ’] } 复制代码`

此时，我们就可以简单地 通过import ’button ’ 导入 。

注意：请分清modules和alias的区别，modules是用来配置一些公共模块，这些公共模块和nodemodules类似，配置以后，我们就可以直接引用模块，前面不需要再加路径，而alias作用是配置路径别名，目的是可以让路径简化。两者是不一样的。

除此之外，还有

* descriptionFiles：配置描述第三方模块的文件名称：默认是package.json
* enforceExtension：配置后缀名是否必须加上

## 五. plugin ##

plugins 其实包括webpack本身自带的插件，也有开源的其他插件，都可以使用，它的作用就是解决loader之外的其他任何相关构建的事情。

` const CommonsChunkPlugin =require ( 'webpack/lib/optimize/ CommonsChunkPlugin' ) ; modules: { plugins: [ new CommonsChunkPlugin (( name :’ coπunon ’ ， chunks: [’a’,’b’] }) , }), //也可以配置其他插件 ] } 复制代码`

至此，我们需要明白：

* plugins的值是一个数组，可以传入多个插件实例。
* plugins如何配置并不是难点，难点是我们需要清楚常用的一些插件，分别解决了什么样的问题，以及这些插件本身的配置项，当然，目前只需知道plugins的作用即可。

## 六. devServer ##

devServe 主要用于本地开发的时候，配置本地服务的各种特性,下面列举一些常见的配置项

* hot：true/false; //是否开启模块热替换
* inline: true/false; //是否开启实时刷新，即代码更改以后，浏览器自动刷新
* contentBase //用于配置本地服务的文件根目录
* header //设置请求头
* host //设置域名
* port //设置端口
* allowedHosts: []//只有请求的域名在该属性所配置的范围内，才可以访问。
* https: true/false;// 使用使用https服务，默认为false
* compress: true/false; //是否启用Gzip压缩，默认为false.
* open //是否开启新窗口
* devtool : ’ source-map ’// 配置webpack是否生成source Map，以方便调试。
* watch： true //默认为true，表示是否监听文件修改以后，自动编译。

## 七：三种易混淆知识点 ##

#### module，chunk， bundle的区别 ####

module，chunk 和 bundle 其实就是同一份逻辑代码在不同转换场景下的取了三个名字：我们直接写出来的是 module，webpack 处理时是 chunk，最后生成浏览器可以直接运行的 bundle。

#### filename， chunkFilename ####

filename 指列在 entry 中，打包后输出的文件的名称。 chunkFilename 指未列在 entry 中，却又需要被打包出来的文件的名称。

#### hash, chunkhash, contenthash ####

具体可以参考： [juejin.im/post/5cede8…]( https://juejin.im/post/5cede821f265da1bbd4b5630#heading-4 )

## 总结 ##

通过本节介绍的6个常见的基础配置项：entry，output，module，plugins, resolve, devServe , 我们掌握的目标就是要清楚这几个配置项的基本功能，以及我们可以在哪些场景下使用他们即可，不需要去记具体怎么配，这些在实际开发中如果需要用到再查文档即可。

接下， 第二节 [webpack 16个常见的实际场景] 我们将结合具体实际案例，去近一步熟悉和掌握我们第一节学到的这些配置，敬请期待吧。