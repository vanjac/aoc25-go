package main

import (
	"bufio"
	"fmt"
	"os"
)

const part2 = true

func main() {
	beams := make([]int, 0)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	for i, ch := range scanner.Text() {
		beams = append(beams, 0)
		if ch == 'S' {
			beams[i] = 1
		}
	}
	numSplits := 0
	if part2 {
		numSplits = 1
	}
	for scanner.Scan() {
		newBeams := make([]int, len(beams))
		for i, ch := range scanner.Text() {
			if ch == '.' {
				newBeams[i] += beams[i]
			} else if ch == '^' {
				if i > 0 {
					newBeams[i-1] += beams[i]
				}
				if i+1 < len(beams) {
					newBeams[i+1] += beams[i]
				}
				if !part2 && beams[i] != 0 {
					numSplits++
				} else if part2 {
					numSplits += beams[i]
				}
			}
		}
		beams = newBeams
	}
	fmt.Println(numSplits)
}
