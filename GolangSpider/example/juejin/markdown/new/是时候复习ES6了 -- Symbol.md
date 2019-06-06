# 是时候复习ES6了 -- Symbol #

### 前言 ###

ES6之中新增了许多的的新的概念还使用方式，Symbol作为新的基础数据类型之一，我们也需要重点的去理解学习。今天就然我们来学习这一新的知识点吧

### 正文-Symbol ###

这是JS之中一种新的基本类型数据，数据类型就是Symbol，虽然类似于字符串，但是其表示的意思是独一无二的意思。ES6引入这一类型的用意在于，常常我们在命名的时候会有相同名称冲突的问题，所以Symbol成为我们这个问题的直接方式。

* 我们可以通过Symbol(String) 来获取一个Symbol类型的变量，这里不适用new的形式是应为Symbol作为基础类型，不能视为对象变量。所以就使用这一方式来获取。其中可以传递一个String类型的参数，表示对symbol实例的描述，同时也方便区分。当然我们参数也可以传递对象内容，但是会先自动的调用Object的toString方法转化为字符串，然后再生成symbol数据。
* typeof在判断所有的Symbol的值的时候返回类型将会是Symbol
* symbol类型的数据，哪怕描述符是相同的，但是只要是两个变量的比对，永远是false。
* symbol的值不可以和其他类型的值做计算。但是其有toString的方法，可以用来转化成为字符串。
* Symbol.prototype.description:这个参数获取的是当前symbol变量的描述符，即我们传递如Symbol方法的那个字符串参数。
* Symbol的变量可以作为对象内容，从而表示我们的对象变量抿成的独一无二，我们再对象的扩展概念之中有提及到。对象的属性的名称可以使用[name]的方式来进行确定，name可以是变量也可以是表达式。所以我们可以 **使用Symbol的变量来作为Object的属性，这样可以完全的防止对象属性的重名的问题** 。但是我们要注意的一点是，我们使用Symbol作为实行名称的时候是不能通过点操作符来书写的。
* Symbol还可以用于常量的控制，保证单个常量并不会与其他的常量的值相同。
* Symbol还可以将与程序强耦合的字符串或者数字，脱离耦合，改成更有语义化的程序代码段。
* Symbol作为对象属性的时候，再所有的遍历方法之中都是没有办法获取他的信息的。这个时候Object之中有一 **个getOwnPropertySymbols方法，可以用来返回我们对象之中所有的Symbol属性组成的数组** 。但是我们要注意的是Symbol属性对然不容易遍历到，但是它本身还是共有属性而不是私有的。但是我们也可以通过这个特性为我们的对象定义一些非私有但是只在对象内部使用的方法。

接下来我们来看一下Symbol为我们提供一些什么属性和方法吧：

* Symbol.for(String):有的时候我们希望使用同一个Symbol的值，这是后我们可以传递相关的Symbol描述符，如果 **方法比对到全局环境下有登记了的描述符相同的Symbol变量，则方法返回已有的Symbol变量，如果没有将返回一个新的Symbol变量，且描述符是传递的参数** 。这一方法创建的对象内容会有登记再全局环境之中，但是我们上面提到的Symbol()这个方法是不会登记的。
* Symbol.forKey(Symbol)：传递的symbol类型的值，返回的是当前值的说明符（key）。这里的Symbol参数需要时登记了的变量，如果时没有登记的变量将会返回undefined。
* Symbol.hasInstance:这个属性指向的时一个内部的方法，当我们再使用instanceOf来做判断的时候，实际上时再函数的内部调用了这个属性对应的方法，所以我们实际可以重写当前对象的InstanceOf的操作。看个例子

` class MyClass { [Symbol.hasInstance](foo) { return foo instanceof Array; } } [1, 2, 3] instanceof new MyClass() // true 但是再上面我们可以看到instanceof之后需要一个实例化的 过程，但是我们使用instanceof的时候往往想要直接使用类 的名称就好了，这是后我们可以吧这个方法写成静态的，如下： class Even { static [Symbol.hasInstance](obj) { return Number(obj) % 2 === 0; } } 1 instanceof Even // false 2 instanceof Even // true 复制代码`

