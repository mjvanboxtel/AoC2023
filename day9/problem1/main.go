package main

import (
	"bufio"
	"log"
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

	for fileScanner.Scan() {
		var input []int
		values := strings.Fields(fileScanner.Text())
		for _, val := range values {
			iVal, _ := strconv.Atoi(val)
			input = append(input, iVal)
		}
		p := predict(input)
		output += p
	}

	log.Printf("Value: %d\n", output)
}

func predict(sequence []int) int {
	var out []int
	var zeros int
	for i := 0; i < len(sequence)-1; i++ {
		diff := sequence[i+1] - sequence[i]
		if diff == 0 {
			zeros++
		}
		out = append(out, diff)
	}
	var prediction int
	if zeros != len(out) {
		prediction = predict(out)
	} else {
		prediction = 0
	}
	out = append(out, prediction)
	sequence = append(sequence, out[len(out)-1]+sequence[len(sequence)-1])
	return sequence[len(sequence)-1]
}
