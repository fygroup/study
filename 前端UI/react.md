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
		{a==true && 
			<h1>aaaa</h1>
			<h1>aaaa</h1>
			<h1>aaaa</h1>
		}
	</div>
	
// 3
	{a == true 
	?	<div>
			<h1>aaaaaa</h1>
			<h1>aaaaaa</h1>
		</div>
	:	<div>
			<h1>bbbbbb</h1>
			<h1>bbbbbb</h1>
		</div>
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

// 2


```