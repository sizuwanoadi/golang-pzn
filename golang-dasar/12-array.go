package main

import "fmt"

func main() {
	var names [3]string

	names[0] = "Calliope"
	names[1] = "Adi"
	names[2] = "Siswanto"

	fmt.Println(names)
	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var values = [3]int{
		10,
		20,
		30,
	}
	fmt.Println(values)

	var bilek [10]string

	fmt.Println(len(names))
	fmt.Println(len(values))
	fmt.Println(len(bilek))
}
