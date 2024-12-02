package main

import (
	"fmt"
	"os"

	"github.com/kossadda/APG1_Bootcamp/pkg/param"
)

func main() {
	prm, err := param.New()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	fmt.Println(prm)

	fmt.Println(prm.IsSetSl())
	fmt.Println(prm.IsSetD())
	fmt.Println(prm.IsSetF())
	fmt.Println(prm.IsSetExt())
}
