package main

import "fmt"

func main() {
	goodBye := getGoodBye

	result := goodBye("Bimlek")
	fmt.Println(result)
}

func getGoodBye(name string) string {
	return "Good Bye " + name
}
