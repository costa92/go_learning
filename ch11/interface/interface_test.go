package interface_test

import "testing"

// 接口
type Programmer interface {
	WriterHelloWord() string
}

// 定义类
type GoProgrammer struct {

}

// 实现接口
func (Go *GoProgrammer) WriterHelloWord()  string {
	return "fmt.Println(\"Hello World\")"
}

func TestClient(t *testing.T) {
	var p Programmer
	p = new(GoProgrammer)
	t.Log(p.WriterHelloWord())
}