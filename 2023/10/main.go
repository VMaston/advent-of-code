package main

import (
	"fmt"
	"os"
	"strings"
)

func parseInput(input []byte) ([][]string, []int) {
	lines := strings.Split(string(input), "\n")
	var arr [][]string
	var start []int
	for i, v := range lines {
		arr = append(arr, []string{})
		for ii, vv := range v {
			if vv == 'S' {
				start = append(start, i, ii)
			}
			arr[i] = append(arr[i], string(vv))
		}
	}

	return arr, start
}

func getStartDirection(start []int, pipes [][]string) [2]int { //Height, Width
	var startDirections [2]int

	north := [2]int{start[0] - 1, start[1]}
	west := [2]int{start[0], start[1] - 1}
	east := [2]int{start[0], start[1] + 1}
	south := [2]int{start[0] + 1, start[1]}

	directions := [][2]int{north, west, east, south}

	for i, v := range directions {
		//Boundary Limits
		if v[0] < 0 || v[0] > len(pipes) || v[1] < 0 || v[1] > len(pipes) {
			continue
		}

		symbol := pipes[v[0]][v[1]]

		switch i {
		case 0:
			switch symbol {
			case "|", "7", "F":
				fmt.Println("Compatible with North")
				startDirections = north
			}
		case 1:
			switch symbol {
			case "-", "L", "F":
				startDirections = west
			}
		case 2:
			switch symbol {
			case "-", "J", "7":
				startDirections = east
			}
		case 3:
			switch symbol {
			case "|", "L", "J":
				startDirections = south
			}
		}
	}

	return startDirections
}

func getDirection(prev []int, step [2]int, symbol string) [2]int {

	if symbol == "S" {
		fmt.Println("End Reached.")
		return [2]int{}
	}

	north := [2]int{step[0] - 1, step[1]}
	west := [2]int{step[0], step[1] - 1}
	east := [2]int{step[0], step[1] + 1}
	south := [2]int{step[0] + 1, step[1]}

	directions := [][2]int{north, west, east, south}
	var startDirection [2]int

	for _, v := range directions {
		if v[0] == prev[0] && v[1] == prev[1] {
			startDirection = v
			break
		}
	}

	fmt.Println(step, startDirection, symbol, prev)

	switch startDirection {
	case north:
		switch symbol {
		case "|":
			return south
		case "L":
			return east
		case "J":
			return west
		default:
			fmt.Println("Prev was North but symbol didn't match.")
		}
	case west:
		switch symbol {
		case "-":
			return east
		case "J":
			return north
		case "7":
			return south
		default:
			fmt.Println("Prev was West but symbol didn't match.")
		}
	case east:
		switch symbol {
		case "-":
			return west
		case "L":
			return north
		case "F":
			return south
		default:
			fmt.Println("Prev was East but symbol didn't match.")
		}
	case south:
		switch symbol {
		case "|":
			return north
		case "F":
			return east
		case "7":
			return west
		default:
			fmt.Println("Prev was South but symbol didn't match.", symbol)
		}
	default:
		fmt.Println("Could not anticipate direction.")
	}
	panic("Did not have start direction!")
}

var sample = []byte(
	`..F7.
.FJ|.
SJ.L7
|F--J
LJ...`)

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("There has been an error reading the file.")
		panic(err)
	}

	pipes, start := parseInput(input)

	running := true

	steps := 1
	prev := start
	next := getStartDirection(start, pipes)
	for iterations := 0; running; iterations++ {
		for ii, line := range pipes {
			for iii, pipe := range line {
				if ii == next[0] && iii == next[1] {
					dir := getDirection(prev, next, pipe)
					if dir == [2]int{} {
						running = false
					}
					prev = []int{ii, iii}
					next[0], next[1] = dir[0], dir[1]
					pipes[ii][iii] = fmt.Sprint(steps)
					steps++
				}
			}
		}
	}

	for _, pipe := range pipes {
		fmt.Println(pipe)
	}
	fmt.Println(steps / 2)
}
