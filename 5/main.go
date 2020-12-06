package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

func main() {
	var boardingPassesByID []int
	boardingPasses := openFile("data")

	for _, boardingPass := range boardingPasses {
		currentSitID := getSeatID(boardingPass)
		boardingPassesByID = append(boardingPassesByID, currentSitID)
	}

	fmt.Println("Max seat ID is", max(boardingPassesByID))
	findMySeatID(boardingPassesByID) // part 2

}

func findMySeatID(passesID []int) {

	sort.Ints(passesID)

	previousID := passesID[0]
	for _, value := range passesID {
		if (value-previousID) == 2 || (value-previousID) == 14 { //2 for when is a seat between columns, or 14 for when is a set in the beggining or the end of row
			fmt.Println("There is a free seat between", value, "and", previousID, ".") // to get the intuition of places
		}
		previousID = value
	}
}

func max(i []int) int {
	maxValue := 0
	for _, value := range i {
		if value > maxValue {
			maxValue = value
		}
	}
	return maxValue
}

func getSeatID(code string) int {
	rowCode := string(code[:len(code)-3])
	columnCode := string(code[len(code)-3:])
	row := getSeatRow(rowCode)
	column := getSeatColumn(columnCode)
	return (row * 8) + column
}

func getSeatRow(code string) int {
	maxRows := 128
	currentRow := 0
	for i := range code {
		if string(code[i]) == "F" {

		} else if string(code[i]) == "B" {
			currentRow += maxRows / int(math.Pow(2, float64(i+1)))
		} else {
			return -1
		}
	}
	return currentRow
}

func getSeatColumn(code string) int {
	maxColumn := 8
	currentColumn := 0
	for i := range code {
		if string(code[i]) == "L" {

		} else if string(code[i]) == "R" {
			currentColumn += maxColumn / int(math.Pow(2, float64(i+1)))
		} else {
			return -1
		}
	}
	return currentColumn
}

func openFile(filename string) []string {
	file, _ := os.Open(filename)

	scanner := bufio.NewScanner(file)
	var fileData []string

	for scanner.Scan() {
		currentLine := scanner.Text()
		fileData = append(fileData, currentLine)
	}

	return fileData
}
