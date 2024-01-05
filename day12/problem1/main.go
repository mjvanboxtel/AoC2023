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
	unknown     = '?'
	broken      = '#'
	operational = '.'
)

func main() {
	start := time.Now()
	defer func() { log.Println(time.Since(start)) }()

	readFile, err := os.Open("input.txt")
	if err != nil {
		log.Println(err)
	}
	defer readFile.Close()
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	var output int
	for fileScanner.Scan() {
		line := strings.Fields(fileScanner.Text())
		perms := findPermutations([]rune(line[0]))
		for i := 0; i < len(perms); i++ {
			if validPermutation(perms[i], getConditionRecords(line[1])) {
				output++
			}
		}
	}

	log.Printf("Value: %d\n", output)
}

func validPermutation(perm []rune, conditionRecords []int) bool {
	var cdnIdx int
	var currentSprings int
	for i := 0; i < len(perm); i++ {
		if perm[i] == broken {
			currentSprings++
		}
		if currentSprings > 0 && perm[i] == operational {
			if cdnIdx == len(conditionRecords) || currentSprings != conditionRecords[cdnIdx] {
				return false
			}
			currentSprings = 0
			cdnIdx++
		}
		if i == len(perm)-1 {
			if perm[i] == broken {
				if cdnIdx == len(conditionRecords) || conditionRecords[cdnIdx] != currentSprings || cdnIdx < len(conditionRecords)-1 {
					return false
				}
			} else {
				if cdnIdx <= len(conditionRecords)-1 {
					return false
				}
			}
		}
	}
	return true
}

func findPermutations(chars []rune) [][]rune {
	if len(chars) == 1 {
		if chars[0] == unknown {
			return [][]rune{{operational}, {broken}}
		}
		return [][]rune{{chars[0]}}
	}

	var allPermutations [][]rune

	var initial []rune
	initial = append(initial, chars[1:]...)
	permutations := findPermutations(initial)
	for j := 0; j < len(permutations); j++ {
		if chars[0] == unknown {
			remaining := []rune{operational, broken}
			for k := 0; k < len(remaining); k++ {
				var newP []rune
				newP = append(newP, remaining[k])
				newP = append(newP, permutations[j]...)
				allPermutations = append(allPermutations, newP)
			}
		} else {
			var newP []rune
			newP = append(newP, chars[0])
			newP = append(newP, permutations[j]...)
			allPermutations = append(allPermutations, newP)
		}
	}

	return allPermutations
}

func getConditionRecords(line string) []int {
	records := strings.Split(line, ",")
	var out []int
	for i := 0; i < len(records); i++ {
		val, _ := strconv.Atoi(records[i])
		out = append(out, val)
	}
	return out
}
