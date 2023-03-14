package main

import "fmt"

type Car struct {
	Type   string
	FuelIn float64 // jumlah bahan bakar dalam liter
}

func (car *Car) EstimateDistance() float64 {
	// asumsikan konsumsi bahan bakar 1.5 liter per mill
	// jarak dapat dihitung dengan mengalikan jumlah bahan bakar dengan konsumsi
	distance := car.FuelIn / 1.5
	return distance
}

func main() {
	// contoh penggunaan
	car := Car{Type: "Sedan", FuelIn: 3} // mobil sedan dengan ... liter bahan bakar
	distance := car.EstimateDistance()
	fmt.Printf("Mobil %s dengan bahan bakar %v liter dapat menempuh jarak perkiraan %v mill", car.Type, car.FuelIn, distance)
}
