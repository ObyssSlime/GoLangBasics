// Package main provides ...
package main

import "fmt"

func main() {
	nama := "dimas"

	switch nama {

	case "ade":
		fmt.Println("hello", nama)

	case "joko":
		fmt.Println("hello", nama)

	default:
		fmt.Println("hello new entity")
	}

}
