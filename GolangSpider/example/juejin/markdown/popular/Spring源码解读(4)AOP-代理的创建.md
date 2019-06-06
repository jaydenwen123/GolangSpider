# Spring源码解读(4)AOP-代理的创建 #

## 1、概述 ##

Spring AOP的核心是基于代理实现的，代理有jdk基于接口的动态代理和基于asm实现的允许类没有接口的cglib代理，上一小结，已经分析过Spring封装了使用了 ` @Aspect` 注解的类，并将切面方法封装成实现了 ` AbstractAspectJAdvice` 的类放入缓存中，等待创建代理对象时使用， ` AbstractAspectJAdvice` 有三个参数:

` //切面方法 诸如使用了@Before这类的方法 Method aspectJAdviceMethod //封装了expression表达式，给定一个类和方法名，能返回是否匹配 AspectJExpressionPointcut pointcut //可以创建一个切面类，然后通过反射执行切面方法 AspectInstanceFactory aspectInstanceFactory 复制代码`

通过这个对象就可以拦截所有的类的创建找出符合条件的bean创建代理执行增强操作，这也是spring的实现原理。

## 2、jdk代理 ##

jdk代理是基于接口的代理，所以被代理的对象必须是有接口实现的类，代理创建时通过 ` Proxy.newProxyInstance` 实现的，这个方法有三个参数:

` //指定要使用的类加载器 ClassLoader loader, //被代理的类所实现的接口，增强接口的方法 Class<?>[] interfaces, //方法处理器，会拦截所有方法，然后执行增强参数。 InvocationHandler inoker 复制代码`

简单实例

订单操作接口

` public interface OrderUpdateService { /** * 订单付款 * @param orderAmt */ void payOrder (String orderAmt) ; } 复制代码`

订单操作实现类

` @Slf 4j public class OrderUpdateServiceImpl implements OrderUpdateService { public void payOrder (String orderAmt) { log.info( "订单付款中....." ); } } 复制代码`

订单代理类

` @Slf 4j public class OrderUpdateServiceImplProxy { //目标类 要增强的类 private OrderUpdateService updateService; public OrderUpdateServiceImplProxy (OrderUpdateService updateService) { this.updateService = updateService; } public OrderUpdateService getProxy () { return (OrderUpdateService) Proxy.newProxyInstance(OrderUpdateServiceImplProxy.class.getClassLoader(), new Class[]{OrderUpdateService.class}, new InvocationHandler() { public Object invoke (Object proxy, Method method, Object[] args) throws Throwable { if ( "payOrder".equals(method.getName())) { log.info( "付款前处理" ); method.invoke(updateService, args); log.info( "付款后处理" ); } return proxy; } }); } } 复制代码`

demo验证

` public static void main (String[] args) { OrderUpdateService orderUpdateService = new OrderUpdateServiceImpl(); OrderUpdateServiceImplProxy proxy = new OrderUpdateServiceImplProxy(orderUpdateService); orderUpdateService = proxy.getProxy(); orderUpdateService.payOrder( "100" ); } 复制代码`

验证结果

` 11 : 18 : 07.859 [main] INFO s.a.e.p.OrderUpdateServiceImplProxy - 付款前处理 11 : 18 : 07.890 [main] INFO s.a.e.p.s.OrderUpdateServiceImpl - 订单付款中..... 11 : 18 : 07.890 [main] INFO s.a.e.p.OrderUpdateServiceImplProxy - 付款后处理 复制代码`

payOrder方法的确时被增强了。

## 3、cglib代理 ##

Cglib是一个强大的、高性能的 **代码生成包** ，它广泛被许多AOP框架使用，为他们 **提供方法的拦截** 。如下图所示Cglib与Spring等应用的关系。

![](图片格式不对)

* 最底层的是字节码 ` Bytecode` ，字节码是Java为了保证“一次编译、到处运行”而产生的一种虚拟指令格式，例如iload_0、iconst_1、if_icmpne、dup等
* 位于字节码之上的是 ` ASM` ，这是一种直接操作字节码的框架，应用ASM需要对Java字节码、Class结构比较熟悉
* 位于 ` ASM` 之上的是 ` CGLIB` 、 ` Groovy` 、 ` BeanShell` ，后两种并不是Java体系中的内容而是脚本语言，它们通过ASM框架生成字节码变相执行Java代码，这说明 **在JVM中执行程序并不一定非要写Java代码----只要你能生成Java字节码，JVM并不关心字节码的来源** ，当然通过Java代码生成的JVM字节码是通过编译器直接生成的，算是最“正统”的JVM字节码
* 位于 ` CGLIB` 、 ` Groovy` 、 ` BeanShell` 之上的就是 ` Hibernate` 、 ` Spring AOP` 这些框架了，这一层大家都比较熟悉
* 最上层的是Applications，即具体应用，一般都是一个Web项目或者本地跑一个程序

所以，Cglib的实现是在字节码的基础上的，并且使用了开源的ASM读取字节码，对类实现增强功能的。

Cglib可以通过 ` Callback` 回调函数完成对方法的增强，通过 ` CallbackFilter` 函数过滤符合条件的函数，Spring也是基于这两个接口完成bean的增强功能的，所以可以猜测前面 ` Advice` 参数的 ` pointCut` 方法匹配器就是在 ` CallbackFilter` 中起作用的。下来看看cglib的简单使用：

下面的验证是拦截 ` GameService` 的 ` playGames` 方法，在 ` playGames` 方法之前，执行 ` TransactionManager` 类的 ` start` 和 ` commit` 方法。

` @Test public void testProxyFilter () throws Exception { Enhancer enhancer = new Enhancer(); enhancer.setSuperclass(GameService.class); // NoOp.INSTANCE 表示没有被匹配的方法不执行任何操作 如果CallbackFilter返回0 //代表使用callbacks集合中的第一个元素执行方法拦截策略 1 则使用第二个。 Callback[] callbacks = { new TransactionInterceptor(), NoOp.INSTANCE}; //设置回调函数集合 enhancer.setCallbacks(callbacks); //根据回调过滤器返回指定回调函数索引 enhancer.setCallbackFilter( new TransactionFilter()); GameService gameService = (GameService) enhancer.create(); Person person = new Person( "lijinpeng" , 26 , false , new UserDao()); gameService.setPerson(person); gameService.playGames(); } 复制代码`

