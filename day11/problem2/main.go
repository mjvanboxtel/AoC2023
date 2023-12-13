package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"time"
)

const expansionCount = 1000000

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

	var galaxies [][]int

	var universeWidth int
	var idx int
	for fileScanner.Scan() {
		lineRunes := []rune(fileScanner.Text())
		universeWidth = len(lineRunes)
		currentGalaxies := findGalaxies(lineRunes, idx)
		galaxies = append(galaxies, currentGalaxies...)
		idx++
	}

	galaxies = expandUniverse(galaxies, universeWidth, idx)

	var output int

	for i := 0; i < len(galaxies); i++ {
		for j := 0; j < len(galaxies); j++ {
			distance := findDistance(galaxies[i], galaxies[j])
			if distance == 0 {
				continue
			}
			output += distance
		}
	}

	if output%2 == 0 {
		output = output / 2
	} else {
		output = int(math.Ceil(float64(output / 2)))
	}

	log.Printf("Value: %d\n", output)
}

func findDistance(g1 []int, g2 []int) int {
	dY := math.Abs(float64(g1[0]) - float64(g2[0]))
	dX := math.Abs(float64(g1[1]) - float64(g2[1]))
	return int(dY + dX)
}

func findGalaxies(line []rune, row int) [][]int {
	var galaxies [][]int
	for i := 0; i < len(line); i++ {
		if line[i] == '#' {
			galaxies = append(galaxies, []int{i, row})
		}
	}
	return galaxies
}

func expandUniverse(galaxies [][]int, initialWidth int, initialLength int) [][]int {
	rowGalaxies := make([]bool, initialWidth)
	columnGalaxies := make([]bool, initialLength)

	for i := 0; i < len(galaxies); i++ {
		rowGalaxies[galaxies[i][1]] = true
		columnGalaxies[galaxies[i][0]] = true
	}

	for i := 0; i < len(galaxies); i++ {
		rowIncrease := getExpansionCount(columnGalaxies, galaxies[i][0]) * (expansionCount - 1)
		colIncrease := getExpansionCount(rowGalaxies, galaxies[i][1]) * (expansionCount - 1)
		galaxies[i][0] += rowIncrease
		galaxies[i][1] += colIncrease
	}

	return galaxies
}

func getExpansionCount(galaxyExist []bool, end int) int {
	var out int
	for i := 0; i < end; i++ {
		if galaxyExist[i] == false {
			out++
		}
	}
	return out
}
