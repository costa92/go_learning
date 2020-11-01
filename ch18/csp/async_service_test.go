package csp

import (
	"fmt"
	"testing"
	"time"
)

func service() string {
	time.Sleep(time.Millisecond*50)
	return  "Done"
}

func otherTask()  {
	fmt.Println("working on something else")
	time.Sleep(time.Millisecond * 1000)
	fmt.Println("Task is Done.")
}

func TestService(t *testing.T){
	fmt.Println(service())
	otherTask()
}

func AsyncService() chan string {
	//retCh := make(chan string)
	// 设置容量的大小，如果不设置空间那边会一直等待，后面协程会堵塞
	retCh := make(chan string,1)

	go func() {
		ret := service()
		fmt.Println("returned result.")
		retCh <- ret
		fmt.Println("service exited.")
	}()
	return retCh
}

func TestAsynService(t *testing.T)  {
	retCh := AsyncService()
	otherTask()
	fmt.Println(<-retCh)
	time.Sleep(time.Second *1)
}