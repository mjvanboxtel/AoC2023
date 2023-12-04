package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
	"unicode"
)

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

	for fileScanner.Scan() {
		output += parseLine(strings.Split(fileScanner.Text(), ":")[1])
	}

	log.Printf("Value: %d\n", output)
}

func parseLine(line string) int {
	lineRunes := []rune(line)
	winningNumbers := map[int]bool{}
	hits := 0
	wSwitch := true
	idx := 0
	for {
		if idx == len(lineRunes) {
			break
		}
		if lineRunes[idx] == '|' {
			wSwitch = false
		}
		if unicode.IsNumber(lineRunes[idx]) {
			nLoc := findNumberLocation(lineRunes, idx)
			var nValue int
			if nLoc == nil {
				nValue, _ = strconv.Atoi(line[idx:len(lineRunes)])
			} else if len(nLoc) == 2 {
				nValue, _ = strconv.Atoi(line[nLoc[0]:nLoc[1]])
				idx++
			} else {
				nValue = int(line[nLoc[0]] - '0')
			}
			if wSwitch {
				winningNumbers[nValue] = true
			} else {
				if _, ok := winningNumbers[nValue]; ok {
					hits++
				}
			}
		}
		idx++
	}
	if hits == 0 {
		return 0
	}
	return int(math.Pow(float64(2), float64(hits-1)))
}

func findNumberLocation(lineRunes []rune, idx int) []int {
	if idx == len(lineRunes)-1 {
		return nil
	}
	if unicode.IsNumber(lineRunes[idx+1]) {
		return []int{idx, idx + 2}
	} else {
		return []int{idx}
	}
}
