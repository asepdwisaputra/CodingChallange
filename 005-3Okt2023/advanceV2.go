package main

import "fmt"

func PangkatV2(dasar int) int {
	nilai := 1
	for i := 1; i <= 3; i++ {
		nilai = nilai * dasar
	}
	return nilai
}

func main() {
	nilaiDasar := 6

	// Cetak nilai dari 1^3 sampai 6^3
	for i := 1; i <= nilaiDasar; i++ {
		fmt.Printf("Curent Number is: %d and the cube is %d\n", i, PangkatV2(i))
	}
}
