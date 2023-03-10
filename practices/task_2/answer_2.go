package main

import "fmt"

// 2 menentukan bilangan kelipatan 7
func main() {
	for i := 1; i <= 100; i++ {
		if i%7 == 0 {
			fmt.Printf("%d ", i)
		}
	}
}
