# 【有趣】这段java代码太古怪 #

首先呢，来一段java代码来开点胃。等等等等，耍我呢，这是java代码？

` \u0070\u0075\u0062\u006c\u0069\u0063\u0020\u0063\u006c\u0061\u0073\u0073\u0020\u0058\u004a\u004a\u0020\u007b \u0020\u0020\u0020\u0020\u0070\u0075\u0062\u006c\u0069\u0063\u0020\u0073\u0074\u0061\u0074\u0069\u0063\u0020\u0076\u006f\u0069\u0064\u0020\u006d\u0061\u0069\u006e\u0028\u0053\u0074\u0072\u0069\u006e\u0067\u005b\u005d\u0020\u0061\u0072\u0067\u0073\u0029\u0020\u007b \u0020\u0020\u0020\u0020\u0020\u0020\u0020\u0020\u0053\u0079\u0073\u0074\u0065\u006d\u002e\u006f\u0075\u0074\u002e\u0070\u0072\u0069\u006e\u0074\u006c\u006e\u0028\u0022\u5c0f\u59d0\u59d0\u6211\u7231\u4f60\u0022\u0029\u003b \u0020\u0020\u0020\u0020\u007d \u007d 复制代码`

非常负责任的告诉你，是的！不信请看下图。纯纯正正的java代码，class为XJJ的java源码，执行后打印 ` 小姐姐我爱你` 。

![](https://user-gold-cdn.xitu.io/2019/4/1/169d7ad30f1de1c2?imageView2/0/w/1280/h/960/ignore-error/1)

还是不信？自个儿拷贝下去执行一下。不过，IDEA是会报错的，用命令行哦。

**好隐晦的表白方式，是暗恋么？**

其实没什么神奇的，我们不过是将正常的源代码翻译成了unicode编码方式。就是这段java代码。

` private static String toUnicode(String str) { StringBuilder sb = new StringBuilder(); for (int i = 0; i < str.length(); i++) { if (str.charAt(i) != '\n' ) { int cp = Character.codePointAt(str, i); int charCount = Character.charCount(cp); if (charCount > 1) { i += charCount - 1; if (i >= str.length()) { throw new IllegalArgumentException( "truncated unexpectedly" ); } } sb.append(String.format( "\\u%04x" , cp)); } else { sb.append( "\n" ); } } return sb.toString(); } 复制代码`

耍到这里，我突然有了一个好主意。我要将我的java项目，全部编码成这种方式，然后传到github，嘿嘿。能编译但不可读，比base64更冷门。

所以以下几行python代码诞生了(仅用于python3)：

` #!/usr/bin/env python # -*- coding: utf-8 -*- import sys java = sys.argv[ 1 ] s = sb = u"" with open(java, 'r' , encoding= 'utf-8' ) as f: s = f.read() for _c in s: sb += '\\u%04x' % ord(_c) with open(java, 'w' , encoding= 'utf-8' ) as f: f.write(sb) print(java) 复制代码`

在命令行中执行以下命令，将会将指定目录(test)中的所有java文件翻译成我们所想要的。

` find ./ test | grep \\.java$ | xargs -I '{}' python3 uni.py {} 复制代码`

是不是很简单？

那改完的java文件怎么恢复呢？我只管编码不管解码，剩下的要靠自己啦，这可是了解unicode编码的好机会。

码农世界可能是太过寂寥，无聊的项目也是频出。比如这个，判断数字是不是13，竟然接近4k星了。 [github.com/jezen/is-th…]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fjezen%2Fis-thirteen )

贴上它的API感受下来自码农世界深深的空虚感吧。

` var is = require( 'is-thirteen' ); // Now with elegant syntax. is(13).thirteen(); // true is(12.8).roughly.thirteen(); // true is(6).within(10).of.thirteen(); // true is(2003).yearOfBirth(); // true // check your math skillz is(4).plus(5).thirteen(); // false is(12).plus(1).thirteen(); // true is(4).minus(12).thirteen(); // false is(14).minus(1).thirteen(); // true is(1).times(8).thirteen(); // false is(26).divideby(2).thirteen(); // true 复制代码`

![](https://user-gold-cdn.xitu.io/2019/4/1/169d7ad6cf3946c8?imageView2/0/w/1280/h/960/ignore-error/1)