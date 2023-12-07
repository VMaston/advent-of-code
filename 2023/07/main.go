package main

import (
	"bytes"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type handType struct {
	cards []int
	bid   int
	set   int
}

func parseText(input []byte) []handType {
	lines := bytes.Split(input, []byte{'\n'})
	var hand []handType
	for _, line := range lines {
		sep := bytes.Index(line, []byte{' '})
		num, _ := strconv.Atoi(string(line[sep+1:]))

		var cardInts []int
		// Part 1
		// faceNum := map[string]int{"T": 10, "J": 11, "Q": 12, "K": 13, "A": 14}
		// Part 2
		faceNum := map[string]int{"T": 10, "J": 0, "Q": 12, "K": 13, "A": 14}
		for _, card := range line[:sep] {
			cardInt, err := strconv.Atoi(string(card))
			if err != nil {
				cardInt = faceNum[string(card)]
			}
			cardInts = append(cardInts, cardInt)
		}
		hand = append(hand, handType{cards: cardInts, bid: num})
	}
	return hand
}

func calculateHand(hand handType) handType {
	cardCounts := make(map[int]int)
	highest := struct {
		card  int
		count int
	}{}
	for _, card := range hand.cards {
		cardCounts[card]++
		if cardCounts[card] > highest.count && card != 0 {
			highest.count = cardCounts[card]
			highest.card = card
		}
	}

	if highest.card != 0 {
		cardCounts[highest.card] += cardCounts[0]
		delete(cardCounts, 0)
	}

	switch len(cardCounts) {
	case 1:
		fmt.Println("Five of a Kind!")
		hand.set = 7
	case 2:
		for _, v := range cardCounts {
			if v == 1 || v == 4 {
				fmt.Println("Four of a Kind!")
				hand.set = 6
				break
			} else {
				fmt.Println("Full House!")
				hand.set = 5
				break
			}
		}
	case 3:
		for _, v := range cardCounts {
			if v == 3 {
				fmt.Println("Three of a Kind!")
				hand.set = 4
				break
			}
			if v == 2 {
				fmt.Println("Two Pair!")
				hand.set = 3
				break
			}
		}
	case 4:
		fmt.Println("One Pair!")
		hand.set = 2
	case 5:
		fmt.Println("High Card!")
		hand.set = 1
	}

	return hand
}

var sample = `2345A 1
Q2KJJ 13
Q2Q2Q 19
T3T3J 17
T3Q33 11
2345J 3
J345A 2
32T3K 5
T55J5 29
KK677 7
KTJJT 34
QQQJA 31
JJJJJ 37
JAAAA 43
AAAAJ 59
AAAAA 61
2AAAA 23
2JJJJ 53
JJJJ2 41`

func main() {
	input, err := os.ReadFile("./input.txt")
	if err != nil {
		fmt.Println("There has been an error reading this file.")
	}

	hands := parseText([]byte(input))
	fmt.Println(hands)

	for i := range hands {
		hands[i] = calculateHand(hands[i])
		fmt.Printf("%v\n\n", hands[i])
	}

	sort.Slice(hands, func(i, j int) bool {
		if hands[i].set == hands[j].set {
			for k := 0; k < len(hands[i].cards); k++ {
				if hands[i].cards[k] == hands[j].cards[k] {
					continue
				} else {
					return hands[i].cards[k] < hands[j].cards[k]
				}
			}
		}
		return hands[i].set < hands[j].set
	})

	var total int
	for i, v := range hands {
		total += v.bid * (i + 1)
	}

	fmt.Println("Total Winnings:", total)
}
