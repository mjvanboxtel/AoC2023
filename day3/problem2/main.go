package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
	"unicode"
)

func main() {
	start := time.Now()
	defer log.Printf("Execution time %s", time.Since(start))

	readFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer readFile.Close()
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	var output int
	var prevLine string
	var nextLine string

	fileScanner.Scan()
	currentLine := fileScanner.Text()
	for fileScanner.Scan() {
		nextLine = fileScanner.Text()
		output += findGearRatios(currentLine, nextLine, prevLine)
		prevLine = currentLine
		currentLine = nextLine
	}
	output += findGearRatios(currentLine, "", prevLine)

	log.Printf("Value: %d\n", output)
}

func findGearRatios(line string, nextLine string, prevLine string) int {
	lineRunes := []rune(line)
	nextLineRunes := []rune(nextLine)
	prevLineRunes := []rune(prevLine)
	var sum int
	idx := 0
	for {
		if idx == 0 {
			idx++
			continue
		} else if idx >= len(lineRunes)-1 {
			return sum
		}
		if lineRunes[idx] == '*' {
			var partNumbers []int
			// look above
			if len(prevLineRunes) != 0 {
				partNumbers = append(partNumbers, identifyPartNumbers(idx, prevLine, prevLineRunes)...)
			}
			// look below
			if nextLine != "" {
				partNumbers = append(partNumbers, identifyPartNumbers(idx, nextLine, nextLineRunes)...)
			}
			// look adjacent
			partNumbers = append(partNumbers, identifyPartNumbers(idx, line, lineRunes)...)
			if len(partNumbers) == 2 {
				sum += partNumbers[0] * partNumbers[1]
			}
		}
		idx++
	}
}

func identifyPartNumbers(idx int, line string, lineRunes []rune) []int {
	partNumbersSet := map[int]bool{}
	var partNumbers []int
	for _, val := range []int{idx - 1, idx, idx + 1} {
		if unicode.IsNumber(lineRunes[val]) {
			pNumberLocation := findNumberLocation(lineRunes, val)
			pNumber, _ := strconv.Atoi(line[pNumberLocation[0]:pNumberLocation[1]])
			partNumbersSet[pNumber] = true
		}
	}
	for partNumber := range partNumbersSet {
		partNumbers = append(partNumbers, partNumber)
	}
	return partNumbers
}

func findNumberLocation(lineRunes []rune, startingIndex int) []int {
	indexMap := map[string]int{}
	currentIndex := startingIndex
	for {
		if currentIndex == len(lineRunes)-1 {
			indexMap["last"] = currentIndex + 1
		}
		if currentIndex == 0 {
			indexMap["first"] = currentIndex
		}
		if _, lastExists := indexMap["last"]; !lastExists {
			if unicode.IsNumber(lineRunes[currentIndex+1]) {
				currentIndex++
			} else {
				indexMap["last"] = currentIndex + 1
				currentIndex = startingIndex
			}
		} else if _, firstExists := indexMap["first"]; !firstExists {
			if unicode.IsNumber(lineRunes[currentIndex-1]) {
				currentIndex--
			} else {
				indexMap["first"] = currentIndex
				currentIndex = startingIndex
			}
		} else {
			break
		}
	}
	return []int{indexMap["first"], indexMap["last"]}
}
