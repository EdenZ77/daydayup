package main

import (
	"fmt"
	"reflect"
	"time"
)

type Config struct {
	Timeout   time.Duration
	Endpoints []string
	Metadata  map[string]interface{}
}

func main() {
	current := Config{
		Timeout:   5 * time.Second,
		Endpoints: []string{"api1", "api2"},
		Metadata:  map[string]interface{}{"version": 1.2},
	}

	updated := Config{
		Timeout:   5 * time.Second,
		Endpoints: []string{"api1", "api2"},
		Metadata:  map[string]interface{}{"version": 1.2},
	}

	// 检查配置是否变更
	if !DeepEqual(current, updated) {
		fmt.Println("配置已变更，需要更新")
	} else {
		fmt.Println("配置未变更")
	}
}

func DeepEqual(a, b interface{}) bool {
	// 处理nil值
	if a == nil || b == nil {
		return a == b
	}

	ta, tb := reflect.TypeOf(a), reflect.TypeOf(b)
	if ta != tb {
		return false
	}

	va, vb := reflect.ValueOf(a), reflect.ValueOf(b)

	// 根据类型分类处理
	switch ta.Kind() {
	case reflect.Struct:
		for i := 0; i < ta.NumField(); i++ {
			field := ta.Field(i)
			fa := va.FieldByName(field.Name)
			fb := vb.FieldByName(field.Name)

			// 跳过不可导出字段
			if !fa.CanInterface() || !fb.CanInterface() {
				continue
			}

			if !DeepEqual(fa.Interface(), fb.Interface()) {
				return false
			}
		}
		return true

	case reflect.Ptr:
		// 比较指针指向的值
		if va.IsNil() || vb.IsNil() {
			return va.IsNil() == vb.IsNil()
		}
		return DeepEqual(va.Elem().Interface(), vb.Elem().Interface())

	case reflect.Slice, reflect.Array:
		// 比较长度
		if va.Len() != vb.Len() {
			return false
		}
		// 比较每个元素
		for i := 0; i < va.Len(); i++ {
			if !DeepEqual(va.Index(i).Interface(), vb.Index(i).Interface()) {
				return false
			}
		}
		return true

	case reflect.Map:
		// 比较长度
		if va.Len() != vb.Len() {
			return false
		}
		// 比较键值对
		for _, key := range va.MapKeys() {
			av := va.MapIndex(key)
			bv := vb.MapIndex(key)
			if !bv.IsValid() || !DeepEqual(av.Interface(), bv.Interface()) {
				return false
			}
		}
		return true

	case reflect.Interface:
		// 解包接口值
		return DeepEqual(va.Elem().Interface(), vb.Elem().Interface())

	default:
		// 基本类型直接比较
		return a == b
	}
}
