package main

import "fmt"

type Payment interface {
	Pay(float64)
}

type Alipay struct{}
type WeChatPay struct{}
type BankCard struct{}

func (a Alipay) Pay(amount float64) {
	fmt.Println("支付宝支付", amount)
}

func (w WeChatPay) Pay(amount float64) {
	fmt.Println("微信支付", amount)
}

func (b BankCard) Pay(amount float64) {
	fmt.Println("银行卡支付", amount)
}

func (a Alipay) Refund() {
	fmt.Println("支付宝退款")
}

func Checkout(payment Payment, amount float64) {
	payment.Pay(amount)

	if payment, ok := payment.(Alipay); ok {
		payment.Refund()
	}
}

func PrintData(data any) {
	switch v := data.(type) {
	case int:
		fmt.Println("数字:", v)

	case string:
		fmt.Println("字符串:", v)

	case Alipay:
		fmt.Println("支付宝:", v)
	}
}

func main() {
	Checkout(Alipay{}, 500)
	Checkout(WeChatPay{}, 200)
	Checkout(BankCard{}, 1000)
}
