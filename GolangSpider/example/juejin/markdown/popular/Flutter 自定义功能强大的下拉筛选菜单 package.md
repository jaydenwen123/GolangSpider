# Flutter 自定义功能强大的下拉筛选菜单 package #

# gzx_dropdown_menu #

自定义功能强大的下拉筛选菜单flutter package

* Custom dropdown header
* Custom dropdown header item
* Custom dropdown menu
* Custom dropdown menu show animation time
* Control dropdown menu show or hide

#### 如果对您有帮助，麻烦给个Star，您的支持是我持续更新的动力。 ####

## 导航 ##

* [Gif效果图]( #Gif%E6%95%88%E6%9E%9C%E5%9B%BE )
* [如何使用]( #%E5%A6%82%E4%BD%95%E4%BD%BF%E7%94%A8 )

## Gif效果图 ##

分别是仿美团和淘宝的效果图
美团的代码就在这个仓库的example目录下
淘宝的代码在 [Flutter 淘宝，点我打开]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2FGanZhiXiong%2FGZXTaoBaoAppFlutter )

![](https://user-gold-cdn.xitu.io/2019/6/3/16b1cdc41baf199a?imageslim) ![](https://user-gold-cdn.xitu.io/2019/6/3/16b1cdc87142fe7c?imageslim)

## 如何使用 ##

由于最近被qiang，所以没有发布到Pub，后续会发布到Pub

### 1、添加gzx_dropdown_menu package ###

打开pubspec.yaml文件 添加如下代码

` gzx_dropdown_menu : git: url: https://github.com/GanZhiXiong/gzx_dropdown_menu.git 复制代码`

添加位置如下图所示

![](https://user-gold-cdn.xitu.io/2019/6/3/16b1cde777abbc34?imageView2/0/w/1280/h/960/ignore-error/1)

添加后打开Terminal，执行flutter packages get

### 2、使用 ###

打开本仓库example项目下的gzx_dropdown_menu_test_page.dart文件自己看。

没空编辑文字了，而且说这么多还不如你直接运行下看下效果，然后看下代码，就知道如何使用了。

**算了 🤪🤪🤪🙄还是简单说下吧！！！**
你只需要将GZXDropDownHeader和GZXDropDownMenu嵌套到你的代码中即可

#### GZXDropDownHeader ####

**这里要注意了，这些参数不是必须要要写的，我写出来只是让你知道强大的自定义功能，实际上就前面三个参数是必填的**

` // 下拉菜单头部 GZXDropDownHeader( // 下拉的头部项，目前每一项，只能自定义显示的文字、图标、图标大小修改 items: [ GZXDropDownHeaderItem(_dropDownHeaderItemStrings[0]), GZXDropDownHeaderItem(_dropDownHeaderItemStrings[1]), GZXDropDownHeaderItem(_dropDownHeaderItemStrings[2]), GZXDropDownHeaderItem(_dropDownHeaderItemStrings[3], iconData: Icons.filter_frames, iconSize: 18), ], // GZXDropDownHeader对应第一父级Stack的key stackKey: _stackKey, // controller用于控制menu的显示或隐藏 controller: _dropdownMenuController, // 当点击头部项的事件，在这里可以进行页面跳转或openEndDrawer onItemTap: (index) { if (index == 3) { _scaffoldKey.currentState.openEndDrawer(); _dropdownMenuController.hide(); } }, // 头部的高度 height: 40, // 头部背景颜色 color: Colors.red, // 头部边框宽度 borderWidth: 1, // 头部边框颜色 borderColor: Color(0xFFeeede6), // 分割线高度 dividerHeight: 20, // 分割线颜色 dividerColor: Color(0xFFeeede6), // 文字样式 style: TextStyle(color: Color(0xFF666666), fontSize: 13), // 下拉时文字样式 dropDownStyle: TextStyle( fontSize: 13, color: Theme.of(context).primaryColor, ), // 图标大小 iconSize: 20, // 图标颜色 iconColor: Color(0xFFafada7), // 下拉时图标颜色 iconDropDownColor: Theme.of(context).primaryColor, ), 复制代码`

#### GZXDropDownMenu ####

` // 下拉菜单 GZXDropDownMenu( // controller用于控制menu的显示或隐藏 controller: _dropdownMenuController, // 下拉菜单显示或隐藏动画时长 animationMilliseconds: 500, // 下拉菜单，高度自定义，你想显示什么就显示什么，完全由你决定，你只需要在选择后调用_dropdownMenuController.hide();即可 menus: [ GZXDropdownMenuBuilder( dropDownHeight: 40 * 8.0, dropDownWidget: _buildQuanChengWidget((selectValue) { _dropDownHeaderItemStrings[0] = selectValue; _dropdownMenuController.hide(); set State(() {}); })), GZXDropdownMenuBuilder( dropDownHeight: 40 * 8.0, dropDownWidget: _buildConditionListWidget(_brandSortConditions, (value) { _selectBrandSortCondition = value; _dropDownHeaderItemStrings[1] = _selectBrandSortCondition.name == '全部' ? '品牌' : _selectBrandSortCondition.name; _dropdownMenuController.hide(); set State(() {}); })), GZXDropdownMenuBuilder( dropDownHeight: 40.0 * _distanceSortConditions.length, dropDownWidget: _buildConditionListWidget(_distanceSortConditions, (value) { _dropDownHeaderItemStrings[2] = _selectDistanceSortCondition.name; _selectDistanceSortCondition = value; _dropdownMenuController.hide(); set State(() {}); })), ], ) 复制代码`