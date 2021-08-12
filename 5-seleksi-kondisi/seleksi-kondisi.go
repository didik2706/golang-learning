package main

import "fmt"

func main() {
	nilai := 10

	// ini adalah if statement
	if nilai == 10 {
		fmt.Println("lulus dengan nilai sempurna")
	} else if nilai > 5 {
		fmt.Println("lulus")
	} else if nilai == 4 {
		fmt.Println("hampir lulus")
	} else {
		fmt.Printf("tidak lulus, nilai anda %d", nilai)
	}

	// variabel temporary
	point := 8000

	if percent := point / 100; percent >= 100 {
		fmt.Printf("%.1f%s perfect!\n", percent, "%")
	} else if percent >= 70 {
		fmt.Printf("%.1f%s good \n", percent, "%")
	} else {
		fmt.Printf("%.1f%s not bad \n", percent, "%")
	}

	// switch case
	point = 2
	switch point {
	case 8:
		fmt.Println("mantab")
	case 10:
		fmt.Println("mantab banget")
	default:
		fmt.Println("sugoi")
	}

	switch point {
	case 2:
		fmt.Println("2 dan 4 adalah ")
	default:
		{
			fmt.Println("Selain 2 dan 4")
			fmt.Println("angka lainnya adalah ganjil")
		}
	}

	// swith case style if-else
	switch {
	case point%2 == 0:
		{
			fmt.Println("Result :")
			fmt.Println("point bernilai genap")
		}
	default:
		{
			fmt.Println("Result :")
			fmt.Println("point bernilai ganjil")
		}
	}

	// penggunaan fallthrough
	switch {
	case point == 2:
		fmt.Println("point bernilai 2")
		fallthrough
	case point >= 1 && point <= 2:
		fmt.Println("lanjut bossqu")
	default:
		fmt.Println("Oghey")
	}
}
