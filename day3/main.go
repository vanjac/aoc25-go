package main

import (
	"bufio"
	"fmt"
	"os"
)

const part2 = true

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		lastDigitIdx := -1
		jolts := 0
		battCount := 2
		if part2 {
			battCount = 12
		}
		for ; battCount > 0; battCount-- {
			maxDigit := byte(0)
			for i, end := lastDigitIdx+1, len(line); i < end-battCount+1; i++ {
				digit := line[i] - '0'
				if digit > maxDigit {
					lastDigitIdx = i
					maxDigit = digit
				}
			}
			jolts = jolts*10 + int(maxDigit)
		}
		sum += jolts
	}
	fmt.Println(sum)
}
