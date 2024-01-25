package main

import "fmt"

/*
参考资料：GPT
*/

// Handler 定义了处理器的接口
type Handler interface {
	HandleRequest(request string)
	SetNext(handler Handler)
}

// BaseHandler 提供了一个基础的处理器实现，它包含对下一个处理器的引用
type BaseHandler struct {
	next Handler
}

// SetNext 设置链中的下一个处理器
func (h *BaseHandler) SetNext(handler Handler) {
	h.next = handler
}

// HandleRequest 如果当前处理器不处理请求，则将请求传递给下一个处理器
func (h *BaseHandler) HandleRequest(request string) {
	if h.next != nil {
		h.next.HandleRequest(request)
	}
}

// ConcreteHandlerA 是一个具体的处理器，它只处理请求中包含"A"的情况
type ConcreteHandlerA struct {
	BaseHandler
}

func (h *ConcreteHandlerA) HandleRequest(request string) {
	if canHandle := h.canHandleRequest(request); canHandle {
		fmt.Println("Handled by ConcreteHandlerA")
	} else {
		fmt.Println("Cannot handle by ConcreteHandlerA, passing to next")
		h.BaseHandler.HandleRequest(request)
	}
}

func (h *ConcreteHandlerA) canHandleRequest(request string) bool {
	return request == "A"
}

// ConcreteHandlerB 是另一个具体的处理器，它只处理请求中包含"B"的情况
type ConcreteHandlerB struct {
	BaseHandler
}

func (h *ConcreteHandlerB) HandleRequest(request string) {
	if canHandle := h.canHandleRequest(request); canHandle {
		fmt.Println("Handled by ConcreteHandlerB")
	} else {
		fmt.Println("Cannot handle by ConcreteHandlerB, passing to next")
		h.BaseHandler.HandleRequest(request)
	}
}

func (h *ConcreteHandlerB) canHandleRequest(request string) bool {
	return request == "B"
}

func main() {
	// 初始化处理器，并设置责任链
	handlerA := &ConcreteHandlerA{}
	handlerB := &ConcreteHandlerB{}

	handlerA.SetNext(handlerB)

	/*
		输出：Handled by ConcreteHandlerA
	*/
	//handlerA.HandleRequest("A")

	/*
		输出：
		Cannot handle by ConcreteHandlerA, passing to next
		Handled by ConcreteHandlerB
	*/
	//handlerA.HandleRequest("B")

	/*
		输出：
			Cannot handle by ConcreteHandlerA, passing to next
			Cannot handle by ConcreteHandlerB, passing to next
	*/
	handlerA.HandleRequest("C")
}
