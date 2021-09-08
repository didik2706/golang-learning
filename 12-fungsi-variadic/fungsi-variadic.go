package main

import "fmt"

func main()  {
	total := total(1,2,3,22,3,2)
	fmt.Println("Totalnya adalah ", total)
}

func total(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}

	return total
}