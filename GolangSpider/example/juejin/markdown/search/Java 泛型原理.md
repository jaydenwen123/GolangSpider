# Java 泛型原理 #

## 泛型是什么？ ##

考虑以下场景：您希望开发一个用于在应用中传递对象的容器。但对象类型并不总是相同。因此，需要开发一个能够存储各种类型对象的容器。

鉴于这种情况，要实现此目标，显然最好的办法是开发一个能够存储和检索 Object 类型本身的容器，然后在将该对象用于各种类型时进行类型转换。

实例1中的类演示了如何开发此类容器。

` public class ObjectContainer { private Object obj; /** * @ return the obj */ public Object getObj () { return obj; } /** * @param obj the obj to set */ public void set Obj(Object obj) { this.obj = obj; } } ObjectContainer myObj = new ObjectContainer(); // store a string myObj.setObj( "Test" ); System.out.println( "Value of myObj:" + myObj.getObj()); // store an int ( which is autoboxed to an Integer object) myObj.setObj(3); System.out.println( "Value of myObj:" + myObj.getObj()); List objectList = new ArrayList(); objectList.add(myObj); // We have to cast and must cast the correct type to avoid ClassCastException! String myStr = (String) ((ObjectContainer)objectList.get(0)).getObj(); System.out.println( "myStr: " + myStr); 复制代码`
> 
> 
> 
> 
> 虽然这个容器会达到预期效果，但就我们的目的而言，它并不是最合适的解决方案。 `
> 它不是类型安全的，并且要求在检索封装对象时使用显式类型转换，因此有可能引发异常` 。
> 
> 

**` 使用泛型` 可以开发一个更好的解决方案，在实例化时为所使用的容器分配一个类型，也称泛型类型，这样就可以创建一个对象来存储所分配类型的对象。**

**` 泛型类型是一种类型参数化的类或接口` ，这意味着可以通过执行泛型类型调用 分配一个类型，将用分配的具体类型替换泛型类型。然后，所分配的类型将用于限制容器内使用的值，这样就无需进行类型转换，还可以在编译时提供更强的类型检查。**

实例2演示了如何创建与先前创建的容器相同的容器，但这次使用泛型类型参数，而不是 Object 类型。

` public class GenericContainer<T> { private T obj; public GenericContainer (){ } // Pass type in as parameter to constructor public GenericContainer(T t){ obj = t; } /** * @ return the obj */ public T getObj () { return obj; } /** * @param obj the obj to set */ public void set Obj(T t) { obj = t; } } //要使用泛型容器，必须在实例化时使用尖括号表示法指定容器类型。 //因此，以下代码将实例化一个 Integer 类型的GenericContainer，并将其分配给 myInt 字段。 GenericContainer<Integer> myInt = new GenericContainer<>(); //或者 GenericContainer<Integer> myInt = new GenericContainer<Integer>(); //如果我们尝试在已经实例化的容器中存储其他类型的对象，代码将无法编译 myInt.setObj(3); // OK myInt.setObj( "Int" ); // Won 't Compile 复制代码`
> 
> 
> 
> 
> 最显著的差异是类定义包含 ，类字段 obj 不再是 Object 类型，而是泛型类型
> T。类定义中的尖括号之间是类型参数部分，介绍类中将要使用的类型参数（或多个参数）。T 是与此类中定义的泛型类型关联的参数。
> 
> 

## 使用泛型的好处 ##

> 
> 
> 
> 一个 ` 最重要的好处是更强的类型检查` ，因为避开运行时可能引发的 ClassCastException 可以节省时间。
> 
> 

> 
> 
> 
> 另一个好处是 ` 消除了类型转换` ，这意味着可以用更少的代码，因为编译器确切知道集合中存储的是何种类型。
> 
> 

## 如何使用泛型 ##

泛型有许多不同用例。本文在前面的示例中介绍了生成泛型对象类型的用例。这对于在类和接口层面了解泛型语法是个很好的起点。

**` 类签名` 包含一个 ` 类型参数` 部分，包括在 ` 类名后的尖括号` (< >) 内**

