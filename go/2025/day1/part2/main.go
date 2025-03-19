package main

import (
	inputs "advent-of-code"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	data := inputs.GetInput(2024, 1)
	lines := strings.Split(data, "\n")

	leftList := make([]int, 0)
	rightListNumFreq := make(map[int]int)

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
		rightListNumFreq[rightNum]++
	}

	similaritySum := 0
	for _, leftNum := range leftList {
		similarityScore := leftNum * rightListNumFreq[leftNum]
		similaritySum += similarityScore
	}
	fmt.Println("Answer:", similaritySum)
}