执行结果： ` playGames` 得到了增强。

` [ 2019 - 06 - 03 15 : 13 : 59 ] [INFO ][s.r.beans.models.aop.TransactionManager][ 15 ][start][]start tx [ 2019 - 06 - 03 15 : 13 : 59 ] [INFO ][s.road.beans.models.scan.GameService][ 29 ][playGames][]person-name:lijinpeng play games [ 2019 - 06 - 03 15 : 13 : 59 ] [INFO ][s.r.beans.models.aop.TransactionManager][ 20 ][commit][]commit tx 复制代码`

方法增强器

` @Slf 4j public class TransactionManager { public void start () { log.info( "start tx" ); MessageTracerUtils.addMessage( "start tx" ); } public void commit () { log.info( "commit tx" ); MessageTracerUtils.addMessage( "commit tx" ); } public void rollback () { log.info( "rollback tx" ); MessageTracerUtils.addMessage( "rollback tx" ); } public Object getAspectInstance () { return new TransactionManager(); } } 复制代码`

方法拦截器

` //方法拦截器 会拦截 所有方法 ，所以需要加判断 cglib 还提供了filter过滤器 可以用于过滤指定方法 public class TransactionInterceptor implements MethodInterceptor { TransactionManager tx = new TransactionManager(); public Object intercept (Object o, Method method, Object[] objects, MethodProxy methodProxy) throws Throwable { tx.start(); Object value = methodProxy.invokeSuper(o, objects); tx.commit(); return value; } } 复制代码`

通过 ` CallbackFilter` 接口，可以自定义拦截增强策略，返回的是前面设置的 ` Callbacks` 集合中的回调函数索引，这里的含义是,如果方法名是 ` playGames` 则使用Callbacks集合中索引为0的回调函数即 ` TransactionInterceptor` ，该回调函数就是要在 ` playGames` 前执行 ` TransactionManager` 的 ` start()` 方法， ` playGames` 方法后执行 ` commit()` ,否则使用 ` NoOp.INSTANCE` 表示什么也不做。

` public class TransactionFilter implements CallbackFilter { //返回回调函数在集合中的索引 public int accept (Method method) { if ( "playGames".equals(method.getName())) { return 0 ; } else { return 1 ; } } } 复制代码`

cglib代理通过Callback方法控制目标方法的增强逻辑，通过CallbackFilter用来指定适配的方法使用不通的回调函数完成不通的功能增强，cglib还可以为类动态得创建新得方法,不过知道cglib如何实现代理对Spring AOP的源码学习来说已经足够了。

## 4、代理对象的校验 ##

在Spring中，并不是所有的bean都需要创建代理，在Spring Aop中，只有被我们配置的Aspect切面类的PointCut表达式匹配的类才会创建代理，所以在创建代理之前需要对要创建bean进行校验以判断该类是否需要被创建为代理对象。上一篇的结尾处分析过，当解析完@Aspect注解的切面类后，就会调用 ` BeanPostProcessor` 的 ` postProcessAfterInitialization` 的方法去创建一个代理对象，这个方法会调用 ` wrapIfNecessary` 方法，这个方法会完成bean是否需要被创建为代理的校验，在这个方法里面，有一个核心的方法 ` getAdvicesAndAdvisorsForBean` ，这个方法会返回一个 ` Advisor` 数组集合，这个数组集合的Advisor即表示要创建的bean是否能被Advisor的Advice中的PointCut匹配到，如果可以匹配到，则为bean创建代理。 ` AbstractAdvisorAutoProxyCreator` 完成这个方法的功能实现：

` @Override protected Object[] getAdvicesAndAdvisorsForBean(Class<?> beanClass, String beanName, TargetSource targetSource) { //查找所有需要作用于beanClass的Advisor List<Advisor> advisors = findEligibleAdvisors(beanClass, beanName); if (advisors.isEmpty()) { //直接返回null return DO_NOT_PROXY; } return advisors.toArray(); } 复制代码`

核心方法是 ` findEligibleAdvisors(beanClass, beanName)` ：

` protected List<Advisor> findEligibleAdvisors (Class<?> beanClass, String beanName) { //核心方法1：查找容器中所有已经生成了的Advisor，这是第二次执行 List<Advisor> candidateAdvisors = findCandidateAdvisors(); //核心方法2:从容器中所有的Advisro找出能够作用于beanClass的Advisor List<Advisor> eligibleAdvisors = findAdvisorsThatCanApply(candidateAdvisors, beanClass, beanName); //为Advisor提供钩子方法 方便扩展Advisor extendAdvisors(eligibleAdvisors); if (!eligibleAdvisors.isEmpty()) { //根据Order排序Advisor eligibleAdvisors = sortAdvisors(eligibleAdvisors); } return eligibleAdvisors; } 复制代码`

这个方法首先又再次调用了 ` findCandidateAdvisors` 方法，这个方法在解析切面类的时候也调用了一次，目的是为切面类的切面方法生成 ` Advisor` ，并根据注解为每个 ` Advisor` 再生成一个 ` Advice` 对象。这里第二次调用的目的是获取容器中所有生成过的 ` Advisor` ,然后调用findAdvisorsThatCanApply方法找出所有能够作用于目标bean的 ` Advisor` ， ` extendAdvisors` 方法是为了方便扩展 ` Advisor` ， ` sortAdvisors` 是基于 ` Order` 接口排序 ` Advisor` 的。接下来主要来看下前两个方法： ` AnnotationAwareAspectJAutoProxyCreator` 重写了 ` findCandidateAdvisors` 方法，在重写了的方法里面，先调用了父类的，然后调用了基于注解的获取 ` Advisor` 的方法。

` @Override protected List<Advisor> findCandidateAdvisors () { //查找构建所有基于XML配置的切面类的切面方法 List<Advisor> advisors = super.findCandidateAdvisors(); //查找构建所有基于注解的切面类的切面方法 advisors.addAll( this.aspectJAdvisorsBuilder.buildAspectJAdvisors()); return advisors; } 复制代码`

这次我们来看看基于xml的切面类是如何获取容器中的 ` Advisor` 的。

