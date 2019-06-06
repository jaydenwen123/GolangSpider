# React Native实现一个带筛选功能的搜房列表（3） #

原文链接 [React Native实现一个带筛选功能的搜房列表（3）]( https://link.juejin.im?target=https%3A%2F%2Fwww.neroxie.com%2F2019%2F06%2F06%2FReact-Native%25E5%25AE%259E%25E7%258E%25B0%25E4%25B8%2580%25E4%25B8%25AA%25E5%25B8%25A6%25E7%25AD%259B%25E9%2580%2589%25E5%258A%259F%25E8%2583%25BD%25E7%259A%2584%25E6%2590%259C%25E6%2588%25BF%25E5%2588%2597%25E8%25A1%25A8%25EF%25BC%25883%25EF%25BC%2589%2F )

在前两篇文章中已经介绍了如何实现一个支持下拉刷新和上拉加载更多的列表以及如何使用Redux进行单向数据流，那这篇就会介绍下最后一个模块筛选功能的开发即React Native与原生iOS的通信。

开始之前，还是先看一下最终实现的效果

![search_house](https://user-gold-cdn.xitu.io/2019/6/6/16b2ae02b72f2324?imageslim)

[代码传送门--NNHybrid]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2FYiHuaXie%2FNNHybrid )

关于如何进行React Native与原生iOS的通信，在官网中有很明确的教程，这里我就不细说了。我分享的主要是如何利用那些接口实现这样的一个效果。

在项目中，筛选条和对应的子菜单都是使用原生代码实现的，而列表使用js代码实现的。我们要做的就是将这些原生页面通过桥接的方式添加到js页面上。

桥接实现在 ` FHTFilterMenuManager` 这个类中，其实现如下：

* 原生实现

` // .h #import <React/RCTViewManager.h> #import "FHTFilterMenu.h" @interface FHTFilterMenu (RNBridge) @property (nonatomic, copy) RCTBubblingEventBlock onUpdateParameters; @property (nonatomic, copy) RCTBubblingEventBlock onChangeParameters; @end @interface RCTConvert (FHTFilterMenu) @end @interface FHTFilterMenuManager : RCTViewManager <RCTBridgeModule> @end // .m #import "FHTFilterMenuManager.h" #import <React/RCTUIManager.h> #import <objc/runtime.h> #import "FilterMenuRentTypeController.h" #import "FilterMenuGeographicController.h" #import "FilterMenuOrderByController.h" #import "FilterMenuMoreController.h" #import "FilterMenuRentalController.h" static ConstString kFilterParams = @ "filterParams" ; typedef NS_ENUM(NSInteger, FilterMenuType) { FilterMenuTypeNone, FilterMenuTypeEntireRent, //整租 FilterMenuTypeSharedRent, //合租 FilterMenuTypeApartment, //独栋公寓 FilterMenuTypeBelowThousand, //千元房源 FilterMenuTypePayMonthly, //月付 FilterMenuTypeVR, //VR }; @implementation RCTConvert (FHTFilterMenu) RCT_ENUM_CONVERTER(FilterMenuType, (@{@ "None" : @(FilterMenuTypeNone), @ "EntireRent" : @(FilterMenuTypeEntireRent), @ "SharedRent" : @(FilterMenuTypeSharedRent), @ "Apartment" : @(FilterMenuTypeApartment), @ "BelowThousand" : @(FilterMenuTypeBelowThousand), @ "PayMonthly" : @(FilterMenuTypePayMonthly), @ "VR" :@(FilterMenuTypeVR)}), FilterMenuTypeNone, integer Value); @end @implementation FHTFilterMenu (RNBridge) #pragma mark - Setter & Getter - (void) set OnUpdateParameters:(RCTBubblingEventBlock)onUpdateParameters { objc_setAssociatedObject(self, @selector(onUpdateParameters), onUpdateParameters, OBJC_ASSOCIATION_COPY_NONATOMIC); } - (RCTBubblingEventBlock)onUpdateParameters { return objc_getAssociatedObject(self, @selector(onUpdateParameters)); } - (void) set OnChangeParameters:(RCTBubblingEventBlock)onChangeParameters { objc_setAssociatedObject(self, @selector(onChangeParameters), onChangeParameters, OBJC_ASSOCIATION_COPY_NONATOMIC); } - (RCTBubblingEventBlock)onChangeParameters { return objc_getAssociatedObject(self, @selector(onChangeParameters)); } - (void) set FilterMenuType:(FilterMenuType)filterMenuType { switch (filterMenuType) { case FilterMenuTypeEntireRent: { UIViewController *vc = (UIViewController *)self.filterControllers[0]; [vc presetWithOptionTitles:@[@ "整租" ]]; } break ; case FilterMenuTypeSharedRent: { UIViewController *vc = (UIViewController *)self.filterControllers[0]; [vc presetWithOptionTitles:@[@ "合租" ]]; } break ; case FilterMenuTypeApartment: { UIViewController *vc = (UIViewController *)self.filterControllers[3]; [vc presetWithOptionTitles:@[@ "房源类型/独栋公寓" ]]; } break ; case FilterMenuTypeBelowThousand: { UIViewController *vc = (UIViewController *)self.filterControllers[2]; [vc presetWithOptionTitles:@[@ "1500以下" ]]; } break ; case FilterMenuTypePayMonthly: { UIViewController *vc = (UIViewController *)self.filterControllers[3]; [vc presetWithOptionTitles:@[@ "房源亮点/月付" ]]; } break ; case FilterMenuTypeVR: { UIViewController *vc = (UIViewController *)self.filterControllers[3]; [vc presetWithOptionTitles:@[@ "房源亮点/VR" ]]; } break ; default: break ; } }; @end @implementation FHTFilterMenuManager RCT_EXPORT_MODULE(); RCT_EXPORT_VIEW_PROPERTY(onUpdateParameters, RCTBubblingEventBlock); RCT_EXPORT_VIEW_PROPERTY(onChangeParameters, RCTBubblingEventBlock); RCT_EXPORT_VIEW_PROPERTY(filterMenuType, FilterMenuType); RCT_CUSTOM_VIEW_PROPERTY(cityId, NSString, FHTFilterMenu) { FilterMenuGeographicController *vc = view.filterControllers[1]; vc.cityId = (NSString *)json; }; RCT_CUSTOM_VIEW_PROPERTY(subwayData, NSArray, FHTFilterMenu) { FilterMenuGeographicController *vc = view.filterControllers[1]; vc.originalSubwayData = (NSArray *)json; }; RCT_EXPORT_METHOD(showFilterMenuOnView:(nonnull NSNumber *)containerTag filterMenuTag:(nonnull NSNumber *)filterMenuTag) { RCTUIManager *uiManager = self.bridge.uiManager; dispatch_async(uiManager.methodQueue, ^{ [uiManager addUIBlock:^(RCTUIManager *uiManager, NSDictionary<NSNumber *,UIView *> *viewRegistry) { UIView *view = viewRegistry[containerTag]; FHTFilterMenu *filterMenu = (FHTFilterMenu *)viewRegistry[filterMenuTag]; [filterMenu showFilterMenuOnView:view]; }]; }); } - (dispatch_queue_t)methodQueue { return dispatch_get_main_queue(); } - (UIView *)view { FilterMenuRentTypeController *rentTypeVC = [[FilterMenuRentTypeController alloc] initWithStyle:UITableViewStylePlain]; FilterMenuGeographicController *geographicVC = [FilterMenuGeographicController new]; FilterMenuRentalController *rentalVC = [[FilterMenuRentalController alloc] initWithStyle:UITableViewStylePlain]; FilterMenuMoreController *moreVC = [FilterMenuMoreController new]; FilterMenuOrderByController *orderByVC = [[FilterMenuOrderByController alloc] initWithStyle:UITableViewStylePlain]; CGRect frame = CGRectMake(0, FULL_NAVIGATION_BAR_HEIGHT, SCREEN_WIDTH, 44); FHTFilterMenu *filterMenu = [[FHTFilterMenu alloc] initWithFrame:frame]; filterMenu.filterControllers = @[rentTypeVC, geographicVC, rentalVC, moreVC, orderByVC]; [filterMenu dismissSubmenu:NO]; [filterMenu resetFilter]; [rentTypeVC presetWithOptionTitles:@[]]; __weak FHTFilterMenu *weakFilterMenu = filterMenu; rentTypeVC.didSetFilterHandler = ^(NSDictionary * _Nonnull params) { BLOCK_EXEC(weakFilterMenu.onUpdateParameters, @{kFilterParams: params}); }; geographicVC.didSetFilterHandler = ^(NSDictionary * _Nonnull params) { NSMutableDictionary *tmpParams = [@{@ "regionId" : nn_makeSureString(params[@ "regionId" ]), @ "zoneIds" : nn_makeSureArray(params[@ "zoneIds" ]), @ "subwayRouteId" : nn_makeSureString(params[@ "subwayRouteId" ]), @ "subwayStationCodes" : nn_makeSureArray(params[@ "subwayStationCodes" ])} mutableCopy]; BLOCK_EXEC(weakFilterMenu.onUpdateParameters, @{kFilterParams: [tmpParams copy]}); }; rentalVC.didSetFilterHandler = ^(NSDictionary * _Nonnull params) { NSDictionary *tmpParams = @{@ "minPrice" : nn_makeSureString(params[@ "minPrice" ]), @ "maxPrice" : nn_makeSureString(params[@ "maxPrice" ])}; BLOCK_EXEC(weakFilterMenu.onUpdateParameters, @{kFilterParams: tmpParams}); }; moreVC.didSetFilterHandler = ^(NSDictionary * _Nonnull params) { NSArray * type Array = params[@ "typeArray" ]; NSString * type = type Array.count == 1 ? ( type Array.lastObject)[@ "type" ] : @ "" ; NSDictionary *tmpParams = @{@ "roomAttributeTags" : params[@ "highlightArray" ], @ "chamberCounts" : params[@ "chamberArray" ], @ "type" : type }; BLOCK_EXEC(weakFilterMenu.onUpdateParameters, @{kFilterParams: tmpParams}); }; orderByVC.didSetFilterHandler = ^(NSDictionary * _Nonnull params) { BLOCK_EXEC(weakFilterMenu.onUpdateParameters, @{kFilterParams: params}); }; filterMenu.filterDidChangedHandler = ^(FHTFilterMenu * _Nonnull filterMenu, id<FHTFilterController> _Nonnull filterController) { BLOCK_EXEC(weakFilterMenu.onChangeParameters, nil); }; return filterMenu; } @end 复制代码`

* JS实现

` import React, { Component } from 'react' ; import { requireNativeComponent, NativeModules, findNodeHandle } from 'react-native' ; const FilterMenu = requireNativeComponent( 'FHTFilterMenu' , SearchFilterMenu); const filterMenuManager = NativeModules.FHTFilterMenuManager; export const FilterMenuType = { NONE: 'None' , ENTIRERENT: 'EntireRent' , SHAREDRENT: 'SharedRent' , APARTMENT: 'Apartment' , BELOWTHOUSAND: 'BelowThousand' , PAYMONTHLY: 'PayMonthly' , VR: 'VR' } export default class SearchFilterMenu extends Component { componentDidUpdate () { const filterMenuTag = findNodeHandle(this.refs.filterMenu); const containerTag = findNodeHandle(this.props.containerRef); if (filterMenuTag && containerTag) filterMenuManager.showFilterMenuOnView(containerTag, filterMenuTag); } render () { return <FilterMenu ref= 'filterMenu' {...this.props} />; } } 复制代码`

* JS调用

` <SearchFilterMenu style={styles.filterMenu} cityId={` ${home.cityId} `} subwayData={home.subwayData} containerRef={this.refs.container} filterMenuType={this.params.filterMenuType} onChangeParameters={() => this._loadData( true )} onUpdateParameters={({ nativeEvent: { filterParams } }) => { this.filterParams = { ...this.filterParams, ...filterParams, }; }} /> 复制代码`

## 创建ViewManager ##

` FHTFilterMenuManager` 是继承自 ` RCTViewManager` ，每一个原生UI都需要被一个 ` RCTViewManager` 的子类来创建和管理。在程序运行过程中， ` RCTViewManager` 会创建原生UI并把视图提供给 ` RCTUIManager` ， ` RCTUIManager` 则反过来委托 ` RCTViewManager` 在需要的时候去设置和更新视图的属性。这里有一个注意点： **ViewManager的命名格式是原生组件名字+Manager** 。

在ViewManager中最重要的是必须实现 ` - (UIView *)view` ，用来返回你想要桥接的原生UI。

## 从React Native传递属性到原生组件 ##

我们知道属性是最简单的跨组件通信，如果RN组件接受到一个属性的时候，可以通过 ` RCT_EXPORT_VIEW_PROPERTY` 和 ` RCT_CUSTOM_VIEW_PROPERTY` 的方式传递给原生组件。

` RCT_EXPORT_VIEW_PROPERTY` 可以将原生组件自带的属性暴露给JS。所以在 ` FHTFilterMenuManager` 中，我暴露三个自带的属性给JS，分别是

` // 点击确定按钮执行网络请求的block RCT_EXPORT_VIEW_PROPERTY(onUpdateParameters, RCTBubblingEventBlock); // 点击子菜单item，参数变更的block RCT_EXPORT_VIEW_PROPERTY(onChangeParameters, RCTBubblingEventBlock); // 筛选菜单的类型， RCT_EXPORT_VIEW_PROPERTY(filterMenuType, FilterMenuType); 复制代码`

这里我使用了分类的方式对原生组件的属性进行拓展，因为在项目中，我并不希望原生组件有太多RN桥接相关的代码，所以使用分类拓展，这样也可以对代码进行解耦。

这里有一个注意点： **如果自带的属性是block类型的话，属性名必须以on开头** 。

` RCT_CUSTOM_VIEW_PROPERTY` 可以让我们添加一些更为复杂的属性。由于 ` cityId` 和 ` subwayData` 是 ` FilterMenuGeographicController` 这个子菜单才需要的，如果把它们设置为 ` FHTFilterMenu` 的属性并不合适，而 ` FilterMenuGeographicController` 我们并没有暴露给JS，所以我们可以使用 ` RCT_CUSTOM_VIEW_PROPERTY` 对没有暴露给JS的对象进行属性传递。

` RCT_CUSTOM_VIEW_PROPERTY(cityId, NSString, FHTFilterMenu) { FilterMenuGeographicController *vc = view.filterControllers[1]; vc.cityId = (NSString *)json; }; RCT_CUSTOM_VIEW_PROPERTY(subwayData, NSArray, FHTFilterMenu) { FilterMenuGeographicController *vc = view.filterControllers[1]; vc.originalSubwayData = (NSArray *)json; }; 复制代码`

## React Native调用原生组件的方法 ##

` RCT_EXPORT_METHOD` 用来提供原生方法给JS调用。 ` RCT_EXPORT_METHOD(showFilterMenuOnView:(nonnull NSNumber *)containerTag filterMenuTag:(nonnull NSNumber *)filterMenuTag)` 用来实现将子菜单添加到 ` SearchHousePage` ，其实现如下：

` RCT_EXPORT_METHOD(showFilterMenuOnView:(nonnull NSNumber *)containerTag filterMenuTag:(nonnull NSNumber *)filterMenuTag) { RCTUIManager *uiManager = self.bridge.uiManager; dispatch_async(uiManager.methodQueue, ^{ [uiManager addUIBlock:^(RCTUIManager *uiManager, NSDictionary<NSNumber *,UIView *> *viewRegistry) { UIView *view = viewRegistry[containerTag]; FHTFilterMenu *filterMenu = (FHTFilterMenu *)viewRegistry[filterMenuTag]; [filterMenu showFilterMenuOnView:view]; }]; }); } 复制代码`

其中 ` containerTag` 代码用来表示 ` SearchHousePage` ， ` filterMenuTag` 则表示筛选条。这里有一个注意点： **不能使用self.view的方式因为会创建出一个新的FHTFilterMenu对象，更不能在 ` - (UIView *)view` 中使用一个指针对创建出来的View进行引用，如果你的组件在多个页面都用使用的话，是会出问题的** 。

另外，由于iOS的UI操作是要放在主线程完成的，所以最好 ` methodQueue` 指定为主线程。

## 结尾 ##

到这里整个 ` SearchHousePage` 页面的开发已经完成，如果需要查看完整代码的话，在 [代码传送门--NNHybrid]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2FYiHuaXie%2FNNHybrid ) 中。

参考： [Native Modules]( https://link.juejin.im?target=https%3A%2F%2Ffacebook.github.io%2Freact-native%2Fdocs%2Fnative-modules-ios ) [Native UI Components]( https://link.juejin.im?target=https%3A%2F%2Ffacebook.github.io%2Freact-native%2Fdocs%2Fnative-components-ios )