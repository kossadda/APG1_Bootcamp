// This program implements a simple utility to count words, lines, or characters in a given text file.
// It uses flags (-w, -l, -m) to determine the type of count and processes the files concurrently using goroutines.
package main

import (
	"fmt"
	"os"

	"github.com/kossadda/APG1_Bootcamp/pkg/wc"
)

func main() {
	args := os.Args[1:]
	w, err := wc.New("wc", &args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", os.Args[0], err)
		return
	}

	ch := wc.Output(args, w)

	for str := range ch {
		fmt.Println(str)
	}
}
