package array_test

import "testing"

func TestArratInit(t *testing.T){
	var arr [3] int  // 声明
	arr1 :=[4]int{1,2,3,4}
	arr2 :=[...]int{1,3,4,5}
	t.Log(arr[1],arr[2])
	t.Log(arr1,arr2)
}

// 循环
func TestArrayTravel(t *testing.T){
	arr3 := [...] int{1,2,3,4,5}
	for i:=0;i<len(arr3);i++  {
		t.Log(arr3[i])
	}

	for ind,e := range arr3{
		t.Log(ind)
		t.Log(e)
	}


	for _,e := range arr3{
		t.Log(e)
	}
}


func TestArraySection(t *testing.T) {
	arr3 := [...]int{1,2,3,4,5}
	arr3_sec := arr3[0:1] // 不支持 负数切片
	t.Log(arr3_sec)
}