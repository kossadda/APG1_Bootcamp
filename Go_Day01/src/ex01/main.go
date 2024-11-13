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
	old_path := flag.String("old", "", "Original database (xml or json)")
	new_path := flag.String("new", "", "Stolen database (xml or json)")
	flag.Parse()

	comparedb.CompareRecipes(getRecipe(old_path), getRecipe(new_path))
}

func getRecipe(path *string) (rec recipes.Recipes) {
	var reader readdb.DBReader
	new_file, err := readdb.DefineFile(&reader, path)

	if err == nil {
		rec, err = reader.Read(new_file)
	}

	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}

	return rec
}
