package err_test

import (
	"errors"
	"fmt"
	"testing"
)

// 定义常见错误类型
var LessThanTwoError = errors.New("n should be not less than 2")
var LargerThanHundredError  = errors.New("n should be not larger than 100")

func GetFibonacci(n int) ([]int,error) {

	if n < 2 { // 验证参数
		return nil,LessThanTwoError // 返回错误提示
	}

	if n > 100 { // 验证参数
		return nil,LargerThanHundredError // 返回错误提示
	}
	if n < 2 || n >100 { // 验证参数
		return nil,errors.New("n should be in [2,2000]") // 返回错误提示
	}

	fibList :=[]int{1,1}

	for i:=2 ;i < n ; i++{
		fibList = append(fibList,fibList[i-2]+fibList[i-1])
	}

	return fibList,nil
}

func TestGetFibonacci(t *testing.T) {
	if v,err := GetFibonacci(1);err !=nil {
		if err == LessThanTwoError{
			fmt.Println("It is less .")
		}
		t.Error(err)
	}else{
		t.Log(v)
	}
}