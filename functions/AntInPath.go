package functions

import (
	"fmt"
	"math"
)

func DistributeAnts(paths [][]string, numAnts int) [][]int {
	distribution := make([][]int, len(paths))
	pathLengths := make([]int, len(paths))
	for i, path := range paths {
		pathLengths[i] = len(path) - 1
	}

	// Distribute ants in a specific order
	for i := 1; i <= numAnts; i++ {
		bestPathIndex := 0
		bestArrivalTime := math.MaxInt32
		for j := range paths {
			arrivalTime := len(distribution[j]) + pathLengths[j]
			if arrivalTime < bestArrivalTime {
				bestPathIndex = j
				bestArrivalTime = arrivalTime
			}
		}
		distribution[bestPathIndex] = append(distribution[bestPathIndex], i)
	}
	fmt.Println(distribution)
	return distribution
}
