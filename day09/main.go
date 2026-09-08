package main

import "fmt"

type User struct {
	Name string

	Age int
}

func updateUser(user *User) {
	fmt.Println("函数收到的地址:")
	fmt.Println(user)
	user.Name = "Bob"
	user.Age = 22
}

func addAge(user *User) {
	user.Age = 21
}
func main() {
	user := User{
		Name: "Tom",

		Age: 20,
	}
	fmt.Println("用户地址:")
	fmt.Println(&user)
	fmt.Println("修改前:")
	fmt.Println(user)

	updateUser(&user)

	fmt.Println("修改后:")
	fmt.Println(user)

	addAge(&user)
}
