package main

import (
	"fmt"
	"os"
)

//go 中 main 函数不支持任何返回值
func main() {
	if len(os.Args) > 1 {
		fmt.Println("hello world !!",os.Args[1])
	}
	os.Exit(0)
}