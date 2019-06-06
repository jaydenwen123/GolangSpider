# Android 事件分发机制 #

# Android 事件分发机制 #

[TOC]

## 前言 ##

Android 分发机制是每个 Android 开发者所要必须了解的知识点，了解了分发机制以后就可以很轻松的解决开发中遇到的问题，比如：

* 实现键盘弹出时，点击空白处隐藏键盘
* 解决滑动冲突
* 自定义 View 实现仿微信录音
* 还有一些其他的应用等。

我会根据源码的方式进行讲解，尽量描述的清楚些。

## Android 事件分发中的事件是什么？ ##

首先，我们在操作移动设备的时候，大多数的操作都是通过手指在屏幕上点击、滑动进行的，为什么我们在点击一个按钮、滑动一个列表的时候能够操控页面呢？

其实这里就是 Android 事件分发机制的体现。

在我们点击按钮的时候，会有按下、抬起的操作，其实这两个操作对应的就是 ACTION_DOWN 事件和 ACTION_UP 事件。

在我们滑动页面的时候，会有按下、移动、移动。。。移动、抬起的操作，这这就对应事件分发里面的 ACTION_DOWN 事件、ACTION_MOVE、ACTION_MOVE。。。ACTION_MOVE、 ACTION_UP 事件。

不管是上面的点击还是滑动，在进行的时候，都是构成了一个事件序列，一个事件序列，一般都是从 ACTION_DOWN 开始、 ACTION_UP 结束的。

在 Android 中，这些事件序列被封装到了 MotionEvent 中。在 MotionEvent 中一般有以下几种重要的事件：

+---------------------------+--------------------------+
|         事件类型          |         触发条件         |
+---------------------------+--------------------------+
| MotionEvent.ACTION_DOWN   | 手指接触到屏幕           |
| MotionEvent.ACTION_MOVE   | 手指在屏幕上滑动         |
| MotionEvent.ACTION_UP     | 手指离开屏幕             |
| MotionEvent.ACTION_CANCEL | 意外原因导致事件序列终止 |
+---------------------------+--------------------------+

大体上，就是这四种不同的事件来构成事件分发里面的事件序列。

## Android 事件分发是什么？ ##

上面讲了关于事件的概念，那么事件到底是怎么分发的呢？

这要从我们应用页面讲起。Android 中的页面构成是这样的：

