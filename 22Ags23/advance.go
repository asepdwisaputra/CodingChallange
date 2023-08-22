/*
	Output:
	54321
	4321
	321
	21
	1
*/

package main

import "fmt"

func main() {
	fmt.Print("Masukan Input Angka:")
	var angka int
	fmt.Scan(&angka)
	n := &angka
	for i := 1; i <= *n; i++ { // cetak n baris
		for j := 0; j < *n-i+1; j++ { // isi
			fmt.Print(*n - j - i + 1)
		}
		fmt.Println()
	}
}
