# 【肥朝】如何手写实现简易的Dubbo？ #

## 前言 ##

结束了 ` 集群容错` 和 ` 服务发布原理` 这两个小专题之后,有朋友问我 ` 服务引用` 什么时候开始,本篇为 ` 服务引用` 的启蒙篇.之前是一直和大家一起看源码,鉴于 ` Talk is cheap.Show me your code` ,所以本篇将和大家一起写写代码.

## 插播面试题 ##

* 

dubbo的原理是怎么样的?请简单谈谈

* 

有没有考虑过自己实现一个类似dubbo的RPC框架,如果有,请问你会如果着手实现? ` (面试高频题,区分度高)`

* 

你说你用过mybatis,那你知道Mapper接口的原理吗?(如果回答得不错,并且提到动态代理这个关键词会继续往下问,那这个动态代理又是如何通过依赖注入到Mapper接口的呢?)

## 直入主题 ##

### 简单原理 ###

谈到dubbo的原理,我们就必须首先要知道,dubbo的基本概念,通俗的说,就是dubbo是干嘛的

> 
> 
> 
> dubbo是一个分布式服务框架，致力于提供高性能和透明化的RPC远程服务调用方案，以及SOA服务治理方案
> 
> 

在此之前,就必须要讲讲以下几个简单又容易混淆的概念

* 集群

> 
> 
> 
> 同一个业务，部署在多个服务器上
> 
> 

* 分布式

> 
> 
> 
> 一个业务分拆多个子业务，部署在不同的服务器上
> 
> 

* RPC

> 
> 
> 
> RPC（Remote Procedure Call Protocol）---远程过程调用
> 
> 

我们捕捉到几个重要的关键词, ` 分布式` , ` 透明化` , ` RPC`.

既然各服务是部署在不同的服务器上,那服务间的调用就是要通过网络通信,简单的用图描述如下:

