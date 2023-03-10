package main

import (
	"fmt"
)

// menentukan bilangan prima
func isPrime(num int) bool {
	if num <= 1 {
		return false
	}
	for i := 2; i <= num/2; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	var num int
	fmt.Print("Masukkan sebuah bilangan: ")
	fmt.Scanln(&num)

	if isPrime(num) {
		fmt.Println(num, "adalah bilangan prima.")
	} else {
		fmt.Println(num, "bukan bilangan prima.")
	}
}
