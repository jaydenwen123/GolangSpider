# 最全的Eslint配置模板，从此统一团队的编程习惯 #

随着项目的不断增加，急切需要统一每个项目的代码规范，将一些低级错误在萌芽状态下掐死。所以特此结合当前项目使用的一些规范，再根据社区推荐的规范，整合成这个 [repo]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Flinxiaowu66%2Feslint-config-templates ) 。里面集成了React和Nodejs的编程规范，所有的规范都是基于 [airbnb]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fairbnb%2Fjavascript ) ，里面细分了js版本和ts版本，满足大家的编程需求。

另外其他框架的代码规范没有实际项目经验，所以没能集齐所有”龙珠“，因此在此欢迎大家贡献出平时使用的标准编码规范(尽量是基于airbnb的)，共享给社区其他童鞋。

## Eslint生态依赖包介绍 ##

在说明Eslint配置之前，我们先来掌握Eslint配置的生态圈中涉及到的一些依赖包的作用，这样我们方可以知其所以然。

### 最基础 ###

* [eslint]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Feslint%2Feslint ) : lint代码的主要工具，所以的一切都是基于此包

### 解析器(parser) ###

* 

[babel-eslint]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fbabel%2Fbabel-eslint ) : 该依赖包允许你使用一些实验特性的时候，依然能够用上Eslint语法检查。反过来说，当你代码并没有用到Eslint不支持的实验特性的时候是不需要安装此依赖包的。

* 

[@typescript-eslint/parser]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Ftypescript-eslint%2Ftypescript-eslint%2Ftree%2Fmaster%2Fpackages%2Fparser ) : Typescript语法的解析器，类似于 ` babel-eslint` 解析器一样。对应 ` parserOptions` 的配置参考官方的README。

### 扩展的配置 ###

* 

[eslint-config-airbnb]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fairbnb%2Fjavascript%2Ftree%2Fmaster%2Fpackages%2Feslint-config-airbnb ) : 该包提供了所有的Airbnb的ESLint配置，作为一种扩展的共享配置，你是可以修改覆盖掉某些不需要的配置的， **该工具包包含了react的相关Eslint规则(eslint-plugin-react与eslint-plugin-jsx-a11y)，所以安装此依赖包的时候还需要安装刚才提及的两个插件**

* 

[eslint-config-airbnb-base]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fairbnb%2Fjavascript%2Ftree%2Fmaster%2Fpackages%2Feslint-config-airbnb-base ) : 与上一个包的区别是，此依赖包不包含react的规则，一般用于服务端检查。

* 

[eslint-config-jest-enzyme]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2FFormidableLabs%2Fenzyme-matchers%2Ftree%2Fmaster%2Fpackages%2Feslint-config-jest-enzyme ) : jest和enzyme专用的校验规则，保证一些断言语法可以让Eslint识别而不会发出警告。

* 

[eslint-config-prettier]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fprettier%2Feslint-config-prettier ) : 将会禁用掉所有那些非必须或者和 [prettier]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fprettier%2Fprettier ) 冲突的规则。这让您可以使用您最喜欢的shareable配置，而不让它的风格选择在使用Prettier时碍事。请注意该配置 **只是** 将规则 **off** 掉,所以它只有在和别的配置一起使用的时候才有意义。

### 插件 ###

* 

[eslint-plugin-babel]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fbabel%2Feslint-plugin-babel ) : 和babel-eslint一起用的一款插件.babel-eslint在将eslint应用于Babel方面做得很好，但是它不能更改内置规则来支持实验性特性。eslint-plugin-babel重新实现了有问题的规则，因此就不会误报一些错误信息

* 

[eslint-plugin-import]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fbenmosher%2Feslint-plugin-import ) : 该插件想要支持对ES2015+ (ES6+) import/export语法的校验, 并防止一些文件路径拼错或者是导入名称错误的情况

* 

[eslint-plugin-jsx-a11y]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fevcohen%2Feslint-plugin-jsx-a11y ) : 该依赖包专注于检查JSX元素的可访问性。

* 

