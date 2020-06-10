### state
```
this.state = {a: null, b: null}

// 写法1
	this.setState({
		a: 1,
		b: 2
	})

// 写法2
	this.setState(({a})=>({
		a: a+1,
		b: 2
	}))	

// 写法3
	this.setState((state, props)=>({
		a: state.a + 1,
		b: state.b
	}))	
```

### setState
```
// 修改state
	this.setState({
		a: a
	})

// state更新可能是异步的
	// React 为了优化性能，有可能会将多个 setState() 调用合并为一次更新
	// 因为this.props和this.state 可能是异步更新的，你不能依赖他们的值计算下一个state
	this.setState((proState, prop) => ({
		a: prop.a
	}))

// 修改state后，怎么拿到最新的值
	// 通过回调方式
	this.setState((state, props)=>({
		a: props.a
	}), ()=>{
		console.log(this.state)
	})
```

### 条件渲染
```
// 1
	if () {
		<>
	}else{
		<>
	}

// 2
	let a = true
	<div>
		{a==true && (
			<h1>aaaa</h1>
			<h1>aaaa</h1>
			<h1>aaaa</h1>
		)}
	</div>
	
// 3
	{a == true 
	?	(<div>
			<h1>aaaaaa</h1>
			<h1>aaaaaa</h1>
		</div>)
	:	(<div>
			<h1>bbbbbb</h1>
			<h1>bbbbbb</h1>
		</div>)
	}
```

### 阻止渲染
```
render () {
	return null
}

```

### 传递组件
```
// 用props.children传递
    function FancyBorder (props) {
        return (
            <div class={"fancyclass"+props.color}>
                {props.children}
            </div>
        )
    }

    function WelcomeDialog() {
        return (
            <FancyBorder color="red">
                <h1>dsadsaa</h1>
                <h1>dsadsaa</h1>
                <h1>dsadsaa</h1>
            </Fancy>
        )
    }

// 利用参数传递
    function SplitPane(props) {
        return(
            <div>
                {props.left}
                {props.right}
            </div>
        )
    }

    function App(){
        return (
            <SplitPane
                left={
                    <h1>zzzz</h1>
                }
                right={
                    <h1>ssss</h1>
                }
            />
        )
    }



```


### React.lazy
```
http://www.ruanyifeng.com/blog/2016/09/redux_tutorial_part_one_basic_usages.html

// React.lazy 函数能让你像渲染常规组件一样处理动态引入（的组件）

const OtherComponent = React.lazy(() => import('./OtherComponent'))
```

### Context
```
// React.createContext
    const MyContext = React.createContext(defaultValue)

    创建一个 Context 对象。当 React 渲染一个订阅了这个 Context 对象的组件，这个组件会从组件树中离自身最近的那个匹配的 Provider 中读取到当前的 context 值
    

// Context.Provider
    <MyContext.Provider value={/* 某个值 */}>

    每个 Context 对象都会返回一个 Provider React 组件，它允许消费组件订阅 context 的变化
    
    Provider 接收一个 value 属性，传递给消费组件。一个 Provider 可以和多个消费组件有对应关系。多个 Provider 也可以嵌套使用，里层的会覆盖外层的数据。
    
    当 Provider 的 value 值发生变化时，它内部的所有消费组件都会重新渲染。Provider 及其内部 consumer 组件都不受制于 shouldComponentUpdate 函数，因此当 consumer 组件在其祖先组件退出更新的情况下也能更新。
    
// Class.contextType
    // 写法1
    class Myclass extends React.Component {
        static contextType = MyContext      // 声明
        let aaa = this.context                            // 引用  
    }
    // 或者
    Myclass.contextType = MyContext  //外部声明

// Context.Consumer
    <MyContext.Consumer>
    {value => /* 基于 context 值进行渲染*/}    //箭头函数，返回组件 (<aaa>)
    </MyContext.Consumer>

// Context.displayName


```

