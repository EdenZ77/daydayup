package main

/*
参考资料：GPT

在实际应用中，通常会有一个事件总线（EventBus）或事件分发器（Event Dispatcher），
它们管理事件和订阅者之间的关系，并负责将事件从发布者分发给所有的订阅者。以下是一个简单的实现示例
*/
import (
	"fmt"
	"sync"
)

// Event 定义了一个基础事件
type Event struct {
	Data string
}

// EventBus 事件总线
type EventBus struct {
	subscribers map[string][]Observer
	lock        sync.RWMutex
}

// NewEventBus 创建一个新的事件总线
func NewEventBus() *EventBus {
	return &EventBus{
		subscribers: make(map[string][]Observer),
	}
}

// Subscribe 订阅事件
func (bus *EventBus) Subscribe(eventType string, observer Observer) {
	bus.lock.Lock()
	defer bus.lock.Unlock()

	bus.subscribers[eventType] = append(bus.subscribers[eventType], observer)
}

// Publish 发布事件
func (bus *EventBus) Publish(eventType string, event Event) {
	bus.lock.RLock()
	defer bus.lock.RUnlock()

	if observers, found := bus.subscribers[eventType]; found {
		for _, observer := range observers {
			observer.Update(event)
		}
	}
}

// Unsubscribe 取消订阅事件
func (bus *EventBus) Unsubscribe(eventType string, observer Observer) {
	bus.lock.Lock()
	defer bus.lock.Unlock()

	if observers, found := bus.subscribers[eventType]; found {
		for i, o := range observers {
			// 接口变量比较，对于*T和T是不同的。
			if o == observer {
				// 移除订阅者
				bus.subscribers[eventType] = append(observers[:i], observers[i+1:]...)
				break
			}
		}
	}
}

// Observer 观察者接口
type Observer interface {
	Update(event Event)
}

// ObserverImpl 观察者实现
type ObserverImpl struct {
	ID int
}

// Update 实现观察者接口
func (o ObserverImpl) Update(event Event) {
	fmt.Printf("Observer %d received: %s\n", o.ID, event.Data)
}

func main() {
	bus := NewEventBus()

	observer1 := ObserverImpl{ID: 1}
	observer11 := ObserverImpl{ID: 1}
	observer2 := ObserverImpl{ID: 2}
	observer3 := ObserverImpl{ID: 3}
	observer4 := ObserverImpl{ID: 4}
	observer44 := ObserverImpl{ID: 4}

	// 订阅者订阅不同类型的事件
	bus.Subscribe("topic:hello", observer1)
	bus.Subscribe("topic:hello", observer2)
	bus.Subscribe("topic:goodbye", observer3)
	bus.Subscribe("topic:goodbye", observer4)

	// 分别发布不同类型的事件
	bus.Publish("topic:hello", Event{Data: "Hello Event"})
	bus.Publish("topic:goodbye", Event{Data: "Goodbye Event"})

	// 取消订阅
	bus.Unsubscribe("topic:hello", observer11)
	bus.Unsubscribe("topic:goodbye", observer44)

	// 再次发布事件，验证取消订阅后，observer1 和 observer4 不再收到通知
	bus.Publish("topic:hello", Event{Data: "Hello Event Again"})
	bus.Publish("topic:goodbye", Event{Data: "Goodbye Event Again"})
}