` public List<Advisor> findAdvisorBeans () { String[] advisorNames = null ; synchronized ( this ) { //这个阶段如果xml配置了切面类 cachedAdvisorBeanNames 应该已经包含了切面方法生成的切面类的 //beanName了 advisorNames = this.cachedAdvisorBeanNames; if (advisorNames == null ) { //Aspect 可能是基于FactroyBean创建的，这里不实例化Aspect，还是应该交给代理实现 advisorNames = BeanFactoryUtils.beanNamesForTypeIncludingAncestors( this.beanFactory, Advisor.class, true , false ); this.cachedAdvisorBeanNames = advisorNames; } } if (advisorNames.length == 0 ) { return new LinkedList<Advisor>(); } List<Advisor> advisors = new LinkedList<Advisor>(); for (String name : advisorNames) { if (isEligibleBean(name)) { if ( this.beanFactory.isCurrentlyInCreation(name)) { if (logger.isDebugEnabled()) { logger.debug( "Skipping currently created advisor '" + name + "'" ); } } else { try { //直接通过切面方法生成的BeanDefinition 获取bean实例 advisors.add( this.beanFactory.getBean(name, Advisor.class)); } catch (BeanCreationException ex) { //异常处理...... } return advisors; } 复制代码`

从上面的代码看获取xml配置的切面方法生成的Advisor很简单，其实不然，真正复杂的逻辑是读取xml生成 ` BeanDefinition` 和 ` getBean()` 的过程，这里就不展开分析了。 这个方法过程就是找出由XMl配置的切面类生成的所有切面方法的 ` Advisor` ，然后遍历 ` beanName` ，生成实例对象，然会返回生成后的 ` Advisor` 实例对象。

### 4.1 获取所有Advisor ###

接下来再来看看基于注解的 ` Advisor` 是如何获取的。

` public List<Advisor> buildAspectJAdvisors () { List<String> aspectNames = null ; synchronized ( this ) { aspectNames = this.aspectBeanNames; if (aspectNames == null ) { //此时的aspectNames已经不为空了，省略之前生成Advisor的部分....... return advisors; } } if (aspectNames.isEmpty()) { return Collections.emptyList(); } List<Advisor> advisors = new LinkedList<Advisor>(); for (String aspectName : aspectNames) { //获取所有基于注解生成Advisor List<Advisor> cachedAdvisors = this.advisorsCache.get(aspectName); if (cachedAdvisors != null ) { advisors.addAll(cachedAdvisors); } else { //如果当时生成的Advisor是由工厂生成的，这个时候从工厂获取 MetadataAwareAspectInstanceFactory factory = this.aspectFactoryCache.get(aspectName); advisors.addAll( this.advisorFactory.getAdvisors(factory)); } } return advisors; } 复制代码`

这个方法有两部分，第一部分省略了，内容是上一篇分析过的生成基于注解的 ` Advisor` 的过程，第二部分是获取已经生成过的 ` Advisor` 的过程，由于基于注解生成的Advisor已经在创建的同时完成了bean的初始化，所以这里直接从缓存中获取即可，如果 ` Advisor` 是由工厂类生成，则此时需要通过工厂类获取 ` Advisor` 。

### 4.2 获取beanClass的Advisor ###

上面分析过程已经从xml和注解的方式获取到了容器中的所有 ` Advisor` ，接下来会执行第二步操作，校验容器中的 ` Advisor` 是否能够作用与要创建的bean上，回头看核心方法2: ` findAdvisorsThatCanApply`

` public static List<Advisor> findAdvisorsThatCanApply (List<Advisor> candidateAdvisors, Class<?> clazz) { if (candidateAdvisors.isEmpty()) { return candidateAdvisors; } List<Advisor> eligibleAdvisors = new LinkedList<Advisor>(); for (Advisor candidate : candidateAdvisors) { //核心点1:通过Introduction实现的advice if (candidate instanceof IntroductionAdvisor && canApply(candidate, clazz)) { eligibleAdvisors.add(candidate); } } boolean hasIntroductions = !eligibleAdvisors.isEmpty(); for (Advisor candidate : candidateAdvisors) { if (candidate instanceof IntroductionAdvisor) { // already processed continue ; } //核心点2:匹配能够作用于clazz的Advice if (canApply(candidate, clazz, hasIntroductions)) { eligibleAdvisors.add(candidate); } } return eligibleAdvisors; } 复制代码`

这个方法的核心点1做了一次匹配，这个匹配是通过 ` IntroductionAdvisor` 实现的，首先简单了解下这个问题的背景：

如果需要对一些组件新增一些功能，比如在付款接口完成后，需要记录这笔交易耗时时间，但是此时根据业务不同，可能会有很多个付款实现类，这个时候的做法就是新增一个统计耗时的接口，这样做虽然可以解决问题，但是污染了业务代码，而且每个类中都要实现一遍，这就违反了面向对象的 **开放封闭** 原则，这个问题通过 **装饰者** 模式解决，Spring为了能够完成这个功能，使用 ` Introduction+advice` 实现的，Spring通过一个特殊的拦截器 ` IntroductionInterceptor` 实现。与 ` PointCut` 作用于接口层次上不同，这种方式作用于类的层次，但是很不建议在生产环境使用这种粗粒度的实现，只需要知道Spring中有这种实现的方式就可以了，这里也不作为重点分析。

我们来看核心点2， ` canApply(candidate, clazz, hasIntroductions)` 从方法参数可以看出参数candidate是一个Advisor,claszz是目标类对象，由于 ` Advisor` 实现类里面有 ` PointCut` ,所以就可以匹配clazz类的方法是否属于切入方法，需要被增强。

` public static boolean canApply (Advisor advisor, Class<?> targetClass, boolean hasIntroductions) { if (advisor instanceof IntroductionAdvisor) { //因为作用于类 直接匹配类就行了。 return ((IntroductionAdvisor) advisor).getClassFilter().matches(targetClass); } else if (advisor instanceof PointcutAdvisor) { //转换成成PointCut接口 PointcutAdvisor pca = (PointcutAdvisor) advisor; return canApply(pca.getPointcut(), targetClass, hasIntroductions); } else { //如果advisor没有一个pointcut，默认它对所有的bean都是生效的 return true ; } } 复制代码`

上面的方法首先校验了 ` advisor` 是否是 ` IntroductionAdvisor` ，上面分析过， ` IntroductionAdvisor` 是基于类型的，所以这里直接针对这种bean直接调用了 ` ClassFilter` 接口去匹配，然后如果 ` advisor` 是一个 ` PointcutAdvisor` ，则转换成 ` PointcutAdvisor` 再调用 ` canApply` 方法，最后如果 ` Advisor` 并没有配置成一个 ` PointcutAdvisor` ，就默认对所有的bean都是生效的，进入 ` canApply` 方法继续看

