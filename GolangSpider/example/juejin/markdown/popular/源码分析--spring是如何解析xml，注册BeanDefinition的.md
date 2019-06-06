# 源码分析--spring是如何解析xml，注册BeanDefinition的 #

个人网站原文： [chenmingyu.top/spring-sour…]( https://link.juejin.im?target=https%3A%2F%2Fchenmingyu.top%2Fspring-source-base%2F )

## spring源码解析 ##

本文首先提供了一个实现了spring aop的demo，通过demo进行源码分析

通过读源码我们可以学习到spring是 **如何解析xml的** ，如何 **加载bean** 的，如何 **创建bean** 的，又是如何 **实现aop** 操作的，及其中各种操作的细节是如何实现的

讲源码的时候我会进行一些取舍，根据上面的问题结合demo对主要流程进行讲解，争取能把上述的问题说明白

**源码地址： [github.com/mingyuHub/s…]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2FmingyuHub%2Fspring-shared )**

### Aop demo ###

代码：

* 一个类 ` UserController` ，提供一个方法 ` login`
* 一个切面 ` UserAspect` ，切入点为 ` login` 方法
* 一个配置文件 ` spring-aop.xml` 将类加载到spring容器中

创建 ` UserController` 类

` public class UserController { public void login () { System.out.println( "登录" ); } } 复制代码`

