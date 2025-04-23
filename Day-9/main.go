package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// partOne()
	partTwo()

}

func partTwo() {

	input := getInput()
	// str := "0..111........22"
	// input := []rune("00...111...2...333.44.5555.6666.777.888899")

	// fmt.Println(string(input[:]))
	i, j := 0, len(input)-1
	result := 0

	fileLength := 1
	for {
		if j < 1 {
			break
		}
		if input[j] == input[j-1] {
			fileLength++
		}

		if input[j-1] != input[j] {
			free := 0
			i = 0
			for i < j {
				if input[i] == '.' {
					free++
				}
				if input[i] != '.' {
					free = 0
				}
				i++

				if fileLength <= free {
					r := j - 1 + fileLength
					l := i - free
					for range fileLength {
						input[l] = input[r]
						input[r] = '.'
						r--
						l++
					}
					free = 0
					fileLength = 1
					i = 0
					// fmt.Println(string(input[:]))
					break
				}
				if i >= j {
					free = 0
					fileLength = 1
					i = 0
					break
				}
			}
		}
		if i >= j {
			break
		}

		j--
	}

	for k := range input {
		if input[k] != '.' {
			result = result + k*int(input[k]-'0')
		}
	}
	fmt.Println(result)
}

func partOne() {

	input := getInput()
	i, j := 0, len(input)-1
	result := 0

	for {
		for input[i] != '.' {
			i++
		}
		for input[j] == '.' {
			j--
		}
		if i >= j {
			break
		}
		input[i] = input[j]
		input[j] = '.'
	}
	// fmt.Println(string(input))

	i = 0
	for input[i] != '.' {
		result = result + i*int(input[i]-'0')
		i++
	}
	fmt.Println(result)
}

func getInput() []rune {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	isFile := true
	id := 0
	result := []rune{}

	for scanner.Scan() {
		line := scanner.Text()
		for _, char := range line {
			l := int(char - '0')
			if isFile {
				for range l {
					result = append(result, rune('0'+id))
				}
				id++
			} else {
				for range l {
					result = append(result, '.')
				}
			}

			isFile = !isFile
		}
	}
	return result
}