` public static boolean canApply (Pointcut pc, Class<?> targetClass, boolean hasIntroductions) { Assert.notNull(pc, "Pointcut must not be null" ); if (!pc.getClassFilter().matches(targetClass)) { return false ; } //前面说过 PointCut接口getMethodMatcher的返回一个方法匹配器 MethodMatcher methodMatcher = pc.getMethodMatcher(); IntroductionAwareMethodMatcher introductionAwareMethodMatcher = null ; if (methodMatcher instanceof IntroductionAwareMethodMatcher) { introductionAwareMethodMatcher = (IntroductionAwareMethodMatcher) methodMatcher; } //获取目标类得所有接口接口 Set<Class<?>> classes = new LinkedHashSet<Class<?>>(ClassUtils.getAllInterfacesForClassAsSet(targetClass)); classes.add(targetClass); for (Class<?> clazz : classes) { //获取目标类及其接口的所有方法 Method[] methods = clazz.getMethods(); for (Method method : methods) { if ((introductionAwareMethodMatcher != null && introductionAwareMethodMatcher.matches(method, targetClass, hasIntroductions)) || methodMatcher.matches(method, targetClass)) { return true ; } } } return false ; } 复制代码`

这个方法其实很明确了，就是通过 ` PointCut` 获取到一个方法匹配器 ` MethodMatcher` ，这个匹配器通过 ` match` 方法可以判断当前的方法是否能够被匹配到， ` PointCut` 的封装其实就是aspectj的，但我们关注点并不在这，就不进入里面分析了， 然后获取目标类的所有接口，循环所有的接口然后获取各个接口的所有方法，调用 ` MethodMatcher` 的 ` match` 方法用来判断当前方法是否能够被匹配，只要有一个方法被匹配到，则但会true，表明当前的 ` advisor` 能够作用于当前bean。

Spring在bean的创建过程中，就是通过上面的过程查找能够作用于bean的 ` Advisor` 的，如果能够找到，就为bean创建代理，否则就当普通的bean创建，经过上面的判断，我们案例中的bean肯定是有符合条件的 ` Advisor` ，所以接下来看看Spring是如何创建代理的。

## 5、代理模式的查找 ##

再来回忆下这个方法：

` protected Object wrapIfNecessary (Object bean, String beanName, Object cacheKey) { //如果bean是通过TargetSource接口获取 则直接返回 if (beanName != null && this.targetSourcedBeans.contains(beanName)) { return bean; } //如果bean是切面类 直接返回 if (Boolean.FALSE.equals( this.advisedBeans.get(cacheKey))) { return bean; } //如果bean是Aspect 而且允许跳过创建代理， 加入advise缓存 返回 if (isInfrastructureClass(bean.getClass()) || shouldSkip(bean.getClass(), beanName)) { this.advisedBeans.put(cacheKey, Boolean.FALSE); return bean; } //如果前面生成的advisor缓存中存在能够匹配到目标类方法的Advisor 则创建代理 Object[] specificInterceptors = getAdvicesAndAdvisorsForBean(bean.getClass(), beanName, null ); if (specificInterceptors != DO_NOT_PROXY) { this.advisedBeans.put(cacheKey, Boolean.TRUE); //创建代理 Object proxy = createProxy( bean.getClass(), beanName, specificInterceptors, new SingletonTargetSource(bean)); this.proxyTypes.put(cacheKey, proxy.getClass()); return proxy; } this.advisedBeans.put(cacheKey, Boolean.FALSE); return bean; } 复制代码`

经过前面的分析，通过 ` getAdvicesAndAdvisorsForBean` 方法可以获取到容器中所有能被作用于当前bean的 ` Advisor` ，此时的 ` specificInterceptors` 肯定是不为null，所以就会执行 ` createProxy` 创建代理的方法了。

` protected Object createProxy ( Class<?> beanClass, String beanName, Object[] specificInterceptors, TargetSource targetSource) { if ( this.beanFactory instanceof ConfigurableListableBeanFactory) { //如果是基于IntroductionAdvisor，直接指定代理的目标类即可 AutoProxyUtils.exposeTargetClass((ConfigurableListableBeanFactory) this.beanFactory, beanName, beanClass); } //创建代理工厂 ProxyFactory proxyFactory = new ProxyFactory(); proxyFactory.copyFrom( this ); //校验当前代理工厂是否可以直接代理目标类和目标接口 if (!proxyFactory.isProxyTargetClass()) { if (shouldProxyTargetClass(beanClass, beanName)) { proxyFactory.setProxyTargetClass( true ); } else { //将beanClass的接口设置到代理工厂上 evaluateProxyInterfaces(beanClass, proxyFactory); } } //构建能够作用于beanClass类的Advisor集合 Advisor[] advisors = buildAdvisors(beanName, specificInterceptors); for (Advisor advisor : advisors) { proxyFactory.addAdvisor(advisor); } proxyFactory.setTargetSource(targetSource); customizeProxyFactory(proxyFactory); proxyFactory.setFrozen( this.freezeProxy); if (advisorsPreFiltered()) { proxyFactory.setPreFiltered( true ); } //创建代理对象 return proxyFactory.getProxy(getProxyClassLoader()); } 复制代码`

由于 ` IntroductionAdvisor` 这样特殊的拦截器已经指定了增强的类对象，也就是说该类实现了 ` IntroductionInterceptor` 拦截器，定义过了类的方法需要如何增强，所以上面的方法第一步就是给他指定了增强类，然后创建 ` ProxyFactory` 代理工厂，先校验代理工厂是否能够直接代理目标类及其接口，如果不能则 ` evaluateProxyInterfaces` 方法会将 ` beanClass` 的接口设置到代理工厂上，接下来通过 ` buildAdvisors` 方法构建一个能够作用于 ` beanClass` 类的 ` Advisor` 集合放入到代理工厂中，以便后续创建代理时使用，最后调用 ` proxyFactory.getProxy` 方法一个代理对象，在这里 ` buildAdvisors` 方法没可看的，就是将 ` Advisor` 封装到一个集合里等待备用：

