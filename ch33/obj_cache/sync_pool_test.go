package obj_cache

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
)

func TestSyncPool(t *testing.T){
	pool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("Create a new Object")
			return 100
		},
	}

	v := pool.Get().(int)
	fmt.Println(v)
	pool.Put(3)
	runtime.GC() // GC 会清除 sync.pool 中缓存的对象
	v1 , _ := pool.Get().(int)
	fmt.Println(v1)
}

func TestSyncPoolMultiGroutine(t *testing.T)  {
	pool := &sync.Pool{
		New: func() interface{} {
			t.Log("Create a new object")
			return 10
		},
	}

	pool.Put(100)
	pool.Put(100)
	pool.Put(100)

	var wg sync.WaitGroup
	for i:=0; i <10  ; i++  {
		wg.Add(1)
		go func(i int) {
          t.Log(pool.Get())
          wg.Done()
		}(i)
	}
	wg.Wait()
}