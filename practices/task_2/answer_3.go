package main

import "fmt"

// program untuk menghitung luas trapesium
func main() {
	var a, b, t float64

	fmt.Print("Masukkan panjang sisi atas: ")
	fmt.Scanln(&a)

	fmt.Print("Masukkan panjang sisi bawah: ")
	fmt.Scanln(&b)

	fmt.Print("Masukkan tinggi: ")
	fmt.Scanln(&t)

	luas := (a + b) * t / 2

	fmt.Printf("Luas trapesium: %.2f", luas)
}
