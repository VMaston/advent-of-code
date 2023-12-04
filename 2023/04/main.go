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

type cardType struct {
	instances int
	winners map[string][]int
	numbers []string
}

func createScratchcard(line string) cardType{
	card := strings.Split(string(line), "|")
	winningNums := strings.Fields(strings.Split(card[0], ":")[1])
	winners := make(map[string][]int)
	playerNums := strings.Fields(card[1])
	for _, num := range winningNums {
		number, _ := strconv.Atoi(num)
		winners[num] = append(winners[num], number)
	}

	return cardType{
		instances: 1,
		winners: winners,
		numbers: playerNums,
	}
}

func main() {
	input, err := os.ReadFile("input.txt")
	if err!= nil {
		fmt.Println("There has been an error reading the input file.")
	}
	lines := strings.Split(string(input), "\n")
	
	cards := make(map[int][]cardType)
	var totalPoints int
	var iterationTotal int

	//Organize each scratchcard into Map with value of cardType struct.
	for i, v := range lines {
		cards[i] = append(cards[i], createScratchcard(v))
	}

	for i := range lines { //We iterate on the lines, just for the index - to get the relevant map key. 
		//Grab properties of the current scratchcard
		playerNums := cards[i][0].numbers
		winners := cards[i][0].winners
		instances := cards[i][0].instances

		var cardPoints int
		var winnings int
		
		//Part 1
		for _, num := range playerNums {
			if winners[num] != nil { //Winners is a key:value map for easy lookup.
				winnings++
				if cardPoints == 0 {
					cardPoints = 1
					} else {
					cardPoints = cardPoints * 2
				}
			}
		}
		totalPoints += cardPoints
		
		//Part 2
		for y := i; y < i+winnings; y++ { //Iterate through the scratchcard collection for the amount you won.
			cards[y+1][0].instances += instances //We add instances to the next card based on how many instances we have on our active card. Instead of iterating +1, we just add on the total.
		}
		iterationTotal += cards[i][0].instances
	}

	fmt.Println("All Scratchcard Points:", totalPoints)
	fmt.Println("Iteration Total:", iterationTotal)
}