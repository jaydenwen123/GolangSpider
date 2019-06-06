# 你需要知道的那些 Java 字节码知识 #

> 
> 
> 
> 作者简介
> 
> 
> 
> 茂功，蜂鸟物流最早的一批骨干，前后参与/主导多个重点系统设计与开发工作，目前负责代理商基础服务、网格商圈、配送范围产线，平时喜欢专研技术，主攻Java，擅长线上排障，稳定性治理。
> 
> 
> 

## 1. class文件中的数据类型 ##

每个class文件都是由8个字节为单位的字节流构成，class文件格式采用类似于C语言结构体的伪结构来描述，在这种伪结构中只有两种数据类型：无符号数和表。

* 无符号数
无符号数使用u1、u2、u4和u8分别表示1个字节、2个字节、4个字节和8个字节的无符号数。
* 表
表是由无符号数和其他表作为数据项构成的数据结构。表经常以“_info”后缀表示。

## 2. class文件结构 ##

class文件结构如下表：

+----------------+---------------------+-----------------------+------------------------------+
|      类型      |        名称         |         数量          |             说明             |
+----------------+---------------------+-----------------------+------------------------------+
| u4             | magic               |                     1 | 魔数                         |
| u2             | minor_version       |                     1 | 副版本号                     |
| u2             | major_version       |                     1 | 主版本号                     |
| u2             | constant_pool_count |                     1 | 常量池计数器                 |
| cp_info        | constant_pool       | constant_pool_count-1 | 常量池                       |
| u2             | access_flags        |                     1 | 类的访问标志                 |
| u2             | this_class          |                     1 | 类在常量池中的索引           |
| u2             | super_class         |                     1 | 父类在常量池中的索引         |
| u2             | interfaces_count    |                     1 | 接口计数器，直接父接口的数量 |
| u2             | interfaces          | interfaces_count      | 接口表                       |
| u2             | fields_count        |                     1 | 字段计数器                   |
| field_info     | fields              | fields_count          | 字段表                       |
| u2             | methods_count       |                     1 | 方法计数器                   |
| method_info    | methods             | methods_count         | 方法表                       |
| u2             | attributes_count    |                     1 | 属性计数器                   |
| attribute_info | attributes          | attributes_count      | 属性表                       |
+----------------+---------------------+-----------------------+------------------------------+

下面根据一个HelloWorld程序具体分析下class文件。
源码HelloWorld.java

` package com.xh.hello; public class HelloWorld { private static int abc = 123; public static void main(String[] args) { print ABC(); } private static void printABC () { System.out.println(abc); } } 复制代码`

使用javac编译该源文件 ` javac com/xh/hello/HelloWorld.java` ，得到HelloWorld.class文件。使用十六进制文件查看器查看此文件内容。

` cafe babe 0000 0034 0023 0a00 0700 140a 0006 0015 0900 1600 1709 0006 0018 0a00 1900 1a07 001b 0700 1c01 0003 6162 6301 0001 4901 0006 3c69 6e69 743e 0100 0328 2956 0100 0443 6f64 6501 000f 4c69 6e65 4e75 6d62 6572 5461 626c 6501 0004 6d61 696e 0100 1628 5b4c 6a61 7661 2f6c 616e 672f 5374 7269 6e67 3b29 5601 0008 7072 696e 7441 4243 0100 083c 636c 696e 6974 3e01 000a 536f 7572 6365 4669 6c65 0100 0f48 656c 6c6f 576f 726c 642e 6a61 7661 0c00 0a00 0b0c 0010 000b 0700 1d0c 001e 001f 0c00 0800 0907 0020 0c00 2100 2201 0017 636f 6d2f 7868 2f68 656c 6c6f 2f48 656c 6c6f 576f 726c 6401 0010 6a61 7661 2f6c 616e 672f 4f62 6a65 6374 0100 106a 6176 612f 6c61 6e67 2f53 7973 7465 6d01 0003 6f75 7401 0015 4c6a 6176 612f 696f 2f50 7269 6e74 5374 7265 616d 3b01 0013 6a61 7661 2f69 6f2f 5072 696e 7453 7472 6561 6d01 0007 7072 696e 746c 6e01 0004 2849 2956 0021 0006 0007 0000 0001 000a 0008 0009 0000 0004 0001 000a 000b 0001 000c 0000 001d 0001 0001 0000 0005 2ab7 0001 b100 0000 0100 0d00 0000 0600 0100 0000 0300 0900 0e00 0f00 0100 0c00 0000 2000 0000 0100 0000 04b8 0002 b100 0000 0100 0d00 0000 0a00 0200 0000 0700 0300 0800 0a00 1000 0b00 0100 0c00 0000 2600 0200 0000 0000 0ab2 0003 b200 04b6 0005 b100 0000 0100 0d00 0000 0a00 0200 0000 0b00 0900 0c00 0800 1100 0b00 0100 0c00 0000 1e00 0100 0000 0000 0610 7bb3 0004 b100 0000 0100 0d00 0000 0600 0100 0000 0400 0100 1200 0000 0200 13 复制代码`

