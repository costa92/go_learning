package mutex_test

import (
	"fmt"
	"sync"
	"testing"
)

type  Foo struct {
	mu sync.Mutex
	count int
}

func (f *Foo) Bar() {
	f.mu.Lock()
	if f.count < 1000 {
		f.count += 3
		f.mu.Unlock()  // 此处释放锁

		return
	}

	f.count++
	f.mu.Unlock()
	fmt.Println(f.count)
	return

}

func (f *Foo) Bars () {
	f.mu.Lock()
	defer f.mu.Unlock()

	if f.count < 1000 {
		f.count +=3
		return
	}
	f.count++
	return
}

func TestLock(t *testing.T){
	var foo = Foo{}
	foo.Bar()

	fmt.Println(foo.count)
}