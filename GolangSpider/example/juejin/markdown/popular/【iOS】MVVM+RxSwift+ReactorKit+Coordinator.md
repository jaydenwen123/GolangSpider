# 【iOS】MVVM+RxSwift+ReactorKit+Coordinator #

### MVVM + RxSwift ###

**iOS** 中的 **MVVM** 架构早就是个老生常谈的问题，相比于传统的 **MVC** 架构方式， **MVVM** 比较核心的地方在于双向绑定的过程，即 **View** 和 **ViewModel** 之间的绑定，而建立绑定关系最优方案是通过响应式的方式构建， **iOS** 原生方面可以通过 **KVO** + **KVC** 的方式去搭建响应式，缺点是API相对复杂，操作不方便，纯 **Swift** 对象需要标记为 ` dynamic` ，需要手动管理 **KVO** 的生命周期。

**RxSwift** 属于 **ReactiveX** 系列，目前存在多个语言版本，基本覆盖全部主流编程语言，其专注于异步编程与控制可观察数据(或者事件)流的API，背后是微软的团队在开发维护，所以稳定性较高。 **RxSwift** 是一种响应式的编程思想，故非常适合与 **MVVM** 架构配合使用。

### ViewModel + ReactorKit ###

**MVVM** 架构最核心的部分无疑是 **ViewModel** ，其主要负责模块的逻辑处理、状态的维护等。在 **iOS** 开发中，状态一词提及的相对较少，不像 **React** 中组件对状态的依赖那么强烈。其实每一个具有交互功能的控件都会依赖于状态，状态决定控件的表示方式，故状态在 **iOS** 开发中同样重要。传统的 **MVC** 方式，状态的管理主要是 **Controller** 负责，在 **MVVM** 中由 **ViewModel** 管理。

**ReactorKit** 是一个轻量级的响应式框架，其依赖于 **RxSwift** 并结合了 **Flux** 。 **Flux** 是 faceBook 提出的一种架构思想，其核心概念是 **数据的单向流动** ， 同样适用于 **iOS** ，在 **iOS** 中主要表现为 **Action** 和 **State** 两个部分：

