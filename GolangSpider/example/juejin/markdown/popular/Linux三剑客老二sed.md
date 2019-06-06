# Linux三剑客老二sed #

**“** 我才不要手动改配置。——编程三分钟 **”**

### 概述 ###

sed命令是用来批量修改文本内容的，比如批量替换配置中的某个ip。
sed命令在处理时，会先读取一行，把当前处理的行存储在临时缓冲区中，处理完缓冲区中的内容后，打印到屏幕上。然后再读入下一行，执行下一个循环。不断的重复，直到文件末尾。
语法：

` sed [参数] [文本或文件] 复制代码`

由于不加 ` -i` 参数只会输出到控制台不会写入到文件中，所以以下例子默认加 ` -i`

### 插入 ###

* 在某行前面插入一行
` $ sed -i "1a insert after" file.txt $ cat file.txt 1 insert after 2 3 复制代码`

其中 ` 1a` 表示在第1行后（after）插入

* 在某行后面插入一行
` $ sed -i "1i insert before" file.txt $ cat file.txt insert before 1 2 3 复制代码`

其中 ` 1i` 表示在第1行前插入

### 删除 ###

` $ sed -i '2,3d' file.txt $ cat file.txt 1 复制代码`

删除行可以删除一行 （ ` 3d` 删除第三行），也可以写一个范围（ ` 2,3d` 删除2-3行，闭区间）， ` $` 符号代表末尾
缺点是只能多次连续删除行，不能一次性删除匹配到的行，可以用正则删除（ ` /^2/d` 代表删除所有内容以2开头的行）

### 替换行 ###

` $ sed -i '2c replace' file.txt $ cat file.txt 1 replace 3 复制代码`

` 2c replace` 表示替换第2行的内容为 ` replace`
缺点是只能多次替换行，不能一次性替换全部匹配到的行，可以用正则替换（ ` /^2/c replace` 代表替换所有以2开头的行为 ` replace` ）

### 仅替换匹配的字符串 ###

为了便于演示修改文件内容为

` $ cat -n config.txt 1 name=coding3min 2 age=0 3 email=coding3min@foxmail.com 4 name=coding3min 5 age=0 6 email=coding3min@foxmail.com 复制代码`

使用命令批量替换 ` 3-4` 行之间 ` coding3min` 字符串为 ` tom`

` $ sed -i '3,4s/coding3min/tom/g' config.txt $ config.txt name=coding3min age=0 email=tom@foxmail.com name=tom age=0 email=coding3min@foxmail.com 复制代码`

` s/coding3min/top/g` 代表全文匹配不限制行，去掉 ` g` 代表只替换匹配到的第一个如 ` s/coding3min/top`

### 查找与输出 ###

输出3-4行的内容

` sed -n 3,4p config.txt email=coding3min@foxmail.com name=coding3min 复制代码`

查找所有以name开头的行

` sed -n '/^name/p' config.txt name=coding3min name=coding3min 复制代码`

可以看到只要用 ` -n` 参数+匹配p模式就可以sj查找并输出

### 自动创建备份文件 ###

当然了，直接 ` sed -i` 很容易造成替换错误，哭都没办法哭！所以需要事先用 ` -n+p` 也就是上一节说的方法先校验下结果。但是每个都校验显然是不实际的。所以可以用 ` sed -i备份文件后缀的方式` 例如 ` sed -i.bak` 或者 ` sed -i.backup`

` $ sed -i.bak 's/coding3min/kitty/g' config.txt $ ls config.txt config.txt.bak $ cat config.txt name=kitty age=0 email=kitty@foxmail.com $ cat config.txt.bak name=coding3min age=0 email=coding3min@foxmail.com 复制代码`

### 与grep的结合使用 ###

与 ` grep` 结合使用最爽的点就在可以提前校验和批量替换，提高容错率和效率，不会的赶紧Get了

` sed -i 's/coding/kitty/g' `grep -rl coding *` $cat config.txt name=conding3min age=0 email=conding3min@foxmail.com $cat test /config.txt name=conding3min age=0 email=conding3min@foxmail.com 复制代码`

看明白了吗？上一节说的 ` grep -rl` 递归找到匹配的文件，并把文件名输出，前后加上了 ` 反引号，就是键盘左上角数字1左边那个符号，代码提前执行。
然后再使用替换文件内容。

### 其他技巧 ###

使用sed把DOS格式的文件转换为Unix格式 ` sed 's/.$//' filename`

匹配所有包含邮箱的行,( ` -n` 选项让sed仅仅是输出经过处理之后的那些行)

` sed -n '/[A-Za-z0-9]\+\@[a-zA-Z0-9_-]\+\(\.[a-zA-Z0-9_-]\+\)/p' config.txt email=coding3min@foxmail.com email=coding3min@foxmail.com 复制代码`

去掉所有的html标签

` $ cat html.txt <b>hi!</b><span>I 'm</span> $ sed ' s/<[^>]*>//g ' html.txt hi!I' m father 复制代码`

### 推荐阅读 ###

（点击标题可跳转阅读）

[linux三剑客之老三grep]( https://link.juejin.im?target=https%3A%2F%2Flink.zhihu.com%2F%3Ftarget%3Dhttp%253A%2F%2Fmp.weixin.qq.com%2Fs%253F__biz%253DMzAxOTc1OTY4NA%253D%253D%2526mid%253D2650855336%2526idx%253D1%2526sn%253Dd19b4edc9359d1bcfa3bca0ead2fbd7b%2526chksm%253D80366083b741e9951040eb26c78acd588ac30eaf5dce74c369ff42b0fac2637676039e84973f%2526scene%253D21%2523wechat_redirect )

[我的服务器怎么老这么慢，难道说是被挖矿了？linux开机启动项自查]( https://link.juejin.im?target=https%3A%2F%2Flink.zhihu.com%2F%3Ftarget%3Dhttp%253A%2F%2Fmp.weixin.qq.com%2Fs%253F__biz%253DMzAxOTc1OTY4NA%253D%253D%2526mid%253D2650855164%2526idx%253D1%2526sn%253D155cd030a0dbfd541564a72a08f4ca1a%2526chksm%253D803661d7b741e8c1939b38064f9043d956e19105876c27e36c492859e8f309ce4ef0f578d511%2526scene%253D21%2523wechat_redirect )

[我偷偷挖了一条网络隧道，差点被公司激活]( https://link.juejin.im?target=https%3A%2F%2Flink.zhihu.com%2F%3Ftarget%3Dhttp%253A%2F%2Fmp.weixin.qq.com%2Fs%253F__biz%253DMzAxOTc1OTY4NA%253D%253D%2526mid%253D2650855243%2526idx%253D1%2526sn%253D5d00658368439b602f2f815269af8145%2526chksm%253D80366060b741e976a0966c541ae0dd9beb462d6be85d475e23b1e4ffbaee4a8c3f4e72358f60%2526scene%253D21%2523wechat_redirect )

![](https://user-gold-cdn.xitu.io/2019/6/1/16b11e67f460a47e?imageView2/0/w/1280/h/960/ignore-error/1)

**如果有帮助别忘了分享给朋友哦~**