### 错误边界
```
可以捕获并打印发生在其子组件树任何位置的 JavaScript 错误，并且，它会渲染出备用 UI

定义了 static getDerivedStateFromError() 或 componentDidCatch() 这两个生命周期方法

class ErrorBound extends React.component {
	constructor (super) {
		super()
		this.state = {
			error: null,
			errorInfo: null
		}
	},

	componentDidCatch(error, errorInfo){
		this.setState({
			error, errorInfo
		})
	},

	render() {
		if (this.state.error) {
			return (
				<div>
					<h1>something is wrong</h1>
					<details>
						{this.state.error && this.state.errorInfo} 
						<br />
            			{this.state.errorInfo.componentStack}			// 错误组件栈
					</details>
				</div>
			)
		}
	}
}

class App extends React.component {
	constructor (super) {
		super()
		this.state = {count: 0}
		this.handClick = this.handClick.bind(this)
	},
	handClick () {
		this.setState({
			count: this.state.count + 1
		})
	}
	render () {
		if (this.state.count > 3) {
			throw new Error("wrong")
		}
		return (
			<ErrorBound>
				<h1 onClick={this.handClick}>
					{this.state.count}
				</h1>
			</ErrorBound>
		)
	}
}
``` 

### redux
```
// 同步Actions

1、Action
	Action 本质上是 JavaScript 普通对象。我们约定，action 内必须使用一个字符串类型的 type 字段来表示将要执行的动作

	// action 创建函数只是简单的返回一个 action，更容易被移植和测试
	function addTodo(text) {
		return {
			type: ADD_TODO,
			text
		}
	}

	// dispatch
		store.dispatch()是 View 发出 Action 的唯一方法
		import { createStore } from 'redux'
		const store = createStore(fn)
		store.dispatch({
			type: 'ADD_TODO',
			payload: 'Learn Redux'
		});

		// 上面代码中，store.dispatch接受一个 Action 对象作为参数，将它发送出去。
		// 结合 Action Creator，这段代码可以改写如下。
		store.dispatch(addTodo('Learn Redux'));

2、Reduce
	Store 收到 Action 以后，必须给出一个新的 State，这样 View 才会发生变化
	Actions 只是描述了有事情发生了这一事实，并没有描述应用如何更新 state
	Reducers 指定了应用状态的变化如何响应 actions 并发送到 store 的

	1) Reduce函数
		Reducer 是一个函数，它接受 Action 和当前 State 作为参数，返回一个新的 State
		(previousState, action) => newState
		> Reduce的纯粹性
			> 不要修改previousState
			> 返回的newState是在previousState的副本基础上修改
			> 在default情况下返回旧的previousState
		import {
			ADD_TODO,
			TOGGLE_TODO,
			SET_VISIBILITY_FILTER,
			VisibilityFilters
		} from './actions'

		function todoApp(preState, action){
			switch(action.type) {
				case ADD_TODO:
				case TOGGLE_TODO:
				case SET_VISIBILITY_FILTER:
				default:
			}
		}

		2) reducer合并
			redux提供combineReducers方法
			定义各个子 Reducer 函数，然后用这个方法，将它们合成一个大的 Reducer

			import { combineReducers } from 'redux';

			const chatReducer = combineReducers({
				chatLog,
				statusMessage,
				userName
			})

			// State 的属性名必须与子 Reducer 同名。如果不同名，就要采用下面的写法
			const chatReducer = combineReducers({
				a:	doSomeA,
				b:	doSomeB,
				c:	doSomeC
			})

3、store
	(1) createStore
		// createStore简单实现
		const createStore = (reducer)=>{
			let state
			let listeners = []

			const getState = ()=>state

			const dispatch = (action)=>{
				state = reducer(state, action)
				listeners.forEach(e => e())
			}

			const subscribe = (listener)=>{
				listeners.push(listener)
				return ()=>{
					listeners = listeners.filter((e)=>e !== listener)
				}
			}

			dispatch({})

			return {getState,dispatch, subscribe }
		}

	(2)	store.getState()
    (3) store.dispatch()
		// 发起actions
		store.dispatch(addTodo(''))
    (4) store.subscribe()
		// 注册监听器
		const unsubscribe = store.subscribe(()=>{
			console.log(store.getState())
		})
		//停止监听器
		unsubscribe()

		// listener可以通过store.getState()得到当前状态，这时可以触发重新渲染 View
		const listener = ()=>{
			let newState = store.getState()
			component.setState(newState)
		}

// 实例
	const Counter = ({value, onIn, onDe}) => {
		return (
			<div>
				<h1> {value} </h1>
				<button onClick={onIn}> + </button>
				<button onClick={onDe}> - </button>
			</div>
		)
	}

	const reducer = (state, action)=>{
		switch (action.type) {
			case 'INCREMENT':
				return state+1
			case 'DECREMENT':
				return state-1
			default:
				return state
		}
	}

	const store = createStore(reduce)

	const render=()=>{
		ReactDOM.render(
			<Counter
				value={store.getState()}
				onIn={store.dispatch({type: 'INCREMENT'})}
				onDe={store.dispatch({type: 'DECREMENT'})}
			/>
			document.getElementById('root')
		)
	}

	render()
	store.subscribe(render)

// 异步Actions
	中间件就是一个函数，对store.dispatch方法进行了改造，在发出 Action 和执行 Reducer 这两步之间，添加了其他功能

4、redux-chunk
	(1) 注册
		// middleware 
		import chunk from 'redux-chunk'
		import {createStore, applyMiddleware} from 'redux'
		createStore(reducer, applyMiddleware(thunk))

	(2) 异步Action
		同步dispatch(Actor)						// Actor返回一个对象
		异步dispatch(ActorCreator)		// ActorCreator返回一个函数

		> ActorCreator返回一个函数
		> 返回的函数的参数是dispatch和getState

		// index.js
		import thunkMiddleware from 'redux-thunk'
		import {createLogger} from 'redux-logger'
		import {createStore, applyMiddleware} from 'redux'
		import {actionNormal, actionCreator} from './actions'
		import reducers from './reducers'
		const initState = {...}
		const loggerMiddleware = createLogger
		createStore(reducers, 
								initState, 
								applyMiddleware(
									thunkMiddleware,		// 允许我们 dispatch() 函数
									loggerMiddleware
								))
		store.dispatch(actionNormal())
		store.dispatch(actionCreator()).then((res)=>{
			console.log(res)
		}).catch(err=>{
			
		})
		// actions.js
		const actionNormal = ()=>{
			return {
				type: 'xxxx'
			}
		}
		
		const actionCreate = ()=>{
			...
			return (dispatch, getState)=>{
				dispatch(actionNormal)
				return fetch('xxx.html')
				.then(res=>res)
				.catch(err=>err)
			}
		}

5、redux-saga
	 saga vs chunk
	 // 特点
	 	> saga使用异步的方式更优雅，dispatch 的参数依然是一个纯粹的 action，而不是thunk function
		> 每一个 saga 都是 一个 generator function，代码采用 同步书写 的方式来处理 异步逻辑
		> 同样是受益于 generator function 的 saga 实现，代码异常/请求失败 都可以直接通过 try/catch 语法直接捕获处理
		> 业务逻辑被转移到单独的saga.js中，不会参杂在action.js或comment.js中

	(1) Effects
		> call: 用来创建 effect 对象，被称作是 effect factory
		> put: 相当于在 saga 中调用 store.dispatch(action)
		> take: 阻塞当前 saga，直到接收到指定的 action，代码才会继续往下执行，有种 Event.once() 事件监听的感觉
		> fork: 类似于 call effect，区别在于它不会阻塞当前 saga，如同后台运行一般，它的返回值是一个 task 对象
		> cancel：针对 fork 方法返回的 task ，可以进行取消关闭
		> select: 类似于getState

	(2) take
		yield takeLatest(pattern, saga, ...args)
		// takeLatest 不允许多个 saga 任务并行地执行。一旦接收到新的发起的 action，它就会取消前面所有 fork 过的任务（如果这些任务还在执行的话）
		yield takeEvery(pattern, saga, ...args)
		// takeEvery 的情况中，被调用的任务无法控制何时被调用， 它们将在每次 action 被匹配时一遍又一遍地被调用。并且它们也无法控制何时停止监听
    	yield take(pattern)
		// 最基础的监听

	(3) fork
		import {take, call, put, cancelled} from 'redux-sage/effects'
		import Api from '...'

		function *loginFlow(){
			const {user, password} = yield take('LOGIN_REQUEST')
			const task = yield fork(authorize, user, password)
			const action = yield take(['LOGOUT', 'LOGIN_ERROR'])
			if (action.type === 'LOGOUT') {
				yield cancle(task)
			}
			yield call(Api.clearItem('token'))
		}

		function *authorize(user, password) {
			try {
				const token = yield call(Api.authorize, user, password)
				yield put({type: 'LOGIN_SUCCESS', token})
				yield call(Api.storeItem, {token})
				return token
			}catch(error) {
				yield put({type: 'LOGIN_ERROR', error})
			} finally {
				if (yield cancelled()) {  // 如果外部取消
					...
				}
			}
		}

	(4)	 多个任务同时执行
		const [users, repos] = yield [call(fetch, '/users'), call(fetch, '/repos')]
		// generator 会被阻塞直到所有的 effects 都执行完毕，或者当一个 effect 被拒绝 

	(5) 组合器
		1) race
			多个任务同时执行时，有一个任务 promise或preject 就立即返回，然后会取消那些还没完成的任务
			yield race([call(), call(), ...])
			yield race({call1: call(), call2: call()...})

			const {posts, timeout} = yield race({
				posts: call(fetch, '/posts'),
				timeout: call(delay, 1000)
			})
			// 模仿超时
			if (posts) {
				put({type: 'POSTS_RECEIVED', posts})
			}else{
				put({type: 'TIMEOUT_ERROR'})
			}

		2) all
			用来命令 middleware 并行地运行多个 Effect，并等待它们全部完成
			yield all([call(), call(), ...])
			yield all({call1: call(), call2: call()...})

	(6) yield*
		function* playLevelOne() { ... }

		function* game() {
		const score1 = yield* playLevelOne()
		yield put(showScore(score1))
		}

	(7) yield call(generator)
		当 yield 一个 call 至 Generator，Saga 将等待 Generator 处理结束， 然后以返回的值恢复执行
		yield call(genertor)
		yield [call(genertor), call(genertor1)]

	(8) cancel
		一旦任务被 fork，可以使用 yield cancel(task) 来中止任务执行。取消正在运行的任务

```

