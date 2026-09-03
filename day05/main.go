package main

import "fmt"

func addUser(users []string) []string {
	var username string
	fmt.Println("请输入新用户的名字：")
	fmt.Scan(&username)
	users = append(users, username)
	return users
}

func listUsers(users []string) {
	fmt.Println("用户列表:")

	for index, value := range users {

		fmt.Println(index, value)

	}
}

func searchUsers(users []string) {
	var searchname string
	found := false
	fmt.Println("请输入搜索用户名:")
	fmt.Scan(&searchname)
	for index, value := range users {
		if searchname == value {
			found = true
			fmt.Println("找到该用户了", index, value)
		}
	}
	if found == false {
		fmt.Println("该用户不存在")
	}
}

func deleteUser(users []string) []string {
	var deleteName string
	fmt.Println("请输入你想删除用户的名字：")
	fmt.Scan(&deleteName)
	deleteIndex := -1
	for index, value := range users {
		if value == deleteName {
			deleteIndex = index
			break
		}
	}
	if deleteIndex != -1 {
		users = append(users[:deleteIndex], users[deleteIndex+1:]...)
		fmt.Println("删除成功")
	} else {
		fmt.Println("该用户不存在")
	}
	return users
}
func menu(users []string) {
	var choice int
MenuLoop:
	for {

		fmt.Println("====== User System ======")
		fmt.Println("1. 添加用户")
		fmt.Println("2. 查看用户")
		fmt.Println("3. 搜索用户")
		fmt.Println("4. 删除用户")
		fmt.Println("5. 退出")
		fmt.Println("请选择:")
		fmt.Scan(&choice)
		switch choice {
		case 1:
			users = addUser(users)
		case 2:
			listUsers(users)
		case 3:
			searchUsers(users)
		case 4:
			users = deleteUser(users)
		case 5:
			break MenuLoop

		default:
			fmt.Println("功能开发中")
		}

	}

}

func main() {
	users := []string{}
	users = append(users,
		"Tom",
		"Jack",
		"Alice",
	)
	menu(users)
}
