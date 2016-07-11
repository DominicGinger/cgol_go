package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

const size = 50

type cell struct {
	populated  bool
	neighbours int
}

func setupGrid(grid *[size][size]cell) {
	rand.Seed(time.Now().UTC().UnixNano())
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if rand.Float32() > 0.5 {
				grid[i][j].populated = true
			} else {
				grid[i][j].populated = false
			}
		}
	}
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

func updateGrid(grid *[size][size]cell) {
	for i := 0; i < size; i++ {
		for j, v := range grid[i] {
			n := v.neighbours
			switch {
			case n == 0 || n == 1 || n >= 4:
				v.populated = false
			case n == 3:
				v.populated = true
			}
			grid[i][j] = v
			if v.populated {
				fmt.Printf("%v ", "#")
			} else {
				fmt.Printf("%v ", " ")
			}
		}
		fmt.Println()
	}
}

func clearConsole() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {
	grid := [size][size]cell{}
	setupGrid(&grid)

	for i := 0; i < 1000; i++ {
		clearConsole()
		updateNeighbours(&grid)
		updateGrid(&grid)
		time.Sleep(500 * time.Millisecond)
	}
}
