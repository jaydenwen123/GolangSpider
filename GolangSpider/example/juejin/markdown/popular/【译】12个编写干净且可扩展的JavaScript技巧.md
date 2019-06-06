# 【译】12个编写干净且可扩展的JavaScript技巧 #

JavaScript起源于早期的网络。 从作为脚本语言开始，到现在它已经发展成为一种完全成熟的编程语言，并且支持服务器端执行。

现代Web应用程序严重依赖JavaScript，尤其是单页应用程序（SPA）。借助于React，AngularJS和Vue.js等新兴框架，Web应用程序主要使用JavaScript构建。

扩展这些应用程序有时候会比较棘手，通过简单的设置，您最终可能会遇到限制并迷失在混乱的海洋中。我想分享一些小技巧，这些技巧将帮助您以有效的方式编写干净的代码。

本文面向任何技能水平的JavaScript开发人员。 但是，至少具有JavaScript中级知识的开发人员将从这些技巧中获益最多。

> 
> 
> 
> 原文链接： [blog.logrocket.com/12-tips-for…](
> https://link.juejin.im?target=https%3A%2F%2Fblog.logrocket.com%2F12-tips-for-writing-clean-and-scalable-javascript-3ffe30abfe20%2F
> )
> 
> 

### 分隔您的代码 ###

我建议保持代码库清洁和可读的最重要的事情是具有按主题分隔的特定逻辑块（通常是函数）。如果你编写一个函数，该函数应该默认只有一个目的，不应该一次做多个事情。

此外，您应避免引起副作用，这意味着在大多数情况下，您不应更改在函数外声明的任何内容。 您将数据接收到带参数的函数中；其他一切都不应该被访问。如果您希望从函数中获取某些内容，请返回新值。

### 模块化 ###

当然，如果以类似的方式使用这些函数或执行类似的操作，您可以将多个函数分组到一个模块（and/or 的类中）。例如，如果要进行许多不同的计算，请将它们拆分为可以链接的独立步骤（函数）。但是，这些函数都可以在一个文件（模块）中声明。 以下是JavaScript中的示例：

` function add ( a, b ) { return a + b } function subtract ( a, b ) { return a - b } module.exports = { add, subtract } const { add, subtract } = require ( './calculations' ) console.log(subtract( 5 , add( 3 , 2 )) 复制代码`

如果您正在编写前端JavaScript，请务必使用默认导出作为最重要的项目，并为次要项目命名导出。

### 多个参数优先于单个对象参数 ###

声明一个函数时，您应该总是喜欢多个参数而不是一个期望对象的参数：

` // GOOD function displayUser ( firstName, lastName, age ) { console.log( `This is ${firstName} ${lastName}. She is ${age} years old.` ) } // BAD function displayUser ( user ) { console.log( `This is ${user.firstName} ${user.lastName}. She is ${user.age} years old.` ) } 复制代码`

这背后的原因是，当您查看函数声明的第一行时，您能确切知道需要传递给函数的内容。

尽管函数应该受到限制 - 只做一项工作 - 但是它可能会变得更大。在函数体中扫描需要传递的变量（嵌套在对象中）将花费更多时间。有时，使用整个对象并将其传递给函数似乎更容易，但为了扩展应用程序，此设置肯定会有所帮助。

在某种程度上，声明特定参数没有意义。对我来说，它超过四个或五个功能参数。如果你的函数变大，你应该转向使用对象参数。

这里的主要原因是参数需要以特定顺序传递。 如果您有可选参数，则需要传递undefined或null。 使用对象参数，您可以简单地传递整个对象，其中顺序和未定义的值无关紧要。

### 解构（Destructuring） ###

解构是ES6引入的一个很好的工具。它允许您从对象中获取特定字段并立即将其分配给变量。 您可以将它用于任何类型的对象或模块。

` // EXAMPLE FOR MODULES const { add, subtract } = require ( './calculations' ) 复制代码`

