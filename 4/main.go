package main

// byr (Birth Year)
// iyr (Issue Year)
// eyr (Expiration Year)
// hgt (Height)
// hcl (Hair Color)
// ecl (Eye Color)
// pid (Passport ID)
// cid (Country ID)

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	nbValid := 0
	totalPassports := 0
	passports := getPassportsFromFile("data")

	for _, passport := range passports {
		totalPassports++
		if checkPassValidity(passport) {
			fmt.Println(passport)
			nbValid++
		}
	}

	fmt.Println(nbValid, totalPassports)
}

func checkPassValidity(passport map[string]string) bool {
	fields := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"} //cid doesnt matter
	fieldsNumber := 0
	for field, value := range passport {
		exists, index := find(fields, field)
		valid := checkField(field, value) // part 2
		if exists && valid {              // valid only for part 2
			fieldsNumber++
			fields = append(fields[:index], fields[index+1:]...)
		}
	}

	if fieldsNumber == 7 {
		return true
	}
	return false
}

func checkField(field string, value string) bool {
	// byr (Birth Year) - four digits; at least 1920 and at most 2002.
	// iyr (Issue Year) - four digits; at least 2010 and at most 2020.
	// eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
	// hgt (Height) - a number followed by either cm or in:
	//     If cm, the number must be at least 150 and at most 193.
	//     If in, the number must be at least 59 and at most 76.
	// hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
	// ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
	// pid (Passport ID) - a nine-digit number, including leading zeroes.
	// cid (Country ID) - ignored, missing or not.
	if field == "byr" {
		intValue, _ := strconv.Atoi(value)
		if intValue >= 1920 && intValue <= 2002 {
			return true
		}
		return false
	}

	if field == "iyr" {
		intValue, _ := strconv.Atoi(value)
		if intValue >= 2010 && intValue <= 2020 {
			return true
		}
		return false
	}

	if field == "eyr" {
		intValue, _ := strconv.Atoi(value)
		if intValue >= 2020 && intValue <= 2030 {
			return true
		}
		return false
	}

	if field == "hgt" {
		measureSystem := value[len(value)-2:]
		height, _ := strconv.Atoi(value[:len(value)-2])

		if measureSystem == "cm" {
			if height >= 150 && height <= 193 {
				return true
			}
		} else if measureSystem == "in" {
			if height >= 59 && height <= 76 {
				return true
			}
		}
		return false
	}

	if field == "ecl" {
		validColors := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
		found, _ := find(validColors, string(value))
		if found {
			return true
		}

		return false
	}

	if field == "hcl" {
		validCharacters := []string{"a", "b", "c", "d", "e", "f", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
		if string(value[0]) == "#" {
			if len(value) == 7 {
				for i := 1; i < 7; i++ {
					found, _ := find(validCharacters, string(value[i]))
					if !found {
						return false
					}
				}
				return true
			}
		}
		return false
	}

	if field == "pid" {
		validCharacters := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
		if len(value) == 9 {
			for i := 1; i < 9; i++ {
				found, _ := find(validCharacters, string(value[i]))
				if !found {
					return false
				}
			}
			return true
		}

		return false
	}

	return false
}

func find(s []string, e string) (bool, int) {
	for i, a := range s {
		if a == e {
			return true, i
		}
	}
	return false, -1
}

func getPassportsFromFile(filename string) []map[string]string {
	file, _ := os.Open(filename)

	scanner := bufio.NewScanner(file)

	currentPass := make(map[string]string)

	var passports []map[string]string

	for scanner.Scan() {
		currentLine := scanner.Text()

		if len(currentLine) != 0 { //So we know is the same pass
			values := strings.Split(currentLine, " ")
			for _, value := range values {
				field := strings.Split(value, ":")
				currentPass[field[0]] = field[1]
			}
		} else {
			passports = append(passports, currentPass)
			currentPass = make(map[string]string)
		}
	}
	passports = append(passports, currentPass)

	return passports
}
