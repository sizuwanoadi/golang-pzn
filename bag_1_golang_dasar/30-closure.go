package main

import "fmt"

func main() {
	name := "Adi"
	counter := 0

	increment := func() {
		name := "Budi"
		counter++
		fmt.Println(name, counter)
	}

	increment()
	increment()

	fmt.Println(counter, name)
}
