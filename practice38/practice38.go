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
