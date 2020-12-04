package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	content := openStringFromFile("data")
	count := 0
	count2 := 0

	for _, value := range content {
		rules := getRule(value)
		// rules[0] - Min
		// rules[1] - Max
		// rules[2] - Char
		// rules[3] - Password
		if checkValidation1(rules) {
			count++
		}
		if checkValidation2(rules) {
			count2++
		}
	}

	fmt.Println("Old rule:", count, "New rule:", count2)
}

func openStringFromFile(filename string) []string {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	return strings.Split(string(content), "\r\n")
}

func getRule(s string) []string {

	s2 := strings.Split(s, "-")
	min := s2[0]

	s2 = strings.SplitAfterN(s2[1], " ", 2)
	max := s2[0]

	s2 = strings.Split(s2[1], ":")
	char := s2[0]

	password := s2[1][1:]

	return []string{min, max, char, password}
}

func checkValidation1(rules []string) bool {
	// rules[0] - Min
	// rules[1] - Max
	// rules[2] - Char
	// rules[3] - Password

	k := strings.Count(rules[3], rules[2])
	min, _ := strconv.Atoi(rules[0])
	max, _ := strconv.Atoi(rules[1][0 : len(rules[1])-1])
	//fmt.Println(rules)

	//fmt.Println(min, max, rules[2], rules[3], k)
	if k >= min && k <= max {
		return true
	}
	return false
}

func checkValidation2(rules []string) bool {
	// rules[0] - pos1
	// rules[1] - pos2
	// rules[2] - Char
	// rules[3] - Password

	i := 0
	pos1, _ := strconv.Atoi(rules[0])
	pos2, _ := strconv.Atoi(rules[1][0 : len(rules[1])-1])

	if string(rules[3][pos1-1]) == rules[2] {
		i++
	}

	if string(rules[3][pos2-1]) == rules[2] {
		i++
	}

	if i == 1 {
		fmt.Println(rules[2], pos1, pos2, rules[3])
		return true
	}
	return false
}
