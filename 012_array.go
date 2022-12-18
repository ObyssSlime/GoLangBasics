// Package main provides ...
package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "ade"
	names[1] = "arya"
	names[2] = "bimantara"

	var values = [3]int{ // array langsung dideklarasikan saat membuat variable
		10, 11, 12,
	}

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(len(names))
	values[0] = 100
	fmt.Println(values)
}
