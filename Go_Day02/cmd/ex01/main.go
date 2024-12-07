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
		fmt.Fprintln(os.Stderr, err)
		return
	}

	ch := wc.Output(args, w)

	for str := range ch {
		fmt.Println(str)
	}
}
