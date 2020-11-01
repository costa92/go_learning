package util_all_done

// 获取最先计算完的结果
import (
	"fmt"
	"runtime"
	"testing"
	"time"
)


func  runTask(id int) string {
	time.Sleep(10*time.Millisecond)
	return  fmt.Sprintf("The result is from %d",id)
}

func AllResponse() string{
	numOfRunner := 10
	//ch := make(chan string)
	ch := make(chan string,numOfRunner) // 防止协程泄露，利用channel 机制
	for i:=0;i < numOfRunner;i++ {
		go func(i int) {
			ret := runTask(i)
			ch <-ret
		}(i)
	}
	fianlRet := ""
	for j:=0;j< numOfRunner ; j++  {
		fianlRet += <-ch + "\n"
	}
	return fianlRet
}

func TestFirstResponse(t *testing.T)  {
	t.Log("Before:",runtime.NumGoroutine())  // NumGoroutine 获取协程数量
	t.Log(AllResponse())
	time.Sleep(time.Second)
	t.Log("After:",runtime.NumGoroutine())  // NumGoroutine 获取协程数量
}