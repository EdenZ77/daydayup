package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
)

// 参考资料： https://mp.weixin.qq.com/s/js3m_Fe6k4ys4aBSFDJhLQ
func main() {
	//jsonDemo()
	decoderDemo()
}

func jsonDemo() {
	// map[string]interface{} -> json string
	var m = make(map[string]interface{}, 1)
	m["count"] = 1 // int
	b, _ := json.Marshal(m)

	fmt.Printf("str:%#v\n", string(b)) // str:"{\"count\":1}"

	// json string -> map[string]interface{}
	var m2 map[string]interface{}
	json.Unmarshal(b, &m2)

	fmt.Printf("value:%v\n", m2["count"]) // 1
	fmt.Printf("type:%T\n", m2["count"])  // float64
}

func decoderDemo() {
	// map[string]interface{} -> json string
	var m = make(map[string]interface{}, 1)
	m["count"] = 1 // int
	b, _ := json.Marshal(m)

	fmt.Printf("str:%#v\n", string(b)) // str:"{\"count\":1}"
	os.Stdout.Write(b)                 // {"count":1}
	fmt.Println()
	// json string -> map[string]interface{}
	var m2 map[string]interface{}
	// 使用decoder方式反序列化，指定使用number类型
	decoder := json.NewDecoder(bytes.NewReader(b))
	decoder.UseNumber()
	decoder.Decode(&m2)

	fmt.Printf("value:%v\n", m2["count"]) // 1
	fmt.Printf("type:%T\n", m2["count"])  // json.Number

	// 将m2["count"]转为json.Number之后调用Int64()方法获得int64类型的值
	// 我们在处理number类型的json字段时需要先得到json.Number类型，然后根据该字段的实际类型调用Float64()或Int64()。
	count, _ := m2["count"].(json.Number).Int64()

	fmt.Printf("type:%T\n", int(count)) // int
}
