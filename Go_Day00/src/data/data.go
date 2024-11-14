// Package data provides functions for reading and processing input data.
package data

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

const (
	MinValue = -100000 // Minimum value for input numbers
	MaxValue = 100000  // Maximum value for input numbers
)

// NumberData reads a sequence of numbers from standard input and returns them as a slice of integers.
func NumberData() (numbers []int) {
	reader := bufio.NewReader(os.Stdin)

	for {
		input, err := reader.ReadString('\n')
		if err == io.EOF {
			fmt.Println("Stopped")
			break
		} else if err != nil {
			fmt.Fprintln(os.Stderr, "Error when entering a number:", err)
			continue
		}

		value, err := convertValue(input)

		if err != nil {
			fmt.Fprintln(os.Stderr, "Error when entering a number:", err)
		} else {
			if value >= MinValue && value <= MaxValue {
				numbers = append(numbers, value)
			} else {
				err := fmt.Errorf("number must be in the range [-100000:100000]")
				fmt.Fprintln(os.Stderr, "Error:", err)
				continue
			}
		}
	}

	return numbers
}

// convertValue converts a string input to an integer.
func convertValue(input string) (result int, err error) {
	input = input[:len(input)-1]

	result, err = strconv.Atoi(input)

	return result, err
}
