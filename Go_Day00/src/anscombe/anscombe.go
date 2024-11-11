package anscombe

import "math"

func GetMean(massive []int) (mean float64) {
	for _, num := range massive {
		mean += float64(num)
	}

	return mean / float64(len(massive))
}

func GetMedian(massive []int) (median float64) {
	length := len(massive)

	if length%2 == 0 {
		median = float64(massive[length/2-1]+massive[length/2]) / 2
	} else {
		median = float64(massive[length/2])
	}

	return median
}

func GetMode(massive []int) (mode float64) {
	repeat_max := 0
	repeat_cnt := 0
	prev := massive[0]
	mode = float64(massive[0])

	for i, value := range massive {
		if i > 0 && value == prev {
			repeat_cnt++
		} else {
			if repeat_cnt > repeat_max {
				repeat_max = repeat_cnt
				mode = float64(prev)
			}

			repeat_cnt = 0
		}

		prev = value
	}

	if repeat_cnt > repeat_max {
		mode = float64(massive[len(massive)-1])
	}

	return mode
}

func GetDeviation(massive []int) (stddev float64) {
	mean := GetMean(massive)

	for _, value := range massive {
		stddev += math.Pow(float64(value)-mean, 2)
	}

	stddev = math.Sqrt(stddev / float64(len(massive)))

	return stddev
}
