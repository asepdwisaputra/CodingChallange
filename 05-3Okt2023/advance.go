/*
	Menulis program:
	untuk memunculkan nilai pangkat dari 1^3 - n^3

	Input:
	n = 2

	Output:
	Nilai saat ini : 1 pangkatnya 1
	Nilai saat ini : 2 pangkatnya 2
*/

package main

import (
	"fmt"
	"math"
)

func Pangkat(n int) {
	result := math.Pow(float64(n), float64(3))
	fmt.Printf("Curent Number is: %v and the cube is %v\n", n, result)
}

func main() {
	n := 6

	for i := 1; i <= n; i++ {
		Pangkat(i)
	}
}
