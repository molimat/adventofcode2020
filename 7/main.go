package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type bag struct {
	name   string
	inside map[string]int
}

func main() {
	fileData := readFile("data")
	allRelations := getHierarchy(fileData)
	// fathersMap := make(map[string]int) // part 1
	// total := countPaths(allRelations, "shiny gold", fathersMap) // part 1
	total := countChildren(allRelations, "shiny gold")
	fmt.Println(total)

}

func countChildren(b []bag, s string) int {
	//s is the bag willing to be counted
	children := getChildren(b, s)
	total := 0
	for _, value := range children {
		total += value
	}
	for key, value := range children {
		total += value * countChildren(b, key)
	}
	return total

}

func countPaths(b []bag, s string, m map[string]int) map[string]int {
	//s is the bag willing to be counted

	fathers, _ := getFathers(b, s)
	for _, father := range fathers {
		m[father] = 1
		m = countPaths(b, father[:len(father)-6], m)
	}
	return m

}

func getFathers(b []bag, s string) ([]string, int) {
	var fathers []string
	counter := 0
	for _, bag := range b {
		if len(bag.inside) != 0 {
			for key := range bag.inside {
				if strings.Contains(key, s) {
					fathers = append(fathers, bag.name) //dad
					counter++
				}
			}
		}
	}
	return fathers, counter
}

func getChildren(b []bag, s string) map[string]int {
	children := make(map[string]int)
	for _, bag := range b {
		if strings.Contains(bag.name, s) && len(bag.inside) != 0 {
			for key, value := range bag.inside {
				children[key] = value
			}
		}
	}
	return children
}

func getHierarchy(s []string) []bag {
	relations := []bag{}
	for _, value := range s {
		name, children := getOrder(value)
		relations = append(relations, bag{name, children})
	}
	return relations
}

func getOrder(s string) (string, map[string]int) { // return dad bag and children
	i := strings.Index(s, " contain")
	dad := string(s[:i])
	children := string(s[i+8:])
	childrenMap := make(map[string]int)

	if strings.Contains(s, "no other bags") {
		return "", childrenMap
	}

	for strings.ContainsAny(children, ",") || strings.ContainsAny(children, ".") {
		if strings.Contains(children, "bags,") {
			k := strings.Index(children, ",")
			childrenMap[string(children[3:k-5])], _ = strconv.Atoi(string(children[1]))
			//fmt.Println(string(children[3:k]))
			children = string(children[k+1:])
		} else if strings.Contains(children, "bag,") {
			k := strings.Index(children, ",")
			childrenMap[string(children[3:(k-4)])], _ = strconv.Atoi(string(children[1]))
			//fmt.Println(string(children[3:k]))
			children = string(children[k+1:])
		} else if strings.Contains(children, "bags.") {
			k := strings.Index(children, ".")
			childrenMap[string(children[3:k-5])], _ = strconv.Atoi(string(children[1]))
			children = ""
		} else {
			k := strings.Index(children, ".")
			childrenMap[string(children[3:k-4])], _ = strconv.Atoi(string(children[1]))
			children = ""
		}
	}

	return dad, childrenMap
}

func readFile(filename string) []string {
	file, _ := os.Open(filename)

	scanner := bufio.NewScanner(file)

	var fileData []string

	for scanner.Scan() {
		currentLine := scanner.Text()
		fileData = append(fileData, string(currentLine))
	}

	return fileData
}
