package main

import "fmt"

func main() {
	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Perhitungan ke -", counter)
	}

	slice := []string{"Sizu", "Adi", "Siswanto", "Bilek", "Slur"}
	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	for index, value := range slice {
		fmt.Println("index =", index, "value =", value)
	}

	for _, value := range slice {
		fmt.Println("value =", value)
	}

	person := make(map[string]string)
	person["firstName"] = "Adi"
	person["lastName"] = "Siswanto"
	person["company"] = "Sizuwano"

	for key, value := range person {
		fmt.Println("key =", key, "value =", value)
	}
}
