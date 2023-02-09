package main

import "fmt"

func main() {
	total := sumAll(10, 11, 12, 13)
	fmt.Println(total)

	slice := []int{1, 2, 3, 4, 5}
	total = sumAll(slice...)
	fmt.Println(total)
}

func sumAll(numbers ...int) int {
	total := 0
	for _, v := range numbers {
		total += v
	}
	return total
}
