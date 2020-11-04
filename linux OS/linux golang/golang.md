### context.WithCancel
```
// 父协程控制子协程退出

fmt.Println("main 函数 开始...")
go func() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	fmt.Println("父 协程 开始...")
	go func(ctx context.Context) {
		for {
			for {
				select {
				case <-ctx.Done():
					fmt.Println("子 协程 接受停止信号...")
					return
				default:
					fmt.Println("子 协程 执行中...")
					timer := time.NewTimer(time.Second * 2)
					<-timer.C
				}
			}
		}
	}(ctx)
	time.Sleep(time.Second*5)
	fmt.Println("父 协程 退出...")
}()
time.Sleep(time.Second*10)
fmt.Println("main 函数 退出")
```

### 任意类型 T() 都能够调用 *T 的方法吗？反过来呢？
```
(1) *T 调用 T()
	*T类型的值可以调用为类型T声明的方法，这是因为解引用指针总是合法的
	func(t T) f(){}
	var a *T = new(T)
	a.f()	// 合法

(2) T 调用 *T()
	T类型的值可以调用为类型*T声明的方法，但是仅限于T的值可寻址
	编译器在调用指针属主方法前，会自动取此T值的地址。因为不是任何T值都是可寻址的，所以并非任何T值都能够调用为类型*T声明的方法
	func(t *T) f() {}
	T a
	a.f()	// 合法，前提是a可寻址

(3) 不可寻址的值
	字符串中的字节
	map 对象中的元素(slice 对象中的元素是可寻址的，slice的底层是数组)
	常量
	包级别的函数等

	type T string
	func(t *T) f(){}

	var a T = "dasdsa"
	a.f()	// 正确
	const a1 T = "dadas"
	a1.f()	// 错误，常量不可寻址
	
```

### const group 自动补全
```
const (
	a, b = "golang", 100
	d, e
	f bool = true
	g
)

// 自动补全
const (
	a, b = "golang", 100
	d, e = "golang", 100
	f bool = true
	g bool = true
)
```

### 无类型常量和有类型常量
```
const N = 100
var x int = N	// 正确

const M int32 = 100
var y int = M	// 错误

无类型常量，赋值给其他变量时，如果字面量能够转换为对应类型的变量，则赋值成功
有类型的常量，赋值给其他变量时，需要类型匹配才能成功，所以显示地类型转换

var y int = int(M) // 正确

```

### 类型转换
```
Go语言不存在隐式类型转换，因此所有的类型转换都必须显式的声明

type MyInt int
var a int = 1
var  MyInt = MyInt(a)	// 必须显示转换

```

### 常量的定义
```
常量的值必须在编译期间确定

const (
	a int = 1
	b error = error.New("dasdaa")	// 错误
)

```

### 常量转换不允许溢出
```
func main() {
	var a int8 = -1
	var b int8 = -128 / a
	fmt.Println(b)	// -128
}


func main() {
	const a int8 = -1
	var b int8 = -128 / a
	fmt.Println(b)	
}
编译失败：constant 128 overflows int8

```

### defer延迟调用
```
// defer 延迟调用时，需要保存函数指针和参数，因此链式调用的情况下，除了最后一个函数外都会在调用时直接执行
type T struct{}
func(t T) f(n int){
	fmt.Println(n)
	return t
}
func main(){
	var t T
	defer t.f(1).f(2)
	fmt.Println(3)
}
结果：132


// defer 语句执行时，会将需要延迟调用的函数和参数保存起来
func f(i int){
	defer fmt.Println(i)
	i += 100
}
func main(){
	f(1)
}
结果：1

func f(i int){
	defer func(){
		fmt.Println(i)
	}()
	i += 1
}
结果：101

// defer 的作用域是函数，而不是代码块
func main(){
	a := 1
	if a == 1 {
		defer fmt.Println(a)
		a += 100
	}
	fmt.Println(a)
}
结果：101 1
```

### sync.Pool
```
type Pool struct {
	...
	New func() interface{}
}

// 从 Pool 中获取元素，当Pool中没有元素时，会调用 New 生成元素，新元素不会放入 Pool 中，若 New 未定义，则返回 nil
func (p *Pool) Get() interface{}

//往 Pool 中添加元素 x
func (p *Pool) Put(x interface{})

type A struct{}

var bufPool = sync.Pool{
	New: func()interface{}{
		return new(A)
	}
}

b, _ = bufPool.Get().(*A)

bufPool.Put(b)
```

### 指针
```
go中不能对指针进行自增或自减运算
不能对指针进行下标运算
```

### 接口赋值
```
type A interface {
	f(int)error
}

type B struct{}
func (b B) f(int)error{
	return nil
}

var x A = B{}
var x A1 = &B{}
var x A2 = new(B)

```