package main

import "fmt"

// menerima pointer dari array numbers yang berisi 6 angka inputan, dan mengembalikan 2 nilai berupa nilai maksimum dan minimum
// referencing
// numbers alamat memory
func findMinMax(numbers *[6]int) (int, int) {
	// Variabel max dan min diinisialisasi dengan nilai pertama pada array numbers
	var max, min int
	// dereferencing *numbers
	// agar nilai aktual dari array yang diwakili oleh pointer numbers dapat diakses.
	for i, n := range *numbers {
		if i == 0 {
			max, min = n, n
		} else {
			if n > max {
				max = n
			}
			if n < min {
				min = n
			}
		}
	}
	return max, min
}

func main() {
	var numbers [6]int
	for i := 0; i < 6; i++ {
		fmt.Printf("Masukkan angka ke-%d: ", i+1)
		// alamat memori dari elemen ke-i dari array numbers
		fmt.Scan(&numbers[i])
	}

	max, min := findMinMax(&numbers)
	fmt.Printf("Nilai maksimum: %d\n", max)
	fmt.Printf("Nilai minimum: %d\n", min)
}
