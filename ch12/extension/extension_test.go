package extension

import (
	"fmt"
	"testing"
)

type Pet struct {

}

func (p *Pet) Speak(){
	fmt.Print("...")
}

func (p *Pet) SpeakTo(host string){
	p.Speak()
	fmt.Print(" ",host)
}

type Dog struct {
	//p *Pet
	Pet // 匿名方式
}
//
//func (d *Dog) Speak(){
//	fmt.Print("Wang!")
//}
//
//func (d *Dog) SpeakTo(host string){
//	d.Speak()
//	fmt.Println(" ",host)
//	//d.p.SpeakTo(host)
//}


func (d *Dog) Speak(){
	fmt.Print("Wang!")
}

func TestDog(t *testing.T) {
	dog := new(Dog)
	dog.SpeakTo("Chao")
}


