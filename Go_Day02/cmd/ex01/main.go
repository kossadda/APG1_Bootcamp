package main

import (
	"fmt"
	"github.com/kossadda/APG1_Bootcamp/pkg/wc"
	"os"
)

func main() {
	args := os.Args[1:]
	w, err := wc.New("wc", &args)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	wc.Output(args, w)
}
