package main

import (
	"fmt"
	"main/encode"
	"main/readdb"
	"os"
)

func main() {
	var reader readdb.DBReader
	var recipes encode.Recipes
	file, err := readdb.DefineFile(&reader)

	if err == nil {
		recipes, err = reader.Read(file)
	} else {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}

	reader.Print(recipes)
}
