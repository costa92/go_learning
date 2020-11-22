package reflect

import (
	"fmt"
	"reflect"
	"testing"
)
type Employee struct {
	EmployeeId string
	Name  string `format:"normal"`
	Age  int
}

func (e *Employee) UpdateAge(newVal int){
	e.Age = newVal
}


type Coustomer struct {
	CookieId string
	Name     string
	Age     int
}


func TestInvokeByName(t *testing.T){
	e :=&Employee{"1","Mike",30}
	// 按名字获取成员
	t.Logf("Name: value(%[1]v),Type(%[1]T)",reflect.ValueOf(*e).FieldByName("Name"))
	if nameField,ok := reflect.TypeOf(*e).FieldByName("Name");!ok {
		t.Error("Failed to get 'name' field.")
	}else{
		t.Log("Tag:format",nameField.Tag.Get("format"))
	}

	reflect.ValueOf(e).MethodByName("UpdateAge").Call([]reflect.Value{reflect.ValueOf(6)})
	t.Log("Update Age",e)
}

func TestTypeAndValue(t *testing.T){
	var f int64 = 10
	t.Log(reflect.TypeOf(f),reflect.TypeOf(f))
	t.Log(reflect.ValueOf(f).Type())
}

func CheckType(v interface{}){
	t := reflect.TypeOf(v)
	switch t.Kind() {
	case reflect.Float32,reflect.Float64:
		fmt.Println("Float")
	case reflect.Int,reflect.Int32,reflect.Int64:
		fmt.Println("Integer")
	default:
		fmt.Println("Unknown",t)
	}
}

func TestBasicType(t *testing.T){
	var f float64 =12

	CheckType(f)
	CheckType(&f)
}