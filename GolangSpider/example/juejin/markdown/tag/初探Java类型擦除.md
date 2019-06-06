# 初探Java类型擦除 #

本篇博客主要介绍了Java类型擦除的定义，详细的介绍了类型擦除在Java中所出现的场景。

## 1. 什么是类型擦除 ##

为了让你们快速的对类型擦除有一个印象，首先举一个很简单也很经典的例子。

` // 指定泛型为String List<String> list1 = new ArrayList<>(); // 指定泛型为Integer List<Integer> list2 = new ArrayList<>(); System.out.println(list1.getClass() == list2.getClass()); // true 复制代码`

上面的判断结果是 ` true` 。代表了两个传入了不同泛型的List最终都编译成了ArrayList，成为了同一种类型，原来的泛型参数String和Integer被擦除掉了。这就是类型擦除的一个典型的例子。

而如果我们说到类型擦除为什么会出现，我们就必须要了解泛型。

## 2. 泛型 ##

### 2.1. 泛型的定义 ###

随着2004年9月30日，工程代号为Tiger的JDK 1.5发布，泛型从此与大家见面。JDK 1.5在Java语法的易用性上作出了非常大的改进。除了泛型，同版本加入的还有自动装箱、动态注解、枚举、可变长参数、foreach循环等等。

而在1.5之前的版本中，为了让Java的类具有通用性，参数类型和返回类型通常都设置为Object，可见，如果需要不用的类型，就需要在相应的地方，对其进行强制转换，程序才可以正常运行，十分麻烦，稍不注意就会出错。

泛型的本质就是参数化类型。也就是，将一个数据类型指定为参数。引入泛型有什么好处呢？

泛型可以将JDK 1.5之前在运行时才能发现的错误，提前到编译期。也就是说，泛型提供了编译时类型安全的检测机制。例如，一个变量本来是Integer类型，我们在代码中设置成了String，没有使用泛型的时候只有在代码运行到这了，才会报错。

而引入泛型之后就不会出现这个问题。这是因为通过泛型可以知道该参数的规定类型，然后在编译时，判断其类型是否符合规定类型。

泛型总共有三种使用方法，分别使用于类、方法和接口。

## 3. 泛型的使用方法 ##

### 3.1 泛型类 ###

#### 3.1.1 定义泛型类 ####

简单的泛型类可以定义为如下。

` public class Generic < T > { T data; public Generic (T data) { setData(data); } public T getData () { return data; } public void setData (T data) { this.data = data; } } 复制代码`

其中的T代表参数类型，代表任何类型。当然，并不是一定要写成T，这只是大家约定俗成的习惯而已。有了上述的泛型类之后我们就可以像如下的方式使用了。

#### 3.1.2 使用泛型类 ####

` // 假设有这样一个具体的类 public class Hello { private Integer id; private String name; private Integer age; private String email; } // 使用泛型类 Hello hello = new Hello(); Generic<Hello> result = new Generic<>(); resule.setData(hello); // 通过泛型类获取数据 Hello data = result.getData(); 复制代码`

当然如果泛型类不传入指定的类型的话，泛型类中的方法或者成员变量定义的类型可以为任意类型，如果打印 ` result.getClass()` 的话，会得到 ` Generic` 。

### 3.2. 泛型方法 ###

#### 3.2.1 定义泛型方法 ####

首先我们看一下不带返回值的泛型方法，可以定义为如下结构。

` // 定义不带返回值的泛型方法 public <T> void genericMethod (T field) { System.out.println(field.getClass().toString()); } // 定义带返回值的泛型方法 private <T> T genericWithReturnMethod (T field) { System.out.println(field.getClass().toString()); return field; } 复制代码`

#### 3.2.2 调用泛型方法 ####

` // 调用不带返回值泛型方法 genericMethod( "This is string" ); // class java.lang.String genericMethod( 56L ); // class java.lang.Long // 调用带返回值的泛型方法 String test = genericWithReturnMethod( "TEST" ); // TEST class java.lang.String 复制代码`

带返回值的方法中，T就是当前函数的返回类型。

### 3.3. 泛型接口 ###

泛型接口定义如下

` public interface genericInterface < T > { } 复制代码`

使用的方法与泛型类类似，这里就不再赘述。

## 4. 泛型通配符 ##

什么是泛型通配符？官方一点的解释是

> 
> 
> 
> Type of unknown.
> 
> 

也就是无限定的通配符，可以代表任意类型。用法也有三种，<?>，<? extends T>和<? super T>。

既然已经有了T这样的代表任意类型的通配符，为什么还需要这样一个无限定的通配符呢？是因为其主要解决的问题是泛型继承带来的问题。

### 4.1. 泛型的继承问题 ###

首先来看一个例子

` List<Integer> integerList = new ArrayList<>(); List<Number> numberList = integerList; 复制代码`

我们知道， ` Integer` 是继承自 ` Number` 类的。

> 
> 
> 
> public final class Integer extends Number implements Comparable { .... }
> 
> 

那么上述的代码能够通过编译吗？肯定是不行的。Integer继承自Number不代表List 和 List之间有继承关系。那通配符的应用场景是什么呢？

### 4.2. 通配符的应用场景 ###

在其他函数中，例如JavaScript中，一个函数的参数可以是任意的类型，而不需要进行任意的类型转换，所以这样的函数在某些应用场景下，就会具有很强的通用性。

而在Java这种强类型语言中，一个函数的参数类型是固定不变的。那如果想要在Java中实现类似于JavaScript那样的通用函数该怎么办呢？这也就是为什么我们需要泛型的通配符。

假设我们有很多动物的类, 例如Dog, Pig和Cat三个类，我们需要有一个通用的函数来计算动物列表中的所有动物的腿的总数，如果在Java中，要怎么做呢？

