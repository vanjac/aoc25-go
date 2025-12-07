package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	beams := make([]bool, 0)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	for _, ch := range scanner.Text() {
		beams = append(beams, ch == 'S')
	}
	numSplits := 0
	for scanner.Scan() {
		newBeams := make([]bool, len(beams))
		for i, ch := range scanner.Text() {
			if ch == '.' {
				newBeams[i] = newBeams[i] || beams[i]
			} else if ch == '^' {
				if beams[i] {
					numSplits++
					if i > 0 {
						newBeams[i-1] = true
					}
					if i+1 < len(beams) {
						newBeams[i+1] = true
					}
				}
			}
		}
		beams = newBeams
	}
	fmt.Println(numSplits)
}
