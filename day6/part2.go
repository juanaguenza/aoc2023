package main

import "fmt"

// Whenever we hold down the button for a second longer at the beginning of the race, our boat's speed increases by 1 millisecond
const speed_increase = 1

func main() {
	// Given that our input file is so small, im not going to bother reading from it...

	time_distance := map[int]int{63789468: 411127420471035}

	total := calculate_winners(time_distance)

	fmt.Print("Total:", total)
}

func calculate_winners(t_d map[int]int) int {
	for time, distance_record := range t_d {
		var race_total int
		for i := 0; i < time; i++ {
			speed := i * speed_increase
			remaining_time := time - i
			distance_travelled := speed * remaining_time

			if distance_travelled > distance_record {
				race_total += 1
			}
		}
		return race_total
	}
	return -1
}
