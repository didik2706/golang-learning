package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	names := []string{"Didik", "Hafidz", "Nurul"}

	// memanggil fungsi printMessage
	printMessage("Hallo", names)

	// memanggi fungsi randomWithRange
	fmt.Println(randomWithRange(3, 10))
	fmt.Println(randomWithRange(1, 10))
	fmt.Println(randomWithRange(1, 10))

	rand.Seed(time.Now().Unix())
}

func printMessage(msg string, arr []string) {
	for _, name := range arr {
		fmt.Println(msg, name)
	}
}

func randomWithRange(min, max int) int {
	value := rand.Int()%(max-min+1) + min
	return value
}
