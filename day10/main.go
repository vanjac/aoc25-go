package main

import (
	"bufio"
	"fmt"
	"math"
	"math/bits"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	totalButtons := 0
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		targetLights := make([]bool, len(fields[0])-2)
		for i := range targetLights {
			targetLights[i] = (fields[0][i+1] == '#')
		}
		buttons := make([][]int, len(fields)-2)
		for i := range buttons {
			wires := fields[i+1]
			parts := strings.Split(wires[1:len(wires)-1], ",")
			button := make([]int, len(parts))
			for i, str := range parts {
				index, _ := strconv.ParseInt(str, 10, 64)
				button[i] = int(index)
			}
			buttons[i] = button
		}

		numCombos := 1 << len(buttons)
		minButtons := math.MaxInt
		for combo := range numCombos {
			buttonCount := bits.OnesCount(uint(combo))
			if buttonCount >= minButtons {
				continue
			}
			lights := make([]bool, len(targetLights))
			for b, button := range buttons {
				if ((combo >> b) & 1) != 0 {
					for _, l := range button {
						lights[l] = !lights[l]
					}
				}
			}
			if slices.Equal(lights, targetLights) {
				minButtons = buttonCount
			}
		}
		totalButtons += minButtons
	}
	fmt.Println(totalButtons)
}
