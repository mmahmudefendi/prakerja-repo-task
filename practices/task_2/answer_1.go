package main

import "fmt"

//  menentukan bilang prima
func isPrime(n int) bool {
	if n <= 1 {
		return false
	}

	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func main() {
	for i := 1; i <= 100; i++ {
		if isPrime(i) {
			fmt.Printf("%d ", i)
		}
	}
}