![-w346](https://user-gold-cdn.xitu.io/2019/6/6/16b2ab181b098c71?imageView2/0/w/1280/h/960/ignore-error/1)

其中 Activity 是在最外面的，而我们平时 ` setContentView` 的布局是在最里面的，而我们布局正常情况下是最外面是 ViewGroup ，里面包裹了 ViewGroup 或者 View 。

比如在登录页面，我们点击了登录按钮，那么就完成了一次完整的事件分发。

![image](https://user-gold-cdn.xitu.io/2019/6/6/16b2ab181b325bf3?imageView2/0/w/1280/h/960/ignore-error/1)

这个时候事件走向是这样的

> 
> 
> 
> Activity -> PhoneWindow -> DecorView -> ViewGroup -> LinearLayout ->
> Button
> 
> 

可以看到：事件是从最外面的 Activity 传递到最里层的 Button 按钮的。

其实事件分发就是将我们手指产生的一些事件，传递到一个具体的 View 上去并且处理的过程。

## 为什么会有事件分发机制？ ##

Android 上的View是树形结构的，View 可能会重叠到一起，当我们点击的时候，可能会有多个 View 进行响应，事件分发机制主要是解决事件该交给谁去处理的。

比如上面，我们点击了 Button，但是 Button 是在 LinearLayout 里面的，那我们点击 Button 的时候，到底是 交给 LinearLayout 处理还是 Button 处理，由事件分发机制说了算。

还有就是 Android 中的滑动冲突等，都要需要依据事件分发机制去解决。

## 事件分发里面重要的三个方法 ##

* dispatchTouchEvent() 在 Activity 、 ViewGroup、 View 中都是有这个方法的，事件接收是从 Activity 的 dispatchTouchEvent(） 开始的。
* onInterceptTouchEvent() 这个方法是 ViewGroup 独有的，在 ViewGroup 中可以通过这个方法确定是否拦截该事件，如果拦截，那么就由 ViewGroup 的 onTouchEvent（） 方法接管事件序列。这个方法默认是返回 false 的，也就是默认不拦截事件
* onTouchEvent() 这个方法也是在 Activity 、 ViewGroup 、 View 中都有的，如果如果确定在 Activity、ViewGroup、View 中处理事件，一般是在这个方法处理的。

## 事件分发讲解 ##

先看 Activity 的。

### Activity 的事件分发 ###

前面讲到事件序列是从最外层的 Activity 开始接收的，然后依次分发到具体的 View 中的，那就先从最外层的 Activity 层开始看起，也就是从 Activity 的 dispatchTouchEvent() 方法看起。

` public boolean dispatchTouchEvent (MotionEvent ev) { // 首先当我们每次手指触控屏幕的时候，都会去调用 onUserInteraction() 方法， // 如果你想知道用户用某种方式和你正在运行的 activity 交互，可以重写此方法。 // 因为在每次事件分发的时候都会调用到该方法。 if (ev.getAction() == MotionEvent.ACTION_DOWN) { onUserInteraction(); } if (getWindow().superDispatchTouchEvent(ev)) { return true ; } return onTouchEvent(ev); } 复制代码`

可以看到会先调用 ` getWindow().superDispatchTouchEvent(ev)` 方法，来看下这个 getWindow() 是什么:

` public Window getWindow () { return mWindow; } 复制代码`

这里返回的是一个 Window 对象，在 Android 里面的 Window 是一个抽象类，唯一的实现是 PhoneWindow 类。也就是调用的 PhoneWindow 的 superDispatchTouchEvent(ev) 方法：

` @Override public boolean superDispatchTouchEvent (MotionEvent event) { // 这里调用了 mDecor 的 superDispatchTouchEvent 方法 return mDecor.superDispatchTouchEvent(event); } // mDecor 是一个DecorView对象 private DecorView mDecor; public boolean superDispatchTouchEvent (MotionEvent event) { // 调用父类的 super.dispatchTouchEvent 方法 return super.dispatchTouchEvent(event); } // DecorView 继承于 FrameLayout public class DecorView extends FrameLayout implements RootViewSurfaceTaker , WindowCallbacks {} // FrameLayout 继承于ViewGroup public class FrameLayout extends ViewGroup {} 复制代码`

可以看到，DecorView 是继承于 FrameLayout 的，FrameLayout是继承于 ViewGroup 的，所以最终调用了 ViewGroup 的 dispatchTouchEvent（） 方法

这个时候，事件已经被传递到了 ViewGroup 。也就是说如果 getWindow().superDispatchTrackballEvent(ev) 这行代码返回的是 true ，表示 事件被消费掉了，那么本次事件分发就结束了。直接 return 。不会执行到

` return onTouchEvent(ev); 复制代码`

也就是不会执行 Activity 的 onTouchEvent 方法，如果是getWindow().superDispatchTrackballEvent(ev) 返回的 false，那么就表示ViewGroup 和 View 均没有对事件进行处理，调用 Activity 的 onTouchEvent 方法。

上面已经讲了 onInterceptTouchEvent() 是只存在 ViewGroup 中的，上面的源码也验证了这一点。下面看看 ViewGroup 的 dispatchTouchEvent（） 方法

### ViewGroup 的事件分发。 ###

ViewGroup 的事件分发，也是从 dispatchTouchEvent() 开始的，来看下关键代码：

` @Override public boolean dispatchTouchEvent (MotionEvent ev) { // 省略n行代码。。。 // Handle an initial down. // 这里在每次事件是 MotionEvent.ACTION_DOWN 的时候，调用 resetTouchState() // 因为事件序列是从 down 事件开始的，所以每次接收到 down 事件，就是一个新的事件序列 if (actionMasked == MotionEvent.ACTION_DOWN) { // Throw away all previous state when starting a new touch gesture. // The framework may have dropped the up or cancel event for the previous gesture // due to an app switch, ANR, or some other state change. cancelAndClearTouchTargets(ev); resetTouchState(); } // Check for interception. final boolean intercepted; // mFirstTouchTarget 在子 View 接管事件的时候会赋值，否则为 null // 如果是 ACTION_DOWN 事件,或者 mFirstTouchTarget 不为空,表明ACTION_DOWN事件没有被消费， // 走 ViewGroup 的分发流程， if (actionMasked == MotionEvent.ACTION_DOWN || mFirstTouchTarget != null ) { // 是否允许拦截 disallowIntercept = 是否禁用事件拦截的功能(默认是false)， // 可通过调用 requestDisallowInterceptTouchEvent（）修改 FLAG_DISALLOW_INTERCEPT这个标志位 final boolean disallowIntercept = (mGroupFlags & FLAG_DISALLOW_INTERCEPT) != 0 ; // 如果不允许拦截,就会 执行 onInterceptTouchEvent(ev) if (!disallowIntercept) { // 这里会根据 onInterceptTouchEvent 的返回值判断当前的 ViewGroup 是否进行拦截. // onInterceptTouchEvent(ev)的源码显示，默认是会返回 false 的，如果有需要我们可以返回 true intercepted = onInterceptTouchEvent(ev); ev.setAction(action); // restore action in case it was changed } else { // ViewGroup 默认是不拦截的,所以置为 false intercepted = false ; } } else { // There are no touch targets and this action is not an initial down // so this view group continues to intercept touches. // 如果不是 Down 事件，并且没有子 View 接管事件,那么 ViewGroup 会阻止后面的事件向后传递 // intercepted置为 true ,拦截事件序列，不需要调用onInterceptTouchEvent(ev)，自己处理 intercepted = true ; } // If intercepted, start normal event dispatch. Also if there is already // a view that is handling the gesture, do normal event dispatch. // 如果拦截,或者一个 View 接管了该事件序列,那么就走正常的分发流程 if (intercepted || mFirstTouchTarget != null ) { ev.setTargetAccessibilityFocus( false ); } // 如果没有触发取消事件并且没有拦截 if (!canceled && !intercepted) { //省略部分代码... // for 循环遍历 View 注意，这里是倒叙的，也就是从最里面的 View 进行遍历 for ( int i = childrenCount - 1 ; i >= 0 ; i--) { final int childIndex = getAndVerifyPreorderedIndex( childrenCount, i, customOrder); final View child = getAndVerifyPreorderedView( preorderedList, children, childIndex); // If there is a view that has accessibility focus we want it // to get the event first and if not handled we will perform a // normal dispatch. We may do a double iteration but this is // safer given the timeframe. if (childWithAccessibilityFocus != null ) { if (childWithAccessibilityFocus != child) { continue ; } childWithAccessibilityFocus = null ; i = childrenCount - 1 ; } // 检查能不能接收到事件，检查触摸位置是否在View区域内，并且在不在播放动画 // 如果都不满足，执行 continue 继续循环下个 View if (!canViewReceivePointerEvents(child) || !isTransformedTouchPointInView(x, y, child, null )) { ev.setTargetAccessibilityFocus( false ); continue ; } // 看是否有 View 接管，如果有，跳出循环。 newTouchTarget = getTouchTarget(child); if (newTouchTarget != null ) { // Child is already receiving touch within its bounds. // Give it the new pointer in addition to the ones it is handling. newTouchTarget.pointerIdBits |= idBitsToAssign; break ; } resetCancelNextUpFlag(child); // 这里进行处理，如果返回 true ，那么表示有 View 接管了 事件，那么就给 newTouchTarget // 在 addTouchTarget(child, idBitsToAssign)中赋值，View接管该事件。结束循环完成分发。 if (dispatchTransformedTouchEvent(ev, false , child, idBitsToAssign)) { // Child wants to receive touch within its bounds. mLastTouchDownTime = ev.getDownTime(); if (preorderedList != null ) { // childIndex points into presorted list, find original index for ( int j = 0 ; j < childrenCount; j++) { if (children[childIndex] == mChildren[j]) { mLastTouchDownIndex = j; break ; } } } else { mLastTouchDownIndex = childIndex; } mLastTouchDownX = ev.getX(); mLastTouchDownY = ev.getY(); newTouchTarget = addTouchTarget(child, idBitsToAssign); alreadyDispatchedToNewTouchTarget = true ; break ; } // The accessibility focus didn't handle the event, so clear // the flag and do a normal dispatch to all children. ev.setTargetAccessibilityFocus( false ); } } // 省略n行代码。。。 } 复制代码`

可以看到，在执行 dispatchTouchEvent 的开始，如果是 ACTIOPN_DOWN 的话，调用 resetTouchState() 来重置所有的触摸状态，这里会将 mFirstTouchTarget 设置为 null 。然后准备新的周期，这样做主要是因为事件序列是从 down 事件开始的，所以每次接收到 down 事件，就是一个新的事件序列。要重新开始处理。

然后执行

` if (actionMasked == MotionEvent.ACTION_DOWN || mFirstTouchTarget != null ) { // 是否允许拦截 disallowIntercept = 是否禁用事件拦截的功能(默认是false)， // 可通过调用 requestDisallowInterceptTouchEvent（）修改 FLAG_DISALLOW_INTERCEPT这个标志位 final boolean disallowIntercept = (mGroupFlags & FLAG_DISALLOW_INTERCEPT) != 0 ; // 如果不允许拦截,就会 执行 onInterceptTouchEvent(ev) if (!disallowIntercept) { // 这里会根据 onInterceptTouchEvent 的返回值判断当前的 ViewGroup 是否进行拦截. // onInterceptTouchEvent(ev)的源码显示，默认是会返回 false 的，如果有需要我们可以返回 true intercepted = onInterceptTouchEvent(ev); ev.setAction(action); // restore action in case it was changed } else { // ViewGroup 默认是不拦截的,所以置为 false intercepted = false ; } } else { // There are no touch targets and this action is not an initial down // so this view group continues to intercept touches. // 如果不是 Down 事件，并且没有子 View 接管事件,那么 ViewGroup 会阻止后面的事件向后传递 // intercepted置为 true ,拦截事件序列，不需要调用onInterceptTouchEvent(ev)，自己处理 intercepted = true ; } 复制代码`

这里会判断是不是 DOWN 事件或者是 mFirstTouchTarget 不是 null。mFirstTouchTarge != null 也就是已经找到能够接收 touch 事件的 View。这个时候会进入 if 内部。

在内部会先创建 disallowIntercept (禁止拦截) 标志位。这个标志位可以在 子 view 中使用 requestDisallowInterceptTouchEvent(boolean disallowIntercept) 方法去设置。以便请求父 View 不拦截事件。

如果 disallowIntercept = false ，也就是在子 view 执行 ` requestDisallowInterceptTouchEvent(false)` 也就是请求拦截，那么就会执行 onInterceptTouchEvent(ev); 方法去判断当前的 ViewGroup 是否对该事件进行拦截。

` public boolean onInterceptTouchEvent (MotionEvent ev) { if (ev.isFromSource(InputDevice.SOURCE_MOUSE) && ev.getAction() == MotionEvent.ACTION_DOWN && ev.isButtonPressed(MotionEvent.BUTTON_PRIMARY) && isOnScrollbarThumb(ev.getX(), ev.getY())) { return true ; } return false ; } 复制代码`

onInterceptTouchEvent 方法内部前面的是对一些特殊事件的处理，就先不管了，如果不满足上面的四个条件的话，就会直接返回 false ，也就是 ViewGroup 默认是不拦截的。所以这个时候的 intercepted 值为 false 不拦截；

如果 disallowIntercept = false ，也就是在子 view 执行 ` requestDisallowInterceptTouchEvent(true)` 也就是请求不拦截，那么ViewGroup 就不会去走 onInterceptTouchEvent 方法，直接将 intercepted 值置为 false，表示不拦截。

我们看到最后一个 else 把 intercepted 的值直接置为 true，表示拦截，那么是什么时候触发的呢？

也就是说在 当前传递来的事件不是 DOWN 并且没有 View 接管事件的时候，ViewGroup 默认是进行拦截的。 **很好理解，事件的开端没有子 View 进行处理，那么之后的事件也不会交给子 View 去处理了。**

再往后的话就是通过一个循环，倒序遍历所有的子 View，依次判断能不能接收到事件或者是正在播放动画，然后满足的话依次去执行 ` if (dispatchTransformedTouchEvent(ev, false, child, idBitsToAssign))` 。

` private boolean dispatchTransformedTouchEvent (MotionEvent event, boolean cancel, View child, int desiredPointerIdBits) { final boolean handled; // Canceling motions is a special case. We don't need to perform any transformations // or filtering. The important part is the action, not the contents. // 如果 取消，或者 事件为 ACTION_CANCEL，不做任何过滤或转换 final int oldAction = event.getAction(); if (cancel || oldAction == MotionEvent.ACTION_CANCEL) { event.setAction(MotionEvent.ACTION_CANCEL); if (child == null ) { handled = super.dispatchTouchEvent(event); } else { handled = child.dispatchTouchEvent(event); } event.setAction(oldAction); return handled; } // Calculate the number of pointers to deliver. final int oldPointerIdBits = event.getPointerIdBits(); final int newPointerIdBits = oldPointerIdBits & desiredPointerIdBits; // If for some reason we ended up in an inconsistent state where it looks like we // might produce a motion event with no pointers in it, then drop the event. if (newPointerIdBits == 0 ) { return false ; } // If the number of pointers is the same and we don't need to perform any fancy // irreversible transformations, then we can reuse the motion event for this // dispatch as long as we are careful to revert any changes we make. // Otherwise we need to make a copy. final MotionEvent transformedEvent; if (newPointerIdBits == oldPointerIdBits) { if (child == null || child.hasIdentityMatrix()) { if (child == null ) { handled = super.dispatchTouchEvent(event); } else { final float offsetX = mScrollX - child.mLeft; final float offsetY = mScrollY - child.mTop; event.offsetLocation(offsetX, offsetY); handled = child.dispatchTouchEvent(event); event.offsetLocation(-offsetX, -offsetY); } return handled; } transformedEvent = MotionEvent.obtain(event); } else { transformedEvent = event.split(newPointerIdBits); } // Perform any necessary transformations and dispatch. if (child == null ) { handled = super.dispatchTouchEvent(transformedEvent); } else { final float offsetX = mScrollX - child.mLeft; final float offsetY = mScrollY - child.mTop; transformedEvent.offsetLocation(offsetX, offsetY); if (!child.hasIdentityMatrix()) { transformedEvent.transform(child.getInverseMatrix()); } // 如果有子 View ，那么去调用子 View 的 dispatchTouchEvent 去处理该事件。也就 //把事件传递到了下一级 handled = child.dispatchTouchEvent(transformedEvent); } // Done. transformedEvent.recycle(); return handled; } 复制代码`

在 dispatchTransformedTouchEvent 内部可以看到，如果是 child 不为 null 的时候，调用子 View 的 dispatchTouchEvent 去处理，这样就把把事件交给子 View 去处理，需要注意的是： **这个地方的返回值 handled 是 最后子 View 的 onTouchEvent 的返回值。**

**如果返回值为 true ，**那么表示子 View 消费掉了事件，那么就会在ViewGroup 的 dispatchTouchEvent 方法的内部通过：

` newTouchTarget = addTouchTarget(child, idBitsToAssign); alreadyDispatchedToNewTouchTarget = true ; private TouchTarget addTouchTarget (@NonNull View child, int pointerIdBits) { final TouchTarget target = TouchTarget.obtain(child, pointerIdBits); target.next = mFirstTouchTarget; mFirstTouchTarget = target; return target; } 复制代码`

对newTouchTarget 进行赋值。在 addTouchTarget 内部对 mFirstTouchTarget 进行赋值。

**如果返回的是 false，**那么就跳过 if 内部。最终会执行到下面的代码：

(这里的代码是 ViewGroup 的 dispatchTouchEvent(event) 方法的一部分)

` // Dispatch to touch targets. if (mFirstTouchTarget == null ) { // No touch targets so treat this as an ordinary view. // 没有 View 处理，那么交给自己处理，在 View 的 dispatchTouchEvent方法中调用 // onTouchEvent，执行 if (child == null) { // handled = super.dispatchTouchEvent(event); // } handled = dispatchTransformedTouchEvent(ev, canceled, null , TouchTarget.ALL_POINTER_IDS); } else { // Dispatch to touch targets, excluding the new touch target if we already // dispatched to it. Cancel touch targets if necessary. // 走到这里说明了，已经把事件交给了具体的 View 处理，那么会将MotionEvent.ACTION_DOWN // 的后续事件分发给mFirstTouchTarget指向的View去处理 TouchTarget predecessor = null ; TouchTarget target = mFirstTouchTarget; while (target != null ) { final TouchTarget next = target.next; // 该TouchTarget已经在前面的情况中被分发处理了，避免重复处理 if (alreadyDispatchedToNewTouchTarget && target == newTouchTarget) { handled = true ; } else { // 如果被当前ViewGroup拦截，向下分发cancel事件 final boolean cancelChild = resetCancelNextUpFlag(target.child) || intercepted; // dispatchTransformedTouchEvent()方法成功向下分发取消事件或分发正常事件 if (dispatchTransformedTouchEvent(ev, cancelChild, target.child, target.pointerIdBits)) { handled = true ; } // 如果发送了取消事件，则移除分发记录（链表移动操作) if (cancelChild) { if (predecessor == null ) { mFirstTouchTarget = next; } else { predecessor.next = next; } target.recycle(); target = next; continue ; } } predecessor = target; target = next; } } 复制代码`

这里会判断有没有 View 处理事件。

如果没有 View 处理，就调用 dispatchTransformedTouchEvent 方法，传入 child = null，执行 ` handled = super.dispatchTouchEvent(event);` 这个时候 ViewGroup super 的 dispatchTouchEvent方法，也就是调用 View 的 dispatchTouchEvent 去处理，进而在View 的 dispatchTouchEvent 中调用 onTouchEvent。这个时候调用的 onTouchEvent就是 ViewGroup 内部的 onTouchEvent 方法。

如果有 View 处理，那么就把后续事件交给子 View 继续处理。这个时候，事件就传递到了 View 里面。

这里有个点就是：

如果在子 View 中的 onTouchEvent 返回了 true，那么 mFirstTouchTarget 就不为空，就轮不到父 View 处理。

如果在子 View 中的 onTouchEvent 返回了 false，子 View 不处理事件，那么 mFirstTouchTarget 就为空，就需要父 View 处理。

### View 的事件分发 ###

其实 View 的分发机制就比较简单了，View 的分发机制也是从 dispatchTouchEvent(event) 开始的：

` public boolean dispatchTouchEvent (MotionEvent event) { // If the event should be handled by accessibility focus first. if (event.isTargetAccessibilityFocus()) { // We don't have focus or no virtual descendant has it, do not handle the event. if (!isAccessibilityFocusedViewOrHost()) { return false ; } // We have focus and got the event, then use normal event dispatch. event.setTargetAccessibilityFocus( false ); } boolean result = false ; if (mInputEventConsistencyVerifier != null ) { mInputEventConsistencyVerifier.onTouchEvent(event, 0 ); } final int actionMasked = event.getActionMasked(); if (actionMasked == MotionEvent.ACTION_DOWN) { // Defensive cleanup for new gesture stopNestedScroll(); } if (onFilterTouchEventForSecurity(event)) { if ((mViewFlags & ENABLED_MASK) == ENABLED && handleScrollBarDragging(event)) { result = true ; } //noinspection SimplifiableIfStatement ListenerInfo li = mListenerInfo; // 可以看到,比如满足四个条件，才会消费事件。 //1、ListenerInfo 不为 null //2、mOnTouchListener 不为 null //3、View 的 enable 是可用的 //4、onTouch 方法返回的是 true if (li != null && li.mOnTouchListener != null && (mViewFlags & ENABLED_MASK) == ENABLED && li.mOnTouchListener.onTouch( this , event)) { result = true ; } // 从上面的代码可以看到，是先执行了 onTouch 方法的，如果在 onTouch 方法 // 里面已经把事件消费掉，那么久不会执行 onTouchEvent 方法，如果没消费， // 就判断是否在 onTouchEvent 中消费掉，如果消费掉，也是直接返回。如果没没消费，继续向下执行 if (!result && onTouchEvent(event)) { result = true ; } } if (!result && mInputEventConsistencyVerifier != null ) { mInputEventConsistencyVerifier.onUnhandledEvent(event, 0 ); } // Clean up after nested scrolls if this is the end of a gesture; // also cancel it if we tried an ACTION_DOWN but we didn't want the rest // of the gesture. if (actionMasked == MotionEvent.ACTION_UP || actionMasked == MotionEvent.ACTION_CANCEL || (actionMasked == MotionEvent.ACTION_DOWN && !result)) { stopNestedScroll(); } return result; } 复制代码`

再来看下 View 的 onTouchEvent ：

` public boolean onTouchEvent (MotionEvent event) { final float x = event.getX(); final float y = event.getY(); final int viewFlags = mViewFlags; final int action = event.getAction(); // 判断view 是可点击的 final boolean clickable = ((viewFlags & CLICKABLE) == CLICKABLE || (viewFlags & LONG_CLICKABLE) == LONG_CLICKABLE) || (viewFlags & CONTEXT_CLICKABLE) == CONTEXT_CLICKABLE; // 如果 View 可点击，但是处于不可用状态，仍然会消费掉事件。 if ((viewFlags & ENABLED_MASK) == DISABLED) { if (action == MotionEvent.ACTION_UP && (mPrivateFlags & PFLAG_PRESSED) != 0 ) { setPressed( false ); } mPrivateFlags3 &= ~PFLAG3_FINGER_DOWN; // A disabled view that is clickable still consumes the touch // events, it just doesn't respond to them. return clickable; } // 如果 View 设置了代理，会调用代理的 onTouchEvent 方法 if (mTouchDelegate != null ) { if (mTouchDelegate.onTouchEvent(event)) { return true ; } } // 如果 View 是可用的，是可点击的，并且没有被 onTouchListener 的 onTouch 方法 // 消费事件，那么就会执行到 点击事件 performClick() if (clickable || (viewFlags & TOOLTIP) == TOOLTIP) { switch (action) { case MotionEvent.ACTION_UP: ... if (!post(mPerformClick)) { performClick(); ... break ; return true ; } return false ; } 复制代码`

从上面的源码分析可以看出，View 的 onTouchListener 优先级高于 onTouchEvent 的优先级，onTouchEvent 的优先级高于 onClickListener onClick 的优先级。

并且从源码中也看到了，如果我们设置了 View 是 disabled 的，那么也就没有下面的 onTouch，因为 && 判断的时候，前面为 false，后面就不去判断，所以 onTouch 就不会执行。但是不会影响执行 onTouchEvent。

在 onTouchEvent 的内部，如果得到 View 是 disabled 的，就会返回 View 是不是 clickable 或者 longClickable 的，如果有一个是 true ，就消费掉该事件。不会执行onCLick 事件了。

onTouch 执行条件

* 父 ViewGroup 没有拦截事件
* 当前 View 是可可用的

onClick 执行的条件如下：

* 父 ViewGroup 没有拦截事件
* 当前 View 是可可用的
* 当前 View 是可点击的
* 当前 View 内部对 ACTION_UP 事件处理
* onTouch不消费事件

为了便于理解，可以参考这张图：

![image](https://user-gold-cdn.xitu.io/2019/6/6/16b2ab181c7a21e1?imageView2/0/w/1280/h/960/ignore-error/1)

## 关于事件分发的总结 ##

* 事件分发里面的事件通常指 ACTION_DOWN、ACTION_MOVE、ACTION_UP、ACTION_Cancel 这四种，他们连在一起就构成了一个时间序列。比如 ACTION_DOWN、ACTION_MOVE、ACTION_MOVE…ACTION_MOVE、ACTION_UP，这就构成了一个事件序列。
* 事件序列的传递是从 Activity 开始，依次经过、PhoneWindow、DecorView、ViewGroup、View。如果是最终的 View 也没有处理的话，就依次向上移交，最终会在 Activity 的 onTouchEvent 方法中处理。
* 如果事件从 ViewGroup 中传递给 View 去处理的时候，如果 View 没有处理掉，在 onTouchEvent 方法中返回了 false，那么该事件就重新交给 ViewGroup 处理，并且后续的事件都不会再传递给该 View。
* onInterceptTouchEvent 方法只有在 ViewGroup 中存在，并且默认返回 false，代表 ViewGroup 不拦截事件。
* 正常情况下，一个事件序列只能由一个 View 处理。如果一个 View 接管了事件，不管是具体的子 View还是 ViewGroup，后续的事件都会让这个 View 处理，除非人为干预事件的分发过程。
* 子 View 可以通过调用 requestDisallowInterceptTouchEvent(true) ，干预父元素的除了 ACTION_DOWN 事件以外的事件走向。 一般用于处理滑动冲突中，子控件请求父控件不拦截ACTION_DOWN以外的其他事件，ACTION_DOWN事件不受影响。
* View 的 onTouchEvent 方法默认是返回 true 的，也就是会默认拦截事件。除非 它是不可点击的，(clickable、longClickable 都为 false)。View 的 longClickable 默认均为 false，Button、ImageButton 的 clickable 默认为 true，TextView clickable 默认为false
* View 的 enable 属性不会影响 onTouchEvent 的返回值，只要 clickable、longClickable 有一个为 true，那么onTouchEvent就默认会返回 true
* View 的点击事件是在 ACTION_UP 事件处理的时候执行的，所以要执行，必须要有 ACTION_DOWN 和 ACTION_UP 两个事件。

![欢迎关注我的公众号](https://user-gold-cdn.xitu.io/2019/6/6/16b2ab181c447ff1?imageView2/0/w/1280/h/960/ignore-error/1)