### react-redux
```
展示组件、容器组件、其他组件

1、展示组件
	> 只负责 UI 的呈现，不带有任何业务逻辑
	> 没有状态（即不使用this.state这个变量）
	> 所有数据都由参数（this.props）提供
    > 不使用任何 Redux 的 API

2、容器组件
	容器组件就是使用 store.subscribe() 从 Redux state 树中读取部分数据，并通过 props 来把这些数据提供给要渲染的组件
	> 负责管理数据和业务逻辑，不负责 UI 的呈现
	> 带有内部状态
	> 使用 Redux 的 API
	(1) mapStateToProps
		用来建立一个从state对象props对象的映射关系

		mapStateToProps是一个函数，第一个参数是state对象，第二个参数是props对象
		返回一个对象，里面的每一个键值对就是一个映射

		connect方法可以省略mapStateToProps参数，UI 组件就不会订阅Store，就是说 Store 的更新不会引起 UI 组件的更新

		const mapStateToProps = (state, props) => {
			return {
				todos: getVisibleTodos(state.todos, state.visibilityFilter),
				action: state.filter ===state.visibilitityFilter
			}
		}

	(2) mapDispatchToProps
		用来建立 UI 组件的参数到store.dispatch方法的映射

		mapDispatchToProp是一个函数，有两个参数dispatch和props

		const mapDispatchToProps = (dispatch, props)=>{
			return {
				onClick: ()=>{
					dispatch({
						type: 'SET_VISIBILITY_FILTER',
						filter: props.filter
					})
				}
			}
		}

		// bindActionCreators
		

	(3) connect
		使用 connect() 创建 VisibleTodoList，并传入上面两个函数
		TodoList是展示组件

		import { connect } from 'react-redux'
		const VisibleTodoList = connect(
			mapStateToProps,
			mapDispatchToProps
		)(TodoList)

		export default VisibleTodoList

3、provider组件
	provider的原理是React组件的context实现的

	import { Provider } from 'react-redux'
	import { createStore } from 'redux'
	import todoApp from './reducers'
	import App from './components/App'

	let store = createStore(todoApp);

	render(
	<Provider store={store}>
		<App />
	</Provider>,
	document.getElementById('root')
	)


```

