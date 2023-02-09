package main

import "fmt"

func main() {
	name := getHello("Bimlek")
	fmt.Println(name)

	fmt.Println(getHello(""))
}

func getHello(name string) string {
	if name == "" {
		return "Hello Lur"
	} else {
		return "Hello " + name
	}
}
