package main

import "fmt"

type HasName interface {
	GetName() string
}

func SayHello(hasName HasName) {
	fmt.Println("Hello", hasName.GetName())
}

type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

func (animal Animal) GetName() string {
	return animal.Name
}

type Animal struct {
	Name string
}

func main() {
	var adi Person
	adi.Name = "Adi Siswanto"

	SayHello(adi)

	animal := Animal{Name: "Puss"}
	SayHello(animal)

}
