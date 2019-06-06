# 【肥朝】图解源码 | MyBatis的Mapper原理 #

> 
> 
> 
> 提到看源码,很多同学内心的恐惧的,其实这个从人性的角度来说是非常正常的,因为人们对未知的事物,都是非常恐惧的,其次,你内心可能始终觉得,好像不会原理也还是能工作啊,你的潜意识里没有强烈的欲望.从阅读源码的经历来说,Java三大框架SSM中,Mybatis的源码是最适合入门的.
> 
> 
> 

## 简单使用 ##

这是一个简单的Mybatis保存对象的例子

` @Test public void testSave () throws Exception { //创建sessionFactory对象 SqlSessionFactory sf = new SqlSessionFactoryBuilder(). build(Resources.getResourceAsStream( "mybatis-config.xml" )); //获取session对象 SqlSession session = sf.openSession(); //创建实体对象 User user = new User(); user.setUsername( "toby" ); user.setPassword( "123" ); user.setAge( 23 ); //保存数据到数据库中 session.insert( "com.toby.mybatis.domain.UserMapper.add" , user); //提交事务,这个是必须要的,否则即使sql发了也保存不到数据库中 session.commit(); //关闭资源 session.close(); } 复制代码` ` < mapper namespace = "com.toby.mybatis.domain.UserMapper" > <!--#{}在传入的对象中找对应的属性值--> <!--parameterType传入的参数是什么类型--> < insert id = "add" parameterType = "com.toby.mybatis.domain.User" > INSERT INTO USER (username,password,age) VALUES (#{username},#{password},#{age}) </ insert > </ mapper > 复制代码`

## 引出主题 ##

但是在实际中,我们都不是这样操作的,我们是通过Mapper接口,调用接口方法,就能实现CRUD操作,那么关键是,这个接口究竟做了什么事,才是我们关心的.

只要把下面这段代码究竟发生了什么事弄明白,就明白,这个Mapper接口究竟做了什么事.

` public void testGetObject () throws Exception { SqlSession session = MybatisUtil.openSession(); UserMapper mapper = session.getMapper(UserMapper.class); User user = mapper.get( 5L ); System.out.println(user); session.close(); } 复制代码` ` public interface UserMapper { public void add (User user) ; public User get (Long id) ; } 复制代码`

## 流程图 ##

但是我认为,一张流程图和时序图就看明白这期间所发生的事

![](https://user-gold-cdn.xitu.io/2019/4/16/16a2692db88d50cb?imageView2/0/w/1280/h/960/ignore-error/1)

![](https://user-gold-cdn.xitu.io/2019/4/16/16a26930d32b563e?imageView2/0/w/1280/h/960/ignore-error/1)

## 写在最后 ##

**肥朝 是一个专注于 原理、源码、开发技巧的技术公众号，号内原创专题式源码解析、真实场景源码原理实战（重点）。 扫描下面二维码 关注肥朝，让本该造火箭的你，不再拧螺丝！**

![](https://user-gold-cdn.xitu.io/2019/4/16/16a2693c8f0d038a?imageslim)