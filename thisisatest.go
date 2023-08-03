package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("请输入数据，以两个回车符为结束标志：")

	scanner := bufio.NewScanner(os.Stdin)
	lines := []string{}

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			// 遇到两个回车符结束输入
			break
		}

		lines = append(lines, line)
	}

	fmt.Println("您输入的数据为：")
	fmt.Printf("长度为: %d\n", len(lines))
	for _, line := range lines {
		fmt.Println(line)
	}
}
