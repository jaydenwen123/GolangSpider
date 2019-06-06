# Java Bean Copy 性能大比拼 #

## 简介 ##

Bean 拷贝在工作中被大量使用，可以大幅度的减少工作量。本文对常用的 Bean copy 工具进行了压力测试，方便大家选择更加适合自己的工具。本篇文章是 [mica cglib 增强——【01】cglib bean copy 介绍]( https://link.juejin.im?target=https%3A%2F%2Fwww.yuque.com%2Fdreamlu%2Fmica%2Fbean-copy-01 ) 续篇，该专栏会持续更新，感兴趣的朋友请订阅我们。

## bean 拷贝工具 ##

* [MapStruct (编译期生成 Mapper 实现)]( https://link.juejin.im?target=http%3A%2F%2Fmapstruct.org%2F )
* [Selma (编译期生成 Mapper 实现)]( https://link.juejin.im?target=http%3A%2F%2Fwww.selma-java.org%2F )
* [yangtu222 - BeanUtils (第一次生成 copy 实现字节码)]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fyangtu222%2FBeanUtils )
* [mica (第一次生成 copy 实现字节码)]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Flets-mica%2Fmica )
* [hutool (反射)]( https://link.juejin.im?target=https%3A%2F%2Fgitee.com%2Floolly%2Fhutool )

## 模型 ##

#### 无类型转换 ####

` /** * 来源用户 * * @author L.cm */ @Data public class FormUser { private Long id; private String nickName; private Integer age; private String phone; private String email; private String password; private Integer gender; private String avatar; } /** * 转换的用户 * * @author L.cm */ @Data public class ToUser { private String nickName; private String phone; private String email; private Integer gender; private String avatar; } 复制代码`

#### 带类型转换 ####

` /** * 附带类型转换的 用户模型 * * @author L.cm */ @Data @Accessors (chain = true ) public class FormConvertUser { private Long id; private String nickName; private Integer age; private String phone; private String email; private String password; private Integer gender; private String avatar; @DateTimeFormat (pattern = DateUtil.PATTERN_DATETIME) private LocalDateTime birthday; } /** * 附带类型转换的 用户模型 * * @author L.cm */ @Data @Accessors (chain = true ) public class ToConvertUser { private String nickName; private Integer age; private String phone; private String email; private String password; private Integer gender; private String avatar; private String birthday; } 复制代码`

## Bean copy 压测结果 ##

#### 环境 ####

* OS: macOS Mojave
* CPU: 2.8 GHz Intel Core i5
* RAM: 8 GB 1600 MHz DDR3
* JVM: Oracle 1.8.0_201 64 bits

#### 简单模型 ####

+-------------+------------+----------+--------+
|  BENCHMARK  |   SCORE    |  ERROR   | UNITS  |
+-------------+------------+----------+--------+
| hutool      |   1939.092 |   26.747 | ops/ms |
| spring      |   3569.035 |   39.607 | ops/ms |
| cglib       |   9112.785 |  560.503 | ops/ms |
| mica        |  17753.409 |  393.245 | ops/ms |
| yangtu222   |  18201.997 |  119.189 | ops/ms |
| cglibMapper |  37679.510 | 3544.624 | ops/ms |
| mapStruct   |  50328.045 |  529.707 | ops/ms |
| selma       | 200859.561 | 2370.531 | ops/ms |
+-------------+------------+----------+--------+

#### 附带类型转换(日期) ####

+-----------+------------+----------+--------+
| BENCHMARK |   SCORE    |  ERROR   | UNITS  |
+-----------+------------+----------+--------+
| mica      |   1186.375 |   64.686 | ops/ms |
| mapStruct |   1623.478 |   13.894 | ops/ms |
| selma     | 160020.595 | 2570.747 | ops/ms |
+-----------+------------+----------+--------+

#### 列表模型(100 item) ####

+-----------+---------+-------+--------+
| BENCHMARK |  SCORE  | ERROR | UNITS  |
+-----------+---------+-------+--------+
| spring    |  35.974 | 0.555 | ops/ms |
| mica      | 169.066 | 5.460 | ops/ms |
+-----------+---------+-------+--------+

#### Map 拷贝到 bean ####

+-----------+-----------+--------+--------+
| BENCHMARK |   SCORE   | ERROR  | UNITS  |
+-----------+-----------+--------+--------+
| hutool    |  1338.551 | 16.746 | ops/ms |
| mica      | 13577.056 | 27.795 | ops/ms |
+-----------+-----------+--------+--------+

## 结论 ##

和 [java-object-mapper-benchmark]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Farey%2Fjava-object-mapper-benchmark ) 测试结果有些出入。

Selma 的表现反而比 MapStruct 更好，可能是模型不一样导致的。

#### 功能比较 ####

+------------------------+----------------+---------+---------------+--------------------------+------+
|         工具包         | 需要编写MAPPER | 支持MAP | 支持LIST、SET |         类型转换         | 性能 |
+------------------------+----------------+---------+---------------+--------------------------+------+
| Selma                  | 是             | 否      | 否            | 需要手写转换             | 极高 |
| MapStruct              | 是             | 否      | 否            | 支持常用类型和复杂表达式 | 极高 |
| BeanUtils（yangtu222） | 否             | 否      | 是            | 需要手写转换             | 极高 |
| mica                   | 否             | 是      | 是            | 是用 Spring 的类型转换   | 极高 |
| Spring                 | 否             | 否      | 否            | 不支持                   | 高   |
| hutool                 | 否             | 是      | 否            | 不支持                   | 高   |
+------------------------+----------------+---------+---------------+--------------------------+------+

## 链接 ##

本项目源码： [github.com/lets-mica/m…]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Flets-mica%2Fmica-jmh )

## 开源推荐 ##

* Spring boot 微服务高效开发 ` mica` 工具集： [gitee.com/596392912/m…]( https://link.juejin.im?target=https%3A%2F%2Fgitee.com%2F596392912%2Fmica )
* ` Avue` 一款基于vue可配置化的神奇框架： [gitee.com/smallweigit…]( https://link.juejin.im?target=https%3A%2F%2Fgitee.com%2Fsmallweigit%2Favue )
* ` pig` 宇宙最强微服务（架构师必备）： [gitee.com/log4j/pig]( https://link.juejin.im?target=https%3A%2F%2Fgitee.com%2Flog4j%2Fpig )
* ` SpringBlade` 完整的线上解决方案（企业开发必备）： [gitee.com/smallc/Spri…]( https://link.juejin.im?target=https%3A%2F%2Fgitee.com%2Fsmallc%2FSpringBlade )
* ` IJPay` 支付SDK让支付触手可及： [gitee.com/javen205/IJ…]( https://link.juejin.im?target=https%3A%2F%2Fgitee.com%2Fjaven205%2FIJPay )

## 关注我们 ##

![如梦技术-公众号.jpg](https://user-gold-cdn.xitu.io/2019/3/29/169c6b332dc94750?imageView2/0/w/1280/h/960/ignore-error/1)

扫描上面二维码，更多 **精彩内容** 每天推荐！