# 从源码解读Spring的IOC #

### 概念 ###

IOC（ **I** nversion **O** f **C** ontrol），即控制反转，或者说DI（ **D** ependency **I** njection），依赖注入，都属于Spring中的一种特点，可以统称为IOC

* 控制反转，即控制权的反转，也就是说将内置对象创建的控制权交给第三方容器，而不是本身的对象
* 依赖注入，即将依赖对象进行注入，也就是说不主动创建对象，而是通过第三方容器将依赖的对象注入进来

不管是IOC还是DI，它们都有一个共同的特点，即通过 **第三方容器** 来管理对象的创建、初始化、注入、销毁，在Spring中，这些容器中的对象统称为 **bean** ，使用的时候只需要通过xml或Java的配置就能够很方便地将一个对象声明为容器中的bean，并且可以通过@Autowired注解等方法将这些bean注入到需要的地方，我们接下来就要深入地了解一下Spring中IOC到底是怎么实现的

### BeanFactory ###

顾名思义，BeanFactory就是Bean的工厂，也就是Bean的容器，BeanFactory作为容器接口，我们暂不关心它的实现类，先了解一下它本身的特性，点进BeanFactory的源码，我们只看以下这几个重要的方法，其余方法用到的时候再说

` /** 通过name获取bean实例 */ Object getBean (String name) throws BeansException ; /** 通过name和对象类型获取bean实例 */ <T> T getBean (String name, Class<T> requiredType) throws BeansException ; /** 得到bean的别名，如果通过别名索引，则原名也会被检索出 */ String[] getAliases(String name); 复制代码`

这几个方法看起来很简单，别急，这只是在接口中的定义，Spring提供了很多BeanFactory的实现类，比如ApplicationContext等

### BeanDefinition ###

bean当然不能用普通的对象来描述，在spring中，bean被封装成BeanDefinition，如下：

