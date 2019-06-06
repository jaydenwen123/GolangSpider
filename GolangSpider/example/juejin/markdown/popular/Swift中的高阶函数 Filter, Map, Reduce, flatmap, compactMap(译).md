# Swift中的高阶函数: Filter, Map, Reduce, flatmap, compactMap(译) #

自从我详细了解函数和闭包后，我就想去知道它们的优点和在编程中的使用， 我的理解是高阶函数的使用是基于集合类型的。

根据我的理解，高阶函数就是把另一个函数或者闭包当做参数并且返回值。

首先，我来解释一下。思考下面的代码将会使你对高阶函数的理解更深：

` // 将函数当做参数传递给另一个函数 func addation(num1: Double, num2: Double) -> Double { return num1 + num2 } func multiply(num1: Double, num2: Double) -> Double { return num1 * num2 } func do MathOperation(operation: (_ x: Double, _ y: Double) -> Double, num1: Double, num2: Double) -> Double { return operation(num1, num2) } // print 12 do MathOperation(operation: addation(num1:num2:), num1: 10, num2: 2) // print 20 do MathOperation(operation: multiply(num1:num2:), num1: 10, num2: 2) 复制代码`

` addation(num1:num2:)` 、 ` multiply(num1:num2:)` 都是 ` (Double,Double)->Double` 的函数。 ` addation(num1:num2:)` 接受两个 Double 的值，并返回它们的和； ` multiply(num1:num2:)` 接受两个 Double 的值，并返回它们的乘积。 ` doMathOperation(operation:num1:num2:)` 则是一个接受三个参数的高阶函数，它接受两个 Double 的参数和一个 ` (Double,Double)->Double` 类型的函数。看一下该函数的调用，你应该就理解了高阶函数的工作原理。

在Swift中，函数和闭包是 ` 一等公民` 。它们可以存储在变量中被传递。

` // 函数返回值为另一个函数 func do ArithmeticOperation(isMultiply: Bool) -> (Double, Double) -> Double { func addation(num1: Double, num2: Double) -> Double { return num1 + num2 } func multiply(num1: Double, num2: Double) -> Double { return num1 * num2 } return isMultiply ? multiply : addation } let operationToPerform1 = do ArithmeticOperation(isMultiply: true ) let operationToPerform2 = do ArithmeticOperation(isMultiply: false ) operationToPerform1(10, 2) //20 operationToPerform2(10, 2) //12 复制代码`

在上面的代码中， ` doArithmeticOperation(isMultiply:)` 是一个返回值类型为 ` (Double,Double)->Double` 的高阶函数。它基于一个布尔值来判断返回值。operationToPerform1 执行乘的操作，operationToPerform2 执行加的操作。通过上面的函数定义和调用，你就理解所有的事情了。

当然，你可以通过多种不同的方法来实现一样的东西。你可以使用闭包来代替函数，你可以使用枚举来判断函数的操作。在这里，我只是通过上面的代码来解释一下什么是高阶函数。