使用 ` javap -verbose com.xh.hello.HelloWorld` 指令解析该类，得到如下内容，配合class文件一起分析。

` Classfile /Users/maogong.han/java_tmp/com/xh/hello/HelloWorld.class Last modified 2019-3-21; size 555 bytes MD5 checksum 4b275a3e082827230300dcb233141209 Compiled from "HelloWorld.java" public class com.xh.hello.HelloWorld minor version: 0 major version: 52 flags: ACC_PUBLIC, ACC_SUPER Constant pool: #1 = Methodref #7.#20 // java/lang/Object."<init>":()V #2 = Methodref #6.#21 // com/xh/hello/HelloWorld.printABC:()V #3 = Fieldref #22.#23 // java/lang/System.out:Ljava/io/PrintStream; #4 = Fieldref #6.#24 // com/xh/hello/HelloWorld.abc:I #5 = Methodref #25.#26 // java/io/PrintStream.println:(I)V #6 = Class #27 // com/xh/hello/HelloWorld #7 = Class #28 // java/lang/Object #8 = Utf8 abc #9 = Utf8 I #10 = Utf8 <init> #11 = Utf8 ()V #12 = Utf8 Code #13 = Utf8 LineNumberTable #14 = Utf8 main #15 = Utf8 ([Ljava/lang/String;)V #16 = Utf8 printABC #17 = Utf8 <clinit> #18 = Utf8 SourceFile #19 = Utf8 HelloWorld.java #20 = NameAndType #10:#11 // "<init>":()V #21 = NameAndType #16:#11 // printABC:()V #22 = Class #29 // java/lang/System #23 = NameAndType #30:#31 // out:Ljava/io/PrintStream; #24 = NameAndType #8:#9 // abc:I #25 = Class #32 // java/io/PrintStream #26 = NameAndType #33:#34 // println:(I)V #27 = Utf8 com/xh/hello/HelloWorld #28 = Utf8 java/lang/Object #29 = Utf8 java/lang/System #30 = Utf8 out #31 = Utf8 Ljava/io/PrintStream; #32 = Utf8 java/io/PrintStream #33 = Utf8 println #34 = Utf8 (I)V { public com.xh.hello.HelloWorld(); descriptor: ()V flags: ACC_PUBLIC Code: stack=1, locals=1, args_size=1 0: aload_0 1: invokespecial #1 // Method java/lang/Object."<init>":()V 4: return LineNumberTable: line 3: 0 public static void main(java.lang.String[]); descriptor: ([Ljava/lang/String;)V flags: ACC_PUBLIC, ACC_STATIC Code: stack=0, locals=1, args_size=1 0: invokestatic #2 // Method printABC:()V 3: return LineNumberTable: line 7: 0 line 8: 3 static {}; descriptor: ()V flags: ACC_STATIC Code: stack=1, locals=0, args_size=0 0: bipush 123 2: putstatic #4 // Field abc:I 5: return LineNumberTable: line 4: 0 } SourceFile: "HelloWorld.java" 复制代码`

### 2.1 魔数与class文件版本号 ###

* 魔数是class文件的前4个字节，是一个固定值:0xcafebabe。该值唯一的作用就是表示文件是否可以被JVM接受，不严格的说就是表示该文件是否是class文件。
* 版本号

