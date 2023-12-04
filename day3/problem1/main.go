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
		output += findPartNumbers(currentLine, []rune(nextLine), []rune(prevLine))
		prevLine = currentLine
		currentLine = nextLine
	}
	output += findPartNumbers(currentLine, nil, []rune(prevLine))

	log.Printf("Value: %d\n", output)
}

func isSymbol(r rune) bool {
	if !unicode.IsNumber(r) && r != '.' {
		return true
	}
	return false
}

func findPartNumbers(line string, nextLineRunes []rune, prevLineRunes []rune) int {
	lineRunes := []rune(line)
	var sum int
	idx := 0
	for {
		if idx == 0 {
			idx++
			continue
		} else if idx >= len(lineRunes)-1 {
			return sum
		}
		if unicode.IsNumber(lineRunes[idx]) {
			if len(prevLineRunes) == 0 {
				if isSymbol(lineRunes[idx-1]) || isSymbol(lineRunes[idx+1]) {
					pNumberLocation := findNumberLocation(lineRunes, idx)
					pNumber, _ := strconv.Atoi(line[pNumberLocation[0]:pNumberLocation[1]])
					sum += pNumber
					idx = pNumberLocation[1] + 1
				} else if isSymbol(nextLineRunes[idx-1]) || isSymbol(nextLineRunes[idx]) || isSymbol(nextLineRunes[idx+1]) {
					pNumberLocation := findNumberLocation(lineRunes, idx)
					pNumber, _ := strconv.Atoi(line[pNumberLocation[0]:pNumberLocation[1]])
					sum += pNumber
					idx = pNumberLocation[1] + 1
				} else {
					idx++
				}
			} else if nextLineRunes == nil {
				if isSymbol(lineRunes[idx-1]) || isSymbol(lineRunes[idx+1]) {
					pNumberLocation := findNumberLocation(lineRunes, idx)
					pNumber, _ := strconv.Atoi(line[pNumberLocation[0]:pNumberLocation[1]])
					sum += pNumber
					idx = pNumberLocation[1] + 1
				} else if isSymbol(prevLineRunes[idx-1]) || isSymbol(prevLineRunes[idx]) || isSymbol(prevLineRunes[idx+1]) {
					pNumberLocation := findNumberLocation(lineRunes, idx)
					pNumber, _ := strconv.Atoi(line[pNumberLocation[0]:pNumberLocation[1]])
					sum += pNumber
					idx = pNumberLocation[1] + 1
				} else {
					idx++
				}
			} else {
				if isSymbol(lineRunes[idx-1]) || isSymbol(lineRunes[idx+1]) {
					pNumberLocation := findNumberLocation(lineRunes, idx)
					pNumber, _ := strconv.Atoi(line[pNumberLocation[0]:pNumberLocation[1]])
					sum += pNumber
					idx = pNumberLocation[1] + 1
				} else if isSymbol(nextLineRunes[idx-1]) || isSymbol(nextLineRunes[idx]) || isSymbol(nextLineRunes[idx+1]) {
					pNumberLocation := findNumberLocation(lineRunes, idx)
					pNumber, _ := strconv.Atoi(line[pNumberLocation[0]:pNumberLocation[1]])
					sum += pNumber
					idx = pNumberLocation[1] + 1
				} else if isSymbol(prevLineRunes[idx-1]) || isSymbol(prevLineRunes[idx]) || isSymbol(prevLineRunes[idx+1]) {
					pNumberLocation := findNumberLocation(lineRunes, idx)
					pNumber, _ := strconv.Atoi(line[pNumberLocation[0]:pNumberLocation[1]])
					sum += pNumber
					idx = pNumberLocation[1] + 1
				} else {
					idx++
				}
			}
		} else {
			idx++
		}
	}
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
