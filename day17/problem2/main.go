package main

import (
	"bufio"
	"container/heap"
	"log"
	"os"
	"strconv"
	"time"
)

const (
	maxSteps = 10
	minSteps = 4
)

type Vector2 struct {
	x int
	y int
}

type Vertex struct {
	position  Vector2
	heading   Vector2
	heat      int
	previousX int
	previousY int
}

type Item struct {
	priority int
	index    int
	vertex   Vertex
}

type MaxHeap []*Item

func (mh *MaxHeap) Len() int { return len(*mh) }

func (mh *MaxHeap) Less(i, j int) bool {
	val := *mh
	return val[i].priority < val[j].priority
}

func (mh *MaxHeap) Swap(i, j int) {
	val := *mh
	val[i], val[j] = val[j], val[i]
	val[i].index = i
	val[j].index = j
}

func (mh *MaxHeap) Push(x any) {
	n := len(*mh)
	v := x.(*Item)
	v.index = n
	*mh = append(*mh, v)
}

func (mh *MaxHeap) Pop() any {
	old := *mh
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*mh = old[0 : n-1]
	return item
}

func (mh *MaxHeap) update(n *Item, priority int) {
	n.priority = priority
	heap.Fix(mh, n.index)
}

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
	var distances MaxHeap
	var grid [][]Vertex

	var row int
	for fileScanner.Scan() {
		grid = append(grid, make([]Vertex, 0))
		line := fileScanner.Text()
		for i := 0; i < len(line); i++ {
			heat, _ := strconv.Atoi(string(line[i]))
			grid[row] = append(grid[row], Vertex{
				position: Vector2{
					x: i,
					y: row,
				},
				heat: heat,
			})
		}
		row++
	}

	shortest := findShortestPath(grid, distances)
	output += shortest

	log.Printf("Value: %d\n", output)
}

func findShortestPath(grid [][]Vertex, mh MaxHeap) int {
	seen := map[Vertex]bool{}
	current := &Item{
		priority: 0,
		vertex:   grid[0][0],
	}
	current.vertex.heading = Vector2{
		x: 2,
		y: 2,
	}

	for {
		var adjacentVertices []*Vertex
		headings := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
		for i := 0; i < len(headings); i++ {
			if (current.vertex.position.y+headings[i][1] >= 0 && current.vertex.position.x+headings[i][0] >= 0) && (current.vertex.position.y+headings[i][1] < len(grid) && current.vertex.position.x+headings[i][0] < len(grid[0])) {
				adjacentVertices = append(adjacentVertices, &grid[current.vertex.position.y+headings[i][1]][current.vertex.position.x+headings[i][0]])
			}
		}
		for i := 0; i < len(adjacentVertices); i++ {
			last := Vector2{
				x: current.vertex.position.x - current.vertex.heading.x,
				y: current.vertex.position.y - current.vertex.heading.y,
			}
			if adjacentVertices[i].position != last {
				if current.vertex.previousX == maxSteps && adjacentVertices[i].position.y == current.vertex.position.y || current.vertex.previousY == maxSteps && adjacentVertices[i].position.x == current.vertex.position.x {
					continue
				}
				if (current.vertex.previousX > 0 && current.vertex.previousX < minSteps && adjacentVertices[i].position.x == current.vertex.position.x) || (current.vertex.previousY > 0 && current.vertex.previousY < minSteps && adjacentVertices[i].position.y == current.vertex.position.y) {
					continue
				}
				priority := current.priority + adjacentVertices[i].heat
				item := &Item{
					priority: priority,
					vertex: Vertex{
						position: Vector2{
							x: adjacentVertices[i].position.x,
							y: adjacentVertices[i].position.y,
						},
						heading: Vector2{
							x: adjacentVertices[i].position.x - current.vertex.position.x,
							y: adjacentVertices[i].position.y - current.vertex.position.y,
						},
						heat: adjacentVertices[i].heat,
					},
				}
				if adjacentVertices[i].position.x == current.vertex.position.x {
					item.vertex.previousY = current.vertex.previousY + 1
				} else if adjacentVertices[i].position.y == current.vertex.position.y {
					item.vertex.previousX = current.vertex.previousX + 1
				}
				if _, ok := seen[item.vertex]; !ok {
					heap.Push(&mh, item)
					seen[item.vertex] = true
					mh.update(item, priority)
				}
			}
		}
		if current.vertex.position.x == len(grid[0])-1 && current.vertex.position.y == len(grid)-1 && (current.vertex.previousY >= minSteps || current.vertex.previousX >= minSteps) {
			break
		}
		current = heap.Pop(&mh).(*Item)
	}
	return current.priority
}
