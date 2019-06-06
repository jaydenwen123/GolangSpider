# 我该如何学习spring源码以及解析bean定义的注册 #

## 如何学习spring源码 ##

### 前言 ###

本文属于spring源码解析的系列文章之一，文章主要是介绍如何学习spring的源码，希望能够最大限度的帮助到有需要的人。文章总体难度不大，但比较繁重，学习时一定要耐住性子坚持下去。

### 获取源码 ###

源码的获取有多种途径

#### GitHub ####

[spring-framework]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fspring-projects%2Fspring-framework )

[spring-wiki]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fspring-projects%2Fspring-framework%2Fwiki )

可以从GitHub上获取源代码，然后自行编译

#### maven ####

使用过maven的都知道可以通过maven下载相关的源代码和相关文档，简单方便。

这里推荐通过maven的方式构建一个web项目。通过对实际项目的运行过程中进行调试来学习更好。

### 如何开始学习 ###

#### 前置条件 ####

如果想要开始学习spring的源码，首先要求本身对spring框架的使用基本了解。明白spring中的一些特性如ioc等。了解spring中各个模块的作用。

#### 确定目标 ####

首先我们要知道spring框架本身经过多年的发展现在已经是一个庞大的家族。可能其中一个功能的实现依赖于多个模块多个类的相互配合，这样会导致我们在阅读代码时难度极大。多个类之间进行跳跃很容易让我们晕头转向。

所以在阅读spring的源代码的时候不能像在JDK代码时一行一行的去理解代码，需要把有限的精力更多的分配给重要的地方。而且我们也没有必要这样去阅读。

在阅读spring某一功能的代码时应当从一个上帝视角来总览全局。只需要知道某一个功能的实现流程即可，而且幸运的是spring的代码规范较好，大多数方法基本都能见名知意，这样也省去了我们很多的麻烦。

#### 利用好工具 ####

阅读代码最好在idea或者eclipse中进行，这类IDE提供的很多功能很有帮助。

在阅读时配合spring文档更好(如果是自行编译源码直接看注释更好)。

#### 笔记和复习 ####

这个过程及其重要，我以前也看过一些spring的源码，但是好几次都是感觉比较吃力在看过一些后就放弃了。而由于没有做笔记和没有复习的原因很快就忘了。下次想看的时候还要重新看一遍，非常的浪费时间。

下面以IOC为例说明下我是怎么看的，供参考。

## IOC ##

### 入口：ApplicationContext ###

在研究源码时首先要找到一个入口，这个入口怎么选择可以自己定，当一定要和你需要看的模块有关联。

比如在IOC中，首先我们想到创建容器是在什么过程？

在程序启动的时候就创建了，而且在启动过程中大多数的bean实例就被注入了。

那问题来了，在启动的时候是从那个类开始的呢？熟悉spring的应该都知道我们平时在做单元测试时如果要获取bean实例，一个是通过注解，另外我们还可以通过构建一个ApplicationContext来获取：

` ApplicationContext applicationContext = new ClassPathXmlApplicationContext( "classpath*:application.xml" ); XxxService xxxService = applicationContext.getBean( "xxxService" ); 复制代码`

在实例化ApplicationContext后既可以获取bean，那么实例化的这个过程就相当于启动的过程了，所以我们可以将ApplicationContext当成我们的入口。

### ApplicationContext是什么 ###

首先我们要明白的事我们平时一直说的IOC容器在Spring中实际上指的就是 ` ApplicationContext` 。

如果有看过我之前手写Spring系列文章的同学肯定知道在当时文章中充当ioc容器的是BeanFactory，每当有bean需要注入时都是由BeanFactory保存，取bean实例时也是从BeanFactory中获取。

那为什么现在要说ApplicationContext才是IOC容器呢？

因为在spring中BeanFactory实际上是被隐藏了的。ApplicationContext是对BeanFactory的一个封装，也提供了获取bean实例等功能。因为BeanFactory本身的能力实在太强，如果可以让我们随便使用可能会对spring功能的运行造成破坏。于是就封装了一个提供查询ioc容器内容的ApplicationContext供我们使用。

如果项目中需要用到ApplicationContext，可以直接使用spring提供的注解获取：

` @Autowired private ApplicationContext applicationContext; 复制代码`

#### 如何使用ApplicationContext ####

