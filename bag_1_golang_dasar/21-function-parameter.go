package main

import "fmt"

func main() {
	firstName := "Gwej"
	sayHelloTo(firstName, "Al-slur")
}

func sayHelloTo(firstName string, lastName string) {
	fmt.Println("Hello", firstName, lastName)
}