可能会有人说，用泛型啊，泛型不就是解决这个问题的吗？泛型必须指定一个特定的类型。正式因为泛型解决不了...才提出了泛型的通配符。

### 4.3. 无界通配符 ###

无界通配符就是 ` ?` 。看到这你可能会问，这不是跟T一样吗？为啥还要搞个 ` ?` 。他们主要区别在于，T主要用于声明一个泛型类或者方法，?主要用于使用泛型类和泛型方法。下面举个简单的例子。

` // 定义打印任何类型列表的函数 public static void printList (List<?> list) { for (Object elem: list) { System.out.print(elem + " " ); } } // 调用上述函数 List<Integer> intList = Arrays.asList( 1 , 2 , 3 ); List<String> stringList = Arrays.asList( "one" , "two" , "three" ); printList(li); // 1 2 3 printList(ls); // one two three 复制代码`

上述函数的目的是打印任何类型的列表。可以看到在函数内部，并没有关心List中的泛型到底是什么类型的，你可以将<?>理解为只提供了一个只读的功能，它去除了增加具体元素的能力，只保留与具体类型无关的功能。从上述的例子可以看出，它只关心元素的数量以及其是否为空，除此之外不关心任何事。

再反观T，上面我们也列举了如何定义泛型的方法以及如果调用泛型方法。泛型方法内部是要去关心具体类型的，而不仅仅是数量和不为空这么简单。

### 4.4. 上界通配符<? extends T> ###

既然 ` ?` 可以代表任何类型，那么extends又是干嘛的呢？

假设有这样一个需求，我们只允许某一些特定的类型可以调用我们的函数（例如，所有的Animal类以及其派生类），但是目前使用 ` ?` ，所有的类型都可以调用函数，无法满足我们的需求。

` private int countLength (List< ? extends Animal> list) {...} 复制代码`

使用了上界通配符来完成这个公共函数之后，就可以使用如下的方式来调用它了。

` List<Pig> pigs = new ArrayList<>(); List<Dog> dogs = new ArrayList<>(); List<Cat> cats = new ArrayList<>(); // 假装写入了数据 int sum = 0 ; sum += countLength(pigs); sum += countLength(dogs); sum += countLength(cats); 复制代码`

看完了例子，我们就可以简单的得出一个结论。上界通配符就是一个可以处理任何特定类型以及是该特定类型的派生类的通配符。

可能会有人看的有点懵逼，我结合上面的例子，再简单的用人话解释一下：上界通配符就是一个啥动物都能放的盒子。

### 4.5. 下界通配符<? super Animal> ###

上面我们聊了上界通配符，它将未知的类型限制为特定类型或者该特定的类型的子类型（也就是上面讨论过的动物以及一切动物的子类）。而下界通配符则将未知的类型限制为特定类型或者该特定的类型的超类型，也就是超类或者基类。

在上述的上界通配符中，我们举了一个例子。写了一个可以处理任何动物类以及是动物类的派生类的函数。而现在我们要写一个函数，用来处理任何是Integer以及是Integer的超类的函数。

` public static void addNumbers (List<? super Integer> list) { for ( int i = 1 ; i <= 10 ; i++) { list.add(i); } } 复制代码`

## 5. 类型擦除 ##

简单的了解了泛型的几种简单的使用方法之后，我们回到本篇博客的主题上来——类型擦除。泛型虽然有上述所列出的一些好处，但是泛型的生命周期只限于编译阶段。

本文最开始的给出的样例就是一个典型的例子。在经过编译之后会采取去泛型化的措施，编译的过程中，在检测了泛型的结果之后会将泛型的相关信息进行擦除操作。就像文章最开始提到的例子一样，我们使用上面定义好的Generic泛型类来举个简单的例子。

` Generic<String> generic = new Generic<>( "Hello" ); Field[] fs = generic.getClass().getDeclaredFields(); for (Field f : fs) { System.out.println( "type: " + f.getType().getName()); // type: java.lang.Object } 复制代码`

` getDeclaredFields` 是反射中的方法，可以获取当前类已经声明的各种字段，包括public，protected以及private。

可以看到我们传入的泛型String已经被擦除了，取而代之的是Object。那之前的String和Integer的泛型信息去哪儿了呢？可能这个时候你会灵光一闪，那是不是所有的泛型在被擦除之后都会变成Object呢？别着急，继续往下看。

当我们在泛型上面使用了上界通配符以后，会有什么情况发生呢？我们将Generic类改成如下形式。

` public class Generic < T extends String > { T data; public Generic (T data) { setData(data); } public T getData () { return data; } public void setData (T data) { this.data = data; } } 复制代码`

然后再次使用反射来查看泛型擦除之后类型。这次控制台会输出 ` type: java.lang.String` 。可以看到，如果我们给泛型类制定了上限，泛型擦除之后就会被替换成类型的上限。而如果没有指定，就会统一的被替换成Object。相应的，泛型类中定义的方法的类型也是如此。

## 6. 写在最后 ##

如果各位发现文章中有问题的，欢迎大家不吝赐教，我会及时的更正。

> 
> 
> 
> 参考：
> 
> * [Java语言类型擦除](
> https://link.juejin.im?target=https%3A%2F%2Fwww.ibm.com%2Fdeveloperworks%2Fcn%2Fjava%2Fjava-language-type-erasure%2Findex.html
> )
> * [下界通配符](
> https://link.juejin.im?target=https%3A%2F%2Fwww.zhihu.com%2Fquestion%2F20400700
> )
> * [List<?>和List的区别](
> https://link.juejin.im?target=https%3A%2F%2Fwww.zhihu.com%2Fquestion%2F31429113
> )
>