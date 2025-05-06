package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func canSort(matrix [][]int64) bool {
	n := len(matrix)
	totalBallsInContainer := make([]int64, n)
	totalBallsOfSameColor := make([]int64, n)

	for i := range n {
		for j := range n {
			totalBallsInContainer[i] += matrix[i][j]
			totalBallsOfSameColor[j] += matrix[i][j]
		}
	}

	slices.Sort(totalBallsInContainer)
	slices.Sort(totalBallsOfSameColor)

	for i := range n {
		if totalBallsInContainer[i] != totalBallsOfSameColor[i] {
			return false
		}
	}

	return true
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n int = 0
	fmt.Fscan(reader, &n)

	matrix := make([][]int64, n)
	for i := range matrix {
		matrix[i] = make([]int64, n)
		for j := range matrix[i] {
			fmt.Fscan(reader, &matrix[i][j])
		}
	}

	if canSort(matrix) {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
}
