package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var totalProcessedCards int

func main() {
	// Open our .txt file containing our input
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Create a new scanner to read each line in our input file
	scanner := bufio.NewScanner(file)

	// Create a map to keep track of each card
	cardsNums := make(map[int]int, 188)

	// Keep track of our cards that are yet to be processed
	var cards []int

	// Read line by line
	for scanner.Scan() {
		input := strings.SplitAfter(scanner.Text(), `:`)

		// Get the card number:
		card_input_end := strings.Index(input[0], `:`)
		card_input := input[0][5:card_input_end]
		card_num := strings.TrimSpace(card_input)
		card_num_int, err := strconv.Atoi(card_num)
		if err != nil {
			log.Fatal(err)
		}
		nums := input[1]

		matching_nums := matchingNums(nums)
		cardsNums[card_num_int] = matching_nums // key: card ID value: # of matches

		// Increment the amount of cards we have processed
		totalProcessedCards++

		// Add the matching cards to our list of cards now
		for i := 0; i < matching_nums; i++ {
			card_num_int++
			cards = append(cards, card_num_int)
		}
		// fmt.Println(cards)
	}

	// Now we should have a dictionary of {cardNum : matches}... We need to process those

	for len(cards) > 0 {
		// Get the first card from the list -> Process it
		cardNum := cards[0]
		matches := cardsNums[cardNum]

		// Loop through and add to cards...
		for i := 0; i < matches; i++ {
			cardNum++
			cards = append(cards, cardNum)
		}

		// Delete the first element of our slice (the current card we are processing)
		cards = cards[1:]

		// Increment amount of processed cards
		totalProcessedCards++
	}
	// Print our result
	fmt.Println("Total Processed Cards:", totalProcessedCards)
}

// Returns an int indicating the amount of matching numbers
func matchingNums(str string) int {
	nums_split := strings.SplitAfter(str, `|`)

	winning_nums := nums_split[0]
	my_nums := nums_split[1]

	winning_nums_slice := strings.Fields(winning_nums)
	my_nums_slice := strings.Fields(my_nums)

	var matching_nums int

	for _, my_num := range my_nums_slice {
		for _, winning_num := range winning_nums_slice {
			if my_num == winning_num {
				matching_nums++
			}
		}
	}
	return matching_nums
}
