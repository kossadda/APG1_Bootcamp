// Package main is the entry point for the application.
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/kossadda/APG1_Bootcamp/Go_Day01/src/comparefs"
	"github.com/kossadda/APG1_Bootcamp/Go_Day01/src/readdb"
)

func main() {
	snapshot1 := flag.String("old", "", "Old filesystem base")
	snapshot2 := flag.String("new", "", "New filesystem base")
	flag.Parse()

	if *snapshot1 == "" || *snapshot2 == "" {
		fmt.Fprintln(os.Stderr, "Please provide both old and new filesystem snapshots.")
		os.Exit(1)
	}

	base1, err := Snapshot(*snapshot1)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}

	base2, err := Snapshot(*snapshot2)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}

	comparefs.Compare(base1, base2)
}

// Snapshot reads the filesystem snapshot from the given file path.
func Snapshot(path string) (map[string]bool, error) {
	file, err := readdb.DefineFile(nil, path)
	if err != nil {
		return nil, err
	}

	return comparefs.MapBase(file), nil
}
