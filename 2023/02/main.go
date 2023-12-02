package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var red, green, blue int = 12, 13, 14
	var gameCount int
	var colourCount int

	input, err := os.ReadFile("input.txt")
	if err!= nil {
		fmt.Println("There has been an error reading the input file.")
	}

	lines := strings.Split(string(input), "\n")
	for i, v := range lines {
		gameValid := true
		var redCount, greenCount, blueCount int

		game := strings.Split(v, ":")[1]
		grabs := strings.Split(game, ";")

		for _, v := range grabs {
			grab := strings.Split(v, ",")
			for _, v := range grab {
				s := strings.Fields(v)
				amount, err := strconv.Atoi(s[0])
				if err != nil {
					fmt.Printf("There has been an error converting %s - %s\n", s[0], err)
				}
				colour := s[1]

				switch colour {
				case "red":
					if amount > red {
						gameValid = false
					}
					if amount > redCount {
						redCount = amount
					}
				case "green":
					if amount > green {
						gameValid = false
					}
					if amount > greenCount {
						greenCount = amount
					}
				case "blue":
					if amount > blue {
						gameValid = false
					}
					if amount > blueCount {
						blueCount = amount
					}
				}
			}
		}

		if gameValid {
			gameCount += i+1
		}

		colourCount += redCount * greenCount * blueCount
	}
	fmt.Println(gameCount)
	fmt.Println(colourCount)
}