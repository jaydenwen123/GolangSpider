# 【译】Node.js 日志打印指南 #

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2ae67c0b39fc0?imageView2/0/w/1280/h/960/ignore-error/1) 当你开始使用JavaScript开发时，可能要学习的第一个技能就是如何使用 ` console.log` 将内容打印到控制台。如果你搜索如何调试JavaScript，将会发现数百篇博客和StackOverflow文章指向 ` console.log` 。因为这是一种很常见的方法，我们甚至开始使用像 ` no-console` 这样的linter规则来确保我们不会在生产代码中留下意外的日志语句。但是如果我们真的想通过打印一些东西来提供更多的信息呢?

在这篇文章中，我们将探讨各种需要打印信息的场景；Node.js中 ` console.log` 和 ` console.error` 之间有什么区别；以及如何在不扰乱用户控制台的情况下在库中记录日志。

` console.log( `Let's go!` ); 复制代码`

## 首要理论：Node.js的重要细节 ##

假如我们能在浏览器或者Node.js中使用 ` console.log` 或 ` console.error` ，那么在使用Node.js时，有一件重要的事情是要记住的。在一个名为index.js的文件中使用Node.js编写以下代码时:

` console.log( 'Hello there' ); console.error( 'Bye bye' ); 复制代码`

并在终端使用 ` node index.js` 执行它，你会看到它们并排输出了:

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2ae6d75112c08?imageView2/0/w/1280/h/960/ignore-error/1)

然而，尽管这两者看起来相同，但实际上系统对它们的处理是不同的。查看 [Node.js文档关于console的介绍]( https://link.juejin.im?target=https%3A%2F%2Fnodejs.org%2Fapi%2Fconsole.html ) ，我们可以看到， **` console.log` 输出到 ` stdout` ， ` console.error` 输出到 ` stderr`** 。

每个进程默认有三个可以使用的流： ` stdin` , ` stdout` 和 ` stderr` 。 ` stdin` 流处理进入程序的输入。例如，按钮按下或重定向输出(我们稍后会讲到)。 ` stdout` 流用于应用程序的输出。最后， ` stderr` 用于错误消息。如果你想了解为什么存在 ` stderr` 以及何时使用它，请参阅这篇 [文章]( https://link.juejin.im?target=https%3A%2F%2Fwww.jstorimer.com%2Fblogs%2Fworkingwithcode%2F7766119-when-to-use-stderr-instead-of-stdout ) 。

简而言之，我们可以使用重定向( ` >` )和管道( ` |` )操作符将应用程序的错误和诊断信息结果分离显示。 ` >` 操作符允许我们将 ` stdout` 的输出重定向到文件中，而 ` 2>` 允许我们将 ` stderr` 的输出重定向到文件中。例如，这个命令将“Hello there”导入一个名为 ` Hello.log` 的文件，并将“Bye Bye”导入一个名为 ` error.log` 的文件。

` node index.js > hello.log 2> error.log 复制代码`

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2aee912321eb1?imageView2/0/w/1280/h/960/ignore-error/1)

## 什么情况下我们需要log? ##

既然我们已经了解了一些关于日志的底层技术方面的知识，那么接下来让我们来讨论一下可能需要记录日志的场景。通常这些场景包含以下类别:

* 开发过程中快速调试异常
* 基于浏览器的日志记录，用于分析或诊断
* 记录服务应用日志，以记录传入的请求以及可能发生的任何故障
* 库的可选调试日志，以协助用户排查问题
* 通过CLI打印进度、确认信息或错误

下面的内容我们将跳过的前两个场景，重点介绍跟Node.js有关的后面三个场景。

## 服务应用日志 ##

在服务器上记录日志的原因可能有很多。例如，通过记录传入的请求我们可以用来做信息统计，比如用户遇到有多少404请求，这些请求可能是什么，或者正在使用什么 ` User-Agent` 。我们也想知道什么时候出了问题，原因是什么。

如果你想尝试本文下面内容，请创建一个新的项目目录。在项目目录中创建 ` index.js` 用于编写代码的程序运行入口。运行以下代码初始化项目并安装 ` express` :

` npm init -y npm install express 复制代码`

让我们设置一个带有console.log的中间件的服务器。将以下内容放入index.js文件中:

` const express = require ( 'express' ); const PORT = process.env.PORT || 3000 ; const app = express(); app.use( ( req, res, next ) => { console.log( '%O' , req); next(); }); app.get( '/' , (req, res) => { res.send( 'Hello World' ); }); app.listen(PORT, () => { console.log( 'Server running on port %d' , PORT); }); 复制代码`

我们使用 ` console.log('$0',req)` 用于记录整个对象。 ` console.log` 底层使用 ` util.format` 方法支持 ` %O` 占位符。详细信息可以在 [Node.js官方文档]( https://link.juejin.im?target=https%3A%2F%2Fnodejs.org%2Fapi%2Futil.html%23util_util_format_format_args ) 中了解。

当执行 ` node index.js` 来执行服务器并访问http://localhost:3000时，你会注意到它将打印出许多我们并不真正需要的信息。

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2aef2942db8be?imageView2/0/w/1280/h/960/ignore-error/1)

即使我们将其更改为 ` console.log('%s', req)` 不打印整个对象，也不会得到太多有用的信息。

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2aef4439586e6?imageView2/0/w/1280/h/960/ignore-error/1)

