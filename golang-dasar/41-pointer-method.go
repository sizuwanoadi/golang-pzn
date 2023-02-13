package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	adi := Man{"Adi"}
	adi.Married()
	fmt.Println(adi)
}
