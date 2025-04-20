package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
)

type Pos struct {
	x int
	y int
}

func main() {

	mp, _ := getInput()
	marked := []Pos{}
	start_pos := Pos{x: 0, y: 0}
	for y, r := range mp {
		for x, c := range r {
			if c == '^' {
				var total int
				marked, total = followRoute(x, y, "UP", mp, marked, 0)
				start_pos.x = x
				start_pos.y = y
				fmt.Println(total)
			}
		}
		// x := strings.Index(r, "^")
		// if x > -1 {
		// 	var total int
		// 	marked, total = followRoute(x, y, "UP", route, marked, 0)
		// 	fmt.Println(total)
		// }
	}

	b := 0
	for i, r := range marked {
		if i != 0 {
			mp[r.y][r.x] = '#'
			if followRoutePart2(start_pos.x, start_pos.y, "UP", mp, []Pos{}, []string{}, 0) {
				b++
			}

			mp[r.y][r.x] = '.'
		}
	}
	fmt.Println(b)

}

func followRoutePart2(x int, y int, dir string, route [][]rune, marked []Pos, markedDir []string, total int) bool {

	pos := Pos{x: x, y: y}
	if !slices.Contains(marked, pos) {
		marked = append(marked, pos)
		total++
	}

	if slices.Contains(markedDir, strconv.Itoa(x)+"-"+strconv.Itoa(y)+"-"+dir) {
		return true
	}

	markedDir = append(markedDir, strconv.Itoa(x)+"-"+strconv.Itoa(y)+"-"+dir)
	// if !slices

	if dir == "UP" {
		if y-1 < 0 {
			return false
		}
		if route[y-1][x] == '#' {
			dir = "RIGHT"
			return followRoutePart2(x, y, dir, route, marked, markedDir, total)
		}
		y -= 1
		return followRoutePart2(x, y, dir, route, marked, markedDir, total)
	}

	if dir == "DOWN" {
		if y+1 >= len(route) {
			return false
		}
		if route[y+1][x] == '#' {
			dir = "LEFT"
			return followRoutePart2(x, y, dir, route, marked, markedDir, total)
		}
		y += 1
		return followRoutePart2(x, y, dir, route, marked, markedDir, total)
	}

	if dir == "RIGHT" {
		if x+1 >= len(route[0]) {
			return false
		}
		if route[y][x+1] == '#' {
			dir = "DOWN"
			return followRoutePart2(x, y, dir, route, marked, markedDir, total)
		}
		x += 1
		return followRoutePart2(x, y, dir, route, marked, markedDir, total)
	}

	if dir == "LEFT" {
		if x-1 < 0 {
			return false
		}
		if route[y][x-1] == '#' {
			dir = "UP"
			return followRoutePart2(x, y, dir, route, marked, markedDir, total)
		}
		x -= 1
		return followRoutePart2(x, y, dir, route, marked, markedDir, total)
	}

	panic("no dir")
}

func followRoute(x int, y int, dir string, route [][]rune, marked []Pos, total int) ([]Pos, int) {

	pos := Pos{x: x, y: y}
	if !slices.Contains(marked, pos) {
		marked = append(marked, pos)
		total++
	}

	if dir == "UP" {
		if y-1 < 0 {
			return marked, total
		}
		if route[y-1][x] == '#' {
			dir = "RIGHT"
			return followRoute(x, y, dir, route, marked, total)
		}
		y -= 1
		return followRoute(x, y, dir, route, marked, total)
	}

	if dir == "DOWN" {
		if y+1 >= len(route) {
			return marked, total
		}
		if route[y+1][x] == '#' {
			dir = "LEFT"
			return followRoute(x, y, dir, route, marked, total)
		}
		y += 1
		return followRoute(x, y, dir, route, marked, total)
	}

	if dir == "RIGHT" {
		if x+1 >= len(route[0]) {
			return marked, total
		}
		if route[y][x+1] == '#' {
			dir = "DOWN"
			return followRoute(x, y, dir, route, marked, total)
		}
		x += 1
		return followRoute(x, y, dir, route, marked, total)
	}

	if dir == "LEFT" {
		if x-1 < 0 {
			return marked, total
		}
		if route[y][x-1] == '#' {
			dir = "UP"
			return followRoute(x, y, dir, route, marked, total)
		}
		x -= 1
		return followRoute(x, y, dir, route, marked, total)
	}

	panic("no dir")
}

func getInput() ([][]rune, error) {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var result [][]rune
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		runeSlice := []rune(line)
		result = append(result, runeSlice)
	}
	return result, scanner.Err()
}
