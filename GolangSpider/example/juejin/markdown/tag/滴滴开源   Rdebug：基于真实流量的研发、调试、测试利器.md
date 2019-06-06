# 滴滴开源 | Rdebug：基于真实流量的研发、调试、测试利器 #

出品 | 滴滴技术

![](https://user-gold-cdn.xitu.io/2019/5/6/16a8c5dcd3a88c50?imageView2/0/w/1280/h/960/ignore-error/1)

前言：近日，滴滴在 GitHub 上开源后端研发、调试、测试的实用工具 Rdebug，全称 Real Debugger，中文称作真 · Debugger 。使用真实的线上流量进行线下回放测试，提升研发效率、保障代码质量，进而减少事故。一起来具体了解吧。

### ▍背景 ###

随着微服务架构的普及和应用，一个复杂的单体服务通常会被拆分成多个小而美的微服务。在享受微服务带来便利的同时，也要接受因为微服务改造带来的问题：需要维护的服务数变多、服务之间 RPC 调用次数增加。

这就造成线下环境维护成本大大增加，其次线下环境涉及到的部门较多，维护一个长期稳定的线下环境也是一个挑战；业务快速发展、需求不断迭代，手写单测又因复杂的业务逻辑以及复杂的服务调用需要 mock 多个下游服务，导致手写和维护单测成本特别的高；手动构造数据，又不够全面真实。以上问题都严重影响 RD 的研发效率，并且增加线上产生事故的隐患。

![](https://user-gold-cdn.xitu.io/2019/5/6/16a8c5fd13ffe1b6?imageView2/0/w/1280/h/960/ignore-error/1)

RD 迫切需要一个只需在本地部署代码、不用搭建下游依赖、使用真实数据，进行快速开发、调试、测试的解决方案。Rdebug 基于流量录制、流量回放的思路，能够巧妙的实现上述方案。

### ******▍****** **宗旨** ###

提升研发效率、降低测试成本、缩短产品研发周期，保障代码质量、减少线上事故。

### ******▍****** **使用** ###

**全景图**

![](https://user-gold-cdn.xitu.io/2019/5/6/16a8c660843dd185?imageView2/0/w/1280/h/960/ignore-error/1)

### ▍全新的研发体验 ###

* 

只需部署模块代码，无需搭建下游服务；

* 

在 macOS 本地回放，开发、调试、测试无需登录远程服务器；

* 

流量录制支持常用协议，FastCGI、HTTP、Redis、Thrift、MySQL 等；

* 

回放速度快，单次回放秒级别。

### ******▍****** **录路径重定向** ###

为了方便 RD 在本地开发、测试，Rdebug 支持路径重定向。

当线上部署路径和本地代码路径不一致时，当代码中存在大量线上路径硬编码时，无需入侵式修改代码，只需要简单的配置即可实现路径重定向。

即代码可以存放在任何路径下回放。

### ******▍****** **时间偏移** ###

流量回放时会自动把时间偏移到流量录制的时间点。

在代码中获取时间时，会获得录制时间点之后的时间。所以，当业务接口对时间敏感时，也无需担心。

### ******▍****** **文件 Mock**
###

流量回放支持文件 Mock，指定文件路径和 Mock 的内容，即可快速实现文件 Mock。

结合录制上报功能，在线上上报配置读取，在线下使用文件Mock实现配置“重现”。

### ▍Elastic 搜索 ###

对存储在 Elastic 中的流量，支持 URI、输入输出关键词、下游调用等多维度搜索。

回放支持指定文件，也支持上述搜索回放，使用体验更佳。

### ▍Xdebug 调试 ###

最高效的功能是 Xdebug 联动调试，通过对代码设置断点即可使用线上流量进行调试。通过这种方式，可以用来研究代码、排查问题、查看下游接口响应格式及数据等，是一个开发调试利器。

![](https://user-gold-cdn.xitu.io/2019/5/6/16a8c67ee6033c08?imageView2/0/w/1280/h/960/ignore-error/1)

### ******▍****** **丰富的报告** ###

回放报告，汇总线上线下的输入、输出、结果对比，一目了然。

![](https://user-gold-cdn.xitu.io/2019/5/6/16a8c68b65770084?imageView2/0/w/1280/h/960/ignore-error/1)

下游调用报告，会列举出所有的下游调用，包括协议、请求内容、匹配上的响应以及相识度。通过不同的背景颜色，标记出完全匹配的流量、存在噪点的调用、缺失的调用、新增的调用等。

![](https://user-gold-cdn.xitu.io/2019/5/6/16a8c6a96e3efd34?imageView2/0/w/1280/h/960/ignore-error/1)

结合 Xdebug 生成覆盖率报告，能够清楚的看到哪些代码被执行、哪些代码未被执行以及接口的覆盖率情况。

![](https://user-gold-cdn.xitu.io/2019/5/6/16a8c6ff7fb43119?imageView2/0/w/1280/h/960/ignore-error/1)

![](https://user-gold-cdn.xitu.io/2019/5/6/16a8c707e459afcf?imageView2/0/w/1280/h/960/ignore-error/1)

有关安装、使用过程以及常见问题解答，请查看以下链接：

**GitHub：** [github.com/didi/rdebug]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fdidi%2Frdebug )

**Wiki：** [github.com/didi/rdebug…]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fdidi%2Frdebug%2Fwiki )

**Documentation：** [github.com/didi/rdebug…]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fdidi%2Frdebug%2Fblob%2Fmaster%2Fdoc%2FDocList.md )

同时欢迎加入 **「Rdebug 用户交流群」**

请在滴滴技术公众号后台回复 **「Rdebug」** 即可加入

### ********▍******** ****END**** ###

![](https://user-gold-cdn.xitu.io/2019/5/6/16a8c714bc5b2b95?imageView2/0/w/1280/h/960/ignore-error/1)

![](https://user-gold-cdn.xitu.io/2019/5/6/16a8c71a616da390?imageView2/0/w/1280/h/960/ignore-error/1)

![](https://user-gold-cdn.xitu.io/2019/5/6/16a8c7211fff1d47?imageView2/0/w/1280/h/960/ignore-error/1)

**扫码关注“滴滴技术”公众号，获取更多最新最热技术干货！定制礼品不定期奉上！**

![](https://user-gold-cdn.xitu.io/2019/5/6/16a8c72809ad842f?imageView2/0/w/1280/h/960/ignore-error/1)