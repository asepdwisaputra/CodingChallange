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
	"os"
)

func main() {
	// Deklarasi variabel yang dibutuhkan(gunakan variabel seminimal mungkin untuk meminimalkan penggunaan RAM)
	// BigO(1) Konstant
	var inputA, inputB uint8

	// Pengisian nilai input
	// Karena saya hanya mengalokasikan 8-bit uint maka cakupan nilai input hanya 0-255
	inputA = 86
	inputB = 85

	// Pencarian Hasil dengan BigO(1) Konstant
	if inputA%3 == 1 && inputB%3 == 2 || inputA%3 == 2 && inputB%3 == 1 {
		fmt.Println("YES")
		os.Exit(0) // OPSIONAL - Fungsinya: Menghentikan Program yang berjalan
	}
	if (inputA+inputB)%3 == 0 {
		fmt.Println("YES")
		os.Exit(0)
	}
	fmt.Println("NO")
}

/*
KESIMPULAN
	 Secara keseluruhan, baik time complexity maupun space complexity dari kode yang adalah
	 konstan (O(1)), yang berarti kinerjanya tidak bergantung pada ukuran atau kompleksitas
	 dari input yang diberikan.
*/
