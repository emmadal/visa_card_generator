package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

func main() {
	var numbers = []int{4}
	var result int

	fmt.Println("Visa Credit card generator")

	for i := 0; i < 15; i++ {
		numbers = append(numbers, rand.Intn(10))
	}
	resultSlice := computeNumber(numbers)

	for _, value := range resultSlice {
		result += value
	}

	card := formatCard(numbers)

	if result%10 == 0 {
		fmt.Printf("%s is a valid VISA card\n", card)
	} else {
		fmt.Printf("%s is a invalid VISA card\n", card)
	}
}

func computeNumber(numbers []int) []int {
	card := append([]int{}, numbers...)
	for i, v := range card {
		if i%2 == 0 {
			num := v * 2
			if num >= 10 {
				num = num - 9
			}
			card[i] = num
		}
	}
	return card
}

func formatCard(numbers []int) string {
	var randomString []string
	for _, v := range numbers {
		randomString = append(randomString, strconv.Itoa(v))
	}
	result := strings.Join(randomString, "")
	return result
}
