package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/kossadda/APG1_Bootcamp/Go_Day01/src/readdb"
	"github.com/kossadda/APG1_Bootcamp/Go_Day01/src/recipes"
)

func main() {
	var reader readdb.DBReader
	var recipes recipes.Recipes
	filename := flag.String("f", "", "Filename to read ('xml' or 'json')")
	flag.Parse()

	file, err := readdb.DefineFile(&reader, *filename)

	if err == nil {
		recipes, err = reader.Read(file)
	}

	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}

	reader.Print(recipes)
}
