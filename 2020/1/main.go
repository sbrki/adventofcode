package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

var target = 2020

// solve solves part 1 in O(n) time complexity
func solve(values []int, target int) (int, int) {
	valuesMap := make(map[int]bool)
	for _, value := range values {
		complement := target - value
		if value == complement {
			return complement, value
		}
		if _, ok := valuesMap[complement]; ok {
			return complement, value
		} else {
			valuesMap[value] = true
		}
	}

	// should be unreachable if the problem input contains a solution
	return 0, 0
}

// solve2 solves part 2 in O(n^3) time complexity :(
func solve2(values []int, target int) (int, int, int) {
	for _, a := range values {
		for _, b := range values {
			for _, c := range values {
				if a+b+c == target {
					return a, b, c
				}

			}
		}
	}
	return 0, 0, 0
}

func main() {
	inputFile, err := os.Open("input.txt")
	if err != nil {
		panic(err.Error())
	}
	bytes, _ := ioutil.ReadAll(inputFile)
	input := string(bytes)
	lines := strings.Split(input, "\n")
	lines = lines[:len(lines)-1] // remove last "fake" line (empty string)
	values := make([]int, len(lines))
	for idx, line := range lines {
		currVal, err := strconv.Atoi(line)
		if err != nil {
			panic(err.Error())
		}
		values[idx] = currVal

	}

	a, b := solve(values, target)
	fmt.Println("match:", a, b)
	fmt.Println("solution:", a*b)

	fmt.Println("part 2")
	a, b, c := solve2(values, target)
	fmt.Println("match:", a, b, c)
	fmt.Println("solution:", a*b*c)
}
