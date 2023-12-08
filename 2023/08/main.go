package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func parseText(input []byte) ([]byte, map[string][]string, []string) {
	firstLine := bytes.Index(input, []byte("\n"))
	directions := input[:firstLine]
	lines := strings.Split(string(input[firstLine+2:]), "\n")
	locations := make(map[string][]string)

	var starts []string
	for _, v := range lines {
		if v[2:3] == "A" {
			starts = append(starts, v[0:3])
		}
		locations[v[0:3]] = []string{v[7:10], v[12:15]}
	}

	return directions, locations, starts
}

// Greatest Common Divisor
func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

var sample = []byte(`RL

AAA = (BBB, CCC)
BBB = (DDD, EEE)
CCC = (ZZZ, GGG)
DDD = (DDD, DDD)
EEE = (EEE, EEE)
GGG = (GGG, GGG)
ZZZ = (ZZZ, ZZZ)`)

var sample2 = []byte(`LLR

AAA = (BBB, BBB)
BBB = (AAA, ZZZ)
ZZZ = (ZZZ, ZZZ)`)

var sample3 = []byte(`LR

11A = (11B, XXX)
11B = (XXX, 11Z)
11Z = (11B, XXX)
22A = (22B, XXX)
22B = (22C, 22C)
22C = (22Z, 22Z)
22Z = (22B, 22B)
XXX = (XXX, XXX)`)

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("There has been an error reading the file.")
		panic(err)
	}

	directions, locations, starts := parseText(input)

	//Part 1
	nextLocation := "AAA"
	var steps int
	for nextLocation != "ZZZ" {
		for _, v := range directions {
			switch v {
			case 'L':
				nextLocation = locations[nextLocation][0]
			case 'R':
				nextLocation = locations[nextLocation][1]
			}
			steps++
		}
	}

	fmt.Println("(Part 1) Steps:", steps)

	// (Part 2) Brute Force Method - Estimated Time 2,000 Days...

	// var steps2 int
	// for !checkZ(starts) {
	// 	for _, v := range directions {
	// 		for i := range starts {
	// 			switch v {
	// 			case 'L':
	// 				starts[i] = locations[starts[i]][0]
	// 			case 'R':
	// 				starts[i] = locations[starts[i]][1]
	// 			}
	// 		}
	// 		steps2++
	// 	}
	// 	fmt.Println(steps2)
	// }

	// Part 2 - Utilizing LCM.
	steps2 := make([]int, len(starts))
	for i := range starts {
		for starts[i][2:3] != "Z" {
			for _, v := range directions {
				switch v {
				case 'L':
					starts[i] = locations[starts[i]][0]
				case 'R':
					starts[i] = locations[starts[i]][1]
				}
				steps2[i]++
			}
		}
	}

	//https://www.geeksforgeeks.org/lcm-of-given-array-elements/
	lcm := 1
	for _, v := range steps2 {
		lcm = (v * lcm) / gcd(v, lcm)
	}
	fmt.Println("(Part 2) Steps:", lcm)
}
