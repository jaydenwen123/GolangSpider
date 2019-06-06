# 有助于理解前端工具的 node 知识 #

## 缘起 ##

平时写惯了业务代码之后，如果想要了解下 webpack 或者 vue-cli，好像是件很难上手的事情🙁 。拿 webpack 来说，我们可能会对配置熟悉点，但常常一段时间过后又忘了，感觉看起来不是很好懂。其实类似这种打包工具、构建工具我们最好应该先去学习一下 node 的一些基础知识，然后再回过头来看这些工具，就会有柳暗花明又一村的感觉，因为这些工具是用 node 写出来的🤯。
想想我们是不是时常看到过这种东西： ` const path = require('path');` 。假设你学过前端框架但没学过 node，你看到这句话的时候就会一头雾水，好像知道它是弄路径的，但具体这是哪里来的，常用来做什么就不得而知了，我起初看的感觉就是这样🤨。
后来才知道这其实是 node 的内置模块，因为这些构建工具或打包工具是用 node 来执行的，只要我们有装 node，它里面的内置模块就能直接引用，不用另外安装。所以强烈建议大家要是想了解这类工具最好先学习一下 node，不然会总是懵逼的🧐。
言归正传，本篇就来简要讲述一下 node 的一些常用内置模块。

## node 初识 ##

### node 是什么 ###

首先 node 不是一门后台语言而是一个环境，一个能够让 js 运行在服务器的环境，这个环境就好比是服务器上的浏览器（虽然不是很恰当），但正是因为有了它才使得 js 变成了一门后台语言。

### node 遵循的规范 ###

其次 node 遵循的是 CommonJs 规范，什么意思？其实就是规定了导入导出的方式😬，就向下面这样：

` require ( './module' ) module.exports = { a : 1 , } exports.a = 1 ; 复制代码`

这就是 node 的规范，用 ` require` 导入、用 ` module.exports` 导出。那 node 为什么不支持 ESM（就是用 ` import` 导入、用 ` export` 导出）规范呢，因为它出现的比较早，仅此而已，然后一时半会儿还改不过来，以后应就会支持了。另外，我们时常在 webpack 里看到 ` require()` 字样却没有看见 ` import()` 就是因为 webpack 是要用 node 来执行的，而 node 目前只支持 ` require()` 。
这里顺带来一张各种规范图（这种东西容易忘，当作历史看看就行🙄），如下：