我们可以写我们自己的log函数，只输出我们关心的东西。但是在此之前，我们先讨论下通常需要关心什么。虽然太多信息分散我们注意力的集中，但实际上我们也需要充分的信息。如:

> 
> 
> 
> * 时间戳——知道事情发生的时间
> * 计算机/服务器名称——如果你正在运行一个分布式系统
> * 进程ID——如果你正在使用类似pm2的东西运行多个节点进程
> * 消息——包含一些内容的实际消息
> * 堆栈跟踪——用于记录错误的场景
> * 其他一些额外的变量/信息
> 
> 
> 

此外，既然我们已经知道所有内容都将进入stdout和stderr流，那么我们可以借助它们实现不同级别的日志，以及根据它们配置和筛选日志的能力。

我们可以通过访问进程的各个部分并编写一堆JavaScript来实现所有这些功能，但Node.js最棒的一点是，拥有npm生态系统，而且已经有各种库可供我们使用。例如:

> 
> 
> 
> * pino
> * winston
> * roarr
> * bunyan (注意：这个已经两年没有更新了)
> 
> 
> 

我个人喜欢 ` pino` ，因为它速度快，生态也很好。让我们看看如何使用pino帮助我们进行日志记录。奇妙的是已经有一个 ` express-pino-logger` 包，我们可以使用它来记录请求。

安装 ` pino` 和 ` express-pino-logger` :

` npm install pino express-pino-logger 复制代码`

然后更新 ` index.js` 文件，使用日志记录器和中间件:

` const express = require ( 'express' ); const pino = require ( 'pino' ); const expressPino = require ( 'express-pino-logger' ); const logger = pino({ level : process.env.LOG_LEVEL || 'info' }); const expressLogger = expressPino({ logger }); const PORT = process.env.PORT || 3000 ; const app = express(); app.use(expressLogger); app.get( '/' , (req, res) => { logger.debug( 'Calling res.send' ); res.send( 'Hello World' ); }); app.listen(PORT, () => { logger.info( 'Server running on port %d' , PORT); }); 复制代码`

在这个代码片段中，我们创建了一个 ` pino` 的日志程序实例，并将其传递到 ` express-pino-logger` 中来创建一个新的日志程序中间件以便 ` app.use` 调用。此外，我们在服务启动的时候用 ` logger.info` 替换 ` console.log` ，并向路由添加了一个额外的 ` logger.debug` ，以显示不同级别的日志。

如果通过再次运行 ` node index.js` 启动服务器。你会看到一个非常不同的输出，它每一行输出一个JSON。再次访问http://localhost:3000，你将看到添加了一行新的JSON。

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2b75f8e6ddc7c?imageView2/0/w/1280/h/960/ignore-error/1)

如果检查这个JSON，你将看到它包含前面提到的所有信息，比如时间戳。你还可能注意到我们的 ` logger.debug` 语句没有打印出来。这是因为我们使用了默认的日志级别。创建logger实例时，我们通过设置 ` process.env.LOG_LEVEL` 的值改变日志级别，默认值为 ` info` 。通过运行 ` LOG_LEVEL=debug node index.js` ，我们可以调整日志级别显示debug类型的日志。

在此之前，让我们先讨论这样一个事实:现在的输出实际上可读性很差。然而这是故意的。pino遵循一种原则，即更高的性能。我们也可以通过管道(使用 ` |` )将所有进程的日志移动到一个单独的进程中，用于提高其可读性或将数据上载到云服务器。这个过程叫做 [transports]( https://link.juejin.im?target=transports ) 。查看 [关于transports的文档]( https://link.juejin.im?target=http%3A%2F%2Fgetpino.io%2F%23%2Fdocs%2Ftransports ) ，还可以了解为什么 ` pino` 中的错误没有被写入 ` stderr` 。

让我们使用工具 ` pino-pretty` 查看更具可读性的日志版本。在终端执行一下命令:

` npm install --save-dev pino-pretty LOG_LEVEL=debug node index.js | ./node_modules/.bin/pino-pretty 复制代码`

现在，使用 ` |` 操作符，所有的日志将通过管道传输到 ` pino-pretty` ，你的输出应该变得清晰，包含了关键信息并且被着色。再次访问http://localhost:3000，还应该能看到 ` debug` 级别的消息。

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2b7df5f380e2f?imageView2/0/w/1280/h/960/ignore-error/1)

