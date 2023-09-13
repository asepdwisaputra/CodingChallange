/*
	ATTENTION PLEASE - ATTENTION PLEASE

	Halo mas/mba penilai, ini codingan saya bnyak bnyt tanya ke mbah GPT(ChatGPT).
	Kalo dibilang mengasah kemampuan ini mengasah bgt si.
	Tapi dibilang kemampuan sendiri ini ga 100% kemampuan sendiri.

	Jadi, kalo bisa 'jangan dikasih menang' saya ya. Ga puas saya kalo menang jalur nanya GPT.
	Mohon maaf kelancangannya, senang bisa meramaikan Coding Challenge🙌😊

	btw hampir 12jam mas, cuma buat ginian😭😿
*/

package main

import (
	"fmt"
	"regexp"
	"time"
)

func format(input string) bool {
	pattern := `^(0[0-9]|1[0-2]):(0[0-9]|[1-5][0-9])(AM|PM)$`

	// Compile regex
	re := regexp.MustCompile(pattern)

	// Mencocokkan string dengan pola regex
	if re.MatchString(input) {
		return true
	} else {
		return false
	}
}

func kapanMakanLagi(jamMakan time.Time) {
	layout := time.Kitchen
	breakfast, _ := time.Parse(layout, "07:00AM")
	lunch, _ := time.Parse(layout, "00:00PM")
	dinner, _ := time.Parse(layout, "07:00PM")

	// Mencari sisa waktu dari jam makan ke makan selanjutnya
	var sisaWaktu time.Duration
	switch {
	case jamMakan.After(dinner):
		waktuBesokPagi := breakfast.Add(24 * time.Hour)
		sisaWaktu = waktuBesokPagi.Sub(jamMakan)
		waktu := pecahWaktu(sisaWaktu)
		fmt.Printf("%d Hour and %d minutes until the next meal, breakfast tomorrow morning", waktu[0], waktu[1])
	case jamMakan.Before(breakfast):
		sisaWaktu = breakfast.Sub(jamMakan)
		waktu := pecahWaktu(sisaWaktu)
		fmt.Printf("%d Hour and %d minutes until the next meal, breakfast", waktu[0], waktu[1])
	case jamMakan.After(lunch):
		sisaWaktu = dinner.Sub(jamMakan)
		waktu := pecahWaktu(sisaWaktu)
		fmt.Printf("%d Hour and %d minutes until the next meal, dinner", waktu[0], waktu[1])
	case jamMakan.After(breakfast):
		sisaWaktu = lunch.Sub(jamMakan)
		waktu := pecahWaktu(sisaWaktu)
		fmt.Printf("%d Hour and %d minutes until the next meal, lunch", waktu[0], waktu[1])
	default:
		fmt.Println("Nampaknya ada yang salah!😵‍💫")
	}
}

func pecahWaktu(sisaWaktu time.Duration) [2]int {
	jam := int(sisaWaktu.Hours())
	menit := int(sisaWaktu.Minutes()) % 60
	result := [2]int{jam, menit}
	return result
}

func main() {
	var input string

	fmt.Print("🍴 Jam makan terakhir(contoh-> 01:00PM) : ")
	fmt.Scanln(&input)
	if format(input) == true {
		// kondisi true
		layout := time.Kitchen
		jamMakan, _ := time.Parse(layout, input)
		fmt.Println(jamMakan)
		kapanMakanLagi(jamMakan)
	} else {
		// false kondisi
		fmt.Printf("Inputan kamu '%s' nampaknya salah😵‍💫\n\n", input)
		ketentuan := "Input yang kamu masukan terdeteksi tidak sesuai ketentuan/contoh!\nKetentuan(contoh-> 01:00PM):\n- Penulisan jam dan menit dipisahkan dengan tanda ':'\n- Penulisan waktu dengan format 12 Jam\n- Gunakan 'AM' KAPITAL untuk menandakan sebelum tengah hari\n- Gunakan 'PM' KAPITAL untuk menandakan sesudah tengah hari\n- Tidak ada karakter/simbol tambahan ya, Good Luck🫡"
		fmt.Print(ketentuan)
	}
}
