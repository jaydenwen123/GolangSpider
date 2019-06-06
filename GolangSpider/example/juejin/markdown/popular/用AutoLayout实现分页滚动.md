# 用AutoLayout实现分页滚动 #

### 滚动视图分页 ###

UIScrollView的pagingEnabled属性用于控制是否按分页进行滚动。在一些应用中会应用到这一个特性，最典型的就是手机桌面的应用图标列表。这些界面中往往每一页功能都比较独立，系统也提供了UIPageViewController来实现这种分页滚动的功能。 实现分页滚动的UI实现一般是最外层一个UIScrollView。然后UIScrollView里面是一个总体的容器视图containerView。容器视图添加N个页视图，对于水平分页滚动来说容器视图的高度和滚动视图一样，而宽度则是滚动视图的宽度乘以页视图的数量，页视图的尺寸则和滚动视图保持一致，对于垂直分页滚动来说容器视图的宽度和滚动视图一样，而高度则是滚动视图的高度乘以页视图的数量，页视图的尺寸则和滚动视图保持一致。每个页视图中在添加各自的条目视图。整体效果图如下：

![分页滚动UI布局](https://user-gold-cdn.xitu.io/2019/6/3/16b1ae827eac2e77?imageView2/0/w/1280/h/960/ignore-error/1)

### AutoLayout实现分页滚动的方法 ###

根据上面的UI结构这里用AutoLayout的代码来实现水平分页的滚动。这里的约束设置代码是iOS9以后提供的相关API。

` - (void)loadView { UIScrollView *scrollView = [[UIScrollView alloc] init]; if (@available(iOS 11.0, *)) { scrollView.contentInsetAdjustmentBehavior = UIScrollViewContentInsetAdjustmentNever; } else { // Fallback on earlier versions } scrollView.pagingEnabled = YES; scrollView.backgroundColor = [UIColor whiteColor]; self.view = scrollView; //建立容器视图 UIView *containerView = [UIView new]; containerView.translatesAutoresizingMaskIntoConstraints = NO; [scrollView addSubview:containerView]; //设置容器的四个边界和滚动视图保持一致的约束。 [containerView.leftAnchor constraintEqualToAnchor:scrollView.leftAnchor].active = YES; [containerView.topAnchor constraintEqualToAnchor:scrollView.topAnchor].active = YES; [containerView.rightAnchor constraintEqualToAnchor:scrollView.rightAnchor].active = YES; [containerView.bottomAnchor constraintEqualToAnchor:scrollView.bottomAnchor].active = YES; //容器视图的高度和滚动视图保持一致。 [containerView.heightAnchor constraintEqualToAnchor:scrollView.heightAnchor].active = YES; //添加页视图 NSArray<UIColor*> *colors = @[[UIColor redColor],[UIColor greenColor], [UIColor blueColor]]; NSMutableArray<UIView*> *pageViews = [NSMutableArray arrayWithCapacity:colors.count]; NSLayoutXAxisAnchor *prevLeftAnchor = containerView.leftAnchor; for (int i = 0; i < colors.count; i++) { //建立页视图 UIView *pageView = [UIView new]; pageView.backgroundColor = colors[i]; pageView.translatesAutoresizingMaskIntoConstraints = NO; [containerView addSubview:pageView]; //页视图分别从左往右排列，第1页的左边约束是容器视图的左边，其他页的左边约束则是前面兄弟视图的右边。 [pageView.leftAnchor constraintEqualToAnchor:prevLeftAnchor].active = YES; //每页的顶部约束是容器视图。 [pageView.topAnchor constraintEqualToAnchor:containerView.topAnchor].active = YES; //每页的宽度约束是滚动视图 [pageView.widthAnchor constraintEqualToAnchor:scrollView.widthAnchor].active = YES; //每页的高度约束是滚动视图 [pageView.heightAnchor constraintEqualToAnchor:scrollView.heightAnchor].active = YES; prevLeftAnchor = pageView.rightAnchor; [pageViews addObject:pageView]; } //关键的一步，如果需要左右滚动则将容器视图中的最右部子视图这里是B的右边边界依赖于容器视图的右边边界。 [pageViews.lastObject.rightAnchor constraintEqualToAnchor:containerView.rightAnchor].active = YES; //这里可以为每个页视图添加不同的条目视图，具体实现大家自行添加代码吧。 } 复制代码`

下面是运行时的效果图：

![分页滚动](https://user-gold-cdn.xitu.io/2019/6/3/16b1ae827f4d8bde?imageslim)

### MyLayout实现分页滚动的方法 ###

你也可以用 [MyLayout]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fyoungsoft%2FMyLinearLayout ) 布局库来实现分页滚动的能力。MyLayout布局库是笔者开源的一套功能强大的UI布局库。 您可以从github地址: [github.com/youngsoft/M…]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fyoungsoft%2FMyLinearLayout ) 下载或者从podfile中导入:

` pod 'MyLayout' 复制代码`

来使用MyLayout。下面是具体用MyLayout来实现分页滚动的代码。

