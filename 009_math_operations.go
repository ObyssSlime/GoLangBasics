// Package main provides ...
package main

import "fmt"

func main() {
	var a = 1
	var b = 3
	var c = a + b
	var d = 10
	var e = 20

	d += 11 // augmented assignments
	e++
	e++ // unary operators

	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
}
