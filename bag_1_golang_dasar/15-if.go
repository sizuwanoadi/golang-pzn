package main

import "fmt"

func main() {
	var name = "12"

	if name == "Bilek" {
		fmt.Println("Jadi kamu yang namanya Bilek")
	} else if name == "Pov" {
		fmt.Println("Jadi kamu yang namanya Pov")
	} else {
		fmt.Println("Sorry bang gak kenal")
	}

	if length := len(name); length > 3 {
		fmt.Println("Nama lebih dari 3 huruf")
	} else {
		fmt.Println("Nama kependekan")
	}
}
