package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type eqauation struct {
	total int
	nums  []int
}

func main() {
	equations := getInput()
	total := 0

	for _, e := range equations {

		if IsEqauatable(e.total, "MUL", 0, 0, e.nums) {
			total = total + e.total
		}
	}

	fmt.Println(total)

}

func IsEqauatable(total int, o string, c int, p int, nums []int) bool {

	if c > total {
		return false
	}

	if c == total && p >= len(nums) {
		return true
	}

	if c < total && p >= len(nums) {
		return false
	}

	if o == "ADD" {
		c = c + nums[p]
	}
	if o == "MUL" {
		if p == 0 {
			c = 1 * nums[p]
		} else {
			c = c * nums[p]
		}
	}
	if o == "CON" {
		if p == 0 {
			c = nums[p]
		} else {
			cs := strconv.Itoa(c)
			ns := strconv.Itoa(nums[p])
			cs = cs + ns
			cc, _ := strconv.Atoi(cs)
			c = cc
		}
	}

	p++

	if IsEqauatable(total, "ADD", c, p, nums) {
		return true
	} else if IsEqauatable(total, "MUL", c, p, nums) {
		return true
	} else {
		return IsEqauatable(total, "CON", c, p, nums)
	}

}

func getInput() []eqauation {
	file, _ := os.Open("input.txt")
	result := []eqauation{}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		i := strings.Index(line, ":")
		total, _ := strconv.Atoi(line[:i])
		numsString := strings.Split(line[i+2:], " ")
		nums := []int{}

		for _, n := range numsString {
			ni, _ := strconv.Atoi(n)
			nums = append(nums, ni)
		}

		result = append(result, eqauation{total: total, nums: nums})
	}

	return result

}
