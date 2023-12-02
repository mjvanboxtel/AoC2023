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
		gameId, _ := strconv.Atoi(strings.Split(gameData[0], " ")[1])
		trials := strings.Split(gameData[1], ";")
		failed := false
		for _, trial := range trials {
			if failedTrial(trial) {
				failed = true
				break
			}
		}
		if !failed {
			idSum += gameId
		}
		fmt.Printf("game: %s--%t\n", trials, failed)
	}
	fmt.Printf("Value: %d\n", idSum)
}

func failedTrial(trial string) bool {
	limits := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}
	trial = strings.ReplaceAll(trial, ",", "")
	data := strings.Split(trial, " ")[1:]
	for i := 0; i < len(data)/2; i++ {
		actualIdx := i * 2
		value, _ := strconv.Atoi(data[actualIdx])
		colour := data[actualIdx+1]
		if value > limits[colour] {
			return true
		}
	}
	return false
}
