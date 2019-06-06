# 在JavaScript中将值转换为字符串的5种方法 #

![](https://user-gold-cdn.xitu.io/2019/6/5/16b26e9ce43810b7?imageView2/0/w/1280/h/960/ignore-error/1) 如果您关注Airbnb的样式指南，首选方法是使用 ` “String（）”` 👍

它也是我使用的那个，因为它是最明确的 - 让其他人轻松地遵循你的代码的意图🤓

请记住，最好的代码不一定是最聪明的方式，它是最能将代码理解传达给他人的代码💯

` const value = 12345; // Concat Empty String value + '' ; // Template Strings ` ${value} `; // JSON.stringify JSON.stringify(value); // toString() value.toString(); // String() String(value); // RESULT // '12345' 复制代码`

## 比较5种方式 ##

好吧，让我们用不同的值测试5种方式。以下是我们要对其进行测试的变量：

` const string = "hello" ; const number = 123; const boolean = true ; const array = [1, "2" , 3]; const object = {one: 1 }; const symbolValue = Symbol( '123' ); const undefinedValue = undefined; const nullValue = null; 复制代码`

## 结合空字符串 ##

` string + '' ; // 'hello' number + '' ; // '123' boolean + '' ; // 'true' array + '' ; // '1,2,3' object + '' ; // '[object Object]' undefinedValue + '' ; // 'undefined' nullValue + '' ; // 'null' // ⚠️ symbolValue + '' ; // ❌ TypeError 复制代码`

从这里，您可以看到如果值为一个Symbol ，此方法将抛出 ` TypeError` 。否则，一切看起来都不错。

## 模板字符串 ##

` ` ${string} `; // 'hello' ` ${number} `; // '123' ` ${boolean} `; // 'true' ` ${array} `; // '1,2,3' ` ${object} `; // '[object Object]' ` ${undefinedValue} `; // 'undefined' ` ${nullValue} `; // 'null' // ⚠️ ` ${symbolValue} `; // ❌ TypeError 复制代码`

使用模版字符串的结果与结合空字符串的结果基本相同。同样，这可能不是理想的处理方式，因为 ` Symbol` 它会抛出一个 ` TypeError` 。

如果你很好奇，那就是 ` TypeError： TypeError: Cannot convert a Symbol value to a string`

## JSON.stringify（） ##

` // ⚠️ JSON.stringify(string); // '"hello"' JSON.stringify(number); // '123' JSON.stringify(boolean); // 'true' JSON.stringify(array); // '[1,"2",3]' JSON.stringify(object); // '{"one":1}' JSON.stringify(nullValue); // 'null' JSON.stringify(symbolValue); // undefined JSON.stringify(undefinedValue); // undefined 复制代码`

因此，您通常不会使用 ` JSON.stringify` 将值转换为字符串。而且这里真的没有强制发生。因此，您了解可用的所有工具。然后你可以决定使用什么工具而不是根据具体情况使用👍

有一点我想指出，因为你可能没有注意它。当您在实际string值上使用它时，它会将其更改为带引号的字符串。

##.toString（） ##

` string.toString(); // 'hello' number.toString(); // '123' boolean.toString(); // 'true' array.toString(); // '1,2,3' object.toString(); // '[object Object]' symbolValue.toString(); // 'Symbol(123)' // ⚠️ undefinedValue.toString(); // ❌ TypeError nullValue.toString(); // ❌ TypeError 复制代码`

所以PK其实就是在 ` toString()` 和 ` String()` ，当你想把一个值转换为字符串。除了它会为 ` undefined` 和 ` null` 抛出一个错误，其他表现都很好。所以一定要注意这一点。

## String（） ##

` String(string); // 'hello' String(number); // '123' String(boolean); // 'true' String(array); // '1,2,3' String(object); // '[object Object]' String(symbolValue); // 'Symbol(123)' String(undefinedValue); // 'undefined' String(nullValue); // 'null' 复制代码`

好吧，我想我们找到了胜利者🏆

正如你所看到的， ` String()` 处理 ` null` 和 ` undefined` 相当不错。不会抛出任何错误 - 除非这是你想要的。一般来说记住我的建议。您将最了解您的应用程序，因此您应该选择最适合您情况的方式。

## 结论：String（）🏆 ##

在向您展示了所有不同方法如何处理不同类型的值之后。希望您了解这些差异，并且您将知道下次处理代码时要使用的工具。如果你不确定， ` String()` 总是一个很好的默认选择👍