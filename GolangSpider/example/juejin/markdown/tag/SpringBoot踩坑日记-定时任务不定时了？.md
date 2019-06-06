# SpringBoot踩坑日记-定时任务不定时了？ #

## 问题描述 ##

springboot定时任务用起来大家应该都会用，加两注解，加点配置就可以运行。但是如果仅仅处在应用层面的话，有很多内在的问题开发中可能难以察觉。话不多说，我先用一种 **极度夸张** 的手法，描述一下遇到的一个问题。

` @Component public class ScheduleTest { @Scheduled(initialDelay = 1000,fixedRate = 2*1000) public void test_a (){ System.out.println( "123" ); } @Scheduled(initialDelay = 2*1000,fixedRate = 2*1000) public void test_b (){ while ( true ){ try { Thread.sleep(2*1000); System.out.println( "456" ); } catch (InterruptedException e) { e.printStackTrace(); } } } } 复制代码`

上面代码是一个项目中的两个定时任务，test_a是正常的方法，test_b是发生异常的方法，为了凸显异常，我搞了个死循环。

在这种情况下，使用默认的定时任务配置运行，会发生什么现象呢？试试看就知道了， **定时任务一直在方法b中循环着，方法a永远执行不到！！！**

## 问题原因 ##

查看源代码后发现 [SpringBoot源码解析-Scheduled定时器的原理]( https://juejin.im/post/5caef144f265da03bb6fa07f ) ，springboot中， **默认的定时任务线程池是只有一个线程的** ，所以如果在一堆定时任务中，有一个发生了延时或者死循环之类的异常，很大可能会影响到其他的定时任务。

## 解决方案 ##

既然问题出在线程池数量上，那么为了让各个任务之间不会互相干扰，那就配置相应的线程池就好了。

### 方案一 异步执行 ###

` @Scheduled(initialDelay = 1000,fixedRate = 2*1000) @Async public void test_a (){ System.out.println( "123" ); } @Scheduled(initialDelay = 2*1000,fixedRate = 2*1000) @Async public void test_b (){ while ( true ){ try { Thread.sleep(2*1000); System.out.println( "456" ); } catch (InterruptedException e) { e.printStackTrace(); } } } 复制代码`

既然在单线程中因为一个任务卡住而影响到其他任务，那么把这个任务异步执行，问题就解决啦。

### 方案二 自定义定时任务线程池数量 ###

` @Configuration public class GlobalConfiguration { @Bean public TaskScheduler schedule (){ return new ConcurrentTaskScheduler(new ScheduledThreadPoolExecutor(2)); } } 复制代码`

既然单线程会互相干扰，那么分配足够的线程，让他们各自分开运行，也是可以解决的。

### 两种方案对比 ###

两种方案都可以解决各个任务之间互相干扰的问题，但是需要根据实际情况选择合适的。我们就以上面出现死循环的代码来分析。

如果在定时任务中真的发生了死循环，那么使用异步执行则会带来灾难性的后果。因为在定时任务这个线程中，每次任务执行完毕后，他会计算下次时间，再次添加一个任务进入异步线程池。而 **添加进异步线程池的任务因为死循环而一直占用着线程资源** 。随着时间的增加异步线程池的所有线程资源都会被死循环的任务占据，导致其他服务全部阻塞。

而使用自定义定时任务线程池则会好一点，因为只有当任务执行完成后，才会计算时间，在执行下次任务。虽然因为死循环任务一直在执行，但是也顶多占据一个线程的资源，不至于更大范围的影响。

## **返回目录** ( https://juejin.im/post/5c8a4458f265da2da23d703c ) ##