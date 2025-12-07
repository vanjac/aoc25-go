package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/vanjac/aoc25-go/util"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	inputs := make([][]int64, 0)
	var tokens []string
	for scanner.Scan() {
		line := scanner.Text()
		tokens = strings.Fields(line)
		if tokens[0] == "*" || tokens[0] == "+" {
			break
		}
		values := make([]int64, len(tokens))
		for i, t := range tokens {
			v, err := strconv.ParseInt(t, 10, 64)
			util.Check(err)
			values[i] = v
		}
		inputs = append(inputs, values)
	}
	total := int64(0)
	for c, t := range tokens {
		if t == "+" {
			sum := int64(0)
			for _, row := range inputs {
				sum += row[c]
			}
			total += sum
		} else if t == "*" {
			product := int64(1)
			for _, row := range inputs {
				product *= row[c]
			}
			total += product
		}
	}
	fmt.Println(total)
}
