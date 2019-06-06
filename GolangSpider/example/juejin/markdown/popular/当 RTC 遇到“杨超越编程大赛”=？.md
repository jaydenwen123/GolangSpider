# 当 RTC 遇到“杨超越编程大赛”=？ #

## 引言 ##

前些日子，我司做了个小的内部分享活动，无心插柳，没想到还有个意外收获。

有位小哥哥，参加了前两个月的 **“杨超越编程大赛”** ，提交了一份作品，作品主题思路清奇，还很有诚意地花了650元“巨款”，租了个服务器。

虽然最终获没获奖，我们也不知道，我们也不敢问。但就冲着他这份“诚意”，我们也想把他的作品分享出来，供大家娱乐娱乐。

本着“在技术社区要分享技术”的原则，我们让他将开发过程也写了下来。

接下来，请看他的“表演”。

![](https://user-gold-cdn.xitu.io/2019/6/5/16b263a3d7ddd298?imageslim)

朋友们可以直接去 [ycy.dev]( https://link.juejin.im?target=https%3A%2F%2Fycy.dev ) 体验。

体验后记得去我的B站视频 [www.bilibili.com/video/av493…]( https://link.juejin.im?target=https%3A%2F%2Fwww.bilibili.com%2Fvideo%2Fav49364693 ) 下素质三连哦！

github： [github.com/scaret/ycy-…]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fscaret%2Fycy-wishing-machine )

## 应用介绍 ##

> 
> 
> 
> 灵感来自于超越的幸运A体质。通过向超越许愿的形式来分享超越的幸运值。
> 该项目通过网络摄像头直播一个放在PO主家的打印机。你可以通过H5页面向打印机发送许愿内容。许愿内容会通过网页直播实时传回，并被全世界看到。许愿纸攒到一定程度会寄给超越。
> 
> 
> 

> 
> 
> 
> 另外，最近买了个小爱同学，打算也放入摄像头范围，加上音视频互通能力，就可以通过实时音视频远程操控小爱同学啦。
> 
> 

## 线上环境 ##

因为是个人项目，也不好意思用公司机器，于是蹭了AWS Free Tier，使用了AWS日本的Ubuntu 16.04，micro类型的实例。

![](https://user-gold-cdn.xitu.io/2019/6/5/16b26421db0a01ed?imageView2/0/w/1280/h/960/ignore-error/1)

基本上服务器只是起到托管页面的作用，由于嫌备案麻烦而放在日本，还比较慢。其他设置如下：

* nginx/1.10.3
* node.js/10.15.3
* pm2 3.5.0
* express/4.16.4

域名是花650块重金从Google买的。证书是免费的LetEncrypt证书。咕咕机是去年买的在家吃灰了一年翻出来的。

![](https://user-gold-cdn.xitu.io/2019/6/5/16b264268ff6c788?imageView2/0/w/1280/h/960/ignore-error/1)

实时音视频技术用的当然是Agora WebRTC SDK啦。主播端和观众端都使用的是Web。

## WebRTC主播端 ##

主播端比较简单，就是一个发送端网页。获取摄像头权限后发动到指定频道即可。

为了兼容不同设备的观众端的Codec，所以主播端发送H264、VP8两路视频。两路视频都是720P的设置。

` var codecs = [ "h264" , "vp8" ]; codecs.forEach( function (codec){ var client = AgoraRTC.createClient({mode: "rtc" , codec: codec}); client.init( "<appId>" , function () { client.join(null, cname + "_" + codec, 8888, function (){ console.log( "Client joined" ); const spec = {video: true , audio: true }; const local Stream = AgoraRTC.createStream(spec); local Stream.setVideoProfile( "720p_2" ); console.log( "spec" , spec); local Stream.init( function (){ window.vt = local Stream.stream.getVideoTracks()[0] console.log( "LocalStream Inited" ); client.publish( local Stream); client.on( "stream-published" , function (){ console.log( "stream published" ); local Stream.play( "local-container" ); }); }); }); }); }); 复制代码`

之后我额外做了混音功能，在频道内不间断地播放【卡路里之歌】。

` localStream.audioMixing.inEarMonitoring = "NONE" ; if (codec === "vp8" ){ localStream.startAudioMixing({ filePath : '/music/kll.mp3' , replace : true , playTime : 0 , loop : true }); } else { localStream.startAudioMixing({ filePath : '/music/pickme.mp4' , replace : true , playTime : 0 , loop : true }); } 复制代码`

## WebRTC观众端 ##

观众端首先会侦测当前环境支持的视频编码格式。 由于安卓设备的H264实现比较不一致，在这里是VP8优先，在不支持VP8的环境中Fallback到H264视频编码。

` var cname = "<cname>" ; var codec = "vp8" ; AgoraRTC.getSupportedCodec().then( function (codecs){ console.log( "codecs" , codecs); if (codecs.video.indexOf( "VP8" !== -1)){ codec = "vp8" ; } else { codec = "h264" ; } }).catch( function (e){ console.error(e); codec = "h264" ; }); 复制代码`

在Chrome、Safari浏览器中，有一个叫做【自动播放策略】的东西。简单地说，浏览器会阻止媒体自动播放声音。 解决的方法也很简单，在播放之前额外设置了一个确认框，引导用于点击确认时，允许视频的播放。

` var nickname = window.localStorage && local Storage.getItem( "nickname" ); function updateNickname (){ bootbox.prompt({ closeButton: false , title: "你的名字是？" , callback: function (result){ if (result){ nickname = result; window.localStorage && local Storage.setItem( "nickname" , nickname); ConnectIO(); } else { updateNickname(); } elem.play(); } }); } 复制代码`

iOS的微信也支持WebRTC了！真是普天同庆。

## 其他的一些零碎 ##

还有一些零零碎碎的工作，包括：

* 分享页面的优化
* 在线人数统计
* 连接咕咕机打印
* 打印太多导致纸张翻车，需要时常捋一捋

相对来说这个应用还是蛮简单的。

希望大家给我点赞！

![](https://user-gold-cdn.xitu.io/2019/6/5/16b2642d827d6d8b?imageView2/0/w/1280/h/960/ignore-error/1)

![](https://user-gold-cdn.xitu.io/2019/6/5/16b2643f481360a2?imageslim)