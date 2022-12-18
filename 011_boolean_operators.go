// Package main provides ...
package main

import "fmt"

func main() {
	var (
		nilai   = 90
		absensi = 100

		lulusNilai   = nilai > 70
		lulusAbsensi = absensi > 80

		hasil = lulusNilai && lulusAbsensi
	)

	fmt.Println(hasil)
}
