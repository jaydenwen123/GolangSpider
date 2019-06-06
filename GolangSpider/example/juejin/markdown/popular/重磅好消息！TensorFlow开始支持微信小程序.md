# 重磅好消息！TensorFlow开始支持微信小程序 #

在昨天的推送《 [一文带你众览Google I/O 2019上的人工智能主题演讲]( https://link.juejin.im?target=https%3A%2F%2Fmp.weixin.qq.com%2Fs%3F__biz%3DMzI3NTQyMzEzNQ%3D%3D%26amp%3Bmid%3D2247485232%26amp%3Bidx%3D1%26amp%3Bsn%3De6ac6fc029e16e6341ffb9d316116d23%26amp%3Bchksm%3Deb044dc0dc73c4d638261fa9b70dffea25eabec6b78407aa2869eb59852bbdaa17fd78bec577%26amp%3Btoken%3D1176368429%26amp%3Blang%3Dzh_CN%23rd ) 》中，回顾了Google I/0 2019大会上的TensorFlow专题演讲，不知道朋友有没有注意到在TensorFlow.js介绍部分，重点提到了 **TensorFlow.js开始支持微信小程序** 。今天我将这部分的视频截取出来，请大家观看：

++此处应有视频，请前往公众号观看。++

视频没中文字幕，不过大致可以看懂，这是一个通过头部姿势控制吃豆人的小游戏。这个小游戏最初是作为web小游戏出现在TensorFlow.js的官方示例程序中，源代码位于github： [github.com/tensorflow/…]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Ftensorflow%2Ftfjs-examples%2Ftree%2Fmaster%2Fwebcam-transfer-learning ) 。这次则作为微信小游戏进行演示。我上微信上搜索这款小游戏，但没有搜到，可能是 **没有公开发布，源代码也未提供** ，想必将web版的源码，移植到微信小程序，难度不会太大吧。

接着我去翻看了tfjs-core的提交记录，看到有如下一条提交：

` commit c211b496a5ee7f88f7bf4ab21a2bc5054f485175 Author: Ping Yu <4018+pyu10055@users.noreply.github.com> Date: Tue Jan 29 07:40:48 2019 -0800 Support WeChat mini app environment ( #1510) To compensate the differences between browser and WeChat mini app: - WeChat mini app runs on JS core (ios) which does not have document, window, and set Immediate function or objects. - When creating a GPGUContext with a existing context, it needs to store the context for the GL version, otherwise it would be picked later. This PR also fix the inconsistency issue with GPGPUContext constructor, it should always cache the rendering context. 复制代码`

可以确定主干分支上的tfjs已经支持微信小程序了，但最新的稳定分支1.1.2是否支持，还无法确定，大家可以尝试一下。看提交，应该是在今年年初，不知道为啥在官方文档上没有提及，也没有媒体进行报道。

之前开发过一款人工智能微信小程序： **识狗君** 。采用的是小程序+TensorFlow Serving的架构，虽然说现在手机联网基本上不成问题，但是服务器部署对于个人开发者还是一件麻烦事，如果能够在手机端完成推理，开发工作可以减少很多。后面有时间我会将识狗君微信小程序用TensorFlow.js改写。

你会在微信小程序中采用TensorFlow.js吗？欢迎大家一起交流！

![](https://user-gold-cdn.xitu.io/2019/6/3/16b1d13ac18d3336?imageView2/0/w/1280/h/960/ignore-error/1)