### Immutable
```
Immutable 实现的原理是 Persistent Data Structure（持久化数据结构），也就是使用旧数据创建新数据时，要保证旧数据同时可用且不变

同时为了避免 deepCopy 把所有节点都复制一遍带来的性能损耗，Immutable 使用了 Structural Sharing（结构共享），即如果对象树中一个节点发生变化，只修改这个节点和受它影响的父节点，其它节点则进行共享

```

### refs
```
Ref 转发是一项将 ref 自动地通过组件传递到其一子组件的技巧

// 1
const FancyButton = React.forwardRef((props, ref) => (
	<button ref={ref} className="FancyButton">
	    {props.children}
  	</button>
))

// 可以直接获取 DOM button 的 ref：
const ref = React.createRef();
<FancyButton ref={ref}>Click me!</FancyButton>;
// ref 挂载完成，ref.current 将指向 <button> DOM 节点


```

### react生命周期
```
			Mounting							Updating					Unmounting
				|																|
				|					NewProps	setState	forceUpdate			|
				↓						|			|			|				|			
			constructor					|			|			|				|
				|						|			|			|				|
				↓						↓			↓			↓				|
			+--------------------------------------------------------+			|
			|				getDerivedStateFromProps  				 |			|
			+--------------------------------------------------------+			|
				|						↓			↓			|				|
				|				 +--------------------------+	|				|
				|				 |	shouldComponentUpdate   |	|				|
				|				 +--------------------------+	|				|
				↓							  ↓ true/false		↓				|
			+--------------------------------------------------------+			|
			|							render						 |			|
			+--------------------------------------------------------+			|
				|									↓							|
				|					+---------------------------------+			|
				|					|		getSnapshotBeforeUpdate	  |			|
				|					+---------------------------------+			|
				|									|							|
				↓									↓							|
			+--------------------------------------------------------+			|
			|				React updates DOM and refs				 |			|
			+--------------------------------------------------------+	 		|
				|									|							|
				↓									↓							↓
			componentDidMount				componentDidUpdate			componentWillUnmount


// mount(挂载)
	当组件实例被创建并插入 DOM 中时，其生命周期调用顺序如下
	(1) constructor(props)
		1) 可以忽略
			如果不初始化 state 或不进行方法绑定，则不需要为 React 组件实现构造函数
		2) 禁止setState
			在 constructor() 函数中不要调用 setState() 方法。如果你的组件需要使用内部 state，请直接在构造函数中为 this.state 赋值初始 state
		3) 避免将props的值复制给state
			props的改变不会改变state

	(2) static getDerivedStateFromProps(props, state)
		在调用 render 方法之前调用，并且在初始挂载及后续更新时都会被调用
		它应"返回一个对象"来更新 state，如果返回 null 则不更新任何内容
		此为静态函数，不能通过this访问成员

	(3) render()
	(4) componentDidMount()
		componentDidMount()会在组件挂载后(插入 DOM 树中)立即调用
		1) 调用setState的代价
			在 componentDidMount() 里直接调用 setState()，它将触发额外渲染，但此渲染会发生在浏览器更新屏幕之前
			如此保证了即使在 render() 两次调用的情况下，用户也不会看到中间状态请谨慎使用该模式，因为它会导致性能问题
			通常应该在 constructor() 中初始化 state

// update(更新)
	当组件的 props 或 state 发生变化时会触发更新。组件更新的生命周期调用顺序如下
	(1) static getDerivedStateFromProps()
		props的改变会执行这个函数，state的改变不会
	(2) shouldComponentUpdate(nextProps, nextState, nextContext)
		1) 当 props 或 state 发生变化时，shouldComponentUpdate() 会在渲染执行之前被调用。返回值默认为 true。如果为false将不会渲染
		2) 首次渲染或使用 forceUpdate() 时不会调用该方法
		3) 考虑使用内置的 PureComponent 组件
			此方法仅作为性能优化的方式而存在。不要企图依靠此方法来"阻止"渲染，因为这可能会产生 bug。你应该考虑使用内置的 PureComponent 组件，而不是手动编写 shouldComponentUpdate()

	(3) render()
	(4) getSnapshotBeforeUpdate(prevProps, prevState)
		这函数会在render之后执行，而执行之时DOM元素还没有被更新，给了一个机会去获取DOM信息
		getSnapshotBeforeUpdate把snapshot返回，然后DOM改变，然后snapshot传递给componentDidUpdate

	(5) componentDidUpdate(prevProps, prevState, snapshot)
		componentDidUpdate() 会在更新后会被立即调用。首次渲染不会执行此方法
		1) 调用setState需要注意
			在 componentDidUpdate() 中直接调用 setState()，但请注意它必须被包裹在一个条件语句里
				否则可能会导致死循环
				可能会导致额外的重新渲染，影响组件性能
			不要将props镜像给state，请考虑直接使用 props
		2) 如果 shouldComponentUpdate() 返回值为 false，则不会调用 componentDidUpdate()
		3) 尽量使用 componentDidUpdate 生命周期，因为它保证每次更新只调用一次
// 卸载
	当组件从 DOM 中移除时会调用如下方法：
	componentWillUnmount()
		componentWillUnmount() 会在组件卸载及销毁之前直接调用
		1) 不应该调用setState
			componentWillUnmount() 中不应调用 setState()，因为该组件将永远不会重新渲染。组件实例卸载后，将永远不会再挂载它

// 错误处理
	当渲染过程，生命周期，或子组件的构造函数中抛出错误时，会调用如下方法：
	static getDerivedStateFromError()
	componentDidCatch()

// 其他API
	forceUpdate()
		当组件的 state 或 props 发生变化时，组件将重新渲染。如果 render() 方法依赖于其他数据，则可以调用 forceUpdate() 强制让组件重新渲染
		调用 forceUpdate() 将致使组件调用 render() 方法，此操作会跳过该组件的 shouldComponentUpdate()。但其子组件会触发正常的生命周期方法，包括 shouldComponentUpdate() 方法
		应该避免使用 forceUpdate()
	setState()
		setState() 并不总是立即更新组件，调用 setState() 后立即读取 this.state 成为了隐患
		请使用 componentDidUpdate 或者 setState 的回调函数（setState(updater, callback)）


```

