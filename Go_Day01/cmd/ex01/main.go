// Package main is the entry point for the application.
package main

import (
	"flag"
	"log"

	"github.com/kossadda/APG1_Bootcamp/pkg/comparedb"
	"github.com/kossadda/APG1_Bootcamp/pkg/readdb"
	"github.com/kossadda/APG1_Bootcamp/pkg/recipes"
)

func main() {
	oldPath := flag.String("old", "", "Original database (xml or json)")
	newPath := flag.String("new", "", "Stolen database (xml or json)")
	flag.Parse()

	oldRecipe, err := Recipe(*oldPath)
	if err != nil {
		log.Fatal(err)
		return
	}

	newRecipe, err := Recipe(*newPath)
	if err != nil {
		log.Fatal(err)
		return
	}

	comparedb.Compare(oldRecipe, newRecipe)
}

// Recipe reads the recipe data from the given file path.
func Recipe(path string) (recipes.Recipes, error) {
	var reader readdb.DBReader

	newFile, err := readdb.DefineFile(&reader, path)
	if err != nil {
		return recipes.Recipes{}, err
	}

	err = reader.DBRead(newFile)
	if err != nil {
		return recipes.Recipes{}, err
	}

	return reader.Recipe(), nil
}
