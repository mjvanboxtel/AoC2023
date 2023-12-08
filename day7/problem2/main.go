package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	HighCard = iota
	OnePair
	TwoPair
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

type hand struct {
	handType int
	runes    []rune
	str      string
	bidValue int
}

func main() {
	start := time.Now()
	defer log.Printf("Execution time %s", time.Since(start))

	readFile, err := os.Open("input.txt")
	if err != nil {
		log.Println(err)
	}
	defer readFile.Close()
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	var output int

	var hands []hand
	for fileScanner.Scan() {
		fields := strings.Fields(fileScanner.Text())
		handString := fields[0]
		bid, _ := strconv.Atoi(fields[1])
		hands = append(hands, hand{handType: getHandType([]rune(handString)), runes: []rune(handString), str: handString, bidValue: bid})
	}

	c := make(chan []hand)
	go sortHands(hands, c)
	sortedHands := <-c

	for idx, h := range sortedHands {
		output += h.bidValue * (idx + 1)
	}

	log.Printf("Value: %d\n", output)
}

func getHandType(hand []rune) int {
	cardMap := map[rune]int{}
	for _, c := range hand {
		if _, ok := cardMap[c]; !ok {
			cardMap[c] = 1
		} else {
			cardMap[c]++
		}
	}
	distinctCards := len(cardMap)
	var handType int
	if distinctCards == 1 {
		handType = FiveOfAKind
	} else if distinctCards == 2 {
		for _, c := range cardMap {
			if c == 4 || c == 1 {
				handType = FourOfAKind
			} else {
				handType = FullHouse
			}
		}
	} else if distinctCards == 3 {
		for _, c := range cardMap {
			if c == 3 {
				handType = ThreeOfAKind
			} else if c == 2 {
				handType = TwoPair
			}
		}
	} else if distinctCards == 4 {
		handType = OnePair
	} else {
		handType = HighCard
	}
	if _, ok := cardMap['J']; ok {
		return handleJokers(handType, cardMap['J'])
	}
	return handType
}

func handleJokers(handType int, jokerCount int) int {
	if handType == HighCard {
		return OnePair
	} else if handType == OnePair {
		return ThreeOfAKind
	} else if handType == TwoPair {
		if jokerCount == 2 {
			return FourOfAKind
		} else {
			return FullHouse
		}
	}
	if handType == ThreeOfAKind {
		return FourOfAKind
	}
	return FiveOfAKind
}

func sortHands(hands []hand, c chan []hand) {
	if len(hands) <= 1 {
		c <- hands
		return
	}
	middle := len(hands) / 2
	leftChan := make(chan []hand)
	rightChan := make(chan []hand)
	go sortHands(hands[:middle], leftChan)
	go sortHands(hands[middle:], rightChan)
	left := <-leftChan
	right := <-rightChan
	c <- mergeHands(left, right)
}

func mergeHands(left, right []hand) []hand {
	result := make([]hand, len(left)+len(right))
	cardValues := getCardStrengthMap()
	i, j := 0, 0
	for k := 0; k < len(result); k++ {
		if i >= len(left) {
			result[k] = right[j]
			j++
		} else if j >= len(right) {
			result[k] = left[i]
			i++
		} else if left[i].handType < right[j].handType {
			result[k] = left[i]
			i++
		} else if left[i].handType > right[j].handType {
			result[k] = right[j]
			j++
		} else {
			for idx, _ := range left[i].runes {
				if cardValues[left[i].runes[idx]] < cardValues[right[j].runes[idx]] {
					result[k] = left[i]
					i++
					break
				} else if cardValues[left[i].runes[idx]] > cardValues[right[j].runes[idx]] {
					result[k] = right[j]
					j++
					break
				}
			}
		}
	}
	return result
}

func getCardStrengthMap() map[rune]int {
	return map[rune]int{
		'A': 13,
		'K': 12,
		'Q': 11,
		'T': 10,
		'9': 9,
		'8': 8,
		'7': 7,
		'6': 6,
		'5': 5,
		'4': 4,
		'3': 3,
		'2': 2,
		'J': 1,
	}
}
