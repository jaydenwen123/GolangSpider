# 翻译： JavaScript中对象解构的3种实际应用 #

(渣渣翻译，如果看翻译不开心可以看-> [原文]( https://link.juejin.im?target=https%3A%2F%2Fmedium.com%2Fbetter-programming%2F3-practical-uses-of-object-destructuring-in-javascript-a2c34ce3367b ) ， )

现在你可能非常熟悉js中解构的概念了！解构这个概念来自2015年ES6草案， 但如果你需要更深一步的了解它， Mozilla有一篇更深入的文章说明它是如何工作的。(文章底部)

但是，了解解构如何工作并不等于我们了解了如何使用它。使用这三个解构模式可以让你的代码更加清晰， 更强大， 更具可读性。

## Named Function Arguments (命名函数参数) ##

相较于我们通过传入参数的位置来控制参数,解构模式用于形式参数是一个很好的代替方式.你只需按名称指定参数，而不是按照与函数签名相同的顺序排序参数。例如，在Python中：

` def sum (a= 1 , b= 2 , c= 3 ) : return a + b + c sum(b= 5 , c= 10 ) 复制代码`

就像你所看到的一样， 参数的顺序并不是问题， 你通过名字指定了他们。命名参数相较于基于位置的参数命名有以下的好处：

* 调用函数时， 可以省略一个或者多个参数
* 当传参的时候，顺序并不是问题
* 调用可能存在于其他地方的函数时，代码更具可读性

虽然JavaScript中不存在真正的命名参数，但我们可以使用解构模式来实现所有3个相同的好处。这是和上面python相同功能的代码， 但是在js中我们可以这样：

` const sum = ( {a= 1 , b= 2 , c= 3 } ) => { return a + b + c; } sum({ b : 5 , a : 1 }); // 9 复制代码`

这种模式符合我们命名参数的所有目标。我们能够省去参数c，顺序无关紧要，我们通过名称引用它们来分配我们的参数。这一切都可以通过对象解构来实现。

## Cleanly Parse a Server Response (清晰解析服务响应) ##

通常我们只关注服务响应内容里的data，或者只关注在data中的一个特定字段。在这个例子中，你可以使用解构来仅获取该字段的值，同时忽略服务器通常发回的许多其他内容。这是一个代码示例：

` const mockServer = () => { return new Promise ( ( resolve, reject ) => { setTimeout( () => { resolve({ 'status' : 200 , 'content-type' : 'application/json' , 'data' : { dataInfo : 42 } }) }, 1000 ); }) } mockServer().then( ( {data: {dataInfo = 100 }} ) => { console.log(dataInfo); }) 复制代码`

此模式允许您在解析参数时从对象中提取值。你还可以自由的设置默认值！

## Setting Default Values During Assignment (在赋值的时候给默认值) ##

给变量或常量赋值时的常见情况是，如果范围中当前不存在时，则给一个默认值。

在没有解构之前， 你可能通过以下代码来实现这个功能：

` var nightMode = userSettings.nightMode || false ; 复制代码`

但这需要为每个赋值写一行代码。通过解构，您可以同时处理所有的赋值及提供默认值。

` const userSettings = { fontSize : 'large' , nightMode : true }; const { nightMode = false , language = 'en' , fontSize = 'normal' } = userSettings; console.log(nightMode, language, fontSize); 复制代码`

解构模式也能应用于react组件中的state.

我希望你能够将一些这些模式应用到你的代码中！查看下面的链接，了解有关解构的更多信息。

[ES6 In Depth: Destructuring - Mozilla Hacks - the Web developer blog]( https://link.juejin.im?target=https%3A%2F%2Fhacks.mozilla.org%2F2015%2F05%2Fes6-in-depth-destructuring%2F )

[Learn the basics of destructuring props in React]( https://link.juejin.im?target=https%3A%2F%2Fmedium.freecodecamp.org%2Fthe-basics-of-destructuring-props-in-react-a196696f5477 )

(完)

[原文链接]( https://link.juejin.im?target=https%3A%2F%2Fmedium.com%2Fbetter-programming%2F3-practical-uses-of-object-destructuring-in-javascript-a2c34ce3367b )