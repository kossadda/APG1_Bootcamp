// This program provides the main logic for scanning a directory
// based on various filters specified by command-line flags. The package
// uses the 'param' package to parse flags and the 'find' package to scan
// the directory. It then prints the results to the standard output.
package main

import (
	"fmt"
	"os"

	"github.com/kossadda/APG1_Bootcamp/pkg/find"
	"github.com/kossadda/APG1_Bootcamp/pkg/param"
)

func main() {
	prm, err := param.New("test", os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	sys, err := find.Scan(prm)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	for i := range sys {
		fmt.Println(sys[i])
	}
}
