package main

import (
	"bufio"
	"fmt"
	"os"
)

func readInputs(scanner *bufio.Scanner) []int64 {
	columns := make([]int64, 0)
	for scanner.Scan() {
		for c, ch := range scanner.Text() {
			if c >= len(columns) {
				columns = append(columns, 0)
			}
			if ch == '*' || ch == '+' {
				return columns
			}
			if ch >= '0' && ch <= '9' {
				columns[c] = columns[c]*10 + int64(ch-'0')
			}
		}
	}
	return columns
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	columns := readInputs(scanner)
	total := int64(0)
	line := scanner.Text()
	for c := 0; c < len(columns); c++ {
		if line[c] == '+' {
			sum := int64(0)
			for ; c < len(columns) && columns[c] != 0; c++ {
				sum += columns[c]
			}
			total += sum
		} else if line[c] == '*' {
			product := int64(1)
			for ; c < len(columns) && columns[c] != 0; c++ {
				product *= columns[c]
			}
			total += product
		}
	}
	fmt.Println(total)
}
