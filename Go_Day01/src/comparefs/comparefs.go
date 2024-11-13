package comparefs

import (
	"fmt"
	"os"
	"strings"
)

func Compare(first string, second string) {
	base_1, err := getFilebase(first)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error reading first file:", err)
		return
	}

	base_2, err := getFilebase(second)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error reading second file:", err)
		return
	}

	for path := range base_2 {
		if !base_1[path] {
			fmt.Printf("ADDED %s\n", path)
		}
	}

	for path := range base_1 {
		if !base_2[path] {
			fmt.Printf("REMOVED %s\n", path)
		}
	}
}

func getFilebase(path string) (map[string]bool, error) {
	byte_base, err := os.ReadFile(path)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		return nil, err
	}

	base := string(byte_base)
	map_base := make(map[string]bool)

	lines := strings.Split(base, "\n")
	for _, line := range lines {
		map_base[line] = true
	}

	return map_base, nil
}
