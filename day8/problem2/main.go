package main

import (
	"bufio"
	"log"
	"os"
	"strings"
	"sync"
	"time"
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

	fileScanner.Scan()
	instructions := []rune(fileScanner.Text())
	fileScanner.Scan()

	historyMap := map[string]string{}
	var startingKeys []string

	for fileScanner.Scan() {
		line := strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(fileScanner.Text(), "(", ""), ")", ""), ",", "")
		lineSplit := strings.Split(strings.ReplaceAll(line, "=", ""), " ")
		historyMap[lineSplit[0]] = strings.Join(lineSplit[1:4], " ")
		if lineSplit[0][2] == 'A' {
			startingKeys = append(startingKeys, lineSplit[0])
		}
	}

	var stepCounts []int
	steps := 1
	var idx int
	for {
		if len(instructions) == idx {
			idx = 0
		}

		// we are a ghost, so do simultaneously?
		wg := &sync.WaitGroup{}
		for cIdx, _ := range startingKeys {
			ch := make(chan string)
			wg.Add(1)
			go getNextStep(wg, historyMap[startingKeys[cIdx]], instructions[idx], ch)
			startingKeys[cIdx] = <-ch
			if startingKeys[cIdx][2] == 'Z' {
				stepCounts = append(stepCounts, steps)
			}
			close(ch)
		}

		wg.Wait()

		if len(stepCounts) == len(startingKeys) {
			break
		}

		steps++
		idx++
	}

	output = findCommonProduct(stepCounts)
	log.Printf("Value: %d\n", output)
}

func findCommonProduct(stepCounts []int) int {
	largest := stepCounts[len(stepCounts)-1]
	idx := 1
	for {
		largestProduct := largest * idx
		values := stepCounts[0 : len(stepCounts)-1]
		matches := 0
		for i := 0; i < len(values); i++ {
			if largestProduct%values[i] == 0 {
				matches++
			} else {
				break
			}
			if matches == len(values) {
				return largestProduct
			}
		}
		idx++
	}
}

func getNextStep(wg *sync.WaitGroup, stepString string, instruction rune, out chan string) {
	defer wg.Done()
	steps := strings.Fields(stepString)
	if instruction == 'L' {
		out <- steps[0]
		return
	}
	out <- steps[1]
}
