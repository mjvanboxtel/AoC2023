package main

import (
	"bufio"
	"log"
	"os"
	"time"
)

const (
	cubeRock  = '#'
	roundRock = 'O'
	ground    = '.'
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

	var idx int
	var platform [][]rune
	for fileScanner.Scan() {
		line := fileScanner.Text()
		if line == "" {
			break
		}
		platform = append(platform, []rune(line))
		idx++
	}

	output += scorePlatform(tiltNorth(platform))

	log.Printf("Value: %d\n", output)
}

func tiltNorth(platform [][]rune) [][]rune {
	out := make([][]rune, len(platform))
	copy(out, platform)
	for i := 0; i < len(platform[0]); i++ {
		var lastValid int
		for j := 0; j < len(platform); j++ {
			if platform[j][i] == cubeRock {
				lastValid = j + 1
				continue
			}
			if platform[j][i] == roundRock {
				out[j][i] = ground
				out[lastValid][i] = roundRock
				lastValid = lastValid + 1
			}
			if lastValid == len(platform)-1 {
				break
			}
		}
	}
	return out
}

func scorePlatform(platform [][]rune) int {
	var score int
	for i := 0; i < len(platform); i++ {
		var count int
		for j := 0; j < len(platform[i]); j++ {
			if platform[i][j] == roundRock {
				count++
			}
		}
		score += count * (len(platform) - i)
	}
	return score
}
