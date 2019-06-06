# Kotlin艺术探索之单例和伴生对象 #

### 单例 ###

Kotlin中的单例写法相比较于Java要简单许多，只需用到一个关键字就可以实现，那就是object关键字

比如下面DataProviderManager就是一个单例类

` object DataProviderManager { fun registerDataProvider(provider: DataProvider) { // ... } val allDataProviders: Collection<DataProvider> get() = // ... } 复制代码`

调用单例类中的方法也很简单

` DataProviderManager.registerDataProvider(...) 复制代码`

调用格式很像Java的静态类调用它的静态方法。那么Kotlin中的静态类和静态方法是不是和Java一样呢？

下面就来说Kotlin中的静态实现

### 伴生对象 ###

相信你看到这个标题，也猜到Kotlin的静态实现和Java不一样了，Java的静态实现需要用到Static关键字，但是Kotlin不是这样的呢，它用的是 companion object，翻译过来就是伴生对象

举个例子

` class MyClass { companion object{ fun create(): MyClass = MyClass() var a = 1 } } 复制代码`

可以看到有一个companion object代码块，在这里面可以编写方法和属性，那如何调用这个类的create()方法呢？

` val instance = MyClass.create() 复制代码`

既然是和静态方法一样，就可以直接类名.方法了

> 
> 
> 
> 注意:
> 
> 
> 
> * 每个类可以对应一个伴生对象
> * 伴生对象的成员全局只有一个
> * 伴生对象的成员类似于java的静态成员
> 
> 
>