package main

import (
	"fmt"
	"os"
)

func readFile(filename string) (string, error) {
	file, err := os.Open(filename)

	if err != nil {
		return "", err
	}

	defer file.Close()

	buffer := make([]byte, 1024)

	n, err := file.Read(buffer)

	if err != nil {
		return "", err
	}

	return string(buffer[:n]), nil
}

func main() {

	content, err := readFile("test.txt")

	if err != nil {
		fmt.Println("读取失败:", err)
		return
	}

	fmt.Println(content)
}