![](https://user-gold-cdn.xitu.io/2019/5/6/16a8ac2cb3f54821?imageView2/0/w/1280/h/960/ignore-error/1)

之前在[dubbo源码解析-本地暴露]的时候就有很多朋友留言问到,这个本地暴露有什么用.首先,dubbo作为一个被广泛运用的框架,点滴的性能提升,那么受益者都是很大一个数量.这也就是为什么JDK的源码,都喜欢用 ` 位运算`.比如图中的 ` UserService` 和 ` RoleService` 服务是在同一模块内的,他们直接的通信通过JVM性能肯定要比通过网络通信要好得多.这就是为什么dubbo在设计上,既有 ` 远程暴露` ,又有 ` 本地暴露` 的原因.

既然涉及到了网络通信,那么服务消费者调用服务之前,都要写各种网络请求,编解码之类的相关代码,明显是很不友好的.dubbo所说的 ` 透明` ,就是指,让调用者对网络请求,编解码之类的细节透明,让我们像调用本地服务一样调用远程服务,甚至感觉不到自己在调用远程服务.

说了这么多,那到底怎么做?要实现这个需求,我们很容易想到一个关键词,那就是动态代理

` public interface MenuService { void sayHello () ; } 复制代码` ` public class MenuServiceImpl implements MenuService { @Override public void sayHello () { } } 复制代码` ` public class ProxyFactory implements InvocationHandler { private Class interfaceClass; public ProxyFactory (Class interfaceClass) { this.interfaceClass = interfaceClass; } //返回代理对象,此处用泛型为了调用时不用强转,用Object需要强转 public <T> T getProxyObject () { return (T) Proxy.newProxyInstance( this.getClass().getClassLoader(), //类加载器 new Class[]{interfaceClass}, //为哪些接口做代理(拦截哪些方法) this ); //(把这些方法拦截到哪处理) } @Override public Object invoke (Object proxy, Method method, Object[] args) throws Throwable { System.out.println(method); System.out.println( "进行编码" ); System.out.println( "发送网络请求" ); System.out.println( "将网络请求结果进行解码并返回" ); return null ; } } 复制代码` ` public void test () throws Exception { ProxyFactory proxyFactor = new ProxyFactory(MenuService.class); MenuService menuService = proxyFactor.getProxyObject(); menuService.sayHello(); //输出结果如下: //public abstract void com.toby.rpc.MenuService.sayHello() //进行编码 //发送网络请求 //将网络请求结果进行解码并返回 } 复制代码`

看到这里可能有朋友要吐槽了,我都看了你几个月的源码解析了,上面说的那些我早就懂了,那我还关注 ` 肥朝公众号` 干嘛.我要的是整出一个类似dubbo的框架,性能上差点没关系,至少外观使用上要差不多,比如我们平时使用dubbo都是先在配置文件配置这么个东西

` < dubbo:reference id = "demoService" interface = "com.alibaba.dubbo.demo.DemoService" /> 复制代码`

然后用在用 ` @Autowired` 依赖注入来使用的,说白了,逼格要有!

![](https://user-gold-cdn.xitu.io/2019/5/6/16a8ac2fd5eab028?imageView2/0/w/1280/h/960/ignore-error/1)

### 与spring融合 ###

我们假如要写一个简单的RPC,就取名叫 ` tobyRPC` (肥朝英文名为toby),其实我个人是比较喜欢截图,但是部分朋友和我反复强调贴代码,那这里我就贴代码吧

1.设计配置属性和JavaBean

` public class ReferenceBean < T > extends ReferenceConfig < T > implements FactoryBean { @Override public Object getObject () throws Exception { return get(); } @Override public Class<?> getObjectType() { return getInterfaceClass(); } @Override public boolean isSingleton () { return true ; } } 复制代码` ` public class ReferenceConfig < T > { private Class<?> interfaceClass; // 接口代理类引用 private transient volatile T ref; public synchronized T get () { if (ref == null ) { init(); } return ref; } private void init () { ref = new ProxyFactory(interfaceClass).getProxyObject(); } public Class<?> getInterfaceClass() { return interfaceClass; } public void setInterfaceClass (Class<?> interfaceClass) { this.interfaceClass = interfaceClass; } } 复制代码`

2.编写XSD文件

` <? xml version= "1.0" encoding= "UTF-8" standalone= "no" ?> < xsd:schema xmlns = "http://toby.com/schema/tobyRPC" xmlns:xsd = "http://www.w3.org/2001/XMLSchema" xmlns:beans = "http://www.springframework.org/schema/beans" xmlns:tool = "http://www.springframework.org/schema/tool" targetNamespace = "http://toby.com/schema/tobyRPC" > < xsd:import namespace = "http://www.w3.org/XML/1998/namespace" /> < xsd:import namespace = "http://www.springframework.org/schema/beans" /> < xsd:import namespace = "http://www.springframework.org/schema/tool" /> < xsd:complexType name = "referenceType" > < xsd:complexContent > < xsd:extension base = "beans:identifiedType" > < xsd:attribute name = "interface" type = "xsd:token" use = "required" > < xsd:annotation > < xsd:documentation > <![CDATA[ The service interface class name. ]]> </ xsd:documentation > < xsd:appinfo > < tool:annotation > < tool:expected-type type = "java.lang.Class" /> </ tool:annotation > </ xsd:appinfo > </ xsd:annotation > </ xsd:attribute > </ xsd:extension > </ xsd:complexContent > </ xsd:complexType > < xsd:element name = "reference" type = "referenceType" > < xsd:annotation > < xsd:documentation > <![CDATA[ Reference service config ]]> </ xsd:documentation > </ xsd:annotation > </ xsd:element > </ xsd:schema > 复制代码`

3.编写 ` NamespaceHandler` 和 ` BeanDefinitionParser` 完成解析工作

` public class TobyRPCBeanDefinitionParser extends AbstractSingleBeanDefinitionParser { protected Class getBeanClass (Element element) { return ReferenceBean.class; } protected void doParse (Element element, BeanDefinitionBuilder bean) { String interfaceClass = element.getAttribute( "interface" ); if (StringUtils.hasText(interfaceClass)) { bean.addPropertyValue( "interfaceClass" , interfaceClass); } } } 复制代码` ` public class TobyRPCNamespaceHandler extends NamespaceHandlerSupport { public void init () { registerBeanDefinitionParser( "reference" , new TobyRPCBeanDefinitionParser()); } } 复制代码`

4.编写 ` spring.handlers` 和 ` spring.schemas` 串联起所有部件

**spring.handlers**

` http\://toby.com/schema/tobyRPC=com.toby.config.TobyRPCNamespaceHandler 复制代码`

**spring.schemas**

` http\://toby.com/schema/tobyRPC.xsd=META-INF/tobyRPC.xsd 复制代码`

5.创建配置文件

` <? xml version= "1.0" encoding= "UTF-8" ?> < beans xmlns = "http://www.springframework.org/schema/beans" xmlns:xsi = "http://www.w3.org/2001/XMLSchema-instance" xmlns:tobyRPC = "http://toby.com/schema/tobyRPC" xsi:schemaLocation = "http://www.springframework.org/schema/beans http://www.springframework.org/schema/beans/spring-beans-2.5.xsd http://toby.com/schema/tobyRPC http://toby.com/schema/tobyRPC.xsd" > < tobyRPC:reference id = "menuService" interface = "com.toby.rpc.MenuService" /> </ beans > 复制代码`

demo结构截图如下:

![](https://user-gold-cdn.xitu.io/2019/5/6/16a8ac3295b0c0ca?imageView2/0/w/1280/h/960/ignore-error/1)

万事俱备,那我们跑个单元测试看看

![](https://user-gold-cdn.xitu.io/2019/5/6/16a8ac35ceb0e64f?imageView2/0/w/1280/h/960/ignore-error/1)

运行结果如我们所料.但是具体要怎么编码,怎么发送请求,又如何解码好像也没说啊.没说?没说就对了.在完结 ` 服务引用` 这个小专题后,还会重点和大家看一下dubbo中的 ` 编解码` , ` spi` , ` javassist` 等重点内容源码,等粗略把整个框架的思想都掌握后,再手把手临摹一个五脏俱全(包含设计模式,dubbo架构设计)的简易dubbo框架.总之一句话,关注肥朝公众号即可.

## 敲黑板划重点 ##

为什么面试都喜欢问原理,难道都是为了装逼?当然不是,明白了原理,很多东西都是一通百通的.我们来看mybatis的这道面试题.首先Mapper接口的原理,可以参考我之前的 [图解源码 | MyBatis的Mapper原理]( https://link.juejin.im?target=https%3A%2F%2Fmp.weixin.qq.com%2Fs%2FKi1f_mkoOkPMQiGbJWu4Sw ) ,其实说白了,就是给Mapper接口注入一个代理对象,然后动态代理对象调用方法会被拦截到 ` invoke` 中,然后在这个 ` invoke` 方法中,做了一些不可描述的事情(老司机可以尽情YY).而这一切的前提,都是要无声无息的把动态代理对象注入进去.其实注入进去的原理和dubbo也是一样的,我们简单看两个图

![](https://user-gold-cdn.xitu.io/2019/5/6/16a8ac37a00e2cfb?imageView2/0/w/1280/h/960/ignore-error/1)

![](https://user-gold-cdn.xitu.io/2019/5/6/16a8ac394eb426c1?imageView2/0/w/1280/h/960/ignore-error/1)

## 写在最后 ##

![](https://user-gold-cdn.xitu.io/2019/5/6/16a8ac3af00eed8c?imageView2/0/w/1280/h/960/ignore-error/1)