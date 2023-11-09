// BACKLOG
// BACKLOG
// BACKLOG

// /*
// 	Soal include dalam folder 006

// 	Mulai:
// 	Inisialisasikan papan catur berukuran K x K.

// 	Iterasi Kuda Pertama:
// 	Lakukan iterasi untuk setiap sel di papan catur.
// 	Misalnya, gunakan loop bersarang untuk menentukan posisi kuda pertama.

// 	Iterasi Kuda Kedua:
// 	Lakukan iterasi untuk setiap sel di papan catur lagi.
// 	Pastikan sel untuk kuda kedua tidak sama dengan sel untuk kuda pertama.

// 	Cek Konflik:
// 	Periksa apakah posisi kuda kedua menyerang kuda pertama atau tidak.
// 	Gunakan aturan langkah kuda (L-shape) untuk menentukan kemungkinan serangan.

// 	Tindakan Berdasarkan Hasil Pengecekan:
// 	Jika tidak ada konflik, catat atau tampilkan posisi kuda pertama dan kuda kedua yang valid.
// 	Jika ada konflik, lanjutkan ke langkah berikutnya.

// 	Akhir Iterasi:
// 	Selesaikan iterasi untuk semua kemungkinan posisi kuda pertama dan kuda kedua.

// 	Tampilkan Hasil:
// 	Tampilkan hasil atau jumlah total posisi yang valid.

// 	Selesai:
// 	Selesaikan algoritma.
// */

// package main

// import "fmt"

// func isAttacking(x1, y1, x2, y2 int) bool {
// 	// Periksa apakah dua kuda saling menyerang
// 	dx := x1 - x2
// 	dy := y1 - y2

// 	return (dx == 1 && dy == 2) || (dx == 2 && dy == 1)
// }

// func countValidPositions(n int) int {
// 	// Hitung posisi valid untuk dua kuda di papan catur berukuran NxN
// 	count := 0

// 	// Loop mewakili baris knight ke-1
// 	for i := 1; i <= n; i++ {
// 		// Loop mewakili kolom knight ke-1
// 		for j := 1; j <= n; j++ {
// 			// Loop mewakili baris knight ke-2
// 			for k := 1; k <= n; k++ {
// 				// Loop mewakili kolom knight ke-2
// 				for l := 1; l <= n; l++ {
// 					if !isAttacking(i, j, k, l) {
// 						count++
// 					}
// 				}
// 			}
// 		}
// 	}

// 	return count
// }

// func main() {
// 	// Contoh dengan papan catur berukuran 5x5
// 	boardSize := 8
// 	result := countValidPositions(boardSize)

//		fmt.Printf("Jumlah posisi valid untuk dua kuda di papan catur %dx%d: %d\n", boardSize, boardSize, result)
//	}

// ------------------------------------------------------------------------------------------------

// package main

// import "fmt"

// func isAttacking(x1, y1, x2, y2 int) bool {
// 	// Check if two knights are attacking each other
// 	dx := x1 - x2
// 	dy := y1 - y2

// 	if (dx == 1 && dy == 2) || (dx == 2 && dy == 1) {
// 		fmt.Printf("Kuda pada posisi (%d, %d) menyerang kuda pada posisi (%d, %d)\n", x1, y1, x2, y2)
// 		return true
// 	}

// 	return false
// }

// func main() {
// 	// Example with a 8x8 chessboard
// 	boardSize := 8

// 	for i := 1; i <= boardSize; i++ {
// 		for j := 1; j <= boardSize; j++ {
// 			for k := 1; k <= boardSize; k++ {
// 				for l := 1; l <= boardSize; l++ {
// 					isAttacking(i, j, k, l)
// 				}
// 			}
// 		}
// 	}
// }
