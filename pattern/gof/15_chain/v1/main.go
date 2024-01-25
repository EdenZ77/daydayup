package main

import "fmt"

/*
参考资料：https://mp.weixin.qq.com/s/zCh12E10JM24EGTyFS7hPQ

它是一种行为型设计模式。使用这个模式，我们能为请求创建一条由多个处理器组成的链路，每个处理器各自负责自己的职责，
相互之间没有耦合，完成自己任务后请求对象即传递到链路的下一个处理器进行处理。

我们可以确定：假如一个流程的步骤不固定，为了在流程中增加步骤时，不必修改原有已经开发好，经过测试的流程，
我们需要让整个流程中的各个步骤解耦，来增加流程的扩展性，这种时候就可以使用职责链模式啦，这个模式可以让我们先设置流程链路中有哪些步骤，再去执行。

看病的具体流程如下：挂号—>诊室看病—>收费处缴费—>药房拿药
我们的目标是利用责任链模式，实现这个流程中的每个步骤，且相互间不耦合，还支持向流程中增加步骤。

*/

type PatientHandler interface {
	Execute(*patient) error
	SetNext(PatientHandler) PatientHandler
	Do(*patient) error
}

type Next struct {
	nextHandler PatientHandler
}

func (n *Next) SetNext(handler PatientHandler) PatientHandler {
	n.nextHandler = handler
	return handler
}

func (n *Next) Execute(patient *patient) (err error) {
	if n.nextHandler != nil {
		if err = n.nextHandler.Do(patient); err != nil {
			return
		}

		return n.nextHandler.Execute(patient)
	}

	return
}

func (n *Next) Do(patient *patient) (err error) {
	fmt.Println("Next Do===")
	return nil
}

// Pharmacy 药房处理器
type Pharmacy struct {
	Next
}

func (m *Pharmacy) Do(p *patient) (err error) {
	if p.MedicineDone {
		fmt.Println("Medicine already given to patient")
		return
	}
	fmt.Println("Pharmacy giving medicine to patient")
	p.MedicineDone = true
	return
}

// Cashier 收费处处理器
type Cashier struct {
	Next
}

func (c *Cashier) Do(p *patient) (err error) {
	if p.PaymentDone {
		fmt.Println("Payment Done")
		return
	}
	fmt.Println("Cashier getting money from patient patient")
	p.PaymentDone = true
	return
}

// Clinic 诊室处理器--用于医生给病人看病
type Clinic struct {
	Next
}

func (d *Clinic) Do(p *patient) (err error) {
	if p.DoctorCheckUpDone {
		fmt.Println("Doctor checkup already done")
		return
	}
	fmt.Println("Doctor checking patient")
	p.DoctorCheckUpDone = true
	return
}

// Reception 挂号处处理器
type Reception struct {
	Next
}

func (r *Reception) Do(p *patient) (err error) {
	if p.RegistrationDone {
		fmt.Println("Patient registration already done")
		return
	}
	fmt.Println("Reception registering patient")
	p.RegistrationDone = true
	return
}

// StartHandler 不做操作，作为第一个Handler向下转发请求
type StartHandler struct {
	Next
}

// Do 空Handler的Do
//func (h *StartHandler) Do(c *patient) (err error) {
//	// 空Handler 这里什么也不做 只是载体 do nothing...
//	fmt.Println("StartHandler Do===")
//	return
//}

type patient struct {
	Name              string
	RegistrationDone  bool
	DoctorCheckUpDone bool
	MedicineDone      bool
	PaymentDone       bool
}

func main() {
	patientHealthHandler := StartHandler{}
	//
	patient := &patient{Name: "abc"}
	// 设置病人看病的链路
	patientHealthHandler.SetNext(&Reception{}). // 挂号
							SetNext(&Clinic{}).  // 诊室看病
							SetNext(&Cashier{}). // 收费处交钱
							SetNext(&Pharmacy{}) // 药房拿药

	// 执行上面设置好的业务流程
	if err := patientHealthHandler.Execute(patient); err != nil {
		// 异常
		fmt.Println("Fail | Error:" + err.Error())
		return
	}
	// 成功
	fmt.Println("Success")
}
