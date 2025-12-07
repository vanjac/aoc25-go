package main

import (
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func numDigits(n int64) int {
	return int(math.Floor(math.Log10(float64(n)))) + 1
}

func nextInvalidPattern(n int64) int64 {
	digits := numDigits(n)
	halfBase := int64(math.Pow10((digits + 1) / 2))
	if digits%2 == 0 {
		low := n % halfBase
		high := n / halfBase
		return min(max(low, high), high+1)
	} else {
		return halfBase / 10
	}
}

func prevInvalidPattern(n int64) int64 {
	digits := numDigits(n)
	halfBase := int64(math.Pow10(digits / 2))
	if digits%2 == 0 {
		low := n % halfBase
		high := n / halfBase
		return max(min(low, high), high-1)
	} else {
		return halfBase - 1
	}
}

func repeatPattern(pattern int64) int64 {
	return pattern*int64(math.Pow10(numDigits(pattern))) + pattern
}

func main() {
	data, err := io.ReadAll(os.Stdin)
	check(err)
	dataStr := strings.Trim(string(data), "\n")
	sum := int64(0)
	for _, r := range strings.Split(dataStr, ",") {
		parts := strings.SplitN(r, "-", 2)
		minId, err := strconv.ParseInt(parts[0], 10, 64)
		check(err)
		maxId, err := strconv.ParseInt(parts[1], 10, 64)
		check(err)
		minPattern := nextInvalidPattern(minId)
		maxPattern := prevInvalidPattern(maxId)
		for pat := minPattern; pat <= maxPattern; pat++ {
			sum += repeatPattern(pat)
		}
	}
	fmt.Println(sum)
}
