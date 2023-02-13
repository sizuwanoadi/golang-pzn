package main

import "fmt"

func main() {
	type NoKTP string
	type Married bool

	var noKtpAdi NoKTP = "10102032103941"
	var marriedStatus Married = false

	fmt.Println(noKtpAdi)
	fmt.Println(marriedStatus)
}
