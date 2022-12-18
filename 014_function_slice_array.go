// Package main provides ...
package main

import "fmt"

func main() {
	days := [...]string{"senin", "selasa", "rabu"}
	daySlice := days[1:]
	daySlice[0] = "kamis"
	daySlice[1] = "jumat"

	fmt.Println(days)
	fmt.Println(daySlice)

	daySlice2 := append(daySlice, "sabtu") // append slice function

	fmt.Println(daySlice2)

	daySlice2[0] = "minggu"

	fmt.Println(daySlice2)
	fmt.Println(days)

	newslice := make([]string, 2, 5) // make slice function
	newslice[0] = "ade"
	newslice[1] = "arya"

	fmt.Println(newslice)
	fmt.Println(len(newslice))
	fmt.Println(cap(newslice))

	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))

	copy(toSlice, fromSlice) // copy slice function

	fmt.Println(fromSlice)
	fmt.Println(toSlice)
}
