package main

import (
	"fmt"
	"os"
)

func readFile() (string, error) {

	file, err := os.Open("test.txt")

	if err != nil {

		return "", err
	}

	defer file.Close()

	buffer := make([]byte, 1024)
	n, err := file.Read(buffer)

	return string(buffer[:n]), nil
}

func main() {

	content, err := readFile()

	if err != nil {

		fmt.Println("读取失败:", err)

		return
	}

	fmt.Println(content)

}