![](https://user-gold-cdn.xitu.io/2019/5/20/16ad4be6bc803f68?imageView2/0/w/1280/h/960/ignore-error/1)

### require 寻找依赖 ###

` require()` 里面的参数有两种写法，一种带路径一种不带路径。就像下面这样：

` require ( './module' ); // 带相对路径 require ( '/module' ); // 带绝对路径 require ( 'module' ); // 不带路径 复制代码`

这种不带路径的 ` require('module')` 引入方式，可能是内置模块，也可能是第三方模块，内置模块优先查找，没有的话就是第三方模块了，它会先从当前目录的 ` node_modules` 里面查找，没有的话就到父目录下的 ` node_modules` 里面去找，如此向上追溯，直到根目录下的 ` node_modules` 目录，要是还没有的话就会到全局里面去找，大概是这么一个搜索过程。
另外一种带路径的方式，就会沿着路径去找，如果没有找到则会尝试将当前目录作一个包来加载。此外，使用绝对路径的速度查找最快，当然了，node 也对路径查找做了缓存机制。

### node 模块包装 ###

node 在解析每个模块（js 文件）时，会对每个模块进行包装，就是在代码外面加一个闭包，并且向里传递五个参数，这样就保证了每个模块之间的独立，就像下面这样：

` ( function ( exports, require, module, __filename, __dirname ) { // module: 表示当前模块 // __filename: 当前模块的带有完整绝对路径的文件名 // __dirname: 当前模块的完整绝对路径 module.exports = exports = this = {}; // 我们的代码就在这里... return module.exports; })() 复制代码`

想想我们平时是不是常在 webpack 里面看到 ` __dirname` 这种东西，我们既没有引入也没有声明它，为什么能够直接使用呢，就是因为这个原因😮。

### node 的应用场景 ###

一般来说，node 主要应用于以下几个方面：

* 自动化构建等工具
* 中间层
* 小项目

第一点对于前端同学来说应该是重中之重了，什么工程化、自动构建工具就是用 node 写出来的，它是前端的一大分水岭之一，是块难啃的骨头，所以我们必须拿下，不然瓶颈很快就到了。如果你能熟练应用 node 的各种模块（系统模块 + 第三方模块），那么恭喜你，你又比别人牛逼了一截😎。

### node 的优点 ###

* 适合前端大大们
* 基于事件驱动和无阻塞的I/O（适合处理并发请求）
* 性能较好（别人做过性能分析）

## node 内置模块 ##

ok，废话了这么多，咱们赶紧来看看一些常见的 node 基础模块吧。相信掌握这些对你学习 webpack 和 vue-cli 等工具是有很大帮助的✊ 。

### http 模块 ###

这是 node 最最基础的功能了，我们用 ` node http.js` 运行一下下面的文件就能开启一个服务器，在浏览器中输入 ` http://localhost:8888` 即可访问，http.js 具体内容如下：

` // http.js const http = require ( 'http' ); http.createServer( ( req, res ) => { // 开启一个服务 console.log( '请求来了' ); // 如果你打开 http://localhost:8888，控制台就会打印此消息 res.write( 'hello' ); // 返回给页面的值，也就是页面会显示 hello res.end(); // 必须有结束的标识，否则页面会一直处于加载状态 }).listen( 8888 ); // 端口号 复制代码`

### fs 文件系统 ###

由于 js 一开始是用来开发给浏览器用的，所以它的能力就局限于浏览器，不能直接对客户端的本地文件进行操作，这样做的目的是为了保证客户端的信息安全，当然了，通过一些手段也可以操作客户端内容（就像 ` <input type='file'>` ），但是需要用户手动操作才行。
但是当 js 作为后台语言时，就可以直接对服务器上的资源文件进行 I/O 操作了。这也是 node 中尤为重要的模块之一（操作文件的能力），这在自动化构建和工程化中是很常用的。它的主要职责就是读写文件，或者移动复制删除等。fs 就好比对数据库进行增删改查一样，不同的是它操作的是文件。下面我们来具体看看代码用例：

` const fs = require ( 'fs' ); // 写入文件：fs.writeFile(path, fileData, cb); fs.writeFile( './text.txt' , 'hello xr!' , err => { if (err) { console.log( '写入失败' , err); } else { console.log( '写入成功' ); } }); // 读取文件：fs.readFile(path, cb); fs.readFile( './text.txt' , (err, fileData) => { if (err) { console.log( '读取失败' , err); } else { console.log( '读取成功' , fileData.toString()); // fileData 是二进制文件，非媒体文件可以用 toString 转换一下 } }); 复制代码`

需要注意的是 readFile 里面的 fileData 是原始的二进制文件🤨（em...就是计算机才看的懂的文件格式），对于非媒体类型（如纯文本）的文件可以用 ` toString()` 转换一下，媒体类型的文件以后则会以流的方式进行读取，要是强行用 ` toString()` 转换的话会丢失掉原始信息，所以不能乱转。二进制和 ` toString` 的效果就像下面这样：

![](https://user-gold-cdn.xitu.io/2019/5/7/16a910dbe0821a57?imageView2/0/w/1280/h/960/ignore-error/1) 另外，和 fs.readFile（异步） 和 fs.writeFile（异步）相对应的还有 fs.readFileSync（同步）和 fs.writeFileSync（同步），fs 的大多方法也都有同步异步两个版本，具体取决于业务选择，一般都用异步，不知道用啥的话也用异步。

### path 路径 ###

这个模块想必大家应该都并不陌生，🧐瞟过 webpack 的都应该看过这个东东。很显然，path 就是来处理路径相关东西的，我们直接看下面的常见用例就能够体会到：

` const path = require ( 'path' ); let str = '/root/a/b/index.html' ; console.log(path.dirname(str)); // 路径 // /root/a/b console.log(path.extname(str)); // 后缀名 // .html console.log(path.basename(str)); // 文件名 // index.html // path.resolve() 路径解析，简单来说就是拼凑路径，最终返回一个绝对路径 let pathOne = path.resolve( 'rooot/a/b' , '../c' , 'd' , '..' , 'e' ); // 一般用来打印绝对路径，就像下面这样，其中 __dirname 指的就是当前目录 let pathTwo = path.resolve(__dirname, 'build' ); // 这个用法很常见，你应该在 webpack 中有见过 console.log(pathOne, pathTwo, __dirname); // pathOne => /Users/lgq/Desktop/node/rooot/a/c/e // pathTwo => /Users/lgq/Desktop/node/build // __dirname => /Users/lgq/Desktop/node 复制代码`

嗯，下次看到 path 这个东西就不会迷茫了。

### url 模块 ###

很显然这是个用来处理网址相关东西的，也是我们必须要掌握的，主要用来获取地址路径和参数的，就像下面这样：

` const url = require ( 'url' ); let site = 'http://www.xr.com/a/b/index.html?a=1&b=2' ; let { pathname, query } = url.parse(site, true ); // url.parse() 解析网址，true 的意思是把参数解析成对象 console.log(pathname, query); // /a/b/index.html { a: '1', b: '2' } 复制代码`

### querystring 查询字符串 ###

这个主要是用来把形如这样的字符串 ` a=1&b=2&c=3` （&和=可以换成别的）解析成 ` { a: '1', b: '2', c: '3' }` 对象，反过来也可以把对象拼接成字符串，上面的 url 参数也可以用 querystring 来解析，具体演示如下：

` const querystring = require ( 'querystring' ); let query = 'a=1&b=2&c=3' ; // 形如这样的字符串就能被解析 let obj = querystring.parse(query); console.log(obj, obj.a); // { a: '1', b: '2', c: '3' } '1' query = 'a=1&b=2&c=3&a=3' ; // 如果参数重复，其所对应的值会变成数组 obj = querystring.parse(query); console.log(obj); // { a: [ '1', '3' ], b: '2', c: '3' } // 相反的我们可以用 querystring.stringify() 把对象拼接成字符串 query = querystring.stringify(obj); console.log(query); // a=1&a=3&b=2&c=3 复制代码`

### assert 断言 ###

这个我们直接看下面代码就知道它的作用了：

` // assert.js const assert = require ( 'assert' ); // assert(条件，错误消息)，条件这部分会返回一个布尔值 assert( 2 < 1 , '断言失败' ); 复制代码`

` node assert.js` 运行一下代码就能看到如下结果：

![](https://user-gold-cdn.xitu.io/2019/5/7/16a91f99a6b29abc?imageView2/0/w/1280/h/960/ignore-error/1) 上图是断言失败的例子，如果断言正确的话，则不会有任何提示，程序会继续默默往下执行。所以断言的作用就是先判断条件是否正确（有点像 if），如果条件返回值为 ` false` 则阻止程序运行，并抛出一个错误，如果返回值为 ` true` 则继续执行，一般用于函数中间和参数判断。
另外，这里再介绍两种 equal 用法（assert 里面有好多种 equal，这里举例其中的两种）：

` // assert.js const assert = require ( 'assert' ); const obj1 = { a : { b : 1 } }; const obj2 = { a : { b : 1 } }; const obj3 = { a : { b : '1' } }; // assert.deepEqual(变量，预期值，错误信息) 变量 == 预期值 // assert.deepStrictEqual(变量，预期值，错误信息) 变量 === 预期值 // 同样也是错误的时候抛出信息，正确的时候继续默默执行 assert.deepEqual(obj1, obj2, '不等哦' ); // true assert.deepEqual(obj1, obj3, '不等哦' ); // true assert.deepStrictEqual(obj1, obj2, '不等哦' ); // true assert.deepStrictEqual(obj1, obj3, '不等哦' ); // false，这个会抛出错误信息 复制代码`

### stream 流 ###

stream 又叫做流，大家或多或少应该有听过这个概念，那具体是什么意思呢？在这里，你可以把它当做是前面说过的 ` fs.readFile` 和 ` fs.writeFile` 的升级版。
我们要知道 ` readFile` 和 ` writeFile` 的工作流程 是先把整个文件读取到内存中，然后再一次写入，这种方式对于稍大的文件就不适用了，因为这样容易导致内存不足，所以更好的方式是什么呢？就是边读边写啦，业界常说成管道流，就像水流经过水管一样，进水多少，出水就多少，这个水管就是占用的资源（内存），就那么大，这我们样就能合理利用内存分配啦，而不是一口气吃成个胖子，有吃撑的风险（就是内存爆了🤐）。

` const fs = require ( 'fs' ); // 读取流：fs.createReadStream(); // 写入流：fs.createWriteStream(); let rs = fs.createReadStream( 'a.txt' ); // 要读取的文件 let ws = fs.createWriteStream( 'a2.txt' ); // 输出的文件 rs.pipe(ws); // 用 pipe 将 rs 和 ws 衔接起来，将读取流的数据传到输出流（就是这么简单的一句话就能搞定） rs.on( 'error' , err => { console.log(err); }); ws.on( 'finish' , () => { console.log( '成功' ); }) 复制代码`

流式操作，就是一直读取，它是个连续的过程，如果一边快一边慢，或者一边出错没衔接上也没关系，它会自动处理，不用我们自己去调整其中的误差，是个优秀的模块没错了👍。另外，我们没有直接使用 stream 模块，是因为 fs 模块引用了它并对其做了封装，所以用 fs 即可。

### zlib 压缩 ###

这个用法简单，作用也明了，直接看下面的代码就能理解：

` const fs = require ( 'fs' ); const zlib = require ( 'zlib' ); let rs = fs.createReadStream( 'tree.jpg' ); let gz = zlib.createGzip(); let ws = fs.createWriteStream( 'tree.jpg.gz' ); rs.pipe(gz).pipe(ws); // 原始文件 => 压缩 => 写入 rs.on( 'error' , err => { console.log(err); }); ws.on( 'finish' , () => { console.log( '成功' ); }) 复制代码`

## 小结 ##

ok👌，以上就是本章要讲的一些 node 知识（比较基础，大家凑合看看）。当然除此之外，还有 util、Buffer、Event、crypto 和 process 等其他内置模块，这里就不一一赘述了，希望大家能够多动手多敲两下代码多实践，毕竟纸上得来终觉浅嘛💪。如果你能用好 node 的各种模块，那么转后端也就拥有了无限可能性😋（其实前端的坑大的超乎你想像😭）。
最后的最后，安利一下自己的文章，勿喷，哈哈！
[1、基于 vue-cli3 打造属于自己的 UI 库]( https://juejin.im/post/5c95c61f6fb9a070c40acf65 )
[2、仿 vue-cli 搭建属于自己的脚手架]( https://juejin.im/post/5c94fef7f265da60fd0c15e8 )
[3、this.$toast() 了解一下？]( https://juejin.im/post/5ca20e426fb9a05e42555d1d#comment )