package reflect_test

import (
	"fmt"
	"reflect"
	"testing"
)



type Student struct {
	Name string
	Age  int
}

func (stu Student) Sum(a int , b int) int {
	return b + a
}

type Enum int

const (
	Zero Enum = 0
)


func TestReflectStruct(t *testing.T) {
	var stu interface{} = Student{}


	typeOfStu := reflect.TypeOf(stu)
	fmt.Println(typeOfStu.Name(),typeOfStu.Kind())
	typeOfZero := reflect.TypeOf(Zero)
	fmt.Println(typeOfZero.Name(),typeOfZero.Kind())


	valueOfStu := reflect.ValueOf(stu)
	sum := []reflect.Value{reflect.ValueOf(20),reflect.ValueOf(30)}

	method := valueOfStu.MethodByName("Sum")

	resultList := method.Call(sum)
	fmt.Println(resultList[0])
}

