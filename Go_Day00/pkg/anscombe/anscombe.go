// Package anscombe provides functions for calculating statistical metrics.
package anscombe

import "math"

// Mean calculates the mean of a slice of integers.
func Mean(massive []int) (mean float64) {
	for _, num := range massive {
		mean += float64(num)
	}

	return mean / float64(len(massive))
}

// Median calculates the median of a sorted slice of integers.
func Median(massive []int) (median float64) {
	length := len(massive)

	if length%2 == 0 {
		median = float64(massive[length/2-1]+massive[length/2]) / 2
	} else {
		median = float64(massive[length/2])
	}

	return median
}

// Mode calculates the mode of a slice of integers.
func Mode(massive []int) (mode float64) {
	repeatMax := 0
	repeatCnt := 0
	prev := massive[0]
	mode = float64(massive[0])

	for i, value := range massive {
		if i > 0 && value == prev {
			repeatCnt++
		} else {
			if repeatCnt > repeatMax {
				repeatMax = repeatCnt
				mode = float64(prev)
			}

			repeatCnt = 0
		}

		prev = value
	}

	if repeatCnt > repeatMax {
		mode = float64(massive[len(massive)-1])
	}

	return mode
}

// Deviation calculates the standard deviation of a slice of integers.
func Deviation(massive []int) (stddev float64) {
	mean := Mean(massive)

	for _, value := range massive {
		stddev += math.Pow(float64(value)-mean, 2)
	}

	stddev = math.Sqrt(stddev / float64(len(massive)))

	return stddev
}
