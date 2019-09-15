package extension

import (
	"fmt"
	"testing"
)

type Pet struct {
}

func (p *Pet) speak() {
	fmt.Println("...")
}

func (p *Pet) speakTo(name string) {
	p.speak()
	fmt.Println(" ", name)
}

type Dog struct {
	Pet
}

func (d *Dog) speak() {
	//d.speak()
	fmt.Println("wang.")
}

//
//func (d *Dog) speakTo(name string){
//	//d.p.speakTo(name)
//	d.speak()
//	fmt.Println(" ", name)
//}

func TestDog(t *testing.T) {
	//dog := new(Dog)
	dog := new(Dog)
	//p:= dog.(Pet)
	//dog.speakTo("Chao")
	dog.speak()
}
