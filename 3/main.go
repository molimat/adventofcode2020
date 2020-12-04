package main

import (
	"bufio"
	"fmt"
	"os"
)

type geomap struct {
	geography []string
	size      int //size to right, so we can repeat
}

func main() {

	var inputMap geomap
	(&inputMap).loadMap("data")
	//slope := []int{3, 1} part 1
	// treesFound := checkTrees(inputMap, slope) part1
	treesFound :=
		checkTrees(inputMap, []int{1, 1}) *
			checkTrees(inputMap, []int{3, 1}) *
			checkTrees(inputMap, []int{5, 1}) *
			checkTrees(inputMap, []int{7, 1}) *
			checkTrees(inputMap, []int{1, 2})
		// part2

	fmt.Println("Found", treesFound, "trees.")

}

func checkTrees(g geomap, slope []int) int { //slope[0] - Right Slope - slope[1]- Down slope
	actualX := 0 //x position
	trees := 0
	row := 0
	for _, line := range g.geography {
		if row%slope[1] == 0 {
			actualX = actualX % g.size // This will emulate a right replication of the map
			here := string(line[actualX])
			if here == "#" {
				trees++
			}
			actualX += slope[0]
		}
		row++
	}
	return trees
}

func (g *geomap) loadMap(filename string) {
	file, _ := os.Open(filename)

	scanner := bufio.NewScanner(file)

	var geographyMap []string

	for scanner.Scan() {
		geographyMap = append(geographyMap, scanner.Text())
	}

	(*g).geography = geographyMap
	(*g).size = len(geographyMap[0])
}
