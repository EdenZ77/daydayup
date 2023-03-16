package main

import "fmt"

// 两种发送消息的方法

type SendMessage interface {
	send(text, to string)
}

type sms struct {
}

func NewSms() SendMessage {
	return &sms{}
}

func (s *sms) send(text, to string) {
	fmt.Println(fmt.Sprintf("send %s to %s sms", text, to))
}

type email struct {
}

func NewEmail() SendMessage {
	return &email{}
}

func (e *email) send(text, to string) {
	fmt.Println(fmt.Sprintf("send %s to %s email", text, to))
}

// 两种发送系统

type systemA struct {
	method SendMessage
}

func NewSystemA(method SendMessage) *systemA {
	return &systemA{
		method: method,
	}
}

func (m *systemA) SendMessage(text, to string) {
	m.method.send(fmt.Sprintf("[System A] %s", text), to)
}

type systemB struct {
	method SendMessage
}

func NewSystemB(method SendMessage) *systemB {
	return &systemB{
		method: method,
	}
}

func (m *systemB) SendMessage(text, to string) {
	m.method.send(fmt.Sprintf("[System B] %s", text), to)
}

func ExampleSystemA() {
	NewSystemA(NewSms()).SendMessage("hi", "baby")
	NewSystemA(NewEmail()).SendMessage("hi", "baby")
	// Output:
	// send [System A] hi to baby sms
	// send [System A] hi to baby email
}

func ExampleSystemB() {
	NewSystemB(NewSms()).SendMessage("hi", "baby")
	NewSystemB(NewEmail()).SendMessage("hi", "baby")
	// Output:
	// send [System B] hi to baby sms
	// send [System B] hi to baby email
}
func main() {

}
