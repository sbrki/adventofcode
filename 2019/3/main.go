package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isHorizontal(a [2][2]int) bool {
	if a[0][1] == a[1][1] {
		return true
	}
	return false
}

func doesIntersect(a [2][2]int, b [2][2]int) bool {
	if isHorizontal(a) == isHorizontal(b) {
		return false
	}

	if isHorizontal(a) {
		if ((a[0][0] > b[0][0]) == (a[1][0] > b[0][0])) || ((a[0][1] > b[0][1]) == (a[0][1] > b[1][1])) {
			return false
		}
		return true
	}

	if ((a[0][0] > b[0][0]) == (a[0][0] > b[1][0])) || ((a[0][1] > b[0][1]) == (a[1][1] > b[0][1])) {
		return false
	}
	return true
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func manhattanIntersect(a [2][2]int, b [2][2]int) int {
	if isHorizontal(a) {
		return abs(a[0][1]) + abs(b[0][0])
	}
	return abs(a[0][0]) + abs(b[0][1])
}

func intersection(a [2][2]int, b [2][2]int) (r [2]int) {
	if isHorizontal(a) {
		r[0] = b[0][0]
		r[1] = a[0][1]
		return r
	}

	r[0] = a[0][0]
	r[1] = b[0][1]
	return r
}

func pointIsOnLine(p [2]int, a [2][2]int) bool {
	if isHorizontal(a) {
		if (p[1] == a[0][1]) && ((a[0][0] > p[0]) != (a[1][0] > p[0])) {
			return true
		}
		return false
	}

	if (p[0] == a[0][0]) && ((a[0][1] > p[1]) != (a[1][1] > p[1])) {
		return true
	}
	return false
}

func distanceInPath(a [2]int, path [][2][2]int) int {
	sum := 0
	for _, v := range path {
		if pointIsOnLine(a, v) {
			sum += abs(v[0][1]-a[1]) + abs(v[0][0]-a[0])
			return sum
		}

		sum += abs(v[0][0]-v[1][0]) + abs(v[0][1]-v[1][1])
	}
	return -9999999999
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	paths := make([][][2][2]int, 0)

	for i := 0; i < 2; i++ {
		scanner.Scan()
		text := scanner.Text()

		strList := strings.Split(text, ",")

		path := make([][2][2]int, 0)
		x, y := 0, 0

		for _, token := range strList {
			delta, _ := strconv.Atoi(token[1:])
			var jump [2][2]int
			//fmt.Println(string(token[0]), delta)

			switch string(token[0]) {
			case "U":
				jump[0][0] = x
				jump[0][1] = y
				jump[1][0] = x
				jump[1][1] = y + delta
				path = append(path, jump)
				y += delta
			case "D":
				jump[0][0] = x
				jump[0][1] = y
				jump[1][0] = x
				jump[1][1] = y - delta
				path = append(path, jump)
				y -= delta
			case "L":
				jump[0][0] = x
				jump[0][1] = y
				jump[1][0] = x - delta
				jump[1][1] = y
				path = append(path, jump)
				x -= delta
			case "R":
				jump[0][0] = x
				jump[0][1] = y
				jump[1][0] = x + delta
				jump[1][1] = y
				path = append(path, jump)
				x += delta
			}
		}
		paths = append(paths, path)
	}

	// find the intersections
	min := -1
	for _, v1 := range paths[0] {
		for _, v2 := range paths[1] {
			if doesIntersect(v1, v2) {
				fmt.Println(v1, v2)
				p := intersection(v1, v2)
				l := distanceInPath(p, paths[0]) + distanceInPath(p, paths[1])
				if l < min || min == -1 {
					min = l
				}
			}
		}
	}

	fmt.Println(min)
}
