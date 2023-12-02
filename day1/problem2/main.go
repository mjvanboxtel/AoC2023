package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	defer log.Printf("Execution time %s", time.Since(start))

	readFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer readFile.Close()
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	var calibrationValue int
	for fileScanner.Scan() {
		final := matchFirstAndLast(fileScanner.Text())
		value, err := strconv.Atoi(final)
		if err != nil {
			panic(err)
		}
		calibrationValue += value
	}

	log.Printf("Value: %d\n", calibrationValue)
}

func matchFirstAndLast(s string) string {
	re := regexp.MustCompile("(one)|(1)|(two)|(2)|(three)|(3)|(four)|(4)|(five)|(5)|(six)|(6)|(seven)|(7)|(eight)|(8)|(nine)|(9)")
	matches := []string{}
	idx := 0
	for {
		m := re.FindStringIndex(s[idx:])
		if m == nil {
			break
		}
		nextIndex := idx + m[0]
		matches = append(matches, s[idx:][m[0]:m[1]])
		idx = nextIndex + 1
	}
	return strings.Join([]string{mapNumber(matches[0]), mapNumber(matches[len(matches)-1])}, "")
}

func mapNumber(n string) string {
	numberMap := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}
	val, ok := numberMap[n]
	if ok {
		return val
	}
	return n
}
