package main

import "fmt"

func main() {
	name := "Jawi"

	switch name {
	case "Bilek":
		fmt.Println("Halo Bilek salam kenal")
	case "Jawir":
		fmt.Println("Halo Jawir, sepertinya aku mengenal mu")
	default:
		fmt.Println("Kenalan dulu ya bang")
	}

	// switch lenght := len(name); lenght > 4 {
	// case true:
	// 	fmt.Println("Nama terlalu panjang")
	// case false:
	// 	fmt.Println("Nama sudah benar")
	// }

	length := len(name)
	switch {
	case length > 7:
		fmt.Println("Nama Terlalu Panjang")
	case length > 4:
		fmt.Println("Nama Lumayan Panjang")
	default:
		fmt.Println("Nama sudah benar")
	}
}