[eslint-import-resolver-webpack]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fbenmosher%2Feslint-plugin-import%23resolvers ) : 可以借助webpack的配置来辅助eslint解析，最有用的就是alias，从而避免unresolved的错误

* 

[eslint-import-resolver-typescript]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Falexgorbatchev%2Feslint-import-resolver-typescript ) ：和eslint-import-resolver-webpack类似，主要是为了解决alias的问题

* 

[eslint-plugin-react]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fyannickcr%2Feslint-plugin-react ) : React专用的校验规则插件.

* 

[eslint-plugin-jest]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fjest-community%2Feslint-plugin-jest ) : Jest专用的Eslint规则校验插件.

* 

[eslint-plugin-prettier]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fprettier%2Feslint-plugin-prettier ) : 该插件辅助Eslint可以平滑地与Prettier一起协作，并将Prettier的解析作为Eslint的一部分，在最后的输出可以给出修改意见。这样当Prettier格式化代码的时候，依然能够遵循我们的Eslint规则。如果你禁用掉了所有和代码格式化相关的Eslint规则的话，该插件可以更好得工作。所以你可以使用eslint-config-prettier禁用掉所有的格式化相关的规则(如果其他有效的Eslint规则与prettier在代码如何格式化的问题上不一致的时候，报错是在所难免的了)

* 

[@typescript-eslint/eslint-plugin]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Ftypescript-eslint%2Ftypescript-eslint%2Ftree%2Fmaster%2Fpackages%2Feslint-plugin ) ：Typescript辅助Eslint的插件。

* 

：promise规范写法检查插件，附带了一些校验规则。

### 辅助优化流程 ###

* 

[husky]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Ftypicode%2Fhusky ) : git命令hook专用配置.

* 

[lint-staged]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fokonet%2Flint-staged ) : 可以定制在特定的git阶段执行特定的命令。

### Prettier ###

Prettier相关的包有好多个，除了上面列举的两个，你可能还会用到下面的三个：

* 

[prettier]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fprettier%2Fprettier ) ：原始实现版本，定义了prettier规则并实现这些规则。支持的规则参考： [传送门]( https://link.juejin.im?target=https%3A%2F%2Fprettier.io%2Fdocs%2Fen%2Foptions.html )

* 

[prettier-eslint]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fprettier%2Fprettier-eslint ) ：输入代码，执行prettier后再eslint --fix输出格式化后的代码。仅支持字符串输入。

* 

[prettier-eslint-cli]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fprettier%2Fprettier-eslint-cli ) ：顾名思义，支持CLI命令执行 [prettier-eslint]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fprettier%2Fprettier-eslint ) 的操作

那么Prettier这么多工具包，都是些什么关系呢？太容易让人混淆了。这里用一段话简单介绍一下：

**最基础的是prettier，然后你需要用eslint-config-prettier去禁用掉所有和prettier冲突的规则，这样才可以使用eslint-plugin-prettier去以符合eslint规则的方式格式化代码并提示对应的修改建议。为了让prettier和eslint结合起来，所以就诞生了prettier-eslint这个工具，但是它只支持输入代码字符串，不支持读取文件，因此又有了prettier-eslint-cli**

这就是5个工具包互相之间的关系。加上prettier之后的提示可读性高一点，如下图：

