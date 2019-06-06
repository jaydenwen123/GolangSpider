# SpringBoot实现动态控制定时任务-支持多参数 #

> 
> 
> 
> 由于工作上的原因，需要进行定时任务的动态增删改查，网上大部分资料都是整合quertz框架实现的。本人查阅了一些资料，发现springBoot本身就支持实现定时任务的动态控制。并进行改进，现支持任意多参数定时任务配置
> 
> 
> 

实现结果如下图所示：

![](https://user-gold-cdn.xitu.io/2019/5/31/16b0be286b6a031a?imageView2/0/w/1280/h/960/ignore-error/1) 后台测试显示如下：

![](https://user-gold-cdn.xitu.io/2019/5/31/16b0bf1a0e118b4b?imageView2/0/w/1280/h/960/ignore-error/1)

github 简单demo地址如下： [springboot-dynamic-task]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fcaotinging%2Fsimple-demo%2Ftree%2Fmaster%2Fspringboot-dynamic-task )

### 1.定时任务的配置类：SchedulingConfig ###

` import org.springframework.context.annotation.Bean; import org.springframework.context.annotation.Configuration; import org.springframework.scheduling.TaskScheduler; import org.springframework.scheduling.concurrent.ThreadPoolTaskScheduler; /** * @program : simple-demo * @description : 定时任务配置类 * @author : CaoTing * @date : 2019/5/23 **/ @Configuration public class SchedulingConfig { @Bean public TaskScheduler taskScheduler () { ThreadPoolTaskScheduler taskScheduler = new ThreadPoolTaskScheduler(); // 定时任务执行线程池核心线程数 taskScheduler.setPoolSize( 4 ); taskScheduler.setRemoveOnCancelPolicy( true ); taskScheduler.setThreadNamePrefix( "TaskSchedulerThreadPool-" ); return taskScheduler; } } 复制代码`

### 2.定时任务注册类：CronTaskRegistrar ###

> 
> 
> 
> 这个类包含了新增定时任务，移除定时任务等等核心功能方法
> 
> 

` import com.caotinging.demo.task.ScheduledTask; import org.springframework.beans.factory.DisposableBean; import org.springframework.beans.factory.annotation.Autowired; import org.springframework.scheduling.TaskScheduler; import org.springframework.scheduling.config.CronTask; import org.springframework.stereotype.Component; import java.util.Map; import java.util.concurrent.ConcurrentHashMap; /** * @program : simple-demo * @description : 添加定时任务注册类，用来增加、删除定时任务。 * @author : CaoTing * @date : 2019/5/23 **/ @Component public class CronTaskRegistrar implements DisposableBean { private final Map<Runnable, ScheduledTask> scheduledTasks = new ConcurrentHashMap<>( 16 ); @Autowired private TaskScheduler taskScheduler; public TaskScheduler getScheduler () { return this.taskScheduler; } /** * 新增定时任务 * @param task * @param cronExpression */ public void addCronTask (Runnable task, String cronExpression) { addCronTask( new CronTask(task, cronExpression)); } public void addCronTask (CronTask cronTask) { if (cronTask != null ) { Runnable task = cronTask.getRunnable(); if ( this.scheduledTasks.containsKey(task)) { removeCronTask(task); } this.scheduledTasks.put(task, scheduleCronTask(cronTask)); } } /** * 移除定时任务 * @param task */ public void removeCronTask (Runnable task) { ScheduledTask scheduledTask = this.scheduledTasks.remove(task); if (scheduledTask != null ) scheduledTask.cancel(); } public ScheduledTask scheduleCronTask (CronTask cronTask) { ScheduledTask scheduledTask = new ScheduledTask(); scheduledTask.future = this.taskScheduler.schedule(cronTask.getRunnable(), cronTask.getTrigger()); return scheduledTask; } @Override public void destroy () { for (ScheduledTask task : this.scheduledTasks.values()) { task.cancel(); } this.scheduledTasks.clear(); } } 复制代码`

### 3.定时任务执行类：SchedulingRunnable ###

` import com.caotinging.demo.utils.SpringContextUtils; import org.slf4j.Logger; import org.slf4j.LoggerFactory; import org.springframework.util.ReflectionUtils; import java.lang.reflect.Method; import java.util.Objects; /** * @program : simple-demo * @description : 定时任务运行类 * @author : CaoTing * @date : 2019/5/23 **/ public class SchedulingRunnable implements Runnable { private static final Logger logger = LoggerFactory.getLogger(SchedulingRunnable.class); private String beanName; private String methodName; private Object[] params; public SchedulingRunnable (String beanName, String methodName) { this (beanName, methodName, null ); } public SchedulingRunnable (String beanName, String methodName, Object...params ) { this.beanName = beanName; this.methodName = methodName; this.params = params; } @Override public void run () { logger.info( "定时任务开始执行 - bean：{}，方法：{}，参数：{}" , beanName, methodName, params); long startTime = System.currentTimeMillis(); try { Object target = SpringContextUtils.getBean(beanName); Method method = null ; if ( null != params && params.length > 0 ) { Class<?>[] paramCls = new Class[params.length]; for ( int i = 0 ; i < params.length; i++) { paramCls[i] = params[i].getClass(); } method = target.getClass().getDeclaredMethod(methodName, paramCls); } else { method = target.getClass().getDeclaredMethod(methodName); } ReflectionUtils.makeAccessible(method); if ( null != params && params.length > 0 ) { method.invoke(target, params); } else { method.invoke(target); } } catch (Exception ex) { logger.error(String.format( "定时任务执行异常 - bean：%s，方法：%s，参数：%s " , beanName, methodName, params), ex); } long times = System.currentTimeMillis() - startTime; logger.info( "定时任务执行结束 - bean：{}，方法：{}，参数：{}，耗时：{} 毫秒" , beanName, methodName, params, times); } @Override public boolean equals (Object o) { if ( this == o) return true ; if (o == null || getClass() != o.getClass()) return false ; SchedulingRunnable that = (SchedulingRunnable) o; if (params == null ) { return beanName.equals(that.beanName) && methodName.equals(that.methodName) && that.params == null ; } return beanName.equals(that.beanName) && methodName.equals(that.methodName) && params.equals(that.params); } @Override public int hashCode () { if (params == null ) { return Objects.hash(beanName, methodName); } return Objects.hash(beanName, methodName, params); } } 复制代码`

### 4.定时任务控制类：ScheduledTask ###

` import java.util.concurrent.ScheduledFuture; /** * @program : simple-demo * @description : 定时任务控制类 * @author : CaoTing * @date : 2019/5/23 **/ public final class ScheduledTask { public volatile ScheduledFuture<?> future; /** * 取消定时任务 */ public void cancel () { ScheduledFuture<?> future = this.future; if (future != null ) { future.cancel( true ); } } } 复制代码`

### 5.定时任务的测试 ###

> 
> 
> 
> 编写一个需要用于测试的任务类
> 
> 

` import org.springframework.stereotype.Component; /** * @program : simple-demo * @description : * @author : CaoTing * @date : 2019/5/23 **/ @Component ( "demoTask" ) public class DemoTask { public void taskWithParams (String param1, Integer param2) { System.out.println( "这是有参示例任务：" + param1 + param2); } public void taskNoParams () { System.out.println( "这是无参示例任务" ); } } 复制代码`
> 
> 
> 
> 
> 进行单元测试
> 
> 

` import com.caotinging.demo.application.DynamicTaskApplication; import com.caotinging.demo.application.SchedulingRunnable; import com.caotinging.demo.config.CronTaskRegistrar; import org.junit.Test; import org.junit.runner.RunWith; import org.springframework.beans.factory.annotation.Autowired; import org.springframework.boot.test.context.SpringBootTest; import org.springframework.test.context.junit4.SpringJUnit4ClassRunner; /** * @program : simple-demo * @description : 测试定时任务 * @author : CaoTing * @date : 2019/5/23 **/ @RunWith (SpringJUnit4ClassRunner.class) @SpringBootTest (classes = DynamicTaskApplication.class) public class TaskTest { @Autowired CronTaskRegistrar cronTaskRegistrar; @Test public void testTask () throws InterruptedException { SchedulingRunnable task = new SchedulingRunnable( "demoTask" , "taskNoParams" , null ); cronTaskRegistrar.addCronTask(task, "0/10 * * * * ?" ); // 便于观察 Thread.sleep( 3000000 ); } @Test public void testHaveParamsTask () throws InterruptedException { SchedulingRunnable task = new SchedulingRunnable( "demoTask" , "taskWithParams" , "haha" , 23 ); cronTaskRegistrar.addCronTask(task, "0/10 * * * * ?" ); // 便于观察 Thread.sleep( 3000000 ); } } 复制代码`

### 6.工具类：SpringContextUtils ###

` import org.springframework.beans.BeansException; import org.springframework.context.ApplicationContext; import org.springframework.context.ApplicationContextAware; import org.springframework.stereotype.Component; /** * @program : simple-demo * @description : spring获取bean工具类 * @author : CaoTing * @date : 2019/5/23 **/ @Component public class SpringContextUtils implements ApplicationContextAware { private static ApplicationContext applicationContext = null ; @Override public void setApplicationContext (ApplicationContext applicationContext) throws BeansException { if (SpringContextUtils.applicationContext == null ) { SpringContextUtils.applicationContext = applicationContext; } } //获取applicationContext public static ApplicationContext getApplicationContext () { return applicationContext; } //通过name获取 Bean. public static Object getBean (String name) { return getApplicationContext().getBean(name); } //通过class获取Bean. public static <T> T getBean (Class<T> clazz) { return getApplicationContext().getBean(clazz); } //通过name,以及Clazz返回指定的Bean public static <T> T getBean (String name, Class<T> clazz) { return getApplicationContext().getBean(name, clazz); } } 复制代码`

### 7.我的pom依赖 ###

` <dependencies> <dependency> <groupId>org.springframework.boot</groupId> <artifactId>spring-boot-starter-jdbc</artifactId> </dependency> <dependency> <groupId>com.baomidou</groupId> <artifactId>mybatisplus-spring-boot-starter</artifactId> <version>1.0.5</version> </dependency> <dependency> <groupId>com.baomidou</groupId> <artifactId>mybatis-plus</artifactId> <version>2.1.9</version> </dependency> <dependency> <groupId>mysql</groupId> <artifactId>mysql-connector-java</artifactId> <scope>runtime</scope> </dependency> <dependency> <groupId>com.alibaba</groupId> <artifactId>druid-spring-boot-starter</artifactId> <version>1.1.9</version> </dependency> <dependency> <groupId>org.springframework.boot</groupId> <artifactId>spring-boot-starter-aop</artifactId> </dependency> <dependency> <groupId>org.springframework.boot</groupId> <artifactId>spring-boot-starter-web</artifactId> </dependency> <!-- 数据库--> <!--<dependency> <groupId>mysql</groupId> <artifactId>mysql-connector-java</artifactId> <scope>runtime</scope> </dependency>--> <!-- https://mvnrepository.com/artifact/com.hynnet/oracle-driver-ojdbc --> <!--<dependency> <groupId>com.oracle</groupId> <artifactId>ojdbc6</artifactId> <version>11.2.0.1.0</version> </dependency>--> <!-- 单元测试 --> <dependency> <groupId>org.springframework.boot</groupId> <artifactId>spring-boot-starter-test</artifactId> <scope>provided</scope> </dependency> <!--redisTemplate --> <dependency> <groupId>org.springframework.boot</groupId> <artifactId>spring-boot-starter-data-redis</artifactId> </dependency> <dependency> <groupId>redis.clients</groupId> <artifactId>jedis</artifactId> <version>2.7.3</version> </dependency> <!-- http连接 restTemplate --> <dependency> <groupId>org.apache.httpcomponents</groupId> <artifactId>httpclient</artifactId> </dependency> <dependency> <groupId>org.apache.httpcomponents</groupId> <artifactId>httpclient-cache</artifactId> </dependency> <!-- 工具--> <dependency> <groupId>org.projectlombok</groupId> <artifactId>lombok</artifactId> <optional>true</optional> </dependency> <dependency> <groupId>com.alibaba</groupId> <artifactId>fastjson</artifactId> <version>1.2.31</version> </dependency> <dependency> <groupId>org.apache.commons</groupId> <artifactId>commons-lang3</artifactId> </dependency> <dependency> <groupId>commons-lang</groupId> <artifactId>commons-lang</artifactId> <version>2.6</version> </dependency> <!-- https://mvnrepository.com/artifact/com.google/guava --> <dependency> <groupId>com.google.guava</groupId> <artifactId>guava</artifactId> <version>10.0.1</version> </dependency> <!-- pinyin4j --> <dependency> <groupId>com.belerweb</groupId> <artifactId>pinyin4j</artifactId> <version>2.5.0</version> </dependency> </dependencies> 复制代码`

### 8.总结 ###

建议移步github获取简单demo上手实践哦，在本文文首哦。有帮助的话点个赞吧，笔芯。