` protected Advisor[] buildAdvisors(String beanName, Object[] specificInterceptors) { Advisor[] commonInterceptors = resolveInterceptorNames(); List<Object> allInterceptors = new ArrayList<Object>(); if (specificInterceptors != null ) { //将所有Advisor放入到allInterceptors集合中 allInterceptors.addAll(Arrays.asList(specificInterceptors)); if (commonInterceptors.length > 0 ) { if ( this.applyCommonInterceptorsFirst) { allInterceptors.addAll( 0 , Arrays.asList(commonInterceptors)); } else { allInterceptors.addAll(Arrays.asList(commonInterceptors)); } } } if (logger.isDebugEnabled()) { int nrOfCommonInterceptors = commonInterceptors.length; int nrOfSpecificInterceptors = (specificInterceptors != null ? specificInterceptors.length : 0 ); Advisor[] advisors = new Advisor[allInterceptors.size()]; for ( int i = 0 ; i < allInterceptors.size(); i++) { //根据Advice类型转换成不用的Advisor实现 advisors[i] = this.advisorAdapterRegistry.wrap(allInterceptors.get(i)); } return advisors; } 复制代码`

下面看看Spring时如何创建代理的

` public Object getProxy (ClassLoader classLoader) { //createAopProxy方法创建一个代理对象 return createAopProxy().getProxy(classLoader); } 复制代码`

` createAopProxy` 方法创建一个代理对象，这个方法由 ` DefaultAopProxyFactory` 类实现:

` @Override public AopProxy createAopProxy (AdvisedSupport config) throws AopConfigException { if (config.isOptimize() || config.isProxyTargetClass() || hasNoUserSuppliedProxyInterfaces(config)) { Class<?> targetClass = config.getTargetClass(); if (targetClass == null ) { throw new AopConfigException( "TargetSource cannot determine target class: " + "Either an interface or a target is required for proxy creation." ); } if (targetClass.isInterface() || Proxy.isProxyClass(targetClass)) { return new JdkDynamicAopProxy(config); } return new ObjenesisCglibAopProxy(config); } else { return new JdkDynamicAopProxy(config); } } 复制代码`

这个方法决定了是使用jdk代理还是cglib代理，上面的判断由三个校验组成，只要有任何一个结果为true都会进入else使用jdk代理，来看下这个三个判断条件:

` //自定义指定配置 config.isOptimize() //是否使用了proxy-target-class="true" config.isProxyTargetClass() //目标类是否实现了接口 而且接口不能与SpringProxy类相同 hasNoUserSuppliedProxyInterfaces(config)) 复制代码`

这三个方法中的 ` config.isOptimize()` ， ` config.isProxyTargetClass()` 默认都会返回false，针对最后一个方法是判断目标类是否没有接口，当目标类有接口实现的时候就会走到最后一个else，使用 ` JdkDynamicAopProxy` 创建代理，也就是jdk代理。

` private boolean hasNoUserSuppliedProxyInterfaces (AdvisedSupport config) { Class<?>[] ifcs = config.getProxiedInterfaces(); //如果目标类是实现了接口 会直接返回false return (ifcs.length == 0 || (ifcs.length == 1 && SpringProxy.class.isAssignableFrom(ifcs[ 0 ]))); } 复制代码`

然后再来看 ` config.isOptimize()` ，这个属性的初始值是false，我并没有在源码中发现有哪里可以修改这个值，但是如果我们自定义一个实现了 ` AbstractAdvisorAutoProxyCreator` 类的处理器，像下面这样配置，就可以设置这个值为true了。这个应该是基于Sping的扩展的，但这里对普通注解bean都是false的，所以不需要太纠结。

` <bean class="org.springframework.aop.framework.autoproxy.DefaultAdvisorAutoProxyCreator"> <property name="optimize" value="true"/> </bean> 复制代码`

然后再看一个比较关键的判断 ` config.isProxyTargetClass()` ，这个配置实际上就是aop基于自动注解配置里面的一个属性

` <aop:aspectj-autoproxy proxy-target- class = "true" > 复制代码`

这个值就是 ` proxy-target-class` 的值，这个值的注入是注册 ` AnnotationAwareAspectJAutoProxyCreator` 时候设置的，来看下面的代码，这部分是注册 ` AnnotationAwareAspectJAutoProxyCreator` 时候，设置 ` proxyTargetClass` 属性的

` private static void useClassProxyingIfNecessary (BeanDefinitionRegistry registry, Element sourceElement) { if (sourceElement != null ) { // PROXY_TARGET_CLASS_ATTRIBUTE就是proxy-target-class属性 boolean proxyTargetClass = Boolean.valueOf(sourceElement.getAttribute(PROXY_TARGET_CLASS_ATTRIBUTE)); if (proxyTargetClass) { //主要执行这个方法 AopConfigUtils.forceAutoProxyCreatorToUseClassProxying(registry); } boolean exposeProxy = Boolean.valueOf(sourceElement.getAttribute(EXPOSE_PROXY_ATTRIBUTE)); if (exposeProxy) { AopConfigUtils.forceAutoProxyCreatorToExposeProxy(registry); } } } 复制代码` ` public static void forceAutoProxyCreatorToUseClassProxying (BeanDefinitionRegistry registry) { if (registry.containsBeanDefinition(AUTO_PROXY_CREATOR_BEAN_NAME)) { BeanDefinition definition = registry.getBeanDefinition(AUTO_PROXY_CREATOR_BEAN_NAME); //如果proxy-target-class属性配置的是true 直接设置这个值的类型为true definition.getPropertyValues().add( "proxyTargetClass" , Boolean.TRUE); } } 复制代码`

所以说如果我们通过 ` <aop:aspectj-autoproxy proxy-target-class="true">` 方式如果配置了 ` proxy-target-class` 为true，则就会直接使用cglib代理了，当然下面还会有一些校验，不过最终还是会走到cglib代理，这次先来看看cglib代理吧

## 6、Cglib模式的代理 ##

### 6.1 Enhancer的创建 ###

cglib代理就是由 ` ObjenesisCglibAopProxy` 这个类实现的，获取代理对象直接调用其父类 ` CglibAopProxy的getProxy` 方法获取。