只导入您需要在文件中使用的函数而不是整个模块，然后从中访问特定的函数。 同样，当您确定您确实需要一个对象作为函数参数时，也可以使用destructuring。 这仍将为您提供函数内所需内容的概述：

` function logCountry ( {name, code, language, currency, population, continent} ) { let msg = `The official language of ${name} ` if (code) msg += `( ${code} ) ` msg += `is ${language}. ${population} inhabitants pay in ${currency}.` if (contintent) msg += ` The country is located in ${continent} ` } logCountry({ name : 'Germany' , code : 'DE' , language 'german' , currency : 'Euro' , population : '82 Million' , }) logCountry({ name : 'China' , language 'mandarin' , currency : 'Renminbi' , population : '1.4 Billion' , continent : 'Asia' , }) 复制代码`

正如你所看到的，我仍然知道我需要传递什么给函数 - 即使它被包装在一个对象中。要解决了解所需内容的问题，请参阅下一个提示！（顺便说一句，这也适用于React功能组件。）

### 使用默认值 ###

解构的默认值甚至基本函数参数都非常有用。首先，它们为您提供了一个可以传递给函数的值的示例。其次，您可以指出哪些值是必需的，哪些值不是。使用前面的示例，该函数的完整设置如下所示：

` function logCountry ( { name = 'United States' , code, language = 'English' , currency = 'USD' , population = '327 Million' , continent, } ) { let msg = `The official language of ${name} ` if (code) msg += `( ${code} ) ` msg += `is ${language}. ${population} inhabitants pay in ${currency}.` if (contintent) msg += ` The country is located in ${continent} ` } logCountry({ name : 'Germany' , code : 'DE' , language 'german' , currency : 'Euro' , population : '82 Million' , }) logCountry({ name : 'China' , language 'mandarin' , currency : 'Renminbi' , population : '1.4 Billion' , continent : 'Asia' , }) 复制代码`

显然，有时您可能不想使用默认值，而是在未传递值时抛出错误。 然而，这通常是一个方便的技巧。

### 数据稀缺性 ###

前面的技巧引出了一个结论：最好不要传递您不需要的数据。同样，在设置函数时，这可能意味着更多的工作。但是，从长远来看，它肯定会为您提供更具可读性的代码库。确切地知道在特定位置使用哪些值是非常有价值的。

### 行数和缩进限制 ###

我见过大文件 - 非常大的文件。实际上，超过3,000行代码。在这些文件中查找逻辑块是非常困难的。

因此，您应该将文件大小限制为一定数量的行。我倾向于将我的文件保存在100行代码之下。 有时候，很难分解文件，它们会增长到200-300行，在极少数情况下会增加到400行。

超过此临界值，意味着文件太杂乱，难以维护。随意创建新的模块和文件夹。您的项目应该看起来像一个森林，由树（模块部分）和分支（模块和模块文件组）组成。避免试图模仿阿尔卑斯山，在密闭区域堆积代码。

相比之下，你的实际文件应该看起来像Shire，这里和那里都有一些山丘（小水平的缩进），但一切都相对平坦。 尽量将压痕水平保持在四级以下。

也许为这些提示启用eslint规则是有帮助的！

### 使用prettier ###

在团队中工作需要清晰的样式指南和格式。ESLint提供了一个巨大的规则集，您可以根据自己的需求进行自定义。还有 ` eslint--fix` ，它可以纠正一些错误，但不是全部。

