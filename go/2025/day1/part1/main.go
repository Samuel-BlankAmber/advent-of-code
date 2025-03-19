package main

import (
	inputs "advent-of-code"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	data := inputs.GetInput(2024, 1)
	lines := strings.Split(data, "\n")

	leftList := make([]int, 0)
	rightList := make([]int, 0)

	for _, line := range lines {
		if line == "" {
			continue
		}

		nums := strings.Split(line, "   ")
		leftNum, err := strconv.Atoi(nums[0])
		if err != nil {
			fmt.Println(err)
		}
		rightNum, err := strconv.Atoi(nums[1])
		if err != nil {
			fmt.Println(err)
		}

		leftList = append(leftList, leftNum)
		rightList = append(rightList, rightNum)
	}

	sort.Ints(leftList)
	sort.Ints(rightList)

	distanceSum := 0
	for i := 0; i < len(leftList); i++ {
		leftNum := leftList[i]
		rightNum := rightList[i]

		distance := 0
		if rightNum > leftNum {
			distance = rightNum - leftNum
		} else {
			distance = leftNum - rightNum
		}
		distanceSum += distance
	}

	fmt.Println("Answer:", distanceSum)
}
