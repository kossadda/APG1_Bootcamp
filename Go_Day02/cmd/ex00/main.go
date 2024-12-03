package main

import (
	"fmt"
	"os"

	"github.com/kossadda/APG1_Bootcamp/pkg/param"
)

func main() {
	args := []string{"-f", "-ext", "", "/home"}

	prm, err := param.New("test", args)
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
