// Package main is the entry point for the archiving program. It parses command-line arguments,
// creates an Archiver instance, and uses it to rotate (archive) the specified files into tar.gz
// archives. Any errors during the archiving process are printed to the standard error output.
package main

import (
	"fmt"
	"os"

	"github.com/kossadda/APG1_Bootcamp/pkg/archiver"
)

func main() {
	args := os.Args[1:]
	arch, err := archiver.New(&args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", os.Args[0], err)
		return
	}

	arch.RotateFiles(args)
}
