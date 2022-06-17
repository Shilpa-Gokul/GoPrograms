package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// GenerateLargestNumber generates the largest number based on the elements in the array
func GenerateLargestNumber(input []string) []string {
	if len(input) == 0 {
		fmt.Println("Input is empty")
		return input
	}
	if len(input) == 1 {
		return input
	}
	for i := range input {
		for j := i + 1; j < len(input); j++ {
			// concatenate val1+val2 and val2+val1
			val1_val2, err := strconv.Atoi(input[i] + input[j])
			if err != nil {
				fmt.Printf("Error converting to integer %v", err)
				return []string{}
			}
			val2_val1, err := strconv.Atoi(input[j] + input[i])
			if err != nil {
				fmt.Printf("Error converting to integer %v", err)
				return []string{}
			}
			if val2_val1 > val1_val2 {
				input[i], input[j] = input[j], input[i]
			}
		}
	}
	return input
}
func main() {
	// Fetch input from user
	fmt.Println("Enter input array:")
	reader := bufio.NewReader(os.Stdin)
	readInput, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("Error while reading input %v", err)
		return
	}
	input := strings.Fields(readInput)
	fmt.Println(GenerateLargestNumber(input))
}
