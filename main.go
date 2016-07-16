package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

const size = 50

var displayChars = make(map[bool]string)

type cell struct {
	populated  bool
	neighbours int
}

func newGrid() *[size][size]cell {
	grid := [size][size]cell{}
	rand.Seed(time.Now().UTC().UnixNano())
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			grid[i][j].populated = rand.Float32() > 0.5
		}
	}
	return &grid
}

func checkCoordinates(x int, y int, grid [size][size]cell) int {
	if x > 0 && y > 0 && x < (size-1) && y < (size-1) && grid[x][y].populated {
		return 1
	}
	return 0
}

func countNeighbours(x int, y int, grid [size][size]cell) int {
	count := 0
	xCo := []int{-1, 1, 0, 0, -1, -1, 1, 1}
	yCo := []int{0, 0, -1, 1, -1, 1, -1, 1}
	for i := range xCo {
		count += checkCoordinates(x+xCo[i], y+yCo[i], grid)
	}
	return count
}

func updateNeighbours(grid *[size][size]cell) {
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			grid[i][j].neighbours = countNeighbours(i, j, *grid)
		}
	}
}

func updateGrid(grid *[size][size]cell, c chan string) {
	str := ""
	for i := 0; i < size; i++ {
		for j, v := range grid[i] {
			n := v.neighbours
			if n < 2 || n >= 4 {
				v.populated = false
			} else if n == 3 {
				v.populated = true
			}
			grid[i][j] = v
			str += displayChars[v.populated]
		}
		str += "\n"
	}
	c <- str
}

func clearConsole() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func run(grid *[size][size]cell, c chan string) {
	for {
		clearConsole()
		updateNeighbours(grid)
		updateGrid(grid, c)
		time.Sleep(300 * time.Millisecond)
	}
}

func main() {
	displayChars[true], displayChars[false] = " # ", "   "
	c := make(chan string)
	go run(newGrid(), c)

	for {
		select {
		case v := <-c:
			fmt.Println(v)
		}
	}
}
