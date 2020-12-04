package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	var intv1 int
	var intv2 int
	var intv3 int

	content, _ := ioutil.ReadFile("data")
	input := strings.Split(string(content), "\r\n")
	input = input[:len(input)-1]
	for _, v1 := range input {
		intv1, _ = strconv.Atoi(v1)
		for _, v2 := range input {
			intv2, _ = strconv.Atoi(v2)
			for _, v3 := range input {
				intv3, _ = strconv.Atoi(v3)
				//fmt.Println(intv1 + intv2)
				if (intv1 + intv2 + intv3) == 2020 {
					fmt.Println(intv1 * intv2 * intv3)
				}
			}
		}
	}

}
