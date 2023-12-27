package main

import (
	"bufio"
	"log"
	"os"
	"strings"
	"time"
)

const (
	multiplier = 17
	divisor    = 256
)

func main() {
	start := time.Now()
	defer log.Printf("Execution time %s", time.Since(start))

	readFile, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer readFile.Close()
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	var output int

	fileScanner.Scan()
	line := strings.Split(fileScanner.Text(), ",")
	for i := 0; i < len(line); i++ {
		output += hash(line[i])
	}

	log.Printf("Value: %d\n", output)
}

func hash(s string) int {
	var val int
	for i := 0; i < len(s); i++ {
		val += int(s[i])
		val *= multiplier
		val = val % divisor
	}
	return val
}
