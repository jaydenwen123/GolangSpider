# Java 8新的时间日期库的20个使用示例 #

## Java 8是如何处理时间及日期的 ##

有人问我学习一个新库的最佳途径是什么？我的回答是，就是在实际项目中那样去使用它。在一个真实的项目中会有各种各样的需求，这会促使开发人员去探索和研究这个新库。简言之，只有任务本身才会真正促使你去探索及学习。java 8的新的日期及时间API也是一样。为了学习Java 8的这个新库，这里我创建了20个以任务为导向的例子。我们先从一个简单的任务开始，比如说如何用Java 8的时间日期库来表示今天，接着再进一步生成一个带时间及时区的完整日期，然后再研究下如何完成一些更实际的任务，比如说开发一个提醒类的应用，来找出距离一些特定日期比如生日，周日纪念日，下一个帐单日，下一个溢价日或者信用卡过期时间还有多少天。

## 示例1 如何 在Java 8中获取当天的日期 ##

Java 8中有一个叫LocalDate的类，它能用来表示今天的日期。这个类与java.util.Date略有不同，因为它只包含日期，没有时间。因此，如果你只需要表示日期而不包含时间，就可以使用它。

` LocalDate today = LocalDate.now(); System.out.println( "Today's Local date : " + today); 输出 Today 's Local date : 2018-02-11 复制代码`

你可以看到它创建了今天的日期却不包含时间信息。它还将日期格式化完了再输出出来，不像之前的Date类那样，打印出来的数据都是未经格式化的。

## 示例2 如何在Java 8中获取当前的年月日 ##

LocalDate类中提供了一些很方便的方法可以用于提取出年月日以及其它的日期属性。使用这些方法，你可以获取到任何你所需要的日期属性，而不再需要使用java.util.Calendar这样的类了：

` LocalDate today = LocalDate.now(); int year = today.getYear(); int month = today.getMonthValue(); int day = today.getDayOfMonth(); System.out.printf( "Year : %d Month : %d day : %d \t %n" , year, month, day); 输出 Today 's Local date : 2018-02-11 Year : 2018 Month : 2 day : 11 复制代码`

可以看到，在Java 8中获取年月信息非常简单，只需使用对应的getter方法就好了，无需记忆，非常直观。你可以拿它和Java中老的获取当前年月日的写法进行一下比较。

## 示例3 在Java 8中如何获取某个特定的日期 ##

在第一个例子中，我们看到通过静态方法now()来生成当天日期是非常简单的，不过通过另一个十分有用的工厂方法LocalDate.of()，则可以创建出任意一个日期，它接受年月日的参数，然后返回一个等价的LocalDate实例。关于这个方法还有一个好消息就是它没有再犯之前API中的错，比方说，年只能从1900年开始，月必须从0开始，等等。这里的日期你写什么就是什么，比如说，下面这个例子中它代表的就是1月14日，没有什么隐藏逻辑。

` LocalDate dateOfBirth = LocalDate.of(2010, 01, 14); System.out.println( "Your Date of birth is : " + dateOfBirth); 输出 : Your Date of birth is : 2010-01-14 复制代码`

可以看出，创建出来的日期就是我们所写的那样，2014年1月14日。

## 示例4 在Java 8中如何检查两个日期是否相等 ##

如果说起现实中实际的处理时间及日期的任务，有一个常见的就是要检查两个日期是否相等。你可能经常会碰到要判断今天是不是某个特殊的日子，比如生日啊，周年纪念日啊，或者假期之类。有的时候，会给你一个日期，让你检查它是不是某个日子比方说假日。下面这个例子将会帮助你在Java 8中完成这类任务。正如你所想的那样，LocalDate重写了equals方法来进行日期的比较，如下所示：

` LocalDate date1 = LocalDate.of(2014, 01, 14); if (date1.equals(today)){ System.out.printf( "Today %s and date1 %s are same date %n" , today, date1); } 输出 today 2014-01-14 and date1 2014-01-14 are same date 复制代码`

在本例中我们比较的两个日期是相等的。同时，如果在代码中你拿到了一个格式化好的日期串，你得先将它解析成日期然后才能比较。你可以将这个例子与Java之前比较日期的方式进行下比较，你会发现它真是爽多了。

