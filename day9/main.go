package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/vanjac/aoc25-go/util"
)

const part2 = true

func abs(i int64) int64 {
	if i < 0 {
		return -i
	}
	return i
}

func rectIntersectShape(r1, r2 [2]int64, shape [][2]int64) bool {
	minX := min(r1[0], r2[0])
	maxX := max(r1[0], r2[0])
	minY := min(r1[1], r2[1])
	maxY := max(r1[1], r2[1])
	for i := 0; i < len(shape); i++ {
		s1 := shape[i]
		s2 := shape[(i+1)%len(shape)]
		if s1[0] == s2[0] {
			if rectIntersectVLine(minX, maxX, minY, maxY, s1[0], s1[1], s2[1]) {
				return true
			}
		} else if s1[1] == s2[1] {
			if rectIntersectHLine(minX, maxX, minY, maxY, s1[0], s2[0], s2[1]) {
				return true
			}
		}
	}
	return false
}

func rectIntersectHLine(minX, maxX, minY, maxY, lx1, lx2, ly int64) bool {
	return rectIntersectVLine(minY, maxY, minX, maxX, ly, lx1, lx2)
}

func rectIntersectVLine(minX, maxX, minY, maxY, lx, ly1, ly2 int64) bool {
	if lx > minX && lx < maxX {
		if (ly1 > minY && ly1 < maxY) || (ly2 > minY && ly2 < maxY) {
			return true
		}
		if (ly1 <= minY && ly2 > minY) || (ly2 <= minY && ly1 > minY) {
			return true
		}
		if (ly1 < maxY && ly2 >= maxY) || (ly2 < maxY && ly1 >= maxY) {
			return true
		}
	}
	return false
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
			if part2 && rectIntersectShape(redCoords[i], redCoords[j], redCoords) {
				continue
			}
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
