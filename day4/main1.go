package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	// Open our .txt file containing our input
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Create a new scanner to read each line in our input file
	scanner := bufio.NewScanner(file)

	var total int

	// Read line by line
	for scanner.Scan() {
		input := strings.SplitAfter(scanner.Text(), `:`)

		nums := input[1]
		nums_split := strings.SplitAfter(nums, `|`)

		winning_nums := nums_split[0]
		my_nums := nums_split[1]

		winning_nums_slice := strings.Fields(winning_nums)
		my_nums_slice := strings.Fields(my_nums)

		var card_total int

		for _, my_num := range my_nums_slice {
			for _, winning_num := range winning_nums_slice {
				if my_num == winning_num {
					if card_total == 0 {
						card_total = 1
					} else {
						card_total *= 2
					}
				}
			}
		}
		total += card_total
	}
	fmt.Println("Total:", total)
}
