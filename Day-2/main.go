package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// partOne()
	partTwo()
}

func partTwo() {
	reports, _ := getInput()
	fmt.Println(len(reports))
	result := 0

	for _, report := range reports {
		isSafe := isReportValid(report, true)
		fmt.Printf("report: %v, \n is valid: %t\n", report, isSafe)

		if isSafe {
			result++
		}
	}
	fmt.Println(result)
}

func partOne() {
	reports, _ := getInput()
	result := 0
	for _, report := range reports {
		isSafe := isReportValid(report, false)
		fmt.Printf("report: %v, \n is valid: %t\n", report, isSafe)

		if isSafe {
			result++
		}
	}
	fmt.Println(result)
}

func getDirection(lvl int, nLvl int) string {
	var newDirection string
	if lvl < nLvl {
		newDirection = "up"
	}
	if lvl > nLvl {
		newDirection = "down"
	}
	if lvl == nLvl {
		newDirection = "none"
	}
	return newDirection
}

func isReportValid(report []int, tolerate bool) bool {
	var direction string
	var newDirection string

	isValid := true

	for i := 0; i < len(report)-1; i++ {
		lvl := report[i]
		nLvl := report[i+1]

		newDirection = getDirection(lvl, nLvl)

		if i == 0 {
			direction = newDirection
		}

		if newDirection == "none" {
			if tolerate {
				return isReportValid(replicate(report, i+1), false) || isReportValid(replicate(report, i), false)
			} else {
				return false
			}
		}

		if i != 0 {
			if newDirection != direction {
				if tolerate {
					return isReportValid(replicate(report, i+1), false) || isReportValid(replicate(report, i), false)
				} else {
					return false
				}

			}
		}

		if newDirection == "up" && nLvl-lvl > 3 {
			if tolerate {
				return isReportValid(replicate(report, i+1), false) || isReportValid(replicate(report, i), false)
			} else {
				return false
			}

		}

		if newDirection == "down" && lvl-nLvl > 3 {
			if tolerate {
				return isReportValid(replicate(report, i+1), false) || isReportValid(replicate(report, i), false)
			} else {
				return false
			}

		}
	}
	return isValid
}

func getInput() ([][]int, error) {

	file, err := os.Open("input.txt")

	if err != nil {
		panic(err)
	}

	var result [][]int

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		sNums := strings.Fields(scanner.Text())
		var nums []int
		for _, sNum := range sNums {
			n, _ := strconv.Atoi(sNum)
			nums = append(nums, n)
		}

		result = append(result, nums)
	}

	return result, scanner.Err()

}
func replicate(report []int, i int) []int {
	dst := make([]int, len(report))
	copy(dst, report)
	dst = append(dst[:i], dst[i+1:]...)
	return dst
}