* Symbol.isConcatSpreadable : 我们可以从字面上看，is开头的属性往往都时布尔值类型的，concat再数组之中的常常出现用于合并内容的情况，Spreadable中文意思表示可拆分。所以我们可以猜测一下，这个属性时用于再合并方法中，当前对象是否可以拆分的判断。这个属性的实际用途就是再concat函数之中才能体现，合并的时候数组对象默认可以展开，当设置这一属性为false的时候，合并的时候，数组将会以一个元素添加到新的合并数组之中。类数组对象则刚好相反，如默认情况下不拆开，如果设置当前这一属性时true的时候则拆开内容。我们看个例子吧：

` let arr1 = [ 'c' , 'd' ]; [ 'a' , 'b' ].concat(arr1, 'e' ) // [ 'a' , 'b' , 'c' , 'd' , 'e' ] arr1[Symbol.isConcatSpreadable] // undefined let arr2 = [ 'c' , 'd' ]; arr2[Symbol.isConcatSpreadable] = false ; [ 'a' , 'b' ].concat(arr2, 'e' ) // [ 'a' , 'b' , [ 'c' , 'd' ], 'e' ] let obj = {length: 2, 0: 'c' , 1: 'd' }; [ 'a' , 'b' ].concat(obj, 'e' ) // [ 'a' , 'b' , obj, 'e' ] obj[Symbol.isConcatSpreadable] = true ; [ 'a' , 'b' ].concat(obj, 'e' ) // [ 'a' , 'b' , 'c' , 'd' , 'e' ] 复制代码`

* Symbol.species 这一属性指向的是一个函数，创建衍生对象的时候，会自动的调用。什么是衍生对象就是通过对象的某一些方法，创建出来的新的，但是类型依旧是原本对象的类型，这一类对象我们可以程之为衍生对象，例如：Array之中的map和filter方法，可以说就是创建了衍生对象内容。而species这个指向的方法则使用与再衍生对象创造的时候返回的值。他需要是一个getter方法，返回的是当前的衍生对象的类型内容。上例子

` class T1 extends Promise { } class T2 extends Promise { static get [Symbol.species]() { return Promise; } } new T1(r => r()).then(v => v) instanceof T1 // true new T2(r => r()).then(v => v) instanceof T2 // false 复制代码`

* Symbol.match 这个属性指向的是一个方法内容，我们再调用String.prototype.match实际上就相当于调用了，RegExp[Symbol.match](this)，这一调用返回值就是match方法最终结果。
* Symbol.replace 这个属性指代的也是一个方法，向上面的match一样，它实际上就是String.prototype.replace方法调用的真实的返回值。
* Symbol.search属性，指向一个方法，该对象调用String.prototype.search方法的时候，会返回该方法的返回值。
* Symbol.split属性，指向一个方法，该对象被String.prototype.split方法调用时，会返回该方法的返回值。
* Symbol.iterator属性，指向该方法的默认遍历器方法。
* Symbol.toPrimitive属性，指向一个方法。再该对象转换成原始类型的值的时候，会调用这个方法，返回该对象原始类型的值。被调用的时候接受一个参数，表明当前运算的模式，一共是三种模式，Number转换成数值类型，String转换成字符串类型，Default转换成数值或者字符串类型。
* Symbol.toStringTag属性，指向一个方法。主要是用于定制，当调用Object.prototype.toString()方法的时候，将当前函数返回的内容放在原本的[object Object]之中的第二个Object上面。
* Symbol.unscopables属性，指向一个方法，该对象指定了使用with关键字时，哪些属性会被with环境排除。由于with关键字不常用，所以我们的对于这一属性的使用相当稀少，我这里就不再多说，大家想要了解清楚的可以自行的到官网上面进行学习。

### 结束 ###

Symbol是一个比较关键的内容，所以这里啰嗦了一些，概念型的东西总是复杂的。如果有什么错误的地方望帮个忙，指出一下。希望共同进步，共同成长。拜了个拜。