package main

import (
	"fmt"
	"github.com/vanjac/aoc25-go/util"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

const part2 = true

func numDigits(n int64) int {
	return int(math.Floor(math.Log10(float64(n)))) + 1
}

func nextInvalidPattern(n int64, repeats int) int64 {
	digits := numDigits(n)
	patDigits := (digits + repeats - 1) / repeats
	if digits%repeats == 0 {
		pattern := n / int64(math.Pow10(patDigits*(repeats-1)))
		if repeatPattern(pattern, repeats) < n {
			return pattern + 1
		}
		return pattern
	} else {
		return int64(math.Pow10(patDigits - 1))
	}
}

func prevInvalidPattern(n int64, repeats int) int64 {
	digits := numDigits(n)
	patDigits := digits / repeats
	if digits%repeats == 0 {
		pattern := n / int64(math.Pow10(patDigits*(repeats-1)))
		if repeatPattern(pattern, repeats) > n {
			return pattern - 1
		}
		return pattern
	} else {
		return int64(math.Pow10(patDigits)) - 1
	}
}

func repeatPattern(pattern int64, repeats int) int64 {
	sum := int64(0)
	patBase := int64(math.Pow10(numDigits(pattern)))
	for range repeats {
		sum += pattern
		pattern *= patBase
	}
	return sum
}

func main() {
	data, err := io.ReadAll(os.Stdin)
	util.Check(err)
	dataStr := strings.Trim(string(data), "\n")
	sum := int64(0)
	invalidIds := make(map[int64]bool)
	for _, r := range strings.Split(dataStr, ",") {
		parts := strings.SplitN(r, "-", 2)
		minId, err := strconv.ParseInt(parts[0], 10, 64)
		util.Check(err)
		maxId, err := strconv.ParseInt(parts[1], 10, 64)
		util.Check(err)
		maxRepeats := 2
		if part2 {
			maxRepeats = numDigits(maxId)
		}
		for repeats := 2; repeats <= maxRepeats; repeats++ {
			minPattern := nextInvalidPattern(minId, repeats)
			maxPattern := prevInvalidPattern(maxId, repeats)
			for pat := minPattern; pat <= maxPattern; pat++ {
				invalid := repeatPattern(pat, repeats)
				if !invalidIds[invalid] {
					sum += invalid
					invalidIds[invalid] = true
				}
			}
		}

	}
	fmt.Println(sum)
}
