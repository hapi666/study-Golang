package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"reflect"
)

type Host struct {
	Name string `json:"name"`
	Port Port   `json:"port"`
}

type Type int

const (
	Int Type = iota
	String
)

type Port struct {
	Type   Type
	IntVal int
	StrVal string
}

type Test interface {
	umter(int, string) int
}

func (Host) umter(a int, b string) int {
	return a
}

/*
*之前我一直特别疑惑，因为这是我从别人博客里看到的嘛，
*我不太理解：为什么实现了这个Unmarshaller接口里面的UnmarshalJSON()方法
*就是实现了能够自己预想的 反序列化/序列化 呢
*第二天。。。
*我翻阅源码（之前粗略翻阅并没有发现），忽然看到了在json包里面，

// To unmarshal JSON into a value implementing the Unmarshaler interface,
// Unmarshal calls that value's UnmarshalJSON method, including
// when the input is a JSON null.

*在反序列化之前是要判断这个待反序列化的json的值类型的（在d.value()这个方法里判断值类型）
*紧接着，判断好了就call对应的函数比方说如果是数组就call d.array() 这个函数
*在这个array函数里面就调用了UnmarshalJSON这个函数

*我又tm产生了一个疑惑。。那源码包里也没有实现了Unmarshaller的数据结构啊，
*那 那个u.UnmarshalJSON(d.data[start:d.off])调用的是啥呢？？
*。。。
*我终于看懂了，源码的流程是这样的：首先我们自己不是call Unmarshal()函数嘛，
*源码会call d.unmarshal(v)将v的动态值取出来，通过reflect判断它的动态值是不是指针或者是不是nil
*如果它的动态值不是指针，或者如果它是Nil，那么直接报错（InvalidUnmarshalError）意思是说这个参数是无效的
*因为啊，传递给Unmarshal()的参数必须是非空指针！
*接着，
*/

// 实现 json.Unmarshaller 接口
func (port *Port) UnmarshalJSON(value []byte) error {
	if value[0] == '"' {
		port.Type = String
		return json.Unmarshal(value, &port.StrVal)
	}
	port.Type = Int
	return json.Unmarshal(value, &port.IntVal)
}

// 实现 json.Marshaller 接口
func (port Port) MarshalJSON() ([]byte, error) {
	switch port.Type {
	case Int:
		return json.Marshal(port.IntVal)
	case String:
		return json.Marshal(port.StrVal)
	default:
		return []byte{}, fmt.Errorf("impossible Port.Type")
	}
}

func main() {
	j, err := ioutil.ReadFile("test.json")
	if err != nil {
		fmt.Println(err)
	}
	test := &Host{}
	err = json.Unmarshal(j, test)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(test)
	host := &Host{
		Name: "yulibaozi",
		Port: Port{
			Type:   Int,
			IntVal: 8080,
		},
	}
	var t Test
	t = host
	f, ok := t.(*Host)
	if ok {
		fmt.Println(f)
	}
	fmt.Println(f.umter(1, "1"))
	var t1 *Host
	var interf Test
	interf = t1
	if interf == nil {
		fmt.Println("yes!")
	}
	//rr := &Host{}
	// var i Test = rr
	v := reflect.ValueOf(interf)
	vv := v.Elem()
	// vv := reflect.ValueOf(v)
	vvv := reflect.ValueOf(vv)
	/*
	   Elem()是对reflect.Value类型的变量不断解引用
	*/
	fmt.Println(v.Kind())
	fmt.Println(vv.Kind())
	fmt.Println(vv.IsValid())
	fmt.Println(vvv.Kind())
	fmt.Println(vvv.IsValid())
	a := 2
	x := reflect.ValueOf(&a).Elem()

	fmt.Println(x.Addr())
	// fmt.Println(vv.Elem().Kind())
	// ee := reflect.ValueOf(rr)
	// if ee.Kind() == reflect.Ptr {
	// 	fmt.Println(ee)
	// }
	var as = []string{"hi ", "hapi!"}
	fmt.Println(as)
}
