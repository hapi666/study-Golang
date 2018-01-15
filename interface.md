#### 接口(interface)

接口从本质上说就是几个方法的集合，接口描述了类型的行为

接口的命名一般是：...r/...able等(由r/able后缀组成)

在Go语言中接口可以有值，一个接口类型的变量就是一个接口值(var ai Namer)，接口值是一个多字数据结构，其初始化值为nil，虽然和指针不是完全相同的东西，但实际上是一个指针。故指向接口值的指针是非法的，指向接口的指针不仅一点用处没有，还会导致代码错误。所以在函数参数列表里面用接口指针是错误的！

对于任何一个类型，只要实现了接口中的所有方法即为实现了该接口。当然它还可以有其他方法。

类型不需要显式声明它实现了某个接口，接口隐式地被实现。

多个类型可以实现同一接口。

一个类型可以实现多个接口。

即使接口在类型之后才定义，二者处于不同的包中，被单独编译：只要类型实现了接口中的方法，它就实现了此接口。

```go
package main

import (
	"fmt"
)

type easy struct {
	value int
}

type easier interface {
	name(p1 int, p2 int) int
}

func (e *easy) name(p1 int, p2 int) int {
	fmt.Println(p1 * p2)
	return p1 * p2
}

func main() {
	test := new(easy)
	tester := easier(test)
	tester.name(6, 6)
}
```

输出：36

接口类型的数组，展示多态。

一个接口可以嵌套到另一个接口中。

假设varM是一个接口类型的变量，检测它的动态类型(运行时在变量中存储的值的实际类型)的方式是：v:=varM.(T)    //检测它是否是T类型

更安全的断言方式是这样的：

```go
if v,ok:=varM.(T);ok {
  Process(v)
  return
}

```

也可使用type-switch来断言,但是不允许使用fallthrough

```go
switch t := testinterface.(type) {
	case *Simple:
		fmt.Printf("Type Simple %T with value %v\n", t, t)
	case nil:
		fmt.Printf("nil value: nothing to check?\n")
	default:
		fmt.Printf("Unexpected type %T\n", t)
	}
```

测试一个值是否实现了某个接口，假定v是一个值(任意类型的值)：

```go
type Stringer interface {
String() string
}
if sv, ok := v.(Stringer); ok {
fmt.Printf("v implements String(): %s\n", sv.String()) // note: sv, not v
}
```

Print函数就是如此检测类型是否可以打印自身的

再来个例子

```go
package main
import (
"fmt"
)
type List []int
func (l List) Len() int {
return len(l)
}
func (l *List) Append(val int) {
*l = append(*l, val)
}
type Appender interface {
Append(int)
}
func CountInto(a Appender, start, end int) {
for i := start; i <= end; i++ {
a.Append(i)
}
}
type Lener interface {
Len() int
}
func LongEnough(l Lener) bool {
return l.Len()*10 > 42
}
func main() {
// A bare value
var lst List
// compiler error:
// cannot use lst (type List) as type Appender in argument to CountInto:
// List does not implement Appender (Append method has pointer receiver)
// CountInto(lst, 1, 10)
if LongEnough(lst) { // VALID:Identical receiver type
fmt.Printf("- lst is long enough\n")
}
// A pointer value
plst := new(List)
CountInto(plst, 1, 10) //VALID:Identical receiver type
if LongEnough(plst) {
// VALID: a *List can be dereferenced for the receiver
fmt.Printf("- plst is long enough\n")
}
}
```

讨论
在 lst 上调用 CountInto 时会导致一个编译器错误，因为 CountInto 需要一个
Appender ，而它的方法 Append 只定义在指针上。 在 lst 上调用 LongEnough 是可以的因
为 'Len' 定义在值上。
在 plst 上调用 CountInto 是可以的，因为 CountInto 需要一个 Appender ，并且它的方法 Append 定义在指针上。在 plst 上调用 LongEnough 也是可以的，因为指针会被自动解引用。
总结：
在接口上调用方法时，必须有和方法定义时相同的接收者类型或者是可以从具体类型 P 直接可以辨识的：
·指针方法可以通过指针调用
·值方法可以通过值调用
·接收者是值的方法可以通过指针调用，因为指针会首先被解引用
·接收者是指针的方法不可以通过值调用，因为存储在接口中的值没有地址
将一个值赋值给一个接口赋值时，编译器会确保所有可能的接口方法都可以在此值上被调用，因此不正确的赋值在编译期就会失败。

Go 语言规范定义了接口方法集的调用规则：
类型 T 的可调用方法集包含接受者为 T 或 T 的所有方法集
类型 T 的可调用方法集包含接受者为 T 的所有方法
类型 T 的可调用方法集不包含接受者为 *T 的方法