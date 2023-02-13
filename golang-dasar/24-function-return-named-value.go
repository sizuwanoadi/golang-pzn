package main

import "fmt"

func main() {
	firstName, _, lastName := getFullname()
	fmt.Println(firstName, lastName)
}

func getFullname() (firstName string, middleName string, lastName string) {
	firstName = "Adi"
	middleName = "Siswanto"
	lastName = "BImlek"

	return
}
