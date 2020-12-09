package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type command struct {
	instruction string
	value       int
}

func main() {
	instructions := readFile("data")
	correctBoot(instructions)

}

func correctBoot(instructions []command) int {
	i := 0
	old := ""
	oldIndex := 0
	value := -1
	for value == -1 {

		if old != "" {
			instructions[oldIndex].instruction = old
		}

		if instructions[i].instruction == "jmp" || instructions[i].instruction == "nop" {
			if instructions[i].instruction == "jmp" {
				instructions[i].instruction = "nop"
				old = "jmp"
				oldIndex = i
			} else {
				instructions[i].instruction = "jmp"
				old = "nop"
				oldIndex = i
			}
			fmt.Println("Changing instruction", old, "on index", oldIndex, "to instruction", instructions[i].instruction)
		}
		i++

		value = boot(instructions)
	}
	return value
}

func boot(commands []command) int {
	executedCommands := make(map[int]string)
	i := 0
	counter := 0
	for i < len(commands) {
		instruct := commands[i].instruction
		value := commands[i].value
		if executedCommands[i] == "" {
			executedCommands[i] = instruct
		} else {
			fmt.Println("Infinite loop begun on line:", i, "Current counter value:", counter)
			return -1
		}

		if instruct == "jmp" {
			i = jmp(i, value)
		} else if instruct == "acc" {
			counter += value
			i++
		} else {
			i++
		}
	}
	fmt.Println("System booted properly. Value:", counter)
	return counter
}

func jmp(pos int, value int) int {
	return pos + value
}

func nop(pos int) int {
	return pos + 1
}

func acc(counter int, value int) int {
	return counter + value
}

func readFile(filename string) []command {
	file, _ := os.Open(filename)

	scanner := bufio.NewScanner(file)

	var instruction []command

	for scanner.Scan() {
		currentLine := scanner.Text()
		value, _ := strconv.Atoi(currentLine[4:])
		instruction = append(instruction, command{string(currentLine[0:3]), value})
	}

	return instruction
}
