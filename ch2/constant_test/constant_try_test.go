package constant_test

import (
	"testing"
)

const (
	Monday  = iota + 1
	Tuesday
	Wednesday
)

const(
	Readable = 1 << iota
	Writable
	Executable
)

func TestConstantTry(t *testing.T)  {
	t.Log(Monday,Tuesday)
}

func TestConstanTry1(t *testing.T){
	//a := 7 // 0111  二进制
	a :=1  // 0001
	t.Log(a&Readable == Readable,a&Writable == Writable,a&Executable == Executable)
}