package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
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
	raceTime := strings.Join(strings.Fields(fileScanner.Text())[1:], "")
	fileScanner.Scan()
	distance := strings.Join(strings.Fields(fileScanner.Text())[1:], "")

	output = 1
	tI, _ := strconv.Atoi(raceTime)
	dI, _ := strconv.Atoi(distance)
	first, second := getStartingPoint(tI)
	output *= findBreakCount(first, second, dI)

	log.Printf("Value: %d\n", output)
}

func getStartingPoint(t int) (int, int) {
	if t%2 == 0 {
		return t / 2, t / 2
	}
	return int(math.Ceil(float64(t) / 2)), t / 2
}

func findBreakCount(first int, second int, distance int) int {
	var count int
	start := first
	for i := 1; i < start; i++ {
		if first*second > distance {
			count++
			first += 1
			second -= 1
		} else {
			break
		}
	}
	if start%2 != 0 {
		return count*2 - 1
	}
	return count * 2
}
