package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// var sample string =
// `Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
// Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
// Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
// Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
// Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
// Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11`

func main() {
	input, err := os.ReadFile("input.txt")
	if err!= nil {
		fmt.Println("There has been an error reading the input file.")
	}
	lines := strings.Split(string(input), "\n")
	
	var totalPoints int

	for i, v := range lines {
		fmt.Println("Card ", i+1)
		var cardPoints int
		winners := make(map[string][]int)
		card := strings.Split(string(v), "|")
		winningNums := strings.Fields(strings.Split(card[0], ":")[1])
		playerNums := strings.Fields(card[1])
		for _, num := range winningNums {
			number, _ := strconv.Atoi(num)
			winners[num] = append(winners[num], number)
		}

		for _, num := range playerNums {
			if winners[num] != nil {
				fmt.Println("Winner!", num, "and", winners[num])
				if cardPoints == 0 {
					cardPoints = 1
				} else {
					cardPoints = cardPoints * 2
				}
				fmt.Println("Total Card Points:", cardPoints)
			}
		}
		totalPoints += cardPoints

	}

	fmt.Println("All Scratchcard Points:", totalPoints)
}