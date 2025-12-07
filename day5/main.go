package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/vanjac/aoc25-go/util"
)

type Range struct {
	min int64
	max int64
}

func withinAnyRange(ranges []Range, id int64) bool {
	for _, r := range ranges {
		if id >= r.min && id <= r.max {
			return true
		}
	}
	return false
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	ranges := make([]Range, 0)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			break
		}
		parts := strings.SplitN(line, "-", 2)
		min, err := strconv.ParseInt(parts[0], 10, 64)
		util.Check(err)
		max, err := strconv.ParseInt(parts[1], 10, 64)
		util.Check(err)
		ranges = append(ranges, Range{min, max})
	}
	fresh := 0
	for scanner.Scan() {
		line := scanner.Text()
		id, err := strconv.ParseInt(line, 10, 64)
		util.Check(err)
		if withinAnyRange(ranges, id) {
			fresh++
		}
	}
	fmt.Println(fresh)
}
