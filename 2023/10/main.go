package main

import (
	"fmt"
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
			if vv == '7' {
				vv = '¬'
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
		if v[0] < 0 || v[0] > len(pipes[0]) || v[1] < 0 || v[1] > len(pipes[1]) {
			continue
		}

		symbol := pipes[v[0]][v[1]]

		switch i {
		case 0:
			switch symbol {
			case "|", "¬", "F":
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
			case "-", "J", "¬":
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
		case "¬":
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
		case "¬":
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
	`...........
.S-------7.
.|F-----7|.
.||.....||.
.||.....||.
.|L-7.F-J|.
.|..|.|..|.
.L--J.L--J.
...........`)

func boundaries(arr [][]string, i int) bool {
	if i < 0 || i > len(arr) {
		return false
	}
	return true
}

func main() {
	// input, err := os.ReadFile("input.txt")
	// if err != nil {
	// 	fmt.Println("There has been an error reading the file.")
	// 	panic(err)
	// }

	pipes, start := parseInput(sample)

	running := true
	steps := 1
	prev := start
	next := getStartDirection(start, pipes)
	mapY := make(map[int][2]int)
	mapX := make(map[int][2]int)

	for running {
		for i, line := range pipes {
			for ii, pipe := range line {
				if i == next[0] && ii == next[1] {
					dir := getDirection(prev, next, pipe)
					if dir == [2]int{} {
						running = false
					}
					prev = []int{i, ii}
					next[0], next[1] = dir[0], dir[1]
					pipes[i][ii] = fmt.Sprint(steps)
					steps++
				}
			}
		}
	}

	// Part 2 - Unfinished (Perimeter algorithm did not account for irregular shapes)
	for i, line := range pipes {
		mapY[i] = [2]int{999, 0}
		for ii, pipe := range line {
			if _, ok := mapX[ii]; !ok {
				mapX[ii] = [2]int{999, 0}
			}
			if strings.ContainsAny(pipe, "0123456789") {
				if ii > mapY[i][1] {
					mapY[i] = [2]int{mapY[i][0], ii}
				}
				if ii < mapY[i][0] {
					mapY[i] = [2]int{ii, mapY[i][1]}
				}
				if i > mapX[ii][1] {
					mapX[ii] = [2]int{mapX[ii][0], i}
				}
				if i < mapX[ii][0] {
					mapX[ii] = [2]int{i, mapX[ii][1]}
				}
			}
		}
	}

	var iCount int
	for i, line := range pipes {
		for ii, pipe := range line {
			if !strings.ContainsAny(pipe, "0123456789") {
				if ii > mapY[i][0] && ii < mapY[i][1] && i > mapX[ii][0] && i < mapX[ii][1] {
					pipes[i][ii] = "I"
					iCount++
				} else {
					pipes[i][ii] = "O"
				}
			} else {
				pipes[i][ii] = "*"
			}
		}
	}

	for _, line := range pipes {
		fmt.Println(line)
	}
	fmt.Println(iCount)

}