### Component 和 PureComponent
```
React.PureComponent 是一个和 React.Component 几乎相同，唯一不同的是 React.PureComponent 帮助我们完成了 shouldComponentUpdate 的一些交浅的比较

React创建了PureComponent组件创建了默认的shouldComponentUpdate行为。这个默认的shouldComponentUpdate行为会一一比较props和state中所有的属性，只有当其中任意一项发生改变是，才会进行重绘
```

### Component 和 FC
```
interface InitProps {

}

interface InitState {

}

// 写法一，类的写法
class App extend React.Component<InitProps, Initstate> {

}

// 写法二，函数的写法
// 注意：函数类型组件要干净，仅通过props来渲染
const App: React.FC<InitProps> = (props)=>{

}

```

### ComponentClass 和 FC 类型的声明
```
interface T {
	name: string
}

const App: React.ComponentClass<T> | React.FC<T>

// 实例
	interface AutoHeightProps {
		height?: number;
	}

	// 模板函数  <>():type => {}
	const App = <P extends AutoHeightProps>(
		WrappedComponent: React.ComponentClass<P> | React.FC<P>,
	): React.ComponentClass<P> => {
		class AutoHeightComponent extends React.Component<P> {
			state = {

			};
			componentDidMount(){

			};
			render(){
				return <>...
			}
		}

		return AutoHeightComponent
	}
```

