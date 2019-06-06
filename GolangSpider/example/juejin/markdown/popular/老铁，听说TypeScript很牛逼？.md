# 老铁，听说TypeScript很牛逼？ #

TypeScript是啥，有人说TypeScript = Type + Script，实际我觉得更准确的应该是TS = Java(JS)或者 TS = C#(JS)，使用Java/C#的语法写JS，并且为了能让JSer能更容易接受，它的语法又不能直接把Java/C#的那套搬过来，要贴近于JS. 所以官方说法是TS是JS的超集，超的地方就是引入了Java/C#的语法特性（但是官方又不承认）。

TypeScript出现的目的是什么，如下 [官方说法]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2FMicrosoft%2FTypeScript%2Fblob%2Fmaster%2Fdoc%2Fspec.md%231-introduction ) ：

> 
> 
> 
> We designed TypeScript to meet the needs of the JavaScript programming
> teams that build and maintain large JavaScript programs.
> 
> 

即给大型的JS应用程序使用和维护，并且它不是为了提供一个必定是正确的类型系统，而是在正确性和生产力之间找到了 [一个平衡]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2FMicrosoft%2FTypeScript%2Fwiki%2FTypeScript-Design-Goals ) ：

> 
> 
> 
> Not Apply a sound or "provably correct" type system. Instead, **strike a
> balance between correctness and productivity**.
> 
> 

所以这里官方似乎已经对要不要用TypeScript给了一个答案：适合于大型的JS项目，不用的话直接写JS生产力会更高，用了的话牺牲了部分生产力但是换来了正确性，但也不是保证100%准确的。这个正确性体现在哪里呢？

## 1. TypeScript的类型正确性 ##

TS是一种强类型语言，变量需要指定类型（内置或者自定义），不同类型的变量不能相互赋值，这个可以提前发现一些运行时错误，例如当往localStorage写数据的时候，类型必须得是字符串，这也是很多人不小心就犯的错误，如下代码：

` const storageManager = { set (key, value) { try { window.localStorage.setItem(key, value); } catch (e) { console.error(e); } }, get (key) { return window.localStorage[key]; } }; storageManager.set( 'pos' , { x : 5 , y : 8 }); storageManager.get( 'pos' ); // 这里取出来的值是[object Object] 复制代码`

而借助TS能够解决这种情况，就是通过设定参数的类型，如下代码所示：

` const storageManager = { // value必须得是string类型 set (key, value: string ) { try { window.localStorage.setItem(key, value); } catch (e) { console.error(e); } } }; storageManager.set( 'pos' , {x: 5 , y: 8 }); 复制代码`

给参数value添加了一个string类型的限定，这段代码在使用ts编译成js的时候将会报错：

