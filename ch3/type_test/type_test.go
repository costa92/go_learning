package type_test

import "testing"
// 自定义的
type MyInt int64
func  TestImplicit( t *testing.T)  {
	 var  a int32  = 1
	 var  b int64
	 // 显示转换 不支持 隐转换
	 b = int64(a)

	 var c MyInt
	 c = MyInt(b)

	 t.Log(a,b,c)
}

func TestPoint(t *testing.T){
	a := 1
	aPtr := &a
	// 指针不支持运算
	//aPtr =  aPtr + 1
	t.Log(a,aPtr)
	t.Logf("%T %T",a,aPtr)
}

func TestString(t *testing.T){
 var s string
 // 字符串默认是 "" 而不是空字符串
 t.Log("*" + s+ "*")
 t.Log(len(s)) // 输出为 0
}
