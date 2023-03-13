package main

import "fmt"

func ArrayMerge(arrayA, arrayB []string) []string {
	// menggabungkan 2 array
	arr := append(arrayA, arrayB...)

	// buat map untuk menyimpan nilai unik dari array
	index := make(map[string]bool)
	for _, item := range arr {
		index[item] = true
	}
	// konversi map menjadi array
	result := make([]string, 0, len(index))
	for item := range index {
		result = append(result, item)
	}

	return result
}

func main() {
	// Test cases
	fmt.Println(ArrayMerge([]string{"king", "devil jin", "akuma"}, []string{"eddie", "steve", " geese"}))
	//  [ "king" , " devil jin" , "akuma" , "eddie" , "steve'' , "geese" ]
	fmt.Println(ArrayMerge([]string{"sergei", "jin"}, []string{"jin", "steve", "bryan"}))
	// // [ "sergei", " jin" , "steve" , "bryan"]
	fmt.Println(ArrayMerge([]string{"alisa", "yoshimitsu"}, []string{"devil jin", "yoshimitsu", "alisa", "law"}))
	// // ["alisa", "yoshimitsu", "devil jin", "1aw"]
	fmt.Println(ArrayMerge([]string{}, []string{"devil jin", "sergei"}))
	// //  ["devil jin", "sergei"]
	fmt.Println(ArrayMerge([]string{"hwoarang"}, []string{}))
	// //  [ "hwoarang" ]
	fmt.Println(ArrayMerge([]string{}, []string{}))
	//  []
}
