package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func word(inputString string) (int, int, int, int)  {
	var firstIndex = -1 
	var firstVal, lastIndex, lastVal int
	numMap := map[string]int{
		"one": 1,
		"two": 2,
		"three": 3,
		"four": 4,
		"five": 5,
		"six": 6,
		"seven": 7,
		"eight": 8,
		"nine": 9,
	}
	for k, v := range numMap {
		if strings.Contains(inputString, k) && (firstIndex == -1 || firstIndex >= strings.Index(inputString, k)) {
			firstIndex = strings.Index(inputString, k)
			firstVal = v 
		}

		if lastIndex < strings.LastIndex(inputString, k) {
			lastIndex = strings.LastIndex(inputString, k)
			lastVal = v
		}
	}

	return firstIndex, firstVal, lastIndex, lastVal
}

func num(inputString string) (int, int)  {
	var first, last int
	nums := "0123456789"
	first = strings.IndexAny(inputString, nums)
	last = strings.LastIndexAny(inputString, nums)

	return first, last
}


func main() {
	var total int
	input, err := os.ReadFile("input.txt")
	if err!= nil {
		fmt.Println("There has been an error reading the input file.")
	}

	array := strings.Split(string(input), "\n")

	for _, v := range array {
		var first, last int
		firstWordIndex, firstWordVal, lastWordIndex, lastWordVal := word(v)
		firstNumIndex, lastNumIndex := num(v)

		if firstWordIndex < firstNumIndex && firstWordIndex != -1 {
			first = firstWordVal
		} else {
			first, _ = strconv.Atoi(string(v[firstNumIndex]))
		}

		if lastWordIndex > lastNumIndex {
			last = lastWordVal
		} else {
			last, _ = strconv.Atoi(string(v[lastNumIndex]))
		}

		string := fmt.Sprint(first) + fmt.Sprint(last)
		num, err := strconv.Atoi(string)

		if err != nil {
			fmt.Printf("There has been an error converting %s - %s\n", string, err)
		}

		total += num
	}
	fmt.Println(total)
}