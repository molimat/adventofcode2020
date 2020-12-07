package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	//teste := []string{"red|d", "abc", "opq|qpe", "acd|fr"}
	//ansByGroups := readFile("data") //each value within has all answers of the group toguether PART 1
	ansByGroups := readFilep2("data") //part 2
	yesCountingByGroup := countAnswers(ansByGroups)
	total := 0

	for _, value := range yesCountingByGroup {
		total += value
	}

	fmt.Println(total)
}

func countAnswers(s []string) map[int]int { //map[group]yesAns
	table := make(map[int]int)
	i := 0
	for _, answers := range s {
		// yesCounting := countUnique(answers) //part1
		yesCounting := countConsistent(answers) // part 2
		table[i] = yesCounting
		i++
	}

	return table
}

func countConsistent(s string) int {
	i := strings.Index(s, "|")
	if i == -1 {
		k := countUnique(s)
		//fmt.Println(s, k)
		return k
	}
	lastCheckedPerson := string(s[:i])

	duplicates := lastCheckedPerson

	group := strings.Split(s, "|")

	for _, currentPerson := range group {

		duplicates = duplicated(duplicates, currentPerson)

		//fmt.Println("Current Person", currentPerson, "Last Person", lastCheckedPerson, "duplicated", duplicates)
		lastCheckedPerson = currentPerson
	}
	k := countUnique(duplicates)
	//fmt.Println(duplicates, k)
	return k
}

func duplicated(a string, b string) string {
	var duplicated string
	for i := 0; i < len(b); i++ {
		k := strings.Index(a, string(b[i]))
		if k != -1 {
			duplicated += string(b[i])
		}
	}
	return duplicated
}

func countUnique(s string) int {
	mapping := make(map[string]int)
	for i := 0; i < len(s); i++ {
		mapping[string(s[i])] = i
	}

	return len(mapping)
}

func readFilep2(filename string) []string {
	file, _ := os.Open(filename)

	scanner := bufio.NewScanner(file)

	var currentLine string

	var groupAnswer string

	var ansByGroups []string

	for scanner.Scan() {
		currentLine = scanner.Text()
		if currentLine != "" {
			if groupAnswer == "" {
				groupAnswer = string(currentLine)
			} else {
				groupAnswer = groupAnswer + "|" + string(currentLine) // append nas strings do mesmo grupo
			}
		} else {
			ansByGroups = append(ansByGroups, groupAnswer)
			groupAnswer = ""
		}
	}

	ansByGroups = append(ansByGroups, groupAnswer) // lastLine

	return ansByGroups
}

func readFile(filename string) []string {
	file, _ := os.Open(filename)

	scanner := bufio.NewScanner(file)

	var currentLine string

	var groupAnswer string

	var ansByGroups []string

	for scanner.Scan() {
		currentLine = scanner.Text()
		if currentLine != "" {
			groupAnswer = groupAnswer + string(currentLine) // append nas strings do mesmo grupo
		} else {
			ansByGroups = append(ansByGroups, groupAnswer)
			groupAnswer = ""
		}
	}

	ansByGroups = append(ansByGroups, groupAnswer) // lastLine

	return ansByGroups
}
