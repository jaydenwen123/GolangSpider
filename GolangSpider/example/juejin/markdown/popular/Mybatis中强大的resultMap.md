# Mybatis中强大的resultMap #

## 前言 ##

在 ` Mybatis` 中，有一个强大的功能元素 ` resultMap` 。当我们希望将 ` JDBC ResultSets` 中的数据，转化为合理的Java对象时，你就能感受到它的非凡之处。正如其官方所述的那样：

> 
> 
> 
> ` resultMap` 元素是 ` MyBatis` 中最重要最强大的元素。它可以让你从 90% 的 ` JDBC ResultSets` 数据提取代码中解放出来，并在一些情形下允许你进行一些
> JDBC 不支持的操作。实际上，在为一些比如连接的复杂语句编写映射代码的时候，一份 ` resultMap` 能够代替实现同等功能的长达数千行的代码。
> ` ResultMap` 的设计思想是，对于简单的语句根本不需要配置显式的结果映射，而对于复杂一点的语句只需要描述它们的关系就行了。
> 
> 

## 一、字段映射 ##

在 ` Mybatis` 中，最简单的结果映射方式，就是通过类型别名 ` typeAliases` 来处理。

如果要这样做，那么第一步需要配置实体类包的路径：

` mybatis.type-aliases-package=com.xxx.entity`

该路径下的所有类，就会被注册到 ` TYPE_ALIASES容器` 。我们在指定返回值类型的时候，就直接用别名即可。

比如，我们有一个 ` User` 类：

` @Data public class User { private String id; private String username; private String password; private String address; private String email; } 复制代码`

如果数据库中表的字段与 ` User` 类的属性名称一致，我们就可以使用 ` resultType` 来返回。

` <select id= "getUsers" resultType= "User" > SELECT u.id, u.username, u.password, u.address, u.email FROM USER u </select> 复制代码`

当然，这是理想状态下，属性和字段名都完全一致的情况。但事实上，不一致的情况是有的，这时候我们的 ` resultMap` 就要登场了。

如果 ` User` 类保持不变，但 ` SQL` 语句发生了变化，将 ` id` 改成了 ` uid` 。

` <select id= "getUsers" resultType= "User" > SELECT u.id as uid, u.username, u.password, u.address, u.email FROM USER u </select> 复制代码`

那么，在结果集中，我们将会丢失 ` id` 数据。这时候我们就可以定义一个 ` resultMap` ，来映射不一样的字段。

` <resultMap id= "getUserByIdMap" type = "User" > <result property= "id" column= "uid" ></result> </resultMap> 复制代码`

然后，我们把上面的 ` select` 语句中的 ` resultType` 修改为 ` resultMap="getUserByIdMap"` 。

这里面 ` column` 对应的是数据库的列名或别名； ` property` 对应的是结果集的字段或属性。

这就是 ` resultMap` 最简单，也最基础的用法：字段映射。

下面，我们看看其他几种标签都是怎么应用的。

+-------------+----------------------------------------+
|  元素名称   |                  描述                  |
+-------------+----------------------------------------+
| constructor | 用于在实例化类时，注入结果到构造方法中 |
| association | 关联一个对象                           |
| collection  | 关联多个对象                           |
+-------------+----------------------------------------+

## 二、构造方法 ##

如果你希望将结果注入构造方法里，就可以用到 ` constructor` 元素。

比如，我们的 ` User` 类增加了一个构造方法：

` public User(String id, String name) { this.id = id+ "--------" ; this.username = name+ "--------" ; } 复制代码`

我们需要在 ` resultMap` 中定义 ` constructor` 元素：

` <resultMap id= "getUserByIdMap" type = "User" > <constructor> <idArg column= "id" name= "id" javaType= "string" ></idArg> <arg column= "username" name= "name" javaType= "string" ></arg> </constructor> </resultMap> 复制代码`

其中， ` column` 代表数据库字段名称或者别名； ` name` 则是构造方法中的参数名称； ` javaType` 指定了参数的类型。

如你所想，这样指定构造方法后，我们结果集中的 ` id和username` 属性都会发生变化。

` { "id" : "1001--------" , "username" : "后羿--------" , "password" : "123456" , "address" : "北京市海淀区" , "email" : "510273027@qq.com" } 复制代码`

## 三、关联 ##

在实际的业务中，我们的用户一般都会有一个角色。那么在 ` User` 类里面一般也是以一个实体类来表示。

