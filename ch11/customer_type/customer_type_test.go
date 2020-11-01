package customer_type

import (
	"fmt"
	"testing"
	"time"
)

// 定义别名
type IntCove func(op int) int

// 函数包一下 类似面向对象的 装饰者模式
func timeSpent(inner IntCove) func(op int) int{
	return func(n int) int{
		start := time.Now()
		ret := inner(n)
		//打印输出时间
		fmt.Println("time spent:",time.Since(start).Seconds())
		return ret
	}
}

func slowFun(op int) int {
	time.Sleep(time.Second*1)
	return op
}

func TestFn(t *testing.T){

	tsSf := timeSpent(slowFun) //slowFun 函数
	t.Log(tsSf(10))
}