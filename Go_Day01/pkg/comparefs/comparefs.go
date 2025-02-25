// Package comparefs provides functions for comparing filesystem snapshots.
package comparefs

import (
	"fmt"
	"strings"
)

// Compare compares two filesystem snapshots and prints the differences.
func Compare(base1, base2 map[string]bool) {
	for path := range base2 {
		if !base1[path] {
			fmt.Printf("ADDED %s\n", path)
		}
	}

	for path := range base1 {
		if !base2[path] {
			fmt.Printf("REMOVED %s\n", path)
		}
	}
}

// MapBase creates a map of file paths from the given filesystem snapshot.
func MapBase(base []byte) map[string]bool {
	mapBase := make(map[string]bool)

	lines := strings.Split(string(base), "\n")
	for _, line := range lines {
		if line != "" {
			mapBase[line] = true
		}
	}

	return mapBase
}
