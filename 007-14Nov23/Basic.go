package main

import "fmt"

func temukanAngka(angka []int) int {
	panjangAngka := len(angka)
	totalPerkiraan := (panjangAngka + 1) * (panjangAngka + 2) / 2
	totalYangAda := 0

	for _, nilai := range angka {
		totalYangAda += nilai
	}

	return totalPerkiraan - totalYangAda
}

func main() {
	angka := []int{2, 3, 1, 5}
	fmt.Println(temukanAngka(angka)) // 4
}
