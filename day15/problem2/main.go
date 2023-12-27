package main

import (
	"bufio"
	"log"
	"os"
	"slices"
	"strconv"
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

	boxMap := map[int][]string{}

	fileScanner.Scan()
	line := strings.Split(fileScanner.Text(), ",")
	for i := 0; i < len(line); i++ {
		labelBox(line[i], boxMap)
	}

	output += focusingPower(boxMap)

	log.Printf("Value: %d\n", output)
}

func focusingPower(boxMap map[int][]string) int {
	var power int
	for i, labels := range boxMap {
		boxPower := i + 1
		for j := 0; j < len(labels); j++ {
			slotPower := j + 1
			focalLength, _ := strconv.Atoi(strings.Split(labels[j], " ")[1])
			power += boxPower * slotPower * focalLength
		}
	}
	return power
}

func labelBox(label string, boxMap map[int][]string) map[int][]string {
	if strings.Contains(label, "-") {
		labelStrings := strings.Split(label, "-")
		box := hash(labelStrings[0])
		for i := 0; i < len(boxMap[box]); i++ {
			if strings.Split(boxMap[box][i], " ")[0] == labelStrings[0] { // buggy af
				boxMap[box] = slices.Delete(boxMap[box], i, i+1)
			}
		}
	} else if strings.Contains(label, "=") {
		labelStrings := strings.Split(label, "=")
		box := hash(labelStrings[0])
		for i := 0; i < len(boxMap[box]); i++ {
			if strings.Split(boxMap[box][i], " ")[0] == labelStrings[0] {
				boxMap[box][i] = strings.Join(labelStrings, " ")
				return boxMap
			}
		}
		boxMap[box] = append(boxMap[box], strings.Join(labelStrings, " "))
	}
	return boxMap
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
