package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
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
}
