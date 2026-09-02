package main

import "fmt"

func Login(username string, password string, correctPassword string) bool {

	correctUsername := "admin"

	if username == correctUsername &&
		password == correctPassword {

		return true
	}

	return false
}

func UserCenter(username string, password string, loginCount int) string {
MenuLoop:
	for {
		var choice int
		fmt.Println("===== User Center =====")
		fmt.Println("1. 查看用户信息")
		fmt.Println("2. 修改密码")
		fmt.Println("3. 查看登录次数")
		fmt.Println("4. 退出")
		fmt.Println("请选择:")
		fmt.Scan(&choice)
		switch choice {
		case 1:
			fmt.Println("用户名:", username)
			fmt.Println("状态：在线")
		case 2:
			var oldPassword string
			var newPassword string

			fmt.Println("请输入旧密码：")
			fmt.Scan(&oldPassword)

			if oldPassword == password {
				fmt.Println("请输入新密码:")
				fmt.Scan(&newPassword)

				password = newPassword
				fmt.Println("修改成功")
			} else {
				fmt.Println("旧密码错误")
			}
		case 3:
			fmt.Println("登录次数:", loginCount)
		case 4:
			fmt.Println("退出系统")
			break MenuLoop
		default:
			fmt.Println("请输入正确选项")

		}
	}
	return password
}

func main() {

	var username string
	var password string
	failCount := 0
	loginCount := 0
	correctPassword := "123456"

	for failCount < 5 {
		fmt.Println("请输入用户名：")
		fmt.Scan(&username)

		fmt.Println("请输入密码：")
		fmt.Scan(&password)

		if Login(username, password, correctPassword) {

			fmt.Println("登录成功！")
			loginCount++
			correctPassword = UserCenter(
				username,
				correctPassword,
				loginCount,
			)
			break

		} else {

			fmt.Println("登录失败！")
			failCount++
		}
	}
	if failCount >= 5 {
		fmt.Println("账号已锁定")
	}

}
