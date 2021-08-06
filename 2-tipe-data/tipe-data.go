package main

import "fmt"

func main() {

	var positiveNumber uint8 = 80
	var negativeNumber = -1243423644

	fmt.Println("bilangan positif :", positiveNumber)
	fmt.Println("bilangan negatif :", negativeNumber)

	var decimalNumber = 2.63
	fmt.Printf("bilangan desimal : %f\n", decimalNumber)
	fmt.Printf("bilangan desimal : %.3f\n", decimalNumber)

	var exist bool = true
	fmt.Printf("exist? %t \n", exist)
}
