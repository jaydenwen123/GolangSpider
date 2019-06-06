# Spring-Boot-操作-Redis，三种方案全解析！ #

在 Redis 出现之前，我们的缓存框架各种各样，有了 Redis ，缓存方案基本上都统一了，关于 Redis，松哥之前有一个系列教程，尚不了解 Redis 的小伙伴可以参考这个教程：

* [Redis 教程合集]( https://link.juejin.im?target=https%3A%2F%2Fmp.weixin.qq.com%2Fs%2F0Bukv4bFGMWosG8SWWoDDw )

使用 Java 操作 Redis 的方案很多，Jedis 是目前较为流行的一种方案，除了 Jedis ，还有很多其他解决方案，如下：

![](https://user-gold-cdn.xitu.io/2019/6/3/16b1ad5055e2a45b?imageView2/0/w/1280/h/960/ignore-error/1)

除了这些方案之外，还有一个使用也相当多的方案，就是 Spring Data Redis。

在传统的 SSM 中，需要开发者自己来配置 Spring Data Redis ，这个配置比较繁琐，主要配置 3 个东西：连接池、连接器信息以及 key 和 value 的序列化方案。

在 Spring Boot 中，默认集成的 Redis 就是 Spring Data Redis，默认底层的连接池使用了 lettuce ，开发者可以自行修改为自己的熟悉的，例如 Jedis。

Spring Data Redis 针对 Redis 提供了非常方便的操作模板 RedisTemplate 。这是 Spring Data 擅长的事情，那么接下来我们就来看看 Spring Boot 中 Spring Data Redis 的具体用法。

# 方案一：Spring Data Redis #

## 创建工程 ##

创建工程，引入 Redis 依赖：

![](https://user-gold-cdn.xitu.io/2019/6/3/16b1ad50142350c8?imageView2/0/w/1280/h/960/ignore-error/1)

创建成功后，还需要手动引入 commos-pool2 的依赖，因此最终完整的 pom.xml 依赖如下：

` < dependencies > < dependency > < groupId > org.springframework.boot </ groupId > < artifactId > spring-boot-starter-data-redis </ artifactId > </ dependency > < dependency > < groupId > org.springframework.boot </ groupId > < artifactId > spring-boot-starter-web </ artifactId > </ dependency > </ dependencies > < dependency > < groupId > org.apache.commons </ groupId > < artifactId > commons-pool2 </ artifactId > </ dependency > 复制代码`

这里主要就是引入了 Spring Data Redis + 连接池。

## 配置 Redis 信息 ##

接下来配置 Redis 的信息，信息包含两方面，一方面是 Redis 的基本信息，另一方面则是连接池信息:

` spring.redis.database=0 spring.redis.password=123 spring.redis.port=6379 spring.redis.host=192.168.66.128 spring.redis.lettuce.pool.min-idle=5 spring.redis.lettuce.pool.max-idle=10 spring.redis.lettuce.pool.max-active=8 spring.redis.lettuce.pool.max-wait=1ms spring.redis.lettuce.shutdown-timeout=100ms 复制代码`

## 自动配置 ##

当开发者在项目中引入了 Spring Data Redis ，并且配置了 Redis 的基本信息，此时，自动化配置就会生效。

我们从 Spring Boot 中 Redis 的自动化配置类中就可以看出端倪：

` @Configuration @ConditionalOnClass (RedisOperations.class) @EnableConfigurationProperties (RedisProperties.class) @Import ({ LettuceConnectionConfiguration.class, JedisConnectionConfiguration.class }) public class RedisAutoConfiguration { @Bean @ConditionalOnMissingBean (name = "redisTemplate" ) public RedisTemplate<Object, Object> redisTemplate ( RedisConnectionFactory redisConnectionFactory) throws UnknownHostException { RedisTemplate<Object, Object> template = new RedisTemplate<>(); template.setConnectionFactory(redisConnectionFactory); return template; } @Bean @ConditionalOnMissingBean public StringRedisTemplate stringRedisTemplate ( RedisConnectionFactory redisConnectionFactory) throws UnknownHostException { StringRedisTemplate template = new StringRedisTemplate(); template.setConnectionFactory(redisConnectionFactory); return template; } } 复制代码`

这个自动化配置类很好理解：

* 首先标记这个是一个配置类，同时该配置在 RedisOperations 存在的情况下才会生效(即项目中引入了 Spring Data Redis)
* 然后导入在 application.properties 中配置的属性
* 然后再导入连接池信息（如果存在的话）
* 最后，提供了两个 Bean ，RedisTemplate 和 StringRedisTemplate ，其中 StringRedisTemplate 是 RedisTemplate 的子类，两个的方法基本一致，不同之处主要体现在操作的数据类型不同，RedisTemplate 中的两个泛型都是 Object ，意味者存储的 key 和 value 都可以是一个对象，而 StringRedisTemplate 的 两个泛型都是 String ，意味者 StringRedisTemplate 的 key 和 value 都只能是字符串。如果开发者没有提供相关的 Bean ，这两个配置就会生效，否则不会生效。

## 使用 ##

接下来，可以直接在 Service 中注入 StringRedisTemplate 或者 RedisTemplate 来使用：

` @Service public class HelloService { @Autowired RedisTemplate redisTemplate; public void hello () { ValueOperations ops = redisTemplate.opsForValue(); ops.set( "k1" , "v1" ); Object k1 = ops.get( "k1" ); System.out.println(k1); } } 复制代码`

Redis 中的数据操作，大体上来说，可以分为两种：

* 针对 key 的操作，相关的方法就在 RedisTemplate 中
* 针对具体数据类型的操作，相关的方法需要首先获取对应的数据类型，获取相应数据类型的操作方法是 opsForXXX

调用该方法就可以将数据存储到 Redis 中去了，如下：

![](https://user-gold-cdn.xitu.io/2019/6/3/16b1ad50467a36ad?imageView2/0/w/1280/h/960/ignore-error/1)

k1 前面的字符是由于使用了 RedisTemplate 导致的，RedisTemplate 对 key 进行序列化之后的结果。

RedisTemplate 中，key 默认的序列化方案是 JdkSerializationRedisSerializer 。

而在 StringRedisTemplate 中，key 默认的序列化方案是 StringRedisSerializer ，因此，如果使用 StringRedisTemplate ，默认情况下 key 前面不会有前缀。

不过开发者也可以自行修改 RedisTemplate 中的序列化方案，如下:

` @Service public class HelloService { @Autowired RedisTemplate redisTemplate; public void hello () { redisTemplate.setKeySerializer( new StringRedisSerializer()); ValueOperations ops = redisTemplate.opsForValue(); ops.set( "k1" , "v1" ); Object k1 = ops.get( "k1" ); System.out.println(k1); } } 复制代码`

当然也可以直接使用 StringRedisTemplate：

` @Service public class HelloService { @Autowired StringRedisTemplate stringRedisTemplate; public void hello2 () { ValueOperations ops = stringRedisTemplate.opsForValue(); ops.set( "k2" , "v2" ); Object k1 = ops.get( "k2" ); System.out.println(k1); } } 复制代码`

另外需要注意 ，Spring Boot 的自动化配置，只能配置单机的 Redis ，如果是 Redis 集群，则所有的东西都需要自己手动配置，关于如何操作 Redis 集群，松哥以后再来和大家分享。

# 方案二：Spring Cache #

通过 Spring Cache 的形式来操作 Redis，Spring Cache 统一了缓存江湖的门面，这种方案，松哥之前有过一篇专门的文章介绍，小伙伴可以移步这里： [Spring Boot中，Redis缓存还能这么用！]( https://link.juejin.im?target=https%3A%2F%2Fmp.weixin.qq.com%2Fs%2FUpTewC66iJyzq0osm_0cfw ) 。

# 方案三：回归原始时代 #

第三种方案，就是直接使用 Jedis 或者 其他的客户端工具来操作 Redis ，这种方案在 Spring Boot 中也是支持的，虽然操作麻烦，但是支持，这种操作松哥之前也有介绍的文章，因此这里就不再赘述了，可以参考 [Jedis 使用]( https://link.juejin.im?target=https%3A%2F%2Fmp.weixin.qq.com%2Fs%2F4ghSyd8BIq01LYhhjToXrg ) 。

# 总结 #

Spring Boot 中，Redis 的操作，这里松哥给大家总结了三种方案，实际上前两个使用广泛一些，直接使用 Jedis 还是比较少，基本上 Spring Boot 中没见过有人直接这么搞。

好了，本文就说到这里，有问题欢迎留言讨论。

关注公众号牧码小子，专注于 Spring Boot+微服务，定期视频教程分享，关注后回复 Java ，领取松哥为你精心准备的 Java 干货！

![](https://user-gold-cdn.xitu.io/2019/5/28/16afbe80e51bd0ee?imageView2/0/w/1280/h/960/ignore-error/1)