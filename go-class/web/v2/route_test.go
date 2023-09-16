package web

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"reflect"
	"testing"
)

func TestRouter_AddRoute(t *testing.T) {
	testRoutes := []struct {
		method string
		path   string
	}{
		{
			method: http.MethodGet,
			path:   "/",
		},
		{
			method: http.MethodGet,
			path:   "/user",
		},
		{
			method: http.MethodGet,
			path:   "/user/home",
		},
		{
			method: http.MethodGet,
			path:   "/order/detail",
		},
		{
			method: http.MethodPost,
			path:   "/order/create",
		},
		{
			method: http.MethodPost,
			path:   "/login",
		},
	}
	mockHandler := func(ctx *Context) {}
	got := newRouter()
	for _, tr := range testRoutes {
		got.addRoute(tr.method, tr.path, mockHandler)
	}

	wantRouter := &router{
		trees: map[string]*node{
			http.MethodGet: {path: "/", children: map[string]*node{
				"user": {path: "user", children: map[string]*node{
					"home": {path: "home", handler: mockHandler},
				}, handler: mockHandler},
				"order": {path: "order", children: map[string]*node{
					"detail": {path: "detail", handler: mockHandler},
				}},
			}, handler: mockHandler},
			http.MethodPost: {path: "/", children: map[string]*node{
				"order": {path: "order", children: map[string]*node{
					"create": {path: "create", handler: mockHandler},
				}},
				"login": {path: "login", handler: mockHandler},
			}},
		},
	}
	msg, ok := wantRouter.equal(got)
	assert.True(t, ok, msg)

	// ============== 非法用例 ==================
	// 空字符串
	assert.PanicsWithValue(t, "web: 路由是空字符串", func() {
		got.addRoute(http.MethodGet, "", mockHandler)
	})
	// 前导没有 /
	assert.PanicsWithValue(t, "web: 路由必须以 / 开头", func() {
		got.addRoute(http.MethodGet, "a/b/c", mockHandler)
	})
	// 后缀有 /
	assert.PanicsWithValue(t, "web: 路由不能以 / 结尾", func() {
		got.addRoute(http.MethodGet, "/a/b/c/", mockHandler)
	})

}

func (r *router) equal(y router) (string, bool) {
	for k, v := range r.trees {
		yv, ok := y.trees[k]
		if !ok {
			return fmt.Sprintf("目标 router 里面没有方法 %s 的路由树", k), false
		}
		str, ok := v.equal(yv)
		if !ok {
			return k + "-" + str, ok
		}
	}
	return "", true
}

func (n *node) equal(y *node) (string, bool) {
	if y == nil {
		return "目标节点为 nil", false
	}
	if n.path != y.path {
		return fmt.Sprintf("%s 节点 path 不相等 x %s, y %s", n.path, n.path, y.path), false
	}

	nhv := reflect.ValueOf(n.handler)
	yhv := reflect.ValueOf(y.handler)
	if nhv != yhv {
		return fmt.Sprintf("%s 节点 handler 不相等 x %s, y %s", n.path, nhv.Type().String(), yhv.Type().String()), false
	}

	if len(n.children) != len(y.children) {
		return fmt.Sprintf("%s 子节点长度不等", n.path), false
	}
	//
	if len(n.children) == 0 {
		return "", true
	}

	for k, v := range n.children {
		yv, ok := y.children[k]
		if !ok {
			return fmt.Sprintf("%s 目标节点缺少子节点 %s", n.path, k), false
		}
		str, ok := v.equal(yv)
		if !ok {
			return n.path + "-" + str, ok
		}
	}
	return "", true
}
