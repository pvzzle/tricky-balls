package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n int = 0
	fmt.Fscan(reader, &n)

	totalBallsInContainer := make([]int64, n)
	totalBallsOfSameColor := make([]int64, n)

	for i := range n {
		for j := range n {
			var value int64 = 0
			fmt.Fscan(reader, &value)

			totalBallsInContainer[i] += value
			totalBallsOfSameColor[j] += value
		}
	}

	slices.Sort(totalBallsInContainer)
	slices.Sort(totalBallsOfSameColor)

	for i := range n {
		if totalBallsInContainer[i] != totalBallsOfSameColor[i] {
			fmt.Println("no")
			return
		}
	}

	fmt.Println("yes")
}