![ReactorKit](https://user-gold-cdn.xitu.io/2019/6/2/16b1875546641022?imageView2/0/w/1280/h/960/ignore-error/1)

> 
> 
> 
> **View** 发出的 **Action** ，经由 **Reactor** 处理后，由 **State** 抛出后绑定到 **View** 上，即每一个状态的改变都要派发一个
> **Action** 。也就是说 **View** 以怎样的方式显示是被动的，如想改变自身的渲染方式需要自己派发 **Action** 。
> 
> 

* 

#### View ####

**ReactorKit** 将 **Controller** 和 **View** 都归类为 ` View` ，使用方式是需要实现 ` View` 协议：

` class ReactorViewController : UIViewController , View { ... } 复制代码`
* 

#### Reactor + State ####

可以理解为 **Reactor** 就是 **ViewModel** 。 **Reactor** 同样是协议，其限定了 **ViewModel** 的行为（代码片段来自网络）：

` class ReactorViewModel : Reactor { /// - Action: View 派发的事件 enum Action { case refreshFollowingStatus( Int ) case follow( Int ) } /// - Mutation：Action 和 State 之间的过渡 enum Mutation { case setFollowing( Bool ) } /// - State：状态管理器 struct State { var isFollowing: Bool = false } /// - 状态管理器初始化 let initialState: State = State () } 复制代码` ` func mutate (action: Action) -> Observable < Mutation > { switch action { /// - View 派发的 Action 会在这里被响应 case let.refreshFollowingStatus(userID): return UserAPI.isFollowing(userID) // create an API stream. map { (isFollowing: Bool ) -> Mutation in /// - 派发一个 Mutation return Mutation.setFollowing(isFollowing) } /// - 同理 case let.follow(userID): return UserAPI.follow() . map { _ -> Mutation in return Mutation.setFollowing( true ) } } func reduce (state: State, mutation: Mutation) -> State { /// - 对 State 做一份拷贝，因为 State 是 let 声明的结构体 var state = state switch mutation { /// - 将 Mutation 所关联的数据映射到 State 上 case let.setFollowing(isFollowing): /// - 改变 State state.isFollowing = isFollowing /// - 返回一个新的 State return state } 复制代码`

以上是 **ViewModel** 的大致工作流程： **Action** -> **Mutatuin** -> **State** ，还有一些可选的API，如 **transform()** ，可翻阅官方文档查看。通过以上代码可知， **ReactorKit** 符合 **Flux** 编程思想，简单来说 **State** 改变需通过 **Action** 。

* 

#### 完善 View ####

众所周知， **MVVM** 中 **Controller** 需持有 **ViewModel** ，同样 **ReactorKit** 中的 **View** 协议规定需显示的指定 **Reactor** 的类型，并提供了 ` bind()` 方法，可以在这个方法中建立 **View** 和 **ViewModel** 之间的绑定关系。

` class ReactorViewController : UIViewController , View { func bind (reactor: ReactorViewModel) { /// - View 发出的 Action 绑定到了 ViewModel(Reactor) 的 action 上 refreshButton.rx.tap. map { Reactor. Action.refresh } .bindTo(reactor.action) .addDisposableTo( self.disposeBag) /// - ViewModel(Reactor) 的 State 绑定到了 View 上，并立即根据 State 渲染 reactor.state. map { $ 0.isFollowing } .bindTo(followButton.rx.isSelected) .addDisposableTo( self.disposeBag) } } 复制代码`

有了 **ReactorKit** 的协助， **ViewModel** 的行为更清晰明了。

### Coordinator ###

**Coordinator** 导航层。传统的开发模式下，页面间的跳转是通过 ` navigationController` 的 ` push()` 方法，这种方法固然便捷，但是实现跳转存在页面间耦合。 **Coordinator** 的诞生就是为了解决这一问题，当然 **路由** 或者引入 **中间管理层** 也可以实现解耦，但 **Coordintaor** 更轻量。

引入 **Coordinator** 后跳转逻辑对页面不可见，由 **Coordinator** 管理，其提供了 ` navigationController` 的接口并持有 **Controller** ，跳转逻辑隐藏在了 **Coordinator** 中。 **Coordinator** 独立与 **MVVM** 之外，是一个附加层，不依赖于 **MVVM** 中的任何一部分。可以理解为 **Coordinator** 是每个组件对外暴露的接口，固然页面间的交互，只能通过 **Coordinator** ，同样依赖于 **RxSwift** 。

* 实现 **Coordinator**

` /// - 页面 Pop 后触发的事件 enum PopResult { case reload case cancel } final class ExampleCoordinator : BaseCoordinator < PopResult > { override func start () -> Observable < PopResult > { let controller = ReactorViewController () let reactor = ReactorViewModel () let cancel = controller.popedAction .asObserver() . map { PopResult.cancel } let delete = reactor.state . map { $ 0.isDelete } . filter { (bool) -> Bool in return bool } . map { _ in PopResult.reload } return Observable.merge(cancel, delete) .take( 1 ). do (onNext: { [ weak self ] (result) in if let ` self ` = self { switch result { case.reload: self.navigationController.popViewController(animated: true ) break default : break } } }) } } 复制代码`

* Coordinator 交互（页面跳转）

` self.coordinate(to: ExampleCoordinator ()) .subscribe(onNext: { result in switch result { case reload: ... case cancel: ... } }) .disposed(by: self.disposeBag) 复制代码`

* 

Coordinator 源码同样很简单

` public protocol CoordinatorType : NSObjectProtocol { var identifier: UUID { get } var childCoordinators: [ UUID : Any ] { get set } var navigationController: RTRootNavigationController ! { get set } } public extension CoordinatorType { func store <T> (coordinator: BaseCoordinator<T>) { coordinator.navigationController = navigationController childCoordinators[coordinator.identifier] = coordinator } func free <T> (coordinator: BaseCoordinator<T>) { childCoordinators[coordinator.identifier] = nil } @discardableResult public func coordinate <T> (to coordinator: BaseCoordinator<T>) -> Observable < T > { store(coordinator: coordinator) return coordinator.start() . do (onNext: { [ weak self ] _ in if let ` self ` = self { self.free(coordinator: coordinator) } }) } } public class BaseCoordinator < ResultType >: NSObject , UINavigationControllerDelegate , CoordinatorType { /// Typealias which will allows to access a ResultType of the Coordainator by `CoordinatorName.CoordinationResult`. typealias CoordinationResult = ResultType public var navigationController: RTRootNavigationController ! func start () -> Observable < ResultType > { fatalError ( "Coordinator Start method should be implemented by subclass." ) } func noResultStart () { fatalError ( "Coordinator noResultStart method should be implemented by subclass." ) } /// UINavigationControllerDelegate public func navigationController ( _ navigationController: UINavigationController, didShow viewController: UIViewController, animated: Bool) { // ensure the view controller is popping if let transitionCoordinator = navigationController.transitionCoordinator, let fromViewController = transitionCoordinator.viewController(forKey: .from), !navigationController.viewControllers. contains (fromViewController) { fromViewController.popedAction.onNext(()) fromViewController.popedAction.onCompleted() } } let disposeBag = DisposeBag () public let identifier = UUID () public var childCoordinators = [ UUID : Any ]() } 复制代码`

### 完！ ###