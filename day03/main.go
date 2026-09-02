package main

import "fmt"

func Operate(a float64, b float64, operator string) float64 {

	if operator == "+" {
		return a + b
	} else if operator == "-" {
		return a - b
	} else if operator == "*" {
		return a * b
	} else if operator == "/" {
		if b != 0 {
			return a / b
		} else {
			return -1
		}
	} else {
		return -1
	}

}

func main() {
	var a float64
	var b float64
	var operator string

	fmt.Println("请输入第一个数字")
	fmt.Scan(&a)
	fmt.Println("请输入运算符")
	fmt.Scan(&operator)
	fmt.Println("请输入第二个数字")
	fmt.Scan(&b)
	result := Operate(a, b, operator)

	fmt.Print("结果为：", result)
}
