# 【第三期】使用lerna管理常用工具库 #

![](https://user-gold-cdn.xitu.io/2019/6/5/16b2663ee842e719?imageView2/0/w/1280/h/960/ignore-error/1)

在工作中我们有时会写一些常用的库，比如包含数据类型判断、 ` cookie` 存储模块的工具库等，但可能在某些业务场景中，并不需要用到所有的模块。

我们通常会将这个库拆分成多个，分别创建 ` git` 仓库，分别打包上传到 ` npm` ，这样做看起来并没有什么问题。

但当多个库之间产生依赖的时候，问题就就会显露出来；你需要打包发布修改后的库，还需要修改所有依赖库的版本号，重新发包。

可想而知，当库多起来后，这个过程将会变得多么繁琐。

那么有什么好的方式来解决呢？ [Lerna]( https://link.juejin.im?target=https%3A%2F%2Fwww.npmjs.com%2Fpackage%2Flerna ) 正适合这样的应用场景。

## lerna是什么 ##

Lerna是一个用于管理具有多个包的JavaScript项目的工具，它采用 ` monorepo` (单代码仓库)的管理方式。

将所有相关 ` module` 都放到一个 ` repo` 里，每个 ` module` 独立发布，（例如 [Babel]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fbabel%2Fbabel%2Ftree%2Fv7.0.0-beta.37%2Fpackages ) 、 [React]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Ffacebook%2Freact%2Ftree%2Fv16.2.0%2Fpackages ) 和 [jest]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Ffacebook%2Fjest%2Ftree%2Fmaster%2Fpackages ) 等），issue和PR都集中到该repo中。

你不需要手动去维护每个包的依赖关系，当发布时，会自动更新相关包的版本号，并自动发布。

Lerna项目文件结构：

` ├── lerna.json ├── package.json └── packages ├── package -a │   ├── index.js │   └── package.json └── package-b ├── index.js └── package.json 复制代码`

## lerna主要做了什么 ##

* 通过 ` lerna bootstrap` 命令安装依赖并将代码库进行 ` npm link` 。
* 通过 ` lerna publish` 发布最新改动的库

## 如何使用 ##

### 安装 ###

` npm install --global lerna #or yarn global add lerna 复制代码`

### 初始化一个项目 ###

` mkdir demo cd demo lerna init 复制代码`

执行后将生成以下目录：

` ├── lerna.json # lerna配置文件 ├── package.json └── packages # 包存放文件夹 复制代码`

Lerna有两种管理项目的模式：固定模式或独立模式

#### 固定模式 ####

固定模式是默认的模式，版本号使用 ` lerna.json` 文件中的 ` version` 属性。执行 ` lerna publish` 时，如果代码有更新，会自动更新此版本号的值。

#### 独立模式 ####

独立模式，允许维护人员独立的增加修改每个包的版本，每次发布，所有更改的包都会提示输入指定版本号。

使用方式：

` lerna init --independent 复制代码`
> 
> 
> 
> 修改 ` lerna.json` 中的 ` version` 值为 ` independent` ，可将固定模式改为独立模式运行。
> 
> 

### lerna配置解析 ###

` { "npmClient": "yarn", // 执行命令所用的客户端，默认为npm "command": { // 命令相关配置 "publish": { // 发布时配置 "ignoreChanges": ["ignored-file", "*.md"], // 发布时忽略的文件 "message": "chore(release): publish" // 发布时的自定义提示消息 }, "bootstrap": { // 安装依赖配置 "ignore": "component-*", // 忽略项 "npmClientArgs": ["--no-package-lock"] // 执行 lerna bootstrap命令时传的参数 } }, "packages": [ // 指定存放包的位置 "packages/*" ], "version": "0.0.0" // 当前版本号 } 复制代码`

### 共用 ` devDependencies` ###

开发过程中，很多模块都会依赖 ` babel` 、 ` eslint` 等模块，这些大多都是可以共用的，

我们可以通过 ` lerna link convert` 命令，将它们自动放到根目录的 ` package.json` 文件中去。

这样做即可以保证每个依赖的版本统一，也可以减少存储空间，减少依赖安装的速度。

**注意：** 一些 ` npm` 可执行的包，仍然需要安装到使用模块的包中，才能正常执行，例如 ` jest` 。

### 使用yarn Workspaces ###

工作区是设置软件包体系结构的一种新方式，只需要运行一次 ` yarn install` 便可将指定工作区中所有依赖包全部安装。

#### 优势 ####

* 依赖包可以链接在一起，这意味着你的工作区可以相互依赖，同时始终使用最新的可用代码。 这也是一个比 ` yarn link` 更好的机制，因为它只影响你工作区的依赖树，而不会影响整个系统。
* 所有的项目依赖将被安装在一起，这样可以让 Yarn 来更好地优化它们。
* Yarn 将使用一个单一的 lock 文件，而不是每个包都有一个，这意味着拥有更少的冲突和更容易的进行代码检查。

