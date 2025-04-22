package main

import (
	"bufio"
	"os"
	"strconv"
)

type antenna struct {
	x  int
	y  int
	id string
}

func main() {
	w, l, antennasTypes := getInput()
	m := make(map[string]bool)

	for _, antennaType := range antennasTypes {
		for _, antenna := range antennaType {
			for _, otherAntena := range antennaType {

				if antenna.id == otherAntena.id {
					continue
				}

				var x int
				var y int

				if antenna.x > otherAntena.x {
					x = otherAntena.x - (antenna.x - otherAntena.x)
				} else if antenna.x < otherAntena.x {
					x = otherAntena.x + (otherAntena.x - antenna.x)
				} else {
					x = antenna.x
				}

				if antenna.y > otherAntena.y {
					y = otherAntena.y - (antenna.y - otherAntena.y)
				} else if antenna.y < otherAntena.y {
					y = otherAntena.y + (otherAntena.y - antenna.y)
				} else {
					y = antenna.y
				}

				if x < 0 || y < 0 || x >= w || y >= l {
					continue
				}

				m[strconv.Itoa(x)+"-"+strconv.Itoa(y)] = true
			}
		}
	}
	println(len(m))

}

func getInput() (int, int, map[rune][]antenna) {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	antannas := make(map[rune][]antenna)

	w, l := 0, 0

	for scanner.Scan() {
		line := scanner.Text()
		w = len(line)
		for i, char := range line {
			if char != '.' {
				_, ok := antannas[char]
				if ok {
					id := strconv.Itoa(i) + "-" + strconv.Itoa(l)
					antannas[char] = append(antannas[char], antenna{x: i, y: l, id: id})
				} else {
					id := strconv.Itoa(i) + "-" + strconv.Itoa(l)
					a := []antenna{}
					a = append(a, antenna{x: i, y: l, id: id})
					antannas[char] = a
				}
			}
		}
		l++
	}
	return w, l, antannas
}
