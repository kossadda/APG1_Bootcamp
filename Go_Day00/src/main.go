package main

import (
	"flag"
	"fmt"
	"github.com/kossadda/APG1_Bootcamp/Go_Day00/src/anscombe"
	"github.com/kossadda/APG1_Bootcamp/Go_Day00/src/data"
	"sort"
)

type Flags struct {
	Mean   bool
	Median bool
	Mode   bool
	StdDev bool
}

func main() {
	flags := readFlags()
	massive := data.NumberData()

	outResults(massive, &flags)
}

func readFlags() Flags {
	var flags Flags

	flag.BoolVar(&flags.Mean, "mean", true, "Mean value")
	flag.BoolVar(&flags.Median, "median", true, "Median value")
	flag.BoolVar(&flags.Mode, "mode", true, "Mode value")
	flag.BoolVar(&flags.StdDev, "stddev", true, "Standart deviation value")

	flag.Parse()

	return flags
}

func outResults(massive []int, flags *Flags) {
	fmt.Printf("\nSequence of numbers: %v\n", massive)

	sort.Ints(massive)

	if len(massive) > 0 {
		if flags.Mean {
			fmt.Printf("Mean: %.2f\n", anscombe.Mean(massive))
		}

		if flags.Median {
			fmt.Printf("Median: %.2f\n", anscombe.Median(massive))
		}

		if flags.Mode {
			fmt.Printf("Mode: %.2f\n", anscombe.Mode(massive))
		}

		if flags.StdDev {
			fmt.Printf("SD: %.2f\n", anscombe.Deviation(massive))
		}
	}
}
