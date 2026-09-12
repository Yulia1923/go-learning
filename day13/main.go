package main

import "fmt"

type Payment interface {
	Pay(amount float64) //只要一个类型具有 Pay(float64) 方法，它就可以作为 Payment 使用
}

type Alipay struct{}
type WechatPay struct{}

func (Alipay) Pay(amount float64) {
	fmt.Println("支付宝支付:", amount)
}

func (WechatPay) Pay(amount float64) {
	fmt.Println("微信支付：", amount)
}

func Checkout(payment Payment, amount int) {
	payment.Pay(float64(amount))
}

func main() {
	Checkout(Alipay{}, 500)
	Checkout(WechatPay{}, 200)
}
