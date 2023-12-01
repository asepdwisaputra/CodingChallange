package main

import "fmt"

func main() {
	input := "dasjgjidasgnasfjsjsaldjksfjdfjffffj"

	nilaiTerbesar := 0
	nilaiKini := 1 // Karena otomatis jika tidak ada perulangan  berarti semuanya 1x

	// Memberikan index pada karakter string input
	// (agar menjadi ASCII code ketika dipanggil dengan index)
	for i := 0; i < len(input)-1; i++ {

		// konversi ke string untuk melihat secara normal
		// fmt.Println(string(input[i]), string(input[i+1]))

		if input[i] == input[i+1] { // Pengecekan ASCII code ke-1 dan ke-1+1
			nilaiKini++
		} else {
			if nilaiKini > nilaiTerbesar {
				nilaiTerbesar = nilaiKini
			}
			nilaiKini = 1 // ulang dari 1
		}
	}
	if nilaiKini > nilaiTerbesar {
		nilaiTerbesar = nilaiKini
	}
	fmt.Println(nilaiTerbesar)
}
