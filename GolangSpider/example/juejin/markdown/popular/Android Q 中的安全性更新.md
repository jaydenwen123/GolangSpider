# Android Q 中的安全性更新 #

![](https://user-gold-cdn.xitu.io/2019/6/3/16b1b0d2764b63c1?imageView2/0/w/1280/h/960/ignore-error/1)

> 
> 
> 
> 作者: Rene Mayrhofer 和 Xiaowen Xin, Android 安全与隐私团队
> 
> 

每次发布 Android 的新版本，我们的首要任务之一就是提高平台的安全防护。在过去几年，安全方面的优化在整个生态圈都取得了喜人的成绩，2018 年亦是如此。

在 2018 年第四季度，接收安全更新的设备数量比去年同期增长了 84%。与此同时，在 2018 年全年，任何对 Android 平台造成威胁的重要安全漏洞在公开披露之前，团队均提供了相应的安全更新或缓解措施。另外，我们还发现安装 [潜在危险应用]( https://link.juejin.im?target=https%3A%2F%2Fdevelopers.google.cn%2Fandroid%2Fplay-protect%2Fpotentially-harmful-applications ) 的设备数量同比下降了 20%。本着透明公开的原则，除了以上数据，我们还在《 [Android 安全及隐私 2018 年度报告]( https://link.juejin.im?target=https%3A%2F%2Fsource.android.google.cn%2Fsecurity%2Freports%2FGoogle_Android_Security_2018_Report_Final.pdf ) 》公布了更多安全方面的细节与回顾，有兴趣的朋友可前往阅读。

不过，大家可能会问，那 Android 接下来又有什么计划呢？

在五月上旬举办的 [ Google I/O’19 ]( https://link.juejin.im?target=https%3A%2F%2Fmp.weixin.qq.com%2Fs%2FIV4beRNCuccS3PkDMe9Nkg ) 上，我们揭晓了Android 中新集成的所有安全特性。我们将在接下来的几周和数月内继续优化这些特性，不过我们想先在这篇文章中与您快速分享一下我们为平台做了哪些安全升级。

## 加密 ##

储存加密属于最基础 (也最有效) 的安全技术之一，不过当前的加密标准有一定的硬件要求，即设备需搭载加密加速硬件，因而许多硬件受限的设备便无法使用该技术。 [Adiamtum ]( https://link.juejin.im?target=https%3A%2F%2Fsecurity.googleblog.com%2F2019%2F02%2Fintroducing-adiantum-encryption-for.html ) 的推出改变了 Android Q 的加密方式。我们早在今年二月就推出了 Adiantum 加密模式，让所有 Android 设备——从智能手表到联网医疗器械——即便在缺少特定硬件的情况下依旧能够实现数据加密。

我们在 Android Q 中继续践行对加密重要性的承诺。所有出厂系统为 Android Q 的兼容设备都必须对用户数据进行加密处理，无一例外。这个要求的涵盖类型包括手机、平板、电视及车载设备。这有助于确保下一代设备比之前的设备更加安全，让亿万新用户从使用 Android 系统的第一天起就免受安全隐患的威胁。

不过，储存加密仅仅构成了我们安全版图的一部分，因此，我们还在 Android Q 中默认启用了 TLS 1.3 支持。TLS 1.3 是 TLS 标准的一次重要更新， IETF (互联网工程组) 于去年 8 月正式完成了 TLS 1.3 的升级工作。与之前几个版本相比，TLS 1.3 在速度、安全性和隐私性三方面均有显著提升。

TLS 1.3 一般通过几轮数据往返即可完成握手流程，将建立会话连接的速度加快了 40%。从安全角度来看，TLS 1.3 移除了对较弱加密算法以及一些不安全或过时特性的支持。TLS 1.3 使用了新设计的握手协议，该协议修复了 1.2 版本中一些不足的地方，更为清晰，也不容易出错，而且对密钥泄露的防御性也有所提高。从隐私角度来看，TLS 1.3 对握手内更多的数据进行了加密处理，从而更好地保护了参与方的身份。

## 强化平台 ##

Android 采用深度防御 (defense-in-depth) 策略，为的是确保实现层面的单个错误无法绕过整个安全系统。我们借助进程隔离、攻击面缩减、架构分离及漏洞利用缓解技术，让漏洞更难或根本无法利用，而且攻击者需要同时利用更多的漏洞才能达成他们的目标。

在 Android Q 中，我们将这些策略实践至多个关键安全领域的研发工作中，例如: 媒体、蓝牙以及系统内核。我们在《 [Android 平台安全增强项详览]( https://link.juejin.im?target=https%3A%2F%2Fsecurity.googleblog.com%2F2019%2F05%2Fqueue-hardening-enhancements.html ) 》一文中提供了详实的介绍，其中的部分更新重点包括:

* 供软件编码器使用的受限沙箱；
* 增加排错程序 (sanitizer) 在生产环境中的使用: 当某组件处理不受信任的内容时，排错程序可用于缓解该组件内所有类别的漏洞；
* Shadow Call Stack (影调用堆栈): 该调用堆栈提供了 backward-edge 控制流完整性 (Control Flow Integrity - CFI)， 并加强了基于 LLVM CFI 的 forward-edge 保护；
* 使用 XOM (仅可执行内存) 来保护地址空间布局随机化 (ASRL)，以防止信息泄露；
* 引入 Scudo 强化内存分配器，增加堆 (heap) 相关漏洞的利用难度。

## 身份验证 ##

Android Pie 引入了 BiometricPrompt API 协助应用通过生物识别技术进行用户身份验证，如面部识别、指纹识别及虹膜识别。该 API 自推出以来便深受欢迎，我们在许许多多应用上都看到了它的身影。随着 Android Q 的发布，我们更新了 BiometricPrompt 底层框架，增强了对面部识别和指纹识别的支持。此外，我们还对该 API 进行了扩展，增加了支持用例的数量，如隐式和显式验证。

在显式流程中，用户必须通过明确的操作，如触摸指纹传感器，才能完成后续的身份验证工作。如果用户使用面部或虹膜进行验证，那么他们需要再点击其他按钮才能继续。显式流程为默认验证流程，所有高价值事务 (如付款) 均需通过显式流程完成。

隐式流程则不要求用户进行额外操作。借助隐式流程，开发者可以为简单的可撤销型事务提供更加轻量和无缝的体验，例如登录和自动填充。

BiometricPrompt 另外还增加了一项十分实用的新功能——在触发 BiometricPrompt 之前，检查设备是否提供生物验证支持。如果应用想在登录界面或应用内设置菜单中显示诸如 “启用生物验证登录” 一类的信息，那么，这项新功能便尤为有用。为了提供支持，我们新添加了一个名为 [BiometricManager]( https://link.juejin.im?target=https%3A%2F%2Fdeveloper.android.google.cn%2Freference%2Fandroid%2Fhardware%2Fbiometrics%2FBiometricManager ) 的类。您可调用其中的 [ canAuthenticate()]( https://link.juejin.im?target=https%3A%2F%2Fdeveloper.android.google.cn%2Freference%2Fandroid%2Fhardware%2Fbiometrics%2FBiometricManager.html%23canAuthenticate() ) 方法，来判定设备是否支持生物验证，以及用户是否已经同意使用。

## 下一步 ##

在 Android Q 之后，我们计划为移动应用添加数字身份证件 (Electronic ID) 支持，从而允许用户把手机当做身份证件 (如驾驶证) 来使用。此类应用需要符合多项安全规定，而且持证用户设备上的客户端应用、读取/认证设备，以及发证机构用于颁发、更新及撤销证件的后台系统三者间的集成工作也很重要。

项目的顺利推进需要 ISO (国际标准化组织) 在加密及标准化方面的专业支持，目前该项目由 Android 安全及隐私团队负责牵头。我们将会为 Android 设备提供相关 API 和 HAL 参考实现，以确保平台为类似的安全及隐私敏感应用提供核心支持。我们会在第一时间发布有关电子身份证支持的最新消息，感兴趣的小伙伴们，千万不要错过！

> 
> 
> 
> 致谢: 感谢 Jeff Vander Stoep 和 Shawn Willden 对本文的贡献。
> 
> 

[点击这里]( https://link.juejin.im?target=http%3A%2F%2Fservices.google.cn%2Ffb%2Fforms%2Fyourquestions%2F ) **提交产品反馈建议**

![](https://user-gold-cdn.xitu.io/2019/6/3/16b1b121049bff79?imageView2/0/w/1280/h/960/ignore-error/1)