` @Data public class User { //省略用户属性... //角色信息 private Role role; } 复制代码`

我们在查询用户的时候，如果也希望看到它的角色信息，我们会这样来写查询语句：

` <select id= "getUserById" resultType= "User" > SELECT u.id, u.username, u.password, u.address, u.email, r.id as 'role_id' , r.name as 'role_name' FROM USER u LEFT JOIN user_roles ur ON u.id = ur.user_id LEFT JOIN role r ON r.id = ur.role_id where u.id= #{id} </select> 复制代码`

如上，就要查询单个用户以及用户的角色信息。不过在这里，我们不能用 ` resultType=User` 来返回。

毕竟， ` User` 类中只有一个 ` Role` 对象，并没有 ` role_id和role_name` 字段属性。

所以，我们要使用 ` association` 来关联它们。

` <resultMap id= "userMap" type = "User" > <id property= "id" column= "id" ></id> <result property= "username" column= "username" ></result> <result property= "password" column= "password" ></result> <result property= "address" column= "address" ></result> <result property= "email" column= "email" ></result> <association property= "role" javaType= "Role" > <id property= "id" column= "role_id" ></id> <result property= "name" column= "role_name" ></result> </association> </resultMap> 复制代码`

最后我们就可以将角色信息一块显示出来：

` { "id" : "1001" , "username" : "后羿" , "password" : "123456" , "address" : "北京市海淀区" , "email" : "510273027@qq.com" , "role" : { "id" : "3" , "name" : "射手" } } 复制代码`

事实上，如果你确定关联信息是 ` 一对一` 的情况，有个更简便的方法可以替代 ` association` ，我们在本文的 ` 第五部分-自动填充关联对象` 再看它是怎么实现的。

## 四、集合 ##

#### 1、集合的嵌套结果映射 ####

上面我们看到一个用户 ` 后羿` ，它的角色是 ` 射手` ；但大部分时候，我们每个人都不可能只拥有一种角色。所以，我们需要将 ` User` 类中的角色属性的类型改成 ` List` 。

` @Data public class User { //省略用户属性... //角色信息 private List<Role> roles; } 复制代码`

现在就变成了一个用户对应多个角色，所以就不是简单的 ` association` 。

因为 ` association` 处理的是 ` 有一个` 类型的关联；而我们这里是 ` 有多个` 类型的关联，所以就需要用到 ` collection` 属性。

我们整体的 ` resultMap` 会变成下面这样：

` <resultMap id= "userMap" type = "User" > <id property= "id" column= "id" ></id> <result property= "username" column= "username" ></result> <result property= "password" column= "password" ></result> <result property= "address" column= "address" ></result> <result property= "email" column= "email" ></result> <collection property= "roles" ofType= "Role" > <id property= "id" column= "role_id" ></id> <result property= "name" column= "role_name" ></result> </collection> </resultMap> 复制代码`

这样的话，即便你有多个角色也可以被正确显示：

` { "id" : "1003" , "username" : "貂蝉" , "password" : "123456" , "address" : "北京市东城区" , "email" : "510273027@qq.com" , "roles" : [ { "id" : "1" , "name" : "中单" }, { "id" : "2" , "name" : "打野" } ] } 复制代码`

#### 2、集合的嵌套 Select 查询 ####

在大部分业务系统中，我们都会有一个菜单的表，比如像下面这样，一张 ` Menu` 表：

+------+----------+-----------------+-----------+
|  ID  |   NAME   |       URL       | PARENT ID |
+------+----------+-----------------+-----------+
|    1 | 系统管理 |                 |         0 |
| 1001 | 用户管理 | /user           |         1 |
| 1002 | 角色管理 | /role           |         1 |
| 1003 | 单位管理 | /employer       |         1 |
|    2 | 平台监控 |                 |         0 |
| 2001 | 系统监控 | /system/monitor |         2 |
| 2002 | 数据监控 | /data/monitor   |         2 |
+------+----------+-----------------+-----------+

这里我们给菜单分为两级。我们给前端返回菜单的时候，也是需要分级的，不可能将这7条数据平级展示。那么，在这里我们的 ` Menu` 实体类如下：

` @Data public class Menu { private String id; private String name; private String url; private String parent_id; private List<Menu> childMenu; } 复制代码`

一级菜单，包含一个二级菜单的列表，这里用 ` childMenu` 来表示。

