package web

import "net/http"

// 确保一定实现了 Server 接口，这点常在开源框架中看到
//var _ Server = &HTTPServer{}

type Context struct {
	Req  *http.Request
	Resp http.ResponseWriter
}

type HandleFunc func(ctx *Context)

type Server interface {
	http.Handler
	// Start 组合 http.Handler 并且增加 Start 方法。
	// Start 启动服务器
	// addr 是监听地址，如果只指定端口，可以使用 “:8081”
	// 或者 “localhost:8082”
	Start(addr string) error

	// AddRoute 注册一个路由
	// method 是 HTTP 方法
	// path 是路径，必须以 / 为开头
	// AddRoute 最终会和路由树交互，我们后面再考虑
	AddRoute(method string, path string, handler HandleFunc)
	//我们并不采取这种设计方案
	//AddRoute1(method string, path string, handlers ...HandleFunc)
}

type HTTPServer struct {
	*router
}

func NewHTTPServer() *HTTPServer {
	return &HTTPServer{
		router: newRouter(),
	}
}

//func (s *HTTPServer) AddRoute(method string, path string, handler HandleFunc) {
//	panic("implement me")
//}

// Post 针对不同 HTTP 方法的注册 API，都可以委托给 Handle 方法。这种设计思路很常用。
func (s *HTTPServer) Post(path string, handler HandleFunc) {
	s.AddRoute(http.MethodPost, path, handler)
}

func (s *HTTPServer) Get(path string, handler HandleFunc) {
	s.AddRoute(http.MethodGet, path, handler)
}

/*
ServeHTTP 则是我们整个 Web 框架的核心入口。我们将在整个方法内部完成：
• Context 构建
• 路由匹配
• 执行业务逻辑
*/
func (s *HTTPServer) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	ctx := &Context{
		Req:  request,
		Resp: writer,
	}
	s.serve(ctx)
}

// Start
/*
该实现直接使用 http.ListenAndServe 来启动，后续可以根据需要替换为：
• 内部创建 http.Server 来启动
• 使用 http.Serve 来启动，换取更大的灵活性，如将端口监听和服务器启动分离等
*/
func (s *HTTPServer) Start(addr string) error {
	return http.ListenAndServe(addr, s)
}

func (s *HTTPServer) serve(ctx *Context) {

}
