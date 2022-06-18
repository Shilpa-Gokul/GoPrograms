package main

import (
	"fmt"
	"strconv"
)

// GenerateLargestNumber generates the largest number based on the elements in the array
func GenerateLargestNumber(input []int) []int {
	if len(input) == 0 {
		fmt.Println("Input is empty")
		return input
	}
	if len(input) == 1 {
		return input
	}
	for i := range input {
		for j := i + 1; j < len(input); j++ {
			// convert int to string and concatenate
			val1_val2 := strconv.Itoa(input[i]) + strconv.Itoa(input[j])
			val2_val1 := strconv.Itoa(input[j]) + strconv.Itoa(input[i])
			// check if the concatenated value val2_val1 is greater and swap elements accordingly
			if val2_val1 > val1_val2 {
				input[i], input[j] = input[j], input[i]
			}
		}
	}
	return input
}
func main() {
	input := []int{9, 93, 24, 6}
	fmt.Println(GenerateLargestNumber(input))
}
