package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

type pos struct {
	y     int
	x     int
	angle int
}

func main() {
	var ship pos
	waypoint := pos{1, 10, 0}

	commands := readFile("data")

	for _, command := range commands {
		handleMovements(command, &ship, &waypoint, true)
	}

	fmt.Println()
	fmt.Println("Final ship location:", ship)
	fmt.Println("Manhattan distance:", math.Abs(float64(ship.x))+math.Abs(float64(ship.y)))

}

func handleMovements(command string, ship *pos, waypoint *pos, part2 bool) {
	if !part2 {
		if string(command[0]) == "R" || string(command[0]) == "L" {
			turn(ship, command, part2)
		} else {
			move(ship, command)
		}
		fmt.Println("Ferry moved. Current ship location: ", *ship)
	} else {
		if string(command[0]) == "F" {
			toWaypoint(ship, waypoint, command)
			fmt.Println("Ferry moved. Current ship location: ", *ship)
		} else if string(command[0]) == "R" || string(command[0]) == "L" {
			turn(waypoint, command, part2)
			fmt.Println("Waypont turned. Current waypoint location: ", *waypoint)
		} else {
			move(waypoint, command)
			fmt.Println("Waypont moved. Current waypoint location: ", *waypoint)
		}
	}
}

func toWaypoint(ship *pos, waypoint *pos, command string) {
	value, _ := strconv.Atoi(string(command[1:]))
	(*ship).x += (*waypoint).x * value
	(*ship).y += (*waypoint).y * value
}

func turn(object *pos, command string, part2 bool) {
	direction := string(command[0])
	angle, _ := strconv.Atoi(string(command[1:]))

	if direction == "R" {
		(*object).angle += angle
	} else {
		(*object).angle -= angle
	}

	if (*object).angle > 359 {
		(*object).angle -= 360
	}

	if (*object).angle < 0 {
		(*object).angle += 360
	}

	if part2 {
		if direction == "R" {
			if angle == 90 {
				aux := (*object).y
				(*object).y = -(*object).x
				(*object).x = aux
			} else if angle == 180 {
				(*object).y = -(*object).y
				(*object).x = -(*object).x
			} else if angle == 270 {
				aux := (*object).y
				(*object).y = (*object).x
				(*object).x = -aux
			}
		} else {
			if angle == 90 {
				aux := (*object).y
				(*object).y = (*object).x
				(*object).x = -aux
			} else if angle == 180 {
				(*object).y = -(*object).y
				(*object).x = -(*object).x
			} else if angle == 270 {
				aux := (*object).y
				(*object).y = -(*object).x
				(*object).x = aux
			}
		}
	}

}

func move(pos *pos, command string) {
	direction := string(command[0])
	value, _ := strconv.Atoi(string(command[1:]))

	if direction == "N" {
		(*pos).y += value
	} else if direction == "S" {
		(*pos).y -= value
	} else if direction == "E" {
		(*pos).x += value
	} else if direction == "W" {
		(*pos).x -= value
	} else { // case F
		if (*pos).angle == 0 { // begin to east
			(*pos).x += value
		} else if (*pos).angle == 90 {
			(*pos).y -= value
		} else if (*pos).angle == 180 {
			(*pos).x -= value
		} else if (*pos).angle == 270 {
			(*pos).y += value
		} else {
			os.Exit(200)
		}
	}

}

func readFile(fileName string) []string {
	file, _ := os.Open(fileName)

	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}