` // #import <MyLayout.h> - (void)loadView { UIScrollView *scrollView = [[UIScrollView alloc] init]; if (@available(iOS 11.0, *)) { scrollView.contentInsetAdjustmentBehavior = UIScrollViewContentInsetAdjustmentNever; } else { // Fallback on earlier versions } scrollView.pagingEnabled = YES; scrollView.backgroundColor = [UIColor whiteColor]; self.view = scrollView; //建立一个水平线性布局容器视图 MyLinearLayout *containerView = [MyLinearLayout linearLayoutWithOrientation:MyOrientation_Horz]; containerView.myVertMargin = 0; //水平线性布局的上下边界和滚动视图保持一致，这里也会确定线性布局的高度。 containerView.gravity = MyGravity_Vert_Fill | MyGravity_Horz_Fill; //设置线性布局中的所有子视图均分和填充线性布局的高度和宽度。 [scrollView addSubview:containerView]; //添加页视图 NSArray<UIColor*> *colors = @[[UIColor redColor],[UIColor greenColor], [UIColor blueColor]]; NSMutableArray<UIView*> *pageViews = [NSMutableArray arrayWithCapacity:colors.count]; for (int i = 0; i < colors.count; i++) { //建立页视图 UIView *pageView = [UIView new]; pageView.backgroundColor = colors[i]; [containerView addSubview:pageView]; //因为线性布局通过属性gravity的设置就可以确定子页视图的高度和宽度，再加上线性布局的特性，所以页视图不需要设置任何附加的约束。 [pageViews addObject:pageView]; } //关键的一步, 设置线性布局的宽度是滚动视图的倍数 containerView.widthSize.equalTo(scrollView.widthSize).multiply(colors.count); //这里可以为每个页视图添加不同的条目视图，具体实现大家自行添加代码吧。 } 复制代码`

### MyLayout实现桌面的图标列表分页功能 ###

MyLayout中的流式布局MyFlowLayout所具备的能力和flex-box相似，甚至有些特性要强于后者。流式布局用于一些子视图有规律排列的场景，就比如本例子中的滚动分页的图标列表的能力。下面就是具体的实现代码。

` - (void)loadView { UIScrollView *scrollView = [[UIScrollView alloc] init]; if (@available(iOS 11.0, *)) { scrollView.contentInsetAdjustmentBehavior = UIScrollViewContentInsetAdjustmentNever; } else { // Fallback on earlier versions } scrollView.pagingEnabled = YES; scrollView.backgroundColor = [UIColor whiteColor]; self.view = scrollView; //建立一个垂直数量约束流式布局:每列展示3个子视图,每页展示9个子视图，整体从左往右滚动。 MyFlowLayout *containerView = [MyFlowLayout flowLayoutWithOrientation:MyOrientation_Vert arrangedCount:3]; containerView.pagedCount = 9; //pagedCount设置为非0时表示开始分页展示的功能，这里表示每页展示9个子视图，这个数量必须是arrangedCount的倍数。 containerView.wrapContentWidth = YES; //设置布局视图的宽度由子视图包裹，当垂直流式布局的这个属性设置为YES，并和pagedCount搭配使用会产生分页从左到右滚动的效果。 containerView.myVertMargin = 0; //容器视图的高度和滚动视图保持一致。 containerView.subviewHSpace = 10; containerView.subviewVSpace = 10; //设置子视图的水平和垂直间距。 containerView.padding = UIEdgeInsetsMake(5, 5, 5, 5); //布局视图的内边距设置。 [scrollView addSubview:containerView]; //建立条目视图 for (int i = 0; i < 40; i++) { UILabel *label = [UILabel new]; label.textAlignment = NSTextAlignmentCenter; label.backgroundColor = [UIColor greenColor]; label.text = [NSString stringWithFormat:@ "%d" ,i]; [containerView addSubview:label]; } //获取流式布局的横屏size classes，并且设置设备处于横屏时,每排数量由3个变为6个，每页的数量由9个变为18个。 MyFlowLayout *containerViewSC = [containerView fetchLayoutSizeClass:MySizeClass_Landscape copyFrom:MySizeClass_wAny | MySizeClass_hAny]; containerViewSC.arrangedCount = 6; containerViewSC.pagedCount = 18; 复制代码`

从上面的代码可以看出要实现分页滚动的图标列表的能力，主要是对充当容器视图的流式布局设置一些属性即可，不需要为条目设置任何约束，而且还支持横竖屏下每页的不同数量的展示能力。整个功能代码量少，对比用UICollectionView来实现相同的功能要简洁和容易得多。下面是程序运行的效果：

![分页图标效果图](https://user-gold-cdn.xitu.io/2019/6/3/16b1ae827f31d57c?imageslim)

### 横竖屏切换 ###

对于带有分页功能的滚动视图来说，当需要支持横竖屏时就有可能会出现横竖屏切换时界面停留在两个页面中间而不是按页进行滚动的效果。其原因是无论是分页滚动还是不分页滚动，在滚动时都是通过调整滚动视图的contentOffset来实现的。而当滚动视图进行横竖屏切换时不会调整对应的contentOffset值，这样就导致了在屏幕方向切换时的滚动位置出现异常。解决的办法就是在屏幕滚动时的相应回调处理方法中修正这个contentOffset的值来解决这个问题。比如我们可以在屏幕切换的sizeclass变化的视图控制器的协议方法中添加如下代码：

` - (void)traitCollectionDidChange:(nullable UITraitCollection *)previousTraitCollection { [super traitCollectionDidChange:previousTraitCollection]; UIScrollView *scrollView = (UIScrollView*)self.view; //根据当前的contentOffset调整到正确的contentOffset int pageIndex = scrollView.contentOffset.x / scrollView.frame.size.width; int pages = scrollView.contentSize.width / scrollView.frame.size.width; if (pageIndex >= pages) pageIndex = pages - 1; if (pageIndex < 0) pageIndex = 0; scrollView.contentOffset = CGPointMake(pageIndex * scrollView.frame.size.width, scrollView.contentOffset.y); } 复制代码`