![](https://user-gold-cdn.xitu.io/2019/6/4/16b207285676b803?imageView2/0/w/1280/h/960/ignore-error/1)

## Eslint配置文件 ##

* 

env: 预定义那些环境需要用到的全局变量，可用的参数是： ` es6` 、 ` broswer` 、 ` node` 等。

` es6` 会使能所有的ECMAScript6的特性除了模块(这个功能在设置ecmaVersion版本为6的时候会自动设置)

` browser` 会添加所有的浏览器变量比如Windows

` node` 会添加所有的全局变量比如 ` global`

更多环境配置参考 [Specifying Environments]( https://link.juejin.im?target=https%3A%2F%2Feslint.org%2Fdocs%2Fuser-guide%2Fconfiguring%23specifying-environments )

* 

extends: 指定扩展的配置，配置支持递归扩展，支持规则的覆盖和聚合。

* 

plugins: 配置那些我们想要Linting规则的插件。

* 

parser: 默认ESlint使用 [Espree]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Feslint%2Fespree ) 作为解析器，但是一旦我们使用babel的话，我们需要用babel-eslint。

* 

parserOptions: 当我们将默认的解析器从Espree改为babel-eslint的时候，我们需要指定parseOptions，这个是必须的。

ecmaVersion: 默认值是5，可以设置为3、5、6、7、8、9、10，用来指定使用哪一个ECMAScript版本的语法。也可以设置基于年份的JS标准，比如2015(ECMA 6)

sourceType: 如果你的代码是ECMAScript 模块写的，该字段配置为 ` module` ，否则为 ` script` (默认值)

ecmaFeatures：该对象指示你想使用的额外的语言特性

` globalReturn：允许全局范围内的`return`语句 impliedStrict：使能全局`strict`模式 jsx：使能JSX 复制代码`
* 

rules: 自定义规则，可以覆盖掉extends的配置。

* 

settings：该字段定义的数据可以在所有的插件中共享。这样每条规则执行的时候都可以访问这里面定义的数据

更多配置选项参考官方文档 [Eslint]( https://link.juejin.im?target=https%3A%2F%2Feslint.org%2Fdocs%2Fuser-guide%2Fconfiguring )

### Eslint配置文件解析 ###

介绍了这么多，我们以模板提供的一个配置例子 来说(说明都写在注释里了~)：

` module.exports = { parser : 'babel-eslint' , // Specifies the ESLint parser parserOptions: { ecmaVersion : 2015 , // specify the version of ECMAScript syntax you want to use: 2015 => (ES6) sourceType: 'module' , // Allows for the use of imports ecmaFeatures: { jsx : true , // enable JSX impliedStrict: true // enable global strict mode } }, extends : [ 'airbnb' , // Uses airbnb, it including the react rule(eslint-plugin-react/eslint-plugin-jsx-a11y) 'plugin:promise/recommended' , // 'prettier', // Use prettier, it can disable all rules which conflict with prettier // 'prettier/react' // Use prettier/react to pretty react syntax ], settings : { 'import/resolver' : { // This config is used by eslint-import-resolver-webpack webpack: { config : './webpack/webpack-common-config.js' } }, }, env : { browser : true // enable all browser global variables }, 'plugins' : [ 'react-hooks' , 'promise' ], // ['prettier', 'react-hooks'] rules: { // Place to specify ESLint rules. Can be used to overwrite rules specified from the extended configs // e.g. '@typescript-eslint/explicit-function-return-type': 'off', "react-hooks/rules-of-hooks" : "error" , "semi" : [ "error" , "never" ], "react/jsx-one-expression-per-line" : 0 , /** * @description rules of eslint-plugin-prettier */ // 'prettier/prettier': [ // 'error', { // 'singleQuote': true, // 'semi': false // } // ] }, }; 复制代码`

因为我们在代码的保存以及提交阶段都会进行prettier的格式化，所以在Eslint中禁用掉了所有跟prettier的配置，如果你需要的话，可以重新enable掉。

下面对比一下打开prettier和没有打开prettier的区别：

![](https://user-gold-cdn.xitu.io/2019/6/4/16b2072859c52520?imageView2/0/w/1280/h/960/ignore-error/1)

和

![](https://user-gold-cdn.xitu.io/2019/6/4/16b207285a1f1608?imageView2/0/w/1280/h/960/ignore-error/1)

## Prettier影响的规则 ##

上面prettier介绍了那么多，那么这小节简单介绍以下prettier的一些重要的的规则：

* printWidth: 确保你的代码每行的长度不会超过100个字符
* singleQuote: 转换所有的双引号为单引号
* trailingComma: 确保对象的最后一个属性后会有一个逗号
* bracketSpacing: 自动为对象字面量之间添加空格，比如：{ foo: bar }
* jsxBracketSameLine: 在多行JSX元素的最后一行追加 >
* tabWidth：指定tab的宽度是几个空格
* semi: 是否在每行声明代码之后都要添加;

更多规则请参考： [Options]( https://link.juejin.im?target=https%3A%2F%2Fprettier.io%2Fdocs%2Fen%2Foptions.html )

## 让优美的代码深入到骨髓里~ ##

### 保存代码的时候自动格式化(Vscode版本) ###

* 

安装Eslint插件

* 

Vscode配置：

2.1. ` editor.formatOnSave` 置为false，以防默认的文件格式化配置和Eslint和Prettier冲突

2.2. ` eslint.autoFixOnSave` 置为true，这样当我们每次保存文件的时候就可以自动fix文件的错误格式。

如下图：

![](https://user-gold-cdn.xitu.io/2019/6/4/16b2075a90ac5765?imageView2/0/w/1280/h/960/ignore-error/1)

![](https://user-gold-cdn.xitu.io/2019/6/4/16b207621a52594d?imageView2/0/w/1280/h/960/ignore-error/1)

### Lint-staged ###

[Lint-staged]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fokonet%2Flint-staged ) 帮助你在暂存文件的时候能够让错误格式代码不会提交到你分支。

#### 为什么使用Lint-staged? ####

因为提交代码前的检查是最后一个管控代码质量的一个环节，所以在提交代码之前进行lint检查意义重大。这样可以确保没有错误的语法和代码样式被提交到仓库上。但是在整个项目上执行Lint进程会很低效，所以最好的做法就是检查那个被改动的文件。而Lint-staged就是做这个的。

根据上面我们提供的生态圈依赖包，在 ` package.json` 中配置该字段：

` "lint-staged": { "**/*.{tsx,ts}": [ // 这里的文件后缀可以修改成自己需要的文件后缀 "prettier-eslint --write", "git add" ] } 复制代码`

#### 与Husky结合使用 ####

为了让lint-staged可以在change被staged之前执行，我们这时候需要借助git的钩子功能，而提供钩子功能的社区解决方案就是 [husky]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Ftypicode%2Fhusky ) ，该工具提供了git在多个阶段前执行的操作，比如我们这次要在预提交的时候进行Lint检查，配置如下：

` "husky": { "hooks": { "pre-commit": "lint-staged" } } 复制代码`

这样每次commit的时候会执行lint操作，如之前所说，prettier-eslint-cli会将代码prettier一遍后再eslint --fix，如果没有错误，那么就会直接执行 ` git add` ，否则报错退出。

### EditorConfig ###

因为并不是所有的人都使用VS code，所以为了在同样的项目下保持一致，比如tab的宽度或者行尾要不要有分号，我们可以使用.editorconfig来统一这些配置。

支持EditorConfig的编辑器列表请走 [这里]( https://link.juejin.im?target=https%3A%2F%2Feditorconfig.org%2F ) 。

下面是模板配置里面推荐的editorconfig的配置

` # EditorConfig is awesome: http://EditorConfig.org # top-most EditorConfig file root = true [*.md] trim_trailing_whitespace = false [*.js] trim_trailing_whitespace = true # Unix-style newlines with a newline ending every file [*] indent_style = space indent_size = 2 # 保证在任何操作系统上都有统一的行尾结束字符 end_of_line = lf charset = utf-8 insert_final_newline = true max_line_length = 100 复制代码`

## 最后 ##

至此Eslint的全解析就完美落幕了，最后再安利一波 [eslint-config-templates]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Flinxiaowu66%2Feslint-config-templates ) ，也欢迎大家PR。

## 参考 ##

* [These tools will help you write clean code]( https://link.juejin.im?target=https%3A%2F%2Fwww.freecodecamp.org%2Fnews%2Fthese-tools-will-help-you-write-clean-code-da4b5401f68e%2F )
* [Prettier]( https://link.juejin.im?target=https%3A%2F%2Fprettier.io%2F )
* [Eslint]( https://link.juejin.im?target=https%3A%2F%2Feslint.org%2Fdocs%2Fuser-guide%2Fconfiguring )