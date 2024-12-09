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
