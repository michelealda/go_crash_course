package main

import "fmt"

// Person struct
type Person struct {
	firstName string
	lastName  string
}

// Greeting method {value receiver}
func (p Person) greet() string {
	return "Hello " + p.firstName
}

// Pointer receiver
func (p *Person) withName(name string) {
	p.firstName = name
}

func main() {
	person1 := Person{firstName: "Alice", lastName: "Cooper"}

	person2 := Person{"John", "Doe"}

	fmt.Println(person2.greet())
	person1.withName("lalala")

	fmt.Println(person1, person2)
}