例如：

` public class GenericContainer<T> { ... 复制代码`
> 
> 
> 
> ` 类型参数` （又称类型变量）用作占位符，指示在运行时为类分配类型。根据需要，可能有一个或多个类型参数，并且可以用于整个类。根据惯例，类型参数是单个大写字母，该字母用于指示所定义的参数类型。下面列出每个用例的标准类型参数：
> 
> 
> 
> 
> * E：元素
> * K：键
> * N：数字
> * T：类型
> * V：值
> * S、U、V 等：多参数情况中的第 2、3、4 个类型
> 
> 
> 

在上面的示例中，T 指示将分配的类型，因此可在实例化时为 GenericContainer 分配任何有效类型。注意，T 参数用于整个类，指示实例化时指定的类型。使用下面这行代码实例化对象时，将用 String 类型替换所有 T 参数：

` GenericContainer<String> stringContainer = new GenericContainer<String>(); 复制代码`

泛型也可用于构造函数中，传递类域初始化所需的类型参数。GenericContainer 的构造函数允许在实例化时传递任意类型：

` GenericContainer gc1 = new GenericContainer(3); GenericContainer gc2 = new GenericContainer( "Hello" ); 复制代码`

注意，未分配类型的泛型称为原始类型。例如，要创建原始类型的 GenericContainer，可以使用以下代码：

` GenericContainer rawContainer = new GenericContainer(); 复制代码`
> 
> 
> 
> 原始类型有时对于实现向后兼容很有用，但并不适用于日常代码。原始类型在编译时无需执行类型检查，导致代码在运行时易于出错。
> 
> 

## 多种泛型类型 ##

有时，能够在类或接口中使用多种泛型类型很有帮助。通过在尖括号之间放置一个逗号分隔的类型列表，可在类或接口中使用多个类型参数。

下面实例中的类使用一个接受以下两种类型的类演示了此概念：T 和 S。

` public class MultiGenericContainer<T, S> { private T firstPosition; private S secondPosition; public MultiGenericContainer(T firstPosition, S secondPosition){ this.firstPosition = firstPosition; this.secondPosition = secondPosition; } public T getFirstPosition (){ return firstPosition; } public void set FirstPosition(T firstPosition){ this.firstPosition = firstPosition; } public S getSecondPosition (){ return secondPosition; } public void set SecondPosition(S secondPosition){ this.secondPosition = secondPosition; } } 复制代码`
> 
> 
> 
> 
> MultiGenericContainer 类可用于存储两个不同对象，每个对象的类型可在实例化时指定。
> 
> 

容器的用法如下

` MultiGenericContainer<String, String> mondayWeather = new MultiGenericContainer<String, String>( "Monday" , "Sunny" ); MultiGenericContainer<Integer, Double> dayOfWeekDegrees = new MultiGenericContainer<Integer, Double>(1, 78.0); String mondayForecast = mondayWeather.getFirstPosition(); // The Double type is unboxed--to double, in this case. More on this in next section! double sundayDegrees = dayOfWeekDegrees.getSecondPosition(); 复制代码`

## 有界类型 ##

我们经常会遇到这种情况，需要指定泛型类型，但又希望可以控制指定的类型，而非不加限制。 ` 有界类型` 在类型参数部分指定 ` extends` 或 ` super` 关键字，分别用上限或下限限制类型，从而限制泛型类型的边界。

` 如果希望将某类型限制为特定类型或特定类型的子类型` ，请使用以下表示法：

` <T extends UpperBoundType> 复制代码`

同样，如果希望 ` 将某个类型限制为特定类型或特定类型的超类型` ，请使用以下表示法：

` <T super LowerBoundType> 复制代码`

### 什么是PECS？ ###

PECS指“Producer Extends，Consumer Super”。 如果你是想遍历collection，并对每一项元素操作时，此时这个集合是生产者（生产元素），应该使用 Collection<? extends Thing>。 如果你是想添加元素到collection中去，那么此时集合是消费者（消费元素）应该使用Collection<? super Thing>。

