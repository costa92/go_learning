package singleton

import (
	"fmt"
	"sync"
	"testing"
	"unsafe"
)

type Singleton struct {

}

var singleInstance *Singleton  // 单例模式
var once sync.Once // 只创建一次

func GetSingletonObj() *Singleton{
	once.Do(func() {
		fmt.Println("Create Obj")
		singleInstance = new(Singleton)
	})
	return singleInstance;
}

func TestGetSingletonObj(t *testing.T)  {
	var wg sync.WaitGroup
	for i:=0; i< 10 ; i++  {
		wg.Add(1)
		go func() {  // 创建协程
			obj:=GetSingletonObj()
			//fmt.Printf("%x",obj)
			fmt.Printf("%d\n",unsafe.Pointer(obj))
			wg.Done()
		}()
	}
	wg.Wait()
}
