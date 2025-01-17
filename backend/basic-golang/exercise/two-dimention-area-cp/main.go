package main

import (
	"fmt"
	"math"
)

// Check Point:
// Two Dimention Area
// - Input: Select Choice
// - Output: Result Operation

// Contoh:
// Input:
// 1: Rectange Area
// 2: Rectangular Area
// 3: Triangle Area
// 4: Circle Area
// - Enter choice 1, 2, 3, or 4: 1 | 2 | 3 | 4
//   - (1) Rectange Area
//   	- Masukkan sisi : 5
//   - (2) Rectangular Area
// 		- Masukkan panjang : 5
// 		- Masukkan lebar : 10
// 	 - (3) Triangle Area
// 		- Masukkan panjang alas segitiga: 5
// 		- Masukkan tinggi segitiga: 10
// 	 - (4) Circle Area
//      - Masukkan jari-jari : 4

// Output:
// - (1) Luas Persegi adalah : 25
// - (2) Luas Persegi Panjang adalah : 50
// - (3) Luas Segitiga adalah : 25
// - (4) Luas Lingkaran adalah : 50.240000

func main() {
	// TODO: answer here
	var s,s1,input float32
	
	fmt.Scan(&input)
	switch input {
	case 1:
		fmt.Scan(&s)
		fmt.Print(s*s)
		break
	case 2 :
		fmt.Scan(&s)
		fmt.Scan(&s1)
		fmt.Println(s*s1)
		break
	case 3 :
		fmt.Scan(&s)
		fmt.Scan(&s1)
		fmt.Println(s*s1/2)
		break
	case 4 :
		fmt.Scan(&s)
		fmt.Println(math.Pi*s*s)
		break
	}



}
