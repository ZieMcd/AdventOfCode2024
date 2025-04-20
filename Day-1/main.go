package main

import (
	"bufio"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	lines, err := getInput()

	if err != nil {
		panic(err)
	}
	// println(dayOne(lines))
	println(dayTwo(lines))
}

// Super inefficient. T(N) = (O)N^2
// could maybe fist get numbers from nums2 and put it in a hash, but to lazy
func dayTwo(lines [2][]int) int {
	result := 0
	var nums1 = lines[0]
	var nums2 = lines[1]

	for _, num := range nums1 {
		xnum := 0
		for _, num2 := range nums2 {
			if num == num2 {
				xnum += 1
			}
		}
		result += num * xnum
	}
	return result
}

func dayOne(lines [2][]int) int {
	result := 0

	var nums1 = lines[0]
	var nums2 = lines[1]

	sort.Ints(nums1)
	sort.Ints(nums2)

	for i := 0; i < len(nums1); i++ {
		dif := math.Abs(float64(nums1[i] - nums2[i]))

		result += int(dif)
	}
	return result
}

func getInput() ([2][]int, error) {

	file, err := os.Open("input.txt")

	if err != nil {
		panic(err)
	}

	var result [2][]int

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		nums := strings.Fields(scanner.Text())
		num1, _ := strconv.Atoi(nums[0])
		num2, _ := strconv.Atoi(nums[1])
		result[0] = append(result[0], num1)
		result[1] = append(result[1], num2)
	}

	return result, scanner.Err()

}
