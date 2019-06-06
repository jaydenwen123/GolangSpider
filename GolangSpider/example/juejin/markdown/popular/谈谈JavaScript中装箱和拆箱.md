# 谈谈JavaScript中装箱和拆箱 #

在 ` JavaScript` 里面有个引用类型叫做 **基本包装类型** ，它包括 ` String、Number和Boolean` 。那么它和基本的类型 ` String、Number和Boolean` 是啥关系呢？接着往下看👀

### 装箱操作 ###

所谓的 **装箱** ，是指将基本数据类型转换为对应的引用类型的操作。而装箱又分为 ` 隐式装箱和显式装箱` 。

#### 隐式装箱 ####

对于隐式装箱，我们看下下面的代码：

` var s1 = 'call_me_R' ; // 隐式装箱 var s2 = s1.substring( 2 ); 复制代码`

上面代码的执行步骤其实是这样的：

* 创建String类型的一个实例；
* 在实例中调用制定的方法；
* 销毁这个实例。

上面的三个步骤转换为代码，如下：

` # 1 var s1 = new String('call_me_R'); # 2 var s2 = s1.substring(2); # 3 s1 = null; 复制代码`

**隐式装箱** 当读取一个基本类型值时，后台会创建一个该基本类型所对应的 ` 基本包装类型` 对象。在这个基本类型的对象上调用方法，其实就是在这个基本类型对象上调用方法。这个基本类型的对象是临时的，它只存在于方法调用那一行代码执行的瞬间，执行方法后立即被销毁。这也是在基本类型上添加属性和方法会不识别或报错的原因了，如下：

` var s1 = 'call_me_R' ; s1.job = 'frontend engineer' ; s1.sayHello = function ( ) { console.log( 'hello kitty' ); } console.log(s1.job); // undefined s1.sayHello(); // Uncaught TypeError: s1.sayHello is not a function 复制代码`

#### 显示装箱 ####

装箱的另一种方式是 **显示装箱** ，这个就比较好理解了，这是通过 ` 基本包装类型` 对象对基本类型进行显示装箱，如下：

` var name = new String ( 'call_me_R' ); 复制代码`

显示装箱的操纵可以对 ` new` 出来的对象进行属性和方法的添加啦，因为通过 ` 通过new操作符创建的引用类型的实例，在执行流离开当前作用域之前一直保留在内存中` 。

` var objStr = new String ( 'call_me_R' ); objStr.job = 'frontend engineer' ; objStr.sayHi = function ( ) { console.log( 'hello kitty' ); } console.log(objStr.job); // frontend engineer objStr.sayHi(); // hello kitty 复制代码`

### 拆箱操作 ###

拆箱就和装箱相反了。 **拆箱** 是指把引用类型转换成基本的数据类型。通常通过引用类型的 ` valueOf()和toString()` 方法来实现。

在下面的代码中，留意下 ` valueOf()和toString()` 返回值的区别：

` var objNum = new Number(64); var objStr = new String('64'); console.log(typeof objNum); // object console.log(typeof objStr); // object # 拆箱 console.log(typeof objNum.valueOf()); // number 基本的数字类型，想要的 console.log(typeof objNum.toString()); // string 基本的字符类型，不想要的 console.log(typeof objStr.valueOf()); // string 基本的数据类型，不想要的 console.log(typeof objStr.toString()); // string 基本的数据类型，想要的 复制代码`

所以，在进行拆箱操作的过程中，还得结合下实际的情况进行拆箱，别盲目来 -- 吃力不讨好就很尴尬了😅

### 后话 ###

文章首发： [github.com/reng99/blog…]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Freng99%2Fblogs%2Fissues%2F28 )

更多内容： [github.com/reng99/blog…]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Freng99%2Fblogs )

### 参考 ###

[JavaScript 基本类型的装箱与拆箱]( https://link.juejin.im?target=https%3A%2F%2Fwww.jianshu.com%2Fp%2Fd66cf6f711a1 )

《JavaScript高级程序设计》