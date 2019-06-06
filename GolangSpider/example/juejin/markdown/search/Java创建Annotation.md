# Java创建Annotation #

注解是Java很强大的部分，但大多数时候我们倾向于使用而不是去创建注解。例如，在Java源代码里不难找到Java编译器处理的@Override注解， [Spring框架]( https://link.juejin.im?target=https%3A%2F%2Fspring.io%2F ) 的@Autowired注解， 或 [Hibernate框架]( https://link.juejin.im?target=http%3A%2F%2Fhibernate.org%2F ) 使用的@Entity 注解，但我们很少看到自定义注解。虽然自定义注解是Java语言中经常被忽视的一个方面，但在开发可读性代码时它可能是非常有用的资产，同样有助于理解常见框架（如Spring或Hibernate）如何简洁地实现其目标。
在本文中，我们将介绍注解的基础知识，包括注解是什么，它们如何在示例中使用，以及如何处理它们。为了演示注解在实践中的工作原理，我们将创建一个Javascript Object Notation（JSON）序列化程序，用于处理带注解的对象并生成表示每个对象的JSON字符串。在此过程中，我们将介绍许多常见的注解块，包括Java反射框架和注解可见性问题。感兴趣的读者可以在 [GitHub]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Falbanoj2%2Fdzone-json-serializer ) 上 [找到]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Falbanoj2%2Fdzone-json-serializer ) 已完成的JSON序列化程序的源代码。

## 什么是注解？ ##

