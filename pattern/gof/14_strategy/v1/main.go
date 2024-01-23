package main

import "fmt"

type PayBehavior interface {
	OrderPay(payParams map[string]interface{})
}

// WxPay 具体支付策略实现
type WxPay struct{}

func (*WxPay) OrderPay(payParams map[string]interface{}) {
	fmt.Printf("Wx支付加工支付请求 %v\n", payParams)
	fmt.Println("正在使用Wx支付进行支付")
}

// ThirdPay 三方支付
type ThirdPay struct{}

func (*ThirdPay) OrderPay(payParams map[string]interface{}) {
	fmt.Printf("三方支付加工支付请求 %v\n", payParams)
	fmt.Println("正在使用三方支付进行支付")
}

type PayCtx struct {
	// 提供支付能力的接口实现
	payBehavior PayBehavior
}

func (px *PayCtx) setPayBehavior(p PayBehavior) {
	px.payBehavior = p
}

func (px *PayCtx) Pay(payParams map[string]interface{}) {
	px.payBehavior.OrderPay(payParams)
}

func NewPayCtx(p PayBehavior) *PayCtx {
	return &PayCtx{
		payBehavior: p,
	}
}

func main() {
	wxPayParams := map[string]interface{}{
		"appId": "234fdfdngj4",
		"mchId": 123456,
		// 微信支付特有参数
		"openId": "user-open-id",
	}
	thPayParams := map[string]interface{}{
		"vendor": "third_party",
		"token":  "secure-token",
		// 三方支付特有参数
		"callbackUrl": "https://example.com/callback",
	}

	wxPay := &WxPay{}
	px := NewPayCtx(wxPay)
	px.Pay(wxPayParams)

	// 假设现在发现微信支付没钱，改用三方支付进行支付
	thPay := &ThirdPay{}
	px.setPayBehavior(thPay)
	px.Pay(thPayParams)
}
