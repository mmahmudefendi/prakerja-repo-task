package main

import "fmt"

// 2 menentukan bilangan kelipatan 7
func main() {
	var num int
	fmt.Print("Masukkan sebuah bilangan: ")
	fmt.Scanln(&num)

	if num%7 == 0 {
		fmt.Println(num, "adalah bilangan kelipatan 7.")
	} else {
		fmt.Println(num, "bukan bilangan kelipatan 7.")
	}
}
