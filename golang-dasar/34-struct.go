package main

import "fmt"

type Customer struct {
	Name, Adress string
	Age          int
}

func main() {
	var adi Customer
	adi.Name = "Adi Siswanto"
	adi.Adress = "West Java, Indonesia"
	adi.Age = 18
	fmt.Println(adi)

	joko_anwar := Customer{
		Name:   "Joko Anwar",
		Adress: "Jl. Gundala",
		Age:    41,
	}
	fmt.Println(joko_anwar)

	bimlek := Customer{Name: "Bimlek", Adress: "Kp. Mukbang", Age: 100}
	fmt.Println(bimlek)
	adi.sayHi()
	adi.sayHuuu()
}

func (customer Customer) sayHi() {
	fmt.Println("Hi,", customer.Name)
}

func (customer Customer) sayHuuu() {
	fmt.Println("Huuuu,", customer.Name)
}
