# Spring源码学习---IOC容器 #

## 前言 ##

在之前的5个文章中我们已经对 IOC, DI, AOP 和配置相关进行了一些了解，相信在此基础上可以帮助大家更好地去阅读 Spring 的源码

## 源码学习-IOC容器 ##

### ① 注意事项 ###

1.这里的源码版本号为 [version 5.1.3.RELEASE]( https://link.juejin.im?target=https%3A%2F%2Fdocs.spring.io%2Fspring%2Fdocs%2F5.1.3.RELEASE%2Fspring-framework-reference%2F )

2.源码获取地址 [github.com/spring-proj…]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fspring-projects%2Fspring-framework )

3.jdk要在1.8以上，spring5中大量使用了lambda表达式，而lambda表达式在1.8后开始支持

4.使用指南： [spring.io/guides]( https://link.juejin.im?target=https%3A%2F%2Fspring.io%2Fguides )

5.各个版本的介绍： [github.com/spring-proj…]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fspring-projects%2Fspring-framework%2Fwiki%2FSpring-Framework-Versions )

6.Spring5.x的版本新特性： [github.com/spring-proj…]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fspring-projects%2Fspring-framework%2Fwiki%2FWhat%2527s-New-in-Spring-Framework-5.x )

### ② IOC的步骤整理（如果已经忘记，请回到 [手写Spring---IOC容器（1）]( https://juejin.im/post/5cb1c9c4e51d456e770bdc9c ) ） ###

` 1.用户配置bean定义 ---> 2.IOC容器加载bean定义 ---> 3.IOC容器创建bean实例 ---> 4.使用IOC容器 复制代码`

### ③ 使用Spring的入口和如何做好阅读源码的准备 ###

` 使用Spring的入口为：ApplicationContext ApplicationContext context = new ClassPathXmlApplicationContext(...Service.xml); 前期准备： 创建一个maven工程，仅引入Spring-Context包（注意版本号）即可 <dependency> <groupId>org.springframework</groupId> <artifactId>spring-context</artifactId> <version>5.1.3.RELEASE</version> </dependency> 复制代码`

### ④ 围绕 ApplicationContext 我们需要了解的问题（会先大致阐述，还并未涉及具体实现） ###

#### 1. ApplicationContext 就是IOC容器 ####

内部隐藏的BeanFactory是很少被关注的，提供给用户直观看到的就是ApplicationContext，我使用的是IDEA，可以点击类后打开navigate选项---Type Hierarchy---选择SuperTypes Hierarchy，观看它的继承体系，以下是它作为一个接口，又继承了什么接口

