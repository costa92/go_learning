package polymorphsim

import (
	"fmt"
	"testing"
)

type Code string

type Programmer interface {
	WriterHelloWorld() Code
}

// 定义 go
type GoProgrammer struct {

}

func (p *GoProgrammer) WriteHelloWorld() Code {
	panic("implement me")
}

//  go writerHelloWorld 方法
func (p *GoProgrammer) WriterHelloWorld() Code{
	return "fmt.PrintLn(\"Hello World !\")"
}

// 定义java
type JavaProgrammer struct {

}

func (p *JavaProgrammer) WriteHelloWorld() Code {
	panic("implement me")
}

func (p *JavaProgrammer) WriterHelloWorld() Code {
	return "System.out.Println(\"Hello World !\")"
}

func writeFirstProgram(p Programmer)  {
	fmt.Printf("%T %v\n",p,p.WriterHelloWorld())
}

func TestPolyMorphs(t *testing.T) {
  goProg := &GoProgrammer{}
  javaProg := new(JavaProgrammer)

  writeFirstProgram(goProg)
  writeFirstProgram(javaProg)
}
