package pathReader

import (
	"fmt"
	"os"
)

func PathRead(path *string) (file []byte, ext string, err error) {
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
