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

	// 输入第一个数字
	fmt.Println("请输入第一个数字")
	fmt.Scan(&a)
	// 输入
	fmt.Println("请输入运算符")
	fmt.Scan(&operator)
	fmt.Println("请输入第二个数字")
	// 根据 operator 计算结果
	fmt.Scan(&b)
	// 输出结果
	result := Operate(a, b, operator)

	fmt.Print("结果为：", result)
}
