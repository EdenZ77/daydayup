package main

import (
	"fmt"
	bbb "hello/protobuftest/address/pbs"

	"github.com/golang/protobuf/proto"
)

// protobuf demo
func main() {
	var cb bbb.ContactBook
	p1 := bbb.Person{
		Id:     3,
		Name:   "小王子",
		Gender: bbb.GenderType_MALE,
		Number: "7878778",
	}
	//fmt.Println(p1)
	cb.Persons = append(cb.Persons, &p1)
	// 序列化
	data, err := proto.Marshal(&p1)
	if err != nil {
		fmt.Printf("marshal failed,err:%v\n", err)
		return
	}
	//err = ioutil.WriteFile("./proto.dat", data, 0644)
	//if err != nil {
	//	return
	//}
	//
	//data2, err := ioutil.ReadFile("./proto.dat")
	//if err != nil {
	//	fmt.Printf("read file failed, err:%v\n", err)
	//	return
	//}
	//var p2 bbb.Person
	var p2 Person1
	proto.Unmarshal(data, &p2)
	fmt.Printf("%#v", p2)
}

// Person1 通过测试发现，当我们name属性与序列化对象的不一致时不影响protobuf反序列化，但是标识号必须一致
type Person1 struct {
	Id     int64          `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name   string         `protobuf:"bytes,21,opt,name=namexx,proto3" json:"name,omitempty"`
	Gender bbb.GenderType `protobuf:"varint,3,opt,name=gender,proto3,enum=GenderType" json:"gender,omitempty"`
	Number string         `protobuf:"bytes,4,opt,name=numberxx,proto3" json:"number,omitempty"`
}

func (x *Person1) Reset() {
	*x = Person1{}
}

func (x *Person1) String() string {
	return "person1"
}

func (*Person1) ProtoMessage() {}
