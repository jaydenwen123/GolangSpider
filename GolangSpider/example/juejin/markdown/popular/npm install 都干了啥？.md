# npm install 都干了啥？ #

## 起因 ##

本周我在将egg项目发布到生产环境之后，发现生产环境无法执行 ` npm run start` ，提示如下图

![image](https://user-gold-cdn.xitu.io/2019/6/5/16b253218d27bd04?imageView2/0/w/1280/h/960/ignore-error/1)

## 摸索原因 ##

### script 命令是如何运行的？ ###

作为一个前端工程师，我每天都会和各种环境打交道，经常在命令行执行 ` cross-env NODE_ENV=dev` 等操作，突然有一天，我发现执行 ` cross-env NODE_ENV=dev` 居然给我报错了？

![image](https://user-gold-cdn.xitu.io/2019/6/5/16b2532193a78f01?imageView2/0/w/1280/h/960/ignore-error/1)

检查了下原来是包没有使用全局安装，现在全局安装下 ` npm i cross-env -g`

![image](https://user-gold-cdn.xitu.io/2019/6/5/16b253218f6184c3?imageView2/0/w/1280/h/960/ignore-error/1)

嗯，果然不报错了，为啥啊？其实我们看上面的安装反馈就可以知道，全局安装会创建一个软连接

![image](https://user-gold-cdn.xitu.io/2019/6/5/16b253218f2ca5af?imageView2/0/w/1280/h/960/ignore-error/1)

而使用项目依赖安装则不会有这步操作，则导致我们不能直接使用简写去执行命令。

![image](https://user-gold-cdn.xitu.io/2019/6/5/16b253218f36d861?imageView2/0/w/1280/h/960/ignore-error/1)

那么问题来了，难道我所有需要在命令行执行的包都需要安装到全局，才能使用？

* 第一种方法，可以执行文件 `./node_modules/.bin/cross-env NODE=dev`
* 作为一个懒人，当然不会用方法一啦，这个时候 ` npm script` 就派上用场了

在 ` package.json` 的 ` script` 中新增命令

` "scripts": { "dev": "cross-env NODE=dev" } 复制代码`

执行 ` npm run dev` ，看到这里，我想大多数人肯定会说，我靠，你说的方法我几百年前就知道了啊，还需要你来BB？

但是我想还是有一部分人和我一样不懂为啥使用 ` npm script` 执行上面的命令就不会报错呢，难道也会创建软连接？其实是因为我们执行 ` npm run` 的时候会将 ` node_modules/.bin` 下面所有文件都添加到系统的环境变量中，这样在执行期间就可以使用缩写，等执行结束会自动删除这些环境变量

理清这些之后再来看我一开始说的问题，提示找不到模块，看了下 .bin 下面的代码

![image](https://user-gold-cdn.xitu.io/2019/6/5/16b2532196afbaa9?imageView2/0/w/1280/h/960/ignore-error/1)

看样子是引入上级目录的 ` index.js` ，但是我们并没有在上级目录找到该文件，导致报错，难道是别人的包写错了？当然这是不可能的，经过我不断的折腾，最后才发现！！！

其实我们在 ` npm install` 的时候首先会下载对应资源包的压缩包放在用户目录下的 `.npm` 文件夹下，然后解压到项目的 ` node_modules` 中，并且提取依赖包中指定的bin文件，在 ` linux` 下会创建一条软连接，所以在 ` linux` 下我们真正执行的是 `.bin` 文件夹下文件指向的文件，看下图。

![image](https://user-gold-cdn.xitu.io/2019/6/5/16b25321fddc855a?imageView2/0/w/1280/h/960/ignore-error/1)

而我遇到的问题就是执行 ` npm install` 和 ` npm start` 是两台机器，生产机上的文件实际上是从发布机 copy 过来的，导致软连接没了，所以才会报错。

### 吸取经验 ###

从这么一个小小的问题就可以暴露出来，虽然我们每天都会执行 ` npm script` 但是却对他的执行过程很模糊，导致遇到问题却找不到原因，最终在看了一些资料之后我总结了如下几个点。

## 总结 ##

* ` npm install` 会先查找本地已经下载过的包，不论版本是多少，找到了就不会去下载，所以如果要升级依赖，可以使用 ` npm update` 或者显示安装 ` npm install cross-env --save`
* ` npm install` 会先下载项目中的依赖包，然后下载依赖的依赖，这样就会导致，生成的文件是树形结构，并且存在许多重复的包，所以这个时候 ` npm` 就会将依赖扁平化，将依赖的依赖提取到第一层，遇到版本号不一致的也会保留，遇到完全一致的就会删除。
* 最后还会提取依赖中的 ` bin` 文件， ` windows` 操作系统生成 ` cmd` 文件， ` linux` 系统生成软连接

## 查阅的资料 ##

* [2018 年了，你还是只会 npm install 吗？]( https://juejin.im/post/5ab3f77df265da2392364341#heading-13 )
* [npm 模块安装机制简介]( https://link.juejin.im?target=http%3A%2F%2Fwww.ruanyifeng.com%2Fblog%2F2016%2F01%2Fnpm-install.html )

> 
> 
> 
> 每一个小问题深入下去都会头皮发麻，学不动了
> 
>