` @Override public Object getProxy (ClassLoader classLoader) { try { Class<?> rootClass = this.advised.getTargetClass(); Class<?> proxySuperClass = rootClass; if (ClassUtils.isCglibProxyClass(rootClass)) { proxySuperClass = rootClass.getSuperclass(); Class<?>[] additionalInterfaces = rootClass.getInterfaces(); for (Class<?> additionalInterface : additionalInterfaces) { this.advised.addInterface(additionalInterface); } } //校验类 validateClassIfNecessary(proxySuperClass, classLoader); //开始创建cglib代理 Enhancer enhancer = createEnhancer(); if (classLoader != null ) { enhancer.setClassLoader(classLoader); if (classLoader instanceof SmartClassLoader && ((SmartClassLoader) classLoader).isClassReloadable(proxySuperClass)) { enhancer.setUseCache( false ); } } enhancer.setSuperclass(proxySuperClass); enhancer.setInterfaces(AopProxyUtils.completeProxiedInterfaces( this.advised)); enhancer.setNamingPolicy(SpringNamingPolicy.INSTANCE); enhancer.setStrategy( new ClassLoaderAwareUndeclaredThrowableStrategy(classLoader)); //核心点1：获取回调集合 Callback[] callbacks = getCallbacks(rootClass); Class<?>[] types = new Class<?>[callbacks.length]; for ( int x = 0 ; x < types.length; x++) { types[x] = callbacks[x].getClass(); } //核心点2：设置过滤器 enhancer.setCallbackFilter( new ProxyCallbackFilter( this.advised.getConfigurationOnlyCopy(), this.fixedInterceptorMap, this.fixedInterceptorOffset)); enhancer.setCallbackTypes(types); // 创建代理对象 return createProxyClassAndInstance(enhancer, callbacks); } catch (Exception ex) { //异常捕获抛出......... } } 复制代码`

cglib代理的使用方式和我们之前做的简单案例的使用是一样的：

> 
> 
> 
> 首先创建一个Enhancer，setSuperclass设置要代理的对象，
> 
> 

> 
> 
> 
> setInterfaces可以不设置，因为是基于cglib代理的，不需要接口实现,这里的含义是指定创建的代理继承AopProxy的一些接口。
> 
> 

> 
> 
> 
> setNamingPolicy设置cglib的命名策略，这里是以BySpringCGLIB为前缀：
> 
> 

> 
> 
> 
> *** ` getCallbacks` ***方法用于获取增强器，会详细解析。
> 
> 

> 
> 
> 
> *** ` enhancer.setCallbackFilter()` ***用于过滤器用于过滤不需要增强的方法的，也会详细解析
> 
> 

> 
> 
> 
> 最后通过*** ` createProxyClassAndInstance` ***创建一个代理实例对象。
> 
> 

从上面单独对cglib单例的使用讲解可以知道，cglib代理是基于asm在字节码上的实现，可以动态的创建一个继承于目标类的代理对象，然后执行增强操作，此外cglib还提供了 ` Callback` 接口用于执行对方法的拦截，由 ` MethodInterceptor` 接口的 ` intercept` 方法拦截，然后cglib还可以通过 ` CallbackFilter` 类来指定对不同的类或者方法执行不同的增强操作， ` accept` 方法会返回一个索引，这个索引是设置的 ` CallbackFilter` 集合中集合的索引，通过 ` accpet` 方法如果返回0就会使用 ` Callback` 集合中第一个增强拦截器，下面会针对Spring设置的 ` Callback` 集合和 ` CallbackFilter` 分别分析。

### 6.2 Callback拦截器集合 ###

首先来看下Spring是如何获取 ` CallbackFilter` 集合。

` private Callback[] getCallbacks(Class<?> rootClass) throws Exception { // Parameters used for optimisation choices... boolean exposeProxy = this.advised.isExposeProxy(); boolean isFrozen = this.advised.isFrozen(); boolean isStatic = this.advised.getTargetSource().isStatic(); Callback aopInterceptor = new DynamicAdvisedInterceptor( this.advised); Callback targetInterceptor; if (exposeProxy) { //这个方法是表示处理嵌套的，也就是在代理方法中调用本次的其他方法 那么针对该方法的增强效果将会失效........ } else { targetInterceptor = isStatic ? new StaticUnadvisedInterceptor( this.advised.getTargetSource().getTarget()) : new DynamicUnadvisedInterceptor( this.advised.getTargetSource()); } Callback targetDispatcher = isStatic ? new StaticDispatcher( this.advised.getTargetSource().getTarget()) : new SerializableNoOp(); //这个事核心点，从这里可以看出Spring里面总共设置了6个方法增强器 //其中第一个就是我们我们配置的增强器 Callback[] mainCallbacks = new Callback[] { aopInterceptor, // for normal advice targetInterceptor, // invoke target without considering advice, if optimized new SerializableNoOp(), // no override for methods mapped to this targetDispatcher, this.advisedDispatcher, new EqualsInterceptor( this.advised), new HashCodeInterceptor( this.advised) }; Callback[] callbacks; if (isStatic && isFrozen) { ......... } else { callbacks = mainCallbacks; } return callbacks; } 复制代码`

首先Spring将创建的的 ` ProxyFactroy` 对象封装到 ` DynamicAdvisedInterceptor` ，这个代理类里面有之前匹配到的 ` Advisor` 集合，然后判断 ` exposeProxy` 是否是true，这个属性也是在xml配置文件中配置的，背景就是如果在事务A中使用了代理，事务A调用了目标类的的方法a，在方法a中又调用目标类的方法b，方法a，b同时都是要被增强的方法，如果不配置 ` exposeProxy` 属性，方法b的增强将会失效，如果配置 ` exposeProxy` ，方法b在方法a的执行中也会被增强了。接下来再看 ` mainCallbacks` 这个集合，这个集合里面封装了6个是实现了 ` Callback` 或 ` MethodInterceptor` 接口的类，其中第一个 ` aopInterceptor` 就是封装了目标类的要增强的逻辑也就是Advisor集合，第二个是针对之前配置的 ` optimize` 属性使用的，后面的四个基本上不会做太多业务逻辑上的拦截，Spring最终将这6个增强器集合返回作为cglib的拦截器链，之后通过 ` CallbackFilter` 的 ` accpet` 方法返回的索引从这个集合中返回对应的拦截增强器执行增强操作。

### 6.3 CallbackFilter 过滤器 ###

