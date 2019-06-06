# 复习webpack4之如何编写loader #

> 
> 
> 
> 之前学习过webpack3的知识，但是webpack4升级后还是有很多变动的，所以这次重新整理一下webpack4的知识点，方便以后复习。
> 
> 

这次学习webpack4不仅仅要会配置，记住核心API，最好还要理解一下webpack更深层次的知识，比如打包原理等等，所以可能会省略一些比较基础的内容，但是希望我可以通过此次学习掌握webpack，更好地应对以后的工作。

# 1.编写入门级loader #

我在之前的文章中，已经把webpack基础的内容基本上都过了一遍，现在开始准备复习更高级的webpack知识了，首先从loader开始。

首先初始化一个项目

` npm init 复制代码`

然后安装依赖

` cnpm install -D webpack webpack-cli 复制代码`

创建一个src目录，里面创建一个index.js

![](https://user-gold-cdn.xitu.io/2019/6/5/16b27a4f4f84861d?imageView2/0/w/1280/h/960/ignore-error/1)

新建一个webpack.config.js，写入最基本的配置

![](https://user-gold-cdn.xitu.io/2019/6/4/16b22d0e6933eb1e?imageView2/0/w/1280/h/960/ignore-error/1)

如果此时，我们有个需求，中打包过程中，需要把world替换成mark，我们就可以借助loader来实现。首先在src同级目录新建一个loader文件夹，里面新建一个replaceLoader.js。

![](https://user-gold-cdn.xitu.io/2019/6/4/16b22e81796c61e1?imageView2/0/w/1280/h/960/ignore-error/1)

replaceLoader.js需要导出一个函数，注意：这个函数不能是箭头函数，因为webpack调用loader的时候会对this做一些变更，上面有一些方法，如果使用箭头函数，this指向就会有问题，没有办法调用this上的一些方法。

函数可以接受一个参数，参数是我们源代码的内容，所以可以对source进行操作后，return source，就可以改变源代码了。

![](https://user-gold-cdn.xitu.io/2019/6/4/16b22e961ca695b5?imageView2/0/w/1280/h/960/ignore-error/1)

然后使用我们自己写的loader，use就不填写loader名称了，需要写我们编写的loader的路径。

![](https://user-gold-cdn.xitu.io/2019/6/4/16b22ebbc1dd8912?imageView2/0/w/1280/h/960/ignore-error/1)

这样我们打包后发现，world已经被替换成mark了，这样我们就实现了一个最简单的loader。

# 2.给loader配置参数 #

loader中常常可以配置一些参数，那么我们如果想配置参数，要怎么做呢？

![](https://user-gold-cdn.xitu.io/2019/6/4/16b22f2058c7ddd4?imageView2/0/w/1280/h/960/ignore-error/1)

此时在replaceLoader中，可以通过this.query访问到参数。

![](https://user-gold-cdn.xitu.io/2019/6/4/16b22f48a9f41df5?imageView2/0/w/1280/h/960/ignore-error/1)

![](https://user-gold-cdn.xitu.io/2019/6/4/16b22f49c06b054d?imageView2/0/w/1280/h/960/ignore-error/1)

![](https://user-gold-cdn.xitu.io/2019/6/4/16b22f5699a245c6?imageView2/0/w/1280/h/960/ignore-error/1)

这样打包后，结果就会替换成我们的参数，但是官方推荐我们使用loader-utils来传参。

` cnpm install --save-dev loader-utils 复制代码`

![](https://user-gold-cdn.xitu.io/2019/6/4/16b2300c8007ccdd?imageView2/0/w/1280/h/960/ignore-error/1)

这样打包的结果也是我们传入的参数。

# 3.this.callback #

有时候我们不止要return一个resource，还可能要返回多个结果，就需要用到callback。

` this.callback( err: Error | null, content: string | Buffer, source Map?: SourceMap, meta?: any ); 复制代码`

第一个参数是错误，第二个是结果，第三个是sourcemap，第四个可以是任何内容（比如元数据）

![](https://user-gold-cdn.xitu.io/2019/6/4/16b2305453c1e5cc?imageView2/0/w/1280/h/960/ignore-error/1)

# 4. this.async #

在loader中，如果我们直接调用setTimeout，就会报错，那么如果我们想进行异步操作要怎么做呢？

![](https://user-gold-cdn.xitu.io/2019/6/5/16b27bcd8aceceba?imageView2/0/w/1280/h/960/ignore-error/1)

当要使用异步的时候，需要先把callback变为this.callback，然后再返回结果（和this.callback一样）。

这样再打包就不会有任何问题。

额外知识点：我们现在配置loader的时候，需要使用path.resolve，有没有什么方法可以像其他loader一样引用呢？

![](https://user-gold-cdn.xitu.io/2019/6/5/16b27c9294e13dd0?imageView2/0/w/1280/h/960/ignore-error/1)

这样只写loader名称，webpack就会先到node_modules里面找，找不到就去当前目录下的loaders中去找。

![](https://user-gold-cdn.xitu.io/2019/6/5/16b27ca98d8bee7b?imageView2/0/w/1280/h/960/ignore-error/1)