/*
	Cecep ada seorang programmer di sebuah startup di australia,
	kemudian dia ingin membuat kalkulator tip
	yang sangat sederhana untuk setiap kali dia makan di restoran.
	Di indonesia cecep tidak pernah mengalami hal ini tapi di australia,
	biasanya memberikan tip sebesar 15% jika nilai tagihan berada di antara 50 dan 300.
	Jika nilainya berbeda, maka tipsnya adalah 20%.
*/

package main

import "fmt"

func main() {
	var tagihan, tip, total, X float32

	// Go tidak suport Ternary Operator

	fmt.Print("Nilai Tagihan:")
	fmt.Scanln(&tagihan)

	if X >= 50 && X <= 300 {
		tip = tagihan * 0.2
	} else {
		tip = tagihan * 0.15
	}

	total = tagihan + tip

	fmt.Printf("Tagihannya %.2f, Tip-nya %.2f, dan Total Nilainya %.2f", tagihan, tip, total)

	// Data 1: Uji untuk nilai tagihan 275, 40, dan 430
}
