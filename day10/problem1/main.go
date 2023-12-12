package main

import (
	"bufio"
	"log"
	"os"
	"strings"
	"time"
)

const (
	vertical         = '|'
	horizontal       = '-'
	nintyDegBendNE   = 'L'
	nintyDegBendNW   = 'J'
	nintyDegBendSW   = '7'
	nintyDegBendSE   = 'F'
	startingPosition = 'S'
	headingSouth     = "South"
	headingNorth     = "North"
	headingEast      = "East"
	headingWest      = "West"
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

	var pipeMap [][]rune
	startingCoordinates := make([]int, 2)

	scanIdx := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		if strings.ContainsRune(line, startingPosition) {
			startingCoordinates[1] = strings.IndexRune(line, startingPosition)
			startingCoordinates[0] = scanIdx
		}
		pipeMap = append(pipeMap, []rune(fileScanner.Text()))
		scanIdx++
	}

	idx := 0
	var heading string
	for {
		startingCoordinates, heading = navigate(startingCoordinates, pipeMap, heading)
		if pipeMap[startingCoordinates[0]][startingCoordinates[1]] == startingPosition {
			if idx%2 == 0 {
				output = idx / 2
			} else {
				output = idx/2 + 1
			}
			break
		}
		idx++
	}

	log.Printf("Value: %d\n", output)
}

func navigate(position []int, pipeMap [][]rune, heading string) ([]int, string) {
	var nextPosition []int
	firstRune := pipeMap[position[0]][position[1]]
	if position[0] == 0 {
		if pathSouth(firstRune, pipeMap, position, heading) {
			return []int{position[0] + 1, position[1]}, headingSouth
		}
		if position[1] == 0 {
			// look around origin
			if pathEast(firstRune, pipeMap, position, heading) {
				return []int{position[0], position[1] + 1}, headingEast
			}
		} else if position[1] == len(pipeMap[0])-1 {
			// look around top right
			if pathWest(firstRune, pipeMap, position, heading) {
				return []int{position[0], position[1] - 1}, headingWest
			}
		} else {
			// look around top edge
			if pathEast(firstRune, pipeMap, position, heading) {
				return []int{position[0], position[1] + 1}, headingEast
			} else if pathWest(firstRune, pipeMap, position, heading) {
				return []int{position[0], position[1] - 1}, headingWest
			}
		}
	} else if position[0] == len(pipeMap)-1 {
		// look around bottom
		if pathNorth(firstRune, pipeMap, position, heading) {
			return []int{position[0] - 1, position[1]}, headingNorth
		}
		if position[1] == 0 {
			// look around bottom left
			if pathEast(firstRune, pipeMap, position, heading) {
				return []int{position[0], position[1] + 1}, headingEast
			}
		} else if position[1] == len(pipeMap[0])-1 {
			// look around bottom right
			if pathWest(firstRune, pipeMap, position, heading) {
				return []int{position[0], position[1] - 1}, headingWest
			}
		} else {
			// look around bottom edge
			if pathEast(firstRune, pipeMap, position, heading) {
				return []int{position[0], position[1] + 1}, headingEast
			} else if pathWest(firstRune, pipeMap, position, heading) {
				return []int{position[0], position[1] - 1}, headingWest
			}
		}
	} else if position[1] == 0 {
		// look around left edge
		if pathEast(firstRune, pipeMap, position, heading) {
			return []int{position[0], position[1] + 1}, headingEast
		} else if pathNorth(firstRune, pipeMap, position, heading) {
			return []int{position[0] - 1, position[1]}, headingNorth
		} else if pathSouth(firstRune, pipeMap, position, heading) {
			return []int{position[0] + 1, position[1]}, headingSouth
		}
	} else if position[1] == len(pipeMap[0])-1 {
		// look around right edge
		if pathWest(firstRune, pipeMap, position, heading) {
			return []int{position[0], position[1] - 1}, headingWest
		} else if pathSouth(firstRune, pipeMap, position, heading) {
			return []int{position[0] + 1, position[1]}, headingSouth
		} else if pathNorth(firstRune, pipeMap, position, heading) {
			return []int{position[0] - 1, position[1]}, headingNorth
		}
	} else {
		// look around node
		if pathNorth(firstRune, pipeMap, position, heading) {
			return []int{position[0] - 1, position[1]}, headingNorth
		} else if pathSouth(firstRune, pipeMap, position, heading) {
			return []int{position[0] + 1, position[1]}, headingSouth
		} else if pathWest(firstRune, pipeMap, position, heading) {
			return []int{position[0], position[1] - 1}, headingWest
		} else if pathEast(firstRune, pipeMap, position, heading) {
			return []int{position[0], position[1] + 1}, headingEast
		}
	}
	return nextPosition, ""
}

func pathWest(first rune, pipeMap [][]rune, position []int, heading string) bool {
	if heading == headingEast {
		return false
	}
	if first == horizontal || first == nintyDegBendNW || first == nintyDegBendSW || first == startingPosition {
		second := pipeMap[position[0]][position[1]-1]
		if second == horizontal || second == nintyDegBendNE || second == nintyDegBendSE || second == startingPosition {
			return true
		}
	}
	return false
}

func pathEast(first rune, pipeMap [][]rune, position []int, heading string) bool {
	if heading == headingWest {
		return false
	}
	if first == horizontal || first == nintyDegBendNE || first == nintyDegBendSE || first == startingPosition {
		second := pipeMap[position[0]][position[1]+1]
		if second == horizontal || second == nintyDegBendNW || second == nintyDegBendSW || second == startingPosition {
			return true
		}
	}
	return false
}

func pathNorth(first rune, pipeMap [][]rune, position []int, heading string) bool {
	if heading == headingSouth {
		return false
	}
	if first == vertical || first == nintyDegBendNE || first == nintyDegBendNW || first == startingPosition {
		second := pipeMap[position[0]-1][position[1]]
		if second == vertical || second == nintyDegBendSW || second == nintyDegBendSE || second == startingPosition {
			return true
		}
	}
	return false
}

func pathSouth(first rune, pipeMap [][]rune, position []int, heading string) bool {
	if heading == headingNorth {
		return false
	}
	if first == vertical || first == nintyDegBendSW || first == nintyDegBendSE || first == startingPosition {
		second := pipeMap[position[0]+1][position[1]]
		if second == vertical || second == nintyDegBendNE || second == nintyDegBendNW || second == startingPosition {
			return true
		}
	}
	return false
}
