package main

import (
	"fmt"
	"math"
)

func main() {
	luas, keliling := calculate(5)

	fmt.Println(luas)
	fmt.Println(keliling)
}

func calculate(d float64) (float64, float64) {
	// hitung luas
	area := math.Pi * math.Pow(d/2, 2)
	// hitung keliling
	keliling := math.Pi * d

	return area, keliling
}
