package main

import (
	"bufio"
	"cmp"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/vanjac/aoc25-go/util"
)

const part2 = true
const maxConnections = 1000

type Connection struct {
	box1, box2 int // index
	distance   float64
}

func main() {
	coords := make([][3]int, 0)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		var coord [3]int
		for i, str := range strings.SplitN(scanner.Text(), ",", 3) {
			val, err := strconv.ParseInt(str, 10, 64)
			util.Check(err)
			coord[i] = int(val)
		}
		coords = append(coords, coord)
	}

	availConn := make([]Connection, 0)
	for a := 0; a < len(coords); a++ {
		for b := a + 1; b < len(coords); b++ {
			sqDist := 0.0
			for i := range 3 {
				delta := float64(coords[a][i] - coords[b][i])
				sqDist += delta * delta
			}
			availConn = append(availConn, Connection{a, b, math.Sqrt(sqDist)})
		}
	}
	slices.SortFunc(availConn, func(a, b Connection) int {
		return cmp.Compare(a.distance, b.distance)
	})

	circuitIds := make([]int, len(coords))   // indexed by box
	circuitSizes := make([]int, len(coords)) // indexed by circuit
	for i := range coords {
		circuitIds[i] = i
		circuitSizes[i] = 1
	}

	numConn := maxConnections
	if part2 {
		numConn = len(availConn)
	}
	for c := 0; c < numConn; c++ {
		conn := availConn[c]
		id1 := circuitIds[conn.box1]
		id2 := circuitIds[conn.box2]
		if id1 != id2 {
			for i, id := range circuitIds {
				if id == id2 {
					circuitIds[i] = id1
				}
			}
			circuitSizes[id1] += circuitSizes[id2]
			circuitSizes[id2] = 0
			if part2 && circuitSizes[id1] == len(coords) {
				// last connection
				fmt.Println(coords[conn.box1][0] * coords[conn.box2][0])
				break
			}
		}
	}
	if !part2 {
		slices.Sort(circuitSizes)
		l := len(circuitSizes)
		fmt.Println(circuitSizes[l-1] * circuitSizes[l-2] * circuitSizes[l-3])
	}
}
