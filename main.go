package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

const size = 32

type Cell struct {
	populated  bool
	neighbours int
}

func setupGrid(grid *[size][size]Cell) {
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

func checkCoordinates(x int, y int, grid [size][size]Cell) int {
	if x > 0 && y > 0 && x < (size-1) && y < (size-1) && grid[x][y].populated {
		return 1
	}
	return 0
}

func countNeighbours(x int, y int, grid [size][size]Cell) int {
	count := 0
	xCo := []int{-1, 1, 0, 0, -1, -1, 1, 1}
	yCo := []int{0, 0, -1, 1, -1, 1, -1, 1}
	for i, _ := range xCo {
		count += checkCoordinates(x+xCo[i], y+yCo[i], grid)
	}
	return count
}

func updateNeighbours(grid *[size][size]Cell) {
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			grid[i][j].neighbours = countNeighbours(i, j, *grid)
		}
	}
}

func updateGrid(grid *[size][size]Cell) {
	for i := 0; i < size; i++ {
		for j, v := range grid[i] {
			n := v.neighbours
			switch {
			case n == 0 || n == 1 || n == 4 || n == 5 || n == 6 || n == 7 || n == 8:
				v.populated = false
			case n == 3:
				v.populated = true
			case n == 2:
				v.populated = v.populated
			}
			grid[i][j] = v
			if v.populated {
				fmt.Printf("%v ", "#")
			} else {
				fmt.Printf("%v ", "_")
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
	grid := [size][size]Cell{}
	setupGrid(&grid)

	for i := 0; i < 1000; i++ {
		clearConsole()
		updateNeighbours(&grid)
		updateGrid(&grid)
		time.Sleep(500 * time.Millisecond)
	}
}
