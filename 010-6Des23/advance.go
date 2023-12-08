/*
*
KRITERIA PENILAIAN
  - Nilai Testing peserta akan disortir secara DESC (tertinggi ke terendah)
  - Dokumentasi kode yang baik, seperti:
    1. TODO setiap baris kode solusi
    2. Ringkas, namun tepat
    3. Penjelasan mandiri mengenai Time Complexity dan Space Complexity yang diimplementasikan pada solusi yang diberikan

*
*/

package main

import (
	"fmt"
	"math"
)

type Coordinate struct {
	X, Y float64
}

func main() {
	Koordinat := []Coordinate{
		{0, 0}, // titik 1
		{0, 2}, // titik 2
		{2, 2}, // titik 3
		{2, 0}, // titik 4
	}

	fmt.Println(NilaiJarak(Koordinat))
}

// Kita memerlukan rumus mencari jarak antar titik
// Jarak = akar dari (x2-x1)(x2-x1)+(y2-y1)(y2-y1)
func RumusJarak(x1, x2, y1, y2 float64) float64 {
	nilaiKuadrat := math.Pow((x2-x1), 2) + math.Pow((y2-y1), 2) // Total nilai kuadrat
	nilaiAkhir := math.Sqrt(float64(nilaiKuadrat))              // Nilai akar dari nilai kuadrat

	return float64(nilaiAkhir)
}

func NilaiJarak(input []Coordinate) float64 {
	var result float64 = 0

	for i := 0; i < len(input); i++ {
		if i == len(input)-1 { // Jika berada di titik terakhir
			result = result + RumusJarak(input[i].X, input[i].Y, input[0].X, input[0].Y)
		} else {
			result = result + RumusJarak(input[i].X, input[i].Y, input[i+1].X, input[i+1].Y)
		}
		//fmt.Println("Perulangan Ke", i+1, "Bernilai", result) // Jika ingin melihat nilai berjalan
	}

	return result
}

/*
KESIMPULAN
	 Secara keseluruhan, space complexity dari kode yang adalah
	 konstan O(1), sedangkan time complexity di O(n) karena
	 terdapat iterasi perulangan sebanyak n di fungsi NilaiJarak.
*/
