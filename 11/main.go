package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	input := readFile("data")
	paddedInput := padding(input)

	getPermanentState(paddedInput)

}

func getPermanentState(p []string) {

	places := make([]string, len(p))
	copy(places, p)
	changed := -1

	for changed != 0 {
		changed, places = updateState(places)
		fmt.Println(changed, "were changed this epoch. Curretly arr:")
		fmt.Print(strings.Join(places[:], "\n"))
		fmt.Println(" ")
		fmt.Println(" ")
	}

	occupied := getOccupiedNb(places)

	fmt.Println("The permanent state have", occupied, "occupied places")

}

func getOccupiedNb(places []string) int {
	occupied := 0
	for _, rowValues := range places {
		occupied += strings.Count(rowValues, "#")

	}
	return occupied
}

func updateState(p []string) (int, []string) {

	places := make([]string, len(p))
	copy(places, p)
	tolerance := 5 // part 2 - 5  | part 1 - 4
	changedSits := 0
	for rowNb, rowValues := range p {
		for columnNb, place := range rowValues {
			//occupiedNearby := getOccupiedAdjacents1(rowNb, columnNb, p)// Part1
			occupiedNearby := getOccupiedAdjacents2(rowNb, columnNb, p) // part2
			if occupiedNearby == 0 && strings.Contains(string(place), "L") {
				places[rowNb] = places[rowNb][:columnNb] + string("#") + places[rowNb][columnNb+1:]
				changedSits++
			} else if occupiedNearby >= tolerance && strings.Contains(string(place), "#") {
				places[rowNb] = places[rowNb][:columnNb] + string("L") + places[rowNb][columnNb+1:]
				changedSits++
			}
		}
	}
	return changedSits, places
}
func getOccupiedAdjacents2(row int, column int, places []string) int {

	var adjacents []string

	counter := 0

	if row != 0 && column != 0 && column != len(places[0])-1 && row != len(places)-1 { //check
		adjacents = append(adjacents, getDiagonals(row, column, places)...)
		adjacents = append(adjacents, getHorizontals(row, column, places)...)
		adjacents = append(adjacents, getVerticals(row, column, places)...)
	}

	for _, value := range adjacents {
		if value == "#" { //count all occupied
			counter++
		}
	}

	return counter
}

func getHorizontals(row int, column int, places []string) []string {
	i := 1
	var horizontals []string
	west, east := false, false
	for column-i > 0 && !west {
		if string(places[row][column-i]) == "#" {
			horizontals = append(horizontals, "#")
			west = true
		} else if string(places[row][column-i]) == "L" {
			west = true
		}
		i++
	}

	i = 1
	for column+i < len(places[0]) && !east {
		if string(places[row][column+i]) == "#" {
			horizontals = append(horizontals, "#")
			east = true
		} else if string(places[row][column+i]) == "L" {
			east = true
		}
		i++
	}

	return horizontals
}

func getVerticals(row int, column int, places []string) []string {
	i := 1
	var verticals []string
	north, south := false, false
	for row-i > 0 && !north {
		if string(places[row-i][column]) == "#" {
			verticals = append(verticals, "#")
			north = true
		} else if string(places[row-i][column]) == "L" {
			north = true
		}
		i++
	}

	i = 1
	for row+i < len(places) && !south {
		if string(places[row+i][column]) == "#" {
			verticals = append(verticals, "#")
			south = true
		} else if string(places[row+i][column]) == "L" {
			south = true
		}
		i++
	}

	return verticals
}

func getDiagonals(row int, column int, places []string) []string {
	i := 1
	var diagonals []string
	se, ne, nw, sw := false, false, false, false
	for i+row < len(places) && i+column < len(places[0]) && !se {
		if string(places[row+i][column+i]) == "#" {
			diagonals = append(diagonals, "#")
			se = true
		} else if string(places[row+i][column+i]) == "L" {
			se = true
		}
		i++
	}

	i = 1
	for row-i > 0 && i+column < len(places[0]) && !ne {
		if string(places[row-i][column+i]) == "#" {
			diagonals = append(diagonals, "#")
			ne = true
		} else if string(places[row-i][column+i]) == "L" {
			ne = true
		}
		i++

	}

	i = 1
	for i+row < len(places) && column-i > 0 && !nw {
		if string(places[row+i][column-i]) == "#" {
			diagonals = append(diagonals, "#")
			nw = true
		} else if string(places[row+i][column-i]) == "L" {
			nw = true
		}
		i++

	}

	i = 1
	for row-i > 0 && column-i > 0 && !sw {
		if string(places[row-i][column-i]) == "#" {
			diagonals = append(diagonals, "#")
			sw = true
		} else if string(places[row-i][column-i]) == "L" {
			sw = true
		}
		i++
	}

	return diagonals
}

func getOccupiedAdjacents1(row int, column int, places []string) int {
	var adjacents []string
	counter := 0

	if row != 0 && column != 0 && column != len(places[0])-1 && row != len(places)-1 { //check
		adjacents = append(adjacents, string(places[row-1][column-1]))
		adjacents = append(adjacents, string(places[row-1][column]))
		adjacents = append(adjacents, string(places[row-1][column+1]))
		adjacents = append(adjacents, string(places[row][column-1]))
		adjacents = append(adjacents, string(places[row][column+1]))
		adjacents = append(adjacents, string(places[row+1][column-1]))
		adjacents = append(adjacents, string(places[row+1][column]))
		adjacents = append(adjacents, string(places[row+1][column+1]))
	}

	for _, value := range adjacents {
		if value == "#" { //count all occupied
			counter++
		}
	}

	return counter
}

func padding(s []string) []string {

	fill := []byte(".")

	lenght := len(s[0])

	padding := make([]byte, lenght)

	i := 0
	for i < lenght {
		padding[i] = fill[0]
		i++
	}

	output := []string{string(padding)}
	output = append(output, s...)
	output = append(output, string(padding))

	for i := range output {
		output[i] = "." + output[i] + "."
	}

	return output
}

func readFile(filename string) []string {
	file, _ := os.Open(filename)

	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		currentLine := scanner.Text()
		lines = append(lines, currentLine)
	}

	return lines
}
