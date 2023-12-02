package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	// Open our txt file containing our input
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	// Create a new scanner to read each line in our input file
	scanner := bufio.NewScanner(file)


	gameIdIndex, _ := regexp.Compile(`:`)

	gameSetEnd, _ := regexp.Compile(`;`)

	findNums := regexp.MustCompile("[0-9]+")

	var total int

	// Read line by line
	for scanner.Scan() {
		lineContent := scanner.Text()
		var gameInfo []int
		// var games []string
		// parse the input:
		// Game #
		sepIndex := gameIdIndex.FindStringIndex(lineContent)
		gameIdStr := lineContent[5:sepIndex[0]]
		// Remove the Game ID from the string
		lineContent = lineContent[sepIndex[1] + 1:]

		gameId, _ := strconv.Atoi(gameIdStr)
		gameInfo = append(gameInfo, gameId)

		// Then each game.

		var redCount int
		var greenCount int
		var blueCount int

		possibleGame := true

		for gameSepIndex := gameSetEnd.FindStringIndex(lineContent); gameSepIndex != nil; gameSepIndex = gameSetEnd.FindStringIndex(lineContent) {

			gameContent := lineContent[:gameSepIndex[0]]
			lineContent = lineContent[gameSepIndex[1] + 1:]

			colorsCount := colorCount(gameContent, findNums)

			redCount = colorsCount[0]
			greenCount = colorsCount[1]
			blueCount = colorsCount[2]

			if redCount > 12 {
				possibleGame = false
				break
			}
			if greenCount > 13 {
				possibleGame = false
				break
			}
			if blueCount > 14 {
				possibleGame = false
				break
			}
		}

		colorsCount := colorCount(lineContent, findNums)

		redCount = colorsCount[0]
		greenCount = colorsCount[1]
		blueCount = colorsCount[2]

		if redCount > 12 {
			possibleGame = false
		}
		if greenCount > 13 {
			possibleGame = false
		}
		if blueCount > 14 {
			possibleGame = false
		}

		if possibleGame {
			total += gameId
		}


	}
	fmt.Println("sum of all possible game id's:", total)
}


func colorCount (str string, regx *regexp.Regexp) []int {

	findRed, _ := regexp.Compile(`red`)
	findGreen, _ := regexp.Compile(`green`)
	findBlue, _ := regexp.Compile(`blue`)

	var result []int

	redCount := 0
	greenCount := 0
	blueCount := 0

	for colorCountIndex := regx.FindStringIndex(str); colorCountIndex != nil; colorCountIndex = regx.FindStringIndex(str) {
		countStr := str[colorCountIndex[0] : colorCountIndex[1]]
		count, _ := strconv.Atoi(countStr)

		redIndex := findRed.FindStringIndex(str)
		greenIndex := findGreen.FindStringIndex(str)
		blueIndex := findBlue.FindStringIndex(str)

		if redIndex != nil && colorCountIndex[1] + 1 == redIndex[0] {
			redCount = count
			if redIndex[1] != len(str) {
				str = str[redIndex[1] + 1 : ]
			} else {
				str = ""
			}
		} else if greenIndex != nil && colorCountIndex[1] + 1 == greenIndex[0] {
			greenCount = count
			if greenIndex[1] != len(str) {
				str = str[greenIndex[1] + 1 : ]
			} else {
				str = ""
			}
		} else if blueIndex != nil && colorCountIndex[1] + 1 == blueIndex[0] {
			blueCount = count
			if blueIndex[1] != len(str) {
				str = str[blueIndex[1] + 1 : ]
			} else {
				str = ""
			}
		}
	}

	result = append(result, redCount)
	result = append(result, greenCount)
	result = append(result, blueCount)

	return result
}