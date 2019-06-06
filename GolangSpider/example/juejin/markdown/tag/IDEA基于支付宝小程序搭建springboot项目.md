# IDEA基于支付宝小程序搭建springboot项目 #

# 服务端 #

## 在平台上创建springboot小程序应用 ##

* 

### 创建小程序 ###

* 登录 [蚂蚁金服开放平台]( https://link.juejin.im?target=https%3A%2F%2Fopenhome.alipay.com%2Fplatform%2Fhome.htm ) ，扫码登录填写信息后，点击 [支付宝小程序]( https://link.juejin.im?target=https%3A%2F%2Fmini.open.alipay.com%2Fchannel%2FminiIndex.htm ) ，选择立即接入 > 同意个人公测 > 开始创建 。
* 填写好小程序基本信息后，点击创建按钮，创建名为xxx小程序。
* PS：一个账号下最多可以创建10个小程序；未提交过审核的小程序可以删除，删除的小程序不在计数范围。

* 

### 创建云应用后端服务 ###

* 在小程序页面选择刚创建的小程序，点击查看，进入开发者页面。
* 在左侧导航栏选择云服务(免费)，点击创建云服务，选择创建云应用，技术栈选SpringBoot，填写好应用名，描述后即可完成创建云应用。

* 

### 构建环境 ###

* 返回 云服务(公测) 页面，点击刚创建的云服务卡片中的 构建环境 按钮
* 在 购买环境资源 页面，选择合适的环境配置方案，点击 同意《产品服务协议》 > 确认配置。

#### 说明： 此处选择小程序云应用入门（Mysql版），当前测试环境该方案免费提供，但若连续 7 日未部署过代码，环境会被自动回收。 ####

* 在 确认订单 页面，点击 确认购买。 购买成功后会自动进入 构建环境 页面。构建过程会耗时几分钟，构建成功后，您可以选择 查看应用详情 ，或者 返回应用列表。

## 在IDEA中安装支付宝小程序开发工具 ##

* 点击 [这里]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Falipay%2Falipay-intellij-plugin%2Freleases ) 前往下载最新的插件版本(com.alipay.devtools.idea-1.0.6.zip )进行下载。
* 下载完成后，IDE 中 Preferences (Windows 下为 Settings) => Plugins => Install plugin from disk…，点选已下载的 Zip 包进行安装，依照提示重启 IDE 生效。
* 中文乱码解决方法： Appearance & Behavior => Appearance => UI Options -> Name 里设置成中文字体，如 微软雅黑（microsoft yahei light）、文泉驿（linux）

## 创建项目 ##

* file > new > project > Alipay DevTools，选择springboot > next > finish;
* 创建完后的效果： ![](https://user-gold-cdn.xitu.io/2019/4/28/16a6269379fd0637?imageView2/0/w/1280/h/960/ignore-error/1)
* HelloController的代码: ![](https://user-gold-cdn.xitu.io/2019/4/28/16a62a8691f0a937?imageView2/0/w/1280/h/960/ignore-error/1)

## 本地项目与平台的云应用对接 ##

* 在IDEA的小程序云应用视图中点击登录账号然后用具有开发者权限的用户扫码授权登录
PS：添加开发者账号的方式：我的小程序 =>查看 => 成员管理 => 添加 => 要添加的账号在客户端找到对应的提示信息并点击确认
* 开始对接 ![](https://user-gold-cdn.xitu.io/2019/4/28/16a627393a49204b?imageView2/0/w/1280/h/960/ignore-error/1)
* 弹出扫码框，请使用具有开发者权限的用户支付宝扫码登录。
* 点击下方底栏的alipay devtools，先于自己的小程序关联。 ![](https://user-gold-cdn.xitu.io/2019/4/28/16a627fe37350d42?imageView2/0/w/1280/h/960/ignore-error/1)
* 选中自己的小程序后，出现如下即关联小程序成功。 ![](https://user-gold-cdn.xitu.io/2019/4/28/16a628677805ac7b?imageView2/0/w/1280/h/960/ignore-error/1)
* 选中关联云应用 ![](https://user-gold-cdn.xitu.io/2019/4/28/16a627801d52f1d5?imageView2/0/w/1280/h/960/ignore-error/1)
* 弹出框选中之前在平台创建的云应用名称 ![](https://user-gold-cdn.xitu.io/2019/4/28/16a6279a5a6f72f0?imageView2/0/w/1280/h/960/ignore-error/1)
* 在下面alipay devtools中选择云端部署，出现如下即关联成功 ![](https://user-gold-cdn.xitu.io/2019/4/28/16a628808bed8075?imageView2/0/w/1280/h/960/ignore-error/1)
* 开始部署 ![](https://user-gold-cdn.xitu.io/2019/4/28/16a6289fade6104a?imageView2/0/w/1280/h/960/ignore-error/1) ![](https://user-gold-cdn.xitu.io/2019/4/28/16a628c7ec6224bb?imageView2/0/w/1280/h/960/ignore-error/1) ![](https://user-gold-cdn.xitu.io/2019/4/28/16a628de346af53f?imageView2/0/w/1280/h/960/ignore-error/1)
* 开始部署后，云应用管理 视窗会打出部署日志。部署结束后会有消息提示部署完成。

# 客户端 #

* 下载 [小程序开发者工具]( https://link.juejin.im?target=https%3A%2F%2Fdocs.alipay.com%2Fmini%2Fdeveloper%2Ftodo-demo ) 并安装 ![](https://user-gold-cdn.xitu.io/2019/4/28/16a62986c44012a7?imageView2/0/w/1280/h/960/ignore-error/1) ![](https://user-gold-cdn.xitu.io/2019/4/28/16a629a123be814a?imageView2/0/w/1280/h/960/ignore-error/1)
* 创建完成后编辑page/index/index.js ![](https://user-gold-cdn.xitu.io/2019/4/28/16a62a0b68011c8c?imageView2/0/w/1280/h/960/ignore-error/1)
* 启动模拟器测试如下： ![](https://user-gold-cdn.xitu.io/2019/4/28/16a62a310b490cba?imageView2/0/w/1280/h/960/ignore-error/1)
* 打开控制台如下： ![](https://user-gold-cdn.xitu.io/2019/4/28/16a62a446574f523?imageView2/0/w/1280/h/960/ignore-error/1)

### 至此IDEA支付宝小程序搭建springboot项目完成，后续有时间会推出支付宝授权，支付等文章，感谢大家的观看，如果错误，请在评论区指出来，我会校正，蟹蟹各位。 ###