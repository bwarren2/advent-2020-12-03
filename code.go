package advent20201203

import (
	"bufio"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// TerrainFromFile gets a 2D array of data to reflect trees in terrain
func TerrainFromFile(filename string) (result [][]int) {
	file, err := os.Open(filename)
	check(err)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		var tempList []int
		for _, char := range scanner.Text() {
			if char == '.' {
				tempList = append(tempList, 0)
			} else {
				tempList = append(tempList, 1)
			}
		}
		result = append(result, tempList)
	}
	return
}

// IndexerFn creates a function to do count-by-in-mod
// We work off 1-based inputs
func IndexerFn(coeff, mod int) func(int) int {
	return func(input int) int {
		return ((input) * coeff) % mod
	}
}

// Part1 answers part 1 of today's advent problem
func Part1(filename string, right int, down int) (sumTrees int) {
	terrain := TerrainFromFile(filename)
	rightFn := IndexerFn(right, len(terrain[0]))
	for idx, lst := range terrain {
		if (idx)%down == 0 {
			sumTrees += lst[rightFn(idx/down)]
		}
	}
	return
}
