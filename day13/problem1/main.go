package main

import (
	"bufio"
	"log"
	"os"
	"reflect"
	"slices"
	"time"
)

const (
	multiplier = 100
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

	var row int
	var landscapes [][]rune
	for fileScanner.Scan() {
		line := fileScanner.Text()
		lineRunes := []rune(line)
		if line == "" {
			output += findReflection(landscapes)
			landscapes = [][]rune{}
			row = 0
			continue
		}
		landscapes = append(landscapes, []rune{})
		for i := 0; i < len(lineRunes); i++ {
			landscapes[row] = append(landscapes[row], lineRunes[i])
		}
		row++
	}
	output += findReflection(landscapes)

	log.Printf("Value: %d\n", output)
}

func findReflection(landscapes [][]rune) int {
	var rows int
	var reflectionPoints []int
	for i := 0; i < len(landscapes)-1; i++ {
		if reflect.DeepEqual(landscapes[i], landscapes[i+1]) {
			reflectionPoints = append(reflectionPoints, i)
		}
	}

	for j := 0; j < len(reflectionPoints); j++ {
		for i := 0; i < len(landscapes)-1; i++ {
			if reflectionPoints[j]+1+i == len(landscapes) {
				rows = len(landscapes) - i
				break
			}
			if !reflect.DeepEqual(landscapes[reflectionPoints[j]-i], landscapes[reflectionPoints[j]+1+i]) {
				rows = 0
				break
			}
			if reflectionPoints[j]-i == 0 {
				rows = i + 1
				break
			}
		}
		if rows != 0 {
			return rows * multiplier
		}
	}

	reflectionPoints = []int{}
	var columns int
	for i := 0; i < len(landscapes[0])-1; i++ {
		for j := 0; j < len(landscapes); j++ {
			if landscapes[j][i] != landscapes[j][i+1] {
				break
			}
			if j == len(landscapes)-1 {
				reflectionPoints = append(reflectionPoints, i)
			}
		}
	}

	for j := 0; j < len(reflectionPoints); j++ {
		for k := 0; k < len(landscapes); k++ {
			var begin []rune
			var end []rune
			if reflectionPoints[j] < (len(landscapes[k])-1)/2 {
				begin = append(begin, landscapes[k][0:reflectionPoints[j]+1]...)
				end = append(end, landscapes[k][reflectionPoints[j]+1:reflectionPoints[j]*2+2]...)
				slices.Reverse(begin)
				if !reflect.DeepEqual(begin, end) {
					columns = 0
					break
				}
			} else {
				begin = append(begin, landscapes[k][reflectionPoints[j]-(len(landscapes[k])-1-reflectionPoints[j])+1:reflectionPoints[j]+1]...)
				end = append(end, landscapes[k][reflectionPoints[j]+1:]...)
				slices.Reverse(begin)
				if !reflect.DeepEqual(begin, end) {
					columns = 0
					break
				}
			}

			if k == len(landscapes)-1 {
				columns = reflectionPoints[j] + 1
				break
			}
		}
		if columns != 0 {
			break
		}
	}
	return columns
}
