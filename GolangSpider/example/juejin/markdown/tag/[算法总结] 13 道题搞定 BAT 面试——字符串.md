# [算法总结] 13 道题搞定 BAT 面试——字符串 #

> 
> 
> 
> 本文首发于我的个人博客： [尾尾部落](
> https://link.juejin.im?target=https%3A%2F%2Fwww.weiweiblog.cn%2F13string%2F
> )
> 
> 

### 1. KMP 算法 ###

谈到字符串问题，不得不提的就是 KMP 算法，它是用来解决字符串查找的问题，可以在一个字符串（S）中查找一个子串（W）出现的位置。KMP 算法把字符匹配的时间复杂度缩小到 O(m+n) ,而空间复杂度也只有O(m)。因为“暴力搜索”的方法会反复回溯主串，导致效率低下，而KMP算法可以利用已经部分匹配这个有效信息，保持主串上的指针不回溯，通过修改子串的指针，让模式串尽量地移动到有效的位置。

具体算法细节请参考：

* [字符串匹配的KMP算法]( https://link.juejin.im?target=http%3A%2F%2Fwww.ruanyifeng.com%2Fblog%2F2013%2F05%2FKnuth%25E2%2580%2593Morris%25E2%2580%2593Pratt_algorithm.html )
* [从头到尾彻底理解KMP]( https://link.juejin.im?target=https%3A%2F%2Fblog.csdn.net%2Fv_july_v%2Farticle%2Fdetails%2F7041827 )
* [如何更好的理解和掌握 KMP 算法?]( https://link.juejin.im?target=https%3A%2F%2Fwww.zhihu.com%2Fquestion%2F21923021 )
* [KMP 算法详细解析]( https://link.juejin.im?target=https%3A%2F%2Fblog.sengxian.com%2Falgorithms%2Fkmp )
* [图解 KMP 算法]( https://link.juejin.im?target=http%3A%2F%2Fblog.jobbole.com%2F76611%2F )
* [汪都能听懂的KMP字符串匹配算法【双语字幕】]( https://link.juejin.im?target=https%3A%2F%2Fwww.bilibili.com%2Fvideo%2Fav3246487%2F%3Ffrom%3Dsearch%26amp%3Bseid%3D17173603269940723925 )
* [KMP字符串匹配算法1]( https://link.juejin.im?target=https%3A%2F%2Fwww.bilibili.com%2Fvideo%2Fav11866460%3Ffrom%3Dsearch%26amp%3Bseid%3D12730654434238709250 )

### 1.1 BM 算法 ###

BM算法也是一种精确字符串匹配算法，它采用从右向左比较的方法，同时应用到了两种启发式规则，即坏字符规则 和好后缀规则 ，来决定向右跳跃的距离。基本思路就是从右往左进行字符匹配，遇到不匹配的字符后从坏字符表和好后缀表找一个最大的右移值，将模式串右移继续匹配。 [字符串匹配的KMP算法]( https://link.juejin.im?target=http%3A%2F%2Fwww.ruanyifeng.com%2Fblog%2F2013%2F05%2FKnuth%25E2%2580%2593Morris%25E2%2580%2593Pratt_algorithm.html )

### 2. 替换空格 ###

> 
> 
> 
> 剑指offer： [替换空格](
> https://link.juejin.im?target=https%3A%2F%2Fwww.weiweiblog.cn%2Freplacespace%2F
> ) 请实现一个函数，将一个字符串中的每个空格替换成“%20”。例如，当字符串为We Are
> Happy.则经过替换之后的字符串为We%20Are%20Happy。
> 
> 

` public class Solution { public String replaceSpace (StringBuffer str) { StringBuffer res = new StringBuffer(); int len = str.length() - 1 ; for ( int i = len; i >= 0 ; i--){ if (str.charAt(i) == ' ' ) res.append( "02%" ); else res.append(str.charAt(i)); } return res.reverse().toString(); } } 复制代码`

### 3. 最长公共前缀 ###

> 
> 
> 
> Leetcode: [最长公共前缀](
> https://link.juejin.im?target=https%3A%2F%2Fleetcode-cn.com%2Fproblems%2Flongest-common-prefix%2Fdescription%2F
> ) 编写一个函数来查找字符串数组中的最长公共前缀。如果不存在公共前缀，返回空字符串 ""。
> 
> 

![](https://user-gold-cdn.xitu.io/2018/9/5/165a8f5403979cd8?imageView2/0/w/1280/h/960/ignore-error/1) 首先对字符串数组进行排序，然后拿数组中的第一个和最后一个字符串进行比较，从第 0 位开始，如果相同，把它加入 res 中，不同则退出。最后返回 res

` class Solution { public String longestCommonPrefix (String[] strs) { if (strs == null || strs.length == 0 ) return "" ; Arrays.sort(strs); char [] first = strs[ 0 ].toCharArray(); char [] last = strs[strs.length - 1 ].toCharArray(); StringBuffer res = new StringBuffer(); int len = first.length < last.length ? first.length : last.length; int i = 0 ; while (i < len){ if (first[i] == last[i]){ res.append(first[i]); i++; } else break ; } return res.toString(); } } 复制代码`

### 4. 最长回文串 ###

> 
> 
> 
> LeetCode: [最长回文串](
> https://link.juejin.im?target=https%3A%2F%2Fleetcode-cn.com%2Fproblems%2Flongest-palindrome%2Fdescription%2F
> ) 给定一个包含大写字母和小写字母的字符串，找到通过这些字母构造成的最长的回文串。在构造过程中，请注意区分大小写。比如 "Aa"
> 不能当做一个回文字符串。
> 
> 

![](https://user-gold-cdn.xitu.io/2018/9/5/165a8f54037ced50?imageView2/0/w/1280/h/960/ignore-error/1) 统计字母出现的次数即可，双数才能构成回文。因为允许中间一个数单独出现，比如“abcba”，所以如果最后有字母落单，总长度可以加 1。

` class Solution { public int longestPalindrome (String s) { HashSet<Character> hs = new HashSet<>(); int len = s.length(); int count = 0 ; if (len == 0 ) return 0 ; for ( int i = 0 ; i<len; i++){ if (hs.contains(s.charAt(i))){ hs.remove(s.charAt(i)); count++; } else { hs.add(s.charAt(i)); } } return hs.isEmpty() ? count * 2 : count * 2 + 1 ; } } 复制代码`

### 4.1 验证回文串 ###

> 
> 
> 
> Leetcode: [验证回文串](
> https://link.juejin.im?target=https%3A%2F%2Fleetcode-cn.com%2Fproblems%2Fvalid-palindrome%2Fdescription%2F
> ) 给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。 说明：本题中，我们将空字符串定义为有效的回文串。
> 
> 

两个指针比较头尾。要注意只考虑字母和数字字符，可以忽略字母的大小写。

![](https://user-gold-cdn.xitu.io/2018/9/5/165a8f5413dd17a7?imageView2/0/w/1280/h/960/ignore-error/1)

` class Solution { public boolean isPalindrome (String s) { if (s.length() == 0 ) return true ; int l = 0 , r = s.length() - 1 ; while (l < r){ if (!Character.isLetterOrDigit(s.charAt(l))){ l++; } else if (!Character.isLetterOrDigit(s.charAt(r))){ r--; } else { if (Character.toLowerCase(s.charAt(l)) != Character.toLowerCase(s.charAt(r))) return false ; l++; r--; } } return true ; } } 复制代码`

### 4.2 最长回文子串 ###

> 
> 
> 
> LeetCode: [最长回文子串](
> https://link.juejin.im?target=https%3A%2F%2Fleetcode-cn.com%2Fproblems%2Flongest-palindromic-substring%2Fdescription%2F
> ) 给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为1000。
> 
> 

![](https://user-gold-cdn.xitu.io/2018/9/5/165a8f5403a10d31?imageView2/0/w/1280/h/960/ignore-error/1) 以某个元素为中心，分别计算偶数长度的回文最大长度和奇数长度的回文最大长度。

` class Solution { private int index, len; public String longestPalindrome (String s) { if (s.length() < 2 ) return s; for ( int i = 0 ; i < s.length()- 1 ; i++){ PalindromeHelper(s, i, i); PalindromeHelper(s, i, i+ 1 ); } return s.substring(index, index+len); } public void PalindromeHelper (String s, int l, int r) { while (l >= 0 && r < s.length() && s.charAt(l) == s.charAt(r)){ l--; r++; } if (len < r - l - 1 ){ index = l + 1 ; len = r - l - 1 ; } } } 复制代码`

### 4.3 最长回文子序列 ###

> 
> 
> 
> LeetCode: [最长回文子序列](
> https://link.juejin.im?target=https%3A%2F%2Fleetcode-cn.com%2Fproblems%2Flongest-palindromic-subsequence%2Fdescription%2F
> ) 给定一个字符串s，找到其中最长的回文子序列。可以假设s的最大长度为1000。 **最长回文子序列和上一题最长回文子串的区别是，子串是字符串中连续的一个序列，而子序列是字符串中保持相对位置的字符序列，例如，"bbbb"可以使字符串"bbbab"的子序列但不是子串。**
> 
> 
> 

![](https://user-gold-cdn.xitu.io/2018/9/5/165a8f5413b37612?imageView2/0/w/1280/h/960/ignore-error/1) 动态规划： dp[i][j] = dp[i+1][j-1] + 2 if s.charAt(i) == s.charAt(j) otherwise, dp[i][j] = Math.max(dp[i+1][j], dp[i][j-1])

` class Solution { public int longestPalindromeSubseq (String s) { int len = s.length(); int [][] dp = new int [len][len]; for ( int i = len - 1 ; i>= 0 ; i--){ dp[i][i] = 1 ; for ( int j = i+ 1 ; j < len; j++){ if (s.charAt(i) == s.charAt(j)) dp[i][j] = dp[i+ 1 ][j- 1 ] + 2 ; else dp[i][j] = Math.max(dp[i+ 1 ][j], dp[i][j- 1 ]); } } return dp[ 0 ][len- 1 ]; } } 复制代码`

### 5. 字符串的排列 ###

> 
> 
> 
> Leetcode: [字符串的排列](
> https://link.juejin.im?target=https%3A%2F%2Fleetcode-cn.com%2Fproblems%2Fpermutation-in-string%2Fdescription%2F
> ) 给定两个字符串 s1 和 s2，写一个函数来判断 s2 是否包含 s1 的排列。 换句话说，第一个字符串的排列之一是第二个字符串的子串。
> 
> 

![](https://user-gold-cdn.xitu.io/2018/9/5/165a8f54149700d3?imageView2/0/w/1280/h/960/ignore-error/1) 我们不用真的去算出s1的全排列，只要统计字符出现的次数即可。可以使用一个哈希表配上双指针来做。

` class Solution { public boolean checkInclusion (String s1, String s2) { int l1 = s1.length(); int l2 = s2.length(); int [] count = new int [ 128 ]; if (l1 > l2) return false ; for ( int i = 0 ; i<l1; i++){ count[s1.charAt(i) - 'a' ]++; count[s2.charAt(i) - 'a' ]--; } if (allZero(count)) return true ; for ( int i = l1; i<l2; i++){ count[s2.charAt(i) - 'a' ]--; count[s2.charAt(i-l1) - 'a' ]++; if (allZero(count)) return true ; } return false ; } public boolean allZero ( int [] count) { int l = count.length; for ( int i = 0 ; i < l; i++){ if (count[i] != 0 ) return false ; } return true ; } } 复制代码`

### 6. 打印字符串的全排列 ###

> 
> 
> 
> 剑指offer： [字符串的排列](
> https://link.juejin.im?target=https%3A%2F%2Fwww.weiweiblog.cn%2Fpermutation%2F
> ) 输入一个字符串,按字典序打印出该字符串中字符的所有排列。例如输入字符串abc,则打印出由字符a,b,c所能排列出来的所有字符串abc,acb,bac,bca,cab和cba。
> 
> 
> 

把问题拆解成简单的步骤： 第一步求所有可能出现在第一个位置的字符（即把第一个字符和后面的所有字符交换[相同字符不交换]）； 第二步固定第一个字符，求后面所有字符的排列。这时候又可以把后面的所有字符拆成两部分（第一个字符以及剩下的所有字符），依此类推。这样，我们就可以用递归的方法来解决。

` public class Solution { ArrayList<String> res = new ArrayList<String>(); public ArrayList<String> Permutation (String str) { if (str == null ) return res; PermutationHelper(str.toCharArray(), 0 ); Collections.sort(res); return res; } public void PermutationHelper ( char [] str, int i) { if (i == str.length - 1 ){ res.add(String.valueOf(str)); } else { for ( int j = i; j < str.length; j++){ if (j!=i && str[i] == str[j]) continue ; swap(str, i, j); PermutationHelper(str, i+ 1 ); swap(str, i, j); } } } public void swap ( char [] str, int i, int j) { char temp = str[i]; str[i] = str[j]; str[j] = temp; } } 复制代码`

### 7. 第一个只出现一次的字符 ###

> 
> 
> 
> 剑指offer: [第一个只出现一次的字符](
> https://link.juejin.im?target=https%3A%2F%2Fwww.weiweiblog.cn%2Ffirstnotrepeatingchar%2F
> ) 在一个字符串(0<=字符串长度<=10000，全部由字母组成)中找到第一个只出现一次的字符,并返回它的位置, 如果没有则返回 -1.
> 
> 

先在hash表中统计各字母出现次数，第二次扫描直接访问hash表获得次数。也可以用数组代替hash表。

` import java.util.HashMap; public class Solution { public int FirstNotRepeatingChar (String str) { int len = str.length(); if (len == 0 ) return - 1 ; HashMap<Character, Integer> map = new HashMap<>(); for ( int i = 0 ; i < len; i++){ if (map.containsKey(str.charAt(i))){ int value = map.get(str.charAt(i)); map.put(str.charAt(i), value+ 1 ); } else { map.put(str.charAt(i), 1 ); } } for ( int i = 0 ; i < len; i++){ if (map.get(str.charAt(i)) == 1 ) return i; } return - 1 ; } } 复制代码`

### 8. 翻转单词顺序列 ###

> 
> 
> 
> 剑指offer: [翻转单词顺序列](
> https://link.juejin.im?target=https%3A%2F%2Fwww.weiweiblog.cn%2Freversesentence%2F
> ) LeetCode: [翻转字符串里的单词](
> https://link.juejin.im?target=https%3A%2F%2Fleetcode-cn.com%2Fproblems%2Freverse-words-in-a-string%2Fdescription%2F
> )
> 
> 

借助trim()和 split()就很容易搞定

` public class Solution { public String reverseWords (String s) { if (s.trim().length() == 0 ) return s.trim(); String [] temp = s.trim().split( " +" ); String res = "" ; for ( int i = temp.length - 1 ; i > 0 ; i--){ res += temp[i] + " " ; } return res + temp[ 0 ]; } } 复制代码`

### 9. 旋转字符串 ###

> 
> 
> 
> Leetcode: [旋转字符串](
> https://link.juejin.im?target=https%3A%2F%2Fleetcode-cn.com%2Fproblems%2Frotate-string%2Fdescription%2F
> ) 给定两个字符串, A 和 B。 A 的旋转操作就是将 A 最左边的字符移动到最右边。 例如, 若 A =
> 'abcde'，在移动一次之后结果就是'bcdea' 。如果在若干次旋转操作之后，A 能变成B，那么返回True。
> 
> 

![](https://user-gold-cdn.xitu.io/2018/9/5/165a8f54a6bda702?imageView2/0/w/1280/h/960/ignore-error/1) 一行代码搞定

` class Solution { public boolean rotateString (String A, String B) { return A.length() == B.length() && (A+A).contains(B); } } 复制代码`

### 9.1 左旋转字符串 ###

> 
> 
> 
> 剑指offer: [左旋转字符串](
> https://link.juejin.im?target=https%3A%2F%2Fwww.weiweiblog.cn%2Fleftrotatestring%2F
> ) 汇编语言中有一种移位指令叫做循环左移（ROL），现在有个简单的任务，就是用字符串模拟这个指令的运算结果。对于一个给定的字符序列S，请你把其循环左移K位后的序列输出。例如，字符序列S=”abcXYZdef”,要求输出循环左移3位后的结果，即“XYZdefabc”。是不是很简单？OK，搞定它！
> 
> 
> 

在第 n 个字符后面将切一刀，将字符串分为两部分，再重新并接起来即可。注意字符串长度为 0 的情况。

` public class Solution { public String LeftRotateString (String str, int n) { int len = str.length(); if (len == 0 ) return "" ; n = n % len; String s1 = str.substring(n, len); String s2 = str.substring( 0 , n); return s1+s2; } } 复制代码`

### 9.2 反转字符串 ###

> 
> 
> 
> LeetCode: [反转字符串](
> https://link.juejin.im?target=https%3A%2F%2Fleetcode-cn.com%2Fproblems%2Freverse-string%2Fdescription%2F
> ) 编写一个函数，其作用是将输入的字符串反转过来。
> 
> 

![](https://user-gold-cdn.xitu.io/2018/9/5/165a8f54a90b808a?imageView2/0/w/1280/h/960/ignore-error/1)

` class Solution { public String reverseString (String s) { if (s.length() < 2 ) return s; int l = 0 , r = s.length() - 1 ; char [] strs = s.toCharArray(); while (l < r){ char temp = strs[l]; strs[l] = strs[r]; strs[r] = temp; l++; r--; } return new String(strs); } } 复制代码`

### 10. 把字符串转换成整数 ###

> 
> 
> 
> 剑指offer: [把字符串转换成整数](
> https://link.juejin.im?target=https%3A%2F%2Fwww.weiweiblog.cn%2Fstrtoint%2F
> ) 将一个字符串转换成一个整数(实现Integer.valueOf(string)的功能，但是string不符合数字要求时返回0)，要求不能使用字符串转换整数的库函数。
> 数值为0或者字符串不是一个合法的数值则返回0。
> 
> 

` public class Solution { public int StrToInt (String str) { if (str.length() == 0 ) return 0 ; int flag = 0 ; if (str.charAt( 0 ) == '+' ) flag = 1 ; else if (str.charAt( 0 ) == '-' ) flag = 2 ; int start = flag > 0 ? 1 : 0 ; long res = 0 ; while (start < str.length()){ if (str.charAt(start) > '9' || str.charAt(start) < '0' ) return 0 ; res = res * 10 + (str.charAt(start) - '0' ); start ++; } return flag == 2 ? -( int )res : ( int )res; } } 复制代码`

### 11. 正则表达式匹配 ###

> 
> 
> 
> 剑指offer： [正则表达式匹配](
> https://link.juejin.im?target=https%3A%2F%2Fwww.weiweiblog.cn%2Fmatch%2F )
> 请实现一个函数用来匹配包括’.’和’*’的正则表达式。模式中的字符’.’表示任意一个字符，而’*’表示它前面的字符可以出现任意次（包含0次）。
> 在本题中，匹配是指字符串的所有字符匹配整个模式。例如，字符串”aaa”与模式”a.a”和”ab*ac*a”匹配，但是与”aa.a”和”ab*a”均不匹配
> 
> 
> 

![](https://user-gold-cdn.xitu.io/2018/9/5/165a8f54b2d1312a?imageView2/0/w/1280/h/960/ignore-error/1) 动态规划： 这里我们采用dp[i+1][j+1]代表s[0..i]匹配p[0..j]的结果，结果自然是采用布尔值True/False来表示。 首先，对边界进行赋值，显然dp[0][0] = true，两个空字符串的匹配结果自然为True; 接着，我们对dp[0][j+1]进行赋值，因为 i=0 是空串，如果一个空串和一个匹配串想要匹配成功，那么只有可能是p.charAt(j) == '*' && dp[0][j-1] 之后，就可以愉快地使用动态规划递推方程了。

` public boolean isMatch (String s, String p) { if (s == null || p == null ) { return false ; } boolean [][] dp = new boolean [s.length()+ 1 ][p.length()+ 1 ]; dp[ 0 ][ 0 ] = true ; for ( int j = 0 ; i < p.length(); j++) { if (p.charAt(j) == '*' && dp[ 0 ][j- 1 ]) { dp[ 0 ][j+ 1 ] = true ; } } for ( int i = 0 ; i < s.length(); i++) { for ( int j = 0 ; j < p.length(); j++) { if (p.charAt(j) == '.' ) { dp[i+ 1 ][j+ 1 ] = dp[i][j]; } if (p.charAt(j) == s.charAt(i)) { dp[i+ 1 ][j+ 1 ] = dp[i][j]; } if (p.charAt(j) == '*' ) { if (p.charAt(j- 1 ) != s.charAt(i) && p.charAt(j- 1 ) != '.' ) { dp[i+ 1 ][j+ 1 ] = dp[i+ 1 ][j- 1 ]; } else { dp[i+ 1 ][j+ 1 ] = (dp[i+ 1 ][j] || dp[i][j+ 1 ] || dp[i+ 1 ][j- 1 ]); } } } } return dp[s.length()][p.length()]; } 复制代码`

### 12. 表示数值的字符串 ###

> 
> 
> 
> 剑指offer: [表示数值的字符串](
> https://link.juejin.im?target=https%3A%2F%2Fwww.weiweiblog.cn%2Fisnumeric%2F
> ) 请实现一个函数用来判断字符串是否表示数值（包括整数和小数）。例如，字符串”+100″,”5e2″,”-123″,”3.1416″和”-1E-16″都表示数值。
> 但是”12e”,”1a3.14″,”1.2.3″,”+-5″和”12e+4.3″都不是。
> 
> 

设置三个标志符分别记录“+/-”、“e/E”和“.”是否出现过。

` public class Solution { public boolean isNumeric ( char [] str) { int len = str.length; boolean sign = false , decimal = false , hasE = false ; for ( int i = 0 ; i < len; i++){ if (str[i] == '+' || str[i] == '-' ){ if (!sign && i > 0 && str[i- 1 ] != 'e' && str[i- 1 ] != 'E' ) return false ; if (sign && str[i- 1 ] != 'e' && str[i- 1 ] != 'E' ) return false ; sign = true ; } else if (str[i] == 'e' || str[i] == 'E' ){ if (i == len - 1 ) return false ; if (hasE) return false ; hasE = true ; } else if (str[i] == '.' ){ if (hasE || decimal) return false ; decimal = true ; } else if (str[i] < '0' || str[i] > '9' ) return false ; } return true ; } } 复制代码`

### 13. 字符流中第一个不重复的字符 ###

> 
> 
> 
> 剑指offer: [字符流中第一个不重复的字符](
> https://link.juejin.im?target=https%3A%2F%2Fwww.weiweiblog.cn%2Ffirstappearingonce%2F
> ) 请实现一个函数用来找出字符流中第一个只出现一次的字符。例如，当从字符流中只读出前两个字符”go”时，第一个只出现一次的字符是”g”。当从该字符流中读出前六个字符“google”时，第一个只出现一次的字符是”l”。
> 
> 
> 

用一个哈希表来存储每个字符及其出现的次数，另外用一个字符串 s 来保存字符流中字符的顺序。

` import java.util.HashMap; public class Solution { HashMap<Character, Integer> map = new HashMap<Character, Integer>(); StringBuffer s = new StringBuffer(); //Insert one char from stringstream public void Insert ( char ch) { s.append(ch); if (map.containsKey(ch)){ map.put(ch, map.get(ch)+ 1 ); } else { map.put(ch, 1 ); } } //return the first appearence once char in current stringstream public char FirstAppearingOnce () { for ( int i = 0 ; i < s.length(); i++){ if (map.get(s.charAt(i)) == 1 ) return s.charAt(i); } return '#' ; } } 复制代码`