## 示例5 在Java 8中如何检查重复事件，比如说生日 ##

在Java中还有一个与时间日期相关的实际任务就是检查重复事件，比如说每月的帐单日，结婚纪念日，每月还款日或者是每年交保险费的日子。如果你在一家电商公司工作的话，那么肯定会有这么一个模块，会去给用户发送生日祝福并且在每一个重要的假日给他们捎去问候，比如说圣诞节，感恩节，在印度则可能是万灯节（Deepawali）。如何在Java中判断是否是某个节日或者重复事件？使用MonthDay类。这个类由月日组合，不包含年信息，也就是说你可以用它来代表每年重复出现的一些日子。当然也有一些别的组合，比如说YearMonth类。它和新的时间日期库中的其它类一样也都是不可变且线程安全的，并且它还是一个值类（value class）。我们通过一个例子来看下如何使用MonthDay来检查某个重复的日期：

` LocalDate dateOfBirth = LocalDate.of(2010, 01, 14); MonthDay birthday = MonthDay.of(dateOfBirth.getMonth(), dateOfBirth.getDayOfMonth()); MonthDay currentMonthDay = MonthDay.from(today); if (currentMonthDay.equals(birthday)){ System.out.println( "Many Many happy returns of the day !!" ); } else { System.out.println( "Sorry, today is not your birthday" ); } 输出: Many Many happy returns of the day !! 复制代码`

虽然年不同，但今天就是生日的那天，所以在输出那里你会看到一条生日祝福。你可以调整下系统的时间再运行下这个程序看看它是否能提醒你下一个生日是什么时候，你还可以试着用你的下一个生日来编写一个JUnit单元测试看看代码能否正确运行。

## 示例6 如何在Java 8中获取当前时间 ##

这与第一个例子中获取当前日期非常相似。这次我们用的是一个叫LocalTime的类，它是没有日期的时间，与LocalDate是近亲。这里你也可以用静态工厂方法now()来获取当前时间。默认的格式是hh:mm:ss:nnn，这里的nnn是纳秒。可以和Java 8以前如何获取当前时间做一下比较。

` LocalTime time = LocalTime.now(); System.out.println( "local time now : " + time); 输出 local time now : 16:33:33.369 // in hour, minutes, seconds, nano seconds 复制代码`

可以看到，当前时间是不包含日期的，因为LocalTime只有时间，没有日期。

## 示例7 如何增加时间里面的小时数 ##

很多时候我们需要增加小时，分或者秒来计算出将来的时间。Java 8不仅提供了不可变且线程安全的类，它还提供了一些更方便的方法譬如plusHours()来替换原来的add()方法。顺便说一下，这些方法返回的是一个新的LocalTime实例的引用，因为LocalTime是不可变的，可别忘了存储好这个新的引用。

` LocalTime time = LocalTime.now(); LocalTime newTime = time.plusHours(2); // adding two hours System.out.println( "Time after 2 hours : " + newTime); 输出 : Time after 2 hours : 18:33:33.369 复制代码`

可以看到当前时间2小时后是16:33:33.369。现在你可以将它和Java中增加或者减少小时的老的方式进行下比较。一看便知哪种方式更好。

## 示例8 如何获取1周后的日期 ##

这与前一个获取2小时后的时间的例子类似，这里我们将学会如何获取到1周后的日期。LocalDate是用来表示无时间的日期的，它有一个plus()方法可以用来增加日，星期，或者月，ChronoUnit则用来表示这个时间单位。由于LocalDate也是不可变的，因此任何修改操作都会返回一个新的实例，因此别忘了保存起来。

` LocalDate nextWeek = today.plus(1, ChronoUnit.WEEKS); System.out.println( "Today is : " + today); System.out.println( "Date after 1 week : " + nextWeek); 输出: Today is : 2018-01-14 Date after 1 week : 2018-01-21 复制代码`

可以看到7天也就是一周后的日期是什么。你可以用这个方法来增加一个月，一年，一小时，一分钟，甚至是十年，查看下Java API中的ChronoUnit类来获取更多选项。

