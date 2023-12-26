package main

import (
	"bufio"
	"log"
	"os"
	"slices"
	"time"
)

const (
	cubeRock  = '#'
	roundRock = 'O'
	ground    = '.'
	cycles    = 1000000000
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

	entryPoint, platform := findCycle(platform)

	for i := 0; i < entryPoint-((cycles-entryPoint)%(entryPoint+1)); i++ {
		platform = runSpinCycle(platform)
	}

	output += scorePlatform(platform)

	log.Printf("Value: %d\n", output)
}

func copyPlatform(platform [][]rune) [][]rune {
	rows := make([][]rune, len(platform))

	for i := range platform {
		rows[i] = make([]rune, len(platform[i]))
		copy(rows[i], platform[i])
	}
	return rows
}

func findCycle(platform [][]rune) (int, [][]rune) {
	slow := copyPlatform(platform)
	fast := copyPlatform(platform)

	var idx int
	for {
		slow = runSpinCycle(slow)
		fast = runSpinCycle(fast)
		fast = runSpinCycle(fast)
		if platformsEqual(fast, slow) {
			break
		}
		if idx == cycles {
			return 0, nil
		}
		idx++
	}

	slow = copyPlatform(platform)

	var entryPoint int
	for {
		slow = runSpinCycle(slow)
		fast = runSpinCycle(fast)
		fast = runSpinCycle(fast)
		if platformsEqual(slow, fast) {
			break
		}
		entryPoint++
	}
	return entryPoint, slow
}

func platformsEqual(p1 [][]rune, p2 [][]rune) bool {
	for i := 0; i < len(p1); i++ {
		if !slices.Equal(p1[i], p2[i]) {
			return false
		}
	}
	return true
}

func runSpinCycle(platform [][]rune) [][]rune {
	return tiltEast(tiltSouth(tiltWest(tiltNorth(platform))))
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

func tiltSouth(platform [][]rune) [][]rune {
	out := make([][]rune, len(platform))
	copy(out, platform)
	for i := 0; i < len(platform[0]); i++ {
		lastValid := len(platform) - 1
		for j := len(platform) - 1; j >= 0; j-- {
			if platform[j][i] == cubeRock {
				lastValid = j - 1
				continue
			}
			if platform[j][i] == roundRock {
				out[j][i] = ground
				out[lastValid][i] = roundRock
				lastValid = lastValid - 1
			}
			if lastValid == 0 {
				break
			}
		}
	}
	return out
}

func tiltEast(platform [][]rune) [][]rune {
	out := make([][]rune, len(platform))
	copy(out, platform)
	for i := 0; i < len(platform); i++ {
		lastValid := len(platform[i]) - 1
		for j := len(platform[i]) - 1; j >= 0; j-- {
			if platform[i][j] == cubeRock {
				lastValid = j - 1
				continue
			}
			if platform[i][j] == roundRock {
				out[i][j] = ground
				out[i][lastValid] = roundRock
				lastValid = lastValid - 1
			}
			if lastValid == 0 {
				break
			}
		}
	}
	return out
}

func tiltWest(platform [][]rune) [][]rune {
	out := make([][]rune, len(platform))
	copy(out, platform)
	for i := 0; i < len(platform); i++ {
		var lastValid int
		for j := 0; j < len(platform[i]); j++ {
			if platform[i][j] == cubeRock {
				lastValid = j + 1
				continue
			}
			if platform[i][j] == roundRock {
				out[i][j] = ground
				out[i][lastValid] = roundRock
				lastValid = lastValid + 1
			}
			if lastValid == len(platform[i])-1 {
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
