package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

// 定义自定义键类型（避免键冲突的最佳实践）
type contextKey string

const (
	requestIDKey contextKey = "requestID"
	userKey      contextKey = "user"
	traceIDKey   contextKey = "traceID"
)

// User 用户信息结构
type User struct {
	ID    int
	Name  string
	Email string
}

// 模拟中间件：设置请求ID和TraceID
func middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 创建初始上下文
		ctx := r.Context()

		// 生成唯一请求ID
		reqID := "req-" + fmt.Sprintf("%d", time.Now().UnixNano())
		ctx = context.WithValue(ctx, requestIDKey, reqID)

		// 生成追踪ID
		traceID := "trace-" + fmt.Sprintf("%x", time.Now().UnixNano())
		ctx = context.WithValue(ctx, traceIDKey, traceID)

		// 记录请求开始
		logInfo(ctx, "请求开始", map[string]interface{}{
			"method": r.Method,
			"path":   r.URL.Path,
		})

		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)

		// 记录请求完成
		logInfo(ctx, "请求处理完成")
	})
}

// 模拟认证中间件
func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		// 模拟用户认证
		user := &User{
			ID:    123,
			Name:  "张明",
			Email: "zhangming@example.com",
		}

		// 将用户信息添加到上下文
		ctx = context.WithValue(ctx, userKey, user)
		logInfo(ctx, "用户认证成功", map[string]interface{}{
			"user_id": user.ID,
			"name":    user.Name,
		})

		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

// 业务处理函数
func handleRequest(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// 获取请求ID
	reqID, ok := ctx.Value(requestIDKey).(string)
	if !ok {
		reqID = "unknown"
	}

	// 获取用户信息
	user, ok := ctx.Value(userKey).(*User)
	if !ok {
		logError(ctx, "未认证用户尝试访问")
		http.Error(w, "未认证用户", http.StatusUnauthorized)
		return
	}

	// 获取追踪ID
	traceID, ok := ctx.Value(traceIDKey).(string)
	if !ok {
		traceID = "not-set"
	}

	// 记录业务处理开始
	logInfo(ctx, "开始处理订单")

	// 模拟业务处理
	result := processOrder(ctx, 1001)

	// 构造响应
	response := fmt.Sprintf(
		"请求处理成功!\n用户: %s (%d)\n请求ID: %s\n追踪ID: %s\n订单结果: %s",
		user.Name, user.ID, reqID, traceID, result,
	)

	w.Write([]byte(response))

	// 记录业务处理完成
	logInfo(ctx, "订单处理完成")
}

// 业务处理函数（使用上下文中的值）
func processOrder(ctx context.Context, orderID int) string {
	// 记录处理开始
	logInfo(ctx, "处理订单中", map[string]interface{}{
		"order_id": orderID,
	})

	// 从上下文中获取请求ID
	reqID, ok := ctx.Value(requestIDKey).(string)
	if !ok {
		reqID = "unknown"
	}

	// 模拟处理时间
	time.Sleep(100 * time.Millisecond)

	// 记录处理完成
	logInfo(ctx, "订单处理完成")

	return fmt.Sprintf("订单 #%d 已处理 (请求ID: %s)", orderID, reqID)
}

// 增强版日志函数
func logInfo(ctx context.Context, message string, fields ...map[string]interface{}) {
	// 从上下文获取元数据
	reqID, _ := ctx.Value(requestIDKey).(string)
	traceID, _ := ctx.Value(traceIDKey).(string)
	user, _ := ctx.Value(userKey).(*User)

	// 构建基础日志
	logEntry := fmt.Sprintf("[%s][%s] %s", reqID, traceID, message)

	// 添加用户信息（如果存在）
	if user != nil {
		logEntry += fmt.Sprintf(" | 用户: %s(%d)", user.Name, user.ID)
	}

	// 添加额外字段（如果提供）
	if len(fields) > 0 {
		for key, value := range fields[0] {
			logEntry += fmt.Sprintf(" | %s: %v", key, value)
		}
	}

	fmt.Println(logEntry)
}

// 错误日志函数
func logError(ctx context.Context, message string) {
	reqID, _ := ctx.Value(requestIDKey).(string)
	traceID, _ := ctx.Value(traceIDKey).(string)

	fmt.Printf("[ERROR][%s][%s] %s\n", reqID, traceID, message)
}

func main() {
	// 设置HTTP路由
	mux := http.NewServeMux()
	mux.HandleFunc("/order", handleRequest)

	// 应用中间件
	handler := middleware(authMiddleware(mux))

	// 启动服务器
	fmt.Println("服务器启动在 :8080")
	http.ListenAndServe(":8080", handler)
}