## 示例9 一年前后的日期 ##

这是上个例子的续集。上例中，我们学习了如何使用LocalDate的plus()方法来给日期增加日，周或者月，现在我们来学习下如何用minus()方法来找出一年前的那天。

` LocalDate previousYear = today.minus(1, ChronoUnit.YEARS); System.out.println( "Date before 1 year : " + previousYear); LocalDate nextYear = today.plus(1, YEARS); System.out.println( "Date after 1 year : " + nextYear); 输出: Date before 1 year : 2013-01-14 Date after 1 year : 2015-01-14 复制代码`

可以看到现在一共有两年，一个是2013年，一个是2015年，分别是2014的前后那年。

## 示例10 在Java 8中使用时钟 ##

Java 8中自带了一个Clock类，你可以用它来获取某个时区下当前的瞬时时间，日期或者时间。可以用Clock来替代 **System.currentTimeInMillis()** 与 **TimeZone.getDefault()** 方法。

` // Returns the current time based on your system clock and set to UTC. Clock clock = Clock.systemUTC(); System.out.println( "Clock : " + clock); // Returns time based on system clock zone Clock defaultClock = Clock.systemDefaultZone(); System.out.println( "Clock : " + clock); 输出: Clock : SystemClock[Z] Clock : SystemClock[Z] 复制代码`

你可以用指定的日期来和这个时钟进行比较，比如下面这样：

` public class MyClass { private Clock clock; // dependency inject ... public void process(LocalDate eventDate) { if (eventDate.isBefore(LocalDate.now(clock)) { ... } } } 复制代码`

如果你需要对不同时区的日期进行处理的话这是相当方便的。

## 示例11 在Java中如何判断某个日期是在另一个日期的前面还是后面 ##

这也是实际项目中常见的一个任务。你怎么判断某个日期是在另一个日期的前面还是后面，或者正好相等呢？在Java 8中，LocalDate类有一个isBefore()和isAfter()方法可以用来比较两个日期。如果调用方法的那个日期比给定的日期要早的话，isBefore()方法会返回true。

` LocalDate tomorrow = LocalDate.of(2014, 1, 15); 、 if (tommorow.isAfter(today)){ System.out.println( "Tomorrow comes after today" ); } LocalDate yesterday = today.minus(1, DAYS); if (yesterday.isBefore(today)){ System.out.println( "Yesterday is day before today" ); } 输出: Tomorrow comes after today Yesterday is day before today 复制代码`

可以看到在Java 8中进行日期比较非常简单。不需要再用像Calendar这样的另一个类来完成类似的任务了。

## 示例12 在Java 8中处理不同的时区 ##

Java 8不仅将日期和时间进行了分离，同时还有时区。现在已经有好几组与时区相关的类了，比如ZonId代表的是某个特定的时区，而ZonedDateTime代表的是带时区的时间。它等同于Java 8以前的GregorianCalendar类。使用这个类，你可以将本地时间转换成另一个时区中的对应时间，比如下面这个例子：

` // Date and time with timezone in Java 8 ZoneId america = ZoneId.of( "America/New_York" ); LocalDateTime localtDateAndTime = LocalDateTime.now(); ZonedDateTime dateAndTimeInNewYork = ZonedDateTime.of(localtDateAndTime, america ); System.out.println( "Current date and time in a particular timezone : " + dateAndTimeInNewYork); 输出 : Current date and time in a particular timezone : 2014-01-14T16:33:33.373-05:00[America/New_York] 复制代码`

可以拿它跟之前将本地时间转换成GMT时间的方式进行下比较。顺便说一下，正如Java 8以前那样，对应时区的那个文本可别弄错了，否则你会碰到这么一个异常：

` Exception in thread "main" java.time.zone.ZoneRulesException: Unknown time-zone ID: ASIA/Tokyo at java.time.zone.ZoneRulesProvider.getProvider(ZoneRulesProvider.java:272) at java.time.zone.ZoneRulesProvider.getRules(ZoneRulesProvider.java:227) at java.time.ZoneRegion.ofId(ZoneRegion.java:120) at java.time.ZoneId.of(ZoneId.java:403) at java.time.ZoneId.of(ZoneId.java:351) 复制代码`

