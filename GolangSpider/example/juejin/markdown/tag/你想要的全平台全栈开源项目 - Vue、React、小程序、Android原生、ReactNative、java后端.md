# 你想要的全平台全栈开源项目 - Vue、React、小程序、Android原生、ReactNative、java后端 #

> 
> 
> 
> 2018.11.22 更新
> 
> 

感谢大家对 [coderiver]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fcachecats%2Fcoderiver ) 项目的关注和支持！

上了掘金首页推荐之后流量暴涨，截止目前，项目在 github 上已经有 575 个 ` Star` ，82 个 ` Fork` ，58 个 ` Watch` ，感谢掘金，感谢大佬们~

很多人还不太明白项目到底是干什么的，还有很多疑问。为此我们整理了两篇简单的文档介绍：

[大家关心的一些问题整理]( https://link.juejin.im?target=https%3A%2F%2Fshimo.im%2Fdocs%2F7UrBDsLgd8AGeZk4 )

[Coderiver 项目简介]( https://link.juejin.im?target=https%3A%2F%2Fshimo.im%2Fdocs%2FMDxvwTXPoK0WiSyL )

**项目最新动态：**

最近几天跟多位大佬沟通，对项目未来发展、使命和规划有了新的理解和计划。 目前正在快马加鞭筹建团队，邀请了经验丰富的架构师指导，每个技术栈都会由该领域专业的大佬把关，尽全力做精品开源项目，为大家献上健壮、优美的代码。

团队筹建完成之后各种规范文档都会相继公布，敬请期待~

欢迎持续关注， [coderiver]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fcachecats%2Fcoderiver ) 团队不会让各位大佬失望的！

> 
> 
> 
> 原文
> 
> 

全平台全栈开源项目 ` coderiver` 今天终于开始前后端联调了~

首先感谢大家的支持， [coderiver]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fcachecats%2Fcoderiver ) 在 GitHub 上开源两周，获得了 54 个 ` Star` ，9 个 ` Fork` ，5 个 ` Watch` 。

这些鼓励和认可也更加坚定了我继续写下去的决心~ 再次感谢各位大佬！

项目地址： [github.com/cachecats/c…]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fcachecats%2Fcoderiver )

靠业余时间从产品立项，到画原型图设计功能，到前端实现，再到后端实现，断断续续写了几个月，今天终于可以调试接口啦！一路走来，感谢大家的鼓励与陪伴~

## [coderiver]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fcachecats%2Fcoderiver ) 是什么？ ##

致力于打造全平台全栈精品开源项目，计划做成包含 pc端（Vue、React）、移动H5（Vue、React）、ReactNative混合开发、Android原生、微信小程序、java后端的全平台型全栈项目。

` coderiver` 中文名 河码，是一个为程序员和设计师提供项目协作的平台。无论你是前端、后端、移动端开发人员，或是设计师、产品经理，都可以在平台上发布项目，与志同道合的小伙伴一起协作完成项目。

` coderiver` 河码 类似程序员客栈，但主要目的是方便各细分领域人才之间技术交流，共同成长，多人协作完成项目。暂不涉及金钱交易。

## 技术架构 ##

目前只做了基于 Vue 的 PC 端，和基于 java 的后端。

前端的技术架构是 Vue 家族，UI 框架用的是饿了么的 [Element-ui]( https://link.juejin.im?target=http%3A%2F%2Felement-cn.eleme.io%2F%23%2Fzh-CN ).

后端采用了基于 SpringCloud 的微服务架构。整个项目分为了五个服务：

* 

注册中心 ` eureka_server`

* 

用户服务 ` user_service`

* 

项目服务 ` project_service`

* 

评论服务 ` comments_service`

* 

服务网关 ` api_gateway`

服务网关用了 Zuul ，所有接口都经过网关访问，便于统一做用户鉴权、负载均衡等操作。

各服务间通信用 Feign。多个场景都使用了 Redis ，主要是作为缓存容器使用。数据库操作暂时用的是 JPA，后期还会用 Mybatis 实现一版。消息队列暂时还没用到，后面会用 RabbitMQ。

部署的时候应该还会用 Nginx 和 Docker。

项目中用到的技术和关键的业务逻辑，都会总结出来写成博客方便大家学习参考，也希望各位大佬多多提意见，共同使项目更完善、优雅、质量更高。

## 博客汇总 ##

博客主页： [juejin.im/user/5b06d5…]( https://juejin.im/user/5b06d578f265da0de02f3b0c/posts )

已经发表的项目相关博客：

### java后端 ###

[点赞模块设计 - Redis缓存 + 定时写入数据库实现高性能点赞功能]( https://juejin.im/post/5bdc257e6fb9a049ba410098 )

[评论模块 - 后端数据库设计及功能实现]( https://juejin.im/post/5be2c213e51d453dfe02d406 )

[服务网关 Zuul 与 Redis 结合实现 Token 权限校验]( https://juejin.im/post/5bec39206fb9a049e062e4a0 )

[评论模块优化 - 数据表优化、添加缓存及用 Feign 与用户服务通信]( https://juejin.im/post/5beea202e51d451f5b54cdc4 )

### Vue pc端 ###

[vue + element-ui + scss 仿简书评论模块]( https://juejin.im/post/5b41fb58f265da0f6d72b917 )

[element-ui 的Dialog被蒙板遮住原因及解决办法]( https://juejin.im/post/5b3ec5b2f265da0f96286b4f )

## 规划 ##

对项目的规划是做成包含 pc端（Vue、React）、移动H5（Vue、React）、ReactNative混合开发、Android原生、微信小程序、java后端的全平台型全栈项目，具体平台和技术实现方案、进度如下表：

+--------------+--------------------+--------+
|     平台     |      实现方案      |  进度  |
+--------------+--------------------+--------+
| pc 端        | Vue + Element      | 90%    |
| pc 端        | React 技术栈       | 未开始 |
| 移动端 H5    | Vue 技术栈         | 未开始 |
| 移动端 H5    | React 技术栈       | 未开始 |
| 小程序       | Wepy 或 小程序原生 | 未开始 |
| 混合开发     | ReactNative        | 未开始 |
| Android 原生 | 安卓原生开发       | 未开始 |
| 后端         | java + SpringCloud | 90%    |
+--------------+--------------------+--------+

其中除了 React 技术栈，其他的我都可以做。

但考虑到时间和项目周期，以后可能会邀请其他贡献者加入。如果遇到合适的小伙伴，也可能会追加实现，比如 IOS 原生应用，Flutter 混合开发等…

所有平台，都会用当下最流行最热门的技术方案实现，代码的质量也会尽全力做到最优。

## 结语 ##

路漫漫其修远兮，吾将上下而求索。

再次感谢大家的鼓励与支持，我会继续努力，保持全速更新，争取早日实现全平台覆盖~

项目地址： [github.com/cachecats/c…]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fcachecats%2Fcoderiver )

项目讨论群：

![](https://user-gold-cdn.xitu.io/2018/11/21/16733f11a4161a75?imageView2/0/w/1280/h/960/ignore-error/1)

如果扫码进不了，加我V: ` douglas1840`

您的鼓励是我前行最大的动力，欢迎点赞，欢迎送小星星✨ ~

![](https://user-gold-cdn.xitu.io/2018/11/19/1672b2d031e9fc9a?imageslim)