## 泛型方法 ##

有时，我们可能不知道传入方法的参数类型。在方法级别应用泛型可以解决此类问题。方法参数可以包含泛型类型，方法也可以包含泛型返回类型。

假设我们要开发一个接受 Number 类型的计算器类。泛型可用于确保可将任何 Number 类型作为参数传递给此类的计算方法。

例如，如下示例中的 add() 方法演示了如何使用泛型限制两个参数的类型，确保其包含 Number 的上限：

` public static <N extends Number> double add(N a, N b){ double sum = 0; sum = a.doubleValue() + b.doubleValue(); return sum; } 复制代码`

通过将类型限制为 Number，您可以将 Number 子类的任何对象作为参数传递。此外，通过将类型限制为 Number，我们还可以确保传递给该方法的任何参数将包含 doubleValue() 方法。要查看实际效果，如果您想添加一个 Integer 和一个 Float，可以按如下所示调用该方法：

` double genericValue1 = Calculator.add(3, 3f); 复制代码`

## 通配符 ##

某些情况下，编写 ` 指定未知类型` 的代码很有用。问号 ` ?` 通配符可用于使用泛型代码 ` 表示未知类型` 。通配符可用于参数、字段、局部变量和返回类型。但最好不要在返回类型中使用通配符，因为确切知道方法返回的类型更安全。

假设我们想编写一个方法来验证指定的 List 中是否存在指定的对象。我们希望该方法接受两个参数：一个是未知类型的 List，另一个是任意类型的对象。

` public static <T> void checkList(List<?> myList, T obj){ if (myList.contains(obj)){ System.out.println( "The list contains the element: " + obj); } else { System.out.println( "The list does not contain the element: " + obj); } } 复制代码`

使用示例

` // Create List of type Integer List<Integer> intList = new ArrayList<Integer>(); intList.add(2); intList.add(4); intList.add(6); // Create List of type String List<String> strList = new ArrayList<String>(); strList.add( "two" ); strList.add( "four" ); strList.add( "six" ); // Create List of type Object List<Object> objList = new ArrayList<Object>(); objList.add( "two" ); objList.add( "four" ); objList.add(strList); checkList(intList, 3); // Output: The list [2, 4, 6] does not contain the element: 3 checkList(objList, strList); /* Output: The list [two, four, [two, four, six]] contains the element: [two, four, six] */ checkList(strList, objList); /* Output: The list [two, four, six] does not contain the element: [two, four, [two, four, six]] */ 复制代码`
> 
> 
> 
> 
> 有时要使用上限或下限限制通配符。与指定带边界的泛型类型极其相似，指定 extends 或 super
> 关键字加上通配符，后面跟用于上限或下限的类型，即可声明带边界的通配符类型。
> 
> 

例如，如果我们要更改 checkList 方法使其只接受扩展 Number 类型的 List，可按清单 14 所示编写代码。

` public static <T> void checkNumber(List<? extends Number> myList, T obj){ if (myList.contains(obj)){ System.out.println( "The list " + myList + " contains the element: " + obj); } else { System.out.println( "The list " + myList + " does not contain the element: " + obj); } } 复制代码`

## 总结 ##

` 泛型` 其实说白了就是 ` 应用在编译时期` ， ` 是给编译器使用的技术` ，到了运行时期，泛型就不存在啦。这是因为，编辑器检查了泛型的类型正确之后，再生成的类文件中是没有泛型的。

## 泛型使用注意事项 ##

* 对象实例化时不指定泛型的话，默认为：Object。
* ` 泛型的指定中不能使用基本数据类型，可以使用包装类替换` 。
* **` 静态方法中不能使用类的泛型`**
* **可以同时绑定多个绑定,用 ` &` 连接**
* 泛型类可能多个参数，此时应将多个参数一起放在尖括号内。比如<E1,E2,E3>
* 从泛型类派生子类，泛型类型需具体化：继承泛型类后，子类类型对应类型需要具体化
* 如果泛型类是一个接口或抽象类，则不可创建泛型类的对象。