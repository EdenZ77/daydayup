package main

import "fmt"

// Request 是传递给拦截器的请求对象
type Request struct {
	Data string
}

// Response 是拦截器处理后返回的响应对象
type Response struct {
	Data string
}

// Filter 是拦截器的接口，拦截器可以处理传入请求和输出响应
type Filter interface {
	DoFilter(request *Request, response *Response, chain *FilterChain)
}

// FilterChain 定义了拦截器链的执行逻辑
type FilterChain struct {
	filters []Filter
	index   int
}

// AddFilter 向链中添加一个拦截器
func (fc *FilterChain) AddFilter(filter Filter) {
	fc.filters = append(fc.filters, filter)
}

// DoFilter 执行拦截器链。如果链中还有拦截器未执行，就递归地调用下一个拦截器。
func (fc *FilterChain) DoFilter(request *Request, response *Response) {
	if fc.index == len(fc.filters) {
		// 所有拦截器都已执行，此处可以执行链末尾的逻辑
		return
	}

	// 获取下一个拦截器并增加索引
	currentFilter := fc.filters[fc.index]
	fc.index++

	// 执行当前拦截器的逻辑
	currentFilter.DoFilter(request, response, fc)
}

// LoggingFilter 打印请求和响应信息
type LoggingFilter struct{}

func (f *LoggingFilter) DoFilter(request *Request, response *Response, chain *FilterChain) {
	fmt.Println("LoggingFilter: 请求信息:", request.Data)

	// 调用链中的下一个拦截器
	chain.DoFilter(request, response)

	fmt.Println("LoggingFilter: 响应信息:", response.Data)
}

// AuthenticationFilter 模拟认证检查
type AuthenticationFilter struct{}

func (f *AuthenticationFilter) DoFilter(request *Request, response *Response, chain *FilterChain) {
	fmt.Println("AuthenticationFilter: 开始认证")

	// 模拟认证逻辑
	if request.Data != "authenticated" {
		response.Data = "认证失败"
		fmt.Println("AuthenticationFilter: 认证失败")
		return // 认证失败时停止链的执行
	}

	// 认证成功则继续执行链中的下一个拦截器
	chain.DoFilter(request, response)

	fmt.Println("AuthenticationFilter: 认证成功")
}

// ModificationFilter 修改请求和响应数据
type ModificationFilter struct{}

func (f *ModificationFilter) DoFilter(request *Request, response *Response, chain *FilterChain) {
	fmt.Println("ModificationFilter: 修改请求数据")
	request.Data += " modified"

	// 调用链中的下一个拦截器
	chain.DoFilter(request, response)

	fmt.Println("ModificationFilter: 修改响应数据")
	response.Data += " modified"
}

func main() {
	// 创建请求和响应实例
	request := &Request{Data: "authenticated"}
	response := &Response{}

	// 创建拦截器链并添加拦截器
	chain := FilterChain{}
	chain.AddFilter(&LoggingFilter{})
	chain.AddFilter(&AuthenticationFilter{})
	chain.AddFilter(&ModificationFilter{})

	// 执行拦截器链
	chain.DoFilter(request, response)

	fmt.Println("响应最终数据:", response.Data)
}
