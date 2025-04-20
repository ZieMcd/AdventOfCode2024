package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	result := 0
	orders := get_orders()

	afterRules, _ := get_rules()

	for _, pages := range orders {
		// result = result + isValid(pages, afterRules, beforeRules)
		result = result + correct(pages, afterRules)

	}
	print(result)
}

func isValid(pages []string, afterRules map[string][]string, beforeRule map[string][]string) int {

	for i := range len(pages) {
		before := pages[:i]
		// after := pages[i+1:]

		x := pages[i]
		mustBeAfter := afterRules[x]
		// allowedBefore := beforeRule[strconv.Itoa(i)]

		for _, b := range before {
			if slices.Contains(mustBeAfter, b) {
				fmt.Printf("%v not valid \n", pages)
				return 0
			}
		}

		// for _, a := range after {
		// 	if slices.Contains(allowedBefore, a) {
		// 		fmt.Printf("%v not valid \n", pages)
		// 		return 0
		// 	}
		// }

	}

	num, _ := strconv.Atoi(pages[len(pages)/2])

	fmt.Printf("%v valid \n", pages)
	return num
}

func correct(pages []string, afterRules map[string][]string) int {
	fmt.Printf("started %v \n", pages)
	isFucked := false
	i := 0
	for {
		hasBeenSwaped := false
		x := pages[i]
		mustBeAfter := afterRules[x]

		for j := i - 1; j >= 0; j-- {
			n := pages[j]
			if slices.Contains(mustBeAfter, n) {
				pages[j] = x
				pages[i] = n
				isFucked = true
				hasBeenSwaped = true
				break
			}
		}

		if hasBeenSwaped == true {

			fmt.Printf("edited %v \n", pages)
			i = 0
		} else {
			i++
		}

		if i >= len(pages) {
			break
		}
	}

	if isFucked {
		fmt.Printf("%v\n", pages)
		num, _ := strconv.Atoi(pages[len(pages)/2])
		return num
	}
	return 0
}

func get_orders() [][]string {
	result := [][]string{}
	file, _ := os.Open("orders.txt")
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		a := strings.Split(scanner.Text(), ",")
		result = append(result, a)
	}
	return result
}

func get_rules() (map[string][]string, map[string][]string) {
	afterRules := make(map[string][]string)
	beforeRules := make(map[string][]string)
	file, err := os.Open("rules.txt")

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		before := line[0:2]
		after := line[3:]

		if afterRules[before] == nil {
			afterRules[before] = []string{after}
		} else {
			afterRules[before] = append(afterRules[before], after)
		}

		if beforeRules[after] == nil {
			beforeRules[after] = []string{before}
		} else {
			beforeRules[after] = append(beforeRules[after], before)
		}

	}
	return afterRules, beforeRules

}