注解是应用于Java结构的装饰器，例如将元数据与类，方法或字段相关联。这些装饰器是良性的，不会自行执行任何代码，但运行时，框架或编译器可以使用它们来执行某些操作。更正式地说，Java语言规范（JLS） [第9.7节]( https://link.juejin.im?target=https%3A%2F%2Fzhuanlan.zhihu.com%2Fwrite%23jls-9.7 ) 提供了以下定义：
注解是信息与程序结构相关联的标记，但在运行时没有任何影响。
请务必注意此定义中的最后一句：注解在运行时对程序没有影响。这并不是说框架不会基于注解的存在而改变其运行时行为，而是包含注解本身的程序不会改变其运行时行为。虽然这可能看起来是细微差别，但为了掌握注解的实用性，理解这一点非常重要。
例如，某个实例的字段添加了@Autowired注解，其本身不会改变程序的运行时行为：编译器只是在运行时包含注解，但注解不执行任何代码或注入任何逻辑来改变程序的正常行为（忽略注解时的预期行为）。一旦我们在运行时引入Spring框架，我们就可以在解析程序时获得强大的依赖注入（DI）功能。通过引入注解，我们已经指示Spring框架向我们的字段注入适当的依赖项。我们将很快看到（当我们创建JSON序列化程序时）注解本身并没有完成此操作，而是充当标记，通知Spring框架我们希望将依赖项注入到带注解的字段中。

### Retention和Target ###

创建注解需要两条信息：（1）retention策略和（2）target。保留策略（retention）指定了在程序的生命周期注解应该被保留多长时间。例如，注解可以在编译时或运行时期间保留，具体取决于与注解关联的保留策略。从Java 9开始，有 [三种标准保留策略]( https://link.juejin.im?target=https%3A%2F%2Fdocs.oracle.com%2Fjavase%2F9%2Fdocs%2Fapi%2Fjava%2Flang%2Fannotation%2FRetentionPolicy.html ) ，总结如下：

+---------------+-----------------------------------------------------------------------------------------+
| **策略** **** | **描述** ****                                                                           |
| Source        | 编译器会丢弃注解                                                                        |
| Class         | 注解是在编译器生成的类文件中记录的，但不需要在运行时处理类文件的Java虚拟机（JVM）保留。 |
| Runtime       | 注解由编译器记录在类文件中，并由JVM在运行时保留                                         |
+---------------+-----------------------------------------------------------------------------------------+

正如我们稍后将看到的，注解保留的运行时选项是最常见的选项之一，因为它允许Java程序反射访问注解并基于存在的注解执行代码，以及访问与注解相关联的数据。请注意，注解只有一个关联的保留策略。
注解的目标（target）指定注解可以应用于哪个Java结构。例如，某些注解可能仅对方法有效，而其他注解可能对类和字段都有效。从Java 9开始，有 [11个标准注解目标]( https://link.juejin.im?target=https%3A%2F%2Fdocs.oracle.com%2Fjavase%2F9%2Fdocs%2Fapi%2Fjava%2Flang%2Fannotation%2FElementType.html ) ，如下表所示：

+-----------------+---------------------------------------------------------------------------------------------------------------------+
| **目标** ****   | **描述** ****                                                                                                       |
| Annotation Type | 注解另一个注解                                                                                                      |
| Constructor     | 注解构造函数                                                                                                        |
| Field           | 注解一个字段，例如类的实例变量或 [枚举常量](                                                                        |
|                 | https://link.juejin.im?target=https%3A%2F%2Fdocs.oracle.com%2Fjavase%2Ftutorial%2Fjava%2FjavaOO%2Fenum.html         |
|                 | )                                                                                                                   |
| Local variable  | 注解局部变量                                                                                                        |
| Method          | 注解类的方法                                                                                                        |
| Module          | 注解模块（Java 9中的新增功能）                                                                                      |
| Package         | 注解包                                                                                                              |
| Parameter       | 注解到方法或构造函数的参数                                                                                          |
| Type            | 注解一个类型，例如类，接口，注解类型或枚举声明                                                                      |
| Type Parameter  | 注解类型参数，例如用作通用参数形式的参数                                                                            |
| Type Use        | 注解类型的使用，例如当使用new关键字创建类型的对象时                                                                 |
|                 | ，当对象强制转换为指定类型时，类实现接口时，或者使用throws关键字声明throwable对象的类型时（有关更多信息，请参阅Type |
|                 | Annotations and Pluggable Type Systems Oracle tutorial）                                                            |
+-----------------+---------------------------------------------------------------------------------------------------------------------+

有关这些目标的更多信息，请参见 [JLS的第9.7.4节]( https://link.juejin.im?target=https%3A%2F%2Fzhuanlan.zhihu.com%2Fwrite%23jls-9.7.4 ) 。要注意，注解可以关联一个或多个目标。例如，如果字段和构造函数目标与注解相关联，则可以在字段或构造函数上使用注解。另一方面，如果注解仅关联方法目标，则将注解应用于除方法之外的任何构造都会在编译期间导致错误。

### 注解参数 ###

注解也可以具有参数。这些参数可以是基本类型（例如int或double），String，类，枚举，注解或前五种类型中任何一种的数组（参见 [JLS的第9.6.1节]( https://link.juejin.im?target=https%3A%2F%2Fzhuanlan.zhihu.com%2Fwrite%23jls-9.6.1 ) ）。将参数与注解相关联允许注解提供上下文信息或者可以参数化注解的处理器。例如，在我们的JSON序列化程序实现中，我们将允许一个可选的注解参数，该参数在序列化时指定字段的名称（如果没有指定名称，则默认使用字段的变量名称）。

## 如何创建注解？ ##

对于我们的JSON序列化程序，我们将创建一个字段注解，允许开发人员在序列化对象时标记要转换的字段名。例如，如果我们创建汽车类，我们可以使用我们的注解来注解汽车的字段（例如品牌和型号）。当我们序列化汽车对象时，生成的JSON将包括make和model键，其中值分别代表make和model字段的值。为简单起见，我们假设此注解仅用于String类型的字段，确保字段的值可以直接序列化为字符串。
要创建这样的字段注解，我们使用@interface 关键字声明一个新的注解：

` @Retention (RetentionPolicy.RUNTIME) @Target (ElementType.FIELD) public @interface JsonField { public String value () default "" ; } 复制代码`

我们声明的核心是public @interface JsonField，声明带有public修饰符的注解——允许我们的注解在任何包中使用（假设在另一个模块中正确导入包）。注解声明一个String类型value的参数，默认值为空字符串。

请注意，变量名称value具有特殊含义：它定义单元素注解（ [JLS的第9.7.3节]( https://link.juejin.im?target=https%3A%2F%2Fzhuanlan.zhihu.com%2Fwrite%23jls-9.7.3 ) ），并允许我们的注解用户向注解提供单个参数，而无需指定参数的名称。例如，用户可以使用@JsonField("someFieldName")并且不需要将注解声明为注解@JsonField(value = "someFieldName")，尽管后者仍然可以使用（但不是必需的）。包含默认值空字符串允许省略该值，value如果没有显式指定值，则导致值为空字符串。例如，如果用户使用表单声明上述注解@JsonField，则该value参数设置为空字符串。
注解声明的保留策略和目标分别使用@Retention和@Target注解指定。保留策略使用 [java.lang.annotation.RetentionPolicy]( https://link.juejin.im?target=https%3A%2F%2Fdocs.oracle.com%2Fjavase%2F9%2Fdocs%2Fapi%2Fjava%2Flang%2Fannotation%2FRetentionPolicy.html ) 枚举指定，并包含三个标准保留策略的常量。同样，指定目标为 [java.lang.annotation.ElementType]( https://link.juejin.im?target=https%3A%2F%2Fdocs.oracle.com%2Fjavase%2F9%2Fdocs%2Fapi%2Fjava%2Flang%2Fannotation%2FElementType.html ) 枚举，包括11种标准目标类型中每种类型的常量。
总之，我们创建了一个名为JsonField的public单元素注解，它在运行时由JVM保留，并且只能应用于字段。此注解只有单个参数，类型String的value，默认值为空字符串。通过创建注解，我们现在可以注解要序列化的字段。

## 如何使用注解？ ##

使用注解仅需要将注解放在适当的结构（注解的任何有效目标）之前。例如，我们可以创建一个Car类：

` public class Car { @JsonField ( "manufacturer" ) private final String make; @JsonField private final String model; private final String year; public Car (String make, String model, String year) { this.make = make; this.model = model; this.year = year; } public String getMake () { return make; } public String getModel () { return model; } public String getYear () { return year; } @Override public String toString () { return year + " " + make + " " + model; } } 复制代码`

该类使用@JsonField注解的两个主要用途：（1）具有显式值，（2）具有默认值。我们也可以使用@JsonField(value = "someName")注解一个字段，但这种样式过于冗长，并没有助于代码的可读性。因此，除非在单元素注解中包含注解参数名称可以增加代码的可读性，否则应该省略它。对于具有多个参数的注解，需要显式指定每个参数的名称来区分参数（除非仅提供一个参数，在这种情况下，如果未显式提供名称，则参数将映射到value参数）。

鉴于@JsonField注解的上述用法，我们希望将Car序列化为JSON字符串{"manufacturer":"someMake", "model":"someModel"} （注意，我们稍后将会看到，我们将忽略键manufacturer 和model在此JSON字符串的顺序）。在这之前，重要的是要注意添加@JsonField注解不会改变类Car的运行时行为。如果编译这个类，包含@JsonField注解不会比省略注解时增强类的行为。类的类文件中只是简单地记录这些注解以及参数的值。改变系统的运行时行为需要我们处理这些注解。

## 如何处理注解？ ##

处理注解是通过 [Java反射应用程序编程接口（API）完成的]( https://link.juejin.im?target=https%3A%2F%2Fdocs.oracle.com%2Fjavase%2Ftutorial%2Freflect%2F ) 。反射API允许我们编写代码来访问对象的类、方法、字段等。例如，如果我们创建一个接受Car对象的方法，我们可以检查该对象的类（即Car），并发现该类有三个字段：（1）make，（2）model和（3）year。此外，我们可以检查这些字段以发现每个字段是否都使用特定注解进行注解。
这样，我们可以遍历传递给方法的参数对象关联类的每个字段，并发现哪些字段使用@JsonField注解。如果该字段使用了@JsonField注解，我们将记录该字段的名称及其值。处理完所有字段后，我们就可以使用这些字段名称和值创建JSON字符串。
确定字段的名称需要比确定值更复杂的逻辑。如果@JsonField包含value参数的提供值（例如"manufacturer"之前使用的@JsonField("manufacturer")），我们将使用提供的字段名称。如果value参数的值是空字符串，我们知道没有显式提供字段名称（因为这是value参数的默认值），否则，显式提供了一个空字符串。后面这几种情况下，我们都将使用字段的变量名作为字段名称（例如，在private final String model声明中）。
将此逻辑组合到一个JsonSerializer类中：

` public class JsonSerializer { public String serialize (Object object) throws JsonSerializeException { try { Class<?> objectClass = requireNonNull(object).getClass(); Map<String, String> jsonElements = new HashMap<>(); for (Field field : objectClass.getDeclaredFields()) { field.setAccessible( true ); if (field.isAnnotationPresent(JsonField.class)) { jsonElements.put(getSerializedKey(field), (String) field.get(object)); } } System.out.println(toJsonString(jsonElements)); return toJsonString(jsonElements); } catch (IllegalAccessException e) { throw new JsonSerializeException(e.getMessage()); } } private String toJsonString (Map<String, String> jsonMap) { String elementsString = jsonMap.entrySet().stream().map(entry -> "\"" + entry.getKey() + "\":\"" + entry.getValue() + "\"" ).collect(Collectors.joining( "," )); return "{" + elementsString + "}" ; } private static String getSerializedKey (Field field) { String annotationValue = field.getAnnotation(JsonField.class).value(); if (annotationValue.isEmpty()) { return field.getName(); } else { return annotationValue; } } } 复制代码`

请注意，为简洁起见，已将多个功能合并到该类中。有关此序列化程序类的重构版本，请参阅codebase存储库中的 [此分支]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Falbanoj2%2Fdzone-json-serializer%2Ftree%2Fsrp_generalization ) （https://github.com/albanoj2/dzone-json-serializer/tree/srp_generalization）。我们还创建了一个异常，用于表示在serialize方法处理对象时是否发生了错误：

` public class JsonSerializeException extends Exception { private static final long serialVersionUID = - 8845242379503538623L ; public JsonSerializeException (String message) { super (message); } } 复制代码`

尽管JsonSerializer该类看起来很复杂，但它包含三个主要任务：（1）查找使用@JsonField注解的所有字段，（2）记录包含@JsonField注解的所有字段的名称（或显式提供的字段名称）和值，以及（3）将所记录的字段名称和值的键值对转换成JSON字符串。

requireNonNull(object).getClass()检查提供的对象不是null （如果是，则抛出一个NullPointerException）并获得与提供的对象关联的 [Class]( https://link.juejin.im?target=https%3A%2F%2Fdocs.oracle.com%2Fjavase%2F9%2Fdocs%2Fapi%2Fjava%2Flang%2FClass.html ) 对象。并使用此对象关联的类来获取关联的字段。接下来，我们创建String到String的Map，存储字段名和值的键值对。
随着数据结构的建立，接下来遍历类中声明的每个字段。对于每个字段，我们配置为在访问字段时禁止Java语言访问检查。这是非常重要的一步，因为我们注解的字段是私有的。在标准情况下，我们将无法访问这些字段，并且尝试获取私有字段的值将导致IllegalAccessException抛出。为了访问这些私有字段，我们必须禁止对该字段的标准Java访问检查。setAccessible(boolean) 定义如下：
返回值true 表示反射对象应禁止Java语言访问检查。false 表示反射对象应强制执行Java语言访问检查。
请注意，随着Java 9中模块的引入，使用setAccessible 方法要求将包含访问其私有字段的类的包在其模块定义中声明为open。有关更多信息，请参阅 [this explanation by Michał Szewczyk]( https://link.juejin.im?target=https%3A%2F%2Fstackoverflow.com%2Fa%2F46482376%2F2403253 ) 和 [Accessing Private State of Java 9 Modules by Gunnar Morling]( https://link.juejin.im?target=http%3A%2F%2Fin.relation.to%2F2017%2F04%2F11%2Faccessing-private-state-of-java-9-modules%2F ) 。
在获得对该字段的访问权限之后，我们检查该字段是否使用了注解@JsonField。如果是，我们确定字段的名称（通过@JsonField注解中提供的显式名称或默认名称），并在我们先前构造的map中记录名称和字段值。处理完所有字段后，我们将字段名称映射转换为JSON字符串。
处理完所有记录后，我们将所有这些字符串与逗号组合在一起。这会产生一个字符串"<fieldName1>":"<fieldValue1>","<fieldName2>":"<fieldValue2>",...。一旦这个字符串被连接起来，我们用花括号括起来，创建一个有效的JSON字符串。
为了测试这个序列化器，我们可以执行以下代码：

` Car car= new Car( "Ford" , "F150" , "2018" ); JsonSerializer serializer= new JsonSerializer(); serializer.serialize(car); 复制代码`

输出：

` { "model" : "F150" , "manufacturer" : "Ford" } 复制代码`

正如预期的那样，Car对象的maker和model字段已经被序列化，使用字段的名称作为键，字段的值作为值。请注意，JSON元素的顺序可能与上面看到的输出相反。发生这种情况是因为对于类的声明字段数组没有明确的排序，如 [getDeclaredFields文档]( https://link.juejin.im?target=https%3A%2F%2Fzhuanlan.zhihu.com%2Fwrite%23getDeclaredFields-- ) 中所述：
返回数组中的元素未排序，并且不按任何特定顺序排列。
由于此限制，JSON字符串中元素的顺序可能会有所不同。为了使元素的顺序具有确定性，我们必须自己强加排序。由于JSON对象被定义为一组无序的键值对，因此根据 [JSON标准]( https://link.juejin.im?target=http%3A%2F%2Fjson.org%2F ) ，不需要强制排序。但请注意，序列化方法的测试用例应该输出{"model":"F150","manufacturer":"Ford"} 或者{"manufacturer":"Ford","model":"F150"}。

## 结论 ##

Java注解是Java语言中非常强大的功能，但大多数情况下，我们使用标准注解（例如@Override）或通用框架注解（例如@Autowired），而不是开发人员。虽然不应使用注解来代替以面向对象的方式，但它们可以极大地简化重复逻辑。例如，我们可以注解每个可序列化字段而不是在接口中的方法创建一个toJsonString以及所有可以序列化的类实现此接口。它还将序列化逻辑与域逻辑分离，从域逻辑的简洁性中消除了手动序列化的混乱。
虽然在大多数Java应用程序中不经常使用自定义注解，但是对于Java语言的任何中级或高级用户来说，需要了解此功能。这个特性的知识不仅增强了开发人员的知识储备，同样也有助于理解最流行的Java框架中的常见注解。
[点击英文原文链接]( https://link.juejin.im?target=https%3A%2F%2Fdzone.com%2Farticles%2Fcreating-custom-annotations-in-java )
[更多文章欢迎访问:]( https://link.juejin.im?target=http%3A%2F%2Fwww.apexyun.com%2F ) http://www.apexyun.com
公众号:银河系1号
联系邮箱：public@space-explore.com
(未经同意，请勿转载)