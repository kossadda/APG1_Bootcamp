package main

import (
	"fmt"
	"github.com/kossadda/APG1_Bootcamp/Go_Day01/src/ex00/encode"
	"github.com/kossadda/APG1_Bootcamp/Go_Day01/src/ex00/readdb"
	"os"
)

func main() {
	var reader readdb.DBReader
	var recipes encode.Recipes
	file, err := readdb.DefineFile(&reader)

	if err == nil {
		recipes, err = reader.Read(file)
	}

	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}

	reader.Print(recipes)
}
