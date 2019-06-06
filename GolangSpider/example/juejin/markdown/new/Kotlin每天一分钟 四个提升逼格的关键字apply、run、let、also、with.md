# Kotlin每天一分钟 四个提升逼格的关键字apply、run、let、also、with #

使用kotlin的时候对这几个关键字的使用不是很明确,找了很多文章和书籍来加深一下印象,感谢mikyou

开篇看结论

![image.png](https://user-gold-cdn.xitu.io/2019/6/6/16b2b4575bd20600?imageView2/0/w/1280/h/960/ignore-error/1)

# let #

let扩展函数的实际上是一个作用域函数，当你需要去定义一个变量在一个特定的作用域范围内，let函数的是一个不错的选择；let函数另一个作用就是可以避免写一些判断null的操作。

* 

## let函数的一般结构 ##

` object.let{ it.todo()//在函数体内使用it替代object对象去访问其公有的属性和方法 ... } //另一种用途 判断object为null的操作 object?.let{//表示object不为null的条件下，才会去执行 let 函数体 it.todo() } 复制代码`

* 

## let函数的kotlin和Java转化 ##

` //kotlin fun main(args: Array<String>) { val result = "testLet".let { println(it.length) 1000 } println(result) } //java public final class LetFunctionKt { public static final void main(@NotNull String[] args) { Intrinsics.checkParameterIsNotNull(args, "args" ); String var2 = "testLet" ; int var4 = var2.length(); System.out.println(var4); int result = 1000; System.out.println(result); } } 复制代码`

* ##let函数使用前后的对比

` mVideoPlayer?.setVideoView(activity.course_video_view) mVideoPlayer?.setControllerView(activity.course_video_controller_view) mVideoPlayer?.setCurtainView(activity.course_video_curtain_view) ------------------------------------------------------------------------------------------------------------------------------ mVideoPlayer?.let { it.setVideoView(activity.course_video_view) it.setControllerView(activity.course_video_controller_view) it.setCurtainView(activity.course_video_curtain_view) } 复制代码`

* let函数适用的场景

#### 场景一: 最常用的场景就是使用let函数处理需要针对一个可null的对象统一做判空处理。 ####

#### 场景二: 然后就是需要去明确一个变量所处特定的作用域范围内可以使用 ####

# with #

* 

## with函数使用的一般结构 ##

` with(object){ //todo } 复制代码`

* 

## with函数的kotlin和Java转化 ##

` //kotlin fun main(args: Array<String>) { val user = User( "Kotlin" , 1, "1111111" ) val result = with(user) { println( "my name is $name , I am $age years old, my phone number is $phoneNum " ) 1000 } println( "result: $result " ) } ------------------------------------------------------------------------------------------------------------------------------ //java public static final void main(@NotNull String[] args) { Intrinsics.checkParameterIsNotNull(args, "args" ); User user = new User( "Kotlin" , 1, "1111111" ); String var4 = "my name is " + user.getName() + ", I am " + user.getAge() + " years old, my phone number is " + user.getPhoneNum(); System.out.println(var4); int result = 1000; String var3 = "result: " + result; System.out.println(var3); } 复制代码`

* 

## with函数使用前后的对比 ##

` override fun onBindViewHolder(holder: ViewHolder, position: Int){ val item = getItem(position)?: return with(item){ holder.tvNewsTitle.text = StringUtils.trimToEmpty(titleEn) holder.tvNewsSummary.text = StringUtils.trimToEmpty(summary) holder.tvExtraInf.text = "难度： $gradeInfo | 单词数： $length | 读后感: $numReviews " } } ------------------------------------------------------------------------------------------------------------------------------ @Override public void onBindViewHolder(ViewHolder holder, int position) { ArticleSnippet item = getItem(position); if (item == null) { return ; } holder.tvNewsTitle.setText(StringUtils.trimToEmpty(item.titleEn)); holder.tvNewsSummary.setText(StringUtils.trimToEmpty(item.summary)); String gradeInfo = "难度：" + item.gradeInfo; String wordCount = "单词数：" + item.length; String reviewNum = "读后感：" + item.numReviews; String extraInfo = gradeInfo + " | " + wordCount + " | " + reviewNum; holder.tvExtraInfo.setText(extraInfo); } 复制代码`

* with函数的适用的场景 **适用于调用同一个类的多个方法时，可以省去类名重复，直接调用类的方法即可，经常用于Android中RecyclerView中onBinderViewHolder中，数据model的属性映射到UI上**

# run #

* 

## run函数使用的一般结构 ##

` object.run{ //todo } 复制代码`

* 

## run函数的kotlin和Java转化 ##

` //java public static final void main(@NotNull String[] args) { Intrinsics.checkParameterIsNotNull(args, "args" ); User user = new User( "Kotlin" , 1, "1111111" ); String var5 = "my name is " + user.getName() + ", I am " + user.getAge() + " years old, my phone number is " + user.getPhoneNum(); System.out.println(var5); int result = 1000; String var3 = "result: " + result; System.out.println(var3); } ------------------------------------------------------------------------------------------------------------------------------ //kotlin fun main(args: Array<String>) { val user = User( "Kotlin" , 1, "1111111" ) val result = user.run { println( "my name is $name , I am $age years old, my phone number is $phoneNum " ) 1000 } println( "result: $result " ) } 复制代码`

* 

## run函数使用前后对比 ##

` override fun onBindViewHolder(holder: ViewHolder, position: Int){ val item = getItem(position)?: return with(item){ holder.tvNewsTitle.text = StringUtils.trimToEmpty(titleEn) holder.tvNewsSummary.text = StringUtils.trimToEmpty(summary) holder.tvExtraInf = "难度： $gradeInfo | 单词数： $length | 读后感: $numReviews "... } } // 使用后 override fun onBindViewHolder(holder: ViewHolder, position: Int){ getItem(position)?.run{ holder.tvNewsTitle.text = StringUtils.trimToEmpty(titleEn) holder.tvNewsSummary.text = StringUtils.trimToEmpty(summary) holder.tvExtraInf = "难度： $gradeInfo | 单词数： $length | 读后感: $numReviews "... } } 复制代码`

* 

## run函数使用场景 ##

**适用于let,with函数任何场景。因为run函数是let,with两个函数结合体，准确来说它弥补了let函数在函数体内必须使用it参数替代对象，在run函数中可以像with函数一样可以省略，直接访问实例的公有属性和方法，另一方面它弥补了with函数传入对象判空问题，在run函数中可以像let函数一样做判空处理**

# apply #

* 

## apply函数使用的一般结构 ##

` object.apply{ //todo } 复制代码`

* 

## apply函数的kotlin和Java转化 ##

` //java public final class ApplyFunctionKt { public static final void main(@NotNull String[] args) { Intrinsics.checkParameterIsNotNull(args, "args" ); User user = new User( "Kotlin" , 1, "1111111" ); String var5 = "my name is " + user.getName() + ", I am " + user.getAge() + " years old, my phone number is " + user.getPhoneNum(); System.out.println(var5); String var3 = "result: " + user; System.out.println(var3); } } //kotlin fun main(args: Array<String>) { val user = User( "Kotlin" , 1, "1111111" ) val result = user.apply { println( "my name is $name , I am $age years old, my phone number is $phoneNum " ) 1000 } println( "result: $result " ) } 复制代码`

* 

## apply函数使用前后的对比 ##

` //使用前 mSheetDialogView = View.inflate(activity, R.layout.biz_exam_plan_layout_sheet_inner, null) mSheetDialogView.course_comment_tv_label.paint.isFakeBoldText = true mSheetDialogView.course_comment_tv_score.paint.isFakeBoldText = true mSheetDialogView.course_comment_tv_cancel.paint.isFakeBoldText = true mSheetDialogView.course_comment_tv_confirm.paint.isFakeBoldText = true mSheetDialogView.course_comment_seek_bar.max = 10 mSheetDialogView.course_comment_seek_bar.progress = 0 //使用后 mSheetDialogView = View.inflate(activity, R.layout.biz_exam_plan_layout_sheet_inner, null).apply{ course_comment_tv_label.paint.isFakeBoldText = true course_comment_tv_score.paint.isFakeBoldText = true course_comment_tv_cancel.paint.isFakeBoldText = true course_comment_tv_confirm.paint.isFakeBoldText = true course_comment_seek_bar.max = 10 course_comment_seek_bar.progress = 0 } //多级判空 if (mSectionMetaData == null || mSectionMetaData.questionnaire == null || mSectionMetaData.section == null) { return ; } if (mSectionMetaData.questionnaire.userProject != null) { renderAnalysis(); return ; } if (mSectionMetaData.section != null && !mSectionMetaData.section.sectionArticles.isEmpty()) { fetchQuestionData(); return ; } mSectionMetaData?.apply{ //mSectionMetaData不为空的时候操作mSectionMetaData }?.questionnaire?.apply{ //questionnaire不为空的时候操作questionnaire }?.section?.apply{ //section不为空的时候操作section }?.sectionArticle?.apply{ //sectionArticle不为空的时候操作sectionArticle } 复制代码`

# also #

* 

## also函数使用的一般结构 ##

` object.also{ //todo } 复制代码`

* 

## also函数编译后的class文件 ##

` //java public final class AlsoFunctionKt { public static final void main(@NotNull String[] args) { Intrinsics.checkParameterIsNotNull(args, "args" ); String var2 = "testLet" ; int var4 = var2.length(); System.out.println(var4); System.out.println(var2); } } //kotlin fun main(args: Array<String>) { val result = "testLet".also { println(it.length) 1000 } println(result) } 复制代码`

* 

## also函数的适用场景 ##

**适用于let函数的任何场景，also函数和let很像，只是唯一的不同点就是let函数最后的返回值是最后一行的返回值而also函数的返回值是返回当前的这个对象。一般可用于多个扩展函数链式调用**