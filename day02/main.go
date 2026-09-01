package main

import "fmt"

func main() {
	var name string
	var age int8
	fmt.Print("请输入姓名：")
	fmt.Scan(&name)
	fmt.Print("请输入年龄：")
	fmt.Scan(&age)
	fmt.Println("===== Student =====")
	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
}
