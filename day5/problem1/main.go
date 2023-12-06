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

	seedLoc := map[int]int{}
	var output int

	fileScanner.Scan()
	seedLine := strings.Split(fileScanner.Text(), ":")
	seeds := strings.Fields(seedLine[1])
	for _, seed := range seeds {
		seedI, _ := strconv.Atoi(seed)
		seedLoc[seedI] = seedI
	}
	found := map[int]bool{}

	for fileScanner.Scan() {
		line := fileScanner.Text()
		lineRunes := []rune(line)
		for idx, _ := range seedLoc {
			if len(lineRunes) > 0 {
				if unicode.IsNumber(lineRunes[0]) {
					if found[idx] != true {
						locMap := strings.Fields(line)
						source, _ := strconv.Atoi(locMap[1])
						destination, _ := strconv.Atoi(locMap[0])
						rangeVal, _ := strconv.Atoi(locMap[2])
						if source <= seedLoc[idx] && seedLoc[idx] <= (source+rangeVal) {
							seedLoc[idx] = seedLoc[idx] + (destination - source)
							found[idx] = true
						}
					}
				} else {
					found = map[int]bool{}
				}
			}
		}
	}
	lowest := math.MaxInt64
	for _, loc := range seedLoc {
		if loc < lowest {
			lowest = loc
		}
	}
	output = lowest
	log.Printf("Value: %d\n", output)
}
