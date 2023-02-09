package main

import "fmt"

func newMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main() {
	person := newMap("")
	fmt.Println(person)

	data := newMap("")
	if data == nil {
		fmt.Println("Data kosong")
	} else {
		fmt.Println(data)
	}
}
