// Main package demonstrates the usage of the xargs utility.
// It reads arguments from stdin and executes the given command with those arguments.
package main

import (
	"fmt"
	"os"

	"github.com/kossadda/APG1_Bootcamp/pkg/xargs"
)

func main() {
	xg, err := xargs.New()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		return
	}

	ex, err := xg.Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", os.Args[0], err)
		return
	}

	fmt.Print(string(ex))
}
