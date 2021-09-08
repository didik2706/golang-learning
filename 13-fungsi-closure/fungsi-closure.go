package main

import "fmt"

func main() {
	// getMinMax := func(n []int) (int, int) {
	// 	var min, max int

	// 	for i, num := range n {
	// 		switch {
	// 		case i == 0:
	// 			min, max = num, num
	// 		case num > max:
	// 			max = num
	// 		case num < min:
	// 			min = num
	// 		}
	// 	}

	// 	return min, max
	// }

	// numbers := []int{4,5,33,5,6,8,10}
	// min, max := getMinMax(numbers)
	// fmt.Printf("Min Value : %v \nMax Value : %v", min, max)

	// findMax

    var max = 3
    var numbers = []int{2, 3, 0, 4, 3, 2, 0, 4, 2, 0, 3}
    var howMany, getNumbers = findMax(numbers, max)
    var theNumbers = getNumbers()

    fmt.Println("numbers\t:", numbers)
    fmt.Printf("find \t: %d\n\n", max)

    fmt.Println("found \t:", howMany)    // 9
    fmt.Println("value \t:", theNumbers) // [2 3 0 3 2 0 2 0 3]
}

func findMax(numbers []int, max int) (int, func() []int) {
	var res []int
	for _, e := range numbers {
		if e <= max {
			res = append(res, e)
		}
	}

	return len(res), func() []int {
		return res
	}
}