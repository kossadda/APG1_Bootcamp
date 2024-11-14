package readdb

import (
	"fmt"
	"os"

	"github.com/kossadda/APG1_Bootcamp/Go_Day01/src/encode"
	"github.com/kossadda/APG1_Bootcamp/Go_Day01/src/recipes"
)

type DBReader interface {
	Read(file []byte) (recipes.Recipes, error)
	Print(recipes recipes.Recipes)
}

func DefineFile(reader *DBReader, filename string) ([]byte, error) {
	file, err := pathRead(filename)

	ext := fileExtension(filename)
	if err == nil {
		if reader != nil {
			if ext == "xml" {
				*reader = &encode.RecipesXML{}
			} else if ext == "json" {
				*reader = &encode.RecipesJSON{}
			} else {
				err = fmt.Errorf("'%s' is wrong file extension (required 'json' or 'xml')", ext)
			}
		} else {
			if ext != "txt" {
				err = fmt.Errorf("'%s' is wrong file extension (required 'txt')", ext)
			}
		}
	}

	return file, err
}

func pathRead(path string) (file []byte, err error) {
	if path == "" {
		err = fmt.Errorf("please provide a filename using the flag (look at flag list with -h)")
		return
	}

	file, err = os.ReadFile(path)

	return file, err
}

func fileExtension(path string) string {
	for i := len(path) - 1; i >= 0; i-- {
		if path[i] == '.' {
			return path[i+1:]
		}
	}

	return ""
}
