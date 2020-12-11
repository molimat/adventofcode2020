package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	input := readFile("data")
	p := 25 // preamble
	k := 0
	for p < len(input) {
		validation := checkValidity(input[k:p], input[p])
		if !validation {
			fmt.Println("Number", input[p], "is not valid for the preamble: \n", input[k:p])
			fmt.Println("Trying to crack the code down") //part 2
			sum := crack(input[:p], input[p])
			fmt.Println("The cracked value is", sum)
			os.Exit(-1)
		}
		p++
		k++
	}

}

func crack(input []int, failure int) int {
	cracked := false
	k := 2
	for !cracked {
		for i := range input {
			slice := input[i : i+k]
			//fmt.Println(slice)
			//fmt.Println(i, k)
			if sum(slice) == failure {
				cracked = true
				fmt.Println(input[i : i+k])
				min := getMin(input[i : i+k])
				max := getMax(input[i : i+k])
				return min + max
			}
		}
		k++
	}

	return -1
}

func getMin(s []int) int {
	min := s[0]
	for _, value := range s {
		if value < min {
			min = value
		}
	}
	return min
}

func getMax(s []int) int {
	max := s[0]
	for _, value := range s {
		if value > max {
			max = value
		}
	}
	return max
}

func sum(s []int) int {
	result := 0
	for _, value := range s {
		result += value
	}
	return result
}

func checkValidity(preamble []int, currentValue int) bool {
	for i, number := range preamble {
		for i < len(preamble) {
			if (number + preamble[i]) == currentValue {
				return true
			}
			i++
		}
	}
	return false
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
