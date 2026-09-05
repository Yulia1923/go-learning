package main

import (
	"fmt"
	"os"
)

type User struct {
	ID       int
	Username string
	Role     string
}

type Product struct {
	ID    int
	Name  string
	Price float64
	Stock int
}

func addUser(users map[int]User, user User) {
	users[user.ID] = user
	fmt.Println("用户添加成功")
}

func addProduct(products map[int]Product, product Product) {
	products[product.ID] = product
	fmt.Println("商品添加成功")
}

func getUser(users map[int]User, id int) (User, bool) {
	user, ok := users[id]
	if ok {
		fmt.Println("用户存在：", user)
		return user, true
	} else {
		fmt.Println("该用户不存在")
		return User{}, false
	}
}

func getProduct(products map[int]Product, id int) (Product, bool) {
	product, ok := products[id]
	if ok {
		fmt.Println("该商品存在：", product)
		return product, true
	} else {
		fmt.Println("该商品不存在")
		return Product{}, false
	}
}
func updateUser(users map[int]User, id int, newUser User) bool {
	_, ok := users[id]
	if !ok {
		fmt.Println("用户不存在")
		return false
	}
	users[id] = newUser
	fmt.Println("用户修改成功")
	return true
}
func updateStock(products map[int]Product, id int, stock int) bool {
	product, ok := products[id]
	if !ok {
		fmt.Println("该商品不存在")
		return false
	}
	product.Stock = stock
	products[id] = product
	fmt.Println("修改库存成功")
	return true
}
func deleteUser(users map[int]User, id int) bool {
	_, ok := users[id]

	if !ok {
		fmt.Println("用户不存在")
		return false
	}

	delete(users, id)

	fmt.Println("用户删除成功")
	return true
}
func deleteProduct(products map[int]Product, id int) bool {
	_, ok := products[id]

	if !ok {
		fmt.Println("该商品不存在")
		return false
	}
	delete(products, id)

	fmt.Println("商品删除成功")
	return true
}

func showMenu(
	users map[int]User,
	products map[int]Product,
	currentUser User,
) {
	var choice int
	for {
		fmt.Println("====================")
		fmt.Println("商品库存管理系统")
		fmt.Println("====================")
		fmt.Println("当前用户:", currentUser.Username)
		fmt.Println("权限:", currentUser.Role)
		fmt.Println("1. 添加用户")
		fmt.Println("2. 查询用户")
		fmt.Println("3. 修改用户")
		fmt.Println("4. 删除用户")
		fmt.Println("5. 添加商品")
		fmt.Println("6. 查询商品")
		fmt.Println("7. 修改库存")
		fmt.Println("8. 删除商品")
		fmt.Println("0. 退出")
		fmt.Println("请选择你的操作：")
		fmt.Scan(&choice)

		switch choice {

		case 1:
			if currentUser.Role != "admin" {

				fmt.Println("权限不足")

				break
			}
			fmt.Println("添加用户")
			var username string
			var role string
			fmt.Println("请输入用户名：")
			fmt.Scan(&username)
			fmt.Println("请输入角色：")
			fmt.Scan(&role)
			user := User{
				ID:       len(users) + 1,
				Username: username,
				Role:     role,
			}
			addUser(users, user)

		case 2:
			fmt.Println("查询用户")
			var id int
			fmt.Println("请输入要查询的用户id:")
			fmt.Scan(&id)
			getUser(users, id)

		case 3:
			if currentUser.Role != "admin" {

				fmt.Println("权限不足")

				break
			}
			fmt.Println("修改用户")
			var id int
			var username string
			var role string
			fmt.Println("请输入要修改的用户id:")
			fmt.Scan(&id)
			fmt.Println("请输入要修改的用户名字:")
			fmt.Scan(&username)
			fmt.Println("请输入要修改的用户角色:")
			fmt.Scan(&role)

			newUser := User{
				ID:       id,
				Username: username,
				Role:     role,
			}
			updateUser(users, id, newUser)
		case 4:
			if currentUser.Role != "admin" {

				fmt.Println("权限不足")

				break
			}
			fmt.Println("删除用户")
			var id int
			fmt.Println("请输入要删除的用户id:")
			fmt.Scan(&id)
			deleteUser(users, id)

		case 5:
			if currentUser.Role != "admin" {

				fmt.Println("权限不足")

				break
			}
			fmt.Println("添加商品")
			var productname string
			var price float64
			var stock int
			fmt.Println("请输入商品名称:")
			fmt.Scan(&productname)
			fmt.Println("请输入商品价格:")
			fmt.Scan(&price)
			fmt.Println("请输入商品库存:")
			fmt.Scan(&stock)
			product := Product{
				ID:    len(products) + 1,
				Name:  productname,
				Price: price,
				Stock: stock,
			}
			addProduct(products, product)

		case 6:
			fmt.Println("查询商品")
			var id int
			fmt.Println("请输入要查询的商品id:")
			fmt.Scan(&id)
			getProduct(products, id)

		case 7:
			if currentUser.Role != "admin" {

				fmt.Println("权限不足")

				break
			}
			fmt.Println("修改库存")
			var id int
			var newstock int
			fmt.Println("请输入要修改的商品id:")
			fmt.Scan(&id)
			fmt.Println("请输入新库存:")
			fmt.Scan(&newstock)
			updateStock(products, id, newstock)
		case 8:
			if currentUser.Role != "admin" {

				fmt.Println("权限不足")

				break
			}
			fmt.Println("删除商品")
			var id int
			fmt.Println("请输入要删除的商品id:")
			fmt.Scan(&id)
			deleteProduct(products, id)
		case 0:
			fmt.Println("系统退出")
			os.Exit(0)

		default:
			fmt.Println("请输入正确选项")
		}
	}
}

func login(users map[int]User) (User, bool) {
	var id int

	fmt.Println("请输入登录用户ID:")
	fmt.Scan(&id)

	user, ok := getUser(users, id)

	if ok {
		fmt.Println("登录成功")
		return user, true
	}
	fmt.Println("登录失败")
	return User{}, false
}

func main() {
	users := make(map[int]User)
	products := make(map[int]Product)

	addUser(users, User{
		ID:       1,
		Username: "admin",
		Role:     "admin",
	})

	addUser(users, User{
		ID:       2,
		Username: "Tom",
		Role:     "user",
	})

	currentUser := users[1]
	for {
		user, ok := login(users)
		if ok {
			showMenu(users, products, user)
		} else {
			fmt.Println("请重新登录")
		}
	}
}
