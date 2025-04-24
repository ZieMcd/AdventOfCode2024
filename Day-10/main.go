package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Pos struct {
	x int
	y int
}

var partTwo int

func main() {

	input := getInput()
	partTwo = 0
	partOne := 0
	for y, line := range input {
		for x, num := range line {
			if num == 0 {
				start := Pos{y: y, x: x}
				// fmt.Printf("y=%d, x=%d", y, x)
				f := make(map[Pos]bool)
				f = trail(input, start, -1, f, len(input[0]), len(input))
				partOne = partOne + len(f)
				// fmt.Println(len(f))
			}
		}
	}
	fmt.Println(partOne)
	fmt.Println(partTwo)
}

func trail(input [][]int, pos Pos, p int, f map[Pos]bool, w int, l int) map[Pos]bool {
	if pos.y < 0 {
		return f
	}
	if pos.y >= l {
		return f
	}
	if pos.x < 0 {
		return f
	}
	if pos.x >= w {
		return f
	}

	c := int(input[pos.y][pos.x])
	if p+1 != c {
		return f
	}

	if c == 9 {
		f[pos] = true
		partTwo++
		return f
	}

	left := Pos{x: pos.x - 1, y: pos.y}
	f = trail(input, left, c, f, w, l)

	right := Pos{x: pos.x + 1, y: pos.y}
	f = trail(input, right, c, f, w, l)

	up := Pos{x: pos.x, y: pos.y - 1}
	f = trail(input, up, c, f, w, l)

	down := Pos{x: pos.x, y: pos.y + 1}
	f = trail(input, down, c, f, w, l)

	//aka no where to go
	return f
}

func getInput() [][]int {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	result := [][]int{}

	for scanner.Scan() {
		toAdd := []int{}
		nums := strings.Split(scanner.Text(), "")
		for _, num := range nums {
			n, _ := strconv.Atoi(num)
			toAdd = append(toAdd, n)
		}
		result = append(result, toAdd)
		// result = append(result, scanner.Text())
	}
	return result
}
