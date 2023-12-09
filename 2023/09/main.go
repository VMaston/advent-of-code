package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func parseInput(input []byte) [][][]int {
	lines := strings.Split(string(input), "\n")
	sequence := make([][][]int, len(lines))
	for i, line := range lines {
		chars := strings.Split(line, " ")
		var nums []int
		for _, char := range chars {
			num, _ := strconv.Atoi(char)
			nums = append(nums, num)
		}
		sequence[i] = append(sequence[i], nums)
	}

	return sequence
}

var sample = []byte(`0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45`)

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("There has been an error reading this file.")
		panic(err)
	}

	sequences := parseInput(input)

	for i := 0; i < len(sequences); i++ {
		for ii := 0; ii < len(sequences[i]); ii++ {
			var nextLine []int
			for iii := 0; iii < len(sequences[i][ii])-1; iii++ {
				nextLine = append(nextLine, sequences[i][ii][iii+1]-sequences[i][ii][iii])
			}
			if slices.Max(sequences[i][ii]) != 0 || slices.Min(sequences[i][ii]) != 0 {
				sequences[i] = append(sequences[i], nextLine)
			}
		}
	}

	var ans1 int
	for _, v := range sequences {
		for i := len(v) - 1; i >= 0; i-- {
			ans1 += v[i][len(v[i])-1]
		}
	}

	var ans2 int
	for _, v := range sequences {
		var total int
		for i := len(v) - 1; i >= 0; i-- {
			total = v[i][0] - total
		}
		ans2 += total
	}

	fmt.Println("(Part 1) Sum:", ans1)
	fmt.Println("(Part 2) Sum:", ans2)

}
