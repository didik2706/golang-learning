package main

import "fmt"

func main() {
	// perulangan pertama
	for i := 1; i <= 5; i++ {
		fmt.Println("Perulangan pertama :", i)
	}

	// perulangan kedua
	i := 1
	for i <= 5 {
		fmt.Println("Perulangan kedua :", i)
		i++
	}

	// perulangan ketiga
	i = 1
	for {
		fmt.Println("Perulangan ketiga :", i)

		if i == 5 {
			break
		} else {
			i++
		}
	}

	// continue dan break
	for i := 1; i <= 10; i++ {
		if i%2 == 1 {
			continue
		}

		if i > 8 {
			break
		}

		fmt.Println("Angka :", i)
	}

	// perulangan bersarang
	for i := 0; i < 5; i++ {
		for j := i; j < 5; j++ {
			fmt.Print(j, " ")
		}
		fmt.Println()
	}

outerLoop:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i == 3 {
				break outerLoop
			}
			fmt.Print("matriks [", i, "][", j, "]", "\n")
		}
	}
}
