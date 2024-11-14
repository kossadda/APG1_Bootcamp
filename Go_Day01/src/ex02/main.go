package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/kossadda/APG1_Bootcamp/Go_Day01/src/comparefs"
	"github.com/kossadda/APG1_Bootcamp/Go_Day01/src/readdb"
)

func main() {
	snapshot1 := flag.String("old", "", "Old filesystem (txt)")
	snapshot2 := flag.String("new", "", "New filesystem (txt)")
	flag.Parse()

	file1 := Snapshot(*snapshot1)
	file2 := Snapshot(*snapshot2)

	comparefs.Compare(string(file1), string(file2))
}

func Snapshot(path string) []byte {
	file, err := readdb.DefineFile(nil, path)

	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}

	return file
}
