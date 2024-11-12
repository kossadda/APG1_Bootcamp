package main

import (
	"flag"
	"fmt"
	"main/encode"
	"main/pathReader"
	"os"
)

type DBReader interface {
	Read(file []byte) (encode.Recipes, error)
	Print(recipes encode.Recipes)
}

func main() {
	file, reader := defineFile()

	recipes, err := reader.Read(file)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}

	reader.Print(recipes)
}

func defineFile() (file []byte, reader DBReader) {
	filename := flag.String("f", "", "Filename to read (xml or json)")
	flag.Parse()

	file, ext, err := pathReader.PathRead(filename)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}

	if ext == "xml" {
		reader = &encode.RecipesXML{}
	} else {
		reader = &encode.RecipesJSON{}
	}

	return file, reader
}