相反，我建议使用 [Prettier]( https://link.juejin.im?target=https%3A%2F%2Fprettier.io%2F ) 格式化代码。这样，开发人员不必担心代码格式化，而只需编写高质量的代码。 外观将一致并且格式自动化。

### 使用有意义的变量名 ###

理想情况下，应根据其内容命名变量。 以下是一些有助于您声明有意义的变量名称的指南。

#### 函数 ####

函数通常执行某种操作。为了解释这一点，人类使用动词 - 转换或显示，例如。在开头用动词命名函数是个好主意，例如 ` convertCurrency` 或 ` displayUserName` 。

#### 数组 ####

这些通常会包含一系列项目; 因此，将s附加到变量名称。 例如：

` const students = [ 'Eddie' , 'Julia' , 'Nathan' , 'Theresa' ] 复制代码`

#### 布尔 ####

简单地说就是尽量多接近于自然语言，这样好理解。你会问“这个人是教师吗？”→“是”或“否”。同样：

` const isTeacher = true // OR false 复制代码`

#### 数组函数 ####

` forEach` , ` map` , ` reduce` , ` filter` 等是很好的原生JavaScript函数，用于处理数组和执行某些操作。 我看到很多人只是将 ` el` 或 ` element` 作为参数传递给回调函数。 虽然这很简单快捷，但您还应根据其值来命名。 例如：

` const cities = [ 'Berlin' , 'San Francisco' , 'Tel Aviv' , 'Seoul' ] cities.forEach( function ( city ) { ... }) 复制代码`

#### 标识 ####

通常，您必须跟踪特定数据集和对象的ID。当嵌套id时，只需将其保留为id即可。在这里，我喜欢在将对象返回到前端之前将MongoDB ` _id` 映射到 ` id` 。从对象中提取 ` id` 时，请预先添加对象的类型。例如：

` const studentId = student.id // OR const { id : studentId } = student // destructuring with renaming 复制代码`

该规则的一个例外是模型中的MongoDB引用。 在这里，只需在引用的模型之后命名字段即可。 这将在填充参考文档时保持清晰：

` const StudentSchema = new Schema({ teacher : { type : Schema.Types.ObjectId, ref : 'Teacher' , required : true , }, name : String , ... }) 复制代码`

### 尽可能使用async / await ###

在可读性方面，回调是最糟糕的 - 特别是在嵌套时。Promises是一个很好的改进，但在我看来，async / await具有最好的可读性。即使对于初学者或来自其他语言的人来说，这也会有很大帮助。但是，请确保您了解其背后的概念，并且不要盲目地在任何地方使用它。

### 模块导入顺序 ###

正如我们在技巧1和2中看到的那样，将逻辑保持在正确的位置是可维护性的关键。 同样，导入不同模块的方式可以减少文件中的混淆。 我在导入不同模块时遵循一个简单的结构：

` // 3rd party packages import React from 'react' import styled from 'styled-components' // Stores import Store from '~/Store // reusable components import Button from '~/components/Button' // utility functions import { add, subtract } from '~/utils/calculate' // submodules import Intro from './Intro' import Selector from './Selector' 复制代码`

我在这里使用了React组件作为示例，因为有更多类型的导入。 您应该能够根据您的具体用例进行调整。

### 摆脱控制台 ###

` console.log` 是一种很好的调试方式 - 非常简单，快速，完成工作。显然，有更复杂的工具，但我认为每个开发人员仍然使用它。如果您忘记清理日志，您的控制台最终将陷入巨大的混乱。然后，您确实要在代码库中保留日志; 例如，警告和错误。

要解决此问题，您仍然可以使用 ` console.log` 进行调试，但对于持久日志，请使用 [loglevel]( https://link.juejin.im?target=https%3A%2F%2Fwww.npmjs.com%2Fpackage%2Floglevel ) 或 [winston]( https://link.juejin.im?target=https%3A%2F%2Fwww.npmjs.com%2Fpackage%2Fwinston ) 等库。此外，您可以使用ESLint警告控制台语句。这样你就可以轻松地全局查找 ` console...` 并删除这些语句。

遵循这些指导原则确实帮助我保持代码库的清洁和可扩展性。 你觉得还有什么提示特别有用的吗？请在评论中告诉我您在编码工作过程中值得推荐的内容，并请分享您用于帮助代码结构的任何其他提示！谢谢~

**如果觉得文章对你有些许帮助，欢迎在 [我的GitHub博客]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fmiqilin21%2Fmiqilin21.github.io ) 点赞和关注，感激不尽！**