package main

import (
	"fmt"
	"sync"
)

// 复制锁

type Counter struct {
	sync.Mutex
	Count int
}

func main() {
	var c Counter
	c.Lock()
	defer c.Unlock()
	c.Count ++
	foo(c) // 复制锁
}

func foo(c Counter) {
	c.Lock()
	defer c.Unlock()
	fmt.Println("in foo")
}

//  go vet copy.go  vet 检查复用锁的工具
