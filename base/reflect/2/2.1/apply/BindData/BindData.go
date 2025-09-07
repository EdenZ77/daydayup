package main

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

/*
实际应用场景
1.Web 框架请求处理：
•将 HTTP 请求参数自动绑定到结构体
•支持表单数据、查询参数、路径参数等

2.配置文件解析：
•将 INI/YAML/JSON 配置映射到结构体
•支持嵌套配置和类型转换

3.数据库结果集映射：
•将 SQL 查询结果绑定到结构体
•支持字段名映射和类型转换

4.命令行参数解析：
•将命令行参数绑定到配置结构体
•支持标志(flag)与结构体字段的映射
*/

// RegisterForm 定义用户注册结构体
type RegisterForm struct {
	Username string `form:"user"`
	Email    string `form:"email"`
	Age      int    `form:"age"`
	Accepted bool   `form:"accept"`
}

func main() {
	// 模拟来自 HTTP 请求的表单数据
	formData := map[string]string{
		"user":   "john_doe",
		"email":  "john@example.com",
		"age":    "30",
		"accept": "true",
		"extra":  "ignored", // 多余的字段会被忽略
	}

	// 创建目标结构体实例
	form := &RegisterForm{}

	// 执行数据绑定
	if err := BindData(form, formData); err != nil {
		fmt.Println("绑定错误:", err)
		return
	}

	// 输出绑定结果
	fmt.Printf("用户名: %s\n", form.Username)   // john_doe
	fmt.Printf("邮箱: %s\n", form.Email)       // john@example.com
	fmt.Printf("年龄: %d\n", form.Age)         // 30
	fmt.Printf("已接受条款: %t\n", form.Accepted) // true
}

func BindData(dest interface{}, data map[string]string) error {
	v := reflect.ValueOf(dest).Elem()
	t := v.Type()

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		key := field.Tag.Get("form")
		if key == "" {
			key = strings.ToLower(field.Name)
		}

		if value, ok := data[key]; ok {
			fieldValue := v.Field(i)

			// 根据字段类型进行转换和赋值
			switch fieldValue.Kind() {
			case reflect.String:
				fieldValue.SetString(value)

			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				if intValue, err := strconv.ParseInt(value, 10, 64); err == nil {
					fieldValue.SetInt(intValue)
				} else {
					return fmt.Errorf("字段 %s 转换 int 失败: %v", field.Name, err)
				}

			case reflect.Bool:
				if boolValue, err := strconv.ParseBool(value); err == nil {
					fieldValue.SetBool(boolValue)
				} else {
					return fmt.Errorf("字段 %s 转换 bool 失败: %v", field.Name, err)
				}

			case reflect.Float32, reflect.Float64:
				if floatValue, err := strconv.ParseFloat(value, 64); err == nil {
					fieldValue.SetFloat(floatValue)
				} else {
					return fmt.Errorf("字段 %s 转换 float 失败: %v", field.Name, err)
				}

			// 可以添加更多类型支持...

			default:
				return fmt.Errorf("不支持的字段类型: %s", fieldValue.Kind())
			}
		}
	}
	return nil
}
