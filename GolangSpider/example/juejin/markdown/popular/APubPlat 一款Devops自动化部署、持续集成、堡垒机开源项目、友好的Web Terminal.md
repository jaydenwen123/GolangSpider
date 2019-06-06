# APubPlat 一款Devops自动化部署、持续集成、堡垒机开源项目、友好的Web Terminal #

嗨、很高心你能进入这里，我是zane, 在这里给你介绍一款完整的Devops自动化部署工具

APubPlat - 一款完整的Devops自动化部署、持续集成、堡垒机、并且友好的Web Terminal开源项目。

如果你对它感兴趣，就给一个小小的关注吧，一款好的产品更需要碰撞和火花。：

github address : [github.com/wangweiange…]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fwangweianger%2FAPubPlat )

document : [apub-wiki.seosiwei.com]( https://link.juejin.im?target=http%3A%2F%2Fapub-wiki.seosiwei.com%2F )

接下来我还会持续的更新和迭代。

## 功能描述 ##

* 资产管理： 方便快捷的管理资产，可为资产分组，为应用分配不同的资产，快捷控制台管理等。
* 

应用管理：可建立各种应用任务，前端，后端发布任务，可同时执行单机和多机任务，并实时显示任务日志。

* 

WEB控制台： 一套强大的Web Terminal，可直接替代Xshell等工具，可单个或批量打开窗口或执行命令（已支持linux系统，后期版本支持windows系统）。

* 

脚本管理：可为单个或者多个资产预装各种软装或者执行各种命令，可自由自定义各种预装脚本，例如安装nginx

* 

单|多机脚本生成：可同时为单机或者多机器同时生成shell脚本到指定的目录，方便统一管理和操作。

* 

备份还原：单多机可同时备份，并按详细日期进行备份，可随时随意一键恢复任意历史版本。

## 应用场景 ##

* 各种前端静态发布（例如：vue,react,jquery之类的纯前端持续集成）
* 前端中间层发布（例如：使用node.js开发的前端中间层之类的服务持续集成）
* 后端发布 （不限制后端语言，只依赖于shell脚本）
* 单机 | 多台机器 同时发布、备份、还原
* web版本的xshell，让你不管何时何地都能方便的管理服务器资源
* 强大的权限管理能力，为不同角色分配不同的管理权限，让我们的持续集成更灵活更方便

## 安装环境 ##

APubPlat依赖的环境并不复杂，对软硬件的要求也并不高，一台1G双核的服务器都能搞定。

APubPlat 开发技术基于egg.js、vue.js, 因此只需要安装node环境，node.js版本推荐 8.9.0 ~ 10.15.1 之间

数据库基于mongudb、环境数据库基于redis、web服务器基于nginx，所有的软件和服务你都可以安装在一台机器中。

如果想了解更多你可以选择去查看项目文档： [apub-wiki.seosiwei.com]( https://link.juejin.im?target=http%3A%2F%2Fapub-wiki.seosiwei.com%2F )

## 项目预览 ##

### 登录界面、第一次使用时请注册admin账号，其他账号在后台中进行新增和编辑管理 ###

![](https://user-gold-cdn.xitu.io/2019/6/5/16b26f033a78380f?imageView2/0/w/1280/h/960/ignore-error/1)

### 你可以自定义任何适合你的项目环境 ###

![](https://user-gold-cdn.xitu.io/2019/6/5/16b26f043d671082?imageView2/0/w/1280/h/960/ignore-error/1)

### ###

### 资产管理是项目的一个核心能力，所有持续集成都依赖于资产，也是Web Terminal的入口之一 ###

![](https://user-gold-cdn.xitu.io/2019/6/5/16b26f033b7080e4?imageView2/0/w/1280/h/960/ignore-error/1)

### ###

### 你可以新建任何需要发布和管理的应用，分配相应的资产，可以选择单机部署、部分部署或者全量部署 ###

![](https://user-gold-cdn.xitu.io/2019/6/5/16b26f033c80e96b?imageView2/0/w/1280/h/960/ignore-error/1)

### ###

### 在这里你可以查看任何时候的应用构建状态、备份状态、生成配置状态 ###

![](https://user-gold-cdn.xitu.io/2019/6/5/16b26f04bfd23215?imageView2/0/w/1280/h/960/ignore-error/1)

### ###

### 一切的部署都依赖于shell脚本，脚本的正确与否，决定了你的应用是否能部署成功 ###

![](https://user-gold-cdn.xitu.io/2019/6/5/16b26f0336ac492a?imageView2/0/w/1280/h/960/ignore-error/1)

### 友好的web化界面部署日志，支持多机，你可以随时掌控部署状态，也可随时终端某台机器的发布 ###

![](https://user-gold-cdn.xitu.io/2019/6/5/16b26f03f7c76c3e?imageView2/0/w/1280/h/960/ignore-error/1)

### 强大的Web Terminal能力，跟xshell工具一样的体验，随时随地管理你的资产吧 ###

![](https://user-gold-cdn.xitu.io/2019/6/5/16b26f03f7d96915?imageView2/0/w/1280/h/960/ignore-error/1)

## 感兴趣 ##

如果你有那么一点感兴趣，别犹豫先star或者watch，我会持续的更新和迭代，让它成为你开发中的神器吧

github address: [github.com/wangweiange…]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fwangweianger%2FAPubPlat )

如果你也认可我，那也可以给我一个following额

你还可以加入QQ群来尽情的交流吧，一款好的产品更需要碰撞和火花。

![](https://user-gold-cdn.xitu.io/2019/6/5/16b26f03fb187d4a?imageView2/0/w/1280/h/960/ignore-error/1)