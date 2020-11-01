package encap

import (
	"fmt"
	"testing"
	"unsafe"
)


type Employee struct {
	Id string
	Name string
	Age int
}

//func TestCreateEmployeeObj(t *testing.T) {
//
//	e := Employee{"1","11",11}
//	e1 := Employee{Name:"Mike",Age: 20}
//	e2 := new(Employee) // 返回指针
//
//	e2.Id = "1"
//	e2.Name = "Rose"
//	e2.Age = 22
//
//	t.Log(e)
//	t.Log(e1)
//	t.Log(e2)
//	t.Logf("e is %T",&e) // 转换成指针类型
//	t.Logf("e2 is %T",e2) // 转换成指针类型
//}

// 第一种定义方式在实例对应方法被调用时，实例的成员进行赋值
func (e Employee) String() string{
	fmt.Printf("Address is %x\n",unsafe.Pointer(&e.Name))
	return fmt.Sprintf("ID:%s-Name:%s-Age:%d",e.Id,e.Name,e.Age)
}
// 通常情况下为了避免内存拷贝我们使用第二种定义方式
//func (e *Employee) String() string{
//	fmt.Printf("Address is %x\n",unsafe.Pointer(&e.Name))
//	return fmt.Sprintf("ID:%s/Name:%s/Age:%d",e.Id,e.Name,e.Age)
//}

func TestStructOperations(t *testing.T) {
	e := Employee{"0","Bob",20}
	fmt.Printf("Address is %x\n",unsafe.Pointer(&e.Name))
	t.Log(e.String())
}