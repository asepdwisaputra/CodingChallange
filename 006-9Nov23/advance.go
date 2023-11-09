/*
	HAPPY NUMBER

	Tugas : Menentukan apakah sebuah anggka BAHAGIA

	Ciri angka bahagia :
		- Dimulai dengan bilangan bulat positif apa pun, ganti bilangan tersebut dengan jumlah kuadrat angka-angkanya.
		- Ulangi proses ini sampai angkanya sama dengan 1 (di mana angka tersebut akan tetap ada), atau angka tersebut berputar tanpa henti dalam siklus yang tidak menyertakan 1.
		- Angka-angka yang proses ini berakhir dengan 1 adalah angka-angka bahagia.

	True jika angka = BAHAGIA atau sebaliknya

	Input --> n integer

	Output dan Penjelasan :
		Input --> 19
		Output --> true
		Penjelasan --> 	1^2 + 9^2 = 82
						8^2 + 2^2 = 68
						6^2 + 8^2 = 100
						1^2 + 0^2 + 0^2 = 1

	KRITERIA PENILAINAN
	- Nilai akan disortir dari besar-kecil(DESC)
	- Dokument Code yang baik:
		- TO-DO setiap baris kode solusi
		- Ringkas, namun tepat
		- Penjelasan mandiri mengenai Time Complexity dan Space Complexity yang diimplementasikan pada solusi yang diberikan.
*/

package main

import (
	"fmt"
)

// Global variable untuk melihat urutan perulangan
var Perulangan int

// Menjadikan nilai input menjadi digit-digit angka
func CariDigit(n int) []int {
	var digits []int
	var digit int

	for n > 0 {
		digit = n % 10
		digits = append([]int{digit}, digits...)
		n /= 10
	}

	return digits
}

// Mencari nilai apakah bernilai akhir 1 dan bahagia
func CariNilai(digits []int) bool {
	var result int

	for _, digit := range digits {
		result = result + digit*digit
	}

	// Melihat Urutan perulangan dan nilainya
	fmt.Println("Perulangan ke", Perulangan, " bernilai", result)

	// Kondisi nilai 4 artinya perulangan tak terbatas(false)
	if result == 4 {
		return false
	}

	// Kondisi perulangan mencapai 100 dianggap false
	if Perulangan == 100 {
		return false
	}

	if result == 1 {
		return true
	} else {
		Perulangan = Perulangan + 1
		return CariNilai(CariDigit(result))
	}
}

func main() {
	// Membuat variabel input yang bernilai int(bulat positif)
	var input int = 5
	Perulangan = 0

	fmt.Println(CariNilai(CariDigit(input)))
}

// DONE
