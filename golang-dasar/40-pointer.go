package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func ChangeCountryToIndonesia(address *Address) {
	address.Country = "Indonesia"
}

func main() {
	// pass by value
	address1 := Address{City: "Bogor", Province: "Jawa Barat", Country: "Indonesia"}
	address2 := address1

	address2.City = "Sukabumi"

	fmt.Println(address1)
	fmt.Println(address2)

	// pass by reference

	address3 := Address{City: "Bandung", Province: "Jawa Barat", Country: "Indonesia"}
	address4 := &address3
	var address5 *Address = &address3

	*address4 = Address{City: "Bilek", Province: "Lur", Country: "wkwkkw"}

	addressNew := new(Address)
	addressNew.City = "Bimlek"

	fmt.Println(address3)
	fmt.Println(address4)
	fmt.Println(address5)
	fmt.Println(addressNew)

	var alamat = Address{
		City:     "Bogor",
		Province: "Jawa Barat",
		Country:  "",
	}
	ChangeCountryToIndonesia(&alamat)

	fmt.Println(alamat)
}
