package testing

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSquare(t *testing.T) {
	inputs := [...]int{1,2,3}
	expected := [...]int{1,4,9}

	for i:=0; i< len(inputs); i++  {
		ret := square(inputs[i])
		if ret != expected[i]{
			t.Errorf("input is %d,the expected is %d, the actual %d",inputs[i],expected[i],ret)
		}
	}
}

func TestErrorInCode(t *testing.T){
	fmt.Println("Start")
	//t.Error("Error")
	t.Fatal("Error")  // 中断了
	fmt.Println("End")

	// go test -v -cover   查询测试的覆盖率
}

func TestSquareWithAssert(t *testing.T) {
	inputs := [...]int{1,2,3}
	expected := [...]int{1,4,9}

	for i:=0; i< len(inputs); i++  {
		ret := square(inputs[i])
		assert.Equal(t,expected[i],ret)
	}
}