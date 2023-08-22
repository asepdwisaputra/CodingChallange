package main

import (
	"fmt"
)

func nilaiBMI(massa, tinggi float32) float32 {
	return massa / tinggi * tinggi
}

func perbandingan(Udin, Nanang float32) bool {
	if Udin > Nanang {
		return true
	} else {
		return false
	}
}

func main() {
	var BUdin, TUdin, BNanang, TNanang float32
	fmt.Print("Berat Udin:")
	fmt.Scan(&BUdin)
	fmt.Print("Tinggi Udin:")
	fmt.Scan(&TUdin)
	fmt.Print("Berat Nanang:")
	fmt.Scan(&BNanang)
	fmt.Print("Tinggi Nanang:")
	fmt.Scan(&TNanang)

	UdinBMI := nilaiBMI(BUdin, TUdin)
	fmt.Println(UdinBMI)

	NanangBMI := nilaiBMI(BNanang, TNanang)
	fmt.Println(NanangBMI)

	Tertinggi := perbandingan(UdinBMI, NanangBMI)
	if Tertinggi == true {
		fmt.Printf("BMI Udin(%f) lebih tinggi dari Nanang(%f)", UdinBMI, NanangBMI)
	} else {
		fmt.Printf("BMI Nanang(%f) lebih tinggi dari Udin(%f)", NanangBMI, UdinBMI)
	}

	//Test1
	//BUdin = 78, TUdin = 1.69 & BNanang = 92, TNanang = 1,95

	//Test
	//BUdin = 92, TUdin = 1.88 & BNanang = 85, TNanang = 1,76

}