#### 如何使用 ####

在 ` package.json` 文件中添加以下内容：

**package.json**

` { "private" : true , "workspaces" : [ "packages/*" ] } 复制代码`

**注意:** ` private: true` 是必需的！工作区本身不应当被发布出去，所以我们添加了这个安全措施以确保它不会被意外暴露。

#### lerna中使用 ####

需要在 ` lerna.json` 文件中增加以下配置来启用yarn workspaces：

` { "useWorkspaces" : true } 复制代码`

### 创建模块 ###

` lerna create package -a 复制代码`

执行上面的命令，会在 ` package` 文件夹下创建模块，并根据交互提示生成对应的 ` package.json` 。

生成目录结构如下：

` ├── lerna.json ├── package.json └── packages └── package -a ├── __tests__ │ └── name.test.js ├── lib │ └── name.js ├── package.json └── README.md 复制代码`

### 添加依赖 ###

将模块 ` package-a` 添加到 ` package-b` 模块依赖中

` larna add package -a --scope=package-b 复制代码`

添加完成后会在 ` package-b` 的 ` package.json` 中增加以下依赖项

` { "dependencies" : { "package-a" : "file:../package-a" } } 复制代码`

包依赖使用 ` file:` 来指定本地路径文件

## 发布 ##

发布时，需要先提交 ` commit` 代码，然后执行 ` lerna publish` 命令，提示选择版本号：

