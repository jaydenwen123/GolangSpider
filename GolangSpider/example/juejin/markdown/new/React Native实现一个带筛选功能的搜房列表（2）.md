# React Native实现一个带筛选功能的搜房列表（2） #

原文链接 [React Native实现一个带筛选功能的搜房列表（2）]( https://link.juejin.im?target=https%3A%2F%2Fwww.neroxie.com%2F2019%2F06%2F06%2FReact-Native%25E5%25AE%259E%25E7%258E%25B0%25E4%25B8%2580%25E4%25B8%25AA%25E5%25B8%25A6%25E7%25AD%259B%25E9%2580%2589%25E5%258A%259F%25E8%2583%25BD%25E7%259A%2584%25E6%2590%259C%25E6%2588%25BF%25E5%2588%2597%25E8%25A1%25A8%25EF%25BC%25882%25EF%25BC%2589%2F )

在 [上一篇]( https://link.juejin.im?target=https%3A%2F%2Fwww.neroxie.com%2F2019%2F06%2F06%2FReact-Native%25E5%25AE%259E%25E7%258E%25B0%25E4%25B8%2580%25E4%25B8%25AA%25E5%25B8%25A6%25E7%25AD%259B%25E9%2580%2589%25E5%258A%259F%25E8%2583%25BD%25E7%259A%2584%25E6%2590%259C%25E6%2588%25BF%25E5%2588%2597%25E8%25A1%25A8%25EF%25BC%25881%25EF%25BC%2589%2F ) 中，我们实现了一个下拉刷新和上拉加载更多的列表，那根据一般的开发步骤，接着应该就是进行网络请求，在网络请求之后更新列表数据和列表的刷新状态。

这篇文章会向大家介绍一下Redux的基本概念以及在页面中如何使用Redux进行状态管理。

文章中的代码都来自 [代码传送门--NNHybrid]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2FYiHuaXie%2FNNHybrid ) 。开始之前，我们先看一下最终实现的效果

![search_house](https://user-gold-cdn.xitu.io/2019/6/6/16b2ae02b72f2324?imageslim)

## Redux概念 ##

首先先简单介绍一下Redux的一些概念。Redux是JavaScript状态容器，提供可预测化的状态管理，其工作流程如下：

整个工作流程为：

![reduxProcess](https://user-gold-cdn.xitu.io/2019/6/6/16b2ae29623f19be?imageView2/0/w/1280/h/960/ignore-error/1)

* View需要订阅Store中的state；
* 操作View（点击了View上的一个按钮或者执行一个网络请求），发出Action；
* Store自动调用Reducer，并且传入两个参数(Old State和Action)，Reducer会返回新的State，如果有Middleware，Store会将Old State和Action传递给Middleware，Middleware会调用Reducer 然后返回新的State；
* State一旦有变化，Store就会调用监听函数，来更新View；

### Store ###

Store是存储state的容器，负责提供所有的状态。整个应用只能有一个Store，这么做的目的是为了让组件之间的通信更加简单。

![reduxCommunication](https://user-gold-cdn.xitu.io/2019/6/6/16b2ae296301dee4?imageView2/0/w/1280/h/960/ignore-error/1)

在没有Store的情况下，组件之间需要通信就比较麻烦，如果一个父组件要将状态传递到子组件，就需要通过props一层一层往下传，一个子组件的状态发生改变并且要让父组件知道，则必须暴露一个事件出去才能通信。这就使得组件之间通信依赖于组件的层次结构。此时如果有两个平级的节点想要通信，就需要通过它们的父组件进行中转。 有了这个全局的Store之后，所有的组件变成了和Store进行通信。这样组件之间通信就会变少，当Store发生变化，对应的组件也能拿到相关的数据。当组件内部有时间触发Store的变化时，更新Store即可。这也就是所谓的单向数据流过程。

Store的职责如下：

* 维持应用的state；
* 提供 ` getState()` 方法获取state；
* 提供 ` dispatch(action)` 方法更新state；
* 通过 ` subscribe(listener)` 注册监听器；
* 通过 ` subscribe(listener)` 返回的函数注销监听器。

### Action ###

当我们想要更改store中的state时，我们便需要使用Action。Action是Store数据的唯一来源，每一次修改state便要发起一次Action。

Action可以理解为是一个Javascript对象。其内部必须包含一个 ` type` 字段来表示将要执行的动作，除了 ` type` 字段外，Action的结构完全由自己决定。多数情况下， ` type` 字段会被定义成字符串常量。

Action举例：

` { type : Types.SEARCH_HOUSE_LOAD_DATA_SUCCESS, currentPage: ++currentPage, houseList, hasMoreData, } 复制代码`

### Action创建函数 ###

Action创建函数就是生成action的方法。“action” 和 “action 创建函数” 这两个概念很容易混在一起，使用时最好注意区分。

Action创建函数举例：

` export function init(storeName) { return dispatch => { dispatch({ type : Types.HOUSE_DETAIL_INIT, storeName }); } } 复制代码`

### Reducer ###

Store收到Action以后，必须给出一个新的State，这样View才会发生变化。 这种State的计算过程就叫做Reducer。Reducer是一个纯函数，它只接受Action和当前State作为参数，返回一个新的State。

由于Reducer是一个纯函数，所以我们不能在reducer里执行以下操作：

* 修改传入的参数；
* 执行有副作用的操作；
* 调用非纯函数；
* 不要修改state；
* 遇到未知的action时，一定要返回旧的state；

Reducer举例：

` const defaultState = { locationCityName: '' , visitedCities: [], hotCities: [], sectionCityData: [], sectionTitles: [] }; export function cityListReducer(state = defaultState, action) { switch (action.type) { case Types.CITY_LIST_LOAD_DATA: return { ...state, visitedCities: action.visitedCities, hotCities: action.hotCities, sectionCityData: action.sectionCityData, sectionTitles: action.sectionTitles, } case Types.CITY_LIST_START_LOCATION: case Types.CITY_LIST_LOCATION_FINISHED: return { ...state, locationCityName: action.locationCityName }; default: return state; } } 复制代码`

### 拆分与合并reducer ###

在开发过程中，由于有的功能是相互独立的，所以我们需要拆分reducer。一般情况下，针对一个页面可以设置一个reducer。但redux原则是只允许一个根reducer，接下来我们需要将每个页面的的reducer聚合到一个根reducer中。

合并reducer代码如下：

` const appReducers = combineReducers({ nav: navReducer, home: homeReducer, cityList: cityListReducer, apartments: apartmentReducer, houseDetails: houseDetailReducer, searchHouse: searchHouseReducer, }); export default (state, action) => { switch (action.type) { case Types.APARTMENT_WILL_UNMOUNT: delete state.apartments[action.storeName]; break ; case Types.HOUSE_DETAIL_WILL_UNMOUNT: delete state.houseDetails[action.storeName]; break ; case Types.SEARCH_HOUSE_WILL_UNMOUNT: delete state.searchHouse; break ; } return appReducers(state, action); } 复制代码`

## SearchHousePage使用Redux ##

### Action类型定义 ###

` SEARCH_HOUSE_LOAD_DATA: 'SEARCH_HOUSE_LOAD_DATA' , SEARCH_HOUSE_LOAD_MORE_DATA: 'SEARCH_HOUSE_LOAD_MORE_DATA' , SEARCH_HOUSE_LOAD_DATA_SUCCESS: 'SEARCH_HOUSE_LOAD_DATA_SUCCESS' , SEARCH_HOUSE_LOAD_DATA_FAIL: 'SEARCH_HOUSE_LOAD_DATA_FAIL' , SEARCH_HOUSE_WILL_UNMOUNT: 'SEARCH_HOUSE_WILL_UNMOUNT' , 复制代码`

### Action创建函数 ###

` export function loadData(params, currentPage, errorCallBack) { return dispatch => { dispatch({ type : currentPage == 1 ? Types.SEARCH_HOUSE_LOAD_DATA : Types.SEARCH_HOUSE_LOAD_MORE_DATA }); set Timeout(() => { Network .my_request({ apiPath: ApiPath.SEARCH, apiMethod: 'searchByPage' , apiVersion: '1.0' , params: { ...params, pageNo: currentPage, pageSize: 10 } }) .then(response => { const tmpResponse = AppUtil.makeSureObject(response); const hasMoreData = currentPage < tmpResponse.totalPages; const houseList = AppUtil.makeSureArray(tmpResponse.resultList); dispatch({ type : Types.SEARCH_HOUSE_LOAD_DATA_SUCCESS, currentPage: ++currentPage, houseList, hasMoreData, }); }) .catch(error => { if (errorCallBack) errorCallBack(error.message); const action = { type : Types.SEARCH_HOUSE_LOAD_DATA_FAIL }; if (currentPage == 1) { action.houseList = [] action.currentPage = 1; }; dispatch(action); }); }, 300); } } 复制代码`

### 创建reducer ###

` // 默认的state const defaultState = { houseList: [], headerIsRefreshing: false , footerRefreshState: FooterRefreshState.Idle, currentPage: 1, } export function searchHouseReducer(state = defaultState, action) { switch (action.type) { case Types.SEARCH_HOUSE_LOAD_DATA: { return { ...state, headerIsRefreshing: true } } case Types.SEARCH_HOUSE_LOAD_MORE_DATA: { return { ...state, footerRefreshState: FooterRefreshState.Refreshing, } } case Types.SEARCH_HOUSE_LOAD_DATA_FAIL: { return { ...state, headerIsRefreshing: false , footerRefreshState: FooterRefreshState.Failure, houseList: action.houseList ? action.houseList : state.houseList, currentPage: action.currentPage, } } case Types.SEARCH_HOUSE_LOAD_DATA_SUCCESS: { const houseList = action.currentPage <= 2 ? action.houseList : state.houseList.concat(action.houseList); let footerRefreshState = FooterRefreshState.Idle; if (AppUtil.isEmptyArray(houseList)) { footerRefreshState = FooterRefreshState.EmptyData; } else if (!action.hasMoreData) { footerRefreshState = FooterRefreshState.NoMoreData; } return { ...state, houseList, currentPage: action.currentPage, headerIsRefreshing: false , footerRefreshState, } } default: return state; } } 复制代码`

### 包装组件 ###

` class SearchHousePage extends Component { // ...代码省略 componentDidMount () { this._loadData( true ); } componentWillUnmount () { NavigationUtil.dispatch(Types.SEARCH_HOUSE_WILL_UNMOUNT); } _loadData(isRefresh) { const { loadData, searchHouse } = this.props; const currentPage = isRefresh ? 1 : searchHouse.currentPage; loadData(this.filterParams, currentPage, error => Toaster.autoDisapperShow(error)); } render () { const { home, searchHouse } = this.props; return ( <View style={styles.container} ref= 'container' > <RefreshFlatList ref= 'flatList' style={{ marginTop: AppUtil.fullNavigationBarHeight + 44 }} showsHorizontalScrollIndicator={ false } data={searchHouse.houseList} keyExtractor={item => ` ${item.id} `} renderItem={({ item, index }) => this._renderHouseCell(item, index)} headerIsRefreshing={searchHouse.headerIsRefreshing} footerRefreshState={searchHouse.footerRefreshState} onHeaderRefresh={() => this._loadData( true )} onFooterRefresh={() => this._loadData( false )} footerRefreshComponent={footerRefreshState => this.footerRefreshComponent(footerRefreshState, searchHouse.houseList)} /> <NavigationBar navBarStyle={{ position: 'absolute' }} backOrCloseHandler={() => NavigationUtil.goBack()} title= '搜房' /> <SearchFilterMenu style={styles.filterMenu} cityId={` ${home.cityId} `} subwayData={home.subwayData} containerRef={this.refs.container} filterMenuType={this.params.filterMenuType} onChangeParameters={() => this._loadData( true )} onUpdateParameters={({ nativeEvent: { filterParams } }) => { this.filterParams = { ...this.filterParams, ...filterParams, }; }} /> </View> ); } } const mapStateToProps = state => ({ home: state.home, searchHouse: state.searchHouse }); const mapDispatchToProps = dispatch => ({ loadData: (params, currentPage, errorCallBack) => dispatch(loadData(params, currentPage, errorCallBack)), }); export default connect(mapStateToProps, mapDispatchToProps)(SearchHousePage); 复制代码`

从上面的代码使用了一个 ` connect` 函数， ` connect` 连接React组件与Redux store，连接操作会返回一个新的与Redux store连接的组件类，并且连接操作不会改变原来的组件类。

` mapStateToProps` 中订阅了home节点和searchHouse节点，该页面主要使用searchHouse节点，那订阅home节点是用来方便组件间通信，这样页面进行网络请求所需的cityId，就不需要从前以页面传入，也不需要从缓存中读取。

列表的刷新状态由 ` headerIsRefreshing` 和 ` footerRefreshState` 进行管理。

## 综上 ##

redux已经帮我们完成了页面的状态管理，再总结一下Redux需要注意的点：

* Redux应用只有一个单一的Store。当需要拆分数据处理逻辑时，你应该使用拆分与合并reducer而不是创建多个Store；
* redux一个特点是：状态共享，所有的状态都放在一个Store中，任何组件都可以订阅Store中的数据，但是不建议组件订阅过多Store中的节点；
* 不要将所有的State都适合放在Store中，这样会让Store变得非常庞大；

到这里，我们实现了列表的下拉刷新、加载更多以及如何使用redux，还差一个筛选栏和子菜单页面的开发，这里涉及到React Native与原生之间的通信，我会在 [React Native实现一个带筛选功能的搜房列表（3）]( https://link.juejin.im?target=https%3A%2F%2Fwww.neroxie.com%2F2019%2F06%2F06%2FReact-Native%25E5%25AE%259E%25E7%258E%25B0%25E4%25B8%2580%25E4%25B8%25AA%25E5%25B8%25A6%25E7%25AD%259B%25E9%2580%2589%25E5%258A%259F%25E8%2583%25BD%25E7%259A%2584%25E6%2590%259C%25E6%2588%25BF%25E5%2588%2597%25E8%25A1%25A8%25EF%25BC%25883%25EF%25BC%2589%2F ) 中分享下如何进行React Native与原生的桥接开发。

另外上面提供的代码均是从项目当中截取的，如果需要查看完整代码的话，在 [代码传送门--NNHybrid]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2FYiHuaXie%2FNNHybrid ) 中。

上述相关代码路径：

` redux文件夹: /NNHybridRN/redux SearchHousePage: /NNHybridRN/sections/searchHouse/SearchHousePage.js 复制代码`

参考资料：

[Redux 中文文档]( https://link.juejin.im?target=https%3A%2F%2Fwww.redux.org.cn%2F )