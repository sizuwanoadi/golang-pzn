package main

import "fmt"

func main() {
	person := map[string]string{
		"Name":  "Adi Siswanto",
		"Age":   "??",
		"Hobby": "Gaming",
	}

	person["Job"] = "Programmer coy"

	fmt.Println(person)
	fmt.Println(person["Name"])
	fmt.Println(person["Hobby"])

	var book = make(map[string]string)
	book["title"] = "Adalah gwej"
	book["author"] = "Seorang yang bernama bilek"
	book["ups"] = "Waduh, pakek nanya"

	fmt.Println(book)
	fmt.Println(len(book))
	delete(book, "ups")
	fmt.Println(book)
	fmt.Println(len(book))
}