有许多现成的传输工具可以美化或转换日志。你甚至可以用 ` pino-colada` 工具使其支持表情符号的显示。这些将对你本地的开发非常有用。在生产环境中运行服务器之后，你可能希望将日志导入到另一个传输中，使用 ` >` 将日志写入磁盘，以便稍后处理它们，或者使用 ` tee` ( https://link.juejin.im?target=https%3A%2F%2Fen.wikipedia.org%2Fwiki%2FTee_(command) ) 之类的命令进行处理。

[官方文档]( https://link.juejin.im?target=https%3A%2F%2Fgetpino.io%2F ) 还介绍关于日志文件归档、过滤和将日志写入不同文件等内容。

## 你的库日志 ##

既然我们已经了解了如何为服务器应用程序高效地编写日志，为什么不为我们编写的库使用相同的技术呢?

问题是，我们希望打印出库用于调试的内容，但也不能混淆使用者的应用程序。如果需要调试某些东西，使用者应该能够启用日志。你的库在默认情况下应该是静默的，并将是否打印日志留给使用者决定。

` express` 就是一个很好的例子。 ` express` 的底层做了很多事情，在调试应用程序时，你可能想了解一下底层的情况。如果我们查阅 ` express` 文档，便会注意到启动相关日志只需要在命令前加上 ` DEBUG=express:*` :

` DEBUG=express:* node index.js 复制代码`

使用现有的应用程序运行该命令，你将看到许多额外的输出，这些输出将帮助你调试问题。

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2b7e3013cb4e5?imageView2/0/w/1280/h/960/ignore-error/1)

如果没有启用调试日志记录，就不会看到这些。这是是通过一个名为 ` debug` 的包来实现的。它允许我们在指定“名称空间”下编写消息，如果库的使用者在调试环境变量中包含与之匹配的名称空间或通配符，它将输出这些消息。

要使用debug库，首先安装它:

` npm install debug 复制代码`

让我们通过创建一个新文件来尝试它，该文件将模拟我们的库 ` random-id.js` ，并在其中放置以下代码:

` const debug = require ( 'debug' ); const log = debug( 'mylib:randomid' ); log( 'Library loaded' ); function getRandomId ( ) { log( 'Computing random ID' ); const outcome = Math.random() .toString( 36 ) .substr( 2 ); log( 'Random ID is "%s"' , outcome); return outcome; } module.exports = { getRandomId }; 复制代码`

以上代码创建一个名称空间为 ` mylib:randomid` 的新调试日志实例，其打印了两条消息。让我们在上一章的 ` index.js` 中使用它:

` const express = require ( 'express' ); const pino = require ( 'pino' ); const expressPino = require ( 'express-pino-logger' ); const randomId = require ( './random-id' ); const logger = pino({ level : process.env.LOG_LEVEL || 'info' }); const expressLogger = expressPino({ logger }); const PORT = process.env.PORT || 3000 ; const app = express(); app.use(expressLogger); app.get( '/' , (req, res) => { logger.debug( 'Calling res.send' ); const id = randomId.getRandomId(); res.send( `Hello World [ ${id} ]` ); }); app.listen(PORT, () => { logger.info( 'Server running on port %d' , PORT); }); 复制代码`

重新运行服务器，但是这次使用 ` DEBUG=mylib:randomid node index.js` ，它将打印我们的“库”的调试日志。

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2b7ece91dcd24?imageView2/0/w/1280/h/960/ignore-error/1)

有趣的是，如果你的库使用者希望将此调试信息放入他们的 ` pino` 日志中，他们可以使用pino团队提供的 ` pino-debug` 库来正确格式化这些日志。

使用以下方法安装库:

` npm install pino-debug 复制代码`

在第一次使用debug之前，需要初始化 ` pino-debug` 。最简单的方法是在启动脚本之前使用Node.js的 ` -r` 或 ` ——require` 标志 ( https://link.juejin.im?target=https%3A%2F%2Fnodejs.org%2Fapi%2Fcli.html%23cli_r_require_module ) 来引入模块。使用如下命令重新运行服务器(假设已经安装了 ` pino-colada` ):

` DEBUG=mylib:randomid node -r pino-debug index.js | ./node_modules/.bin/pino-colada 复制代码`

现在，你将看到库的调试日志与应用程序日志的格式相同。

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2b7f4ddaf00fc?imageView2/0/w/1280/h/960/ignore-error/1)

## CLI(命令行界面)输出 ##