Spring实现的是通过 ` ProxyCallbackFilter` 实现 ` CallbackFilter` 接口，然后赋值给代理对象，上面分析过，这个接口主要是通过 ` accpet` 接口实现的，来看看是如何匹配的:

` @Override public int accept (Method method) { if (AopUtils.isFinalizeMethod(method)) { logger.debug( "Found finalize() method - using NO_OVERRIDE" ); //不使用代理 return NO_OVERRIDE; } if (! this.advised.isOpaque() && method.getDeclaringClass().isInterface() && method.getDeclaringClass().isAssignableFrom(Advised.class)) { if (logger.isDebugEnabled()) { logger.debug( "Method is declared on Advised interface: " + method); } return DISPATCH_ADVISED; } // We must always proxy equals, to direct calls to this. if (AopUtils.isEqualsMethod(method)) { logger.debug( "Found 'equals' method: " + method); return INVOKE_EQUALS; } // We must always calculate hashCode based on the proxy. if (AopUtils.isHashCodeMethod(method)) { logger.debug( "Found 'hashCode' method: " + method); return INVOKE_HASHCODE; } Class<?> targetClass = this.advised.getTargetClass(); // Proxy is not yet available, but that shouldn't matter. List<?> chain = this.advised.getInterceptorsAndDynamicInterceptionAdvice(method, targetClass); boolean haveAdvice = !chain.isEmpty(); boolean exposeProxy = this.advised.isExposeProxy(); boolean isStatic = this.advised.getTargetSource().isStatic(); boolean isFrozen = this.advised.isFrozen(); if (haveAdvice || !isFrozen) { // If exposing the proxy, then AOP_PROXY must be used. if (exposeProxy) { if (logger.isDebugEnabled()) { logger.debug( "Must expose proxy on advised method: " + method); } //使用封装了Advisor的拦截增强器 return AOP_PROXY; } String key = method.toString(); // Check to see if we have fixed interceptor to serve this method. // Else use the AOP_PROXY. if (isStatic && isFrozen && this.fixedInterceptorMap.containsKey(key)) { // We know that we are optimising so we can use the FixedStaticChainInterceptors. int index = this.fixedInterceptorMap.get(key); return (index + this.fixedInterceptorOffset); } else { if (logger.isDebugEnabled()) { logger.debug( "Unable to apply any optimisations to advised method: " + method); } //使用封装了Advisor的拦截增强器 return AOP_PROXY; } } else { if (exposeProxy || !isStatic) { return INVOKE_TARGET; } Class<?> returnType = method.getReturnType(); if (targetClass == returnType) { return INVOKE_TARGET; } else if (returnType.isPrimitive() || !returnType.isAssignableFrom(targetClass)) { return DISPATCH_TARGET; } else { return INVOKE_TARGET; } } } 复制代码`

这个方法里面，我们只需要关心 ` AOP_PROXY` 这个值就可以了，这个值是 ` 0` ，也就是对应上面生成的 ` Callback` 集合中的第一个拦截增强器，也就是 ` aopInterceptor` ，这个拦截器里面封装了所有能够作用于目标类的 ` Advisor` 增强类，所以就会调用到切面类定义的切面方法来进行增强操作了。

### 6.4 Advisor调用链 ###

上面分析过，Spring使用的是 ` aopInterceptor` 作为拦截增强器，这个增强器被封装进了 ` DynamicAdvisedInterceptor` 类中，这个类实现了 ` MethodInterceptor` 方法，所以被拦截到的方法会进入 ` intercept` 方法中，由于这个类里面封装了前面我们匹配到的所有能够作用于这个类的 ` Advisor` 集合，但是可能针对这些 ` Advisor` 可能执行顺序有些疑问，虽然Spring允许通过继承 ` Order` 接口实现排序，但是比如 ` @Before` ， ` @After` , ` @Round` 等是需要在某个方法的不同时机执行的，所以这需要构造一个调用链，废话不多，还是先来看aopInterceptor的intercept方法是如何做的:

` @Override public Object intercept (Object proxy, Method method, Object[] args, MethodProxy methodProxy) throws Throwable { try { if () { //省略不重要代码........ } else { //创建一个方法执行器 主要逻辑在这里 retVal = new CglibMethodInvocation(proxy, target, method, args, targetClass, chain, methodProxy).proceed(); } retVal = processReturnType(proxy, target, method, retVal); return retVal; } finally { if (target != null ) { releaseTarget(target); } if (setProxyContext) { // Restore old proxy. AopContext.setCurrentProxy(oldProxy); } } } 复制代码`

在这个方法里面，我们只需要关心 ` CglibMethodInvocation` 的 ` proceed` 方法，调用链就是在这里构建的

` @Override public Object proceed () throws Throwable { //首先获取到所有的Advisor 然后去除一个下标-1 if ( this.currentInterceptorIndex == this.interceptorsAndDynamicMethodMatchers.size() - 1 ) { return invokeJoinpoint(); } //获取当前Advisor Object interceptorOrInterceptionAdvice = this.interceptorsAndDynamicMethodMatchers.get(++ this.currentInterceptorIndex); //如果是动态方法匹配器 if (interceptorOrInterceptionAdvice instanceof InterceptorAndDynamicMethodMatcher) { InterceptorAndDynamicMethodMatcher dm = (InterceptorAndDynamicMethodMatcher) interceptorOrInterceptionAdvice; if (dm.methodMatcher.matches( this.method, this.targetClass, this.arguments)) { return dm.interceptor.invoke( this ); } else { //匹配失败直接跳过，再次执行proceed方法 return proceed(); } } else { //调用拦截器增强集合 return ((MethodInterceptor) interceptorOrInterceptionAdvice).invoke( this ); } } 复制代码`

这个方法是构造调用链的实现原理，首先获取到Advisor集合，然后用 ` currentInterceptorIndex` 表示当前集合剩余未执行的 ` Advisor` ， ` currentInterceptorIndex=0` 的时候，集合中所有的 ` Advisor` 已经执行完毕了，这种方法还是很巧妙的，我们从整体上看下是如何形成调用链的，该方法最终调用最后一行代码，也就是 ` MethodInterceptor` 类的 ` invoke` 方法，把当前对象作为参数传进去，当前对象也就是 ` aopInterceptor` ，这里面封装了所有符合条件的 ` Advisor` （每次都要强调这个类里面的数据就是让我们始终能保持在主线上），该类实现了 ` ProxyMethodInvocation` 接口间接实现了 ` MethodInvocation` 接口，最后一行的代码调用 ` invoke` 方法，也就是 ` MethodInterceptor` 接口的方法，来看下这个方法的定义:

