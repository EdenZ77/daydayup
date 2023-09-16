package web

type router struct {
	// trees 是按照HTTP方法来组织的
	// 如 GET => *node
	trees map[string]*node
}

func newRouter() *router {
	return &router{
		trees: map[string]*node{},
	}
}

func (r *router) AddRoute(method string, path string, handler HandleFunc) {
	panic("implement me")
}

type node struct {
	path string
	// children 子节点
	// 子节点的 path => node
	children map[string]*node
	// handler 命中路由之后执行的逻辑
	handler HandleFunc
}
