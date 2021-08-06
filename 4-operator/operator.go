package main

import "fmt"

func main() {
	// operator aritmatika
	var value = (5 + 5) / 2

	fmt.Println(value % 3)
	fmt.Println(value)

	// operator perbandingan
	fmt.Println("Operator Perbandingan :", value%2 == 0)

	// operator logika
	// 1. AND
	fmt.Println("Operator Logika AND :", value%2 == 1 && value%3 == 2)

	// 2. OR
	fmt.Println("Operator OR :", value%2 == 0 || value%3 == 2)

	// 3. NOT
	fmt.Println("Operator NOT :", !(value%2 == 0))

}
