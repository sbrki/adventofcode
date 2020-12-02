package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func solve(lines []string) int {
	valid := 0

	for _, line := range lines {
		parts := strings.Split(line, " ")
		if len(parts) != 3 {
			panic("len(parts) != 3")
		}

		targetLetter := parts[1]
		targetLetter = targetLetter[:len(targetLetter)-1]

		password := parts[2]

		bounds := strings.Split(parts[0], "-")
		if len(bounds) != 2 {
			panic("len(bounds) != 2")
		}
		lowerBound, _ := strconv.Atoi(bounds[0])
		upperBound, _ := strconv.Atoi(bounds[1])

		// check the password
		counter := 0
		for _, character := range password {
			if string(character) == targetLetter {
				counter++
			}

			// slight optimisation
			if counter > upperBound {
				continue
			}
		}

		if counter >= lowerBound && counter <= upperBound {
			valid++
		}
	}
	return valid
}

func solve2(lines []string) int {
	valid := 0

	for _, line := range lines {
		parts := strings.Split(line, " ")
		if len(parts) != 3 {
			panic("len(parts) != 3")
		}

		targetLetter := parts[1]
		targetLetter = targetLetter[:len(targetLetter)-1]

		password := parts[2]

		bounds := strings.Split(parts[0], "-")
		if len(bounds) != 2 {
			panic("len(bounds) != 2")
		}
		firstIndex, _ := strconv.Atoi(bounds[0])
		secondIndex, _ := strconv.Atoi(bounds[1])

		// check the password
		firstLetter := ""
		secondLetter := ""
		if len(password) >= firstIndex {
			firstLetter = string(password[firstIndex-1])
		}
		if len(password) >= secondIndex {
			secondLetter = string(password[secondIndex-1])
		}

		if (firstLetter == targetLetter && secondLetter != targetLetter) || (firstLetter != targetLetter && secondLetter == targetLetter) {
			valid++
		}

	}
	return valid
}

func main() {
	inputFile, err := os.Open("input.txt")
	if err != nil {
		panic(err.Error())
	}
	bytes, _ := ioutil.ReadAll(inputFile)
	input := string(bytes)
	lines := strings.Split(input, "\n")
	if lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1] // remove last "fake" line (empty string)
	}

	fmt.Println(solve(lines))
	fmt.Println("part 2")
	fmt.Println(solve2(lines))
}
