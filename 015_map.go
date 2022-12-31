// Package main provides ...
package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "Ade",
		"address": "Tangerang",
	}

	book := make(map[string]string)
	book["title"] = "Buku Go-Lang"
	book["author"] = "Eko Kurniawan"
	book["wrong"] = "oops"

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	delete(book, "wrong")

	fmt.Println(book)
	fmt.Println(book["title"])
	fmt.Println(book["author"])

}