![](https://user-gold-cdn.xitu.io/2019/6/5/16b2658c168474bf?imageView2/0/w/1280/h/960/ignore-error/1)

这里选择 ` Patch` ，然后会提示，哪些包会升级到 ` 1.0.1` 。

接着根据提示选择确认即可发布成功。

也可以使用 ` lerna publish -y` 默认选项全部选择 ` Yes` ，并根据 ` commit` 信息自动升级版本号。

## Lerna Changelog ##

` lerna` 自带生成 ` Changelog` 的功能，只需要通过简单的配置就可以生成 ` CHANGELOG.md` 文件。

配置如下：

` { "command": { "publish": { "allowBranch": "master", // 只在master分支执行publish "conventionalCommits": true, // 生成changelog文件 "exact": true // 准确的依赖项 } } } 复制代码`

配置后，当我们执行 ` lerna publish` 后会在项目根目录以及每个 ` packages` 包下，生成 ` CHANGELOG.md` 。

**注意:** 只有符合 [约定]( https://link.juejin.im?target=https%3A%2F%2Fwww.conventionalcommits.org%2Fzh%2Fv1.0.0-beta.3%2F ) 的 ` commit` 提交才能正确生成 ` CHANGELOG.md` 文件。

如果提交的 ` commit` 为 ` fix` 会自动升级版本的 **修订号** ;

如果为 ` feat` 则自动更新 **次版本号** ;

如果有破坏性的更改，则会修改 **主版本号** 。

## Lerna与Jest集成 ##

在包发布之前，为了保证代码的质量，都需要来编写单元测试，为了提高效率并方便测试运行，我们想要做到以下功能：

* 所有包只维护一份公共的jest配置文件
* 可以整体运行所有单元测试
* 可以只对某个包执行单元测试

### jest配置 ###

在项目根目录配置 ` jest.config.js` 文件如下：

` const path = require ( 'path' ) module.exports = { collectCoverage : true , // 收集测试时的覆盖率信息 coverageDirectory: path.resolve(__dirname, './coverage' ), // 指定输出覆盖信息文件的目录 collectCoverageFrom: [ // 指定收集覆盖率的目录文件，只收集每个包的lib目录，不收集打包后的dist目录 '**/lib/**' , '!**/dist/**' ], testURL : 'https://www.shuidichou.com/jd' , // 设置jsdom环境的URL testMatch: [ // 测试文件匹配规则 '**/__tests__/**/*.test.js' ], testPathIgnorePatterns : [ // 忽略测试路径 '/node_modules/' ], coverageThreshold : { // 配置测试最低阈值 global: { branches : 100 , functions : 100 , lines : 100 , statements : 100 } } } 复制代码`

### 编写测试脚本 ###

新增 ` scripts` 文件夹，添加 ` test.js` 文件:

` const minimist = require ( 'minimist' ) const rawArgs = process.argv.slice( 2 ) const args = minimist(rawArgs) const path = require ( 'path' ) let rootDir = path.resolve(__dirname, '../' ) // 指定包测试 if (args.p) { rootDir = rootDir + '/packages/' + args.p } const jestArgs = [ '--runInBand' , '--rootDir' , rootDir ] console.log( `\n===> running: jest ${jestArgs.join( ' ' )} ` ) require ( 'jest' ).run(jestArgs) 复制代码`

该脚本通过解析命令行参数 ` -p` 来决定执行指定包的测试用例，如果没有指定 ` -p` 参数，则执行全部测试用例。

修改根目录下 ` package.json` 的 ` script` 增加如下命令：

` { "scripts" : { "ut" : "node scripts/test.js" } } 复制代码`

运行测试脚本：

` # 执行全部测试 yarn ut # 执行某个包测试 yarn ut -p package -a 复制代码`

## Lerna与webpack集成 ##

发包时，会需要使用 ` webpack` 进行 ` es6` 转码或压缩打包，如果每个都维护一份配置文件，就会很繁琐；我们有与 ` jest` 相同的需求：

* 只用一份 ` webpack` 配置文件
* 可以一次性将所有的模块分别打包
* 也可以单独对指定模块打包

### webpack配置 ###

在根目录创建 ` webpack.config.js` 文件，如下：

` var path = require ( 'path' ) const CleanWebpackPlugin = require ( 'clean-webpack-plugin' ) module.exports = ( opt ) => { return { mode : 'production' , entry : path.resolve(opt.path, './lib/index.js' ), output : { path : path.resolve(opt.path, './dist' ), filename : ` ${opt.name}.min.js` , library : opt.name, libraryTarget : 'umd' , umdNamedDefine : true }, externals : opt.externals, plugins : [ new CleanWebpackPlugin() ], module : { rules : [ { test : /\.js$/ , loader : 'babel-loader' , include : [path.resolve(opt.path, './lib' )], options : { // 指定babel配置文件 configFile: path.resolve(__dirname, '.babelrc' ) } } ] }, optimization : { minimize : true } } } 复制代码`

这个配置文件是一个函数，通过接受一个参数对象，来返回最终的配置内容，

### 编写build脚本 ###

思路：

* 读取 ` packages` 目录下的所有模块，获取模块的路径
* 读取模块下的 ` package.json` ，获取 ` name` 及依赖项
* 通过模块路径、包名和 ` package.json` 中的 ` dependencies` 参数来获取 ` webpack` 配置
* 通过 ` webpack` 的 ` Node API` 执行配置编译打包
* 根据命令行参数，判断执行需要打包的配置文件（单独打包）

具体实现如下：

` /scripts/build.js`

` const minimist = require ( 'minimist' ) const rawArgs = process.argv.slice( 2 ) const args = minimist(rawArgs) const webpack = require ( 'webpack' ) const webpackConfig = require ( '../webpack.config' ) const fs = require ( 'fs' ) const path = require ( 'path' ) const packages = fs.readdirSync(path.resolve(__dirname, '../packages/' )) // 获取外部依赖配置 function getExternals ( dependencies ) { let externals = {} if (dependencies) { Object.keys(dependencies).forEach( p => { externals[p] = `commonjs ${p} ` }) return externals } } const packageWebpackConfig = {} // 遍历所有的包生成配置参数 packages.forEach( item => { let packagePath = path.resolve(__dirname, '../packages/' , item) const { name, dependencies } = require (path.resolve(packagePath, 'package.json' )) packageWebpackConfig[item] = { path : packagePath, name, externals : getExternals(dependencies) } }) function build ( configs ) { // 遍历执行配置项 configs.forEach( config => { webpack(webpackConfig(config), (err, stats) => { if (err) { console.error(err) return } console.log(stats.toString({ chunks : false , // 使构建过程更静默无输出 colors: true // 在控制台展示颜色 })) if (stats.hasErrors()) { return } console.log( ` ${config.name} build successed!` ) }) }) } console.log( '\n===> running build' ) // 根据 -p 参数获取执行对应的webpack配置项 if (args.p) { if (packageWebpackConfig[args.p]) { build([packageWebpackConfig[args.p]]) } else { console.error( ` ${args.p} package is not find!` ) } } else { // 执行所有配置 build( Object.values(packageWebpackConfig)) } 复制代码`

然后在根目录下 ` package.json` 的 ` script` 增加如下命令：

` { "scripts" : { "build" : "node scripts/build.js" } } 复制代码`

运行构建脚本：

` # 全部打包 yarn build # 指定打包 yarn build -p package -a 复制代码`

至此， ` Lerna` 的使用方法就介绍完成了。

水滴前端团队招募伙伴，欢迎投递简历到邮箱：fed@shuidihuzhu.com

![](https://user-gold-cdn.xitu.io/2019/5/30/16b06976f9a2f135?imageView2/0/w/1280/h/960/ignore-error/1)