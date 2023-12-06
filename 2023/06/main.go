package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sample = `Time:      7  15   30
Distance:  9  40  200`

func parseText(input string) [][]int {
	nl := strings.Index(input, "\n")
	timeString := strings.Fields(input[strings.Index(input[:nl], ":")+1 : nl])
	distanceString := strings.Fields(input[strings.Index(input[nl:], ":")+nl+1:])
	var races [][]int
	for i := 0; i < len(timeString); i++ {
		time, _ := strconv.Atoi(timeString[i])
		distance, _ := strconv.Atoi(distanceString[i])
		races = append(races, []int{time, distance})
	}
	return races
}

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("There has been an error reading this file:", err)
	}
	races := parseText(string(input)) //[0] = time, [1] = distance
	fmt.Println("(Part 1) Races:", races)

	var totalRecords []int
	var concatTime, concatDistance string
	// Part 1
	for _, race := range races {
		var records int
		for ms := 1; ms < race[0]-1; ms++ { //Handling edges that return 0.
			distance := ms * (race[0] - ms)
			// fmt.Println("Held Button for", ms, "ms - Travelling a distance of", distance, "mm")
			if distance > race[1] {
				// fmt.Println("Wow! This beat our old record.")
				records++
			}
		}
		totalRecords = append(totalRecords, records)
		concatTime += strconv.Itoa(race[0])
		concatDistance += strconv.Itoa(race[1]) //Concat
	}

	fmt.Println("(Part 1) Total Records for Each Race:", totalRecords)
	var total int
	for i, records := range totalRecords {
		if i == 0 {
			total += records
		} else {
			total = total * records //i * i * i * i...
		}
	}
	fmt.Printf("(Part 1) Total Records Multiplied: %d\n\n", total)

	// Part 2
	fmt.Println("(Part 2) Race Time:", concatTime)
	fmt.Println("(Part 2) Race Distance:", concatDistance)
	fullTime, _ := strconv.Atoi(concatTime)
	fullDistance, _ := strconv.Atoi(concatDistance)
	var fullRecord int
	for ms := 1; ms < fullTime-1; ms++ {
		distance := ms * (fullTime - ms)
		//Print statements removed to avoid print-based slowdown.
		if distance > fullDistance {
			fullRecord++
		}
	}

	fmt.Println("(Part 2) Full Race Record:", fullRecord)
}
