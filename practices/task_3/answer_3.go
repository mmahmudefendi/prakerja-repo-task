package main

import (
	"fmt"
)

func munculSekali(angka string) []int {
	// inisialisasi map untuk menghitung jumlah kemunculan setiap angka
	count := make(map[int]int)
	for _, digit := range angka {
		count[int(digit-'0')]++
	}

	// buat slice untuk menyimpan angka yang hanya muncul 1 kali
	var result []int
	for _, digit := range angka {
		if count[int(digit-'0')] == 1 {
			result = append(result, int(digit-'0'))
		}
	}

	fmt.Println()
	fmt.Println("Input: " + angka)
	fmt.Printf("Output: ")
	return result
}

func main() {
	fmt.Println(munculSekali("1234123"))    // [4]
	fmt.Println(munculSekali("76523752"))   // [6, 3]
	fmt.Println(munculSekali("12345"))      // [1 2 3 4 5]
	fmt.Println(munculSekali("1122334455")) // []
	fmt.Println(munculSekali("0872504"))    // [8 7 2 5 4]
}
