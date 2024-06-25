package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)


func main() {
	var array []int
	var elementToFind int

	// Converts all input arguments to integers and adds them to the list to be sorted
	for _, value := range os.Args[1:] {
		number, err := strconv.Atoi(value)

		if err != nil {
			log.Fatal(err)
		}

		array = append(array, number)
	}

	// Sort array of ints in increasing order
	sort.Ints(array)

	// Asks the user to enter the number they want to find within the array
	fmt.Print("Enter the number to find: ")
	fmt.Scanf("%d", &elementToFind)

	fmt.Println(fmt.Sprintf("Element %d was found at position %d", elementToFind, search(array, elementToFind)))
}

/*
I => Array: 12 19 20 33 44 68 79 80
	 Element to find: 68
O => 5

Time complexity: O(log n)
Space complexity: O(1)
*/
