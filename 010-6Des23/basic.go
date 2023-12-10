/*
	Pada problem ini kamu harus menemukan
	total maksimum jumlah bilangan dari deret sebuah integer secara berurutan.

	Sample Test Case
		Input: [-2, 1, -3, 4, -1, 2, 1, -5, 4]
		Output: 6
		Penjelasan: 6 adalah hasil penambahan dari deret 4, -1, 2, 1

	Sample Test Case
		Input: [-2, -5, 6, -2, -3, 1, 5, -6]
		Output: 7
		Penjelasan: 7 adalah hasil penambahan dari deret 6, -2, -3, 1, 5
*/

// Algoritma Kadane => https://www.youtube.com/watch?v=GrNSGC8Z2T0

package main

import "fmt"

func MaxSequence(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	nilaiMax := arr[0]
	nilaiSejauhIni := arr[0]

	for i := 1; i < len(arr); i++ { // Kenapa dari index 1 = karena kita akn membandingkan dengan index 0
		nilaiSejauhIni = ManaMax(arr[i], arr[i]+nilaiSejauhIni)
		// nilaiSejauhIni = nilai maks antara (index saat ini, ATAU, index saat ini + nilai maks sejauh ini)
		nilaiMax = ManaMax(nilaiMax, nilaiSejauhIni)
		// nilai maks = nilai tertinggi antara nilai maks yang ada sebelumnya dibanding nilai sejauh ini(biasanya akan ada perubahan di nilai sejauh ini)
	}
	return nilaiMax
}

func ManaMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(MaxSequence([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})) // 6
	fmt.Println(MaxSequence([]int{-2, -5, 6, -2, -3, 1, 5, -6}))   // 7
	fmt.Println(MaxSequence([]int{-2, -3, 4, -1, -2, 1, 5, -3}))   // 7
	fmt.Println(MaxSequence([]int{-2, -5, 6, -2, -3, 1, 6, -6}))   // 8
	fmt.Println(MaxSequence([]int{-2, -5, 6, 2, -3, 1, 6, -6}))    // 12
}
