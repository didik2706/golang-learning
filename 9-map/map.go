package main

import "fmt"

func main() {
	// cara membuat map
	var foods map[string]int
	foods = map[string]int{}

	foods["chicken"] = 5000
	foods["burger"] = 15000

	fmt.Println("Price burger :", foods["burger"])
	fmt.Println("Price chicken :", foods["chicken"])

	// cara menggunakan map di golang
	data := map[string]int{}

	data["satu"] = 1
	fmt.Println(data["satu"])

	// inisialisasi nilai map di golang
	fruits := map[int]string{
		1: "Apel",
		2: "Mangga",
		3: "Manggis",
	}

	fmt.Println(fruits[1])

	// cara menampilkan seluruh nilai dari map
	date := map[string]int{
		"januari":  31,
		"februari": 28,
		"maret":    30,
		"april":    31,
		"mei":      30,
		"juni":     31,
	}

	for key, val := range date {
		fmt.Println(key, ":", val)
	}

	// cara menghapus key dan value di map
	delete(date, "januari")
	fmt.Println(date)

	// cara mendeteksi item di map
	value, isExist := date["maret"]

	if isExist {
		fmt.Println(value)
	} else {
		fmt.Println("item not found")
	}

	students := []map[string]string{
		map[string]string{"name": "Didik Nur Hidayat", "prodi": "Informatika"},
		map[string]string{"name": "M. Hafidurrahman", "prodi": "Informatika"},
	}

	fmt.Println(students[0]["name"])
}