下面是 Swift 库中自带的一些高阶函数，如果我理解没错的话，下面的函数接受闭包类型的参数。你可以在集合类型(Array Set Dictionary)中使用它们。如果你对闭包不太了解，你可以通过 [这篇文章]( https://link.juejin.im?target=https%3A%2F%2Fmedium.com%2F%40abhimuralidharan%2Ffunctional-swift-all-about-closures-310bc8af31dd ) 来学习。

### Map ###

**map 函数的作用就是对集合进行一个循环，循环内部再对每个元素做同一个操作。** 它返回一个包含映射后元素的数组。

#### Map on Array： ####

假设我们有一个整数型数组：

` let arrayOfInt = [2,3,4,5,4,7,2] 复制代码`

我们如何让每个数都乘10呢？我们通常使用 ` for-in` 来遍历每个元素，然后执行相关操作。

` var newArr: [Int] = [] for value in arrayOfInt { newArr.append(value*10) } print (newArr) // prints [20, 30, 40, 50, 40, 70, 20] 复制代码`

上面的代码看着有点冗余。它包含创建一个新数组的样板代码，我们可以通过使用 map 来避免。我们通过 Swift 的自动补全功能可以看到 map 函数接受有一个 Int 类型的参数并返回一个泛型的闭包。

![](https://user-gold-cdn.xitu.io/2019/6/4/16b208da0b5297d3?imageView2/0/w/1280/h/960/ignore-error/1)

` let newArrUsingMap = arrayOfInt.map { $0 * 10 } 复制代码`

对于一个整型数组，这是 map 的极简版本。我们可以在闭包中使用 ` $` 操作符来代指遍历的每个元素。

下面的代码作用都是一样的，它们表示了一个从繁到简的过程。通过下面的代码，你应该对闭包有了一个清晰的认识。

![](https://user-gold-cdn.xitu.io/2019/6/4/16b20912ab040c6e?imageView2/0/w/1280/h/960/ignore-error/1)

**map 的工作原理** ：map 函数接受一个闭包作为参数，在迭代集合时调用该闭包。这个闭包映射集合中的元素，并将结果返回。map 函数再将结果放在数组中返回。

#### Map on Dictionary： ####

假设我们有一个书名当做 key ，书的价格当做 value 的字典。

` let bookAmount = [“harrypotter”:100.0, “junglebook”:100.0] 复制代码`

如果你试图 map 这个字典，Swift 的自动补全将是这样：

![](https://user-gold-cdn.xitu.io/2019/6/5/16b255140a22e344?imageView2/0/w/1280/h/960/ignore-error/1)

` let bookAmount = [ "harrypotter" : 100.0, "junglebook" : 100.0] let return FormatMap = bookAmount.map { (key, value) in return key.capitalized } print ( return FormatMap) //[ "Junglebook" , "Harrypotter" ] 复制代码`

我们通过上面的代码，对一个字典进行遍历，每次遍历在闭包中都有一个 String 类型的 key ，和一个 Double 类型的 value 。返回值为一个大写首字母的字符串数组，数组的值还可以是价格或者元组，这取决于你的需求。

` 注意：map 函数的返回值类型总是一个泛型数组。你可以返回包含任意类型的数组。`

#### Map on set: ####

` let lengthInMeters: Set = [4.0, 6.2, 8.9] let lengthInFeet = lengthInMeters.map { $0 * 3.2808 } print (lengthInFeet) //[20.340960000000003, 13.1232, 29.199120000000004] 复制代码`

在上面的代码中，我们有一个值类型为 Double 的 set，我们的闭包返回值也是 Double 类型。 ` lengthInMeters` 是一个 set，而 ` lengthInFeet` 是一个数组。

#### 如果你想在 map 的时候获取 index 应该怎么做？ ####

答案很简单，你必须在 map 之前调用 ` enumerate` 。

![](https://user-gold-cdn.xitu.io/2019/6/5/16b255fbf97ae77c?imageView2/0/w/1280/h/960/ignore-error/1)

下面是示例代码：

` let numbers = [1, 2, 4, 5] let indexAndNum = numbers.enumerated().map { (index,element) in return "\(index):\(element)" } print (indexAndNum) // [“0:1”, “1:2”, “2:4”, “3:5”] 复制代码`

### Filter ###

` filter` 会遍历集合，返回一个包含符合条件元素的数组。

#### Filter on array ####

假设我们要筛选一个整型数组中包含的偶数，你可能会写下面的代码：

` let arrayOfIntegers = [1, 2, 3, 4, 5, 6, 7, 8, 9] var newArray = [Int]() for integer in arrayOfIntegers { if integer % 2 == 0 { newArray.append( integer ) } } print (newArray) //[2, 4, 6, 8] 复制代码`

就像 map ，这是一个简单的函数去筛选集合中的元素。

如果我们对整形数组使用 filter ，Swift 的自动补全展示如下：

![](https://user-gold-cdn.xitu.io/2019/6/5/16b2569a261ae6fc?imageView2/0/w/1280/h/960/ignore-error/1)

如你所见， filter 函数调用了一个接受一个 Int 类型参数、返回值为 Bool 类型、名字为 ` isIncluded` 的闭包。 ` isIncluded` 在每次遍历都会返回一个布尔值，然后基于布尔值将创建一个新的包含筛选结果的数组。

我们可以通过 filter 将上面的代码修改为：

` var newArray = arrayOfIntegers.filter { (value) -> Bool in return value % 2 == 0 } 复制代码`

filter 闭包也可以被简化，就像 map：

![](https://user-gold-cdn.xitu.io/2019/6/5/16b256f340acc0f1?imageView2/0/w/1280/h/960/ignore-error/1)

#### Filter on dictionary ####

假设有一个书名当做 key ，书的价格当做 value 的字典。

` let bookAmount = [ "harrypotter" :100.0, "junglebook" :1000.0] 复制代码`

如果你想对这个字典调用 filter 函数，Swift 的自动补全将是这样：

![](https://user-gold-cdn.xitu.io/2019/6/5/16b25718b662bc53?imageView2/0/w/1280/h/960/ignore-error/1)

filter 函数会调用一个名字为 isIncluded 的闭包，该闭包接受一个键值对作为参数，并返回一个布尔值。最终，基于返回的布尔值，filter 函数将决定是否将键值对添加到数组中。

` <重要> 原文中作者写道：对字典调用 Filter 函数，将返回一个包含元组类型的数组。但译者在 playground 中发现 返回值实际为字典类型的数组。`

` let bookAmount = [ "harrypotter" : 100.0, "junglebook" : 1000.0] let results = bookAmount.filter { (key, value) -> Bool in return value > 100 } print (results) //[ "junglebook" : 1000.0] 复制代码`

![](https://user-gold-cdn.xitu.io/2019/6/5/16b257814370c7b5?imageView2/0/w/1280/h/960/ignore-error/1)

还可以将上述代码简化：

` // $0 为 key $1 为 value let results = bookAmount.filter { $1 > 100 } 复制代码`

#### Filter on set ####

` let lengthInMeters: Set = [1, 2, 3, 4, 5, 6, 7, 8, 9] let lengthInFeet = lengthInMeters.filter { $0 > 5 } print (lengthInFeet) //[9, 8, 7, 6] 复制代码`

在每次遍历时， filter 闭包接受一个 Double 的参数，返回一个布尔值。筛选数组中包含的元素基于返回的布尔值。

### Reduce ###

reduce ：联合集合中所有的值，并返回一个新值。

Apple 的官方文档如下：

` func reduce<Result>(_ initialResult: Result, _ nextPartialResult: (Result, Element) throws -> Result) rethrows -> Result 复制代码`

reduce 函数接受两个参数：

* 第一个为初始值，它用来存储初始值和每次迭代中的返回值。
* 另一个参数是一个闭包，闭包包含两个参数：初始值或者当前操作的结果、集合中的下一个 item 。

#### Reduce on arrays ####

让我们通过一个例子来理解 reduce 的具体作用：

` let numbers = [1, 2, 3, 4] let numberSum = numbers.reduce(0) { (x, y) in return x + y } print (numberSum) // 10 复制代码`

reduce 函数会迭代4次。

` 1.初始值为0，x为0，y为1 -> 返回 x + y 。所以初始值或者结果变为 1。 2.初始值或者结果变为 1，x为1，y为2 -> 返回 x + y 。所以初始值或者结果变为 3。 3.初始值或者结果变为 3，x为3，y为3 -> 返回 x + y 。所以初始值或者结果变为 6。 4.初始值或者结果变为 6，x为6，y为4 -> 返回 x + y 。所以初始值或者结果变为 10。 复制代码`

reduce 函数可以简化为：

` let reducedNumberSum = numbers.reduce(0) { $0 + $1 } print (reducedNumberSum) // prints 10 复制代码`

![](https://user-gold-cdn.xitu.io/2019/6/5/16b259771af1183f?imageView2/0/w/1280/h/960/ignore-error/1)

在本例中，闭包的类型为 ` (Int,Int)->Int` 。所以，我们可以传递类型为 ` (Int,Int)->Int` 的任意函数或者闭包。比如我们可以把操作符替换为 -, *, / 等。

` let reducedNumberSum = numbers.reduce(0,+) // returns 10 复制代码`

我们可以在闭包里添加 ` *` 或者其他的操作符。

` let reducedNumberSum = numbers.reduce(0) { $0 * $1 } // reducedNumberSum is 0... 复制代码`

上面的代码也可以写成这样：

` let reducedNumberSum = numbers.reduce(0,*) 复制代码`

reduce 也可以通过 + 操作符来合并字符串。

` let codes = [ "abc" , "def" , "ghi" ] let text = codes.reduce( "" ) { $0 + $1 } //the result is "abcdefghi" or let text = codes.reduce( "" ,+) //the result is "abcdefghi" 复制代码`

#### Reduce on dictionary ####

让我们来 reduce bookAmount。

` let bookAmount = [ "harrypotter" :100.0, "junglebook" :1000.0] let reduce1 = bookAmount.reduce(10) { (result, tuple) in return result + tuple.value } print (reduce1) //1110.0 let reduce2 = bookAmount.reduce( "book are " ) { (result, tuple) in return result + tuple.key + " " } print (reduce2) //book are junglebook harrypotter 复制代码`

对于字典，reduce 的闭包接受两个参数。

` 1.一个应该被 reduce 的初始值或结果 2.一个当前键值对的元组 复制代码`

reduce2 可以被简化为：

` let reducedBookNamesOnDict = bookAmount.reduce( "Books are " ) { $0 + $1.key + " " } //or $0 + $1.0 + " " 复制代码`

#### Reduce on set ####

Set 中的 reduce 使用和数组中的一致。

` let lengthInMeters: Set = [4.0, 6.2, 8.9] let reducedSet = lengthInMeters.reduce(0) { $0 + $1 } print (reducedSet) //19.1 复制代码`

闭包中的返回值类型为 Double。

### Flatmap ###

Flatmap 用来铺平 collections 中的 collection 。在铺平 collection 之前，我们对每一个元素进行 map 操作。 ` Apple docs 解释: 返回一个对序列的每个元素进行形变的串级结果( Returns an array containing the concatenated results of calling the given transformation with each element of this sequence.)`

` 解读 : map + (Flat the collection) 复制代码`

![](https://user-gold-cdn.xitu.io/2019/6/5/16b25ab438f3251e?imageView2/0/w/1280/h/960/ignore-error/1) ` 图1 对 flatmap 进行代码说明`

![](https://user-gold-cdn.xitu.io/2019/6/5/16b25ab6137d5395?imageView2/0/w/1280/h/960/ignore-error/1) ` 图2`

在图2 中，flatMap 迭代 collections 中的所有 collection 进行大写操作。在这个例子中，每个 collection 是字符串。下面是执行步骤：

* 对所有的字符串执行 ` upperCased()` 函数，这类似于:
` [“abc”,”def”,”ghi”].map { $0.uppercased() } 复制代码`

输出：

` output: [“ABC”, “DEF”, “GHI”] 复制代码` * 将 collections 铺平为一个 collection。
` output: [ "A" , "B" , "C" , "D" , "E" , "F" , "G" , "H" , "I" ] 复制代码`

![](https://user-gold-cdn.xitu.io/2019/6/5/16b25b2c67bee6a9?imageView2/0/w/1280/h/960/ignore-error/1)

` 注意：在 Swift3 中，flatMap 还可以自动过滤 nil 值。但是现在已经废弃该功能。现在用 compactMap 来实现这一功能，稍后在文章中我们会讲到。`

![](https://user-gold-cdn.xitu.io/2019/6/5/16b25b52b702b52b?imageView2/0/w/1280/h/960/ignore-error/1)

现在，你应该明白了 flatMap 是做什么得了。

#### Flatmap on array ####

` let arrs = [[1,2,3], [4, 5, 6]] let flat1 = arrs.flatMap { return $0 } print (flat1) //[1, 2, 3, 4, 5, 6] 复制代码`

#### Flatmap on array of dictionaries ####

因为在铺平之后的返回数组包含元素的类型为元组。所以我们不得不转换为字典。

` let arrs = [[ "key1" : 0, "key2" : 1], [ "key3" : 3, "key4" : 4]] let flat1 = arrs.flatMap { return $0 } print (flat1) //[(key: "key2" , value: 1), (key: "key1" , value: 0), (key: "key3" , value: 3), (key: "key4" , value: 4)] var dict = [String: Int]() flat1.forEach { (key, value) in dict[key] = value } print (dict) //[ "key4" : 4, "key2" : 1, "key1" : 0, "key3" : 3] 复制代码`

#### Flatmap on set ####

![](https://user-gold-cdn.xitu.io/2019/6/5/16b25bdb9c68ce0e?imageView2/0/w/1280/h/960/ignore-error/1)

#### Flatmap by filtering or mapping ####

我们可以用 flatMap 来实现将一个二维数组铺平为一维数组。 flatMap 的闭包接受一个集合类型的参数，在闭包中我们还可以进行 filter map reduce 等操作。

` let collections = [[5, 2, 7], [4, 8], [9, 1, 3]] let onlyEven = collections.flatMap { (intArray) in intArray.filter({ $0 % 2 == 0}) } print (onlyEven) //[2, 4, 8] 复制代码`

上述代码的简化版：

` let onlyEven = collections.flatMap { $0.filter { $0 % 2 == 0 } } 复制代码`

### 链式 : (map + filter + reduce) ###

我们可以链式调用高阶函数。 ` 不要链接太多，不然执行效率会慢。下面的代码我在playground中就执行不了。`

` let arrayOfArrays = [[1, 2, 3, 4], [5, 6, 7, 8, 4]] let sumOfSquareOfEvenNums = arrayOfArrays.flatMap{ $0 }.filter{ $0 % 2 == 0}.map{ $0 * $0 }.reduce {0, +} print (sumOfSquareOfEvenNums) // 136 //这样可以运行 let SquareOfEvenNums = arrayOfArrays.flatMap{ $0 }.filter{ $0 % 2 == 0}.map{ $0 * $0 } let sum = SquareOfEvenNums.reduce(0 , +) // 136 复制代码`

### CompactMap ###

在迭代完集合中的每个元素的映射操作后，返回一个非空的数组。

` let arr = [1, nil, 3, 4, nil] let result = arr.compactMap{ $0 } print (result) //[1, 3, 4] 复制代码`

它对于 Set ,和数组是一样的作用。

` let nums: Set = [1, 2, nil] let r1 = nums.compactMap { $0 } print (r1) //[2, 1] 复制代码`

而对于 Dictionary ，它是没有任何作用的，它只会返回一个元组类型的数组。所以我们需要使用 compactMapValues 函数。(该函数在Swift5发布)

` let dict = [ "key1" : nil, "key2" : 20] let result = dict.compactMap{ $0 } print (result) //[(key: "key1" , value: nil), (key: "key2" , value: Optional(20))] let dict = [ "key1" : nil, "key2" : 20] let result = dict.compactMapValues{ $0 } print (result) 复制代码`

#### Tip ####

` let arr = [1, nil, 3, 4, nil] let result = arr.map { $0 } print (result) //[Optional(1), nil, Optional(3), Optional(4), nil] 复制代码`

使用 map 是这样的。

### 总结 ###

好了，到这里这篇文章也就接近尾声了。

我们要尽可能的使用高阶函数：

* 它可以提高你的 Swift 技能
* 代码可读性更高
* 更加函数式