### Fragments
```
public render(){
  return (
    <React.Fragment>
      <div></div>
      <div></div>
    </React.Fragment>
  )
}

//or

public render(){
  return (
    <>
      <div></div>
      <div></div>
    </>
  )
}
```

### memo
```
// useMemo
	用于性能优化
	// 例如
	const App = ()=>{
		const [index, setIndex] = useState(0);
		const [str, setStr] = useState('');
		const add = ()=>{
			return index * 100;
		}
		return (
			<div>{index}-{str}-{add()}</div>
		)
	}
	// 无论如何修改 index 或 str 都会引发 add() 的执行，这对于性能来说是很难接受的
	
	// add() 只依赖于 index ，因此我们可以使用 useMemo 来优化此项
	const App = ()=>{
		const [index, setIndex] = useState(0);
		const [str, setStr] = useState('');
		const add = useMemo(()=>{
			return index * 100;
		}, [index])
		return (
			<div>{index}-{str}-{add}</div>
		)
	}
	// 此时add只依赖index的改变而渲染

// useCallback
	那么 useCallback 的使用和 useMemo 比较类似，但它返回的是缓存函数

// React.memo
	React.memo 为高阶组件。它与 React.PureComponent 非常相似，但只适用于函数组件，而不适用 class 组件
	
	function MyComponent(props) {
	/* 使用 props 渲染 */
	}

	function areEqual(prevProps, nextProps) {
	/*
	如果把 nextProps 传入 render 方法的返回结果与
	将 prevProps 传入 render 方法的返回结果一致则返回 true，
	否则返回 false
	*/
	}

	export default React.memo(MyComponent, areEqual);

	//注意
	与 class 组件中 shouldComponentUpdate() 方法不同的是，如果 props 相等，areEqual 会返回 true；如果 props 不相等，则返回 false。这与 shouldComponentUpdate 方法的返回值相反。
```

