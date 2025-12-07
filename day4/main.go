package main

import (
	"bufio"
	"fmt"
	"os"
)

const part2 = true

func hasRoll(grid [][]bool, r, c int) bool {
	if r < 0 || r >= len(grid) {
		return false
	}
	row := grid[r]
	if c < 0 || c >= len(row) {
		return false
	}
	return grid[r][c]
}

func main() {
	grid := make([][]bool, 0)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		row := make([]bool, len(line))
		for c, ch := range line {
			row[c] = (ch == '@')
		}
		grid = append(grid, row)
	}
	totalRolls := 0
	for {
		numRemoved := 0
		for r := 0; r < len(grid); r++ {
			for c := 0; c < len(grid[0]); c++ {
				if !hasRoll(grid, r, c) {
					continue
				}
				neighbors := -1
				for dR := -1; dR <= 1; dR++ {
					for dC := -1; dC <= 1; dC++ {
						if hasRoll(grid, r+dR, c+dC) {
							neighbors++
						}
					}
				}
				if neighbors < 4 {
					numRemoved++
					if part2 {
						grid[r][c] = false
					}
				}
			}
		}
		totalRolls += numRemoved
		if !part2 || numRemoved == 0 {
			break
		}
	}
	fmt.Println(totalRolls)
}
