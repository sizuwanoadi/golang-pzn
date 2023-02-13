package main

import "fmt"

func main() {
	runApp(true)
}

func endApp() {
	message := recover()
	if message != nil {
		fmt.Println("Error dengan message", message)
	}
	fmt.Println("End App")
}

func runApp(error bool) {
	defer endApp()
	fmt.Println("App running")
	if error {
		panic("ERROR")
	}
}
