package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	adapters := readFile("data")
	cache := make(map[string]int) // as the part 2 takes too much, we need a cache to handle the recursivity

	slice := append([]int{0}, adapters...)
	sort.Ints(slice)
	slice = append(slice, slice[len(slice)-1]+3)

	fmt.Println(countDistinctArrangements(slice, cache) + 1) // the first case is not count, thats why plus 1

}

//part 2 - Get different arrangements number
func countDistinctArrangements(slice []int, cache map[string]int) int {
	stringedSlice := toString(slice)

	if cache[stringedSlice] != 0 {
		return cache[stringedSlice]
	}

	arrangements := 0
	for i := range slice {
		k := 2
		for k < 4 && i+k < len(slice) {
			if slice[i+k]-slice[i] < 4 {
				if len(slice) == 1 {
					return 1
				}
				currArrangValue := countDistinctArrangements(slice[i+k:], cache)
				cache[toString(slice[i+k:])] = currArrangValue
				arrangements += currArrangValue
				arrangements++
			}
			k++

		}
	}
	return arrangements

}

func toString(i []int) string {
	var stringSlice []string
	for _, value := range i {
		stringSlice = append(stringSlice, strconv.Itoa(value))
	}
	return strings.Join(stringSlice, "")
}

// part 1
func adapterOrder(adapters []int) (int, int, int) {
	currentJoint, jointOne, jointTwo, jointThree := 0, 0, 0, 0
	sort.Ints(adapters)
	for _, value := range adapters {
		if value-currentJoint == 1 {
			jointOne++
			currentJoint = value
		} else if value-currentJoint == 2 {
			jointTwo++
			currentJoint = value
		} else if value-currentJoint == 3 {
			jointThree++
			currentJoint = value
		} else {
			return -1, -1, -1
		}
	}
	return jointOne, jointTwo, jointThree + 1 //+1 to the device joint
}

func readFile(filename string) []int {
	file, _ := os.Open(filename)

	scanner := bufio.NewScanner(file)
	var numbers []int

	for scanner.Scan() {
		currentLine := scanner.Text()
		number, _ := strconv.Atoi(currentLine)
		numbers = append(numbers, number)
	}

	return numbers
}
