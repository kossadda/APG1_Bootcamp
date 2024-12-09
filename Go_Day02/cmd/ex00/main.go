// This program provides the main logic for scanning a directory
// based on various filters specified by command-line flags. The package
// uses the 'param' package to parse flags and the 'find' package to scan
// the directory. It then prints the results to the standard output.
package main

import (
	"fmt"
	"os"

	"github.com/kossadda/APG1_Bootcamp/pkg/find/find"
	"github.com/kossadda/APG1_Bootcamp/pkg/find/param"
)

func main() {
	prm, err := param.New(os.Args[1:])
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", os.Args[0], err)
		return
	}

	sys, err := find.Scan(prm)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", os.Args[0], err)
		return
	}

	for i := range sys {
		fmt.Println(sys[i])
	}
}
