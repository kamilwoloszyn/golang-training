package interfaces

import "fmt"

type Person struct {
	name    string
	surname string
}

type greetings interface {
	print()
}

func (p *Person) print() {
	fmt.Print("Hello %s %s", p.name, p.surname)
}
