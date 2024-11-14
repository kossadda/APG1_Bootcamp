package comparefs

import (
	"fmt"
	"strings"
)

func Compare(first string, second string) {
	base1 := MapBase(first)
	base2 := MapBase(second)

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

func MapBase(base string) map[string]bool {
	mapBase := make(map[string]bool)

	lines := strings.Split(base, "\n")
	for _, line := range lines {
		if line != "" {
			mapBase[line] = true
		}
	}

	return mapBase
}