如果我们要使用ApplicationContext可以通过new该类的一个实例即可，定义好相应的xml文件。然后通过下面的代码即可：

` @Test public void testClassPathXmlApplicationContext () { //1.准备配置文件，从当前类加载路径中获取配置文件 //2.初始化容器 ApplicationContext applicationContext = new ClassPathXmlApplicationContext( "classpath*:application.xml" ); //2、从容器中获取Bean HelloApi helloApi = applicationContext.getBean( "hello" , HelloApi.class); //3、执行业务逻辑 helloApi.sayHello(); } 复制代码`

#### ApplicationContext的体系 ####

了解一个类，首先可以来看看它的继承关系来了解其先天的提供哪些功能。然后在看其本身又实现了哪些功能。

![ApplicationContext继承体系](https://user-gold-cdn.xitu.io/2019/6/6/16b2a54bbe1b0234?imageView2/0/w/1280/h/960/ignore-error/1)

上图中继承关系从左至右简要介绍其功能。

* ApplicationEventPublisher：提供发布监听事件的功能，接收一个监听事件实体作为参数。需要了解的可以通过这篇文章： [事件监听]( https://link.juejin.im?target=https%3A%2F%2Fwww.xttblog.com%2F%3Fp%3D3584 )
* ResourcePatternResolver：用于解析一些传入的文件路径(比如ant风格的路径)，然后将文件加载为resource。
* HierarchicalBeanFactory：提供父子容器关系，保证子容器能访问父容器，父容器无法访问子容器。
* ListableBeanFactory：继承自BeanFactory，提供访问IOC容器的方法。
* EnvironmentCapable：获取环境变量相关的内容。
* MessageSource：提供国际化的message的解析

### 配置文件的加载 ###

Spring中每一个功能都是很大的一个工程，所以在阅读时也要分为多个模块来理解。要理解IOC容器，我们首先需要了解spring是如何加载配置文件的。

#### 纵览大局 ####

idea或者eclipse提供了一个很好的功能就是能在调试模式下看到整个流程的调用链。利用这个功能我们可以直接观察到某一功能实现的整体流程，也方便在阅读代码时在不同类切换。

以加载配置文件为例，这里给出整个调用链。

![配置文件加载流程](https://user-gold-cdn.xitu.io/2019/6/6/16b2a54bac40d3eb?imageView2/0/w/1280/h/960/ignore-error/1)

上图中下面的红框是我们写的代码，即就是我们应该开始的地方。下面的红框就是加载配置文件结束的地方。中间既是整体流程的实现过程。在阅读配置文件加载的源码时我们只需要关心这一部分的内容即可。

需要知道的是这里展示出来的仅仅只是跟这个过程密切相关的一些方法。实际上在这个过程中还有需要的方法被执行，只不过执行完毕后方法栈弹出所以不显示在这里。不过大多数方法都是在为这个流程做准备，所以基本上我们也不用太在意这部分内容

#### refresh() ####

前面的关于 ` ClassPathXmlApplicationContext` 的构造函数部分没有啥好说的，在构造函数中调用了一个方法 ` AbstractApplicationContext#refresh` 。该方法非常重要，在创建IOC容器的过程中该方法基本上是全程参与。主要功能为用于加载配置或这用于刷新已经加载完成的容器配置。通过该方法可以在运行过程中动态的加入配置文件等：

` ClassPathXmlApplicationContext ctx = new ClassPathXmlApplicationContext(); ctx.setConfigLocation( "application-temp.xml" ); ctx.refresh(); 复制代码`

AbstractApplicationContext#refresh

` public void refresh() throws BeansException, IllegalStateException { synchronized (this.startupShutdownMonitor) { prepareRefresh(); ConfigurableListableBeanFactory beanFactory = obtainFreshBeanFactory(); // more statement ... } } 复制代码`

这里将于当前功能不相关的部分删除掉了，可以看到进入方法后就会进入一个同步代码块。这是为了防止在同一时间有多个线程开始创建IOC容器造成重复实例化。

` prepareRefresh();` 方法主要用于设置一些日志相关的信息，比如容器启动时间用于计算启动容器整体用时，以及设置一些变量用来标识当前容器已经被激活，后续不会再进行创建。

` obtainFreshBeanFactory();` 方法用于获取一个BeanFactory，在这一过程中便会加载配置文件和解析用于生成一个BeanFactory。

#### refreshBeanFactory ####

refreshBeanFactory方法有obtainFreshBeanFactory方法调用

` protected final void refreshBeanFactory() throws BeansException { if (hasBeanFactory()) { destroyBeans(); closeBeanFactory(); } try { DefaultListableBeanFactory beanFactory = createBeanFactory(); beanFactory.setSerializationId(getId()); customizeBeanFactory(beanFactory); loadBeanDefinitions(beanFactory); synchronized (this.beanFactoryMonitor) { this.beanFactory = beanFactory; } }catch (IOException ex) { throw new ApplicationContextException( "I/O error parsing bean definition source for " + getDisplayName(), ex); } } 复制代码`

该方法首先判断是否已经实例化好BeanFactory，如果已经实例化完成则将已经实例化好的BeanFactory销毁。

然后通过new关键字创建一个BeanFactory的实现类实例，设置好相关信息。 ` customizeBeanFactory(beanFactory)` 方法用于设置是否运行当beanName重复是修改bean的名称(allowBeanDefinitionOverriding)和是否运行循环引用(allowCircularReferences)。

` loadBeanDefinitions(beanFactory)` 方法既是开始加载bean定义的方法。当BeanFactory在加载完所有配置信息后创建，然后将创建好的BeanFactory赋值给当前context下的BeanFactory。

#### loadBeanDefinitions ####

` protected void loadBeanDefinitions(DefaultListableBeanFactory beanFactory) throws BeansException, IOException { XmlBeanDefinitionReader beanDefinitionReader = new XmlBeanDefinitionReader(beanFactory); beanDefinitionReader.setEnvironment(this.getEnvironment()); beanDefinitionReader.setResourceLoader(this); beanDefinitionReader.setEntityResolver(new ResourceEntityResolver(this)); initBeanDefinitionReader(beanDefinitionReader); loadBeanDefinitions(beanDefinitionReader); } 复制代码`

` loadBeanDefinitions` 见名知意其就是用于加载bean定义的方法，在 ` AbstractXmlApplicationContext` 中定义了一系列该方法的重载方法。上面的方法主要便是引入 ` XmlBeanDefinitionReader` 。 ` XmlBeanDefinitionReader` 是一个用于读取xml文件中bean定义的类，其提供了一些诸如BeanFactory和BeanDefinitionRegistery类的属性以供使用。但其实真正的读取操作并没该类完成，其也是作为一个代理存在。

在spring中如果是完成一些类似操作的类的命名都是有迹可循的，比如这里读取xml文件就是以reader结尾，类似的读取注解中bean定义也有如 ` AnnotatedBeanDefinitionReader` 。如果需要向类中注入一些Spring中的bean，一般是以Aware结尾如 ` BeanFactoryAware` 等。所以在阅读spring源码时如果遇到这样的类很多时候我们可以直接根据其命名了解其大概的实现方式。

` public int loadBeanDefinitions(String location, Set<Resource> actualResources) throws BeanDefinitionStoreException { ResourceLoader resourceLoader = getResourceLoader(); if (resourceLoader == null) { throw new BeanDefinitionStoreException( "Cannot import bean definitions from location [" + location + "]: no ResourceLoader available" ); } if (resourceLoader instanceof ResourcePatternResolver) { try { Resource[] resources = ((ResourcePatternResolver) resourceLoader).getResources(location); int loadCount = loadBeanDefinitions(resources); if (actualResources != null) { for (Resource resource : resources) { actualResources.add(resource); } } //logging return loadCount; }catch (IOException ex) { throw new BeanDefinitionStoreException( "Could not resolve bean definition resource pattern [" + location + "]" , ex); } } else { Resource resource = resourceLoader.getResource(location); int loadCount = loadBeanDefinitions(resource); if (actualResources != null) { actualResources.add(resource); } //logging return loadCount; } } 复制代码`

上面代码是 ` loadBeanDefinitions` 的一个实现类，该方法的主要注意点在于三个地方。

一个是方法中抛出的两个异常，前一个异常时因为ResourceLoader定义的问题，一般来说不需要我们关注。后一个就是配置文件出错了，可能是因为文件本身xml格式出错或者是由于循环引用等原因，具体的原因也会通过日志打印。我们需要对这些异常信息有印象，也不用刻意去记，遇到了能快速定位问题即可。

另一个就是代码中的一个 ` if(){}else{}` 语句块，判断语句快中都是用于解析配置文件，不同之处在于if中支持解析匹配风格的location，比如 ` classpath*:spring.xml` 这种，该功能的实现由 ` ResourcePatternResolver` 提供， ` ResourcePatternResolver` 对 ` ResourceLoader` 的功能进行了增强，支持解析ant风格等模式的location。而else中仅仅只能解析指定的某一文件如 ` spring.xml` 这种。实际上在 ` ApplicationContext` 中实现了 ` ResourcePatternResolver` ，如果也按照 ` spring.xml` 配置，也是按照 ` ResourceLoader` 提供的解析方式解析。

最后一处就是 ` Resource` 类， ` Resource` 是spring为了便于加载文件而特意设计的接口。其提供了大量对传入的location操作方法，支持对不同风格的location(比如文件系统或者ClassPath)。其本身还有许多不同的实现类，本质上是对 ` File，URL，ClassPath` 等不同方式获取location的一个整合，功能十分强大。即使我们的项目不依赖spring，如果涉及到Resource方面的操作也可以使用Spring中的Resource。

` public int loadBeanDefinitions(EncodedResource encodedResource) throws BeanDefinitionStoreException { // log and assert Set<EncodedResource> currentResources = this.resourcesCurrentlyBeingLoaded.get(); if (currentResources == null) { currentResources = new HashSet<EncodedResource>(4); this.resourcesCurrentlyBeingLoaded.set(currentResources); } if (!currentResources.add(encodedResource)) { throw new BeanDefinitionStoreException( "Detected cyclic loading of " + encodedResource + " - check your import definitions!" ); } try { InputStream inputStream = encodedResource.getResource().getInputStream(); try { InputSource inputSource = new InputSource(inputStream); if (encodedResource.getEncoding() != null) { inputSource.setEncoding(encodedResource.getEncoding()); } return do LoadBeanDefinitions(inputSource, encodedResource.getResource()); } finally { inputStream.close(); } }catch (IOException ex) { throw new BeanDefinitionStoreException( "IOException parsing XML document from " + encodedResource.getResource(), ex); }finally { currentResources.remove(encodedResource); if (currentResources.isEmpty()) { this.resourcesCurrentlyBeingLoaded.remove(); } } } 复制代码`

该方法依旧是loadBeanDefinitions的重载方法。

方法传入一个EncodedResource，该类可以通过制定的字符集对Resource进行编码，利于统一字符编码格式。

然后try语句块上面的代码也是比较重要的，主要功能便是判断是否有配置文件存在循环引用的问题。

循环应用问题出现在比如我加载一个配置文件 ` application.xml` ，但是在该文件内部又通过 ` import` 标签引用了自身。在解析到 ` import` 时会加载 ` import` 指定的文件。这样就造成了一个死循环，如果不解决程序就会永远启动不起来。

解决的方法也很简单，通过一个 ` ThreadLocal` 记录下当前正在加载的配置文件名称(包括路径)，每一次在加载新的配置文件时从 ` ThreadLocal` 中取出放入到set集合中，通过set自动去重的特性判断是否循环加载了。当一个文件加载完成后，就从 ` ThreadLocal` 中去掉(finally)。这里是判断xml文件时否重复加载，而在spring中判断bean是否循环引用是虽然实现上有点差别，但基本思想也是这样的。

#### doLoadBeanDefinitions(InputSource, Resource) ####

到了这一步基本上才算是真正开始解析了。该方法虽然代码行数较多，但是大多都是异常处理，异常代码已经省略。我们需要关注的就是try中的两句代码。

` protected int do LoadBeanDefinitions(InputSource inputSource, Resource resource) throws BeanDefinitionStoreException { try { Document doc = do LoadDocument(inputSource, resource); return registerBeanDefinitions(doc, resource); }catch (Exception ex) { //多个catch语句块 } } 复制代码`

` Document doc = doLoadDocument(inputSource, resource)` 就是读取配置文件并将其内容解析为一个Document的过程。解析xml一般来说并不需要我们特别的去掌握，稍微有个了解即可，spring这里使用的解析方式为Sax解析，有兴趣的可以直接搜索相关文章，这里不进行介绍。下面的 ` registerBeanDefinitions` 才是我们需要关注的地方。

#### registerBeanDefinitions(Document, Resource) ####

` public int registerBeanDefinitions(Document doc, Resource resource) throws BeanDefinitionStoreException { BeanDefinitionDocumentReader documentReader = createBeanDefinitionDocumentReader(); documentReader.setEnvironment(getEnvironment()); int countBefore = getRegistry().getBeanDefinitionCount(); documentReader.registerBeanDefinitions(doc, createReaderContext(resource)); return getRegistry().getBeanDefinitionCount() - countBefore; } 复制代码`

在进入该方法后首先创建了一个 ` BeanDefinitionDocumentReader` 的实例，这和之前的用于读取xml的reader类一样，只不过该类是用于从xml文件中读取 ` BeanDefinition` 。

##### Environment #####

在上面的代码中给Reader设置了Environment，这里谈一下关于Environment。

Environment是对spring程序中涉及到环境有关的一个描述集合，主要分为profile和properties。

profile是一组bean定义的集合，通过profile可以指定不同的配置文件用以在不同的环境中，如测试环境，生产环境的配置分开。在部署时只需要配置好当前所处环境值即可按不同分类加载不同的配置。

profile支持xml配置和注解的方式。

` <?xml version= "1.0" encoding= "UTF-8" ?> <beans xmlns= "http://www.springframework.org/schema/beans" //more > <!-- 定义开发环境的profile --> <beans profile= "development" > <!-- 只扫描开发环境下使用的类 --> <context:component-scan base-package= "com.demo.service" /> <!-- 加载开发使用的配置文件 --> <util:properties id= "config" location= "classpath:dev/config.properties" /> </beans> <!-- 定义生产环境的profile --> <beans profile= "produce" > <!-- 只扫描生产环境下使用的类 --> <context:component-scan base-package= "com.demo.service" /> <!-- 加载生产使用的配置文件 --> <util:properties id= "config" location= "classpath:produce/config.properties" /> </beans> </beans> 复制代码`

也可以通过注解配置：

` @Service @Profile( "dev" ) public class ProductRpcImpl implements ProductRpc { public String productBaseInfo(int id) { return "success" ; } } 复制代码`

然后在启动时根据传入的环境值加载相应的配置。

properties是一个很宽泛的定义，其来源很多如properties文件，JVM系统变量，系统环境变量，JNDI，servlet上下文参数，Map等。spring会读取这些配置并在environment接口中提供了方便对其进行操作的方法。

总之就是设计到跟环境有关的直接来找Environment即可。

#### handler ####

代码接着往下走， ` documentReader.registerBeanDefinitions(doc, createReaderContext(resource))` 这一步很明显就是从解析好的document对象中读取BeanDefinition的过程，但是在此之前我们先要关注一下 ` createReaderContext(resource)` 方法。

先来看一个XML文件。

` <beans xmlns= "http://www.springframework.org/schema/beans" xmlns:xsi= "http://www.w3.org/2001/XMLSchema-instance" xmlns:context= "http://www.springframework.org/schema/context" xmlns:mvc= "http://www.springframework.org/schema/mvc" > </beans> 复制代码`

上面是xml中根元素定义部分，可能平时并没有太多人注意。其属性中的xmlns是XML NameSpace的缩写。namespace的作用主要是为了防止在xml定义的节点存在冲突的问题。比如上面声明了mvc的namespace： ` xmlns:mvc="http://www.springframework.org/schema/mvc"` 。在xml文件中我们就可以使用mvc了：

` <mvc:annotation-driven /> <mvc:default-servlet-handler/> 复制代码`

而实际上在spring中还根据上面定义的namespace来准备了各自的处理类。这里因为解析过程就是将xml定义的每一个节点取出根据配置好的属性和值来初始化或注册bean，为了保证代码可读性和明确的分工，每一个namespace通过一个专有的handler来处理。

跟踪 ` createReaderContext(resource)` 方法，最终来到 ` DefaultNamespaceHandlerResolver` 类的构造方法中。

![handler匹配](https://user-gold-cdn.xitu.io/2019/6/6/16b2a54bb423bc00?imageView2/0/w/1280/h/960/ignore-error/1)

` public DefaultNamespaceHandlerResolver(ClassLoader classLoader) { //DEFAULT_HANDLER_MAPPINGS_LOCATION = "META-INF/spring.handlers" this(classLoader, DEFAULT_HANDLER_MAPPINGS_LOCATION); } 复制代码`

可以看到默认的handler是通过一个本地文件来进行映射的。该文件存在于被依赖jar包下的META-INF文件夹下的spring.handlers文件中。

![handler mapping](https://user-gold-cdn.xitu.io/2019/6/6/16b2a54bafb7ea52?imageView2/0/w/1280/h/960/ignore-error/1)

` http\://www.springframework.org/schema/c=org.springframework.beans.factory.xml.SimpleConstructorNamespaceHandler http\://www.springframework.org/schema/p=org.springframework.beans.factory.xml.SimplePropertyNamespaceHandler http\://www.springframework.org/schema/util=org.springframework.beans.factory.xml.UtilNamespaceHandler 复制代码`

这里只是展示了beans包下的映射文件，其他如aop包，context包下都有相应的映射文件。通过读取这些配置文件来映射相应的处理类。在解析xml时会根据使用的namespace前缀使用对应的handler类解析。这种实现机制其实就是所谓的SPI(Service Provider Interface)，目前很多的应用都在实现过程中使用了spi，如dubbo，mysql的jdbc实现等，有兴趣的可以取了解一下。

#### doRegisterBeanDefinitions(Element) ####

到这一步中间省略了一个方法，很简单没有分析的必要。

` protected void do RegisterBeanDefinitions(Element root) { BeanDefinitionParserDelegate parent = this.delegate; this.delegate = createDelegate(getReaderContext(), root, parent); if (this.delegate.isDefaultNamespace(root)) { String profileSpec = root.getAttribute(PROFILE_ATTRIBUTE); if (StringUtils.hasText(profileSpec)) { String[] specifiedProfiles = StringUtils.tokenizeToStringArray( profileSpec, BeanDefinitionParserDelegate.MULTI_VALUE_ATTRIBUTE_DELIMITERS); if (!getReaderContext().getEnvironment().acceptsProfiles(specifiedProfiles)) { return ; } } } preProcessXml(root); parseBeanDefinitions(root, this.delegate); postProcessXml(root); this.delegate = parent; } 复制代码`

##### delegate #####

这里的delegate是对BeanDefinitionParse功能的代理，提供了一些支持解析过程的方法。我们可以看到上面有一个重新创建delegate同时又将之前的delegate保存的代码。注释上说是为了防止嵌套的beans标签递归操作导致出错，但是注释后面又说这并不需要这样处理，这个操作真的看不懂了，实际上我认为即使递归应该也是没有影响的。还是我理解错了？

创建好delegate后下面的if语句块就是用来判断当前加载的配置文件是否是当前使用的profile指定的配置文件。上面在介绍Environment的时候已经介绍过来，如果这里加载的配置文件和profile指定的不符则直接结束。

` preProcessXml(root)` 方法是一个空实现，并且当前spring框架中好像也没有对这个方法实现的，这里不管了。同理还有下面的 ` postProcessXml(root)` 。

#### parseBeanDefinitions(Element, BeanDefinitionParserDelegate) ####

` protected void parseBeanDefinitions(Element root, BeanDefinitionParserDelegate delegate) { if (delegate.isDefaultNamespace(root)) { NodeList nl = root.getChildNodes(); for (int i = 0; i < nl.getLength(); i++) { Node node = nl.item(i); if (node instanceof Element) { Element ele = (Element) node; if (delegate.isDefaultNamespace(ele)) { parseDefaultElement(ele, delegate); } else { delegate.parseCustomElement(ele); } } } } else { delegate.parseCustomElement(root); } } 复制代码`

该方法的理解起来并不难，首先读取根节点(beans)下的所有子节点，然后对这些节点进行解析。这里需要注意的即使是对节点的解析也有一个判断语句。

主要来看一下 ` delegate.isDefaultNamespace(ele)` ,

` public boolean isDefaultNamespace(String namespaceUri) { //BEANS_NAMESPACE_URI = "http://www.springframework.org/schema/beans" return (!StringUtils.hasLength(namespaceUri) || BEANS_NAMESPACE_URI.equals(namespaceUri)); } 复制代码`

也就是说beans命名空间下的标签一个解析方法，而另外的标签一个解析方法。

#### parseDefaultElement(Element, BeanDefinitionParserDelegate) ####

` private void parseDefaultElement(Element ele, BeanDefinitionParserDelegate delegate) { if (delegate.nodeNameEquals(ele, IMPORT_ELEMENT)) {//import importBeanDefinitionResource(ele); } else if (delegate.nodeNameEquals(ele, ALIAS_ELEMENT)) {// alias processAliasRegistration(ele); } else if (delegate.nodeNameEquals(ele, BEAN_ELEMENT)) {//bean processBeanDefinition(ele, delegate); } else if (delegate.nodeNameEquals(ele, NESTED_BEANS_ELEMENT)) { // recurse do RegisterBeanDefinitions(ele); } } 复制代码`

可以看到对于每一个标签，都提供了一个方法来进行解析，最后一个方法用于对嵌套标签进行解析，这里以bean标签的解析为例。

` protected void processBeanDefinition(Element ele, BeanDefinitionParserDelegate delegate) { BeanDefinitionHolder bdHolder = delegate.parseBeanDefinitionElement(ele); if (bdHolder != null) { //用于将一些属性值塞进BeanDefinition中如lazy-init //以及子节点中的值 如bean节点下的property bdHolder = delegate.decorateBeanDefinitionIfRequired(ele, bdHolder); try { BeanDefinitionReaderUtils.registerBeanDefinition(bdHolder, getReaderContext().getRegistry()); } catch (BeanDefinitionStoreException ex) { getReaderContext().error( "Failed to register bean definition with name '" + bdHolder.getBeanName() + "'" , ele, ex); } getReaderContext().fireComponentRegistered(new BeanComponentDefinition(bdHolder)); } } 复制代码`

Holder对象也是spring中的一个系列性的对象，主要就是对某一些实例进行包装，比如 ` BeanDefinitionHolder` 就是对 ` BeanDefinition` 进行包装，主要就是持有BeanDefinition以及它的名称和别名等(BeanDefinition为接口，无法提供名称等属性)。

` public static void registerBeanDefinition( BeanDefinitionHolder definitionHolder, BeanDefinitionRegistry registry) throws BeanDefinitionStoreException { String beanName = definitionHolder.getBeanName(); registry.registerBeanDefinition(beanName, definitionHolder.getBeanDefinition()); String[] aliases = definitionHolder.getAliases(); if (aliases != null) { for (String alias : aliases) { registry.registerAlias(beanName, alias ); } } } 复制代码`

接下来的 ` BeanDefinitionReaderUtils.registerBeanDefinition(bdHolder, getReaderContext().getRegistry())` 可以算是我们本次目标最重要的一步了，前面所有的流程都是给这一步铺垫。通过该方法将解析出来的BeanDefinition注册到容器中，方便实例化。

方法接收两个参数holder和registry，如果有看过我手写系列的文章(IOC篇)应该知道，当时为了将bean定义和容器关联起来以及为了将 ` beanfactory` 的功能简化，所以我们定义了一个 ` BeanDefinitionRegistry` 接口用于将 ` BeanDefinition` 注册到容器中和从容器中取 ` BeanDefinition` ，这里的registry功能也是一样的。( ` BeanDefinitionRegistry` 被 ` DefaultListableBeanFactory` 实现，而 ` DefaultListableBeanFactory` 实际就是容器)

而且可以看到实际上这里也是通过beanName来区分BeanDefinition是否重复(实际上肯定是我仿的spring的(笑))，只不过为了运行名称相同的BeanDefinition注册提供了alias，之前在实现ioc时没有实现这一步。

在 ` processBeanDefinition` 方法的最后一步实际上是注册了一个listener，在一个 ` BeanDefinition` 被注册后触发，只不过上spring中实际触发方法是一个空方法，如果我们需要在 ` BeanDefinition` 注册完成后做一些什么工作可以直接继承 ` EmptyReaderEventListener` 后实现 ` componentRegistered(componentDefinition)` 方法即可。

到这里基本上关于BeanDefinition的加载就完成了，后面就是重复上面的流程加载多个配置文件。

### 小结 ###

本节主要介绍了我关于学习spring源码的一些方法，以及以spring的BeanDefinition的加载为例分析了其整体的流程，希望对大家能有所帮助。还要提的是spring源码很复杂，如果只是开断点一路调试下去肯定是不够的，看的过程中需要多做笔记。由于该文章内容较多以及本人水平问题，文章中可能会存在错误，如果有可以指出来方便修改。