定义一个切面 ` UserAspect` ，不了解aop概念的可以看一下： [chenmingyu.top/springboot-…]( https://link.juejin.im?target=https%3A%2F%2Fchenmingyu.top%2Fspringboot-aop%2F )

` /** * @author : chenmingyu * @date : 2019/3/19 18:29 * @description : */ @Aspect public class UserAspect { /** * 切入点 */ @Pointcut ( "execution(public * com.my.spring.*.*(..))" ) public void execute () { } /** * 前置通知 * @param joinPoint */ @Before (value = "execute()" ) public void Before (JoinPoint joinPoint) { System.out.println( "执行方法之前" ); } /** * 后置通知 * @param joinPoint */ @After (value = "execute()" ) public void After (JoinPoint joinPoint) { System.out.println( "执行方法之后" ); } } 复制代码`

自定义一个xml文件，名为 ` spring-aop.xml`

` <?xml version="1.0" encoding="UTF-8"?> <beans xmlns="http://www.springframework.org/schema/beans" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:context="http://www.springframework.org/schema/context" xmlns:aop="http://www.springframework.org/schema/aop" xsi:schemaLocation="http://www.springframework.org/schema/beans http://www.springframework.org/schema/beans/spring-beans-2.5.xsd http://www.springframework.org/schema/aop http://www.springframework.org/schema/aop/spring-aop-3.1.xsd "> <aop:aspectj-autoproxy proxy-target-class="true"/> <bean id="userController" class="com.my.spring.UserController"/> <bean id="userAspect" class="com.my.spring.UserAspect"/> </beans> 复制代码`

调试的代码已经准备好，首先写一个测试类测一下

` @Test public void test () { ApplicationContext applicationContext = new ClassPathXmlApplicationContext( "spring-aop.xml" ); UserController userController = (UserController) applicationContext.getBean( "userController" ); userController.login(); } 复制代码`

正常情况下的输出如下：

![](https://user-gold-cdn.xitu.io/2019/6/5/16b263e7b9905689?imageView2/0/w/1280/h/960/ignore-error/1)

### 核心类 ###

开始学习spring的源码之前，有必要先了解一下spring中几个较为核心的类

先大概了解一下这些类是干什么的，不必深究，后续读源码的时候碰到会着重讲解一下

先看一个spring容器的类图：

![](https://user-gold-cdn.xitu.io/2019/6/5/16b263ec5d7ffadf?imageView2/0/w/1280/h/960/ignore-error/1)

* 

**BeanFactory**

工厂类的顶级接口，用于获取bean及bean的各种属性，提供了ioc容器最基本的形式，给具体的IOC容器的实现提供了规范

ListableBeanFactory、HierarchicalBeanFactory 和 AutowireCapableBeanFactory是BeanFactory接口的子接口，最终的默认实现类是 DefaultListableBeanFactory，定义这多接口主要是为了区分在 Spring 内部在操作过程中对象的传递和转化过程中，对对象的数据访问所做的限制

* 

**DefaultListableBeanFactory**

ioc容器的实现，DefaultListableBeanFactory作为一个可以独立使用的ioc容器，是整个Bean加载的核心部分，是spring注册及加载bean的默认实现

* 

**xmlBeanFactory**

xmlBeanFactory继承DefaultListableBeanFactory，对其进行了扩展，增加了自定义的xml读取器 ` XmlBeanDefinitionReader` ，实现了个性化的BeanDefinitionReader读取，主要作用就是将xml配置解析成BeanDefinition

* 

**ApplicationContext**

ApplicationContext继承自BeanFactory，包含BeanFactory的所有功能，情况下都使用这个

* 

**BeanDefinition**

Spring中用于包装Bean的数据结构

* 

**BeanDefinitionRegistory**

定义对BeanDefinition的各种增删操作

* 

**BeanDefinitionReader**

定义了读取BeanDefinition的接口，主要作用是从资源文件中读取Bean定义，XmlBeanDefinitionReader是其具体的实现类

* 

**SingletonBeanRegistry**

定义对单例的注册及获取

* 

**AliasRegistry**

定义对alias的简单增删改操作

了解一些核心类之后我们就要开始读源码了

### 源码分析 ###

我们的源码以测试类为入口开始分析

` ApplicationContext applicationContext = new ClassPathXmlApplicationContext( "spring-aop.xml" ); 复制代码`

` ApplicationContext` 和 ` BeanFactory` 都是用于加载Bean的，相比之下 ` ApplicationContext` 提供了更多的扩展功能

` ClassPathXmlApplicationContext` 最终调用下面这个构造函数

` public ClassPathXmlApplicationContext (String[] configLocations, boolean refresh, ApplicationContext parent) throws BeansException { super (parent); //设置配置文件 setConfigLocations(configLocations); if (refresh) { refresh(); } } 复制代码`

` refresh()` 是 ` ApplicationContext` 的核心方法，这个方法基本包含了 ` ApplicationContext` 提供的全部功能

` @Override public void refresh () throws BeansException, IllegalStateException { synchronized ( this.startupShutdownMonitor) { // 准备刷新此上下文，不重要. prepareRefresh(); // 初始化bean工厂，加载xml，解析默认标签，解析自定义标签，注册BeanDefinitions ConfigurableListableBeanFactory beanFactory = obtainFreshBeanFactory(); // 设置bean工厂的属性，进行功能填充. prepareBeanFactory(beanFactory); try { // 子类覆盖方法做额外的处理. postProcessBeanFactory(beanFactory); // 激活bean处理器. invokeBeanFactoryPostProcessors(beanFactory); // 注册拦截bean创建的bean处理器，只是注册 registerBeanPostProcessors(beanFactory); // 初始化message源. initMessageSource(); // 初始化应用消息广播器 initApplicationEventMulticaster(); // 留给子类加载其他bean. onRefresh(); // 注册Listeners bean，到消息广播器 registerListeners(); // 实例化所有剩余的（非lazy init）单例. finishBeanFactoryInitialization(beanFactory); // 刷新通知. finishRefresh(); } catch (BeansException ex) { logger.warn( "Exception encountered during context initialization - cancelling refresh attempt" , ex); // Destroy already created singletons to avoid dangling resources. destroyBeans(); // Reset 'active' flag. cancelRefresh(ex); // Propagate exception to caller. throw ex; } } } 复制代码`

#### spring是如何加载解析xml，注册BeanDefinition的 ####

加载解析xml和注册BeanDefinition的逻辑都在 ` obtainFreshBeanFactory()` 方法中，这个方法的作用是初始化bean工厂，加载xml，解析默认标签和自定义标签，将解析出来的 ` BeanDefinition` 注册到容器中

` protected ConfigurableListableBeanFactory obtainFreshBeanFactory () { // 核心逻辑：创建bean工厂，解析xml，注册BeanDefinition refreshBeanFactory(); ConfigurableListableBeanFactory beanFactory = getBeanFactory(); if (logger.isDebugEnabled()) { logger.debug( "Bean factory for " + getDisplayName() + ": " + beanFactory); } return beanFactory; } ------------------调用下面这个方法------------------ //调用AbstractRefreshableApplicationContext的refreshBeanFactory() protected final void refreshBeanFactory () throws BeansException { if (hasBeanFactory()) { destroyBeans(); closeBeanFactory(); } try { //创建bean工厂，类型为DefaultListableBeanFactory DefaultListableBeanFactory beanFactory = createBeanFactory(); beanFactory.setSerializationId(getId()); customizeBeanFactory(beanFactory); //核心逻辑：加载BeanDefinitions，进入这个方法 loadBeanDefinitions(beanFactory); synchronized ( this.beanFactoryMonitor) { this.beanFactory = beanFactory; } } catch (IOException ex) { throw new ApplicationContextException( "I/O error parsing bean definition source for " + getDisplayName(), ex); } } 复制代码`

` createBeanFactory()` 方法逻辑特别简单，我们详细说一下 ` loadBeanDefinitions(beanFactory)` 方法

在 ` loadBeanDefinitions()` 这个方法实例化了一个 **XmlBeanDefinitionReader** ，介绍 ` XmlBeanDefinitionReader` 之前需要介绍一下什么是 ` BeanDefinition`

**BeanDefinition** ：Bean的定义主要由 ` BeanDefinition` 来描述的。作为Spring中用于包装Bean的数据结构

` BeanDefinition` 作为顶级接口 ，拥有三种实现： ` RootBeanDefinition` ， ` ChildBeanDefinition` ， ` GenericBeanDefinition` ，三种 ` BeanDefinition` 均继承了 ` AbstractBeanDefinition`

![](https://user-gold-cdn.xitu.io/2019/6/5/16b263f16ba11559?imageView2/0/w/1280/h/960/ignore-error/1)

spring通过 ` BeanDefinition` 将配置文件中的标签转换为容器的内部表示，并将 ` BeanDefinition` 注册到 ` BeandefinitionRegistry` 中，spring中的容器主要以map的形式进行存储 ` BeanDefinition`

再介绍一下 **BeanDefinitionReader** ：

![](https://user-gold-cdn.xitu.io/2019/6/5/16b263f2ef156509?imageView2/0/w/1280/h/960/ignore-error/1)

**BeanDefinitionReader** 解决的是从资源文件（xml,propert）解析到 ` BeanDefinition` 的过程，所以 **XmlBeanDefinitionReader** 的作用就很明显了，将xml转为 ` BeanDefinition`

` protected void loadBeanDefinitions (DefaultListableBeanFactory beanFactory) throws BeansException, IOException { // 根据beanFactory创建XmlBeanDefinitionReader XmlBeanDefinitionReader beanDefinitionReader = new XmlBeanDefinitionReader(beanFactory); // 加载环境变量啥的 beanDefinitionReader.setEnvironment( this.getEnvironment()); beanDefinitionReader.setResourceLoader( this ); beanDefinitionReader.setEntityResolver( new ResourceEntityResolver( this )); // initBeanDefinitionReader initBeanDefinitionReader(beanDefinitionReader); //核心逻辑在这个方法里，加载beanDefinitions loadBeanDefinitions(beanDefinitionReader); } 复制代码`

在 ` loadBeanDefinitions(beanDefinitionReader);` 方法中我们层层调用，发现最终解析xml调用的方法是 ` doLoadBeanDefinitions(InputSource inputSource, Resource resource)`

` protected int doLoadBeanDefinitions (InputSource inputSource, Resource resource) throws BeanDefinitionStoreException { try { //将xml转成document对象 Document doc = doLoadDocument(inputSource, resource); //核心逻辑在这个方法里，解析Document，注册beanDefinition return registerBeanDefinitions(doc, resource); } catch (BeanDefinitionStoreException ex) { throw ex; } ......省略其他异常信息 } 复制代码`

` registerBeanDefinitions(doc, resource);` 方法中最终调用 ` doRegisterBeanDefinitions(Element root)` 将由 ` xml` 转出来的 ` Document` （通过 ` doc.getDocumentElement()` 取到Element 元素）解析成 ` beandefinition` 并注册

` protected void doRegisterBeanDefinitions (Element root) { BeanDefinitionParserDelegate parent = this.delegate; this.delegate = createDelegate(getReaderContext(), root, parent); if ( this.delegate.isDefaultNamespace(root)) { String profileSpec = root.getAttribute(PROFILE_ATTRIBUTE); if (StringUtils.hasText(profileSpec)) { String[] specifiedProfiles = StringUtils.tokenizeToStringArray( profileSpec, BeanDefinitionParserDelegate.MULTI_VALUE_ATTRIBUTE_DELIMITERS); if (!getReaderContext().getEnvironment().acceptsProfiles(specifiedProfiles)) { return ; } } } //解析前置操作，留给子类实现 preProcessXml(root); // 核心逻辑：解析并注册BeanDefinition parseBeanDefinitions(root, this.delegate); //解析后置操作，留给子类实现 postProcessXml(root); this.delegate = parent; } 复制代码`

经历层层调用我们终于找到了核心方法，解析 ` beanDefinition` ，我们看看它是如何进行解析的

spring中的标签包括默认标签和自定义标签两种，但是两种标签的解析方式有很大的区别

` protected void parseBeanDefinitions (Element root, BeanDefinitionParserDelegate delegate) { //根据root元素的namespace是否等于http://www.springframework.org/schema/beans //通过上面的判断，确定是否是属于spring的默认标签 if (delegate.isDefaultNamespace(root)) { NodeList nl = root.getChildNodes(); for ( int i = 0 ; i < nl.getLength(); i++) { Node node = nl.item(i); if (node instanceof Element) { Element ele = (Element) node; if (delegate.isDefaultNamespace(ele)) { //默认标签解析 parseDefaultElement(ele, delegate); } else { //自定义标签解析 delegate.parseCustomElement(ele); } } } } else { //自定义标签解析 delegate.parseCustomElement(root); } } 复制代码`

#### 默认标签解析 ####

` parseDefaultElement()` 方法提供了对import，alias，bean，beans标签的解析

` private void parseDefaultElement (Element ele, BeanDefinitionParserDelegate delegate) { //解析<import>标签 if (delegate.nodeNameEquals(ele, IMPORT_ELEMENT)) { importBeanDefinitionResource(ele); } //解析<alias>标签 else if (delegate.nodeNameEquals(ele, ALIAS_ELEMENT)) { processAliasRegistration(ele); } //解析<bean>标签 else if (delegate.nodeNameEquals(ele, BEAN_ELEMENT)) { processBeanDefinition(ele, delegate); } //解析<beans>标签 else if (delegate.nodeNameEquals(ele, NESTED_BEANS_ELEMENT)) { // recurse doRegisterBeanDefinitions(ele); } } 复制代码`

详细的讲解一下对 ` bean` 标签的解析

` protected void processBeanDefinition (Element ele, BeanDefinitionParserDelegate delegate) { //步骤一，将Element解析成BeanDefinitionHolder. BeanDefinitionHolder bdHolder = delegate.parseBeanDefinitionElement(ele); if (bdHolder != null ) { // 对 bdHolder 进行装饰，针对自定义的属性进行解析，根据自定义标签找到对应的处理器，进行解析（自定义解析方式下面会细说） bdHolder = delegate.decorateBeanDefinitionIfRequired(ele, bdHolder); try { //步骤二，注册解析出来的BeanDefinition. BeanDefinitionReaderUtils.registerBeanDefinition(bdHolder, getReaderContext().getRegistry()); } catch (BeanDefinitionStoreException ex) { getReaderContext().error( "Failed to register bean definition with name '" + bdHolder.getBeanName() + "'" , ele, ex); } // Send registration event. getReaderContext().fireComponentRegistered( new BeanComponentDefinition(bdHolder)); } } 复制代码`

**步骤一** ：先介绍下 **BeanDefinitionHolder** ：这个类是一个工具类，作用是承载 **BeanDefinition** 数据的

` public class BeanDefinitionHolder implements BeanMetadataElement { private final BeanDefinition beanDefinition; private final String beanName; private final String[] aliases; ... } 复制代码`

` delegate.parseBeanDefinitionElement(ele);` 方法中将 ` Element` 转化为 **BeanDefinitionHolder** ，并且识别出bean的beanName和aliases(别名)，得到 ` BeanDefinitionHolder` 其实默认标签的解析就已经结束了

这个方法没有啥复杂逻辑，挺清晰的

` public BeanDefinitionHolder parseBeanDefinitionElement (Element ele) { return parseBeanDefinitionElement(ele, null ); } ----------------------调用下面这个方法------------------------------ public BeanDefinitionHolder parseBeanDefinitionElement (Element ele, BeanDefinition containingBean) { //获取id属性 String id = ele.getAttribute(ID_ATTRIBUTE); //获取name属性 String nameAttr = ele.getAttribute(NAME_ATTRIBUTE); List<String> aliases = new ArrayList<String>(); //如果name属性配置的不为空 if (StringUtils.hasLength(nameAttr)) { //按,; 分割成字符串数组 String[] nameArr = StringUtils.tokenizeToStringArray(nameAttr, MULTI_VALUE_ATTRIBUTE_DELIMITERS); //加到别名的集合里 aliases.addAll(Arrays.asList(nameArr)); } //把id属性赋值给beanName，如果id为空就在aliases别名的集合里取第一个 String beanName = id; if (!StringUtils.hasText(beanName) && !aliases.isEmpty()) { beanName = aliases.remove( 0 ); if (logger.isDebugEnabled()) { logger.debug( "No XML 'id' specified - using '" + beanName + "' as bean name and " + aliases + " as aliases" ); } } if (containingBean == null ) { //检查beanName和aliases是否已经使用，如果使用了就报异常，没使用就加到一个 checkNameUniqueness(beanName, aliases, ele); } //创建了一个GenericBeanDefinition类型的BeanDefinition，并对个属性进行填充 AbstractBeanDefinition beanDefinition = parseBeanDefinitionElement(ele, beanName, containingBean); if (beanDefinition != null ) { if (!StringUtils.hasText(beanName)) { try { if (containingBean != null ) { beanName = BeanDefinitionReaderUtils.generateBeanName( beanDefinition, this.readerContext.getRegistry(), true ); } else { beanName = this.readerContext.generateBeanName(beanDefinition); String beanClassName = beanDefinition.getBeanClassName(); if (beanClassName != null && beanName.startsWith(beanClassName) && beanName.length() > beanClassName.length() && ! this.readerContext.getRegistry().isBeanNameInUse(beanClassName)) { aliases.add(beanClassName); } } if (logger.isDebugEnabled()) { logger.debug( "Neither XML 'id' nor 'name' specified - " + "using generated bean name [" + beanName + "]" ); } } catch (Exception ex) { error(ex.getMessage(), ele); return null ; } } String[] aliasesArray = StringUtils.toStringArray(aliases); //创建一个BeanDefinitionHolder返回 return new BeanDefinitionHolder(beanDefinition, beanName, aliasesArray); } return null ; } 复制代码`

得到 ` BeanDefinitionHolder` 后，剩下的就是注册 ` BeanDefinition` 了

**步骤二** ，调用 ` BeanDefinitionReaderUtils.registerBeanDefinition(bdHolder, getReaderContext().getRegistry())` 方法，注册 ` BeanDifinition`

` public static void registerBeanDefinition ( BeanDefinitionHolder definitionHolder, BeanDefinitionRegistry registry) throws BeanDefinitionStoreException { // 使用beanName做唯一标识注册 String beanName = definitionHolder.getBeanName(); registry.registerBeanDefinition(beanName, definitionHolder.getBeanDefinition()); // 使用所有别名进行注册 String[] aliases = definitionHolder.getAliases(); if (aliases != null ) { for (String alias : aliases) { registry.registerAlias(beanName, alias); } } } 复制代码`

通过 **beanName** 进行注册： ` definitionHolder.getBeanName()` ，默认标签和自定义标签都使用这个方法进行 ` BeanDefinition` 的注册

` public void registerBeanDefinition (String beanName, BeanDefinition beanDefinition) throws BeanDefinitionStoreException { Assert.hasText(beanName, "Bean name must not be empty" ); Assert.notNull(beanDefinition, "BeanDefinition must not be null" ); // beanDefinition是否属于AbstractBeanDefinition的实例 if (beanDefinition instanceof AbstractBeanDefinition) { try { // 进行校验，主要是校验methodOverrides与工程方法是否存在以及methodOverrides对应的方法存不存在 ((AbstractBeanDefinition) beanDefinition).validate(); } catch (BeanDefinitionValidationException ex) { throw new BeanDefinitionStoreException(beanDefinition.getResourceDescription(), beanName, "Validation of bean definition failed" , ex); } } BeanDefinition oldBeanDefinition; // 通过beanName获取BeanDefinition是否已经注册 oldBeanDefinition = this.beanDefinitionMap.get(beanName); //如果已经注册并且不允许覆盖就抛出异常 if (oldBeanDefinition != null ) { if (!isAllowBeanDefinitionOverriding()) { throw new BeanDefinitionStoreException(beanDefinition.getResourceDescription(), beanName, "Cannot register bean definition [" + beanDefinition + "] for bean '" + beanName + "': There is already [" + oldBeanDefinition + "] bound." ); } else if (oldBeanDefinition.getRole() < beanDefinition.getRole()) { // e.g. was ROLE_APPLICATION, now overriding with ROLE_SUPPORT or ROLE_INFRASTRUCTURE if ( this.logger.isWarnEnabled()) { this.logger.warn( "Overriding user-defined bean definition for bean '" + beanName + "' with a framework-generated bean definition: replacing [" + oldBeanDefinition + "] with [" + beanDefinition + "]" ); } } else { if ( this.logger.isInfoEnabled()) { this.logger.info( "Overriding bean definition for bean '" + beanName + "': replacing [" + oldBeanDefinition + "] with [" + beanDefinition + "]" ); } } } else { //注册BeanDefinition this.beanDefinitionNames.add(beanName); this.manualSingletonNames.remove(beanName); this.frozenBeanDefinitionNames = null ; } this.beanDefinitionMap.put(beanName, beanDefinition); // 如果oldBeanDefinition通过上述校验没抛出异常或者beanName是单例 if (oldBeanDefinition != null || containsSingleton(beanName)) { //则更新对应的缓存 resetBeanDefinition(beanName); } } 复制代码`

通过 ` this.beanDefinitionMap.put(beanName, beanDefinition)` 这行代码我们就知道，spring使用一个叫 ` beanDefinitionMap` 的 **ConcurrentHashMap** 来存储解析出来的 ` beanDefinition`

` Map<String, BeanDefinition> beanDefinitionMap = new ConcurrentHashMap<String, BeanDefinition>(64);`

通过别名注册的方式跟通过beanName注册的区别不大，仔细看下 ` registry.registerAlias(beanName, alias);` 方法应该就能了解

spring默认标签的解析大致流程就是这样，细枝末节并没有特别详细的讲解，不过这并不会对我们理解spring的整体流程有阻碍，大家可自行看一下，代码逻辑也不是特别复杂

#### 自定义标签解析 ####

自定义标签解析的时候会先根据从 ` Element` 获取到的 ` namespaceUri` 获取到对应的 ` NamespaceHandler` ，根据 ` NamespaceHandler` 进行自定义的解析，以aop为例，我们在配置文件中配置了 ` <aop:aspectj-autoproxy proxy-target-class="true"/>` ，解析会根据 ` aspectj-autoproxy` 找到对应的处理器，然后调用其 ` parse` 方法创建

` delegate.parseCustomElement(root);` 定义了自定义标签解析的流程

` public BeanDefinition parseCustomElement (Element ele) { return parseCustomElement(ele, null ); } -----------------------调用下面这个方法----------------------- public BeanDefinition parseCustomElement (Element ele, BeanDefinition containingBd) { //获取命名空间 String namespaceUri = getNamespaceURI(ele); //步骤一，根据命名空间找到对应的NamespaceHandler NamespaceHandler handler = this.readerContext.getNamespaceHandlerResolver().resolve(namespaceUri); if (handler == null ) { error( "Unable to locate Spring NamespaceHandler for XML schema namespace [" + namespaceUri + "]" , ele); return null ; } //步骤二，根据自定义的NamespaceHandler进行解析 return handler.parse(ele, new ParserContext( this.readerContext, this , containingBd)); } 复制代码`

**步骤一** ：有了 ` namespaceUri` 我们就可以根据 ` this.readerContext.getNamespaceHandlerResolver().resolve(namespaceUri)` 方法获取对应 ` NamespaceHandler` 的实例

` //DefaultNamespaceHandlerResolver.java public NamespaceHandler resolve (String namespaceUri) { //1,获取到所有的解析器 Map<String, Object> handlerMappings = getHandlerMappings(); //2,根据 namespaceUri 获取到对应的handle Object handlerOrClassName = handlerMappings.get(namespaceUri); if (handlerOrClassName == null ) { return null ; } else if (handlerOrClassName instanceof NamespaceHandler) { //3,强转返回 return (NamespaceHandler) handlerOrClassName; } else { //4,根据handlerOrClassName实例化NamespaceHandler String className = (String) handlerOrClassName; try { Class<?> handlerClass = ClassUtils.forName(className, this.classLoader); if (!NamespaceHandler.class.isAssignableFrom(handlerClass)) { throw new FatalBeanException( "Class [" + className + "] for namespace [" + namespaceUri + "] does not implement the [" + NamespaceHandler.class.getName() + "] interface" ); } NamespaceHandler namespaceHandler = (NamespaceHandler) BeanUtils.instantiateClass(handlerClass); //调用自定义的NamespaceHandler的初始化方法 namespaceHandler.init(); //添加到缓存里 handlerMappings.put(namespaceUri, namespaceHandler); return namespaceHandler; } catch (ClassNotFoundException ex) { throw new FatalBeanException( "NamespaceHandler class [" + className + "] for namespace [" + namespaceUri + "] not found" , ex); } catch (LinkageError err) { throw new FatalBeanException( "Invalid NamespaceHandler class [" + className + "] for namespace [" + namespaceUri + "]: problem with handler class file or dependent class" , err); } } } 复制代码`

这个过程还是比较清晰的，debug一下：

![](https://user-gold-cdn.xitu.io/2019/6/5/16b263f8347f9275?imageView2/0/w/1280/h/960/ignore-error/1)

这个 ` handlerMappings.get(namespaceUri)` 取到是字符串: ` org.springframework.aop.config.AopNamespaceHandler` ，接下来按流程走调用 ` BeanUtils.instantiateClass(handlerClass)` 就是实例化这个类，然后调用它的 ` init` 方法，然后加到缓存里，然后返回这个 ` NamespaceHandler` 的实例

` public class AopNamespaceHandler extends NamespaceHandlerSupport { /** * Register the { @link BeanDefinitionParser BeanDefinitionParsers} for the * '{ @code config}', '{ @code spring-configured}', '{ @code aspectj-autoproxy}' * and '{ @code scoped-proxy}' tags. */ @Override public void init () { // In 2.0 XSD as well as in 2.1 XSD. registerBeanDefinitionParser( "config" , new ConfigBeanDefinitionParser()); registerBeanDefinitionParser( "aspectj-autoproxy" , new AspectJAutoProxyBeanDefinitionParser()); registerBeanDefinitionDecorator( "scoped-proxy" , new ScopedProxyBeanDefinitionDecorator()); // Only in 2.0 XSD: moved to context namespace as of 2.1 registerBeanDefinitionParser( "spring-configured" , new SpringConfiguredBeanDefinitionParser()); } } 复制代码`

注册 ` BeanDefinitionParser` ，就是放到一个叫 ` parsers` 的 **HashMap** 里，总共4个 ` config` ， ` aspectj-autoproxy` ， ` scoped-proxy` ， ` spring-configured`

` protected final void registerBeanDefinitionParser (String elementName, BeanDefinitionParser parser) { this.parsers.put(elementName, parser); } 复制代码`

**步骤二** ：返回 ` NamespaceHandler` 实例之后调用它的 ` parse` 方法

` handler.parse(ele, new ParserContext(this.readerContext, this, containingBd));`

我们的配置文件中只有一个自定义标签： ` <aop:aspectj-autoproxy proxy-target-class="true"/>`

所以 ` findParserForElement(element, parserContext)` 这个方法根据标签 ` aspectj-autoproxy` 取到的是取到的 ` BeanDefinition` 是： ` AspectJAutoProxyBeanDefinitionParser`

` //NamespaceHandlerSupport.java //获取到对应解析器的BeanDefinition，调用其parse方法 //比如aspectj-autoproxy标签对应AspectJAutoProxyBeanDefinitionParser public BeanDefinition parse (Element element, ParserContext parserContext) { return findParserForElement(element, parserContext).parse(element, parserContext); } --------------------调用下面这个方法----------------------- private BeanDefinitionParser findParserForElement (Element element, ParserContext parserContext) { String localName = parserContext.getDelegate().getLocalName(element); BeanDefinitionParser parser = this.parsers.get(localName); if (parser == null ) { parserContext.getReaderContext().fatal( "Cannot locate BeanDefinitionParser for element [" + localName + "]" , element); } return parser; } 复制代码`

` AspectJAutoProxyBeanDefinitionParser` 实现 ` BeanDefinitionParser` 接口

` BeanDefinitionParser` 接口中只定义了一个 ` parse` 方法，所有自定义处理器都需要实现 ` BeanDefinitionParser` 接口进行自定标签的解析

接下来我们看下 ` AspectJAutoProxyBeanDefinitionParser` 类的 ` parse` 方法

` //AspectJAutoProxyBeanDefinitionParser.java public BeanDefinition parse (Element element, ParserContext parserContext) { //注册 AspectJAnnotationAutoProxyCreator AopNamespaceUtils.registerAspectJAnnotationAutoProxyCreatorIfNecessary(parserContext, element); extendBeanDefinition(element, parserContext); return null ; } 复制代码`

主要的逻辑就在注册 ` AspectJAnnotationAutoProxyCreator` 这个方法上

` public static void registerAspectJAnnotationAutoProxyCreatorIfNecessary ( ParserContext parserContext, Element sourceElement) { //核心逻辑：注册或升级AutoProxyCreator定义beanName为org.springframework.aop.config.internalAutoProxyCreator的BeanDefinition BeanDefinition beanDefinition = AopConfigUtils.registerAspectJAnnotationAutoProxyCreatorIfNecessary( parserContext.getRegistry(), parserContext.extractSource(sourceElement)); useClassProxyingIfNecessary(parserContext.getRegistry(), sourceElement); registerComponentIfNecessary(beanDefinition, parserContext); } 复制代码`

这个方法调用了 ` registerOrEscalateApcAsRequired(AnnotationAwareAspectJAutoProxyCreator.class, registry, source)` 方法

调用的这个方法逻辑也特别清晰

` public static BeanDefinition registerAspectJAnnotationAutoProxyCreatorIfNecessary (BeanDefinitionRegistry registry, Object source) { return registerOrEscalateApcAsRequired(AnnotationAwareAspectJAutoProxyCreator.class, registry, source); } -----------------------调用下面这个方法--------------------- private static BeanDefinition registerOrEscalateApcAsRequired (Class<?> cls, BeanDefinitionRegistry registry, Object source) { Assert.notNull(registry, "BeanDefinitionRegistry must not be null" ); //判断BeanDefinitionRegistry是否包含AUTO_PROXY_CREATOR_BEAN_NAME这个静态变量 if (registry.containsBeanDefinition(AUTO_PROXY_CREATOR_BEAN_NAME)) { //包含说明就注册过，将这个BeanDefinition取出来，然后判断BeanClassName如果不相等，重新BeanClassName为cls.getName() BeanDefinition apcDefinition = registry.getBeanDefinition(AUTO_PROXY_CREATOR_BEAN_NAME); if (!cls.getName().equals(apcDefinition.getBeanClassName())) { int currentPriority = findPriorityForClass(apcDefinition.getBeanClassName()); int requiredPriority = findPriorityForClass(cls); if (currentPriority < requiredPriority) { apcDefinition.setBeanClassName(cls.getName()); } } return null ; } //如果不包含就创建一个RootBeanDefinition，填充属性然后注册 RootBeanDefinition beanDefinition = new RootBeanDefinition(cls); beanDefinition.setSource(source); beanDefinition.getPropertyValues().add( "order" , Ordered.HIGHEST_PRECEDENCE); beanDefinition.setRole(BeanDefinition.ROLE_INFRASTRUCTURE); //这个注册方法我们前面讲默认标签注册BeanDefinition的时候讲过，用的一个方法 registry.registerBeanDefinition(AUTO_PROXY_CREATOR_BEAN_NAME, beanDefinition); return beanDefinition; } 复制代码`

自定义标签解析注册 ` BeanDefinition` 的过程我们也讲解完了

现在我们知道了spring是如果解析默认标签和自定义标签的了，整体流程还是比较清晰的

总结一下spring如何加载xml及注册BeanDefinition：

首先将xml文件转化为Element对象，获取命名空间，根据命名空间判断是spring的默认标签还是自定义标签

* 

默认标签：使用spring的流程进行处理，遇到默认标签首先判断是哪种标签，import，alias，bean，beans标签都有着不同的解析处理逻辑，解析成BeanDefinition之后进行注册，注册的过程就是放到一个 ` ConcurrentHashMap` 里

* 

自定义标签：使用自定义的命名空间处理器（实现了 ` NamespaceHandler` 接口）进行解析注册处理

首先根据 ` namespaceUri` 找到对应的 ` NamespaceHandler` 处理器

然后调用它的init方法，注册对应自定义标签的解析器（比如 ` aspectj-autoproxy` 对应 ` AspectJAutoProxyBeanDefinitionParser` ）

调用 ` NamespaceHandler` 的 ` parse` 方法，在这个方法里根据自定义标签找到对应的解析器，调用对应的解析器的 ` parse` 方法进行注册 ` BeanDefinition`

本想一篇文章把所有的问题都说明白，发现写完一个问题篇幅就比较长了

那关于spring是如何加载bean的，如何创建bean的，又是如何实现aop操作的，我们下篇分解

觉得写得还可以记得点关注，下次更新文章直接就能看到了