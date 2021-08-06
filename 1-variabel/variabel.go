package main

import "fmt"

func main() {
	// ada dua cara mendeklarasikan variabel
	// cara pertama
	// var firstName string = "Didik"

	// // cara kedua
	// var lastName string
	// lastName = " Nur Hidayat"

	// cara deklarasi variabel tanpa tipe data
	// type inferece
	firstName := "Didk"
	firstName = "Didik"
	var lastName = "Nur Hidayat"

	// deklarasi multi variabel
	// ada 3 cara
	// cara pertama
	var satu, dua, tiga string = "satu", "dua", "tiga"
	// cara kedua
	var empat, lima, enam string
	empat, lima, enam = "empat", "lima", "enam"
	// cara ketiga
	tujuh, delapan, isMale := "tujuh", "delapan", true

	// dekalrasi variabel menggunakan keyword new
	name := new(string)

	// mencetak nilai dari variabel
	fmt.Println(*name)
	fmt.Println(satu, dua, tiga, empat, lima, enam, tujuh, delapan, isMale)
	fmt.Printf("Hello, %s %s! \n", firstName, lastName)
}
