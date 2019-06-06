# 实现一个简单的Webpack #

大家好，我是神三元，今天通过一道面试题来和大家聊一聊webpack。

### 一、什么是Webpack ###

我相信，尽管很多开发者会根据官方文档进行webpack的相关配置，但仍然并不了解Webpack究竟是起什么作用的，在前端工程化扮演者什么角色，观念仍然简单地停留在“代码打包工具”上。真的是这样吗？ 让我们来看看官方定义:

> 
> 
> 
> 本质上，webpack 是一个现代 JavaScript 应用程序的静态模块打包器(module bundler)。当 webpack
> 处理应用程序时，它会递归地构建一个依赖关系图(dependency graph)，其中包含应用程序需要的每个模块，然后将所有这些模块打包成一个或多个
> bundle。
> 
> 

相信这个定义已经说的非常清楚了。首先，它的本质是一个模块打包器，其工作是将每个模块打包成相应的bundle。那么在这中间究竟做了什么事情呢？

### 二、场景引入 ###

假如你是一个面试者，请看题：

在src目录下有以下文件：

` //word.js export const word = 'hello' 复制代码` ` //message.js import {word} from './word.js' ; const message = `say ${word} ` export default message; 复制代码` ` //index.js import message from './message.js' console.log(message) 复制代码`

请编写一个bundler.js,将其中的ES6代码转换为ES5代码，并将这些文件打包，生成一段能在浏览器正确运行起来的代码。(最后输出say hello)

如果你真正理解了Webpack的定义，那么这里思路应该非常清晰：

* 1、利用babel完成代码转换,并生成单个文件的依赖
* 2、生成依赖图谱
* 3、生成最后打包代码

接下来，让我们来一步步解开bundler的面纱。

### 三、步步为营 ###

#### 第一步:转换代码、 生成依赖 ####

转换代码需要利用@babel/parser生成AST抽象语法树，然后利用@babel/traverse进行AST遍历，记录依赖关系，最后通过@babel/core和@babel/preset-env进行代码的转换

` //先安装好相应的包 npm install @babel/parser @babel/traverse @babel/core @babel/preset-env -D 复制代码` ` //导入包 const fs = require ( 'fs' ) const path = require ( 'path' ) const parser = require ( '@babel/parser' ) const traverse = require ( '@babel/traverse' ).default const babel = require ( '@babel/core' ) 复制代码` ` function stepOne ( filename ) { //读入文件 const content = fs.readFileSync(filename, 'utf-8' ) const ast = parser.parse(content, { sourceType : 'module' //babel官方规定必须加这个参数，不然无法识别ES Module }) const dependencies = {} //遍历AST抽象语法树 traverse(ast, { //获取通过import引入的模块 ImportDeclaration({node}){ const dirname = path.dirname(filename) const newFile = './' + path.join(dirname, node.source.value) //保存所依赖的模块 dependencies[node.source.value] = newFile } }) //通过@babel/core和@babel/preset-env进行代码的转换 const {code} = babel.transformFromAst(ast, null , { presets : [ "@babel/preset-env" ] }) return { filename, //该文件名 dependencies, //该文件所依赖的模块集合(键值对存储) code //转换后的代码 } } 复制代码`

#### 第二步：生成依赖图谱。 ####

` //entry为入口文件 function stepTwo ( entry ) { const entryModule = stepOne(entry) //这个数组是核心，虽然现在只有一个元素，往后看你就会明白 const graphArray = [entryModule] for ( let i = 0 ; i < graphArray.length; i++){ const item = graphArray[i]; const {dependencies} = item; //拿到文件所依赖的模块集合(键值对存储) for ( let j in dependencies){ graphArray.push( stepOne(dependencies[j]) ) //敲黑板！关键代码，目的是将入口模块及其所有相关的模块放入数组 } } //接下来生成图谱 const graph = {} graphArray.forEach( item => { graph[item.filename] = { dependencies : item.dependencies, code : item.code } }) return graph } 复制代码` ` //测试一下 console.log(stepTwo( './src/index.js' )) //结果如下，是不是很神奇鸭 { './src/index.js' : { dependencies : { './message.js' : './src\\message.js' }, code : '"use strict";\n\nvar _message = _interopRequireDefault(require("./message.js"));\n\nfunction _interopRequireDefault(obj) { return obj && obj.__esModule ? obj : { "default": obj }; }\n\nconsole.log(_message["default"]);' }, './src\\message.js' : { dependencies : { './word.js' : './src\\word.js' }, code : '"use strict";\n\nObject.defineProperty(exports, "__esModule", {\n value: true\n});\nexports["default"] = void 0;\n\nvar _word = require("./word.js");\n\nvar message = "say ".concat(_word.word);\nvar _default = message;\nexports["default"] = _default;' }, './src\\word.js' : { dependencies : {}, code : '"use strict";\n\nObject.defineProperty(exports, "__esModule", {\n value: true\n});\nexports.word = void 0;\nvar word = \'hello\';\nexports.word = word;' } } 复制代码`

#### 第三步: 生成代码字符串 ####

` //下面是生成代码字符串的操作，仔细看，不要眨眼睛哦！ function step3 ( entry ) { //要先把对象转换为字符串，不然在下面的模板字符串中会默认调取对象的toString方法，参数变成[Object object],显然不行 const graph = JSON.stringify(stepTwo(entry)) return ` (function(graph) { //require函数的本质是执行一个模块的代码，然后将相应变量挂载到exports对象上 function require(module) { //localRequire的本质是拿到依赖包的exports变量 function localRequire(relativePath) { return require(graph[module].dependencies[relativePath]); } var exports = {}; (function(require, exports, code) { eval(code); })(localRequire, exports, graph[module].code); return exports;//函数返回指向局部变量，形成闭包，exports变量在函数执行后不会被摧毁 } require(' ${entry} ') })( ${graph} )` } 复制代码` ` //最终测试 const code = step3( './src/index.js' ) console.log(code) 复制代码`

将生成的这段代码字符串放在浏览器端执行，

![](https://user-gold-cdn.xitu.io/2019/6/1/16b12e88c959607a?imageView2/0/w/1280/h/960/ignore-error/1)

大功告成！其实你再新建一个dist目录，将这些字符串放在main.js文件里，是不是跟你平日里开发npm run build一样的效果呢？

那这个时候就有人要"发炎"了，说你这题目不是让人手写一个Webpack吗？确实，但是真正意义上的Webpack需要考虑非常多的因素，事实上要庞大很多，不过通过这一波实践你应该更加理解了Webpack所做的事情，对Webpack有了一个清晰的认知，这样我的目的也就达到了。中间会有一部分代码比较绕，但没关系，相信你很快就能啃下来，一定会收获满满。

我是神三元，希望这篇文章能够帮助到更多的同学。加油吧！