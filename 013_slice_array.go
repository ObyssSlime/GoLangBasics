// Package main provides ...
package main

import "fmt"

func main() {
	names := [...]string{"ade", "arya", "bimantara", "universitas", "raharja"}
	slice := names[0:2]

	fmt.Println(slice[0])
	fmt.Println(slice[1])
}
