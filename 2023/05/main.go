package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// var sample string =
// `seeds: 79 14 55 13

// seed-to-soil map:
// 50 98 2
// 52 50 48

// soil-to-fertilizer map:
// 0 15 37
// 37 52 2
// 39 0 15

// fertilizer-to-water map:
// 49 53 8
// 0 11 42
// 42 0 7
// 57 7 4

// water-to-light map:
// 88 18 7
// 18 25 70

// light-to-temperature map:
// 45 77 23
// 81 45 19
// 68 64 13

// temperature-to-humidity map:
// 0 69 1
// 1 0 69

// humidity-to-location map:
// 60 56 37
// 56 93 4`

func main() {
	input, err := os.ReadFile("input.txt")
	if err!= nil {
		fmt.Println("There has been an error reading the input file.")
	}
	atlas := strings.Split(string(input), "\n\n")
	
	var masterRange [][][]int

	for _, chart := range atlas {
		var chartRange [][]int
		rangeStrings := strings.Split(chart, "\n")
		for _, rangeString := range rangeStrings {
			nums := strings.Fields(rangeString)
			var currentRange []int
			for _, v := range nums {
				num, err := strconv.Atoi(v)
				if err != nil {
					continue
				}
				currentRange = append(currentRange, num)
			}
			if currentRange != nil {
				chartRange = append(chartRange, currentRange)
			}
		}
		masterRange = append(masterRange, chartRange)
	}

	seeds := masterRange[0][0]
	mapped := make(map[int]int)
	
	for i, ranges := range masterRange {
		if i == 0 {
			for _, seed := range seeds {
				mapped[seed] = seed
			}
			fmt.Printf("Initial Seeds: %v\n\n", mapped)
			continue
		}

		localMapped := make(map[int]int)
		for _, v := range ranges {
			rangeLength := v[2]
			sRangeStart := v[1]
			dRangeStart := v[0]

			for _, seed := range seeds {
				if (mapped[seed] >= sRangeStart && mapped[seed] < sRangeStart + rangeLength)  {
					localMapped[seed] = mapped[seed] + (dRangeStart - sRangeStart)
					mapped[seed] = -1
				} 
			}
		}
		for k, v := range localMapped {
			mapped[k] = v
		}
	}

	fmt.Printf("Processed Seeds: %v\n\n", mapped)

	min := math.MaxInt32
	for _, v := range mapped {
		if (v < min) {
			min = v
		}
	}

	fmt.Println("Lowest Seed:", min)
	
}