![](https://user-gold-cdn.xitu.io/2019/6/4/16b1f8dbb630cdc7?imageView2/0/w/1280/h/960/ignore-error/1)

既然它继承了上面6个接口，那它必定会有这6个接口相关的行为

#### 2. ApplicationContext 所继承的接口 ####

` EnvironmentCapable：取环境相关的参数，.properties文件 ListableBeanFactory：提供BeanFactorys行为 HierarchicalBeanFactory：父子容器 ---提供bean分层管理的方式 且父容器无法访问子容器，子容器可以访问父容器,就比如只有儿子问老爸拿钱，没有父亲问儿子要钱的 MethodSource：国际化 ApplicationEventPublisher：应用的事件发布，比如应用的开启，结束，销毁等等 ResourcePatternResolver：加载Resource 复制代码`

这些在后面我们再展开来讲···

#### 3.继承了 ApplicationContext 的接口 ####

![](https://user-gold-cdn.xitu.io/2019/6/5/16b256602eb2d464?imageView2/0/w/1280/h/960/ignore-error/1)

那如何先大致地查看这幅图

` 先看 ConfigurableApplicationContext 这一大块， ConfigurableApplicationContext 以下全是抽象的，一直到 FileSystemXmlApplicationContext 和 ClassPathXmlApplicationContext 两个xml配置方式的具体实现， 而下面的 GenericApplicationContext 则是通用实现，其中包括了通用xml，静态，动态语言Groovy和注解， 通用xml实现 GenericXmlApplicationContext 支持 FileSystemXmlApplicationContext 和 ClassPathXmlApplicationContext， 则无论放到文件系统或者classpath都可以 两大块 AbstractRefreshableApplicationContext 和 GenericApplicationContext 都继承了 AbstractApplicationContext 复制代码`

##### 一、ConfigurableApplicationContext --- 可配置的 ApplicationContext #####

![](https://user-gold-cdn.xitu.io/2019/6/5/16b2577501b26b96?imageView2/0/w/1280/h/960/ignore-error/1)

作为一个接口肯定加入了某些行为，我们参考图左 structure 处，然后比照源码片段

##### ① void addApplicationListener(ApplicationListener<?> listener); #####

` 此处加入了应用监听器，肯定是使用了观察者模式，里面发生的事件都可以往外提供发布 通过这个 listener 就可以获取到 复制代码`

##### ② void addBeanFactoryPostProcessor(BeanFactoryPostProcessor postProcessor); #####

此处看名字是否会有熟悉的感觉，如果没有，可以回顾下 [手写Spring---AOP面向切面编程（4）]( https://juejin.im/post/5cc83ae5e51d456e7d18a000#heading-14 )

` 在手写AOP的时候，postProcessor已经被提及，讲到了beanFactory如何能够灵活扩展， 但是我们当时讲的是BeanPostProcessor，是对bean的创建过程实现阶段动态增强 那为什么现在这个是 BeanFactoryPostProcessor 呢，那就是它把工厂的创建过程也引入了各个阶段 而且提供了各个阶段支持动态增强的功能 ApplicationContext会帮我们完成bean定义的加载，解析等一系列过程，在此过程中我们可能需要灵活加入一些处理 复制代码`

##### ③ void refresh() throws BeansException, IllegalStateException; #####

` 这里的refresh方法是刷新，刷新bean定义，IOC容器里面的bean实例 复制代码`

##### 二、AbstractApplicationContext --- 抽象的 ApplicationContext 的实现 #####

此类中已经提供了很多的接口方法的实现，而且里面的定义都普遍具有了数据结构的支持

需要注意的是 registerBeanPostProcessors() 方法是保护类型的，只能供子类调用

` protected void registerBeanPostProcessors(ConfigurableListableBeanFactory beanFactory) { PostProcessorRegistrationDelegate.registerBeanPostProcessors(beanFactory, this); } 复制代码`

之前提及到的 HierarchicalBeanFactory --- 父子容器在这里也有体现

比如它的一个构造器，此时这里父容器就已经给进来了

` @Nullable private ApplicationContext parent; public AbstractApplicationContext(@Nullable ApplicationContext parent) { this(); setParent(parent); } 复制代码`

此时再看看在setParent()方法中，大致就是取得父容器的环境参数，然后进行一个比较 instanceof 与合并 merge 的事情

` @Override public void setParent(@Nullable ApplicationContext parent) { this.parent = parent; if (parent != null) { Environment parentEnvironment = parent.getEnvironment(); if (parentEnvironment instanceof ConfigurableEnvironment) { getEnvironment().merge((ConfigurableEnvironment) parentEnvironment); } } } 复制代码`

而在此实现后的继续深入扩展就是 AbstractRefreshableApplicationContext 和 AbstractRefreshableConfigApplicationContext ， 刚刚在 ConfigurableApplicationContext 的第 ③ 点不是提及了一个 refresh() 方法吗，第一个就是可支持刷新的，第二个就是可支持刷新且又被配置的，再之后就是 xml 的了，我们现在先不细说

##### 三、GenericApplicationContext --- 通用的 ApplicationContext #####

` public class GenericApplicationContext extends AbstractApplicationContext implements BeanDefinitionRegistry 复制代码`

此时往下看

##### 一、构造方法 #####

![](https://user-gold-cdn.xitu.io/2019/6/5/16b283bc85d4d0d4?imageView2/0/w/1280/h/960/ignore-error/1)

不难发现构造方法中，只有默认的 beanFactory 和父容器 parent 等作为参数，并没有提及我们可以给入 beanDefinition 的来源，比如 xml文件等等

` private final DefaultListableBeanFactory beanFactory; // 构造方法: // 如果你没有传入beanFactory，那就是默认的 public GenericApplicationContext() { this.beanFactory = new DefaultListableBeanFactory(); } // 也可以自己提供 public GenericApplicationContext(DefaultListableBeanFactory beanFactory) { Assert.notNull(beanFactory, "BeanFactory must not be null"); this.beanFactory = beanFactory; } 复制代码`

和上面一样的套路，提供了setParent()方法

` @Override public void setParent(@Nullable ApplicationContext parent) { super.setParent(parent); this.beanFactory.setParentBeanFactory(getInternalParentBeanFactory()); } 复制代码`

##### 二、模板方法 refreshBeanFactory() --- 实现了父类的保护类型的方法 #####

` //--------------------------------------------------------------------- // Implementations of AbstractApplicationContext's template methods //--------------------------------------------------------------------- /** * Do nothing: We hold a single internal BeanFactory and rely on callers * to register beans through our public methods (or the BeanFactory's). * @see #registerBeanDefinition */ @Override protected final void refreshBeanFactory() throws IllegalStateException { if (!this.refreshed.compareAndSet(false, true)) { throw new IllegalStateException( "GenericApplicationContext does not support multiple refresh attempts: just call 'refresh' once"); } this.beanFactory.setSerializationId(getId()); } @Override protected void cancelRefresh(BeansException ex) { this.beanFactory.setSerializationId(null); super.cancelRefresh(ex); } 复制代码`

##### 三、接口 BeanDefinitionRegistry 要求提供的方法，包括实现注册bean定义等等 #####

![](https://user-gold-cdn.xitu.io/2019/6/5/16b2829f7f803910?imageView2/0/w/1280/h/960/ignore-error/1)

##### 四、其他 #####

![](https://user-gold-cdn.xitu.io/2019/6/5/16b283a8374e8302?imageView2/0/w/1280/h/960/ignore-error/1)

这样下来，我们能大致了解最外层的 ApplciationContext 是如何一步步加入哪些参数的

## Finally ##

此篇只是大致地看了一下 ApplicationContext 的子类所拥有的一些东西，可能大家看起来会觉得一头雾水，下一篇会结合实例去进行使用然后一步步进行分析，可能篇幅就会增加。

有人反映了篇幅过长的问题，所以现在秉承少吃多餐的原则去慢更，望多多总结，互相进步··