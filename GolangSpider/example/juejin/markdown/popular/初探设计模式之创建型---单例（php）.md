# 初探设计模式之创建型---单例（php） #

## 序言 ##

首先，设计模式出现在眼里太多次了，看了又看，忘了又忘，咋看咋忘，在做前端开发之时就看过相关型文章，水平有限，业务有限，看过的知识好比过往云烟，匆匆而已，不留下一点痕迹，如今作为一个php，我决定深深的探一下设计模式，对于项目的开发以及代码的规范有着大大的好处。

## 为什么先说单例 ##

1、单例是设计模式里最基础的一种
2、我用过单例
3、我不是什么大神，只是一个初转后端的小白。

## 正文 ##

### 单例是什么 ###

` 1、创建型设计模式的一种。 2、字面意思（单个例子，单个实例）。 复制代码`

### 单例的目的 ###

` 1、在特定的场景里，特定的类只能有一个实例。 2、节约内存。 复制代码`

### 单例的条件 ###

` 1、自行实例化。 2、只有一个实例。 3、向全局暴露实例。 复制代码`

### 粗暴理解单例 ###

* 

首先先模拟一个故事。

小时候最头疼写作业，语文作业，英语作业等等，非常多，然后写作的时候需要用到笔，一般我是直接用一个笔写完作业，不可能是写语文作业用一支笔，写英语作业用另一只笔。然后以这个故事用代码来描述一下。

* 

代码部分

1、先抽象一个笔类

` class Pen { public function __construct () { echo "制作一支笔" ; } public function write () { echo "我可以写字" ; } } 复制代码`

这个类叫做pen，有一个方法叫写，通俗的讲，抽象一个笔，这个笔有一个功能叫做写字。

2、 然后使用它

` class Xiaochen { public function writeEnglish () : Pen { $pen = new Pen(); $pen->write(); echo "写英语<br>" ; return $pen; } public function writeChinese () : Pen { $pen = new Pen(); $pen->write(); echo "写英语<br>" ; return $pen; } } 复制代码`

抽象一个人，叫小陈，小陈今天要写语文和数学作业(writeEnglish,writeChinese)，然后小陈写作业的时候肯定要用到笔($pen=new Pen)。这样就可以完成了使用。 结果是这样:

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2ba06a196923e?imageView2/0/w/1280/h/960/ignore-error/1)

3、然后我们来判断一下写语文和写作业是不是用的一支笔

` var_dump($writeChinesePen === $writeChinesePen); 答案是false。 复制代码`

因为虽然是统一的pen类实例出来的对象，但是两个对象却是不同的存在，至于为什么，不懂的可以去补一下基础。

4、修改pen类

` class Pen { // 定义静态，全局可访问的实例对象 private static $pen; // 暴露一个实例化的方法，只能通过这个方法实例化 public static function getInstance () { if (is_null( static ::$pen)) { static ::$pen = new static ; } return static ::$pen; } public function write () { echo "写字" ; } } 复制代码`

1、首先预定义一个全局的属性用来存储实例对象，使用static，因为不必实例化就可以使用。

2、制作一个方法，代替实例化，判断，如果实例过，直接返回实例对象，else，进行实例化。

3、如果有人不知道，使用new操作符来实例这个类怎么办。(不懂的请观看访问控制相关的文档)

答案是修改构造函数:

` private function __construct () { } 复制代码`

4、由于类是继承，那么单例也就代表着单例类不可继承，并且不可复制。

修改php内置的clone:

` private function __clone () { } 复制代码`

修改类不可继承 **final** class Pen

这样的话，这个笔类大致就完成了。不可通过new操作符进行实例，也不会创建多个实例。因为是static，所以全局都可以使用。

5、最终测试

` final class Pen { // 定义静态，全局可访问的实例对象 private static $pen; // 暴露一个实例化的方法，只能通过这个方法实例化 public static function getInstance () { if (is_null( static ::$pen)) { static ::$pen = new static ; } return static ::$pen; } public function write () { echo "写字" ; } private function __construct () { } private function __clone () { } } class Xiaochen { public function writeEnglish () : Pen { $pen = Pen::getInstance(); $pen->write(); echo "写英语<br>" ; return $pen; } public function writeChinese () : Pen { $pen = Pen::getInstance(); $pen->write(); echo "写数学<br>" ; return $pen; } } $xiaochen = new Xiaochen(); $writeEnglistPen = $xiaochen->writeEnglish(); $writeChinesePen = $xiaochen->writeChinese(); var_dump($writeChinesePen === $writeEnglistPen); 复制代码`

这样就可以保证全局只有一个实例，不会创建多个实例，只要实例过一次，就不会在创建第二个实例。

### 总结 ###

以单例开始设计模式之门，以上例子纯属个人编造，可能有写不恰当，大神勿喷，单例在应用种我一般都用来初始化数据库连接，日志对象，以及需要设定全局只有一个实例的需求。