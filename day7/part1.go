package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	var result int

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Create a new scanner to read each line in our input file
	scanner := bufio.NewScanner(file)

	cards_bids := make(map[string]int)
	cards_type := make(map[string]int)
	// Read each line
	for scanner.Scan() {
		input := strings.Split(scanner.Text(), " ")
		// input [0] cards, input [1] bid
		input_bid, err := strconv.Atoi(input[1])
		if err != nil {
			log.Fatal(err)
		}
		cards_bids[input[0]] = input_bid
		cards_type[input[0]] = handType(input[0])
	}

	var five_kind []string
	var four_kind []string
	var full_house []string
	var three_kind []string
	var two_pair []string
	var one_pair []string
	var high_card []string

	for k, v := range cards_type {
		switch v {
		case 1:
			five_kind = append(five_kind, k)
		case 2:
			four_kind = append(four_kind, k)
		case 3:
			full_house = append(full_house, k)
		case 4:
			three_kind = append(three_kind, k)
		case 5:
			two_pair = append(two_pair, k)
		case 6:
			one_pair = append(one_pair, k)
		case 7:
			high_card = append(high_card, k)
		}
	}

	slices.Sort(five_kind)
	slices.Sort(four_kind)
	slices.Sort(full_house)
	slices.Sort(three_kind)
	slices.Sort(two_pair)
	slices.Sort(one_pair)
	slices.Sort(high_card)

	fmt.Println(five_kind)
	fmt.Println(four_kind)

	rank := 1
	for i := range high_card {
		i = len(high_card) - 1 - i
		// fmt.Printf("%d * %d\n", cards_bids[high_card[i]], rank)
		result += cards_bids[high_card[i]] * rank
		rank++
	}

	for i := range one_pair {
		i = len(one_pair) - 1 - i
		// fmt.Printf("%d * %d\n", cards_bids[one_pair[i]], rank)
		result += cards_bids[one_pair[i]] * rank
		rank++
	}

	for i := range two_pair {
		i = len(two_pair) - 1 - i
		// fmt.Printf("%d * %d\n", cards_bids[two_pair[i]], rank)
		result += cards_bids[two_pair[i]] * rank
		rank++
	}

	for i := range three_kind {
		i = len(three_kind) - 1 - i
		// fmt.Printf("%d * %d\n", cards_bids[three_kind[i]], rank)
		result += cards_bids[three_kind[i]] * rank
		rank++
	}

	for i := range full_house {
		i = len(full_house) - 1 - i
		// fmt.Printf("%d * %d", cards_bids[full_house[i]], rank)
		result += cards_bids[full_house[i]] * rank
		rank++
	}

	for i := range four_kind {
		i = len(four_kind) - 1 - i
		// fmt.Printf("%d * %d\n", cards_bids[four_kind[i]], rank)
		result += cards_bids[four_kind[i]] * rank
		rank++
	}

	for i := range five_kind {
		i = len(five_kind) - 1 - i
		// fmt.Printf("%d * %d\n", cards_bids[five_kind[i]], rank)
		result += cards_bids[five_kind[i]] * rank
		rank++
	}

	fmt.Printf("result is: %d", result)
}

// 5 is 53... f is 102 h is 104
func handType(hand string) int {
	// five of a kind
	if hand[0] == hand[1] && hand[1] == hand[2] && hand[2] == hand[3] && hand[3] == hand[4] {
		return 1
	}

	// four of a kind
	for i := 0; i < len(hand); i++ {
		matches := 1
		for j := i + 1; j != i; {
			if j == len(hand) {
				j = 0
			}
			if hand[i] == hand[j] {
				matches++
			}
			j++
			if j == len(hand) {
				j = 0
			}
		}
		if matches == 4 {
			return 2
		}
	}

	// full house
	labels := make(map[rune]int)
	for _, card := range hand {
		labels[card] = 0
	}
	if len(labels) == 2 {
		return 3
	}

	// three of a kind
	for i := 0; i < len(hand); i++ {
		matches := 1
		for j := i + 1; j != i; {
			if j == len(hand) {
				j = 0
			}
			if hand[i] == hand[j] {
				matches++
			}
			j++
			if j == len(hand) {
				j = 0
			}
		}
		if matches == 3 {
			return 4
		}
	}

	// two pair -- reuse labels from full house
	if len(labels) == 3 {
		return 5
	}

	// one pair -- reuse labels from full house
	if len(labels) == 4 {
		return 6
	}

	// high card
	return 7
}

// Take in two hands, return the hand with the higher rank
func higherRank(hand1 string, hand2 string) string {
	return "a"
}
