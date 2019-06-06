# TypeScript - 一种思维方式 #

**摘要：** 学会TS思考方式。

* 原文： [TypeScript - 一种思维方式]( https://link.juejin.im?target=https%3A%2F%2Fzhuanlan.zhihu.com%2Fp%2F63346965 )
* 作者： [zhangwang]( https://link.juejin.im?target=https%3A%2F%2Fwww.zhihu.com%2Fpeople%2Fzhang-wang-79%2Factivities )

**[Fundebug]( https://link.juejin.im?target=https%3A%2F%2Fwww.fundebug.com%2F ) 经授权转载，版权归原作者所有。**

电影《降临》中有一个观点，语言会影响人的思维方式，对于前端工程师来说，使用 typescript 开发无疑就是在尝试换一种思维方式做事情。

其实直到最近，我才开始系统的学习 typescript ，前后大概花了一个月左右的时间。在这之前，我也在一些项目中模仿他人的写法用过 TS，不过平心而论，在这一轮系统的学习之前，我并不理解 TS。一个多月前，我理解的 TS 是一种可以对类型进行约束的工具，但是现在才发现 TS 并不简单是一个工具，使用它，会影响我写代码时的思考方式。

### TS 怎么影响了我的思考方式 ###

对前端开发者来说，TS 能强化了「面向接口编程」这一理念。我们知道稍微复杂一点的程序都离不开不同模块间的配合，不同模块的功能理应是更为清晰的，TS 能帮我们梳理清不同的接口。

#### 明确的模块抽象过程 ####

TS 对我的思考方式的影响之一在于，我现在会把考虑抽象和拓展看作写一个模块前的必备环节了。当然一个好的开发者用任何语言写程序，考虑抽象和拓展都会是一个必备环节，不过如果你在日常生活中使用过清单，你就会明白 TS 通过接口将这种抽象明确为具体的内容的意义所在了，任何没有被明确的内容，其实都有点像是可选的内容，往往就容易被忽略。

举例来说，比如说我们用 TS 定义一个函数，TS 会要求我们对函数的参数及返回值有一个明确的定义，简单的定义一些类型，却能帮助我们定位函数的作用，比如说我们设置其返回值类型为 ` void` ，就明确的表明了我们想利用这个函数的副作用；

把抽象明确下来，对后续代码的修改也非常有意义，我们不用再担心忘记了之前是怎么构想的呢，对多人协作的团队来说，这一点也许更为重要。

> 
> 
> 
> 当然使用 jsdoc 等工具也能把对函数的抽象明确下来，不过并没有那么强制，所以效果不一定会很好，不过 jsdoc 反而可以做为 TS 的一种补充。
> 
> 
> 

#### 更自信的写代码 ####

TS 还能让我更自信的写前端代码，这种自信来自 TS 可以帮我们避免很多可能由于自己的忽略造成的 bug。实际上，关于 TS 辅助避免 bug 方面存在专门的研究，一篇名为 [To Type or Not to Type: Quantifying Detectable Bugs in JavaScript]( https://link.juejin.im?target=https%3A%2F%2Flink.zhihu.com%2F%3Ftarget%3Dhttp%253A%2F%2Fttendency.cs.ucl.ac.uk%2Fprojects%2Ftype_study%2Fdocuments%2Ftype_study.pdf ) 的论文，表明使用 TS 进行静态类型检查能帮我们至少减少 15% 以上的 bug （这篇论文的研究过程也很有意思，感兴趣可以点击链接阅读）。

可以举一个例子来说明，TS 是怎么给我带来这种自信的。

下面这条语句，大家都很熟悉，是 DOM 提供依据 id 获取元素的方法。

` const a = document.getElementById( "a" ) 复制代码`

对我自己来说，使用 TS 之前，我忽略了 ` document.getElementById` 的返回值还可能是 ` null` ，这种不经意的忽略也许在未来就会造成一个意想不到的 bug。

使用 TS，在编辑器中就会明确的提醒我们 ` a` 的值可能为 ` null` 。

![](https://user-gold-cdn.xitu.io/2019/5/11/16aa4cc9cce5ca4e?imageView2/0/w/1280/h/960/ignore-error/1)

我们并不一定要处理值 ` null` 的情况，使用 ` const a = document.getElementById('id')!` 可以明确告诉 TS ，它不会是 ` null` ，不过至少，这时候我们清楚的知道自己想做什么。

#### 使用 TS 的过程就是一种学习的过程 ####

使用 TS 后，感觉自己通过浏览器查文档的时间明显少了很多。无论是库还是原生的 js 或者 nodejs，甚至是自己团队其它成员定义的类型。结合 VSCode ，会有非常智能的提醒，也可以很方便看到相应的接口的确切定义。使用的过程就是在加深理解的过程，确实「面向接口编程」天然和静态类型更为亲密。

比如说，我们使用 Color 这个库，VSCode 会有下面这类提醒：

![2019-05-11-002](https://user-gold-cdn.xitu.io/2019/5/11/16aa4cc9c2966a06?imageView2/0/w/1280/h/960/ignore-error/1)

不用去查文档，我们就能看到其提供的 API。 如果我们去看这个库的源文件会发现，能有提醒的原因在于存在下面这样的定义：

` // @types/color/index.d.TS interface Color { toString(): string; toJSON(): Color; string(places?: number): string; percenTString(places?: number): string; array(): number[]; object(): { alpha?: number } & { [key: string]: number }; unitArray(): number[]; unitObject(): { r : number, g : number, b : number, alpha?: number }; ... } 复制代码`

这种提醒无疑能增强开发的效率，虽然定义类型在早期会花费一定的时间，但是对于一个长期维护的比较大型的项目，使用 TS 非常值得。

### 一种学习 typescript 的路径 ###

也许是因为，我之前从未系统的学习过一门静态语言，所以从开始学到感觉自己基本入门了 TS 花的精力还挺多的。 学习 TS 的过程中，主要参考了以下这些资料，你可以直接点击链接查看，也可以继续看后文，我对这些资料有着一些简单的分析。

* [TypeScript handbook — book]( https://link.juejin.im?target=https%3A%2F%2Flink.zhihu.com%2F%3Ftarget%3Dhttps%253A%2F%2Fwww.tslang.cn%2Fdocs%2Fhandbook%2Fbasic-types.html )
* [TypeScript Deep Dive — book]( https://link.juejin.im?target=https%3A%2F%2Flink.zhihu.com%2F%3Ftarget%3Dhttps%253A%2F%2Flegacy.gitbook.com%2Fbook%2Fbasarat%2Ftypescript )
* [TypeScript-React-Starter — github]( https://link.juejin.im?target=https%3A%2F%2Flink.zhihu.com%2F%3Ftarget%3Dhttps%253A%2F%2Fgithub.com%2FMicrosoft%2FTypeScript-React-Starter )
* [react-typescript-cheaTSheet — github]( https://link.juejin.im?target=https%3A%2F%2Flink.zhihu.com%2F%3Ftarget%3Dhttps%253A%2F%2Fgithub.com%2Fsw-yx%2Freact-typescript-cheatsheet )
* [Advanced Static Types in TypeScript — egghead.io]( https://link.juejin.im?target=https%3A%2F%2Flink.zhihu.com%2F%3Ftarget%3Dhttps%253A%2F%2Fegghead.io%2Fcourses%2Fadvanced-static-types-in-typescript )
* [Use TypeScript to develop React Applications — egghead.io]( https://link.juejin.im?target=https%3A%2F%2Flink.zhihu.com%2F%3Ftarget%3Dhttps%253A%2F%2Fegghead.io%2Fcourses%2Fuse-typescript-to-develop-react-applications )
* [Practical Advanced TypeScript — egghead.io]( https://link.juejin.im?target=https%3A%2F%2Flink.zhihu.com%2F%3Ftarget%3Dhttps%253A%2F%2Fegghead.io%2Fcourses%2Fpractical-advanced-typescript )
* [Ultimate React Component Patterns with Typescript 2.8 — medium]( https://link.juejin.im?target=https%3A%2F%2Flink.zhihu.com%2F%3Ftarget%3Dhttps%253A%2F%2Flevelup.gitconnected.com%2Fultimate-react-component-patterns-with-typescript-2-8-82990c516935 )
* [The TypeScript Tax — medium]( https://link.juejin.im?target=https%3A%2F%2Flink.zhihu.com%2F%3Ftarget%3Dhttps%253A%2F%2Fmedium.com%2Fjavascript-scene%2Fthe-typescript-tax-132ff4cb175b )

在阅读上述资料的过程中，我使用 TS 重写了一个基于 CRA 的简单但是很完整的前端项目，现在觉得，使用 TS 来开发工作中的常见需求，应该都能应对了。如果你是刚刚开始学 TS，不妨参照下面的路径学习。

#### 搭建 TS 运行环境 ####

不要误解，并非从零搭建。学习实践性很强的内容时，边学边练习可以帮我们更快的掌握。如果你使用 React，借助 ` yarn` 或者 ` create-react-app` ，可轻易的构造一个基于 TS 的项目。

在命令行中执行下述命令即可生产可直接使用的项目：

` # 使用 yarn $ yarn create react-app TS-react-playground --typescript # 使用 npx $ npx create-react-app TS-react-playground --typescript 复制代码`

随后如果需要，可以在 ` tsconfig.json` 中添加额外的配置。

就我个人而言，我喜欢同步配置 TS-lint 与 prettier，已免去之后练习过程中格式的烦恼。配置方法可以参考 [Configure TypeScript, TSLint, and Prettier in VS Code for React Native Development]( https://link.juejin.im?target=https%3A%2F%2Flink.zhihu.com%2F%3Ftarget%3Dhttps%253A%2F%2Fmedium.com%2F%2540sgroff04%2Fconfigure-typescript-tslint-and-prettier-in-vs-code-for-react-native-development-7f31f0068d2 ) 这篇文章，或者 [看我的配置记录]( https://link.juejin.im?target=https%3A%2F%2Flink.zhihu.com%2F%3Ftarget%3Dhttps%253A%2F%2Fgist.github.com%2FVal-Zhang%2Ff34a318ffd46b037f00b9a29c90a5082 ) 。

如果你不使用 React， [TypeScript 官方文档]( https://link.juejin.im?target=https%3A%2F%2Flink.zhihu.com%2F%3Ftarget%3Dhttps%253A%2F%2Fwww.typescriptlang.org%2Fsamples%2Findex.html ) 首页就提供了 TS 配合其它框架的使用方法。

#### 理解关键的概念 ####

我一直觉得，学习一项新的技能，清楚其边界很重要，相关的细节知识则可以在后续的使用过程中逐步的了解。我们都知道，TS 是 JS 的超集，所以学习 TS 的第一件事情就是要找到「超」的边界在哪里。

这个阶段，推荐阅读 [TypeScript handbook — book]( https://link.juejin.im?target=https%3A%2F%2Flink.zhihu.com%2F%3Ftarget%3Dhttps%253A%2F%2Fwww.tslang.cn%2Fdocs%2Fhandbook%2Fbasic-types.html ) ，这本书其实也是官方推荐的入门手册。这里给的链接是中文翻译版的链接，翻译的质量非常好，虽然内容没有 [英文官方文档]( https://link.juejin.im?target=https%3A%2F%2Flink.zhihu.com%2F%3Ftarget%3Dhttps%253A%2F%2Fwww.typescriptlang.org%2Fdocs%2Fhandbook%2Fbasic-types.html ) 新，不过学习新的东西最好还是从自己最熟悉的内容入手，所以不妨先看中文文档。阅读过程中遇到的示例，都可以在上面搭建的 TS-playground 中练习一下，熟悉一下。

TS 做为 JS 的超集，其「超」其实主要在两方面

* TS 为 JS 引入了一套类型系统；
* TS 支持一些非 ECMAScript 正式标准的语法，比如装饰器；

关于第二点，TS 做的事情有些类似 babel，所以也有人说 TS 是 babel 最大的威胁。不过这些新语法，很可能你早就使用过，本文不再赘述。

比较难理解的其实是这套类型系统，这套类型系统有着自己的 [声明空间（Declaration Spaces）]( https://link.juejin.im?target=https%3A%2F%2Flink.zhihu.com%2F%3Ftarget%3Dhttps%253A%2F%2Fbasarat.gitbooks.io%2Ftypescript%2Fcontent%2Fdocs%2Fproject%2Fdeclarationspaces.html ) ，具有自己的一些关键字和语法。

对我来说，学习 TS 最大的难点就在于这套类型系统中有着一些我之前很少了解的概念，在这里可以大致的梳理一下。

#### 一些 TS 中的新概念 ####

编程实际上就是对数据进行操作和加工的过程。类型系统能辅助我们对数据进行更为准确的操作。TypeScript 的核心就在于其提供一套类型系统，让我们对数据类型有所约束。约束有时候很简单，有时候很抽象。

TS 支持的类型如下： ` boolean` , ` number` , ` string` , ` []` , ` Tuple` , ` enum` , ` any` , ` void` , ` null` , ` undefined` , ` never` , ` Object` 。

TS 中更复杂的数据结构其实都是针对上述类型的组合，关于类型的基础知识，推荐先阅读 [基础类型]( https://link.juejin.im?target=https%3A%2F%2Flink.zhihu.com%2F%3Ftarget%3Dhttps%253A%2F%2Fwww.tslang.cn%2Fdocs%2Fhandbook%2Fbasic-types.html ) 一节，这里只讨论最初对我造成困扰的概念：

* enum： 现在想想 ` enum` 枚举类型非常实用，很多其它的语言都内置了这一类型，合理的使用枚举，能让我们的代码可读性更高，比如：

` const enum MediaTypes { JSON = "application/json" } fetch( "https://swapi.co/api/people/1/" , { headers: { Accept: MediaTypes.JSON } }) .then( ( res ) => res.json()) 复制代码`

* never： ` never` 代表代码永远不会执行到这里，常常可以应用在 ` switch case` 的 ` default` 中，防止我们遗漏 ` case` 未处理，比如：

` enum ShirTSize { XS, S, M, L, XL } function assertNever ( value: never ): never { console.log( Error ( `Unexpected value ' ${value} '` )); } function prettyPrint ( size: ShirTSize ) { switch (size) { case ShirTSize.S: console.log( "small" ); case ShirTSize.M: return "medium" ; case ShirTSize.L: return "large" ; case ShirTSize.XL: return "extra large" ; // case ShirTSize.XS: return "extra small"; default : return assertNever(size); } } 复制代码`

下面是上述代码在我的编辑器中的截图，编辑器会通过报错告知我们还有未处理的情况。

![](https://user-gold-cdn.xitu.io/2019/5/11/16aa4cc9c0a78609?imageView2/0/w/1280/h/960/ignore-error/1)

* 类型断言： 类型断言其实就是你告诉编译器，某个值具备某种类型。有两种不同的方式可以添加类型断言：
* ` <string>someValue`
* ` someValue as string`

关于类型断言，我看文档时的疑惑点在于，我想不到什么情况下会使用它。后来发现，当你知道有这么一个功能，在实际使用过程中，就会发现能用得着，比如说迁移遗留项目时。

* Generics(泛型)： 泛型让我们的数据结构更为抽象可复用，因为这种抽象，也让它有时候不是那么好理解。泛型的应用场景非常广泛，比如：

` type Nullable<T> = { [P in keyof T]: T[P] | null ; }; 复制代码`

能够让某一种接口的子类型都可以为 null。我记得我第一次看到泛型时也觉得它很不好理解，不过后来多用了几次后，就觉得还好了。

* interface 和 type ` interface` 和 ` type` 都可以用来定义一些复杂的类型结构，最很多情况下是通用的，最初我一直没能理解它们二者之间区别在哪里，后来发现，二者的区别在于：
* 

* ` interface` 创建了一种新的类型，而 type 仅仅是别名，是一种引用；
* 如果 type 使用了 union operator （|） 操作符，则不能将 type implements 到 class 上；
* 如果 type 使用了 ` union` （|） 操作符 ，则不能被用以 ` extends` interface
* type 不能像 interface 那样合并，其在作用域内唯一； [ [1]]( https://link.juejin.im?target=https%3A%2F%2Fzhuanlan.zhihu.com%2Fp%2F63346965%23ref_144719_1 )

在视频 [Use Types vs. Interfaces from @volkeron on @eggheadio]( https://link.juejin.im?target=https%3A%2F%2Flink.zhihu.com%2F%3Ftarget%3Dhttps%253A%2F%2Fegghead.io%2Flessons%2Ftypescript-use-types-vs-interfaces ) 中，通过实例对二者的区别有更细致的说明。

> 
> 
> 
> 值得指出的是，TypeScript handbook 关于 type 和 interface 的区别还停留在 TS 2.0
> 版本，对应章节现在的描述并不准确，想要详细了解，可参考 [Interface vs Type alias in TypeScript 2.7](
> https://link.juejin.im?target=https%3A%2F%2Flink.zhihu.com%2F%3Ftarget%3Dhttps%253A%2F%2Fmedium.com%2F%2540martin_hotell%2Finterface-vs-type-alias-in-typescript-2-7-2a8f1777af4c
> ) 这篇文章。
> 
> 

* 类型保护 TS 编译器会分析我们的程序并为某一个变量在指定的作用域来指明尽可能确切的类型，类型保护就是一种辅助确定类型的方法，下面的语句都可以用作类型保护：
* 

* ` typeof padding === "number"`
* ` padder instanceof SpaceRepeatingPadder`

一个应用实例是结合 redux 中的 reducer 中依据不同的 type，TS 能分别出不同作用域内 action 应有的类型。

* 类型映射 类型映射是 TypeScript 提供的从旧类型中创建新类型的一种方式。它们非常实用。比如说，我们想要快速让某个接口中的所有属性变为可选的，可以按照下面这样写：

` interface Person { name: string ; age: number ; } type PartialPerson = { [P in keyof Person]?: Person[P] } 复制代码`

还有一个概念叫做 **映射类型** ，TS 内置一些映射类型（实际上是一些语法糖），让我们可以方便的进行类型映射。比如通过内置的映射类型 Partial ，上面的表达式可以按照下面这样写：

` interface Person { name: string ; age: number ; } type PartialPerson = Partial<Person> 复制代码`

常见的映射类型，可以参看这篇文章 — [TS 一些工具泛型的使用及其实现]( https://link.juejin.im?target=https%3A%2F%2Fzhuanlan.zhihu.com%2Fp%2F40311981 ) ，除了做为语法糖内置在 TS 中的映射类型（如 ` Readonly` ），这篇文章中也提到了一些未内置最 TS 中但是很实用的映射类型（比如 ` Omit` ）。

* 第三方的库，如何得到类型支持 我们很难保证，第三方的库都原生支持 TS 类型，在你使用过一段时间 TS 后，你肯定安装过类似 ` @types/xxx` 的类型库，安装类似这样的库，实际上就安装了某个库的描述文件，对于这些第三方库的类型的定义，都存储在 [DefinitelyTyped]( https://link.juejin.im?target=https%3A%2F%2Flink.zhihu.com%2F%3Ftarget%3Dhttps%253A%2F%2Fgithub.com%2FDefinitelyTyped%2FDefinitelyTyped ) 这个仓库中，常用的第三方库在这里面都有定义了。在 [TypeSearch]( https://link.juejin.im?target=https%3A%2F%2Flink.zhihu.com%2F%3Ftarget%3Dhttps%253A%2F%2Fmicrosoft.github.io%2FTypeSearch%2F ) 中可以搜索第三方库的类型定义包。

关于类型，还有一些很多其它的知识点，不过一些没有那么常用，一些没有那么难理解，在此暂不赘述。

#### 消化学到的新概念 ####

我首次看完《 [TypeScript handbook]( https://link.juejin.im?target=https%3A%2F%2Flink.zhihu.com%2F%3Ftarget%3Dhttps%253A%2F%2Fwww.tslang.cn%2Fdocs%2Fhandbook%2Fbasic-types.html ) 》时，确实觉得自己懂了不少，但是发现动手写代码，还是会经常卡住。追其原因，可能在于一下子接收了太多的新概念，一些概念并没有来得及消化，这时候我推荐看下面这门网课：

* [Advanced Static Types in TypeScript — egghead.io]( https://link.juejin.im?target=https%3A%2F%2Flink.zhihu.com%2F%3Ftarget%3Dhttps%253A%2F%2Fegghead.io%2Fcourses%2Fadvanced-static-types-in-typescript )

看视频算是一种比较轻松的学习方式，这门课时长大概是一个小时。会把 TypeScript handbook 这本书中的一些比较重要的概念，配合实例讲解一次。可以跟着教程把示例敲一次，在 vscode 中多看看给出的提示，看完之后，对 TS 的一些核心概念，肯定会有更深的理解。

#### 模仿和实践 ####

想要真的掌握 TS，少不了实践。模仿也是一种好的实践方式，已 React + TypeScript 为例，比较推荐的模仿内容如下：

* [TypeScript-React-Starter]( https://link.juejin.im?target=https%3A%2F%2Flink.zhihu.com%2F%3Ftarget%3Dhttps%253A%2F%2Fgithub.com%2FMicrosoft%2FTypeScript-React-Starter ) ，这是微软为 TS 初学者提供的一个非常好的资料，可以继续使用我们上面构建的 playground ，参照这个仓库的 readme 写一次，差不多就能知道 TS 结合 React 的基本用法了；
* [GitHub - react-typescript-cheaTSheet]( https://link.juejin.im?target=https%3A%2F%2Flink.zhihu.com%2F%3Ftarget%3Dhttps%253A%2F%2Fgithub.com%2Fsw-yx%2Freact-typescript-cheatsheet ) ，这个教程也比较简单，不过上面那个教程更近了一步，依据其 readme 继续改造我们的 playground 后，我们能知道，React + Redux + TypeScript 该如何配合使用；
* [react-redux-typescript-guide]( https://link.juejin.im?target=https%3A%2F%2Flink.zhihu.com%2F%3Ftarget%3Dhttps%253A%2F%2Fgithub.com%2Fpiotrwitek%2Freact-redux-typescript-guide ) ，这个教程则展示了基于 TypeScript 如何应用一些更复杂的模式，我们也可以模仿其提供的用法，将其应用到我们自己的项目中；
* [Ultimate React Component Patterns with Typescript 2.8]( https://link.juejin.im?target=https%3A%2F%2Flink.zhihu.com%2F%3Ftarget%3Dhttps%253A%2F%2Flevelup.gitconnected.com%2Fultimate-react-component-patterns-with-typescript-2-8-82990c516935 ) ，这篇文章则可以做为上述内容的补充，其在掘金上有 [汉语翻译]( https://link.juejin.im?target=https%3A%2F%2Flink.zhihu.com%2F%3Ftarget%3Dhttps%253A%2F%2Fjuejin.im%2Fpost%2F5b07caf16fb9a07aa83f2977 ) ，点赞量非常高，看完之后，差不多就能了解到如果使用 TS 应对各种 React 组件模式了。
* [Use TypeScript to develop React Applications — egghead.io]( https://link.juejin.im?target=https%3A%2F%2Flink.zhihu.com%2F%3Ftarget%3Dhttps%253A%2F%2Fegghead.io%2Fcourses%2Fuse-typescript-to-develop-react-applications ) ，随后如果想再轻松一点，则可以再看看这个网课，跟着别人的讲解，回头看看自己模仿着写的一些代码，也许会有不同的感触；

至此，你肯定就已经具备了基础的 TS 开发能力，可以独立的结合 TS 开发相对复杂的应用了。

#### 更深的理解 ####

当然也许你并不会满足于会用 TS，你还想知道 TS 的工作原理是什么。这时候推荐阅读下面两篇内容：

* [TypeScript Compiler Internals · TypeScript Deep Dive]( https://link.juejin.im?target=https%3A%2F%2Flink.zhihu.com%2F%3Ftarget%3Dhttps%253A%2F%2Fbasarat.gitbooks.io%2Ftypescript%2Fcontent%2Fdocs%2Fcompiler%2Foverview.html ) ，TS 编译的核心还是 AST，这篇文章讲解了 TS 编译的五个阶段（ Scanner /Parser / Binder /Checker /Emitter ）分别是怎么工作的；
* [Learn how to contribute to the TypeScript compiler on GitHub through a real-world example]( https://link.juejin.im?target=https%3A%2F%2Flink.zhihu.com%2F%3Ftarget%3Dhttps%253A%2F%2Fdev.to%2Fremojansen%2Flearn-how-to-contribute-to-the-typescript-compiler-on-github-through-a-real-world-example-4df0 ) ，则是另外一篇比较好的了解 TS 运行原理的资料。

关于 TS 的原理，我还没有来得及仔细去看。不过 AST 在前端中的应用还真是多，待我补充更多的相关知识后，也许会对 AST 有一个更全面的总结。

TS 当然也不是没有缺点， [The TypeScript Tax]( https://link.juejin.im?target=https%3A%2F%2Flink.zhihu.com%2F%3Ftarget%3Dhttps%253A%2F%2Fmedium.com%2Fjavascript-scene%2Fthe-typescript-tax-132ff4cb175b ) [ [2]]( https://link.juejin.im?target=https%3A%2F%2Fzhuanlan.zhihu.com%2Fp%2F63346965%23ref_144719_2 ) 是一篇非常优秀的文章，阅读这篇文章能让我们更为客观看待 TS，虽然站在作者的角度看，TS 弊大于利，主要原因是 TS 提供的功能大多都可以用其它工具配合在一定程度上代替，而且类型系统会需要写太多额外的代码，类型系统在一定程度上也破坏了动态语言的灵活性，让一些动态语言特有的模式很难在其中被应用。作者最终的结论带有很强的主观色彩，我并不是非常认可，但是这篇文章的分析过程非常精彩，就 TS 的各种特性和现在的 JS 生态进行了对比，能让我们对 TS 有一个更全面的了解，非常推荐阅读，也许你会和我一样，看完这个分析过程，会对 TS 更感兴趣。

TS 每隔几个月就会发布一个新的小版本，每个小版本在 [TypeScript 官方博客]( https://link.juejin.im?target=https%3A%2F%2Flink.zhihu.com%2F%3Ftarget%3Dhttps%253A%2F%2Fdevblogs.microsoft.com%2Ftypescript%2F ) [ [3]]( https://link.juejin.im?target=https%3A%2F%2Fzhuanlan.zhihu.com%2Fp%2F63346965%23ref_144719_3 ) 上都会有专门的说明，可用用作跟进学习 TS 的参考。

### 推荐阅读 ###

下述参考内容在文中，都有链接，如果都看过，则无需再重复查看了。

* [TypeScript handbook — book]( https://link.juejin.im?target=https%3A%2F%2Flink.zhihu.com%2F%3Ftarget%3Dhttps%253A%2F%2Fwww.tslang.cn%2Fdocs%2Fhandbook%2Fbasic-types.html )
* [TypeScript Deep Dive — book]( https://link.juejin.im?target=https%3A%2F%2Flink.zhihu.com%2F%3Ftarget%3Dhttps%253A%2F%2Flegacy.gitbook.com%2Fbook%2Fbasarat%2Ftypescript )
* [TypeScript-React-Starter — github]( https://link.juejin.im?target=https%3A%2F%2Flink.zhihu.com%2F%3Ftarget%3Dhttps%253A%2F%2Fgithub.com%2FMicrosoft%2FTypeScript-React-Starter )
* [react-typescript-cheaTSheet — github]( https://link.juejin.im?target=https%3A%2F%2Flink.zhihu.com%2F%3Ftarget%3Dhttps%253A%2F%2Fgithub.com%2Fsw-yx%2Freact-typescript-cheatsheet )
* [Advanced Static Types in TypeScript — egghead.io]( https://link.juejin.im?target=https%3A%2F%2Flink.zhihu.com%2F%3Ftarget%3Dhttps%253A%2F%2Fegghead.io%2Fcourses%2Fadvanced-static-types-in-typescript )
* [Use TypeScript to develop React Applications — egghead.io]( https://link.juejin.im?target=https%3A%2F%2Flink.zhihu.com%2F%3Ftarget%3Dhttps%253A%2F%2Fegghead.io%2Fcourses%2Fuse-typescript-to-develop-react-applications )
* [Practical Advanced TypeScript — egghead.io]( https://link.juejin.im?target=https%3A%2F%2Flink.zhihu.com%2F%3Ftarget%3Dhttps%253A%2F%2Fegghead.io%2Fcourses%2Fpractical-advanced-typescript )
* [Ultimate React Component Patterns with Typescript 2.8 — medium]( https://link.juejin.im?target=https%3A%2F%2Flink.zhihu.com%2F%3Ftarget%3Dhttps%253A%2F%2Flevelup.gitconnected.com%2Fultimate-react-component-patterns-with-typescript-2-8-82990c516935 )

### 参考 ###

* 关于interface 和 type 二者的区别，handbook 的描述过时了，可参照 [Interface vs Type alias in TypeScript 2.7]( https://link.juejin.im?target=https%3A%2F%2Fmedium.com%2F%40martin_hotell%2Finterface-vs-type-alias-in-typescript-2-7-2a8f1777af4c )
* The TypeScript Tax 对 TS 的优缺点有非常详细的描述，非常推荐阅读 [The TypeScript Tax]( https://link.juejin.im?target=https%3A%2F%2Fmedium.com%2Fjavascript-scene%2Fthe-typescript-tax-132ff4cb175b )
* ts 更新时，官方博客地址可参看这里 [ [devblogs.microsoft.com/typescript/]( https://link.juejin.im?target=https%3A%2F%2Fdevblogs.microsoft.com%2Ftypescript%2F )