package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// Part A
func main() {
	// Open our txt file containing our input
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Create a new scanner to read each line in our input file
	scanner := bufio.NewScanner(file)

	var total int

	// Read line by line
	for scanner.Scan() {
		var arr []string
		// If we encounter a char that can be converted to a string
		// Do so and append it to a list
		for i := 0; i < len(scanner.Text()); i++ {
			line := scanner.Text()
			_, err := strconv.Atoi(string(line[i]))
			if err == nil {
				arr = append(arr, string(line[i]))
			}
		}
		// The first and the last element of the list are the two numbers to add
		// In the case there was only 1 number, the same number will get added to each other
		total_str := arr[0] + arr[len(arr)-1]
		// Convert it to an Int (kept as strings before so it was easier to do 8+8 -> 88)
		currTotal, _ := strconv.Atoi(total_str)
		total += currTotal
	}

	file.Close()

	fmt.Print(total)
}
