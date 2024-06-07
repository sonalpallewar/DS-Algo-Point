package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

/*
Sorts an array of integers using the bubble sort algorithm.
*/
func sort(array []int) []int {
	arrayLen := len(array)

	for i := 0; i < arrayLen-1; i++ {
		for j := 0; j < arrayLen-i-1; j++ {
			if array[j] > array[j+1] {
				// Swap the values without using a temporary variable
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}

	return array
}

func main() {
	var array []int

	// Iterate over command-line arguments starting from the first one (ignoring the program path)
	for _, value := range os.Args[1:] {
		// Convert the parameter from string to integer
		number, err := strconv.Atoi(value)
		if err != nil {
			log.Fatal(err)
		}

		// Add number to array
		array = append(array, number)
	}

	fmt.Println("Unsorted array =>", array)
	fmt.Println("Sorted array =>", sort(array))
}

/*
Example usage:
Input: -2 45 0 11 -9 3 92 88234 8 -14234 234 215 52
Output: -14234 -9 -2 0 3 8 11 45 52 92 215 234 88234

Time complexity: O(n^2)
Space complexity: O(1)
*/
