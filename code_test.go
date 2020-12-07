package advent20201203_test

import (
	"fmt"
	"testing"

	advent "github.com/bwarren2/advent20201203"
)

func TestPart1(t *testing.T) {
	answer := 1
	answer *= advent.Part1("terrain.txt", 1, 1)
	answer *= advent.Part1("terrain.txt", 3, 1)
	answer *= advent.Part1("terrain.txt", 5, 1)
	answer *= advent.Part1("terrain.txt", 7, 1)
	answer *= advent.Part1("terrain.txt", 1, 2)
	fmt.Println(answer)
}

func TestTerrainFromFile(t *testing.T) {
	advent.TerrainFromFile("terrain.txt")
}
func TestIndexerFn(t *testing.T) {
	fn := advent.IndexerFn(2, 5)
	fn(1)
}
