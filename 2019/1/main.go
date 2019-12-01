package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var sum int
	for {
		scanner.Scan()
		text := scanner.Text()
		if len(text) == 0 {
			fmt.Println(sum)
			return
		}

		mass, _ := strconv.Atoi(text)
		fuel := (mass / 3) - 2
		sum += fuel

		for {
			fuel = (fuel / 3) - 2
			if fuel >= 0 {
				sum += fuel
			} else {
				break
			}
		}
	}

}
