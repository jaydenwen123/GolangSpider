# java23种设计模式-门面模式（外观模式） #

## title: java-23种设计模式-门面模式（外观模式） catalog: true date: 2018-11-28 15:07:16 subtitle: header-img: tags: ##

### 1 介绍 ###

外观模式（Facade）,他隐藏了系统的复杂性，并向客户端提供了一个可以访问系统的接口。这种类型的设计模式属于结构性模式。为子系统中的一组接口提供了一个统一的访问接口，这个接口使得子系统更容易被访问或者使用。

### 2 角色和使用场景 ###

![](https://user-gold-cdn.xitu.io/2019/6/4/16b21aee451f2586?imageView2/0/w/1280/h/960/ignore-error/1)

简单来说，该模式就是把一些复杂的流程封装成一个接口供给外部用户更简单的使用。这个模式中，设计到3个角色。

* 

门面角色：外观模式的核心。它被客户角色调用，它熟悉子系统的功能。内部根据客户角色的需求预定了几种功能的组合。

* 

子系统角色:实现了子系统的功能。它对客户角色和Facade时未知的。它内部可以有系统内的相互交互，也可以由供外界调用的接口。

* 

客户角色:通过调用Facede来完成要实现的功能。

#### 2.1 使用场景 ####

* 

为复杂的模块或子系统提供外界访问的模块；

* 

子系统相互独立；

* 

在层析结构中，可以使用外观模式定义系统的每一层的入口。

### 3 代码实例 ###

#### 3.1 原始状态 ####

一写信为例，比如给女朋友写情书什么的，写信 的过程大家都还记得吧，先写信的内容，然后写信封，然后把信放到信封中，封好，投递到信箱中进行邮 递，这个过程还是比较简单的，虽然简单，这四个步骤都是要跑的呀，信多了还是麻烦，比如到了情人节， 为了大海捞针，给十个女孩子发情书，都要这样跑一遍，你不要累死，更别说你要发个广告信啥的，一下 子发 1 千万封邮件，那不就完蛋了？那怎么办呢？还好，现在邮局开发了一个新业务，你只要把信件的必 要信息告诉我，我给你发，我来做这四个过程，你就不要管了，只要把信件交给我就成了。

类图：

![](https://user-gold-cdn.xitu.io/2019/6/4/16b21af1b5d48ed6?imageView2/0/w/1280/h/960/ignore-error/1)

先看写信的过程接口，定义了写信的四个步骤：

` public interface LetterProcess { //定义一个写信的过程 //首先要写信的内容 public void writeContext(String context); //其次写信封 public void fillEnvelope(String address); //把信放到信封里 public void letterInotoEnvelope(); //然后邮递 public void sendLetter(); } 复制代码`

写信过程的具体实现：

` public class LetterProcessImpl implements LetterProcess{ //写信 @Override public void writeContext(String context) { System.out.println( "填写信的内容...." + context); } //在信封上填写必要的信息 @Override public void fillEnvelope(String address) { System.out.println( "填写收件人地址及姓名...." + address); } //把信放到信封中，并封好 @Override public void letterInotoEnvelope () { System.out.println( "把信放到信封中...." ); } //塞到邮箱中，邮递 @Override public void sendLetter () { System.out.println( "邮递信件..." ); } } 复制代码`

然后就有人开始用这个过程写信了：

` public interface LetterProcess { //定义一个写信的过程 //首先要写信的内容 public void writeContext(String context); //其次写信封 public void fillEnvelope(String address); //把信放到信封里 public void letterInotoEnvelope(); //然后邮递 public void sendLetter(); } 复制代码`

#### 3.2 使用门面模式 ####

那这个过程与高内聚的要求相差甚远，你想，你要知道这四个步骤，而且还要知道这四个步骤的顺序， 一旦出错，信就不可能邮寄出去，那我们如何来改进呢？先看类图：

![](https://user-gold-cdn.xitu.io/2019/6/4/16b21af4747ba9bf?imageView2/0/w/1280/h/960/ignore-error/1)

这就是门面模式，还是比较简单的，Sub System(子系统) 比较复杂，为了让调用者更方便的调用，就对 Sub System 进行了封装，增加了一个门面，Client 调用时，直接调用门面的方法就可以了，不用了解具体的实现方法 以及相关的业务顺序，我们来看程序的改变，LetterProcess 接口和实现类都没有改变，只是增加了一个 ModenPostOffice 类，我们这个 java 程序清单如下：

` public class ModenPostOffice { private LetterProcess letterProcess = new LetterProcessImpl(); //写信，封装，投递，一体化了 public void sendLetter(String context, String address) { //帮你写信 letterProcess.writeContext(context); //写好信封 letterProcess.fillEnvelope(address); //把信放到信封中 letterProcess.letterInotoEnvelope(); //邮递信件 letterProcess.sendLetter(); } } 复制代码`

个类是什么意思呢，就是说现在有一个提供了一种新型的 服务，客户只要把信的内容以及收信地址给他们，他们就会把信写好，封好，并发送出去，这种服务提出 时大受欢迎呀，这简单呀，客户减少了很多工作，那我们看看客户是怎么调用的，Client.java 的程序清单 如下：

` public class Client { public static void main(String[] args) { //现代化的邮局，有这项服务 ModenPostOffice hellRoadPostOffice = new ModenPostOffice(); //你只要把信的内容和收信人地址给他，他会帮你完成一系列的工作； String address = "Happy Road No. 666,God Province,Heaven" ; //定义一个地址 String context = "Hello,It's me,do you know who I am? I'm your old lover." + "I'd like to...." ; hellRoadPostOffice.sendLetter(context, address); } } 复制代码`

### 4 优点 ###

* 松散耦合：使得客户端和子系统之间解耦，让子系统内部的模块功能更容易扩展和维护。
* 简单易用：客户端根本不需要知道子系统内部的实现，或者根本不需要知道子系统内部的构成，它只需要跟Facade类交互即可。
* 更好的划分访问层次：有些方法是对系统外的，有些方法是系统内部相互交互的使用的。子系统把那些暴露给外部的功能集中到门面中，这样就可以实现客户端的使用，很好的隐藏了子系统内部的细节。