package main

import (
	"bufio"
	"log"
	"os"
	"strings"
	"time"
)

const (
	startNode = "AAA"
	endNode   = "ZZZ"
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

	for fileScanner.Scan() {
		line := strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(fileScanner.Text(), "(", ""), ")", ""), ",", "")
		lineSplit := strings.Split(strings.ReplaceAll(line, "=", ""), " ")
		historyMap[lineSplit[0]] = strings.Join(lineSplit[1:4], " ")
	}

	steps := 1
	var idx int
	nextKey := startNode
	for {
		if len(instructions) == idx {
			idx = 0
		}
		nextKey = getNextStep(historyMap[nextKey], instructions[idx])
		if nextKey == endNode {
			break
		}
		steps++
		idx++
	}

	output = steps
	log.Printf("Value: %d\n", output)
}

func getNextStep(stepString string, instruction rune) string {
	steps := strings.Fields(stepString)
	if instruction == 'L' {
		return steps[0]
	}
	return steps[1]
}
