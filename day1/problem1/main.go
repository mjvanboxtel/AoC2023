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
		re := regexp.MustCompile("[0-9]+")
		line := strings.Join(re.FindAllString(fileScanner.Text(), -1), "")
		val, err := strconv.Atoi(string(line[0]) + string(line[len(line)-1]))
		if err != nil {
			panic(err)
		}
		calibrationValue += val
	}

	log.Printf("Value: %d\n", calibrationValue)
}
