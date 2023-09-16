package split

import (
	"reflect"
	"testing"
)

//func TestSplit(t *testing.T) {
//	got := Split("a:b:c", ":")
//	want := []string{"a", "b", "c"}
//	if !reflect.DeepEqual(want, got) {
//		t.Errorf("expected:%v, got:%v", want, got)
//	}
//}

func TestSplit(t *testing.T) {
	// 定义一个测试用例类型
	type test struct {
		input string
		sep   string
		want  []string
	}
	// v1 定义一个存储测试用例的切片
	// v2 测试用例使用map存储
	tests := map[string]test{
		"test1 simple":             {input: "a:b:c", sep: ":", want: []string{"a", "b", "c"}},
		"test2 wrong sep":          {input: "a:b:c", sep: ",", want: []string{"a:b:c"}},
		"test3 end more sep":       {input: "a//b//c//", sep: "/", want: []string{"a", "b", "c"}},
		"test4 start end more sep": {input: "//a//b//c//", sep: "//", want: []string{"a", "b", "c"}},
		"test5 more sep":           {input: "abcd", sep: "bc", want: []string{"a", "d"}},
		"test6 chinese test":       {input: "沙河有沙又有河", sep: "沙", want: []string{"河有", "又有河"}},
	}

	// 遍历切片，逐一执行测试用例
	//for _, tc := range tests {
	//	got := Split(tc.input, tc.sep)
	//	if !reflect.DeepEqual(got, tc.want) {
	//		//t.Errorf("expected:%v, got:%v", tc.want, got)
	//		t.Errorf("expected:%#v, got:%#v", tc.want, got)
	//	}
	//}

	for name, tc := range tests {
		// 使用 t.Run() 执行子测试
		t.Run(name, func(t *testing.T) {
			got := Split(tc.input, tc.sep)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("expected:%#v, got:%#v", tc.want, got)
			}
		})
	}
}

func TestMoreSplit(t *testing.T) {
	got := Split("abcd", "bc")
	want := []string{"a", "d"}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("expected:%v, got:%v", want, got)
	}
}
