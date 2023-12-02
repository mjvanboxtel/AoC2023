package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	defer fmt.Printf("Execution time %s", time.Since(start))

	readFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer readFile.Close()
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	var idSum int
	for fileScanner.Scan() {
		gameData := strings.Split(fileScanner.Text(), ":")
		trials := strings.ReplaceAll(gameData[1], ";", "")
		trials = strings.ReplaceAll(trials, ",", "")[1:]
		minCounts := minCubeCounts(trials)
		idSum += minCounts["red"] * minCounts["green"] * minCounts["blue"]
	}
	fmt.Printf("Value: %d\n", idSum)
}

func minCubeCounts(gameData string) map[string]int {
	data := strings.Split(gameData, " ")
	highest := map[string]int{
		"red":   0,
		"green": 0,
		"blue":  0,
	}
	for i := 0; i < len(data)/2; i++ {
		actualIdx := i * 2
		value, _ := strconv.Atoi(data[actualIdx])
		colour := data[actualIdx+1]
		if highest[colour] < value {
			highest[colour] = value
		}
	}
	return highest
}