![图1](https://user-gold-cdn.xitu.io/2019/6/4/16b228530d0faf67?imageView2/0/w/1280/h/960/ignore-error/1)

### Bean资源的加载过程 ###

资源的加载，也可以是认为是容器的初始化，可以分为以下三个部分：

* 定位资源
* 载入资源
* 注册资源

比如XmlWebApplicationContext就是从xml文件中加载资源，我们这里以ClassPathXmlApplicationContext为例，了解一下xml文件中的配置是怎么加载到Spring容器中的

首先是构造方法，如下：

` /** * @param configLocations 资源路径 * @param refresh 是否自动刷新容器 * @parent parent 容器的父类 */ public ClassPathXmlApplicationContext ( String[] configLocations, boolean refresh, @Nullable ApplicationContext parent) throws BeansException { super (parent); setConfigLocations(configLocations); if (refresh) { refresh(); } } 复制代码`

先来看这个super(parent)，我们一路找上去，发现每一层都调用了super(parent)，直到进入AbstractApplicationContext类中，发现调用了一个this()方法，这个方法详细如下：

` public AbstractApplicationContext () { this.resourcePatternResolver = getResourcePatternResolver(); } 复制代码`

从字面上理解，应该是类似设置资源解析器之类的方法，我们进入这个方法，发现其实际上创建了一个PathMatchingResourcePatternResolver对象，同时设置我们的最顶层容器为resourceLoader资源加载器，看到这里就差不多了解了，super(parent)实际上就是设置了bean的资源加载器

我们接着看setConfigLocations(configLocations)方法，源码如下：

` public void setConfigLocations (@Nullable String... locations) { if (locations != null ) { Assert.noNullElements(locations, "Config locations must not be null" ); this.configLocations = new String[locations.length]; for ( int i = 0 ; i < locations.length; i++) { this.configLocations[i] = resolvePath(locations[i]).trim(); } } else { this.configLocations = null ; } } 复制代码`

这个方法是继承而来的，是AbstractRefreshableConfigApplicationContext的一个方法，在这个方法内部，设置了configLocations的值为资源路径（进行环境变量填补充并去除空格），可以理解为对资源进行定位

也就是说，容器在创建出来时，做了以下两件事（不包括刷新容器操作）：

* 设置资源解析器
* 设置资源路径，进行资源定位

然后我们再来看这个可选的refresh()方法，这是从AbstractApplicationContext继承而来的方法，源码如下：

` @Override public void refresh () throws BeansException, IllegalStateException { synchronized ( this.startupShutdownMonitor) { // 获取当前时间，同时设置同步标识，避免多线程下冲突 prepareRefresh(); // 实际调用了子类的refreshBeanFactory方法，同时返回子类的beanFactory ConfigurableListableBeanFactory beanFactory = obtainFreshBeanFactory(); // 设置容器属性 prepareBeanFactory(beanFactory); try { // 为子类beanFactory指定BeanPost事件处理器 postProcessBeanFactory(beanFactory); // 调用注册为bean的事件处理器 invokeBeanFactoryPostProcessors(beanFactory); // 注册BeanPost事件处理器，用于监听容器创建 registerBeanPostProcessors(beanFactory); // 初始化消息源 initMessageSource(); // 初始化事件传播器 initApplicationEventMulticaster(); // 在特定的子类中初始化其他特殊的bean onRefresh(); // 检查并注册监听器 registerListeners(); // 初始化剩余的单例 finishBeanFactoryInitialization(beanFactory); // 最后一步：初始化容器生命周期处理器，并发布容器生命周期事件 finishRefresh(); } catch (BeansException ex) { if (logger.isWarnEnabled()) { logger.warn( "Exception encountered during context initialization - " + "cancelling refresh attempt: " + ex); } // 销毁创建的单例，以避免悬置资源 destroyBeans(); // 重置同步标识 cancelRefresh(ex); throw ex; } finally { // 因为不需要单例bean中元数据，所以重置spring的自检缓存 resetCommonCaches(); } } } 复制代码`

配合注释，就能差不多了解了执行过程，实际就是初始化并注册一系列处理器和监听器的过程，有人可能会发现，怎么没有加载资源的过程，别急，我们进入obtainFreshBeanFactory()方法，其中有一个refreshBeanFactory()方法，我们点开AbstractRefreshableApplicationContext中的实现：

` @Override protected final void refreshBeanFactory () throws BeansException { if (hasBeanFactory()) { destroyBeans(); closeBeanFactory(); } try { DefaultListableBeanFactory beanFactory = createBeanFactory(); beanFactory.setSerializationId(getId()); customizeBeanFactory(beanFactory); loadBeanDefinitions(beanFactory); synchronized ( this.beanFactoryMonitor) { this.beanFactory = beanFactory; } } catch (IOException ex) { throw new ApplicationContextException( "I/O error parsing bean definition source for " + getDisplayName(), ex); } } 复制代码`

发现了吗，这里有一个loadBeanDefinitions()方法，会根据选用的xml解析还是注解解析调用子类中的方法，再接下来的解析过程就不是我们分析的重点了，如果以后有时间，我会再写一篇来专门分析解析过程的文章

到这里，整个加载过程就清晰了

### 依赖注入的过程 ###

一开始，我们就介绍了getBean(name)方法，那么我们接下来，就要详细的来进行分析这个方法到底是怎么把我们需要的bean创建并交给我们的

这个方法有两种常见的实现，AbstractBeanFactory和AbstractApplicationContext，而实际上AbstractApplicationContext也是调用了AbstractBeanFactory的方法，所以我们就只看AbstractBeanFactory即可

在这个方法内部调用了doGetBean方法，我们进入方法内部，如下：

` protected <T> T doGetBean ( final String name, @Nullable final Class<T> requiredType, @Nullable final Object[] args, boolean typeCheckOnly) throws BeansException { final String beanName = transformedBeanName(name); Object bean; // Eagerly check singleton cache for manually registered singletons. Object sharedInstance = getSingleton(beanName); if (sharedInstance != null && args == null ) { if (logger.isTraceEnabled()) { if (isSingletonCurrentlyInCreation(beanName)) { logger.trace( "Returning eagerly cached instance of singleton bean '" + beanName + "' that is not fully initialized yet - a consequence of a circular reference" ); } else { logger.trace( "Returning cached instance of singleton bean '" + beanName + "'" ); } } bean = getObjectForBeanInstance(sharedInstance, name, beanName, null ); } else { // Fail if we're already creating this bean instance: // We're assumably within a circular reference. if (isPrototypeCurrentlyInCreation(beanName)) { throw new BeanCurrentlyInCreationException(beanName); } // Check if bean definition exists in this factory. BeanFactory parentBeanFactory = getParentBeanFactory(); if (parentBeanFactory != null && !containsBeanDefinition(beanName)) { // Not found -> check parent. String nameToLookup = originalBeanName(name); if (parentBeanFactory instanceof AbstractBeanFactory) { return ((AbstractBeanFactory) parentBeanFactory).doGetBean( nameToLookup, requiredType, args, typeCheckOnly); } else if (args != null ) { // Delegation to parent with explicit args. return (T) parentBeanFactory.getBean(nameToLookup, args); } else if (requiredType != null ) { // No args -> delegate to standard getBean method. return parentBeanFactory.getBean(nameToLookup, requiredType); } else { return (T) parentBeanFactory.getBean(nameToLookup); } } if (!typeCheckOnly) { markBeanAsCreated(beanName); } try { final RootBeanDefinition mbd = getMergedLocalBeanDefinition(beanName); checkMergedBeanDefinition(mbd, beanName, args); // Guarantee initialization of beans that the current bean depends on. String[] dependsOn = mbd.getDependsOn(); if (dependsOn != null ) { for (String dep : dependsOn) { if (isDependent(beanName, dep)) { throw new BeanCreationException(mbd.getResourceDescription(), beanName, "Circular depends-on relationship between '" + beanName + "' and '" + dep + "'" ); } registerDependentBean(dep, beanName); try { getBean(dep); } catch (NoSuchBeanDefinitionException ex) { throw new BeanCreationException(mbd.getResourceDescription(), beanName, "'" + beanName + "' depends on missing bean '" + dep + "'" , ex); } } } // Create bean instance. if (mbd.isSingleton()) { sharedInstance = getSingleton(beanName, () -> { try { return createBean(beanName, mbd, args); } catch (BeansException ex) { // Explicitly remove instance from singleton cache: It might have been put there // eagerly by the creation process, to allow for circular reference resolution. // Also remove any beans that received a temporary reference to the bean. destroySingleton(beanName); throw ex; } }); bean = getObjectForBeanInstance(sharedInstance, name, beanName, mbd); } else if (mbd.isPrototype()) { // It's a prototype -> create a new instance. Object prototypeInstance = null ; try { beforePrototypeCreation(beanName); prototypeInstance = createBean(beanName, mbd, args); } finally { afterPrototypeCreation(beanName); } bean = getObjectForBeanInstance(prototypeInstance, name, beanName, mbd); } else { String scopeName = mbd.getScope(); final Scope scope = this.scopes.get(scopeName); if (scope == null ) { throw new IllegalStateException( "No Scope registered for scope name '" + scopeName + "'" ); } try { Object scopedInstance = scope.get(beanName, () -> { beforePrototypeCreation(beanName); try { return createBean(beanName, mbd, args); } finally { afterPrototypeCreation(beanName); } }); bean = getObjectForBeanInstance(scopedInstance, name, beanName, mbd); } catch (IllegalStateException ex) { throw new BeanCreationException(beanName, "Scope '" + scopeName + "' is not active for the current thread; consider " + "defining a scoped proxy for this bean if you intend to refer to it from a singleton" , ex); } } } catch (BeansException ex) { cleanupAfterBeanCreationFailure(beanName); throw ex; } } // Check if required type matches the type of the actual bean instance. if (requiredType != null && !requiredType.isInstance(bean)) { try { T convertedBean = getTypeConverter().convertIfNecessary(bean, requiredType); if (convertedBean == null ) { throw new BeanNotOfRequiredTypeException(name, requiredType, bean.getClass()); } return convertedBean; } catch (TypeMismatchException ex) { if (logger.isTraceEnabled()) { logger.trace( "Failed to convert bean '" + name + "' to required type '" + ClassUtils.getQualifiedName(requiredType) + "'" , ex); } throw new BeanNotOfRequiredTypeException(name, requiredType, bean.getClass()); } } return (T) bean; } 复制代码`

整个方法相当之长，我们分部分来看，先来看第一部分：

` // 转换为规范名称（主要针对别名） final String beanName = transformedBeanName(name); Object bean; // 检查缓存，避免重复创建单例 Object sharedInstance = getSingleton(beanName); // 如果不为空，就返回缓存中的单例 if (sharedInstance != null && args == null ) { // 如果开启了trace日志，就根据当前的状态打印日志 if (logger.isTraceEnabled()) { if (isSingletonCurrentlyInCreation(beanName)) { logger.trace( "Returning eagerly cached instance of singleton bean '" + beanName + "' that is not fully initialized yet - a consequence of a circular reference" ); } else { logger.trace( "Returning cached instance of singleton bean '" + beanName + "'" ); } } // 从缓存中返回 bean = getObjectForBeanInstance(sharedInstance, name, beanName, null ); } 复制代码`

这个部分相当于一个预检操作，如果缓存中已经有单例了，就直接返回，避免重复创建

接下来就是真正的创建过程，如下

` else { // 发现bean正在被创建，说明缓存中已经有原型bean， // 可能是由于循环引用导致，这里抛出异常 if (isPrototypeCurrentlyInCreation(beanName)) { throw new BeanCurrentlyInCreationException(beanName); } // 查找容器中是否有指定bean的定义 BeanFactory parentBeanFactory = getParentBeanFactory(); if (parentBeanFactory != null && !containsBeanDefinition(beanName)) { // ... } // 判断是否需要类型验证，这个值默认为false if (!typeCheckOnly) { // 在容器中标记指定的bean已经被创建 markBeanAsCreated(beanName); } try { // 获取父级bean定义，合并公共属性 final RootBeanDefinition mbd = getMergedLocalBeanDefinition(beanName); checkMergedBeanDefinition(mbd, beanName, args); // 获取bean的依赖，保证其依赖的bean提前被正常的初始化 String[] dependsOn = mbd.getDependsOn(); // 如果依赖有其他bean，就先初始化其依赖的bean if (dependsOn != null ) { // ... } // 以单例模式创建 if (mbd.isSingleton()) { // ... } // 以原型模式创建 else if (mbd.isPrototype()) { // ... } // 如果是其他模式，就用bean定义资源中配置的生命周期范围（request、session、application等） else { // ... } } catch (BeansException ex) { cleanupAfterBeanCreationFailure(beanName); throw ex; } } 复制代码`

配合注释，核心部分分为以下几个小部分：

* 预检查
* 前置工作
* 按指定模式进行实例化

这样一看好像很简单，确实如此，如果你仅仅是想了解一个大致的创建过程，接下来的部分可以略过，直接进入下一部分，如果你想详细地了解整个创建过程，那么就请跟着我再进一步分析在代码中省略的部分

##### 查找容器中bean的定义 #####

` // 查找容器中是否有指定bean的定义 BeanFactory parentBeanFactory = getParentBeanFactory(); // 如果当前容器中不存在bean的定义，且父容器不为空，就进入父容器中查找 if (parentBeanFactory != null && !containsBeanDefinition(beanName)) { String nameToLookup = originalBeanName(name); // 如果父容器是AbstractBeanFactory，说明已经到最顶层容器了， // 直接调用其doGetBean方法 if (parentBeanFactory instanceof AbstractBeanFactory) { return ((AbstractBeanFactory) parentBeanFactory).doGetBean( nameToLookup, requiredType, args, typeCheckOnly); } // 如果指定参数，就根据显式参数查找 else if (args != null ) { return (T) parentBeanFactory.getBean(nameToLookup, args); } // 如果没有指定参数，就根据指定类型名查找 else if (requiredType != null ) { return parentBeanFactory.getBean(nameToLookup, requiredType); } // 否则就采用默认查找方式 else { return (T) parentBeanFactory.getBean(nameToLookup); } } 复制代码`

查找bean定义的过程实际上是一个递归的操作，如果子类中不存在bean定义，就从父类中寻找，如果父类不存在，就去父类的父类中寻找，...，直到抵达最顶层的父类

##### 获取bean的依赖 #####

` // 获取bean的依赖，保证其依赖的bean提前被正常的初始化 String[] dependsOn = mbd.getDependsOn(); if (dependsOn != null ) { // 遍历并初始化所有依赖 for (String dep : dependsOn) { // 如果存在循环依赖，就抛出异常 if (isDependent(beanName, dep)) { throw new BeanCreationException(mbd.getResourceDescription(), beanName, "Circular depends-on relationship between '" + beanName + "' and '" + dep + "'" ); } // 注册依赖 registerDependentBean(dep, beanName); try { // 调用getBean方法创建依赖的bean getBean(dep); } catch (NoSuchBeanDefinitionException ex) { throw new BeanCreationException(mbd.getResourceDescription(), beanName, "'" + beanName + "' depends on missing bean '" + dep + "'" , ex); } } } 复制代码`

这个方法没什么好说的，配合注释应该能很轻松地看懂，重点是其中包含一个检查循环依赖的过程

##### 以单例模式创建bean #####

接下来就是整个方法的核心了，这里的三种创建模式大同小异，这里只讲最经典的单例模式，其余创建模式可以自行查阅

` // 以单例模式创建 if (mbd.isSingleton()) { // 创建单例对象 sharedInstance = getSingleton(beanName, () -> { try { return createBean(beanName, mbd, args); } catch (BeansException ex) { // 从缓存中删除单例，同时删除接收到该bean临时引用的bean destroySingleton(beanName); throw ex; } }); // 获取给定的bean的实例 bean = getObjectForBeanInstance(sharedInstance, name, beanName, mbd); } 复制代码`

整个方法看起来非常清晰，很好理解，无非就是创建单例，然后返回（不理解‘()->{}’这样的lambda表达式的，可以参考我的上一篇博文： [函数式编程——Java中的lambda表达式]( https://link.juejin.im?target=https%3A%2F%2Fblog.csdn.net%2Fqq_37435078%2Farticle%2Fdetails%2F89048253 ) ）

这段代码虽然短，但是包含了整个方法中核心的内容：创建bean实例，我们点进createBean方法，这是一个抽象方法，实现部分在AbstractAutowireCapableBeanFactory中，具体源码如下：

` @Override protected Object createBean (String beanName, RootBeanDefinition mbd, @Nullable Object[] args) throws BeanCreationException { if (logger.isTraceEnabled()) { logger.trace( "Creating instance of bean '" + beanName + "'" ); } RootBeanDefinition mbdToUse = mbd; Class<?> resolvedClass = resolveBeanClass(mbd, beanName); if (resolvedClass != null && !mbd.hasBeanClass() && mbd.getBeanClassName() != null ) { mbdToUse = new RootBeanDefinition(mbd); mbdToUse.setBeanClass(resolvedClass); } try { mbdToUse.prepareMethodOverrides(); } catch (BeanDefinitionValidationException ex) { throw new BeanDefinitionStoreException(mbdToUse.getResourceDescription(), beanName, "Validation of method overrides failed" , ex); } try { Object bean = resolveBeforeInstantiation(beanName, mbdToUse); if (bean != null ) { return bean; } } catch (Throwable ex) { throw new BeanCreationException(mbdToUse.getResourceDescription(), beanName, "BeanPostProcessor before instantiation of bean failed" , ex); } try { Object beanInstance = doCreateBean(beanName, mbdToUse, args); if (logger.isTraceEnabled()) { logger.trace( "Finished creating instance of bean '" + beanName + "'" ); } return beanInstance; } catch (BeanCreationException | ImplicitlyAppearedSingletonException ex) { throw ex; } catch (Throwable ex) { throw new BeanCreationException( mbdToUse.getResourceDescription(), beanName, "Unexpected exception during bean creation" , ex); } } 复制代码`

看起来非常臃肿，为了便于分析，我们把日志和异常都删掉，再来看：

` @Override protected Object createBean (String beanName, RootBeanDefinition mbd, @Nullable Object[] args) throws BeanCreationException { RootBeanDefinition mbdToUse = mbd; // 判断给定的bean是否可以被实例化（即是否可以被当前的类加载器加载） Class<?> resolvedClass = resolveBeanClass(mbd, beanName); // 如果不可以，就委派给其父类进行查找 if (resolvedClass != null && !mbd.hasBeanClass() && mbd.getBeanClassName() != null ) { mbdToUse = new RootBeanDefinition(mbd); mbdToUse.setBeanClass(resolvedClass); } // 准备覆盖bean中的方法 mbdToUse.prepareMethodOverrides(); // 如果设置了初始化前后的处理器，就返回一个代理对象 Object bean = resolveBeforeInstantiation(beanName, mbdToUse); if (bean != null ) { return bean; } // 创建bean Object beanInstance = doCreateBean(beanName, mbdToUse, args); return beanInstance; } 复制代码`

是不是一下子就简单了很多，这也是分析源码常用的方式，可以更方便地理解程序结构。好了不多说，我们看这段程序，首先是传入了三个参数：bean名称、父类bean，以及参数列表，然后就是一些常规操作，我们这里只看核心方法，发现实际这里并没有创建bean的代码，毕竟连new都没有，别急，点进doCreateBean方法，接着看：

` protected Object doCreateBean ( final String beanName, final RootBeanDefinition mbd, final @Nullable Object[] args) throws BeanCreationException { BeanWrapper instanceWrapper = null ; if (mbd.isSingleton()) { // 移除缓存（单例模式的一种实现方式）中beanName的映射，并返回这个bean instanceWrapper = this.factoryBeanInstanceCache.remove(beanName); } if (instanceWrapper == null ) { // 如果缓存中没有该bean，就创建该bean的实例 instanceWrapper = createBeanInstance(beanName, mbd, args); } // 对bean进行封装 final Object bean = instanceWrapper.getWrappedInstance(); // 获取bean类型 Class<?> beanType = instanceWrapper.getWrappedClass(); if (beanType != NullBean.class) { mbd.resolvedTargetType = beanType; } // 对后置处理器加同步锁 // 允许后置处理器修改合并后bean定义 synchronized (mbd.postProcessingLock) { // 判断后置处理器是否处理完成，如果没有就进行【合并bean定义】后的处理操作 if (!mbd.postProcessed) { try { applyMergedBeanDefinitionPostProcessors(mbd, beanType, beanName); } catch (Throwable ex) { throw new BeanCreationException(mbd.getResourceDescription(), beanName, "Post-processing of merged bean definition failed" , ex); } mbd.postProcessed = true ; } } // 立即将单例缓存起来，以便于依赖对象的循环引用 boolean earlySingletonExposure = (mbd.isSingleton() && this.allowCircularReferences && isSingletonCurrentlyInCreation(beanName)); if (earlySingletonExposure) { if (logger.isTraceEnabled()) { logger.trace( "Eagerly caching bean '" + beanName + "' to allow for resolving potential circular references" ); } // 让容器尽早持有对象的引用，以便于依赖对象的循环引用 addSingletonFactory(beanName, () -> getEarlyBeanReference(beanName, mbd, bean)); } // 初始化bean实例，实际触发依赖的地方 Object exposedObject = bean; try { // 用参数填充bean实例 populateBean(beanName, mbd, instanceWrapper); // 初始化bean对象 exposedObject = initializeBean(beanName, exposedObject, mbd); } catch (Throwable ex) { if (ex instanceof BeanCreationException && beanName.equals(((BeanCreationException) ex).getBeanName())) { throw (BeanCreationException) ex; } else { throw new BeanCreationException( mbd.getResourceDescription(), beanName, "Initialization of bean failed" , ex); } } // 父类是单例对象 && 该bean允许循环引用 && 该bean正在创建 if (earlySingletonExposure) { // 获取已注册的单例bean Object earlySingletonReference = getSingleton(beanName, false ); if (earlySingletonReference != null ) { // 如果已注册的bean和正在创建的bean是同一个，则直接返回这个bean if (exposedObject == bean) { exposedObject = earlySingletonReference; } // 如果该bean依赖于其他bean，且不允许在循环依赖的情况下注入bean else if (! this.allowRawInjectionDespiteWrapping && hasDependentBean(beanName)) { String[] dependentBeans = getDependentBeans(beanName); Set<String> actualDependentBeans = new LinkedHashSet<>(dependentBeans.length); for (String dependentBean : dependentBeans) { // 检查类型并添加依赖 if (!removeSingletonIfCreatedForTypeCheckOnly(dependentBean)) { actualDependentBeans.add(dependentBean); } } if (!actualDependentBeans.isEmpty()) { throw new BeanCurrentlyInCreationException(beanName, "Bean with name '" + beanName + "' has been injected into other beans [" + StringUtils.collectionToCommaDelimitedString(actualDependentBeans) + "] in its raw version as part of a circular reference, but has eventually been " + "wrapped. This means that said other beans do not use the final version of the " + "bean. This is often the result of over-eager type matching - consider using " + "'getBeanNamesOfType' with the 'allowEagerInit' flag turned off, for example." ); } } } } // 添加bean到工厂中的一次性bean列表中，仅适用于单例模式 try { registerDisposableBeanIfNecessary(beanName, bean, mbd); } catch (BeanDefinitionValidationException ex) { throw new BeanCreationException( mbd.getResourceDescription(), beanName, "Invalid destruction signature" , ex); } return exposedObject; } 复制代码`

具体执行流程已经详细地注释在代码中，我也不准备再复述一遍，相信认真看的都能读懂，我们这里关注一个很有意思的点，可以发现代码中有两个地方涉及了bean的加载：

` // ... instanceWrapper = createBeanInstance(beanName, mbd, args); // ... exposedObject = initializeBean(beanName, exposedObject, mbd); // ... 复制代码`

先来看createBeanInstance方法，这里为了避免很多人看不下去，就不打算放源码了，我简单地讲一下方法的流程：

* 预检查
* 调用instantiateUsingFactoryMethod工厂方法对bean进行实例化
* 使用自动装配方法进行实例化 -- 设置同步标记 -- 如果设置了自动装配属性，就调用autowireConstructor方法根据参数类型自动匹配构造方法 -- 否则使用默认的无参构造方法
* 如果没有设置自动装配，就使用构造方法进行实例化

肯定有人会问，这个方法不是已经实例化对象了吗，那后面的方法是干什么的？别急，我们直接进入initializeBean方法中，对源码有兴趣的可以自行查阅，我这里也是简要说下流程：

* 获取安全管理接口
* 根据设定的实例化策略来创建对象

我们可以发现，在这个方法里触发了applyBeanPostProcessors **Before** Initialization和applyBeanPostProcessors **After** Initialization两个方法，从字面意思上也很好理解，就是前置处理器和后置处理器，在这两个方法之间，调用了invokeInitMethods方法来执行初始化方法

那么问题来了，initializeBean和createBeanInstance有什么区别呢，不都是初始化吗？实际上，真正实例化Java Bean的是createBeanInstance方法，而initializeBean则相当于我们的自定义初始化操作，同时在其中也会执行一些前置处理和后置处理

createBeanInstance方法执行完，就相当于我们简单new出来一个对象而已，但是这个对象没有添加事务，没有添加aop，没有进行url的映射等等操作，所以就需要initializeBean来进行初始化

别忘了，之前我们说的那个特别长的doGetBean方法还没完呢，最后还有一段，如下：

` // 检查所需类型是否与bean类型匹配（如果没有设置检查类型匹配，最后会进行强制类型转换） if (requiredType != null && !requiredType.isInstance(bean)) { try { T convertedBean = getTypeConverter().convertIfNecessary(bean, requiredType); if (convertedBean == null ) { throw new BeanNotOfRequiredTypeException(name, requiredType, bean.getClass()); } return convertedBean; } catch (TypeMismatchException ex) { if (logger.isTraceEnabled()) { logger.trace( "Failed to convert bean '" + name + "' to required type '" + ClassUtils.getQualifiedName(requiredType) + "'" , ex); } throw new BeanNotOfRequiredTypeException(name, requiredType, bean.getClass()); } } return (T) bean; 复制代码`

##### 阶段总结 #####

依赖注入阶段有点长，我们这里做一下总结，总共步骤如下：

* 执行getBean，内部调用了doGetBean方法
* 检查缓存，如果缓存中有单例，就直接返回
* 查找bean的定义，同时进行类型验证
* 合并父级公有属性，创建依赖对象
* 根据设置的创建模式，选择创建方法
* 执行方法的覆盖，准备创建bean的实例
* 检查注册表，如果注册表中存在，就直接返回
* 否则，创建bean的实例
* 对单例进行缓存，以便循环依赖使用
* 注入依赖属性
* 执行初始化操作，同时会执行前后处理器的方法
* 添加bean到工厂中的一次性列表中（仅适用于单例模式）
* 进行类型检查（如果没有设置类型检查，最后会进行强制类型转换）

### 总结 ###

我习惯把总结写成要点的形式，因为这种方式比较清晰，所以尽量习惯一下

* 类似于ClassPathXmlApplicationContext这种具体的容器，在初始化时会首先将最顶层容器AbstractApplicationContext设置为资源加载器，然后会设置资源路径
* 容器在初始化时，如果设置了refresh参数，则会在每一次初始化时会重新注册处理器
* 容器加载资源是通过obtainFreshBeanFactory方法进行加载的，这个方式实际调用了AbstractRefreshableApplicationContext的loadBeanDefinitions方法进行加载
* getBean方法的实现在AbstractBeanFactory中，内部调用了doGetBean方法（doGetBean方法的执行流程就在上面）