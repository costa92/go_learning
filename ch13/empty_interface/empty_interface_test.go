package empty_interface

// 多态
import (
	"fmt"
	"testing"
)

func DoSomething(p interface{}){
	//if i,ok:=p.(int); ok{
	//	fmt.Println("Integer",i)
	//	return
	//}
	//
	//if s,ok :=p.(string);ok {
	//	fmt.Println("string",s)
	//	return
	//}
	//
	//fmt.Println("Unknow Type")

	// switch

	switch v :=p.(type) {
	case int:
		fmt.Println("Integer",v)
	case string:
		fmt.Println("string",v)
	default:
		fmt.Println("Unknow type")


	}
}

func TestEmptyInterfaceAssertion(t *testing.T){
	DoSomething(100)
}


