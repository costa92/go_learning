package operator_test

import "testing"

const(
	Readable = 1 << iota
	Writable
	Executable
)


func TestCompareArray(t *testing.T){
	a := [...]int{1,2,3,4}
	b := [...]int{1,2,3,5}
	//c := [...]int{1,2,3,4,5}
	d := [...]int{1,2,3,4}
	t.Log(a == b)
	//t.Log(a == c)
	t.Log(a == d)
}

// &^ 按位清零
func TestBitClear(t *testing.T){
	a := 7 // 0111  二进制
	a = a&^ Readable
	a = a&^ Writable   // 按位清零运行
	t.Log(a&Readable == Readable,a&Writable == Writable,a&Executable == Executable)
}