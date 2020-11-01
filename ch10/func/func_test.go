package fn_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func returnMultiValues() (int ,int) { // 返回两个值
	return rand.Intn(10),rand.Intn(20)
}

func slowFun(op int) int {
	time.Sleep(time.Second*1)
	return op
}
// 函数包一下 类似面向对象的 装饰者模式
func timeSpent(inner func(op int) int ) func(op int) int{
	return func(n int) int{
		start := time.Now()
		ret := inner(n)
		//打印输出时间
		fmt.Println("time spent:",time.Since(start).Seconds())
		return ret
	}
}

func TestFn(t *testing.T){
	a ,_ := returnMultiValues() // 忽略第二个值
	t.Log(a)

	tsSf := timeSpent(slowFun) //slowFun 函数
	t.Log(tsSf(10))
}

func Sum(ops ...int) int  {
	ret:=0
	for _,op:= range ops{
		ret += op
	}
	return ret;
}

func TestVarParams(t *testing.T){
   sum1 := Sum(1,1,2,3)
   sum2 := Sum(1,1,2,2,33)
   t.Log(sum1);
   t.Log(sum2);
}


func timeText(inner func(op int) int) func(op int) int{
	return func(n int) int{
		ret := inner(n)
		return ret
	}
}

func Clear() {
	fmt.Println("clear resources.")
}
func TestDefer(t *testing.T){
	defer Clear() // 延迟执行
	fmt.Println("start")
	panic("err")
    fmt.Println("end") // 在panic 后面代码不会执行 因此需要 defer来释放资源与解锁
}