![](https://user-gold-cdn.xitu.io/2019/4/1/169d6e633b26dee8?imageView2/0/w/1280/h/960/ignore-error/1)

### 2.2 常量池 ###

* 常量池计数器
常量池计数器表示常量池中的项的个数，在class文件中的位置是主版本号之后的2个字节，也就是第9和第10个字节。常量池计数器是从1开始的，在constant_pool表中，只有索引大于0且小于constant_pool_count的项才是有效的。
本例中constant_pool_count的值是0x0023，换算成十进制是35，表示constant_pool中有34项，有效索引是1到34，刚好是用javap解析出来的34个常量。
* 常量池
class文件中，constant_pool_count之后紧接着就是常量池的内容。每个常量池项(cp_info)都是由一个u1类型的tag和一个具体类型的表构成，具体类型由tag的值决定。如下表：

+-------+----------------------------------+
| TAG值 |            对应的类型            |
+-------+----------------------------------+
|     7 | CONSTANT_Class_info              |
|     9 | CONSTANT_Fieldref_info           |
|    10 | CONSTANT_Methodref_info          |
|    11 | CONSTANT_InterfaceMethodref_info |
|     8 | CONSTANT_String_info             |
|     3 | CONSTANT_Integer_info            |
|     4 | CONSTANT_Float_info              |
|     5 | CONSTANT_Long_info               |
|     6 | CONSTANT_Double_info             |
|    12 | CONSTANT_NameAndType_info        |
|     1 | CONSTANT_Utf8_info               |
|    15 | CONSTANT_MethodHandle_info       |
|    16 | CONSTANT_MethodType_info         |
|    18 | CONSTANT_InvokeDynamic_info      |
+-------+----------------------------------+

截取上文反编译出来的常量池部分信息，来分析常量池中的第一个常量。

` #1 = Methodref #7.#20 // java/lang/Object."<init>":()V #7 = Class #28 // java/lang/Object #10 = Utf8 <init> #11 = Utf8 ()V #20 = NameAndType #10:#11 // "<init>":()V #28 = Utf8 java/lang/Object 复制代码`

![](https://user-gold-cdn.xitu.io/2019/4/1/169d6e521ac081a8?imageView2/0/w/1280/h/960/ignore-error/1)

"#1"表示常量池中索引是1。class文件中的0x0a位置开始。类型是Methodref。Methodref类型的结构如下：

` CONSTANT_Methodref_info { u1 tag; u2 class_index; u2 name_and_type_index; } 复制代码`

Methodref中tag的值为0x0a，十进制为10，正好表示CONSTANT_Methodref_info类型。
class_index的值为0x0007，十进制为7，指向索引为7的常量池的项。#7是CONSTANT_Class_info类型，指向CONSTANT_Utf8_info类型的#28，表示此常量属于java/lang/Object的。
name_and_type_index的值为0x0014，十进制为20，指向索引为20的常量池的项，此项是NameAndType（字段或方法）类型，方法名索引（name_index）指向常量池的#10，为一个CONSTANT_Utf8_info类型，表示方法名为"<init>"；NameAndType的方法描述索引（descriptor_index）指向常量池的#11，表示无参类型。

CONSTANT_Class_info类型结构:

` CONSTANT_Class_info { u1 tag; u2 name_index; } 复制代码`

tag值为7，表示是CONSTANT_Class_info类型。 name_index是指向常量池中一个类型为CONSTANT_Utf8_info的常量索引，表示类或者接口的名字。

CONSTANT_Utf8_info类型结构：

` CONSTANT_Utf8_info { u1 tag; u2 length; u1 bytes[length]; } 复制代码`

tag值为1，表示CONSTANT_Utf8_info类型。bytes指的是字符串值的bytes数组。 bytes表示的字符串和十六进制转换可由下程序完成：

` public static String print HexString(byte[] b) { StringBuilder sb = new StringBuilder(); for (int i = 0; i < b.length; i++) { String hex = Integer.toHexString(b[i] & 0xFF); if (hex.length() == 1) { hex = '0' + hex; } sb.append(hex); } return sb.toString(); } 复制代码`

上文提到的各项类型结构和说明可参考 [《Java虚拟机规范》]( https://link.juejin.im?target=https%3A%2F%2Fdocs.oracle.com%2Fjavase%2Fspecs%2Findex.html ) 。

### 2.3 访问标志符 ###

在常量池之后，紧挨着是占2个字节的访问标志符：0x0021。
ACC_PUBLIC(0x0001)+ACC_SUPER (0x0020)。
access_flags表示类或接口的访问权限。其取值和含义见下表：

+----------------+--------+-------------------------------------------------------+
|     标记名     |   值   |                         含义                          |
+----------------+--------+-------------------------------------------------------+
| ACC_PUBLIC     | 0x0001 | 为public类型                                          |
| ACC_FINAL      | 0x0010 | 是否为final类型，只有类可设置                         |
| ACC_SUPER      | 0x0020 | 当用到invokespecial指令时，是否需要特殊处理的父类方法 |
| ACC_INTERFACE  | 0x0200 | 标识接口，不是类                                      |
| ACC_ABSTRACT   | 0x0400 | 标识是否为abstract，是否可以实例化                    |
| ACC_SYNTHETIC  | 0x1000 | 标识并非由Java源码生成的代码，而是由编译器生成的      |
| ACC_ANNOTATION | 0x2000 | 注解类型                                              |
| ACC_ENUM       | 0x4000 | 枚举类型                                              |
+----------------+--------+-------------------------------------------------------+

### 2.4 类索引、父类索引和接口索引 ###

访问标记符之后，紧接着是类索引、父类索引和接口索引。
类索引和父类索引都是一个u2类型的数据，接口索引是一组u2类型数据的集合。他们的值都表示在常量池中的索引。另外这三项数据确定了类的关系：单继承、多实现。

* 类索引（this_class）
类索引在常量池中的索引值为0x0006，十进制为6，指向一个CONSTANT_Class_info类型的常量，其tag值为7，name_index指向27的索引。

` #6 = Class #27 // com/xh/hello/HelloWorld #27 = Utf8 com/xh/hello/HelloWorld 复制代码`

![](https://user-gold-cdn.xitu.io/2019/4/1/169d6e55f1da48ff?imageView2/0/w/1280/h/960/ignore-error/1)

* 父类索引（super_class）
类索引之后，是父类索引。在常量池中的索引值为0x0007，十进制为7，指向一个CONSTANT_Class_info类型的常量，其tag值为7，name_index指向28的索引。

` #7 = Class #28 // java/lang/Object #28 = Utf8 java/lang/Object 复制代码`

![](https://user-gold-cdn.xitu.io/2019/4/1/169d6e5861278fe4?imageView2/0/w/1280/h/960/ignore-error/1)

* 接口索引(interfaces) 在父类索引之后的内容是接口索引计数器(interfaces_count)和接口索引表(interfaces)，class文件中interfaces_count的值为0x0000，表示未实现任何接口，这里不在讨论。

### 2.5 字段 ###

在接口索引表之后是字段索引计数器和字段索引表。字段索引计数器是一个u2类型的数值，class文件中的值为0x0001，表示有一个字段。
字段表中的每项都表示指向常量池中的一个索引，该索引指向一个field_info结构的数据。字段表描述当前类或接口声明的所有字段，但不包括从父类或接口中继承过来的。

` field_info { u2 access_flags; u2 name_index; u2 descriptor_index; u2 attributes_count; attribute_info attrubutes[attributes_count]; } 复制代码`

* access_flags项定义字段的访问权限和基础属性，如下表：

字段access_flags表：

+---------------+--------+-------------------------------------+
|    标记名     |   值   |                说明                 |
+---------------+--------+-------------------------------------+
| ACC_PUBLIC    | 0x0001 | public，字段可以被从任何package访问 |
| ACC_PRIVATE   | 0x0002 | private，字段只可以被该类自身访问   |
| ACC_PROTECTED | 0x0004 | protected，字段可以被子类访问       |
| ACC_STATIC    | 0x0008 | static，静态字段                    |
| ACC_FINAL     | 0x0010 | final，字段定义后无法修改           |
| ACC_VOLATILE  | 0x0040 | volatile字段                        |
| ACC_TRANSIENT | 0x0080 | transient，是否被序列化             |
| ACC_SYNTHETIC | 0x1000 | 是否编译器自动生成                  |
| ACC_ENUM      | 0x4000 | 是否为枚举                          |
+---------------+--------+-------------------------------------+

* name_index项是常量池的一个索引，该索引指向一个CONSTANT_Utf8_info类型，表示字段的非全限定名。
* descriptor_index项是常量池的一个索引，该索引指向一个CONSTANT_Utf8_info类型，表示字段的描述符。
字段描述符如下表：

+------------+-----------+---------------------+
|    字符    |   类型    |        说明         |
+------------+-----------+---------------------+
| B          | byte      |                     |
| C          | char      |                     |
| D          | double    |                     |
| F          | float     |                     |
| I          | int       |                     |
| J          | long      |                     |
| S          | short     |                     |
| Z          | boolean   |                     |
| LClassname | reference | 一个Classname的实例 |
| [          | reference | 一个一维数组        |
+------------+-----------+---------------------+

* attribute_info表示的是字段的附加属性。

本例class文件中access_flags的值为0x000a：ACC_PRIVATE(0x0002) + ACC_STATIC(0x0008)。name_index的值为0x0008，指向常量池的索引为8。descriptor_index的值为0x0009，指向常量池的索引为9。附加属性的值为0x0000，表示没有属性。综上该字段是一个被private和static修饰的int类型的字段，名称是"abc"。

` #8 = Utf8 abc #9 = Utf8 I 复制代码`

![](https://user-gold-cdn.xitu.io/2019/4/1/169d6e5c107317c3?imageView2/0/w/1280/h/960/ignore-error/1)

### 2.6 方法区域 ###

字段之后，紧接着是方法区域。有方法计数器(methods_count)和方法表(methods)。方法计数器是一个u2类型的数值，本例class中值为0x0004，表示有4个方法。方法表中每一项都是method_info结构。

` method_info { u2 access_flags; u2 name_index; u2 descriptor_index; u2 attributes_count; attribute_info attributes[attributes_count]; } 复制代码`

* access_flags表示方法的访问权限和基本属性。如下表：
方法access_flags表：

+------------------+--------+--------------------------------------+
|      标记名      |   值   |                 说明                 |
+------------------+--------+--------------------------------------+
| ACC_PUBLIC       | 0x0001 | public，方法可以被从任何package访问  |
| ACC_PRIVATE      | 0x0002 | private，方法只可以被该类自身访问    |
| ACC_PROTECTED    | 0x0004 | protected，方法可以被子类访问        |
| ACC_STATIC       | 0x0008 | static，静态方法                     |
| ACC_FINAL        | 0x0010 | final，方法不能被重写                |
| ACC_SYNCHRONIZED | 0x0020 | synchronized，方法加同步             |
| ACC_BRIDGE       | 0x0040 | bridge，方法由编译器生成             |
| ACC_VARARGS      | 0x0080 | 方法有可变参数                       |
| ACC_NATIVE       | 0x0100 | native，方法引用非Java语言的本地方法 |
| ACC_ABSTRACT     | 0x0400 | abstract，抽象方法                   |
| ACC_STRICT       | 0x0800 | strictfp，方法使用FP-strict浮点格式  |
| ACC_SYNTHETIC    | 0x1000 | 方法在源文件中不出现，由编译器产生   |
+------------------+--------+--------------------------------------+

* name_index项是常量池的一个索引，该索引指向一个CONSTANT_Utf8_info类型，表示方法的非全限定名或者初始化方法的名字(<init>或<clinit>方法)。
* descriptor_index项是常量池的一个索引，该索引指向一个CONSTANT_Utf8_info类型，表示方法的描述符。
* attributes_count和attributes分别表示方法附加属性的计数器和附加属性表。

这里分析第一个方法。方法计数器0x0004之后，是第一个方法的access_flags，值为0x0001，表示public类型。接下来是name_index，值为0x0001，指向常量池索引为1的项，该项表示的是java/lang/Object."<init>":()V方法。接下来是descriptor_index，值为0x000a，指向常量池索引为10的项，表示方法的非全限定名。接下来是attributes_count，值为0x000b，表示有11个附加属性，之后是这11个附加属性的数据，包含code和操作符，这里不在展开，之后会专门写解析的内容。

` #1 = Methodref #7.#20 // java/lang/Object."<init>":()V #7 = Class #28 // java/lang/Object #10 = Utf8 <init> #11 = Utf8 ()V #20 = NameAndType #10:#11 // "<init>":()V #28 = Utf8 java/lang/Object 复制代码`

![](https://user-gold-cdn.xitu.io/2019/4/1/169d6e5f118c3c9b?imageView2/0/w/1280/h/960/ignore-error/1)

### 2.7 属性 ###

这里主要记录文件的属性。有属性计数器attributes_count和属性表attributes。
本例class文件中，attributes_count值为0x0001，表示有一个属性。属性表中的每一项都常量池中的一个索引，该索引处的格式为：

` attribute_info { u2 attribute_name_index; u4 attribute_length; u2 source_file_index; } 复制代码`

attribute_name_index的值为0x0012，指向索引为18的常量池的项，该项是一个CONSTANT_Utf8_info结构，表示“SourceFile”。然后是attribute_length的值为0x00000002，表示紧跟其后的有2个字节，source_file_index值为0x0013，指向索引为19的常量池项，该项是一个CONSTANT_Utf8_info结构，表示“HelloWorld.java”。至此class文件简单分析完毕。

` #18 = Utf8 SourceFile #19 = Utf8 HelloWorld.java 复制代码`

### 2.8 附class文件结构备注 ###

` 魔数: cafe babe 副版本号和主版本号: 0000 0034 常量池： 0023 0a00 0700 140a 0006 0015 0900 1600 1709 0006 0018 0a00 1900 1a07 001b 0700 1c01 0003 6162 6301 0001 4901 0006 3c69 6e69 743e 0100 0328 2956 0100 0443 6f64 6501 000f 4c69 6e65 4e75 6d62 6572 5461 626c 6501 0004 6d61 696e 0100 1628 5b4c 6a61 7661 2f6c 616e 672f 5374 7269 6e67 3b29 5601 0008 7072 696e 7441 4243 0100 083c 636c 696e 6974 3e01 000a 536f 7572 6365 4669 6c65 0100 0f48 656c 6c6f 576f 726c 642e 6a61 7661 0c00 0a00 0b0c 0010 000b 0700 1d0c 001e 001f 0c00 0800 0907 0020 0c00 2100 2201 0017 636f 6d2f 7868 2f68 656c 6c6f 2f48 656c 6c6f 576f 726c 6401 0010 6a61 7661 2f6c 616e 672f 4f62 6a65 6374 0100 106a 6176 612f 6c61 6e67 2f53 7973 7465 6d01 0003 6f75 7401 0015 4c6a 6176 612f 696f 2f50 7269 6e74 5374 7265 616d 3b01 0013 6a61 7661 2f69 6f2f 5072 696e 7453 7472 6561 6d01 0007 7072 696e 746c 6e01 0004 2849 2956 访问标记符: 0021 类在常量池中的索引: 0006 父类在常量池中的索引: 0007 接口索引计数器: 0000 字段索引计数器: 0001 第一个字段field_info: 000a access_flags 0008 name_index 0009 descriptor_index 0000 附加属性 方法计数器: 0004 方法表: 0001 000a 000b 0001 000c 0000 001d 0001 0001 0000 0005 2ab7 0001 b100 0000 0100 0d00 0000 0600 0100 0000 0300 0900 0e00 0f00 0100 0c00 0000 2000 0000 0100 0000 04b8 0002 b100 0000 0100 0d00 0000 0a00 0200 0000 0700 0300 0800 0a00 1000 0b00 0100 0c00 0000 2600 0200 0000 0000 0ab2 0003 b200 04b6 0005 b100 0000 0100 0d00 0000 0a00 0200 0000 0b00 0900 0c00 0800 1100 0b00 0100 0c00 0000 1e00 0100 0000 0000 0610 7bb3 0004 b100 0000 0100 0d00 0000 0600 0100 0000 04 属性计数器: 00 01 第一个属性: 00 12 attribute_name_index 00 0000 02 attribute_length 00 13 source_file_index 复制代码`

## 3. 参考资料 ##

* [The Java Virtual Machine Instruction Set]( https://link.juejin.im?target=https%3A%2F%2Fdocs.oracle.com%2Fjavase%2Fspecs%2Fjvms%2Fse8%2Fhtml%2Fjvms-6.html )
* [《Java虚拟机规范》]( https://link.juejin.im?target=https%3A%2F%2Fbook.douban.com%2Fsubject%2F25792515%2F )

阅读博客还不过瘾？

> 
> 
> 
> 欢迎大家扫二维码通过添加群助手，加入交流群，讨论和博客有关的技术问题，还可以和博主有更多互动
> 
> ![](https://user-gold-cdn.xitu.io/2018/12/26/167e9cc24048932b?imageView2/0/w/1280/h/960/ignore-error/1)
> 博客转载、线下活动及合作等问题请邮件至 shadowfly_zyl@hotmail.com 进行沟通
> 
> 
> 
>