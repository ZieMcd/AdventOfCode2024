package main

import (
	"bufio"
	// "fmt"
	"os"
)

const xmas = "XMAS"

func main() {

	input := getInput()
	result := 0
	for i := 1; i < len(input)-1; i++ {
		for j := 1; j < len(input[i])-1; j++ {
			if input[i][j] == 'A' {
				if (input[i+1][j-1] == 'M' && input[i-1][j+1] == 'S') || (input[i+1][j-1] == 'S' && input[i-1][j+1] == 'M') {
					if (input[i-1][j-1] == 'M' && input[i+1][j+1] == 'S') || (input[i-1][j-1] == 'S' && input[i+1][j+1] == 'M') {
						result++
					}
				}
			}
		}
	}
	print(result)
}

func part1() {
	result := 0
	input := getInput()
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			result = find_xmas(input, "r", i, j, result, 0)
			result = find_xmas(input, "l", i, j, result, 0)
			result = find_xmas(input, "d", i, j, result, 0)
			result = find_xmas(input, "u", i, j, result, 0)
			result = find_xmas(input, "ur", i, j, result, 0)
			result = find_xmas(input, "ul", i, j, result, 0)
			result = find_xmas(input, "dr", i, j, result, 0)
			result = find_xmas(input, "dl", i, j, result, 0)
		}
	}
	print(result)

}
func find_xmas(input [][]rune, dir string, y int, x int, result int, i int) int {
	want := rune(xmas[i])
	item := input[y][x]
	// fmt.Printf("dir: %s, y: %d, x: %d, want: %c,  item: %c \n", dir, y, x, want, item)

	if item == want && item == 'S' {
		return result + 1
	}

	if item != want {
		return result
	}

	switch dir {
	case "r":
		x++
	case "l":
		x--
	case "u":
		y++
	case "d":
		y--
	case "ur":
		x++
		y++
	case "ul":
		x--
		y++
	case "dr":
		x++
		y--
	case "dl":
		x--
		y--
	}

	if y < 0 || y >= len(input) || x < 0 || x >= len(input[y]) {
		return result
	}

	i++
	return find_xmas(input, dir, y, x, result, i)
}

func getInput() [][]rune {

	file, err := os.Open("input.txt")
	result := [][]rune{}

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		result = append(result, []rune(scanner.Text()))
	}
	return result
}
