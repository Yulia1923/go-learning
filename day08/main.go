package main

import "fmt"

type User struct {
	ID int

	Name string

	Age int
}

var users = make(map[int]User)

func AddUser() {
	var id int
	var name string
	var age int

	fmt.Println("请输入id:")
	fmt.Scan(&id)

	fmt.Println("请输入名字:")
	fmt.Scan(&name)

	fmt.Println("请输入年龄:")
	fmt.Scan(&age)

	user := User{
		ID:   id,
		Name: name,
		Age:  age,
	}
	users[id] = user

	fmt.Println("添加成功！")
}

func GetUser() {
	var id int

	fmt.Println("请输入你想查询的用户:")

	fmt.Scan(&id)

	user, ok := users[id]

	if ok {

		fmt.Println("用户信息:")
		fmt.Println("ID:", user.ID)
		fmt.Println("Name:", user.Name)
		fmt.Println("Age:", user.Age)

	} else {

		fmt.Println("用户不存在")
	}

}

func UpdateUser() {
	var id int

	fmt.Println("请输入修改用户id:")

	fmt.Scan(&id)

	user, ok := users[id]

	if ok {
		var age int

		fmt.Println("请输入新的年龄:")

		fmt.Scan(&age)

		user.Age = age

		users[id] = user

		fmt.Println("修改成功")
	} else {
		fmt.Println("用户不存在")
	}
}

func DeleteUser() {

	var id int

	fmt.Println("请输入删除用户ID:")

	fmt.Scan(&id)

	_, ok := users[id]

	if ok {
		delete(users, id)
		fmt.Println("删除成功")
	} else {
		fmt.Println("用户不存在")
	}
}

func main() {

}
