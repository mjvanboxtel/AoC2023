package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
	"unicode"
)

type sourceDestMap struct {
	header  string
	entries []entry
}

type entry struct {
	sourceRange []int
	destRange   []int
}

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
	seedLine := strings.Split(fileScanner.Text(), ":")
	seeds := strings.Fields(seedLine[1])

	var sourceRanges [][]int
	sourceRanges = getSeedRanges(seeds, sourceRanges)

	almanac := mapSourceDestRanges(fileScanner)

	for _, sDR := range almanac {
		var newRanges [][]int
		for _, sRange := range sourceRanges {
			newRanges = append(newRanges, toDestRanges(sDR, sRange)...)
		}
		sourceRanges = newRanges
	}

	lowest := math.MaxInt64
	for i := 0; i < len(sourceRanges); i++ {
		if sourceRanges[i][0] < lowest {
			lowest = sourceRanges[i][0]
		}
	}

	output = lowest
	log.Printf("Value: %d\n", output)
}

func toDestRanges(sDM sourceDestMap, sourceRange []int) [][]int {
	var out [][]int
	for _, e := range sDM.entries {
		if rangesOverlap(sourceRange, e.sourceRange) {
			// the source range is encapsulated
			if sourceRange[0] >= e.sourceRange[0] && sourceRange[1] <= e.sourceRange[1] {
				sRange := []int{sourceRange[0] + (e.destRange[0] - e.sourceRange[0]), sourceRange[1] + (e.destRange[1] - e.sourceRange[1])}
				out = append(out, sRange)
				sourceRange[0] = sourceRange[1]
				continue
			}
			// the higher end of the source range overlaps
			if sourceRange[0] < e.sourceRange[0] {
				out = append(out, []int{e.destRange[0], sourceRange[1] + (e.destRange[1] - e.sourceRange[1])})
				sourceRange[1] -= sourceRange[1] - e.sourceRange[0] + 1
			}
			// the lower end of the source range overlaps
			if sourceRange[1] > e.sourceRange[1] {
				out = append(out, []int{sourceRange[0] - (e.sourceRange[0] - e.destRange[0]), e.destRange[0] + sourceRange[1] - e.sourceRange[1] - 1})
				sourceRange[0] = e.sourceRange[1] + 1
			}
		}
	}
	if sourceRange[0] != sourceRange[1] && sDM.header == "seed-to-soil" {
		out = append(out, sourceRange)
	}
	return out
}

func mapSourceDestRanges(fs *bufio.Scanner) []sourceDestMap {
	var out []sourceDestMap
	sDM := sourceDestMap{}
	fs.Scan()
	for fs.Scan() {
		line := fs.Text()
		lineRunes := []rune(line)
		if len(lineRunes) == 0 {
			out = append(out, sDM)
			sDM = sourceDestMap{}
		} else if unicode.IsNumber(lineRunes[0]) {
			lineVals := strings.Fields(line)
			source, _ := strconv.Atoi(lineVals[1])
			destination, _ := strconv.Atoi(lineVals[0])
			rangeVal, _ := strconv.Atoi(lineVals[2])
			e := entry{
				sourceRange: []int{source, source + rangeVal},
				destRange:   []int{destination, destination + rangeVal},
			}
			sDM.entries = append(sDM.entries, e)
		} else {
			lineVals := strings.Fields(line)
			sDM.header = lineVals[0]
		}
	}
	out = append(out, sDM)
	return out
}

func rangesOverlap(range1 []int, range2 []int) bool {
	if range1[0] <= range2[len(range2)-1] && range1[len(range1)-1] >= range2[0] {
		return true
	}
	return false
}

func getSeedRanges(seeds []string, seedRanges [][]int) [][]int {
	var startSeed int
	for idx, seed := range seeds {
		val, _ := strconv.Atoi(seed)
		if (idx+1)%2 != 0 {
			startSeed = val
			continue
		}
		seedRanges = append(seedRanges, []int{startSeed, startSeed + val})
	}
	return seedRanges
}
