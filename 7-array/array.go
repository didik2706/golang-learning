package main

import "fmt"

func main() {
	// cara pertama inisialisasi array do golang
	var names [4]string

	names[0] = "Didik"
	names[1] = "Hafidz"
	names[2] = "Nurul"
	names[3] = "Dzun"

	fmt.Println(names[0], "dan", names[2])

	// inisialisasi array secara horizontal
	var fruits = [4]string{"Pisang", "Apel", "Mangga", "Pir"}
	// menampilkan output array
	fmt.Println(len(fruits))
	fmt.Println(fruits)

	// inisialisasi array secara vertikal
	fruits = [4]string{
		"Strowberi",
		"Ceri",
		"blueberry",
		"Kiwi",
	}

	fmt.Println(fruits)

	// inisialisasi nilai awal array tanpa menentukan jumkah arraynya
	numbers := [...]int{1, 2, 3, 4, 5, 6}
	fmt.Println("Data Array :", numbers)
	fmt.Println("Jumlah Array :", len(numbers))

	// array multidemensi
	// cara pertama
	numbers2 := [2][3]int{[3]int{1, 2, 3}, [3]int{4, 5, 6}}
	// cara kedua
	numbers3 := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(numbers2)
	fmt.Println(numbers3)

	// menggunakan perulangan untuk mendapatkan elemen
	for i := 0; i < len(fruits); i++ {
		fmt.Println("elemen :", fruits[i])
	}

	// menggunakan perulangan for - range untuk mendapatkan elemen
	for i, fruit := range fruits {
		fmt.Println("Elemen :", fruit, "pada index ke-", i)
	}
}
