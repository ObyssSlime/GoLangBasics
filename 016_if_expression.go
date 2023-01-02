// Package main provides ...
package main

import "fmt"

func main() {
	name := "yug"

	if name == "ade" {
		fmt.Println("hello", name)
	} else if name == "joko" {
		fmt.Println("hello", name)
	} else if length := len(name); length > 4 {
		fmt.Println("hello new entity")
	} else {
		fmt.Println("anda siapa ya?")
	}

}
