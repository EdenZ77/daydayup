package main

import "fmt"

/*
参考资料：https://mp.weixin.qq.com/s?__biz=MzUzNzAzMTc3MA==&mid=2247484407&idx=1&sn=6f4b17f3a396c5ff48ef39ef70d47438&chksm=faec66c2cd9befd4ba2653eb0c555838e0e9a49a11006cbfa9498a50ff16fdb522dfc50c324e&scene=178&cur_album_id=1908992469812199431#rd

桥接模式一般用于有多种分类的情况，如果实现系统可能有多角度分类，每一种分类都有可能变化，那么就把这种多角度分离出来让他们独立变化，减少他们之间的耦合。
将抽象部分与它的实现部分分离，使他们都可以独立地变化。


这个触达系统的业务场景是：已经定义好触达的紧急情况，触达需要的数据来源不同，当运营使用的时候，根据触达紧急情况，配置好数据（文案、收件人等）即可。
可以看出：一个分类是触达方式、一个分类是触达紧急情况。

一定学习极客时间关于API监控报警的例子，讲解了演变过程。
https://time.geekbang.org/column/article/202786?screen=full
*/

// MessageSend 消息发送接口
type MessageSend interface {
	send(msg string)
}

// SMS 短信消息
type SMS struct {
}

func (s *SMS) send(msg string) {
	fmt.Println("sms 发送的消息内容为: " + msg)
}

// Email 邮件消息
type Email struct {
}

func (e *Email) send(msg string) {
	fmt.Println("email 发送的消息内容为: " + msg)
}

// AppPush AppPush消息
type AppPush struct {
}

func (a *AppPush) send(msg string) {
	fmt.Println("appPush 发送的消息内容为: " + msg)
}

// Letter 站内信消息
type Letter struct {
}

func (l *Letter) send(msg string) {
	fmt.Println("站内信 发送的消息内容为: " + msg)
}

// Touch 用户触达父类，包含触达方式数组messageSends
type Touch struct {
	messageSends []MessageSend
}

/**
 * @Description: 触达方法，调用每一种方式进行触达
 * @receiver t
 * @param msg
 */
func (t *Touch) do(msg string) {
	for _, s := range t.messageSends {
		s.send(msg)
	}
}

// TouchUrgent 紧急消息做用户触达
type TouchUrgent struct {
	base Touch
}

/**
 * @Description: 紧急消息，先从db中获取各种信息，然后使用各种触达方式通知用户
 * @receiver t
 * @param msg
 */
func (t *TouchUrgent) do(msg string) {
	// 封装复杂的业务逻辑，试图将一个类的复杂代码拆分到更细小的类，然后再通过某种更合理的结构组装在一起
	fmt.Println("touch urgent 从db获取接收人等信息")
	t.base.do(msg)
}

// TouchNormal 普通消息做用户触达
type TouchNormal struct {
	base Touch
}

/**
 * @Description: 普通消息，先从文件中获取各种信息，然后使用各种触达方式通知用户
 * @receiver t
 * @param msg
 */
func (t *TouchNormal) do(msg string) {
	fmt.Println("touch normal 从文件获取接收人等信息")
	t.base.do(msg)
}

/*
	根据多个角度分类，每一种分类都有可能变化，例如本例中触达方式和紧急程度这两种分类都有可能增删，那么就把这个多角度分离出来，让他们独立变化，减少他们之间的耦合

使用桥接模式，一定要看一下场景中是否有多种分类、且分类之间有一定关联。如果符合的话，建议用桥接模式，这样不同分类可以独立变化，相互之间不影响。
*/
func main() {
	//触达方式
	sms := &SMS{}
	appPush := &AppPush{}
	letter := &Letter{}
	email := &Email{}

	//根据触达类型选择触达方式
	fmt.Println("-------------------touch urgent")
	touchUrgent := TouchUrgent{
		base: Touch{
			messageSends: []MessageSend{sms, appPush, letter, email},
		},
	}
	touchUrgent.do("urgent情况")
	fmt.Println("-------------------touch normal")
	touchNormal := TouchNormal{ //
		base: Touch{
			messageSends: []MessageSend{sms, appPush, letter, email},
		},
	}
	touchNormal.do("normal情况")
}
