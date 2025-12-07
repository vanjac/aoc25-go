package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		firstIdx := -1
		firstDigit := byte(0)
		for i, end := 0, len(line); i < end-1; i++ {
			digit := line[i] - '0'
			if digit > firstDigit {
				firstIdx = i
				firstDigit = digit
			}
		}
		secondDigit := byte(0)
		for i, end := firstIdx+1, len(line); i < end; i++ {
			digit := line[i] - '0'
			if digit > secondDigit {
				secondDigit = digit
			}
		}
		sum += int(firstDigit*10 + secondDigit)
	}
	fmt.Println(sum)
}