### ReactNode
```
React.ReactNode 是组件的返回值

ReactNode 可以是 ReactElement, ReactFragment, string, number, boolean, 数组 ReactNodes, null, undefined

render(): React.ReactNode {

}
```

### Hook

```
useState： setState
useReducer： setState
useRef: ref
useImperativeMethods: ref
useContext: context
useCallback: 可以对setState的优化
useMemo: useCallback的变形
useLayoutEffect: 类似componentDidMount/Update, componentWillUnmount
useEffect: 类似于setState(state, cb)中的cb，总是在整个更新周期的最后才执行

// useState

// useContext

// useReducer

// useEffect
	https://juejin.im/post/5c9827745188250ff85afe50

	Function Component 是更彻底的状态驱动抽象，甚至没有 Class Component 生命周期的概念，只有一个状态，而 React 负责同步到 DOM

	useEffect Hook 看做 componentDidMount，componentDidUpdate 和 componentWillUnmount 这三个函数的组合

	(1) 避免无限执行
		const [data, setData] = useState()
		useEffect(()=>{
			setData(...)
		})
		这个会无限循环

	(2) 仅在mount时执行
		const [data, setData] = useState()
		useEffect(()=>{
			setData(...)
		}, [])

	(3) 仅在data不同时才执行
		const [data, setData] = useState()
		useEffect(()=>{
			setData(...)
		}, [data])
		
// useLayoutEffect


// useCallback
```
