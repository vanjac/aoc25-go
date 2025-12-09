package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/vanjac/aoc25-go/util"
)

func abs(i int64) int64 {
	if i < 0 {
		return -i
	}
	return i
}

func main() {
	redCoords := make([][2]int64, 0)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		xStr, yStr, _ := strings.Cut(scanner.Text(), ",")
		x, err := strconv.ParseInt(xStr, 10, 64)
		util.Check(err)
		y, err := strconv.ParseInt(yStr, 10, 64)
		redCoords = append(redCoords, [2]int64{x, y})
	}

	maxArea := int64(0)
	for i := 0; i < len(redCoords); i++ {
		for j := 0; j < len(redCoords); j++ {
			w := abs(redCoords[i][0]-redCoords[j][0]) + 1
			h := abs(redCoords[i][1]-redCoords[j][1]) + 1
			area := w * h
			if area > maxArea {
				maxArea = area
			}
		}
	}
	fmt.Println(maxArea)
}
