package readdb

import (
	"flag"
	"fmt"
	"main/encode"
	"os"
)

type DBReader interface {
	Read(file []byte) (encode.Recipes, error)
	Print(recipes encode.Recipes)
}

func DefineFile(reader *DBReader) (file []byte, err error) {
	filename := flag.String("f", "", "Filename to read (xml or json)")
	flag.Parse()

	file, ext, err := pathRead(filename)
	if err == nil {
		if ext == "xml" {
			*reader = &encode.RecipesXML{}
		} else {
			*reader = &encode.RecipesJSON{}
		}
	}

	return file, err
}

func pathRead(path *string) (file []byte, ext string, err error) {
	if *path == "" {
		err = fmt.Errorf("please provide a filename using the -f flag")
		return
	}

	ext = fileExtension(*path)
	if ext != "json" && ext != "xml" {
		err = fmt.Errorf("file extension must be 'json' or 'xml'")
		return
	}

	file, err = os.ReadFile(*path)

	return file, ext, err
}

func fileExtension(path string) string {
	for i := len(path) - 1; i >= 0; i-- {
		if path[i] == '.' {
			return path[i+1:]
		}
	}

	return ""
}
