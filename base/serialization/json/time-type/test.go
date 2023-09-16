package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

/*
Go语言内置的 json 包使用 RFC3339 标准中定义的时间格式，对我们序列化时间字段的时候有很多限制。
*/
func main() {
	timeFieldDemo()
}

//type Post struct {
//	CreateTime time.Time `json:"create_time"`
//}

/*
也就是内置的json包不识别我们常用的字符串时间格式，如2023-06-01 12:25:42。
不过我们通过实现 json.Marshaler/json.Unmarshaler 接口来实现自定义的时间格式解析。
*/
//func timeFieldDemo() {
//	p1 := Post{CreateTime: time.Now()}
//	b, _ := json.Marshal(p1) //这里会输出RFC3339格式的时间
//
//	fmt.Printf("str:%s\n", b) // str:{"create_time":"2023-06-06T09:43:45.2318246+08:00"}
//	jsonStr := `{"create_time":"2020-04-05 12:25:42"}`
//	var p2 Post
//	//  反序列化时会报错
//	if err := json.Unmarshal([]byte(jsonStr), &p2); err != nil {
//		fmt.Printf("json.Unmarshal failed, err:%v\n", err)
//		return
//	}
//	fmt.Printf("p2:%#v\n", p2)
//}

// CustomTime ===================================
// ==============================================
type CustomTime struct {
	time.Time
}

const ctLayout = "2006-01-02 15:04:05"

var nilTime = (time.Time{}).UnixNano()

// UnmarshalJSON 实现了json.Unmarshaler接口中的方法
func (ct *CustomTime) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")
	if s == "null" {
		ct.Time = time.Time{}
		return
	}
	ct.Time, err = time.Parse(ctLayout, s)
	return
}

// MarshalJSON 实现了json.Marshaler接口中的方法
func (ct *CustomTime) MarshalJSON() ([]byte, error) {
	if ct.Time.UnixNano() == nilTime {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf("\"%s\"", ct.Time.Format(ctLayout))), nil
}

func (ct *CustomTime) IsSet() bool {
	return ct.UnixNano() != nilTime
}

type Post struct {
	CreateTime CustomTime `json:"create_time"`
}

func timeFieldDemo() {
	p1 := Post{CreateTime: CustomTime{time.Now()}}
	b, err := json.Marshal(p1)
	if err != nil {
		fmt.Printf("json.Marshal p1 failed, err:%v\n", err)
		return
	}
	fmt.Printf("str:%s\n", b)
	jsonStr := `{"create_time":"2020-04-05 12:25:42"}`
	var p2 Post
	if err := json.Unmarshal([]byte(jsonStr), &p2); err != nil {
		fmt.Printf("json.Unmarshal failed, err:%v\n", err)
		return
	}
	fmt.Printf("p2:%#v\n", p2)
}
