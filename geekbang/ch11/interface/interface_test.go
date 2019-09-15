package _interface

import "testing"

type Programmer interface {
	WriteHelloworld() string
}

type GoProgrammer struct {
}

func (p *GoProgrammer) WriteHelloworld() string {
	return "fmt.Print(\"HelloWorld\")"
}

func TestClient(t *testing.T) {
	var p Programmer = new(GoProgrammer)
	t.Log(p.WriteHelloworld())
}
