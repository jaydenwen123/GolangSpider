# Flutter学习篇(三)—— MobX的使用和原理 #

### 导航 ###

* [Flutter学习篇(一)—— Dialog的简单使用]( https://juejin.im/post/5cf236d8f265da1bc23f5fbf )
* [Flutter学习篇(二)—— Drawer和水纹按压效果]( https://juejin.im/post/5cf3a503e51d4555fd20a2e2 )

## 前言 ##

[MobX]( https://link.juejin.im?target=https%3A%2F%2Fcn.mobx.js.org%2F ) 是前端一个很流行的函数响应式编程，让状态变得简单可扩展。背后的哲学是：

> 
> 
> 
> 任何源自应用状态的东西都应该自动地获得
> 
> 

基于观察者的MVVM框架完成了数据到UI的双向绑定。Google2017年也发布了类似思想的MVVM框架ViewModel。MVVM是数据驱动更新的框架，可以很方便地把页面和逻辑抽开，在前端很受欢迎。所以MobX也出了dart的版本用来支持Flutter的使用。下面我们就开始动手在Flutter引入MobX。

## 使用 ##

先放出 [官网]( https://link.juejin.im?target=https%3A%2F%2Fmobx.pub%2F ) ，使用分几步走：

### 1. 首先，引入依赖： ###

` mobx: ^0.2.0 flutter_mobx: ^0.2.0 mobx_codegen: ^0.2.0 复制代码`

### 2. 添加一个store： ###

` import 'package:mobx/mobx.dart' ; import 'package:shared_preferences/shared_preferences.dart' ; // 自动生成的类 part 'settings_store.g.dart' ; class SettingsStore = _SettingsStore with _ $ SettingsStore ; abstract class _SettingsStore implements Store { var key = { "showPage" : "showPage" , }; @observable String showPage = "" ; @action getPrefsData() async { SharedPreferences prefs = await SharedPreferences.getInstance(); showPage = prefs. get (key[ "showPage" ]) ?? "首页" ; } @action saveShowPage( String showPage) async { if (showPage == null ) { return ; } this.showPage = showPage; SharedPreferences prefs = await SharedPreferences.getInstance(); prefs.setString(key[ "showPage" ], showPage); } } 复制代码`

对于dart版本的mobx，是通过生成新的类来实现双向绑定的效果，所以需要在store里面加上生成类的一些定义:

` part 'settings_store.g.dart' ; class SettingsStore = _SettingsStore with _ $ SettingsStore ; 复制代码`

_$SettingsStore是待生成的类，SettingsStore则是混合了两个store的新类。如下是自动生成的类：

` // GENERATED CODE - DO NOT MODIFY BY HAND part of 'settings_store.dart' ; // ************************************************************************** // StoreGenerator // ************************************************************************** // ignore_for_file: non_constant_identifier_names, unnecessary_lambdas, prefer_expression_function_bodies, lines_longer_than_80_chars mixin _$SettingsStore on _SettingsStore, Store { final _$showPageAtom = Atom(name: '_SettingsStore.showPage' ); @override String get showPage { _$showPageAtom.reportObserved(); return super.showPage; } @override set showPage( String value) { _$showPageAtom.context.checkIfStateModificationsAreAllowed(_$showPageAtom); super.showPage = value; _$showPageAtom.reportChanged(); } final _$getPrefsDataAsyncAction = AsyncAction( 'getPrefsData' ); @override Future getPrefsData() { return _$getPrefsDataAsyncAction.run(() => super.getPrefsData()); } final _$saveShowPageAsyncAction = AsyncAction( 'saveShowPage' ); @override Future saveShowPage( String showPage) { return _$saveShowPageAsyncAction.run(() => super.saveShowPage(showPage)); } } 复制代码`

要实现上面的效果还需要分几步走：

* 

在需要被观察的数据增加@observable注解，需要执行操作的方法增加@action注解,

* 

接着执行 ` flutter packages pub run build_runner build`

* 

就会自动生成上述的类，特别的是，如果需要实时跟踪store的变化从而实时改变新生成的类，需要执行一个命令:

` flutter packages pub run build_runner watch` , 如果操作失败了，可以尝试下面的clean命令:

` flutter packages pub run build_runner watch --delete-conflicting-outputs`

### 3. 在widget中使用： ###

在需要观察数据变化的widget套上一层Observer widget，

` _buildShowPageLine(BuildContext context) { return GestureDetector( onTap: () { showDialog< String >( context: context, builder: (context) { String selectValue = ' ${settingsStore.showPage} ' ; List < String > valueList = [ "首页" , "生活" ]; return RadioAlertDialog(title: "选择展示页面" , selectValue: selectValue, valueList: valueList); }).then((value) { print (value); settingsStore.saveShowPage(value); }); }, // 在需要观察变化的widget套上一层Observer widget， child: Observer( builder: (_) => ListTile( title: Common.primaryTitle(content: "默认展示页面" ), subtitle: Common.primarySubTitle(content: ' ${settingsStore.showPage} ' ), ) )); } 复制代码`

完成上述步骤就可以通过对store的数据进行操作，从而自动刷新widget。

## 原理 ##

看完上述的使用之后，相信读者会感到又疑惑又神奇。别急，接下来就进入原理的剖析。
首先看到新生成的代码_$SettingsStore，其中有几处关键的插桩代码，

` @override String get showPage { _$showPageAtom.reportObserved(); return super.showPage; } @override set showPage( String value) { _$showPageAtom.context.checkIfStateModificationsAreAllowed(_$showPageAtom); super.showPage = value; _$showPageAtom.reportChanged(); } 复制代码`

可以看到在获取变量时，会调用 ` dart reportObserved()` , 设置变量会调用 ` dart reportChanged` , 从名字就可以看出获取变量就是将变量上报，变为被观察的状态，设置变量其实就是上报数据变化，进行通知。
我们先看看reportObserved()做了什么，

` // atom可以理解为对应的被观察变量的封装 void _reportObserved(Atom atom) { final derivation = _state.trackingDerivation; if (derivation != null ) { derivation._newObservables.add(atom); if (!atom._isBeingObserved) { atom .._isBeingObserved = true.._notifyOnBecomeObserved(); } } } 复制代码`

可以看出核心就是把当前的变量加入被观察的队列中去。

reportChanged做的是啥呢，

` void propagateChanged(Atom atom) { if (atom._lowestObserverState == DerivationState.stale) { return ; } atom._lowestObserverState = DerivationState.stale; for ( final observer in atom._observers) { if (observer._dependenciesState == DerivationState.upToDate) { observer._onBecomeStale(); } observer._dependenciesState = DerivationState.stale; } } 复制代码`

关键的代码是

` if (observer._dependenciesState == DerivationState.upToDate) { observer._onBecomeStale(); } 复制代码`

当数据需要更新的时候，调用观察者的_onBecomeStale方法，看到这里，相信聪明的读者应该会记起观察者的存在了。 那就是我们用了被观察数据的widget上面套着的Observer的widget。源码如下：

` library flutter_mobx; // ignore_for_file:implementation_imports import 'package:flutter/widgets.dart' ; import 'package:mobx/mobx.dart' ; import 'package:mobx/src/core.dart' show ReactionImpl; /// Observer observes the observables used in the `builder` function and rebuilds the Widget /// whenever any of them change. There is no need to do any other wiring besides simply referencing /// the required observables. /// /// Internally, [Observer] uses a `Reaction` around the `builder` function. If your `builder` function does not contain /// any observables, [Observer] will throw an [AssertionError]. This is a debug-time hint to let you know that you are not observing any observables. class Observer extends StatefulWidget { /// Returns a widget that rebuilds every time an observable referenced in the /// [builder] function is altered. /// /// The [builder] argument must not be null. Use the [context] to specify a ReactiveContext other than the `mainContext`. /// Normally there is no need to change this. [name] can be used to give a debug-friendly identifier. const Observer({ @required this.builder, Key key, this.context, this.name}) : assert (builder != null ), super (key: key); final String name; final ReactiveContext context; final WidgetBuilder builder; @visibleForTesting Reaction createReaction( Function () onInvalidate) { final ctx = context ?? mainContext; return ReactionImpl(ctx, onInvalidate, name: name ?? 'Observer@ ${ctx.nextId} ' ); } @override State<Observer> createState() => _ObserverState(); void log( String msg) { debugPrint(msg); } } class _ObserverState extends State < Observer > { ReactionImpl _reaction; @override void initState() { super.initState(); _reaction = widget.createReaction(_invalidate); } void _invalidate() => setState(noOp); static void noOp() {} @override Widget build(BuildContext context) { Widget built; dynamic error; _reaction.track(() { try { built = widget.builder(context); } on Object catch (ex) { error = ex; } }); if (!_reaction.hasObservables) { widget.log( 'There are no observables detected in the builder function for ${_reaction.name} ' ); } if (error != null ) { throw error; } return built; } @override void dispose() { _reaction.dispose(); super.dispose(); } } 复制代码`

猜猜我们看到了什么， Observer继承自StatefulWidget，看到这里应该就豁然开朗了吧，其实就是在我们的widget上面套了一个父的widget，并且是StatefulWidget类型的，这样一来，只要更新了父widget，同样的我们的widget也就可以进行更新了。
在build的过程，可以看到调用了track方法，跟踪源码可以发现就是先调用了传入的方法（这里对应的是我们widget的构建），然后就是把Observer插入观察者队列:

` void _bindDependencies(Derivation derivation) { final staleObservables = derivation._observables.difference(derivation._newObservables); final newObservables = derivation._newObservables.difference(derivation._observables); var lowestNewDerivationState = DerivationState.upToDate; // Add newly found observables for ( final observable in newObservables) { observable._addObserver(derivation); // Computed = Observable + Derivation if (observable is Computed) { if (observable._dependenciesState.index > lowestNewDerivationState.index) { lowestNewDerivationState = observable._dependenciesState; } } } // Remove previous observables for ( final ob in staleObservables) { ob._removeObserver(derivation); } if (lowestNewDerivationState != DerivationState.upToDate) { derivation .._dependenciesState = lowestNewDerivationState .._onBecomeStale(); } derivation .._observables = derivation._newObservables .._newObservables = {}; // No need for newObservables beyond this point } 复制代码`

接着我们需要找出观察者的_onBecomeStale方法，如果跟踪_onBecomeStale方法，可以发现最终调用的是reaction的run方法:

` @override void _run() { if (_isDisposed) { return ; } _context.startBatch(); _isScheduled = false ; if (_context._shouldCompute( this )) { try { _onInvalidate(); } on Object catch (e) { // Note: "on Object" accounts for both Error and Exception _errorValue = MobXCaughtException(e); _reportException(e); } } _context.endBatch(); } 复制代码`

其中的 ` _onInvalidate()` 就是在observer构成的时候传入的方法：

` void _invalidate() => setState(noOp); static void noOp() {} 复制代码`

看到这里，其实已经水落石出了，就是通过调用的setState从而刷新了widget。

## 总结 ##

对于Mobx，本质就是在使用了被观察数据的widget上面套了一个父的widget，而这个父的widget是一个StatefulWidget。 然后通过观察者模式，发现数据更改时，通知观察者，然后观察者调用了setState了，更新了Observer，从而最后达到刷新子widget的效果。

## 仓库 ##

点击 [flutter_demo]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Ftaoszu%2Fflutter_demo%2Ftree%2Fmaster%2Flib%2Fsettings ) ，查看完整代码。