## 示例13 如何表示固定的日期，比如信用卡过期时间 ##

正如MonthDay表示的是某个重复出现的日子的，YearMonth又是另一个组合，它代表的是像信用卡还款日，定期存款到期日，options到期日这类的日期。你可以用这个类来找出那个月有多少天，lengthOfMonth()这个方法返回的是这个YearMonth实例有多少天，这对于检查2月到底是28天还是29天可是非常有用的。

` YearMonth currentYearMonth = YearMonth.now(); System.out.printf( "Days in month year %s: %d%n" , currentYearMonth, currentYearMonth.lengthOfMonth()); YearMonth creditCardExpiry = YearMonth.of(2018, Month.FEBRUARY); System.out.printf( "Your credit card expires on %s %n" , creditCardExpiry); 输出: Days in month year 2014-01: 31 Your credit card expires on 2018-02 复制代码`

## 示例14 如何在Java 8中检查闰年 ##

这并没什么复杂的，LocalDate类有一个isLeapYear()的方法能够返回当前LocalDate对应的那年是否是闰年。如果你还想重复造轮子的话，可以看下这段代码，这是纯用Java编写的判断某年是否是闰年的逻辑。

` if (today.isLeapYear()){ System.out.println( "This year is Leap year" ); } else { System.out.println( "2018 is not a Leap year" ); } 输出: 2018 is not a Leap year 复制代码`

你可以多检查几年看看结果是否正确，最好写一个单元测试来对正常年份和闰年进行下测试。

## 示例15 两个日期之间包含多少天，多少个月 ##

还有一个常见的任务就是计算两个给定的日期之间包含多少天，多少周或者多少年。你可以用java.time.Period类来完成这个功能。在下面这个例子中，我们将计算当前日期与将来的一个日期之前一共隔着几个月。

` LocalDate java8Release = LocalDate.of(2014, Month.MARCH, 14); Period periodToNextJavaRelease = Period.between(today, java8Release); System.out.println( "Months left between today and Java 8 release : " + periodToNextJavaRelease.getMonths() ); 输出: Months left between today and Java 8 release : 2 复制代码`

可以看到，本月是1月，而Java 8的发布日期是3月，因此中间隔着2个月。

## 示例16 带时区偏移量的日期与时间 ##

在Java 8里面，你可以用ZoneOffset类来代表某个时区，比如印度是GMT或者UTC5：30，你可以使用它的静态方法ZoneOffset.of()方法来获取对应的时区。只要获取到了这个偏移量，你就可以拿LocalDateTime和这个偏移量创建出一个OffsetDateTime。

` LocalDateTime datetime = LocalDateTime.of(2014, Month.JANUARY, 14, 19, 30); ZoneOffset offset = ZoneOffset.of( "+05:30" ); OffsetDateTime date = OffsetDateTime.of(datetime, offset); System.out.println( "Date and Time with timezone offset in Java : " + date); 输出 : Date and Time with timezone offset in Java : 2014-01-14T19:30+05:30 复制代码`

可以看到现在时间日期与时区是关联上了。还有一点就是，OffSetDateTime主要是给机器来理解的，如果是给人看的，可以使用ZoneDateTime类。

## 示例17 在Java 8中如何获取当前时间戳 ##

如果你还记得在Java 8前是如何获取当前时间戳的，那现在这简直就是小菜一碟了。Instant类有一个静态的工厂方法now()可以返回当前时间戳，如下：

` Instant timestamp = Instant.now(); System.out.println( "What is value of this instant " + timestamp); 输出 : What is value of this instant 2014-01-14T08:33:33.379Z 复制代码`

可以看出，当前时间戳是包含日期与时间的，与java.util.Date很类似，事实上Instant就是Java 8前的Date，你可以使用这两个类中的方法来在这两个类型之间进行转换，比如Date.from(Instant)是用来将Instant转换成java.util.Date的，而Date.toInstant()是将Date转换成Instant的。

## 示例18 如何在Java 8中使用预定义的格式器来对日期进行解析/格式化 ##

