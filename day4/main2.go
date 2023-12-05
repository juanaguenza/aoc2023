package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
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

	// Keep track of all our cards
	var cards []int

	// Create a map to keep track of each card
	cardsNums := make(map[int]string, 188)

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
		cardsNums[card_num_int] = nums

		matching_nums := matchingNums(nums)

		if err != nil {
			log.Fatal(err)
		}
		for i := 0; i < matching_nums; i++ {
			card_num_int++
			cards = append(cards, card_num_int)
		}
	}
	// For our remaining cards in cards
	winner := true
	var resulting_cards []int
	for winner || len(cards) > 0 {
		for i, card := range cards {
			nums, ok := cardsNums[card]
			if ok {
				matching_nums := matchingNums(nums)
				winner = true
				// Remove that card from our current stash
				cards = append(cards[:i], cards[i+1:]...)

			} else {
				winner = false
			}

		}
	}
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
