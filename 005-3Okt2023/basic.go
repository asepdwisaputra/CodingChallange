/*
	Tuti dan Nining sedang melakukan studi seputar kebiasaan kucing.
	Masing-masing dari mereka bertanya kepada 5 pemilik kucing
	tentang usia Kucing Mereka, dan menyimpan datanya ke dalam sebuah
	array (satu array untuk masing-masing).
	Saat ini, mereka hanya tertarik untuk mengetahui apakah seekor Kucing Adalah
	Kucing Dewasa atau Kucing Kecil (Kitten).
	Sebuah Kucing Dianggap dewasa jika usianya setidaknya 3 tahun,
	dan dianggap Kucing Kecil (Kitten) jika usianya kurang dari 3 tahun.

	Buatlah sebuah fungsi ‘checkCats’, yang menerima 2 array usia
	Kucing(‘CatsTuti’ dan ‘CatsNining’), dan melakukan hal-hal berikut:

	1. Tuti mengetahui bahwa pemilik dari Kucing Pertama dan dua Kucing
	Terakhir sebenarnya memiliki anjing, bukan kucing!
	Jadi, buatlah salinan dari array milik Tuti, dan hapus usia kucing
	dari array yang disalin tersebut.

	2. Buatlah sebuah array yang berisi data Tuti (yang sudah dikoreksi)
	dan data Nining.

	3. Untuk setiap Kucing Yang tersisa, tampilkan di konsol
	apakah itu Kucing Dewasa (“Kucing Nomor 1 adalah Kucing Dewasa,
	dan berusia 5 tahun”) atau Kucing Kecil (“Kucing Nomor 2 masih anak-anak  ”).

	Jalankan fungsi tersebut untuk kedua set data uji.

	Data uji:
	Data 1: Data Tuti [3, 5, 2, 12, 7], Data Nining [4, 1, 15, 8, 3]
	Data 2: Data Tuti [9, 16, 6, 8, 3], Data Nining [10, 5, 6, 1, 4]
*/

package main

import "fmt"

func checkCats(CatsTuti, CatsNining []int) {
	for i, v := range CatsTuti {
		if v >= 3 {
			fmt.Printf("Data Tuti: Kucing Nomor %d adalah Kucing Dewasa, dan berusia %d tahun\n", i+1, v)
		} else {
			fmt.Printf("Data Tuti: Kucing Nomor %d masih anak-anak, dan berusia %d tahun\n", i+1, v)
		}

	}

	for i, v := range CatsNining {
		if v >= 3 {
			fmt.Printf("Data Nining: Kucing Nomor %d adalah Kucing Dewasa, dan berusia %d tahun\n", i+1, v)
		} else {
			fmt.Printf("Data Nining: Kucing Nomor %d masih anak-anak, dan berusia %d tahun\n", i+1, v)
		}
	}
}

func main() {
	DataTutiAwal := [5]int{3, 5, 2, 12, 7}
	DataNiningAwal := [5]int{4, 1, 15, 8, 3}
	//Ternyata data 1 dan 2 terakhir milik tuti tidak valid
	DataTutiFix := DataTutiAwal[1:3]
	DataNiningFix := DataNiningAwal[:]

	checkCats(DataTutiFix, DataNiningFix)
}
