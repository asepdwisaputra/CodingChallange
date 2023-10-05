package main

import (
	"fmt"
	"strconv"
	"strings"
)

func TotalNilai(angka int) int {
	// Saya berencana memasukan nilai dari urutan ke sebuah temporary slice
	wadah := []string{}

	// Disini saya ingin mengkonversi input parameter ke bentuk 2, 22, 222, ...
	var nilai strings.Builder
	for angka > 0 {
		nilai.WriteString("2")
		wadah = append(wadah, nilai.String()) // memasukan ke wadah
		angka--
	}

	// Disini saya ingin memecah setiap wadah tadi
	// Kemudian didalamnya ada konversi string>>int
	// Kemudian Jumlahkan setiap isinya
	var result int
	for _, isi := range wadah {
		nilai, _ := strconv.Atoi(isi)
		result = result + nilai
	}
	return result
}

func main() {
	var input int
	fmt.Print("Masukan Jumlah Angka Urutan: ")
	fmt.Scanln(&input)

	if input <= 0 {
		fmt.Println("Masukan Urutan lebih dari 0!")
	}

	hasil := TotalNilai(input)

	fmt.Printf("Jumlah nilai urutan kamu adalah: %d", hasil)

	/*
		Kelemahan:
		Bila inputan terlalu besar pasti akan membutuhkan waktu yang lama.
		Karena inputan akan dikerjakan satu-persatu(Brute Force).
	*/
}
