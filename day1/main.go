package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"slices"
	"strconv"
)

// part a
// func main() {
// 	// Open our txt file containing our input
// 	file, err := os.Open("input.txt")

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// Create a new scanner to read each line in our input file
// 	scanner := bufio.NewScanner(file)

// 	var total int

// 	// Read line by line
// 	for scanner.Scan() {
// 		var arr []string
// 		// If we encounter a char that can be converted to a string
// 		// Do so and append it to a list
// 		for i := 0; i < len(scanner.Text()); i++ {
// 			line := scanner.Text()
// 			_, err := strconv.Atoi(string(line[i]))
// 			if err == nil {
// 				arr = append(arr, string(line[i]))
// 			}
// 		}
// 		// The first and the last element of the list are the two numbers to add
// 		// In the case there was only 1 number, the same number will get added to each other
// 		total_str := arr[0] + arr[len(arr)-1]
// 		// Convert it to an Int (kept as strings before so it was easier to do 8+8 -> 88)
// 		currTotal, _ := strconv.Atoi(total_str)
// 		total += currTotal
// 	}

// 	file.Close()

// 	fmt.Print(total)
// }

// part b

func main() {

	// Create a dictionary to make our map each Int (1-9) to its corresponding string
	numStr := make(map[int]string)
	numStr[1] = `one`
	numStr[2] = `two`
	numStr[3] = `three`
	numStr[4] = `four`
	numStr[5] = `five`
	numStr[6] = `six`
	numStr[7] = `seven`
	numStr[8] = `eight`
	numStr[9] = `nine`

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
		var indexes []int
		var regIndexes [][][]int
		var nums []string
		var str string

		// If we encounter a char that can be converted to a string,
 		// do so and append it to a list
		// Additionally, keep track of what indexes those chars are at,
		// add it to a separate list
		// These lists are stored parallelly
		for i := 0; i < len(scanner.Text()); i++ {
			line := scanner.Text()
			_, err := strconv.Atoi(string(line[i]))
			if err == nil {
				indexes = append(indexes, i)
				nums = append(nums, string(line[i]))
			}
			str = string(line)
		}

		// Use regex to get the indexes for each num (one, two, three, ..., nine)
		for i := 1; i < 10; i++ {
			re := regexp.MustCompile(numStr[i])
			index := re.FindAllStringIndex(str, -1)
			if index != nil {
				for j := 0; j < len(index); j++ {
					index[j] = append(index[j], i)
				}
				// For each match, append the num that was found to the end of it
				regIndexes = append(regIndexes, index)
			}
		}

		// Add the beginning indexes for each word num to our indexes list
		for i := 0; i < len(regIndexes); i++ {
			for j := 0; j < len(regIndexes[i]); j++ {
				indexes = append(indexes, regIndexes[i][j][0])
			}
		}

		// Sort the indexes list
		slices.Sort(indexes)

		var firstNum int
		var secondNum int

		setFirstNum := false
		setSecondNum := false

		// Check if the first or last sorted index is equivalent to a beginning index of a word num
		// If so, we want to set our first and last num accordingly
		for i := 0; i < len(regIndexes); i++ {
			for j := 0; j < len(regIndexes[i]); j++ {
				if indexes[0] == regIndexes[i][j][0] {
					firstNum = regIndexes[i][j][2]
					setFirstNum = true
				}
				if indexes[len(indexes)-1] == regIndexes[i][j][0] {
					secondNum = regIndexes[i][j][2]
					setSecondNum = true
				}
			}
		}

		// Store the number of our current line total as a string
		var lineNum string

		if setFirstNum && setSecondNum {
			lineNum = strconv.Itoa(firstNum) + strconv.Itoa(secondNum)
		} else if setFirstNum {
			lineNum = strconv.Itoa(firstNum) + nums[len(nums)-1]
		} else if setSecondNum {
			lineNum = nums[0] + strconv.Itoa(secondNum)
		} else {
			lineNum = nums[0] + nums[len(nums)-1]
		}

		lineTotal, err := strconv.Atoi(lineNum)

		if err == nil {
			total += lineTotal
		}

	}

	file.Close()

	fmt.Print(total)
}
