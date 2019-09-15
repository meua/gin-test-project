package encapsulation

import (
	"fmt"
	"testing"
	"unsafe"
)

func TestEncapsulation(t *testing.T) {
	e := Employee{1, "Navy", 26}
	e1 := Employee{Name: "arenas", Age: 27}
	e2 := new(Employee)
	e2.Id = 23
	e2.Name = "Lily"
	e2.Age = 28
	t.Log(e)
	t.Log(e1)
	t.Log(e1.Id)
	t.Log(e2)
	t.Logf("e is %T", e)
	t.Logf("e is %T", e2)
}

type Employee struct {
	Id   int
	Name string
	Age  int
}

func TestStructOparations(t *testing.T) {
	e := Employee{0, "Bob", 20}
	fmt.Printf("Address2 is %x\n", unsafe.Pointer(&e.Name))
	t.Log(e.String())
}

func (e Employee) String() string {
	fmt.Printf("Address1 is %x\n", unsafe.Pointer(&e.Name))
	return fmt.Sprintf("ID:%d,Name:%s,Age:%d", e.Id, e.Name, e.Age)
}

/*func (e *Employee)  String() string {
	fmt.Printf("Address1 is %x\n", unsafe.Pointer(&e.Name))
	return fmt.Sprintf("ID:%d,Name:%s,Age:%d", e.Id, e.Name, e.Age)
}*/