![](https://user-gold-cdn.xitu.io/2019/5/25/16aeeeff53838b49?imageView2/0/w/1280/h/960/ignore-error/1)

意思是说两个类型不匹配，不能赋值。所以我们不用运行也不用取出localStorage里的东西也能够知道这里发生了问题。但是这样就能保证万无一失了吗，如下面的例子：

` // 从某个地方获取到数据，如input输入框，或者是url的参数 let data: string = '{"x": 5, "y": 8}' ; storageManager.set( 'pos' , JSON.parse(data)); storageManager.get( 'pos' ); // 这里取出来的值是[object Object] 复制代码`

我们从输入框或者url参数取到了字符串的数据，然后在传递参数的时候用了JSON.parse（例如误以为需要传Object），这个时候我们再将重新编译，TSC就不会报错，能够通过，但是运行结果依据悲剧。所以你可能还得这么做，在函数里面判断参数类型：

` const storageManager = { set (key, value) { if ( typeof value !== 'string' ) { throw new Error ( 'value should be type of string' ); } // 其它代码略 } }; 复制代码`

一旦有人传进来的参数不是string，那么控制台立刻抛异常。虽然是运行时才发现的错误，但是你总不能说我写了段代码连跑一下都没有就提交代码了吧。

这里我们也看到TypeScript强类型的缺陷，它是一种外挂的强类型，无法判断一些原生API的类型，所以和内置语言层面的强类型还是有差异的。

还有一种常见的错误是把API拼错，如下面的例子：

` let data: string = 'Version-1.3.1' ; storageManager.set( 'pos' , data.toLowerCace()); 复制代码`

编译这段代码将会提示：

![](https://user-gold-cdn.xitu.io/2019/5/25/16aeeeff53bf7f1e?imageView2/0/w/1280/h/960/ignore-error/1)

这里我们把toLowerCase拼错了，TSC给了我们提示说ES5的String是没有toLowerCace的。这确实也是使用TS的好处，就好像你在写Java的时候，如果你写了一个不存在的方法，IDE会直接标红。同样地提前发现了一些需要在运行才能发现的错误。这个还有点像ESLint的 [no-undef]( https://link.juejin.im?target=https%3A%2F%2Fcn.eslint.org%2Fdocs%2Frules%2Fno-undef ) 规则，当我们把一个变量名写错或者忘记import就直接用的时候就会导致该变量undefined，eslint检查便会给出提示。但是和ESLint的区别在于，ESLint检查的仍然是原生JS语法，不需要附加一个TS的语法，只是ESLint不检查方法名是否存在。

那么这种方法名写错的概率到底有多高？实际上JS的API并没有几个，这也正是我写JS一直使用vim编辑器的原因，而之前写Java（Spring MVC）的时候便使用了IDE。

这种强类型附带的好处便是IDE会提示API，如下Sulime编辑提示string的方法：

![](https://user-gold-cdn.xitu.io/2019/5/25/16aeeeff540fdc25?imageView2/0/w/1280/h/960/ignore-error/1)

但实际上这个问题并不大，只要某个单词当前文件出现过一次之后，那么下次便可使用自动补全，不管是vim还是其它编辑器。

另外TS还对API做了DOCS，如下图所示：

![](https://user-gold-cdn.xitu.io/2019/5/25/16aeeeff53ea7d8a?imageView2/0/w/1280/h/960/ignore-error/1)

这样我们就不用去查文档，而对自己写的类型也可以添加相应的DOCS，方便在输入的时候便可知道函数的作用是什么，参数和返回值分别又是什么。但实际上JS也是可以的，如下图所示：

![](https://user-gold-cdn.xitu.io/2019/5/25/16aeeeff60c0278e?imageView2/0/w/1280/h/960/ignore-error/1)

只要你在定义变量的时候给出一个初始值，便能使用TS的文档提示（编辑器安装TS的提示插件），所以从这点看的话只是出于自动提示目的也可以不写TS.

除了类型之外，TS还提供了一套完整的OOP实现。

## 2. 一套完整的OOP ##

ES5的原型继承在其它面向对象语言如Java/C++等来看是一种很独特的存在，它连类的私有变量、公有变量的概念都没有，更别提什么接口类、抽象类、虚函数这些OOP必备的东西。但是不要紧，可以写TS，TS还你一个真正的OOP，让你重新找回OOP的感觉，进而实现各种设计模式。

首先连私有变量的概念都没有，怎么能算是OOP封装呢，所以TS具备公有（public）、私有（private）、保护（protected）三种类型属性，其中公有属性能被类的实例所访问，而私有无法实例所访问，只能在类定义的方法里面被访问到，而保护类型的虽然无法被类实例所访问，但是能够被继承的类所访问。这个特性和其它OOP语言对齐。

如果看过《Head Frist设计模式》这本书的读者应该知道，它里面是用Java写的，几乎所有的设计模式都借助了抽象类和接口类进行实现，所以在Java里面，如果没有抽象类/接口类的存在似乎就玩不了设计模式。例如书中的策略模式的例子是这样的，有这么一个需求，实现两种鸭子，一种是会飞的，另一种是不会飞的，并且一只鸭子可以随时切换飞行的行为，如从不会飞切到会飞，也就是说这种飞行的行为就是一种策略，切换飞行行为就是切换策略，我们可依照书中的代码用TS实现一遍，如下代码所示：

` // 由于有两种不同的飞行行为，所以要有一个接口类 interface FlyBehaviour { fly () : void ; } // 会飞的行为，实现怎么飞 class FlyWithWings implements FlyBehaviour { fly () : void { console.log( 'I\'m flying' ); } } // 不会飞的行为 class FlyWithNoWays implements FlyBehaviour { fly () : void { console.log( 'I can\'t fly' ); } } 复制代码`

飞行的策略都是一种叫做FlyBehaviour的自定义类型，因为它们要能赋值给同一个变量，所以需要使用一个接口类。

然后在鸭子类里面使用这种自定义类型，如下代码所示：

` // 由于有不同类型的鸭子，有相同和不同行为，所以我需要写一个抽象类 abstract class Duck { // 它组合了一个飞行策略 protected flyBehaviour: FlyBehaviour; quak () : void { console.log( 'quak' ); } abstract performFly () : void ; // 必须在派生类中实现 } // 现在有一种野鸭，它是一种鸭子，所以用继承 class MallardDuck extends Duck { constructor ( ) { super (); // 默认不会飞 this.flyBehaviour = new FlyWithNoWays(); } // 可以随时改变策略 setFlyBehaviour (flyBehaviour) { this.flyBehaviour = flyBehaviour; } performFly () : void { this.flyBehaviour.fly(); } } 复制代码`

驱动代码如下所示：

` let mallardDuck = new MallardDuck(); mallardDuck.performFly(); // 输出I'cant fly // 改变飞行策略 mallardDuck.setFlyBehaviour( new FlyWithWings()); mallardDuck.performFly(); // 输出I'm flying 复制代码`

这样我们就实现了一个策略模式，这几段代码看起来非常的OOP，什么接口类、抽象类、继承、实现都用上了，代码看起来也十分地高大上，但实际上，我们真的需要写这么多类的代码吗？如果是我的话，我可能会这么实现：

` // fly-behaviour.js const flyBehaviour = { flyWithWings () { console.log( 'I\'m flying' ); }, flyWithNoWay () { console.log( 'I can\'t fly' ); } }; export default flyBehaviour; // import flyBehaviour from 'fly-behaviour.js'; class MallardDuck { constructor () { this.flyType = 'flyWithNoWay' ; } setFlyBehaviour (type) { if ( typeof flyBehaviour[type] === 'undefined' ) { throw new Error ( 'flying type is not support' ); } this.flyType = type; } performFly () { flyBehaviour[ this.flyType](); } }; let mallardDuck = new MallardDuck(); mallardDuck.performFly(); // 输出I'cant fly mallardDuck.setFlyBehaviour( 'flyWithWings' ); mallardDuck.performFly(); // 输出I'm flying 复制代码`

上面的代码的思想是使用一个Object表示策略模式，通过不同的类型区分不同的策略。实际上在实际的写代码过程中我们发现很多类其实只需要实例化一次，而JS的Object正好可以表示这种实例化一次的对象。所以这里的策略模式我们连一个类都没有写。真正需要写类的可能就是那种每个对象需要有自己的数据，不同实例之间的数据不一样，例如上面的Duck，不同的Duck有不同的飞行行为，在写游戏模拟外部世界，或者是UI的弹框类等经常需要使用类来实现。

还有一个问题是如果TS的强类型真的这么好的话，那么ES标准为什么不规定呢，例如在let/const上再增加string/number之类的定义变量的方式（我们看到私有变量已经有了），而是需要别人给它加一个外挂

你可能已经看了很多篇介绍TS是如何地优秀，如何地牛逼的文章，突然发现我这一篇的味道不太一样。但我并不是说TS不好，而只是反对TS“政治”正确性的言论，写TS就是对的，写JS就是不提倡的，正如写Vue/React就是对的，写jQuery就是不对的。当然我也建议去学习TS，例如如果连什么是枚举类型，什么是泛类型/模板类型，什么是抽象类/虚类这些概念都没有，但是又不想去学习Java/C++这种典型的强类型和OOP语言的话，那么学习TS是很有意义，不管用或不用。