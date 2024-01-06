package main

import (
	"bufio"
	"log"
	"os"
	"sync"
	"time"
)

const (
	vertical   = '|'
	horizontal = '-'
	rDiagonal  = '/'
	lDiagonal  = '\\'
	edge       = 'E'
	right      = "right"
	left       = "left"
	up         = "up"
	down       = "down"
)

type Vector2 struct {
	x    int
	y    int
	name string
}

type Tile struct {
	energized  bool
	identifier rune
	position   Vector2
}

var wg sync.WaitGroup

func main() {
	start := time.Now()
	defer func() { log.Println(time.Since(start)) }()

	readFile, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer readFile.Close()
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	var output int
	var grid [][]Tile

	var row int
	for fileScanner.Scan() {
		line := []rune(fileScanner.Text())
		if row == 0 {
			grid = createEdge(grid, len(line)+2)
			row++
		}
		grid = append(grid, []Tile{})
		grid[row] = append(grid[row], Tile{
			energized:  false,
			identifier: edge,
		})
		for i := 0; i < len(line); i++ {
			grid[row] = append(grid[row], Tile{
				energized:  false,
				identifier: line[i],
				position: Vector2{
					x: i + 1,
					y: row,
				},
			})
		}
		grid[row] = append(grid[row], Tile{
			energized:  false,
			identifier: edge,
		})
		row++
	}
	grid = createEdge(grid, len(grid[0]))

	wg.Add(1)
	go energize(grid, &grid[1][1], Vector2{
		x:    1,
		y:    0,
		name: right,
	})

	wg.Wait()

	output = scoreGrid(grid)

	log.Printf("Value: %d\n", output)
}

func scoreGrid(grid [][]Tile) int {
	var score int
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j].energized == true {
				score++
			}
		}
	}
	return score
}

func energize(grid [][]Tile, start *Tile, initialHeading Vector2) [][]Tile {
	defer wg.Done()
	next := start
	heading := initialHeading
	for {
		grid[next.position.y][next.position.x].energized = true
		heading = navigate(grid, next, heading)
		next = &grid[next.position.y+heading.y][next.position.x+heading.x]
		if next.identifier == edge || (next.identifier == horizontal && next.energized || next.identifier == vertical && next.energized) {
			break
		}
	}
	return grid
}

func navigate(grid [][]Tile, tile *Tile, heading Vector2) Vector2 {
	if tile.identifier == vertical && (heading.name == right || heading.name == left) {
		wg.Add(1)
		go energize(grid, tile, Vector2{
			x:    0,
			y:    -1,
			name: up,
		})
		return Vector2{
			x:    0,
			y:    1,
			name: down,
		}
	}
	if tile.identifier == horizontal && (heading.name == up || heading.name == down) {
		wg.Add(1)
		go energize(grid, tile, Vector2{
			x:    -1,
			y:    0,
			name: left,
		})
		return Vector2{
			x:    1,
			y:    0,
			name: right,
		}
	}
	if tile.identifier == rDiagonal {
		if heading.name == right {
			return Vector2{
				x:    0,
				y:    -1,
				name: up,
			}
		} else if heading.name == left {
			return Vector2{
				x:    0,
				y:    1,
				name: down,
			}
		} else if heading.name == up {
			return Vector2{
				x:    1,
				y:    0,
				name: right,
			}
		} else if heading.name == down {
			return Vector2{
				x:    -1,
				y:    0,
				name: left,
			}
		}
	}
	if tile.identifier == lDiagonal {
		if heading.name == right {
			return Vector2{
				x:    0,
				y:    1,
				name: down,
			}
		} else if heading.name == left {
			return Vector2{
				x:    0,
				y:    -1,
				name: up,
			}
		} else if heading.name == down {
			return Vector2{
				x:    1,
				y:    0,
				name: right,
			}
		} else if heading.name == up {
			return Vector2{
				x:    -1,
				y:    0,
				name: left,
			}
		}
	}
	return heading
}

func createEdge(grid [][]Tile, l int) [][]Tile {
	grid = append(grid, []Tile{})
	for i := 0; i < l; i++ {
		grid[len(grid)-1] = append(grid[len(grid)-1], Tile{
			energized:  false,
			identifier: edge,
		})
	}
	return grid
}
