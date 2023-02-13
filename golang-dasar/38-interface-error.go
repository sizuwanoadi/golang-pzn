package main

import (
	"errors"
	"fmt"
)

func Pembagi(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagi tidak boleh 0")
	}

	result := nilai / pembagi
	return result, nil
}

func main() {
	result, err := Pembagi(0, 10)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}
