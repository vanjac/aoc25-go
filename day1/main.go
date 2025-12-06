package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const part2 bool = true

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	dial := 50
	password := 0
	for scanner.Scan() {
		line := scanner.Text()
		num, err := strconv.ParseInt(line[1:], 10, 32)
		check(err)
		if line[0] == 'L' {
			dial -= int(num)
			if part2 && dial <= 0 {
				password += -dial / 100
				if dial != -int(num) {
					password += 1
				}
			}
		} else {
			dial += int(num)
			if part2 && dial >= 100 {
				password += dial / 100
			}
		}
		dial = (dial % 100 + 100) % 100
		if !part2 && dial == 0 {
			password++
		}
	}
	fmt.Println("Password:", password)
}
