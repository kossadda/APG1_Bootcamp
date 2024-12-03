package main

import (
	"fmt"
	"github.com/kossadda/APG1_Bootcamp/pkg/find"
	"os"

	"github.com/kossadda/APG1_Bootcamp/pkg/param"
)

func main() {
	prm, err := param.New("test", os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	err = find.Scan(prm)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