在Java 8之前，时间日期的格式化可是个技术活，我们的好伙伴SimpleDateFormat并不是线程安全的，而如果用作本地变量来格式化的话又显得有些笨重。多亏了线程本地变量，这使得它在多线程环境下也算有了用武之地，但Java维持这一状态也有很长一段时间了。这次它引入了一个全新的线程安全的日期与时间格式器。它还自带了一些预定义好的格式器，包含了常用的日期格式。比如说，本例 中我们就用了预定义的BASIC_ISO_DATE格式，它会将2014年2月14日格式化成20140114。

` String dayAfterTommorrow = "20140116" ; LocalDate formatted = LocalDate.parse(dayAfterTommorrow, DateTimeFormatter.BASIC_ISO_DATE); System.out.printf( "Date generated from String %s is %s %n" , dayAfterTommorrow, formatted); 输出 : Date generated from String 20140116 is 2014-01-16 复制代码`

你可以看到生成的日期与指定字符串的值是匹配的，就是日期格式上略有不同。

## 示例19 如何在Java中使用自定义的格式器来解析日期 ##

在上例中，我们使用了内建的时间日期格式器来解析日期字符串。当然了，预定义的格式器的确不错但有时候你可能还是需要使用自定义的日期格式，这个时候你就得自己去创建一个自定义的日期格式器实例了。下面这个例子中的日期格式是”MMM dd yyyy”。你可以给DateTimeFormatter的ofPattern静态方法()传入任何的模式，它会返回一个实例，这个模式的字面量与前例中是相同的。比如说M还是代表月，而m仍是分。无效的模式会抛出DateTimeParseException异常，但如果是逻辑上的错误比如说该用M的时候用成m，这样就没办法了。

` String goodFriday = "Apr 18 2014" ; try { DateTimeFormatter formatter = DateTimeFormatter.ofPattern( "MMM dd yyyy" ); LocalDate holiday = LocalDate.parse(goodFriday, formatter); System.out.printf( "Successfully parsed String %s, date is %s%n" , goodFriday, holiday); } catch (DateTimeParseException ex) { System.out.printf( "%s is not parsable!%n" , goodFriday); ex.printStackTrace(); } 输出 : Successfully parsed String Apr 18 2014, date is 2014-04-18 复制代码`

可以看到日期的值与传入的字符串的确是相符的，只是格式不同。

## 示例20 如何在Java 8中对日期进行格式化，转换成字符串 ##

在上两个例子中，尽管我们用到了DateTimeFormatter类但我们主要是进行日期字符串的解析。在这个例子中我们要做的事情正好相反。这里我们有一个LocalDateTime类的实例，我们要将它转换成一个格式化好的日期串。这是目前为止Java中将日期转换成字符串最简单便捷的方式了。下面这个例子将会返回一个格式化好的字符串。与前例相同的是，我们仍需使用指定的模式串去创建一个DateTimeFormatter类的实例，但调用的并不是LocalDate类的parse方法，而是它的format()方法。这个方法会返回一个代表当前日期的字符串，对应的模式就是传入的DateTimeFormatter实例中所定义好的。

` LocalDateTime arrivalDate = LocalDateTime.now(); try { DateTimeFormatter format = DateTimeFormatter.ofPattern( "MMM dd yyyy hh:mm a" ); String landing = arrivalDate.format(format); System.out.printf( "Arriving at : %s %n" , landing); } catch (DateTimeException ex) { System.out.printf( "%s can't be formatted!%n" , arrivalDate); ex.printStackTrace(); } Output : Arriving at : Jan 14 2014 04:33 PM 复制代码`

可以看到，当前时间是用给定的”MMM dd yyyy hh:mm a”模式来表示的，它包含了三个字母表示的月份以及用AM及PM来表示的时间。

## Java 8中日期与时间API的几个关键点 ##

看完了这些例子后，我相信你已经对Java 8这套新的时间日期API有了一定的了解了。现在我们来回顾下关于这个新的API的一些关键的要素。