` public interface MethodInterceptor extends Interceptor { Object invoke (MethodInvocation invocation) throws Throwable ; } 复制代码`

` invoke` 的参数是 ` MethodInvocation` ，由于 ` aopInterceptor` 是 ` MethodInvocation` 的实现类，所以可以将该类作为参数往下传递，前面分析过了切面的注解方法注解对应的类为:

+----------------+-------------------------------+
|    @BEFORE     | METHODBEFOREADVICEINTERCEPTOR |
+----------------+-------------------------------+
| @After         | AspectJAfterReturningAdvice   |
| @Around        | AspectJAroundAdvice           |
| @AfterThrowing | AspectJAfterThrowingAdvice    |
+----------------+-------------------------------+

这些生成的 ` Advice` 切面类毫无疑问都实现了 ` MethodInterceptor` 方法，而 ` MethodInterceptor` 的 ` invoke` 方法需要接收一个类型为 ` MethodInvocation` 的对象，正好 ` aopInterceptor` 实现了MethodInvocation ` ，所以可以作为参数在多个Advisor间流转，每执行完一个` Advisor ` 就进入` aopInterceptor ` 的` process ` 方法再次校验，直至执行完所有的` Advisor`。

假设此时的 ` aopInterceptor` 里面包含了两个切面方法 ` MethodBeforeAdviceInterceptor` 和 ` AspectJAfterReturningAdvice` ，此时先取出 ` AspectJAfterReturningAdvice` ，然后执行 ` invoke` 方法，看下 ` AspectJAfterReturningAdvice` 的 ` invoke` 方法：

` @Override public Object invoke (MethodInvocation mi) throws Throwable { try { //此时的mi对象就是 aopInterceptor ，又回到了proceed方法 return mi.proceed(); } finally { //自身的方法最后执行 invokeAdviceMethod(getJoinPointMatch(), null , null ); } } 复制代码`

` AspectJAfterReturningAdvice` 的 ` invoke` 方法会首先调用参数 mi 的 ` proceed` 方法，此时mi就是 ` aopInterceptor` 对象(这里面封装了所有适用于目标类对象的切面方法类),所以会再次返回到拦截器集合里，取出 ` MethodBeforeAdviceInterceptor` 对象，调用它的 ` invoke` 方法，最后自调用 ` AspectJAfterReturningAdvice` 封装的切面方法也就是通过 ` @After` 注解修饰的方法。再来看看 ` AspectJAfterReturningAdvice` 的invoke方法

` @Override public Object invoke (MethodInvocation mi) throws Throwable { this.advice.before(mi.getMethod(), mi.getArguments(), mi.getThis() ); return mi.proceed(); } 复制代码`

这里会先调用通过 ` @Before` 注解修饰的方法，然后再次调用 ` aopInterceptor` 拦截器链，虽然这两个注解是最简单的，但是确时逻辑最清晰的，也是最容易理解的，其他的切面注解方法的调用过程也是这样，最终完成这个拦截器链的调用。

经过上面的步骤，就可以创建一个代理对象了， 然后Spring将代理对象放入缓存中，我们使用时得到的是一个代理对象，调用目标方法，就会执行到增强器方法，进而找到所有的 ` Advisor` ，构造一个调用链，这样就完成业务的增强了。我相信如果对照着源码看，多调试几遍代码，应该会大致了解Spring实现Cglib的代理的一个流程。关于Spring时如何实现cglib代理的暂时就分析到这里。

## 7、JDK模式的代理 ##

jdk模式的代理是通过JdkDynamicAopProxy这个类实现的，由于jdk的代理的实现比cglib代理稍微简单一点，我这里暂时就先不分析了，关于jdk代理模式的解析估计有很多资料，以后有时间我补上这部分内容

## 8、总结 ##

Spring源码很复杂，如果想把每个细节都搞明白，几乎是不可能的，而且无疑会浪费很多时间，我们只需要搞懂常用的功能的实现方式，遇到问题可以从源码的角度解决问题就行了，说的更高一点，如果我们能学到作者的一些架构思路以及对代码的功能抽象，以后可以对Spring进行扩展，那就再好不过了，废话不多了，总结下一些这部分内容究竟在说些什么吧。

这一部分内容是Spring AOP的第二部分内容，是基于上一步解析 ` @Aspect` 注解，将注入 ` @Before` 等注解方法生成 ` Advice` 切面方法类的过程，这一部分则是使用这些切面方法类，然后找出能够作用于目标类的 ` Advisor` ，然后创建代理对象，具体流程是:

1、从 ` Advisor` 缓存中获取第一步生成的所有 ` Advisor` (包括通过xml配置和注解配置的)

2、遍历 ` Advisor` 集合，从 ` Advisor` 里面获取 ` PointCut` ，获取方法匹配器 ` MethodMatcher` ,通过match方法匹配一个当前 ` Advisor` 是否能够作用于目标类，如果能，则放入集合中

3、如果第二部获取的集合为空，则直接返回，不需要创建代理，如果不为空，也就是 ` Advisor` 缓存中有能够作用于目标类的 ` Advisor` ，则根据配置选择不同的代理的方法(cglib 还是 jdk 代理)

4、假如是 ` cglib` 代理，需要构建 ` Callback` 拦截增强器集合，然后创建一个 ` CallbackFilter` 过滤器，选择合适的拦截增强器对bean实现增强处理

5、返回创建的代理对象。

写完这一篇，我终于把这两个月看Spring的源码对其的理解记录下来了， 即便如此，我感觉我对Spring还是知之甚少，还需要巩固，我认为如果能通过Spring封装自己写的框架，就像 ` Sping-redis` , ` Spring-kafka` 那样直接通过注解或者配置就可以拿来使用了，肯定对Spring有一个更高层次的认识，为了加深对Spring的理解，通过刘欣老师的讲解，我也跟着写了一遍稍微简单的Spring实现，那段代码花了我将近一个月的时间，但是写完再看Spring的源码就会容易理解的多，源码地址是:

[github.com/StringBuild…]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2FStringBuilderSun%2FHi-Spring.git )

感兴趣的可以拉下来看看。