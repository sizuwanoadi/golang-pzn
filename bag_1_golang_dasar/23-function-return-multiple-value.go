package main

import "fmt"

func main() {
	firstName, _, lastName := getFullname()
	fmt.Println(firstName, lastName)
}

func getFullname() (string, string, string) {
	return "Adi", "Sizu", "Siswanto"
}
