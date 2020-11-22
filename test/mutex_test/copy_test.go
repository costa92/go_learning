package mutex_test

import (
	"fmt"
	"sync"
	"testing"
)

// 复制锁

type Counter struct {
	sync.Mutex
	Count int
}

func TestCopy(t *testing.T) {
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