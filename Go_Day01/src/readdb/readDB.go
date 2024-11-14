// Package readdb provides functions for reading and processing database files.
package readdb

import (
	"fmt"
	"os"

	"github.com/kossadda/APG1_Bootcamp/Go_Day01/src/encode"
	"github.com/kossadda/APG1_Bootcamp/Go_Day01/src/recipes"
)

// DBReader is an interface for reading database files.
type DBReader interface {
	Read(file []byte) (recipes.Recipes, error)
	Print(recipes recipes.Recipes)
}

// DefineFile determines the file type and returns the appropriate DBReader implementation.
func DefineFile(reader *DBReader, filename string) ([]byte, error) {
	file, err := pathRead(filename)
	if err != nil {
		return nil, err
	}

	ext := fileExtension(filename)
	if reader != nil {
		if ext == "xml" {
			*reader = &encode.RecipesXML{}
		} else if ext == "json" {
			*reader = &encode.RecipesJSON{}
		} else {
			return nil, fmt.Errorf("'%s' is wrong file extension (required 'json' or 'xml')", ext)
		}
	}

	return file, nil
}

// pathRead reads the file from the given path.
func pathRead(path string) ([]byte, error) {
	if path == "" {
		return nil, fmt.Errorf("please provide a filename using the flag (look at flag list with -h)")
	}

	file, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	return file, nil
}

// fileExtension returns the file extension of the given filename.
func fileExtension(path string) string {
	for i := len(path) - 1; i >= 0; i-- {
		if path[i] == '.' {
			return path[i+1:]
		}
	}

	return ""
}
