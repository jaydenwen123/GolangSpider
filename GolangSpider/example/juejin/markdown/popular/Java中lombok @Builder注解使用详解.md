# Java中lombok @Builder注解使用详解 #

## 简介 ##

**Lombok** 大家都知道，在使用POJO过程中，它给我们带来了很多便利，省下大量写get、set方法、构造器、equal、toString方法的时间。除此之外，通过@Builder注解，lombok还可以方便的时间建造者模式。

只需要定义一个静态公共的内部类即可。代码示例如下：

` public class User { private Integer id; private String name; private String address; private User () { } private User(User origin) { this.id = origin.id; this.name = origin.name; this.address = origin.address; } public static class Builder { private User target; public Builder () { this.target = new User(); } public Builder id(Integer id) { target.id = id; return this; } public Builder name(String name) { target.name = name; return this; } public Builder address(String address) { target.address = address; return this; } public User build () { return new User(target); } } 复制代码`

**如果项目中有使用lombok的话，可以直接使用@Builder注解来实现**

改造上面的类如下：

` import lombok.Builder; import lombok.ToString; /** * @author wulongtao */ @ToString @Builder public class UserExample { private Integer id; private String name; private String address; } 复制代码`

如何使用：

` UserExample userExample = UserExample.builder() .id(1) .name( "aaa" ) .address( "bbb" ) .build(); System.out.println(userExample); 复制代码`

## 遇到问题 ##

### 在使用@Builder过程中，发现了一问题：子类的Builder对象没有父类的属性。这在使用上造成了一定的问题。 ###

对于这个问题，找到了如下解法

* 对于父类，使用@AllArgsConstructor注解
* 对于子类，手动编写全参数构造器，内部调用父类全参数构造器，在子类全参数构造器上使用@Builder注解

通过这种方式，子类Builder对象可以使用父类的所有私有属性。
但是这种解法也有两个副作用：

* 因为使用 ` @AllArgsConstructor` 注解，父类构造函数字段的顺序由声明字段的顺序决定，如果子类构造函数传参的时候顺序不一致，字段类型还一样的话，出了错不好发现
* 如果父类字段有增减，所有子类的构造器都要修改

虽然有这两个副作用，但是这种解法是我找到的唯一一种解决子类使用@Builder，能使用父类属性的方式。

参考博客评论： [Lombok’s @Builder annotation and inheritance]( https://link.juejin.im?target=https%3A%2F%2Freinhard.codes%2F2015%2F09%2F16%2Flomboks-builder-annotation-and-inheritance%2F )

### 如何在使用@Builder的模式中，加入字段的默认值。因为使用了建造者模式，那么一般在类内声明字段的时候给字段默认值的方式就是无效的，需要在建造者上动手脚。 ###

* 自定义静态内部类作为建造者，赋予默认值，再使用@Builder注解，这个时候lombok会补全已有的建造者类，进而使用默认值
* 更新的lombok有@Builder.Default声明，注解在需要默认值的字段上即可。

在评论区也有这种方式的副作用讨论，链接是： [Using Lombok’s @Builder annotation with default values]( https://link.juejin.im?target=https%3A%2F%2Freinhard.codes%2F2016%2F07%2F13%2Fusing-lomboks-builder-annotation-with-default-values%2F )

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2aa83da9fe52e?imageView2/0/w/1280/h/960/ignore-error/1)