package comparefs

import (
	"fmt"
	"os"
	"strings"
)

func Compare(first string, second string) {
	base1, err := getFilebase(first)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error reading first file:", err)
		return
	}

	base2, err := getFilebase(second)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error reading second file:", err)
		return
	}

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

func getFilebase(path string) (map[string]bool, error) {
	byteBase, err := os.ReadFile(path)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		return nil, err
	}

	base := string(byteBase)
	mapBase := make(map[string]bool)

	lines := strings.Split(base, "\n")
	for _, line := range lines {
		mapBase[line] = true
	}

	return mapBase, nil
}
