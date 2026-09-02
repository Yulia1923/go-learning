package main

import "fmt"

func checkScore(score float32) bool {
	return score >= 60

}

func getGrade(score float32) string {
	if score >= 90 {
		return "A"
	} else if score >= 80 {
		return "B"
	} else if score >= 70 {
		return "C"
	} else if score >= 60 {
		return "D"
	} else {
		return "F"
	}
}

func main() {
	var name string
	var age int
	var score float32
	fmt.Print("请输入姓名：")
	fmt.Scan(&name)
	fmt.Print("请输入年龄：")
	fmt.Scan(&age)
	fmt.Print("请输入分数：")
	fmt.Scan(&score)
	passed := checkScore(score)
	grade := getGrade(score)
	fmt.Println("===== Student =====")
	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Score", score)
	fmt.Println("Passed", passed)
	fmt.Println("Grade", grade)
}
