package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func main() {
	pos := 0
	do_pos := 0
	dont_pos := 0
	num1 := ""
	num2 := ""
	input := getInput()
	reset := false
	result := 0
	do := true

	for _, char := range input {
		do, do_pos = find_do(do_pos, char, do)
		do, dont_pos = find_dont(dont_pos, char, do)

		if reset {
			pos = 0
			num1 = ""
			num2 = ""
		}

		if do {
			reset, result, pos, num1, num2 = cal(pos, num1, num2, char, result)
		}
		// fmt.Println(result)
	}
	fmt.Println(result)
}

func find_do(pos int, char rune, do bool) (bool, int) {
	switch pos {
	case 0:
		if char == 'd' {
			pos++
			return do, pos
		} else {
			return do, 0
		}
	case 1:
		if char == 'o' {
			pos++
			return do, pos
		} else {
			return do, 0
		}
	case 2:
		if char == '(' {
			pos++
			return do, pos
		} else {
			return do, 0
		}
	case 3:
		if char == ')' {
			return true, 0
		} else {
			return do, 0
		}
	default:
		panic("aaa")
	}
}

func find_dont(pos int, char rune, do bool) (bool, int) {
	switch pos {
	case 0:
		if char == 'd' {
			pos++
			return do, pos
		} else {
			return do, 0
		}
	case 1:
		if char == 'o' {
			pos++
			return do, pos
		} else {
			return do, 0
		}
	case 2:
		if char == 'n' {
			pos++
			return do, pos
		} else {
			return do, 0
		}
	case 3:
		if char == '\'' {
			pos++
			return do, pos
		} else {
			return do, pos
		}
	case 4:
		if char == 't' {
			pos++
			return do, pos
		} else {
			return do, 0
		}
	case 5:
		if char == '(' {
			pos++
			return do, pos
		} else {
			return do, 0
		}
	case 6:
		if char == ')' {
			return false, 0
		} else {
			return do, 0
		}
	default:
		panic("aaa")
		// return do, pos
	}
	// return do, pos
}
func cal(pos int, num1 string, num2 string, char rune, result int) (bool, int, int, string, string) {
	var reset bool
	switch pos {
	case 0:
		if char == 'm' {
			pos = pos + 1
			reset = false
		} else {
			reset = true
		}
	case 1:
		if char == 'u' {
			pos++
			reset = false
		} else {
			reset = true
		}
	case 2:
		if char == 'l' {
			pos++
			reset = false
		} else {
			reset = true
		}
	case 3:
		if char == '(' {
			pos++
			reset = false
		} else {
			reset = true
		}
	case 4:
		if unicode.IsDigit(char) {
			if len(num1) < 3 {
				num1 += string(char)
				reset = false
			}
			if len(num1) == 3 {
				pos++
				reset = false
			}
			if len(num1) > 3 {
				reset = true
			}
		} else if char == ',' {
			pos = 6
			reset = false
		} else {
			reset = true
		}

	case 5:
		if char == ',' {
			pos++
			reset = false
		} else {
			reset = true
		}
	case 6:
		if unicode.IsDigit(char) {
			if len(num2) < 3 {
				num2 += string(char)
				reset = false
			}
			if len(num2) == 3 {
				pos++
				reset = false
			}
			if len(num2) > 3 {
				reset = true
			}
		} else if char == ')' {
			i, _ := strconv.Atoi(num1)
			y, _ := strconv.Atoi(num2)
			result = result + (i * y)
			reset = true
		} else {
			reset = true
		}
	case 7:
		if char == ')' {
			reset = true
			i, _ := strconv.Atoi(num1)
			y, _ := strconv.Atoi(num2)
			result = result + (i * y)
		}
		reset = true
	}
	return reset, result, pos, num1, num2
}
func getInput() string {

	file, err := os.Open("input.txt")
	result := ""

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		result += scanner.Text()
	}
	return result

}