在这篇文章中，我们将讨论的最后一种情况是CLIs而不是库的特殊日志记录情况。我的原则是将逻辑日志与CLI输出的“日志”分开。对于任何逻辑日志，都应该使用debug之类的库。这样，你或其他人就可以重用逻辑，而不受CLI特定用例的限制。

当你的Node.js应用采用CLI构建时，你可能希望通过添加颜色、标记或以一种特定的具有视觉吸引力的方式格式化内容，使其看起来更漂亮。然而，在使用CLI构建时，你应该记住以下几个场景。

一种场景是，你的CLI可能在持续集成(CI)系统的上下文中使用，因此你可能希望删除颜色或任何花哨的装饰输出。一些CI系统设置了一个称为CI的环境标志。如果你想更安全地检查你是否在CI中，可以使用 ` is-CI` ( https://link.juejin.im?target=https%3A%2F%2Fwww.npmjs.com%2Fpackage%2Fis-ci ) 这样的包，它已经支持许多CI系统。

一些库，如 ` chalk` ，已经为你检测是否CI环境并为你删除颜色。接下来让我们一起看下使用它之后的样子。

使用npm安装chalk并创建一个名为 ` clip .js` 的文件。放入以下代码:

` const chalk = require ( 'chalk' ); console.log( '%s Hi there' ,chalk.cyan( 'INFO' )); 复制代码`

现在，如果你使用 ` node clip.js` 运行这个脚本，你将看到彩色的输出。

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2b88fee088a2f?imageView2/0/w/1280/h/960/ignore-error/1)

但是如果你用 ` CI=true node clip .js` 运行它，你会看到颜色被抑制了:

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2b8921dba4571?imageView2/0/w/1280/h/960/ignore-error/1)

你要记住的另一个场景是，如果你的stdout运行在终端模式中，表示内容写入终端。如果是这种情况，我们可以使用 [boxen]( https://link.juejin.im?target=https%3A%2F%2Fnpm.im%2Fboxen ) 之类的东西来显示所有漂亮的输出。如果不是，很可能输出被重定向到文件或管道的某个地方。

你可以通过检查相应流上的 [isTTY]( https://link.juejin.im?target=https%3A%2F%2Fnodejs.org%2Fapi%2Fprocess.html%23process_a_note_on_process_i_o ) 属性来检查 ` stdin` 、 ` stdout` 或 ` stderr` 是否处于终端模式,例如:process.stdout.isTTY。TTY代表“teletypewriter(电传打字机)”，在本例中特指终端。

根据Node.js进程的启动方式，这三个流的值可能有所不同。你可以在 [Node.js文档的“process I/O”]( https://link.juejin.im?target=https%3A%2F%2Fnodejs.org%2Fapi%2Fprocess.html%23process_a_note_on_process_i_o ) 部分了解更多。

让我们看看 ` process.stdout` 的值。 ` isTTY` 在不同的情况下是不同的。更新你的 ` clil.js` 文件，以检查它:

` const chalk = require ( 'chalk' ); console.log (process.stdout.isTTY); console.log( '%s Hi there' ，chalk.cyan( 'INFO' )); 复制代码`

现在在你的终端中运行 ` node clip.js` ，你会看到true后面跟着我们的彩色消息。

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2b8987f78a178?imageView2/0/w/1280/h/960/ignore-error/1)

之后运行相同的东西，但重定向输出到一个文件，并检查内容后运行:

` node clip .js > output.log cat output.log 复制代码`

你将看到，这一次它打印的是undefined，后面跟着一条纯色的消息，因为stdout的重定向关闭了stdout的终端模式。 ` chalk` 使用了 ` support-color` ( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fchalk%2Fsupports-color%23readme ) 检查相应流上是否支持TTY。

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2b89d93bc4e50?imageView2/0/w/1280/h/960/ignore-error/1)

类似 ` chalk` 这样的工具已经为你处理了这种场景。但是，在开发CLI时，你应该始终了解CLI可能在CI模式下运行或重定向输出的情况。它还可以帮助您进一步获得CLI的体验。例如，你可以在终端中以漂亮的方式排列数据，如果isTTY未定义，则切换到更容易解析的方式。

## 总结 ##

开始使用JavaScript并使用console.log记录第一行代码非常快，但是当你将代码投入生产时，你应该考虑更多关于日志的内容。这篇文章仅仅介绍了各种方法和可用的日志解决方案。但它不包含你需要知道的一切。我建议你查看一些你最感兴趣的开源项目，了解它们如何解决日志记录问题以及使用哪些工具。现在去记录所有的信息,而不是仅仅打印日志吧😉。

原文： [A Guide to Node.js Logging]( https://link.juejin.im?target=https%3A%2F%2Fwww.twilio.com%2Fblog%2Fguide-node-js-logging )