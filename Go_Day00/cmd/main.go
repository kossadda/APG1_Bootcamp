// Copyright 2024 Gabilov Pervin. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package main is the entry point for the application.
package main

import (
	"flag"
	"fmt"
	"github.com/kossadda/APG1_Bootcamp/pkg/anscombe"
	"github.com/kossadda/APG1_Bootcamp/pkg/data"
	"sort"
)

// Flags represents the command-line flags for the program.
type Flags struct {
	Mean   bool // Flag to enable/disable mean calculation
	Median bool // Flag to enable/disable median calculation
	Mode   bool // Flag to enable/disable mode calculation
	StdDev bool // Flag to enable/disable standard deviation calculation
}

func main() {
	flags := readFlags()
	massive := data.NumberData()

	outResults(massive, &flags)
}

// readFlags reads the command-line flags and returns a Flags struct.
func readFlags() Flags {
	var flags Flags

	flag.BoolVar(&flags.Mean, "mean", true, "Mean value")
	flag.BoolVar(&flags.Median, "median", true, "Median value")
	flag.BoolVar(&flags.Mode, "mode", true, "Mode value")
	flag.BoolVar(&flags.StdDev, "stddev", true, "Standard deviation value")

	flag.Parse()

	return flags
}

// outResults prints the statistical results based on the provided flags.
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
