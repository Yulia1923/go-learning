package main

import "fmt"

func showProducts(products []string, stocks []int) {
	for i, value := range products {
		fmt.Println(value, stocks[i])
	}
}

func addProduct(products []string, stocks []int, name string, stock int) ([]string, []int) {
	products = append(products, name)
	stocks = append(stocks, stock)
	return products, stocks
}

func updateStock(products []string, stocks []int, name string, newStock int) {
	label := false
	for index, value := range products {
		if name == value {
			stocks[index] = newStock
			fmt.Println("库存更新成功！")
			label = true
			break
		}
	}
	if label == false {
		fmt.Println("该产品不存在")
	}
}

func deleteProduct(products []string, stocks []int, name string) ([]string, []int) {
	label := false
	for index, value := range products {
		if name == value {
			products = append(products[:index], products[index+1:]...)
			stocks = append(stocks[:index], stocks[index+1:]...)
			label = true
			break
		}
	}
	if label == false {
		fmt.Println("该产品不存在")
	}
	return products, stocks
}

func main() {
	products := []string{"Apple", "Banana", "Orange"}
	stocks := []int{100, 50, 80}
	showProducts(products, stocks)
	products, stocks = addProduct(products, stocks, "Milk", 30)
	updateStock(products, stocks, "Banana", 200)
	products, stocks = deleteProduct(products, stocks, "Banana")
	showProducts(products, stocks)
}
