package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var nums string

func calculateAdjacents(coords []int, lines []string) [8]string {
	boundaries := func(coords []int, lines []string, y, x int ) string {
		if coords[0]+y >= 0 && coords[0]+y < len(lines) && coords[1]+x >= 0 && coords[1]+x < len(lines[0]) {
			return string(lines[coords[0]+y][coords[1]+x])
		}
		return "." //Parse out of bounds co-ordinates as "." for processing.
	}

	return [8]string{ //Calculate Adjacencies
		boundaries(coords, lines, -1, -1), boundaries(coords, lines, -1, 0), boundaries(coords, lines, -1, 1), //height, row
		boundaries(coords, lines, 0, -1), boundaries(coords, lines, 0, 1),
		boundaries(coords, lines, 1, -1), boundaries(coords, lines, 1, 0), boundaries(coords, lines, 1, 1),
	}
}

func checkParts(partIndex int, lines []string, lineNumber int, currentLine string) int {
	var partHalf string
	numCoords := [][]int{}
	for p, n := range currentLine[partIndex:] {
		if strings.ContainsAny(string(n), nums) {
			partHalf += string(n)
			numCoords = append(numCoords, []int{lineNumber, partIndex+p})
		} else {
			break
		}
	}
	partIndex += len(partHalf) //Incrementing where adjacent numbers have been processed.
	for _, v := range numCoords {
		adjacents := calculateAdjacents(v, lines)

		for _, a := range adjacents {
			if !strings.ContainsAny(a, nums + ".") {
				num, _ := strconv.Atoi(partHalf)
				return num
			}
		}
	}
	return 0
}

func checkGear(coords []int, lines []string) int {
	adjacents := calculateAdjacents(coords, lines)

	gear := []int{}

	for i := 0; i < len(adjacents); i++ {
		if strings.ContainsAny(string(adjacents[i]), nums) {
			var gearHalf string
			var line int
			var column int

			//Solve Co-ordinates Location based on i
			switch i {
			case 0:
				line = coords[0]-1 
				column = coords[1]-1
			case 1:
				line = coords[0]-1 
				column = coords[1]
			case 2:
				line = coords[0]-1 
				column = coords[1]+1
			case 3:
				line = coords[0]
				column = coords[1]-1
			case 4:
				line = coords[0]
				column = coords[1]+1
			case 5:
				line = coords[0]+1 
				column = coords[1]-1
			case 6:
				line = coords[0]+1 
				column = coords[1]
			case 7:
				line = coords[0]+1 
				column = coords[1]+1
			}

			for y := column; y < len(lines[line]); y++ { //Iterating forward through the found number adjacencies to get the whole number.
				if strings.ContainsAny(string(lines[line][y]), nums) {
					gearHalf += string(lines[line][y])
					if i < 2 || i > 4 { //Incrementing where adjacent numbers have already been processed.
						i++
					}
				} else {
					break
				}
			}

			for x := column-1; x >= 0; x-- { //Iterating backwards for numbers which appear behind *.
				if strings.ContainsAny(string(lines[line][x]), nums) {
					gearHalf = string(lines[line][x]) + gearHalf
				} else {
					break
				}
			}

			gearCalc, _ := strconv.Atoi(gearHalf)
			gear = append(gear, gearCalc)
		}
	}

	if len(gear) == 2 { //2 adjacencies = a gear
		return gear[0] * gear[1]
	}

	return 0 //failsafe: return nothing
}

func main() {
	nums = "0123456789" //finding criteria
	var totalSum int
	var totalGear int

	input, err := os.ReadFile("input.txt")
	if err!= nil {
		fmt.Println("There has been an error reading the input file.")
	}
	lines := strings.Split(string(input), "\n")

	for i, v := range lines {
		//Task 1
		for o := 0; o < len(v); o++ {
			if strings.ContainsAny(string(v[o]), nums) {
				totalSum += checkParts(o, lines, i, v)
			}
		}
		//Task 2
		for o := 0; o < len(v); o++ {
			if strings.ContainsAny(string(v[o]), "*") {
				gearCoords := []int{i, o}
				totalGear += checkGear(gearCoords, lines)
			}
		}

	}
	fmt.Println("Part Numbers:", totalSum)
	fmt.Println("Gear Ratio:", totalGear)
}