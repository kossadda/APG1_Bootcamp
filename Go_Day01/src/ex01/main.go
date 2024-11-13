package main

import (
	"flag"
	"fmt"
	"github.com/kossadda/APG1_Bootcamp/Go_Day01/src/comparedb"
	"github.com/kossadda/APG1_Bootcamp/Go_Day01/src/readdb"
	"github.com/kossadda/APG1_Bootcamp/Go_Day01/src/recipes"
	"os"
)

func main() {
	oldPath := flag.String("old", "", "Original database (xml or json)")
	newPath := flag.String("new", "", "Stolen database (xml or json)")
	flag.Parse()

	comparedb.Compare(getRecipe(oldPath), getRecipe(newPath))
}

func getRecipe(path *string) (rec recipes.Recipes) {
	var reader readdb.DBReader
	newFile, err := readdb.DefineFile(&reader, path)

	if err == nil {
		rec, err = reader.Read(newFile)
	}

	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}

	return rec
}
