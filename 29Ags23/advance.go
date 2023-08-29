package main

import "fmt"

func BilPrima(angka int) bool {
	if angka <= 1 {
		return false
	}

	if angka == 2 {
		return true
	}

	if angka%2 == 0 {
		return false
	}

	// 5,7 = 35
	for i := 3; i*i <= angka; i += 2 {
		if angka%i == 0 {
			return false
		}
	}
	return true
}

func BilPrimaRentang(angka1, angka2 int) {
	fmt.Printf("Bilangan prima antara %d dan %d adalah : \n", angka1, angka2)
	for i := angka1; i <= angka2; i++ {
		if BilPrima(i) {
			fmt.Println(i)
		}
	}
}

func main() {
	var angka1, angka2 int
	fmt.Print("Masukan Rentang Awal Bilangan Bulat:")
	fmt.Scanln(&angka1)
	fmt.Print("Masukan Rentang Akhir Bilangan Bulat:")
	fmt.Scanln(&angka2)

	BilPrimaRentang(angka1, angka2)
}
