package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

// BankBusinessHandler
/*
参考资料：https://mp.weixin.qq.com/s/-Ysho1jI9MfrAIrplzj7UQ
*/
type BankBusinessHandler interface {
	// TakeRowNumber 排队拿号
	TakeRowNumber()
	// WaitInHead 等位
	WaitInHead()
	// HandleBusiness 处理具体业务
	HandleBusiness()
	// Commentate 对服务作出评价
	Commentate()
	// CheckVipIdentity 钩子方法，判断是不是VIP， VIP不用等位
	CheckVipIdentity() bool
}

type BankBusinessExecutor struct {
	handler BankBusinessHandler
}

// ExecuteBankBusiness 模板方法，处理银行业务
func (b *BankBusinessExecutor) ExecuteBankBusiness() {
	b.handler.TakeRowNumber()
	if !b.handler.CheckVipIdentity() {
		b.handler.WaitInHead()
	}
	b.handler.HandleBusiness()
	b.handler.Commentate()
}

// DepositBusinessHandler 存款业务的流程
type DepositBusinessHandler struct {
	*DefaultBusinessHandler
	userVip bool
}

func (*DepositBusinessHandler) HandleBusiness() {
	fmt.Println("账户存储很多万人民币...")
}

func (dh *DepositBusinessHandler) CheckVipIdentity() bool {
	return dh.userVip
}

func (*DepositBusinessHandler) Commentate() {

	fmt.Println("重写评价")
}

// DefaultBusinessHandler
/*
注意，上面的DefaultBusinessHandler并没有实现我们想要留给具体子类实现的HandleBusiness方法，
这样 DefaultBusinessHandler 就不能算是BankBusinessHandler接口的实现，这么做是为了这个类型只能用于被实现类包装，
让 Go 语言的类型检查能够帮我们强制要求，必须用存款或者取款这样子类去实现HandleBusiness方法，整个银行办理业务的流程的程序才能运行起来。
*/
type DefaultBusinessHandler struct {
}

func (*DefaultBusinessHandler) TakeRowNumber() {
	fmt.Println("请拿好您的取件码：" + strconv.Itoa(rand.Intn(100)) +
		" ，注意排队情况，过号后顺延三个安排")
}

func (dbh *DefaultBusinessHandler) WaitInHead() {
	fmt.Println("排队等号中...")
	time.Sleep(5 * time.Second)
	fmt.Println("请去窗口xxx...")
}

func (*DefaultBusinessHandler) Commentate() {

	fmt.Println("请对我的服务作出评价，满意请按0，满意请按0，(～￣▽￣)～")
}

func (*DefaultBusinessHandler) CheckVipIdentity() bool {
	// 留给具体实现类实现
	return false
}

func NewBankBusinessExecutor(businessHandler BankBusinessHandler) *BankBusinessExecutor {
	return &BankBusinessExecutor{handler: businessHandler}
}

func main() {
	dh := &DepositBusinessHandler{userVip: false}
	bbe := NewBankBusinessExecutor(dh)
	bbe.ExecuteBankBusiness()
}
