# Go借助PProf的一次性能优化 #

> 
> 
> 
> 本文以leetcode的一题为例来讲解如何通过PProf来优化我们的程序，题目如下： [Longest Substring Without
> Repeating Characters](
> https://link.juejin.im?target=https%3A%2F%2Fleetcode.com%2Fproblems%2Flongest-substring-without-repeating-characters%2F
> )
> 
> 

### 首先给出我们一般的解法 ###

` func lengthOfNonRepeatingSubStr(s string) int { lastOccurred := make(map[rune]int) start := 0 maxLength := 0 for i, ch := range []rune(s) { if lastI, ok := lastOccurred[ch]; ok && lastI >= start { start = lastI + 1 } if i-start+1 > maxLength { maxLength = i - start + 1 } lastOccurred[ch] = i } return maxLength } 复制代码`

性能检测可知 200次， 花费时间 6970790 ns/op

![](https://user-gold-cdn.xitu.io/2019/3/25/169b3aa646076c08?imageView2/0/w/1280/h/960/ignore-error/1)

### PProf分析 ###

` go test -bench . -cpuprofile cpu.out //首先生成cpuprofile文件 go tool pprof -http=:8080 ./启动 PProf 可视化界面 分析cpu.out文件 复制代码`

通过访问 [http://localhost:8080/ui/]( https://link.juejin.im?target=http%3A%2F%2Flocalhost%3A8080%2Fui%2F ) 可以查看到如下页面

![](https://user-gold-cdn.xitu.io/2019/3/25/169b3b6af8b9a9ad?imageView2/0/w/1280/h/960/ignore-error/1)

可以看到主要消耗时间在2大块，一个是mapaccess,mapassign,还有一块是decoderune。 decoderune主要是对UFT8字符的解码，将字符串转换成 ` []rune(s)` 这个是不能避免的。所以主要去解决map的访问和赋值问题，也就是代码中的 ` lastOccurred`

### 优化分析 ###

由于map需要进行算hash 判重，分配空间等操作会导致操作慢下来，解决思路就是用空间换时间,通过slice来替换map.

修改后的代码如下：

` func lengthOfNonRepeatingSubStr2(s string) int { lastOccurred := make([]int, 0xffff) //赋给一个初始值 for i := range lastOccurred { lastOccurred[i] = -1 } start := 0 maxLength := 0 for i, ch := range []rune(s) { if lastI := lastOccurred[ch]; lastI != -1 && lastI >= start { start = lastI + 1 } if i-start+1 > maxLength { maxLength = i - start + 1 } lastOccurred[ch] = i } return maxLength } 复制代码`

性能检测可知 500， 花费时间 2578859 ns/op。 相比之前的6970790 ns/op 已经又很大的优化了

![](https://user-gold-cdn.xitu.io/2019/3/25/169b3cc0f8044b00?imageView2/0/w/1280/h/960/ignore-error/1)

### 后续优化 ###

通过pprof查看除了decoderune外还有一部分时间花费在makeslice上面,这是由于每次调用函数都要makeslicke，可以slice移到函数外面进行声明。具体可以自己操作下，然后查看下pprof图上面的makeslice是否有消除。

![](https://user-gold-cdn.xitu.io/2019/3/25/169b3cf53eff0d7d?imageView2/0/w/1280/h/960/ignore-error/1)

> 
> 
> 
> PProf具使用可查看 [Golang 大杀器之性能剖析 PProf](
> https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2FEDDYCJY%2Fblog%2Fblob%2F7f17a29ebb6d2a506601f0e9cc3275ace0aa28cb%2Fgolang%2F2018-09-15-Golang%2520%25E5%25A4%25A7%25E6%259D%2580%25E5%2599%25A8%25E4%25B9%258B%25E6%2580%25A7%25E8%2583%25BD%25E5%2589%2596%25E6%259E%2590%2520PProf.md
> )
> 
> 

本文亦在微信公众号【小道资讯】发布，欢迎扫码关注！

![](https://user-gold-cdn.xitu.io/2019/3/13/1697573b580effc0?imageView2/0/w/1280/h/960/ignore-error/1)