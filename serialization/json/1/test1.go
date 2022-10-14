package main

import (
	"encoding/json"
	"fmt"
)

/*
这个示例主要展示了json序列化与反序列化操作，与此对应的就是protobuf序列化，两者之间有什么差异呢？
对于不同语言的程序，对象的表达方式不同，想要进行传输(可以传输给其他应用，或者传输到磁盘)，就需要序列化为字节流形式，在另一端使用反序列化转换为对象
*/

type Person struct {
	NameGood   string `json:"nameGood"`
	AgeGood    int64  `json:"age"`
	WeightGood float64
}

type Person1 struct {
	Name   string  `json:"name_good"`
	Age    int64   `json:"age"`
	Weight float64 `json:"weight11"`
}

func main() {
	p1 := Person{
		NameGood:   "七米",
		AgeGood:    18,
		WeightGood: 71.5,
	}
	// struct -> json string
	b, err := json.Marshal(&p1)
	if err != nil {
		fmt.Printf("json.Marshal failed, err:%v\n", err)
		return
	}
	fmt.Printf("str:%s\n", b)
	// json string -> struct
	var p2 Person1
	err = json.Unmarshal(b, &p2)
	if err != nil {
		fmt.Printf("json.Unmarshal failed, err:%v\n", err)
		return
	}
	fmt.Printf("p2:%#v\n", p2)
}