` SQL` 语句中，如果没有 ` parent_id` 字段属性，我们就先查所有的一级菜单：

` <select id= "getMenus" resultMap= "menusMap" > SELECT m.id, m.name, m.url, m.parent_id FROM m_menu m where 1=1 <choose> <when test = "parent_id!=null" > and m.parent_id = #{parent_id} </when> <otherwise> and m.parent_id = '0' </otherwise> </choose> </select> 复制代码`

这个查询语句，在不传输任何参数的情况下，我们会得到两条一级菜单的数据。

那么在只调用此方法一次的情况下，怎么把所有的菜单信息查询出来，并按层级展示呢？

我们来看 ` menusMap` 的定义：

` <resultMap id= "menusMap" type = "Menu" > <id property= "id" column= "id" ></id> <result property= "name" column= "name" ></result> <result property= "url" column= "url" ></result> <result property= "m_desc" column= "m_desc" ></result> <result property= "parent_id" column= "parent_id" ></result> <collection property= "childMenu" ofType= "Menu" select= "getMenus" column= "{parent_id=id}" ></collection> </resultMap> 复制代码`

重点来看 ` collection` 元素：

` property="childMenu"` 对应的是菜单中的子级菜单列表；

` ofType="Menu"` 对应返回数据的类型；

` select="getMenus"` 指定了 ` SELECT` 语句的id；

` column="{parent_id=id}"` 则是参数的表达式。

这个 ` collection` 整体的含义可以这样理解：

通过 ` getMenus` 这个 ` SELECT语句` 来获取一级菜单中的 ` childMenu` 属性结果；在上面的 ` SELECT语句` 中，需要传递一个 ` parent_id` 参数；这个参数的值就是一级菜单中的 ` id` 。

通过这种方式，我们就可以得到已分级的所有菜单信息。

` [ { "id" : "1" , "name" : "系统管理" , "parent_id" : "0" , "childMenu" : [ { "id" : "1001" , "name" : "用户管理" , "url" : "/user" , "parent_id" : "1" }, { "id" : "1002" , "name" : "角色管理" , "url" : "/role" , "parent_id" : "1" }, { "id" : "1003" , "name" : "单位管理" , "url" : "/employer" , "parent_id" : "1" } ] }, { "id" : "2" , "name" : "平台监控" , "parent_id" : "0" , "childMenu" : [ { "id" : "2001" , "name" : "系统监控" , "url" : "/system/monitor" , "parent_id" : "2" }, { "id" : "2002" , "name" : "数据监控" , "url" : "/data/monitor" , "parent_id" : "2" } ] } ] 复制代码`

## 五、自动填充关联对象 ##

我们知道，在 ` Mybatis` 解析返回值的时候。

第一步是获取返回值类型，拿到 ` Class` 对象，然后获取构造器，设置可访问并返回实例，然后又把它包装成 ` MetaObject` 对象。

从数据库 ` rs` 中拿到结果之后，会调用 ` MetaObject.setValue(String name, Object value)` 来填充对象。

在这过程中，有趣的是，它会以 `.` 来分隔这个 ` name` 属性。

如果 ` name` 属性中包含 `.符号` ，就找到 `.符号` 之前的属性名称，把它当做一个实体对象来处理。

可能笔者在这里描述的不够直观，我们还是来看例子。

在本文 ` 第三部分` 中，我们有 ` 一个用户对应一个角色` 的例子。

其中， ` User` 类定义如下：

` @Data public class User { //省略用户属性... //角色信息 private Role role; } 复制代码`

在这里，我们无需定义 ` resultMap` ，直接返回 ` resultType=User` 即可。不过需要把 ` role` 信息的别名修改一下，重点是 `.符号`

` <select id= "getUserList" resultType= "User" > SELECT u.id, u.username, u.password, u.address, u.email, r.id as 'role.id' , r.name as 'role.name' FROM USER u LEFT JOIN user_roles ur ON u.id = ur.user_id LEFT JOIN role r ON r.id = ur.role_id </select> 复制代码`

这样，在 ` Mybatis` 解析到 ` role.id` 属性的时候，以 `.符号` 分隔之后发现， ` role` 别名对应的是 ` Role` 对象，则会先初始化 ` Role` 对象，并将值赋予 ` id` 属性。

相关代码如图：

![](https://user-gold-cdn.xitu.io/2019/6/2/16b18039810b4157?imageView2/0/w/1280/h/960/ignore-error/1)