* 它提供了javax.time.ZoneId用来处理时区。
* 它提供了LocalDate与LocalTime类
* Java 8中新的时间与日期API中的所有类都是不可变且线程安全的，这与之前的Date与Calendar API中的恰好相反，那里面像java.util.Date以及SimpleDateFormat这些关键的类都不是线程安全的。
* 新的时间与日期API中很重要的一点是它定义清楚了基本的时间与日期的概念，比方说，瞬时时间，持续时间，日期，时间，时区以及时间段。它们都是基于ISO日历体系的。
* 每个Java开发人员都应该至少了解这套新的API中的这五个类：
* Instant 它代表的是时间戳，比如2014-01-14T02:20:13.592Z，这可以从java.time.Clock类中获取，像这样： Instant current = Clock.system(ZoneId.of(“Asia/Tokyo”)).instant();
* LocalDate 它表示的是不带时间的日期，比如2014-01-14。它可以用来存储生日，周年纪念日，入职日期等。
* LocalTime – 它表示的是不带日期的时间
* LocalDateTime – 它包含了时间与日期，不过没有带时区的偏移量
* ZonedDateTime – 这是一个带时区的完整时间，它根据UTC/格林威治时间来进行时区调整
* 这个库的主包是java.time，里面包含了代表日期，时间，瞬时以及持续时间的类。它有两个子package，一个是java.time.foramt，这个是什么用途就很明显了，还有一个是java.time.temporal，它能从更低层面对各个字段进行访问。
* 时区指的是地球上共享同一标准时间的地区。每个时区都有一个唯一标识符，同时还有一个地区/城市(Asia/Tokyo)的格式以及从格林威治时间开始的一个偏移时间。比如说，东京的偏移时间就是+09:00。
* OffsetDateTime类实际上包含了LocalDateTime与ZoneOffset。它用来表示一个包含格林威治时间偏移量（+/-小时：分，比如+06:00或者 -08：00）的完整的日期（年月日）及时间（时分秒，纳秒）。
* DateTimeFormatter类用于在Java中进行日期的格式化与解析。与SimpleDateFormat不同，它是不可变且线程安全的，如果需要的话，可以赋值给一个静态变量。DateTimeFormatter类提供了许多预定义的格式器，你也可以自定义自己想要的格式。当然了，根据约定，它还有一个parse()方法是用于将字符串转换成日期的，如果转换期间出现任何错误，它会抛出DateTimeParseException异常。类似的，DateFormatter类也有一个用于格式化日期的format()方法，它出错的话则会抛出DateTimeException异常。
* 再说一句，“MMM d yyyy”与“MMm dd yyyy”这两个日期格式也略有不同，前者能识别出”Jan 2 2014″与”Jan 14 2014″这两个串，而后者如果传进来的是”Jan 2 2014″则会报错，因为它期望月份处传进来的是两个字符。为了解决这个问题，在天为个位数的情况下，你得在前面补0，比如”Jan 2 2014″应该改为”Jan 02 2014″。

关于Java 8这个新的时间日期API就讲到这了。这几个简短的示例 对于理解这套新的API中的一些新增类已经足够了。由于它是基于实际任务来讲解的，因此后面再遇到Java中要对时间与日期进行处理的工作时，就不用再四处寻找了。我们学习了如何创建与修改日期实例。我们还了解了纯日期，日期加时间，日期加时区的区别，知道如何比较两个日期，如何找到某天到指定日期比如说下一个生日，周年纪念日或者保险日还有多少天。我们还学习了如何在Java 8中用线程安全的方式对日期进行解析及格式化，而无需再使用线程本地变量或者第三方库这种取巧的方式。新的API能胜任任何与时间日期相关的任务。

## 福利部分： ##

**[《大数据成神之路》]( https://link.juejin.im/?target=https%3A%2F%2Fshimo.im%2Fdocs%2FjdPhrtFwVCAMkoWv%2F )**

[《几百TJava和大数据资源下载》]( https://link.juejin.im/?target=https%3A%2F%2Fshimo.im%2Fdocs%2FsSTmRlbM4FUbUwX4%2F )

![](https://user-gold-cdn.xitu.io/2019/3/26/169b5973fef05422?imageView2/0/w/1280/h/960/ignore-error/1)